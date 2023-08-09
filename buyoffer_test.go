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

var collection Collection = Collection{
	Name:           "Unlinked",
	CollectionName: "unlinked",
	Image:          "QmaUCCtENe4cbryfcyyKpk7Wsytz1CttuFPQCBXAa3PqDY",
	Author:         "unlinked",
	AllowNotify:    true,
	AuthorizedAccounts: []string{
		"unlinked",
		"atomicdropsx",
		"blend.nefty",
		"c.unlinked",
		"atomicpacksx",
		"neftyblocksd",
	},
	NotifyAccounts: []string{
		"c.unlinked",
		"s.unlinked",
	},
	MarketFee:      0.05,
	CreatedAtBlock: "160757701",
	CreatedAtTime:  unixtime.Time(1641834018000),
}

var schema InlineSchema = InlineSchema{
	Name: "distributors",
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
			Name: "info",
			Type: "string",
		},
		{
			Name: "score",
			Type: "string",
		},
		{
			Name: "rarity",
			Type: "string",
		},
		{
			Name: "level",
			Type: "uint64",
		},
	},
	CreatedAtBlock: "167835321",
	CreatedAtTime:  unixtime.Time(1645374414000),
}

var template Template = Template{
	ID:             "443115",
	MaxSupply:      "0",
	IsTransferable: true,
	IsBurnable:     true,
	IssuedSupply:   "4413",
	ImmutableData: map[string]interface{}{
		"img":    "QmPXzxCQREu5yCC8TKc61T4QcpHnrG6dZvq7GuVyGTmccn",
		"info":   "This distributor produces GRAIN daily",
		"name":   "GRAIN Distributor",
		"score":  "https://whitepaper.unlinked.io/the-game/distributors/scores",
		"rarity": "Rare",
	},
	CreatedAtBlock: "167837248",
	CreatedAtTime:  unixtime.Time(1645375378000),
}

