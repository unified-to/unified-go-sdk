// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Scopes string

const (
	ScopesAuthLogin                  Scopes = "auth_login"
	ScopesAccountingAccountRead      Scopes = "accounting_account_read"
	ScopesAccountingAccountWrite     Scopes = "accounting_account_write"
	ScopesAccountingTransactionRead  Scopes = "accounting_transaction_read"
	ScopesAccountingTransactionWrite Scopes = "accounting_transaction_write"
	ScopesAccountingInvoiceRead      Scopes = "accounting_invoice_read"
	ScopesAccountingInvoiceWrite     Scopes = "accounting_invoice_write"
	ScopesAccountingCustomerRead     Scopes = "accounting_customer_read"
	ScopesAccountingCustomerWrite    Scopes = "accounting_customer_write"
	ScopesAccountingPaymentRead      Scopes = "accounting_payment_read"
	ScopesAccountingPaymentWrite     Scopes = "accounting_payment_write"
	ScopesAtsScorecardRead           Scopes = "ats_scorecard_read"
	ScopesAtsScorecardWrite          Scopes = "ats_scorecard_write"
	ScopesAtsApplicationRead         Scopes = "ats_application_read"
	ScopesAtsApplicationWrite        Scopes = "ats_application_write"
	ScopesAtsApplicationstatusRead   Scopes = "ats_applicationstatus_read"
	ScopesAtsCandidateRead           Scopes = "ats_candidate_read"
	ScopesAtsCandidateWrite          Scopes = "ats_candidate_write"
	ScopesAtsInterviewRead           Scopes = "ats_interview_read"
	ScopesAtsInterviewWrite          Scopes = "ats_interview_write"
	ScopesAtsJobRead                 Scopes = "ats_job_read"
	ScopesAtsJobWrite                Scopes = "ats_job_write"
	ScopesAtsDocumentRead            Scopes = "ats_document_read"
	ScopesAtsDocumentWrite           Scopes = "ats_document_write"
	ScopesCrmCompanyRead             Scopes = "crm_company_read"
	ScopesCrmCompanyWrite            Scopes = "crm_company_write"
	ScopesCrmContactRead             Scopes = "crm_contact_read"
	ScopesCrmContactWrite            Scopes = "crm_contact_write"
	ScopesCrmDealRead                Scopes = "crm_deal_read"
	ScopesCrmDealWrite               Scopes = "crm_deal_write"
	ScopesCrmEventRead               Scopes = "crm_event_read"
	ScopesCrmEventWrite              Scopes = "crm_event_write"
	ScopesCrmLeadRead                Scopes = "crm_lead_read"
	ScopesCrmLeadWrite               Scopes = "crm_lead_write"
	ScopesCrmFileRead                Scopes = "crm_file_read"
	ScopesCrmFileWrite               Scopes = "crm_file_write"
	ScopesCrmPipelineRead            Scopes = "crm_pipeline_read"
	ScopesCrmPipelineWrite           Scopes = "crm_pipeline_write"
	ScopesMartechListRead            Scopes = "martech_list_read"
	ScopesMartechListWrite           Scopes = "martech_list_write"
	ScopesMartechMemberRead          Scopes = "martech_member_read"
	ScopesMartechMemberWrite         Scopes = "martech_member_write"
	ScopesTicketingCustomerRead      Scopes = "ticketing_customer_read"
	ScopesTicketingCustomerWrite     Scopes = "ticketing_customer_write"
	ScopesTicketingTicketRead        Scopes = "ticketing_ticket_read"
	ScopesTicketingTicketWrite       Scopes = "ticketing_ticket_write"
	ScopesTicketingNoteRead          Scopes = "ticketing_note_read"
	ScopesTicketingNoteWrite         Scopes = "ticketing_note_write"
	ScopesHrisEmployeeRead           Scopes = "hris_employee_read"
	ScopesHrisEmployeeWrite          Scopes = "hris_employee_write"
	ScopesHrisGroupRead              Scopes = "hris_group_read"
	ScopesHrisGroupWrite             Scopes = "hris_group_write"
	ScopesUcCallRead                 Scopes = "uc_call_read"
	ScopesWebhook                    Scopes = "webhook"
)

func (e Scopes) ToPointer() *Scopes {
	return &e
}

