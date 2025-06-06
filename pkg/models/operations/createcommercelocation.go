// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateCommerceLocationRequest struct {
	CommerceLocation shared.CommerceLocation `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// Raw parameters to include in the 3rd-party request. Encoded as a URL component. eg. raw parameters: foo=bar&zoo=bar -> raw=foo%3Dbar%26zoo%3Dbar
	Raw *string `queryParam:"style=form,explode=true,name=raw"`
}

func (o *CreateCommerceLocationRequest) GetCommerceLocation() shared.CommerceLocation {
	if o == nil {
		return shared.CommerceLocation{}
	}
	return o.CommerceLocation
}

func (o *CreateCommerceLocationRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateCommerceLocationRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *CreateCommerceLocationRequest) GetRaw() *string {
	if o == nil {
		return nil
	}
	return o.Raw
}

type CreateCommerceLocationResponse struct {
	// Successful
	CommerceLocation *shared.CommerceLocation
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateCommerceLocationResponse) GetCommerceLocation() *shared.CommerceLocation {
	if o == nil {
		return nil
	}
	return o.CommerceLocation
}

func (o *CreateCommerceLocationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateCommerceLocationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateCommerceLocationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
