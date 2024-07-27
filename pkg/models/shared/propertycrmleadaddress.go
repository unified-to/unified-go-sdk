// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type PropertyCrmLeadAddress struct {
	Address1    *string `json:"address1,omitempty"`
	Address2    *string `json:"address2,omitempty"`
	City        *string `json:"city,omitempty"`
	Country     *string `json:"country,omitempty"`
	CountryCode *string `json:"country_code,omitempty"`
	PostalCode  *string `json:"postal_code,omitempty"`
	Region      *string `json:"region,omitempty"`
	RegionCode  *string `json:"region_code,omitempty"`
}

func (o *PropertyCrmLeadAddress) GetAddress1() *string {
	if o == nil {
		return nil
	}
	return o.Address1
}

func (o *PropertyCrmLeadAddress) GetAddress2() *string {
	if o == nil {
		return nil
	}
	return o.Address2
}

func (o *PropertyCrmLeadAddress) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *PropertyCrmLeadAddress) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *PropertyCrmLeadAddress) GetCountryCode() *string {
	if o == nil {
		return nil
	}
	return o.CountryCode
}

func (o *PropertyCrmLeadAddress) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *PropertyCrmLeadAddress) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *PropertyCrmLeadAddress) GetRegionCode() *string {
	if o == nil {
		return nil
	}
	return o.RegionCode
}
