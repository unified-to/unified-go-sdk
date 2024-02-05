// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PropertyAtsCompanyAddress struct {
	Address1    *string `json:"address1,omitempty"`
	Address2    *string `json:"address2,omitempty"`
	City        *string `json:"city,omitempty"`
	Country     *string `json:"country,omitempty"`
	CountryCode *string `json:"country_code,omitempty"`
	PostalCode  *string `json:"postal_code,omitempty"`
	Region      *string `json:"region,omitempty"`
	RegionCode  *string `json:"region_code,omitempty"`
}

func (o *PropertyAtsCompanyAddress) GetAddress1() *string {
	if o == nil {
		return nil
	}
	return o.Address1
}

func (o *PropertyAtsCompanyAddress) GetAddress2() *string {
	if o == nil {
		return nil
	}
	return o.Address2
}

func (o *PropertyAtsCompanyAddress) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *PropertyAtsCompanyAddress) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *PropertyAtsCompanyAddress) GetCountryCode() *string {
	if o == nil {
		return nil
	}
	return o.CountryCode
}

func (o *PropertyAtsCompanyAddress) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *PropertyAtsCompanyAddress) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *PropertyAtsCompanyAddress) GetRegionCode() *string {
	if o == nil {
		return nil
	}
	return o.RegionCode
}