func TestGetBuyOffer(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicmarket/v1/buyoffers/10000", req.URL.String())

		payload := `{
			"success": true,
			"data": {
			  "market_contract": "atomicmarket",
			  "assets_contract": "atomicassets",
			  "buyoffer_id": "10000",
			  "seller": "slgz2.wam",
			  "buyer": "scammer",
			  "price": {
				"token_contract": "eosio.token",
				"token_symbol": "WAX",
				"token_precision": 8,
				"amount": "10000"
			  },
			  "assets": [
				{
				  "contract": "atomicassets",
				  "asset_id": "1099590505186",
				  "owner": null,
				  "is_transferable": true,
				  "is_burnable": true,
				  "collection": {
					"collection_name": "unlinked",
					"name": "Unlinked",
					"img": "QmaUCCtENe4cbryfcyyKpk7Wsytz1CttuFPQCBXAa3PqDY",
					"author": "unlinked",
					"allow_notify": true,
					"authorized_accounts": [
					  "unlinked",
					  "atomicdropsx",
					  "blend.nefty",
					  "c.unlinked",
					  "atomicpacksx",
					  "neftyblocksd"
					],
					"notify_accounts": [
					  "c.unlinked",
					  "s.unlinked"
					],
					"market_fee": 0.05,
					"created_at_block": "160757701",
					"created_at_time": "1641834018000"
				  },
				  "schema": {
					"schema_name": "distributors",
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
						"name": "info",
						"type": "string"
					  },
					  {
						"name": "score",
						"type": "string"
					  },
					  {
						"name": "rarity",
						"type": "string"
					  },
					  {
						"name": "level",
						"type": "uint64"
					  }
					],
					"created_at_block": "167835321",
					"created_at_time": "1645374414000"
				  },
				  "template": {
					"template_id": "443115",
					"max_supply": "0",
					"is_transferable": true,
					"is_burnable": true,
					"issued_supply": "4413",
					"immutable_data": {
					  "img": "QmPXzxCQREu5yCC8TKc61T4QcpHnrG6dZvq7GuVyGTmccn",
					  "info": "This distributor produces GRAIN daily",
					  "name": "GRAIN Distributor",
					  "score": "https://whitepaper.unlinked.io/the-game/distributors/scores",
					  "rarity": "Rare"
					},
					"created_at_time": "1645375378000",
					"created_at_block": "167837248"
				  },
				  "mutable_data": {},
				  "immutable_data": {},
				  "template_mint": "2415382",
				  "backed_tokens": [],
				  "burned_by_account": null,
				  "burned_at_block": null,
				  "burned_at_time": null,
				  "updated_at_block": "171265463",
				  "updated_at_time": "1647089599500",
				  "transferred_at_block": "162290829",
				  "transferred_at_time": "1642600627500",
				  "minted_at_block": "155685360",
				  "minted_at_time": "1639297074000",
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
					  "median": "79500000",
					  "average": "246913187",
					  "suggested_median": "68607",
					  "suggested_average": "74859",
					  "min": "1",
					  "max": "637771700000",
					  "sales": "3801879"
					}
				  ],
				  "data": {
					"img": "QmPXzxCQREu5yCC8TKc61T4QcpHnrG6dZvq7GuVyGTmccn",
					"info": "This distributor produces GRAIN daily",
					"name": "GRAIN Distributor",
					"score": "https://whitepaper.unlinked.io/the-game/distributors/scores",
					"rarity": "Rare"
				  },
				  "name": "GRAIN Distributor"
				}
			  ],
			  "maker_marketplace": "",
			  "taker_marketplace": null,
			  "collection": {
				"collection_name": "unlinked",
				"name": "Unlinked",
				"img": "QmaUCCtENe4cbryfcyyKpk7Wsytz1CttuFPQCBXAa3PqDY",
				"author": "unlinked",
				"allow_notify": true,
				"authorized_accounts": [
				  "unlinked",
				  "atomicdropsx",
				  "blend.nefty",
				  "c.unlinked",
				  "atomicpacksx",
				  "neftyblocksd"
				],
				"notify_accounts": [
				  "c.unlinked",
				  "s.unlinked"
				],
				"market_fee": 0.05,
				"created_at_block": "160757701",
				"created_at_time": "1641834018000"
			  },
			  "memo": "You've been offered 117.00 WAX. Click on the confirm button to sell your NFT. This transaction has a 0.000091 WAX fee.",
			  "decline_memo": null,
			  "updated_at_block": "156854333",
			  "updated_at_time": "1639882183500",
			  "created_at_block": "156667517",
			  "created_at_time": "1639788773000",
			  "state": 2
			},
			"query_time": 1701735576000
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetBuyOffer(10000)

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2023, time.December, 5, 0, 19, 36, 0, time.UTC), res.QueryTime.Time())

	expected := BuyOffer{
		ID:             "10000",
		MarketContract: "atomicmarket",
		AssetsContract: "atomicassets",
		Seller:         "slgz2.wam",
		Buyer:          "scammer",
		Price: Token{
			Contract:  "eosio.token",
			Symbol:    "WAX",
			Precision: 8,
			Amount:    "10000",
		},
		Assets: []Asset{
			{
				ID:                "1099590505186",
				Name:              "GRAIN Distributor",
				Contract:          "atomicassets",
				Owner:             "",
				IsTransferable:    true,
				IsBurnable:        true,
				Collection:        collection,
				Schema:            schema,
				Template:          template,
				MutableData:       map[string]interface{}{},
				ImmutableData:     map[string]interface{}{},
				TemplateMint:      "2415382",
				BackedTokens:      []Token{},
				BurnedByAccount:   "",
				BurnedAtBlock:     "",
				BurnedAtTime:      unixtime.Time(0),
				UpdatedAtBlock:    "171265463",
				UpdatedAtTime:     unixtime.Time(1647089599500),
				TransferedAtBlock: "162290829",
				TransferedAtTime:  unixtime.Time(1642600627500),
				MintedAtBlock:     "155685360",
				MintedAtTime:      unixtime.Time(1639297074000),
				Data: map[string]interface{}{
					"img":    "QmPXzxCQREu5yCC8TKc61T4QcpHnrG6dZvq7GuVyGTmccn",
					"info":   "This distributor produces GRAIN daily",
					"name":   "GRAIN Distributor",
					"score":  "https://whitepaper.unlinked.io/the-game/distributors/scores",
					"rarity": "Rare",
				},
			},
		},
		MakerMarketplace: "",
		TakerMarketplace: "",
		Collection:       collection,
		Memo:             "You've been offered 117.00 WAX. Click on the confirm button to sell your NFT. This transaction has a 0.000091 WAX fee.",
		DeclineMemo:      "",
		UpdatedAtBlock:   "156854333",
		UpdatedAtTime:    unixtime.Time(1639882183500),
		CreatedAtBlock:   "156667517",
		CreatedAtTime:    unixtime.Time(1639788773000),
		State:            SalesStateCanceled,
	}

	assert.Equal(t, expected, res.Data)
}

func TestGetBuyOfferLogs(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicmarket/v1/buyoffers/10000/logs", req.URL.String())

		payload := `{
			"success": true,
			"data": [
			  {
				"log_id": "31520835672",
				"name": "cancelbuyo",
				"data": {},
				"txid": "60fcebaaae9d788ff7c303ad83c9256394b1c76b0487fa7d2ad2c2667c1a3bef",
				"created_at_block": "156854333",
				"created_at_time": "1639882183500"
			  },
			  {
				"log_id": "31389623810",
				"name": "lognewbuyo",
				"data": {
				  "collection_fee": 0.05,
				  "maker_marketplace": ""
				},
				"txid": "1f0ff70331e08875a704a6e0b4945f4aa579d4e2a765a5b658d550ac584e11c4",
				"created_at_block": "156667517",
				"created_at_time": "1639788773000"
			  }
			],
			"query_time": 1076821754000
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetBuyOfferLogs(10000, LogRequestParams{})

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2004, time.February, 15, 5, 9, 14, 0, time.UTC), res.QueryTime.Time())

	expected := []Log{
		{
			ID:             "31520835672",
			Name:           "cancelbuyo",
			Data:           map[string]interface{}{},
			TxID:           "60fcebaaae9d788ff7c303ad83c9256394b1c76b0487fa7d2ad2c2667c1a3bef",
			CreatedAtBlock: "156854333",
			CreatedAtTime:  unixtime.Time(1639882183500),
		},
		{
			ID:   "31389623810",
			Name: "lognewbuyo",
			Data: map[string]interface{}{
				"collection_fee":    float64(0.05),
				"maker_marketplace": "",
			},
			TxID:           "1f0ff70331e08875a704a6e0b4945f4aa579d4e2a765a5b658d550ac584e11c4",
			CreatedAtBlock: "156667517",
			CreatedAtTime:  unixtime.Time(1639788773000),
		},
	}

	assert.Equal(t, expected, res.Data)
}

