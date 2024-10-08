// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type RemovePassthroughRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	Path         string `pathParam:"style=simple,explode=false,name=path"`
}

func (o *RemovePassthroughRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *RemovePassthroughRequest) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

type RemovePassthroughResponse struct {
	// Successful
	TwoXXApplicationJSONAny any
	// Successful
	TwoXXTextPlainRes *string
	Body              []byte
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RemovePassthroughResponse) GetTwoXXApplicationJSONAny() any {
	if o == nil {
		return nil
	}
	return o.TwoXXApplicationJSONAny
}

func (o *RemovePassthroughResponse) GetTwoXXTextPlainRes() *string {
	if o == nil {
		return nil
	}
	return o.TwoXXTextPlainRes
}

func (o *RemovePassthroughResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *RemovePassthroughResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RemovePassthroughResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *RemovePassthroughResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RemovePassthroughResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
