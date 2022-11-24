package atomicasset

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var asset1 = Asset{
	ID:             "1099667509880",
	Contract:       "atomicassets",
	Owner:          "farmersworld",
	Name:           "Silver Member",
	IsTransferable: true,
	IsBurnable:     true,
	TemplateMint:   "4433",
	Collection: Collection{
		CollectionName: "farmersworld",
		Name:           "Farmers World",
		Author:         ".jieg.wam",
		AllowNotify:    true,
		AuthorizedAccounts: []string{
			".jieg.wam",
			"farmersworld",
			"atomicdropsx",
			"atomicpacksx",
			"neftyblocksd",
		},
		NotifyAccounts: []string{
			"atomicdropsx",
		},
		MarketFee:      0.05,
		CreatedAtBlock: "123762633",
		CreatedAtTime:  UnixTime(1623323058000),
	},
	Schema: InlineSchema{
		Name: "memberships",
		Format: []SchemaFormat{
			{
				Name: "name",
				Type: "string",
			},
			{
				Name: "img",
				Type: "image",
			},
			{
				Name: "description",
				Type: "string",
			},
			{
				Name: "type",
				Type: "string",
			},
			{
				Name: "rarity",
				Type: "string",
			},
			{
				Name: "level",
				Type: "uint8",
			},
		},
		CreatedAtBlock: "136880914",
		CreatedAtTime:  UnixTime(1629887699000),
	},
	Template: Template{
		ID:             "260629",
		MaxSupply:      "0",
		IsTransferable: true,
		IsBurnable:     true,
		IssuedSupply:   "112195",
		ImmutableData: map[string]interface{}{
			"img":         "QmZWg1mP2UNcSwhrYNVqjk16BnhcWCz3oAva8BfiTNB3J4",
			"name":        "Silver Member",
			"type":        "Wood",
			"level":       float64(2),
			"rarity":      "Uncommon",
			"description": "This is a member card powered by Wood. When used by the farmer, it will increase the power and luck of the wood mining tools, and can mine the Farmer Coin that has been lost since ancient times.",
		},
		CreatedAtBlock: "136882467",
		CreatedAtTime:  UnixTime(1629888476000),
	},
	ImmutableData: map[string]interface{}{
		"asdx": "4321",
	},
	MutableData: map[string]interface{}{
		"asdf": "1234",
	},
	UpdatedAtBlock:    "171080009",
	UpdatedAtTime:     "1646996870500",
	TransferedAtBlock: "171080009",
	TransferedAtTime:  "1646996870500",
	MintedAtBlock:     "171080009",
	MintedAtTime:      "1646996870500",
	BackedTokens:      []Token{},
}

func TestClient_SendError(t *testing.T) {
	client := New("http://0.0.0.0:8080")

	_, err := client.send("GET", "/", nil)

	assert.EqualError(t, err, "Get \"http://0.0.0.0:8080/\": dial tcp 0.0.0.0:8080: connect: connection refused")
}

func TestClient_SendEncodeParametersFail(t *testing.T) {
	client := Client{}

	_, err := client.send("GET", "/", "a string")

	assert.EqualError(t, err, "expects struct input, got string")
}

func TestClient_GetHealth(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.URL.String() == "/health" {
			payload := `{
                "success":true,
                "data":{
                    "version":"1.0.0",
                    "postgres":{
                        "status":"OK",
                        "readers":[
                            {
                                "block_num":"167836036"
                            },
                            {
                                "block_num":"167836034"
                            }
                        ]
                    },
                    "redis":{
                        "status":"OK"
                    },
                    "chain":{
                        "status":"OK",
                        "head_block":167836035,
                        "head_time":1645374771500
                    }
                },
                "query_time":1645374772067
            }`

			res.Header().Add("Content-type", "application/json; charset=utf-8")
			_, err := res.Write([]byte(payload))
			assert.NoError(t, err)
		}
	}))

	client := New(srv.URL)

	h, err := client.GetHealth()

	require.NoError(t, err)
	assert.Equal(t, 200, h.HTTPStatusCode)

	assert.True(t, h.Success)
	assert.Equal(t, time.Time(time.Date(2022, time.February, 20, 16, 32, 52, 67, time.UTC)), h.QueryTime.Time())

	// Data
	assert.Equal(t, "1.0.0", h.Data.Version)

	// Postgres
	assert.Equal(t, "OK", h.Data.Postgres.Status)

	// Redis
	assert.Equal(t, "OK", h.Data.Redis.Status)

	// Chain
	assert.Equal(t, "OK", h.Data.Chain.Status)
	assert.Equal(t, int64(167836035), h.Data.Chain.HeadBlock)
	assert.Equal(t, time.Time(time.Date(2022, time.February, 20, 16, 32, 51, 500, time.UTC)), h.Data.Chain.HeadTime.Time())
}

