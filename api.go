package atomicasset

import (
	null "gopkg.in/guregu/null.v4"
)

type APIError struct {
	Success null.Bool   `json:"success"`
	Message null.String `json:"message"`
}

type APIResponse struct {
	HTTPResponse
	Success   bool     `json:"success"`
	QueryTime UnixTime `json:"query_time"`
}
