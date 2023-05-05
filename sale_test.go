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

var matrix_funko_col = Collection{
	Name:               "The Matrix x Funko",
	CollectionName:     "matrix.funko",
	Image:              "QmZ3gL78m7c9LoEbKpVpKcf4m8UJUXKEpZizQ1P1SWehRs",
	Author:             "matrix.funko",
	AllowNotify:        true,
	AuthorizedAccounts: []string{"matrix.funko"},
	NotifyAccounts:     []string{},
	MarketFee:          0.06,
	CreatedAtBlock:     "213350477",
	CreatedAtTime:      unixtime.Time(1668138328000),
}

var morpheus_data = map[string]interface{}{
	"tid":                        float64(74),
	"name":                       "Morpheus",
	"legal":                      "THE MATRIX and all related characters and elements © & ™ Warner Bros.  Entertainment Inc. (s22)",
	"video":                      "QmPXbmfk6DPTgd4hgtAAp797zpqqUJaGuCfJpxANtTMZXU/Assets/THE-MATRIX_UNCOMMON_MORPHEUS-DOJO-SUIT_BINARY_STATIC.mp4",
	"cardid":                     float64(40),
	"rarity":                     "Uncommon",
	"backimg":                    "QmPXbmfk6DPTgd4hgtAAp797zpqqUJaGuCfJpxANtTMZXU/Assets/back.png",
	"variant":                    "Binary Static",
	"description":                "Morpheus Digital Pop!",
	"release date":               "November 15,  2022",
	"end user license agreement": "https://digital.funko.com/eula/wb ",
}

var sale_100420714_json = `{
	  "market_contract": "atomicmarket",
	  "assets_contract": "atomicassets",
	  "sale_id": "100420714",
	  "seller": "zhl3g.c.wam",
	  "buyer": null,
	  "offer_id": "104058114",
	  "price": {
		"token_contract": "eosio.token",
		"token_symbol": "WAX",
		"token_precision": 8,
		"median": null,
		"amount": "50000000"
	  },
	  "listing_price": "50000000",
	  "listing_symbol": "WAX",
	  "assets": [
		{
		  "contract": "atomicassets",
		  "asset_id": "1099835349411",
		  "owner": "zhl3g.c.wam",
		  "is_transferable": true,
		  "is_burnable": true,
		  "collection": {
			"collection_name": "matrix.funko",
			"name": "The Matrix x Funko",
			"img": "QmZ3gL78m7c9LoEbKpVpKcf4m8UJUXKEpZizQ1P1SWehRs",
			"author": "matrix.funko",
			"allow_notify": true,
			"authorized_accounts": [
			  "matrix.funko"
			],
			"notify_accounts": [],
			"market_fee": 0.06,
			"created_at_block": "213350477",
			"created_at_time": "1668138328000"
		  },
		  "schema": {
			"schema_name": "series1.drop",
			"format": [
			  {
				"name": "name",
				"type": "string"
			  },
			  {
				"name": "rarity",
				"type": "string"
			  },
			  {
				"name": "variant",
				"type": "string"
			  },
			  {
				"name": "cardid",
				"type": "uint8"
			  },
			  {
				"name": "legal",
				"type": "string"
			  },
			  {
				"name": "end user license agreement",
				"type": "string"
			  },
			  {
				"name": "video",
				"type": "image"
			  },
			  {
				"name": "backimg",
				"type": "image"
			  },
			  {
				"name": "tid",
				"type": "uint16"
			  },
			  {
				"name": "release date",
				"type": "string"
			  },
			  {
				"name": "description",
				"type": "string"
			  },
			  {
				"name": "img",
				"type": "image"
			  }
			],
			"created_at_block": "213351916",
			"created_at_time": "1668139047500"
		  },
		  "template": {
			"template_id": "620698",
			"max_supply": "4400",
			"is_transferable": true,
			"is_burnable": true,
			"issued_supply": "4400",
			"immutable_data": {
			  "tid": 74,
			  "name": "Morpheus",
			  "legal": "THE MATRIX and all related characters and elements © & ™ Warner Bros.  Entertainment Inc. (s22)",
			  "video": "QmPXbmfk6DPTgd4hgtAAp797zpqqUJaGuCfJpxANtTMZXU/Assets/THE-MATRIX_UNCOMMON_MORPHEUS-DOJO-SUIT_BINARY_STATIC.mp4",
			  "cardid": 40,
			  "rarity": "Uncommon",
			  "backimg": "QmPXbmfk6DPTgd4hgtAAp797zpqqUJaGuCfJpxANtTMZXU/Assets/back.png",
			  "variant": "Binary Static",
			  "description": "Morpheus Digital Pop!",
			  "release date": "November 15,  2022",
			  "end user license agreement": "https://digital.funko.com/eula/wb "
			},
			"created_at_time": "1668139238500",
			"created_at_block": "213352298"
		  },
		  "mutable_data": {},
		  "immutable_data": {},
		  "template_mint": "3032",
		  "backed_tokens": [],
		  "burned_by_account": null,
		  "burned_at_block": null,
		  "burned_at_time": null,
		  "updated_at_block": "217566006",
		  "updated_at_time": "1670247881000",
		  "transferred_at_block": "217566006",
		  "transferred_at_time": "1670247881000",
		  "minted_at_block": "213359433",
		  "minted_at_time": "1668142807000",
		  "sales": [],
		  "auctions": [],
		  "prices": [
			{
			  "market_contract": "atomicmarket",
			  "token": {
				"token_symbol": "WAX",
				"token_precision": 8,
				"token_contract": "eosio.token"
			  },
			  "median": "60000000",
			  "average": "106240976",
			  "suggested_median": "58212000",
			  "suggested_average": "62726400",
			  "min": "30000000",
			  "max": "2500000000",
			  "sales": "713"
			}
		  ],
		  "data": {
			"tid": 74,
			"name": "Morpheus",
			"legal": "THE MATRIX and all related characters and elements © & ™ Warner Bros.  Entertainment Inc. (s22)",
			"video": "QmPXbmfk6DPTgd4hgtAAp797zpqqUJaGuCfJpxANtTMZXU/Assets/THE-MATRIX_UNCOMMON_MORPHEUS-DOJO-SUIT_BINARY_STATIC.mp4",
			"cardid": 40,
			"rarity": "Uncommon",
			"backimg": "QmPXbmfk6DPTgd4hgtAAp797zpqqUJaGuCfJpxANtTMZXU/Assets/back.png",
			"variant": "Binary Static",
			"description": "Morpheus Digital Pop!",
			"release date": "November 15,  2022",
			"end user license agreement": "https://digital.funko.com/eula/wb "
		  },
		  "name": "Morpheus"
		}
	  ],
	  "maker_marketplace": "nft.hive",
	  "taker_marketplace": null,
	  "collection": {
		"collection_name": "matrix.funko",
		"name": "The Matrix x Funko",
		"img": "QmZ3gL78m7c9LoEbKpVpKcf4m8UJUXKEpZizQ1P1SWehRs",
		"author": "matrix.funko",
		"allow_notify": true,
		"authorized_accounts": [
		  "matrix.funko"
		],
		"notify_accounts": [],
		"market_fee": 0.06,
		"created_at_block": "213350477",
		"created_at_time": "1668138328000"
	  },
	  "is_seller_contract": false,
	  "updated_at_block": "217566088",
	  "updated_at_time": "1670247922000",
	  "created_at_block": "217566088",
	  "created_at_time": "1670247922000",
	  "ordinality": "1",
	  "state": 1
	}`

