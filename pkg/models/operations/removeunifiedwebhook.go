// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type RemoveUnifiedWebhookRequest struct {
	// ID of the Webhook
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *RemoveUnifiedWebhookRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type RemoveUnifiedWebhookResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	Res *string
}

func (o *RemoveUnifiedWebhookResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RemoveUnifiedWebhookResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RemoveUnifiedWebhookResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RemoveUnifiedWebhookResponse) GetRes() *string {
	if o == nil {
		return nil
	}
	return o.Res
}
