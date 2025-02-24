// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchTicketingTicketRequest struct {
	TicketingTicket shared.TicketingTicket `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Ticket
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchTicketingTicketRequest) GetTicketingTicket() shared.TicketingTicket {
	if o == nil {
		return shared.TicketingTicket{}
	}
	return o.TicketingTicket
}

func (o *PatchTicketingTicketRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchTicketingTicketRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *PatchTicketingTicketRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchTicketingTicketResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	TicketingTicket *shared.TicketingTicket
}

func (o *PatchTicketingTicketResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchTicketingTicketResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchTicketingTicketResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchTicketingTicketResponse) GetTicketingTicket() *shared.TicketingTicket {
	if o == nil {
		return nil
	}
	return o.TicketingTicket
}
