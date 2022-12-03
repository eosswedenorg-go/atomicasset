package atomicasset

// Types

type Transfer struct {
	ID        string  `json:"transfer_id"`
	Contract  string  `json:"contract"`
	Sender    string  `json:"sender_name"`
	Recipient string  `json:"recipient_name"`
	Memo      string  `json:"memo"`
	Assets    []Asset `json:"assets"`

	CreatedAtBlock string   `json:"created_at_block"`
	CreatedAtTime  UnixTime `json:"created_at_time"`
}

// Request Parameters

type TransferRequestParams struct {
	Account             ReqStringList `qs:"account,omitempty"`
	Sender              ReqStringList `qs:"sender,omitempty"`
	Recipient           ReqStringList `qs:"recipient,omitempty"`
	Memo                string        `qs:"memo,omitempty"`
	MatchMemo           string        `qs:"match_memo,omitempty"`
	AssetID             ReqIntList    `qs:"asset_id,omitempty"`
	TemplateID          ReqIntList    `qs:"template_id,omitempty"`
	SchemaName          ReqStringList `qs:"schema_name,omitempty"`
	CollectionName      ReqStringList `qs:"collection_name,omitempty"`
	CollectionWhitelist ReqStringList `qs:"collection_whitelist,omitempty"`
	CollectionBlacklist ReqStringList `qs:"collection_blacklist,omitempty"`
	HideContracts       bool          `qs:"hide_contracts,omitempty"`
	IDs                 ReqStringList `qs:"ids,omitempty"`
	LowerBound          string        `qs:"lower_bound,omitempty"`
	UpperBound          string        `qs:"upper_bound,omitempty"`
	Before              int           `qs:"before,omitempty"`
	After               int           `qs:"after,omitempty"`
	Page                int           `qs:"page,omitempty"`
	Limit               int           `qs:"limit,omitempty"`
	Order               SortOrder     `qs:"order,omitempty"`

	// Sort parameter exists but only has one value "created" that also is the default.
	// So skip that for now until more values are added.
}

// Responses

type TransfersResponse struct {
	APIResponse
	Data []Transfer
}

// Client API Functions

// GetTransfers fetches "/atomicassets/v1/transfers" from API
func (c *Client) GetTransfers(params TransferRequestParams) (TransfersResponse, error) {
	var resp TransfersResponse

	r, err := c.fetch("GET", "/atomicassets/v1/transfers", params, &resp.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&resp)
	}
	return resp, err
}
