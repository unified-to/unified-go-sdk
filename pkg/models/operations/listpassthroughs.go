// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ListPassthroughsRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	Path         string `pathParam:"style=simple,explode=false,name=path"`
}

func (o *ListPassthroughsRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListPassthroughsRequest) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

type ListPassthroughsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	Result map[string]any
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListPassthroughsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListPassthroughsResponse) GetResult() map[string]any {
	if o == nil {
		return nil
	}
	return o.Result
}

func (o *ListPassthroughsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListPassthroughsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
