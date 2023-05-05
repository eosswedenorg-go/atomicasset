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

func TestClient_GetOffers(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicassets/v1/offers", req.URL.String())

		payload := `{
			"success": true,
			"data": [
			  {
				"contract": "atomicassets",
				"offer_id": "103763644",
				"sender_name": "wzzwk.wam",
				"recipient_name": "atomicmarket",
				"memo": "sale",
				"state": 0,
				"sender_assets": [
				  {
					"contract": "atomicassets",
					"asset_id": "1099814256937",
					"owner": "wzzwk.wam",
					"is_transferable": true,
					"is_burnable": true,
					"collection": {
					  "collection_name": "pr.funko",
					  "name": "Power Rangers x Funko",
					  "img": "QmXo9XW76dGrLAMB4UxZNE1yomkkm8iMEGNY6UwSQ8wAsw",
					  "author": "pr.funko",
					  "allow_notify": true,
					  "authorized_accounts": [
						"pr.funko"
					  ],
					  "notify_accounts": [],
					  "market_fee": 0.06,
					  "created_at_block": "199268367",
					  "created_at_time": "1661093566500"
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
					  "created_at_block": "199271893",
					  "created_at_time": "1661095329500"
					},
					"template": {
					  "template_id": "583939",
					  "max_supply": "2400",
					  "is_transferable": true,
					  "is_burnable": true,
					  "issued_supply": "2400",
					  "immutable_data": {
						"tid": 53,
						"name": "White Ranger",
						"legal": "TM & ©2022 SCG Power Rangers LLC and Hasbro. Power Rangers and all related logos, characters, names, and distinctive likenesses thereof are the exclusive property of SCG Power Rangers LLC.  All Rights Reserved.  Used Under Authorization.",
						"video": "QmZQXQXksKmh255bieGC14mev8iUVupjWWQeBFYcRfUj4H/Front/MIGHTY-MORPHIN-POWER-RANGERS_EPIC_WHITE-RANGER_OG_ANIMATED.mp4",
						"cardid": 103,
						"rarity": "Epic",
						"backimg": "QmZQXQXksKmh255bieGC14mev8iUVupjWWQeBFYcRfUj4H/Back/back.png",
						"variant": "Animated",
						"description": "White Ranger Digital Pop!",
						"release date": "August 23, 2022",
						"end user license agreement": "https://digital.funko.com/eula/"
					  },
					  "created_at_time": "1661095780000",
					  "created_at_block": "199272794"
					},
					"mutable_data": {},
					"immutable_data": {},
					"template_mint": "1271",
					"backed_tokens": [],
					"burned_by_account": null,
					"burned_at_block": null,
					"burned_at_time": null,
					"updated_at_block": "217067471",
					"updated_at_time": "1669998379500",
					"transferred_at_block": "217067471",
					"transferred_at_time": "1669998379500",
					"minted_at_block": "199294390",
					"minted_at_time": "1661106579000",
					"data": {
					  "tid": 53,
					  "name": "White Ranger",
					  "legal": "TM & ©2022 SCG Power Rangers LLC and Hasbro. Power Rangers and all related logos, characters, names, and distinctive likenesses thereof are the exclusive property of SCG Power Rangers LLC.  All Rights Reserved.  Used Under Authorization.",
					  "video": "QmZQXQXksKmh255bieGC14mev8iUVupjWWQeBFYcRfUj4H/Front/MIGHTY-MORPHIN-POWER-RANGERS_EPIC_WHITE-RANGER_OG_ANIMATED.mp4",
					  "cardid": 103,
					  "rarity": "Epic",
					  "backimg": "QmZQXQXksKmh255bieGC14mev8iUVupjWWQeBFYcRfUj4H/Back/back.png",
					  "variant": "Animated",
					  "description": "White Ranger Digital Pop!",
					  "release date": "August 23, 2022",
					  "end user license agreement": "https://digital.funko.com/eula/"
					},
					"name": "White Ranger"
				  }
				],
				"recipient_assets": [],
				"is_sender_contract": false,
				"is_recipient_contract": true,
				"updated_at_block": "217068480",
				"updated_at_time": "1669998884000",
				"created_at_block": "217068480",
				"created_at_time": "1669998884000"
			  }
			],
			"query_time": 1597948728000
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	expected := Offer{
		ID:        "103763644",
		Contract:  "atomicassets",
		Sender:    "wzzwk.wam",
		Recipient: "atomicmarket",
		Memo:      "sale",
		State:     0,
		SenderAssets: []Asset{
			{
				ID:             "1099814256937",
				Name:           "White Ranger",
				Contract:       "atomicassets",
				Owner:          "wzzwk.wam",
				IsTransferable: true,
				IsBurnable:     true,
				Collection: Collection{
					Name:           "Power Rangers x Funko",
					CollectionName: "pr.funko",
					Image:          "QmXo9XW76dGrLAMB4UxZNE1yomkkm8iMEGNY6UwSQ8wAsw",
					Author:         "pr.funko",
					AllowNotify:    true,
					AuthorizedAccounts: []string{
						"pr.funko",
					},
					NotifyAccounts: []string{},
					MarketFee:      0.06,
					CreatedAtBlock: "199268367",
					CreatedAtTime:  unixtime.Time(1661093566500),
				},
				Schema: InlineSchema{
					Name: "series1.drop",
					Format: []SchemaFormat{
						{
							Name: "name",
							Type: "string",
						},
						{
							Name: "rarity",
							Type: "string",
						},
						{
							Name: "variant",
							Type: "string",
						},
						{
							Name: "cardid",
							Type: "uint8",
						},
						{
							Name: "legal",
							Type: "string",
						},
						{
							Name: "end user license agreement",
							Type: "string",
						},
						{
							Name: "video",
							Type: "image",
						},
						{
							Name: "backimg",
							Type: "image",
						},
						{
							Name: "tid",
							Type: "uint16",
						},
						{
							Name: "release date",
							Type: "string",
						},
						{
							Name: "description",
							Type: "string",
						},
						{
							Name: "img",
							Type: "image",
						},
					},
					CreatedAtBlock: "199271893",
					CreatedAtTime:  unixtime.Time(1661095329500),
				},
				Template: Template{
					ID:             "583939",
					MaxSupply:      "2400",
					IsTransferable: true,
					IsBurnable:     true,
					IssuedSupply:   "2400",
					ImmutableData: map[string]interface{}{
						"tid":                        float64(53),
						"name":                       "White Ranger",
						"legal":                      "TM & ©2022 SCG Power Rangers LLC and Hasbro. Power Rangers and all related logos, characters, names, and distinctive likenesses thereof are the exclusive property of SCG Power Rangers LLC.  All Rights Reserved.  Used Under Authorization.",
						"video":                      "QmZQXQXksKmh255bieGC14mev8iUVupjWWQeBFYcRfUj4H/Front/MIGHTY-MORPHIN-POWER-RANGERS_EPIC_WHITE-RANGER_OG_ANIMATED.mp4",
						"cardid":                     float64(103),
						"rarity":                     "Epic",
						"backimg":                    "QmZQXQXksKmh255bieGC14mev8iUVupjWWQeBFYcRfUj4H/Back/back.png",
						"variant":                    "Animated",
						"description":                "White Ranger Digital Pop!",
						"release date":               "August 23, 2022",
						"end user license agreement": "https://digital.funko.com/eula/",
					},
					CreatedAtBlock: "199272794",
					CreatedAtTime:  unixtime.Time(1661095780000),
				},
				TemplateMint:      "1271",
				BackedTokens:      []Token{},
				BurnedByAccount:   "",
				BurnedAtBlock:     "",
				BurnedAtTime:      unixtime.Time(0),
				UpdatedAtBlock:    "217067471",
				UpdatedAtTime:     unixtime.Time(1669998379500),
				TransferedAtBlock: "217067471",
				TransferedAtTime:  unixtime.Time(1669998379500),
				MintedAtBlock:     "199294390",
				MintedAtTime:      unixtime.Time(1661106579000),
				Data: map[string]interface{}{
					"tid":                        float64(53),
					"name":                       "White Ranger",
					"legal":                      "TM & ©2022 SCG Power Rangers LLC and Hasbro. Power Rangers and all related logos, characters, names, and distinctive likenesses thereof are the exclusive property of SCG Power Rangers LLC.  All Rights Reserved.  Used Under Authorization.",
					"video":                      "QmZQXQXksKmh255bieGC14mev8iUVupjWWQeBFYcRfUj4H/Front/MIGHTY-MORPHIN-POWER-RANGERS_EPIC_WHITE-RANGER_OG_ANIMATED.mp4",
					"cardid":                     float64(103),
					"rarity":                     "Epic",
					"backimg":                    "QmZQXQXksKmh255bieGC14mev8iUVupjWWQeBFYcRfUj4H/Back/back.png",
					"variant":                    "Animated",
					"description":                "White Ranger Digital Pop!",
					"release date":               "August 23, 2022",
					"end user license agreement": "https://digital.funko.com/eula/",
				},
				MutableData:   map[string]interface{}{},
				ImmutableData: map[string]interface{}{},
			},
		},
		RecipientAssets:     []Asset{},
		IsSenderContract:    false,
		IsRecipientContract: true,
		UpdatedAtBlock:      "217068480",
		UpdatedAtTime:       unixtime.Time(1669998884000),
		CreatedAtBlock:      "217068480",
		CreatedAtTime:       unixtime.Time(1669998884000),
	}

	client := New(srv.URL)

	a, err := client.GetOffers(OfferRequestParams{})

	require.NoError(t, err)
	assert.Equal(t, 200, a.HTTPStatusCode)
	assert.True(t, a.Success)
	assert.Equal(t, time.Time(time.Date(2020, time.August, 20, 18, 38, 48, 0, time.UTC)), a.QueryTime.Time())
	assert.Equal(t, []Offer{expected}, a.Data)
}

func TestClient_GetOffer(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicassets/v1/offers/103763600", req.URL.String())

		payload := `{
			"success": true,
			"data": {
			  "contract": "atomicassets",
			  "offer_id": "103763600",
			  "sender_name": "sellotronwax",
			  "recipient_name": "atomicmarket",
			  "memo": "sale",
			  "state": 0,
			  "sender_assets": [
				{
				  "contract": "atomicassets",
				  "asset_id": "1099838096677",
				  "owner": "sellotronwax",
				  "is_transferable": true,
				  "is_burnable": true,
				  "collection": {
					"collection_name": "novarallywax",
					"name": "Nova Rally",
					"img": "QmUiQX1BEDWdtQSf1kzJqje3dDrZoBrne7N2xtXgJYXJ4V",
					"author": "novarallywax",
					"allow_notify": true,
					"authorized_accounts": [
					  "novarallywax",
					  "neftyblocksd",
					  "atomicpacksx",
					  "atomicdropsx",
					  "novarallyrac",
					  "novarallypro",
					  "novarallyapp",
					  "blend.nefty",
					  "blenderizerx",
					  "novarallybuy",
					  "neftyblocksp",
					  "iraces.nova",
					  "promo.nova",
					  "bp.nova",
					  "shares.nova"
					],
					"notify_accounts": [
					  "novarallysnk",
					  "novarallyrac",
					  "atomicpacksx",
					  "bp.nova",
					  "scrap.nova",
					  "iraces.nova",
					  "creator.nova"
					],
					"market_fee": 0.08,
					"created_at_block": "119292230",
					"created_at_time": "1621087532500"
				  },
				  "schema": {
					"schema_name": "boost",
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
						"name": "Description",
						"type": "string"
					  }
					],
					"created_at_block": "140846993",
					"created_at_time": "1631874248500"
				  },
				  "template": {
					"template_id": "289842",
					"max_supply": "0",
					"is_transferable": true,
					"is_burnable": true,
					"issued_supply": "117126",
					"immutable_data": {
					  "img": "Qme2DQmTCgE6z2e8QxwyBtgcyRoVSptGk68AGYknG8PbHR/boost.png",
					  "name": "Boost",
					  "Description": "1 x Boost NFT used to try and enhance your time in the Nova Rally! Beware though, it could backfire!"
					},
					"created_at_time": "1631874315000",
					"created_at_block": "140847126"
				  },
				  "mutable_data": {},
				  "immutable_data": {},
				  "template_mint": "116296",
				  "backed_tokens": [],
				  "burned_by_account": null,
				  "burned_at_block": null,
				  "burned_at_time": null,
				  "updated_at_block": "216271853",
				  "updated_at_time": "1669600180500",
				  "transferred_at_block": "216271853",
				  "transferred_at_time": "1669600180500",
				  "minted_at_block": "216271646",
				  "minted_at_time": "1669600077000",
				  "data": {
					"img": "Qme2DQmTCgE6z2e8QxwyBtgcyRoVSptGk68AGYknG8PbHR/boost.png",
					"name": "Boost",
					"Description": "1 x Boost NFT used to try and enhance your time in the Nova Rally! Beware though, it could backfire!"
				  },
				  "name": "Boost"
				}
			  ],
			  "recipient_assets": [],
			  "is_sender_contract": true,
			  "is_recipient_contract": true,
			  "updated_at_block": "217068401",
			  "updated_at_time": "1669998844500",
			  "created_at_block": "217068401",
			  "created_at_time": "1669998844500"
			},
			"query_time": 1427461238500
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	expected := Offer{
		ID:        "103763600",
		Contract:  "atomicassets",
		Sender:    "sellotronwax",
		Recipient: "atomicmarket",
		Memo:      "sale",
		State:     0,
		SenderAssets: []Asset{
			{
				ID:             "1099838096677",
				Name:           "Boost",
				Contract:       "atomicassets",
				Owner:          "sellotronwax",
				IsTransferable: true,
				IsBurnable:     true,
				Collection: Collection{
					Name:           "Nova Rally",
					CollectionName: "novarallywax",
					Image:          "QmUiQX1BEDWdtQSf1kzJqje3dDrZoBrne7N2xtXgJYXJ4V",
					Author:         "novarallywax",
					AllowNotify:    true,
					AuthorizedAccounts: []string{
						"novarallywax",
						"neftyblocksd",
						"atomicpacksx",
						"atomicdropsx",
						"novarallyrac",
						"novarallypro",
						"novarallyapp",
						"blend.nefty",
						"blenderizerx",
						"novarallybuy",
						"neftyblocksp",
						"iraces.nova",
						"promo.nova",
						"bp.nova",
						"shares.nova",
					},
					NotifyAccounts: []string{
						"novarallysnk",
						"novarallyrac",
						"atomicpacksx",
						"bp.nova",
						"scrap.nova",
						"iraces.nova",
						"creator.nova",
					},
					MarketFee:      0.08,
					CreatedAtBlock: "119292230",
					CreatedAtTime:  unixtime.Time(1621087532500),
				},
				Schema: InlineSchema{
					Name: "boost",
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
							Name: "Description",
							Type: "string",
						},
					},
					CreatedAtBlock: "140846993",
					CreatedAtTime:  unixtime.Time(1631874248500),
				},
				Template: Template{
					ID:             "289842",
					MaxSupply:      "0",
					IsTransferable: true,
					IsBurnable:     true,
					IssuedSupply:   "117126",
					ImmutableData: map[string]interface{}{
						"img":         "Qme2DQmTCgE6z2e8QxwyBtgcyRoVSptGk68AGYknG8PbHR/boost.png",
						"name":        "Boost",
						"Description": "1 x Boost NFT used to try and enhance your time in the Nova Rally! Beware though, it could backfire!",
					},
					CreatedAtBlock: "140847126",
					CreatedAtTime:  unixtime.Time(1631874315000),
				},
				TemplateMint:      "116296",
				BackedTokens:      []Token{},
				BurnedByAccount:   "",
				BurnedAtBlock:     "",
				BurnedAtTime:      unixtime.Time(0),
				UpdatedAtBlock:    "216271853",
				UpdatedAtTime:     unixtime.Time(1669600180500),
				TransferedAtBlock: "216271853",
				TransferedAtTime:  unixtime.Time(1669600180500),
				MintedAtBlock:     "216271646",
				MintedAtTime:      unixtime.Time(1669600077000),
				Data: map[string]interface{}{
					"img":         "Qme2DQmTCgE6z2e8QxwyBtgcyRoVSptGk68AGYknG8PbHR/boost.png",
					"name":        "Boost",
					"Description": "1 x Boost NFT used to try and enhance your time in the Nova Rally! Beware though, it could backfire!",
				},
				MutableData:   map[string]interface{}{},
				ImmutableData: map[string]interface{}{},
			},
		},
		RecipientAssets:     []Asset{},
		IsSenderContract:    true,
		IsRecipientContract: true,
		UpdatedAtBlock:      "217068401",
		UpdatedAtTime:       unixtime.Time(1669998844500),
		CreatedAtBlock:      "217068401",
		CreatedAtTime:       unixtime.Time(1669998844500),
	}

	client := New(srv.URL)

	a, err := client.GetOffer("103763600")

	require.NoError(t, err)
	assert.Equal(t, 200, a.HTTPStatusCode)
	assert.True(t, a.Success)
	assert.Equal(t, time.Time(time.Date(2015, time.March, 27, 13, 0, 38, 500, time.UTC)), a.QueryTime.Time())
	assert.Equal(t, expected, a.Data)
}

