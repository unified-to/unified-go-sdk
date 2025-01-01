// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type AccountingJournal struct {
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	Currency    *string    `json:"currency,omitempty"`
	Description *string    `json:"description,omitempty"`
	ID          *string    `json:"id,omitempty"`
	// new field name
	Lineitems []AccountingJournalLineitem `json:"lineitems,omitempty"`
	Raw       map[string]any              `json:"raw,omitempty"`
	Reference *string                     `json:"reference,omitempty"`
	TaxAmount *float64                    `json:"tax_amount,omitempty"`
	TaxrateID *string                     `json:"taxrate_id,omitempty"`
	UpdatedAt *time.Time                  `json:"updated_at,omitempty"`
}

func (a AccountingJournal) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountingJournal) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountingJournal) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AccountingJournal) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *AccountingJournal) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AccountingJournal) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AccountingJournal) GetLineitems() []AccountingJournalLineitem {
	if o == nil {
		return nil
	}
	return o.Lineitems
}

func (o *AccountingJournal) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *AccountingJournal) GetReference() *string {
	if o == nil {
		return nil
	}
	return o.Reference
}

func (o *AccountingJournal) GetTaxAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TaxAmount
}

func (o *AccountingJournal) GetTaxrateID() *string {
	if o == nil {
		return nil
	}
	return o.TaxrateID
}

func (o *AccountingJournal) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
