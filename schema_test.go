package atomicasset

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
				Image:              "QmYUtzfGWAYrQ43eo1nNRfrYKpUS1cvCWmceCQYxjP7CkS",
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
