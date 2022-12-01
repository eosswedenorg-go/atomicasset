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
	Page            int       `qs:"page,omitempty"`
	Limit           int       `qs:"limit,omitempty"`
	Order           SortOrder `qs:"order,omitempty"`
	ActionWhitelist string    `qs:"action_whitelist,omitempty"`
	ActionBlacklist string    `qs:"action_blacklist,omitempty"`
}

// ReqStringList type is used to encode string slices into a single string
// separated by "," instead of qs default (multiple keys with the same name).

// ReqStringList{"a", "b", "c"} url encodes into: "field=a,b,c"
// []string{"a", "b", "c"} encodes into "field=a&field=b&field=c"

// atomicassets calls usually wants comma separated strings.

type ReqStringList []string

func (cs ReqStringList) EncodeParam() (string, error) {
	return strings.Join(cs, ","), nil
}

func (cs ReqStringList) IsZero() bool {
	return len(cs) < 1
}

// Same as ReqStringList type but for integers

type ReqIntList []int

func (cs ReqIntList) EncodeParam() (string, error) {
	l := []string{}
	for _, v := range cs {
		l = append(l, fmt.Sprint(v))
	}
	return strings.Join(l, ","), nil
}

func (cs ReqIntList) IsZero() bool {
	return len(cs) < 1
}
