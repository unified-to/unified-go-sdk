// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type CommerceCollectionRaw struct {
}

type CommerceCollectionType string

const (
	CommerceCollectionTypeCollection  CommerceCollectionType = "COLLECTION"
	CommerceCollectionTypeSavedSearch CommerceCollectionType = "SAVED_SEARCH"
	CommerceCollectionTypeCategory    CommerceCollectionType = "CATEGORY"
)

func (e CommerceCollectionType) ToPointer() *CommerceCollectionType {
	return &e
}
func (e *CommerceCollectionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "COLLECTION":
		fallthrough
	case "SAVED_SEARCH":
		fallthrough
	case "CATEGORY":
		*e = CommerceCollectionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CommerceCollectionType: %v", v)
	}
}

// CommerceCollection - A collection of items/products/services
type CommerceCollection struct {
	CreatedAt         *time.Time              `json:"created_at,omitempty"`
	Description       *string                 `json:"description,omitempty"`
	ID                *string                 `json:"id,omitempty"`
	IsActive          *bool                   `json:"is_active,omitempty"`
	IsFeatured        *bool                   `json:"is_featured,omitempty"`
	IsVisible         *bool                   `json:"is_visible,omitempty"`
	Media             []CommerceItemMedia     `json:"media,omitempty"`
	Metadata          []CommerceMetadata      `json:"metadata,omitempty"`
	Name              string                  `json:"name"`
	ParentID          *string                 `json:"parent_id,omitempty"`
	PublicDescription *string                 `json:"public_description,omitempty"`
	PublicName        *string                 `json:"public_name,omitempty"`
	Raw               *CommerceCollectionRaw  `json:"raw,omitempty"`
	Tags              []string                `json:"tags,omitempty"`
	Type              *CommerceCollectionType `json:"type,omitempty"`
	UpdatedAt         *time.Time              `json:"updated_at,omitempty"`
}

func (c CommerceCollection) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CommerceCollection) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CommerceCollection) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *CommerceCollection) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CommerceCollection) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CommerceCollection) GetIsActive() *bool {
	if o == nil {
		return nil
	}
	return o.IsActive
}

func (o *CommerceCollection) GetIsFeatured() *bool {
	if o == nil {
		return nil
	}
	return o.IsFeatured
}

func (o *CommerceCollection) GetIsVisible() *bool {
	if o == nil {
		return nil
	}
	return o.IsVisible
}

func (o *CommerceCollection) GetMedia() []CommerceItemMedia {
	if o == nil {
		return nil
	}
	return o.Media
}

func (o *CommerceCollection) GetMetadata() []CommerceMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *CommerceCollection) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CommerceCollection) GetParentID() *string {
	if o == nil {
		return nil
	}
	return o.ParentID
}

func (o *CommerceCollection) GetPublicDescription() *string {
	if o == nil {
		return nil
	}
	return o.PublicDescription
}

func (o *CommerceCollection) GetPublicName() *string {
	if o == nil {
		return nil
	}
	return o.PublicName
}

func (o *CommerceCollection) GetRaw() *CommerceCollectionRaw {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *CommerceCollection) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CommerceCollection) GetType() *CommerceCollectionType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *CommerceCollection) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
