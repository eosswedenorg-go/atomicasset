package atomicasset

import (
	"net/url"
	"testing"

	"github.com/sonh/qs"

	"github.com/stretchr/testify/assert"
)

func TestSchemasRequestParams(t *testing.T) {
	tests := []struct {
		name     string
		input    SchemasRequestParams
		expected url.Values
	}{
		{"Empty", SchemasRequestParams{}, url.Values{}},
		{"Author", SchemasRequestParams{Author: "alice"}, url.Values{"author": []string{"alice"}}},
		{"Match", SchemasRequestParams{Match: "value"}, url.Values{"match": []string{"value"}}},
		{"AuthorizedAccount", SchemasRequestParams{AuthorizedAccount: "bob"}, url.Values{"authorized_account": []string{"bob"}}},
		{"NotifiedAccount", SchemasRequestParams{NotifyAccount: "cesar"}, url.Values{"notify_account": []string{"cesar"}}},

		{"Blacklist", SchemasRequestParams{Blacklist: []string{"col1", "col2"}}, url.Values{"collection_blacklist": []string{"col1,col2"}}},
		{"Whitelist", SchemasRequestParams{Whitelist: []string{"col3", "col4"}}, url.Values{"collection_whitelist": []string{"col3,col4"}}},

		{"IDs", SchemasRequestParams{IDs: []int{1, 2, 3}}, url.Values{"ids": []string{"1,2,3"}}},

		{"LowerBound", SchemasRequestParams{LowerBound: "1000"}, url.Values{"lower_bound": []string{"1000"}}},
		{"UpperBound", SchemasRequestParams{UpperBound: "2000"}, url.Values{"upper_bound": []string{"2000"}}},

		{"Before", SchemasRequestParams{Before: 10}, url.Values{"before": []string{"10"}}},
		{"After", SchemasRequestParams{After: 20}, url.Values{"after": []string{"20"}}},

		{"Limit", SchemasRequestParams{Limit: 50}, url.Values{"limit": []string{"50"}}},
		{"Order None", SchemasRequestParams{Order: SortNone}, url.Values{}},
		{"Order Desc", SchemasRequestParams{Order: SortDescending}, url.Values{"order": []string{"desc"}}},
		{"Order Asc", SchemasRequestParams{Order: SortAscending}, url.Values{"order": []string{"asc"}}},

		{"Sort Default", SchemasRequestParams{Sort: SchemaSortDefault}, url.Values{}},
		{"Sort Created", SchemasRequestParams{Sort: SchemaSortCreated}, url.Values{"sort": []string{"created"}}},
		{"Sort Assets", SchemasRequestParams{Sort: SchemaSortAssets}, url.Values{"sort": []string{"assets"}}},
		{"Sort Name", SchemasRequestParams{Sort: SchemaSortName}, url.Values{"sort": []string{"schema_name"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := qs.NewEncoder().Values(tt.input)

			assert.NoError(t, err)
			assert.EqualValues(t, tt.expected, v)
		})
	}
}
