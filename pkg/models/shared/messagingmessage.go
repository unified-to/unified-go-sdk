// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type MessagingMessage struct {
	AuthorMember       PropertyMessagingMessageAuthorMember `json:"author_member"`
	ChannelID          *string                              `json:"channel_id,omitempty"`
	CreatedAt          *string                              `json:"created_at,omitempty"`
	DestinationMembers []MessagingMember                    `json:"destination_members,omitempty"`
	HiddenMembers      []MessagingMember                    `json:"hidden_members,omitempty"`
	ID                 *string                              `json:"id,omitempty"`
	MentionedMembers   []MessagingMember                    `json:"mentioned_members,omitempty"`
	Message            string                               `json:"message"`
	ParentMessageID    *string                              `json:"parent_message_id,omitempty"`
	Raw                map[string]any                       `json:"raw,omitempty"`
	Subject            *string                              `json:"subject,omitempty"`
	UpdatedAt          *string                              `json:"updated_at,omitempty"`
	WebURL             *string                              `json:"web_url,omitempty"`
}

func (o *MessagingMessage) GetAuthorMember() PropertyMessagingMessageAuthorMember {
	if o == nil {
		return PropertyMessagingMessageAuthorMember{}
	}
	return o.AuthorMember
}

func (o *MessagingMessage) GetChannelID() *string {
	if o == nil {
		return nil
	}
	return o.ChannelID
}

func (o *MessagingMessage) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *MessagingMessage) GetDestinationMembers() []MessagingMember {
	if o == nil {
		return nil
	}
	return o.DestinationMembers
}

func (o *MessagingMessage) GetHiddenMembers() []MessagingMember {
	if o == nil {
		return nil
	}
	return o.HiddenMembers
}

func (o *MessagingMessage) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *MessagingMessage) GetMentionedMembers() []MessagingMember {
	if o == nil {
		return nil
	}
	return o.MentionedMembers
}

func (o *MessagingMessage) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *MessagingMessage) GetParentMessageID() *string {
	if o == nil {
		return nil
	}
	return o.ParentMessageID
}

func (o *MessagingMessage) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *MessagingMessage) GetSubject() *string {
	if o == nil {
		return nil
	}
	return o.Subject
}

func (o *MessagingMessage) GetUpdatedAt() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *MessagingMessage) GetWebURL() *string {
	if o == nil {
		return nil
	}
	return o.WebURL
}
