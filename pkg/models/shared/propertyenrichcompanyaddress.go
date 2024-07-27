// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// PropertyEnrichCompanyAddress - The address of the company
type PropertyEnrichCompanyAddress struct {
	Address1    *string `json:"address1,omitempty"`
	Address2    *string `json:"address2,omitempty"`
	City        *string `json:"city,omitempty"`
	Country     *string `json:"country,omitempty"`
	CountryCode *string `json:"country_code,omitempty"`
	PostalCode  *string `json:"postal_code,omitempty"`
	Region      *string `json:"region,omitempty"`
	RegionCode  *string `json:"region_code,omitempty"`
}

func (o *PropertyEnrichCompanyAddress) GetAddress1() *string {
	if o == nil {
		return nil
	}
	return o.Address1
}

func (o *PropertyEnrichCompanyAddress) GetAddress2() *string {
	if o == nil {
		return nil
	}
	return o.Address2
}

func (o *PropertyEnrichCompanyAddress) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *PropertyEnrichCompanyAddress) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *PropertyEnrichCompanyAddress) GetCountryCode() *string {
	if o == nil {
		return nil
	}
	return o.CountryCode
}

func (o *PropertyEnrichCompanyAddress) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *PropertyEnrichCompanyAddress) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *PropertyEnrichCompanyAddress) GetRegionCode() *string {
	if o == nil {
		return nil
	}
	return o.RegionCode
}
