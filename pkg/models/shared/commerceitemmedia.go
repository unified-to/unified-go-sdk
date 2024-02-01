// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type CommerceItemMediaType string

const (
	CommerceItemMediaTypeImage CommerceItemMediaType = "image"
	CommerceItemMediaTypeVideo CommerceItemMediaType = "video"
)

func (e CommerceItemMediaType) ToPointer() *CommerceItemMediaType {
	return &e
}

func (e *CommerceItemMediaType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "image":
		fallthrough
	case "video":
		*e = CommerceItemMediaType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CommerceItemMediaType: %v", v)
	}
}

type CommerceItemMedia struct {
	Alt      *string                `json:"alt,omitempty"`
	Height   *float64               `json:"height,omitempty"`
	Position *float64               `json:"position,omitempty"`
	Type     *CommerceItemMediaType `json:"type,omitempty"`
	URL      string                 `json:"url"`
	Width    *float64               `json:"width,omitempty"`
}

func (o *CommerceItemMedia) GetAlt() *string {
	if o == nil {
		return nil
	}
	return o.Alt
}

func (o *CommerceItemMedia) GetHeight() *float64 {
	if o == nil {
		return nil
	}
	return o.Height
}

func (o *CommerceItemMedia) GetPosition() *float64 {
	if o == nil {
		return nil
	}
	return o.Position
}

func (o *CommerceItemMedia) GetType() *CommerceItemMediaType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *CommerceItemMedia) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *CommerceItemMedia) GetWidth() *float64 {
	if o == nil {
		return nil
	}
	return o.Width
}
