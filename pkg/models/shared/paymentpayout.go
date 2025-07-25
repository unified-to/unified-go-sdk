// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type PaymentPayoutStatus string

const (
	PaymentPayoutStatusSucceeded PaymentPayoutStatus = "SUCCEEDED"
	PaymentPayoutStatusPending   PaymentPayoutStatus = "PENDING"
	PaymentPayoutStatusFailed    PaymentPayoutStatus = "FAILED"
	PaymentPayoutStatusCanceled  PaymentPayoutStatus = "CANCELED"
)

func (e PaymentPayoutStatus) ToPointer() *PaymentPayoutStatus {
	return &e
}

type PaymentPayout struct {
	CreatedAt   *string              `json:"created_at,omitempty"`
	Currency    *string              `json:"currency,omitempty"`
	ID          *string              `json:"id,omitempty"`
	Notes       *string              `json:"notes,omitempty"`
	Raw         map[string]any       `json:"raw,omitempty"`
	Status      *PaymentPayoutStatus `json:"status,omitempty"`
	TotalAmount float64              `json:"total_amount"`
	UpdatedAt   *string              `json:"updated_at,omitempty"`
}

func (o *PaymentPayout) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *PaymentPayout) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *PaymentPayout) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PaymentPayout) GetNotes() *string {
	if o == nil {
		return nil
	}
	return o.Notes
}

func (o *PaymentPayout) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *PaymentPayout) GetStatus() *PaymentPayoutStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *PaymentPayout) GetTotalAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.TotalAmount
}

func (o *PaymentPayout) GetUpdatedAt() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
