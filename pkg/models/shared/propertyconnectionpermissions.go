// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type PropertyConnectionPermissions string

const (
	PropertyConnectionPermissionsAuthLogin               PropertyConnectionPermissions = "auth_login"
	PropertyConnectionPermissionsAccountingInvoiceRead   PropertyConnectionPermissions = "accounting_invoice_read"
	PropertyConnectionPermissionsAccountingInvoiceWrite  PropertyConnectionPermissions = "accounting_invoice_write"
	PropertyConnectionPermissionsAccountingCustomerRead  PropertyConnectionPermissions = "accounting_customer_read"
	PropertyConnectionPermissionsAccountingCustomerWrite PropertyConnectionPermissions = "accounting_customer_write"
	PropertyConnectionPermissionsAccountingPaymentRead   PropertyConnectionPermissions = "accounting_payment_read"
	PropertyConnectionPermissionsAccountingPaymentWrite  PropertyConnectionPermissions = "accounting_payment_write"
	PropertyConnectionPermissionsAtsScorecardRead        PropertyConnectionPermissions = "ats_scorecard_read"
	PropertyConnectionPermissionsAtsScorecardWrite       PropertyConnectionPermissions = "ats_scorecard_write"
	PropertyConnectionPermissionsAtsApplicationRead      PropertyConnectionPermissions = "ats_application_read"
	PropertyConnectionPermissionsAtsApplicationWrite     PropertyConnectionPermissions = "ats_application_write"
	PropertyConnectionPermissionsAtsCandidateRead        PropertyConnectionPermissions = "ats_candidate_read"
	PropertyConnectionPermissionsAtsCandidateWrite       PropertyConnectionPermissions = "ats_candidate_write"
	PropertyConnectionPermissionsAtsInterviewRead        PropertyConnectionPermissions = "ats_interview_read"
	PropertyConnectionPermissionsAtsInterviewWrite       PropertyConnectionPermissions = "ats_interview_write"
	PropertyConnectionPermissionsAtsJobRead              PropertyConnectionPermissions = "ats_job_read"
	PropertyConnectionPermissionsAtsJobWrite             PropertyConnectionPermissions = "ats_job_write"
	PropertyConnectionPermissionsAtsDocumentRead         PropertyConnectionPermissions = "ats_document_read"
	PropertyConnectionPermissionsAtsDocumentWrite        PropertyConnectionPermissions = "ats_document_write"
	PropertyConnectionPermissionsCrmCompanyRead          PropertyConnectionPermissions = "crm_company_read"
	PropertyConnectionPermissionsCrmCompanyWrite         PropertyConnectionPermissions = "crm_company_write"
	PropertyConnectionPermissionsCrmContactRead          PropertyConnectionPermissions = "crm_contact_read"
	PropertyConnectionPermissionsCrmContactWrite         PropertyConnectionPermissions = "crm_contact_write"
	PropertyConnectionPermissionsCrmDealRead             PropertyConnectionPermissions = "crm_deal_read"
	PropertyConnectionPermissionsCrmDealWrite            PropertyConnectionPermissions = "crm_deal_write"
	PropertyConnectionPermissionsCrmEventRead            PropertyConnectionPermissions = "crm_event_read"
	PropertyConnectionPermissionsCrmEventWrite           PropertyConnectionPermissions = "crm_event_write"
	PropertyConnectionPermissionsCrmLeadRead             PropertyConnectionPermissions = "crm_lead_read"
	PropertyConnectionPermissionsCrmLeadWrite            PropertyConnectionPermissions = "crm_lead_write"
	PropertyConnectionPermissionsCrmFileRead             PropertyConnectionPermissions = "crm_file_read"
	PropertyConnectionPermissionsCrmFileWrite            PropertyConnectionPermissions = "crm_file_write"
	PropertyConnectionPermissionsCrmPipelineRead         PropertyConnectionPermissions = "crm_pipeline_read"
	PropertyConnectionPermissionsCrmPipelineWrite        PropertyConnectionPermissions = "crm_pipeline_write"
	PropertyConnectionPermissionsMartechListRead         PropertyConnectionPermissions = "martech_list_read"
	PropertyConnectionPermissionsMartechListWrite        PropertyConnectionPermissions = "martech_list_write"
	PropertyConnectionPermissionsMartechMemberRead       PropertyConnectionPermissions = "martech_member_read"
	PropertyConnectionPermissionsMartechMemberWrite      PropertyConnectionPermissions = "martech_member_write"
	PropertyConnectionPermissionsTicketingCustomerRead   PropertyConnectionPermissions = "ticketing_customer_read"
	PropertyConnectionPermissionsTicketingCustomerWrite  PropertyConnectionPermissions = "ticketing_customer_write"
	PropertyConnectionPermissionsTicketingTicketRead     PropertyConnectionPermissions = "ticketing_ticket_read"
	PropertyConnectionPermissionsTicketingTicketWrite    PropertyConnectionPermissions = "ticketing_ticket_write"
	PropertyConnectionPermissionsTicketingNoteRead       PropertyConnectionPermissions = "ticketing_note_read"
	PropertyConnectionPermissionsTicketingNoteWrite      PropertyConnectionPermissions = "ticketing_note_write"
	PropertyConnectionPermissionsHrisEmployeeRead        PropertyConnectionPermissions = "hris_employee_read"
	PropertyConnectionPermissionsHrisEmployeeWrite       PropertyConnectionPermissions = "hris_employee_write"
	PropertyConnectionPermissionsHrisGroupRead           PropertyConnectionPermissions = "hris_group_read"
	PropertyConnectionPermissionsHrisGroupWrite          PropertyConnectionPermissions = "hris_group_write"
	PropertyConnectionPermissionsUcCallRead              PropertyConnectionPermissions = "uc_call_read"
	PropertyConnectionPermissionsWebhook                 PropertyConnectionPermissions = "webhook"
)

func (e PropertyConnectionPermissions) ToPointer() *PropertyConnectionPermissions {
	return &e
}

func (e *PropertyConnectionPermissions) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "auth_login":
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
		*e = PropertyConnectionPermissions(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PropertyConnectionPermissions: %v", v)
	}
}
