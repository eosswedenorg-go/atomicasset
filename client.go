package atomicasset

import (
	"fmt"
	"strings"

	"github.com/imroc/req/v3"
	"github.com/sonh/qs"
)

// Client interacts with the api
type Client struct {
	URL  string
	Host string
}

// New Creates a new client object
func New(url string) *Client {
	return &Client{
		URL: url,
	}
}

func isContentType(t string, expected string) bool {
	p := strings.IndexByte(t, ';')
	if p >= 0 {
		t = t[:p]
	}
	return t == expected
}

func (c *Client) send(method string, path string, params interface{}) (*req.Response, error) {
	r := req.C().R()

	if params != nil {
		query, err := qs.NewEncoder().Values(params)
		if err != nil {
			return nil, err
		}
		r.SetQueryString(query.Encode())
	}

	if len(c.Host) > 0 {
		r.SetHeader("Host", c.Host)
	}

	resp, err := r.Send(method, c.URL+path)
	if err != nil {
		return nil, err
	}

	t := resp.GetContentType()
	if !isContentType(t, "application/json") {
		return nil, fmt.Errorf("invalid content-type '%s', expected 'application/json'", t)
	}

	if resp.IsError() {
		apiErr := APIError{}
		if resp.Unmarshal(&apiErr) == nil && apiErr.Success.Valid && !apiErr.Success.Bool {
			return nil, fmt.Errorf("API Error: %s", apiErr.Message.String)
		}
	}

	return resp, err
}

// GetHealth fetches "/health" from API
func (c *Client) GetHealth() (Health, error) {
	var health Health

	r, err := c.send("GET", "/health", nil)
	if err == nil {

		// Set HTTPStatusCode
		health.HTTPStatusCode = r.StatusCode

		// Parse json
		err = r.Unmarshal(&health)
	}
	return health, err
}

// GetAsset fetches "/atomicassets/v1/assets/{asset_id}" from API
func (c *Client) GetAsset(assetID string) (AssetResponse, error) {
	var asset AssetResponse

	r, err := c.send("GET", "/atomicassets/v1/assets/"+assetID, nil)
	if err == nil {

		// Set HTTPStatusCode
		asset.HTTPStatusCode = r.StatusCode

		// Parse json
		err = r.Unmarshal(&asset)
	}
	return asset, err
}

// GetAssets fetches "/atomicassets/v1/assets" from API
func (c *Client) GetAssets(params AssetsRequestParams) (AssetsResponse, error) {
	var assets AssetsResponse

	r, err := c.send("GET", "/atomicassets/v1/assets", params)
	if err == nil {

		// Set HTTPStatusCode
		assets.HTTPStatusCode = r.StatusCode

		// Parse json
		err = r.Unmarshal(&assets)
	}
	return assets, err
}

// GetAssetLog fetches "/atomicassets/v1/assets/{asset_id}/logs" from API
func (c *Client) GetAssetLog(assetID string, params LogRequestParams) (AssetLogResponse, error) {
	var logs AssetLogResponse

	r, err := c.send("GET", "/atomicassets/v1/assets/"+assetID+"/logs", params)
	if err == nil {

		// Set HTTPStatusCode
		logs.HTTPStatusCode = r.StatusCode

		// Parse json
		err = r.Unmarshal(&logs)
	}
	return logs, err
}

// GetAssetSales fetches "/atomicmarket/v1/assets/{asset_id}/sales" from API
func (c *Client) GetAssetSales(assetID string, params AssetSalesRequestParams) (SalesResponse, error) {
	var sales SalesResponse

	r, err := c.send("GET", "/atomicmarket/v1/assets/"+assetID+"/sales", params)
	if err == nil {

		// Set HTTPStatusCode
		sales.HTTPStatusCode = r.StatusCode

		// Parse json
		err = r.Unmarshal(&sales)
	}
	return sales, err
}

// GetCollections fetches "/atomicassets/v1/collections" from API
func (c *Client) GetCollections(params CollectionsRequestParams) (CollectionsResponse, error) {
	var resp CollectionsResponse

	r, err := c.send("GET", "/atomicassets/v1/collections", params)
	if err == nil {

		// Set HTTPStatusCode
		resp.HTTPStatusCode = r.StatusCode

		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}

// GetCollection fetches "/atomicassets/v1/collection/<name>" from API
func (c *Client) GetCollection(name string) (CollectionResponse, error) {
	var resp CollectionResponse

	r, err := c.send("GET", "/atomicassets/v1/collection/"+name, nil)
	if err == nil {

		// Set HTTPStatusCode
		resp.HTTPStatusCode = r.StatusCode

		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}
