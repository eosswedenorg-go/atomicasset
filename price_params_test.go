package atomicasset

import (
	"net/url"
	"testing"

	"github.com/sonh/qs"

	"github.com/stretchr/testify/assert"
)

func TestRequest_PriceSalesRequestParams(t *testing.T) {
	tests := []struct {
		name     string
		input    PriceSalesRequestParams
		expected url.Values
	}{
		{"Empty", PriceSalesRequestParams{}, url.Values{}},

		{"Collection", PriceSalesRequestParams{Collection: "col"}, url.Values{"collection_name": []string{"col"}}},
		{"Schema", PriceSalesRequestParams{Schema: "val"}, url.Values{"schema_name": []string{"val"}}},
		{"TemplateID", PriceSalesRequestParams{TemplateID: 1234}, url.Values{"template_id": []string{"1234"}}},
		{"Burned", PriceSalesRequestParams{Burned: true}, url.Values{"burned": []string{"true"}}},
		{"Symbol", PriceSalesRequestParams{Symbol: "WAX"}, url.Values{"symbol": []string{"WAX"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := qs.NewEncoder().Values(tt.input)

			assert.NoError(t, err)
			assert.EqualValues(t, tt.expected, v)
		})
	}
}

func TestRequest_PriceTemplatesRequestParams(t *testing.T) {
	tests := []struct {
		name     string
		input    PriceTemplatesRequestParams
		expected url.Values
	}{
		{"Empty", PriceTemplatesRequestParams{}, url.Values{}},

		{"Collection", PriceTemplatesRequestParams{Collection: "col"}, url.Values{"collection_name": []string{"col"}}},
		{"Schema", PriceTemplatesRequestParams{Schema: "val"}, url.Values{"schema_name": []string{"val"}}},
		{"TemplateID", PriceTemplatesRequestParams{TemplateID: 1234}, url.Values{"template_id": []string{"1234"}}},
		{"Burned", PriceTemplatesRequestParams{Burned: true}, url.Values{"burned": []string{"true"}}},
		{"Symbol", PriceTemplatesRequestParams{Symbol: "WAX"}, url.Values{"symbol": []string{"WAX"}}},

		{"Page", PriceTemplatesRequestParams{Page: 5}, url.Values{"page": []string{"5"}}},
		{"Limit", PriceTemplatesRequestParams{Limit: 1}, url.Values{"limit": []string{"1"}}},
		{"Order", PriceTemplatesRequestParams{Order: SortAscending}, url.Values{"order": []string{"asc"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := qs.NewEncoder().Values(tt.input)

			assert.NoError(t, err)
			assert.EqualValues(t, tt.expected, v)
		})
	}
}

func TestRequest_PriceAssetsRequestParams(t *testing.T) {
	tests := []struct {
		name     string
		input    PriceAssetsRequestParams
		expected url.Values
	}{
		{"Empty", PriceAssetsRequestParams{}, url.Values{}},

		{"Collection", PriceAssetsRequestParams{CollectionName: "col"}, url.Values{"collection_name": []string{"col"}}},
		{"Schema", PriceAssetsRequestParams{SchemaName: "val"}, url.Values{"schema_name": []string{"val"}}},
		{"TemplateID", PriceAssetsRequestParams{TemplateID: 1234}, url.Values{"template_id": []string{"1234"}}},

		{"Owner", PriceAssetsRequestParams{Owner: "val"}, url.Values{"owner": []string{"val"}}},
		{"Search", PriceAssetsRequestParams{Search: "val"}, url.Values{"search": []string{"val"}}},
		{"MatchImmutableName", PriceAssetsRequestParams{MatchImmutableName: "val"}, url.Values{"match_immutable_name": []string{"val"}}},

		{"MatchMutableName", PriceAssetsRequestParams{MatchMutableName: "val"}, url.Values{"match_mutable_name": []string{"val"}}},

		{"IsTransferable", PriceAssetsRequestParams{IsTransferable: true}, url.Values{"is_transferable": []string{"true"}}},
		{"IsBurnable", PriceAssetsRequestParams{IsBurnable: true}, url.Values{"is_burnable": []string{"true"}}},
		{"Burned", PriceAssetsRequestParams{Burned: true}, url.Values{"burned": []string{"true"}}},
		{"Minter", PriceAssetsRequestParams{Minter: "alice"}, url.Values{"minter": []string{"alice"}}},
		{"InitialReceiver", PriceAssetsRequestParams{InitialReceiver: "bob"}, url.Values{"initial_receiver": []string{"bob"}}},

		{"HideOffers", PriceAssetsRequestParams{HideOffers: true}, url.Values{"hide_offers": []string{"true"}}},
		{"Ids", PriceAssetsRequestParams{Ids: []string{"1", "2"}}, url.Values{"ids": []string{"1,2"}}},
		{"LowerBound", PriceAssetsRequestParams{LowerBound: "100"}, url.Values{"lower_bound": []string{"100"}}},
		{"UpperBound", PriceAssetsRequestParams{UpperBound: "200"}, url.Values{"upper_bound": []string{"200"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := qs.NewEncoder().Values(tt.input)

			assert.NoError(t, err)
			assert.EqualValues(t, tt.expected, v)
		})
	}
}

func TestRequest_PriceInventoryRequestParams(t *testing.T) {
	tests := []struct {
		name     string
		input    PriceInventoryRequestParams
		expected url.Values
	}{
		{"Empty", PriceInventoryRequestParams{}, url.Values{}},

		{"Collection", PriceInventoryRequestParams{CollectionName: "col"}, url.Values{"collection_name": []string{"col"}}},
		{"Schema", PriceInventoryRequestParams{SchemaName: "val"}, url.Values{"schema_name": []string{"val"}}},
		{"TemplateID", PriceInventoryRequestParams{TemplateID: 1234}, url.Values{"template_id": []string{"1234"}}},

		{"Owner", PriceInventoryRequestParams{Owner: "val"}, url.Values{"owner": []string{"val"}}},
		{"Search", PriceInventoryRequestParams{Search: "val"}, url.Values{"search": []string{"val"}}},
		{"MatchImmutableName", PriceInventoryRequestParams{MatchImmutableName: "val"}, url.Values{"match_immutable_name": []string{"val"}}},

		{"MatchMutableName", PriceInventoryRequestParams{MatchMutableName: "val"}, url.Values{"match_mutable_name": []string{"val"}}},

		{"IsTransferable", PriceInventoryRequestParams{IsTransferable: true}, url.Values{"is_transferable": []string{"true"}}},
		{"IsBurnable", PriceInventoryRequestParams{IsBurnable: true}, url.Values{"is_burnable": []string{"true"}}},
		{"Burned", PriceInventoryRequestParams{Burned: true}, url.Values{"burned": []string{"true"}}},
		{"Minter", PriceInventoryRequestParams{Minter: "alice"}, url.Values{"minter": []string{"alice"}}},
		{"InitialReceiver", PriceInventoryRequestParams{InitialReceiver: "bob"}, url.Values{"initial_receiver": []string{"bob"}}},

		{"HideOffers", PriceInventoryRequestParams{HideOffers: true}, url.Values{"hide_offers": []string{"true"}}},
		{"Ids", PriceInventoryRequestParams{Ids: []string{"1", "2"}}, url.Values{"ids": []string{"1,2"}}},
		{"LowerBound", PriceInventoryRequestParams{LowerBound: "100"}, url.Values{"lower_bound": []string{"100"}}},
		{"UpperBound", PriceInventoryRequestParams{UpperBound: "200"}, url.Values{"upper_bound": []string{"200"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := qs.NewEncoder().Values(tt.input)

			assert.NoError(t, err)
			assert.EqualValues(t, tt.expected, v)
		})
	}
}
