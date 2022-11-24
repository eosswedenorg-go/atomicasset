package atomicasset

type HTTPResponse struct {
	HTTPStatusCode int
}

func (resp *HTTPResponse) IsError() bool {
	return resp.HTTPStatusCode == 0 || resp.HTTPStatusCode > 399
}
