package timescale

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"gitlab.com/thorchain/midgard/internal/common"
	"gitlab.com/thorchain/midgard/internal/models"
)

const blockSpeed = 3

// timeOfBlock = ((currentTime - genesisTime) / (currentBlockheight))*blockHeight + genesisTime (edited)
func (s *Client) GetDateCreated(asset common.Asset) (uint64, error) {
	assetBlockHeight, err := s.getBlockHeight(asset)
	if err != nil {
		return 0, err
	}
	dateCreated, err := s.getTimeOfBlock(assetBlockHeight)
	if err != nil {
		return 0, err
	}

	return dateCreated, nil
}

func (s *Client) getTimeOfBlock(assetBlockHeight uint64) (uint64, error) {
	currentTime := uint64(time.Now().Unix())
	getGenesis, err := s.getGenesis()
	if err != nil {
		return 0, errors.Wrap(err, "failed to getGenesis")
	}
	genesisTime := uint64(getGenesis.Unix())
	currentBlockHeight := (currentTime - genesisTime) / blockSpeed

	timeOfBlock := (((currentTime - genesisTime) / currentBlockHeight) * assetBlockHeight) + genesisTime

	return timeOfBlock, nil
}

func (s *Client) getGenesis() (time.Time, error) {
	stmnt := `SELECT genesis_time FROM genesis`

	var genesisTime time.Time
	row := s.db.QueryRow(stmnt)

	if err := row.Scan(&genesisTime); err != nil {
		return time.Time{}, err
	}

	return genesisTime, nil
}

func (s *Client) getBlockHeight(pool common.Asset) (uint64, error) {
	stmnt := fmt.Sprintf(`
    SELECT MAX(height)
    FROM %v
    WHERE pool = $1
  `, models.ModelEventsTable)

	var blockHeight sql.NullInt64
	if err := s.db.Get(&blockHeight, stmnt, pool.String()); err != nil {
		return 0, err
	}

	return uint64(blockHeight.Int64), nil
}

func (s *Client) CreateGenesis(genesis models.Genesis) (int64, error) {
	query := fmt.Sprintf(`
		INSERT INTO %v (
			genesis_time
		)  VALUES ( $1 )
		ON CONFLICT (genesis_time) DO NOTHING;`, models.ModelGenesisTable)

	results, err := s.db.Exec(query, genesis.GenesisTime)

	if err != nil {
		return 0, errors.Wrap(err, "Failed to prepareNamed query for GenesisRecord")
	}

	return results.RowsAffected()
}
