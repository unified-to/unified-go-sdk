// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateCrmEventRequest struct {
	// An event represents an event, activity, or engagement and is always associated with a deal, contact, or company
	CrmEvent *shared.CrmEvent `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

func (o *CreateCrmEventRequest) GetCrmEvent() *shared.CrmEvent {
	if o == nil {
		return nil
	}
	return o.CrmEvent
}

func (o *CreateCrmEventRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type CreateCrmEventResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	CrmEvent *shared.CrmEvent
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateCrmEventResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateCrmEventResponse) GetCrmEvent() *shared.CrmEvent {
	if o == nil {
		return nil
	}
	return o.CrmEvent
}

func (o *CreateCrmEventResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateCrmEventResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