func TestClient_GetHealthFailed(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.URL.String() == "/health" {
			payload := `{
                "success":true,
                "data":{
                    "version":"1.0.0",
                    "postgres":{
                        "status":"ERROR",
                        "readers":[]
                    },
                    "redis":{
                        "status":"ERROR"
                    },
                    "chain":{
                        "status":"ERROR",
                        "head_block":0,
                        "head_time":0
                    }
                },
                "query_time":1645374772067
            }`

			res.Header().Add("Content-type", "application/json")
			_, err := res.Write([]byte(payload))
			assert.NoError(t, err)
		}
	}))

	client := New(srv.URL)

	h, err := client.GetHealth()

	require.NoError(t, err)
	assert.Equal(t, 200, h.HTTPStatusCode)

	assert.True(t, h.Success)
	assert.Equal(t, time.Time(time.Date(2022, time.February, 20, 16, 32, 52, 67, time.UTC)), h.QueryTime.Time())

	// Data
	assert.Equal(t, "1.0.0", h.Data.Version)

	// Postgres
	assert.Equal(t, "ERROR", h.Data.Postgres.Status)

	// Redis
	assert.Equal(t, "ERROR", h.Data.Redis.Status)

	// Chain
	assert.Equal(t, "ERROR", h.Data.Chain.Status)
	assert.Equal(t, int64(0), h.Data.Chain.HeadBlock)

	assert.Equal(t, time.Unix(0, 0).UTC(), h.Data.Chain.HeadTime.Time())
}

func TestClient_APIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		payload := `{
          "success": false,
          "message": "Some internal error"
        }`

		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(500)
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	_, err := client.GetHealth()

	assert.EqualError(t, err, "API Error: Some internal error")
}

func TestClient_APIErrorEmptyPayload(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(404)
		_, err := res.Write([]byte(`{}`))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	health, err := client.GetHealth()

	assert.NoError(t, err)
	assert.Equal(t, 404, health.HTTPStatusCode)
}

func TestClient_ErrorNoPayload(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(200)
		_, err := res.Write([]byte{})
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	_, err := client.GetHealth()

	assert.EqualError(t, err, "unexpected end of JSON input")
}

func TestClient_HostHeader(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "my-custom-host", req.Host)
		res.Header().Add("Content-type", "application/json")
		res.WriteHeader(200)
		_, err := res.Write([]byte{})
		assert.NoError(t, err)
	}))

	client := New(srv.URL)
	client.Host = "my-custom-host"

	_, err := client.send("GET", "/", nil)
	assert.NoError(t, err)
}

func TestClient_InvalidContentType(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Content-type", "some-type")
	}))

	client := New(srv.URL)

	_, err := client.send("GET", "/", nil)

	assert.EqualError(t, err, "invalid content-type 'some-type', expected 'application/json'")
}

