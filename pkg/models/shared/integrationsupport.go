// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type IntegrationSupport struct {
	InboundFields       map[string]any  `json:"inbound_fields,omitempty"`
	ListAccountID       *bool           `json:"list_account_id,omitempty"`
	ListApplicationID   *bool           `json:"list_application_id,omitempty"`
	ListCandidateID     *bool           `json:"list_candidate_id,omitempty"`
	ListChannelID       *bool           `json:"list_channel_id,omitempty"`
	ListCollectionID    *bool           `json:"list_collection_id,omitempty"`
	ListCompanyID       *bool           `json:"list_company_id,omitempty"`
	ListContactID       *bool           `json:"list_contact_id,omitempty"`
	ListCustomerID      *bool           `json:"list_customer_id,omitempty"`
	ListDealID          *bool           `json:"list_deal_id,omitempty"`
	ListInterviewID     *bool           `json:"list_interview_id,omitempty"`
	ListInvoiceID       *bool           `json:"list_invoice_id,omitempty"`
	ListItemID          *bool           `json:"list_item_id,omitempty"`
	ListJobID           *bool           `json:"list_job_id,omitempty"`
	ListLimit           *bool           `json:"list_limit,omitempty"`
	ListLinkID          *bool           `json:"list_link_id,omitempty"`
	ListListID          *bool           `json:"list_list_id,omitempty"`
	ListLocationID      *bool           `json:"list_location_id,omitempty"`
	ListOffset          *bool           `json:"list_offset,omitempty"`
	ListOrder           *bool           `json:"list_order,omitempty"`
	ListParentID        *bool           `json:"list_parent_id,omitempty"`
	ListQuery           *bool           `json:"list_query,omitempty"`
	ListSortByCreatedAt *bool           `json:"list_sort_by_created_at,omitempty"`
	ListSortByName      *bool           `json:"list_sort_by_name,omitempty"`
	ListSortByUpdatedAt *bool           `json:"list_sort_by_updated_at,omitempty"`
	ListSpaceID         *bool           `json:"list_space_id,omitempty"`
	ListTicketID        *bool           `json:"list_ticket_id,omitempty"`
	ListType            *bool           `json:"list_type,omitempty"`
	ListUpdatedGte      *bool           `json:"list_updated_gte,omitempty"`
	ListUserID          *bool           `json:"list_user_id,omitempty"`
	Methods             map[string]bool `json:"methods,omitempty"`
	OutboundFields      map[string]any  `json:"outbound_fields,omitempty"`
	// objects that we map from in the integration
	RawObjects        []string                                 `json:"raw_objects,omitempty"`
	SearchDomain      *bool                                    `json:"search_domain,omitempty"`
	SearchEmail       *bool                                    `json:"search_email,omitempty"`
	SearchLinkedinurl *bool                                    `json:"search_linkedinurl,omitempty"`
	SearchName        *bool                                    `json:"search_name,omitempty"`
	SearchTwitter     *bool                                    `json:"search_twitter,omitempty"`
	WebhookEvents     *PropertyIntegrationSupportWebhookEvents `json:"webhook_events,omitempty"`
}

func (o *IntegrationSupport) GetInboundFields() map[string]any {
	if o == nil {
		return nil
	}
	return o.InboundFields
}

func (o *IntegrationSupport) GetListAccountID() *bool {
	if o == nil {
		return nil
	}
	return o.ListAccountID
}

func (o *IntegrationSupport) GetListApplicationID() *bool {
	if o == nil {
		return nil
	}
	return o.ListApplicationID
}

func (o *IntegrationSupport) GetListCandidateID() *bool {
	if o == nil {
		return nil
	}
	return o.ListCandidateID
}

func (o *IntegrationSupport) GetListChannelID() *bool {
	if o == nil {
		return nil
	}
	return o.ListChannelID
}

func (o *IntegrationSupport) GetListCollectionID() *bool {
	if o == nil {
		return nil
	}
	return o.ListCollectionID
}

func (o *IntegrationSupport) GetListCompanyID() *bool {
	if o == nil {
		return nil
	}
	return o.ListCompanyID
}

func (o *IntegrationSupport) GetListContactID() *bool {
	if o == nil {
		return nil
	}
	return o.ListContactID
}

func (o *IntegrationSupport) GetListCustomerID() *bool {
	if o == nil {
		return nil
	}
	return o.ListCustomerID
}

