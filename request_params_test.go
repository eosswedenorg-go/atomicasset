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

func TestReqStringList_EncodeParam(t *testing.T) {
	tests := []struct {
		name    string
		cs      ReqStringList
		want    string
		wantErr bool
	}{
		{"Empty", []string{}, "", false},
		{"One", []string{"one"}, "one", false},
		{"Many", []string{"one", "two", "three"}, "one,two,three", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.cs.EncodeParam()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReqStringList.EncodeParam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReqStringList.EncodeParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReqStringList_IsZero(t *testing.T) {
	tests := []struct {
		name string
		cs   ReqStringList
		want bool
	}{
		{"Empty", []string{}, true},
		{"Non empty", []string{"random"}, false},
		{"2 elements", []string{"one", "two"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cs.IsZero(); got != tt.want {
				t.Errorf("ReqStringList.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
