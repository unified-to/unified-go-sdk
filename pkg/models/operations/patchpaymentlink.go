// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchPaymentLinkRequest struct {
	PaymentLink *shared.PaymentLink `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Link
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchPaymentLinkRequest) GetPaymentLink() *shared.PaymentLink {
	if o == nil {
		return nil
	}
	return o.PaymentLink
}

func (o *PatchPaymentLinkRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchPaymentLinkRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchPaymentLinkResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	PaymentLink *shared.PaymentLink
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PatchPaymentLinkResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchPaymentLinkResponse) GetPaymentLink() *shared.PaymentLink {
	if o == nil {
		return nil
	}
	return o.PaymentLink
}

func (o *PatchPaymentLinkResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchPaymentLinkResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
