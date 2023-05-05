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

func TestGetMarketprice(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicassets/v1/marketplaces/dablabsstore", req.URL.String())

		payload := `{
			"success": true,
			"data": {
				"marketplace_name": "dablabsstore",
				"creator": "dablabsdotio",
				"created_at_block": "183424843",
				"created_at_time": "1653170778000"
			},
			"query_time": 1069360291000
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetMarketplace("dablabsstore")

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2003, time.November, 20, 20, 31, 31, 0, time.UTC), res.QueryTime.Time())

	expected := Marketplace{
		Name:           "dablabsstore",
		Creator:        "dablabsdotio",
		CreatedAtBlock: "183424843",
		CreatedAtTime:  unixtime.Time(1653170778000),
	}

	assert.Equal(t, expected, res.Data)
}

func TestGetMarketprices(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/atomicassets/v1/marketplaces", req.URL.String())

		payload := `{
			"success": true,
			"data": [
				{
					"marketplace_name": "beastsbazaar",
					"creator": "beastsbazaar",
					"created_at_block": "160233137",
					"created_at_time": "1641571675500"
				},
				{
					"marketplace_name": "cheatcodemrk",
					"creator": "lamovichwaxp",
					"created_at_block": "191007878",
					"created_at_time": "1656962990500"
				}
			],
			"query_time": 1069360291000
		  }`

		res.Header().Add("Content-type", "application/json; charset=utf-8")
		_, err := res.Write([]byte(payload))
		assert.NoError(t, err)
	}))

	client := New(srv.URL)

	res, err := client.GetMarketplaces()

	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPStatusCode)
	assert.True(t, res.Success)
	assert.Equal(t, time.Date(2003, time.November, 20, 20, 31, 31, 0, time.UTC), res.QueryTime.Time())

	expected := []Marketplace{
		{
			Name:           "beastsbazaar",
			Creator:        "beastsbazaar",
			CreatedAtBlock: "160233137",
			CreatedAtTime:  unixtime.Time(1641571675500),
		},
		{
			Name:           "cheatcodemrk",
			Creator:        "lamovichwaxp",
			CreatedAtBlock: "191007878",
			CreatedAtTime:  unixtime.Time(1656962990500),
		},
	}

	assert.Equal(t, expected, res.Data)
}
