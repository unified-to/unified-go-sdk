// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type AccountingBalancesheetItem struct {
	AccountID *string                                      `json:"account_id,omitempty"`
	Amount    *float64                                     `json:"amount,omitempty"`
	Name      *string                                      `json:"name,omitempty"`
	SubItems  []PropertyAccountingBalancesheetItemSubItems `json:"sub_items,omitempty"`
}

func (o *AccountingBalancesheetItem) GetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AccountID
}

func (o *AccountingBalancesheetItem) GetAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *AccountingBalancesheetItem) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *AccountingBalancesheetItem) GetSubItems() []PropertyAccountingBalancesheetItemSubItems {
	if o == nil {
		return nil
	}
	return o.SubItems
}
