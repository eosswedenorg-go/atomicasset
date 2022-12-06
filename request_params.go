package atomicasset

import (
	"fmt"
	"strings"
)

type SortOrder string

const (

	// SortNone order is not specified.
	SortNone SortOrder = ""

	// SortAscending sorts in ascending order
	SortAscending SortOrder = "asc"

	// SortDescending sorts in descending order.
	SortDescending SortOrder = "desc"
)

// LogRequestParams holds the parameters for an Log request
type LogRequestParams struct {
	Page            int             `qs:"page,omitempty"`
	Limit           int             `qs:"limit,omitempty"`
	Order           SortOrder       `qs:"order,omitempty"`
	ActionWhitelist ReqList[string] `qs:"action_whitelist,omitempty"`
	ActionBlacklist ReqList[string] `qs:"action_blacklist,omitempty"`
}

// ReqStringList type is used to encode string slices into a single string
// separated by "," instead of qs default (multiple keys with the same name).

// ReqStringList{"a", "b", "c"} url encodes into: "field=a,b,c"
// []string{"a", "b", "c"} encodes into "field=a&field=b&field=c"

// atomicassets calls usually wants comma separated strings.

type ReqList[T any] []T

func (l ReqList[T]) EncodeParam() (string, error) {
	f := []string{}
	for _, v := range l {
		f = append(f, fmt.Sprint(v))
	}
	return strings.Join(f, ","), nil
}

func (l ReqList[T]) IsZero() bool {
	return len(l) < 1
}
