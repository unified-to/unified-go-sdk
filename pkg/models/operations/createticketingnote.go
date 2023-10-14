// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateTicketingNoteRequest struct {
	TicketingNote *shared.TicketingNote `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the ticket
	TicketID string `pathParam:"style=simple,explode=false,name=ticket_id"`
}

func (o *CreateTicketingNoteRequest) GetTicketingNote() *shared.TicketingNote {
	if o == nil {
		return nil
	}
	return o.TicketingNote
}

func (o *CreateTicketingNoteRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateTicketingNoteRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *CreateTicketingNoteRequest) GetTicketID() string {
	if o == nil {
		return ""
	}
	return o.TicketID
}

type CreateTicketingNoteResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	TicketingNote *shared.TicketingNote
}

func (o *CreateTicketingNoteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateTicketingNoteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateTicketingNoteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateTicketingNoteResponse) GetTicketingNote() *shared.TicketingNote {
	if o == nil {
		return nil
	}
	return o.TicketingNote
}
