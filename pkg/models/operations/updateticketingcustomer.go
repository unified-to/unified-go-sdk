// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateTicketingCustomerRequest struct {
	TicketingCustomer *shared.TicketingCustomer `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Customer
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateTicketingCustomerRequest) GetTicketingCustomer() *shared.TicketingCustomer {
	if o == nil {
		return nil
	}
	return o.TicketingCustomer
}

func (o *UpdateTicketingCustomerRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdateTicketingCustomerRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *UpdateTicketingCustomerRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateTicketingCustomerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	TicketingCustomer *shared.TicketingCustomer
}

func (o *UpdateTicketingCustomerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateTicketingCustomerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateTicketingCustomerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateTicketingCustomerResponse) GetTicketingCustomer() *shared.TicketingCustomer {
	if o == nil {
		return nil
	}
	return o.TicketingCustomer
}
