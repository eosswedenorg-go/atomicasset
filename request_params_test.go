package atomicasset

import (
	"net/url"
	"testing"

	"github.com/sonh/qs"

	"github.com/stretchr/testify/assert"
)

func TestRequest_LogRequestParams(t *testing.T) {
	tests := []struct {
		name     string
		input    LogRequestParams
		expected url.Values
	}{
		{"Empty", LogRequestParams{}, url.Values{}},
		{"Page", LogRequestParams{Page: 134}, url.Values{"page": []string{"134"}}},
		{"Limit", LogRequestParams{Limit: 50}, url.Values{"limit": []string{"50"}}},
		{"Order", LogRequestParams{Order: SortDescending}, url.Values{"order": []string{"desc"}}},
		{"Whitelist", LogRequestParams{ActionWhitelist: "one,two"}, url.Values{"action_whitelist": []string{"one,two"}}},
		{"Blacklist", LogRequestParams{ActionBlacklist: "one,two"}, url.Values{"action_blacklist": []string{"one,two"}}},
		{"PageOrderLimit", LogRequestParams{Page: 2, Limit: 30, Order: SortAscending}, url.Values{"order": []string{"asc"}, "page": []string{"2"}, "limit": []string{"30"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := qs.NewEncoder().Values(tt.input)

			assert.NoError(t, err)
			assert.EqualValues(t, tt.expected, v)
		})
	}
}
