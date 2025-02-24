// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateTicketingTicketRequest struct {
	TicketingTicket shared.TicketingTicket `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Ticket
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateTicketingTicketRequest) GetTicketingTicket() shared.TicketingTicket {
	if o == nil {
		return shared.TicketingTicket{}
	}
	return o.TicketingTicket
}

func (o *UpdateTicketingTicketRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdateTicketingTicketRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *UpdateTicketingTicketRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateTicketingTicketResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	TicketingTicket *shared.TicketingTicket
}

func (o *UpdateTicketingTicketResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateTicketingTicketResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateTicketingTicketResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateTicketingTicketResponse) GetTicketingTicket() *shared.TicketingTicket {
	if o == nil {
		return nil
	}
	return o.TicketingTicket
}
