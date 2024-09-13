// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type UpdatePassthroughRawRequest struct {
	// integration-specific payload
	// This field accepts []byte data or io.Reader implementations, such as *os.File.
	RequestBody *any `request:"mediaType=text/plain"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	Path         string `pathParam:"style=simple,explode=false,name=path"`
}

func (o *UpdatePassthroughRawRequest) GetRequestBody() *any {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *UpdatePassthroughRawRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdatePassthroughRawRequest) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

type UpdatePassthroughRawResponse struct {
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

func (o *UpdatePassthroughRawResponse) GetTwoXXApplicationJSONAny() any {
	if o == nil {
		return nil
	}
	return o.TwoXXApplicationJSONAny
}

func (o *UpdatePassthroughRawResponse) GetTwoXXTextPlainRes() *string {
	if o == nil {
		return nil
	}
	return o.TwoXXTextPlainRes
}

func (o *UpdatePassthroughRawResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *UpdatePassthroughRawResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdatePassthroughRawResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *UpdatePassthroughRawResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdatePassthroughRawResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}