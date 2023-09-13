// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

// Connection - A connection represents a specific authentication of an integration.
type Connection struct {
	// An authentication object that represents a specific authorized user's connection to an integration.
	Auth       *PropertyConnectionAuth `json:"auth,omitempty"`
	AuthAwsArn *string                 `json:"auth_aws_arn,omitempty"`
	// The Integration categories that this connection supports
	Categories      []PropertyConnectionCategories  `json:"categories"`
	CreatedAt       *types.Date                     `json:"created_at,omitempty"`
	Environment     *string                         `json:"environment,omitempty"`
	ExternalXref    *string                         `json:"external_xref,omitempty"`
	ID              *string                         `json:"id,omitempty"`
	IntegrationType string                          `json:"integration_type"`
	IsPaused        *bool                           `json:"is_paused,omitempty"`
	Permissions     []PropertyConnectionPermissions `json:"permissions"`
	UpdatedAt       *types.Date                     `json:"updated_at,omitempty"`
	WorkspaceID     *string                         `json:"workspace_id,omitempty"`
}

func (o *Connection) GetAuth() *PropertyConnectionAuth {
	if o == nil {
		return nil
	}
	return o.Auth
}

func (o *Connection) GetAuthAwsArn() *string {
	if o == nil {
		return nil
	}
	return o.AuthAwsArn
}

func (o *Connection) GetCategories() []PropertyConnectionCategories {
	if o == nil {
		return []PropertyConnectionCategories{}
	}
	return o.Categories
}

func (o *Connection) GetCreatedAt() *types.Date {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Connection) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *Connection) GetExternalXref() *string {
	if o == nil {
		return nil
	}
	return o.ExternalXref
}

func (o *Connection) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Connection) GetIntegrationType() string {
	if o == nil {
		return ""
	}
	return o.IntegrationType
}

func (o *Connection) GetIsPaused() *bool {
	if o == nil {
		return nil
	}
	return o.IsPaused
}

func (o *Connection) GetPermissions() []PropertyConnectionPermissions {
	if o == nil {
		return []PropertyConnectionPermissions{}
	}
	return o.Permissions
}

func (o *Connection) GetUpdatedAt() *types.Date {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *Connection) GetWorkspaceID() *string {
	if o == nil {
		return nil
	}
	return o.WorkspaceID
}
