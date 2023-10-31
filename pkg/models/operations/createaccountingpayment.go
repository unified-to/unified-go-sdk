// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateAccountingPaymentRequest struct {
	AccountingPayment *shared.AccountingPayment `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

func (o *CreateAccountingPaymentRequest) GetAccountingPayment() *shared.AccountingPayment {
	if o == nil {
		return nil
	}
	return o.AccountingPayment
}

func (o *CreateAccountingPaymentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type CreateAccountingPaymentResponse struct {
	// Successful
	AccountingPayment *shared.AccountingPayment
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateAccountingPaymentResponse) GetAccountingPayment() *shared.AccountingPayment {
	if o == nil {
		return nil
	}
	return o.AccountingPayment
}

func (o *CreateAccountingPaymentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateAccountingPaymentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateAccountingPaymentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
