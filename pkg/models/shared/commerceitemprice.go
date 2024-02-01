// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CommerceItemPrice struct {
	CompareAtPrice *float64 `json:"compare_at_price,omitempty"`
	Currency       *string  `json:"currency,omitempty"`
	Price          float64  `json:"price"`
}

func (o *CommerceItemPrice) GetCompareAtPrice() *float64 {
	if o == nil {
		return nil
	}
	return o.CompareAtPrice
}

func (o *CommerceItemPrice) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *CommerceItemPrice) GetPrice() float64 {
	if o == nil {
		return 0.0
	}
	return o.Price
}
