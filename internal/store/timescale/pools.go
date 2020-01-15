package timescale

import (
	"errors"

	"gitlab.com/thorchain/midgard/internal/common"
)

type PoolData struct {
	Status           string // TODO build
	Asset            common.Asset
	AssetDepth       uint64
	AssetROI         float64
	AssetStakedTotal uint64
	BuyAssetCount    uint64
	BuyFeeAverage    uint64
	BuyFeesTotal     uint64
	BuySlipAverage   float64
	BuyTxAverage     uint64
	BuyVolume        uint64
	PoolDepth        uint64
	PoolFeeAverage   uint64
	PoolFeesTotal    uint64
	PoolROI          float64
	PoolROI12        float64
	PoolSlipAverage  float64
	PoolStakedTotal  uint64
	PoolTxAverage    uint64
	PoolUnits        uint64
	PoolVolume       uint64
	PoolVolume24hr   uint64
	Price            float64
	RuneDepth        uint64
	RuneROI          float64
	RuneStakedTotal  uint64
	SellAssetCount   uint64
	SellFeeAverage   uint64
	SellFeesTotal    uint64
	SellSlipAverage  float64
	SellTxAverage    uint64
	SellVolume       uint64
	StakeTxCount     uint64
	StakersCount     uint64
	StakingTxCount   uint64
	SwappersCount    uint64
	SwappingTxCount  uint64
	WithdrawTxCount  uint64
}

var USDPools = []string{
	"BNB.TUSD-000",
	"BNB.BUSD-BD1",
}

func (s *Client) GetPool(asset common.Asset) (common.Asset, error) {
	query := `
		SELECT sub.pool
		FROM (
			SELECT pool, SUM(units) AS total_units
			FROM stakes
			WHERE pool = $1
			GROUP BY pool
		) as sub
		WHERE sub.total_units > 0
	`

	row := s.db.QueryRowx(query, asset.String())

	var a string

	if err := row.Scan(&a); err != nil {
		return common.Asset{}, err
	}

	return common.NewAsset(a)
}

func (s *Client) GetPools() ([]common.Asset, error) {
	var pools []common.Asset

	query := `
		SELECT sub.pool
		From (
			SELECT pool, SUM(units) AS total_units
			FROM stakes
			GROUP BY pool
		) AS sub
		WHERE sub.total_units > 0
	`

	rows, err := s.db.Queryx(query)
	if err != nil {
		return nil, err
	}

	type results struct {
		Pool string
	}

	for rows.Next() {
		var result results
		if err := rows.StructScan(&result); err != nil {
			return nil, err
		}
		pool, err := common.NewAsset(result.Pool)
		if err != nil {
			return nil, err
		}
		pools = append(pools, pool)
	}

	return pools, nil
}

