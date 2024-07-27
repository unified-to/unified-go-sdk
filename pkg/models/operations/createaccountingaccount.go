// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateAccountingAccountRequest struct {
	// Chart of accounts
	AccountingAccount *shared.AccountingAccount `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

func (o *CreateAccountingAccountRequest) GetAccountingAccount() *shared.AccountingAccount {
	if o == nil {
		return nil
	}
	return o.AccountingAccount
}

func (o *CreateAccountingAccountRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type CreateAccountingAccountResponse struct {
	// Successful
	AccountingAccount *shared.AccountingAccount
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateAccountingAccountResponse) GetAccountingAccount() *shared.AccountingAccount {
	if o == nil {
		return nil
	}
	return o.AccountingAccount
}

func (o *CreateAccountingAccountResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateAccountingAccountResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateAccountingAccountResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
