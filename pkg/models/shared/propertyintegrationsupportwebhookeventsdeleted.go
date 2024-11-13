// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type PropertyIntegrationSupportWebhookEventsDeleted string

const (
	PropertyIntegrationSupportWebhookEventsDeletedVirtual PropertyIntegrationSupportWebhookEventsDeleted = "virtual"
	PropertyIntegrationSupportWebhookEventsDeletedNative  PropertyIntegrationSupportWebhookEventsDeleted = "native"
)

func (e PropertyIntegrationSupportWebhookEventsDeleted) ToPointer() *PropertyIntegrationSupportWebhookEventsDeleted {
	return &e
}
func (e *PropertyIntegrationSupportWebhookEventsDeleted) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "virtual":
		fallthrough
	case "native":
		*e = PropertyIntegrationSupportWebhookEventsDeleted(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PropertyIntegrationSupportWebhookEventsDeleted: %v", v)
	}
}