func (s *Client) GetPoolData(asset common.Asset) (PoolData, error) {
	exists, err := s.exists(asset)
	if err != nil {
		return PoolData{}, err
	}
	if !exists {
		return PoolData{}, errors.New("pool does not exist")
	}

	assetDepth, err := s.assetDepth(asset)
	if err != nil {
		return PoolData{}, err
	}

	assetROI, err := s.assetROI(asset)
	if err != nil {
		return PoolData{}, err
	}

	assetStakedTotal, err := s.assetStakedTotal(asset)
	if err != nil {
		return PoolData{}, err
	}

	buyAssetCount, err := s.buyAssetCount(asset)
	if err != nil {
		return PoolData{}, err
	}

	buyFeeAverage, err := s.buyFeeAverage(asset)
	if err != nil {
		return PoolData{}, err
	}

	buyFeesTotal, err := s.buyFeesTotal(asset)
	if err != nil {
		return PoolData{}, err
	}

	buySlipAverage, err := s.buySlipAverage(asset)
	if err != nil {
		return PoolData{}, err
	}

	buyTxAverage, err := s.buyTxAverage(asset)
	if err != nil {
		return PoolData{}, err
	}

	buyVolume, err := s.buyVolume(asset)
	if err != nil {
		return PoolData{}, err
	}

	poolDepth, err := s.poolDepth(asset)
	if err != nil {
		return PoolData{}, err
	}

	poolFeeAverage, err := s.poolFeeAverage(asset)
	if err != nil {
		return PoolData{}, err
	}

	poolFeesTotal, err := s.poolFeesTotal(asset)
	if err != nil {
		return PoolData{}, err
	}

	poolSlipAverage, err := s.poolSlipAverage(asset)
	if err != nil {
		return PoolData{}, err
	}

	poolStakedTotal, err := s.poolStakedTotal(asset)
	if err != nil {
		return PoolData{}, err
	}

	poolTxAverage, err := s.poolTxAverage(asset)
	if err != nil {
		return PoolData{}, err
	}

	poolUnits, err := s.poolUnits(asset)
	if err != nil {
		return PoolData{}, err
	}

	poolVolume, err := s.poolVolume(asset)
	if err != nil {
		return PoolData{}, err
	}

	poolVolume24hr, err := s.poolVolume24hr(asset)
	if err != nil {
		return PoolData{}, err
	}

	GetPriceInRune, err := s.GetPriceInRune(asset)
	if err != nil {
		return PoolData{}, err
	}

	runeDepth, err := s.runeDepth(asset)
	if err != nil {
		return PoolData{}, err
	}

	runeROI, err := s.runeROI(asset)
	if err != nil {
		return PoolData{}, err
	}

	runeStakedTotal, err := s.runeStakedTotal(asset)
	if err != nil {
		return PoolData{}, err
	}

	sellAssetCount, err := s.sellAssetCount(asset)
	if err != nil {
		return PoolData{}, err
	}

	sellFeeAverage, err := s.sellFeeAverage(asset)
	if err != nil {
		return PoolData{}, err
	}

	sellFeesTotal, err := s.sellFeesTotal(asset)
	if err != nil {
		return PoolData{}, err
	}

	sellSlipAverage, err := s.sellSlipAverage(asset)
	if err != nil {
		return PoolData{}, err
	}

	sellTxAverage, err := s.sellTxAverage(asset)
	if err != nil {
		return PoolData{}, err
	}

	sellVolume, err := s.sellVolume(asset)
	if err != nil {
		return PoolData{}, err
	}

	stakeTxCount, err := s.stakeTxCount(asset)
	if err != nil {
		return PoolData{}, err
	}

	stakersCount, err := s.stakersCount(asset)
	if err != nil {
		return PoolData{}, err
	}

	stakingTxCount, err := s.stakingTxCount(asset)
	if err != nil {
		return PoolData{}, err
	}

	swappersCount, err := s.swappersCount(asset)
	if err != nil {
		return PoolData{}, err
	}

	swappingTxCount, err := s.swappingTxCount(asset)
	if err != nil {
		return PoolData{}, err
	}

	withdrawTxCount, err := s.withdrawTxCount(asset)
	if err != nil {
		return PoolData{}, err
	}

	poolROI, err := s.poolROI(asset)
	if err != nil {
		return PoolData{}, err
	}

	poolROI12, err := s.poolROI12(asset)
	if err != nil {
		return PoolData{}, err
	}

	return PoolData{
		Asset:            asset,
		AssetDepth:       assetDepth,
		AssetROI:         assetROI,
		AssetStakedTotal: assetStakedTotal,
		BuyAssetCount:    buyAssetCount,
		BuyFeeAverage:    buyFeeAverage,
		BuyFeesTotal:     buyFeesTotal,
		BuySlipAverage:   buySlipAverage,
		BuyTxAverage:     buyTxAverage,
		BuyVolume:        buyVolume,
		PoolDepth:        poolDepth,
		PoolFeeAverage:   poolFeeAverage,
		PoolFeesTotal:    poolFeesTotal,
		PoolROI:          poolROI,
		PoolROI12:        poolROI12,
		PoolSlipAverage:  poolSlipAverage,
		PoolStakedTotal:  poolStakedTotal,
		PoolTxAverage:    poolTxAverage,
		PoolUnits:        poolUnits,
		PoolVolume:       poolVolume,
		PoolVolume24hr:   poolVolume24hr,
		Price:            GetPriceInRune,
		RuneDepth:        runeDepth,
		RuneROI:          runeROI,
		RuneStakedTotal:  runeStakedTotal,
		SellAssetCount:   sellAssetCount,
		SellFeeAverage:   sellFeeAverage,
		SellFeesTotal:    sellFeesTotal,
		SellSlipAverage:  sellSlipAverage,
		SellTxAverage:    sellTxAverage,
		SellVolume:       sellVolume,
		StakeTxCount:     stakeTxCount,
		StakersCount:     stakersCount,
		StakingTxCount:   stakingTxCount,
		SwappersCount:    swappersCount,
		SwappingTxCount:  swappingTxCount,
		WithdrawTxCount:  withdrawTxCount,
	}, nil
}

