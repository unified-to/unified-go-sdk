// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreatePaymentLinkRequest struct {
	PaymentLink *shared.PaymentLink `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

func (o *CreatePaymentLinkRequest) GetPaymentLink() *shared.PaymentLink {
	if o == nil {
		return nil
	}
	return o.PaymentLink
}

func (o *CreatePaymentLinkRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type CreatePaymentLinkResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	PaymentLink *shared.PaymentLink
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreatePaymentLinkResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreatePaymentLinkResponse) GetPaymentLink() *shared.PaymentLink {
	if o == nil {
		return nil
	}
	return o.PaymentLink
}

func (o *CreatePaymentLinkResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreatePaymentLinkResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