var sale_100420714 = Sale{
	ID:              "100420714",
	AsssetsContract: "atomicassets",
	MarketContract:  "atomicmarket",
	Seller:          "zhl3g.c.wam",
	Buyer:           "",
	OfferID:         "104058114",
	Price: Token{
		Contract:  "eosio.token",
		Symbol:    "WAX",
		Precision: 8,
		Amount:    "50000000",
	},
	ListingPrice:  "50000000",
	ListingSymbol: "WAX",
	Assets: []Asset{
		{
			ID:             "1099835349411",
			Name:           "Morpheus",
			Contract:       "atomicassets",
			Owner:          "zhl3g.c.wam",
			IsTransferable: true,
			IsBurnable:     true,
			Collection:     matrix_funko_col,
			Schema: InlineSchema{
				Name: "series1.drop",
				Format: []SchemaFormat{
					{Name: "name", Type: "string"},
					{Name: "rarity", Type: "string"},
					{Name: "variant", Type: "string"},
					{Name: "cardid", Type: "uint8"},
					{Name: "legal", Type: "string"},
					{Name: "end user license agreement", Type: "string"},
					{Name: "video", Type: "image"},
					{Name: "backimg", Type: "image"},
					{Name: "tid", Type: "uint16"},
					{Name: "release date", Type: "string"},
					{Name: "description", Type: "string"},
					{Name: "img", Type: "image"},
				},
				CreatedAtBlock: "213351916",
				CreatedAtTime:  unixtime.Time(1668139047500),
			},
			Template: Template{
				ID:             "620698",
				MaxSupply:      "4400",
				IsTransferable: true,
				IsBurnable:     true,
				IssuedSupply:   "4400",
				ImmutableData:  morpheus_data,
				CreatedAtBlock: "213352298",
				CreatedAtTime:  unixtime.Time(1668139238500),
			},
			MutableData:       map[string]interface{}{},
			ImmutableData:     map[string]interface{}{},
			TemplateMint:      "3032",
			BackedTokens:      []Token{},
			BurnedByAccount:   "",
			BurnedAtBlock:     "",
			BurnedAtTime:      unixtime.Time(0),
			UpdatedAtBlock:    "217566006",
			UpdatedAtTime:     unixtime.Time(1670247881000),
			TransferedAtBlock: "217566006",
			TransferedAtTime:  unixtime.Time(1670247881000),
			MintedAtBlock:     "213359433",
			MintedAtTime:      unixtime.Time(1668142807000),
			Data:              morpheus_data,
		},
	},
	MakerMarketplace: "nft.hive",
	TakerMarketplace: "",
	Collection:       matrix_funko_col,
	IsSellerContract: false,
	State:            SalesStateListed,
	UpdatedAtBlock:   "217566088",
	UpdatedAtTime:    unixtime.Time(1670247922000),
	CreatedAtBlock:   "217566088",
	CreatedAtTime:    unixtime.Time(1670247922000),
}

