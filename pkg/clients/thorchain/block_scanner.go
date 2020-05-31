package thorchain

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
)

// BlockScanner is a kind of scanner that will fetch events through scanning blocks.
// with websocket or directly by requesting http endpoint.
type BlockScanner struct {
	client   Tendermint
	callback Callback
	interval time.Duration
	stopChan chan struct{}
	wg       sync.WaitGroup
	running  bool
	height   int64
	logger   zerolog.Logger
	pageSize int64
}

// NewBlockScanner will create a new instance of BlockScanner.
func NewBlockScanner(client Tendermint, callback Callback, interval time.Duration) *BlockScanner {
	sc := &BlockScanner{
		client:   client,
		callback: callback,
		interval: interval,
		stopChan: make(chan struct{}),
		logger:   log.With().Str("module", "block_scanner").Logger(),
		pageSize: 20,
	}
	return sc
}

// SetHeight sets the height that scanner will start scanning from.
func (sc *BlockScanner) SetHeight(height int64) error {
	if sc.running {
		return errors.New("scanner is running")
	}

	sc.height = height
	return nil
}

// GetHeight returns the latest processed block height.
func (sc *BlockScanner) GetHeight() int64 {
	return atomic.LoadInt64(&sc.height)
}

// Start will start the scanner.
func (sc *BlockScanner) Start() error {
	if sc.running {
		return errors.New("scanner is already running")
	}

	sc.running = true
	go sc.scan()
	return nil
}

func (sc *BlockScanner) scan() {
	sc.wg.Add(1)
	defer sc.wg.Done()

	for {
		select {
		case <-sc.stopChan:
			return
		default:
			synced, err := sc.processNextBlock()
			if err != nil {
				sc.logger.Error().Int64("height", sc.GetHeight()).Err(err).Msg("failed to process the next block")
			} else {
				if !synced {
					continue
				}
			}

			select {
			case <-time.After(sc.interval):
			case <-sc.stopChan:
				return
			}
		}
	}
}

func (sc *BlockScanner) processNextBlock() (bool, error) {
	height := sc.GetHeight()
	info, err := sc.client.BlockchainInfo(height, height+sc.pageSize)
	if err != nil {
		return false, errors.Wrap(err, "could not get blockchain info")
	}
	for i, j := 0, len(info.BlockMetas)-1; i < j; i, j = i+1, j-1 {
		info.BlockMetas[i], info.BlockMetas[j] = info.BlockMetas[j], info.BlockMetas[i]
	}
	if int64(len(info.BlockMetas)) != sc.pageSize {
		fmt.Println("Oops")
	}
	for _, blockInfo := range info.BlockMetas {
		if blockInfo.NumTxs > 0 {
			block, err := sc.client.BlockResults(&blockInfo.Header.Height)
			if err != nil {
				return false, errors.Wrap(err, "could not get results of block")
			}

			for _, tx := range block.TxsResults {
				events := convertEvents(tx.Events)
				sc.callback.NewTx(blockInfo.Header.Height, events)
			}

			blockTime := info.BlockMetas[0].Header.Time
			beginEvents := convertEvents(block.BeginBlockEvents)
			endEvents := convertEvents(block.EndBlockEvents)
			sc.callback.NewBlock(blockInfo.Header.Height, blockTime, beginEvents, endEvents)
		}
		sc.incrementHeight()
	}
	synced := len(info.BlockMetas) == 0 || info.BlockMetas[len(info.BlockMetas)-1].Header.Height == info.LastHeight
	if synced {
		fmt.Println("Synced!!!")
	}
	return synced, nil
}

func (sc *BlockScanner) incrementHeight() {
	newHeight := atomic.AddInt64(&sc.height, 1)
	sc.logger.Info().Int64("height", newHeight).Msg("new block scanned")
}

// Stop will attempt to stop the scanner (blocking until the scanner stops completely).
func (sc *BlockScanner) Stop() error {
	if !sc.running {
		return errors.New("scanner isn't running")
	}

	sc.stopChan <- struct{}{}
	sc.wg.Wait()

	sc.running = false
	return nil
}

func convertEvents(tes []abcitypes.Event) []Event {
	events := make([]Event, len(tes))
	for i, te := range tes {
		events[i].FromTendermintEvent(te)
	}

	return events
}

// Tendermint represents every method BlockScanner needs to scan blocks.
type Tendermint interface {
	BlockchainInfo(minHeight, maxHeight int64) (*coretypes.ResultBlockchainInfo, error)
	BlockResults(height *int64) (*coretypes.ResultBlockResults, error)
}

// Callback represents methods required by Scanner to notify events.
type Callback interface {
	NewBlock(height int64, blockTime time.Time, begin, end []Event)
	NewTx(height int64, events []Event)
}
