package atomicasset

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_GetTransfers(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicassets/v1/transfers", req.URL.String())

		payload := `{
			"success": true,
			"data": [
			  {
				"transfer_id": "70580694247",
				"contract": "atomicassets",
				"sender_name": "starshipgame",
				"recipient_name": "4awaxaccount",
				"memo": "Starship part replaced",
				"txid": "b036bfecda89eae9d325dc7f895e4d248f0513b032973bd1cfba052a151118b6",
				"assets": [
				  {
					"contract": "atomicassets",
					"asset_id": "1099801566489",
					"owner": "4awaxaccount",
					"is_transferable": true,
					"is_burnable": true,
					"collection": {
					  "collection_name": "starshipnfts",
					  "name": "Starship NFT Game",
					  "img": "QmYWS19pX4nfsq8eubJaMtVWwk6at6Ar4TyujvYXoSaiDc",
					  "author": "starshipnfts",
					  "allow_notify": true,
					  "authorized_accounts": [
						"starshipnfts",
						"atomicdropsx",
						"atomicpacksx",
						"neftyblocksd",
						"starshipgame",
						"blenderizerx",
						"blend.nefty"
					  ],
					  "notify_accounts": [
						"starshipgame"
					  ],
					  "market_fee": 0.08,
					  "created_at_block": "138805797",
					  "created_at_time": "1630851624500"
					},
					"schema": {
					  "schema_name": "component",
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
						  "name": "rarity",
						  "type": "string"
						},
						{
						  "name": "type",
						  "type": "string"
						},
						{
						  "name": "rdata",
						  "type": "uint64"
						},
						{
						  "name": "tdata",
						  "type": "uint64"
						},
						{
						  "name": "video",
						  "type": "ipfs"
						}
					  ],
					  "created_at_block": "144015964",
					  "created_at_time": "1633460071000"
					},
					"template": {
					  "template_id": "536289",
					  "max_supply": "0",
					  "is_transferable": true,
					  "is_burnable": true,
					  "issued_supply": "2458",
					  "immutable_data": {
						"img": "QmSho935ezZYpK2jMd6oNxLvyu5uQmEFPKeSnxkijXQjr2",
						"name": "Kyanite Resonator",
						"type": "Kyanite Resonator",
						"rdata": "3",
						"tdata": "11",
						"rarity": "Epic"
					  },
					  "created_at_time": "1656179852000",
					  "created_at_block": "189441618"
					},
					"mutable_data": {},
					"immutable_data": {},
					"template_mint": "667",
					"backed_tokens": [],
					"burned_by_account": null,
					"burned_at_block": null,
					"burned_at_time": null,
					"updated_at_block": "217211017",
					"updated_at_time": "1670070209000",
					"transferred_at_block": "217211017",
					"transferred_at_time": "1670070209000",
					"minted_at_block": "191339821",
					"minted_at_time": "1657128962500",
					"data": {
					  "img": "QmSho935ezZYpK2jMd6oNxLvyu5uQmEFPKeSnxkijXQjr2",
					  "name": "Kyanite Resonator",
					  "type": "Kyanite Resonator",
					  "rdata": "3",
					  "tdata": "11",
					  "rarity": "Epic"
					},
					"name": "Kyanite Resonator"
				  }
				],
				"created_at_block": "217211017",
				"created_at_time": "1670070209000"
			  }
			],
			"query_time": 1590434103000
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	expected := Transfer{
		ID:        "70580694247",
		Contract:  "atomicassets",
		Sender:    "starshipgame",
		Recipient: "4awaxaccount",
		Memo:      "Starship part replaced",
		// TxID
		Assets: []Asset{
			{
				ID:             "1099801566489",
				Name:           "Kyanite Resonator",
				Contract:       "atomicassets",
				Owner:          "4awaxaccount",
				IsTransferable: true,
				IsBurnable:     true,
				Collection: Collection{
					Name:           "Starship NFT Game",
					CollectionName: "starshipnfts",
					Image:          "QmYWS19pX4nfsq8eubJaMtVWwk6at6Ar4TyujvYXoSaiDc",
					Author:         "starshipnfts",
					AllowNotify:    true,
					AuthorizedAccounts: []string{
						"starshipnfts",
						"atomicdropsx",
						"atomicpacksx",
						"neftyblocksd",
						"starshipgame",
						"blenderizerx",
						"blend.nefty",
					},
					NotifyAccounts: []string{
						"starshipgame",
					},
					MarketFee:      0.08,
					CreatedAtBlock: "138805797",
					CreatedAtTime:  UnixTime(1630851624500),
				},
				Schema: InlineSchema{
					Name: "component",
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
							Name: "rarity",
							Type: "string",
						},
						{
							Name: "type",
							Type: "string",
						},
						{
							Name: "rdata",
							Type: "uint64",
						},
						{
							Name: "tdata",
							Type: "uint64",
						},
						{
							Name: "video",
							Type: "ipfs",
						},
					},
					CreatedAtBlock: "144015964",
					CreatedAtTime:  UnixTime(1633460071000),
				},
				Template: Template{
					ID:             "536289",
					MaxSupply:      "0",
					IsTransferable: true,
					IsBurnable:     true,
					IssuedSupply:   "2458",
					ImmutableData: map[string]interface{}{
						"img":    "QmSho935ezZYpK2jMd6oNxLvyu5uQmEFPKeSnxkijXQjr2",
						"name":   "Kyanite Resonator",
						"type":   "Kyanite Resonator",
						"rdata":  "3",
						"tdata":  "11",
						"rarity": "Epic",
					},
					CreatedAtBlock: "189441618",
					CreatedAtTime:  UnixTime(1656179852000),
				},
				MutableData:       map[string]interface{}{},
				ImmutableData:     map[string]interface{}{},
				TemplateMint:      "667",
				BackedTokens:      []Token{},
				BurnedByAccount:   "",
				BurnedAtBlock:     "",
				BurnedAtTime:      "",
				UpdatedAtBlock:    "217211017",
				UpdatedAtTime:     "1670070209000",
				TransferedAtBlock: "217211017",
				TransferedAtTime:  "1670070209000",
				MintedAtBlock:     "191339821",
				MintedAtTime:      "1657128962500",
				Data: map[string]interface{}{
					"img":    "QmSho935ezZYpK2jMd6oNxLvyu5uQmEFPKeSnxkijXQjr2",
					"name":   "Kyanite Resonator",
					"type":   "Kyanite Resonator",
					"rdata":  "3",
					"tdata":  "11",
					"rarity": "Epic",
				},
			},
		},
		CreatedAtBlock: "217211017",
		CreatedAtTime:  UnixTime(1670070209000),
	}

	client := New(srv.URL)

	a, err := client.GetTransfers(TransferRequestParams{})

	require.NoError(t, err)
	assert.Equal(t, 200, a.HTTPStatusCode)
	assert.True(t, a.Success)
	assert.Equal(t, time.Time(time.Date(2020, time.May, 25, 19, 15, 3, 0, time.UTC)), a.QueryTime.Time())
	assert.Equal(t, []Transfer{expected}, a.Data)
}
