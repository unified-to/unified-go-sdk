// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateTicketingNoteRequest struct {
	TicketingNote *shared.TicketingNote `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Note
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateTicketingNoteRequest) GetTicketingNote() *shared.TicketingNote {
	if o == nil {
		return nil
	}
	return o.TicketingNote
}

func (o *UpdateTicketingNoteRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdateTicketingNoteRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *UpdateTicketingNoteRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateTicketingNoteResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	TicketingNote *shared.TicketingNote
}

func (o *UpdateTicketingNoteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateTicketingNoteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateTicketingNoteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateTicketingNoteResponse) GetTicketingNote() *shared.TicketingNote {
	if o == nil {
		return nil
	}
	return o.TicketingNote
}
