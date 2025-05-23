// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type CommerceItemOption struct {
	ID       *string                            `json:"id,omitempty"`
	Name     string                             `json:"name"`
	Position *float64                           `json:"position,omitempty"`
	Values   []PropertyCommerceItemOptionValues `json:"values"`
}

func (o *CommerceItemOption) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CommerceItemOption) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CommerceItemOption) GetPosition() *float64 {
	if o == nil {
		return nil
	}
	return o.Position
}

func (o *CommerceItemOption) GetValues() []PropertyCommerceItemOptionValues {
	if o == nil {
		return []PropertyCommerceItemOptionValues{}
	}
	return o.Values
}
