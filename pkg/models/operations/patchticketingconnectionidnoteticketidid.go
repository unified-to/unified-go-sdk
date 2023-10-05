// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchTicketingConnectionIDNoteTicketIDIDRequest struct {
	TicketingNote *shared.TicketingNote `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Note
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// ID of the ticket
	TicketID string `pathParam:"style=simple,explode=false,name=ticket_id"`
}

func (o *PatchTicketingConnectionIDNoteTicketIDIDRequest) GetTicketingNote() *shared.TicketingNote {
	if o == nil {
		return nil
	}
	return o.TicketingNote
}

func (o *PatchTicketingConnectionIDNoteTicketIDIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchTicketingConnectionIDNoteTicketIDIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PatchTicketingConnectionIDNoteTicketIDIDRequest) GetTicketID() string {
	if o == nil {
		return ""
	}
	return o.TicketID
}

type PatchTicketingConnectionIDNoteTicketIDIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	TicketingNote *shared.TicketingNote
}

func (o *PatchTicketingConnectionIDNoteTicketIDIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchTicketingConnectionIDNoteTicketIDIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchTicketingConnectionIDNoteTicketIDIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchTicketingConnectionIDNoteTicketIDIDResponse) GetTicketingNote() *shared.TicketingNote {
	if o == nil {
		return nil
	}
	return o.TicketingNote
}