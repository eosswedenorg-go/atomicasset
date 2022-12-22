package atomicasset

import (
	"net/url"
	"testing"

	"github.com/sonh/qs"

	"github.com/stretchr/testify/assert"
)

func TestRequest_LinkRequestParams(t *testing.T) {
	tests := []struct {
		name     string
		input    LinkRequestParams
		expected url.Values
	}{
		{"Empty", LinkRequestParams{}, url.Values{}},

		{"Creator", LinkRequestParams{Creator: "alice"}, url.Values{"creator": []string{"alice"}}},
		{"Claimer", LinkRequestParams{Claimer: "bob"}, url.Values{"claimer": []string{"bob"}}},
		{"PublicKey", LinkRequestParams{PublicKey: "key"}, url.Values{"public_key": []string{"key"}}},

		{"State", LinkRequestParams{State: []LinkState{LinkStateWaiting, LinkStateCanceled}}, url.Values{"state": []string{"0,2"}}},

		{"CollectionBlacklist", LinkRequestParams{CollectionBlacklist: []string{"col1", "col2"}}, url.Values{"collection_blacklist": []string{"col1,col2"}}},
		{"CollectionWhitelist", LinkRequestParams{CollectionWhitelist: []string{"col3", "col4"}}, url.Values{"collection_whitelist": []string{"col3,col4"}}},

		{"IDs", LinkRequestParams{IDs: []int{6, 7, 8}}, url.Values{"ids": []string{"6,7,8"}}},

		{"LowerBound", LinkRequestParams{LowerBound: "1000"}, url.Values{"lower_bound": []string{"1000"}}},
		{"UpperBound", LinkRequestParams{UpperBound: "2000"}, url.Values{"upper_bound": []string{"2000"}}},

		{"Before", LinkRequestParams{Before: 10}, url.Values{"before": []string{"10"}}},
		{"After", LinkRequestParams{After: 20}, url.Values{"after": []string{"20"}}},

		{"Limit", LinkRequestParams{Limit: 50}, url.Values{"limit": []string{"50"}}},
		{"Order", LinkRequestParams{Order: SortDescending}, url.Values{"order": []string{"desc"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := qs.NewEncoder().Values(tt.input)

			assert.NoError(t, err)
			assert.EqualValues(t, tt.expected, v)
		})
	}
}
