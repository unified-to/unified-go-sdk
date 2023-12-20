// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateAccountingTransactionRequest struct {
	AccountingTransaction *shared.AccountingTransaction `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

func (o *CreateAccountingTransactionRequest) GetAccountingTransaction() *shared.AccountingTransaction {
	if o == nil {
		return nil
	}
	return o.AccountingTransaction
}

func (o *CreateAccountingTransactionRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type CreateAccountingTransactionResponse struct {
	// Successful
	AccountingTransaction *shared.AccountingTransaction
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateAccountingTransactionResponse) GetAccountingTransaction() *shared.AccountingTransaction {
	if o == nil {
		return nil
	}
	return o.AccountingTransaction
}

func (o *CreateAccountingTransactionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateAccountingTransactionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateAccountingTransactionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}