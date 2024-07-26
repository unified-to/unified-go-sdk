// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

type MessagingMember struct {
	Email  *string `json:"email,omitempty"`
	Name   *string `json:"name,omitempty"`
	UserID *string `json:"user_id,omitempty"`
}

func (o *MessagingMember) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *MessagingMember) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *MessagingMember) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}
