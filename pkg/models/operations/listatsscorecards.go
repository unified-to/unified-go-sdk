// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"net/http"
	"time"
)

type ListAtsScorecardsRequest struct {
	ApplicationID *string `queryParam:"style=form,explode=true,name=application_id"`
	CandidateID   *string `queryParam:"style=form,explode=true,name=candidate_id"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Comma-delimited fields to return
	Fields      []string `queryParam:"style=form,explode=true,name=fields"`
	InterviewID *string  `queryParam:"style=form,explode=true,name=interview_id"`
	JobID       *string  `queryParam:"style=form,explode=true,name=job_id"`
	Limit       *float64 `queryParam:"style=form,explode=true,name=limit"`
	Offset      *float64 `queryParam:"style=form,explode=true,name=offset"`
	// Query string to search. eg. email address or name
	Query *string `queryParam:"style=form,explode=true,name=query"`
	// Return only results whose updated date is equal or greater to this value
	UpdatedGte *time.Time `queryParam:"style=form,explode=true,name=updated_gte"`
}

func (l ListAtsScorecardsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAtsScorecardsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAtsScorecardsRequest) GetApplicationID() *string {
	if o == nil {
		return nil
	}
	return o.ApplicationID
}

func (o *ListAtsScorecardsRequest) GetCandidateID() *string {
	if o == nil {
		return nil
	}
	return o.CandidateID
}

func (o *ListAtsScorecardsRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListAtsScorecardsRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *ListAtsScorecardsRequest) GetInterviewID() *string {
	if o == nil {
		return nil
	}
	return o.InterviewID
}

func (o *ListAtsScorecardsRequest) GetJobID() *string {
	if o == nil {
		return nil
	}
	return o.JobID
}

func (o *ListAtsScorecardsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListAtsScorecardsRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListAtsScorecardsRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *ListAtsScorecardsRequest) GetUpdatedGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedGte
}

type ListAtsScorecardsResponse struct {
	// Successful
	AtsScorecards []shared.AtsScorecard
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListAtsScorecardsResponse) GetAtsScorecards() []shared.AtsScorecard {
	if o == nil {
		return nil
	}
	return o.AtsScorecards
}

func (o *ListAtsScorecardsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListAtsScorecardsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListAtsScorecardsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
