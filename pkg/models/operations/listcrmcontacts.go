// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"net/http"
	"time"
)

type ListCrmContactsRequest struct {
	CompanyID *string `queryParam:"style=form,explode=true,name=company_id"`
	// ID of the connection
	ConnectionID string  `pathParam:"style=simple,explode=false,name=connection_id"`
	DealID       *string `queryParam:"style=form,explode=true,name=deal_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	Limit  *float64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *float64 `queryParam:"style=form,explode=true,name=offset"`
	// Query string to search. eg. email address or name
	Query *string `queryParam:"style=form,explode=true,name=query"`
	// Return only results whose updated date is equal or greater to this value
	UpdatedGte *time.Time `queryParam:"style=form,explode=true,name=updated_gte"`
	UserID     *string    `queryParam:"style=form,explode=true,name=user_id"`
}

func (l ListCrmContactsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListCrmContactsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListCrmContactsRequest) GetCompanyID() *string {
	if o == nil {
		return nil
	}
	return o.CompanyID
}

func (o *ListCrmContactsRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListCrmContactsRequest) GetDealID() *string {
	if o == nil {
		return nil
	}
	return o.DealID
}

func (o *ListCrmContactsRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *ListCrmContactsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListCrmContactsRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListCrmContactsRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *ListCrmContactsRequest) GetUpdatedGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedGte
}

func (o *ListCrmContactsRequest) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

type ListCrmContactsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	CrmContacts []shared.CrmContact
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListCrmContactsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListCrmContactsResponse) GetCrmContacts() []shared.CrmContact {
	if o == nil {
		return nil
	}
	return o.CrmContacts
}

func (o *ListCrmContactsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListCrmContactsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
