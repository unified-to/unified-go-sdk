// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"net/http"
	"time"
)

type GetHrisConnectionIDEmployeeRequest struct {
	// ID of the connection
	ConnectionID string   `pathParam:"style=simple,explode=false,name=connection_id"`
	Limit        *float64 `queryParam:"style=form,explode=true,name=limit"`
	Offset       *float64 `queryParam:"style=form,explode=true,name=offset"`
	Order        *string  `queryParam:"style=form,explode=true,name=order"`
	// Query string to search. eg. email address or name
	Query *string `queryParam:"style=form,explode=true,name=query"`
	Sort  *string `queryParam:"style=form,explode=true,name=sort"`
	// Return only results whose updated date is equal or greater to this value
	UpdatedGte *time.Time `queryParam:"style=form,explode=true,name=updated_gte"`
}

func (g GetHrisConnectionIDEmployeeRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetHrisConnectionIDEmployeeRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetHrisConnectionIDEmployeeRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetHrisConnectionIDEmployeeRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetHrisConnectionIDEmployeeRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetHrisConnectionIDEmployeeRequest) GetOrder() *string {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *GetHrisConnectionIDEmployeeRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *GetHrisConnectionIDEmployeeRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *GetHrisConnectionIDEmployeeRequest) GetUpdatedGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedGte
}

type GetHrisConnectionIDEmployeeResponse struct {
	ContentType string
	// Successful
	HrisEmployees []shared.HrisEmployee
	StatusCode    int
	RawResponse   *http.Response
}

func (o *GetHrisConnectionIDEmployeeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetHrisConnectionIDEmployeeResponse) GetHrisEmployees() []shared.HrisEmployee {
	if o == nil {
		return nil
	}
	return o.HrisEmployees
}

func (o *GetHrisConnectionIDEmployeeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetHrisConnectionIDEmployeeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
