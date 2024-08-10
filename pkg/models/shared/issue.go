// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type IssueStatus string

const (
	IssueStatusCompleted  IssueStatus = "COMPLETED"
	IssueStatusNew        IssueStatus = "NEW"
	IssueStatusRoadmap    IssueStatus = "ROADMAP"
	IssueStatusInProgress IssueStatus = "IN_PROGRESS"
	IssueStatusOnHold     IssueStatus = "ON_HOLD"
	IssueStatusValidating IssueStatus = "VALIDATING"
	IssueStatusRejected   IssueStatus = "REJECTED"
	IssueStatusUpNext     IssueStatus = "UP_NEXT"
)

func (e IssueStatus) ToPointer() *IssueStatus {
	return &e
}
func (e *IssueStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "COMPLETED":
		fallthrough
	case "NEW":
		fallthrough
	case "ROADMAP":
		fallthrough
	case "IN_PROGRESS":
		fallthrough
	case "ON_HOLD":
		fallthrough
	case "VALIDATING":
		fallthrough
	case "REJECTED":
		fallthrough
	case "UP_NEXT":
		*e = IssueStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for IssueStatus: %v", v)
	}
}

type Issue struct {
	CreatedAt      *string     `json:"created_at,omitempty"`
	ID             *string     `json:"id,omitempty"`
	Importance     *float64    `json:"importance,omitempty"`
	ResolutionTime *float64    `json:"resolution_time,omitempty"`
	Size           *float64    `json:"size,omitempty"`
	Status         IssueStatus `json:"status"`
	TicketRef      string      `json:"ticket_ref"`
	Title          string      `json:"title"`
	Type           []string    `json:"type,omitempty"`
	UpdatedAt      *string     `json:"updated_at,omitempty"`
	URL            *string     `json:"url,omitempty"`
	WorkspaceID    string      `json:"workspace_id"`
}

func (o *Issue) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Issue) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Issue) GetImportance() *float64 {
	if o == nil {
		return nil
	}
	return o.Importance
}

func (o *Issue) GetResolutionTime() *float64 {
	if o == nil {
		return nil
	}
	return o.ResolutionTime
}

func (o *Issue) GetSize() *float64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *Issue) GetStatus() IssueStatus {
	if o == nil {
		return IssueStatus("")
	}
	return o.Status
}

func (o *Issue) GetTicketRef() string {
	if o == nil {
		return ""
	}
	return o.TicketRef
}

func (o *Issue) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *Issue) GetType() []string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Issue) GetUpdatedAt() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *Issue) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *Issue) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
