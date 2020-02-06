// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// AssetDetail defines model for AssetDetail.
type AssetDetail struct {
	Asset       *Asset   `json:"asset,omitempty"`
	DateCreated *int64   `json:"dateCreated,omitempty"`
	Logo        *string  `json:"logo,omitempty"`
	Name        *string  `json:"name,omitempty"`
	PriceRune   *float64 `json:"priceRune,omitempty"`
}

// Error defines model for Error.
type Error struct {
	Error string `json:"error"`
}

// PoolDetail defines model for PoolDetail.
type PoolDetail struct {
	Asset            *Asset   `json:"asset,omitempty"`
	AssetDepth       *int64   `json:"assetDepth,omitempty"`
	AssetROI         *float64 `json:"assetROI,omitempty"`
	AssetStakedTotal *int64   `json:"assetStakedTotal,omitempty"`
	BuyAssetCount    *int64   `json:"buyAssetCount,omitempty"`
	BuyFeeAverage    *int64   `json:"buyFeeAverage,omitempty"`
	BuyFeesTotal     *int64   `json:"buyFeesTotal,omitempty"`
	BuySlipAverage   *float64 `json:"buySlipAverage,omitempty"`
	BuyTxAverage     *int64   `json:"buyTxAverage,omitempty"`
	BuyVolume        *int64   `json:"buyVolume,omitempty"`
	PoolDepth        *int64   `json:"poolDepth,omitempty"`
	PoolFeeAverage   *int64   `json:"poolFeeAverage,omitempty"`
	PoolFeesTotal    *int64   `json:"poolFeesTotal,omitempty"`
	PoolROI          *float64 `json:"poolROI,omitempty"`
	PoolROI12        *float64 `json:"poolROI12,omitempty"`
	PoolSlipAverage  *float64 `json:"poolSlipAverage,omitempty"`
	PoolStakedTotal  *int64   `json:"poolStakedTotal,omitempty"`
	PoolTxAverage    *int64   `json:"poolTxAverage,omitempty"`
	PoolUnits        *int64   `json:"poolUnits,omitempty"`
	PoolVolume       *int64   `json:"poolVolume,omitempty"`
	PoolVolume24hr   *int64   `json:"poolVolume24hr,omitempty"`
	Price            *float64 `json:"price,omitempty"`
	RuneDepth        *int64   `json:"runeDepth,omitempty"`
	RuneROI          *float64 `json:"runeROI,omitempty"`
	RuneStakedTotal  *int64   `json:"runeStakedTotal,omitempty"`
	SellAssetCount   *int64   `json:"sellAssetCount,omitempty"`
	SellFeeAverage   *int64   `json:"sellFeeAverage,omitempty"`
	SellFeesTotal    *int64   `json:"sellFeesTotal,omitempty"`
	SellSlipAverage  *float64 `json:"sellSlipAverage,omitempty"`
	SellTxAverage    *int64   `json:"sellTxAverage,omitempty"`
	SellVolume       *int64   `json:"sellVolume,omitempty"`
	StakeTxCount     *int64   `json:"stakeTxCount,omitempty"`
	StakersCount     *int64   `json:"stakersCount,omitempty"`
	StakingTxCount   *int64   `json:"stakingTxCount,omitempty"`
	Status           *string  `json:"status,omitempty"`
	SwappersCount    *int64   `json:"swappersCount,omitempty"`
	SwappingTxCount  *int64   `json:"swappingTxCount,omitempty"`
	WithdrawTxCount  *int64   `json:"withdrawTxCount,omitempty"`
}

// Stakers defines model for Stakers.
type Stakers string

// StakersAddressData defines model for StakersAddressData.
type StakersAddressData struct {
	PoolsArray  *[]Asset `json:"poolsArray,omitempty"`
	TotalEarned *int64   `json:"totalEarned,omitempty"`
	TotalROI    *float64 `json:"totalROI,omitempty"`
	TotalStaked *int64   `json:"totalStaked,omitempty"`
}