func TestGetSales(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicmarket/v2/sales?limit=1", req.URL.String())

		payload := `{
			"success": true,
			"data": [
			  ` + sale_100420714_json + `
			],
			"query_time": 1298808950000
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetSales(SalesRequestParams{Limit: 1})

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2011, time.February, 27, 12, 15, 50, 0, time.UTC), res.QueryTime.Time())

	assert.Equal(t, []Sale{sale_100420714}, res.Data)
}

func TestGetSalesGroupByTemplate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicmarket/v1/sales/templates?limit=1&symbol=WAX", req.URL.String())

		payload := `{
			"success": true,
			"data": [
			  ` + sale_100420714_json + `
			],
			"query_time": "1516192861000"
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetSalesGroupByTemplate(SalesTemplateRequestParams{Symbol: "WAX", Limit: 1})

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2018, time.January, 17, 12, 41, 1, 0, time.UTC), res.QueryTime.Time())

	assert.Equal(t, []Sale{sale_100420714}, res.Data)
}

func TestGetSale(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicmarket/v1/sales/100420714", req.URL.String())

		payload := `{
			"success": true,
			"data":` + sale_100420714_json + `,
			"query_time": 1298808950000
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetSale(100420714)

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2011, time.February, 27, 12, 15, 50, 0, time.UTC), res.QueryTime.Time())

	assert.Equal(t, sale_100420714, res.Data)
}

func TestGetSaleLogs(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicmarket/v1/sales/100420714/logs?limit=2&order=desc&page=1", req.URL.String())

		payload := `{
			"success": true,
			"data": [
			  {
				"log_id": "70817779067",
				"name": "purchasesale",
				"data": {
				  "taker_marketplace": "",
				  "intended_delphi_median": "0"
				},
				"txid": "3650f8425f0dd7a880a244c77b2c2ad201760648ba77859e26f8bc7bf7d59efc",
				"created_at_block": "217573294",
				"created_at_time": "1670251531000"
			  },
			  {
				"log_id": "70813258501",
				"name": "logsalestart",
				"data": {
				  "offer_id": "104058114"
				},
				"txid": "1ae2c6b0f1ad14c4bf0dcac2175ddded71e773bf5abc3342351034227cafd59c",
				"created_at_block": "217566088",
				"created_at_time": "1670247922000"
			  }
			],
			"query_time": 1634527029000
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetSaleLogs(100420714, LogRequestParams{Page: 1, Limit: 2, Order: SortDescending})

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2021, time.October, 18, 3, 17, 9, 0, time.UTC), res.QueryTime.Time())

	expected := []Log{
		{
			ID:   "70817779067",
			Name: "purchasesale",
			Data: map[string]interface{}{
				"taker_marketplace":      "",
				"intended_delphi_median": "0",
			},
			TxID:           "3650f8425f0dd7a880a244c77b2c2ad201760648ba77859e26f8bc7bf7d59efc",
			CreatedAtBlock: "217573294",
			CreatedAtTime:  unixtime.Time(1670251531000),
		},
		{
			ID:   "70813258501",
			Name: "logsalestart",
			Data: map[string]interface{}{
				"offer_id": "104058114",
			},
			TxID:           "1ae2c6b0f1ad14c4bf0dcac2175ddded71e773bf5abc3342351034227cafd59c",
			CreatedAtBlock: "217566088",
			CreatedAtTime:  unixtime.Time(1670247922000),
		},
	}

	assert.Equal(t, expected, res.Data)
}
