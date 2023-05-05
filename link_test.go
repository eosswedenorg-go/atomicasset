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

var link1 = Link{
	ID:             "1693600",
	ToolsContract:  "atomictoolsx",
	AssetsContract: "atomicassets",
	Creator:        "cmcdrops4all",
	Claimer:        "",
	State:          LinkStateCreated,
	Memo:           "advent of code day 17 winner 1 (msc7#2934) WARNING: Tip bot claimlinks may be cancelled 10 days after issuance.",
	PublicKey:      "PUB_K1_61Sh5Roq63vxrWL7Xn1G5ntFmAUqJJa3Uaj4NhuLvVtjApLxdq",
	CreatedAtBlock: "220155650",
	CreatedAtTime:  unixtime.Time(1671543988000),
	UpdatedAtBlock: "220155650",
	UpdatedAtTime:  unixtime.Time(1671543988000),
	Assets:         []Asset{link_asset1},
}

var link_asset1 Asset = Asset{
	ID:             "1099839311748",
	Name:           "x-mas 2022",
	Contract:       "atomicassets",
	Owner:          "atomictoolsx",
	IsTransferable: true,
	IsBurnable:     true,
	Collection: Collection{
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
		CreatedAtBlock: "66677714",
		CreatedAtTime:  unixtime.Time(1594768100000),
	},
	Schema: InlineSchema{
		Name: "monkeystacks",
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
				Type: "ipfs",
			},
			{
				Name: "card_id",
				Type: "uint64",
			},
			{
				Name: "generation",
				Type: "uint64",
			},
			{
				Name: "rarity",
				Type: "string",
			},
			{
				Name: "credits",
				Type: "string",
			},
			{
				Name: "text",
				Type: "string",
			},
			{
				Name: "memability",
				Type: "uint64",
			},
			{
				Name: "stonks",
				Type: "uint64",
			},
			{
				Name: "pumpamentality",
				Type: "uint64",
			},
			{
				Name: "ape strength",
				Type: "uint64",
			},
			{
				Name: "chimp speed",
				Type: "uint64",
			},
			{
				Name: "potassium diet",
				Type: "uint64",
			},
			{
				Name: "sneakiness",
				Type: "uint64",
			},
			{
				Name: "awesomeness",
				Type: "uint64",
			},
			{
				Name: "powwah",
				Type: "uint64",
			},
		},
		CreatedAtBlock: "148096441",
		CreatedAtTime:  unixtime.Time(1635501352500),
	},
	Template: Template{
		ID:             "637934",
		MaxSupply:      "1919",
		IsTransferable: true,
		IsBurnable:     true,
		IssuedSupply:   "1919",
		ImmutableData: map[string]interface{}{
			"img":            "QmZBunzJYQvLCE7nQo28JGVhjtMZyKrYkG1GoiLLpEgHiw",
			"name":           "x-mas 2022",
			"text":           "Distributed to Xmas 2022 SantaBot claimers at Banano and cryptomonKey discords and generally via NFT tipbots, monKeyslots, monKeymining, monKeymiles and monKeytrains.",
			"powwah":         "6",
			"rarity":         "Uncommon",
			"stonks":         "19",
			"backimg":        "QmYDxr9iCEzi2TPzwt54B2ADMXAZkQnqjjVnXdnJoyRSMw",
			"card_id":        "116",
			"credits":        "Sem Brys a.k.a. Mantichore",
			"generation":     "2",
			"memability":     "17",
			"sneakiness":     "16",
			"awesomeness":    "19",
			"chimp speed":    "13",
			"ape strength":   "11",
			"potassium diet": "4",
			"pumpamentality": "9",
		},
		CreatedAtBlock: "217567668",
		CreatedAtTime:  unixtime.Time(1670248713000),
	},
	MutableData:       map[string]interface{}{},
	ImmutableData:     map[string]interface{}{},
	TemplateMint:      "550",
	BackedTokens:      []Token{},
	BurnedByAccount:   "",
	BurnedAtBlock:     "",
	BurnedAtTime:      unixtime.Time(0),
	UpdatedAtBlock:    "220155650",
	UpdatedAtTime:     unixtime.Time(1671543988000),
	TransferedAtBlock: "220155650",
	TransferedAtTime:  unixtime.Time(1671543988000),
	MintedAtBlock:     "217574806",
	MintedAtTime:      unixtime.Time(1670252287500),
	Data: map[string]interface{}{
		"img":            "QmZBunzJYQvLCE7nQo28JGVhjtMZyKrYkG1GoiLLpEgHiw",
		"name":           "x-mas 2022",
		"text":           "Distributed to Xmas 2022 SantaBot claimers at Banano and cryptomonKey discords and generally via NFT tipbots, monKeyslots, monKeymining, monKeymiles and monKeytrains.",
		"powwah":         "6",
		"rarity":         "Uncommon",
		"stonks":         "19",
		"backimg":        "QmYDxr9iCEzi2TPzwt54B2ADMXAZkQnqjjVnXdnJoyRSMw",
		"card_id":        "116",
		"credits":        "Sem Brys a.k.a. Mantichore",
		"generation":     "2",
		"memability":     "17",
		"sneakiness":     "16",
		"awesomeness":    "19",
		"chimp speed":    "13",
		"ape strength":   "11",
		"potassium diet": "4",
		"pumpamentality": "9",
	},
}

