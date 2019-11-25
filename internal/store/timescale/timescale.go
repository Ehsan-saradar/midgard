package timescale

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	migrate "github.com/rubenv/sql-migrate"

	"gitlab.com/thorchain/midgard/internal/config"
)

type Client struct {
	logger zerolog.Logger
	cfg    config.TimeScaleConfiguration
	db     *sqlx.DB
}


func NewClient(cfg config.TimeScaleConfiguration) *Client {
	time.Sleep(3* time.Second)
	logger := log.With().Str("module", "timescale").Logger()
	connStr := fmt.Sprintf("user=%s sslmode=%v password=%v host=%v port=%v", cfg.UserName, cfg.Sslmode, cfg.Password, cfg.Host, cfg.Port)
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		logger.Err(err).Msg("Open")
		return &Client{}
	}

	if err := CreateDatabase(db, cfg); err != nil {
		logger.Err(err).Msg("CreateDatabase")
	}

	db, err = Open(cfg)
	if err != nil {
		logger.Err(err).Msg("Open")
	}
	if err := MigrationsUp(db, logger, cfg); err != nil {
		logger.Err(err).Msg("MigrationsUp")
	}

	return &Client{
		cfg: cfg,
		db:  db,
		logger: logger,
	}
}

func (s *Client) Ping() error {
	return s.db.Ping()
}

func Open(cfg config.TimeScaleConfiguration) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=%v password=%v host=%v port=%v", cfg.UserName, cfg.Database, cfg.Sslmode, cfg.Password, cfg.Host, cfg.Port)
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return &sqlx.DB{}, err
	}
	return db, nil
}

func (s *Client) Open() (*sqlx.DB, error) {
	return Open(s.cfg)
}

func CreateDatabase(db *sqlx.DB, cfg config.TimeScaleConfiguration) error{
	query := fmt.Sprintf(`CREATE DATABASE %v;`, cfg.Database)
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (s *Client) CreateDatabase() error {
	return CreateDatabase(s.db, s.cfg)
}

func MigrationsUp(db *sqlx.DB, logger zerolog.Logger, cfg config.TimeScaleConfiguration) error {
	n, err := migrate.Exec(db.DB, "postgres", &migrate.FileMigrationSource{Dir:cfg.MigrationsDir}, migrate.Up)

	if err != nil {
		return err
	}
	logger.Debug().Int("Applied migrations", n)
	return nil
}

func (s *Client) MigrationsUp() error {
	return MigrationsUp(s.db, s.logger, s.cfg)
}

func MigrationsDown(db *sqlx.DB, logger zerolog.Logger, cfg config.TimeScaleConfiguration) error {
	n, err := migrate.Exec(db.DB, "postgres", &migrate.FileMigrationSource{Dir:cfg.MigrationsDir}, migrate.Down)
	if err != nil {
		return err
	}
	logger.Debug().Int("Applied migrations", n)
	return nil
}

func (s *Client) MigrationsDown() error {
	return MigrationsDown(s.db, s.logger, s.cfg)
}