func TestGetBuyOffers(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicmarket/v1/buyoffers?limit=1&page=1", req.URL.String())

		payload := `{
			"success": true,
			"data": [
			  {
				"market_contract": "atomicmarket",
				"assets_contract": "atomicassets",
				"buyoffer_id": "1504960",
				"seller": "euoza.wam",
				"buyer": "fu2hw.wam",
				"price": {
				  "token_contract": "eosio.token",
				  "token_symbol": "WAX",
				  "token_precision": 8,
				  "amount": "700000000"
				},
				"assets": [
				  {
					"contract": "atomicassets",
					"asset_id": "1099654773899",
					"owner": "euoza.wam",
					"is_transferable": true,
					"is_burnable": true,
					"collection": {
					  "collection_name": "unlinked",
					  "name": "Unlinked",
					  "img": "QmaUCCtENe4cbryfcyyKpk7Wsytz1CttuFPQCBXAa3PqDY",
					  "author": "unlinked",
					  "allow_notify": true,
					  "authorized_accounts": [
						"unlinked",
						"atomicdropsx",
						"blend.nefty",
						"c.unlinked",
						"atomicpacksx",
						"neftyblocksd"
					  ],
					  "notify_accounts": [
						"c.unlinked",
						"s.unlinked"
					  ],
					  "market_fee": 0.05,
					  "created_at_block": "160757701",
					  "created_at_time": "1641834018000"
					},
					"schema": {
					  "schema_name": "distributors",
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
						  "name": "info",
						  "type": "string"
						},
						{
						  "name": "score",
						  "type": "string"
						},
						{
						  "name": "rarity",
						  "type": "string"
						},
						{
						  "name": "level",
						  "type": "uint64"
						}
					  ],
					  "created_at_block": "167835321",
					  "created_at_time": "1645374414000"
					},
					"template": {
					  "template_id": "443115",
					  "max_supply": "0",
					  "is_transferable": true,
					  "is_burnable": true,
					  "issued_supply": "4413",
					  "immutable_data": {
						"img": "QmPXzxCQREu5yCC8TKc61T4QcpHnrG6dZvq7GuVyGTmccn",
						"info": "This distributor produces GRAIN daily",
						"name": "GRAIN Distributor",
						"score": "https://whitepaper.unlinked.io/the-game/distributors/scores",
						"rarity": "Rare"
					  },
					  "created_at_time": "1645375378000",
					  "created_at_block": "167837248"
					},
					"mutable_data": {},
					"immutable_data": {},
					"template_mint": "1148",
					"backed_tokens": [],
					"burned_by_account": null,
					"burned_at_block": null,
					"burned_at_time": null,
					"updated_at_block": "169238455",
					"updated_at_time": "1646076037500",
					"transferred_at_block": "169238455",
					"transferred_at_time": "1646076037500",
					"minted_at_block": "169238455",
					"minted_at_time": "1646076037500",
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
						"median": "3846153846",
						"average": "3788152303",
						"suggested_median": "580000000",
						"suggested_average": "570000000",
						"min": "100000000",
						"max": "9899000000",
						"sales": "180"
					  }
					],
					"data": {
					  "img": "QmPXzxCQREu5yCC8TKc61T4QcpHnrG6dZvq7GuVyGTmccn",
					  "info": "This distributor produces GRAIN daily",
					  "name": "GRAIN Distributor",
					  "score": "https://whitepaper.unlinked.io/the-game/distributors/scores",
					  "rarity": "Rare"
					},
					"name": "GRAIN Distributor"
				  },
				  {
					"contract": "atomicassets",
					"asset_id": "1099654773836",
					"owner": "euoza.wam",
					"is_transferable": true,
					"is_burnable": true,
					"collection": {
					  "collection_name": "unlinked",
					  "name": "Unlinked",
					  "img": "QmaUCCtENe4cbryfcyyKpk7Wsytz1CttuFPQCBXAa3PqDY",
					  "author": "unlinked",
					  "allow_notify": true,
					  "authorized_accounts": [
						"unlinked",
						"atomicdropsx",
						"blend.nefty",
						"c.unlinked",
						"atomicpacksx",
						"neftyblocksd"
					  ],
					  "notify_accounts": [
						"c.unlinked",
						"s.unlinked"
					  ],
					  "market_fee": 0.05,
					  "created_at_block": "160757701",
					  "created_at_time": "1641834018000"
					},
					"schema": {
					  "schema_name": "distributors",
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
						  "name": "info",
						  "type": "string"
						},
						{
						  "name": "score",
						  "type": "string"
						},
						{
						  "name": "rarity",
						  "type": "string"
						},
						{
						  "name": "level",
						  "type": "uint64"
						}
					  ],
					  "created_at_block": "167835321",
					  "created_at_time": "1645374414000"
					},
					"template": {
					  "template_id": "443130",
					  "max_supply": "0",
					  "is_transferable": true,
					  "is_burnable": true,
					  "issued_supply": "4398",
					  "immutable_data": {
						"img": "QmdaSvs4q8grUCatmVuwrWTfA1mb1ge4k9vLPkm6UwrhZQ",
						"info": "This distributor produces PLASMA daily",
						"name": "PLASMA Distributor",
						"score": "https://whitepaper.unlinked.io/the-game/distributors/scores",
						"rarity": "Rare"
					  },
					  "created_at_time": "1645375945000",
					  "created_at_block": "167838382"
					},
					"mutable_data": {},
					"immutable_data": {},
					"template_mint": "1168",
					"backed_tokens": [],
					"burned_by_account": null,
					"burned_at_block": null,
					"burned_at_time": null,
					"updated_at_block": "169238453",
					"updated_at_time": "1646076036500",
					"transferred_at_block": "169238453",
					"transferred_at_time": "1646076036500",
					"minted_at_block": "169238453",
					"minted_at_time": "1646076036500",
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
						"median": "3900000000",
						"average": "3661324705",
						"suggested_median": "600000000",
						"suggested_average": "585000000",
						"min": "200000000",
						"max": "8400000000",
						"sales": "187"
					  }
					],
					"data": {
					  "img": "QmdaSvs4q8grUCatmVuwrWTfA1mb1ge4k9vLPkm6UwrhZQ",
					  "info": "This distributor produces PLASMA daily",
					  "name": "PLASMA Distributor",
					  "score": "https://whitepaper.unlinked.io/the-game/distributors/scores",
					  "rarity": "Rare"
					},
					"name": "PLASMA Distributor"
				  }
				],
				"maker_marketplace": "",
				"taker_marketplace": null,
				"collection": {
				  "collection_name": "unlinked",
				  "name": "Unlinked",
				  "img": "QmaUCCtENe4cbryfcyyKpk7Wsytz1CttuFPQCBXAa3PqDY",
				  "author": "unlinked",
				  "allow_notify": true,
				  "authorized_accounts": [
					"unlinked",
					"atomicdropsx",
					"blend.nefty",
					"c.unlinked",
					"atomicpacksx",
					"neftyblocksd"
				  ],
				  "notify_accounts": [
					"c.unlinked",
					"s.unlinked"
				  ],
				  "market_fee": 0.05,
				  "created_at_block": "160757701",
				  "created_at_time": "1641834018000"
				},
				"memo": "memo",
				"decline_memo": "declined",
				"updated_at_block": "219983297",
				"updated_at_time": "1671457786000",
				"created_at_block": "219983297",
				"created_at_time": "1671457786000",
				"state": 0
			  }
			],
			"query_time": 1623321161000
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetBuyOffers(AuctionsRequestParams{Limit: 1, Page: 1})

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2021, time.June, 10, 10, 32, 41, 0, time.UTC), res.QueryTime.Time())

	expected := BuyOffer{
		ID:             "1504960",
		MarketContract: "atomicmarket",
		AssetsContract: "atomicassets",
		Seller:         "euoza.wam",
		Buyer:          "fu2hw.wam",
		Price: Token{
			Contract:  "eosio.token",
			Symbol:    "WAX",
			Precision: 8,
			Amount:    "700000000",
		},
		Assets: []Asset{
			{
				ID:                "1099654773899",
				Name:              "GRAIN Distributor",
				Contract:          "atomicassets",
				Owner:             "euoza.wam",
				IsTransferable:    true,
				IsBurnable:        true,
				Collection:        collection,
				Schema:            schema,
				Template:          template,
				MutableData:       map[string]interface{}{},
				ImmutableData:     map[string]interface{}{},
				TemplateMint:      "1148",
				BackedTokens:      []Token{},
				BurnedByAccount:   "",
				BurnedAtBlock:     "",
				BurnedAtTime:      unixtime.Time(0),
				UpdatedAtBlock:    "169238455",
				UpdatedAtTime:     unixtime.Time(1646076037500),
				TransferedAtBlock: "169238455",
				TransferedAtTime:  unixtime.Time(1646076037500),
				MintedAtBlock:     "169238455",
				MintedAtTime:      unixtime.Time(1646076037500),
				Data: map[string]interface{}{
					"img":    "QmPXzxCQREu5yCC8TKc61T4QcpHnrG6dZvq7GuVyGTmccn",
					"info":   "This distributor produces GRAIN daily",
					"name":   "GRAIN Distributor",
					"score":  "https://whitepaper.unlinked.io/the-game/distributors/scores",
					"rarity": "Rare",
				},
			},
			{
				ID:             "1099654773836",
				Name:           "PLASMA Distributor",
				Contract:       "atomicassets",
				Owner:          "euoza.wam",
				IsTransferable: true,
				IsBurnable:     true,
				Collection:     collection,
				Schema:         schema,
				Template: Template{
					ID:             "443130",
					MaxSupply:      "0",
					IsTransferable: true,
					IsBurnable:     true,
					IssuedSupply:   "4398",
					ImmutableData: map[string]interface{}{
						"img":    "QmdaSvs4q8grUCatmVuwrWTfA1mb1ge4k9vLPkm6UwrhZQ",
						"info":   "This distributor produces PLASMA daily",
						"name":   "PLASMA Distributor",
						"score":  "https://whitepaper.unlinked.io/the-game/distributors/scores",
						"rarity": "Rare",
					},
					CreatedAtBlock: "167838382",
					CreatedAtTime:  unixtime.Time(1645375945000),
				},
				MutableData:       map[string]interface{}{},
				ImmutableData:     map[string]interface{}{},
				TemplateMint:      "1168",
				BackedTokens:      []Token{},
				BurnedByAccount:   "",
				BurnedAtBlock:     "",
				BurnedAtTime:      unixtime.Time(0),
				UpdatedAtBlock:    "169238453",
				UpdatedAtTime:     unixtime.Time(1646076036500),
				TransferedAtBlock: "169238453",
				TransferedAtTime:  unixtime.Time(1646076036500),
				MintedAtBlock:     "169238453",
				MintedAtTime:      unixtime.Time(1646076036500),

				Data: map[string]interface{}{
					"img":    "QmdaSvs4q8grUCatmVuwrWTfA1mb1ge4k9vLPkm6UwrhZQ",
					"info":   "This distributor produces PLASMA daily",
					"name":   "PLASMA Distributor",
					"score":  "https://whitepaper.unlinked.io/the-game/distributors/scores",
					"rarity": "Rare",
				},
			},
		},
		MakerMarketplace: "",
		TakerMarketplace: "",
		Collection:       collection,
		Memo:             "memo",
		DeclineMemo:      "declined",
		UpdatedAtBlock:   "219983297",
		UpdatedAtTime:    unixtime.Time(1671457786000),
		CreatedAtBlock:   "219983297",
		CreatedAtTime:    unixtime.Time(1671457786000),
		State:            SalesStateWaiting,
	}

	assert.Equal(t, []BuyOffer{expected}, res.Data)
}
