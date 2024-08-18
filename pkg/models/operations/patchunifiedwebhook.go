// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchUnifiedWebhookRequest struct {
	// A webhook is used to POST new/updated information to your server.
	Webhook *shared.Webhook `request:"mediaType=application/json"`
	// ID of the Webhook
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchUnifiedWebhookRequest) GetWebhook() *shared.Webhook {
	if o == nil {
		return nil
	}
	return o.Webhook
}

func (o *PatchUnifiedWebhookRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchUnifiedWebhookResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	Webhook *shared.Webhook
}

func (o *PatchUnifiedWebhookResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchUnifiedWebhookResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchUnifiedWebhookResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchUnifiedWebhookResponse) GetWebhook() *shared.Webhook {
	if o == nil {
		return nil
	}
	return o.Webhook
}
