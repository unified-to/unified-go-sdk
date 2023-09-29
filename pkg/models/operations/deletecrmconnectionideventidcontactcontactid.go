// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type DeleteCrmConnectionIDEventIDContactContactIDRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the contact
	ContactID string `pathParam:"style=simple,explode=false,name=contact_id"`
	// ID of the Event
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteCrmConnectionIDEventIDContactContactIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *DeleteCrmConnectionIDEventIDContactContactIDRequest) GetContactID() string {
	if o == nil {
		return ""
	}
	return o.ContactID
}

func (o *DeleteCrmConnectionIDEventIDContactContactIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteCrmConnectionIDEventIDContactContactIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	CrmEvent *shared.CrmEvent
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteCrmConnectionIDEventIDContactContactIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteCrmConnectionIDEventIDContactContactIDResponse) GetCrmEvent() *shared.CrmEvent {
	if o == nil {
		return nil
	}
	return o.CrmEvent
}

func (o *DeleteCrmConnectionIDEventIDContactContactIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteCrmConnectionIDEventIDContactContactIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
