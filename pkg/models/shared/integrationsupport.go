// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type WebhookType string

const (
	WebhookTypeVirtual WebhookType = "virtual"
	WebhookTypeNone    WebhookType = "none"
	WebhookTypeNative  WebhookType = "native"
)

func (e WebhookType) ToPointer() *WebhookType {
	return &e
}

func (e *WebhookType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "virtual":
		fallthrough
	case "none":
		fallthrough
	case "native":
		*e = WebhookType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WebhookType: %v", v)
	}
}

type IntegrationSupport struct {
	InboundFields       *PropertyIntegrationSupportInboundFields  `json:"inbound_fields,omitempty"`
	ListAgentID         *bool                                     `json:"list_agent_id,omitempty"`
	ListApplicationID   *bool                                     `json:"list_application_id,omitempty"`
	ListCandidateID     *bool                                     `json:"list_candidate_id,omitempty"`
	ListCompanyID       *bool                                     `json:"list_company_id,omitempty"`
	ListContactID       *bool                                     `json:"list_contact_id,omitempty"`
	ListCustomerID      *bool                                     `json:"list_customer_id,omitempty"`
	ListDealID          *bool                                     `json:"list_deal_id,omitempty"`
	ListInvoiceID       *bool                                     `json:"list_invoice_id,omitempty"`
	ListJobID           *bool                                     `json:"list_job_id,omitempty"`
	ListLimit           *bool                                     `json:"list_limit,omitempty"`
	ListOffset          *bool                                     `json:"list_offset,omitempty"`
	ListOrder           *bool                                     `json:"list_order,omitempty"`
	ListQuery           *bool                                     `json:"list_query,omitempty"`
	ListSortByCreatedAt *bool                                     `json:"list_sort_by_created_at,omitempty"`
	ListSortByName      *bool                                     `json:"list_sort_by_name,omitempty"`
	ListSortByUpdatedAt *bool                                     `json:"list_sort_by_updated_at,omitempty"`
	ListUpdatedGte      *bool                                     `json:"list_updated_gte,omitempty"`
	Methods             map[string]bool                           `json:"methods,omitempty"`
	OutboundFields      *PropertyIntegrationSupportOutboundFields `json:"outbound_fields,omitempty"`
	SearchDomain        *bool                                     `json:"search_domain,omitempty"`
	SearchEmail         *bool                                     `json:"search_email,omitempty"`
	SearchLinkedinURL   *bool                                     `json:"search_linkedin_url,omitempty"`
	SearchName          *bool                                     `json:"search_name,omitempty"`
	SearchTwitter       *bool                                     `json:"search_twitter,omitempty"`
	WebhookEvents       []PropertyIntegrationSupportWebhookEvents `json:"webhook_events,omitempty"`
	WebhookType         *WebhookType                              `json:"webhook_type,omitempty"`
}

func (o *IntegrationSupport) GetInboundFields() *PropertyIntegrationSupportInboundFields {
	if o == nil {
		return nil
	}
	return o.InboundFields
}

func (o *IntegrationSupport) GetListAgentID() *bool {
	if o == nil {
		return nil
	}
	return o.ListAgentID
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

func (o *IntegrationSupport) GetListInvoiceID() *bool {
	if o == nil {
		return nil
	}
	return o.ListInvoiceID
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

func (o *IntegrationSupport) GetListUpdatedGte() *bool {
	if o == nil {
		return nil
	}
	return o.ListUpdatedGte
}

func (o *IntegrationSupport) GetMethods() map[string]bool {
	if o == nil {
		return nil
	}
	return o.Methods
}

func (o *IntegrationSupport) GetOutboundFields() *PropertyIntegrationSupportOutboundFields {
	if o == nil {
		return nil
	}
	return o.OutboundFields
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

func (o *IntegrationSupport) GetSearchLinkedinURL() *bool {
	if o == nil {
		return nil
	}
	return o.SearchLinkedinURL
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

func (o *IntegrationSupport) GetWebhookEvents() []PropertyIntegrationSupportWebhookEvents {
	if o == nil {
		return nil
	}
	return o.WebhookEvents
}

func (o *IntegrationSupport) GetWebhookType() *WebhookType {
	if o == nil {
		return nil
	}
	return o.WebhookType
}
