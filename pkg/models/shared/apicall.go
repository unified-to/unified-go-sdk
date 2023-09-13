// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

type APICallType string

const (
	APICallTypeLogin   APICallType = "login"
	APICallTypeWebhook APICallType = "webhook"
	APICallTypeInbound APICallType = "inbound"
)

func (e APICallType) ToPointer() *APICallType {
	return &e
}

func (e *APICallType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "login":
		fallthrough
	case "webhook":
		fallthrough
	case "inbound":
		*e = APICallType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for APICallType: %v", v)
	}
}

type APICall struct {
	ConnectionID    *string     `json:"connection_id,omitempty"`
	CreatedAt       *types.Date `json:"created_at,omitempty"`
	Environment     *string     `json:"environment,omitempty"`
	Error           *string     `json:"error,omitempty"`
	ExternalXref    *string     `json:"external_xref,omitempty"`
	ID              *string     `json:"id,omitempty"`
	IntegrationType string      `json:"integration_type"`
	IPAddress       *string     `json:"ip_address,omitempty"`
	Method          string      `json:"method"`
	Name            string      `json:"name"`
	Path            string      `json:"path"`
	Size            *float64    `json:"size,omitempty"`
	Status          string      `json:"status"`
	Type            APICallType `json:"type"`
	WorkspaceID     string      `json:"workspace_id"`
}

func (o *APICall) GetConnectionID() *string {
	if o == nil {
		return nil
	}
	return o.ConnectionID
}

func (o *APICall) GetCreatedAt() *types.Date {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *APICall) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *APICall) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *APICall) GetExternalXref() *string {
	if o == nil {
		return nil
	}
	return o.ExternalXref
}

func (o *APICall) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *APICall) GetIntegrationType() string {
	if o == nil {
		return ""
	}
	return o.IntegrationType
}

func (o *APICall) GetIPAddress() *string {
	if o == nil {
		return nil
	}
	return o.IPAddress
}

func (o *APICall) GetMethod() string {
	if o == nil {
		return ""
	}
	return o.Method
}

func (o *APICall) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *APICall) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

func (o *APICall) GetSize() *float64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *APICall) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *APICall) GetType() APICallType {
	if o == nil {
		return APICallType("")
	}
	return o.Type
}

func (o *APICall) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
