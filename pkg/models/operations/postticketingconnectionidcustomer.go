// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PostTicketingConnectionIDCustomerRequest struct {
	TicketingCustomer *shared.TicketingCustomer `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

func (o *PostTicketingConnectionIDCustomerRequest) GetTicketingCustomer() *shared.TicketingCustomer {
	if o == nil {
		return nil
	}
	return o.TicketingCustomer
}

func (o *PostTicketingConnectionIDCustomerRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type PostTicketingConnectionIDCustomerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	TicketingCustomer *shared.TicketingCustomer
}

func (o *PostTicketingConnectionIDCustomerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostTicketingConnectionIDCustomerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostTicketingConnectionIDCustomerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostTicketingConnectionIDCustomerResponse) GetTicketingCustomer() *shared.TicketingCustomer {
	if o == nil {
		return nil
	}
	return o.TicketingCustomer
}