func (e *Scopes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "auth_login":
		fallthrough
	case "accounting_account_read":
		fallthrough
	case "accounting_account_write":
		fallthrough
	case "accounting_transaction_read":
		fallthrough
	case "accounting_transaction_write":
		fallthrough
	case "accounting_invoice_read":
		fallthrough
	case "accounting_invoice_write":
		fallthrough
	case "accounting_customer_read":
		fallthrough
	case "accounting_customer_write":
		fallthrough
	case "accounting_payment_read":
		fallthrough
	case "accounting_payment_write":
		fallthrough
	case "ats_scorecard_read":
		fallthrough
	case "ats_scorecard_write":
		fallthrough
	case "ats_application_read":
		fallthrough
	case "ats_application_write":
		fallthrough
	case "ats_applicationstatus_read":
		fallthrough
	case "ats_candidate_read":
		fallthrough
	case "ats_candidate_write":
		fallthrough
	case "ats_interview_read":
		fallthrough
	case "ats_interview_write":
		fallthrough
	case "ats_job_read":
		fallthrough
	case "ats_job_write":
		fallthrough
	case "ats_document_read":
		fallthrough
	case "ats_document_write":
		fallthrough
	case "crm_company_read":
		fallthrough
	case "crm_company_write":
		fallthrough
	case "crm_contact_read":
		fallthrough
	case "crm_contact_write":
		fallthrough
	case "crm_deal_read":
		fallthrough
	case "crm_deal_write":
		fallthrough
	case "crm_event_read":
		fallthrough
	case "crm_event_write":
		fallthrough
	case "crm_lead_read":
		fallthrough
	case "crm_lead_write":
		fallthrough
	case "crm_file_read":
		fallthrough
	case "crm_file_write":
		fallthrough
	case "crm_pipeline_read":
		fallthrough
	case "crm_pipeline_write":
		fallthrough
	case "martech_list_read":
		fallthrough
	case "martech_list_write":
		fallthrough
	case "martech_member_read":
		fallthrough
	case "martech_member_write":
		fallthrough
	case "ticketing_customer_read":
		fallthrough
	case "ticketing_customer_write":
		fallthrough
	case "ticketing_ticket_read":
		fallthrough
	case "ticketing_ticket_write":
		fallthrough
	case "ticketing_note_read":
		fallthrough
	case "ticketing_note_write":
		fallthrough
	case "hris_employee_read":
		fallthrough
	case "hris_employee_write":
		fallthrough
	case "hris_group_read":
		fallthrough
	case "hris_group_write":
		fallthrough
	case "uc_call_read":
		fallthrough
	case "webhook":
		*e = Scopes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Scopes: %v", v)
	}
}

type GetUnifiedIntegrationAuthRequest struct {
	Env *string `queryParam:"style=form,explode=true,name=env"`
	// Your user identifier to associate with the new Integration
	ExternalXref *string `queryParam:"style=form,explode=true,name=external_xref"`
	// The URL where you want the user to be redirect to after an unsuccessful authentication. An "error" variable will be appended.
	FailureRedirect *string `queryParam:"style=form,explode=true,name=failure_redirect"`
	// Type of the supported integration
	IntegrationType string `pathParam:"style=simple,explode=false,name=integration_type"`
	// Language: en, fr, es, it, pt, zh, hi
	Lang     *string  `queryParam:"style=form,explode=true,name=lang"`
	Redirect *bool    `queryParam:"style=form,explode=true,name=redirect"`
	Scopes   []Scopes `queryParam:"style=form,explode=true,name=scopes"`
	// Extra state to send back to your success URL
	State     *string `queryParam:"style=form,explode=true,name=state"`
	Subdomain *string `queryParam:"style=form,explode=true,name=subdomain"`
	// The URL where you want the user to be redirect to after a successful authentication.  The connection ID will be appended with (id=<connectionId>) to this URL, as will the state that was provided.
	SuccessRedirect *string `queryParam:"style=form,explode=true,name=success_redirect"`
	// The ID of the workspace
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
}

func (o *GetUnifiedIntegrationAuthRequest) GetEnv() *string {
	if o == nil {
		return nil
	}
	return o.Env
}

func (o *GetUnifiedIntegrationAuthRequest) GetExternalXref() *string {
	if o == nil {
		return nil
	}
	return o.ExternalXref
}

func (o *GetUnifiedIntegrationAuthRequest) GetFailureRedirect() *string {
	if o == nil {
		return nil
	}
	return o.FailureRedirect
}

func (o *GetUnifiedIntegrationAuthRequest) GetIntegrationType() string {
	if o == nil {
		return ""
	}
	return o.IntegrationType
}

func (o *GetUnifiedIntegrationAuthRequest) GetLang() *string {
	if o == nil {
		return nil
	}
	return o.Lang
}

func (o *GetUnifiedIntegrationAuthRequest) GetRedirect() *bool {
	if o == nil {
		return nil
	}
	return o.Redirect
}

func (o *GetUnifiedIntegrationAuthRequest) GetScopes() []Scopes {
	if o == nil {
		return nil
	}
	return o.Scopes
}

func (o *GetUnifiedIntegrationAuthRequest) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *GetUnifiedIntegrationAuthRequest) GetSubdomain() *string {
	if o == nil {
		return nil
	}
	return o.Subdomain
}

func (o *GetUnifiedIntegrationAuthRequest) GetSuccessRedirect() *string {
	if o == nil {
		return nil
	}
	return o.SuccessRedirect
}

func (o *GetUnifiedIntegrationAuthRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

type GetUnifiedIntegrationAuthResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	Res *string
}

func (o *GetUnifiedIntegrationAuthResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetUnifiedIntegrationAuthResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetUnifiedIntegrationAuthResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetUnifiedIntegrationAuthResponse) GetRes() *string {
	if o == nil {
		return nil
	}
	return o.Res
}