func TestClient_GetAsset(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicassets/v1/assets/1099667509880", req.URL.String())

		payload := `{
            "success": true,
            "data": {
            "contract": "atomicassets",
            "asset_id": "1099667509880",
            "owner": "farmersworld",
            "is_transferable": true,
            "is_burnable": true,
            "collection": {
                "collection_name": "farmersworld",
                "name": "Farmers World",
                "img": "QmX79zrJsk4DbWQ3krgu41pX3fdvEvWjkMXiNCKpxFXSgj",
                "author": ".jieg.wam",
                "allow_notify": true,
                "authorized_accounts": [
                ".jieg.wam",
                "farmersworld",
                "atomicdropsx",
                "atomicpacksx",
                "neftyblocksd"
                ],
                "notify_accounts": [
                "atomicdropsx"
                ],
                "market_fee": 0.05,
                "created_at_block": "123762633",
                "created_at_time": "1623323058000"
            },
            "schema": {
                "schema_name": "memberships",
                "format": [
                {
                    "name": "name",
                    "type": "string"
                },
                {
                    "name": "img",
                    "type": "image"
                },
                {
                    "name": "description",
                    "type": "string"
                },
                {
                    "name": "type",
                    "type": "string"
                },
                {
                    "name": "rarity",
                    "type": "string"
                },
                {
                    "name": "level",
                    "type": "uint8"
                }
                ],
                "created_at_block": "136880914",
                "created_at_time": "1629887699000"
            },
            "template": {
                "template_id": "260629",
                "max_supply": "0",
                "is_transferable": true,
                "is_burnable": true,
                "issued_supply": "112195",
                "immutable_data": {
                "img": "QmZWg1mP2UNcSwhrYNVqjk16BnhcWCz3oAva8BfiTNB3J4",
                "name": "Silver Member",
                "type": "Wood",
                "level": 2,
                "rarity": "Uncommon",
                "description": "This is a member card powered by Wood. When used by the farmer, it will increase the power and luck of the wood mining tools, and can mine the Farmer Coin that has been lost since ancient times."
                },
                "created_at_time": "1629888476000",
                "created_at_block": "136882467"
            },
            "mutable_data": {
                "asdf": "1234"
            },
            "immutable_data": {
                "asdx": "4321"
            },
            "template_mint": "4433",
            "backed_tokens": [],
            "burned_by_account": null,
            "burned_at_block": null,
            "burned_at_time": null,
            "updated_at_block": "171080009",
            "updated_at_time": "1646996870500",
            "transferred_at_block": "171080009",
            "transferred_at_time": "1646996870500",
            "minted_at_block": "171080009",
            "minted_at_time": "1646996870500",
            "data": {
                "img": "QmZWg1mP2UNcSwhrYNVqjk16BnhcWCz3oAva8BfiTNB3J4",
                "name": "Silver Member",
                "type": "Wood",
                "level": 2,
                "rarity": "Uncommon",
                "description": "This is a member card powered by Wood. When used by the farmer, it will increase the power and luck of the wood mining tools, and can mine the Farmer Coin that has been lost since ancient times."
            },
            "name": "Silver Member"
            },
            "query_time": 1647016614598
        }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	a, err := client.GetAsset("1099667509880")

	require.NoError(t, err)
	assert.Equal(t, 200, a.HTTPStatusCode)
	assert.True(t, a.Success)
	assert.Equal(t, time.Time(time.Date(2022, time.March, 11, 16, 36, 54, 598, time.UTC)), a.QueryTime.Time())
	assert.Equal(t, asset1, a.Data)
}

func TestClient_GetAssets(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicassets/v1/assets?before=100&is_transferable=true&schema_name=test", req.URL.String())

		payload := `{
  "success": true,
  "data": [
    {
      "contract": "atomicassets",
      "asset_id": "1099667509880",
      "owner": "farmersworld",
      "is_transferable": true,
      "is_burnable": true,
      "collection": {
        "collection_name": "farmersworld",
        "name": "Farmers World",
        "img": "QmX79zrJsk4DbWQ3krgu41pX3fdvEvWjkMXiNCKpxFXSgj",
        "author": ".jieg.wam",
        "allow_notify": true,
        "authorized_accounts": [
          ".jieg.wam",
          "farmersworld",
          "atomicdropsx",
          "atomicpacksx",
          "neftyblocksd"
        ],
        "notify_accounts": [
          "atomicdropsx"
        ],
        "market_fee": 0.05,
        "created_at_block": "123762633",
        "created_at_time": "1623323058000"
      },
      "schema": {
        "schema_name": "memberships",
        "format": [
          {
            "name": "name",
            "type": "string"
          },
          {
            "name": "img",
            "type": "image"
          },
          {
            "name": "description",
            "type": "string"
          },
          {
            "name": "type",
            "type": "string"
          },
          {
            "name": "rarity",
            "type": "string"
          },
          {
            "name": "level",
            "type": "uint8"
          }
        ],
        "created_at_block": "136880914",
        "created_at_time": "1629887699000"
      },
      "template": {
        "template_id": "260629",
        "max_supply": "0",
        "is_transferable": true,
        "is_burnable": true,
        "issued_supply": "112195",
        "immutable_data": {
          "img": "QmZWg1mP2UNcSwhrYNVqjk16BnhcWCz3oAva8BfiTNB3J4",
          "name": "Silver Member",
          "type": "Wood",
          "level": 2,
          "rarity": "Uncommon",
          "description": "This is a member card powered by Wood. When used by the farmer, it will increase the power and luck of the wood mining tools, and can mine the Farmer Coin that has been lost since ancient times."
        },
        "created_at_time": "1629888476000",
        "created_at_block": "136882467"
      },
      "mutable_data": {
          "asdf": "1234"
      },
      "immutable_data": {
          "asdx": "4321"
      },
      "template_mint": "4433",
      "backed_tokens": [],
      "burned_by_account": null,
      "burned_at_block": null,
      "burned_at_time": null,
      "updated_at_block": "171080009",
      "updated_at_time": "1646996870500",
      "transferred_at_block": "171080009",
      "transferred_at_time": "1646996870500",
      "minted_at_block": "171080009",
      "minted_at_time": "1646996870500",
      "data": {
        "img": "QmZWg1mP2UNcSwhrYNVqjk16BnhcWCz3oAva8BfiTNB3J4",
        "name": "Silver Member",
        "type": "Wood",
        "level": 2,
        "rarity": "Uncommon",
        "description": "This is a member card powered by Wood. When used by the farmer, it will increase the power and luck of the wood mining tools, and can mine the Farmer Coin that has been lost since ancient times."
      },
      "name": "Silver Member"
    }],
    "query_time":1646996870918
    }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	a, err := client.GetAssets(AssetsRequestParams{Before: 100, SchemaName: "test", IsTransferable: true})

	require.NoError(t, err)
	assert.Equal(t, 200, a.HTTPStatusCode)
	assert.True(t, a.Success)
	assert.Equal(t, time.Time(time.Date(2022, time.March, 11, 11, 7, 50, 918, time.UTC)), a.QueryTime.Time())

	expected := []Asset{asset1}

	assert.Equal(t, expected, a.Data)
}