func TestGetLink(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomictools/v1/links/1693600", req.URL.String())

		payload := `{
			"success": true,
			"data": {
			  "tools_contract": "atomictoolsx",
			  "link_id": "1693600",
			  "assets_contract": "atomicassets",
			  "creator": "cmcdrops4all",
			  "claimer": null,
			  "state": 1,
			  "memo": "advent of code day 17 winner 1 (msc7#2934) WARNING: Tip bot claimlinks may be cancelled 10 days after issuance.",
			  "assets": [
				{
				  "contract": "atomicassets",
				  "asset_id": "1099839311748",
				  "owner": "atomictoolsx",
				  "is_transferable": true,
				  "is_burnable": true,
				  "collection": {
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
					"created_at_block": "66677714",
					"created_at_time": "1594768100000"
				  },
				  "schema": {
					"schema_name": "monkeystacks",
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
						"type": "ipfs"
					  },
					  {
						"name": "card_id",
						"type": "uint64"
					  },
					  {
						"name": "generation",
						"type": "uint64"
					  },
					  {
						"name": "rarity",
						"type": "string"
					  },
					  {
						"name": "credits",
						"type": "string"
					  },
					  {
						"name": "text",
						"type": "string"
					  },
					  {
						"name": "memability",
						"type": "uint64"
					  },
					  {
						"name": "stonks",
						"type": "uint64"
					  },
					  {
						"name": "pumpamentality",
						"type": "uint64"
					  },
					  {
						"name": "ape strength",
						"type": "uint64"
					  },
					  {
						"name": "chimp speed",
						"type": "uint64"
					  },
					  {
						"name": "potassium diet",
						"type": "uint64"
					  },
					  {
						"name": "sneakiness",
						"type": "uint64"
					  },
					  {
						"name": "awesomeness",
						"type": "uint64"
					  },
					  {
						"name": "powwah",
						"type": "uint64"
					  }
					],
					"created_at_block": "148096441",
					"created_at_time": "1635501352500"
				  },
				  "template": {
					"template_id": "637934",
					"max_supply": "1919",
					"is_transferable": true,
					"is_burnable": true,
					"issued_supply": "1919",
					"immutable_data": {
					  "img": "QmZBunzJYQvLCE7nQo28JGVhjtMZyKrYkG1GoiLLpEgHiw",
					  "name": "x-mas 2022",
					  "text": "Distributed to Xmas 2022 SantaBot claimers at Banano and cryptomonKey discords and generally via NFT tipbots, monKeyslots, monKeymining, monKeymiles and monKeytrains.",
					  "powwah": "6",
					  "rarity": "Uncommon",
					  "stonks": "19",
					  "backimg": "QmYDxr9iCEzi2TPzwt54B2ADMXAZkQnqjjVnXdnJoyRSMw",
					  "card_id": "116",
					  "credits": "Sem Brys a.k.a. Mantichore",
					  "generation": "2",
					  "memability": "17",
					  "sneakiness": "16",
					  "awesomeness": "19",
					  "chimp speed": "13",
					  "ape strength": "11",
					  "potassium diet": "4",
					  "pumpamentality": "9"
					},
					"created_at_time": "1670248713000",
					"created_at_block": "217567668"
				  },
				  "mutable_data": {},
				  "immutable_data": {},
				  "template_mint": "550",
				  "backed_tokens": [],
				  "burned_by_account": null,
				  "burned_at_block": null,
				  "burned_at_time": null,
				  "updated_at_block": "220155650",
				  "updated_at_time": "1671543988000",
				  "transferred_at_block": "220155650",
				  "transferred_at_time": "1671543988000",
				  "minted_at_block": "217574806",
				  "minted_at_time": "1670252287500",
				  "data": {
					"img": "QmZBunzJYQvLCE7nQo28JGVhjtMZyKrYkG1GoiLLpEgHiw",
					"name": "x-mas 2022",
					"text": "Distributed to Xmas 2022 SantaBot claimers at Banano and cryptomonKey discords and generally via NFT tipbots, monKeyslots, monKeymining, monKeymiles and monKeytrains.",
					"powwah": "6",
					"rarity": "Uncommon",
					"stonks": "19",
					"backimg": "QmYDxr9iCEzi2TPzwt54B2ADMXAZkQnqjjVnXdnJoyRSMw",
					"card_id": "116",
					"credits": "Sem Brys a.k.a. Mantichore",
					"generation": "2",
					"memability": "17",
					"sneakiness": "16",
					"awesomeness": "19",
					"chimp speed": "13",
					"ape strength": "11",
					"potassium diet": "4",
					"pumpamentality": "9"
				  },
				  "name": "x-mas 2022"
				}
			  ],
			  "created_at_block": "220155650",
			  "created_at_time": "1671543988000",
			  "updated_at_block": "220155650",
			  "updated_at_time": "1671543988000",
			  "public_key": "PUB_K1_61Sh5Roq63vxrWL7Xn1G5ntFmAUqJJa3Uaj4NhuLvVtjApLxdq"
			},
			"query_time": 1720783862000
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetLink(1693600)

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2024, time.July, 12, 11, 31, 2, 0, time.UTC), res.QueryTime.Time())

	assert.Equal(t, link1, res.Data)
}

func TestGetLinkLogs(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomictools/v1/links/1693600/logs?limit=2", req.URL.String())

		payload := `{
			"success": true,
			"data": [
			  {
				"log_id": "71768004842",
				"name": "claimlink",
				"data": {
				  "claimer_signature": "1693600"
				},
				"txid": "8c9ef0b05888d93376b7414ace818178d0d988a1486307a4d13ac8a442996a3e",
				"created_at_block": "220187047",
				"created_at_time": "1671559689500"
			  },
			  {
				"log_id": "71758057739",
				"name": "loglinkstart",
				"data": {},
				"txid": "98edbca5689aef78a1ff31dd6768c0cb8b6cb8db8b0696582b3667f5427d8537",
				"created_at_block": "220155650",
				"created_at_time": "1671543988000"
			  }
			],
			"query_time": 1820816785000
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetLinkLogs(1693600, LogRequestParams{Limit: 2})

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2027, time.September, 13, 6, 26, 25, 0, time.UTC), res.QueryTime.Time())

	expected := []Log{
		{
			ID:   "71768004842",
			Name: "claimlink",
			Data: map[string]interface{}{
				"claimer_signature": "1693600",
			},
			TxID:           "8c9ef0b05888d93376b7414ace818178d0d988a1486307a4d13ac8a442996a3e",
			CreatedAtBlock: "220187047",
			CreatedAtTime:  unixtime.Time(1671559689500),
		},
		{
			ID:             "71758057739",
			Name:           "loglinkstart",
			Data:           map[string]interface{}{},
			TxID:           "98edbca5689aef78a1ff31dd6768c0cb8b6cb8db8b0696582b3667f5427d8537",
			CreatedAtBlock: "220155650",
			CreatedAtTime:  unixtime.Time(1671543988000),
		},
	}

	assert.Equal(t, expected, res.Data)
}

