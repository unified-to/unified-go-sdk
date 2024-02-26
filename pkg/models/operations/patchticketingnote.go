// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchTicketingNoteSecurity struct {
	Jwt string `security:"scheme,type=apiKey,subtype=header,name=authorization"`
}

func (o *PatchTicketingNoteSecurity) GetJwt() string {
	if o == nil {
		return ""
	}
	return o.Jwt
}

type PatchTicketingNoteRequest struct {
	TicketingNote *shared.TicketingNote `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Note
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchTicketingNoteRequest) GetTicketingNote() *shared.TicketingNote {
	if o == nil {
		return nil
	}
	return o.TicketingNote
}

func (o *PatchTicketingNoteRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchTicketingNoteRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchTicketingNoteResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	TicketingNote *shared.TicketingNote
}

func (o *PatchTicketingNoteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchTicketingNoteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchTicketingNoteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchTicketingNoteResponse) GetTicketingNote() *shared.TicketingNote {
	if o == nil {
		return nil
	}
	return o.TicketingNote
}
