// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type RepoCommit struct {
	BranchID  *string        `json:"branch_id,omitempty"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	ID        *string        `json:"id,omitempty"`
	Message   *string        `json:"message,omitempty"`
	Raw       map[string]any `json:"raw,omitempty"`
	RepoID    string         `json:"repo_id"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	UserID    string         `json:"user_id"`
}

func (r RepoCommit) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RepoCommit) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RepoCommit) GetBranchID() *string {
	if o == nil {
		return nil
	}
	return o.BranchID
}

func (o *RepoCommit) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *RepoCommit) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RepoCommit) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *RepoCommit) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *RepoCommit) GetRepoID() string {
	if o == nil {
		return ""
	}
	return o.RepoID
}

func (o *RepoCommit) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *RepoCommit) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}