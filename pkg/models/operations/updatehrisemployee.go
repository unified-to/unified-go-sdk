// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateHrisEmployeeRequest struct {
	HrisEmployee shared.HrisEmployee `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Employee
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateHrisEmployeeRequest) GetHrisEmployee() shared.HrisEmployee {
	if o == nil {
		return shared.HrisEmployee{}
	}
	return o.HrisEmployee
}

func (o *UpdateHrisEmployeeRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdateHrisEmployeeRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *UpdateHrisEmployeeRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateHrisEmployeeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	HrisEmployee *shared.HrisEmployee
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateHrisEmployeeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateHrisEmployeeResponse) GetHrisEmployee() *shared.HrisEmployee {
	if o == nil {
		return nil
	}
	return o.HrisEmployee
}

func (o *UpdateHrisEmployeeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateHrisEmployeeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
