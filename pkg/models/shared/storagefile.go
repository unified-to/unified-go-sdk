// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type StorageFileType string

const (
	StorageFileTypeFile   StorageFileType = "FILE"
	StorageFileTypeFolder StorageFileType = "FOLDER"
)

func (e StorageFileType) ToPointer() *StorageFileType {
	return &e
}
func (e *StorageFileType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "FILE":
		fallthrough
	case "FOLDER":
		*e = StorageFileType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StorageFileType: %v", v)
	}
}

type StorageFile struct {
	CreatedAt   *time.Time          `json:"created_at,omitempty"`
	Description *string             `json:"description,omitempty"`
	DownloadURL *string             `json:"download_url,omitempty"`
	Hash        *string             `json:"hash,omitempty"`
	ID          *string             `json:"id,omitempty"`
	MimeType    *string             `json:"mime_type,omitempty"`
	Name        *string             `json:"name,omitempty"`
	ParentID    *string             `json:"parent_id,omitempty"`
	Permissions []StoragePermission `json:"permissions,omitempty"`
	Raw         map[string]any      `json:"raw,omitempty"`
	Size        *float64            `json:"size,omitempty"`
	Type        *StorageFileType    `json:"type,omitempty"`
	UpdatedAt   *time.Time          `json:"updated_at,omitempty"`
	UserID      *string             `json:"user_id,omitempty"`
}

func (s StorageFile) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *StorageFile) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *StorageFile) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *StorageFile) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *StorageFile) GetDownloadURL() *string {
	if o == nil {
		return nil
	}
	return o.DownloadURL
}

func (o *StorageFile) GetHash() *string {
	if o == nil {
		return nil
	}
	return o.Hash
}

func (o *StorageFile) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *StorageFile) GetMimeType() *string {
	if o == nil {
		return nil
	}
	return o.MimeType
}

func (o *StorageFile) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *StorageFile) GetParentID() *string {
	if o == nil {
		return nil
	}
	return o.ParentID
}

func (o *StorageFile) GetPermissions() []StoragePermission {
	if o == nil {
		return nil
	}
	return o.Permissions
}

func (o *StorageFile) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *StorageFile) GetSize() *float64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *StorageFile) GetType() *StorageFileType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *StorageFile) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *StorageFile) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}
