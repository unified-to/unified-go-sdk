// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type PropertyIntegrationCategories string

const (
	PropertyIntegrationCategoriesPassthrough PropertyIntegrationCategories = "passthrough"
	PropertyIntegrationCategoriesHris        PropertyIntegrationCategories = "hris"
	PropertyIntegrationCategoriesAts         PropertyIntegrationCategories = "ats"
	PropertyIntegrationCategoriesAuth        PropertyIntegrationCategories = "auth"
	PropertyIntegrationCategoriesCrm         PropertyIntegrationCategories = "crm"
	PropertyIntegrationCategoriesEnrich      PropertyIntegrationCategories = "enrich"
	PropertyIntegrationCategoriesMartech     PropertyIntegrationCategories = "martech"
	PropertyIntegrationCategoriesTicketing   PropertyIntegrationCategories = "ticketing"
	PropertyIntegrationCategoriesUc          PropertyIntegrationCategories = "uc"
)

func (e PropertyIntegrationCategories) ToPointer() *PropertyIntegrationCategories {
	return &e
}

func (e *PropertyIntegrationCategories) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "passthrough":
		fallthrough
	case "hris":
		fallthrough
	case "ats":
		fallthrough
	case "auth":
		fallthrough
	case "crm":
		fallthrough
	case "enrich":
		fallthrough
	case "martech":
		fallthrough
	case "ticketing":
		fallthrough
	case "uc":
		*e = PropertyIntegrationCategories(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PropertyIntegrationCategories: %v", v)
	}
}