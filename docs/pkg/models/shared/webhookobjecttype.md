# WebhookObjectType

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.WebhookObjectTypeAccountingAccount

// Open enum: custom values can be created with a direct type cast
custom := shared.WebhookObjectType("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `WebhookObjectTypeAccountingAccount`        | accounting_account                          |
| `WebhookObjectTypeAccountingTransaction`    | accounting_transaction                      |
| `WebhookObjectTypeAccountingJournal`        | accounting_journal                          |
| `WebhookObjectTypeAccountingContact`        | accounting_contact                          |
| `WebhookObjectTypeAccountingInvoice`        | accounting_invoice                          |
| `WebhookObjectTypeAccountingBill`           | accounting_bill                             |
| `WebhookObjectTypeAccountingVendorcredit`   | accounting_vendorcredit                     |
| `WebhookObjectTypeAccountingCreditmemo`     | accounting_creditmemo                       |
| `WebhookObjectTypeAccountingTaxrate`        | accounting_taxrate                          |
| `WebhookObjectTypeAccountingOrganization`   | accounting_organization                     |
| `WebhookObjectTypeAccountingOrder`          | accounting_order                            |
| `WebhookObjectTypeAccountingSalesorder`     | accounting_salesorder                       |
| `WebhookObjectTypeAccountingPurchaseorder`  | accounting_purchaseorder                    |
| `WebhookObjectTypeAccountingReport`         | accounting_report                           |
| `WebhookObjectTypeAccountingBalancesheet`   | accounting_balancesheet                     |
| `WebhookObjectTypeAccountingProfitloss`     | accounting_profitloss                       |
| `WebhookObjectTypeAccountingTrialbalance`   | accounting_trialbalance                     |
| `WebhookObjectTypeAccountingCategory`       | accounting_category                         |
| `WebhookObjectTypeAccountingExpense`        | accounting_expense                          |
| `WebhookObjectTypeAccountingCashflow`       | accounting_cashflow                         |
| `WebhookObjectTypeAccountingQuote`          | accounting_quote                            |
| `WebhookObjectTypeAccountingAgedreceivable` | accounting_agedreceivable                   |
| `WebhookObjectTypeAccountingAgedpayable`    | accounting_agedpayable                      |
| `WebhookObjectTypeAccountingProject`        | accounting_project                          |
| `WebhookObjectTypePaymentPayment`           | payment_payment                             |
| `WebhookObjectTypePaymentLink`              | payment_link                                |
| `WebhookObjectTypePaymentPayout`            | payment_payout                              |
| `WebhookObjectTypePaymentRefund`            | payment_refund                              |
| `WebhookObjectTypePaymentSubscription`      | payment_subscription                        |
| `WebhookObjectTypeCommerceItem`             | commerce_item                               |
| `WebhookObjectTypeCommerceCollection`       | commerce_collection                         |
| `WebhookObjectTypeCommerceInventory`        | commerce_inventory                          |
| `WebhookObjectTypeCommerceLocation`         | commerce_location                           |
| `WebhookObjectTypeCommerceReview`           | commerce_review                             |
| `WebhookObjectTypeCommerceSaleschannel`     | commerce_saleschannel                       |
| `WebhookObjectTypeCommerceItemvariant`      | commerce_itemvariant                        |
| `WebhookObjectTypeCommerceReservation`      | commerce_reservation                        |
| `WebhookObjectTypeCommerceAvailability`     | commerce_availability                       |
| `WebhookObjectTypeVerificationPackage`      | verification_package                        |
| `WebhookObjectTypeVerificationRequest`      | verification_request                        |
| `WebhookObjectTypeAssessmentPackage`        | assessment_package                          |
| `WebhookObjectTypeAssessmentOrder`          | assessment_order                            |
| `WebhookObjectTypeAtsActivity`              | ats_activity                                |
| `WebhookObjectTypeAtsApplication`           | ats_application                             |
| `WebhookObjectTypeAtsApplicationstatus`     | ats_applicationstatus                       |
| `WebhookObjectTypeAtsCandidate`             | ats_candidate                               |
| `WebhookObjectTypeAtsDocument`              | ats_document                                |
| `WebhookObjectTypeAtsInterview`             | ats_interview                               |
| `WebhookObjectTypeAtsJob`                   | ats_job                                     |
| `WebhookObjectTypeAtsScorecard`             | ats_scorecard                               |
| `WebhookObjectTypeAtsCompany`               | ats_company                                 |
| `WebhookObjectTypeCrmCompany`               | crm_company                                 |
| `WebhookObjectTypeCrmContact`               | crm_contact                                 |
| `WebhookObjectTypeCrmDeal`                  | crm_deal                                    |
| `WebhookObjectTypeCrmEvent`                 | crm_event                                   |
| `WebhookObjectTypeCrmLead`                  | crm_lead                                    |
| `WebhookObjectTypeCrmPipeline`              | crm_pipeline                                |
| `WebhookObjectTypeCrmPicklist`              | crm_picklist                                |
| `WebhookObjectTypeHrisEmployee`             | hris_employee                               |
| `WebhookObjectTypeHrisGroup`                | hris_group                                  |
| `WebhookObjectTypeHrisPayslip`              | hris_payslip                                |
| `WebhookObjectTypeHrisTimeoff`              | hris_timeoff                                |
| `WebhookObjectTypeHrisCompany`              | hris_company                                |
| `WebhookObjectTypeHrisLocation`             | hris_location                               |
| `WebhookObjectTypeHrisDevice`               | hris_device                                 |
| `WebhookObjectTypeHrisTimeshift`            | hris_timeshift                              |
| `WebhookObjectTypeHrisDeduction`            | hris_deduction                              |
| `WebhookObjectTypeHrisBenefit`              | hris_benefit                                |
| `WebhookObjectTypeHrisBankaccount`          | hris_bankaccount                            |
| `WebhookObjectTypeHrisDocument`             | hris_document                               |
| `WebhookObjectTypeHrisTaxonomy`             | hris_taxonomy                               |
| `WebhookObjectTypeMartechList`              | martech_list                                |
| `WebhookObjectTypeMartechMember`            | martech_member                              |
| `WebhookObjectTypeMartechCampaign`          | martech_campaign                            |
| `WebhookObjectTypeMartechReport`            | martech_report                              |
| `WebhookObjectTypePassthrough`              | passthrough                                 |
| `WebhookObjectTypeTicketingNote`            | ticketing_note                              |
| `WebhookObjectTypeTicketingTicket`          | ticketing_ticket                            |
| `WebhookObjectTypeTicketingCustomer`        | ticketing_customer                          |
| `WebhookObjectTypeTicketingCategory`        | ticketing_category                          |
| `WebhookObjectTypeUcContact`                | uc_contact                                  |
| `WebhookObjectTypeUcCall`                   | uc_call                                     |
| `WebhookObjectTypeUcComment`                | uc_comment                                  |
| `WebhookObjectTypeUcRecording`              | uc_recording                                |
| `WebhookObjectTypeEnrichPerson`             | enrich_person                               |
| `WebhookObjectTypeEnrichCompany`            | enrich_company                              |
| `WebhookObjectTypeStorageFile`              | storage_file                                |
| `WebhookObjectTypeGenaiModel`               | genai_model                                 |
| `WebhookObjectTypeGenaiPrompt`              | genai_prompt                                |
| `WebhookObjectTypeGenaiEmbedding`           | genai_embedding                             |
| `WebhookObjectTypeMessagingMessage`         | messaging_message                           |
| `WebhookObjectTypeMessagingChannel`         | messaging_channel                           |
| `WebhookObjectTypeMessagingEvent`           | messaging_event                             |
| `WebhookObjectTypeKmsSpace`                 | kms_space                                   |
| `WebhookObjectTypeKmsPage`                  | kms_page                                    |
| `WebhookObjectTypeKmsComment`               | kms_comment                                 |
| `WebhookObjectTypeTaskProject`              | task_project                                |
| `WebhookObjectTypeTaskTask`                 | task_task                                   |
| `WebhookObjectTypeTaskComment`              | task_comment                                |
| `WebhookObjectTypeTaskChange`               | task_change                                 |
| `WebhookObjectTypeScimUsers`                | scim_users                                  |
| `WebhookObjectTypeScimGroups`               | scim_groups                                 |
| `WebhookObjectTypeLmsCourse`                | lms_course                                  |
| `WebhookObjectTypeLmsClass`                 | lms_class                                   |
| `WebhookObjectTypeLmsStudent`               | lms_student                                 |
| `WebhookObjectTypeLmsInstructor`            | lms_instructor                              |
| `WebhookObjectTypeLmsContent`               | lms_content                                 |
| `WebhookObjectTypeLmsCollection`            | lms_collection                              |
| `WebhookObjectTypeLmsActivity`              | lms_activity                                |
| `WebhookObjectTypeRepoOrganization`         | repo_organization                           |
| `WebhookObjectTypeRepoRepository`           | repo_repository                             |
| `WebhookObjectTypeRepoBranch`               | repo_branch                                 |
| `WebhookObjectTypeRepoCommit`               | repo_commit                                 |
| `WebhookObjectTypeRepoPullrequest`          | repo_pullrequest                            |
| `WebhookObjectTypeMetadataMetadata`         | metadata_metadata                           |
| `WebhookObjectTypeCalendarCalendar`         | calendar_calendar                           |
| `WebhookObjectTypeCalendarEvent`            | calendar_event                              |
| `WebhookObjectTypeCalendarBusy`             | calendar_busy                               |
| `WebhookObjectTypeCalendarLink`             | calendar_link                               |
| `WebhookObjectTypeCalendarRecording`        | calendar_recording                          |
| `WebhookObjectTypeCalendarWebinar`          | calendar_webinar                            |
| `WebhookObjectTypeAdsOrganization`          | ads_organization                            |
| `WebhookObjectTypeAdsAd`                    | ads_ad                                      |
| `WebhookObjectTypeAdsCampaign`              | ads_campaign                                |
| `WebhookObjectTypeAdsReport`                | ads_report                                  |
| `WebhookObjectTypeAdsGroup`                 | ads_group                                   |
| `WebhookObjectTypeAdsCreative`              | ads_creative                                |
| `WebhookObjectTypeAdsAsset`                 | ads_asset                                   |
| `WebhookObjectTypeAdsInsertionorder`        | ads_insertionorder                          |
| `WebhookObjectTypeAdsTarget`                | ads_target                                  |
| `WebhookObjectTypeAdsPromoted`              | ads_promoted                                |
| `WebhookObjectTypeAnalyticsProperty`        | analytics_property                          |
| `WebhookObjectTypeAnalyticsEvent`           | analytics_event                             |
| `WebhookObjectTypeAnalyticsSession`         | analytics_session                           |
| `WebhookObjectTypeAnalyticsVisitor`         | analytics_visitor                           |
| `WebhookObjectTypeAnalyticsReport`          | analytics_report                            |
| `WebhookObjectTypeFormsForm`                | forms_form                                  |
| `WebhookObjectTypeFormsSubmission`          | forms_submission                            |
| `WebhookObjectTypeShippingCarrier`          | shipping_carrier                            |
| `WebhookObjectTypeShippingRate`             | shipping_rate                               |
| `WebhookObjectTypeShippingShipment`         | shipping_shipment                           |
| `WebhookObjectTypeShippingLabel`            | shipping_label                              |
| `WebhookObjectTypeShippingTracking`         | shipping_tracking                           |
| `WebhookObjectTypeSigningDocument`          | signing_document                            |
| `WebhookObjectTypeSigningSignatory`         | signing_signatory                           |
| `WebhookObjectTypeSigningTemplate`          | signing_template                            |
| `WebhookObjectTypeClubsGroup`               | clubs_group                                 |
| `WebhookObjectTypeClubsMember`              | clubs_member                                |
| `WebhookObjectTypeClubsActivity`            | clubs_activity                              |
| `WebhookObjectTypeClubsLocation`            | clubs_location                              |
| `WebhookObjectTypeClubsEvent`               | clubs_event                                 |
| `WebhookObjectTypeDatastoreDatabase`        | datastore_database                          |
| `WebhookObjectTypeDatastoreTable`           | datastore_table                             |
| `WebhookObjectTypeDatastoreRecord`          | datastore_record                            |
| `WebhookObjectTypeDatastoreQuery`           | datastore_query                             |
| `WebhookObjectTypeCdpProfile`               | cdp_profile                                 |
| `WebhookObjectTypeCdpSegment`               | cdp_segment                                 |
| `WebhookObjectTypeCdpEvent`                 | cdp_event                                   |
| `WebhookObjectTypeCdpSource`                | cdp_source                                  |
| `WebhookObjectTypeCdpDestination`           | cdp_destination                             |
| `WebhookObjectTypeCdpActivation`            | cdp_activation                              |
| `WebhookObjectTypePerformanceCycle`         | performance_cycle                           |
| `WebhookObjectTypePerformanceReview`        | performance_review                          |
| `WebhookObjectTypePerformanceGoal`          | performance_goal                            |
| `WebhookObjectTypePerformanceFeedback`      | performance_feedback                        |