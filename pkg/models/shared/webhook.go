// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type Event string

const (
	EventUpdated Event = "updated"
	EventCreated Event = "created"
	EventDeleted Event = "deleted"
)

func (e Event) ToPointer() *Event {
	return &e
}
func (e *Event) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "updated":
		fallthrough
	case "created":
		fallthrough
	case "deleted":
		*e = Event(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Event: %v", v)
	}
}

type ObjectType string

const (
	ObjectTypeAccountingAccount      ObjectType = "accounting_account"
	ObjectTypeAccountingTransaction  ObjectType = "accounting_transaction"
	ObjectTypeAccountingJournal      ObjectType = "accounting_journal"
	ObjectTypeAccountingContact      ObjectType = "accounting_contact"
	ObjectTypeAccountingInvoice      ObjectType = "accounting_invoice"
	ObjectTypeAccountingTaxrate      ObjectType = "accounting_taxrate"
	ObjectTypeAccountingOrganization ObjectType = "accounting_organization"
	ObjectTypeAccountingOrder        ObjectType = "accounting_order"
	ObjectTypePaymentPayment         ObjectType = "payment_payment"
	ObjectTypePaymentLink            ObjectType = "payment_link"
	ObjectTypePaymentPayout          ObjectType = "payment_payout"
	ObjectTypePaymentRefund          ObjectType = "payment_refund"
	ObjectTypePaymentSubscription    ObjectType = "payment_subscription"
	ObjectTypeCommerceItem           ObjectType = "commerce_item"
	ObjectTypeCommerceCollection     ObjectType = "commerce_collection"
	ObjectTypeCommerceInventory      ObjectType = "commerce_inventory"
	ObjectTypeCommerceLocation       ObjectType = "commerce_location"
	ObjectTypeAtsActivity            ObjectType = "ats_activity"
	ObjectTypeAtsApplication         ObjectType = "ats_application"
	ObjectTypeAtsApplicationstatus   ObjectType = "ats_applicationstatus"
	ObjectTypeAtsCandidate           ObjectType = "ats_candidate"
	ObjectTypeAtsDocument            ObjectType = "ats_document"
	ObjectTypeAtsInterview           ObjectType = "ats_interview"
	ObjectTypeAtsJob                 ObjectType = "ats_job"
	ObjectTypeAtsScorecard           ObjectType = "ats_scorecard"
	ObjectTypeAtsCompany             ObjectType = "ats_company"
	ObjectTypeCrmCompany             ObjectType = "crm_company"
	ObjectTypeCrmContact             ObjectType = "crm_contact"
	ObjectTypeCrmDeal                ObjectType = "crm_deal"
	ObjectTypeCrmEvent               ObjectType = "crm_event"
	ObjectTypeCrmLead                ObjectType = "crm_lead"
	ObjectTypeCrmPipeline            ObjectType = "crm_pipeline"
	ObjectTypeHrisEmployee           ObjectType = "hris_employee"
	ObjectTypeHrisGroup              ObjectType = "hris_group"
	ObjectTypeHrisPayslip            ObjectType = "hris_payslip"
	ObjectTypeHrisTimeoff            ObjectType = "hris_timeoff"
	ObjectTypeHrisCompany            ObjectType = "hris_company"
	ObjectTypeHrisLocation           ObjectType = "hris_location"
	ObjectTypeMartechList            ObjectType = "martech_list"
	ObjectTypeMartechMember          ObjectType = "martech_member"
	ObjectTypePassthrough            ObjectType = "passthrough"
	ObjectTypeTicketingNote          ObjectType = "ticketing_note"
	ObjectTypeTicketingTicket        ObjectType = "ticketing_ticket"
	ObjectTypeTicketingCustomer      ObjectType = "ticketing_customer"
	ObjectTypeUcContact              ObjectType = "uc_contact"
	ObjectTypeUcCall                 ObjectType = "uc_call"
	ObjectTypeEnrichPerson           ObjectType = "enrich_person"
	ObjectTypeEnrichCompany          ObjectType = "enrich_company"
	ObjectTypeStorageFile            ObjectType = "storage_file"
	ObjectTypeGenaiModel             ObjectType = "genai_model"
	ObjectTypeGenaiPrompt            ObjectType = "genai_prompt"
	ObjectTypeMessagingMessage       ObjectType = "messaging_message"
	ObjectTypeMessagingChannel       ObjectType = "messaging_channel"
	ObjectTypeKmsSpace               ObjectType = "kms_space"
	ObjectTypeKmsPage                ObjectType = "kms_page"
	ObjectTypeTaskProject            ObjectType = "task_project"
	ObjectTypeTaskTask               ObjectType = "task_task"
	ObjectTypeTaskComment            ObjectType = "task_comment"
	ObjectTypeScimUsers              ObjectType = "scim_users"
	ObjectTypeScimGroups             ObjectType = "scim_groups"
	ObjectTypeLmsCourse              ObjectType = "lms_course"
	ObjectTypeLmsClass               ObjectType = "lms_class"
	ObjectTypeLmsStudent             ObjectType = "lms_student"
	ObjectTypeLmsInstructor          ObjectType = "lms_instructor"
	ObjectTypeRepoOrganization       ObjectType = "repo_organization"
	ObjectTypeRepoRepository         ObjectType = "repo_repository"
	ObjectTypeRepoBranch             ObjectType = "repo_branch"
	ObjectTypeRepoCommit             ObjectType = "repo_commit"
	ObjectTypeRepoPullrequest        ObjectType = "repo_pullrequest"
	ObjectTypeMetadataMetadata       ObjectType = "metadata_metadata"
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
	case "accounting_journal":
		fallthrough
	case "accounting_contact":
		fallthrough
	case "accounting_invoice":
		fallthrough
	case "accounting_taxrate":
		fallthrough
	case "accounting_organization":
		fallthrough
	case "accounting_order":
		fallthrough
	case "payment_payment":
		fallthrough
	case "payment_link":
		fallthrough
	case "payment_payout":
		fallthrough
	case "payment_refund":
		fallthrough
	case "payment_subscription":
		fallthrough
	case "commerce_item":
		fallthrough
	case "commerce_collection":
		fallthrough
	case "commerce_inventory":
		fallthrough
	case "commerce_location":
		fallthrough
	case "ats_activity":
		fallthrough
	case "ats_application":
		fallthrough
	case "ats_applicationstatus":
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
	case "ats_company":
		fallthrough
	case "crm_company":
		fallthrough
	case "crm_contact":
		fallthrough
	case "crm_deal":
		fallthrough
	case "crm_event":
		fallthrough
	case "crm_lead":
		fallthrough
	case "crm_pipeline":
		fallthrough
	case "hris_employee":
		fallthrough
	case "hris_group":
		fallthrough
	case "hris_payslip":
		fallthrough
	case "hris_timeoff":
		fallthrough
	case "hris_company":
		fallthrough
	case "hris_location":
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
		fallthrough
	case "storage_file":
		fallthrough
	case "genai_model":
		fallthrough
	case "genai_prompt":
		fallthrough
	case "messaging_message":
		fallthrough
	case "messaging_channel":
		fallthrough
	case "kms_space":
		fallthrough
	case "kms_page":
		fallthrough
	case "task_project":
		fallthrough
	case "task_task":
		fallthrough
	case "task_comment":
		fallthrough
	case "scim_users":
		fallthrough
	case "scim_groups":
		fallthrough
	case "lms_course":
		fallthrough
	case "lms_class":
		fallthrough
	case "lms_student":
		fallthrough
	case "lms_instructor":
		fallthrough
	case "repo_organization":
		fallthrough
	case "repo_repository":
		fallthrough
	case "repo_branch":
		fallthrough
	case "repo_commit":
		fallthrough
	case "repo_pullrequest":
		fallthrough
	case "metadata_metadata":
		*e = ObjectType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ObjectType: %v", v)
	}
}

