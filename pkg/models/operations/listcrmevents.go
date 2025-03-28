// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"net/http"
	"time"
)

type ListCrmEventsRequest struct {
	CompanyID *string `queryParam:"style=form,explode=true,name=company_id"`
	// ID of the connection
	ConnectionID string  `pathParam:"style=simple,explode=false,name=connection_id"`
	ContactID    *string `queryParam:"style=form,explode=true,name=contact_id"`
	DealID       *string `queryParam:"style=form,explode=true,name=deal_id"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	LeadID *string  `queryParam:"style=form,explode=true,name=lead_id"`
	Limit  *float64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *float64 `queryParam:"style=form,explode=true,name=offset"`
	Order  *string  `queryParam:"style=form,explode=true,name=order"`
	// Query string to search. eg. email address or name
	Query *string `queryParam:"style=form,explode=true,name=query"`
	Sort  *string `queryParam:"style=form,explode=true,name=sort"`
	Type  *string `queryParam:"style=form,explode=true,name=type"`
	// Return only results whose updated date is equal or greater to this value
	UpdatedGte *time.Time `queryParam:"style=form,explode=true,name=updated_gte"`
	UserID     *string    `queryParam:"style=form,explode=true,name=user_id"`
}

func (l ListCrmEventsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListCrmEventsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListCrmEventsRequest) GetCompanyID() *string {
	if o == nil {
		return nil
	}
	return o.CompanyID
}

func (o *ListCrmEventsRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListCrmEventsRequest) GetContactID() *string {
	if o == nil {
		return nil
	}
	return o.ContactID
}

func (o *ListCrmEventsRequest) GetDealID() *string {
	if o == nil {
		return nil
	}
	return o.DealID
}

func (o *ListCrmEventsRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *ListCrmEventsRequest) GetLeadID() *string {
	if o == nil {
		return nil
	}
	return o.LeadID
}

func (o *ListCrmEventsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListCrmEventsRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListCrmEventsRequest) GetOrder() *string {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListCrmEventsRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *ListCrmEventsRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListCrmEventsRequest) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ListCrmEventsRequest) GetUpdatedGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedGte
}

func (o *ListCrmEventsRequest) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

type ListCrmEventsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	CrmEvents []shared.CrmEvent
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListCrmEventsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListCrmEventsResponse) GetCrmEvents() []shared.CrmEvent {
	if o == nil {
		return nil
	}
	return o.CrmEvents
}

func (o *ListCrmEventsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListCrmEventsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
