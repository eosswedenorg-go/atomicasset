package atomicasset

import (
	"fmt"

	"github.com/eosswedenorg-go/unixtime"
)

// Types

type LinkState int

const (
	LinkStateWaiting  = LinkState(0)
	LinkStateCreated  = LinkState(1)
	LinkStateCanceled = LinkState(2)
	LinkStateClaimed  = LinkState(3)
)

type Link struct {
	ID             string    `json:"link_id"`
	ToolsContract  string    `json:"tools_contract"`
	AssetsContract string    `json:"assets_contract"`
	Creator        string    `json:"creator"`
	Claimer        string    `json:"claimer,omitempty"`
	State          LinkState `json:"state"`
	PublicKey      string    `json:"public_key"`
	Memo           string    `json:"memo"`
	TxID           string    `json:"txid"`
	Assets         []Asset   `json:"assets"`

	CreatedAtBlock string        `json:"created_at_block"`
	CreatedAtTime  unixtime.Time `json:"created_at_time"`
	UpdatedAtBlock string        `json:"updated_at_block"`
	UpdatedAtTime  unixtime.Time `json:"updated_at_time"`
}

// Request Parameters

type LinkRequestParams struct {
	Creator             string             `qs:"creator,omitempty"`
	Claimer             string             `qs:"claimer,omitempty"`
	PublicKey           string             `qs:"public_key,omitempty"`
	State               ReqList[LinkState] `qs:"state,omitempty"`
	CollectionWhitelist ReqList[string]    `qs:"collection_whitelist,omitempty"`
	CollectionBlacklist ReqList[string]    `qs:"collection_blacklist,omitempty"`
	IDs                 ReqList[int]       `qs:"ids,omitempty"`
	LowerBound          string             `qs:"lower_bound,omitempty"`
	UpperBound          string             `qs:"upper_bound,omitempty"`
	Before              int                `qs:"before,omitempty"`
	After               int                `qs:"after,omitempty"`
	Page                int                `qs:"page,omitempty"`
	Limit               int                `qs:"limit,omitempty"`
	Order               SortOrder          `qs:"order,omitempty"`
}

// Responses

type LinkResponse struct {
	APIResponse
	Data Link
}

type LinksResponse struct {
	APIResponse
	Data []Link
}

// API Client functions

// GetLink fetches "/atomictools/v1/links/{id}" from API
func (c *Client) GetLink(id int64) (LinkResponse, error) {
	var resp LinkResponse

	r, err := c.fetch("GET", fmt.Sprintf("/atomictools/v1/links/%d", id), nil, &resp.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}

// GetLinkLogs fetches "/atomiclinks/v1/links/{id}/logs" from API
func (c *Client) GetLinkLogs(id int64, params LogRequestParams) (LogsResponse, error) {
	var resp LogsResponse

	r, err := c.fetch("GET", fmt.Sprintf("/atomictools/v1/links/%d/logs", id), params, &resp.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}

// GetLinks fetches "/atomiclinks/v1/links" from API
func (c *Client) GetLinks(params LinkRequestParams) (LinksResponse, error) {
	var resp LinksResponse

	r, err := c.fetch("GET", "/atomictools/v1/links", params, &resp.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}
