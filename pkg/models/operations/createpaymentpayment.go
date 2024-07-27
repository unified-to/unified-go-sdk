// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreatePaymentPaymentRequest struct {
	PaymentPayment *shared.PaymentPayment `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

func (o *CreatePaymentPaymentRequest) GetPaymentPayment() *shared.PaymentPayment {
	if o == nil {
		return nil
	}
	return o.PaymentPayment
}

func (o *CreatePaymentPaymentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type CreatePaymentPaymentResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	PaymentPayment *shared.PaymentPayment
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreatePaymentPaymentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreatePaymentPaymentResponse) GetPaymentPayment() *shared.PaymentPayment {
	if o == nil {
		return nil
	}
	return o.PaymentPayment
}

func (o *CreatePaymentPaymentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreatePaymentPaymentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
