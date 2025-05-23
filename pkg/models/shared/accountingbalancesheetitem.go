// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type AccountingBalanceSheetItem struct {
	AccountID *string                                      `json:"account_id,omitempty"`
	Amount    *float64                                     `json:"amount,omitempty"`
	Name      *string                                      `json:"name,omitempty"`
	SubItems  []PropertyAccountingBalanceSheetItemSubItems `json:"sub_items,omitempty"`
}

func (o *AccountingBalanceSheetItem) GetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AccountID
}

func (o *AccountingBalanceSheetItem) GetAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *AccountingBalanceSheetItem) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *AccountingBalanceSheetItem) GetSubItems() []PropertyAccountingBalanceSheetItemSubItems {
	if o == nil {
		return nil
	}
	return o.SubItems
}
