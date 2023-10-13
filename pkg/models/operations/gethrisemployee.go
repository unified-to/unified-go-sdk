// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type GetHrisEmployeeRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Employee
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetHrisEmployeeRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetHrisEmployeeRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetHrisEmployeeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	HrisEmployee *shared.HrisEmployee
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetHrisEmployeeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetHrisEmployeeResponse) GetHrisEmployee() *shared.HrisEmployee {
	if o == nil {
		return nil
	}
	return o.HrisEmployee
}

func (o *GetHrisEmployeeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetHrisEmployeeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}