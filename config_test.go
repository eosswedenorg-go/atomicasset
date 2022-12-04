package atomicasset

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAssetsConfig(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicassets/v1/config", req.URL.String())

		payload := `{
			"success": true,
			"data": {
			  "contract": "atomicassets",
			  "version": "1.2.3",
			  "collection_format": [
				{
				  "name": "name",
				  "type": "string"
				},
				{
				  "name": "img",
				  "type": "ipfs"
				},
				{
				  "name": "description",
				  "type": "string"
				},
				{
				  "name": "url",
				  "type": "string"
				},
				{
				  "name": "images",
				  "type": "string"
				},
				{
				  "name": "socials",
				  "type": "string"
				},
				{
				  "name": "creator_info",
				  "type": "string"
				}
			  ],
			  "supported_tokens": [
				{
				  "token_symbol": "WAX",
				  "token_contract": "eosio.token",
				  "token_precision": 8
				},
				{
				  "token_symbol": "PGL",
				  "token_contract": "prospectorsw",
				  "token_precision": 4
				}
			  ]
			},
			"query_time": 1773123095000
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetAssetsConfig()
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2026, time.March, 10, 6, 11, 35, 0, time.UTC), res.QueryTime.Time())

	expected := AssetsConfig{
		Contract: "atomicassets",
		Version:  "1.2.3",
		CollectionFormat: []SchemaFormat{
			{
				Name: "name",
				Type: "string",
			},
			{
				Name: "img",
				Type: "ipfs",
			},
			{
				Name: "description",
				Type: "string",
			},
			{
				Name: "url",
				Type: "string",
			},
			{
				Name: "images",
				Type: "string",
			},
			{
				Name: "socials",
				Type: "string",
			},
			{
				Name: "creator_info",
				Type: "string",
			},
		},
		SupportedTokens: []PriceToken{
			{
				Symbol:    "WAX",
				Contract:  "eosio.token",
				Precision: 8,
			},
			{
				Symbol:    "PGL",
				Contract:  "prospectorsw",
				Precision: 4,
			},
		},
	}

	assert.Equal(t, expected, res.Data)
}

func TestGetMarketConfig(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicmarket/v1/config", req.URL.String())

		payload := `{
			"success": true,
			"data": {
			  "atomicassets_contract": "atomicassets",
			  "atomicmarket_contract": "atomicmarket",
			  "delphioracle_contract": "delphioracle",
			  "version": "1.3.3",
			  "maker_market_fee": 0.01,
			  "taker_market_fee": 0.01,
			  "minimum_auction_duration": 120,
			  "maximum_auction_duration": 2592000,
			  "minimum_bid_increase": 0.1,
			  "auction_reset_duration": 120,
			  "supported_tokens": [
				{
				  "token_contract": "eosio.token",
				  "token_symbol": "WAX",
				  "token_precision": 8
				}
			  ],
			  "supported_pairs": [
				{
				  "listing_symbol": "USD",
				  "settlement_symbol": "WAX",
				  "delphi_pair_name": "waxpusd",
				  "invert_delphi_pair": false,
				  "data": {
					"median": 595,
					"contract": "delphioracle",
					"base_symbol": "WAXP",
					"quote_symbol": "USD",
					"base_precision": 8,
					"quote_precision": 2,
					"updated_at_time": 1670108470500,
					"delphi_pair_name": "waxpusd",
					"median_precision": 4,
					"updated_at_block": 217287489
				  }
				}
			  ]
			},
			"query_time": 986953317000
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetMarketConfig()

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2001, time.April, 11, 1, 41, 57, 0, time.UTC), res.QueryTime.Time())

	expected := MarketConfig{
		AtomicassetsContract:   "atomicassets",
		AtomicmarketContract:   "atomicmarket",
		DelphioracleContract:   "delphioracle",
		Version:                "1.3.3",
		MakerMarketFee:         0.01,
		TakerMarketFee:         0.01,
		MinimumAuctionDuration: 120,
		MaximumAuctionDuration: 2592000,
		MinimumBidIncrease:     0.1,
		AuctionResetDuration:   120,
		SupportedTokens: []PriceToken{
			{
				Contract:  "eosio.token",
				Symbol:    "WAX",
				Precision: 8,
			},
		},
		SupportedPairs: []TokenPair{
			{
				ListingSymbol:    "USD",
				SettlementSymbol: "WAX",
				DelphiPairName:   "waxpusd",
				InvertDelphiPair: false,
				Data: map[string]interface{}{
					"median":           float64(595),
					"contract":         "delphioracle",
					"base_symbol":      "WAXP",
					"quote_symbol":     "USD",
					"base_precision":   float64(8),
					"quote_precision":  float64(2),
					"updated_at_time":  float64(1670108470500),
					"delphi_pair_name": "waxpusd",
					"median_precision": float64(4),
					"updated_at_block": float64(217287489),
				},
			},
		},
	}

	assert.Equal(t, expected, res.Data)
}

func TestGetToolsConfig(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomictools/v1/config", req.URL.String())

		payload := `{
			"success": true,
			"data": {
			  "atomictools_contract": "atomictoolsx",
			  "atomicassets_contract": "atomicassets",
			  "version": "1.0.0"
			},
			"query_time": 1669034048000
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetToolsConfig()

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2022, time.November, 21, 12, 34, 8, 0, time.UTC), res.QueryTime.Time())

	expected := ToolsConfig{
		AtomictoolsContract:  "atomictoolsx",
		AtomicassetsContract: "atomicassets",
		Version:              "1.0.0",
	}

	assert.Equal(t, expected, res.Data)
}
