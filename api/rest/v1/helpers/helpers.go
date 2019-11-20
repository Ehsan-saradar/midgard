package helpers

import (
	"github.com/openlyinc/pointy"
	"gitlab.com/thorchain/bepswap/chain-service/internal/models"

	api "gitlab.com/thorchain/bepswap/chain-service/api/rest/v1/codegen"
	"gitlab.com/thorchain/bepswap/chain-service/internal/common"
)

func ConvertAssetForAPI(asset common.Asset) *api.Asset {
	return &api.Asset{
		Chain:  pointy.String(asset.Chain.String()),
		Symbol: pointy.String(asset.Symbol.String()),
		Ticker: pointy.String(asset.Ticker.String()),
	}
}

func ConvertEventDataForAPI(events models.Events) *api.Event {
	return &api.Event{
		Fee:        pointy.Int64(int64(events.Fee)),
		Slip:       pointy.Float64(events.Slip),
		StakeUnits: pointy.Int64(int64(events.StakeUnits)),
	}
}

func ConvertGasForAPI(gas models.TxGas) *api.Gas {
	if gas.Amount == 0 {
		return nil
	}

	a, _ := common.NewAsset(gas.Asset.Symbol.String())
	asset := ConvertAssetForAPI(a)

	return &api.Gas{
		Amount: pointy.Int64(int64(gas.Amount)),
		Asset:  asset,
	}
}

func ConvertCoinForAPI(coin common.Coin) *api.Coin {
	a, _ := common.NewAsset(coin.Asset.Symbol.String())
	asset := ConvertAssetForAPI(a)

	return &api.Coin{
		Amount: pointy.Int64(coin.Amount),
		Asset:  asset,
	}
}

func ConvertCoinsForAPI(coins common.Coins) api.Coins {
	var c []api.Coin
	for _, coin := range coins {
		converted := ConvertCoinForAPI(coin)
		c = append(c, *converted)
	}

	return c
}

func ConvertTxForAPI(tx models.TxData) *api.Tx {
	if tx.Address == "" {
		return nil
	}

	coins := ConvertCoinsForAPI(tx.Coin)
	return &api.Tx{
		Address: pointy.String(tx.Address),
		Coins:   &coins,
		Memo:    pointy.String(tx.Memo),
		TxID:    pointy.String(tx.TxID),
	}
}

func ConvertOptionsForAPI(options models.Options) *api.Option {
	return &api.Option{
		Asymmetry:           pointy.Float64(options.Asymmetry),
		PriceTarget:         pointy.Int64(int64(options.PriceTarget)),
		WithdrawBasisPoints: pointy.Int64(int64(options.WithdrawBasisPoints)),
	}
}