func TestGetAssetLog(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicassets/v1/assets/1099667509880/logs?limit=100&order=desc&page=1", req.URL.String())

		payload := `{
            "success":true,
            "data":[
                {
                    "log_id":"41007120919",
                    "name":"logmint",
                    "data":{
                        "new_asset_owner":"farmersworld",
                        "authorized_minter":"farmersworld"
                    },
                    "txid":"4bac45fbb2fd4d5ee434ef0c682683834cec17711d3ab1d0fd44023de5c66ec9",
                    "created_at_block":"171080009",
                    "created_at_time":"1646996870500"
                }
            ],
            "query_time":1669043479123
        }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetAssetLog("1099667509880", LogRequestParams{Page: 1, Limit: 100, Order: SortDescending})

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2022, time.November, 21, 15, 11, 19, 123, time.UTC), res.QueryTime.Time())

	expected := []Log{
		{
			ID:   "41007120919",
			TxID: "4bac45fbb2fd4d5ee434ef0c682683834cec17711d3ab1d0fd44023de5c66ec9",
			Name: "logmint",
			Data: map[string]interface{}{
				"new_asset_owner":   "farmersworld",
				"authorized_minter": "farmersworld",
			},
			CreatedAtBlock: "171080009",
			CreatedAtTime:  UnixTime(1646996870500),
		},
	}

	assert.Equal(t, expected, res.Data)
}

func TestGetAssetSale(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicmarket/v1/assets/1099563680227/sales?order=desc", req.URL.String())

		payload := `{
            "success": true,
            "data": [
              {
                "market_contract": "atomicmarket",
                "sale_id": "35230996",
                "auction_id": null,
                "buyoffer_id": null,
                "token_symbol": "WAX",
                "token_precision": 8,
                "token_contract": "eosio.token",
                "price": "85000000",
                "seller": "rixcm.wam",
                "buyer": "pnbse.wam",
                "block_time": "1633004737000"
              },
              {
                "market_contract": "atomicmarket",
                "sale_id": "31692801",
                "auction_id": null,
                "buyoffer_id": null,
                "token_symbol": "WAX",
                "token_precision": 8,
                "token_contract": "eosio.token",
                "price": "9000000",
                "seller": "ryuri.wam",
                "buyer": "rixcm.wam",
                "block_time": "1630481160000"
              }
            ],
            "query_time": 1669121848963
          }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetAssetSales("1099563680227", AssetSalesRequestParams{Order: SortDescending})

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2022, time.November, 22, 12, 57, 28, 963, time.UTC), res.QueryTime.Time())

	expected := []AssetSale{
		{
			ID:             "35230996",
			MarketContract: "atomicmarket",
			TokenSymbol:    "WAX",
			TokenPrecision: 8,
			TokenContract:  "eosio.token",
			Price:          "85000000",
			Seller:         "rixcm.wam",
			Buyer:          "pnbse.wam",
			BlockTime:      1633004737000,
		},
		{
			ID:             "31692801",
			MarketContract: "atomicmarket",
			TokenSymbol:    "WAX",
			TokenPrecision: 8,
			TokenContract:  "eosio.token",
			Price:          "9000000",
			Seller:         "ryuri.wam",
			Buyer:          "rixcm.wam",
			BlockTime:      1630481160000,
		},
	}

	assert.Equal(t, expected, res.Data)
}

func TestGetAssetSaleFilterSeller(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicmarket/v1/assets/1099563680227/sales?order=asc&seller=rixcm.wam", req.URL.String())

		payload := `{
            "success": true,
            "data": [
              {
                "market_contract": "atomicmarket",
                "sale_id": "35230996",
                "auction_id": null,
                "buyoffer_id": null,
                "token_symbol": "WAX",
                "token_precision": 8,
                "token_contract": "eosio.token",
                "price": "85000000",
                "seller": "rixcm.wam",
                "buyer": "pnbse.wam",
                "block_time": "1633004737000"
              }
            ],
            "query_time": 1669121848963
          }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetAssetSales("1099563680227", AssetSalesRequestParams{Seller: "rixcm.wam", Order: SortAscending})

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2022, time.November, 22, 12, 57, 28, 963, time.UTC), res.QueryTime.Time())

	expected := []AssetSale{
		{
			ID:             "35230996",
			MarketContract: "atomicmarket",
			TokenSymbol:    "WAX",
			TokenPrecision: 8,
			TokenContract:  "eosio.token",
			Price:          "85000000",
			Seller:         "rixcm.wam",
			Buyer:          "pnbse.wam",
			BlockTime:      1633004737000,
		},
	}

	assert.Equal(t, expected, res.Data)
}

func TestGetAssetSaleFilterBuyer(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicmarket/v1/assets/1099563680227/sales?buyer=rixcm.wam", req.URL.String())

		payload := `{
            "success": true,
            "data": [
              {
                "market_contract": "atomicmarket",
                "sale_id": "31692801",
                "auction_id": null,
                "buyoffer_id": null,
                "token_symbol": "WAX",
                "token_precision": 8,
                "token_contract": "eosio.token",
                "price": "9000000",
                "seller": "ryuri.wam",
                "buyer": "rixcm.wam",
                "block_time": "1630481160000"
              }
            ],
            "query_time": 1669121848963
          }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetAssetSales("1099563680227", AssetSalesRequestParams{Buyer: "rixcm.wam"})

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2022, time.November, 22, 12, 57, 28, 963, time.UTC), res.QueryTime.Time())

	expected := []AssetSale{
		{
			ID:             "31692801",
			MarketContract: "atomicmarket",
			TokenSymbol:    "WAX",
			TokenPrecision: 8,
			TokenContract:  "eosio.token",
			Price:          "9000000",
			Seller:         "ryuri.wam",
			Buyer:          "rixcm.wam",
			BlockTime:      1630481160000,
		},
	}

	assert.Equal(t, expected, res.Data)
}

