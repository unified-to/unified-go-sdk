// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"net/http"
	"time"
)

type ListCrmDealsRequest struct {
	CompanyID *string `queryParam:"style=form,explode=true,name=company_id"`
	// ID of the connection
	ConnectionID string  `pathParam:"style=simple,explode=false,name=connection_id"`
	ContactID    *string `queryParam:"style=form,explode=true,name=contact_id"`
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

func (l ListCrmDealsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListCrmDealsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListCrmDealsRequest) GetCompanyID() *string {
	if o == nil {
		return nil
	}
	return o.CompanyID
}

func (o *ListCrmDealsRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListCrmDealsRequest) GetContactID() *string {
	if o == nil {
		return nil
	}
	return o.ContactID
}

func (o *ListCrmDealsRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *ListCrmDealsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListCrmDealsRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListCrmDealsRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *ListCrmDealsRequest) GetUpdatedGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedGte
}

func (o *ListCrmDealsRequest) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

type ListCrmDealsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	CrmDeals []shared.CrmDeal
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListCrmDealsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListCrmDealsResponse) GetCrmDeals() []shared.CrmDeal {
	if o == nil {
		return nil
	}
	return o.CrmDeals
}

func (o *ListCrmDealsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListCrmDealsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
