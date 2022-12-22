package atomicasset

// Basic types

type Token struct {
	Contract  string `json:"token_contract"`
	Symbol    string `json:"token_symbol"`
	Precision int    `json:"token_precision"`
	Amount    string `json:"amount"`
}

type TokenPair struct {
	ListingSymbol    string                 `json:"listing_symbol"`
	SettlementSymbol string                 `json:"settlement_symbol"`
	DelphiPairName   string                 `json:"delphi_pair_name"`
	InvertDelphiPair bool                   `json:"invert_delphi_pair"`
	Data             map[string]interface{} `json:"data"`
}

type Log struct {
	ID             string                 `json:"log_id"`
	TxID           string                 `json:"txid"`
	Name           string                 `json:"name"`
	Data           map[string]interface{} `json:"data"`
	CreatedAtBlock string                 `json:"created_at_block"`
	CreatedAtTime  UnixTime               `json:"created_at_time"`
}

type LogsResponse struct {
	APIResponse
	Data []Log
}

// Offer types

type ListingOffer struct {
	ID                  string `json:"offer_id"`
	Contract            string `json:"contract"`
	Sender              string `json:"sender_name"`
	Recipient           string `json:"recipient_name"`
	Memo                string `json:"memo"`
	State               int64  `json:"state"`
	IsSenderContract    bool   `json:"is_sender_contract"`
	IsRecipientContract bool   `json:"is_recipient_contract"`

	SenderAssets    []ListingAsset `json:"sender_assets"`
	RecipientAssets []ListingAsset `json:"recipient_assets"`

	UpdatedAtBlock string   `json:"updated_at_block"`
	UpdatedAtTime  UnixTime `json:"updated_at_time"`

	CreatedAtBlock string   `json:"created_at_block"`
	CreatedAtTime  UnixTime `json:"created_at_time"`
}

// Transfer types

type ListingTransfer struct {
	ID        string         `json:"transfer_id"`
	Contract  string         `json:"contract"`
	Sender    string         `json:"sender_name"`
	Recipient string         `json:"recipient_name"`
	Memo      string         `json:"memo"`
	Assets    []ListingAsset `json:"assets"`

	CreatedAtBlock string   `json:"created_at_block"`
	CreatedAtTime  UnixTime `json:"created_at_time"`
}

// Price types

type Price struct {
	Average          string     `json:"average"`
	MarketContract   string     `json:"market_contract"`
	Max              string     `json:"max"`
	Median           string     `json:"median"`
	Min              string     `json:"min"`
	Sales            string     `json:"sales"`
	SuggestedAverage string     `json:"suggested_average"`
	SuggestedMedian  string     `json:"suggested_median"`
	Token            PriceToken `json:"token"`
}

type PriceToken struct {
	Contract  string `json:"token_contract"`
	Symbol    string `json:"token_symbol"`
	Precision int    `json:"token_precision"`
}
