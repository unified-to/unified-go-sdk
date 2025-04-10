// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type CalendarLink struct {
	CreatedAt     *string        `json:"created_at,omitempty"`
	Description   *string        `json:"description,omitempty"`
	Duration      *float64       `json:"duration,omitempty"`
	ID            *string        `json:"id,omitempty"`
	IsActive      *bool          `json:"is_active,omitempty"`
	Name          *string        `json:"name,omitempty"`
	PriceAmount   *float64       `json:"price_amount,omitempty"`
	PriceCurrency *string        `json:"price_currency,omitempty"`
	Raw           map[string]any `json:"raw,omitempty"`
	UpdatedAt     *string        `json:"updated_at,omitempty"`
	URL           string         `json:"url"`
}

func (o *CalendarLink) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *CalendarLink) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CalendarLink) GetDuration() *float64 {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *CalendarLink) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CalendarLink) GetIsActive() *bool {
	if o == nil {
		return nil
	}
	return o.IsActive
}

func (o *CalendarLink) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CalendarLink) GetPriceAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.PriceAmount
}

func (o *CalendarLink) GetPriceCurrency() *string {
	if o == nil {
		return nil
	}
	return o.PriceCurrency
}

func (o *CalendarLink) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *CalendarLink) GetUpdatedAt() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *CalendarLink) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}
