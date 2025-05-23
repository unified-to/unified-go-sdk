// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type PropertyAtsJobQuestionOptions string

const (
	PropertyAtsJobQuestionOptionsUnknown PropertyAtsJobQuestionOptions = ""
)

func (e PropertyAtsJobQuestionOptions) ToPointer() *PropertyAtsJobQuestionOptions {
	return &e
}
func (e *PropertyAtsJobQuestionOptions) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "":
		*e = PropertyAtsJobQuestionOptions(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PropertyAtsJobQuestionOptions: %v", v)
	}
}
