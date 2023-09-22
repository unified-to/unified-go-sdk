// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteUnifiedWebhookIDRequest struct {
	// ID of the Webhook
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteUnifiedWebhookIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteUnifiedWebhookIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful
	DeleteUnifiedWebhookIDDefaultApplicationJSONString *string
}

func (o *DeleteUnifiedWebhookIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteUnifiedWebhookIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteUnifiedWebhookIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteUnifiedWebhookIDResponse) GetDeleteUnifiedWebhookIDDefaultApplicationJSONString() *string {
	if o == nil {
		return nil
	}
	return o.DeleteUnifiedWebhookIDDefaultApplicationJSONString
}