func TestClient_GetOfferLog(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicassets/v1/offers/103763600/logs", req.URL.String())

		payload := `{
		"success": true,
		"data": [
			{
				"log_id": "70487516527",
				"name": "lognewoffer",
				"data": {},
				"txid": "ca3dd523426ca7d410e60bbb49eb76675374b107f9b25b908bf9fe70fa3fc640",
				"created_at_block": "217068401",
				"created_at_time": "1669998844500"
			}
		],
		"query_time": 1697235391000
	  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	expected := Log{
		ID:             "70487516527",
		Name:           "lognewoffer",
		TxID:           "ca3dd523426ca7d410e60bbb49eb76675374b107f9b25b908bf9fe70fa3fc640",
		Data:           map[string]interface{}{},
		CreatedAtBlock: "217068401",
		CreatedAtTime:  unixtime.Time(1669998844500),
	}

	client := New(srv.URL)

	a, err := client.GetOfferLog("103763600", LogRequestParams{})

	require.NoError(t, err)
	assert.Equal(t, 200, a.HTTPStatusCode)
	assert.True(t, a.Success)
	assert.Equal(t, time.Time(time.Date(2023, time.October, 13, 22, 16, 31, 0, time.UTC)), a.QueryTime.Time())
	assert.Equal(t, []Log{expected}, a.Data)
}
