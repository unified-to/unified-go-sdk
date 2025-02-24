// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateMessagingMessageRequest struct {
	MessagingMessage shared.MessagingMessage `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
}

func (o *CreateMessagingMessageRequest) GetMessagingMessage() shared.MessagingMessage {
	if o == nil {
		return shared.MessagingMessage{}
	}
	return o.MessagingMessage
}

func (o *CreateMessagingMessageRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateMessagingMessageRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

type CreateMessagingMessageResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	MessagingMessage *shared.MessagingMessage
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateMessagingMessageResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateMessagingMessageResponse) GetMessagingMessage() *shared.MessagingMessage {
	if o == nil {
		return nil
	}
	return o.MessagingMessage
}

func (o *CreateMessagingMessageResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateMessagingMessageResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
