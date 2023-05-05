package atomicasset

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/eosswedenorg-go/unixtime"
)

// Types

// Need to use string here as SalesRequestParams.State needs to
// have a zero value that is not 0. So we use the empty string for that.
type SalesState string

const (
	SalesStateWaiting  = SalesState("0")
	SalesStateListed   = SalesState("1")
	SalesStateCanceled = SalesState("2")
	SalesStateSold     = SalesState("3")
	SalesStateInvalid  = SalesState("4")
)

// And for json, we need to parse the value as and integer
// and then convert it to string.
func (s *SalesState) UnmarshalJSON(b []byte) error {
	var n int64
	err := json.Unmarshal(b, &n)
	if err == nil {
		v := strconv.FormatInt(n, 10)
		*s = SalesState(v)
	}
	return err
}

type Sale struct {
	ID               string      `json:"sale_id"`
	MarketContract   string      `json:"market_contract"`
	AsssetsContract  string      `json:"assets_contract"`
	Seller           string      `json:"seller"`
	Buyer            string      `json:"buyer"`
	OfferID          string      `json:"offer_id"`
	Price            Token       `json:"price"`
	ListingPrice     json.Number `json:"listing_price"`
	ListingSymbol    string      `json:"listing_symbol"`
	Assets           []Asset     `json:"assets"`
	MakerMarketplace string      `json:"maker_marketplace,omitempty"`
	TakerMarketplace string      `json:"taker_marketplace,omitempty"`
	Collection       Collection  `json:"collection"`
	IsSellerContract bool        `json:"is_seller_contract"`
	State            SalesState  `json:"state"`

	UpdatedAtBlock string        `json:"updated_at_block"`
	UpdatedAtTime  unixtime.Time `json:"updated_at_time"`

	CreatedAtBlock string        `json:"created_at_block"`
	CreatedAtTime  unixtime.Time `json:"created_at_time"`
}

// Request Parameters

type SaleSortColumn string

const (
	SaleSortCreated      = SaleSortColumn("created")
	SaleSortUpdated      = SaleSortColumn("updated")
	SaleSortID           = SaleSortColumn("sale_id")
	SaleSortPrice        = SaleSortColumn("price")
	SaleSortTemplateMint = SaleSortColumn("template_mint")
	SaleSortName         = SaleSortColumn("name")
)

type SalesRequestParams struct {
	State               SalesState      `qs:"state,omitempty"`
	MaxAssets           int             `qs:"max_assets,omitempty"`
	MinAssets           int             `qs:"min_assets,omitempty"`
	ShowSellerContract  string          `qs:"show_seller_contract,omitempty"`
	ContractBlacklist   ReqList[string] `qs:"contract_blacklist,omitempty"`
	ContractWhitelist   ReqList[string] `qs:"contract_whitelist,omitempty"`
	SellerBlacklist     ReqList[string] `qs:"seller_blacklist,omitempty"`
	BuyerBlacklist      ReqList[string] `qs:"buyer_blacklist,omitempty"`
	AssetId             int             `qs:"asset_id,omitempty"`
	Marketplace         ReqList[string] `qs:"marketplace,omitempty"`
	MakerMarketplace    ReqList[string] `qs:"maker_marketplace,omitempty"`
	TakerMarketplace    ReqList[string] `qs:"taker_marketplace,omitempty"`
	Symbol              string          `qs:"symbol,omitempty"`
	Account             string          `qs:"account,omitempty"`
	Seller              ReqList[string] `qs:"seller,omitempty"`
	Buyer               ReqList[string] `qs:"buyer,omitempty"`
	MinPrice            int             `qs:"min_price,omitempty"`
	MaxPrice            int             `qs:"max_price,omitempty"`
	MinTemplateMint     int             `qs:"min_template_mint,omitempty"`
	MaxTemplateMint     int             `qs:"max_template_mint,omitempty"`
	CollectionName      string          `qs:"collection_name,omitempty"`
	CollectionBlacklist ReqList[string] `qs:"collection_blacklist,omitempty"`
	CollectionWhitelist ReqList[string] `qs:"collection_whitelist,omitempty"`
	SchemaName          string          `qs:"schema_name,omitempty"`
	TemplateID          int             `qs:"template_id,omitempty"`
	Burned              bool            `qs:"burned,omitempty"`
	Owner               string          `qs:"owner,omitempty"`
	Match               string          `qs:"match,omitempty"`
	Search              string          `qs:"search,omitempty"`
	MatchImmutableName  string          `qs:"match_immutable_name,omitempty"`
	MatchMutableName    string          `qs:"match_mutable_name,omitempty"`
	IsTransferable      bool            `qs:"is_transferable,omitempty"`
	IsBurnable          bool            `qs:"is_burnable,omitempty"`
	Minter              string          `qs:"minter,omitempty"`
	Burner              string          `qs:"burner,omitempty"`
	IDs                 ReqList[int]    `qs:"ids,omitempty"`
	LowerBound          string          `qs:"lower_bound,omitempty"`
	UpperBound          string          `qs:"upper_bound,omitempty"`
	Before              int             `qs:"before,omitempty"`
	After               int             `qs:"after,omitempty"`
	Page                int             `qs:"page,omitempty"`
	Limit               int             `qs:"limit,omitempty"`
	Order               SortOrder       `qs:"order,omitempty"`
	Sort                SaleSortColumn  `qs:"sort,omitempty"`
}

