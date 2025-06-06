// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateLmsClassRequest struct {
	LmsClass shared.LmsClass `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Class
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Raw parameters to include in the 3rd-party request. Encoded as a URL component. eg. raw parameters: foo=bar&zoo=bar -> raw=foo%3Dbar%26zoo%3Dbar
	Raw *string `queryParam:"style=form,explode=true,name=raw"`
}

func (o *UpdateLmsClassRequest) GetLmsClass() shared.LmsClass {
	if o == nil {
		return shared.LmsClass{}
	}
	return o.LmsClass
}

func (o *UpdateLmsClassRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdateLmsClassRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *UpdateLmsClassRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateLmsClassRequest) GetRaw() *string {
	if o == nil {
		return nil
	}
	return o.Raw
}

type UpdateLmsClassResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	LmsClass *shared.LmsClass
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateLmsClassResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateLmsClassResponse) GetLmsClass() *shared.LmsClass {
	if o == nil {
		return nil
	}
	return o.LmsClass
}

func (o *UpdateLmsClassResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateLmsClassResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
