// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type AtsScorecardRecommendation string

const (
	AtsScorecardRecommendationDefinitelyNo AtsScorecardRecommendation = "DEFINITELY_NO"
	AtsScorecardRecommendationNo           AtsScorecardRecommendation = "NO"
	AtsScorecardRecommendationYes          AtsScorecardRecommendation = "YES"
	AtsScorecardRecommendationStrongYes    AtsScorecardRecommendation = "STRONG_YES"
)

func (e AtsScorecardRecommendation) ToPointer() *AtsScorecardRecommendation {
	return &e
}

func (e *AtsScorecardRecommendation) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DEFINITELY_NO":
		fallthrough
	case "NO":
		fallthrough
	case "YES":
		fallthrough
	case "STRONG_YES":
		*e = AtsScorecardRecommendation(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AtsScorecardRecommendation: %v", v)
	}
}

type AtsScorecard struct {
	ApplicationID  *string                     `json:"application_id,omitempty"`
	CandidateID    *string                     `json:"candidate_id,omitempty"`
	CreatedAt      *time.Time                  `json:"created_at,omitempty"`
	ID             *string                     `json:"id,omitempty"`
	InterviewID    *string                     `json:"interview_id,omitempty"`
	InterviewerID  *string                     `json:"interviewer_id,omitempty"`
	JobID          *string                     `json:"job_id,omitempty"`
	Raw            PropertyAtsScorecardRaw     `json:"raw"`
	Recommendation *AtsScorecardRecommendation `json:"recommendation,omitempty"`
	UpdatedAt      *time.Time                  `json:"updated_at,omitempty"`
}

func (a AtsScorecard) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AtsScorecard) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AtsScorecard) GetApplicationID() *string {
	if o == nil {
		return nil
	}
	return o.ApplicationID
}

func (o *AtsScorecard) GetCandidateID() *string {
	if o == nil {
		return nil
	}
	return o.CandidateID
}

func (o *AtsScorecard) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AtsScorecard) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AtsScorecard) GetInterviewID() *string {
	if o == nil {
		return nil
	}
	return o.InterviewID
}

func (o *AtsScorecard) GetInterviewerID() *string {
	if o == nil {
		return nil
	}
	return o.InterviewerID
}

func (o *AtsScorecard) GetJobID() *string {
	if o == nil {
		return nil
	}
	return o.JobID
}

func (o *AtsScorecard) GetRaw() PropertyAtsScorecardRaw {
	if o == nil {
		return PropertyAtsScorecardRaw{}
	}
	return o.Raw
}

func (o *AtsScorecard) GetRecommendation() *AtsScorecardRecommendation {
	if o == nil {
		return nil
	}
	return o.Recommendation
}

func (o *AtsScorecard) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
