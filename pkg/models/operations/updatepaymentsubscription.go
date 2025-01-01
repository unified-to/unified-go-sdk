// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdatePaymentSubscriptionRequest struct {
	PaymentSubscription *shared.PaymentSubscription `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Subscription
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdatePaymentSubscriptionRequest) GetPaymentSubscription() *shared.PaymentSubscription {
	if o == nil {
		return nil
	}
	return o.PaymentSubscription
}

func (o *UpdatePaymentSubscriptionRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdatePaymentSubscriptionRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *UpdatePaymentSubscriptionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdatePaymentSubscriptionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	PaymentSubscription *shared.PaymentSubscription
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdatePaymentSubscriptionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdatePaymentSubscriptionResponse) GetPaymentSubscription() *shared.PaymentSubscription {
	if o == nil {
		return nil
	}
	return o.PaymentSubscription
}

func (o *UpdatePaymentSubscriptionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdatePaymentSubscriptionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