func TestGetCollections(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicassets/v1/collections?limit=2&order=desc&page=1", req.URL.String())

		payload := `{
			"success": true,
			"data": [
			  {
				"contract": "atomicassets",
				"collection_name": "itabettysart",
				"name": "Play Gruond",
				"img": "QmbKHBLk9VTHfLhxy7LJEFUjEBJLB7gGSaBL2UeSRn6oMA",
				"author": "gma3a.c.wam",
				"allow_notify": true,
				"authorized_accounts": [
				  "gma3a.c.wam"
				],
				"notify_accounts": [],
				"market_fee": 0.05,
				"data": {
				  "img": "QmbKHBLk9VTHfLhxy7LJEFUjEBJLB7gGSaBL2UeSRn6oMA",
				  "url": "https://swite.com/bettysart",
				  "name": "Play Gruond",
				  "images": "{\"banner_1920x500\":\"QmXGQBZKXH9p6qjJy7pBnCcNcue5fWZgLThj63Q9cekAim\",\"logo_512x512\":\"QmZxH22w2FVTnE8iZfTFqLdRqD8Y5SLZhH8TQDn8TFB2tV\"}",
				  "socials": "{\"twitter\":\"\",\"medium\":\"\",\"facebook\":\"\",\"github\":\"\",\"discord\":\"\",\"youtube\":\"\",\"telegram\":\"\"}",
				  "description": "real hand-made and digitised works . Italian artist decorator",
				  "creator_info": "{\"country\":\"IT\",\"address\":\"\",\"city\":\"Limite Sull'Arno (FI)\",\"zip\":\"50050\",\"company\":\"Elisabetta Rosa\",\"name\":\"Elisabetta Rosa\",\"registration_number\":\"\"}"
				},
				"created_at_time": "1669204520000",
				"created_at_block": "215481247"
			  },
			  {
				"contract": "atomicassets",
				"collection_name": "pinoydigiart",
				"name": "Filipino Digital Arts",
				"img": "QmbukpNarUzBPfTsWsLNQzKhiowviTBenQS9YgdV8nqtj7",
				"author": "utjsk.wam",
				"allow_notify": true,
				"authorized_accounts": [
				  "utjsk.wam"
				],
				"notify_accounts": [],
				"market_fee": 0.05,
				"data": {
				  "img": "QmbukpNarUzBPfTsWsLNQzKhiowviTBenQS9YgdV8nqtj7",
				  "url": "https://pinoydigiart.entriprenyur.com/",
				  "name": "Filipino Digital Arts",
				  "images": "{\"banner_1920x500\":\"QmQjStJgicWHLA7fLNVFKy1fTVGTgQCEwrhYr5WWr4LXrR\",\"logo_512x512\":\"QmbukpNarUzBPfTsWsLNQzKhiowviTBenQS9YgdV8nqtj7\"}",
				  "socials": "{\"twitter\":\"https://twitter.com/uplandcitizen\",\"medium\":\"\",\"facebook\":\"\",\"github\":\"\",\"discord\":\"\",\"youtube\":\"\",\"telegram\":\"\"}",
				  "description": "The Filipino Digital Arts collective showcases the best works of Pinoy digital artists. You can view our gallery on Pinoy Digital Arts website.",
				  "creator_info": "{\"country\":\"\",\"address\":\"\",\"city\":\"\",\"zip\":\"\",\"company\":\"\",\"name\":\"\",\"registration_number\":\"\"}"
				},
				"created_at_time": "1669190667500",
				"created_at_block": "215453568"
			  }
			],
			"query_time": 1355367264400
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetCollections(CollectionsRequestParams{Page: 1, Limit: 2, Order: SortDescending})

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2012, time.December, 13, 2, 54, 24, 400, time.UTC), res.QueryTime.Time())

	expected := []Collection{
		{
			Name:           "Play Gruond",
			CollectionName: "itabettysart",
			Contract:       "atomicassets",
			Author:         "gma3a.c.wam",
			AuthorizedAccounts: []string{
				"gma3a.c.wam",
			},
			AllowNotify:    true,
			NotifyAccounts: []string{},
			MarketFee:      0.05,
			Data: map[string]interface{}{
				"img":          "QmbKHBLk9VTHfLhxy7LJEFUjEBJLB7gGSaBL2UeSRn6oMA",
				"url":          "https://swite.com/bettysart",
				"name":         "Play Gruond",
				"images":       "{\"banner_1920x500\":\"QmXGQBZKXH9p6qjJy7pBnCcNcue5fWZgLThj63Q9cekAim\",\"logo_512x512\":\"QmZxH22w2FVTnE8iZfTFqLdRqD8Y5SLZhH8TQDn8TFB2tV\"}",
				"socials":      "{\"twitter\":\"\",\"medium\":\"\",\"facebook\":\"\",\"github\":\"\",\"discord\":\"\",\"youtube\":\"\",\"telegram\":\"\"}",
				"description":  "real hand-made and digitised works . Italian artist decorator",
				"creator_info": "{\"country\":\"IT\",\"address\":\"\",\"city\":\"Limite Sull'Arno (FI)\",\"zip\":\"50050\",\"company\":\"Elisabetta Rosa\",\"name\":\"Elisabetta Rosa\",\"registration_number\":\"\"}",
			},
			CreatedAtBlock: "215481247",
			CreatedAtTime:  UnixTime(1669204520000),
		},
		{
			Name:           "Filipino Digital Arts",
			CollectionName: "pinoydigiart",
			Contract:       "atomicassets",
			Author:         "utjsk.wam",
			AuthorizedAccounts: []string{
				"utjsk.wam",
			},
			AllowNotify:    true,
			NotifyAccounts: []string{},
			MarketFee:      0.05,
			Data: map[string]interface{}{
				"img":          "QmbukpNarUzBPfTsWsLNQzKhiowviTBenQS9YgdV8nqtj7",
				"url":          "https://pinoydigiart.entriprenyur.com/",
				"name":         "Filipino Digital Arts",
				"images":       "{\"banner_1920x500\":\"QmQjStJgicWHLA7fLNVFKy1fTVGTgQCEwrhYr5WWr4LXrR\",\"logo_512x512\":\"QmbukpNarUzBPfTsWsLNQzKhiowviTBenQS9YgdV8nqtj7\"}",
				"socials":      "{\"twitter\":\"https://twitter.com/uplandcitizen\",\"medium\":\"\",\"facebook\":\"\",\"github\":\"\",\"discord\":\"\",\"youtube\":\"\",\"telegram\":\"\"}",
				"description":  "The Filipino Digital Arts collective showcases the best works of Pinoy digital artists. You can view our gallery on Pinoy Digital Arts website.",
				"creator_info": "{\"country\":\"\",\"address\":\"\",\"city\":\"\",\"zip\":\"\",\"company\":\"\",\"name\":\"\",\"registration_number\":\"\"}",
			},
			CreatedAtBlock: "215453568",
			CreatedAtTime:  UnixTime(1669190667500),
		},
	}

	assert.Equal(t, expected, res.Data)
}

func TestGetCollection(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicassets/v1/collection/futuredoge11", req.URL.String())

		payload := `{
			"success": true,
			"data": {
				"contract": "atomicassets",
				"collection_name": "futuredoge11",
				"name": "OUT OF THE MATRIX ",
				"img": "QmWfJDKaGXrXCQyCwQZs1PEzoTKU8VBmRW7X7jDv8Vw8ou",
				"author": "d5jnk.c.wam",
				"allow_notify": true,
				"authorized_accounts": [
				  "d5jnk.c.wam"
				],
				"notify_accounts": [],
				"market_fee": 0.02,
				"data": {
				  "img": "QmWfJDKaGXrXCQyCwQZs1PEzoTKU8VBmRW7X7jDv8Vw8ou",
				  "url": "https://mariomojica11198.wixsite.com/www-mike1011989-com",
				  "name": "OUT OF THE MATRIX ",
				  "images": "{\"banner_1920x500\":\"QmSaqnwSQtvCiyaksLXz5uDMRiMee3Er3V41xLAYRZHtMZ\",\"logo_512x512\":\"QmQuQ14uEbmGoPoFhzweJCMpJtNvr6GXCQxXnRzh6dxHi4\"}",
				  "socials": "{\"twitter\":\"https://twitter.com/mike101198/status/1594230977034854400?s=46&t=tArO-6-7eEam6d_qD8BX-A\",\"medium\":\"\",\"facebook\":\"\",\"github\":\"\",\"discord\":\"\",\"youtube\":\"\",\"telegram\":\"\"}",
				  "creator_info": "{\"country\":\"\",\"address\":\"\",\"city\":\"\",\"zip\":\"\",\"company\":\"\",\"name\":\"\",\"registration_number\":\"\"}"
				},
				"created_at_time": "1669185229500",
				"created_at_block": "215442702"
			  },
			"query_time": 1355367264400
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetCollection("futuredoge11")

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2012, time.December, 13, 2, 54, 24, 400, time.UTC), res.QueryTime.Time())

	expected := Collection{
		Name:           "OUT OF THE MATRIX ",
		CollectionName: "futuredoge11",
		Contract:       "atomicassets",
		Author:         "d5jnk.c.wam",
		AuthorizedAccounts: []string{
			"d5jnk.c.wam",
		},
		AllowNotify:    true,
		NotifyAccounts: []string{},
		MarketFee:      0.02,
		Data: map[string]interface{}{
			"img":          "QmWfJDKaGXrXCQyCwQZs1PEzoTKU8VBmRW7X7jDv8Vw8ou",
			"url":          "https://mariomojica11198.wixsite.com/www-mike1011989-com",
			"name":         "OUT OF THE MATRIX ",
			"images":       "{\"banner_1920x500\":\"QmSaqnwSQtvCiyaksLXz5uDMRiMee3Er3V41xLAYRZHtMZ\",\"logo_512x512\":\"QmQuQ14uEbmGoPoFhzweJCMpJtNvr6GXCQxXnRzh6dxHi4\"}",
			"socials":      "{\"twitter\":\"https://twitter.com/mike101198/status/1594230977034854400?s=46&t=tArO-6-7eEam6d_qD8BX-A\",\"medium\":\"\",\"facebook\":\"\",\"github\":\"\",\"discord\":\"\",\"youtube\":\"\",\"telegram\":\"\"}",
			"creator_info": "{\"country\":\"\",\"address\":\"\",\"city\":\"\",\"zip\":\"\",\"company\":\"\",\"name\":\"\",\"registration_number\":\"\"}",
		},
		CreatedAtBlock: "215442702",
		CreatedAtTime:  UnixTime(1669185229500),
	}

	assert.Equal(t, expected, res.Data)
}