type SaleTemplateSortColumn string

const (
	SaleTemplateSortPrice      = SaleTemplateSortColumn("price")
	SaleTemplateSortTemplateID = SaleTemplateSortColumn("template_id")
)

type SalesTemplateRequestParams struct {
	Symbol              string                 `qs:"symbol,omitempty"`
	MinPrice            int                    `qs:"min_price,omitempty"`
	MaxPrice            int                    `qs:"max_price,omitempty"`
	CollectionName      string                 `qs:"collection_name,omitempty"`
	CollectionBlacklist ReqList[string]        `qs:"collection_blacklist,omitempty"`
	CollectionWhitelist ReqList[string]        `qs:"collection_whitelist,omitempty"`
	SchemaName          string                 `qs:"schema_name,omitempty"`
	TemplateID          int                    `qs:"template_id,omitempty"`
	Burned              bool                   `qs:"burned,omitempty"`
	Owner               string                 `qs:"owner,omitempty"`
	Match               string                 `qs:"match,omitempty"`
	Search              string                 `qs:"search,omitempty"`
	MatchImmutableName  string                 `qs:"match_immutable_name,omitempty"`
	MatchMutableName    string                 `qs:"match_mutable_name,omitempty"`
	IsTransferable      bool                   `qs:"is_transferable,omitempty"`
	IsBurnable          bool                   `qs:"is_burnable,omitempty"`
	Minter              string                 `qs:"minter,omitempty"`
	Burner              string                 `qs:"burner,omitempty"`
	InitialReceiver     string                 `qs:"initial_receiver,omitempty"`
	IDs                 ReqList[int]           `qs:"ids,omitempty"`
	LowerBound          string                 `qs:"lower_bound,omitempty"`
	UpperBound          string                 `qs:"upper_bound,omitempty"`
	Before              int                    `qs:"before,omitempty"`
	After               int                    `qs:"after,omitempty"`
	Page                int                    `qs:"page,omitempty"`
	Limit               int                    `qs:"limit,omitempty"`
	Order               SortOrder              `qs:"order,omitempty"`
	Sort                SaleTemplateSortColumn `qs:"sort,omitempty"`
}

// Responses

type SaleResponse struct {
	APIResponse
	Data Sale
}

type SalesResponse struct {
	APIResponse
	Data []Sale
}

// API Client functions

// GetSale fetches "/atomicassets/v1/sales/{sale_id}" from API
func (c *Client) GetSale(sale_id int) (SaleResponse, error) {
	var resp SaleResponse

	r, err := c.fetch("GET", fmt.Sprintf("/atomicmarket/v1/sales/%d", sale_id), nil, &resp.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}

// GetSales fetches "/atomicassets/v2/sales" from API
func (c *Client) GetSales(params SalesRequestParams) (SalesResponse, error) {
	var resp SalesResponse

	r, err := c.fetch("GET", "/atomicmarket/v2/sales", params, &resp.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}

func (c *Client) GetSalesGroupByTemplate(params SalesTemplateRequestParams) (SalesResponse, error) {
	var resp SalesResponse

	r, err := c.fetch("GET", "/atomicmarket/v1/sales/templates", params, &resp.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}

// GetSaleLogs fetches "/atomicassets/v1/sales/{sale_id}/logs" from API
func (c *Client) GetSaleLogs(sale_id int, params LogRequestParams) (LogsResponse, error) {
	var resp LogsResponse

	r, err := c.fetch("GET", fmt.Sprintf("/atomicmarket/v1/sales/%d/logs", sale_id), params, &resp.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}
