// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type PatchAccountingContactRequest struct {
	AccountingContact *shared.AccountingContact `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Contact
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PatchAccountingContactRequest) GetAccountingContact() *shared.AccountingContact {
	if o == nil {
		return nil
	}
	return o.AccountingContact
}

func (o *PatchAccountingContactRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchAccountingContactRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PatchAccountingContactResponse struct {
	// Successful
	AccountingContact *shared.AccountingContact
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PatchAccountingContactResponse) GetAccountingContact() *shared.AccountingContact {
	if o == nil {
		return nil
	}
	return o.AccountingContact
}

func (o *PatchAccountingContactResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchAccountingContactResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchAccountingContactResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
