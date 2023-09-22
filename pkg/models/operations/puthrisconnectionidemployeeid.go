// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PutHrisConnectionIDEmployeeIDRequest struct {
	HrisEmployee *shared.HrisEmployee `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Employee
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PutHrisConnectionIDEmployeeIDRequest) GetHrisEmployee() *shared.HrisEmployee {
	if o == nil {
		return nil
	}
	return o.HrisEmployee
}

func (o *PutHrisConnectionIDEmployeeIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PutHrisConnectionIDEmployeeIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PutHrisConnectionIDEmployeeIDResponse struct {
	ContentType string
	// Successful
	HrisEmployee *shared.HrisEmployee
	StatusCode   int
	RawResponse  *http.Response
}

func (o *PutHrisConnectionIDEmployeeIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutHrisConnectionIDEmployeeIDResponse) GetHrisEmployee() *shared.HrisEmployee {
	if o == nil {
		return nil
	}
	return o.HrisEmployee
}

func (o *PutHrisConnectionIDEmployeeIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutHrisConnectionIDEmployeeIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
