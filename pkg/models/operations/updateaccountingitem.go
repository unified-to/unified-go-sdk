// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateAccountingItemRequest struct {
	// An item or product
	AccountingItem *shared.AccountingItem `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// ID of the Item
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateAccountingItemRequest) GetAccountingItem() *shared.AccountingItem {
	if o == nil {
		return nil
	}
	return o.AccountingItem
}

func (o *UpdateAccountingItemRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdateAccountingItemRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateAccountingItemResponse struct {
	// Successful
	AccountingItem *shared.AccountingItem
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateAccountingItemResponse) GetAccountingItem() *shared.AccountingItem {
	if o == nil {
		return nil
	}
	return o.AccountingItem
}

func (o *UpdateAccountingItemResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateAccountingItemResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateAccountingItemResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
