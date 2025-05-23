// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreatePaymentSubscriptionRequest struct {
	PaymentSubscription shared.PaymentSubscription `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// Raw parameters to include in the 3rd-party request. Encoded as a URL component. eg. raw parameters: foo=bar&zoo=bar -> raw=foo%3Dbar%26zoo%3Dbar
	Raw *string `queryParam:"style=form,explode=true,name=raw"`
}

func (o *CreatePaymentSubscriptionRequest) GetPaymentSubscription() shared.PaymentSubscription {
	if o == nil {
		return shared.PaymentSubscription{}
	}
	return o.PaymentSubscription
}

func (o *CreatePaymentSubscriptionRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreatePaymentSubscriptionRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *CreatePaymentSubscriptionRequest) GetRaw() *string {
	if o == nil {
		return nil
	}
	return o.Raw
}

type CreatePaymentSubscriptionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	PaymentSubscription *shared.PaymentSubscription
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreatePaymentSubscriptionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreatePaymentSubscriptionResponse) GetPaymentSubscription() *shared.PaymentSubscription {
	if o == nil {
		return nil
	}
	return o.PaymentSubscription
}

func (o *CreatePaymentSubscriptionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreatePaymentSubscriptionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
