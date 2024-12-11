// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type PropertyConnectionPermissions string

const (
	PropertyConnectionPermissionsAuthLogin                  PropertyConnectionPermissions = "auth_login"
	PropertyConnectionPermissionsAccountingAccountRead      PropertyConnectionPermissions = "accounting_account_read"
	PropertyConnectionPermissionsAccountingAccountWrite     PropertyConnectionPermissions = "accounting_account_write"
	PropertyConnectionPermissionsAccountingTransactionRead  PropertyConnectionPermissions = "accounting_transaction_read"
	PropertyConnectionPermissionsAccountingTransactionWrite PropertyConnectionPermissions = "accounting_transaction_write"
	PropertyConnectionPermissionsAccountingJournalRead      PropertyConnectionPermissions = "accounting_journal_read"
	PropertyConnectionPermissionsAccountingJournalWrite     PropertyConnectionPermissions = "accounting_journal_write"
	PropertyConnectionPermissionsAccountingInvoiceRead      PropertyConnectionPermissions = "accounting_invoice_read"
	PropertyConnectionPermissionsAccountingInvoiceWrite     PropertyConnectionPermissions = "accounting_invoice_write"
	PropertyConnectionPermissionsAccountingContactRead      PropertyConnectionPermissions = "accounting_contact_read"
	PropertyConnectionPermissionsAccountingContactWrite     PropertyConnectionPermissions = "accounting_contact_write"
	PropertyConnectionPermissionsAccountingTaxrateRead      PropertyConnectionPermissions = "accounting_taxrate_read"
	PropertyConnectionPermissionsAccountingTaxrateWrite     PropertyConnectionPermissions = "accounting_taxrate_write"
	PropertyConnectionPermissionsAccountingOrganizationRead PropertyConnectionPermissions = "accounting_organization_read"
	PropertyConnectionPermissionsAccountingOrderRead        PropertyConnectionPermissions = "accounting_order_read"
	PropertyConnectionPermissionsAccountingOrderWrite       PropertyConnectionPermissions = "accounting_order_write"
	PropertyConnectionPermissionsPaymentPaymentRead         PropertyConnectionPermissions = "payment_payment_read"
	PropertyConnectionPermissionsPaymentPaymentWrite        PropertyConnectionPermissions = "payment_payment_write"
	PropertyConnectionPermissionsPaymentPayoutRead          PropertyConnectionPermissions = "payment_payout_read"
	PropertyConnectionPermissionsPaymentRefundRead          PropertyConnectionPermissions = "payment_refund_read"
	PropertyConnectionPermissionsPaymentLinkRead            PropertyConnectionPermissions = "payment_link_read"
	PropertyConnectionPermissionsPaymentLinkWrite           PropertyConnectionPermissions = "payment_link_write"
	PropertyConnectionPermissionsCommerceItemRead           PropertyConnectionPermissions = "commerce_item_read"
	PropertyConnectionPermissionsCommerceItemWrite          PropertyConnectionPermissions = "commerce_item_write"
	PropertyConnectionPermissionsCommerceCollectionRead     PropertyConnectionPermissions = "commerce_collection_read"
	PropertyConnectionPermissionsCommerceCollectionWrite    PropertyConnectionPermissions = "commerce_collection_write"
	PropertyConnectionPermissionsCommerceInventoryRead      PropertyConnectionPermissions = "commerce_inventory_read"
	PropertyConnectionPermissionsCommerceInventoryWrite     PropertyConnectionPermissions = "commerce_inventory_write"
	PropertyConnectionPermissionsCommerceLocationRead       PropertyConnectionPermissions = "commerce_location_read"
	PropertyConnectionPermissionsCommerceLocationWrite      PropertyConnectionPermissions = "commerce_location_write"
	PropertyConnectionPermissionsAtsActivityRead            PropertyConnectionPermissions = "ats_activity_read"
	PropertyConnectionPermissionsAtsActivityWrite           PropertyConnectionPermissions = "ats_activity_write"
	PropertyConnectionPermissionsAtsApplicationRead         PropertyConnectionPermissions = "ats_application_read"
	PropertyConnectionPermissionsAtsApplicationWrite        PropertyConnectionPermissions = "ats_application_write"
	PropertyConnectionPermissionsAtsApplicationstatusRead   PropertyConnectionPermissions = "ats_applicationstatus_read"
	PropertyConnectionPermissionsAtsCandidateRead           PropertyConnectionPermissions = "ats_candidate_read"
	PropertyConnectionPermissionsAtsCandidateWrite          PropertyConnectionPermissions = "ats_candidate_write"
	PropertyConnectionPermissionsAtsInterviewRead           PropertyConnectionPermissions = "ats_interview_read"
	PropertyConnectionPermissionsAtsInterviewWrite          PropertyConnectionPermissions = "ats_interview_write"
	PropertyConnectionPermissionsAtsJobRead                 PropertyConnectionPermissions = "ats_job_read"
	PropertyConnectionPermissionsAtsJobWrite                PropertyConnectionPermissions = "ats_job_write"
	PropertyConnectionPermissionsAtsCompanyRead             PropertyConnectionPermissions = "ats_company_read"
	PropertyConnectionPermissionsAtsDocumentRead            PropertyConnectionPermissions = "ats_document_read"
	PropertyConnectionPermissionsAtsDocumentWrite           PropertyConnectionPermissions = "ats_document_write"
	PropertyConnectionPermissionsAtsScorecardRead           PropertyConnectionPermissions = "ats_scorecard_read"
	PropertyConnectionPermissionsAtsScorecardWrite          PropertyConnectionPermissions = "ats_scorecard_write"
	PropertyConnectionPermissionsCrmCompanyRead             PropertyConnectionPermissions = "crm_company_read"
	PropertyConnectionPermissionsCrmCompanyWrite            PropertyConnectionPermissions = "crm_company_write"
	PropertyConnectionPermissionsCrmContactRead             PropertyConnectionPermissions = "crm_contact_read"
	PropertyConnectionPermissionsCrmContactWrite            PropertyConnectionPermissions = "crm_contact_write"
	PropertyConnectionPermissionsCrmDealRead                PropertyConnectionPermissions = "crm_deal_read"
	PropertyConnectionPermissionsCrmDealWrite               PropertyConnectionPermissions = "crm_deal_write"
	PropertyConnectionPermissionsCrmEventRead               PropertyConnectionPermissions = "crm_event_read"
	PropertyConnectionPermissionsCrmEventWrite              PropertyConnectionPermissions = "crm_event_write"
	PropertyConnectionPermissionsCrmLeadRead                PropertyConnectionPermissions = "crm_lead_read"
	PropertyConnectionPermissionsCrmLeadWrite               PropertyConnectionPermissions = "crm_lead_write"
	PropertyConnectionPermissionsCrmPipelineRead            PropertyConnectionPermissions = "crm_pipeline_read"
	PropertyConnectionPermissionsCrmPipelineWrite           PropertyConnectionPermissions = "crm_pipeline_write"
	PropertyConnectionPermissionsMartechListRead            PropertyConnectionPermissions = "martech_list_read"
	PropertyConnectionPermissionsMartechListWrite           PropertyConnectionPermissions = "martech_list_write"
	PropertyConnectionPermissionsMartechMemberRead          PropertyConnectionPermissions = "martech_member_read"
	PropertyConnectionPermissionsMartechMemberWrite         PropertyConnectionPermissions = "martech_member_write"
	PropertyConnectionPermissionsTicketingCustomerRead      PropertyConnectionPermissions = "ticketing_customer_read"
	PropertyConnectionPermissionsTicketingCustomerWrite     PropertyConnectionPermissions = "ticketing_customer_write"
	PropertyConnectionPermissionsTicketingTicketRead        PropertyConnectionPermissions = "ticketing_ticket_read"
	PropertyConnectionPermissionsTicketingTicketWrite       PropertyConnectionPermissions = "ticketing_ticket_write"
	PropertyConnectionPermissionsTicketingNoteRead          PropertyConnectionPermissions = "ticketing_note_read"
	PropertyConnectionPermissionsTicketingNoteWrite         PropertyConnectionPermissions = "ticketing_note_write"
	PropertyConnectionPermissionsHrisEmployeeRead           PropertyConnectionPermissions = "hris_employee_read"
	PropertyConnectionPermissionsHrisEmployeeWrite          PropertyConnectionPermissions = "hris_employee_write"
	PropertyConnectionPermissionsHrisGroupRead              PropertyConnectionPermissions = "hris_group_read"
	PropertyConnectionPermissionsHrisGroupWrite             PropertyConnectionPermissions = "hris_group_write"
	PropertyConnectionPermissionsHrisPayslipRead            PropertyConnectionPermissions = "hris_payslip_read"
	PropertyConnectionPermissionsHrisPayslipWrite           PropertyConnectionPermissions = "hris_payslip_write"
	PropertyConnectionPermissionsHrisTimeoffRead            PropertyConnectionPermissions = "hris_timeoff_read"
	PropertyConnectionPermissionsHrisTimeoffWrite           PropertyConnectionPermissions = "hris_timeoff_write"
	PropertyConnectionPermissionsHrisCompanyRead            PropertyConnectionPermissions = "hris_company_read"
	PropertyConnectionPermissionsHrisCompanyWrite           PropertyConnectionPermissions = "hris_company_write"
	PropertyConnectionPermissionsHrisLocationRead           PropertyConnectionPermissions = "hris_location_read"
	PropertyConnectionPermissionsHrisLocationWrite          PropertyConnectionPermissions = "hris_location_write"
	PropertyConnectionPermissionsUcCallRead                 PropertyConnectionPermissions = "uc_call_read"
	PropertyConnectionPermissionsStorageFileRead            PropertyConnectionPermissions = "storage_file_read"
	PropertyConnectionPermissionsStorageFileWrite           PropertyConnectionPermissions = "storage_file_write"
	PropertyConnectionPermissionsWebhook                    PropertyConnectionPermissions = "webhook"
	PropertyConnectionPermissionsGenaiModelRead             PropertyConnectionPermissions = "genai_model_read"
	PropertyConnectionPermissionsGenaiPromptRead            PropertyConnectionPermissions = "genai_prompt_read"
	PropertyConnectionPermissionsGenaiPromptWrite           PropertyConnectionPermissions = "genai_prompt_write"
	PropertyConnectionPermissionsMessagingMessageRead       PropertyConnectionPermissions = "messaging_message_read"
	PropertyConnectionPermissionsMessagingMessageWrite      PropertyConnectionPermissions = "messaging_message_write"
	PropertyConnectionPermissionsMessagingChannelRead       PropertyConnectionPermissions = "messaging_channel_read"
	PropertyConnectionPermissionsKmsSpaceRead               PropertyConnectionPermissions = "kms_space_read"
	PropertyConnectionPermissionsKmsSpaceWrite              PropertyConnectionPermissions = "kms_space_write"
	PropertyConnectionPermissionsKmsPageRead                PropertyConnectionPermissions = "kms_page_read"
	PropertyConnectionPermissionsKmsPageWrite               PropertyConnectionPermissions = "kms_page_write"
	PropertyConnectionPermissionsKmsCommentRead             PropertyConnectionPermissions = "kms_comment_read"
	PropertyConnectionPermissionsKmsCommentWrite            PropertyConnectionPermissions = "kms_comment_write"
	PropertyConnectionPermissionsTaskProjectRead            PropertyConnectionPermissions = "task_project_read"
	PropertyConnectionPermissionsTaskProjectWrite           PropertyConnectionPermissions = "task_project_write"
	PropertyConnectionPermissionsTaskTaskRead               PropertyConnectionPermissions = "task_task_read"
	PropertyConnectionPermissionsTaskTaskWrite              PropertyConnectionPermissions = "task_task_write"
	PropertyConnectionPermissionsScimUsersRead              PropertyConnectionPermissions = "scim_users_read"
	PropertyConnectionPermissionsScimUsersWrite             PropertyConnectionPermissions = "scim_users_write"
	PropertyConnectionPermissionsScimGroupsRead             PropertyConnectionPermissions = "scim_groups_read"
	PropertyConnectionPermissionsScimGroupsWrite            PropertyConnectionPermissions = "scim_groups_write"
	PropertyConnectionPermissionsLmsCourseRead              PropertyConnectionPermissions = "lms_course_read"
	PropertyConnectionPermissionsLmsCourseWrite             PropertyConnectionPermissions = "lms_course_write"
	PropertyConnectionPermissionsLmsClassRead               PropertyConnectionPermissions = "lms_class_read"
	PropertyConnectionPermissionsLmsClassWrite              PropertyConnectionPermissions = "lms_class_write"
	PropertyConnectionPermissionsLmsStudentRead             PropertyConnectionPermissions = "lms_student_read"
	PropertyConnectionPermissionsLmsStudentWrite            PropertyConnectionPermissions = "lms_student_write"
	PropertyConnectionPermissionsLmsInstructorRead          PropertyConnectionPermissions = "lms_instructor_read"
	PropertyConnectionPermissionsLmsInstructorWrite         PropertyConnectionPermissions = "lms_instructor_write"
	PropertyConnectionPermissionsRepoOrganizationRead       PropertyConnectionPermissions = "repo_organization_read"
	PropertyConnectionPermissionsRepoOrganizationWrite      PropertyConnectionPermissions = "repo_organization_write"
	PropertyConnectionPermissionsRepoRepositoryRead         PropertyConnectionPermissions = "repo_repository_read"
	PropertyConnectionPermissionsRepoRepositoryWrite        PropertyConnectionPermissions = "repo_repository_write"
	PropertyConnectionPermissionsRepoBranchRead             PropertyConnectionPermissions = "repo_branch_read"
	PropertyConnectionPermissionsRepoBranchWrite            PropertyConnectionPermissions = "repo_branch_write"
	PropertyConnectionPermissionsRepoCommitRead             PropertyConnectionPermissions = "repo_commit_read"
	PropertyConnectionPermissionsRepoCommitWrite            PropertyConnectionPermissions = "repo_commit_write"
	PropertyConnectionPermissionsRepoPullrequestRead        PropertyConnectionPermissions = "repo_pullrequest_read"
	PropertyConnectionPermissionsRepoPullrequestWrite       PropertyConnectionPermissions = "repo_pullrequest_write"
	PropertyConnectionPermissionsMetadataMetadataRead       PropertyConnectionPermissions = "metadata_metadata_read"
	PropertyConnectionPermissionsMetadataMetadataWrite      PropertyConnectionPermissions = "metadata_metadata_write"
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
	case "accounting_account_read":
		fallthrough
	case "accounting_account_write":
		fallthrough
	case "accounting_transaction_read":
		fallthrough
	case "accounting_transaction_write":
		fallthrough
	case "accounting_journal_read":
		fallthrough
	case "accounting_journal_write":
		fallthrough
	case "accounting_invoice_read":
		fallthrough
	case "accounting_invoice_write":
		fallthrough
	case "accounting_contact_read":
		fallthrough
	case "accounting_contact_write":
		fallthrough
	case "accounting_taxrate_read":
		fallthrough
	case "accounting_taxrate_write":
		fallthrough
	case "accounting_organization_read":
		fallthrough
	case "accounting_order_read":
		fallthrough
	case "accounting_order_write":
		fallthrough
	case "payment_payment_read":
		fallthrough
	case "payment_payment_write":
		fallthrough
	case "payment_payout_read":
		fallthrough
	case "payment_refund_read":
		fallthrough
	case "payment_link_read":
		fallthrough
	case "payment_link_write":
		fallthrough
	case "commerce_item_read":
		fallthrough
	case "commerce_item_write":
		fallthrough
	case "commerce_collection_read":
		fallthrough
	case "commerce_collection_write":
		fallthrough
	case "commerce_inventory_read":
		fallthrough
	case "commerce_inventory_write":
		fallthrough
	case "commerce_location_read":
		fallthrough
	case "commerce_location_write":
		fallthrough
	case "ats_activity_read":
		fallthrough
	case "ats_activity_write":
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
	case "ats_company_read":
		fallthrough
	case "ats_document_read":
		fallthrough
	case "ats_document_write":
		fallthrough
	case "ats_scorecard_read":
		fallthrough
	case "ats_scorecard_write":
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
	case "hris_payslip_read":
		fallthrough
	case "hris_payslip_write":
		fallthrough
	case "hris_timeoff_read":
		fallthrough
	case "hris_timeoff_write":
		fallthrough
	case "hris_company_read":
		fallthrough
	case "hris_company_write":
		fallthrough
	case "hris_location_read":
		fallthrough
	case "hris_location_write":
		fallthrough
	case "uc_call_read":
		fallthrough
	case "storage_file_read":
		fallthrough
	case "storage_file_write":
		fallthrough
	case "webhook":
		fallthrough
	case "genai_model_read":
		fallthrough
	case "genai_prompt_read":
		fallthrough
	case "genai_prompt_write":
		fallthrough
	case "messaging_message_read":
		fallthrough
	case "messaging_message_write":
		fallthrough
	case "messaging_channel_read":
		fallthrough
	case "kms_space_read":
		fallthrough
	case "kms_space_write":
		fallthrough
	case "kms_page_read":
		fallthrough
	case "kms_page_write":
		fallthrough
	case "kms_comment_read":
		fallthrough
	case "kms_comment_write":
		fallthrough
	case "task_project_read":
		fallthrough
	case "task_project_write":
		fallthrough
	case "task_task_read":
		fallthrough
	case "task_task_write":
		fallthrough
	case "scim_users_read":
		fallthrough
	case "scim_users_write":
		fallthrough
	case "scim_groups_read":
		fallthrough
	case "scim_groups_write":
		fallthrough
	case "lms_course_read":
		fallthrough
	case "lms_course_write":
		fallthrough
	case "lms_class_read":
		fallthrough
	case "lms_class_write":
		fallthrough
	case "lms_student_read":
		fallthrough
	case "lms_student_write":
		fallthrough
	case "lms_instructor_read":
		fallthrough
	case "lms_instructor_write":
		fallthrough
	case "repo_organization_read":
		fallthrough
	case "repo_organization_write":
		fallthrough
	case "repo_repository_read":
		fallthrough
	case "repo_repository_write":
		fallthrough
	case "repo_branch_read":
		fallthrough
	case "repo_branch_write":
		fallthrough
	case "repo_commit_read":
		fallthrough
	case "repo_commit_write":
		fallthrough
	case "repo_pullrequest_read":
		fallthrough
	case "repo_pullrequest_write":
		fallthrough
	case "metadata_metadata_read":
		fallthrough
	case "metadata_metadata_write":
		*e = PropertyConnectionPermissions(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PropertyConnectionPermissions: %v", v)
	}
}
