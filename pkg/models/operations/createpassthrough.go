// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type CreatePassthroughRequest struct {
	// integration-specific payload
	RequestBody map[string]any `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	Path         string `pathParam:"style=simple,explode=false,name=path"`
}

func (o *CreatePassthroughRequest) GetRequestBody() map[string]any {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *CreatePassthroughRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreatePassthroughRequest) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

type CreatePassthroughResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	Result map[string]any
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreatePassthroughResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreatePassthroughResponse) GetResult() map[string]any {
	if o == nil {
		return nil
	}
	return o.Result
}

func (o *CreatePassthroughResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreatePassthroughResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
