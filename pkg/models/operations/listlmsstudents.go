// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"net/http"
	"time"
)

type ListLmsStudentsRequest struct {
	ClassID *string `queryParam:"style=form,explode=true,name=class_id"`
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

func (l ListLmsStudentsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListLmsStudentsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListLmsStudentsRequest) GetClassID() *string {
	if o == nil {
		return nil
	}
	return o.ClassID
}

func (o *ListLmsStudentsRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListLmsStudentsRequest) GetCourseID() *string {
	if o == nil {
		return nil
	}
	return o.CourseID
}

func (o *ListLmsStudentsRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *ListLmsStudentsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListLmsStudentsRequest) GetLocationID() *string {
	if o == nil {
		return nil
	}
	return o.LocationID
}

func (o *ListLmsStudentsRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListLmsStudentsRequest) GetOrder() *string {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListLmsStudentsRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *ListLmsStudentsRequest) GetSessionID() *string {
	if o == nil {
		return nil
	}
	return o.SessionID
}

func (o *ListLmsStudentsRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListLmsStudentsRequest) GetUpdatedGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedGte
}

type ListLmsStudentsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful
	LmsStudents []shared.LmsStudent
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListLmsStudentsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListLmsStudentsResponse) GetLmsStudents() []shared.LmsStudent {
	if o == nil {
		return nil
	}
	return o.LmsStudents
}

func (o *ListLmsStudentsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListLmsStudentsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