func TestGetCollectionStats(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicassets/v1/collection/mycoolcollection/stats", req.URL.String())

		payload := `{
			"success": true,
			"data": {
				"assets": "27",
				"burned": "1",
				"burned_by_template": [
					"sometemplate"
				],
				"burned_by_schema": [],
				"templates": "1",
				"schemas": "1"
			},
			"query_time": 1355367264400
		}`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetCollectionStats("mycoolcollection")

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2012, time.December, 13, 2, 54, 24, 400, time.UTC), res.QueryTime.Time())

	expected := CollectionStats{
		Assets:           "27",
		Burned:           "1",
		BurnedByTemplate: []string{"sometemplate"},
		BurnedBySchema:   []string{},
		Templates:        "1",
		Schemas:          "1",
	}

	assert.Equal(t, expected, res.Data)
}

func TestGetCollectionLogs(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicassets/v1/collection/futuredoge11/logs?limit=100&order=asc&page=1", req.URL.String())

		payload := `{
			"success": true,
			"data": [
			  {
				"log_id": "69460695139",
				"name": "createcol",
				"data": {
				  "data": [
					{
					  "key": "name",
					  "value": [
						"string",
						"OUT OF THE MATRIX "
					  ]
					},
					{
					  "key": "img",
					  "value": [
						"string",
						"QmWfJDKaGXrXCQyCwQZs1PEzoTKU8VBmRW7X7jDv8Vw8ou"
					  ]
					},
					{
					  "key": "url",
					  "value": [
						"string",
						"https://mariomojica11198.wixsite.com/www-mike1011989-com"
					  ]
					},
					{
					  "key": "socials",
					  "value": [
						"string",
						"{\"twitter\":\"https://twitter.com/mike101198/status/1594230977034854400?s=46&t=tArO-6-7eEam6d_qD8BX-A\",\"medium\":\"\",\"facebook\":\"\",\"github\":\"\",\"discord\":\"\",\"youtube\":\"\",\"telegram\":\"\"}"
					  ]
					},
					{
					  "key": "creator_info",
					  "value": [
						"string",
						"{\"country\":\"\",\"address\":\"\",\"city\":\"\",\"zip\":\"\",\"company\":\"\",\"name\":\"\",\"registration_number\":\"\"}"
					  ]
					},
					{
					  "key": "images",
					  "value": [
						"string",
						"{\"banner_1920x500\":\"QmSaqnwSQtvCiyaksLXz5uDMRiMee3Er3V41xLAYRZHtMZ\",\"logo_512x512\":\"QmQuQ14uEbmGoPoFhzweJCMpJtNvr6GXCQxXnRzh6dxHi4\"}"
					  ]
					}
				  ],
				  "author": "d5jnk.c.wam",
				  "market_fee": 0.02,
				  "allow_notify": true,
				  "notify_accounts": [],
				  "authorized_accounts": [
					"d5jnk.c.wam"
				  ]
				},
				"txid": "fff8a4b6deebe16f3377498c2832ece181de789c78fabe1dbfe0c97b6547d165",
				"created_at_block": "215442702",
				"created_at_time": "1669185229500"
			  }
			],
			"query_time": 1669212234204
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetCollectionLogs("futuredoge11", CollectionLogsRequestParams{Page: 1, Limit: 100, Order: SortAscending})

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2022, time.November, 23, 14, 3, 54, 204, time.UTC), res.QueryTime.Time())

	expected := []Log{
		{
			ID:   "69460695139",
			TxID: "fff8a4b6deebe16f3377498c2832ece181de789c78fabe1dbfe0c97b6547d165",
			Name: "createcol",
			Data: map[string]interface{}{
				"data": []interface{}{
					map[string]interface{}{
						"key": "name",
						"value": []interface{}{
							"string",
							"OUT OF THE MATRIX ",
						},
					},
					map[string]interface{}{
						"key": "img",
						"value": []interface{}{
							"string",
							"QmWfJDKaGXrXCQyCwQZs1PEzoTKU8VBmRW7X7jDv8Vw8ou",
						},
					},
					map[string]interface{}{
						"key": "url",
						"value": []interface{}{
							"string",
							"https://mariomojica11198.wixsite.com/www-mike1011989-com",
						},
					},
					map[string]interface{}{
						"key": "socials",
						"value": []interface{}{
							"string",
							"{\"twitter\":\"https://twitter.com/mike101198/status/1594230977034854400?s=46&t=tArO-6-7eEam6d_qD8BX-A\",\"medium\":\"\",\"facebook\":\"\",\"github\":\"\",\"discord\":\"\",\"youtube\":\"\",\"telegram\":\"\"}",
						},
					},
					map[string]interface{}{
						"key": "creator_info",
						"value": []interface{}{
							"string",
							"{\"country\":\"\",\"address\":\"\",\"city\":\"\",\"zip\":\"\",\"company\":\"\",\"name\":\"\",\"registration_number\":\"\"}",
						},
					},
					map[string]interface{}{
						"key": "images",
						"value": []interface{}{
							"string",
							"{\"banner_1920x500\":\"QmSaqnwSQtvCiyaksLXz5uDMRiMee3Er3V41xLAYRZHtMZ\",\"logo_512x512\":\"QmQuQ14uEbmGoPoFhzweJCMpJtNvr6GXCQxXnRzh6dxHi4\"}",
						},
					},
				},
				"author":          "d5jnk.c.wam",
				"market_fee":      0.02,
				"allow_notify":    true,
				"notify_accounts": []interface{}{},
				"authorized_accounts": []interface{}{
					"d5jnk.c.wam",
				},
			},

			CreatedAtBlock: "215442702",
			CreatedAtTime:  UnixTime(1669185229500),
		},
	}

	assert.Equal(t, expected, res.Data)
}

func TestGetSchemas(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicassets/v1/schemas?limit=1", req.URL.String())

		payload := `{
			"success": true,
			"data": [
			  {
				"contract": "atomicassets",
				"schema_name": "myschema",
				"format": [
				  {
					"name": "name",
					"type": "string"
				  },
				  {
					"name": "img",
					"type": "image"
				  },
				  {
					"name": "video",
					"type": "string"
				  }
				],
				"collection": {
				  "collection_name": "mycollection",
				  "name": "Some Cool Collection Name ",
				  "img": "QmYUtzfGWAYrQ43eo1nNRfrYKpUS1cvCWmceCQYxjP7CkS",
				  "author": "es2fwuiv5eyf",
				  "allow_notify": true,
				  "authorized_accounts": [
					"es2fwuiv5eyf"
				  ],
				  "notify_accounts": [
					"d11dqs2fnh2v",
					"fg2ng2izgkvl"
				  ],
				  "market_fee": 0.06,
				  "created_at_block": "18683993",
				  "created_at_time": "1427545955000"
				},
				"created_at_time": "1440686620500",
				"created_at_block": "18684000",
				"assets": 22
			  }
			],
			"query_time": 986850311000
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetSchemas(SchemasRequestParams{Limit: 1})

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2001, time.April, 9, 21, 5, 11, 0, time.UTC), res.QueryTime.Time())

	expected := []Schema{
		{
			Name:     "myschema",
			Contract: "atomicassets",
			Format: []SchemaFormat{
				{
					Name: "name",
					Type: "string",
				},
				{
					Name: "img",
					Type: "image",
				},
				{
					Name: "video",
					Type: "string",
				},
			},
			Collection: Collection{
				CollectionName:     "mycollection",
				Name:               "Some Cool Collection Name ",
				Author:             "es2fwuiv5eyf",
				AllowNotify:        true,
				AuthorizedAccounts: []string{"es2fwuiv5eyf"},
				NotifyAccounts:     []string{"d11dqs2fnh2v", "fg2ng2izgkvl"},
				MarketFee:          0.06,
				CreatedAtBlock:     "18683993",
				CreatedAtTime:      UnixTime(1427545955000),
			},
			CreatedAtBlock: "18684000",
			CreatedAtTime:  UnixTime(1440686620500),
		},
	}

	assert.Equal(t, expected, res.Data)
}
