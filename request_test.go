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

func TestRequest_AssetsRequestParams(t *testing.T) {
	tests := []struct {
		name     string
		input    AssetsRequestParams
		expected url.Values
	}{
		{"Empty", AssetsRequestParams{}, url.Values{}},

		{"CollectionName", AssetsRequestParams{CollectionName: "name"}, url.Values{"collection_name": []string{"name"}}},
		{"CollectionBlacklist", AssetsRequestParams{CollectionBlacklist: []string{"col1", "col2"}}, url.Values{"collection_blacklist": []string{"col1", "col2"}}},
		{"CollectionWhitelist", AssetsRequestParams{CollectionWhitelist: []string{"col3", "col4"}}, url.Values{"collection_whitelist": []string{"col3", "col4"}}},
		{"SchemaName", AssetsRequestParams{SchemaName: "schema"}, url.Values{"schema_name": []string{"schema"}}},

		{"TemplateID", AssetsRequestParams{TemplateID: 1337}, url.Values{"template_id": []string{"1337"}}},
		{"TemplateWhitelist", AssetsRequestParams{TemplateWhitelist: []int{1, 2}}, url.Values{"template_whitelist": []string{"1", "2"}}},
		{"TemplateBlacklist", AssetsRequestParams{TemplateBlacklist: []int{3, 4}}, url.Values{"template_blacklist": []string{"3", "4"}}},

		{"Owner", AssetsRequestParams{Owner: "name"}, url.Values{"owner": []string{"name"}}},

		{"Match", AssetsRequestParams{Match: "value"}, url.Values{"match": []string{"value"}}},
		{"MatchImmutableName", AssetsRequestParams{MatchImmutableName: "value"}, url.Values{"match_immutable_name": []string{"value"}}},
		{"MatchMutableName", AssetsRequestParams{MatchMutableName: "value"}, url.Values{"match_mutable_name": []string{"value"}}},

		{"HideTemplatesByAccounts", AssetsRequestParams{HideTemplatesByAccounts: "account"}, url.Values{"hide_templates_by_accounts": []string{"account"}}},
		{"IsTransferable", AssetsRequestParams{IsTransferable: true}, url.Values{"is_transferable": []string{"true"}}},
		{"IsBurnable", AssetsRequestParams{IsBurnable: true}, url.Values{"is_burnable": []string{"true"}}},
		{"Burned", AssetsRequestParams{Burned: true}, url.Values{"burned": []string{"true"}}},
		{"OnlyDuplicatedTemplates", AssetsRequestParams{OnlyDuplicatedTemplates: true}, url.Values{"only_duplicated_templates": []string{"true"}}},

		{"HasBackedTokens", AssetsRequestParams{HasBackedTokens: true}, url.Values{"has_backend_tokens": []string{"true"}}},
		{"HideOffers", AssetsRequestParams{HideOffers: true}, url.Values{"hide_offers": []string{"true"}}},

		{"LowerBound", AssetsRequestParams{LowerBound: "1000"}, url.Values{"lower_bound": []string{"1000"}}},
		{"UpperBound", AssetsRequestParams{UpperBound: "2000"}, url.Values{"upper_bound": []string{"2000"}}},

		{"Before", AssetsRequestParams{Before: 10}, url.Values{"before": []string{"10"}}},
		{"After", AssetsRequestParams{After: 20}, url.Values{"after": []string{"20"}}},

		{"Limit", AssetsRequestParams{Limit: 50}, url.Values{"limit": []string{"50"}}},
		{"Order", AssetsRequestParams{Order: SortDescending}, url.Values{"order": []string{"desc"}}},
		{"Sort", AssetsRequestParams{Sort: "column"}, url.Values{"sort": []string{"column"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := qs.NewEncoder().Values(tt.input)

			assert.NoError(t, err)
			assert.EqualValues(t, tt.expected, v)
		})
	}
}

func TestRequest_AssetSalesRequestParams(t *testing.T) {
	tests := []struct {
		name     string
		input    AssetSalesRequestParams
		expected url.Values
	}{
		{"Empty", AssetSalesRequestParams{}, url.Values{}},
		{"Page", AssetSalesRequestParams{Buyer: "alice"}, url.Values{"buyer": []string{"alice"}}},
		{"Limit", AssetSalesRequestParams{Seller: "bob"}, url.Values{"seller": []string{"bob"}}},
		{"Order None", AssetSalesRequestParams{Order: SortNone}, url.Values{}},
		{"Order Desc", AssetSalesRequestParams{Order: SortDescending}, url.Values{"order": []string{"desc"}}},
		{"Order Asc", AssetSalesRequestParams{Order: SortAscending}, url.Values{"order": []string{"asc"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := qs.NewEncoder().Values(tt.input)

			assert.NoError(t, err)
			assert.EqualValues(t, tt.expected, v)
		})
	}
}

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

		{"Blacklist", CollectionsRequestParams{Blacklist: "col1"}, url.Values{"collection_blacklist": []string{"col1"}}},
		{"Whitelist", CollectionsRequestParams{Whitelist: "col3"}, url.Values{"collection_whitelist": []string{"col3"}}},

		{"IDs", CollectionsRequestParams{IDs: "1,2,3"}, url.Values{"ids": []string{"1,2,3"}}},

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
		{"Blacklist", CollectionLogsRequestParams{ActionBlacklist: "col1"}, url.Values{"action_blacklist": []string{"col1"}}},
		{"Whitelist", CollectionLogsRequestParams{ActionWhitelist: "col3"}, url.Values{"action_whitelist": []string{"col3"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := qs.NewEncoder().Values(tt.input)

			assert.NoError(t, err)
			assert.EqualValues(t, tt.expected, v)
		})
	}
}

func TestRequest_SchemasRequestParams(t *testing.T) {
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

		{"Blacklist", SchemasRequestParams{Blacklist: "col1"}, url.Values{"collection_blacklist": []string{"col1"}}},
		{"Whitelist", SchemasRequestParams{Whitelist: "col3"}, url.Values{"collection_whitelist": []string{"col3"}}},

		{"IDs", SchemasRequestParams{IDs: "1,2,3"}, url.Values{"ids": []string{"1,2,3"}}},

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
