// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"net/http"
	"time"
)

type ListLmsClassesRequest struct {
	// ID of the connection
	ConnectionID string  `pathParam:"style=simple,explode=false,name=connection_id"`
	CourseID     *string `queryParam:"style=form,explode=true,name=course_id"`
	// Comma-delimited fields to return
	Fields     []string `queryParam:"style=form,explode=true,name=fields"`
	Limit      *float64 `queryParam:"style=form,explode=true,name=limit"`
	LocationID *string  `queryParam:"style=form,explode=true,name=location_id"`
	Offset     *float64 `queryParam:"style=form,explode=true,name=offset"`
	Order      *string  `queryParam:"style=form,explode=true,name=order"`
	// Query string to search. eg. email address or name
	Query *string `queryParam:"style=form,explode=true,name=query"`
	Sort  *string `queryParam:"style=form,explode=true,name=sort"`
	// Return only results whose updated date is equal or greater to this value
	UpdatedGte *time.Time `queryParam:"style=form,explode=true,name=updated_gte"`
}

func (l ListLmsClassesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListLmsClassesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListLmsClassesRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListLmsClassesRequest) GetCourseID() *string {
	if o == nil {
		return nil
	}
	return o.CourseID
}

func (o *ListLmsClassesRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *ListLmsClassesRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListLmsClassesRequest) GetLocationID() *string {
	if o == nil {
		return nil
	}
	return o.LocationID
}

func (o *ListLmsClassesRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListLmsClassesRequest) GetOrder() *string {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListLmsClassesRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *ListLmsClassesRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListLmsClassesRequest) GetUpdatedGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedGte
}

type ListLmsClassesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	LmsClasses []shared.LmsClass
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListLmsClassesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListLmsClassesResponse) GetLmsClasses() []shared.LmsClass {
	if o == nil {
		return nil
	}
	return o.LmsClasses
}

func (o *ListLmsClassesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListLmsClassesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
