// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AccountingTransaction struct {
	CreatedAt   *string                         `json:"created_at,omitempty"`
	Currency    *string                         `json:"currency,omitempty"`
	Description *string                         `json:"description,omitempty"`
	ID          string                          `json:"id"`
	LineItems   []AccountingTransactionLineitem `json:"line_items,omitempty"`
	Raw         map[string]interface{}          `json:"raw,omitempty"`
	Reference   *string                         `json:"reference,omitempty"`
	TaxAmount   *float64                        `json:"tax_amount,omitempty"`
	TaxrateID   *string                         `json:"taxrate_id,omitempty"`
	UpdatedAt   *string                         `json:"updated_at,omitempty"`
}

func (o *AccountingTransaction) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AccountingTransaction) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *AccountingTransaction) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AccountingTransaction) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AccountingTransaction) GetLineItems() []AccountingTransactionLineitem {
	if o == nil {
		return nil
	}
	return o.LineItems
}

func (o *AccountingTransaction) GetRaw() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *AccountingTransaction) GetReference() *string {
	if o == nil {
		return nil
	}
	return o.Reference
}

func (o *AccountingTransaction) GetTaxAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TaxAmount
}

func (o *AccountingTransaction) GetTaxrateID() *string {
	if o == nil {
		return nil
	}
	return o.TaxrateID
}

func (o *AccountingTransaction) GetUpdatedAt() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
