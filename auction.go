package atomicasset

import (
	"fmt"

	"github.com/eosswedenorg-go/unixtime"
)

// Types

type Bid struct {
	Number         int           `json:"number"`
	Account        string        `json:"account"`
	Amount         string        `json:"amount"`
	CreatedAtBlock string        `json:"created_at_block"`
	CreatedAtTime  unixtime.Time `json:"created_at_time"`
	TxID           string        `json:"txid"`
}

type Auction struct {
	ID               string        `json:"auction_id"`
	MarketContract   string        `json:"market_contract"`
	AssetsContract   string        `json:"assets_contract"`
	Seller           string        `json:"seller"`
	Buyer            string        `json:"buyer"`
	Price            Token         `json:"price"`
	Assets           []Asset       `json:"assets"`
	Bids             []Bid         `json:"bids"`
	MakerMarketplace string        `json:"maker_marketplace"`
	TakerMarketplace string        `json:"taker_marketplace"`
	ClaimedByBuyer   bool          `json:"claimed_by_buyer"`
	ClaimedBySeller  bool          `json:"claimed_by_seller"`
	Collection       Collection    `json:"collection"`
	EndTime          unixtime.Time `json:"end_time"`
	IsSellerContract bool          `json:"is_seller_contract"`
	UpdatedAtBlock   string        `json:"updated_at_block"`
	UpdatedAtTime    unixtime.Time `json:"updated_at_time"`
	CreatedAtBlock   string        `json:"created_at_block"`
	CreatedAtTime    unixtime.Time `json:"created_at_time"`
	State            SalesState    `json:"state"`
}

type AuctionsRequestParams struct {
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

// Responses

type AuctionResponse struct {
	APIResponse
	Data Auction
}

type AuctionsResponse struct {
	APIResponse
	Data []Auction
}

// API Client functions

// GetAuction fetches "/atomicassets/v1/auctions/{auction_id}" from API
func (c *Client) GetAuction(auction_id int) (AuctionResponse, error) {
	var resp AuctionResponse

	r, err := c.fetch("GET", fmt.Sprintf("/atomicmarket/v1/auctions/%d", auction_id), nil, &resp.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}

// GetAuctionLogs fetches "/atomicassets/v1/auctions/{auction_id}/logs" from API
func (c *Client) GetAuctionLogs(auction_id int, params LogRequestParams) (LogsResponse, error) {
	var resp LogsResponse

	r, err := c.fetch("GET", fmt.Sprintf("/atomicmarket/v1/auctions/%d/logs", auction_id), params, &resp.APIResponse)
	if err == nil {

		// Set HTTPStatusCode
		resp.HTTPStatusCode = r.StatusCode

		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}

// GetAuctions fetches "/atomicassets/v2/auctions" from API
func (c *Client) GetAuctions(params AuctionsRequestParams) (AuctionsResponse, error) {
	var resp AuctionsResponse

	r, err := c.fetch("GET", "/atomicmarket/v2/auctions", params, &resp.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}
