// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type GetMessagingMessageRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Message
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetMessagingMessageRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetMessagingMessageRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *GetMessagingMessageRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetMessagingMessageResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	MessagingMessage *shared.MessagingMessage
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetMessagingMessageResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMessagingMessageResponse) GetMessagingMessage() *shared.MessagingMessage {
	if o == nil {
		return nil
	}
	return o.MessagingMessage
}

func (o *GetMessagingMessageResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMessagingMessageResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
