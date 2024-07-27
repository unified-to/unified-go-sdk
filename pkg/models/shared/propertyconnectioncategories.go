// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type PropertyConnectionCategories string

const (
	PropertyConnectionCategoriesPassthrough PropertyConnectionCategories = "passthrough"
	PropertyConnectionCategoriesHris        PropertyConnectionCategories = "hris"
	PropertyConnectionCategoriesAts         PropertyConnectionCategories = "ats"
	PropertyConnectionCategoriesAuth        PropertyConnectionCategories = "auth"
	PropertyConnectionCategoriesCrm         PropertyConnectionCategories = "crm"
	PropertyConnectionCategoriesEnrich      PropertyConnectionCategories = "enrich"
	PropertyConnectionCategoriesMartech     PropertyConnectionCategories = "martech"
	PropertyConnectionCategoriesTicketing   PropertyConnectionCategories = "ticketing"
	PropertyConnectionCategoriesUc          PropertyConnectionCategories = "uc"
	PropertyConnectionCategoriesAccounting  PropertyConnectionCategories = "accounting"
	PropertyConnectionCategoriesStorage     PropertyConnectionCategories = "storage"
	PropertyConnectionCategoriesCommerce    PropertyConnectionCategories = "commerce"
	PropertyConnectionCategoriesPayment     PropertyConnectionCategories = "payment"
	PropertyConnectionCategoriesGenai       PropertyConnectionCategories = "genai"
	PropertyConnectionCategoriesMessaging   PropertyConnectionCategories = "messaging"
	PropertyConnectionCategoriesKms         PropertyConnectionCategories = "kms"
	PropertyConnectionCategoriesTask        PropertyConnectionCategories = "task"
)

func (e PropertyConnectionCategories) ToPointer() *PropertyConnectionCategories {
	return &e
}
func (e *PropertyConnectionCategories) UnmarshalJSON(data []byte) error {
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
		fallthrough
	case "accounting":
		fallthrough
	case "storage":
		fallthrough
	case "commerce":
		fallthrough
	case "payment":
		fallthrough
	case "genai":
		fallthrough
	case "messaging":
		fallthrough
	case "kms":
		fallthrough
	case "task":
		*e = PropertyConnectionCategories(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PropertyConnectionCategories: %v", v)
	}
}
