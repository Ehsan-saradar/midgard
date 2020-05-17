package models

import (
	"gitlab.com/thorchain/midgard/internal/clients/thorchain/types"
	"gitlab.com/thorchain/midgard/internal/common"
)

type GasPool struct {
	Asset    common.Asset `json:"asset"`
	AssetAmt uint64       `json:"asset_amt,string" mapstructure:"asset_amt"`
	RuneAmt  uint64       `json:"rune_amt,string"  mapstructure:"rune_amt"`
}

type EventGas struct {
	Event
	Pools []GasPool
}

func NewGasEvent(gas types.EventGas, event types.Event) EventGas {
	gasEvents := EventGas{
		Event: newEvent(event),
		Pools: []GasPool{},
	}
	for _, pool := range gas.Pools {
		gasEvents.Pools = append(gasEvents.Pools, GasPool{
			Asset:    pool.Asset,
			AssetAmt: pool.AssetAmt,
			RuneAmt:  pool.RuneAmt,
		})
	}

	return gasEvents
}
