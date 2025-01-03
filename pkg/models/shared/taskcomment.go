// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type TaskComment struct {
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	ID        *string        `json:"id,omitempty"`
	Raw       map[string]any `json:"raw,omitempty"`
	TaskID    string         `json:"task_id"`
	Text      string         `json:"text"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	UserID    *string        `json:"user_id,omitempty"`
	UserName  *string        `json:"user_name,omitempty"`
}

func (t TaskComment) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TaskComment) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TaskComment) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *TaskComment) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TaskComment) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *TaskComment) GetTaskID() string {
	if o == nil {
		return ""
	}
	return o.TaskID
}

func (o *TaskComment) GetText() string {
	if o == nil {
		return ""
	}
	return o.Text
}

func (o *TaskComment) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *TaskComment) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

func (o *TaskComment) GetUserName() *string {
	if o == nil {
		return nil
	}
	return o.UserName
}
