// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateHrisLocationRequest struct {
	HrisLocation *shared.HrisLocation `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
}

func (o *CreateHrisLocationRequest) GetHrisLocation() *shared.HrisLocation {
	if o == nil {
		return nil
	}
	return o.HrisLocation
}

func (o *CreateHrisLocationRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateHrisLocationRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

type CreateHrisLocationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	HrisLocation *shared.HrisLocation
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateHrisLocationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateHrisLocationResponse) GetHrisLocation() *shared.HrisLocation {
	if o == nil {
		return nil
	}
	return o.HrisLocation
}

func (o *CreateHrisLocationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateHrisLocationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
