package atomicasset

import (
	"net/url"
	"testing"

	"github.com/sonh/qs"
	"github.com/stretchr/testify/assert"
)

func TestRequest_AuctionsRequestParams(t *testing.T) {
	tests := []struct {
		name     string
		input    AuctionsRequestParams
		expected url.Values
	}{
		{"Empty", AuctionsRequestParams{}, url.Values{}},

		{"StateWaiting", AuctionsRequestParams{State: SalesStateWaiting}, url.Values{"state": []string{"0"}}},
		{"StateListed", AuctionsRequestParams{State: SalesStateListed}, url.Values{"state": []string{"1"}}},
		{"StateCanceled", AuctionsRequestParams{State: SalesStateCanceled}, url.Values{"state": []string{"2"}}},
		{"StateSold", AuctionsRequestParams{State: SalesStateSold}, url.Values{"state": []string{"3"}}},
		{"StateInvalid", AuctionsRequestParams{State: SalesStateInvalid}, url.Values{"state": []string{"4"}}},

		{"MaxAssets", AuctionsRequestParams{MaxAssets: 25}, url.Values{"max_assets": []string{"25"}}},
		{"MinAssets", AuctionsRequestParams{MinAssets: 30}, url.Values{"min_assets": []string{"30"}}},

		{"ShowSellerContract", AuctionsRequestParams{ShowSellerContract: "contract"}, url.Values{"show_seller_contract": []string{"contract"}}},

		{"ContractBlacklist", AuctionsRequestParams{ContractBlacklist: []string{"one", "two"}}, url.Values{"contract_blacklist": []string{"one,two"}}},
		{"ContractWhitelist", AuctionsRequestParams{ContractWhitelist: []string{"one", "two"}}, url.Values{"contract_whitelist": []string{"one,two"}}},

		{"SellerBlacklist", AuctionsRequestParams{SellerBlacklist: []string{"one", "two"}}, url.Values{"seller_blacklist": []string{"one,two"}}},
		{"BuyerBlacklist", AuctionsRequestParams{BuyerBlacklist: []string{"one", "two"}}, url.Values{"buyer_blacklist": []string{"one,two"}}},

		{"Seller", AuctionsRequestParams{Seller: []string{"alice", "bob"}}, url.Values{"seller": []string{"alice,bob"}}},
		{"Buyer", AuctionsRequestParams{Buyer: []string{"alice", "bob"}}, url.Values{"buyer": []string{"alice,bob"}}},

		{"MinPrice", AuctionsRequestParams{MinPrice: 20}, url.Values{"min_price": []string{"20"}}},
		{"MaxPrice", AuctionsRequestParams{MaxPrice: 40}, url.Values{"max_price": []string{"40"}}},

		{"MinTemplateMint", AuctionsRequestParams{MinTemplateMint: 1}, url.Values{"min_template_mint": []string{"1"}}},
		{"MaxTemplateMint", AuctionsRequestParams{MaxTemplateMint: 1337}, url.Values{"max_template_mint": []string{"1337"}}},

		{"CollectionName", AuctionsRequestParams{CollectionName: "collection"}, url.Values{"collection_name": []string{"collection"}}},
		{"CollectionBlacklist", AuctionsRequestParams{CollectionBlacklist: []string{"a", "b"}}, url.Values{"collection_blacklist": []string{"a,b"}}},
		{"CollectionWhitelsit", AuctionsRequestParams{CollectionWhitelist: []string{"c", "d"}}, url.Values{"collection_whitelist": []string{"c,d"}}},

		{"SchemaName", AuctionsRequestParams{SchemaName: "schema"}, url.Values{"schema_name": []string{"schema"}}},
		{"TemplateID", AuctionsRequestParams{TemplateID: 1337}, url.Values{"template_id": []string{"1337"}}},

		{"Burned", AuctionsRequestParams{Burned: true}, url.Values{"burned": []string{"true"}}},
		{"Owner", AuctionsRequestParams{Owner: "owner"}, url.Values{"owner": []string{"owner"}}},
		{"Match", AuctionsRequestParams{Match: "value"}, url.Values{"match": []string{"value"}}},
		{"Search", AuctionsRequestParams{Search: "value"}, url.Values{"search": []string{"value"}}},

		{"MatchImmutableName", AuctionsRequestParams{MatchImmutableName: "value"}, url.Values{"match_immutable_name": []string{"value"}}},
		{"MatchMutableName", AuctionsRequestParams{MatchMutableName: "value"}, url.Values{"match_mutable_name": []string{"value"}}},

		{"IsTransferable", AuctionsRequestParams{IsTransferable: true}, url.Values{"is_transferable": []string{"true"}}},
		{"IsBurnable", AuctionsRequestParams{IsBurnable: true}, url.Values{"is_burnable": []string{"true"}}},

		{"Minter", AuctionsRequestParams{Minter: "alice"}, url.Values{"minter": []string{"alice"}}},
		{"Burner", AuctionsRequestParams{Burner: "bob"}, url.Values{"burner": []string{"bob"}}},

		{"IDs", AuctionsRequestParams{IDs: []int{1, 2, 3}}, url.Values{"ids": []string{"1,2,3"}}},

		{"LowerBound", AuctionsRequestParams{LowerBound: "1000"}, url.Values{"lower_bound": []string{"1000"}}},
		{"UpperBound", AuctionsRequestParams{UpperBound: "2000"}, url.Values{"upper_bound": []string{"2000"}}},

		{"Before", AuctionsRequestParams{Before: 10}, url.Values{"before": []string{"10"}}},
		{"After", AuctionsRequestParams{After: 20}, url.Values{"after": []string{"20"}}},

		{"Limit", AuctionsRequestParams{Limit: 50}, url.Values{"limit": []string{"50"}}},
		{"Order None", AuctionsRequestParams{Order: SortNone}, url.Values{}},
		{"Order Desc", AuctionsRequestParams{Order: SortDescending}, url.Values{"order": []string{"desc"}}},
		{"Order Asc", AuctionsRequestParams{Order: SortAscending}, url.Values{"order": []string{"asc"}}},

		{"Sort Created", AuctionsRequestParams{Sort: SaleSortCreated}, url.Values{"sort": []string{"created"}}},
		{"Sort Update", AuctionsRequestParams{Sort: SaleSortUpdated}, url.Values{"sort": []string{"updated"}}},
		{"Sort ID", AuctionsRequestParams{Sort: SaleSortID}, url.Values{"sort": []string{"sale_id"}}},
		{"Sort Price", AuctionsRequestParams{Sort: SaleSortPrice}, url.Values{"sort": []string{"price"}}},
		{"Sort TemplateMint", AuctionsRequestParams{Sort: SaleSortTemplateMint}, url.Values{"sort": []string{"template_mint"}}},
		{"Sort Name", AuctionsRequestParams{Sort: SaleSortName}, url.Values{"sort": []string{"name"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := qs.NewEncoder().Values(tt.input)

			assert.NoError(t, err)
			assert.EqualValues(t, tt.expected, v)
		})
	}
}
