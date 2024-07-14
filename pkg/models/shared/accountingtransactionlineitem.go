// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AccountingTransactionLineItem struct {
	AccountID    *string  `json:"account_id,omitempty"`
	Description  *string  `json:"description,omitempty"`
	ID           *string  `json:"id,omitempty"`
	Name         *string  `json:"name,omitempty"`
	ObjectType   *string  `json:"object_type,omitempty"`
	TotalAmount  *float64 `json:"total_amount,omitempty"`
	UnitAmount   *float64 `json:"unit_amount,omitempty"`
	UnitQuantity *float64 `json:"unit_quantity,omitempty"`
}

func (o *AccountingTransactionLineItem) GetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AccountID
}

func (o *AccountingTransactionLineItem) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AccountingTransactionLineItem) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AccountingTransactionLineItem) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *AccountingTransactionLineItem) GetObjectType() *string {
	if o == nil {
		return nil
	}
	return o.ObjectType
}

func (o *AccountingTransactionLineItem) GetTotalAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *AccountingTransactionLineItem) GetUnitAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.UnitAmount
}

func (o *AccountingTransactionLineItem) GetUnitQuantity() *float64 {
	if o == nil {
		return nil
	}
	return o.UnitQuantity
}
