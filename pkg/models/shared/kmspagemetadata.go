// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type KmsPageMetadata struct {
	Name  string         `json:"name"`
	Type  *string        `json:"type,omitempty"`
	Value map[string]any `json:"value,omitempty"`
}

func (o *KmsPageMetadata) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *KmsPageMetadata) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *KmsPageMetadata) GetValue() map[string]any {
	if o == nil {
		return nil
	}
	return o.Value
}