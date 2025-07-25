// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type AccountingTelephoneType string

const (
	AccountingTelephoneTypeWork   AccountingTelephoneType = "WORK"
	AccountingTelephoneTypeHome   AccountingTelephoneType = "HOME"
	AccountingTelephoneTypeOther  AccountingTelephoneType = "OTHER"
	AccountingTelephoneTypeFax    AccountingTelephoneType = "FAX"
	AccountingTelephoneTypeMobile AccountingTelephoneType = "MOBILE"
)

func (e AccountingTelephoneType) ToPointer() *AccountingTelephoneType {
	return &e
}

type AccountingTelephone struct {
	Telephone *string                  `json:"telephone,omitempty"`
	Type      *AccountingTelephoneType `json:"type,omitempty"`
}

func (o *AccountingTelephone) GetTelephone() *string {
	if o == nil {
		return nil
	}
	return o.Telephone
}

func (o *AccountingTelephone) GetType() *AccountingTelephoneType {
	if o == nil {
		return nil
	}
	return o.Type
}
