package atomicasset

// Types

type Asset struct {
	ID             string       `json:"asset_id"`
	Contract       string       `json:"contract"`
	Owner          string       `json:"owner"`
	Name           string       `json:"name"`
	IsTransferable bool         `json:"is_transferable"`
	IsBurnable     bool         `json:"is_burnable"`
	TemplateMint   string       `json:"template_mint"`
	Collection     Collection   `json:"collection"`
	Schema         InlineSchema `json:"schema"`
	Template       Template     `json:"template"`
	BackedTokens   []Token      `json:"backed_tokens"`

	Data          map[string]interface{} `json:"data"`
	ImmutableData map[string]interface{} `json:"immutable_data"`
	MutableData   map[string]interface{} `json:"mutable_data"`

	BurnedByAccount string `json:"burned_by_account"`
	BurnedAtBlock   string `json:"burned_at_block"`
	BurnedAtTime    string `json:"burned_at_time"`

	UpdatedAtBlock string `json:"updated_at_block"`
	UpdatedAtTime  string `json:"updated_at_time"`

	TransferedAtBlock string `json:"transferred_at_block"`
	TransferedAtTime  string `json:"transferred_at_time"`

	MintedAtBlock string `json:"minted_at_block"`
	MintedAtTime  string `json:"minted_at_time"`
}

type ListingAsset struct {
	AssetID        string       `json:"asset_id"`
	Contract       string       `json:"contract"`
	Onwer          string       `json:"owner"`
	Name           string       `json:"name"`
	IsTransferable bool         `json:"is_transferable"`
	IsBurnable     bool         `json:"is_burnable"`
	TemplateMint   string       `json:"template_mint"`
	Collection     Collection   `json:"collection"`
	Schema         InlineSchema `json:"schema"`
	Template       Template     `json:"template"`
	BackedTokens   []Token      `json:"backed_tokens"`

	ImmutableData map[string]interface{} `json:"immutable_data"`
	MutableData   map[string]interface{} `json:"mutable_data"`
	Data          map[string]interface{} `json:"data"`

	BurnedByAccount string   `json:"burned_by_account"`
	BurnedAtBlock   string   `json:"burned_at_block"`
	BurnedAtTime    UnixTime `json:"burned_at_time"`

	UpdatedAtBlock string   `json:"updated_at_block"`
	UpdatedAtTime  UnixTime `json:"updated_at_time"`

	TransferedAtBlock string   `json:"transferred_at_block"`
	TransferedAtTime  UnixTime `json:"transferred_at_time"`

	MintedAtBlock string   `json:"minted_at_block"`
	MintedAtTime  UnixTime `json:"minted_at_time"`

	Sales    []Sale    `json:"sales"`
	Auctions []Auction `json:"actions"`
	Prices   []Price   `json:"prices"`
}

type AssetSale struct {
	ID             string   `json:"sale_id"`
	MarketContract string   `json:"market_contract"`
	AuctionID      string   `json:"auction_id"`
	BuyOfferID     string   `json:"buyoffer_id"`
	Price          string   `json:"price"`
	TokenSymbol    string   `json:"token_symbol"`
	TokenPrecision int64    `json:"token_precision"`
	TokenContract  string   `json:"token_contract"`
	Seller         string   `json:"seller"`
	Buyer          string   `json:"buyer"`
	BlockTime      UnixTime `json:"block_time"`
}

// Request Parameters

// AssetsRequestParams holds the parameters for an Asset request
type AssetsRequestParams struct {
	CollectionName          string        `qs:"collection_name,omitempty"`
	CollectionBlacklist     ReqStringList `qs:"collection_blacklist,omitempty"`
	CollectionWhitelist     ReqStringList `qs:"collection_whitelist,omitempty"`
	SchemaName              string        `qs:"schema_name,omitempty"`
	TemplateID              int           `qs:"template_id,omitempty"`
	TemplateWhitelist       ReqIntList    `qs:"template_whitelist,omitempty"`
	TemplateBlacklist       ReqIntList    `qs:"template_blacklist,omitempty"`
	Owner                   string        `qs:"owner,omitempty"`
	Match                   string        `qs:"match,omitempty"`
	MatchImmutableName      string        `qs:"match_immutable_name,omitempty"`
	MatchMutableName        string        `qs:"match_mutable_name,omitempty"`
	HideTemplatesByAccounts string        `qs:"hide_templates_by_accounts,omitempty"`

	IsTransferable          bool `qs:"is_transferable,omitempty"`
	IsBurnable              bool `qs:"is_burnable,omitempty"`
	Burned                  bool `qs:"burned,omitempty"`
	OnlyDuplicatedTemplates bool `qs:"only_duplicated_templates,omitempty"`
	HasBackedTokens         bool `qs:"has_backend_tokens,omitempty"`
	HideOffers              bool `qs:"hide_offers,omitempty"`

	LowerBound string `qs:"lower_bound,omitempty"`
	UpperBound string `qs:"upper_bound,omitempty"`

	Before int `qs:"before,omitempty"`
	After  int `qs:"after,omitempty"`

	Limit int       `qs:"limit,omitempty"`
	Order SortOrder `qs:"order,omitempty"`
	Sort  string    `qs:"sort,omitempty"`
}

// AssetSalesRequestParams holds the parameters for an AssetSales request
type AssetSalesRequestParams struct {
	Buyer  string    `qs:"buyer,omitempty"`
	Seller string    `qs:"seller,omitempty"`
	Symbol string    `qs:"symbol,omitempty"`
	Order  SortOrder `qs:"order,omitempty"`
}

// Responses

type AssetResponse struct {
	APIResponse
	Data Asset
}

type AssetsResponse struct {
	APIResponse
	Data []Asset
}

type SalesResponse struct {
	APIResponse
	Data []AssetSale
}

type AssetLogResponse struct {
	APIResponse
	Data []Log
}

// API Client functions

// GetAssets fetches "/atomicassets/v1/assets" from API
func (c *Client) GetAssets(params AssetsRequestParams) (AssetsResponse, error) {
	var assets AssetsResponse

	r, err := c.fetch("GET", "/atomicassets/v1/assets", params, &assets.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&assets)
	}
	return assets, err
}

// GetAsset fetches "/atomicassets/v1/assets/{asset_id}" from API
func (c *Client) GetAsset(assetID string) (AssetResponse, error) {
	var asset AssetResponse

	r, err := c.fetch("GET", "/atomicassets/v1/assets/"+assetID, nil, &asset.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&asset)
	}
	return asset, err
}

// GetAssetLog fetches "/atomicassets/v1/assets/{asset_id}/logs" from API
func (c *Client) GetAssetLog(assetID string, params LogRequestParams) (AssetLogResponse, error) {
	var logs AssetLogResponse

	r, err := c.fetch("GET", "/atomicassets/v1/assets/"+assetID+"/logs", params, &logs.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&logs)
	}
	return logs, err
}

// GetAssetSales fetches "/atomicmarket/v1/assets/{asset_id}/sales" from API
func (c *Client) GetAssetSales(assetID string, params AssetSalesRequestParams) (SalesResponse, error) {
	var sales SalesResponse

	r, err := c.fetch("GET", "/atomicmarket/v1/assets/"+assetID+"/sales", params, &sales.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&sales)
	}
	return sales, err
}
