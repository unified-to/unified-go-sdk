// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdatePassthroughRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	Path         string `pathParam:"style=simple,explode=false,name=path"`
	// integration-specific payload
	Undefined *shared.Undefined `request:"mediaType=application/json"`
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

func (o *UpdatePassthroughRequest) GetUndefined() *shared.Undefined {
	if o == nil {
		return nil
	}
	return o.Undefined
}

type UpdatePassthroughResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	Undefined *shared.Undefined
}

func (o *UpdatePassthroughResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
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

func (o *UpdatePassthroughResponse) GetUndefined() *shared.Undefined {
	if o == nil {
		return nil
	}
	return o.Undefined
}
