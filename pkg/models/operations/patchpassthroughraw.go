// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PatchPassthroughRawRequest struct {
	// integration-specific payload
	// This field accepts []byte data or io.Reader implementations, such as *os.File.
	RequestBody *any `request:"mediaType=text/plain"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	Path         string `pathParam:"style=simple,explode=false,name=path"`
}

func (o *PatchPassthroughRawRequest) GetRequestBody() *any {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *PatchPassthroughRawRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchPassthroughRawRequest) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

type PatchPassthroughRawResponse struct {
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

func (o *PatchPassthroughRawResponse) GetTwoXXApplicationJSONAny() any {
	if o == nil {
		return nil
	}
	return o.TwoXXApplicationJSONAny
}

func (o *PatchPassthroughRawResponse) GetTwoXXTextPlainRes() *string {
	if o == nil {
		return nil
	}
	return o.TwoXXTextPlainRes
}

func (o *PatchPassthroughRawResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *PatchPassthroughRawResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchPassthroughRawResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *PatchPassthroughRawResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchPassthroughRawResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
