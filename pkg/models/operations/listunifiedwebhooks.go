// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"net/http"
	"time"
)

type ListUnifiedWebhooksSecurity struct {
	Jwt string `security:"scheme,type=apiKey,subtype=header,name=authorization"`
}

func (o *ListUnifiedWebhooksSecurity) GetJwt() string {
	if o == nil {
		return ""
	}
	return o.Jwt
}

type ListUnifiedWebhooksRequest struct {
	Env   *string  `queryParam:"style=form,explode=true,name=env"`
	Limit *float64 `queryParam:"style=form,explode=true,name=limit"`
	// Filter the results for webhooks for only this object
	Object *string  `queryParam:"style=form,explode=true,name=object"`
	Offset *float64 `queryParam:"style=form,explode=true,name=offset"`
	Order  *string  `queryParam:"style=form,explode=true,name=order"`
	Sort   *string  `queryParam:"style=form,explode=true,name=sort"`
	// Return only results whose updated date is equal or greater to this value
	UpdatedGte *time.Time `queryParam:"style=form,explode=true,name=updated_gte"`
}

func (l ListUnifiedWebhooksRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListUnifiedWebhooksRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListUnifiedWebhooksRequest) GetEnv() *string {
	if o == nil {
		return nil
	}
	return o.Env
}

func (o *ListUnifiedWebhooksRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListUnifiedWebhooksRequest) GetObject() *string {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *ListUnifiedWebhooksRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListUnifiedWebhooksRequest) GetOrder() *string {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListUnifiedWebhooksRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListUnifiedWebhooksRequest) GetUpdatedGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedGte
}

type ListUnifiedWebhooksResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	Webhooks []shared.Webhook
}

func (o *ListUnifiedWebhooksResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListUnifiedWebhooksResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListUnifiedWebhooksResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListUnifiedWebhooksResponse) GetWebhooks() []shared.Webhook {
	if o == nil {
		return nil
	}
	return o.Webhooks
}
