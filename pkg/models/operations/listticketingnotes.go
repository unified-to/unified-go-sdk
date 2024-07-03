// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"net/http"
	"time"
)

type ListTicketingNotesRequest struct {
	// ID of the connection
	ConnectionID string  `pathParam:"style=simple,explode=false,name=connection_id"`
	CustomerID   *string `queryParam:"style=form,explode=true,name=customer_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	Limit  *float64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *float64 `queryParam:"style=form,explode=true,name=offset"`
	// Query string to search. eg. email address or name
	Query    *string `queryParam:"style=form,explode=true,name=query"`
	TicketID *string `queryParam:"style=form,explode=true,name=ticket_id"`
	// Return only results whose updated date is equal or greater to this value
	UpdatedGte *time.Time `queryParam:"style=form,explode=true,name=updated_gte"`
}

func (l ListTicketingNotesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListTicketingNotesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListTicketingNotesRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListTicketingNotesRequest) GetCustomerID() *string {
	if o == nil {
		return nil
	}
	return o.CustomerID
}

func (o *ListTicketingNotesRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *ListTicketingNotesRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListTicketingNotesRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListTicketingNotesRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *ListTicketingNotesRequest) GetTicketID() *string {
	if o == nil {
		return nil
	}
	return o.TicketID
}

func (o *ListTicketingNotesRequest) GetUpdatedGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedGte
}

type ListTicketingNotesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	TicketingNotes []shared.TicketingNote
}

func (o *ListTicketingNotesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListTicketingNotesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListTicketingNotesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListTicketingNotesResponse) GetTicketingNotes() []shared.TicketingNote {
	if o == nil {
		return nil
	}
	return o.TicketingNotes
}
