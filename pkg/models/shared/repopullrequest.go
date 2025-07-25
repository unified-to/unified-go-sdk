// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type RepoPullrequestStatus string

const (
	RepoPullrequestStatusPending  RepoPullrequestStatus = "PENDING"
	RepoPullrequestStatusApproved RepoPullrequestStatus = "APPROVED"
	RepoPullrequestStatusRejected RepoPullrequestStatus = "REJECTED"
)

func (e RepoPullrequestStatus) ToPointer() *RepoPullrequestStatus {
	return &e
}

type RepoPullrequest struct {
	ClosedAt  *time.Time             `json:"closed_at,omitempty"`
	CommitIds []string               `json:"commit_ids,omitempty"`
	CreatedAt *time.Time             `json:"created_at,omitempty"`
	ID        *string                `json:"id,omitempty"`
	Labels    []string               `json:"labels,omitempty"`
	Raw       map[string]any         `json:"raw,omitempty"`
	RepoID    *string                `json:"repo_id,omitempty"`
	Status    *RepoPullrequestStatus `json:"status,omitempty"`
	UpdatedAt *time.Time             `json:"updated_at,omitempty"`
	UserIds   []string               `json:"user_ids,omitempty"`
}

func (r RepoPullrequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RepoPullrequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RepoPullrequest) GetClosedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ClosedAt
}

func (o *RepoPullrequest) GetCommitIds() []string {
	if o == nil {
		return nil
	}
	return o.CommitIds
}

func (o *RepoPullrequest) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *RepoPullrequest) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RepoPullrequest) GetLabels() []string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *RepoPullrequest) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *RepoPullrequest) GetRepoID() *string {
	if o == nil {
		return nil
	}
	return o.RepoID
}

func (o *RepoPullrequest) GetStatus() *RepoPullrequestStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *RepoPullrequest) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *RepoPullrequest) GetUserIds() []string {
	if o == nil {
		return nil
	}
	return o.UserIds
}
