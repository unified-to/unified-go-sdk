// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type GetTicketingConnectionIDCustomerIDRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Customer
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetTicketingConnectionIDCustomerIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetTicketingConnectionIDCustomerIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetTicketingConnectionIDCustomerIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	TicketingCustomer *shared.TicketingCustomer
}

func (o *GetTicketingConnectionIDCustomerIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTicketingConnectionIDCustomerIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTicketingConnectionIDCustomerIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTicketingConnectionIDCustomerIDResponse) GetTicketingCustomer() *shared.TicketingCustomer {
	if o == nil {
		return nil
	}
	return o.TicketingCustomer
}