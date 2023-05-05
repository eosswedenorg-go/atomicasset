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
			Image:          "QmbKHBLk9VTHfLhxy7LJEFUjEBJLB7gGSaBL2UeSRn6oMA",
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
			CreatedAtTime:  unixtime.Time(1669204520000),
		},
		{
			Name:           "Filipino Digital Arts",
			Image:          "QmbukpNarUzBPfTsWsLNQzKhiowviTBenQS9YgdV8nqtj7",
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
			CreatedAtTime:  unixtime.Time(1669190667500),
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
		Image:          "QmWfJDKaGXrXCQyCwQZs1PEzoTKU8VBmRW7X7jDv8Vw8ou",
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
		CreatedAtTime:  unixtime.Time(1669185229500),
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
			CreatedAtTime:  unixtime.Time(1669185229500),
		},
	}

	assert.Equal(t, expected, res.Data)
}
