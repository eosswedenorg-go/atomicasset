package atomicasset

// Types

type Offer struct {
	ID                  string  `json:"offer_id"`
	Contract            string  `json:"contract"`
	Sender              string  `json:"sender_name"`
	Recipient           string  `json:"recipient_name"`
	Memo                string  `json:"memo"`
	State               int64   `json:"state"`
	IsSenderContract    bool    `json:"is_sender_contract"`
	IsRecipientContract bool    `json:"is_recipient_contract"`
	SenderAssets        []Asset `json:"sender_assets"`
	RecipientAssets     []Asset `json:"recipient_assets"`

	UpdatedAtBlock string   `json:"updated_at_block"`
	UpdatedAtTime  UnixTime `json:"updated_at_time"`

	CreatedAtBlock string   `json:"created_at_block"`
	CreatedAtTime  UnixTime `json:"created_at_time"`
}

// Request Parameters

type OfferState string

const (
	OfferStatePending  = OfferState("0")
	OfferStateInvalid  = OfferState("1")
	OfferStateUnknown  = OfferState("2")
	OfferStateAccepted = OfferState("3")
	OfferStateDeclined = OfferState("4")
	OfferStateCanceled = OfferState("5")
)

type OfferSortColumn string

const (
	OfferSortCreated = OfferSortColumn("created")
	OfferSortUpdated = OfferSortColumn("updated")
)

type OfferRequestParams struct {
	Account                 string          `qs:"account,omitempty"`
	Sender                  ReqList[string] `qs:"sender,omitempty"`
	Recipient               ReqList[string] `qs:"recipient,omitempty"`
	Memo                    string          `qs:"memo,omitempty"`
	MatchMemo               string          `qs:"match_memo,omitempty"`
	State                   OfferState      `qs:"state,omitempty"`
	IsRecipientContract     bool            `qs:"is_recipient_contract,omitempty"`
	AssetID                 ReqList[int]    `qs:"asset_id,omitempty"`
	TemplateID              ReqList[int]    `qs:"template_id,omitempty"`
	SchemaName              ReqList[string] `qs:"schema_name,omitempty"`
	CollectionName          ReqList[string] `qs:"collection_name,omitempty"`
	AccountWhitelist        ReqList[string] `qs:"account_whitelist,omitempty"`
	AccountBlacklist        ReqList[string] `qs:"account_blacklist,omitempty"`
	SenderAssetWhitelist    ReqList[string] `qs:"sender_asset_whitelist,omitempty"`
	SenderAssetBlacklist    ReqList[string] `qs:"sender_asset_blacklist,omitempty"`
	RecipientAssetWhitelist ReqList[string] `qs:"recipient_asset_whitelist,omitempty"`
	RecipientAssetBlacklist ReqList[string] `qs:"recipient_asset_blacklist,omitempty"`
	CollectionWhitelist     ReqList[string] `qs:"collection_whitelist,omitempty"`
	CollectionBlacklist     ReqList[string] `qs:"collection_blacklist,omitempty"`
	HideContracts           bool            `qs:"hide_contracts,omitempty"`
	HideEmptyOffers         bool            `qs:"hide_empty_offers,omitempty"`
	IDs                     ReqList[int]    `qs:"ids,omitempty"`
	LowerBound              string          `qs:"lower_bound,omitempty"`
	UpperBound              string          `qs:"upper_bound,omitempty"`
	Before                  int             `qs:"before,omitempty"`
	After                   int             `qs:"after,omitempty"`
	Page                    int             `qs:"page,omitempty"`
	Limit                   int             `qs:"limit,omitempty"`
	Order                   SortOrder       `qs:"order,omitempty"`
	Sort                    OfferSortColumn `qs:"sort,omitempty"`
}

// Responses

type OfferResponse struct {
	APIResponse
	Data Offer
}

type OffersResponse struct {
	APIResponse
	Data []Offer
}

type OfferLogResponse struct {
	APIResponse
	Data []Log
}

// API Client functions

// GetOffers fetches "/atomicassets/v1/offers" from API
func (c *Client) GetOffers(params OfferRequestParams) (OffersResponse, error) {
	var offers OffersResponse

	r, err := c.fetch("GET", "/atomicassets/v1/offers", params, &offers.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&offers)
	}
	return offers, err
}

// GetOffer fetches "/atomicassets/v1/offers/{offers_id}" from API
func (c *Client) GetOffer(offerID string) (OfferResponse, error) {
	var offer OfferResponse

	r, err := c.fetch("GET", "/atomicassets/v1/offers/"+offerID, nil, &offer.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&offer)
	}
	return offer, err
}

// GetOfferLog fetches "/atomicassets/v1/offers/{offers_id}/logs" from API
func (c *Client) GetOfferLog(offerID string, params LogRequestParams) (OfferLogResponse, error) {
	var logs OfferLogResponse

	r, err := c.fetch("GET", "/atomicassets/v1/offers/"+offerID+"/logs", params, &logs.APIResponse)
	if err == nil {
		// Parse json
		err = r.Unmarshal(&logs)
	}
	return logs, err
}
