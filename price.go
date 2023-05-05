package atomicasset

import (
	"github.com/eosswedenorg-go/unixtime"
)

// Types

type PriceSale struct {
	SaleID         string        `json:"sale_id"`
	AuctionID      string        `json:"auction_id"`
	BuyofferID     string        `json:"buyoffer_id"`
	TemplateMint   int64         `json:"template_mint"`
	Price          string        `json:"price"`
	TokenSymbol    string        `json:"token_symbol"`
	TokenPrecision int64         `json:"token_precision"`
	TokenContract  string        `json:"token_contract"`
	BlockTime      unixtime.Time `json:"block_time"`
}

type PriceSaleDay struct {
	Median         string        `json:"median"`
	Average        string        `json:"average"`
	Sales          string        `json:"sales"`
	TokenSymbol    string        `json:"token_symbol"`
	TokenPrecision int64         `json:"token_precision"`
	TokenContract  string        `json:"token_contract"`
	Time           unixtime.Time `json:"time"`
}

type PriceTemplate struct {
	MarketContract   string `json:"market_contract"`
	AssetsContract   string `json:"assets_contract"`
	CollectionName   string `json:"collection_name"`
	TemplateID       string `json:"template_id"`
	TokenSymbol      string `json:"token_symbol"`
	TokenPrecision   int64  `json:"token_precision"`
	TokenContract    string `json:"token_contract"`
	Median           string `json:"median"`
	Average          string `json:"average"`
	Min              string `json:"min"`
	Max              string `json:"max"`
	Sales            string `json:"sales"`
	SuggestedMedian  string `json:"suggested_median"`
	SuggestedAverage string `json:"suggested_average"`
}

type PriceAsset struct {
	TokenSymbol      string `json:"token_symbol"`
	TokenPrecision   int64  `json:"token_precision"`
	TokenContract    string `json:"token_contract"`
	Median           string `json:"median"`
	Average          string `json:"average"`
	Min              string `json:"min"`
	Max              string `json:"max"`
	SuggestedMedian  string `json:"suggested_median"`
	SuggestedAverage string `json:"suggested_average"`
}

type PriceInventory struct {
	Collection Collection   `json:"collection"`
	Prices     []PriceAsset `json:"prices"`
}

// Request Parameters

type PriceSalesRequestParams struct {
	Collection string `qs:"collection_name,omitempty"`
	Schema     string `qs:"schema_name,omitempty"`
	TemplateID int    `qs:"template_id,omitempty"`
	Burned     bool   `qs:"burned,omitempty"`
	Symbol     string `qs:"symbol,omitempty"`
}

type PriceTemplatesRequestParams struct {
	Collection string    `qs:"collection_name,omitempty"`
	Schema     string    `qs:"schema_name,omitempty"`
	TemplateID int       `qs:"template_id,omitempty"`
	Burned     bool      `qs:"burned,omitempty"`
	Symbol     string    `qs:"symbol,omitempty"`
	Page       int       `qs:"page,omitempty"`
	Limit      int       `qs:"limit,omitempty"`
	Order      SortOrder `qs:"order,omitempty"`
}

type PriceAssetsRequestParams struct {
	CollectionName      string          `qs:"collection_name,omitempty"`
	CollectionBlacklist ReqList[string] `qs:"collection_blacklist,omitempty"`
	CollectionWhitelist ReqList[string] `qs:"collection_whitelist,omitempty"`
	SchemaName          string          `qs:"schema_name,omitempty"`
	TemplateID          int             `qs:"template_id,omitempty"`
	Owner               string          `qs:"owner,omitempty"`
	Search              string          `qs:"search,omitempty"`
	Match               string          `qs:"match,omitempty"`
	MatchImmutableName  string          `qs:"match_immutable_name,omitempty"`
	MatchMutableName    string          `qs:"match_mutable_name,omitempty"`
	IsTransferable      bool            `qs:"is_transferable,omitempty"`
	IsBurnable          bool            `qs:"is_burnable,omitempty"`
	Burned              bool            `qs:"burned,omitempty"`
	Minter              string          `qs:"minter,omitempty"`
	Burner              string          `qs:"burner,omitempty"`
	InitialReceiver     string          `qs:"initial_receiver,omitempty"`
	HideOffers          bool            `qs:"hide_offers,omitempty"`
	Ids                 ReqList[string] `qs:"ids,omitempty"`
	LowerBound          string          `qs:"lower_bound,omitempty"`
	UpperBound          string          `qs:"upper_bound,omitempty"`
}

// These "should" have identical fields, so just make an alias for now.
type PriceInventoryRequestParams PriceAssetsRequestParams

// Responses

type SalePricesResponse struct {
	APIResponse
	Data []PriceSale
}

type SalePricesDaysResponse struct {
	APIResponse
	Data []PriceSaleDay
}

type PriceTemplatesResponse struct {
	APIResponse
	Data []PriceTemplate
}

type PriceAssetsResponse struct {
	APIResponse
	Data []PriceAsset
}

type PriceInventoryResponse struct {
	APIResponse
	Data []PriceInventory
}

// API Client functions

// GetSalePrices fetches "/atomicassets/v1/prices/sales" from API
func (c *Client) GetSalePrices(params PriceSalesRequestParams) (SalePricesResponse, error) {
	var resp SalePricesResponse

	r, err := c.fetch("GET", "/atomicassets/v1/prices/sales", params, &resp.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}

// GetSalePricesDays fetches "/atomicassets/v1/prices/sales/days" from API
func (c *Client) GetSalePricesDays(params PriceSalesRequestParams) (SalePricesDaysResponse, error) {
	var resp SalePricesDaysResponse

	r, err := c.fetch("GET", "/atomicassets/v1/prices/sales/days", params, &resp.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}

// GetPriceTemplates fetches "/atomicassets/v1/prices/templates" from API
func (c *Client) GetPriceTemplates(params PriceTemplatesRequestParams) (PriceTemplatesResponse, error) {
	var resp PriceTemplatesResponse

	r, err := c.fetch("GET", "/atomicassets/v1/prices/templates", params, &resp.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}

// GetPriceAssets fetches "/atomicassets/v1/prices/assets" from API
func (c *Client) GetPriceAssets(params PriceAssetsRequestParams) (PriceAssetsResponse, error) {
	var resp PriceAssetsResponse

	r, err := c.fetch("GET", "/atomicassets/v1/prices/assets", params, &resp.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}

// GetPriceInventory fetches "/atomicassets/v1/prices/inventory/{account}" from API
func (c *Client) GetPriceInventory(account string, params PriceInventoryRequestParams) (PriceInventoryResponse, error) {
	var resp PriceInventoryResponse

	// Bit of a hack to extract nested collection array.
	var raw struct {
		APIResponse
		Data struct {
			Collections []PriceInventory `json:"collections"`
		} `json:"data"`
	}

	r, err := c.fetch("GET", "/atomicassets/v1/prices/inventory/"+account, params, &resp.APIResponse)

	if err == nil {
		// Parse json
		err = r.Unmarshal(&raw)
		if err == nil {
			resp.Data = raw.Data.Collections
			resp.Success = raw.Success
			resp.QueryTime = raw.QueryTime
		}
	}
	return resp, err
}
