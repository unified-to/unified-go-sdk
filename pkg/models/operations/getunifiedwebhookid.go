// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type GetUnifiedWebhookIDSecurity struct {
	Jwt string `security:"scheme,type=apiKey,subtype=header,name=authorization"`
}

func (o *GetUnifiedWebhookIDSecurity) GetJwt() string {
	if o == nil {
		return ""
	}
	return o.Jwt
}

type GetUnifiedWebhookIDRequest struct {
	// ID of the Webhook
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetUnifiedWebhookIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetUnifiedWebhookIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful
	Webhook *shared.Webhook
}

func (o *GetUnifiedWebhookIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetUnifiedWebhookIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetUnifiedWebhookIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetUnifiedWebhookIDResponse) GetWebhook() *shared.Webhook {
	if o == nil {
		return nil
	}
	return o.Webhook
}
