// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PatchPassthroughJSONRequest struct {
	// integration-specific payload
	RequestBody any `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	Path         string `pathParam:"style=simple,explode=false,name=path"`
}

func (o *PatchPassthroughJSONRequest) GetRequestBody() any {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *PatchPassthroughJSONRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchPassthroughJSONRequest) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

type PatchPassthroughJSONResponse struct {
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

func (o *PatchPassthroughJSONResponse) GetTwoXXApplicationJSONAny() any {
	if o == nil {
		return nil
	}
	return o.TwoXXApplicationJSONAny
}

func (o *PatchPassthroughJSONResponse) GetTwoXXTextPlainRes() *string {
	if o == nil {
		return nil
	}
	return o.TwoXXTextPlainRes
}

func (o *PatchPassthroughJSONResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *PatchPassthroughJSONResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchPassthroughJSONResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *PatchPassthroughJSONResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchPassthroughJSONResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