// StakersAssetData defines model for StakersAssetData.
type StakersAssetData struct {
	Asset           *Asset   `json:"asset,omitempty"`
	AssetEarned     *int64   `json:"assetEarned,omitempty"`
	AssetROI        *float64 `json:"assetROI,omitempty"`
	AssetStaked     *int64   `json:"assetStaked,omitempty"`
	DateFirstStaked *int64   `json:"dateFirstStaked,omitempty"`
	PoolEarned      *int64   `json:"poolEarned,omitempty"`
	PoolROI         *float64 `json:"poolROI,omitempty"`
	PoolStaked      *int64   `json:"poolStaked,omitempty"`
	RuneEarned      *int64   `json:"runeEarned,omitempty"`
	RuneROI         *float64 `json:"runeROI,omitempty"`
	RuneStaked      *int64   `json:"runeStaked,omitempty"`
	StakeUnits      *int64   `json:"stakeUnits,omitempty"`
}

// StatsData defines model for StatsData.
type StatsData struct {
	DailyActiveUsers   *int64 `json:"dailyActiveUsers,omitempty"`
	DailyTx            *int64 `json:"dailyTx,omitempty"`
	MonthlyActiveUsers *int64 `json:"monthlyActiveUsers,omitempty"`
	MonthlyTx          *int64 `json:"monthlyTx,omitempty"`
	PoolCount          *int64 `json:"poolCount,omitempty"`
	TotalAssetBuys     *int64 `json:"totalAssetBuys,omitempty"`
	TotalAssetSells    *int64 `json:"totalAssetSells,omitempty"`
	TotalDepth         *int64 `json:"totalDepth,omitempty"`
	TotalEarned        *int64 `json:"totalEarned,omitempty"`
	TotalStakeTx       *int64 `json:"totalStakeTx,omitempty"`
	TotalStaked        *int64 `json:"totalStaked,omitempty"`
	TotalTx            *int64 `json:"totalTx,omitempty"`
	TotalUsers         *int64 `json:"totalUsers,omitempty"`
	TotalVolume        *int64 `json:"totalVolume,omitempty"`
	TotalVolume24hr    *int64 `json:"totalVolume24hr,omitempty"`
	TotalWithdrawTx    *int64 `json:"totalWithdrawTx,omitempty"`
}

// ThorchainEndpoint defines model for ThorchainEndpoint.
type ThorchainEndpoint struct {
	Address *string `json:"address,omitempty"`
	Chain   *string `json:"chain,omitempty"`
	PubKey  *string `json:"pub_key,omitempty"`
}

// ThorchainEndpoints defines model for ThorchainEndpoints.
type ThorchainEndpoints struct {
	Current *[]ThorchainEndpoint `json:"current,omitempty"`
}

// TxDetails defines model for TxDetails.
type TxDetails struct {
	Date    *int64  `json:"date,omitempty"`
	Events  *Event  `json:"events,omitempty"`
	Gas     *Gas    `json:"gas,omitempty"`
	Height  *int64  `json:"height,omitempty"`
	In      *Tx     `json:"in,omitempty"`
	Options *Option `json:"options,omitempty"`
	Out     *Tx     `json:"out,omitempty"`
	Pool    *Asset  `json:"pool,omitempty"`
	Status  *string `json:"status,omitempty"`
	Type    *string `json:"type,omitempty"`
}

