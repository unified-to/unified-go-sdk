// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"net/http"
	"time"
)

type GetUnifiedWebhookRequest struct {
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

func (g GetUnifiedWebhookRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetUnifiedWebhookRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetUnifiedWebhookRequest) GetEnv() *string {
	if o == nil {
		return nil
	}
	return o.Env
}

func (o *GetUnifiedWebhookRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetUnifiedWebhookRequest) GetObject() *string {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *GetUnifiedWebhookRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetUnifiedWebhookRequest) GetOrder() *string {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *GetUnifiedWebhookRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *GetUnifiedWebhookRequest) GetUpdatedGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedGte
}

type GetUnifiedWebhookResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful
	Webhooks []shared.Webhook
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

func (o *GetUnifiedWebhookResponse) GetWebhooks() []shared.Webhook {
	if o == nil {
		return nil
	}
	return o.Webhooks
}