func (o *IntegrationSupport) GetListDealID() *bool {
	if o == nil {
		return nil
	}
	return o.ListDealID
}

func (o *IntegrationSupport) GetListInterviewID() *bool {
	if o == nil {
		return nil
	}
	return o.ListInterviewID
}

func (o *IntegrationSupport) GetListInvoiceID() *bool {
	if o == nil {
		return nil
	}
	return o.ListInvoiceID
}

func (o *IntegrationSupport) GetListItemID() *bool {
	if o == nil {
		return nil
	}
	return o.ListItemID
}

func (o *IntegrationSupport) GetListJobID() *bool {
	if o == nil {
		return nil
	}
	return o.ListJobID
}

func (o *IntegrationSupport) GetListLimit() *bool {
	if o == nil {
		return nil
	}
	return o.ListLimit
}

func (o *IntegrationSupport) GetListLinkID() *bool {
	if o == nil {
		return nil
	}
	return o.ListLinkID
}

func (o *IntegrationSupport) GetListListID() *bool {
	if o == nil {
		return nil
	}
	return o.ListListID
}

func (o *IntegrationSupport) GetListLocationID() *bool {
	if o == nil {
		return nil
	}
	return o.ListLocationID
}

func (o *IntegrationSupport) GetListOffset() *bool {
	if o == nil {
		return nil
	}
	return o.ListOffset
}

func (o *IntegrationSupport) GetListOrder() *bool {
	if o == nil {
		return nil
	}
	return o.ListOrder
}

func (o *IntegrationSupport) GetListParentID() *bool {
	if o == nil {
		return nil
	}
	return o.ListParentID
}

func (o *IntegrationSupport) GetListQuery() *bool {
	if o == nil {
		return nil
	}
	return o.ListQuery
}

func (o *IntegrationSupport) GetListSortByCreatedAt() *bool {
	if o == nil {
		return nil
	}
	return o.ListSortByCreatedAt
}

func (o *IntegrationSupport) GetListSortByName() *bool {
	if o == nil {
		return nil
	}
	return o.ListSortByName
}

func (o *IntegrationSupport) GetListSortByUpdatedAt() *bool {
	if o == nil {
		return nil
	}
	return o.ListSortByUpdatedAt
}

func (o *IntegrationSupport) GetListSpaceID() *bool {
	if o == nil {
		return nil
	}
	return o.ListSpaceID
}

func (o *IntegrationSupport) GetListTicketID() *bool {
	if o == nil {
		return nil
	}
	return o.ListTicketID
}

func (o *IntegrationSupport) GetListType() *bool {
	if o == nil {
		return nil
	}
	return o.ListType
}

func (o *IntegrationSupport) GetListUpdatedGte() *bool {
	if o == nil {
		return nil
	}
	return o.ListUpdatedGte
}

func (o *IntegrationSupport) GetListUserID() *bool {
	if o == nil {
		return nil
	}
	return o.ListUserID
}

func (o *IntegrationSupport) GetMethods() map[string]bool {
	if o == nil {
		return nil
	}
	return o.Methods
}

func (o *IntegrationSupport) GetOutboundFields() map[string]any {
	if o == nil {
		return nil
	}
	return o.OutboundFields
}

func (o *IntegrationSupport) GetRawObjects() []string {
	if o == nil {
		return nil
	}
	return o.RawObjects
}

func (o *IntegrationSupport) GetSearchDomain() *bool {
	if o == nil {
		return nil
	}
	return o.SearchDomain
}

func (o *IntegrationSupport) GetSearchEmail() *bool {
	if o == nil {
		return nil
	}
	return o.SearchEmail
}

func (o *IntegrationSupport) GetSearchLinkedinurl() *bool {
	if o == nil {
		return nil
	}
	return o.SearchLinkedinurl
}

func (o *IntegrationSupport) GetSearchName() *bool {
	if o == nil {
		return nil
	}
	return o.SearchName
}

func (o *IntegrationSupport) GetSearchTwitter() *bool {
	if o == nil {
		return nil
	}
	return o.SearchTwitter
}

func (o *IntegrationSupport) GetWebhookEvents() *PropertyIntegrationSupportWebhookEvents {
	if o == nil {
		return nil
	}
	return o.WebhookEvents
}
