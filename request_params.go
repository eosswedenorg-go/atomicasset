package atomicasset

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
