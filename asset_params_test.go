package atomicasset

import (
	"net/url"
	"testing"

	"github.com/sonh/qs"

	"github.com/stretchr/testify/assert"
)

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
