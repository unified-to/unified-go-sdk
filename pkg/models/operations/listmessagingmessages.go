// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type ListMessagingMessagesRequest struct {
	// The channel ID to filter by
	ChannelID *string `queryParam:"style=form,explode=true,name=channel_id"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// The end date to filter by
	EndLe *string `queryParam:"style=form,explode=true,name=end_le"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	Limit  *float64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *float64 `queryParam:"style=form,explode=true,name=offset"`
	Order  *string  `queryParam:"style=form,explode=true,name=order"`
	// The parent ID to filter by
	ParentID *string `queryParam:"style=form,explode=true,name=parent_id"`
	// Query string to search. eg. email address or name
	Query *string `queryParam:"style=form,explode=true,name=query"`
	// Raw parameters to include in the 3rd-party request. Encoded as a URL component. eg. raw parameters: foo=bar&zoo=bar -> raw=foo%3Dbar%26zoo%3Dbar
	Raw  *string `queryParam:"style=form,explode=true,name=raw"`
	Sort *string `queryParam:"style=form,explode=true,name=sort"`
	// The start date to filter by
	StartGte *string `queryParam:"style=form,explode=true,name=start_gte"`
	// Return only results whose updated date is equal or greater to this value
	UpdatedGte *string `queryParam:"style=form,explode=true,name=updated_gte"`
}

func (o *ListMessagingMessagesRequest) GetChannelID() *string {
	if o == nil {
		return nil
	}
	return o.ChannelID
}

func (o *ListMessagingMessagesRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListMessagingMessagesRequest) GetEndLe() *string {
	if o == nil {
		return nil
	}
	return o.EndLe
}

func (o *ListMessagingMessagesRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *ListMessagingMessagesRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListMessagingMessagesRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListMessagingMessagesRequest) GetOrder() *string {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListMessagingMessagesRequest) GetParentID() *string {
	if o == nil {
		return nil
	}
	return o.ParentID
}

func (o *ListMessagingMessagesRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *ListMessagingMessagesRequest) GetRaw() *string {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *ListMessagingMessagesRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListMessagingMessagesRequest) GetStartGte() *string {
	if o == nil {
		return nil
	}
	return o.StartGte
}

func (o *ListMessagingMessagesRequest) GetUpdatedGte() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedGte
}

type ListMessagingMessagesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	MessagingMessages []shared.MessagingMessage
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListMessagingMessagesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListMessagingMessagesResponse) GetMessagingMessages() []shared.MessagingMessage {
	if o == nil {
		return nil
	}
	return o.MessagingMessages
}

func (o *ListMessagingMessagesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListMessagingMessagesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