func (s *Client) GetPriceInRune(asset common.Asset) (float64, error) {
	assetDepth, err := s.assetDepth(asset)
	if err != nil {
		return 0, err
	}
	if assetDepth > 0 {
		runeDepth, err := s.runeDepth(asset)
		if err != nil {
			return 0, err
		}
		return float64(runeDepth) / float64(assetDepth), nil
	}

	return 0, nil
}

func (s *Client) exists(asset common.Asset) (bool, error) {
	staked, err := s.stakeTxCount(asset)
	if err != nil {
		return false, err
	}
	if staked > 0 {
		return true, nil
	}

	return false, nil
}

// assetStakedTotal - total amount of asset staked in given pool
func (s *Client) assetStakedTotal(asset common.Asset) (uint64, error) {
	stmnt := `
		SELECT SUM(assetAmt)
		FROM stakes
		WHERE pool = $1
		`

	var assetStakedTotal uint64
	row := s.db.QueryRow(stmnt, asset.String())

	if err := row.Scan(&assetStakedTotal); err != nil {
		return 0, err
	}

	return assetStakedTotal, nil
}

// assetStakedTotal12 - total amount of asset staked in given pool in the last
// 12 months
func (s *Client) assetStakedTotal12m(asset common.Asset) (uint64, error) {
	stmnt := `
		SELECT SUM(assetAmt)
		FROM stakes
		WHERE pool = $1
		AND time BETWEEN NOW() - INTERVAL '12 MONTHS' AND NOW()
	`

	var assetStakedTotal uint64
	row := s.db.QueryRow(stmnt, asset.String())

	if err := row.Scan(&assetStakedTotal); err != nil {
		return 0, err
	}

	return assetStakedTotal, nil
}

// assetWithdrawnTotal - total amount of asset withdrawn
func (s *Client) assetWithdrawnTotal(asset common.Asset) (int64, error) {
	stmnt := `
		SELECT SUM(assetAmt)
		FROM stakes
		WHERE pool = $1
		AND units < 0
		`

	var assetWithdrawnTotal int64
	row := s.db.QueryRow(stmnt, asset.String())

	if err := row.Scan(&assetWithdrawnTotal); err != nil {
		return 0, err
	}

	return -assetWithdrawnTotal, nil
}

// runeStakedTotal - total amount of rune staked on the network for given pool.
func (s *Client) runeStakedTotal(asset common.Asset) (uint64, error) {
	stmnt := `
		SELECT SUM(runeAmt)
		FROM stakes
		WHERE pool = $1
	`

	var runeStakedTotal uint64
	row := s.db.QueryRow(stmnt, asset.String())

	if err := row.Scan(&runeStakedTotal); err != nil {
		return 0, err
	}

	return runeStakedTotal, nil
}

