// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdatePaymentPaymentRequest struct {
	PaymentPayment *shared.PaymentPayment `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Payment
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdatePaymentPaymentRequest) GetPaymentPayment() *shared.PaymentPayment {
	if o == nil {
		return nil
	}
	return o.PaymentPayment
}

func (o *UpdatePaymentPaymentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdatePaymentPaymentRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdatePaymentPaymentResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	PaymentPayment *shared.PaymentPayment
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdatePaymentPaymentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdatePaymentPaymentResponse) GetPaymentPayment() *shared.PaymentPayment {
	if o == nil {
		return nil
	}
	return o.PaymentPayment
}

func (o *UpdatePaymentPaymentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdatePaymentPaymentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
