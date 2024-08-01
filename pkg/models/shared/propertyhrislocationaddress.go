// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type PropertyHrisLocationAddress struct {
	Address1    *string `json:"address1,omitempty"`
	Address2    *string `json:"address2,omitempty"`
	City        *string `json:"city,omitempty"`
	Country     *string `json:"country,omitempty"`
	CountryCode *string `json:"country_code,omitempty"`
	PostalCode  *string `json:"postal_code,omitempty"`
	Region      *string `json:"region,omitempty"`
	RegionCode  *string `json:"region_code,omitempty"`
}

func (o *PropertyHrisLocationAddress) GetAddress1() *string {
	if o == nil {
		return nil
	}
	return o.Address1
}

func (o *PropertyHrisLocationAddress) GetAddress2() *string {
	if o == nil {
		return nil
	}
	return o.Address2
}

func (o *PropertyHrisLocationAddress) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *PropertyHrisLocationAddress) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *PropertyHrisLocationAddress) GetCountryCode() *string {
	if o == nil {
		return nil
	}
	return o.CountryCode
}

func (o *PropertyHrisLocationAddress) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *PropertyHrisLocationAddress) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *PropertyHrisLocationAddress) GetRegionCode() *string {
	if o == nil {
		return nil
	}
	return o.RegionCode
}
