// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type CommerceItemMetadataExtraData struct {
}

type CommerceItemMetadataValue struct {
}

type CommerceItemMetadata struct {
	ExtraData *CommerceItemMetadataExtraData `json:"extra_data,omitempty"`
	ID        *string                        `json:"id,omitempty"`
	Key       string                         `json:"key"`
	Namespace *string                        `json:"namespace,omitempty"`
	Type      *string                        `json:"type,omitempty"`
	Value     *CommerceItemMetadataValue     `json:"value,omitempty"`
}

func (o *CommerceItemMetadata) GetExtraData() *CommerceItemMetadataExtraData {
	if o == nil {
		return nil
	}
	return o.ExtraData
}

func (o *CommerceItemMetadata) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CommerceItemMetadata) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *CommerceItemMetadata) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *CommerceItemMetadata) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *CommerceItemMetadata) GetValue() *CommerceItemMetadataValue {
	if o == nil {
		return nil
	}
	return o.Value
}
