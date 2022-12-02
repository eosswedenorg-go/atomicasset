package atomicasset

import (
	"net/url"
	"testing"

	"github.com/sonh/qs"

	"github.com/stretchr/testify/assert"
)

func TestRequest_TemplateRequestParams(t *testing.T) {
	tests := []struct {
		name     string
		input    TemplateRequestParams
		expected url.Values
	}{
		{"Empty", TemplateRequestParams{}, url.Values{}},

		{"SchemaName", TemplateRequestParams{SchemaName: "schema"}, url.Values{"schema_name": []string{"schema"}}},

		{"CollectionName", TemplateRequestParams{CollectionName: "name"}, url.Values{"collection_name": []string{"name"}}},
		{"CollectionBlacklist", TemplateRequestParams{CollectionBlacklist: []string{"col1", "col2"}}, url.Values{"collection_blacklist": []string{"col1,col2"}}},
		{"CollectionWhitelist", TemplateRequestParams{CollectionWhitelist: []string{"col3", "col4"}}, url.Values{"collection_whitelist": []string{"col3,col4"}}},

		{"IssuedSupply", TemplateRequestParams{IssuedSypply: 3420}, url.Values{"issued_supply": []string{"3420"}}},
		{"MinIssuedSupply", TemplateRequestParams{MinIssuedSupply: 1233}, url.Values{"min_issued_supply": []string{"1233"}}},
		{"MaxIssuedSupply", TemplateRequestParams{MaxIssuedSupply: 4600}, url.Values{"max_issued_supply": []string{"4600"}}},
		{"MaxSupply", TemplateRequestParams{MaxSupply: 9000}, url.Values{"max_supply": []string{"9000"}}},

		{"HasAssets", TemplateRequestParams{HasAssets: true}, url.Values{"has_assets": []string{"true"}}},
		{"IsBurnable", TemplateRequestParams{IsBurnable: true}, url.Values{"is_burnable": []string{"true"}}},
		{"IsTransferable", TemplateRequestParams{IsTransferable: true}, url.Values{"is_transferable": []string{"true"}}},

		{"AuthorizedAccount", TemplateRequestParams{AuthorizedAccount: "acc"}, url.Values{"authorized_account": []string{"acc"}}},

		{"Match", TemplateRequestParams{Match: "value"}, url.Values{"match": []string{"value"}}},

		{"IDs", TemplateRequestParams{IDs: ReqStringList{"6", "7", "8"}}, url.Values{"ids": []string{"6,7,8"}}},

		{"LowerBound", TemplateRequestParams{LowerBound: "1000"}, url.Values{"lower_bound": []string{"1000"}}},
		{"UpperBound", TemplateRequestParams{UpperBound: "2000"}, url.Values{"upper_bound": []string{"2000"}}},

		{"Before", TemplateRequestParams{Before: 10}, url.Values{"before": []string{"10"}}},
		{"After", TemplateRequestParams{After: 20}, url.Values{"after": []string{"20"}}},

		{"Limit", TemplateRequestParams{Limit: 50}, url.Values{"limit": []string{"50"}}},
		{"Order", TemplateRequestParams{Order: SortDescending}, url.Values{"order": []string{"desc"}}},
		{"Sort", TemplateRequestParams{Sort: "column"}, url.Values{"sort": []string{"column"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := qs.NewEncoder().Values(tt.input)

			assert.NoError(t, err)
			assert.EqualValues(t, tt.expected, v)
		})
	}
}
