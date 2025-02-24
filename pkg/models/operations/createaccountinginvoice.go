// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateAccountingInvoiceRequest struct {
	AccountingInvoice shared.AccountingInvoice `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
}

func (o *CreateAccountingInvoiceRequest) GetAccountingInvoice() shared.AccountingInvoice {
	if o == nil {
		return shared.AccountingInvoice{}
	}
	return o.AccountingInvoice
}

func (o *CreateAccountingInvoiceRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateAccountingInvoiceRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
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