// runeStakedTotal12m - total amount of rune staked on the network for given
// pool in the last 12 months.
func (s *Client) runeStakedTotal12m(asset common.Asset) (uint64, error) {
	stmnt := `
		SELECT SUM(runeAmt)
		FROM stakes
		WHERE pool = $1
		AND time BETWEEN NOW() - INTERVAL '12 MONTHS' AND NOW()
		`

	var runeStakedTotal uint64
	row := s.db.QueryRow(stmnt, asset.String())

	if err := row.Scan(&runeStakedTotal); err != nil {
		return 0, err
	}

	return runeStakedTotal, nil
}

func (s *Client) poolStakedTotal(asset common.Asset) (uint64, error) {
	assetTotal, err := s.assetStakedTotal(asset)
	if err != nil {
		return 0, err
	}
	runeTotal, err := s.runeStakedTotal(asset)
	if err != nil {
		return 0, nil
	}
	priceInRune, err := s.GetPriceInRune(asset)
	if err != nil {
		return 0, err
	}

	stakedPrice := float64(assetTotal) * priceInRune
	stakedTotal := runeTotal + (uint64(stakedPrice))

	return stakedTotal, nil
}

// +stakes
// +incomingSwapAsset
// -outgoingSwapAsset
// -withdraws
func (s *Client) assetDepth(asset common.Asset) (uint64, error) {
	stakes, err := s.assetStakedTotal(asset)
	if err != nil {
		return 0, err
	}
	swaps, err := s.assetSwapTotal(asset)
	if err != nil {
		return 0, nil
	}

	depth := int64(stakes) + swaps
	return uint64(depth), nil
}

func (s *Client) assetDepth12m(asset common.Asset) (uint64, error) {
	stakes, err := s.assetStakedTotal12m(asset)
	if err != nil {
		return 0, err
	}
	swaps, err := s.assetSwapTotal12m(asset)
	if err != nil {
		return 0, err
	}

	depth := int64(stakes) + swaps
	return uint64(depth), nil
}

func (s *Client) runeDepth(asset common.Asset) (uint64, error) {
	stakes, err := s.runeStakedTotal(asset)
	if err != nil {
		return 0, err
	}
	swaps, err := s.runeSwapTotal(asset)
	if err != nil {
		return 0, err
	}

	depth := int64(stakes) + swaps
	return uint64(depth), nil
}

func (s *Client) runeDepth12m(asset common.Asset) (uint64, error) {
	stakes, err := s.runeStakedTotal12m(asset)
	if err != nil {
		return 0, err
	}
	swaps, err := s.runeSwapTotal12m(asset)
	if err != nil {
		return 0, err
	}
	depth := int64(stakes) + swaps
	return uint64(depth), nil
}

// runeSwapTotal - total amount rune swapped through the pool
func (s *Client) runeSwapTotal(asset common.Asset) (int64, error) {
	stmnt := `
		SELECT SUM(runeAmt)
		FROM swaps
		WHERE pool = $1
	`

	var total int64
	row := s.db.QueryRow(stmnt, asset.String())

	if err := row.Scan(&total); err != nil {
		return 0, err
	}

	return total, nil
}

// runeSwapTotal12m - total amount rune swapped through the pool in the last 12
// months
func (s *Client) runeSwapTotal12m(asset common.Asset) (int64, error) {
	stmnt := `
		SELECT SUM(runeAmt)
		FROM swaps
		WHERE pool = $1
		AND time BETWEEN NOW() - INTERVAL '12 MONTHS' AND NOW()
	`

	var total int64
	row := s.db.QueryRow(stmnt, asset.String())

	if err := row.Scan(&total); err != nil {
		return 0, err
	}

	return total, nil
}

func (s *Client) assetSwapTotal(asset common.Asset) (int64, error) {
	stmnt := `
		SELECT SUM(assetAmt)
		FROM swaps
		WHERE pool = $1
	`

	var total int64
	row := s.db.QueryRow(stmnt, asset.String())

	if err := row.Scan(&total); err != nil {
		return 0, err
	}

	return total, nil
}