type WebhookType string

const (
	WebhookTypeVirtual WebhookType = "virtual"
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
	case "native":
		*e = WebhookType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WebhookType: %v", v)
	}
}

// Webhook - A webhook is used to POST new/updated information to your server.
type Webhook struct {
	CheckedAt       *time.Time        `json:"checked_at,omitempty"`
	ConnectionID    string            `json:"connection_id"`
	CreatedAt       *time.Time        `json:"created_at,omitempty"`
	Environment     *string           `default:"Production" json:"environment"`
	Event           Event             `json:"event"`
	Fields          *string           `json:"fields,omitempty"`
	Filters         map[string]string `json:"filters,omitempty"`
	HookURL         string            `json:"hook_url"`
	ID              *string           `json:"id,omitempty"`
	IntegrationType *string           `json:"integration_type,omitempty"`
	Interval        *float64          `json:"interval,omitempty"`
	IsHealthy       *bool             `json:"is_healthy,omitempty"`
	Meta            map[string]any    `json:"meta,omitempty"`
	ObjectType      ObjectType        `json:"object_type"`
	PageMaxLimit    *float64          `json:"page_max_limit,omitempty"`
	// An array of the most revent virtual webhook runs
	Runs        []string     `json:"runs,omitempty"`
	UpdatedAt   *time.Time   `json:"updated_at,omitempty"`
	WebhookType *WebhookType `json:"webhook_type,omitempty"`
	WorkspaceID *string      `json:"workspace_id,omitempty"`
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

func (o *Webhook) GetEvent() Event {
	if o == nil {
		return Event("")
	}
	return o.Event
}

func (o *Webhook) GetFields() *string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *Webhook) GetFilters() map[string]string {
	if o == nil {
		return nil
	}
	return o.Filters
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

func (o *Webhook) GetIntegrationType() *string {
	if o == nil {
		return nil
	}
	return o.IntegrationType
}

func (o *Webhook) GetInterval() *float64 {
	if o == nil {
		return nil
	}
	return o.Interval
}

func (o *Webhook) GetIsHealthy() *bool {
	if o == nil {
		return nil
	}
	return o.IsHealthy
}

func (o *Webhook) GetMeta() map[string]any {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *Webhook) GetObjectType() ObjectType {
	if o == nil {
		return ObjectType("")
	}
	return o.ObjectType
}

func (o *Webhook) GetPageMaxLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.PageMaxLimit
}

func (o *Webhook) GetRuns() []string {
	if o == nil {
		return nil
	}
	return o.Runs
}

func (o *Webhook) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *Webhook) GetWebhookType() *WebhookType {
	if o == nil {
		return nil
	}
	return o.WebhookType
}

func (o *Webhook) GetWorkspaceID() *string {
	if o == nil {
		return nil
	}
	return o.WorkspaceID
}
