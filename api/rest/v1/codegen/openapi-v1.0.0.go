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
	"time"
)

// Asset defines model for Asset.
type Asset struct {
	Chain  *string `json:"chain,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Ticker *string `json:"ticker,omitempty"`
}

// AssetDetail defines model for AssetDetail.
type AssetDetail struct {
	Asset       *Asset     `json:"Asset,omitempty"`
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	Logo        *string    `json:"Logo,omitempty"`
	Name        *string    `json:"Name,omitempty"`
	PriceRune   *float64   `json:"priceRune,omitempty"`
	PriceUSD    *float64   `json:"priceUSD,omitempty"`
}

// BEPSwapData defines model for BEPSwapData.
type BEPSwapData struct {
	DAU             *int `json:"DAU,omitempty"`
	DailyTx         *int `json:"DailyTx,omitempty"`
	MAU             *int `json:"MAU,omitempty"`
	MonthlyTx       *int `json:"MonthlyTx,omitempty"`
	PoolCount       *int `json:"PoolCount,omitempty"`
	TotalAssetBuys  *int `json:"TotalAssetBuys,omitempty"`
	TotalAssetSells *int `json:"TotalAssetSells,omitempty"`
	TotalDepth      *int `json:"TotalDepth,omitempty"`
	TotalEarned     *int `json:"TotalEarned,omitempty"`
	TotalStaked     *int `json:"TotalStaked,omitempty"`
	TotalTx         *int `json:"TotalTx,omitempty"`
	TotalUsers      *int `json:"TotalUsers,omitempty"`
	TotalVolume     *int `json:"TotalVolume,omitempty"`
	TotalVolume24hr *int `json:"TotalVolume24hr,omitempty"`
	TotalStakeTx    *int `json:"totalStakeTx,omitempty"`
	TotalWithdrawTx *int `json:"totalWithdrawTx,omitempty"`
}

// Error defines model for Error.
type Error struct {
	Error string `json:"error"`
}

// PoolDetail defines model for PoolDetail.
type PoolDetail struct {
	Asset            *Asset   `json:"Asset,omitempty"`
	AssetDepth       *int     `json:"AssetDepth,omitempty"`
	AssetROI         *float64 `json:"AssetROI,omitempty"`
	AssetStakedTotal *int     `json:"AssetStakedTotal,omitempty"`
	BuyAssetCount    *int     `json:"BuyAssetCount,omitempty"`
	BuyFeeAverage    *int     `json:"BuyFeeAverage,omitempty"`
	BuyFeesTotal     *int     `json:"BuyFeesTotal,omitempty"`
	BuySlipAverage   *int     `json:"BuySlipAverage,omitempty"`
	BuyTxAverage     *int     `json:"BuyTxAverage,omitempty"`
	BuyVolume        *int     `json:"BuyVolume,omitempty"`
	PoolDepth        *int     `json:"PoolDepth,omitempty"`
	PoolFeeAverage   *int     `json:"PoolFeeAverage,omitempty"`
	PoolFeesTotal    *int     `json:"PoolFeesTotal,omitempty"`
	PoolROI          *float64 `json:"PoolROI,omitempty"`
	PoolROI12        *float64 `json:"PoolROI12,omitempty"`
	PoolSlipAverage  *int     `json:"PoolSlipAverage,omitempty"`
	PoolStakedTotal  *int     `json:"PoolStakedTotal,omitempty"`
	PoolTxAverage    *int     `json:"PoolTxAverage,omitempty"`
	PoolUnits        *int     `json:"PoolUnits,omitempty"`
	PoolVolume       *int     `json:"PoolVolume,omitempty"`
	PoolVolume24hr   *int     `json:"PoolVolume24hr,omitempty"`
	Price            *float64 `json:"Price,omitempty"`
	RuneDepth        *int     `json:"RuneDepth,omitempty"`
	RuneROI          *float64 `json:"RuneROI,omitempty"`
	RuneStakedTotal  *int     `json:"RuneStakedTotal,omitempty"`
	SellAssetCount   *int     `json:"SellAssetCount,omitempty"`
	SellFeeAverage   *int     `json:"SellFeeAverage,omitempty"`
	SellFeesTotal    *int     `json:"SellFeesTotal,omitempty"`
	SellSlipAverage  *int     `json:"SellSlipAverage,omitempty"`
	SellTxAverage    *int     `json:"SellTxAverage,omitempty"`
	SellVolume       *int     `json:"SellVolume,omitempty"`
	StakeTxCount     *int     `json:"StakeTxCount,omitempty"`
	StakersCount     *int     `json:"StakersCount,omitempty"`
	StakingTxCount   *int     `json:"StakingTxCount,omitempty"`
	SwappersCount    *int     `json:"SwappersCount,omitempty"`
	SwappingTxCount  *int     `json:"SwappingTxCount,omitempty"`
	WithdrawTxCount  *int     `json:"WithdrawTxCount,omitempty"`
}

// Stakers defines model for Stakers.
type Stakers string

// StakersAddressData defines model for StakersAddressData.
type StakersAddressData struct {
	StakeArray  *[]Asset `json:"StakeArray,omitempty"`
	TotalEarned *int64   `json:"TotalEarned,omitempty"`
	TotalROI    *int64   `json:"TotalROI,omitempty"`
	TotalStaked *int64   `json:"TotalStaked,omitempty"`
}

// StakersAssetData defines model for StakersAssetData.
type StakersAssetData struct {
	Asset           *Asset     `json:"Asset,omitempty"`
	AssetEarned     *int64     `json:"AssetEarned,omitempty"`
	AssetROI        *float64   `json:"AssetROI,omitempty"`
	AssetStaked     *int64     `json:"AssetStaked,omitempty"`
	DateFirstStaked *time.Time `json:"DateFirstStaked,omitempty"`
	PoolEarned      *int64     `json:"PoolEarned,omitempty"`
	PoolROI         *float64   `json:"PoolROI,omitempty"`
	PoolStaked      *int64     `json:"PoolStaked,omitempty"`
	RuneEarned      *int64     `json:"RuneEarned,omitempty"`
	RuneROI         *float64   `json:"RuneROI,omitempty"`
	RuneStaked      *int64     `json:"RuneStaked,omitempty"`
	StakeUnits      *int64     `json:"StakeUnits,omitempty"`
}

// AssetsDetailedResponse defines model for AssetsDetailedResponse.
type AssetsDetailedResponse AssetDetail

// AssetsResponse defines model for AssetsResponse.
type AssetsResponse []Asset

// BEPSwapResponse defines model for BEPSwapResponse.
type BEPSwapResponse BEPSwapData

// GeneralErrorResponse defines model for GeneralErrorResponse.
type GeneralErrorResponse Error

// PoolsDetailedResponse defines model for PoolsDetailedResponse.
type PoolsDetailedResponse PoolDetail

// StakersAddressDataResponse defines model for StakersAddressDataResponse.
type StakersAddressDataResponse []StakersAddressData

// StakersAssetDataResponse defines model for StakersAssetDataResponse.
type StakersAssetDataResponse StakersAssetData

// StakersResponse defines model for StakersResponse.
type StakersResponse []Stakers

// GetStakerTxParams defines parameters for GetStakerTx.
type GetStakerTxParams struct {
	Staker string  `json:"staker"`
	Limit  *int    `json:"limit,omitempty"`
	Offset *int    `json:"offset,omitempty"`
	Asset  *string `json:"asset,omitempty"`
}

// GetSwapTxParams defines parameters for GetSwapTx.
type GetSwapTxParams struct {
	Asset  string `json:"asset"`
	Sender string `json:"sender"`
	Dest   string `json:"dest"`
	Limit  *int   `json:"limit,omitempty"`
	Offset *int   `json:"offset,omitempty"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// (GET /v1/assets)
	GetAssets(ctx echo.Context) error
	// (GET /v1/assets/{asset})
	GetAssetInfo(ctx echo.Context, asset string) error
	// (GET /v1/bepswap)
	GetBEPSwapData(ctx echo.Context) error
	// This swagger/openapi 3.0 generated documentation// (GET /v1/doc)
	GetDocs(ctx echo.Context) error
	// (GET /v1/health)
	GetHealth(ctx echo.Context) error
	// (GET /v1/pools/{asset})
	GetPoolsData(ctx echo.Context, asset string) error
	// (GET /v1/stakerTx)
	GetStakerTx(ctx echo.Context, params GetStakerTxParams) error
	// (GET /v1/stakers)
	GetStakersData(ctx echo.Context) error
	// (GET /v1/stakers/{address})
	GetStakersAddressData(ctx echo.Context, address string) error
	// (GET /v1/stakers/{address}/{asset})
	GetStakersAddressAndAssetData(ctx echo.Context, address string, asset string) error
	// JSON swagger/openapi 3.0 specification endpoint// (GET /v1/swagger.json)
	GetSwagger(ctx echo.Context) error
	// (GET /v1/swapTx)
	GetSwapTx(ctx echo.Context, params GetSwapTxParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetAssets converts echo context to params.
func (w *ServerInterfaceWrapper) GetAssets(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAssets(ctx)
	return err
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

// GetBEPSwapData converts echo context to params.
func (w *ServerInterfaceWrapper) GetBEPSwapData(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetBEPSwapData(ctx)
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

// GetStakerTx converts echo context to params.
func (w *ServerInterfaceWrapper) GetStakerTx(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetStakerTxParams
	// ------------- Required query parameter "staker" -------------
	if paramValue := ctx.QueryParam("staker"); paramValue != "" {

	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Query argument staker is required, but not found"))
	}

	err = runtime.BindQueryParameter("form", true, true, "staker", ctx.QueryParams(), &params.Staker)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter staker: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------
	if paramValue := ctx.QueryParam("limit"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------
	if paramValue := ctx.QueryParam("offset"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// ------------- Optional query parameter "asset" -------------
	if paramValue := ctx.QueryParam("asset"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "asset", ctx.QueryParams(), &params.Asset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter asset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetStakerTx(ctx, params)
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

// GetSwagger converts echo context to params.
func (w *ServerInterfaceWrapper) GetSwagger(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetSwagger(ctx)
	return err
}

// GetSwapTx converts echo context to params.
func (w *ServerInterfaceWrapper) GetSwapTx(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetSwapTxParams
	// ------------- Required query parameter "asset" -------------
	if paramValue := ctx.QueryParam("asset"); paramValue != "" {

	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Query argument asset is required, but not found"))
	}

	err = runtime.BindQueryParameter("form", true, true, "asset", ctx.QueryParams(), &params.Asset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter asset: %s", err))
	}

	// ------------- Required query parameter "sender" -------------
	if paramValue := ctx.QueryParam("sender"); paramValue != "" {

	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Query argument sender is required, but not found"))
	}

	err = runtime.BindQueryParameter("form", true, true, "sender", ctx.QueryParams(), &params.Sender)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter sender: %s", err))
	}

	// ------------- Required query parameter "dest" -------------
	if paramValue := ctx.QueryParam("dest"); paramValue != "" {

	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Query argument dest is required, but not found"))
	}

	err = runtime.BindQueryParameter("form", true, true, "dest", ctx.QueryParams(), &params.Dest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter dest: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------
	if paramValue := ctx.QueryParam("limit"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------
	if paramValue := ctx.QueryParam("offset"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetSwapTx(ctx, params)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router runtime.EchoRouter, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/v1/assets", wrapper.GetAssets)
	router.GET("/v1/assets/:asset", wrapper.GetAssetInfo)
	router.GET("/v1/bepswap", wrapper.GetBEPSwapData)
	router.GET("/v1/doc", wrapper.GetDocs)
	router.GET("/v1/health", wrapper.GetHealth)
	router.GET("/v1/pools/:asset", wrapper.GetPoolsData)
	router.GET("/v1/stakerTx", wrapper.GetStakerTx)
	router.GET("/v1/stakers", wrapper.GetStakersData)
	router.GET("/v1/stakers/:address", wrapper.GetStakersAddressData)
	router.GET("/v1/stakers/:address/:asset", wrapper.GetStakersAddressAndAssetData)
	router.GET("/v1/swagger.json", wrapper.GetSwagger)
	router.GET("/v1/swapTx", wrapper.GetSwapTx)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
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
