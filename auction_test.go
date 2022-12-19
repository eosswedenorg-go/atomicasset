package atomicasset

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAuction(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicmarket/v1/auctions/10000", req.URL.String())

		payload := `{
			"success": true,
			"data": {
			  "market_contract": "atomicmarket",
			  "assets_contract": "atomicassets",
			  "auction_id": "10000",
			  "seller": "svwqu.wam",
			  "buyer": null,
			  "price": {
				"token_contract": "eosio.token",
				"token_symbol": "WAX",
				"token_precision": 8,
				"amount": "300000000000"
			  },
			  "assets": [
				{
				  "contract": "atomicassets",
				  "asset_id": "1099513564595",
				  "owner": "s.rplanet",
				  "is_transferable": true,
				  "is_burnable": true,
				  "collection": {
					"collection_name": "officialhero",
					"name": "Blockchain Heroes",
					"img": "QmSVTpCkchHSaWJ1VXU16pkVayEouKBFCZZ5xhBGUessiu",
					"author": "heroes",
					"allow_notify": true,
					"authorized_accounts": [
					  "heroes",
					  "air.atomic",
					  "unbox.heroes",
					  "atomicdropsx",
					  "theniftyshop",
					  "atomicpoolsx",
					  "blenderizerx",
					  "heroespoolsx",
					  "onessusdrops"
					],
					"notify_accounts": [],
					"market_fee": 0.06,
					"created_at_block": "65772806",
					"created_at_time": "1594315642000"
				  },
				  "schema": {
					"schema_name": "series2.x",
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
						"name": "backimg",
						"type": "image"
					  },
					  {
						"name": "video",
						"type": "string"
					  },
					  {
						"name": "rarity",
						"type": "string"
					  },
					  {
						"name": "variation",
						"type": "string"
					  },
					  {
						"name": "cardid",
						"type": "uint64"
					  },
					  {
						"name": "description",
						"type": "string"
					  }
					],
					"created_at_block": "93213470",
					"created_at_time": "1608040192000"
				  },
				  "template": {
					"template_id": "42427",
					"max_supply": "0",
					"is_transferable": true,
					"is_burnable": true,
					"issued_supply": "233",
					"immutable_data": {
					  "img": "QmdM2Q1QXXAPtWMUifLVN4ySbqaZL6pE2f82HRrnAwioLq/img/01.gif",
					  "name": "Sentinel-256",
					  "video": "QmdM2Q1QXXAPtWMUifLVN4ySbqaZL6pE2f82HRrnAwioLq/video/01.mp4",
					  "rarity": "collectors",
					  "variation": "Reward"
					},
					"created_at_time": "1608735860500",
					"created_at_block": "94604712"
				  },
				  "mutable_data": {},
				  "immutable_data": {},
				  "template_mint": "15",
				  "backed_tokens": [],
				  "burned_by_account": "acc.wam",
				  "burned_at_block": "112878022",
				  "burned_at_time": "1617878250000",
				  "updated_at_block": "112878020",
				  "updated_at_time": "1617878249000",
				  "transferred_at_block": "112878020",
				  "transferred_at_time": "1617878249000",
				  "minted_at_block": "94619162",
				  "minted_at_time": "1608743086000",
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
					  "median": "7000000000",
					  "average": "8542377561",
					  "suggested_median": "288178803",
					  "suggested_average": "288178803",
					  "min": "300000000",
					  "max": "35000000000",
					  "sales": "48"
					}
				  ],
				  "data": {
					"img": "QmdM2Q1QXXAPtWMUifLVN4ySbqaZL6pE2f82HRrnAwioLq/img/01.gif",
					"name": "Sentinel-256",
					"video": "QmdM2Q1QXXAPtWMUifLVN4ySbqaZL6pE2f82HRrnAwioLq/video/01.mp4",
					"rarity": "collectors",
					"variation": "Reward"
				  },
				  "name": "Sentinel-256"
				}
			  ],
			  "bids": [],
			  "maker_marketplace": "",
			  "taker_marketplace": null,
			  "claimed_by_buyer": false,
			  "claimed_by_seller": false,
			  "collection": {
				"collection_name": "officialhero",
				"name": "Blockchain Heroes",
				"img": "QmSVTpCkchHSaWJ1VXU16pkVayEouKBFCZZ5xhBGUessiu",
				"author": "heroes",
				"allow_notify": true,
				"authorized_accounts": [
				  "heroes",
				  "air.atomic",
				  "unbox.heroes",
				  "atomicdropsx",
				  "theniftyshop",
				  "atomicpoolsx",
				  "blenderizerx",
				  "heroespoolsx",
				  "onessusdrops"
				],
				"notify_accounts": [],
				"market_fee": 0.06,
				"created_at_block": "65772806",
				"created_at_time": "1594315642000"
			  },
			  "end_time": "1611186498000",
			  "is_seller_contract": false,
			  "updated_at_block": "99508886",
			  "updated_at_time": "1611188638000",
			  "created_at_block": "99331913",
			  "created_at_time": "1611100098000",
			  "state": 2
			},
			"query_time": 1081349214000
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetAuction(10000)

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2004, time.April, 7, 14, 46, 54, 0, time.UTC), res.QueryTime.Time())

	collection := Collection{
		Name:           "Blockchain Heroes",
		CollectionName: "officialhero",
		Image:          "QmSVTpCkchHSaWJ1VXU16pkVayEouKBFCZZ5xhBGUessiu",
		Author:         "heroes",
		AllowNotify:    true,
		AuthorizedAccounts: []string{
			"heroes",
			"air.atomic",
			"unbox.heroes",
			"atomicdropsx",
			"theniftyshop",
			"atomicpoolsx",
			"blenderizerx",
			"heroespoolsx",
			"onessusdrops",
		},
		NotifyAccounts: []string{},
		MarketFee:      0.06,
		CreatedAtBlock: "65772806",
		CreatedAtTime:  UnixTime(1594315642000),
	}

	expected := Auction{
		ID:             "10000",
		MarketContract: "atomicmarket",
		AssetsContract: "atomicassets",
		Seller:         "svwqu.wam",
		Buyer:          "",
		Price: Token{
			Contract:  "eosio.token",
			Symbol:    "WAX",
			Precision: 8,
			Amount:    "300000000000",
		},
		Assets: []Asset{
			{
				Name:           "Sentinel-256",
				ID:             "1099513564595",
				Contract:       "atomicassets",
				Owner:          "s.rplanet",
				IsTransferable: true,
				IsBurnable:     true,
				Collection:     collection,
				Schema: InlineSchema{
					Name: "series2.x",
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
							Name: "backimg",
							Type: "image",
						},
						{
							Name: "video",
							Type: "string",
						},
						{
							Name: "rarity",
							Type: "string",
						},
						{
							Name: "variation",
							Type: "string",
						},
						{
							Name: "cardid",
							Type: "uint64",
						},
						{
							Name: "description",
							Type: "string",
						},
					},
					CreatedAtBlock: "93213470",
					CreatedAtTime:  UnixTime(1608040192000),
				},
				Template: Template{
					ID:             "42427",
					MaxSupply:      "0",
					IsTransferable: true,
					IsBurnable:     true,
					IssuedSupply:   "233",
					ImmutableData: map[string]interface{}{
						"img":       "QmdM2Q1QXXAPtWMUifLVN4ySbqaZL6pE2f82HRrnAwioLq/img/01.gif",
						"name":      "Sentinel-256",
						"video":     "QmdM2Q1QXXAPtWMUifLVN4ySbqaZL6pE2f82HRrnAwioLq/video/01.mp4",
						"rarity":    "collectors",
						"variation": "Reward",
					},
					CreatedAtBlock: "94604712",
					CreatedAtTime:  UnixTime(1608735860500),
				},
				MutableData:       map[string]interface{}{},
				ImmutableData:     map[string]interface{}{},
				TemplateMint:      "15",
				BackedTokens:      []Token{},
				BurnedByAccount:   "acc.wam",
				BurnedAtBlock:     "112878022",
				BurnedAtTime:      UnixTime(1617878250000),
				UpdatedAtBlock:    "112878020",
				UpdatedAtTime:     UnixTime(1617878249000),
				TransferedAtBlock: "112878020",
				TransferedAtTime:  UnixTime(1617878249000),
				MintedAtBlock:     "94619162",
				MintedAtTime:      UnixTime(1608743086000),
				Data: map[string]interface{}{
					"img":       "QmdM2Q1QXXAPtWMUifLVN4ySbqaZL6pE2f82HRrnAwioLq/img/01.gif",
					"name":      "Sentinel-256",
					"video":     "QmdM2Q1QXXAPtWMUifLVN4ySbqaZL6pE2f82HRrnAwioLq/video/01.mp4",
					"rarity":    "collectors",
					"variation": "Reward",
				},
			},
		},
		Bids:             []Bid{},
		MakerMarketplace: "",
		TakerMarketplace: "",
		ClaimedByBuyer:   false,
		ClaimedBySeller:  false,
		Collection:       collection,
		EndTime:          UnixTime(1611186498000),
		IsSellerContract: false,
		UpdatedAtBlock:   "99508886",
		UpdatedAtTime:    UnixTime(1611188638000),
		CreatedAtBlock:   "99331913",
		CreatedAtTime:    UnixTime(1611100098000),
		State:            SalesStateCanceled,
	}

	assert.Equal(t, expected, res.Data)
}

func TestGetAuctionLogs(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicmarket/v1/auctions/10000/logs", req.URL.String())

		payload := `{
			"success": true,
			"data": [
			  {
				"log_id": "890251761",
				"name": "cancelauct",
				"data": {},
				"txid": "d73c7955472224a5a077075956d5a0b7e4981658efcf0f120bbc187d4973f471",
				"created_at_block": "99508886",
				"created_at_time": "1611188638000"
			  },
			  {
				"log_id": "883565679",
				"name": "logauctstart",
				"data": {},
				"txid": "0ff331772b222abed459ca4536c3654f3eb3b636aa7f74963e5eca93a74b0c70",
				"created_at_block": "99331913",
				"created_at_time": "1611100098000"
			  },
			  {
				"log_id": "883565673",
				"name": "lognewauct",
				"data": {
				  "starting_bid": "3000.00000000 WAX",
				  "collection_fee": 0.06,
				  "maker_marketplace": ""
				},
				"txid": "0ff331772b222abed459ca4536c3654f3eb3b636aa7f74963e5eca93a74b0c70",
				"created_at_block": "99331913",
				"created_at_time": "1611100098000"
			  }
			],
			"query_time": 1559724233000
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	expected := []Log{
		{
			ID:             "890251761",
			Name:           "cancelauct",
			TxID:           "d73c7955472224a5a077075956d5a0b7e4981658efcf0f120bbc187d4973f471",
			Data:           map[string]interface{}{},
			CreatedAtBlock: "99508886",
			CreatedAtTime:  UnixTime(1611188638000),
		},
		{
			ID:             "883565679",
			Name:           "logauctstart",
			TxID:           "0ff331772b222abed459ca4536c3654f3eb3b636aa7f74963e5eca93a74b0c70",
			Data:           map[string]interface{}{},
			CreatedAtBlock: "99331913",
			CreatedAtTime:  UnixTime(1611100098000),
		},
		{
			ID:   "883565673",
			Name: "lognewauct",
			TxID: "0ff331772b222abed459ca4536c3654f3eb3b636aa7f74963e5eca93a74b0c70",
			Data: map[string]interface{}{
				"starting_bid":      "3000.00000000 WAX",
				"collection_fee":    0.06,
				"maker_marketplace": "",
			},
			CreatedAtBlock: "99331913",
			CreatedAtTime:  UnixTime(1611100098000),
		},
	}

	client := New(srv.URL)

	a, err := client.GetAuctionLogs(10000, LogRequestParams{})

	require.NoError(t, err)
	assert.Equal(t, 200, a.HTTPStatusCode)
	assert.True(t, a.Success)
	assert.Equal(t, time.Date(2019, time.June, 5, 8, 43, 53, 0, time.UTC), a.QueryTime.Time())
	assert.Equal(t, expected, a.Data)
}

func TestGetAuctions(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicmarket/v2/auctions?limit=1&order=desc&page=1", req.URL.String())

		payload := `{
			"success": true,
			"data": [
			  {
				"market_contract": "atomicmarket",
				"assets_contract": "atomicassets",
				"auction_id": "1000940",
				"seller": "irxc4.wam",
				"buyer": "3wkba.wam",
				"price": {
				  "token_contract": "eosio.token",
				  "token_symbol": "WAX",
				  "token_precision": 8,
				  "amount": "3000000000"
				},
				"assets": [
				  {
					"contract": "atomicassets",
					"asset_id": "1099841916930",
					"owner": "atomicmarket",
					"is_transferable": true,
					"is_burnable": true,
					"collection": {
					  "collection_name": "tokengirlslv",
					  "name": "Token Girls",
					  "img": "QmaoBxDmdsGJgSwnmePRrtXKfpsby17AH6YLuffRxKXQyb",
					  "author": "irxc4.wam",
					  "allow_notify": true,
					  "authorized_accounts": [
						"irxc4.wam",
						"atomicdropsx",
						"neftyblocksd",
						"atomicpacksx",
						"blenderizerx",
						"blend.nefty",
						"p2wfe.wam",
						"nfthivedrops",
						"johnfromtglv"
					  ],
					  "notify_accounts": [],
					  "market_fee": 0.06,
					  "created_at_block": "117574616",
					  "created_at_time": "1620228511000"
					},
					"schema": {
					  "schema_name": "pfps.girl",
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
						  "name": "backimg",
						  "type": "image"
						},
						{
						  "name": "description",
						  "type": "string"
						}
					  ],
					  "created_at_block": "158859181",
					  "created_at_time": "1640884681000"
					},
					"template": {
					  "template_id": "642479",
					  "max_supply": "1",
					  "is_transferable": true,
					  "is_burnable": true,
					  "issued_supply": "1",
					  "immutable_data": {
						"img": "Qme4VGrcbGqM5xZwCALaskcMuhRL4rQzeu9WDUHwScJiDX",
						"name": "Neo The Cyber Witch",
						"description": "PFP - Neo The Cyber Witch - Shocked ver."
					  },
					  "created_at_time": "1671207271500",
					  "created_at_block": "219482739"
					},
					"mutable_data": {},
					"immutable_data": {},
					"template_mint": "1",
					"backed_tokens": [],
					"burned_by_account": null,
					"burned_at_block": null,
					"burned_at_time": null,
					"updated_at_block": "219483049",
					"updated_at_time": "1671207427500",
					"transferred_at_block": "219483049",
					"transferred_at_time": "1671207427500",
					"minted_at_block": "219483002",
					"minted_at_time": "1671207403000",
					"sales": [],
					"auctions": [],
					"prices": [],
					"data": {
					  "img": "Qme4VGrcbGqM5xZwCALaskcMuhRL4rQzeu9WDUHwScJiDX",
					  "name": "Neo The Cyber Witch",
					  "description": "PFP - Neo The Cyber Witch - Shocked ver."
					},
					"name": "Neo The Cyber Witch"
				  }
				],
				"bids": [
				  {
					"number": 1,
					"account": "jollewaxacc1",
					"amount": "1100000000",
					"created_at_block": "219484412",
					"created_at_time": "1671208109000",
					"txid": "9311ca95a37510213c6d969586787904bea4197fab594d186704920d806b29eb"
				  },
				  {
					"number": 2,
					"account": "3wkba.wam",
					"amount": "3000000000",
					"created_at_block": "219485027",
					"created_at_time": "1671208416500",
					"txid": "f416c24434dc1bc91850c6b95f21bdf3b4f4b29799b067714fcf1507167e41b8"
				  }
				],
				"maker_marketplace": "",
				"taker_marketplace": "",
				"claimed_by_buyer": false,
				"claimed_by_seller": false,
				"collection": {
				  "collection_name": "tokengirlslv",
				  "name": "Token Girls",
				  "img": "QmaoBxDmdsGJgSwnmePRrtXKfpsby17AH6YLuffRxKXQyb",
				  "author": "irxc4.wam",
				  "allow_notify": true,
				  "authorized_accounts": [
					"irxc4.wam",
					"atomicdropsx",
					"neftyblocksd",
					"atomicpacksx",
					"blenderizerx",
					"blend.nefty",
					"p2wfe.wam",
					"nfthivedrops",
					"johnfromtglv"
				  ],
				  "notify_accounts": [],
				  "market_fee": 0.06,
				  "created_at_block": "117574616",
				  "created_at_time": "1620228511000"
				},
				"end_time": "1671294427000",
				"is_seller_contract": false,
				"updated_at_block": "219485027",
				"updated_at_time": "1671208416500",
				"created_at_block": "219483049",
				"created_at_time": "1671207427500",
				"state": 1
			  }
			],
			"query_time": 1059555360000
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetAuctions(AuctionsRequestParams{Page: 1, Limit: 1, Order: SortDescending})

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2003, time.July, 30, 8, 56, 0, 0, time.UTC), res.QueryTime.Time())

	collection := Collection{
		CollectionName: "tokengirlslv",
		Name:           "Token Girls",
		Image:          "QmaoBxDmdsGJgSwnmePRrtXKfpsby17AH6YLuffRxKXQyb",
		Author:         "irxc4.wam",
		AllowNotify:    true,
		AuthorizedAccounts: []string{
			"irxc4.wam",
			"atomicdropsx",
			"neftyblocksd",
			"atomicpacksx",
			"blenderizerx",
			"blend.nefty",
			"p2wfe.wam",
			"nfthivedrops",
			"johnfromtglv",
		},
		NotifyAccounts: []string{},
		MarketFee:      0.06,
		CreatedAtBlock: "117574616",
		CreatedAtTime:  UnixTime(1620228511000),
	}

	expected := []Auction{
		{
			ID:             "1000940",
			MarketContract: "atomicmarket",
			AssetsContract: "atomicassets",
			Seller:         "irxc4.wam",
			Buyer:          "3wkba.wam",
			Price: Token{
				Contract:  "eosio.token",
				Symbol:    "WAX",
				Precision: 8,
				Amount:    "3000000000",
			},
			Assets: []Asset{
				{
					Name:           "Neo The Cyber Witch",
					ID:             "1099841916930",
					Contract:       "atomicassets",
					Owner:          "atomicmarket",
					IsTransferable: true,
					IsBurnable:     true,
					Collection:     collection,
					Schema: InlineSchema{
						Name: "pfps.girl",
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
								Name: "backimg",
								Type: "image",
							},
							{
								Name: "description",
								Type: "string",
							},
						},
						CreatedAtBlock: "158859181",
						CreatedAtTime:  UnixTime(1640884681000),
					},
					Template: Template{
						ID:             "642479",
						MaxSupply:      "1",
						IsTransferable: true,
						IsBurnable:     true,
						IssuedSupply:   "1",
						ImmutableData: map[string]interface{}{
							"img":         "Qme4VGrcbGqM5xZwCALaskcMuhRL4rQzeu9WDUHwScJiDX",
							"name":        "Neo The Cyber Witch",
							"description": "PFP - Neo The Cyber Witch - Shocked ver.",
						},
						CreatedAtBlock: "219482739",
						CreatedAtTime:  UnixTime(1671207271500),
					},
					MutableData:       map[string]interface{}{},
					ImmutableData:     map[string]interface{}{},
					TemplateMint:      "1",
					BackedTokens:      []Token{},
					BurnedByAccount:   "",
					BurnedAtBlock:     "",
					BurnedAtTime:      UnixTime(0),
					UpdatedAtBlock:    "219483049",
					UpdatedAtTime:     UnixTime(1671207427500),
					TransferedAtBlock: "219483049",
					TransferedAtTime:  UnixTime(1671207427500),
					MintedAtBlock:     "219483002",
					MintedAtTime:      UnixTime(1671207403000),
					Data: map[string]interface{}{
						"img":         "Qme4VGrcbGqM5xZwCALaskcMuhRL4rQzeu9WDUHwScJiDX",
						"name":        "Neo The Cyber Witch",
						"description": "PFP - Neo The Cyber Witch - Shocked ver.",
					},
				},
			},
			Bids: []Bid{
				{
					Number:         1,
					Account:        "jollewaxacc1",
					Amount:         "1100000000",
					CreatedAtBlock: "219484412",
					CreatedAtTime:  UnixTime(1671208109000),
					TxID:           "9311ca95a37510213c6d969586787904bea4197fab594d186704920d806b29eb",
				},
				{
					Number:         2,
					Account:        "3wkba.wam",
					Amount:         "3000000000",
					CreatedAtBlock: "219485027",
					CreatedAtTime:  UnixTime(1671208416500),
					TxID:           "f416c24434dc1bc91850c6b95f21bdf3b4f4b29799b067714fcf1507167e41b8",
				},
			},
			MakerMarketplace: "",
			TakerMarketplace: "",
			ClaimedByBuyer:   false,
			ClaimedBySeller:  false,
			Collection:       collection,
			EndTime:          UnixTime(1671294427000),
			IsSellerContract: false,
			UpdatedAtBlock:   "219485027",
			UpdatedAtTime:    UnixTime(1671208416500),
			CreatedAtBlock:   "219483049",
			CreatedAtTime:    UnixTime(1671207427500),
			State:            SalesStateListed,
		},
	}

	assert.Equal(t, expected, res.Data)
}
