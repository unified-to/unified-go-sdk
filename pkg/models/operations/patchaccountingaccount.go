// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchAccountingAccountRequest struct {
	// A user's bank account
	AccountingAccount *shared.AccountingAccount `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Account
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchAccountingAccountRequest) GetAccountingAccount() *shared.AccountingAccount {
	if o == nil {
		return nil
	}
	return o.AccountingAccount
}

func (o *PatchAccountingAccountRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchAccountingAccountRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchAccountingAccountResponse struct {
	// Successful
	AccountingAccount *shared.AccountingAccount
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PatchAccountingAccountResponse) GetAccountingAccount() *shared.AccountingAccount {
	if o == nil {
		return nil
	}
	return o.AccountingAccount
}

func (o *PatchAccountingAccountResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchAccountingAccountResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchAccountingAccountResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
