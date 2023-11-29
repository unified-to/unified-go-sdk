// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type ObjectType string

const (
	ObjectTypeAccountingAccount     ObjectType = "accounting_account"
	ObjectTypeAccountingTransaction ObjectType = "accounting_transaction"
	ObjectTypeAccountingCustomer    ObjectType = "accounting_customer"
	ObjectTypeAccountingInvoice     ObjectType = "accounting_invoice"
	ObjectTypeAccountingPayment     ObjectType = "accounting_payment"
	ObjectTypeAtsApplication        ObjectType = "ats_application"
	ObjectTypeAtsCandidate          ObjectType = "ats_candidate"
	ObjectTypeAtsDocument           ObjectType = "ats_document"
	ObjectTypeAtsInterview          ObjectType = "ats_interview"
	ObjectTypeAtsJob                ObjectType = "ats_job"
	ObjectTypeAtsScorecard          ObjectType = "ats_scorecard"
	ObjectTypeCrmCompany            ObjectType = "crm_company"
	ObjectTypeCrmContact            ObjectType = "crm_contact"
	ObjectTypeCrmDeal               ObjectType = "crm_deal"
	ObjectTypeCrmEvent              ObjectType = "crm_event"
	ObjectTypeCrmFile               ObjectType = "crm_file"
	ObjectTypeCrmLead               ObjectType = "crm_lead"
	ObjectTypeCrmPipeline           ObjectType = "crm_pipeline"
	ObjectTypeHrisEmployee          ObjectType = "hris_employee"
	ObjectTypeHrisGroup             ObjectType = "hris_group"
	ObjectTypeMartechList           ObjectType = "martech_list"
	ObjectTypeMartechMember         ObjectType = "martech_member"
	ObjectTypePassthrough           ObjectType = "passthrough"
	ObjectTypeTicketingNote         ObjectType = "ticketing_note"
	ObjectTypeTicketingTicket       ObjectType = "ticketing_ticket"
	ObjectTypeTicketingCustomer     ObjectType = "ticketing_customer"
	ObjectTypeUcContact             ObjectType = "uc_contact"
	ObjectTypeUcCall                ObjectType = "uc_call"
	ObjectTypeEnrichPerson          ObjectType = "enrich_person"
	ObjectTypeEnrichCompany         ObjectType = "enrich_company"
)

func (e ObjectType) ToPointer() *ObjectType {
	return &e
}

func (e *ObjectType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "accounting_account":
		fallthrough
	case "accounting_transaction":
		fallthrough
	case "accounting_customer":
		fallthrough
	case "accounting_invoice":
		fallthrough
	case "accounting_payment":
		fallthrough
	case "ats_application":
		fallthrough
	case "ats_candidate":
		fallthrough
	case "ats_document":
		fallthrough
	case "ats_interview":
		fallthrough
	case "ats_job":
		fallthrough
	case "ats_scorecard":
		fallthrough
	case "crm_company":
		fallthrough
	case "crm_contact":
		fallthrough
	case "crm_deal":
		fallthrough
	case "crm_event":
		fallthrough
	case "crm_file":
		fallthrough
	case "crm_lead":
		fallthrough
	case "crm_pipeline":
		fallthrough
	case "hris_employee":
		fallthrough
	case "hris_group":
		fallthrough
	case "martech_list":
		fallthrough
	case "martech_member":
		fallthrough
	case "passthrough":
		fallthrough
	case "ticketing_note":
		fallthrough
	case "ticketing_ticket":
		fallthrough
	case "ticketing_customer":
		fallthrough
	case "uc_contact":
		fallthrough
	case "uc_call":
		fallthrough
	case "enrich_person":
		fallthrough
	case "enrich_company":
		*e = ObjectType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ObjectType: %v", v)
	}
}

// Webhook - A webhook is used to POST new/updated information to your server.
type Webhook struct {
	CheckedAt       *time.Time              `json:"checked_at,omitempty"`
	ConnectionID    string                  `json:"connection_id"`
	CreatedAt       *time.Time              `json:"created_at,omitempty"`
	Environment     *string                 `default:"Production" json:"environment"`
	Events          []PropertyWebhookEvents `json:"events"`
	HookURL         string                  `json:"hook_url"`
	ID              *string                 `json:"id,omitempty"`
	IncludeRaw      *bool                   `json:"include_raw,omitempty"`
	IntegrationType string                  `json:"integration_type"`
	Interval        float64                 `json:"interval"`
	ObjectType      ObjectType              `json:"object_type"`
	// integration-specific subscriptions IDs
	Subscriptions []string   `json:"subscriptions,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
	WorkspaceID   string     `json:"workspace_id"`
}

func (w Webhook) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *Webhook) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Webhook) GetCheckedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CheckedAt
}

func (o *Webhook) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *Webhook) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Webhook) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *Webhook) GetEvents() []PropertyWebhookEvents {
	if o == nil {
		return []PropertyWebhookEvents{}
	}
	return o.Events
}

func (o *Webhook) GetHookURL() string {
	if o == nil {
		return ""
	}
	return o.HookURL
}

func (o *Webhook) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Webhook) GetIncludeRaw() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeRaw
}

func (o *Webhook) GetIntegrationType() string {
	if o == nil {
		return ""
	}
	return o.IntegrationType
}

func (o *Webhook) GetInterval() float64 {
	if o == nil {
		return 0.0
	}
	return o.Interval
}

func (o *Webhook) GetObjectType() ObjectType {
	if o == nil {
		return ObjectType("")
	}
	return o.ObjectType
}

func (o *Webhook) GetSubscriptions() []string {
	if o == nil {
		return nil
	}
	return o.Subscriptions
}

func (o *Webhook) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *Webhook) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
