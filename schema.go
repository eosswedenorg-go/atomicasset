package atomicasset

import (
	"github.com/eosswedenorg-go/unixtime"
)

// Types

type Schema struct {
	Name           string         `json:"schema_name"`
	Contract       string         `json:"contract"`
	Format         []SchemaFormat `json:"format"`
	Collection     Collection     `json:"collection"`
	CreatedAtBlock string         `json:"created_at_block"`
	CreatedAtTime  unixtime.Time  `json:"created_at_time"`
}

type InlineSchema struct {
	Name           string         `json:"schema_name"`
	Format         []SchemaFormat `json:"format"`
	CreatedAtBlock string         `json:"created_at_block"`
	CreatedAtTime  unixtime.Time  `json:"created_at_time"`
}

type SchemaFormat struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// Request Parameters

type SchemaSortColumn string

const (

	// SchemaSortDefault sorts by the default column (created)
	SchemaSortDefault SchemaSortColumn = ""

	// SchemaSortCreated sorts by the created column
	SchemaSortCreated SchemaSortColumn = "created"

	// SchemaSortAssets sorts by the assets column
	SchemaSortAssets SchemaSortColumn = "assets"

	// SchemaSortName sorts by the schema name column
	SchemaSortName SchemaSortColumn = "schema_name"
)

type SchemasRequestParams struct {
	Author            string           `qs:"author,omitempty"`
	Match             string           `qs:"match,omitempty"`
	AuthorizedAccount string           `qs:"authorized_account,omitempty"`
	NotifyAccount     string           `qs:"notify_account,omitempty"`
	Blacklist         ReqList[string]  `qs:"collection_blacklist,omitempty"`
	Whitelist         ReqList[string]  `qs:"collection_whitelist,omitempty"`
	IDs               ReqList[int]     `qs:"ids,omitempty"`
	LowerBound        string           `qs:"lower_bound,omitempty"`
	UpperBound        string           `qs:"upper_bound,omitempty"`
	Before            int              `qs:"before,omitempty"`
	After             int              `qs:"after,omitempty"`
	Page              int              `qs:"page,omitempty"`
	Limit             int              `qs:"limit,omitempty"`
	Order             SortOrder        `qs:"order,omitempty"`
	Sort              SchemaSortColumn `qs:"sort,omitempty"`
}

// Responses

type SchemasResponse struct {
	APIResponse
	Data []Schema
}

// Client API Functions

// GetSchemas fetches "/atomicassets/v1/schemas" from API
func (c *Client) GetSchemas(params SchemasRequestParams) (SchemasResponse, error) {
	var resp SchemasResponse

	r, err := c.fetch("GET", "/atomicassets/v1/schemas", params, &resp.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}
