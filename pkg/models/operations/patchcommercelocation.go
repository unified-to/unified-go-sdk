// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchCommerceLocationRequest struct {
	CommerceLocation *shared.CommerceLocation `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Location
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchCommerceLocationRequest) GetCommerceLocation() *shared.CommerceLocation {
	if o == nil {
		return nil
	}
	return o.CommerceLocation
}

func (o *PatchCommerceLocationRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchCommerceLocationRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchCommerceLocationResponse struct {
	// Successful
	CommerceLocation *shared.CommerceLocation
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PatchCommerceLocationResponse) GetCommerceLocation() *shared.CommerceLocation {
	if o == nil {
		return nil
	}
	return o.CommerceLocation
}

func (o *PatchCommerceLocationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchCommerceLocationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchCommerceLocationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
