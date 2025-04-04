// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"net/http"
	"time"
)

type ListLmsInstructorsRequest struct {
	ClassID   *string `queryParam:"style=form,explode=true,name=class_id"`
	CompanyID *string `queryParam:"style=form,explode=true,name=company_id"`
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
	Query     *string `queryParam:"style=form,explode=true,name=query"`
	SessionID *string `queryParam:"style=form,explode=true,name=session_id"`
	Sort      *string `queryParam:"style=form,explode=true,name=sort"`
	// Return only results whose updated date is equal or greater to this value
	UpdatedGte *time.Time `queryParam:"style=form,explode=true,name=updated_gte"`
}

func (l ListLmsInstructorsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListLmsInstructorsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListLmsInstructorsRequest) GetClassID() *string {
	if o == nil {
		return nil
	}
	return o.ClassID
}

func (o *ListLmsInstructorsRequest) GetCompanyID() *string {
	if o == nil {
		return nil
	}
	return o.CompanyID
}

func (o *ListLmsInstructorsRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListLmsInstructorsRequest) GetCourseID() *string {
	if o == nil {
		return nil
	}
	return o.CourseID
}

func (o *ListLmsInstructorsRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *ListLmsInstructorsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListLmsInstructorsRequest) GetLocationID() *string {
	if o == nil {
		return nil
	}
	return o.LocationID
}

func (o *ListLmsInstructorsRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListLmsInstructorsRequest) GetOrder() *string {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListLmsInstructorsRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *ListLmsInstructorsRequest) GetSessionID() *string {
	if o == nil {
		return nil
	}
	return o.SessionID
}

func (o *ListLmsInstructorsRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListLmsInstructorsRequest) GetUpdatedGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedGte
}

type ListLmsInstructorsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	LmsInstructors []shared.LmsInstructor
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListLmsInstructorsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListLmsInstructorsResponse) GetLmsInstructors() []shared.LmsInstructor {
	if o == nil {
		return nil
	}
	return o.LmsInstructors
}

func (o *ListLmsInstructorsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListLmsInstructorsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
