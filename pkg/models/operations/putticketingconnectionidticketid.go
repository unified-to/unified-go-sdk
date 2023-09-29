// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PutTicketingConnectionIDTicketIDRequest struct {
	TicketingTicket *shared.TicketingTicket `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Ticket
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PutTicketingConnectionIDTicketIDRequest) GetTicketingTicket() *shared.TicketingTicket {
	if o == nil {
		return nil
	}
	return o.TicketingTicket
}

func (o *PutTicketingConnectionIDTicketIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PutTicketingConnectionIDTicketIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PutTicketingConnectionIDTicketIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	TicketingTicket *shared.TicketingTicket
}

func (o *PutTicketingConnectionIDTicketIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutTicketingConnectionIDTicketIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutTicketingConnectionIDTicketIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PutTicketingConnectionIDTicketIDResponse) GetTicketingTicket() *shared.TicketingTicket {
	if o == nil {
		return nil
	}
	return o.TicketingTicket
}
