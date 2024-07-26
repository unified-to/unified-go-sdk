// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type GetPaymentLinkRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the Link
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetPaymentLinkRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetPaymentLinkRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *GetPaymentLinkRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetPaymentLinkResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	PaymentLink *shared.PaymentLink
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetPaymentLinkResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPaymentLinkResponse) GetPaymentLink() *shared.PaymentLink {
	if o == nil {
		return nil
	}
	return o.PaymentLink
}

func (o *GetPaymentLinkResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPaymentLinkResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