func (s *Client) assetSwapTotal12m(asset common.Asset) (int64, error) {
	stmnt := `
		SELECT SUM(runeAmt)
		FROM swaps
		WHERE pool = $1
		AND time BETWEEN NOW() - INTERVAL '12 MONTHS' AND NOW()
	`

	var total int64
	row := s.db.QueryRow(stmnt, asset.String())

	if err := row.Scan(&total); err != nil {
		return 0, err
	}

	return total, nil
}

func (s *Client) poolDepth(asset common.Asset) (uint64, error) {
	runeDepth, err := s.runeDepth(asset)
	if err != nil {
		return 0, err
	}
	return 2 * runeDepth, nil
}

func (s *Client) poolUnits(asset common.Asset) (uint64, error) {
	stmnt := `
		SELECT SUM(units)
		FROM stakes
		WHERE pool = $1
	`

	var units uint64
	row := s.db.QueryRow(stmnt, asset.String())

	if err := row.Scan(&units); err != nil {
		return 0, err
	}

	return units, nil
}

func (s *Client) sellVolume(asset common.Asset) (uint64, error) {
	stmnt := `
		SELECT SUM(assetAmt)
		FROM swaps
		WHERE pool = $1
		AND assetAmt > 0
	`

	var sellVolume uint64
	row := s.db.QueryRow(stmnt, asset.String())

	if err := row.Scan(&sellVolume); err != nil {
		return 0, err
	}

	priceInRune, err := s.GetPriceInRune(asset)
	if err != nil {
		return 0, err
	}
	return uint64(float64(sellVolume) * priceInRune), nil
}

func (s *Client) sellVolume24hr(asset common.Asset) (uint64, error) {
	stmnt := `
		SELECT SUM(assetAmt)
		FROM swaps
		WHERE pool = $1
		AND assetAmt > 0
		AND time BETWEEN NOW() - INTERVAL '24 HOURS' AND NOW()
	`

	var sellVolume uint64
	row := s.db.QueryRow(stmnt, asset.String())

	if err := row.Scan(&sellVolume); err != nil {
		return 0, err
	}

	priceInRune, err := s.GetPriceInRune(asset)
	if err != nil {
		return 0, err
	}
	return uint64(float64(sellVolume) * priceInRune), nil
}

func (s *Client) buyVolume(asset common.Asset) (uint64, error) {
	stmnt := `
		SELECT SUM(runeAmt)
		FROM swaps
		WHERE pool = $1
		AND runeAmt > 0
	`

	var buyVolume uint64
	row := s.db.QueryRow(stmnt, asset.String())

	if err := row.Scan(&buyVolume); err != nil {
		return 0, err
	}

	return buyVolume, nil
}

func (s *Client) buyVolume24hr(asset common.Asset) (uint64, error) {
	stmnt := `
		SELECT SUM(runeAmt)
		FROM swaps
		WHERE pool = $1
		AND runeAmt > 0
		AND time BETWEEN NOW() - INTERVAL '24 HOURS' AND NOW()
	`

	var buyVolume uint64
	row := s.db.QueryRow(stmnt, asset.String())

	if err := row.Scan(&buyVolume); err != nil {
		return 0, err
	}

	return buyVolume, nil
}

func (s *Client) poolVolume(asset common.Asset) (uint64, error) {
	sellVolume, err := s.sellVolume(asset)
	if err != nil {
		return 0, err
	}

	buyVolume, err := s.buyVolume(asset)
	if err != nil {
		return 0, err
	}

	return buyVolume + sellVolume, nil
}

func (s *Client) poolVolume24hr(asset common.Asset) (uint64, error) {
	buyVolume24r, err := s.buyVolume24hr(asset)
	if err != nil {
		return 0, err
	}

	sellVolume24r, err := s.sellVolume24hr(asset)
	if err != nil {
		return 0, err
	}
	return buyVolume24r + sellVolume24r, nil
}

