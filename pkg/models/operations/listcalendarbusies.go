// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"net/http"
	"time"
)

type ListCalendarBusiesRequest struct {
	CalendarID *string `queryParam:"style=form,explode=true,name=calendar_id"`
	// ID of the connection
	ConnectionID string  `pathParam:"style=simple,explode=false,name=connection_id"`
	EndLe        *string `queryParam:"style=form,explode=true,name=end_le"`
	// Comma-delimited fields to return
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	Limit  *float64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *float64 `queryParam:"style=form,explode=true,name=offset"`
	Order  *string  `queryParam:"style=form,explode=true,name=order"`
	// Query string to search. eg. email address or name
	Query    *string `queryParam:"style=form,explode=true,name=query"`
	Sort     *string `queryParam:"style=form,explode=true,name=sort"`
	StartGte *string `queryParam:"style=form,explode=true,name=start_gte"`
	// Return only results whose updated date is equal or greater to this value
	UpdatedGte *time.Time `queryParam:"style=form,explode=true,name=updated_gte"`
}

func (l ListCalendarBusiesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListCalendarBusiesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListCalendarBusiesRequest) GetCalendarID() *string {
	if o == nil {
		return nil
	}
	return o.CalendarID
}

func (o *ListCalendarBusiesRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListCalendarBusiesRequest) GetEndLe() *string {
	if o == nil {
		return nil
	}
	return o.EndLe
}

func (o *ListCalendarBusiesRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *ListCalendarBusiesRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListCalendarBusiesRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListCalendarBusiesRequest) GetOrder() *string {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListCalendarBusiesRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *ListCalendarBusiesRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListCalendarBusiesRequest) GetStartGte() *string {
	if o == nil {
		return nil
	}
	return o.StartGte
}

func (o *ListCalendarBusiesRequest) GetUpdatedGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedGte
}

type ListCalendarBusiesResponse struct {
	// Successful
	CalendarBusies []shared.CalendarBusy
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListCalendarBusiesResponse) GetCalendarBusies() []shared.CalendarBusy {
	if o == nil {
		return nil
	}
	return o.CalendarBusies
}

func (o *ListCalendarBusiesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListCalendarBusiesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListCalendarBusiesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
