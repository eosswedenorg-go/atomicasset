package atomicasset

import (
	"net/url"
	"testing"

	"github.com/sonh/qs"

	"github.com/stretchr/testify/assert"
)

func TestRequest_TransferRequestParams(t *testing.T) {
	tests := []struct {
		name     string
		input    TransferRequestParams
		expected url.Values
	}{
		{"Empty", TransferRequestParams{}, url.Values{}},

		{"Account", TransferRequestParams{Account: []string{"alice", "bob"}}, url.Values{"account": []string{"alice,bob"}}},
		{"Sender", TransferRequestParams{Sender: []string{"alice", "bob"}}, url.Values{"sender": []string{"alice,bob"}}},
		{"Recipient", TransferRequestParams{Recipient: []string{"alice", "bob"}}, url.Values{"recipient": []string{"alice,bob"}}},

		{"Memo", TransferRequestParams{Memo: "value"}, url.Values{"memo": []string{"value"}}},
		{"MemoMatch", TransferRequestParams{MatchMemo: "value"}, url.Values{"match_memo": []string{"value"}}},

		{"AssetID", TransferRequestParams{AssetID: []int{1, 2}}, url.Values{"asset_id": []string{"1,2"}}},
		{"TemplateID", TransferRequestParams{TemplateID: []int{1, 2}}, url.Values{"template_id": []string{"1,2"}}},
		{"SchemeName", TransferRequestParams{SchemaName: []string{"name1", "name2"}}, url.Values{"schema_name": []string{"name1,name2"}}},
		{"CollectionName", TransferRequestParams{CollectionName: []string{"name1", "name2"}}, url.Values{"collection_name": []string{"name1,name2"}}},
		{"CollectionBlacklist", TransferRequestParams{CollectionBlacklist: []string{"col1", "col2"}}, url.Values{"collection_blacklist": []string{"col1,col2"}}},
		{"CollectionWhitelist", TransferRequestParams{CollectionWhitelist: []string{"col3", "col4"}}, url.Values{"collection_whitelist": []string{"col3,col4"}}},

		{"HideContracts", TransferRequestParams{HideContracts: true}, url.Values{"hide_contracts": []string{"true"}}},

		{"IDs", TransferRequestParams{IDs: []int{6, 7, 8}}, url.Values{"ids": []string{"6,7,8"}}},

		{"LowerBound", TransferRequestParams{LowerBound: "1000"}, url.Values{"lower_bound": []string{"1000"}}},
		{"UpperBound", TransferRequestParams{UpperBound: "2000"}, url.Values{"upper_bound": []string{"2000"}}},

		{"Before", TransferRequestParams{Before: 10}, url.Values{"before": []string{"10"}}},
		{"After", TransferRequestParams{After: 20}, url.Values{"after": []string{"20"}}},

		{"Limit", TransferRequestParams{Limit: 50}, url.Values{"limit": []string{"50"}}},
		{"Order", TransferRequestParams{Order: SortDescending}, url.Values{"order": []string{"desc"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := qs.NewEncoder().Values(tt.input)

			assert.NoError(t, err)
			assert.EqualValues(t, tt.expected, v)
		})
	}
}