func (s *Client) sellTxAverage(asset common.Asset) (uint64, error) {
	stmnt := `
		SELECT AVG(assetAmt)
		FROM swaps
		WHERE pool = $1
		AND assetAmt > 0
	`

	var avg float64
	row := s.db.QueryRow(stmnt, asset.String())

	if err := row.Scan(&avg); err != nil {
		return 0, err
	}

	priceInRune, err := s.GetPriceInRune(asset)
	if err != nil {
		return 0, err
	}
	return uint64(avg * priceInRune), nil
}

func (s *Client) buyTxAverage(asset common.Asset) (uint64, error) {
	stmnt := `
		SELECT AVG(runeAmt)
		FROM swaps
		WHERE pool = $1
		AND runeAmt > 0
	`

	var avg uint64
	row := s.db.QueryRow(stmnt, asset.String())

	if err := row.Scan(&avg); err != nil {
		return 0, err
	}

	return avg, nil
}

func (s *Client) poolTxAverage(asset common.Asset) (uint64, error) {
	buyTxAverage, err := s.buyTxAverage(asset)
	if err != nil {
		return 0, err
	}

	sellTxAverage, err := s.sellTxAverage(asset)
	if err != nil {
		return 0, err
	}
	return (buyTxAverage + sellTxAverage) / 2, nil
}

func (s *Client) sellSlipAverage(asset common.Asset) (float64, error) {
	stmnt := `
		SELECT AVG(trade_slip)
		FROM swaps
		WHERE pool = $1
		AND assetAmt > 0
	`

	var sellSlipAverage float64
	row := s.db.QueryRow(stmnt, asset.String())

	if err := row.Scan(&sellSlipAverage); err != nil {
		return 0, err
	}

	return sellSlipAverage, nil
}

func (s *Client) buySlipAverage(asset common.Asset) (float64, error) {
	stmnt := `
		SELECT AVG(trade_slip)
		FROM swaps
		WHERE pool = $1
		AND runeAmt > 0
	`

	var buySlipAverage float64
	row := s.db.QueryRow(stmnt, asset.String())

	if err := row.Scan(&buySlipAverage); err != nil {
		return 0, err
	}

	return buySlipAverage, nil
}

func (s *Client) poolSlipAverage(asset common.Asset) (float64, error) {
	buySlipAverage, err := s.buySlipAverage(asset)
	if err != nil {
		return 0, err
	}

	sellSlipAverage, err := s.sellSlipAverage(asset)
	if err != nil {
		return 0, err
	}
	return (buySlipAverage + sellSlipAverage) / 2, nil
}

func (s *Client) sellFeeAverage(asset common.Asset) (uint64, error) {
	stmnt := `
		SELECT AVG(liquidity_fee)
		FROM swaps
		WHERE pool = $1
		AND assetAmt > 0
	`

	var sellFeeAverage uint64
	row := s.db.QueryRow(stmnt, asset.String())

	if err := row.Scan(&sellFeeAverage); err != nil {
		return 0, err
	}

	priceInRune, err := s.GetPriceInRune(asset)
	if err != nil {
		return 0, err
	}
	return uint64(float64(sellFeeAverage) * priceInRune), nil
}

func (s *Client) buyFeeAverage(asset common.Asset) (uint64, error) {
	stmnt := `
		SELECT AVG(liquidity_fee)
		FROM swaps
		WHERE pool = $1
		AND runeAmt > 0
	`

	var buyFeeAverage uint64
	row := s.db.QueryRow(stmnt, asset.String())

	if err := row.Scan(&buyFeeAverage); err != nil {
		return 0, err
	}

	return buyFeeAverage, nil
}

func (s *Client) poolFeeAverage(asset common.Asset) (uint64, error) {
	sellFeeAverage, err := s.sellFeeAverage(asset)
	if err != nil {
		return 0, err
	}

	buyFeeAverage, err := s.buyFeeAverage(asset)
	if err != nil {
		return 0, err
	}
	return (buyFeeAverage + sellFeeAverage) / 2, nil
}

