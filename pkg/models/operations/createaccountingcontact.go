// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateAccountingContactRequest struct {
	AccountingContact *shared.AccountingContact `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

func (o *CreateAccountingContactRequest) GetAccountingContact() *shared.AccountingContact {
	if o == nil {
		return nil
	}
	return o.AccountingContact
}

func (o *CreateAccountingContactRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type CreateAccountingContactResponse struct {
	// Successful
	AccountingContact *shared.AccountingContact
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateAccountingContactResponse) GetAccountingContact() *shared.AccountingContact {
	if o == nil {
		return nil
	}
	return o.AccountingContact
}

func (o *CreateAccountingContactResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateAccountingContactResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateAccountingContactResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
