// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateAccountingInvoiceSecurity struct {
	Jwt string `security:"scheme,type=apiKey,subtype=header,name=authorization"`
}

func (o *CreateAccountingInvoiceSecurity) GetJwt() string {
	if o == nil {
		return ""
	}
	return o.Jwt
}

type CreateAccountingInvoiceRequest struct {
	AccountingInvoice *shared.AccountingInvoice `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

func (o *CreateAccountingInvoiceRequest) GetAccountingInvoice() *shared.AccountingInvoice {
	if o == nil {
		return nil
	}
	return o.AccountingInvoice
}

func (o *CreateAccountingInvoiceRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type CreateAccountingInvoiceResponse struct {
	// Successful
	AccountingInvoice *shared.AccountingInvoice
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateAccountingInvoiceResponse) GetAccountingInvoice() *shared.AccountingInvoice {
	if o == nil {
		return nil
	}
	return o.AccountingInvoice
}

func (o *CreateAccountingInvoiceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateAccountingInvoiceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateAccountingInvoiceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
