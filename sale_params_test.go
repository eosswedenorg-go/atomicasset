package atomicasset

import (
	"net/url"
	"testing"

	"github.com/sonh/qs"
	"github.com/stretchr/testify/assert"
)

func TestRequest_SalesRequestParams(t *testing.T) {
	tests := []struct {
		name     string
		input    SalesRequestParams
		expected url.Values
	}{
		{"Empty", SalesRequestParams{}, url.Values{}},

		{"StateWaiting", SalesRequestParams{State: SalesStateWaiting}, url.Values{"state": []string{"0"}}},
		{"StateListed", SalesRequestParams{State: SalesStateListed}, url.Values{"state": []string{"1"}}},
		{"StateCanceled", SalesRequestParams{State: SalesStateCanceled}, url.Values{"state": []string{"2"}}},
		{"StateSold", SalesRequestParams{State: SalesStateSold}, url.Values{"state": []string{"3"}}},
		{"StateInvalid", SalesRequestParams{State: SalesStateInvalid}, url.Values{"state": []string{"4"}}},

		{"MaxAssets", SalesRequestParams{MaxAssets: 25}, url.Values{"max_assets": []string{"25"}}},
		{"MinAssets", SalesRequestParams{MinAssets: 30}, url.Values{"min_assets": []string{"30"}}},

		{"ShowSellerContract", SalesRequestParams{ShowSellerContract: "contract"}, url.Values{"show_seller_contract": []string{"contract"}}},

		{"ContractBlacklist", SalesRequestParams{ContractBlacklist: []string{"one", "two"}}, url.Values{"contract_blacklist": []string{"one,two"}}},
		{"ContractWhitelist", SalesRequestParams{ContractWhitelist: []string{"one", "two"}}, url.Values{"contract_whitelist": []string{"one,two"}}},

		{"SellerBlacklist", SalesRequestParams{SellerBlacklist: []string{"one", "two"}}, url.Values{"seller_blacklist": []string{"one,two"}}},
		{"BuyerBlacklist", SalesRequestParams{BuyerBlacklist: []string{"one", "two"}}, url.Values{"buyer_blacklist": []string{"one,two"}}},

		{"Seller", SalesRequestParams{Seller: []string{"alice", "bob"}}, url.Values{"seller": []string{"alice,bob"}}},
		{"Buyer", SalesRequestParams{Buyer: []string{"alice", "bob"}}, url.Values{"buyer": []string{"alice,bob"}}},

		{"MinPrice", SalesRequestParams{MinPrice: 20}, url.Values{"min_price": []string{"20"}}},
		{"MaxPrice", SalesRequestParams{MaxPrice: 40}, url.Values{"max_price": []string{"40"}}},

		{"MinTemplateMint", SalesRequestParams{MinTemplateMint: 1}, url.Values{"min_template_mint": []string{"1"}}},
		{"MaxTemplateMint", SalesRequestParams{MaxTemplateMint: 1337}, url.Values{"max_template_mint": []string{"1337"}}},

		{"CollectionName", SalesRequestParams{CollectionName: "collection"}, url.Values{"collection_name": []string{"collection"}}},
		{"CollectionBlacklist", SalesRequestParams{CollectionBlacklist: []string{"a", "b"}}, url.Values{"collection_blacklist": []string{"a,b"}}},
		{"CollectionWhitelsit", SalesRequestParams{CollectionWhitelist: []string{"c", "d"}}, url.Values{"collection_whitelist": []string{"c,d"}}},

		{"SchemaName", SalesRequestParams{SchemaName: "schema"}, url.Values{"schema_name": []string{"schema"}}},
		{"TemplateID", SalesRequestParams{TemplateID: 1337}, url.Values{"template_id": []string{"1337"}}},

		{"Burned", SalesRequestParams{Burned: true}, url.Values{"burned": []string{"true"}}},
		{"Owner", SalesRequestParams{Owner: "owner"}, url.Values{"owner": []string{"owner"}}},
		{"Match", SalesRequestParams{Match: "value"}, url.Values{"match": []string{"value"}}},
		{"Search", SalesRequestParams{Search: "value"}, url.Values{"search": []string{"value"}}},

		{"MatchImmutableName", SalesRequestParams{MatchImmutableName: "value"}, url.Values{"match_immutable_name": []string{"value"}}},
		{"MatchMutableName", SalesRequestParams{MatchMutableName: "value"}, url.Values{"match_mutable_name": []string{"value"}}},

		{"IsTransferable", SalesRequestParams{IsTransferable: true}, url.Values{"is_transferable": []string{"true"}}},
		{"IsBurnable", SalesRequestParams{IsBurnable: true}, url.Values{"is_burnable": []string{"true"}}},

		{"Minter", SalesRequestParams{Minter: "alice"}, url.Values{"minter": []string{"alice"}}},
		{"Burner", SalesRequestParams{Burner: "bob"}, url.Values{"burner": []string{"bob"}}},

		{"IDs", SalesRequestParams{IDs: []int{1, 2, 3}}, url.Values{"ids": []string{"1,2,3"}}},

		{"LowerBound", SalesRequestParams{LowerBound: "1000"}, url.Values{"lower_bound": []string{"1000"}}},
		{"UpperBound", SalesRequestParams{UpperBound: "2000"}, url.Values{"upper_bound": []string{"2000"}}},

		{"Before", SalesRequestParams{Before: 10}, url.Values{"before": []string{"10"}}},
		{"After", SalesRequestParams{After: 20}, url.Values{"after": []string{"20"}}},

		{"Limit", SalesRequestParams{Limit: 50}, url.Values{"limit": []string{"50"}}},
		{"Order None", SalesRequestParams{Order: SortNone}, url.Values{}},
		{"Order Desc", SalesRequestParams{Order: SortDescending}, url.Values{"order": []string{"desc"}}},
		{"Order Asc", SalesRequestParams{Order: SortAscending}, url.Values{"order": []string{"asc"}}},

		{"Sort Created", SalesRequestParams{Sort: SaleSortCreated}, url.Values{"sort": []string{"created"}}},
		{"Sort Update", SalesRequestParams{Sort: SaleSortUpdated}, url.Values{"sort": []string{"updated"}}},
		{"Sort ID", SalesRequestParams{Sort: SaleSortID}, url.Values{"sort": []string{"sale_id"}}},
		{"Sort Price", SalesRequestParams{Sort: SaleSortPrice}, url.Values{"sort": []string{"price"}}},
		{"Sort TemplateMint", SalesRequestParams{Sort: SaleSortTemplateMint}, url.Values{"sort": []string{"template_mint"}}},
		{"Sort Name", SalesRequestParams{Sort: SaleSortName}, url.Values{"sort": []string{"name"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := qs.NewEncoder().Values(tt.input)

			assert.NoError(t, err)
			assert.EqualValues(t, tt.expected, v)
		})
	}
}

func TestRequest_SalesTemplateRequestParams(t *testing.T) {
	tests := []struct {
		name     string
		input    SalesTemplateRequestParams
		expected url.Values
	}{
		{"Empty", SalesTemplateRequestParams{}, url.Values{}},

		{"Symbol", SalesTemplateRequestParams{Symbol: "WAX"}, url.Values{"symbol": []string{"WAX"}}},

		{"MinPrice", SalesTemplateRequestParams{MinPrice: 20}, url.Values{"min_price": []string{"20"}}},
		{"MaxPrice", SalesTemplateRequestParams{MaxPrice: 40}, url.Values{"max_price": []string{"40"}}},

		{"CollectionName", SalesTemplateRequestParams{CollectionName: "collection"}, url.Values{"collection_name": []string{"collection"}}},
		{"CollectionBlacklist", SalesTemplateRequestParams{CollectionBlacklist: []string{"a", "b"}}, url.Values{"collection_blacklist": []string{"a,b"}}},
		{"CollectionWhitelsit", SalesTemplateRequestParams{CollectionWhitelist: []string{"c", "d"}}, url.Values{"collection_whitelist": []string{"c,d"}}},

		{"SchemaName", SalesTemplateRequestParams{SchemaName: "schema"}, url.Values{"schema_name": []string{"schema"}}},
		{"TemplateID", SalesTemplateRequestParams{TemplateID: 1337}, url.Values{"template_id": []string{"1337"}}},

		{"Burned", SalesTemplateRequestParams{Burned: true}, url.Values{"burned": []string{"true"}}},
		{"Owner", SalesTemplateRequestParams{Owner: "owner"}, url.Values{"owner": []string{"owner"}}},
		{"Match", SalesTemplateRequestParams{Match: "value"}, url.Values{"match": []string{"value"}}},
		{"Search", SalesTemplateRequestParams{Search: "value"}, url.Values{"search": []string{"value"}}},

		{"MatchImmutableName", SalesTemplateRequestParams{MatchImmutableName: "value"}, url.Values{"match_immutable_name": []string{"value"}}},
		{"MatchMutableName", SalesTemplateRequestParams{MatchMutableName: "value"}, url.Values{"match_mutable_name": []string{"value"}}},

		{"IsTransferable", SalesTemplateRequestParams{IsTransferable: true}, url.Values{"is_transferable": []string{"true"}}},
		{"IsBurnable", SalesTemplateRequestParams{IsBurnable: true}, url.Values{"is_burnable": []string{"true"}}},

		{"Minter", SalesTemplateRequestParams{Minter: "alice"}, url.Values{"minter": []string{"alice"}}},
		{"Burner", SalesTemplateRequestParams{Burner: "bob"}, url.Values{"burner": []string{"bob"}}},

		{"InitialReceiver", SalesTemplateRequestParams{InitialReceiver: "recv"}, url.Values{"initial_receiver": []string{"recv"}}},

		{"IDs", SalesTemplateRequestParams{IDs: []int{1, 2, 3}}, url.Values{"ids": []string{"1,2,3"}}},

		{"LowerBound", SalesTemplateRequestParams{LowerBound: "1000"}, url.Values{"lower_bound": []string{"1000"}}},
		{"UpperBound", SalesTemplateRequestParams{UpperBound: "2000"}, url.Values{"upper_bound": []string{"2000"}}},

		{"Before", SalesTemplateRequestParams{Before: 10}, url.Values{"before": []string{"10"}}},
		{"After", SalesTemplateRequestParams{After: 20}, url.Values{"after": []string{"20"}}},

		{"Limit", SalesTemplateRequestParams{Limit: 50}, url.Values{"limit": []string{"50"}}},
		{"Order None", SalesTemplateRequestParams{Order: SortNone}, url.Values{}},
		{"Order Desc", SalesTemplateRequestParams{Order: SortDescending}, url.Values{"order": []string{"desc"}}},
		{"Order Asc", SalesTemplateRequestParams{Order: SortAscending}, url.Values{"order": []string{"asc"}}},

		{"Sort Price", SalesTemplateRequestParams{Sort: SaleTemplateSortPrice}, url.Values{"sort": []string{"price"}}},
		{"Sort TemplateMint", SalesTemplateRequestParams{Sort: SaleTemplateSortTemplateID}, url.Values{"sort": []string{"template_id"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := qs.NewEncoder().Values(tt.input)

			assert.NoError(t, err)
			assert.EqualValues(t, tt.expected, v)
		})
	}
}
