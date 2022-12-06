package atomicasset

import (
	"net/url"
	"testing"

	"github.com/sonh/qs"

	"github.com/stretchr/testify/assert"
)

func TestRequest_CollectionsRequestParams(t *testing.T) {
	tests := []struct {
		name     string
		input    CollectionsRequestParams
		expected url.Values
	}{
		{"Empty", CollectionsRequestParams{}, url.Values{}},
		{"Author", CollectionsRequestParams{Author: "alice"}, url.Values{"author": []string{"alice"}}},
		{"Match", CollectionsRequestParams{Match: "value"}, url.Values{"match": []string{"value"}}},
		{"AuthorizedAccount", CollectionsRequestParams{AuthorizedAccount: "bob"}, url.Values{"authorized_account": []string{"bob"}}},
		{"NotifiedAccount", CollectionsRequestParams{NotifyAccount: "cesar"}, url.Values{"notify_account": []string{"cesar"}}},

		{"Blacklist", CollectionsRequestParams{Blacklist: []string{"col1"}}, url.Values{"collection_blacklist": []string{"col1"}}},
		{"Whitelist", CollectionsRequestParams{Whitelist: []string{"col3"}}, url.Values{"collection_whitelist": []string{"col3"}}},

		{"IDs", CollectionsRequestParams{IDs: []int{1, 2, 3}}, url.Values{"ids": []string{"1,2,3"}}},

		{"LowerBound", CollectionsRequestParams{LowerBound: "1000"}, url.Values{"lower_bound": []string{"1000"}}},
		{"UpperBound", CollectionsRequestParams{UpperBound: "2000"}, url.Values{"upper_bound": []string{"2000"}}},

		{"Before", CollectionsRequestParams{Before: 10}, url.Values{"before": []string{"10"}}},
		{"After", CollectionsRequestParams{After: 20}, url.Values{"after": []string{"20"}}},

		{"Limit", CollectionsRequestParams{Limit: 50}, url.Values{"limit": []string{"50"}}},
		{"Order None", CollectionsRequestParams{Order: SortNone}, url.Values{}},
		{"Order Desc", CollectionsRequestParams{Order: SortDescending}, url.Values{"order": []string{"desc"}}},
		{"Order Asc", CollectionsRequestParams{Order: SortAscending}, url.Values{"order": []string{"asc"}}},

		{"Sort Default", CollectionsRequestParams{Sort: CollectionSortDefault}, url.Values{}},
		{"Sort Created", CollectionsRequestParams{Sort: CollectionSortCreated}, url.Values{"sort": []string{"created"}}},
		{"Sort Name", CollectionsRequestParams{Sort: CollectionSortName}, url.Values{"sort": []string{"collection_name"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := qs.NewEncoder().Values(tt.input)

			assert.NoError(t, err)
			assert.EqualValues(t, tt.expected, v)
		})
	}
}

func TestRequest_CollectionLogsRequestParams(t *testing.T) {
	tests := []struct {
		name     string
		input    CollectionLogsRequestParams
		expected url.Values
	}{
		{"Empty", CollectionLogsRequestParams{}, url.Values{}},
		{"Page", CollectionLogsRequestParams{Page: 134}, url.Values{"page": []string{"134"}}},
		{"Limit", CollectionLogsRequestParams{Limit: 50}, url.Values{"limit": []string{"50"}}},
		{"Order None", CollectionLogsRequestParams{Order: SortNone}, url.Values{}},
		{"Order Desc", CollectionLogsRequestParams{Order: SortDescending}, url.Values{"order": []string{"desc"}}},
		{"Order Asc", CollectionLogsRequestParams{Order: SortAscending}, url.Values{"order": []string{"asc"}}},
		{"Blacklist", CollectionLogsRequestParams{ActionBlacklist: []string{"col1"}}, url.Values{"action_blacklist": []string{"col1"}}},
		{"Whitelist", CollectionLogsRequestParams{ActionWhitelist: []string{"col3"}}, url.Values{"action_whitelist": []string{"col3"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := qs.NewEncoder().Values(tt.input)

			assert.NoError(t, err)
			assert.EqualValues(t, tt.expected, v)
		})
	}
}
