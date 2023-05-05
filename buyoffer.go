package atomicasset

import (
	"fmt"

	"github.com/eosswedenorg-go/unixtime"
)

// Types

type BuyOffer struct {
	ID               string     `json:"buyoffer_id"`
	MarketContract   string     `json:"market_contract"`
	AssetsContract   string     `json:"assets_contract"`
	Seller           string     `json:"seller"`
	Buyer            string     `json:"buyer"`
	Price            Token      `json:"price"`
	Assets           []Asset    `json:"assets"`
	MakerMarketplace string     `json:"maker_marketplace,omitempty"`
	TakerMarketplace string     `json:"taker_marketplace,omitempty"`
	Collection       Collection `json:"collection"`
	State            SalesState `json:"state"`
	Memo             string     `json:"memo"`
	DeclineMemo      string     `json:"decline_memo"`

	UpdatedAtBlock string        `json:"updated_at_block"`
	UpdatedAtTime  unixtime.Time `json:"updated_at_time"`

	CreatedAtBlock string        `json:"created_at_block"`
	CreatedAtTime  unixtime.Time `json:"created_at_time"`
}

// Request Parameters

type BuyOfferSortColumn string

const (
	BuyOfferSortCreated      = BuyOfferSortColumn("created")
	BuyOfferSortUpdated      = BuyOfferSortColumn("updated")
	BuyOfferSortID           = BuyOfferSortColumn("buyoffer_id")
	BuyOfferSortPrice        = BuyOfferSortColumn("price")
	BuyOfferSortTemplateMint = BuyOfferSortColumn("template_mint")
	BuyOfferSortName         = BuyOfferSortColumn("name")
)

type BuyOffersRequestParams struct {
	State               SalesState         `qs:"state,omitempty"`
	MaxAssets           int                `qs:"max_assets,omitempty"`
	MinAssets           int                `qs:"min_assets,omitempty"`
	ShowSellerContract  string             `qs:"show_seller_contract,omitempty"`
	ContractBlacklist   ReqList[string]    `qs:"contract_blacklist,omitempty"`
	ContractWhitelist   ReqList[string]    `qs:"contract_whitelist,omitempty"`
	SellerBlacklist     ReqList[string]    `qs:"seller_blacklist,omitempty"`
	BuyerBlacklist      ReqList[string]    `qs:"buyer_blacklist,omitempty"`
	AssetId             int                `qs:"asset_id,omitempty"`
	Marketplace         ReqList[string]    `qs:"marketplace,omitempty"`
	MakerMarketplace    ReqList[string]    `qs:"maker_marketplace,omitempty"`
	TakerMarketplace    ReqList[string]    `qs:"taker_marketplace,omitempty"`
	Symbol              string             `qs:"symbol,omitempty"`
	Account             string             `qs:"account,omitempty"`
	Seller              ReqList[string]    `qs:"seller,omitempty"`
	Buyer               ReqList[string]    `qs:"buyer,omitempty"`
	MinPrice            int                `qs:"min_price,omitempty"`
	MaxPrice            int                `qs:"max_price,omitempty"`
	MinTemplateMint     int                `qs:"min_template_mint,omitempty"`
	MaxTemplateMint     int                `qs:"max_template_mint,omitempty"`
	CollectionName      string             `qs:"collection_name,omitempty"`
	CollectionBlacklist ReqList[string]    `qs:"collection_blacklist,omitempty"`
	CollectionWhitelist ReqList[string]    `qs:"collection_whitelist,omitempty"`
	SchemaName          string             `qs:"schema_name,omitempty"`
	TemplateID          int                `qs:"template_id,omitempty"`
	Burned              bool               `qs:"burned,omitempty"`
	Owner               string             `qs:"owner,omitempty"`
	Match               string             `qs:"match,omitempty"`
	Search              string             `qs:"search,omitempty"`
	MatchImmutableName  string             `qs:"match_immutable_name,omitempty"`
	MatchMutableName    string             `qs:"match_mutable_name,omitempty"`
	IsTransferable      bool               `qs:"is_transferable,omitempty"`
	IsBurnable          bool               `qs:"is_burnable,omitempty"`
	Minter              string             `qs:"minter,omitempty"`
	Burner              string             `qs:"burner,omitempty"`
	InitialReceiver     string             `qs:"initial_receiver,omitempty"`
	IDs                 ReqList[int]       `qs:"ids,omitempty"`
	LowerBound          string             `qs:"lower_bound,omitempty"`
	UpperBound          string             `qs:"upper_bound,omitempty"`
	Before              int                `qs:"before,omitempty"`
	After               int                `qs:"after,omitempty"`
	Page                int                `qs:"page,omitempty"`
	Limit               int                `qs:"limit,omitempty"`
	Order               SortOrder          `qs:"order,omitempty"`
	Sort                BuyOfferSortColumn `qs:"sort,omitempty"`
}

// Responses

type BuyOfferResponse struct {
	APIResponse
	Data BuyOffer
}

type BuyOffersResponse struct {
	APIResponse
	Data []BuyOffer
}

// API Client functions

// GetBuyOffer fetches "/atomicassets/v1/buyoffers/{buyoffer_id}" from API
func (c *Client) GetBuyOffer(buyoffer_id int) (BuyOfferResponse, error) {
	var resp BuyOfferResponse

	r, err := c.fetch("GET", fmt.Sprintf("/atomicmarket/v1/buyoffers/%d", buyoffer_id), nil, &resp.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}

// GetBuyOfferLogs fetches "/atomicassets/v1/buyoffers/{buyoffer_id}/logs" from API
func (c *Client) GetBuyOfferLogs(buyoffer_id int, params LogRequestParams) (LogsResponse, error) {
	var resp LogsResponse

	r, err := c.fetch("GET", fmt.Sprintf("/atomicmarket/v1/buyoffers/%d/logs", buyoffer_id), params, &resp.APIResponse)
	if err == nil {

		// Set HTTPStatusCode
		resp.HTTPStatusCode = r.StatusCode

		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}

// GetBuyOffers fetches "/atomicassets/v1/buyoffers" from API
func (c *Client) GetBuyOffers(params AuctionsRequestParams) (BuyOffersResponse, error) {
	var resp BuyOffersResponse

	r, err := c.fetch("GET", "/atomicmarket/v1/buyoffers", params, &resp.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}
