package atomicasset

import (
	"net/url"
	"testing"

	"github.com/sonh/qs"

	"github.com/stretchr/testify/assert"
)

func TestRequest_OfferRequestParams(t *testing.T) {
	tests := []struct {
		name     string
		input    OfferRequestParams
		expected url.Values
	}{
		{"Empty", OfferRequestParams{}, url.Values{}},

		{"Account", OfferRequestParams{Account: "alice"}, url.Values{"account": []string{"alice"}}},
		{"Sender", OfferRequestParams{Sender: []string{"alice", "bob"}}, url.Values{"sender": []string{"alice,bob"}}},
		{"Recipient", OfferRequestParams{Recipient: []string{"alice", "bob"}}, url.Values{"recipient": []string{"alice,bob"}}},

		{"Memo", OfferRequestParams{Memo: "value"}, url.Values{"memo": []string{"value"}}},
		{"MemoMatch", OfferRequestParams{MatchMemo: "value"}, url.Values{"match_memo": []string{"value"}}},

		{"StatePending", OfferRequestParams{State: OfferStatePending}, url.Values{"state": []string{"0"}}},
		{"StateInvalid", OfferRequestParams{State: OfferStateInvalid}, url.Values{"state": []string{"1"}}},
		{"StateUnknown", OfferRequestParams{State: OfferStateUnknown}, url.Values{"state": []string{"2"}}},
		{"StateAccepted", OfferRequestParams{State: OfferStateAccepted}, url.Values{"state": []string{"3"}}},
		{"StateDeclined", OfferRequestParams{State: OfferStateDeclined}, url.Values{"state": []string{"4"}}},
		{"StateCanceled", OfferRequestParams{State: OfferStateCanceled}, url.Values{"state": []string{"5"}}},

		{"IsRecipientContract", OfferRequestParams{IsRecipientContract: true}, url.Values{"is_recipient_contract": []string{"true"}}},

		{"AssetID", OfferRequestParams{AssetID: []int{1, 2}}, url.Values{"asset_id": []string{"1,2"}}},
		{"TemplateID", OfferRequestParams{TemplateID: []int{1, 2}}, url.Values{"template_id": []string{"1,2"}}},
		{"SchemeName", OfferRequestParams{SchemaName: []string{"name1", "name2"}}, url.Values{"schema_name": []string{"name1,name2"}}},
		{"CollectionName", OfferRequestParams{CollectionName: []string{"name1", "name2"}}, url.Values{"collection_name": []string{"name1,name2"}}},

		{"AccountWhitelist", OfferRequestParams{AccountWhitelist: []string{"alice", "bob"}}, url.Values{"account_whitelist": []string{"alice,bob"}}},
		{"AccountBlacklist", OfferRequestParams{AccountBlacklist: []string{"alice", "bob"}}, url.Values{"account_blacklist": []string{"alice,bob"}}},
		{"SenderAssetWhitelist", OfferRequestParams{SenderAssetWhitelist: []string{"alice", "bob"}}, url.Values{"sender_asset_whitelist": []string{"alice,bob"}}},
		{"SenderAssetBlacklist", OfferRequestParams{SenderAssetBlacklist: []string{"alice", "bob"}}, url.Values{"sender_asset_blacklist": []string{"alice,bob"}}},
		{"RecipientAssetWhitelist", OfferRequestParams{RecipientAssetWhitelist: []string{"alice", "bob"}}, url.Values{"recipient_asset_whitelist": []string{"alice,bob"}}},
		{"RecipientAssetBlacklist", OfferRequestParams{RecipientAssetBlacklist: []string{"alice", "bob"}}, url.Values{"recipient_asset_blacklist": []string{"alice,bob"}}},
		{"CollectionBlacklist", OfferRequestParams{CollectionBlacklist: []string{"col1", "col2"}}, url.Values{"collection_blacklist": []string{"col1,col2"}}},
		{"CollectionWhitelist", OfferRequestParams{CollectionWhitelist: []string{"col3", "col4"}}, url.Values{"collection_whitelist": []string{"col3,col4"}}},

		{"HideContracts", OfferRequestParams{HideContracts: true}, url.Values{"hide_contracts": []string{"true"}}},
		{"HideEmptyOffers", OfferRequestParams{HideEmptyOffers: true}, url.Values{"hide_empty_offers": []string{"true"}}},

		{"IDs", OfferRequestParams{IDs: ReqStringList{"6", "7", "8"}}, url.Values{"ids": []string{"6,7,8"}}},

		{"LowerBound", OfferRequestParams{LowerBound: "1000"}, url.Values{"lower_bound": []string{"1000"}}},
		{"UpperBound", OfferRequestParams{UpperBound: "2000"}, url.Values{"upper_bound": []string{"2000"}}},

		{"Before", OfferRequestParams{Before: 10}, url.Values{"before": []string{"10"}}},
		{"After", OfferRequestParams{After: 20}, url.Values{"after": []string{"20"}}},

		{"Limit", OfferRequestParams{Limit: 50}, url.Values{"limit": []string{"50"}}},
		{"Order", OfferRequestParams{Order: SortDescending}, url.Values{"order": []string{"desc"}}},
		{"Sort", OfferRequestParams{Sort: "column"}, url.Values{"sort": []string{"column"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := qs.NewEncoder().Values(tt.input)

			assert.NoError(t, err)
			assert.EqualValues(t, tt.expected, v)
		})
	}
}
