// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type PropertyIntegrationTokenInstructions string

const (
	PropertyIntegrationTokenInstructionsUnknown PropertyIntegrationTokenInstructions = ""
)

func (e PropertyIntegrationTokenInstructions) ToPointer() *PropertyIntegrationTokenInstructions {
	return &e
}
func (e *PropertyIntegrationTokenInstructions) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "":
		*e = PropertyIntegrationTokenInstructions(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PropertyIntegrationTokenInstructions: %v", v)
	}
}
