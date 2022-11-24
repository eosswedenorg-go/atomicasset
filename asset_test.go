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
		Image:          "QmX79zrJsk4DbWQ3krgu41pX3fdvEvWjkMXiNCKpxFXSgj",
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
