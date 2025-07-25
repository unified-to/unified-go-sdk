// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type MessagingReaction struct {
	Member   PropertyMessagingReactionMember `json:"member"`
	Reaction string                          `json:"reaction"`
}

func (o *MessagingReaction) GetMember() PropertyMessagingReactionMember {
	if o == nil {
		return PropertyMessagingReactionMember{}
	}
	return o.Member
}

func (o *MessagingReaction) GetReaction() string {
	if o == nil {
		return ""
	}
	return o.Reaction
}
