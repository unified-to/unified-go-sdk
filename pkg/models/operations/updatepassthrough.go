// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type UpdatePassthroughRequest struct {
	// integration-specific payload
	RequestBody map[string]any `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	Path         string `pathParam:"style=simple,explode=false,name=path"`
}

func (o *UpdatePassthroughRequest) GetRequestBody() map[string]any {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *UpdatePassthroughRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdatePassthroughRequest) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

type UpdatePassthroughResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	Result map[string]any
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdatePassthroughResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdatePassthroughResponse) GetResult() map[string]any {
	if o == nil {
		return nil
	}
	return o.Result
}

func (o *UpdatePassthroughResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdatePassthroughResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
