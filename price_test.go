package atomicasset

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/eosswedenorg-go/unixtime"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSalePrices(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicassets/v1/prices/sales", req.URL.String())

		payload := `{
			"success": true,
			"data": [
			  {
				"sale_id": "101853791",
				"auction_id": null,
				"buyoffer_id": null,
				"price": "6000000000",
				"template_mint": 304,
				"token_symbol": "WAX",
				"token_precision": 8,
				"token_contract": "eosio.token",
				"block_time": "1671539853000"
			  },
			  {
				"sale_id": "101718170",
				"auction_id": "1234",
				"buyoffer_id": "1234",
				"price": "65010033",
				"template_mint": 6325,
				"token_symbol": "WAX",
				"token_precision": 8,
				"token_contract": "eosio.token",
				"block_time": "1671539857000"
			  }
			],
			"query_time": 1739985526000
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetSalePrices(PriceSalesRequestParams{})

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2025, time.February, 19, 17, 18, 46, 0, time.UTC), res.QueryTime.Time())

	expected := []PriceSale{
		{
			SaleID:         "101853791",
			AuctionID:      "",
			BuyofferID:     "",
			Price:          "6000000000",
			TemplateMint:   304,
			TokenSymbol:    "WAX",
			TokenPrecision: 8,
			TokenContract:  "eosio.token",
			BlockTime:      unixtime.Time(1671539853000),
		},
		{
			SaleID:         "101718170",
			AuctionID:      "1234",
			BuyofferID:     "1234",
			Price:          "65010033",
			TemplateMint:   6325,
			TokenSymbol:    "WAX",
			TokenPrecision: 8,
			TokenContract:  "eosio.token",
			BlockTime:      unixtime.Time(1671539857000),
		},
	}

	assert.Equal(t, expected, res.Data)
}

func TestSalePricesDays(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicassets/v1/prices/sales/days?template_id=443115", req.URL.String())

		payload := `{
			"success": true,
			"data": [
			  {
				"median": "570000000",
				"average": "575000000",
				"sales": "2",
				"token_symbol": "WAX",
				"token_precision": 8,
				"token_contract": "eosio.token",
				"time": 1670587200000
			  },
			  {
				"median": "600000000",
				"average": "566666667",
				"sales": "3",
				"token_symbol": "WAX",
				"token_precision": 8,
				"token_contract": "eosio.token",
				"time": 1670155200000
			  }
			],
			"query_time": 1761233132000
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetSalePricesDays(PriceSalesRequestParams{TemplateID: 443115})

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2025, time.October, 23, 15, 25, 32, 0, time.UTC), res.QueryTime.Time())

	expected := []PriceSaleDay{
		{
			Median:         "570000000",
			Average:        "575000000",
			Sales:          "2",
			TokenSymbol:    "WAX",
			TokenPrecision: 8,
			TokenContract:  "eosio.token",
			Time:           unixtime.Time(1670587200000),
		},
		{
			Median:         "600000000",
			Average:        "566666667",
			Sales:          "3",
			TokenSymbol:    "WAX",
			TokenPrecision: 8,
			TokenContract:  "eosio.token",
			Time:           unixtime.Time(1670155200000),
		},
	}

	assert.Equal(t, expected, res.Data)
}

func TestGetPriceTemplates(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicassets/v1/prices/templates?limit=2", req.URL.String())

		payload := `{
			"success": true,
			"data": [
			  {
				"market_contract": "atomicmarket",
				"assets_contract": "atomicassets",
				"collection_name": "atomic",
				"template_id": "1",
				"token_symbol": "WAX",
				"token_contract": "eosio.token",
				"token_precision": 8,
				"median": "77700000000",
				"average": "165825790166",
				"min": "500000000",
				"max": "704697986577",
				"sales": "20",
				"suggested_median": "299900000000",
				"suggested_average": "247694911838"
			  },
			  {
				"market_contract": "atomicmarket",
				"assets_contract": "atomicassets",
				"collection_name": "anyo.b1",
				"template_id": "3",
				"token_symbol": "WAX",
				"token_contract": "eosio.token",
				"token_precision": 8,
				"median": "133700000000",
				"average": "133700000000",
				"min": "133700000000",
				"max": "133700000000",
				"sales": "1",
				"suggested_median": "133700000000",
				"suggested_average": "133700000000"
			  }
			],
			"query_time": 1074487326000
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetPriceTemplates(PriceTemplatesRequestParams{Limit: 2})

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2004, time.January, 19, 4, 42, 6, 0, time.UTC), res.QueryTime.Time())

	expected := []PriceTemplate{
		{
			MarketContract:   "atomicmarket",
			AssetsContract:   "atomicassets",
			CollectionName:   "atomic",
			TemplateID:       "1",
			TokenSymbol:      "WAX",
			TokenPrecision:   8,
			TokenContract:    "eosio.token",
			Median:           "77700000000",
			Average:          "165825790166",
			Min:              "500000000",
			Max:              "704697986577",
			Sales:            "20",
			SuggestedMedian:  "299900000000",
			SuggestedAverage: "247694911838",
		},
		{
			MarketContract:   "atomicmarket",
			AssetsContract:   "atomicassets",
			CollectionName:   "anyo.b1",
			TemplateID:       "3",
			TokenSymbol:      "WAX",
			TokenPrecision:   8,
			TokenContract:    "eosio.token",
			Median:           "133700000000",
			Average:          "133700000000",
			Min:              "133700000000",
			Max:              "133700000000",
			Sales:            "1",
			SuggestedMedian:  "133700000000",
			SuggestedAverage: "133700000000",
		},
	}

	assert.Equal(t, expected, res.Data)
}

func TestGetPriceAssets(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicassets/v1/prices/assets?template_id=443110", req.URL.String())

		payload := `{
			"success": true,
			"data": [
			  {
				"token_symbol": "WAX",
				"token_precision": 8,
				"token_contract": "eosio.token",
				"median": "10250500000000",
				"average": "10625664766275",
				"min": "539500000000",
				"max": "26435500000000",
				"suggested_median": "799714936345",
				"suggested_average": "981074594305"
			  }
			],
			"query_time": 1497596433000
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetPriceAssets(PriceAssetsRequestParams{TemplateID: 443110})

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2017, time.June, 16, 7, 0, 33, 0, time.UTC), res.QueryTime.Time())

	expected := []PriceAsset{
		{
			TokenSymbol:      "WAX",
			TokenPrecision:   8,
			TokenContract:    "eosio.token",
			Median:           "10250500000000",
			Average:          "10625664766275",
			Min:              "539500000000",
			Max:              "26435500000000",
			SuggestedMedian:  "799714936345",
			SuggestedAverage: "981074594305",
		},
	}

	assert.Equal(t, expected, res.Data)
}