func (s *Client) sellFeesTotal(asset common.Asset) (uint64, error) {
	stmnt := `
		SELECT SUM(liquidity_fee)
		FROM swaps
		WHERE pool = $1
		AND assetAmt > 0
	`

	var sellFeesTotal uint64
	row := s.db.QueryRow(stmnt, asset.String())

	if err := row.Scan(&sellFeesTotal); err != nil {
		return 0, err
	}

	priceInRune, err := s.GetPriceInRune(asset)
	if err != nil {
		return 0, err
	}
	return uint64(float64(sellFeesTotal) * priceInRune), nil
}

func (s *Client) buyFeesTotal(asset common.Asset) (uint64, error) {
	stmnt := `
		SELECT SUM(liquidity_fee)
		FROM swaps
		WHERE pool = $1
		AND runeAmt > 0
	`

	var buyFeesTotal uint64
	row := s.db.QueryRow(stmnt, asset.String())

	if err := row.Scan(&buyFeesTotal); err != nil {
		return 0, err
	}

	return buyFeesTotal, nil
}

func (s *Client) poolFeesTotal(asset common.Asset) (uint64, error) {
	buyFeesTotal, err := s.buyFeesTotal(asset)
	if err != nil {
		return 0, err
	}

	sellFeesTotal, err := s.sellFeesTotal(asset)
	if err != nil {
		return 0, err
	}
	return buyFeesTotal + sellFeesTotal, nil
}

func (s *Client) sellAssetCount(asset common.Asset) (uint64, error) {
	stmnt := `
		SELECT COUNT(assetAmt)
		FROM swaps
		WHERE pool = $1
		AND assetAmt > 0
	`

	var sellAssetCount uint64
	row := s.db.QueryRow(stmnt, asset.String())

	if err := row.Scan(&sellAssetCount); err != nil {
		return 0, err
	}

	return sellAssetCount, nil
}

func (s *Client) buyAssetCount(asset common.Asset) (uint64, error) {
	stmnt := `
		SELECT COUNT(liquidity_fee)
		FROM swaps
		WHERE pool = $1
		AND runeAmt < 0
	`

	var buyAssetCount uint64
	row := s.db.QueryRow(stmnt, asset.String())

	if err := row.Scan(&buyAssetCount); err != nil {
		return 0, err
	}

	return buyAssetCount, nil
}

func (s *Client) swappingTxCount(asset common.Asset) (uint64, error) {
	stmnt := `
		SELECT COUNT(event_id) FROM swaps WHERE pool = $1
	`

	var swappingTxCount uint64
	row := s.db.QueryRow(stmnt, asset.String())

	if err := row.Scan(&swappingTxCount); err != nil {
		return 0, err
	}

	return swappingTxCount, nil
}

// swappersCount - number of unique swappers on the network
func (s *Client) swappersCount(asset common.Asset) (uint64, error) {
	stmnt := `
		SELECT COUNT(from_address)
		FROM swaps
		WHERE pool = $1
		GROUP BY from_address
	`

	var swappersCount uint64
	row := s.db.QueryRow(stmnt, asset.String())

	if err := row.Scan(&swappersCount); err != nil {
		return 0, err
	}

	return swappersCount, nil
}

// stakeTxCount - number of stakes that occurred on a given pool
func (s *Client) stakeTxCount(asset common.Asset) (uint64, error) {
	stmnt := `
		SELECT COUNT(event_id)
		FROM stakes
		WHERE pool = $1
		AND units > 0
	`

	var stateTxCount uint64
	row := s.db.QueryRow(stmnt, asset.String())

	if err := row.Scan(&stateTxCount); err != nil {
		return 0, err
	}

	return stateTxCount, nil
}

// withdrawTxCount - number of unstakes that occurred on a given pool
func (s *Client) withdrawTxCount(asset common.Asset) (uint64, error) {
	stmnt := `
		SELECT COUNT(event_id)
		FROM stakes
		WHERE pool = $1
		AND units < 0
	`

	var withdrawTxCount uint64
	row := s.db.QueryRow(stmnt, asset.String())

	if err := row.Scan(&withdrawTxCount); err != nil {
		return 0, err
	}

	return withdrawTxCount, nil
}

