// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PatchUnifiedWebhookTriggerSecurity struct {
	Jwt string `security:"scheme,type=apiKey,subtype=header,name=authorization"`
}

func (o *PatchUnifiedWebhookTriggerSecurity) GetJwt() string {
	if o == nil {
		return ""
	}
	return o.Jwt
}

type PatchUnifiedWebhookTriggerRequest struct {
	// ID of the Webhook
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchUnifiedWebhookTriggerRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchUnifiedWebhookTriggerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	Res *string
}

func (o *PatchUnifiedWebhookTriggerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchUnifiedWebhookTriggerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchUnifiedWebhookTriggerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchUnifiedWebhookTriggerResponse) GetRes() *string {
	if o == nil {
		return nil
	}
	return o.Res
}
