// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type PropertyAccountingOrderBillingAddress struct {
	Address1    *string `json:"address1,omitempty"`
	Address2    *string `json:"address2,omitempty"`
	City        *string `json:"city,omitempty"`
	Country     *string `json:"country,omitempty"`
	CountryCode *string `json:"country_code,omitempty"`
	PostalCode  *string `json:"postal_code,omitempty"`
	Region      *string `json:"region,omitempty"`
	RegionCode  *string `json:"region_code,omitempty"`
}

func (o *PropertyAccountingOrderBillingAddress) GetAddress1() *string {
	if o == nil {
		return nil
	}
	return o.Address1
}

func (o *PropertyAccountingOrderBillingAddress) GetAddress2() *string {
	if o == nil {
		return nil
	}
	return o.Address2
}

func (o *PropertyAccountingOrderBillingAddress) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *PropertyAccountingOrderBillingAddress) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *PropertyAccountingOrderBillingAddress) GetCountryCode() *string {
	if o == nil {
		return nil
	}
	return o.CountryCode
}

func (o *PropertyAccountingOrderBillingAddress) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *PropertyAccountingOrderBillingAddress) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *PropertyAccountingOrderBillingAddress) GetRegionCode() *string {
	if o == nil {
		return nil
	}
	return o.RegionCode
}