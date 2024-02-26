// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type GetUnifiedWebhookSecurity struct {
	Jwt string `security:"scheme,type=apiKey,subtype=header,name=authorization"`
}

func (o *GetUnifiedWebhookSecurity) GetJwt() string {
	if o == nil {
		return ""
	}
	return o.Jwt
}

type GetUnifiedWebhookRequest struct {
	// ID of the Webhook
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetUnifiedWebhookRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetUnifiedWebhookResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	Webhook *shared.Webhook
}

func (o *GetUnifiedWebhookResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetUnifiedWebhookResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetUnifiedWebhookResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetUnifiedWebhookResponse) GetWebhook() *shared.Webhook {
	if o == nil {
		return nil
	}
	return o.Webhook
}
