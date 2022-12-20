package atomicasset

import (
	"net/url"
	"testing"

	"github.com/sonh/qs"
	"github.com/stretchr/testify/assert"
)

func TestRequest_BuyOffersRequestParams(t *testing.T) {
	tests := []struct {
		name     string
		input    BuyOffersRequestParams
		expected url.Values
	}{
		{"Empty", BuyOffersRequestParams{}, url.Values{}},

		{"StateWaiting", BuyOffersRequestParams{State: SalesStateWaiting}, url.Values{"state": []string{"0"}}},
		{"StateListed", BuyOffersRequestParams{State: SalesStateListed}, url.Values{"state": []string{"1"}}},
		{"StateCanceled", BuyOffersRequestParams{State: SalesStateCanceled}, url.Values{"state": []string{"2"}}},
		{"StateSold", BuyOffersRequestParams{State: SalesStateSold}, url.Values{"state": []string{"3"}}},
		{"StateInvalid", BuyOffersRequestParams{State: SalesStateInvalid}, url.Values{"state": []string{"4"}}},

		{"MaxAssets", BuyOffersRequestParams{MaxAssets: 25}, url.Values{"max_assets": []string{"25"}}},
		{"MinAssets", BuyOffersRequestParams{MinAssets: 30}, url.Values{"min_assets": []string{"30"}}},

		{"ShowSellerContract", BuyOffersRequestParams{ShowSellerContract: "contract"}, url.Values{"show_seller_contract": []string{"contract"}}},

		{"ContractBlacklist", BuyOffersRequestParams{ContractBlacklist: []string{"one", "two"}}, url.Values{"contract_blacklist": []string{"one,two"}}},
		{"ContractWhitelist", BuyOffersRequestParams{ContractWhitelist: []string{"one", "two"}}, url.Values{"contract_whitelist": []string{"one,two"}}},

		{"SellerBlacklist", BuyOffersRequestParams{SellerBlacklist: []string{"one", "two"}}, url.Values{"seller_blacklist": []string{"one,two"}}},
		{"BuyerBlacklist", BuyOffersRequestParams{BuyerBlacklist: []string{"one", "two"}}, url.Values{"buyer_blacklist": []string{"one,two"}}},

		{"Seller", BuyOffersRequestParams{Seller: []string{"alice", "bob"}}, url.Values{"seller": []string{"alice,bob"}}},
		{"Buyer", BuyOffersRequestParams{Buyer: []string{"alice", "bob"}}, url.Values{"buyer": []string{"alice,bob"}}},

		{"MinPrice", BuyOffersRequestParams{MinPrice: 20}, url.Values{"min_price": []string{"20"}}},
		{"MaxPrice", BuyOffersRequestParams{MaxPrice: 40}, url.Values{"max_price": []string{"40"}}},

		{"MinTemplateMint", BuyOffersRequestParams{MinTemplateMint: 1}, url.Values{"min_template_mint": []string{"1"}}},
		{"MaxTemplateMint", BuyOffersRequestParams{MaxTemplateMint: 1337}, url.Values{"max_template_mint": []string{"1337"}}},

		{"CollectionName", BuyOffersRequestParams{CollectionName: "collection"}, url.Values{"collection_name": []string{"collection"}}},
		{"CollectionBlacklist", BuyOffersRequestParams{CollectionBlacklist: []string{"a", "b"}}, url.Values{"collection_blacklist": []string{"a,b"}}},
		{"CollectionWhitelsit", BuyOffersRequestParams{CollectionWhitelist: []string{"c", "d"}}, url.Values{"collection_whitelist": []string{"c,d"}}},

		{"SchemaName", BuyOffersRequestParams{SchemaName: "schema"}, url.Values{"schema_name": []string{"schema"}}},
		{"TemplateID", BuyOffersRequestParams{TemplateID: 1337}, url.Values{"template_id": []string{"1337"}}},

		{"Burned", BuyOffersRequestParams{Burned: true}, url.Values{"burned": []string{"true"}}},
		{"Owner", BuyOffersRequestParams{Owner: "owner"}, url.Values{"owner": []string{"owner"}}},
		{"Match", BuyOffersRequestParams{Match: "value"}, url.Values{"match": []string{"value"}}},
		{"Search", BuyOffersRequestParams{Search: "value"}, url.Values{"search": []string{"value"}}},

		{"MatchImmutableName", BuyOffersRequestParams{MatchImmutableName: "value"}, url.Values{"match_immutable_name": []string{"value"}}},
		{"MatchMutableName", BuyOffersRequestParams{MatchMutableName: "value"}, url.Values{"match_mutable_name": []string{"value"}}},

		{"IsTransferable", BuyOffersRequestParams{IsTransferable: true}, url.Values{"is_transferable": []string{"true"}}},
		{"IsBurnable", BuyOffersRequestParams{IsBurnable: true}, url.Values{"is_burnable": []string{"true"}}},

		{"Minter", BuyOffersRequestParams{Minter: "alice"}, url.Values{"minter": []string{"alice"}}},
		{"Burner", BuyOffersRequestParams{Burner: "bob"}, url.Values{"burner": []string{"bob"}}},

		{"InitialReceiver", BuyOffersRequestParams{InitialReceiver: "recv"}, url.Values{"initial_receiver": []string{"recv"}}},

		{"IDs", BuyOffersRequestParams{IDs: []int{1, 2, 3}}, url.Values{"ids": []string{"1,2,3"}}},

		{"LowerBound", BuyOffersRequestParams{LowerBound: "1000"}, url.Values{"lower_bound": []string{"1000"}}},
		{"UpperBound", BuyOffersRequestParams{UpperBound: "2000"}, url.Values{"upper_bound": []string{"2000"}}},

		{"Before", BuyOffersRequestParams{Before: 10}, url.Values{"before": []string{"10"}}},
		{"After", BuyOffersRequestParams{After: 20}, url.Values{"after": []string{"20"}}},

		{"Limit", BuyOffersRequestParams{Limit: 50}, url.Values{"limit": []string{"50"}}},
		{"Order None", BuyOffersRequestParams{Order: SortNone}, url.Values{}},
		{"Order Desc", BuyOffersRequestParams{Order: SortDescending}, url.Values{"order": []string{"desc"}}},
		{"Order Asc", BuyOffersRequestParams{Order: SortAscending}, url.Values{"order": []string{"asc"}}},

		{"Sort Created", BuyOffersRequestParams{Sort: BuyOfferSortCreated}, url.Values{"sort": []string{"created"}}},
		{"Sort Update", BuyOffersRequestParams{Sort: BuyOfferSortUpdated}, url.Values{"sort": []string{"updated"}}},
		{"Sort ID", BuyOffersRequestParams{Sort: BuyOfferSortID}, url.Values{"sort": []string{"buyoffer_id"}}},
		{"Sort Price", BuyOffersRequestParams{Sort: BuyOfferSortPrice}, url.Values{"sort": []string{"price"}}},
		{"Sort TemplateMint", BuyOffersRequestParams{Sort: BuyOfferSortTemplateMint}, url.Values{"sort": []string{"template_mint"}}},
		{"Sort Name", BuyOffersRequestParams{Sort: BuyOfferSortName}, url.Values{"sort": []string{"name"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := qs.NewEncoder().Values(tt.input)

			assert.NoError(t, err)
			assert.EqualValues(t, tt.expected, v)
		})
	}
}
