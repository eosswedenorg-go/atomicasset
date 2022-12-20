package atomicasset

// Types

type Marketplace struct {
	Name           string   `json:"marketplace_name"`
	Creator        string   `json:"creator"`
	CreatedAtBlock string   `json:"created_at_block"`
	CreatedAtTime  UnixTime `json:"created_at_time"`
}

// Request Parameters

// Responses

type MarketplaceResponse struct {
	APIResponse
	Data Marketplace
}

type MarketplacesResponse struct {
	APIResponse
	Data []Marketplace
}

// API Client functions

// GetMarketplace fetches "/atomicassets/v1/marketplaces/{name}" from API
func (c *Client) GetMarketplace(name string) (MarketplaceResponse, error) {
	var resp MarketplaceResponse

	r, err := c.fetch("GET", "/atomicassets/v1/marketplaces/"+name, nil, &resp.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}

// GetMarketplaces fetches "/atomicassets/v1/marketplaces" from API
func (c *Client) GetMarketplaces() (MarketplacesResponse, error) {
	var resp MarketplacesResponse

	r, err := c.fetch("GET", "/atomicassets/v1/marketplaces", nil, &resp.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}
