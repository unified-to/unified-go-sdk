// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PropertyMessagingMessageAuthorMember struct {
	Email  *string `json:"email,omitempty"`
	Name   *string `json:"name,omitempty"`
	UserID *string `json:"user_id,omitempty"`
}

func (o *PropertyMessagingMessageAuthorMember) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *PropertyMessagingMessageAuthorMember) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PropertyMessagingMessageAuthorMember) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}