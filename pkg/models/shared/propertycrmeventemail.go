// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// PropertyCrmEventEmail - The email object, when type = email
type PropertyCrmEventEmail struct {
	AttachmentFileIds []string `json:"attachment_file_ids,omitempty"`
	Body              *string  `json:"body,omitempty"`
	// The event email's cc name & email (name )
	Cc      []string `json:"cc,omitempty"`
	From    *string  `json:"from,omitempty"`
	Subject *string  `json:"subject,omitempty"`
	// The event email's "to" name & email (name )
	To []string `json:"to,omitempty"`
}

func (o *PropertyCrmEventEmail) GetAttachmentFileIds() []string {
	if o == nil {
		return nil
	}
	return o.AttachmentFileIds
}

func (o *PropertyCrmEventEmail) GetBody() *string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *PropertyCrmEventEmail) GetCc() []string {
	if o == nil {
		return nil
	}
	return o.Cc
}

func (o *PropertyCrmEventEmail) GetFrom() *string {
	if o == nil {
		return nil
	}
	return o.From
}

func (o *PropertyCrmEventEmail) GetSubject() *string {
	if o == nil {
		return nil
	}
	return o.Subject
}

func (o *PropertyCrmEventEmail) GetTo() []string {
	if o == nil {
		return nil
	}
	return o.To
}
