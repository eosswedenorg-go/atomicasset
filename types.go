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

// Logs

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
