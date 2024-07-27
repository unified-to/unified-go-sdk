// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type GetPaymentPayoutRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Payout
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetPaymentPayoutRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetPaymentPayoutRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *GetPaymentPayoutRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetPaymentPayoutResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	PaymentPayout *shared.PaymentPayout
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetPaymentPayoutResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPaymentPayoutResponse) GetPaymentPayout() *shared.PaymentPayout {
	if o == nil {
		return nil
	}
	return o.PaymentPayout
}

func (o *GetPaymentPayoutResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPaymentPayoutResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
