// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type LmsMediaType string

const (
	LmsMediaTypeImage    LmsMediaType = "IMAGE"
	LmsMediaTypeHeadshot LmsMediaType = "HEADSHOT"
	LmsMediaTypeVideo    LmsMediaType = "VIDEO"
	LmsMediaTypeWeb      LmsMediaType = "WEB"
	LmsMediaTypeDocument LmsMediaType = "DOCUMENT"
	LmsMediaTypeOther    LmsMediaType = "OTHER"
)

func (e LmsMediaType) ToPointer() *LmsMediaType {
	return &e
}
func (e *LmsMediaType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "IMAGE":
		fallthrough
	case "HEADSHOT":
		fallthrough
	case "VIDEO":
		fallthrough
	case "WEB":
		fallthrough
	case "DOCUMENT":
		fallthrough
	case "OTHER":
		*e = LmsMediaType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LmsMediaType: %v", v)
	}
}

type LmsMedia struct {
	Description  *string       `json:"description,omitempty"`
	Name         *string       `json:"name,omitempty"`
	ThumbnailURL *string       `json:"thumbnail_url,omitempty"`
	Type         *LmsMediaType `json:"type,omitempty"`
	URL          string        `json:"url"`
}

func (o *LmsMedia) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *LmsMedia) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LmsMedia) GetThumbnailURL() *string {
	if o == nil {
		return nil
	}
	return o.ThumbnailURL
}

func (o *LmsMedia) GetType() *LmsMediaType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *LmsMedia) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}