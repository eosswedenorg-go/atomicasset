package atomicasset

// Types

type Collection struct {
	CollectionName     string   `json:"collection_name"`
	Contract           string   `json:"contract"`
	Name               string   `json:"name"`
	Image              string   `json:"img"` // Not defined in the spec. but might be included in a response.
	Author             string   `json:"author"`
	AllowNotify        bool     `json:"allow_notify"`
	AuthorizedAccounts []string `json:"authorized_accounts"`
	NotifyAccounts     []string `json:"notify_accounts"`
	MarketFee          float64  `json:"market_fee"`

	Data map[string]interface{} `json:"data"`

	CreatedAtBlock string   `json:"created_at_block"`
	CreatedAtTime  UnixTime `json:"created_at_time"`
}

type CollectionStats struct {
	Assets           string   `json:"assets"`
	Burned           string   `json:"burned"`
	BurnedByTemplate []string `json:"burned_by_template"`
	BurnedBySchema   []string `json:"burned_by_schema"`
	Templates        string   `json:"templates"`
	Schemas          string   `json:"schemas"`
}

// Request Parameters

type CollectionSortColumn string

const (

	// CollectionSortDefault sorts by the default column (created)
	CollectionSortDefault CollectionSortColumn = ""

	// CollectionSortCreated sorts by the created column
	CollectionSortCreated CollectionSortColumn = "created"

	// CollectionSortName sorts by the collection name column
	CollectionSortName CollectionSortColumn = "collection_name"
)

type CollectionsRequestParams struct {
	Author            string               `qs:"author,omitempty"`
	Match             string               `qs:"match,omitempty"`
	AuthorizedAccount string               `qs:"authorized_account,omitempty"`
	NotifyAccount     string               `qs:"notify_account,omitempty"`
	Blacklist         string               `qs:"collection_blacklist,omitempty"`
	Whitelist         string               `qs:"collection_whitelist,omitempty"`
	IDs               string               `qs:"ids,omitempty"`
	LowerBound        string               `qs:"lower_bound,omitempty"`
	UpperBound        string               `qs:"upper_bound,omitempty"`
	Before            int                  `qs:"before,omitempty"`
	After             int                  `qs:"after,omitempty"`
	Page              int                  `qs:"page,omitempty"`
	Limit             int                  `qs:"limit,omitempty"`
	Order             SortOrder            `qs:"order,omitempty"`
	Sort              CollectionSortColumn `qs:"sort,omitempty"`
}

type CollectionLogsRequestParams struct {
	Page            int       `qs:"page,omitempty"`
	Limit           int       `qs:"limit,omitempty"`
	Order           SortOrder `qs:"order,omitempty"`
	ActionBlacklist string    `qs:"action_blacklist,omitempty"`
	ActionWhitelist string    `qs:"action_whitelist,omitempty"`
}

// Responses

type CollectionsResponse struct {
	APIResponse
	Data []Collection
}

type CollectionResponse struct {
	APIResponse
	Data Collection
}

type CollectionStatsResponse struct {
	APIResponse
	Data CollectionStats
}

type CollectionLogsResponse struct {
	APIResponse
	Data []Log
}

// Client API Functions

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

// GetCollectionStats fetches "/atomicassets/v1/collection/<name>/stats" from API
func (c *Client) GetCollectionStats(name string) (CollectionStatsResponse, error) {
	var resp CollectionStatsResponse

	r, err := c.send("GET", "/atomicassets/v1/collection/"+name+"/stats", nil)
	if err == nil {

		// Set HTTPStatusCode
		resp.HTTPStatusCode = r.StatusCode

		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}

// GetCollectionStats fetches "/atomicassets/v1/collection/<name>/stats" from API
func (c *Client) GetCollectionLogs(name string, params CollectionLogsRequestParams) (CollectionLogsResponse, error) {
	var resp CollectionLogsResponse

	r, err := c.send("GET", "/atomicassets/v1/collection/"+name+"/logs", params)
	if err == nil {

		// Set HTTPStatusCode
		resp.HTTPStatusCode = r.StatusCode

		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}