func (s *Client) stakingTxCount(asset common.Asset) (uint64, error) {
	stakeTxCount, err := s.stakeTxCount(asset)
	if err != nil {
		return 0, err
	}
	withdrawTxCount, err := s.withdrawTxCount(asset)
	if err != nil {
		return 0, err
	}
	stakingTxCount := stakeTxCount + withdrawTxCount

	return stakingTxCount, nil
}

// stakersCount - number of addresses staking on a given pool
func (s *Client) stakersCount(asset common.Asset) (uint64, error) {
	stmnt := `
		SELECT COUNT(sub.from_address)
		FROM (
			SELECT from_address, SUM(units) AS total_units
			FROM stakes
			WHERE pool = $1
			GROUP BY from_address
		) AS sub
		WHERE sub.total_units > 0
	`

	var stakersCount uint64
	row := s.db.QueryRow(stmnt, asset.String())

	if err := row.Scan(&stakersCount); err != nil {
		return 0, err
	}

	return stakersCount, nil
}

func (s *Client) assetROI(asset common.Asset) (float64, error) {
	assetDepth, err := s.assetDepth(asset)
	if err != nil {
		return 0, err
	}
	assetStakedTotal, err := s.assetStakedTotal(asset)
	if err != nil {
		return 0, err
	}

	staked := float64(assetStakedTotal)
	depth := float64(assetDepth)

	var roi float64
	if staked > 0 {
		roi = (depth - staked) / staked
	}

	return roi, nil
}

func (s *Client) assetROI12(asset common.Asset) (float64, error) {
	assetDepth12m, err := s.assetDepth12m(asset)
	if err != nil {
		return 0, err
	}
	assetStakedTotal12m, err := s.assetStakedTotal12m(asset)
	if err != nil {
		return 0, err
	}

	staked := float64(assetStakedTotal12m)
	depth := float64(assetDepth12m)

	var roi float64
	if staked > 0 {
		roi = (depth - staked) / staked
	}

	return roi, nil
}

func (s *Client) runeROI(asset common.Asset) (float64, error) {
	runeDepth, err := s.runeDepth(asset)
	if err != nil {
		return 0, err
	}
	runeStakedTotal, err := s.runeStakedTotal(asset)
	if err != nil {
		return 0, err
	}
	staked := float64(runeStakedTotal)
	depth := float64(runeDepth)

	var roi float64
	if staked > 0 {
		roi = (depth - staked) / staked
	}

	return roi, nil
}

func (s *Client) runeROI12(asset common.Asset) (float64, error) {
	runeDepth12m, err := s.runeDepth12m(asset)
	if err != nil {
		return 0, err
	}
	runeStakedTotal12m, err := s.runeStakedTotal12m(asset)
	if err != nil {
		return 0, err
	}
	staked := float64(runeStakedTotal12m)
	depth := float64(runeDepth12m)

	var roi float64
	if staked > 0 {
		roi = (depth - staked) / staked
	}

	return roi, nil
}

func (s *Client) poolROI(asset common.Asset) (float64, error) {
	assetROI, err := s.assetROI(asset)
	if err != nil {
		return 0, err
	}
	runeROI, err := s.runeROI(asset)
	if err != nil {
		return 0, err
	}

	var roi float64
	if runeROI > 0 {
		roi = (assetROI / runeROI) / 2
	}

	return roi, err
}

func (s *Client) poolROI12(asset common.Asset) (float64, error) {
	assetROI12, err := s.assetROI12(asset)
	if err != nil {
		return 0, err
	}
	runeROI12, err := s.runeROI12(asset)
	if err != nil {
		return 0, err
	}

	var roi float64
	if runeROI12 > 0 {
		roi = (assetROI12 / runeROI12) / 2
	}

	return roi, err
}