// Asset defines model for asset.
type Asset struct {
	Chain  *string `json:"chain,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Ticker *string `json:"ticker,omitempty"`
}

// Coin defines model for coin.
type Coin struct {
	Amount *int64 `json:"amount,omitempty"`
	Asset  *Asset `json:"asset,omitempty"`
}

// Coins defines model for coins.
type Coins []Coin

// Event defines model for event.
type Event struct {
	Fee        *int64   `json:"fee,omitempty"`
	Slip       *float64 `json:"slip,omitempty"`
	StakeUnits *int64   `json:"stakeUnits,omitempty"`
}

// Gas defines model for gas.
type Gas struct {
	Amount *int64 `json:"amount,omitempty"`
	Asset  *Asset `json:"asset,omitempty"`
}

// Option defines model for option.
type Option struct {
	Asymmetry           *float64 `json:"asymmetry,omitempty"`
	PriceTarget         *int64   `json:"priceTarget,omitempty"`
	WithdrawBasisPoints *int64   `json:"withdrawBasisPoints,omitempty"`
}

// Tx defines model for tx.
type Tx struct {
	Address *string `json:"address,omitempty"`
	Coins   *Coins  `json:"coins,omitempty"`
	Memo    *string `json:"memo,omitempty"`
	TxID    *string `json:"txID,omitempty"`
}

// AssetsDetailedResponse defines model for AssetsDetailedResponse.
type AssetsDetailedResponse AssetDetail

// GeneralErrorResponse defines model for GeneralErrorResponse.
type GeneralErrorResponse Error

// PoolsDetailedResponse defines model for PoolsDetailedResponse.
type PoolsDetailedResponse PoolDetail

// PoolsResponse defines model for PoolsResponse.
type PoolsResponse []Asset

// StakersAddressDataResponse defines model for StakersAddressDataResponse.
type StakersAddressDataResponse StakersAddressData

// StakersAssetDataResponse defines model for StakersAssetDataResponse.
type StakersAssetDataResponse StakersAssetData

// StakersResponse defines model for StakersResponse.
type StakersResponse []Stakers

// StatsResponse defines model for StatsResponse.
type StatsResponse StatsData

// ThorchainEndpointsResponse defines model for ThorchainEndpointsResponse.
type ThorchainEndpointsResponse ThorchainEndpoints

// TxDetailedResponse defines model for TxDetailedResponse.
type TxDetailedResponse []TxDetails

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get Asset Information// (GET /v1/assets/{asset})
	GetAssetInfo(ctx echo.Context, asset string) error
	// Get Documents// (GET /v1/doc)
	GetDocs(ctx echo.Context) error
	// Get Health// (GET /v1/health)
	GetHealth(ctx echo.Context) error
	// Get Asset Pools// (GET /v1/pools)
	GetPools(ctx echo.Context) error
	// Get Pools Data// (GET /v1/pools/{asset})
	GetPoolsData(ctx echo.Context, asset string) error
	// Get Stakers// (GET /v1/stakers)
	GetStakersData(ctx echo.Context) error
	// Get Staker Data// (GET /v1/stakers/{address})
	GetStakersAddressData(ctx echo.Context, address string) error
	// Get Staker Pool Data// (GET /v1/stakers/{address}/{asset})
	GetStakersAddressAndAssetData(ctx echo.Context, address string, asset string) error
	// Get Global Stats// (GET /v1/stats)
	GetStats(ctx echo.Context) error
	// Get Swagger// (GET /v1/swagger.json)
	GetSwagger(ctx echo.Context) error
	// Get the Proxied Pool Addresses// (GET /v1/thorchain/pool_addresses)
	GetThorchainProxiedEndpoints(ctx echo.Context) error
	// Get transaction// (GET /v1/tx/asset/{asset})
	GetTxDetailsByAsset(ctx echo.Context, asset string) error
	// Get transaction// (GET /v1/tx/{address})
	GetTxDetails(ctx echo.Context, address string) error
	// Get transaction// (GET /v1/tx/{address}/asset/{asset})
	GetTxDetailsByAddressAsset(ctx echo.Context, address string, asset string) error
	// Get transaction// (GET /v1/tx/{address}/txid/{txid})
	GetTxDetailsByAddressTxId(ctx echo.Context, address string, txid string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetAssetInfo converts echo context to params.
func (w *ServerInterfaceWrapper) GetAssetInfo(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "asset" -------------
	var asset string

	err = runtime.BindStyledParameter("simple", false, "asset", ctx.Param("asset"), &asset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter asset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAssetInfo(ctx, asset)
	return err
}

// GetDocs converts echo context to params.
func (w *ServerInterfaceWrapper) GetDocs(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDocs(ctx)
	return err
}

// GetHealth converts echo context to params.
func (w *ServerInterfaceWrapper) GetHealth(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetHealth(ctx)
	return err
}

// GetPools converts echo context to params.
func (w *ServerInterfaceWrapper) GetPools(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPools(ctx)
	return err
}

// GetPoolsData converts echo context to params.
func (w *ServerInterfaceWrapper) GetPoolsData(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "asset" -------------
	var asset string

	err = runtime.BindStyledParameter("simple", false, "asset", ctx.Param("asset"), &asset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter asset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPoolsData(ctx, asset)
	return err
}

// GetStakersData converts echo context to params.
func (w *ServerInterfaceWrapper) GetStakersData(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetStakersData(ctx)
	return err
}

// GetStakersAddressData converts echo context to params.
func (w *ServerInterfaceWrapper) GetStakersAddressData(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameter("simple", false, "address", ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetStakersAddressData(ctx, address)
	return err
}

// GetStakersAddressAndAssetData converts echo context to params.
func (w *ServerInterfaceWrapper) GetStakersAddressAndAssetData(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameter("simple", false, "address", ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	// ------------- Path parameter "asset" -------------
	var asset string

	err = runtime.BindStyledParameter("simple", false, "asset", ctx.Param("asset"), &asset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter asset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetStakersAddressAndAssetData(ctx, address, asset)
	return err
}

// GetStats converts echo context to params.
func (w *ServerInterfaceWrapper) GetStats(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetStats(ctx)
	return err
}

// GetSwagger converts echo context to params.
func (w *ServerInterfaceWrapper) GetSwagger(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetSwagger(ctx)
	return err
}

// GetThorchainProxiedEndpoints converts echo context to params.
func (w *ServerInterfaceWrapper) GetThorchainProxiedEndpoints(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetThorchainProxiedEndpoints(ctx)
	return err
}

// GetTxDetailsByAsset converts echo context to params.
func (w *ServerInterfaceWrapper) GetTxDetailsByAsset(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "asset" -------------
	var asset string

	err = runtime.BindStyledParameter("simple", false, "asset", ctx.Param("asset"), &asset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter asset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTxDetailsByAsset(ctx, asset)
	return err
}

// GetTxDetails converts echo context to params.
func (w *ServerInterfaceWrapper) GetTxDetails(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameter("simple", false, "address", ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTxDetails(ctx, address)
	return err
}

// GetTxDetailsByAddressAsset converts echo context to params.
func (w *ServerInterfaceWrapper) GetTxDetailsByAddressAsset(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameter("simple", false, "address", ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	// ------------- Path parameter "asset" -------------
	var asset string

	err = runtime.BindStyledParameter("simple", false, "asset", ctx.Param("asset"), &asset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter asset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTxDetailsByAddressAsset(ctx, address, asset)
	return err
}

// GetTxDetailsByAddressTxId converts echo context to params.
func (w *ServerInterfaceWrapper) GetTxDetailsByAddressTxId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameter("simple", false, "address", ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	// ------------- Path parameter "txid" -------------
	var txid string

	err = runtime.BindStyledParameter("simple", false, "txid", ctx.Param("txid"), &txid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter txid: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTxDetailsByAddressTxId(ctx, address, txid)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router runtime.EchoRouter, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/v1/assets/:asset", wrapper.GetAssetInfo)
	router.GET("/v1/doc", wrapper.GetDocs)
	router.GET("/v1/health", wrapper.GetHealth)
	router.GET("/v1/pools", wrapper.GetPools)
	router.GET("/v1/pools/:asset", wrapper.GetPoolsData)
	router.GET("/v1/stakers", wrapper.GetStakersData)
	router.GET("/v1/stakers/:address", wrapper.GetStakersAddressData)
	router.GET("/v1/stakers/:address/:asset", wrapper.GetStakersAddressAndAssetData)
	router.GET("/v1/stats", wrapper.GetStats)
	router.GET("/v1/swagger.json", wrapper.GetSwagger)
	router.GET("/v1/thorchain/pool_addresses", wrapper.GetThorchainProxiedEndpoints)
	router.GET("/v1/tx/asset/:asset", wrapper.GetTxDetailsByAsset)
	router.GET("/v1/tx/:address", wrapper.GetTxDetails)
	router.GET("/v1/tx/:address/asset/:asset", wrapper.GetTxDetailsByAddressAsset)
	router.GET("/v1/tx/:address/txid/:txid", wrapper.GetTxDetailsByAddressTxId)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"/xMAAP//vojh2ctDAAA=",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}

