package atomicasset

import (
	"fmt"

	"github.com/eosswedenorg-go/unixtime"
)

// Types

type Template struct {
	ID             string                 `json:"template_id"`
	Contract       string                 `json:"contract"`
	MaxSupply      string                 `json:"max_supply"`
	IssuedSupply   string                 `json:"issued_supply"`
	IsTransferable bool                   `json:"is_transferable"`
	IsBurnable     bool                   `json:"is_burnable"`
	ImmutableData  map[string]interface{} `json:"immutable_data"`

	// Note: Collection.Contract is always non-existant in template responses.
	// But cba to create a new struct just for one field.
	Collection     Collection    `json:"collection"`
	Schema         InlineSchema  `json:"schema"`
	CreatedAtBlock string        `json:"created_at_block"`
	CreatedAtTime  unixtime.Time `json:"created_at_time"`
}

type TemplateStats struct {
	Assets    string `json:"assets"`
	Burned    string `json:"burned"`
	Templates string `json:"templates"`
	Schemas   string `json:"schemas"`
}

// Request Parameters

type TemplateSortColumn string

const (

	// TemplateSortDefault sorts by the default column (created)
	TemplateSortDefault TemplateSortColumn = ""

	// TemplateSortCreated sorts by the created column
	TemplateSortCreated TemplateSortColumn = "created"

	// TemplateSortName sorts by the name column
	TemplateSortName TemplateSortColumn = "name"
)

type TemplateRequestParams struct {
	SchemaName          string           `qs:"schema_name,omitempty"`
	CollectionName      string           `qs:"collection_name,omitempty"`
	CollectionBlacklist ReqList[string]  `qs:"collection_blacklist,omitempty"`
	CollectionWhitelist ReqList[string]  `qs:"collection_whitelist,omitempty"`
	IssuedSypply        int              `qs:"issued_supply,omitempty"`
	MinIssuedSupply     int              `qs:"min_issued_supply,omitempty"`
	MaxIssuedSupply     int              `qs:"max_issued_supply,omitempty"`
	MaxSupply           int              `qs:"max_supply,omitempty"`
	HasAssets           bool             `qs:"has_assets,omitempty"`
	IsBurnable          bool             `qs:"is_burnable,omitempty"`
	IsTransferable      bool             `qs:"is_transferable,omitempty"`
	AuthorizedAccount   string           `qs:"authorized_account,omitempty"`
	Match               string           `qs:"match,omitempty"`
	IDs                 ReqList[int]     `qs:"ids,omitempty"`
	LowerBound          string           `qs:"lower_bound,omitempty"`
	UpperBound          string           `qs:"upper_bound,omitempty"`
	Before              int              `qs:"before,omitempty"`
	After               int              `qs:"after,omitempty"`
	Page                int              `qs:"page,omitempty"`
	Limit               int              `qs:"limit,omitempty"`
	Order               SortOrder        `qs:"order,omitempty"`
	Sort                SchemaSortColumn `qs:"sort,omitempty"`
}

// Responses

type TemplatesResponse struct {
	APIResponse
	Data []Template
}

type TemplateResponse struct {
	APIResponse
	Data Template
}

type TemplateStatsResponse struct {
	APIResponse
	Data TemplateStats
}

// Client API Functions

// GetSchemas fetches "/atomicassets/v1/templates" from API
func (c *Client) GetTemplates(params TemplateRequestParams) (TemplatesResponse, error) {
	var resp TemplatesResponse

	r, err := c.fetch("GET", "/atomicassets/v1/templates", params, &resp.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}

// GetSchemas fetches "/atomicassets/v1/template/{collection}/{template_id}" from API
func (c *Client) GetTemplate(collection, template_id string) (TemplateResponse, error) {
	var resp TemplateResponse

	url := fmt.Sprintf("/atomicassets/v1/templates/%s/%s", collection, template_id)
	r, err := c.fetch("GET", url, nil, &resp.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}

func (c *Client) GetTemplateStats(collection, template_id string) (TemplateStatsResponse, error) {
	var resp TemplateStatsResponse

	url := fmt.Sprintf("/atomicassets/v1/templates/%s/%s/stats", collection, template_id)
	r, err := c.fetch("GET", url, nil, &resp.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}
