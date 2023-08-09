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

func TestClient_GetTemplates(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicassets/v1/templates", req.URL.String())

		payload := `{
			"success": true,
			"data": [
			  {
				"contract": "atomicassets",
				"template_id": "637383",
				"is_transferable": true,
				"is_burnable": true,
				"issued_supply": "77",
				"max_supply": "0",
				"collection": {
				  "collection_name": "mindmastrart",
				  "name": "MindMaster Art",
				  "img": "QmSMgwd6mGVKqmwiAj6iCyZceehiLgiVRaS4cqmFu6qBCW",
				  "author": "mindmaster12",
				  "allow_notify": true,
				  "authorized_accounts": [
					"mindmaster12",
					"atomicpacksx",
					"neftyblocksd",
					"neftyblocksp",
					"blend.nefty",
					"nfthivedrops"
				  ],
				  "notify_accounts": [],
				  "market_fee": 0.09,
				  "created_at_block": "176839276",
				  "created_at_time": "1649876948500"
				},
				"schema": {
				  "schema_name": "wam.collabs",
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
					},
					{
					  "name": "description",
					  "type": "string"
					},
					{
					  "name": "rarity",
					  "type": "string"
					}
				  ],
				  "created_at_block": "197880498",
				  "created_at_time": "1660399568000"
				},
				"immutable_data": {
				  "img": "QmSDfu2HFrfeBgPR74a76GKuFLThtRV5AJwmb8WDnSCmHe",
				  "name": "We're All Monsters - Advent Calender 2022: Geeked For Xmas",
				  "rarity": "Collab",
				  "description": "The Geeked for Christmas PFP was brought to you through 25 Days of NFTs, a mega collab featuring art from 40 Artists, organized by Crackers (Meet The Artist) @Crackers832 on Twitter! Blends available here https://neftyblocks.com/c/mindmastrart/blends"
				},
				"created_at_time": "1669984452500",
				"created_at_block": "217039631",
				"name": "We're All Monsters - Advent Calender 2022: Geeked For Xmas"
			  }
			],
			"query_time": 1212026652500
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	expected := Template{
		ID:             "637383",
		Contract:       "atomicassets",
		MaxSupply:      "0",
		IssuedSupply:   "77",
		IsTransferable: true,
		IsBurnable:     true,
		Schema: InlineSchema{
			Name: "wam.collabs",
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
				{
					Name: "description",
					Type: "string",
				},
				{
					Name: "rarity",
					Type: "string",
				},
			},
			CreatedAtBlock: "197880498",
			CreatedAtTime:  unixtime.Time(1660399568000),
		},
		Collection: Collection{
			Name:           "MindMaster Art",
			CollectionName: "mindmastrart",
			Contract:       "",
			Image:          "QmSMgwd6mGVKqmwiAj6iCyZceehiLgiVRaS4cqmFu6qBCW",
			Author:         "mindmaster12",
			AllowNotify:    true,
			AuthorizedAccounts: []string{
				"mindmaster12",
				"atomicpacksx",
				"neftyblocksd",
				"neftyblocksp",
				"blend.nefty",
				"nfthivedrops",
			},
			NotifyAccounts: []string{},
			MarketFee:      0.09,
			CreatedAtBlock: "176839276",
			CreatedAtTime:  unixtime.Time(1649876948500),
		},
		ImmutableData: map[string]interface{}{
			"img":         "QmSDfu2HFrfeBgPR74a76GKuFLThtRV5AJwmb8WDnSCmHe",
			"name":        "We're All Monsters - Advent Calender 2022: Geeked For Xmas",
			"rarity":      "Collab",
			"description": "The Geeked for Christmas PFP was brought to you through 25 Days of NFTs, a mega collab featuring art from 40 Artists, organized by Crackers (Meet The Artist) @Crackers832 on Twitter! Blends available here https://neftyblocks.com/c/mindmastrart/blends",
		},
		CreatedAtBlock: "217039631",
		CreatedAtTime:  unixtime.Time(1669984452500),
	}

	client := New(srv.URL)

	a, err := client.GetTemplates(TemplateRequestParams{})

	require.NoError(t, err)
	assert.Equal(t, 200, a.HTTPStatusCode)
	assert.True(t, a.Success)
	assert.Equal(t, time.Date(2008, time.May, 29, 2, 4, 12, int(time.Millisecond)*500, time.UTC), a.QueryTime.Time())
	assert.Equal(t, []Template{expected}, a.Data)
}

func TestClient_GetTemplate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicassets/v1/templates/mindmastrart/637383", req.URL.String())

		payload := `{
			"success": true,
			"data": {
			  "contract": "atomicassets",
			  "template_id": "637383",
			  "is_transferable": true,
			  "is_burnable": true,
			  "issued_supply": "77",
			  "max_supply": "0",
			  "collection": {
				"collection_name": "mindmastrart",
				"name": "MindMaster Art",
				"img": "QmSMgwd6mGVKqmwiAj6iCyZceehiLgiVRaS4cqmFu6qBCW",
				"author": "mindmaster12",
				"allow_notify": true,
				"authorized_accounts": [
				  "mindmaster12",
				  "atomicpacksx",
				  "neftyblocksd",
				  "neftyblocksp",
				  "blend.nefty",
				  "nfthivedrops"
				],
				"notify_accounts": [],
				"market_fee": 0.09,
				"created_at_block": "176839276",
				"created_at_time": "1649876948500"
			  },
			  "schema": {
				"schema_name": "wam.collabs",
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
				  },
				  {
					"name": "description",
					"type": "string"
				  },
				  {
					"name": "rarity",
					"type": "string"
				  }
				],
				"created_at_block": "197880498",
				"created_at_time": "1660399568000"
			  },
			  "immutable_data": {
				"img": "QmSDfu2HFrfeBgPR74a76GKuFLThtRV5AJwmb8WDnSCmHe",
				"name": "We're All Monsters - Advent Calender 2022: Geeked For Xmas",
				"rarity": "Collab",
				"description": "The Geeked for Christmas PFP was brought to you through 25 Days of NFTs, a mega collab featuring art from 40 Artists, organized by Crackers (Meet The Artist) @Crackers832 on Twitter! Blends available here https://neftyblocks.com/c/mindmastrart/blends"
			  },
			  "created_at_time": "1669984452500",
			  "created_at_block": "217039631",
			  "name": "We're All Monsters - Advent Calender 2022: Geeked For Xmas"
			},
			"query_time": 1053342787500
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	expected := Template{
		ID:             "637383",
		Contract:       "atomicassets",
		MaxSupply:      "0",
		IssuedSupply:   "77",
		IsTransferable: true,
		IsBurnable:     true,
		Schema: InlineSchema{
			Name: "wam.collabs",
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
				{
					Name: "description",
					Type: "string",
				},
				{
					Name: "rarity",
					Type: "string",
				},
			},
			CreatedAtBlock: "197880498",
			CreatedAtTime:  unixtime.Time(1660399568000),
		},
		Collection: Collection{
			Name:           "MindMaster Art",
			CollectionName: "mindmastrart",
			Contract:       "",
			Image:          "QmSMgwd6mGVKqmwiAj6iCyZceehiLgiVRaS4cqmFu6qBCW",
			Author:         "mindmaster12",
			AllowNotify:    true,
			AuthorizedAccounts: []string{
				"mindmaster12",
				"atomicpacksx",
				"neftyblocksd",
				"neftyblocksp",
				"blend.nefty",
				"nfthivedrops",
			},
			NotifyAccounts: []string{},
			MarketFee:      0.09,
			CreatedAtBlock: "176839276",
			CreatedAtTime:  unixtime.Time(1649876948500),
		},
		ImmutableData: map[string]interface{}{
			"img":         "QmSDfu2HFrfeBgPR74a76GKuFLThtRV5AJwmb8WDnSCmHe",
			"name":        "We're All Monsters - Advent Calender 2022: Geeked For Xmas",
			"rarity":      "Collab",
			"description": "The Geeked for Christmas PFP was brought to you through 25 Days of NFTs, a mega collab featuring art from 40 Artists, organized by Crackers (Meet The Artist) @Crackers832 on Twitter! Blends available here https://neftyblocks.com/c/mindmastrart/blends",
		},
		CreatedAtBlock: "217039631",
		CreatedAtTime:  unixtime.Time(1669984452500),
	}

	client := New(srv.URL)

	a, err := client.GetTemplate("mindmastrart", "637383")

	require.NoError(t, err)
	assert.Equal(t, 200, a.HTTPStatusCode)
	assert.True(t, a.Success)
	assert.Equal(t, time.Date(2003, time.May, 19, 11, 13, 7, int(time.Millisecond)*500, time.UTC), a.QueryTime.Time())
	assert.Equal(t, expected, a.Data)
}

func TestClient_GetTemplateStats(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicassets/v1/templates/mindmastrart/637383/stats", req.URL.String())

		payload := `{
		"success": true,
		"data": {
		  "assets": "77",
		  "burned": "0"
		},
		"query_time": 1117111695000
	  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	expected := TemplateStats{
		Assets: "77",
		Burned: "0",
	}

	client := New(srv.URL)

	a, err := client.GetTemplateStats("mindmastrart", "637383")

	require.NoError(t, err)
	assert.Equal(t, 200, a.HTTPStatusCode)
	assert.True(t, a.Success)
	assert.Equal(t, time.Date(2005, time.May, 26, 12, 48, 15, 0, time.UTC), a.QueryTime.Time())
	assert.Equal(t, expected, a.Data)
}
