// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type TaskTaskStatus string

const (
	TaskTaskStatusOpened     TaskTaskStatus = "OPENED"
	TaskTaskStatusInProgress TaskTaskStatus = "IN_PROGRESS"
	TaskTaskStatusCompleted  TaskTaskStatus = "COMPLETED"
)

func (e TaskTaskStatus) ToPointer() *TaskTaskStatus {
	return &e
}
func (e *TaskTaskStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OPENED":
		fallthrough
	case "IN_PROGRESS":
		fallthrough
	case "COMPLETED":
		*e = TaskTaskStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TaskTaskStatus: %v", v)
	}
}

type TaskTask struct {
	AssignedUserIds []string        `json:"assigned_user_ids,omitempty"`
	CompletedAt     *time.Time      `json:"completed_at,omitempty"`
	CreatedAt       *time.Time      `json:"created_at,omitempty"`
	CreatorUserID   *string         `json:"creator_user_id,omitempty"`
	DueAt           *time.Time      `json:"due_at,omitempty"`
	FollowerUserIds []string        `json:"follower_user_ids,omitempty"`
	ID              *string         `json:"id,omitempty"`
	Name            string          `json:"name"`
	Notes           *string         `json:"notes,omitempty"`
	ParentID        *string         `json:"parent_id,omitempty"`
	Priority        *string         `json:"priority,omitempty"`
	ProjectID       *string         `json:"project_id,omitempty"`
	Raw             map[string]any  `json:"raw,omitempty"`
	Status          *TaskTaskStatus `json:"status,omitempty"`
	Tags            []string        `json:"tags,omitempty"`
	UpdatedAt       *time.Time      `json:"updated_at,omitempty"`
	URL             *string         `json:"url,omitempty"`
}

func (t TaskTask) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TaskTask) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TaskTask) GetAssignedUserIds() []string {
	if o == nil {
		return nil
	}
	return o.AssignedUserIds
}

func (o *TaskTask) GetCompletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CompletedAt
}

func (o *TaskTask) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *TaskTask) GetCreatorUserID() *string {
	if o == nil {
		return nil
	}
	return o.CreatorUserID
}

func (o *TaskTask) GetDueAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DueAt
}

func (o *TaskTask) GetFollowerUserIds() []string {
	if o == nil {
		return nil
	}
	return o.FollowerUserIds
}

func (o *TaskTask) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TaskTask) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *TaskTask) GetNotes() *string {
	if o == nil {
		return nil
	}
	return o.Notes
}

func (o *TaskTask) GetParentID() *string {
	if o == nil {
		return nil
	}
	return o.ParentID
}

func (o *TaskTask) GetPriority() *string {
	if o == nil {
		return nil
	}
	return o.Priority
}

func (o *TaskTask) GetProjectID() *string {
	if o == nil {
		return nil
	}
	return o.ProjectID
}

func (o *TaskTask) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *TaskTask) GetStatus() *TaskTaskStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *TaskTask) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *TaskTask) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *TaskTask) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}