func TestGetPriceInventory(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicassets/v1/prices/inventory/3nt3a.wam", req.URL.String())

		payload := `{
			"success": true,
			"data": {
				"collections": [
					{
						"collection": {
							"contract": "atomicassets",
							"collection_name": "crptomonkeys",
							"name": "cryptomonKeys",
							"img": "QmUEr6noR9tF5qKPT68VWMARykdwT1vhigm9DzjBRCMSZm",
							"author": "crptomonkeys",
							"allow_notify": true,
							"authorized_accounts": [
								"crptomonkeys",
								"bantano1nine",
								"valorisation",
								"neftyblocksp"
							],
							"notify_accounts": [],
							"market_fee": 0.07,
							"data": {
								"img": "QmUEr6noR9tF5qKPT68VWMARykdwT1vhigm9DzjBRCMSZm",
								"url": "https://www.cryptomonkeys.cc/",
								"name": "cryptomonKeys",
								"socials": "{\"twitter\":\"crypt0monKeys\",\"discord\":\"\",\"medium\":\"banano\",\"youtube\":\"cryptomonKeys\",\"telegram\":\"crypt0monKeys\",\"facebook\":\"\"}",
								"description": "cryptomonKeys is a freely distributed, community-driven, meme-rich digital trading card series based on NFT technology, here to disrupt the meme economy.",
								"creator_info": "{\"address\":\"\",\"company\":\"\",\"name\":\"\",\"registration_number\":\"\"}"
							},
							"created_at_time": "1594768100000",
							"created_at_block": "66677714"
						},
						"prices": [
							{
								"token_symbol": "WAX",
								"token_precision": 8,
								"token_contract": "eosio.token",
								"median": "11160000000",
								"average": "12069398817",
								"min": "5390000000",
								"max": "60499000000",
								"suggested_median": "11790000000",
								"suggested_average": "11747200000"
							}
						]
					}
				]
			},
			"query_time": 1434834852000
		}`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetPriceInventory("3nt3a.wam", PriceInventoryRequestParams{})

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2015, time.June, 20, 21, 14, 12, 0, time.UTC), res.QueryTime.Time())

	expected := []PriceInventory{
		{
			Collection: Collection{
				Contract:       "atomicassets",
				CollectionName: "crptomonkeys",
				Name:           "cryptomonKeys",
				Image:          "QmUEr6noR9tF5qKPT68VWMARykdwT1vhigm9DzjBRCMSZm",
				Author:         "crptomonkeys",
				AllowNotify:    true,
				AuthorizedAccounts: []string{
					"crptomonkeys",
					"bantano1nine",
					"valorisation",
					"neftyblocksp",
				},
				NotifyAccounts: []string{},
				MarketFee:      0.07,
				Data: map[string]interface{}{
					"img":          "QmUEr6noR9tF5qKPT68VWMARykdwT1vhigm9DzjBRCMSZm",
					"url":          "https://www.cryptomonkeys.cc/",
					"name":         "cryptomonKeys",
					"socials":      "{\"twitter\":\"crypt0monKeys\",\"discord\":\"\",\"medium\":\"banano\",\"youtube\":\"cryptomonKeys\",\"telegram\":\"crypt0monKeys\",\"facebook\":\"\"}",
					"description":  "cryptomonKeys is a freely distributed, community-driven, meme-rich digital trading card series based on NFT technology, here to disrupt the meme economy.",
					"creator_info": "{\"address\":\"\",\"company\":\"\",\"name\":\"\",\"registration_number\":\"\"}",
				},
				CreatedAtBlock: "66677714",
				CreatedAtTime:  unixtime.Time(1594768100000),
			},
			Prices: []PriceAsset{
				{
					TokenSymbol:      "WAX",
					TokenPrecision:   8,
					TokenContract:    "eosio.token",
					Median:           "11160000000",
					Average:          "12069398817",
					Min:              "5390000000",
					Max:              "60499000000",
					SuggestedMedian:  "11790000000",
					SuggestedAverage: "11747200000",
				},
			},
		},
	}

	assert.Equal(t, expected, res.Data)
}
