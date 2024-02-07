// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AtsScorecardQuestion struct {
	Answer      *string `json:"answer,omitempty"`
	Description *string `json:"description,omitempty"`
	Text        string  `json:"text"`
}

func (o *AtsScorecardQuestion) GetAnswer() *string {
	if o == nil {
		return nil
	}
	return o.Answer
}

func (o *AtsScorecardQuestion) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AtsScorecardQuestion) GetText() string {
	if o == nil {
		return ""
	}
	return o.Text
}
