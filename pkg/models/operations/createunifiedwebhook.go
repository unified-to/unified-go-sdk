// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateUnifiedWebhookRequest struct {
	Webhook *shared.Webhook `request:"mediaType=application/json"`
	// When set, all of the existing data will sent back to your server.
	IncludeAll *bool `queryParam:"style=form,explode=true,name=include_all"`
}

func (o *CreateUnifiedWebhookRequest) GetWebhook() *shared.Webhook {
	if o == nil {
		return nil
	}
	return o.Webhook
}

func (o *CreateUnifiedWebhookRequest) GetIncludeAll() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeAll
}

type CreateUnifiedWebhookResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	Webhook *shared.Webhook
}

func (o *CreateUnifiedWebhookResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateUnifiedWebhookResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateUnifiedWebhookResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateUnifiedWebhookResponse) GetWebhook() *shared.Webhook {
	if o == nil {
		return nil
	}
	return o.Webhook
}
