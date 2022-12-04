package atomicasset

// Types

type AssetsConfig struct {
	Contract         string         `json:"contract"`
	CollectionFormat []SchemaFormat `json:"collection_format"`
	SupportedTokens  []PriceToken   `json:"supported_tokens"`
	Version          string         `json:"version"`
}

type MarketConfig struct {
	AtomicassetsContract   string       `json:"atomicassets_contract"`
	AtomicmarketContract   string       `json:"atomicmarket_contract"`
	DelphioracleContract   string       `json:"delphioracle_contract"`
	MakerMarketFee         float64      `json:"maker_market_fee"`
	TakerMarketFee         float64      `json:"taker_market_fee"`
	MinimumAuctionDuration int          `json:"minimum_auction_duration"`
	MaximumAuctionDuration int          `json:"maximum_auction_duration"`
	MinimumBidIncrease     float64      `json:"minimum_bid_increase"`
	AuctionResetDuration   int          `json:"auction_reset_duration"`
	SupportedTokens        []PriceToken `json:"supported_tokens"`
	SupportedPairs         []TokenPair  `json:"supported_pairs"`
	Version                string       `json:"version"`
}

type ToolsConfig struct {
	AtomictoolsContract  string `json:"atomictools_contract"`
	AtomicassetsContract string `json:"atomicassets_contract"`
	Version              string `json:"version"`
}

// Request Parameters

// Responses

type AssetsConfigResponse struct {
	APIResponse
	Data AssetsConfig
}

type MarketConfigResponse struct {
	APIResponse
	Data MarketConfig
}

type ToolsConfigResponse struct {
	APIResponse
	Data ToolsConfig
}

// Client API Functions

// GetAssetsConfig fetches "/atomicassets/v1/config" from API
func (c *Client) GetAssetsConfig() (AssetsConfigResponse, error) {
	var resp AssetsConfigResponse

	r, err := c.fetch("GET", "/atomicassets/v1/config", nil, &resp.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}

// GetMarketConfig fetches "/atomicmarket/v1/config" from API
func (c *Client) GetMarketConfig() (MarketConfigResponse, error) {
	var resp MarketConfigResponse

	r, err := c.fetch("GET", "/atomicmarket/v1/config", nil, &resp.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}

// GetToolsConfig fetches "/atomictools/v1/config" from API
func (c *Client) GetToolsConfig() (ToolsConfigResponse, error) {
	var resp ToolsConfigResponse

	r, err := c.fetch("GET", "/atomictools/v1/config", nil, &resp.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}
