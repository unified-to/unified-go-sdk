// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type GetMessagingChannelRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Channel
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetMessagingChannelRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetMessagingChannelRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *GetMessagingChannelRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetMessagingChannelResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	MessagingChannel *shared.MessagingChannel
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetMessagingChannelResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMessagingChannelResponse) GetMessagingChannel() *shared.MessagingChannel {
	if o == nil {
		return nil
	}
	return o.MessagingChannel
}

func (o *GetMessagingChannelResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMessagingChannelResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
