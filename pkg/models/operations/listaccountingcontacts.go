// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"net/http"
	"time"
)

type ListAccountingContactsRequest struct {
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	Limit  *float64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *float64 `queryParam:"style=form,explode=true,name=offset"`
	Order  *string  `queryParam:"style=form,explode=true,name=order"`
	// Query string to search. eg. email address or name
	Query *string `queryParam:"style=form,explode=true,name=query"`
	Sort  *string `queryParam:"style=form,explode=true,name=sort"`
	Type  *string `queryParam:"style=form,explode=true,name=type"`
	// Return only results whose updated date is equal or greater to this value
	UpdatedGte *time.Time `queryParam:"style=form,explode=true,name=updated_gte"`
}

func (l ListAccountingContactsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountingContactsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAccountingContactsRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListAccountingContactsRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *ListAccountingContactsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListAccountingContactsRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListAccountingContactsRequest) GetOrder() *string {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListAccountingContactsRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *ListAccountingContactsRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListAccountingContactsRequest) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ListAccountingContactsRequest) GetUpdatedGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedGte
}

type ListAccountingContactsResponse struct {
	// Successful
	AccountingContacts []shared.AccountingContact
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListAccountingContactsResponse) GetAccountingContacts() []shared.AccountingContact {
	if o == nil {
		return nil
	}
	return o.AccountingContacts
}

func (o *ListAccountingContactsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListAccountingContactsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListAccountingContactsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