func TestGetLinks(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomictools/v1/links?limit=1", req.URL.String())

		payload := `{
			"success": true,
			"data": [
			  {
				"tools_contract": "atomictoolsx",
				"link_id": "1693600",
				"assets_contract": "atomicassets",
				"creator": "cmcdrops4all",
				"claimer": null,
				"state": 1,
				"memo": "advent of code day 17 winner 1 (msc7#2934) WARNING: Tip bot claimlinks may be cancelled 10 days after issuance.",
				"assets": [
					{
					"contract": "atomicassets",
					"asset_id": "1099839311748",
					"owner": "atomictoolsx",
					"is_transferable": true,
					"is_burnable": true,
					"collection": {
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
						"created_at_block": "66677714",
						"created_at_time": "1594768100000"
					},
					"schema": {
						"schema_name": "monkeystacks",
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
							"type": "ipfs"
						},
						{
							"name": "card_id",
							"type": "uint64"
						},
						{
							"name": "generation",
							"type": "uint64"
						},
						{
							"name": "rarity",
							"type": "string"
						},
						{
							"name": "credits",
							"type": "string"
						},
						{
							"name": "text",
							"type": "string"
						},
						{
							"name": "memability",
							"type": "uint64"
						},
						{
							"name": "stonks",
							"type": "uint64"
						},
						{
							"name": "pumpamentality",
							"type": "uint64"
						},
						{
							"name": "ape strength",
							"type": "uint64"
						},
						{
							"name": "chimp speed",
							"type": "uint64"
						},
						{
							"name": "potassium diet",
							"type": "uint64"
						},
						{
							"name": "sneakiness",
							"type": "uint64"
						},
						{
							"name": "awesomeness",
							"type": "uint64"
						},
						{
							"name": "powwah",
							"type": "uint64"
						}
						],
						"created_at_block": "148096441",
						"created_at_time": "1635501352500"
					},
					"template": {
						"template_id": "637934",
						"max_supply": "1919",
						"is_transferable": true,
						"is_burnable": true,
						"issued_supply": "1919",
						"immutable_data": {
						"img": "QmZBunzJYQvLCE7nQo28JGVhjtMZyKrYkG1GoiLLpEgHiw",
						"name": "x-mas 2022",
						"text": "Distributed to Xmas 2022 SantaBot claimers at Banano and cryptomonKey discords and generally via NFT tipbots, monKeyslots, monKeymining, monKeymiles and monKeytrains.",
						"powwah": "6",
						"rarity": "Uncommon",
						"stonks": "19",
						"backimg": "QmYDxr9iCEzi2TPzwt54B2ADMXAZkQnqjjVnXdnJoyRSMw",
						"card_id": "116",
						"credits": "Sem Brys a.k.a. Mantichore",
						"generation": "2",
						"memability": "17",
						"sneakiness": "16",
						"awesomeness": "19",
						"chimp speed": "13",
						"ape strength": "11",
						"potassium diet": "4",
						"pumpamentality": "9"
						},
						"created_at_time": "1670248713000",
						"created_at_block": "217567668"
					},
					"mutable_data": {},
					"immutable_data": {},
					"template_mint": "550",
					"backed_tokens": [],
					"burned_by_account": null,
					"burned_at_block": null,
					"burned_at_time": null,
					"updated_at_block": "220155650",
					"updated_at_time": "1671543988000",
					"transferred_at_block": "220155650",
					"transferred_at_time": "1671543988000",
					"minted_at_block": "217574806",
					"minted_at_time": "1670252287500",
					"data": {
						"img": "QmZBunzJYQvLCE7nQo28JGVhjtMZyKrYkG1GoiLLpEgHiw",
						"name": "x-mas 2022",
						"text": "Distributed to Xmas 2022 SantaBot claimers at Banano and cryptomonKey discords and generally via NFT tipbots, monKeyslots, monKeymining, monKeymiles and monKeytrains.",
						"powwah": "6",
						"rarity": "Uncommon",
						"stonks": "19",
						"backimg": "QmYDxr9iCEzi2TPzwt54B2ADMXAZkQnqjjVnXdnJoyRSMw",
						"card_id": "116",
						"credits": "Sem Brys a.k.a. Mantichore",
						"generation": "2",
						"memability": "17",
						"sneakiness": "16",
						"awesomeness": "19",
						"chimp speed": "13",
						"ape strength": "11",
						"potassium diet": "4",
						"pumpamentality": "9"
					},
					"name": "x-mas 2022"
					}
				],
				"created_at_block": "220155650",
				"created_at_time": "1671543988000",
				"updated_at_block": "220155650",
				"updated_at_time": "1671543988000",
				"public_key": "PUB_K1_61Sh5Roq63vxrWL7Xn1G5ntFmAUqJJa3Uaj4NhuLvVtjApLxdq"
				}
			],
			"query_time": 1720783862000
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetLinks(LinkRequestParams{Limit: 1})

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2024, time.July, 12, 11, 31, 2, 0, time.UTC), res.QueryTime.Time())

	assert.Equal(t, []Link{link1}, res.Data)
}
