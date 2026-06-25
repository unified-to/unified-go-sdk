<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://github.com/unified-to/unified-go-sdk/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/bolt-php/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
    
</div>

<!-- Start Summary [summary] -->
## Summary

Unified.to  API: One API to Rule Them All

For more information about the API: [API Documentation](https://docs.unified.to)
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
  * [Authentication](#authentication)
  * [Retries](#retries)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/unified-to/unified-go-sdk
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"log"
)

func main() {
	ctx := context.Background()

	s := unifiedgosdk.New(
		unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Accounting.CreateAccountingAccount2(ctx, operations.CreateAccountingAccount2Request{
		AccountingAccount: shared.AccountingAccount{},
		ConnectionID:      "<id>",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.AccountingAccount != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Account](docs/sdks/account/README.md)

* [CreateAccountingAccount2](docs/sdks/account/README.md#createaccountingaccount2) - Create an account
* [GetAccountingAccount2](docs/sdks/account/README.md#getaccountingaccount2) - Retrieve an account
* [ListAccountingAccounts2](docs/sdks/account/README.md#listaccountingaccounts2) - List all accounts
* [PatchAccountingAccount2](docs/sdks/account/README.md#patchaccountingaccount2) - Update an account
* [RemoveAccountingAccount2](docs/sdks/account/README.md#removeaccountingaccount2) - Remove an account
* [UpdateAccountingAccount2](docs/sdks/account/README.md#updateaccountingaccount2) - Update an account

### [Accounting](docs/sdks/accounting/README.md)

* [CreateAccountingAccount2](docs/sdks/accounting/README.md#createaccountingaccount2) - Create an account
* [CreateAccountingBill2](docs/sdks/accounting/README.md#createaccountingbill2) - Create a bill
* [CreateAccountingCategory2](docs/sdks/accounting/README.md#createaccountingcategory2) - Create a category
* [CreateAccountingContact2](docs/sdks/accounting/README.md#createaccountingcontact2) - Create a contact
* [CreateAccountingCreditmemo2](docs/sdks/accounting/README.md#createaccountingcreditmemo2) - Create a creditmemo
* [CreateAccountingExpense2](docs/sdks/accounting/README.md#createaccountingexpense2) - Create an expense
* [CreateAccountingInvoice2](docs/sdks/accounting/README.md#createaccountinginvoice2) - Create an invoice
* [CreateAccountingJournal2](docs/sdks/accounting/README.md#createaccountingjournal2) - Create a journal
* [CreateAccountingOrder2](docs/sdks/accounting/README.md#createaccountingorder2) - Create an order
* [CreateAccountingPurchaseorder2](docs/sdks/accounting/README.md#createaccountingpurchaseorder2) - Create a purchaseorder
* [CreateAccountingSalesorder2](docs/sdks/accounting/README.md#createaccountingsalesorder2) - Create a salesorder
* [CreateAccountingTaxrate2](docs/sdks/accounting/README.md#createaccountingtaxrate2) - Create a taxrate
* [CreateAccountingTransaction2](docs/sdks/accounting/README.md#createaccountingtransaction2) - Create a transaction
* [GetAccountingAccount2](docs/sdks/accounting/README.md#getaccountingaccount2) - Retrieve an account
* [GetAccountingBalancesheet2](docs/sdks/accounting/README.md#getaccountingbalancesheet2) - Retrieve a balancesheet
* [GetAccountingBill2](docs/sdks/accounting/README.md#getaccountingbill2) - Retrieve a bill
* [GetAccountingCashflow2](docs/sdks/accounting/README.md#getaccountingcashflow2) - Retrieve a cashflow
* [GetAccountingCategory2](docs/sdks/accounting/README.md#getaccountingcategory2) - Retrieve a category
* [GetAccountingContact2](docs/sdks/accounting/README.md#getaccountingcontact2) - Retrieve a contact
* [GetAccountingCreditmemo2](docs/sdks/accounting/README.md#getaccountingcreditmemo2) - Retrieve a creditmemo
* [GetAccountingExpense2](docs/sdks/accounting/README.md#getaccountingexpense2) - Retrieve an expense
* [GetAccountingInvoice2](docs/sdks/accounting/README.md#getaccountinginvoice2) - Retrieve an invoice
* [GetAccountingJournal2](docs/sdks/accounting/README.md#getaccountingjournal2) - Retrieve a journal
* [GetAccountingOrder2](docs/sdks/accounting/README.md#getaccountingorder2) - Retrieve an order
* [GetAccountingOrganization2](docs/sdks/accounting/README.md#getaccountingorganization2) - Retrieve an organization
* [GetAccountingProfitloss2](docs/sdks/accounting/README.md#getaccountingprofitloss2) - Retrieve a profitloss
* [GetAccountingPurchaseorder2](docs/sdks/accounting/README.md#getaccountingpurchaseorder2) - Retrieve a purchaseorder
* [GetAccountingReport2](docs/sdks/accounting/README.md#getaccountingreport2) - Retrieve a report
* [GetAccountingSalesorder2](docs/sdks/accounting/README.md#getaccountingsalesorder2) - Retrieve a salesorder
* [GetAccountingTaxrate2](docs/sdks/accounting/README.md#getaccountingtaxrate2) - Retrieve a taxrate
* [GetAccountingTransaction2](docs/sdks/accounting/README.md#getaccountingtransaction2) - Retrieve a transaction
* [GetAccountingTrialbalance2](docs/sdks/accounting/README.md#getaccountingtrialbalance2) - Retrieve a trialbalance
* [ListAccountingAccounts2](docs/sdks/accounting/README.md#listaccountingaccounts2) - List all accounts
* [ListAccountingBalancesheets2](docs/sdks/accounting/README.md#listaccountingbalancesheets2) - List all balancesheets
* [ListAccountingBills2](docs/sdks/accounting/README.md#listaccountingbills2) - List all bills
* [ListAccountingCashflows2](docs/sdks/accounting/README.md#listaccountingcashflows2) - List all cashflows
* [ListAccountingCategories2](docs/sdks/accounting/README.md#listaccountingcategories2) - List all categories
* [ListAccountingContacts2](docs/sdks/accounting/README.md#listaccountingcontacts2) - List all contacts
* [ListAccountingCreditmemoes2](docs/sdks/accounting/README.md#listaccountingcreditmemoes2) - List all creditmemoes
* [ListAccountingExpenses2](docs/sdks/accounting/README.md#listaccountingexpenses2) - List all expenses
* [ListAccountingInvoices2](docs/sdks/accounting/README.md#listaccountinginvoices2) - List all invoices
* [ListAccountingJournals2](docs/sdks/accounting/README.md#listaccountingjournals2) - List all journals
* [ListAccountingOrders2](docs/sdks/accounting/README.md#listaccountingorders2) - List all orders
* [ListAccountingOrganizations2](docs/sdks/accounting/README.md#listaccountingorganizations2) - List all organizations
* [ListAccountingProfitlosses2](docs/sdks/accounting/README.md#listaccountingprofitlosses2) - List all profitlosses
* [ListAccountingPurchaseorders2](docs/sdks/accounting/README.md#listaccountingpurchaseorders2) - List all purchaseorders
* [ListAccountingReports2](docs/sdks/accounting/README.md#listaccountingreports2) - List all reports
* [ListAccountingSalesorders2](docs/sdks/accounting/README.md#listaccountingsalesorders2) - List all salesorders
* [ListAccountingTaxrates2](docs/sdks/accounting/README.md#listaccountingtaxrates2) - List all taxrates
* [ListAccountingTransactions2](docs/sdks/accounting/README.md#listaccountingtransactions2) - List all transactions
* [ListAccountingTrialbalances2](docs/sdks/accounting/README.md#listaccountingtrialbalances2) - List all trialbalances
* [PatchAccountingAccount2](docs/sdks/accounting/README.md#patchaccountingaccount2) - Update an account
* [PatchAccountingBill2](docs/sdks/accounting/README.md#patchaccountingbill2) - Update a bill
* [PatchAccountingCategory2](docs/sdks/accounting/README.md#patchaccountingcategory2) - Update a category
* [PatchAccountingContact2](docs/sdks/accounting/README.md#patchaccountingcontact2) - Update a contact
* [PatchAccountingCreditmemo2](docs/sdks/accounting/README.md#patchaccountingcreditmemo2) - Update a creditmemo
* [PatchAccountingExpense2](docs/sdks/accounting/README.md#patchaccountingexpense2) - Update an expense
* [PatchAccountingInvoice2](docs/sdks/accounting/README.md#patchaccountinginvoice2) - Update an invoice
* [PatchAccountingJournal2](docs/sdks/accounting/README.md#patchaccountingjournal2) - Update a journal
* [PatchAccountingOrder2](docs/sdks/accounting/README.md#patchaccountingorder2) - Update an order
* [PatchAccountingPurchaseorder2](docs/sdks/accounting/README.md#patchaccountingpurchaseorder2) - Update a purchaseorder
* [PatchAccountingSalesorder2](docs/sdks/accounting/README.md#patchaccountingsalesorder2) - Update a salesorder
* [PatchAccountingTaxrate2](docs/sdks/accounting/README.md#patchaccountingtaxrate2) - Update a taxrate
* [PatchAccountingTransaction2](docs/sdks/accounting/README.md#patchaccountingtransaction2) - Update a transaction
* [RemoveAccountingAccount2](docs/sdks/accounting/README.md#removeaccountingaccount2) - Remove an account
* [RemoveAccountingBill2](docs/sdks/accounting/README.md#removeaccountingbill2) - Remove a bill
* [RemoveAccountingCategory2](docs/sdks/accounting/README.md#removeaccountingcategory2) - Remove a category
* [RemoveAccountingContact2](docs/sdks/accounting/README.md#removeaccountingcontact2) - Remove a contact
* [RemoveAccountingCreditmemo2](docs/sdks/accounting/README.md#removeaccountingcreditmemo2) - Remove a creditmemo
* [RemoveAccountingExpense2](docs/sdks/accounting/README.md#removeaccountingexpense2) - Remove an expense
* [RemoveAccountingInvoice2](docs/sdks/accounting/README.md#removeaccountinginvoice2) - Remove an invoice
* [RemoveAccountingJournal2](docs/sdks/accounting/README.md#removeaccountingjournal2) - Remove a journal
* [RemoveAccountingOrder2](docs/sdks/accounting/README.md#removeaccountingorder2) - Remove an order
* [RemoveAccountingPurchaseorder2](docs/sdks/accounting/README.md#removeaccountingpurchaseorder2) - Remove a purchaseorder
* [RemoveAccountingSalesorder2](docs/sdks/accounting/README.md#removeaccountingsalesorder2) - Remove a salesorder
* [RemoveAccountingTaxrate2](docs/sdks/accounting/README.md#removeaccountingtaxrate2) - Remove a taxrate
* [RemoveAccountingTransaction2](docs/sdks/accounting/README.md#removeaccountingtransaction2) - Remove a transaction
* [UpdateAccountingAccount2](docs/sdks/accounting/README.md#updateaccountingaccount2) - Update an account
* [UpdateAccountingBill2](docs/sdks/accounting/README.md#updateaccountingbill2) - Update a bill
* [UpdateAccountingCategory2](docs/sdks/accounting/README.md#updateaccountingcategory2) - Update a category
* [UpdateAccountingContact2](docs/sdks/accounting/README.md#updateaccountingcontact2) - Update a contact
* [UpdateAccountingCreditmemo2](docs/sdks/accounting/README.md#updateaccountingcreditmemo2) - Update a creditmemo
* [UpdateAccountingExpense2](docs/sdks/accounting/README.md#updateaccountingexpense2) - Update an expense
* [UpdateAccountingInvoice2](docs/sdks/accounting/README.md#updateaccountinginvoice2) - Update an invoice
* [UpdateAccountingJournal2](docs/sdks/accounting/README.md#updateaccountingjournal2) - Update a journal
* [UpdateAccountingOrder2](docs/sdks/accounting/README.md#updateaccountingorder2) - Update an order
* [UpdateAccountingPurchaseorder2](docs/sdks/accounting/README.md#updateaccountingpurchaseorder2) - Update a purchaseorder
* [UpdateAccountingSalesorder2](docs/sdks/accounting/README.md#updateaccountingsalesorder2) - Update a salesorder
* [UpdateAccountingTaxrate2](docs/sdks/accounting/README.md#updateaccountingtaxrate2) - Update a taxrate
* [UpdateAccountingTransaction2](docs/sdks/accounting/README.md#updateaccountingtransaction2) - Update a transaction

### [Activity](docs/sdks/activity/README.md)

* [CreateAtsActivity2](docs/sdks/activity/README.md#createatsactivity2) - Create an activity
* [CreateLmsActivity2](docs/sdks/activity/README.md#createlmsactivity2) - Create an activity
* [GetAtsActivity2](docs/sdks/activity/README.md#getatsactivity2) - Retrieve an activity
* [GetClubsActivity2](docs/sdks/activity/README.md#getclubsactivity2) - Retrieve an activity
* [GetLmsActivity2](docs/sdks/activity/README.md#getlmsactivity2) - Retrieve an activity
* [ListAtsActivities2](docs/sdks/activity/README.md#listatsactivities2) - List all activities
* [ListClubsActivities2](docs/sdks/activity/README.md#listclubsactivities2) - List all activities
* [ListLmsActivities2](docs/sdks/activity/README.md#listlmsactivities2) - List all activities
* [PatchAtsActivity2](docs/sdks/activity/README.md#patchatsactivity2) - Update an activity
* [PatchLmsActivity2](docs/sdks/activity/README.md#patchlmsactivity2) - Update an activity
* [RemoveAtsActivity2](docs/sdks/activity/README.md#removeatsactivity2) - Remove an activity
* [RemoveLmsActivity2](docs/sdks/activity/README.md#removelmsactivity2) - Remove an activity
* [UpdateAtsActivity2](docs/sdks/activity/README.md#updateatsactivity2) - Update an activity
* [UpdateLmsActivity2](docs/sdks/activity/README.md#updatelmsactivity2) - Update an activity

### [Ad](docs/sdks/ad/README.md)

* [CreateAdsAd2](docs/sdks/ad/README.md#createadsad2) - Create an ad
* [GetAdsAd2](docs/sdks/ad/README.md#getadsad2) - Retrieve an ad
* [ListAdsAds2](docs/sdks/ad/README.md#listadsads2) - List all ads
* [PatchAdsAd2](docs/sdks/ad/README.md#patchadsad2) - Update an ad
* [RemoveAdsAd2](docs/sdks/ad/README.md#removeadsad2) - Remove an ad
* [UpdateAdsAd2](docs/sdks/ad/README.md#updateadsad2) - Update an ad

### [Ads](docs/sdks/ads/README.md)

* [CreateAdsAd2](docs/sdks/ads/README.md#createadsad2) - Create an ad
* [CreateAdsCampaign2](docs/sdks/ads/README.md#createadscampaign2) - Create a campaign
* [CreateAdsCreative2](docs/sdks/ads/README.md#createadscreative2) - Create a creative
* [CreateAdsGroup2](docs/sdks/ads/README.md#createadsgroup2) - Create a group
* [CreateAdsInsertionorder2](docs/sdks/ads/README.md#createadsinsertionorder2) - Create an insertionorder
* [CreateAdsOrganization2](docs/sdks/ads/README.md#createadsorganization2) - Create an organization
* [GetAdsAd2](docs/sdks/ads/README.md#getadsad2) - Retrieve an ad
* [GetAdsCampaign2](docs/sdks/ads/README.md#getadscampaign2) - Retrieve a campaign
* [GetAdsCreative2](docs/sdks/ads/README.md#getadscreative2) - Retrieve a creative
* [GetAdsGroup2](docs/sdks/ads/README.md#getadsgroup2) - Retrieve a group
* [GetAdsInsertionorder2](docs/sdks/ads/README.md#getadsinsertionorder2) - Retrieve an insertionorder
* [GetAdsOrganization2](docs/sdks/ads/README.md#getadsorganization2) - Retrieve an organization
* [GetAdsPromoted2](docs/sdks/ads/README.md#getadspromoted2) - Retrieve a promoted
* [GetAdsTarget2](docs/sdks/ads/README.md#getadstarget2) - Retrieve a target
* [ListAdsAds2](docs/sdks/ads/README.md#listadsads2) - List all ads
* [ListAdsCampaigns2](docs/sdks/ads/README.md#listadscampaigns2) - List all campaigns
* [ListAdsCreatives2](docs/sdks/ads/README.md#listadscreatives2) - List all creatives
* [ListAdsGroups2](docs/sdks/ads/README.md#listadsgroups2) - List all groups
* [ListAdsInsertionorders2](docs/sdks/ads/README.md#listadsinsertionorders2) - List all insertionorders
* [ListAdsOrganizations2](docs/sdks/ads/README.md#listadsorganizations2) - List all organizations
* [ListAdsPromoteds2](docs/sdks/ads/README.md#listadspromoteds2) - List all promoteds
* [ListAdsReports2](docs/sdks/ads/README.md#listadsreports2) - List all reports
* [ListAdsTargets2](docs/sdks/ads/README.md#listadstargets2) - List all targets
* [PatchAdsAd2](docs/sdks/ads/README.md#patchadsad2) - Update an ad
* [PatchAdsCampaign2](docs/sdks/ads/README.md#patchadscampaign2) - Update a campaign
* [PatchAdsCreative2](docs/sdks/ads/README.md#patchadscreative2) - Update a creative
* [PatchAdsGroup2](docs/sdks/ads/README.md#patchadsgroup2) - Update a group
* [PatchAdsInsertionorder2](docs/sdks/ads/README.md#patchadsinsertionorder2) - Update an insertionorder
* [PatchAdsOrganization2](docs/sdks/ads/README.md#patchadsorganization2) - Update an organization
* [RemoveAdsAd2](docs/sdks/ads/README.md#removeadsad2) - Remove an ad
* [RemoveAdsCampaign2](docs/sdks/ads/README.md#removeadscampaign2) - Remove a campaign
* [RemoveAdsCreative2](docs/sdks/ads/README.md#removeadscreative2) - Remove a creative
* [RemoveAdsGroup2](docs/sdks/ads/README.md#removeadsgroup2) - Remove a group
* [RemoveAdsInsertionorder2](docs/sdks/ads/README.md#removeadsinsertionorder2) - Remove an insertionorder
* [RemoveAdsOrganization2](docs/sdks/ads/README.md#removeadsorganization2) - Remove an organization
* [UpdateAdsAd2](docs/sdks/ads/README.md#updateadsad2) - Update an ad
* [UpdateAdsCampaign2](docs/sdks/ads/README.md#updateadscampaign2) - Update a campaign
* [UpdateAdsCreative2](docs/sdks/ads/README.md#updateadscreative2) - Update a creative
* [UpdateAdsGroup2](docs/sdks/ads/README.md#updateadsgroup2) - Update a group
* [UpdateAdsInsertionorder2](docs/sdks/ads/README.md#updateadsinsertionorder2) - Update an insertionorder
* [UpdateAdsOrganization2](docs/sdks/ads/README.md#updateadsorganization2) - Update an organization

### [Analytics](docs/sdks/analytics/README.md)

* [CreateAnalyticsEvent2](docs/sdks/analytics/README.md#createanalyticsevent2) - Create an event
* [CreateAnalyticsProperty2](docs/sdks/analytics/README.md#createanalyticsproperty2) - Create a property
* [CreateAnalyticsVisitor2](docs/sdks/analytics/README.md#createanalyticsvisitor2) - Create a visitor
* [GetAnalyticsEvent2](docs/sdks/analytics/README.md#getanalyticsevent2) - Retrieve an event
* [GetAnalyticsProperty2](docs/sdks/analytics/README.md#getanalyticsproperty2) - Retrieve a property
* [GetAnalyticsSession2](docs/sdks/analytics/README.md#getanalyticssession2) - Retrieve a session
* [GetAnalyticsVisitor2](docs/sdks/analytics/README.md#getanalyticsvisitor2) - Retrieve a visitor
* [ListAnalyticsEvents2](docs/sdks/analytics/README.md#listanalyticsevents2) - List all events
* [ListAnalyticsProperties2](docs/sdks/analytics/README.md#listanalyticsproperties2) - List all properties
* [ListAnalyticsReports2](docs/sdks/analytics/README.md#listanalyticsreports2) - List all reports
* [ListAnalyticsSessions2](docs/sdks/analytics/README.md#listanalyticssessions2) - List all sessions
* [ListAnalyticsVisitors2](docs/sdks/analytics/README.md#listanalyticsvisitors2) - List all visitors
* [PatchAnalyticsProperty2](docs/sdks/analytics/README.md#patchanalyticsproperty2) - Update a property
* [PatchAnalyticsVisitor2](docs/sdks/analytics/README.md#patchanalyticsvisitor2) - Update a visitor
* [RemoveAnalyticsProperty2](docs/sdks/analytics/README.md#removeanalyticsproperty2) - Remove a property
* [RemoveAnalyticsVisitor2](docs/sdks/analytics/README.md#removeanalyticsvisitor2) - Remove a visitor
* [UpdateAnalyticsProperty2](docs/sdks/analytics/README.md#updateanalyticsproperty2) - Update a property
* [UpdateAnalyticsVisitor2](docs/sdks/analytics/README.md#updateanalyticsvisitor2) - Update a visitor

### [Apicall](docs/sdks/apicall/README.md)

* [GetUnifiedApicall](docs/sdks/apicall/README.md#getunifiedapicall) - Retrieve specific API Call by its ID
* [ListUnifiedApicalls](docs/sdks/apicall/README.md#listunifiedapicalls) - Returns API Calls

### [Application](docs/sdks/application/README.md)

* [CreateAtsApplication2](docs/sdks/application/README.md#createatsapplication2) - Create an application
* [GetAtsApplication2](docs/sdks/application/README.md#getatsapplication2) - Retrieve an application
* [ListAtsApplications2](docs/sdks/application/README.md#listatsapplications2) - List all applications
* [PatchAtsApplication2](docs/sdks/application/README.md#patchatsapplication2) - Update an application
* [RemoveAtsApplication2](docs/sdks/application/README.md#removeatsapplication2) - Remove an application
* [UpdateAtsApplication2](docs/sdks/application/README.md#updateatsapplication2) - Update an application

### [Applicationstatus](docs/sdks/applicationstatus/README.md)

* [ListAtsApplicationstatuses2](docs/sdks/applicationstatus/README.md#listatsapplicationstatuses2) - List all applicationstatuses

### [Assessment](docs/sdks/assessment/README.md)

* [CreateAssessmentPackage2](docs/sdks/assessment/README.md#createassessmentpackage2) - Create an assessment package
* [GetAssessmentPackage2](docs/sdks/assessment/README.md#getassessmentpackage2) - Get an assessment package
* [ListAssessmentPackages2](docs/sdks/assessment/README.md#listassessmentpackages2) - List assessment packages
* [PatchAssessmentOrder2](docs/sdks/assessment/README.md#patchassessmentorder2) - Update an order
* [PatchAssessmentPackage2](docs/sdks/assessment/README.md#patchassessmentpackage2) - Update an assessment package
* [RemoveAssessmentPackage2](docs/sdks/assessment/README.md#removeassessmentpackage2) - Delete an assessment package
* [UpdateAssessmentOrder2](docs/sdks/assessment/README.md#updateassessmentorder2) - Update an order
* [UpdateAssessmentPackage2](docs/sdks/assessment/README.md#updateassessmentpackage2) - Update an assessment package

### [Ats](docs/sdks/ats/README.md)

* [CreateAtsActivity2](docs/sdks/ats/README.md#createatsactivity2) - Create an activity
* [CreateAtsApplication2](docs/sdks/ats/README.md#createatsapplication2) - Create an application
* [CreateAtsCandidate2](docs/sdks/ats/README.md#createatscandidate2) - Create a candidate
* [CreateAtsCompany2](docs/sdks/ats/README.md#createatscompany2) - Create a company
* [CreateAtsDocument2](docs/sdks/ats/README.md#createatsdocument2) - Create a document
* [CreateAtsInterview2](docs/sdks/ats/README.md#createatsinterview2) - Create an interview
* [CreateAtsJob2](docs/sdks/ats/README.md#createatsjob2) - Create a job
* [CreateAtsScorecard2](docs/sdks/ats/README.md#createatsscorecard2) - Create a scorecard
* [GetAtsActivity2](docs/sdks/ats/README.md#getatsactivity2) - Retrieve an activity
* [GetAtsApplication2](docs/sdks/ats/README.md#getatsapplication2) - Retrieve an application
* [GetAtsCandidate2](docs/sdks/ats/README.md#getatscandidate2) - Retrieve a candidate
* [GetAtsCompany2](docs/sdks/ats/README.md#getatscompany2) - Retrieve a company
* [GetAtsDocument2](docs/sdks/ats/README.md#getatsdocument2) - Retrieve a document
* [GetAtsInterview2](docs/sdks/ats/README.md#getatsinterview2) - Retrieve an interview
* [GetAtsJob2](docs/sdks/ats/README.md#getatsjob2) - Retrieve a job
* [GetAtsScorecard2](docs/sdks/ats/README.md#getatsscorecard2) - Retrieve a scorecard
* [ListAtsActivities2](docs/sdks/ats/README.md#listatsactivities2) - List all activities
* [ListAtsApplications2](docs/sdks/ats/README.md#listatsapplications2) - List all applications
* [ListAtsApplicationstatuses2](docs/sdks/ats/README.md#listatsapplicationstatuses2) - List all applicationstatuses
* [ListAtsCandidates2](docs/sdks/ats/README.md#listatscandidates2) - List all candidates
* [ListAtsCompanies2](docs/sdks/ats/README.md#listatscompanies2) - List all companies
* [ListAtsDocuments2](docs/sdks/ats/README.md#listatsdocuments2) - List all documents
* [ListAtsInterviews2](docs/sdks/ats/README.md#listatsinterviews2) - List all interviews
* [ListAtsJobs2](docs/sdks/ats/README.md#listatsjobs2) - List all jobs
* [ListAtsScorecards2](docs/sdks/ats/README.md#listatsscorecards2) - List all scorecards
* [PatchAtsActivity2](docs/sdks/ats/README.md#patchatsactivity2) - Update an activity
* [PatchAtsApplication2](docs/sdks/ats/README.md#patchatsapplication2) - Update an application
* [PatchAtsCandidate2](docs/sdks/ats/README.md#patchatscandidate2) - Update a candidate
* [PatchAtsCompany2](docs/sdks/ats/README.md#patchatscompany2) - Update a company
* [PatchAtsDocument2](docs/sdks/ats/README.md#patchatsdocument2) - Update a document
* [PatchAtsInterview2](docs/sdks/ats/README.md#patchatsinterview2) - Update an interview
* [PatchAtsJob2](docs/sdks/ats/README.md#patchatsjob2) - Update a job
* [PatchAtsScorecard2](docs/sdks/ats/README.md#patchatsscorecard2) - Update a scorecard
* [RemoveAtsActivity2](docs/sdks/ats/README.md#removeatsactivity2) - Remove an activity
* [RemoveAtsApplication2](docs/sdks/ats/README.md#removeatsapplication2) - Remove an application
* [RemoveAtsCandidate2](docs/sdks/ats/README.md#removeatscandidate2) - Remove a candidate
* [RemoveAtsCompany2](docs/sdks/ats/README.md#removeatscompany2) - Remove a company
* [RemoveAtsDocument2](docs/sdks/ats/README.md#removeatsdocument2) - Remove a document
* [RemoveAtsInterview2](docs/sdks/ats/README.md#removeatsinterview2) - Remove an interview
* [RemoveAtsJob2](docs/sdks/ats/README.md#removeatsjob2) - Remove a job
* [RemoveAtsScorecard2](docs/sdks/ats/README.md#removeatsscorecard2) - Remove a scorecard
* [UpdateAtsActivity2](docs/sdks/ats/README.md#updateatsactivity2) - Update an activity
* [UpdateAtsApplication2](docs/sdks/ats/README.md#updateatsapplication2) - Update an application
* [UpdateAtsCandidate2](docs/sdks/ats/README.md#updateatscandidate2) - Update a candidate
* [UpdateAtsCompany2](docs/sdks/ats/README.md#updateatscompany2) - Update a company
* [UpdateAtsDocument2](docs/sdks/ats/README.md#updateatsdocument2) - Update a document
* [UpdateAtsInterview2](docs/sdks/ats/README.md#updateatsinterview2) - Update an interview
* [UpdateAtsJob2](docs/sdks/ats/README.md#updateatsjob2) - Update a job
* [UpdateAtsScorecard2](docs/sdks/ats/README.md#updateatsscorecard2) - Update a scorecard

### [Auth](docs/sdks/auth/README.md)

* [GetUnifiedIntegrationAuth](docs/sdks/auth/README.md#getunifiedintegrationauth) - Authorize new connection
* [GetUnifiedIntegrationLogin](docs/sdks/auth/README.md#getunifiedintegrationlogin) - Sign in a user

### [Availability](docs/sdks/availability/README.md)

* [ListCommerceAvailabilities2](docs/sdks/availability/README.md#listcommerceavailabilities2) - List all availabilities

### [Balancesheet](docs/sdks/balancesheet/README.md)

* [GetAccountingBalancesheet2](docs/sdks/balancesheet/README.md#getaccountingbalancesheet2) - Retrieve a balancesheet
* [ListAccountingBalancesheets2](docs/sdks/balancesheet/README.md#listaccountingbalancesheets2) - List all balancesheets

### [Bankaccount](docs/sdks/bankaccount/README.md)

* [CreateHrisBankaccount2](docs/sdks/bankaccount/README.md#createhrisbankaccount2) - Create a bankaccount
* [GetHrisBankaccount2](docs/sdks/bankaccount/README.md#gethrisbankaccount2) - Retrieve a bankaccount
* [ListHrisBankaccounts2](docs/sdks/bankaccount/README.md#listhrisbankaccounts2) - List all bankaccounts
* [PatchHrisBankaccount2](docs/sdks/bankaccount/README.md#patchhrisbankaccount2) - Update a bankaccount
* [RemoveHrisBankaccount2](docs/sdks/bankaccount/README.md#removehrisbankaccount2) - Remove a bankaccount
* [UpdateHrisBankaccount2](docs/sdks/bankaccount/README.md#updatehrisbankaccount2) - Update a bankaccount

### [Benefit](docs/sdks/benefit/README.md)

* [CreateHrisBenefit2](docs/sdks/benefit/README.md#createhrisbenefit2) - Create a benefit
* [GetHrisBenefit2](docs/sdks/benefit/README.md#gethrisbenefit2) - Retrieve a benefit
* [ListHrisBenefits2](docs/sdks/benefit/README.md#listhrisbenefits2) - List all benefits
* [PatchHrisBenefit2](docs/sdks/benefit/README.md#patchhrisbenefit2) - Update a benefit
* [RemoveHrisBenefit2](docs/sdks/benefit/README.md#removehrisbenefit2) - Remove a benefit
* [UpdateHrisBenefit2](docs/sdks/benefit/README.md#updatehrisbenefit2) - Update a benefit

### [Bill](docs/sdks/bill/README.md)

* [CreateAccountingBill2](docs/sdks/bill/README.md#createaccountingbill2) - Create a bill
* [GetAccountingBill2](docs/sdks/bill/README.md#getaccountingbill2) - Retrieve a bill
* [ListAccountingBills2](docs/sdks/bill/README.md#listaccountingbills2) - List all bills
* [PatchAccountingBill2](docs/sdks/bill/README.md#patchaccountingbill2) - Update a bill
* [RemoveAccountingBill2](docs/sdks/bill/README.md#removeaccountingbill2) - Remove a bill
* [UpdateAccountingBill2](docs/sdks/bill/README.md#updateaccountingbill2) - Update a bill

### [Branch](docs/sdks/branch/README.md)

* [CreateRepoBranch2](docs/sdks/branch/README.md#createrepobranch2) - Create a branch
* [GetRepoBranch2](docs/sdks/branch/README.md#getrepobranch2) - Retrieve a branch
* [ListRepoBranches2](docs/sdks/branch/README.md#listrepobranches2) - List all branches
* [PatchRepoBranch2](docs/sdks/branch/README.md#patchrepobranch2) - Update a branch
* [RemoveRepoBranch2](docs/sdks/branch/README.md#removerepobranch2) - Remove a branch
* [UpdateRepoBranch2](docs/sdks/branch/README.md#updaterepobranch2) - Update a branch

### [Busy](docs/sdks/busy/README.md)

* [ListCalendarBusies2](docs/sdks/busy/README.md#listcalendarbusies2) - List all busies

### [Calendar](docs/sdks/calendar/README.md)

* [CreateCalendarCalendar2](docs/sdks/calendar/README.md#createcalendarcalendar2) - Create a calendar
* [CreateCalendarEvent2](docs/sdks/calendar/README.md#createcalendarevent2) - Create an event
* [CreateCalendarLink2](docs/sdks/calendar/README.md#createcalendarlink2) - Create a link
* [CreateCalendarWebinar2](docs/sdks/calendar/README.md#createcalendarwebinar2) - Create a webinar
* [GetCalendarCalendar2](docs/sdks/calendar/README.md#getcalendarcalendar2) - Retrieve a calendar
* [GetCalendarEvent2](docs/sdks/calendar/README.md#getcalendarevent2) - Retrieve an event
* [GetCalendarLink2](docs/sdks/calendar/README.md#getcalendarlink2) - Retrieve a link
* [GetCalendarRecording2](docs/sdks/calendar/README.md#getcalendarrecording2) - Retrieve a recording
* [GetCalendarWebinar2](docs/sdks/calendar/README.md#getcalendarwebinar2) - Retrieve a webinar
* [ListCalendarBusies2](docs/sdks/calendar/README.md#listcalendarbusies2) - List all busies
* [ListCalendarCalendars2](docs/sdks/calendar/README.md#listcalendarcalendars2) - List all calendars
* [ListCalendarEvents2](docs/sdks/calendar/README.md#listcalendarevents2) - List all events
* [ListCalendarLinks2](docs/sdks/calendar/README.md#listcalendarlinks2) - List all links
* [ListCalendarRecordings2](docs/sdks/calendar/README.md#listcalendarrecordings2) - List all recordings
* [ListCalendarWebinars2](docs/sdks/calendar/README.md#listcalendarwebinars2) - List all webinars
* [PatchCalendarCalendar2](docs/sdks/calendar/README.md#patchcalendarcalendar2) - Update a calendar
* [PatchCalendarEvent2](docs/sdks/calendar/README.md#patchcalendarevent2) - Update an event
* [PatchCalendarLink2](docs/sdks/calendar/README.md#patchcalendarlink2) - Update a link
* [PatchCalendarWebinar2](docs/sdks/calendar/README.md#patchcalendarwebinar2) - Update a webinar
* [RemoveCalendarCalendar2](docs/sdks/calendar/README.md#removecalendarcalendar2) - Remove a calendar
* [RemoveCalendarEvent2](docs/sdks/calendar/README.md#removecalendarevent2) - Remove an event
* [RemoveCalendarLink2](docs/sdks/calendar/README.md#removecalendarlink2) - Remove a link
* [RemoveCalendarWebinar2](docs/sdks/calendar/README.md#removecalendarwebinar2) - Remove a webinar
* [UpdateCalendarCalendar2](docs/sdks/calendar/README.md#updatecalendarcalendar2) - Update a calendar
* [UpdateCalendarEvent2](docs/sdks/calendar/README.md#updatecalendarevent2) - Update an event
* [UpdateCalendarLink2](docs/sdks/calendar/README.md#updatecalendarlink2) - Update a link
* [UpdateCalendarWebinar2](docs/sdks/calendar/README.md#updatecalendarwebinar2) - Update a webinar

### [Call](docs/sdks/call/README.md)

* [GetUcCall2](docs/sdks/call/README.md#getuccall2) - Retrieve a call
* [ListUcCalls2](docs/sdks/call/README.md#listuccalls2) - List all calls

### [Campaign](docs/sdks/campaign/README.md)

* [CreateAdsCampaign2](docs/sdks/campaign/README.md#createadscampaign2) - Create a campaign
* [CreateMartechCampaign2](docs/sdks/campaign/README.md#createmartechcampaign2) - Create a campaign
* [GetAdsCampaign2](docs/sdks/campaign/README.md#getadscampaign2) - Retrieve a campaign
* [GetMartechCampaign2](docs/sdks/campaign/README.md#getmartechcampaign2) - Retrieve a campaign
* [ListAdsCampaigns2](docs/sdks/campaign/README.md#listadscampaigns2) - List all campaigns
* [ListMartechCampaigns2](docs/sdks/campaign/README.md#listmartechcampaigns2) - List all campaigns
* [PatchAdsCampaign2](docs/sdks/campaign/README.md#patchadscampaign2) - Update a campaign
* [PatchMartechCampaign2](docs/sdks/campaign/README.md#patchmartechcampaign2) - Update a campaign
* [RemoveAdsCampaign2](docs/sdks/campaign/README.md#removeadscampaign2) - Remove a campaign
* [RemoveMartechCampaign2](docs/sdks/campaign/README.md#removemartechcampaign2) - Remove a campaign
* [UpdateAdsCampaign2](docs/sdks/campaign/README.md#updateadscampaign2) - Update a campaign
* [UpdateMartechCampaign2](docs/sdks/campaign/README.md#updatemartechcampaign2) - Update a campaign

### [Candidate](docs/sdks/candidate/README.md)

* [CreateAtsCandidate2](docs/sdks/candidate/README.md#createatscandidate2) - Create a candidate
* [GetAtsCandidate2](docs/sdks/candidate/README.md#getatscandidate2) - Retrieve a candidate
* [ListAtsCandidates2](docs/sdks/candidate/README.md#listatscandidates2) - List all candidates
* [PatchAtsCandidate2](docs/sdks/candidate/README.md#patchatscandidate2) - Update a candidate
* [RemoveAtsCandidate2](docs/sdks/candidate/README.md#removeatscandidate2) - Remove a candidate
* [UpdateAtsCandidate2](docs/sdks/candidate/README.md#updateatscandidate2) - Update a candidate

### [Carrier](docs/sdks/carrier/README.md)

* [GetShippingCarrier2](docs/sdks/carrier/README.md#getshippingcarrier2) - Retrieve a carrier
* [ListShippingCarriers2](docs/sdks/carrier/README.md#listshippingcarriers2) - List all carriers

### [Cashflow](docs/sdks/cashflow/README.md)

* [GetAccountingCashflow2](docs/sdks/cashflow/README.md#getaccountingcashflow2) - Retrieve a cashflow
* [ListAccountingCashflows2](docs/sdks/cashflow/README.md#listaccountingcashflows2) - List all cashflows

### [Category](docs/sdks/category/README.md)

* [CreateAccountingCategory2](docs/sdks/category/README.md#createaccountingcategory2) - Create a category
* [CreateTicketingCategory2](docs/sdks/category/README.md#createticketingcategory2) - Create a category
* [GetAccountingCategory2](docs/sdks/category/README.md#getaccountingcategory2) - Retrieve a category
* [GetTicketingCategory2](docs/sdks/category/README.md#getticketingcategory2) - Retrieve a category
* [ListAccountingCategories2](docs/sdks/category/README.md#listaccountingcategories2) - List all categories
* [ListTicketingCategories2](docs/sdks/category/README.md#listticketingcategories2) - List all categories
* [PatchAccountingCategory2](docs/sdks/category/README.md#patchaccountingcategory2) - Update a category
* [PatchTicketingCategory2](docs/sdks/category/README.md#patchticketingcategory2) - Update a category
* [RemoveAccountingCategory2](docs/sdks/category/README.md#removeaccountingcategory2) - Remove a category
* [RemoveTicketingCategory2](docs/sdks/category/README.md#removeticketingcategory2) - Remove a category
* [UpdateAccountingCategory2](docs/sdks/category/README.md#updateaccountingcategory2) - Update a category
* [UpdateTicketingCategory2](docs/sdks/category/README.md#updateticketingcategory2) - Update a category

### [Change](docs/sdks/change/README.md)

* [GetTaskChange2](docs/sdks/change/README.md#gettaskchange2) - Retrieve a change
* [ListTaskChanges2](docs/sdks/change/README.md#listtaskchanges2) - List all changes

### [Channel](docs/sdks/channel/README.md)

* [GetMessagingChannel2](docs/sdks/channel/README.md#getmessagingchannel2) - Retrieve a channel
* [ListMessagingChannels2](docs/sdks/channel/README.md#listmessagingchannels2) - List all channels

### [Class](docs/sdks/class/README.md)

* [CreateLmsClass2](docs/sdks/class/README.md#createlmsclass2) - Create a class
* [GetLmsClass2](docs/sdks/class/README.md#getlmsclass2) - Retrieve a class
* [ListLmsClasses2](docs/sdks/class/README.md#listlmsclasses2) - List all classes
* [PatchLmsClass2](docs/sdks/class/README.md#patchlmsclass2) - Update a class
* [RemoveLmsClass2](docs/sdks/class/README.md#removelmsclass2) - Remove a class
* [UpdateLmsClass2](docs/sdks/class/README.md#updatelmsclass2) - Update a class

### [Clubs](docs/sdks/clubs/README.md)

* [GetClubsActivity2](docs/sdks/clubs/README.md#getclubsactivity2) - Retrieve an activity
* [GetClubsEvent2](docs/sdks/clubs/README.md#getclubsevent2) - Retrieve an event
* [GetClubsGroup2](docs/sdks/clubs/README.md#getclubsgroup2) - Retrieve a group
* [GetClubsLocation2](docs/sdks/clubs/README.md#getclubslocation2) - Retrieve a location
* [GetClubsMember2](docs/sdks/clubs/README.md#getclubsmember2) - Retrieve a member
* [ListClubsActivities2](docs/sdks/clubs/README.md#listclubsactivities2) - List all activities
* [ListClubsEvents2](docs/sdks/clubs/README.md#listclubsevents2) - List all events
* [ListClubsGroups2](docs/sdks/clubs/README.md#listclubsgroups2) - List all groups
* [ListClubsLocations2](docs/sdks/clubs/README.md#listclubslocations2) - List all locations
* [ListClubsMembers2](docs/sdks/clubs/README.md#listclubsmembers2) - List all members

### [Collection](docs/sdks/collection/README.md)

* [CreateCommerceCollection2](docs/sdks/collection/README.md#createcommercecollection2) - Create a collection
* [CreateLmsCollection2](docs/sdks/collection/README.md#createlmscollection2) - Create a collection
* [GetCommerceCollection2](docs/sdks/collection/README.md#getcommercecollection2) - Retrieve a collection
* [GetLmsCollection2](docs/sdks/collection/README.md#getlmscollection2) - Retrieve a collection
* [ListCommerceCollections2](docs/sdks/collection/README.md#listcommercecollections2) - List all collections
* [ListLmsCollections2](docs/sdks/collection/README.md#listlmscollections2) - List all collections
* [PatchCommerceCollection2](docs/sdks/collection/README.md#patchcommercecollection2) - Update a collection
* [PatchLmsCollection2](docs/sdks/collection/README.md#patchlmscollection2) - Update a collection
* [RemoveCommerceCollection2](docs/sdks/collection/README.md#removecommercecollection2) - Remove a collection
* [RemoveLmsCollection2](docs/sdks/collection/README.md#removelmscollection2) - Remove a collection
* [UpdateCommerceCollection2](docs/sdks/collection/README.md#updatecommercecollection2) - Update a collection
* [UpdateLmsCollection2](docs/sdks/collection/README.md#updatelmscollection2) - Update a collection

### [Comment](docs/sdks/comment/README.md)

* [CreateKmsComment2](docs/sdks/comment/README.md#createkmscomment2) - Create a comment
* [CreateTaskComment2](docs/sdks/comment/README.md#createtaskcomment2) - Create a comment
* [CreateUcComment2](docs/sdks/comment/README.md#createuccomment2) - Create a comment
* [GetKmsComment2](docs/sdks/comment/README.md#getkmscomment2) - Retrieve a comment
* [GetTaskComment2](docs/sdks/comment/README.md#gettaskcomment2) - Retrieve a comment
* [GetUcComment2](docs/sdks/comment/README.md#getuccomment2) - Retrieve a comment
* [ListKmsComments2](docs/sdks/comment/README.md#listkmscomments2) - List all comments
* [ListTaskComments2](docs/sdks/comment/README.md#listtaskcomments2) - List all comments
* [ListUcComments2](docs/sdks/comment/README.md#listuccomments2) - List all comments
* [PatchKmsComment2](docs/sdks/comment/README.md#patchkmscomment2) - Update a comment
* [PatchTaskComment2](docs/sdks/comment/README.md#patchtaskcomment2) - Update a comment
* [PatchUcComment2](docs/sdks/comment/README.md#patchuccomment2) - Update a comment
* [RemoveKmsComment2](docs/sdks/comment/README.md#removekmscomment2) - Remove a comment
* [RemoveTaskComment2](docs/sdks/comment/README.md#removetaskcomment2) - Remove a comment
* [RemoveUcComment2](docs/sdks/comment/README.md#removeuccomment2) - Remove a comment
* [UpdateKmsComment2](docs/sdks/comment/README.md#updatekmscomment2) - Update a comment
* [UpdateTaskComment2](docs/sdks/comment/README.md#updatetaskcomment2) - Update a comment
* [UpdateUcComment2](docs/sdks/comment/README.md#updateuccomment2) - Update a comment

### [Commerce](docs/sdks/commerce/README.md)

* [CreateCommerceCollection2](docs/sdks/commerce/README.md#createcommercecollection2) - Create a collection
* [CreateCommerceInventory2](docs/sdks/commerce/README.md#createcommerceinventory2) - Create an inventory
* [CreateCommerceItem2](docs/sdks/commerce/README.md#createcommerceitem2) - Create an item
* [CreateCommerceItemvariant2](docs/sdks/commerce/README.md#createcommerceitemvariant2) - Create an itemvariant
* [CreateCommerceLocation2](docs/sdks/commerce/README.md#createcommercelocation2) - Create a location
* [CreateCommerceReservation2](docs/sdks/commerce/README.md#createcommercereservation2) - Create a reservation
* [CreateCommerceReview2](docs/sdks/commerce/README.md#createcommercereview2) - Create a review
* [CreateCommerceSaleschannel2](docs/sdks/commerce/README.md#createcommercesaleschannel2) - Create a saleschannel
* [GetCommerceCollection2](docs/sdks/commerce/README.md#getcommercecollection2) - Retrieve a collection
* [GetCommerceInventory2](docs/sdks/commerce/README.md#getcommerceinventory2) - Retrieve an inventory
* [GetCommerceItem2](docs/sdks/commerce/README.md#getcommerceitem2) - Retrieve an item
* [GetCommerceItemvariant2](docs/sdks/commerce/README.md#getcommerceitemvariant2) - Retrieve an itemvariant
* [GetCommerceLocation2](docs/sdks/commerce/README.md#getcommercelocation2) - Retrieve a location
* [GetCommerceReservation2](docs/sdks/commerce/README.md#getcommercereservation2) - Retrieve a reservation
* [GetCommerceReview2](docs/sdks/commerce/README.md#getcommercereview2) - Retrieve a review
* [GetCommerceSaleschannel2](docs/sdks/commerce/README.md#getcommercesaleschannel2) - Retrieve a saleschannel
* [ListCommerceAvailabilities2](docs/sdks/commerce/README.md#listcommerceavailabilities2) - List all availabilities
* [ListCommerceCollections2](docs/sdks/commerce/README.md#listcommercecollections2) - List all collections
* [ListCommerceInventories2](docs/sdks/commerce/README.md#listcommerceinventories2) - List all inventories
* [ListCommerceItems2](docs/sdks/commerce/README.md#listcommerceitems2) - List all items
* [ListCommerceItemvariants2](docs/sdks/commerce/README.md#listcommerceitemvariants2) - List all itemvariants
* [ListCommerceLocations2](docs/sdks/commerce/README.md#listcommercelocations2) - List all locations
* [ListCommerceReservations2](docs/sdks/commerce/README.md#listcommercereservations2) - List all reservations
* [ListCommerceReviews2](docs/sdks/commerce/README.md#listcommercereviews2) - List all reviews
* [ListCommerceSaleschannels2](docs/sdks/commerce/README.md#listcommercesaleschannels2) - List all saleschannels
* [PatchCommerceCollection2](docs/sdks/commerce/README.md#patchcommercecollection2) - Update a collection
* [PatchCommerceInventory2](docs/sdks/commerce/README.md#patchcommerceinventory2) - Update an inventory
* [PatchCommerceItem2](docs/sdks/commerce/README.md#patchcommerceitem2) - Update an item
* [PatchCommerceItemvariant2](docs/sdks/commerce/README.md#patchcommerceitemvariant2) - Update an itemvariant
* [PatchCommerceLocation2](docs/sdks/commerce/README.md#patchcommercelocation2) - Update a location
* [PatchCommerceReservation2](docs/sdks/commerce/README.md#patchcommercereservation2) - Update a reservation
* [PatchCommerceReview2](docs/sdks/commerce/README.md#patchcommercereview2) - Update a review
* [PatchCommerceSaleschannel2](docs/sdks/commerce/README.md#patchcommercesaleschannel2) - Update a saleschannel
* [RemoveCommerceCollection2](docs/sdks/commerce/README.md#removecommercecollection2) - Remove a collection
* [RemoveCommerceInventory2](docs/sdks/commerce/README.md#removecommerceinventory2) - Remove an inventory
* [RemoveCommerceItem2](docs/sdks/commerce/README.md#removecommerceitem2) - Remove an item
* [RemoveCommerceItemvariant2](docs/sdks/commerce/README.md#removecommerceitemvariant2) - Remove an itemvariant
* [RemoveCommerceLocation2](docs/sdks/commerce/README.md#removecommercelocation2) - Remove a location
* [RemoveCommerceReservation2](docs/sdks/commerce/README.md#removecommercereservation2) - Remove a reservation
* [RemoveCommerceReview2](docs/sdks/commerce/README.md#removecommercereview2) - Remove a review
* [RemoveCommerceSaleschannel2](docs/sdks/commerce/README.md#removecommercesaleschannel2) - Remove a saleschannel
* [UpdateCommerceCollection2](docs/sdks/commerce/README.md#updatecommercecollection2) - Update a collection
* [UpdateCommerceInventory2](docs/sdks/commerce/README.md#updatecommerceinventory2) - Update an inventory
* [UpdateCommerceItem2](docs/sdks/commerce/README.md#updatecommerceitem2) - Update an item
* [UpdateCommerceItemvariant2](docs/sdks/commerce/README.md#updatecommerceitemvariant2) - Update an itemvariant
* [UpdateCommerceLocation2](docs/sdks/commerce/README.md#updatecommercelocation2) - Update a location
* [UpdateCommerceReservation2](docs/sdks/commerce/README.md#updatecommercereservation2) - Update a reservation
* [UpdateCommerceReview2](docs/sdks/commerce/README.md#updatecommercereview2) - Update a review
* [UpdateCommerceSaleschannel2](docs/sdks/commerce/README.md#updatecommercesaleschannel2) - Update a saleschannel

### [Commit](docs/sdks/commit/README.md)

* [CreateRepoCommit2](docs/sdks/commit/README.md#createrepocommit2) - Create a commit
* [GetRepoCommit2](docs/sdks/commit/README.md#getrepocommit2) - Retrieve a commit
* [ListRepoCommits2](docs/sdks/commit/README.md#listrepocommits2) - List all commits
* [PatchRepoCommit2](docs/sdks/commit/README.md#patchrepocommit2) - Update a commit
* [RemoveRepoCommit2](docs/sdks/commit/README.md#removerepocommit2) - Remove a commit
* [UpdateRepoCommit2](docs/sdks/commit/README.md#updaterepocommit2) - Update a commit

### [Company](docs/sdks/company/README.md)

* [CreateAtsCompany2](docs/sdks/company/README.md#createatscompany2) - Create a company
* [CreateCrmCompany2](docs/sdks/company/README.md#createcrmcompany2) - Create a company
* [CreateHrisCompany2](docs/sdks/company/README.md#createhriscompany2) - Create a company
* [GetAtsCompany2](docs/sdks/company/README.md#getatscompany2) - Retrieve a company
* [GetCrmCompany2](docs/sdks/company/README.md#getcrmcompany2) - Retrieve a company
* [GetHrisCompany2](docs/sdks/company/README.md#gethriscompany2) - Retrieve a company
* [ListAtsCompanies2](docs/sdks/company/README.md#listatscompanies2) - List all companies
* [ListCrmCompanies2](docs/sdks/company/README.md#listcrmcompanies2) - List all companies
* [ListEnrichCompanies2](docs/sdks/company/README.md#listenrichcompanies2) - Retrieve enrichment information for a company
* [ListHrisCompanies2](docs/sdks/company/README.md#listhriscompanies2) - List all companies
* [PatchAtsCompany2](docs/sdks/company/README.md#patchatscompany2) - Update a company
* [PatchCrmCompany2](docs/sdks/company/README.md#patchcrmcompany2) - Update a company
* [PatchHrisCompany2](docs/sdks/company/README.md#patchhriscompany2) - Update a company
* [RemoveAtsCompany2](docs/sdks/company/README.md#removeatscompany2) - Remove a company
* [RemoveCrmCompany2](docs/sdks/company/README.md#removecrmcompany2) - Remove a company
* [RemoveHrisCompany2](docs/sdks/company/README.md#removehriscompany2) - Remove a company
* [UpdateAtsCompany2](docs/sdks/company/README.md#updateatscompany2) - Update a company
* [UpdateCrmCompany2](docs/sdks/company/README.md#updatecrmcompany2) - Update a company
* [UpdateHrisCompany2](docs/sdks/company/README.md#updatehriscompany2) - Update a company

### [Connection](docs/sdks/connection/README.md)

* [CreateUnifiedConnection](docs/sdks/connection/README.md#createunifiedconnection) - Create connection
* [GetUnifiedConnection](docs/sdks/connection/README.md#getunifiedconnection) - Retrieve connection
* [ListUnifiedConnections](docs/sdks/connection/README.md#listunifiedconnections) - List all connections
* [PatchUnifiedConnection](docs/sdks/connection/README.md#patchunifiedconnection) - Update connection
* [RemoveUnifiedConnection](docs/sdks/connection/README.md#removeunifiedconnection) - Remove connection
* [UpdateUnifiedConnection](docs/sdks/connection/README.md#updateunifiedconnection) - Update connection

### [Contact](docs/sdks/contact/README.md)

* [CreateAccountingContact2](docs/sdks/contact/README.md#createaccountingcontact2) - Create a contact
* [CreateCrmContact2](docs/sdks/contact/README.md#createcrmcontact2) - Create a contact
* [CreateUcContact2](docs/sdks/contact/README.md#createuccontact2) - Create a contact
* [GetAccountingContact2](docs/sdks/contact/README.md#getaccountingcontact2) - Retrieve a contact
* [GetCrmContact2](docs/sdks/contact/README.md#getcrmcontact2) - Retrieve a contact
* [GetUcContact2](docs/sdks/contact/README.md#getuccontact2) - Retrieve a contact
* [ListAccountingContacts2](docs/sdks/contact/README.md#listaccountingcontacts2) - List all contacts
* [ListCrmContacts2](docs/sdks/contact/README.md#listcrmcontacts2) - List all contacts
* [ListUcContacts2](docs/sdks/contact/README.md#listuccontacts2) - List all contacts
* [PatchAccountingContact2](docs/sdks/contact/README.md#patchaccountingcontact2) - Update a contact
* [PatchCrmContact2](docs/sdks/contact/README.md#patchcrmcontact2) - Update a contact
* [PatchUcContact2](docs/sdks/contact/README.md#patchuccontact2) - Update a contact
* [RemoveAccountingContact2](docs/sdks/contact/README.md#removeaccountingcontact2) - Remove a contact
* [RemoveCrmContact2](docs/sdks/contact/README.md#removecrmcontact2) - Remove a contact
* [RemoveUcContact2](docs/sdks/contact/README.md#removeuccontact2) - Remove a contact
* [UpdateAccountingContact2](docs/sdks/contact/README.md#updateaccountingcontact2) - Update a contact
* [UpdateCrmContact2](docs/sdks/contact/README.md#updatecrmcontact2) - Update a contact
* [UpdateUcContact2](docs/sdks/contact/README.md#updateuccontact2) - Update a contact

### [Content](docs/sdks/content/README.md)

* [CreateLmsContent2](docs/sdks/content/README.md#createlmscontent2) - Create a content
* [GetLmsContent2](docs/sdks/content/README.md#getlmscontent2) - Retrieve a content
* [ListLmsContents2](docs/sdks/content/README.md#listlmscontents2) - List all contents
* [PatchLmsContent2](docs/sdks/content/README.md#patchlmscontent2) - Update a content
* [RemoveLmsContent2](docs/sdks/content/README.md#removelmscontent2) - Remove a content
* [UpdateLmsContent2](docs/sdks/content/README.md#updatelmscontent2) - Update a content

### [Course](docs/sdks/course/README.md)

* [CreateLmsCourse2](docs/sdks/course/README.md#createlmscourse2) - Create a course
* [GetLmsCourse2](docs/sdks/course/README.md#getlmscourse2) - Retrieve a course
* [ListLmsCourses2](docs/sdks/course/README.md#listlmscourses2) - List all courses
* [PatchLmsCourse2](docs/sdks/course/README.md#patchlmscourse2) - Update a course
* [RemoveLmsCourse2](docs/sdks/course/README.md#removelmscourse2) - Remove a course
* [UpdateLmsCourse2](docs/sdks/course/README.md#updatelmscourse2) - Update a course

### [Creative](docs/sdks/creative/README.md)

* [CreateAdsCreative2](docs/sdks/creative/README.md#createadscreative2) - Create a creative
* [GetAdsCreative2](docs/sdks/creative/README.md#getadscreative2) - Retrieve a creative
* [ListAdsCreatives2](docs/sdks/creative/README.md#listadscreatives2) - List all creatives
* [PatchAdsCreative2](docs/sdks/creative/README.md#patchadscreative2) - Update a creative
* [RemoveAdsCreative2](docs/sdks/creative/README.md#removeadscreative2) - Remove a creative
* [UpdateAdsCreative2](docs/sdks/creative/README.md#updateadscreative2) - Update a creative

### [Creditmemo](docs/sdks/creditmemo/README.md)

* [CreateAccountingCreditmemo2](docs/sdks/creditmemo/README.md#createaccountingcreditmemo2) - Create a creditmemo
* [GetAccountingCreditmemo2](docs/sdks/creditmemo/README.md#getaccountingcreditmemo2) - Retrieve a creditmemo
* [ListAccountingCreditmemoes2](docs/sdks/creditmemo/README.md#listaccountingcreditmemoes2) - List all creditmemoes
* [PatchAccountingCreditmemo2](docs/sdks/creditmemo/README.md#patchaccountingcreditmemo2) - Update a creditmemo
* [RemoveAccountingCreditmemo2](docs/sdks/creditmemo/README.md#removeaccountingcreditmemo2) - Remove a creditmemo
* [UpdateAccountingCreditmemo2](docs/sdks/creditmemo/README.md#updateaccountingcreditmemo2) - Update a creditmemo

### [Crm](docs/sdks/crm/README.md)

* [CreateCrmCompany2](docs/sdks/crm/README.md#createcrmcompany2) - Create a company
* [CreateCrmContact2](docs/sdks/crm/README.md#createcrmcontact2) - Create a contact
* [CreateCrmDeal2](docs/sdks/crm/README.md#createcrmdeal2) - Create a deal
* [CreateCrmEvent2](docs/sdks/crm/README.md#createcrmevent2) - Create an event
* [CreateCrmLead2](docs/sdks/crm/README.md#createcrmlead2) - Create a lead
* [CreateCrmPipeline2](docs/sdks/crm/README.md#createcrmpipeline2) - Create a pipeline
* [GetCrmCompany2](docs/sdks/crm/README.md#getcrmcompany2) - Retrieve a company
* [GetCrmContact2](docs/sdks/crm/README.md#getcrmcontact2) - Retrieve a contact
* [GetCrmDeal2](docs/sdks/crm/README.md#getcrmdeal2) - Retrieve a deal
* [GetCrmEvent2](docs/sdks/crm/README.md#getcrmevent2) - Retrieve an event
* [GetCrmLead2](docs/sdks/crm/README.md#getcrmlead2) - Retrieve a lead
* [GetCrmPipeline2](docs/sdks/crm/README.md#getcrmpipeline2) - Retrieve a pipeline
* [ListCrmCompanies2](docs/sdks/crm/README.md#listcrmcompanies2) - List all companies
* [ListCrmContacts2](docs/sdks/crm/README.md#listcrmcontacts2) - List all contacts
* [ListCrmDeals2](docs/sdks/crm/README.md#listcrmdeals2) - List all deals
* [ListCrmEvents2](docs/sdks/crm/README.md#listcrmevents2) - List all events
* [ListCrmLeads2](docs/sdks/crm/README.md#listcrmleads2) - List all leads
* [ListCrmPipelines2](docs/sdks/crm/README.md#listcrmpipelines2) - List all pipelines
* [PatchCrmCompany2](docs/sdks/crm/README.md#patchcrmcompany2) - Update a company
* [PatchCrmContact2](docs/sdks/crm/README.md#patchcrmcontact2) - Update a contact
* [PatchCrmDeal2](docs/sdks/crm/README.md#patchcrmdeal2) - Update a deal
* [PatchCrmEvent2](docs/sdks/crm/README.md#patchcrmevent2) - Update an event
* [PatchCrmLead2](docs/sdks/crm/README.md#patchcrmlead2) - Update a lead
* [PatchCrmPipeline2](docs/sdks/crm/README.md#patchcrmpipeline2) - Update a pipeline
* [RemoveCrmCompany2](docs/sdks/crm/README.md#removecrmcompany2) - Remove a company
* [RemoveCrmContact2](docs/sdks/crm/README.md#removecrmcontact2) - Remove a contact
* [RemoveCrmDeal2](docs/sdks/crm/README.md#removecrmdeal2) - Remove a deal
* [RemoveCrmEvent2](docs/sdks/crm/README.md#removecrmevent2) - Remove an event
* [RemoveCrmLead2](docs/sdks/crm/README.md#removecrmlead2) - Remove a lead
* [RemoveCrmPipeline2](docs/sdks/crm/README.md#removecrmpipeline2) - Remove a pipeline
* [UpdateCrmCompany2](docs/sdks/crm/README.md#updatecrmcompany2) - Update a company
* [UpdateCrmContact2](docs/sdks/crm/README.md#updatecrmcontact2) - Update a contact
* [UpdateCrmDeal2](docs/sdks/crm/README.md#updatecrmdeal2) - Update a deal
* [UpdateCrmEvent2](docs/sdks/crm/README.md#updatecrmevent2) - Update an event
* [UpdateCrmLead2](docs/sdks/crm/README.md#updatecrmlead2) - Update a lead
* [UpdateCrmPipeline2](docs/sdks/crm/README.md#updatecrmpipeline2) - Update a pipeline

### [Customer](docs/sdks/customer/README.md)

* [CreateTicketingCustomer2](docs/sdks/customer/README.md#createticketingcustomer2) - Create a customer
* [GetTicketingCustomer2](docs/sdks/customer/README.md#getticketingcustomer2) - Retrieve a customer
* [ListTicketingCustomers2](docs/sdks/customer/README.md#listticketingcustomers2) - List all customers
* [PatchTicketingCustomer2](docs/sdks/customer/README.md#patchticketingcustomer2) - Update a customer
* [RemoveTicketingCustomer2](docs/sdks/customer/README.md#removeticketingcustomer2) - Remove a customer
* [UpdateTicketingCustomer2](docs/sdks/customer/README.md#updateticketingcustomer2) - Update a customer

### [Database](docs/sdks/database/README.md)

* [CreateDatastoreDatabase2](docs/sdks/database/README.md#createdatastoredatabase2) - Create a database
* [GetDatastoreDatabase2](docs/sdks/database/README.md#getdatastoredatabase2) - Retrieve a database
* [ListDatastoreDatabases2](docs/sdks/database/README.md#listdatastoredatabases2) - List all databases
* [PatchDatastoreDatabase2](docs/sdks/database/README.md#patchdatastoredatabase2) - Update a database
* [RemoveDatastoreDatabase2](docs/sdks/database/README.md#removedatastoredatabase2) - Remove a database
* [UpdateDatastoreDatabase2](docs/sdks/database/README.md#updatedatastoredatabase2) - Update a database

### [Datastore](docs/sdks/datastore/README.md)

* [CreateDatastoreDatabase2](docs/sdks/datastore/README.md#createdatastoredatabase2) - Create a database
* [CreateDatastoreQuery2](docs/sdks/datastore/README.md#createdatastorequery2) - Create a query
* [CreateDatastoreRecord2](docs/sdks/datastore/README.md#createdatastorerecord2) - Create a record
* [CreateDatastoreTable2](docs/sdks/datastore/README.md#createdatastoretable2) - Create a table
* [GetDatastoreDatabase2](docs/sdks/datastore/README.md#getdatastoredatabase2) - Retrieve a database
* [GetDatastoreRecord2](docs/sdks/datastore/README.md#getdatastorerecord2) - Retrieve a record
* [GetDatastoreTable2](docs/sdks/datastore/README.md#getdatastoretable2) - Retrieve a table
* [ListDatastoreDatabases2](docs/sdks/datastore/README.md#listdatastoredatabases2) - List all databases
* [ListDatastoreRecords2](docs/sdks/datastore/README.md#listdatastorerecords2) - List all records
* [ListDatastoreTables2](docs/sdks/datastore/README.md#listdatastoretables2) - List all tables
* [PatchDatastoreDatabase2](docs/sdks/datastore/README.md#patchdatastoredatabase2) - Update a database
* [PatchDatastoreRecord2](docs/sdks/datastore/README.md#patchdatastorerecord2) - Update a record
* [PatchDatastoreTable2](docs/sdks/datastore/README.md#patchdatastoretable2) - Update a table
* [RemoveDatastoreDatabase2](docs/sdks/datastore/README.md#removedatastoredatabase2) - Remove a database
* [RemoveDatastoreRecord2](docs/sdks/datastore/README.md#removedatastorerecord2) - Remove a record
* [RemoveDatastoreTable2](docs/sdks/datastore/README.md#removedatastoretable2) - Remove a table
* [UpdateDatastoreDatabase2](docs/sdks/datastore/README.md#updatedatastoredatabase2) - Update a database
* [UpdateDatastoreRecord2](docs/sdks/datastore/README.md#updatedatastorerecord2) - Update a record
* [UpdateDatastoreTable2](docs/sdks/datastore/README.md#updatedatastoretable2) - Update a table

### [Deal](docs/sdks/deal/README.md)

* [CreateCrmDeal2](docs/sdks/deal/README.md#createcrmdeal2) - Create a deal
* [GetCrmDeal2](docs/sdks/deal/README.md#getcrmdeal2) - Retrieve a deal
* [ListCrmDeals2](docs/sdks/deal/README.md#listcrmdeals2) - List all deals
* [PatchCrmDeal2](docs/sdks/deal/README.md#patchcrmdeal2) - Update a deal
* [RemoveCrmDeal2](docs/sdks/deal/README.md#removecrmdeal2) - Remove a deal
* [UpdateCrmDeal2](docs/sdks/deal/README.md#updatecrmdeal2) - Update a deal

### [Deduction](docs/sdks/deduction/README.md)

* [CreateHrisDeduction2](docs/sdks/deduction/README.md#createhrisdeduction2) - Create a deduction
* [GetHrisDeduction2](docs/sdks/deduction/README.md#gethrisdeduction2) - Retrieve a deduction
* [ListHrisDeductions2](docs/sdks/deduction/README.md#listhrisdeductions2) - List all deductions
* [PatchHrisDeduction2](docs/sdks/deduction/README.md#patchhrisdeduction2) - Update a deduction
* [RemoveHrisDeduction2](docs/sdks/deduction/README.md#removehrisdeduction2) - Remove a deduction
* [UpdateHrisDeduction2](docs/sdks/deduction/README.md#updatehrisdeduction2) - Update a deduction

### [Device](docs/sdks/device/README.md)

* [CreateHrisDevice2](docs/sdks/device/README.md#createhrisdevice2) - Create a device
* [GetHrisDevice2](docs/sdks/device/README.md#gethrisdevice2) - Retrieve a device
* [ListHrisDevices2](docs/sdks/device/README.md#listhrisdevices2) - List all devices
* [PatchHrisDevice2](docs/sdks/device/README.md#patchhrisdevice2) - Update a device
* [RemoveHrisDevice2](docs/sdks/device/README.md#removehrisdevice2) - Remove a device
* [UpdateHrisDevice2](docs/sdks/device/README.md#updatehrisdevice2) - Update a device

### [Document](docs/sdks/document/README.md)

* [CreateAtsDocument2](docs/sdks/document/README.md#createatsdocument2) - Create a document
* [CreateSigningDocument2](docs/sdks/document/README.md#createsigningdocument2) - Create a document
* [GetAtsDocument2](docs/sdks/document/README.md#getatsdocument2) - Retrieve a document
* [GetSigningDocument2](docs/sdks/document/README.md#getsigningdocument2) - Retrieve a document
* [ListAtsDocuments2](docs/sdks/document/README.md#listatsdocuments2) - List all documents
* [ListSigningDocuments2](docs/sdks/document/README.md#listsigningdocuments2) - List all documents
* [PatchAtsDocument2](docs/sdks/document/README.md#patchatsdocument2) - Update a document
* [PatchSigningDocument2](docs/sdks/document/README.md#patchsigningdocument2) - Update a document
* [RemoveAtsDocument2](docs/sdks/document/README.md#removeatsdocument2) - Remove a document
* [RemoveSigningDocument2](docs/sdks/document/README.md#removesigningdocument2) - Remove a document
* [UpdateAtsDocument2](docs/sdks/document/README.md#updateatsdocument2) - Update a document
* [UpdateSigningDocument2](docs/sdks/document/README.md#updatesigningdocument2) - Update a document

### [Embedding](docs/sdks/embedding/README.md)

* [CreateGenaiEmbedding2](docs/sdks/embedding/README.md#creategenaiembedding2) - Create an embedding

### [Employee](docs/sdks/employee/README.md)

* [CreateHrisEmployee2](docs/sdks/employee/README.md#createhrisemployee2) - Create an employee
* [GetHrisEmployee2](docs/sdks/employee/README.md#gethrisemployee2) - Retrieve an employee
* [ListHrisEmployees2](docs/sdks/employee/README.md#listhrisemployees2) - List all employees
* [PatchHrisEmployee2](docs/sdks/employee/README.md#patchhrisemployee2) - Update an employee
* [RemoveHrisEmployee2](docs/sdks/employee/README.md#removehrisemployee2) - Remove an employee
* [UpdateHrisEmployee2](docs/sdks/employee/README.md#updatehrisemployee2) - Update an employee

### [Enrich](docs/sdks/enrich/README.md)

* [ListEnrichCompanies2](docs/sdks/enrich/README.md#listenrichcompanies2) - Retrieve enrichment information for a company
* [ListEnrichPeople2](docs/sdks/enrich/README.md#listenrichpeople2) - Retrieve enrichment information for a person

### [Environment](docs/sdks/environment/README.md)

* [CreateUnifiedEnvironment](docs/sdks/environment/README.md#createunifiedenvironment) - Create new environments
* [ListUnifiedEnvironments](docs/sdks/environment/README.md#listunifiedenvironments) - Returns all environments
* [RemoveUnifiedEnvironment](docs/sdks/environment/README.md#removeunifiedenvironment) - Remove an environment

### [Event](docs/sdks/event/README.md)

* [CreateAnalyticsEvent2](docs/sdks/event/README.md#createanalyticsevent2) - Create an event
* [CreateCalendarEvent2](docs/sdks/event/README.md#createcalendarevent2) - Create an event
* [CreateCrmEvent2](docs/sdks/event/README.md#createcrmevent2) - Create an event
* [GetAnalyticsEvent2](docs/sdks/event/README.md#getanalyticsevent2) - Retrieve an event
* [GetCalendarEvent2](docs/sdks/event/README.md#getcalendarevent2) - Retrieve an event
* [GetClubsEvent2](docs/sdks/event/README.md#getclubsevent2) - Retrieve an event
* [GetCrmEvent2](docs/sdks/event/README.md#getcrmevent2) - Retrieve an event
* [ListAnalyticsEvents2](docs/sdks/event/README.md#listanalyticsevents2) - List all events
* [ListCalendarEvents2](docs/sdks/event/README.md#listcalendarevents2) - List all events
* [ListClubsEvents2](docs/sdks/event/README.md#listclubsevents2) - List all events
* [ListCrmEvents2](docs/sdks/event/README.md#listcrmevents2) - List all events
* [PatchCalendarEvent2](docs/sdks/event/README.md#patchcalendarevent2) - Update an event
* [PatchCrmEvent2](docs/sdks/event/README.md#patchcrmevent2) - Update an event
* [PatchMessagingEvent2](docs/sdks/event/README.md#patchmessagingevent2) - Update an event
* [RemoveCalendarEvent2](docs/sdks/event/README.md#removecalendarevent2) - Remove an event
* [RemoveCrmEvent2](docs/sdks/event/README.md#removecrmevent2) - Remove an event
* [UpdateCalendarEvent2](docs/sdks/event/README.md#updatecalendarevent2) - Update an event
* [UpdateCrmEvent2](docs/sdks/event/README.md#updatecrmevent2) - Update an event
* [UpdateMessagingEvent2](docs/sdks/event/README.md#updatemessagingevent2) - Update an event

### [Expense](docs/sdks/expense/README.md)

* [CreateAccountingExpense2](docs/sdks/expense/README.md#createaccountingexpense2) - Create an expense
* [GetAccountingExpense2](docs/sdks/expense/README.md#getaccountingexpense2) - Retrieve an expense
* [ListAccountingExpenses2](docs/sdks/expense/README.md#listaccountingexpenses2) - List all expenses
* [PatchAccountingExpense2](docs/sdks/expense/README.md#patchaccountingexpense2) - Update an expense
* [RemoveAccountingExpense2](docs/sdks/expense/README.md#removeaccountingexpense2) - Remove an expense
* [UpdateAccountingExpense2](docs/sdks/expense/README.md#updateaccountingexpense2) - Update an expense

### [File](docs/sdks/file/README.md)

* [CreateStorageFile2](docs/sdks/file/README.md#createstoragefile2) - Create a file
* [GetStorageFile2](docs/sdks/file/README.md#getstoragefile2) - Retrieve a file
* [ListStorageFiles2](docs/sdks/file/README.md#liststoragefiles2) - List all files
* [PatchStorageFile2](docs/sdks/file/README.md#patchstoragefile2) - Update a file
* [RemoveStorageFile2](docs/sdks/file/README.md#removestoragefile2) - Remove a file
* [UpdateStorageFile2](docs/sdks/file/README.md#updatestoragefile2) - Update a file

### [Form](docs/sdks/form/README.md)

* [CreateFormsForm2](docs/sdks/form/README.md#createformsform2) - Create a form
* [GetFormsForm2](docs/sdks/form/README.md#getformsform2) - Retrieve a form
* [ListFormsForms2](docs/sdks/form/README.md#listformsforms2) - List all forms
* [PatchFormsForm2](docs/sdks/form/README.md#patchformsform2) - Update a form
* [RemoveFormsForm2](docs/sdks/form/README.md#removeformsform2) - Remove a form
* [UpdateFormsForm2](docs/sdks/form/README.md#updateformsform2) - Update a form

### [Forms](docs/sdks/forms/README.md)

* [CreateFormsForm2](docs/sdks/forms/README.md#createformsform2) - Create a form
* [GetFormsForm2](docs/sdks/forms/README.md#getformsform2) - Retrieve a form
* [GetFormsSubmission2](docs/sdks/forms/README.md#getformssubmission2) - Retrieve a submission
* [ListFormsForms2](docs/sdks/forms/README.md#listformsforms2) - List all forms
* [ListFormsSubmissions2](docs/sdks/forms/README.md#listformssubmissions2) - List all submissions
* [PatchFormsForm2](docs/sdks/forms/README.md#patchformsform2) - Update a form
* [RemoveFormsForm2](docs/sdks/forms/README.md#removeformsform2) - Remove a form
* [UpdateFormsForm2](docs/sdks/forms/README.md#updateformsform2) - Update a form

### [Genai](docs/sdks/genai/README.md)

* [CreateGenaiEmbedding2](docs/sdks/genai/README.md#creategenaiembedding2) - Create an embedding
* [CreateGenaiPrompt2](docs/sdks/genai/README.md#creategenaiprompt2) - Create a prompt
* [GetGenaiModel2](docs/sdks/genai/README.md#getgenaimodel2) - Retrieve a model
* [ListGenaiModels2](docs/sdks/genai/README.md#listgenaimodels2) - List all models

### [Group](docs/sdks/group/README.md)

* [CreateAdsGroup2](docs/sdks/group/README.md#createadsgroup2) - Create a group
* [CreateHrisGroup2](docs/sdks/group/README.md#createhrisgroup2) - Create a group
* [CreateScimGroups](docs/sdks/group/README.md#createscimgroups) - Create group
* [GetAdsGroup2](docs/sdks/group/README.md#getadsgroup2) - Retrieve a group
* [GetClubsGroup2](docs/sdks/group/README.md#getclubsgroup2) - Retrieve a group
* [GetHrisGroup2](docs/sdks/group/README.md#gethrisgroup2) - Retrieve a group
* [GetScimGroups](docs/sdks/group/README.md#getscimgroups) - Get group
* [ListAdsGroups2](docs/sdks/group/README.md#listadsgroups2) - List all groups
* [ListClubsGroups2](docs/sdks/group/README.md#listclubsgroups2) - List all groups
* [ListHrisGroups2](docs/sdks/group/README.md#listhrisgroups2) - List all groups
* [ListScimGroups](docs/sdks/group/README.md#listscimgroups) - List groups
* [PatchAdsGroup2](docs/sdks/group/README.md#patchadsgroup2) - Update a group
* [PatchHrisGroup2](docs/sdks/group/README.md#patchhrisgroup2) - Update a group
* [PatchScimGroups](docs/sdks/group/README.md#patchscimgroups) - Update group
* [RemoveAdsGroup2](docs/sdks/group/README.md#removeadsgroup2) - Remove a group
* [RemoveHrisGroup2](docs/sdks/group/README.md#removehrisgroup2) - Remove a group
* [RemoveScimGroups](docs/sdks/group/README.md#removescimgroups) - Delete group
* [UpdateAdsGroup2](docs/sdks/group/README.md#updateadsgroup2) - Update a group
* [UpdateHrisGroup2](docs/sdks/group/README.md#updatehrisgroup2) - Update a group
* [UpdateScimGroups](docs/sdks/group/README.md#updatescimgroups) - Update group

### [Hris](docs/sdks/hris/README.md)

* [CreateHrisBankaccount2](docs/sdks/hris/README.md#createhrisbankaccount2) - Create a bankaccount
* [CreateHrisBenefit2](docs/sdks/hris/README.md#createhrisbenefit2) - Create a benefit
* [CreateHrisCompany2](docs/sdks/hris/README.md#createhriscompany2) - Create a company
* [CreateHrisDeduction2](docs/sdks/hris/README.md#createhrisdeduction2) - Create a deduction
* [CreateHrisDevice2](docs/sdks/hris/README.md#createhrisdevice2) - Create a device
* [CreateHrisEmployee2](docs/sdks/hris/README.md#createhrisemployee2) - Create an employee
* [CreateHrisGroup2](docs/sdks/hris/README.md#createhrisgroup2) - Create a group
* [CreateHrisLocation2](docs/sdks/hris/README.md#createhrislocation2) - Create a location
* [CreateHrisTimeoff2](docs/sdks/hris/README.md#createhristimeoff2) - Create a timeoff
* [CreateHrisTimeshift2](docs/sdks/hris/README.md#createhristimeshift2) - Create a timeshift
* [GetHrisBankaccount2](docs/sdks/hris/README.md#gethrisbankaccount2) - Retrieve a bankaccount
* [GetHrisBenefit2](docs/sdks/hris/README.md#gethrisbenefit2) - Retrieve a benefit
* [GetHrisCompany2](docs/sdks/hris/README.md#gethriscompany2) - Retrieve a company
* [GetHrisDeduction2](docs/sdks/hris/README.md#gethrisdeduction2) - Retrieve a deduction
* [GetHrisDevice2](docs/sdks/hris/README.md#gethrisdevice2) - Retrieve a device
* [GetHrisEmployee2](docs/sdks/hris/README.md#gethrisemployee2) - Retrieve an employee
* [GetHrisGroup2](docs/sdks/hris/README.md#gethrisgroup2) - Retrieve a group
* [GetHrisLocation2](docs/sdks/hris/README.md#gethrislocation2) - Retrieve a location
* [GetHrisPayslip2](docs/sdks/hris/README.md#gethrispayslip2) - Retrieve a payslip
* [GetHrisTimeoff2](docs/sdks/hris/README.md#gethristimeoff2) - Retrieve a timeoff
* [GetHrisTimeshift2](docs/sdks/hris/README.md#gethristimeshift2) - Retrieve a timeshift
* [ListHrisBankaccounts2](docs/sdks/hris/README.md#listhrisbankaccounts2) - List all bankaccounts
* [ListHrisBenefits2](docs/sdks/hris/README.md#listhrisbenefits2) - List all benefits
* [ListHrisCompanies2](docs/sdks/hris/README.md#listhriscompanies2) - List all companies
* [ListHrisDeductions2](docs/sdks/hris/README.md#listhrisdeductions2) - List all deductions
* [ListHrisDevices2](docs/sdks/hris/README.md#listhrisdevices2) - List all devices
* [ListHrisEmployees2](docs/sdks/hris/README.md#listhrisemployees2) - List all employees
* [ListHrisGroups2](docs/sdks/hris/README.md#listhrisgroups2) - List all groups
* [ListHrisLocations2](docs/sdks/hris/README.md#listhrislocations2) - List all locations
* [ListHrisPayslips2](docs/sdks/hris/README.md#listhrispayslips2) - List all payslips
* [ListHrisTimeoffs2](docs/sdks/hris/README.md#listhristimeoffs2) - List all timeoffs
* [ListHrisTimeshifts2](docs/sdks/hris/README.md#listhristimeshifts2) - List all timeshifts
* [PatchHrisBankaccount2](docs/sdks/hris/README.md#patchhrisbankaccount2) - Update a bankaccount
* [PatchHrisBenefit2](docs/sdks/hris/README.md#patchhrisbenefit2) - Update a benefit
* [PatchHrisCompany2](docs/sdks/hris/README.md#patchhriscompany2) - Update a company
* [PatchHrisDeduction2](docs/sdks/hris/README.md#patchhrisdeduction2) - Update a deduction
* [PatchHrisDevice2](docs/sdks/hris/README.md#patchhrisdevice2) - Update a device
* [PatchHrisEmployee2](docs/sdks/hris/README.md#patchhrisemployee2) - Update an employee
* [PatchHrisGroup2](docs/sdks/hris/README.md#patchhrisgroup2) - Update a group
* [PatchHrisLocation2](docs/sdks/hris/README.md#patchhrislocation2) - Update a location
* [PatchHrisTimeoff2](docs/sdks/hris/README.md#patchhristimeoff2) - Update a timeoff
* [PatchHrisTimeshift2](docs/sdks/hris/README.md#patchhristimeshift2) - Update a timeshift
* [RemoveHrisBankaccount2](docs/sdks/hris/README.md#removehrisbankaccount2) - Remove a bankaccount
* [RemoveHrisBenefit2](docs/sdks/hris/README.md#removehrisbenefit2) - Remove a benefit
* [RemoveHrisCompany2](docs/sdks/hris/README.md#removehriscompany2) - Remove a company
* [RemoveHrisDeduction2](docs/sdks/hris/README.md#removehrisdeduction2) - Remove a deduction
* [RemoveHrisDevice2](docs/sdks/hris/README.md#removehrisdevice2) - Remove a device
* [RemoveHrisEmployee2](docs/sdks/hris/README.md#removehrisemployee2) - Remove an employee
* [RemoveHrisGroup2](docs/sdks/hris/README.md#removehrisgroup2) - Remove a group
* [RemoveHrisLocation2](docs/sdks/hris/README.md#removehrislocation2) - Remove a location
* [RemoveHrisTimeoff2](docs/sdks/hris/README.md#removehristimeoff2) - Remove a timeoff
* [RemoveHrisTimeshift2](docs/sdks/hris/README.md#removehristimeshift2) - Remove a timeshift
* [UpdateHrisBankaccount2](docs/sdks/hris/README.md#updatehrisbankaccount2) - Update a bankaccount
* [UpdateHrisBenefit2](docs/sdks/hris/README.md#updatehrisbenefit2) - Update a benefit
* [UpdateHrisCompany2](docs/sdks/hris/README.md#updatehriscompany2) - Update a company
* [UpdateHrisDeduction2](docs/sdks/hris/README.md#updatehrisdeduction2) - Update a deduction
* [UpdateHrisDevice2](docs/sdks/hris/README.md#updatehrisdevice2) - Update a device
* [UpdateHrisEmployee2](docs/sdks/hris/README.md#updatehrisemployee2) - Update an employee
* [UpdateHrisGroup2](docs/sdks/hris/README.md#updatehrisgroup2) - Update a group
* [UpdateHrisLocation2](docs/sdks/hris/README.md#updatehrislocation2) - Update a location
* [UpdateHrisTimeoff2](docs/sdks/hris/README.md#updatehristimeoff2) - Update a timeoff
* [UpdateHrisTimeshift2](docs/sdks/hris/README.md#updatehristimeshift2) - Update a timeshift

### [Insertionorder](docs/sdks/insertionorder/README.md)

* [CreateAdsInsertionorder2](docs/sdks/insertionorder/README.md#createadsinsertionorder2) - Create an insertionorder
* [GetAdsInsertionorder2](docs/sdks/insertionorder/README.md#getadsinsertionorder2) - Retrieve an insertionorder
* [ListAdsInsertionorders2](docs/sdks/insertionorder/README.md#listadsinsertionorders2) - List all insertionorders
* [PatchAdsInsertionorder2](docs/sdks/insertionorder/README.md#patchadsinsertionorder2) - Update an insertionorder
* [RemoveAdsInsertionorder2](docs/sdks/insertionorder/README.md#removeadsinsertionorder2) - Remove an insertionorder
* [UpdateAdsInsertionorder2](docs/sdks/insertionorder/README.md#updateadsinsertionorder2) - Update an insertionorder

### [Instructor](docs/sdks/instructor/README.md)

* [CreateLmsInstructor2](docs/sdks/instructor/README.md#createlmsinstructor2) - Create an instructor
* [GetLmsInstructor2](docs/sdks/instructor/README.md#getlmsinstructor2) - Retrieve an instructor
* [ListLmsInstructors2](docs/sdks/instructor/README.md#listlmsinstructors2) - List all instructors
* [PatchLmsInstructor2](docs/sdks/instructor/README.md#patchlmsinstructor2) - Update an instructor
* [RemoveLmsInstructor2](docs/sdks/instructor/README.md#removelmsinstructor2) - Remove an instructor
* [UpdateLmsInstructor2](docs/sdks/instructor/README.md#updatelmsinstructor2) - Update an instructor

### [Integration](docs/sdks/integration/README.md)

* [GetUnifiedIntegrationAuth](docs/sdks/integration/README.md#getunifiedintegrationauth) - Authorize new connection
* [ListUnifiedIntegrationWorkspaces](docs/sdks/integration/README.md#listunifiedintegrationworkspaces) - Returns all activated integrations in a workspace
* [ListUnifiedIntegrations](docs/sdks/integration/README.md#listunifiedintegrations) - Returns all integrations

### [Interview](docs/sdks/interview/README.md)

* [CreateAtsInterview2](docs/sdks/interview/README.md#createatsinterview2) - Create an interview
* [GetAtsInterview2](docs/sdks/interview/README.md#getatsinterview2) - Retrieve an interview
* [ListAtsInterviews2](docs/sdks/interview/README.md#listatsinterviews2) - List all interviews
* [PatchAtsInterview2](docs/sdks/interview/README.md#patchatsinterview2) - Update an interview
* [RemoveAtsInterview2](docs/sdks/interview/README.md#removeatsinterview2) - Remove an interview
* [UpdateAtsInterview2](docs/sdks/interview/README.md#updateatsinterview2) - Update an interview

### [Inventory](docs/sdks/inventory/README.md)

* [CreateCommerceInventory2](docs/sdks/inventory/README.md#createcommerceinventory2) - Create an inventory
* [GetCommerceInventory2](docs/sdks/inventory/README.md#getcommerceinventory2) - Retrieve an inventory
* [ListCommerceInventories2](docs/sdks/inventory/README.md#listcommerceinventories2) - List all inventories
* [PatchCommerceInventory2](docs/sdks/inventory/README.md#patchcommerceinventory2) - Update an inventory
* [RemoveCommerceInventory2](docs/sdks/inventory/README.md#removecommerceinventory2) - Remove an inventory
* [UpdateCommerceInventory2](docs/sdks/inventory/README.md#updatecommerceinventory2) - Update an inventory

### [Invoice](docs/sdks/invoice/README.md)

* [CreateAccountingInvoice2](docs/sdks/invoice/README.md#createaccountinginvoice2) - Create an invoice
* [GetAccountingInvoice2](docs/sdks/invoice/README.md#getaccountinginvoice2) - Retrieve an invoice
* [ListAccountingInvoices2](docs/sdks/invoice/README.md#listaccountinginvoices2) - List all invoices
* [PatchAccountingInvoice2](docs/sdks/invoice/README.md#patchaccountinginvoice2) - Update an invoice
* [RemoveAccountingInvoice2](docs/sdks/invoice/README.md#removeaccountinginvoice2) - Remove an invoice
* [UpdateAccountingInvoice2](docs/sdks/invoice/README.md#updateaccountinginvoice2) - Update an invoice

### [Issue](docs/sdks/issue/README.md)

* [GetUnifiedIssue](docs/sdks/issue/README.md#getunifiedissue) - Retrieve support issue
* [ListUnifiedIssues](docs/sdks/issue/README.md#listunifiedissues) - List support issues

### [Item](docs/sdks/item/README.md)

* [CreateCommerceItem2](docs/sdks/item/README.md#createcommerceitem2) - Create an item
* [GetCommerceItem2](docs/sdks/item/README.md#getcommerceitem2) - Retrieve an item
* [ListCommerceItems2](docs/sdks/item/README.md#listcommerceitems2) - List all items
* [PatchCommerceItem2](docs/sdks/item/README.md#patchcommerceitem2) - Update an item
* [RemoveCommerceItem2](docs/sdks/item/README.md#removecommerceitem2) - Remove an item
* [UpdateCommerceItem2](docs/sdks/item/README.md#updatecommerceitem2) - Update an item

### [Itemvariant](docs/sdks/itemvariant/README.md)

* [CreateCommerceItemvariant2](docs/sdks/itemvariant/README.md#createcommerceitemvariant2) - Create an itemvariant
* [GetCommerceItemvariant2](docs/sdks/itemvariant/README.md#getcommerceitemvariant2) - Retrieve an itemvariant
* [ListCommerceItemvariants2](docs/sdks/itemvariant/README.md#listcommerceitemvariants2) - List all itemvariants
* [PatchCommerceItemvariant2](docs/sdks/itemvariant/README.md#patchcommerceitemvariant2) - Update an itemvariant
* [RemoveCommerceItemvariant2](docs/sdks/itemvariant/README.md#removecommerceitemvariant2) - Remove an itemvariant
* [UpdateCommerceItemvariant2](docs/sdks/itemvariant/README.md#updatecommerceitemvariant2) - Update an itemvariant

### [Job](docs/sdks/job/README.md)

* [CreateAtsJob2](docs/sdks/job/README.md#createatsjob2) - Create a job
* [GetAtsJob2](docs/sdks/job/README.md#getatsjob2) - Retrieve a job
* [ListAtsJobs2](docs/sdks/job/README.md#listatsjobs2) - List all jobs
* [PatchAtsJob2](docs/sdks/job/README.md#patchatsjob2) - Update a job
* [RemoveAtsJob2](docs/sdks/job/README.md#removeatsjob2) - Remove a job
* [UpdateAtsJob2](docs/sdks/job/README.md#updateatsjob2) - Update a job

### [Journal](docs/sdks/journal/README.md)

* [CreateAccountingJournal2](docs/sdks/journal/README.md#createaccountingjournal2) - Create a journal
* [GetAccountingJournal2](docs/sdks/journal/README.md#getaccountingjournal2) - Retrieve a journal
* [ListAccountingJournals2](docs/sdks/journal/README.md#listaccountingjournals2) - List all journals
* [PatchAccountingJournal2](docs/sdks/journal/README.md#patchaccountingjournal2) - Update a journal
* [RemoveAccountingJournal2](docs/sdks/journal/README.md#removeaccountingjournal2) - Remove a journal
* [UpdateAccountingJournal2](docs/sdks/journal/README.md#updateaccountingjournal2) - Update a journal

### [Kms](docs/sdks/kms/README.md)

* [CreateKmsComment2](docs/sdks/kms/README.md#createkmscomment2) - Create a comment
* [CreateKmsPage2](docs/sdks/kms/README.md#createkmspage2) - Create a page
* [CreateKmsSpace2](docs/sdks/kms/README.md#createkmsspace2) - Create a space
* [GetKmsComment2](docs/sdks/kms/README.md#getkmscomment2) - Retrieve a comment
* [GetKmsPage2](docs/sdks/kms/README.md#getkmspage2) - Retrieve a page
* [GetKmsSpace2](docs/sdks/kms/README.md#getkmsspace2) - Retrieve a space
* [ListKmsComments2](docs/sdks/kms/README.md#listkmscomments2) - List all comments
* [ListKmsPages2](docs/sdks/kms/README.md#listkmspages2) - List all pages
* [ListKmsSpaces2](docs/sdks/kms/README.md#listkmsspaces2) - List all spaces
* [PatchKmsComment2](docs/sdks/kms/README.md#patchkmscomment2) - Update a comment
* [PatchKmsPage2](docs/sdks/kms/README.md#patchkmspage2) - Update a page
* [PatchKmsSpace2](docs/sdks/kms/README.md#patchkmsspace2) - Update a space
* [RemoveKmsComment2](docs/sdks/kms/README.md#removekmscomment2) - Remove a comment
* [RemoveKmsPage2](docs/sdks/kms/README.md#removekmspage2) - Remove a page
* [RemoveKmsSpace2](docs/sdks/kms/README.md#removekmsspace2) - Remove a space
* [UpdateKmsComment2](docs/sdks/kms/README.md#updatekmscomment2) - Update a comment
* [UpdateKmsPage2](docs/sdks/kms/README.md#updatekmspage2) - Update a page
* [UpdateKmsSpace2](docs/sdks/kms/README.md#updatekmsspace2) - Update a space

### [Label](docs/sdks/label/README.md)

* [CreateShippingLabel2](docs/sdks/label/README.md#createshippinglabel2) - Create a label
* [GetShippingLabel2](docs/sdks/label/README.md#getshippinglabel2) - Retrieve a label
* [ListShippingLabels2](docs/sdks/label/README.md#listshippinglabels2) - List all labels
* [PatchShippingLabel2](docs/sdks/label/README.md#patchshippinglabel2) - Update a label
* [RemoveShippingLabel2](docs/sdks/label/README.md#removeshippinglabel2) - Remove a label
* [UpdateShippingLabel2](docs/sdks/label/README.md#updateshippinglabel2) - Update a label

### [Lead](docs/sdks/lead/README.md)

* [CreateCrmLead2](docs/sdks/lead/README.md#createcrmlead2) - Create a lead
* [GetCrmLead2](docs/sdks/lead/README.md#getcrmlead2) - Retrieve a lead
* [ListCrmLeads2](docs/sdks/lead/README.md#listcrmleads2) - List all leads
* [PatchCrmLead2](docs/sdks/lead/README.md#patchcrmlead2) - Update a lead
* [RemoveCrmLead2](docs/sdks/lead/README.md#removecrmlead2) - Remove a lead
* [UpdateCrmLead2](docs/sdks/lead/README.md#updatecrmlead2) - Update a lead

### [Link](docs/sdks/link/README.md)

* [CreateCalendarLink2](docs/sdks/link/README.md#createcalendarlink2) - Create a link
* [CreatePaymentLink2](docs/sdks/link/README.md#createpaymentlink2) - Create a link
* [GetCalendarLink2](docs/sdks/link/README.md#getcalendarlink2) - Retrieve a link
* [GetPaymentLink2](docs/sdks/link/README.md#getpaymentlink2) - Retrieve a link
* [ListCalendarLinks2](docs/sdks/link/README.md#listcalendarlinks2) - List all links
* [ListPaymentLinks2](docs/sdks/link/README.md#listpaymentlinks2) - List all links
* [PatchCalendarLink2](docs/sdks/link/README.md#patchcalendarlink2) - Update a link
* [PatchPaymentLink2](docs/sdks/link/README.md#patchpaymentlink2) - Update a link
* [RemoveCalendarLink2](docs/sdks/link/README.md#removecalendarlink2) - Remove a link
* [RemovePaymentLink2](docs/sdks/link/README.md#removepaymentlink2) - Remove a link
* [UpdateCalendarLink2](docs/sdks/link/README.md#updatecalendarlink2) - Update a link
* [UpdatePaymentLink2](docs/sdks/link/README.md#updatepaymentlink2) - Update a link

### [List](docs/sdks/list/README.md)

* [CreateMartechList2](docs/sdks/list/README.md#createmartechlist2) - Create a list
* [GetMartechList2](docs/sdks/list/README.md#getmartechlist2) - Retrieve a list
* [ListMartechLists2](docs/sdks/list/README.md#listmartechlists2) - List all lists
* [PatchMartechList2](docs/sdks/list/README.md#patchmartechlist2) - Update a list
* [RemoveMartechList2](docs/sdks/list/README.md#removemartechlist2) - Remove a list
* [UpdateMartechList2](docs/sdks/list/README.md#updatemartechlist2) - Update a list

### [Lms](docs/sdks/lms/README.md)

* [CreateLmsActivity2](docs/sdks/lms/README.md#createlmsactivity2) - Create an activity
* [CreateLmsClass2](docs/sdks/lms/README.md#createlmsclass2) - Create a class
* [CreateLmsCollection2](docs/sdks/lms/README.md#createlmscollection2) - Create a collection
* [CreateLmsContent2](docs/sdks/lms/README.md#createlmscontent2) - Create a content
* [CreateLmsCourse2](docs/sdks/lms/README.md#createlmscourse2) - Create a course
* [CreateLmsInstructor2](docs/sdks/lms/README.md#createlmsinstructor2) - Create an instructor
* [CreateLmsStudent2](docs/sdks/lms/README.md#createlmsstudent2) - Create a student
* [GetLmsActivity2](docs/sdks/lms/README.md#getlmsactivity2) - Retrieve an activity
* [GetLmsClass2](docs/sdks/lms/README.md#getlmsclass2) - Retrieve a class
* [GetLmsCollection2](docs/sdks/lms/README.md#getlmscollection2) - Retrieve a collection
* [GetLmsContent2](docs/sdks/lms/README.md#getlmscontent2) - Retrieve a content
* [GetLmsCourse2](docs/sdks/lms/README.md#getlmscourse2) - Retrieve a course
* [GetLmsInstructor2](docs/sdks/lms/README.md#getlmsinstructor2) - Retrieve an instructor
* [GetLmsStudent2](docs/sdks/lms/README.md#getlmsstudent2) - Retrieve a student
* [ListLmsActivities2](docs/sdks/lms/README.md#listlmsactivities2) - List all activities
* [ListLmsClasses2](docs/sdks/lms/README.md#listlmsclasses2) - List all classes
* [ListLmsCollections2](docs/sdks/lms/README.md#listlmscollections2) - List all collections
* [ListLmsContents2](docs/sdks/lms/README.md#listlmscontents2) - List all contents
* [ListLmsCourses2](docs/sdks/lms/README.md#listlmscourses2) - List all courses
* [ListLmsInstructors2](docs/sdks/lms/README.md#listlmsinstructors2) - List all instructors
* [ListLmsStudents2](docs/sdks/lms/README.md#listlmsstudents2) - List all students
* [PatchLmsActivity2](docs/sdks/lms/README.md#patchlmsactivity2) - Update an activity
* [PatchLmsClass2](docs/sdks/lms/README.md#patchlmsclass2) - Update a class
* [PatchLmsCollection2](docs/sdks/lms/README.md#patchlmscollection2) - Update a collection
* [PatchLmsContent2](docs/sdks/lms/README.md#patchlmscontent2) - Update a content
* [PatchLmsCourse2](docs/sdks/lms/README.md#patchlmscourse2) - Update a course
* [PatchLmsInstructor2](docs/sdks/lms/README.md#patchlmsinstructor2) - Update an instructor
* [PatchLmsStudent2](docs/sdks/lms/README.md#patchlmsstudent2) - Update a student
* [RemoveLmsActivity2](docs/sdks/lms/README.md#removelmsactivity2) - Remove an activity
* [RemoveLmsClass2](docs/sdks/lms/README.md#removelmsclass2) - Remove a class
* [RemoveLmsCollection2](docs/sdks/lms/README.md#removelmscollection2) - Remove a collection
* [RemoveLmsContent2](docs/sdks/lms/README.md#removelmscontent2) - Remove a content
* [RemoveLmsCourse2](docs/sdks/lms/README.md#removelmscourse2) - Remove a course
* [RemoveLmsInstructor2](docs/sdks/lms/README.md#removelmsinstructor2) - Remove an instructor
* [RemoveLmsStudent2](docs/sdks/lms/README.md#removelmsstudent2) - Remove a student
* [UpdateLmsActivity2](docs/sdks/lms/README.md#updatelmsactivity2) - Update an activity
* [UpdateLmsClass2](docs/sdks/lms/README.md#updatelmsclass2) - Update a class
* [UpdateLmsCollection2](docs/sdks/lms/README.md#updatelmscollection2) - Update a collection
* [UpdateLmsContent2](docs/sdks/lms/README.md#updatelmscontent2) - Update a content
* [UpdateLmsCourse2](docs/sdks/lms/README.md#updatelmscourse2) - Update a course
* [UpdateLmsInstructor2](docs/sdks/lms/README.md#updatelmsinstructor2) - Update an instructor
* [UpdateLmsStudent2](docs/sdks/lms/README.md#updatelmsstudent2) - Update a student

### [Location](docs/sdks/location/README.md)

* [CreateCommerceLocation2](docs/sdks/location/README.md#createcommercelocation2) - Create a location
* [CreateHrisLocation2](docs/sdks/location/README.md#createhrislocation2) - Create a location
* [GetClubsLocation2](docs/sdks/location/README.md#getclubslocation2) - Retrieve a location
* [GetCommerceLocation2](docs/sdks/location/README.md#getcommercelocation2) - Retrieve a location
* [GetHrisLocation2](docs/sdks/location/README.md#gethrislocation2) - Retrieve a location
* [ListClubsLocations2](docs/sdks/location/README.md#listclubslocations2) - List all locations
* [ListCommerceLocations2](docs/sdks/location/README.md#listcommercelocations2) - List all locations
* [ListHrisLocations2](docs/sdks/location/README.md#listhrislocations2) - List all locations
* [PatchCommerceLocation2](docs/sdks/location/README.md#patchcommercelocation2) - Update a location
* [PatchHrisLocation2](docs/sdks/location/README.md#patchhrislocation2) - Update a location
* [RemoveCommerceLocation2](docs/sdks/location/README.md#removecommercelocation2) - Remove a location
* [RemoveHrisLocation2](docs/sdks/location/README.md#removehrislocation2) - Remove a location
* [UpdateCommerceLocation2](docs/sdks/location/README.md#updatecommercelocation2) - Update a location
* [UpdateHrisLocation2](docs/sdks/location/README.md#updatehrislocation2) - Update a location

### [Login](docs/sdks/login/README.md)

* [GetUnifiedIntegrationLogin](docs/sdks/login/README.md#getunifiedintegrationlogin) - Sign in a user

### [Martech](docs/sdks/martech/README.md)

* [CreateMartechCampaign2](docs/sdks/martech/README.md#createmartechcampaign2) - Create a campaign
* [CreateMartechList2](docs/sdks/martech/README.md#createmartechlist2) - Create a list
* [CreateMartechMember2](docs/sdks/martech/README.md#createmartechmember2) - Create a member
* [GetMartechCampaign2](docs/sdks/martech/README.md#getmartechcampaign2) - Retrieve a campaign
* [GetMartechList2](docs/sdks/martech/README.md#getmartechlist2) - Retrieve a list
* [GetMartechMember2](docs/sdks/martech/README.md#getmartechmember2) - Retrieve a member
* [ListMartechCampaigns2](docs/sdks/martech/README.md#listmartechcampaigns2) - List all campaigns
* [ListMartechLists2](docs/sdks/martech/README.md#listmartechlists2) - List all lists
* [ListMartechMembers2](docs/sdks/martech/README.md#listmartechmembers2) - List all members
* [ListMartechReports2](docs/sdks/martech/README.md#listmartechreports2) - List all reports
* [PatchMartechCampaign2](docs/sdks/martech/README.md#patchmartechcampaign2) - Update a campaign
* [PatchMartechList2](docs/sdks/martech/README.md#patchmartechlist2) - Update a list
* [PatchMartechMember2](docs/sdks/martech/README.md#patchmartechmember2) - Update a member
* [RemoveMartechCampaign2](docs/sdks/martech/README.md#removemartechcampaign2) - Remove a campaign
* [RemoveMartechList2](docs/sdks/martech/README.md#removemartechlist2) - Remove a list
* [RemoveMartechMember2](docs/sdks/martech/README.md#removemartechmember2) - Remove a member
* [UpdateMartechCampaign2](docs/sdks/martech/README.md#updatemartechcampaign2) - Update a campaign
* [UpdateMartechList2](docs/sdks/martech/README.md#updatemartechlist2) - Update a list
* [UpdateMartechMember2](docs/sdks/martech/README.md#updatemartechmember2) - Update a member

### [Member](docs/sdks/member/README.md)

* [CreateMartechMember2](docs/sdks/member/README.md#createmartechmember2) - Create a member
* [GetClubsMember2](docs/sdks/member/README.md#getclubsmember2) - Retrieve a member
* [GetMartechMember2](docs/sdks/member/README.md#getmartechmember2) - Retrieve a member
* [ListClubsMembers2](docs/sdks/member/README.md#listclubsmembers2) - List all members
* [ListMartechMembers2](docs/sdks/member/README.md#listmartechmembers2) - List all members
* [PatchMartechMember2](docs/sdks/member/README.md#patchmartechmember2) - Update a member
* [RemoveMartechMember2](docs/sdks/member/README.md#removemartechmember2) - Remove a member
* [UpdateMartechMember2](docs/sdks/member/README.md#updatemartechmember2) - Update a member

### [Message](docs/sdks/message/README.md)

* [CreateMessagingMessage2](docs/sdks/message/README.md#createmessagingmessage2) - Create a message
* [GetMessagingMessage2](docs/sdks/message/README.md#getmessagingmessage2) - Retrieve a message
* [ListMessagingMessages2](docs/sdks/message/README.md#listmessagingmessages2) - List all messages
* [PatchMessagingMessage2](docs/sdks/message/README.md#patchmessagingmessage2) - Update a message
* [RemoveMessagingMessage2](docs/sdks/message/README.md#removemessagingmessage2) - Remove a message
* [UpdateMessagingMessage2](docs/sdks/message/README.md#updatemessagingmessage2) - Update a message

### [Messaging](docs/sdks/messaging/README.md)

* [CreateMessagingMessage2](docs/sdks/messaging/README.md#createmessagingmessage2) - Create a message
* [GetMessagingChannel2](docs/sdks/messaging/README.md#getmessagingchannel2) - Retrieve a channel
* [GetMessagingMessage2](docs/sdks/messaging/README.md#getmessagingmessage2) - Retrieve a message
* [ListMessagingChannels2](docs/sdks/messaging/README.md#listmessagingchannels2) - List all channels
* [ListMessagingMessages2](docs/sdks/messaging/README.md#listmessagingmessages2) - List all messages
* [PatchMessagingEvent2](docs/sdks/messaging/README.md#patchmessagingevent2) - Update an event
* [PatchMessagingMessage2](docs/sdks/messaging/README.md#patchmessagingmessage2) - Update a message
* [RemoveMessagingMessage2](docs/sdks/messaging/README.md#removemessagingmessage2) - Remove a message
* [UpdateMessagingEvent2](docs/sdks/messaging/README.md#updatemessagingevent2) - Update an event
* [UpdateMessagingMessage2](docs/sdks/messaging/README.md#updatemessagingmessage2) - Update a message

### [Metadata](docs/sdks/metadata/README.md)

* [CreateMetadataMetadata2](docs/sdks/metadata/README.md#createmetadatametadata2) - Create a metadata
* [GetMetadataMetadata2](docs/sdks/metadata/README.md#getmetadatametadata2) - Retrieve a metadata
* [ListMetadataMetadatas2](docs/sdks/metadata/README.md#listmetadatametadatas2) - List all metadatas
* [PatchMetadataMetadata2](docs/sdks/metadata/README.md#patchmetadatametadata2) - Update a metadata
* [RemoveMetadataMetadata2](docs/sdks/metadata/README.md#removemetadatametadata2) - Remove a metadata
* [UpdateMetadataMetadata2](docs/sdks/metadata/README.md#updatemetadatametadata2) - Update a metadata

### [Model](docs/sdks/model/README.md)

* [GetGenaiModel2](docs/sdks/model/README.md#getgenaimodel2) - Retrieve a model
* [ListGenaiModels2](docs/sdks/model/README.md#listgenaimodels2) - List all models

### [Note](docs/sdks/note/README.md)

* [CreateTicketingNote2](docs/sdks/note/README.md#createticketingnote2) - Create a note
* [GetTicketingNote2](docs/sdks/note/README.md#getticketingnote2) - Retrieve a note
* [ListTicketingNotes2](docs/sdks/note/README.md#listticketingnotes2) - List all notes
* [PatchTicketingNote2](docs/sdks/note/README.md#patchticketingnote2) - Update a note
* [RemoveTicketingNote2](docs/sdks/note/README.md#removeticketingnote2) - Remove a note
* [UpdateTicketingNote2](docs/sdks/note/README.md#updateticketingnote2) - Update a note

### [Order](docs/sdks/order/README.md)

* [CreateAccountingOrder2](docs/sdks/order/README.md#createaccountingorder2) - Create an order
* [GetAccountingOrder2](docs/sdks/order/README.md#getaccountingorder2) - Retrieve an order
* [ListAccountingOrders2](docs/sdks/order/README.md#listaccountingorders2) - List all orders
* [PatchAccountingOrder2](docs/sdks/order/README.md#patchaccountingorder2) - Update an order
* [PatchAssessmentOrder2](docs/sdks/order/README.md#patchassessmentorder2) - Update an order
* [RemoveAccountingOrder2](docs/sdks/order/README.md#removeaccountingorder2) - Remove an order
* [UpdateAccountingOrder2](docs/sdks/order/README.md#updateaccountingorder2) - Update an order
* [UpdateAssessmentOrder2](docs/sdks/order/README.md#updateassessmentorder2) - Update an order

### [Organization](docs/sdks/organization/README.md)

* [CreateAdsOrganization2](docs/sdks/organization/README.md#createadsorganization2) - Create an organization
* [CreateRepoOrganization2](docs/sdks/organization/README.md#createrepoorganization2) - Create an organization
* [GetAccountingOrganization2](docs/sdks/organization/README.md#getaccountingorganization2) - Retrieve an organization
* [GetAdsOrganization2](docs/sdks/organization/README.md#getadsorganization2) - Retrieve an organization
* [GetRepoOrganization2](docs/sdks/organization/README.md#getrepoorganization2) - Retrieve an organization
* [ListAccountingOrganizations2](docs/sdks/organization/README.md#listaccountingorganizations2) - List all organizations
* [ListAdsOrganizations2](docs/sdks/organization/README.md#listadsorganizations2) - List all organizations
* [ListRepoOrganizations2](docs/sdks/organization/README.md#listrepoorganizations2) - List all organizations
* [PatchAdsOrganization2](docs/sdks/organization/README.md#patchadsorganization2) - Update an organization
* [PatchRepoOrganization2](docs/sdks/organization/README.md#patchrepoorganization2) - Update an organization
* [RemoveAdsOrganization2](docs/sdks/organization/README.md#removeadsorganization2) - Remove an organization
* [RemoveRepoOrganization2](docs/sdks/organization/README.md#removerepoorganization2) - Remove an organization
* [UpdateAdsOrganization2](docs/sdks/organization/README.md#updateadsorganization2) - Update an organization
* [UpdateRepoOrganization2](docs/sdks/organization/README.md#updaterepoorganization2) - Update an organization

### [Package](docs/sdks/package/README.md)

* [CreateAssessmentPackage2](docs/sdks/package/README.md#createassessmentpackage2) - Create an assessment package
* [GetAssessmentPackage2](docs/sdks/package/README.md#getassessmentpackage2) - Get an assessment package
* [GetVerificationPackage2](docs/sdks/package/README.md#getverificationpackage2) - Retrieve a package
* [ListAssessmentPackages2](docs/sdks/package/README.md#listassessmentpackages2) - List assessment packages
* [ListVerificationPackages2](docs/sdks/package/README.md#listverificationpackages2) - List all packages
* [PatchAssessmentPackage2](docs/sdks/package/README.md#patchassessmentpackage2) - Update an assessment package
* [RemoveAssessmentPackage2](docs/sdks/package/README.md#removeassessmentpackage2) - Delete an assessment package
* [UpdateAssessmentPackage2](docs/sdks/package/README.md#updateassessmentpackage2) - Update an assessment package

### [Page](docs/sdks/page/README.md)

* [CreateKmsPage2](docs/sdks/page/README.md#createkmspage2) - Create a page
* [GetKmsPage2](docs/sdks/page/README.md#getkmspage2) - Retrieve a page
* [ListKmsPages2](docs/sdks/page/README.md#listkmspages2) - List all pages
* [PatchKmsPage2](docs/sdks/page/README.md#patchkmspage2) - Update a page
* [RemoveKmsPage2](docs/sdks/page/README.md#removekmspage2) - Remove a page
* [UpdateKmsPage2](docs/sdks/page/README.md#updatekmspage2) - Update a page

### [Passthrough](docs/sdks/passthrough/README.md)

* [CreatePassthrough2JSON](docs/sdks/passthrough/README.md#createpassthrough2json) - Passthrough POST
* [CreatePassthrough2Raw](docs/sdks/passthrough/README.md#createpassthrough2raw) - Passthrough POST
* [ListPassthroughs2](docs/sdks/passthrough/README.md#listpassthroughs2) - Passthrough GET
* [PatchPassthrough2JSON](docs/sdks/passthrough/README.md#patchpassthrough2json) - Passthrough PUT
* [PatchPassthrough2Raw](docs/sdks/passthrough/README.md#patchpassthrough2raw) - Passthrough PUT
* [RemovePassthrough2](docs/sdks/passthrough/README.md#removepassthrough2) - Passthrough DELETE
* [UpdatePassthrough2JSON](docs/sdks/passthrough/README.md#updatepassthrough2json) - Passthrough PUT
* [UpdatePassthrough2Raw](docs/sdks/passthrough/README.md#updatepassthrough2raw) - Passthrough PUT

### [Payment](docs/sdks/payment/README.md)

* [CreatePaymentLink2](docs/sdks/payment/README.md#createpaymentlink2) - Create a link
* [CreatePaymentPayment2](docs/sdks/payment/README.md#createpaymentpayment2) - Create a payment
* [CreatePaymentSubscription2](docs/sdks/payment/README.md#createpaymentsubscription2) - Create a subscription
* [GetPaymentLink2](docs/sdks/payment/README.md#getpaymentlink2) - Retrieve a link
* [GetPaymentPayment2](docs/sdks/payment/README.md#getpaymentpayment2) - Retrieve a payment
* [GetPaymentPayout2](docs/sdks/payment/README.md#getpaymentpayout2) - Retrieve a payout
* [GetPaymentRefund2](docs/sdks/payment/README.md#getpaymentrefund2) - Retrieve a refund
* [GetPaymentSubscription2](docs/sdks/payment/README.md#getpaymentsubscription2) - Retrieve a subscription
* [ListPaymentLinks2](docs/sdks/payment/README.md#listpaymentlinks2) - List all links
* [ListPaymentPayments2](docs/sdks/payment/README.md#listpaymentpayments2) - List all payments
* [ListPaymentPayouts2](docs/sdks/payment/README.md#listpaymentpayouts2) - List all payouts
* [ListPaymentRefunds2](docs/sdks/payment/README.md#listpaymentrefunds2) - List all refunds
* [ListPaymentSubscriptions2](docs/sdks/payment/README.md#listpaymentsubscriptions2) - List all subscriptions
* [PatchPaymentLink2](docs/sdks/payment/README.md#patchpaymentlink2) - Update a link
* [PatchPaymentPayment2](docs/sdks/payment/README.md#patchpaymentpayment2) - Update a payment
* [PatchPaymentSubscription2](docs/sdks/payment/README.md#patchpaymentsubscription2) - Update a subscription
* [RemovePaymentLink2](docs/sdks/payment/README.md#removepaymentlink2) - Remove a link
* [RemovePaymentPayment2](docs/sdks/payment/README.md#removepaymentpayment2) - Remove a payment
* [RemovePaymentSubscription2](docs/sdks/payment/README.md#removepaymentsubscription2) - Remove a subscription
* [UpdatePaymentLink2](docs/sdks/payment/README.md#updatepaymentlink2) - Update a link
* [UpdatePaymentPayment2](docs/sdks/payment/README.md#updatepaymentpayment2) - Update a payment
* [UpdatePaymentSubscription2](docs/sdks/payment/README.md#updatepaymentsubscription2) - Update a subscription

### [Payout](docs/sdks/payout/README.md)

* [GetPaymentPayout2](docs/sdks/payout/README.md#getpaymentpayout2) - Retrieve a payout
* [ListPaymentPayouts2](docs/sdks/payout/README.md#listpaymentpayouts2) - List all payouts

### [Payslip](docs/sdks/payslip/README.md)

* [GetHrisPayslip2](docs/sdks/payslip/README.md#gethrispayslip2) - Retrieve a payslip
* [ListHrisPayslips2](docs/sdks/payslip/README.md#listhrispayslips2) - List all payslips

### [Person](docs/sdks/person/README.md)

* [ListEnrichPeople2](docs/sdks/person/README.md#listenrichpeople2) - Retrieve enrichment information for a person

### [Pipeline](docs/sdks/pipeline/README.md)

* [CreateCrmPipeline2](docs/sdks/pipeline/README.md#createcrmpipeline2) - Create a pipeline
* [GetCrmPipeline2](docs/sdks/pipeline/README.md#getcrmpipeline2) - Retrieve a pipeline
* [ListCrmPipelines2](docs/sdks/pipeline/README.md#listcrmpipelines2) - List all pipelines
* [PatchCrmPipeline2](docs/sdks/pipeline/README.md#patchcrmpipeline2) - Update a pipeline
* [RemoveCrmPipeline2](docs/sdks/pipeline/README.md#removecrmpipeline2) - Remove a pipeline
* [UpdateCrmPipeline2](docs/sdks/pipeline/README.md#updatecrmpipeline2) - Update a pipeline

### [Profitloss](docs/sdks/profitloss/README.md)

* [GetAccountingProfitloss2](docs/sdks/profitloss/README.md#getaccountingprofitloss2) - Retrieve a profitloss
* [ListAccountingProfitlosses2](docs/sdks/profitloss/README.md#listaccountingprofitlosses2) - List all profitlosses

### [Project](docs/sdks/project/README.md)

* [CreateTaskProject2](docs/sdks/project/README.md#createtaskproject2) - Create a project
* [GetTaskProject2](docs/sdks/project/README.md#gettaskproject2) - Retrieve a project
* [ListTaskProjects2](docs/sdks/project/README.md#listtaskprojects2) - List all projects
* [PatchTaskProject2](docs/sdks/project/README.md#patchtaskproject2) - Update a project
* [RemoveTaskProject2](docs/sdks/project/README.md#removetaskproject2) - Remove a project
* [UpdateTaskProject2](docs/sdks/project/README.md#updatetaskproject2) - Update a project

### [Promoted](docs/sdks/promoted/README.md)

* [GetAdsPromoted2](docs/sdks/promoted/README.md#getadspromoted2) - Retrieve a promoted
* [ListAdsPromoteds2](docs/sdks/promoted/README.md#listadspromoteds2) - List all promoteds

### [Prompt](docs/sdks/prompt/README.md)

* [CreateGenaiPrompt2](docs/sdks/prompt/README.md#creategenaiprompt2) - Create a prompt

### [Property](docs/sdks/property/README.md)

* [CreateAnalyticsProperty2](docs/sdks/property/README.md#createanalyticsproperty2) - Create a property
* [GetAnalyticsProperty2](docs/sdks/property/README.md#getanalyticsproperty2) - Retrieve a property
* [ListAnalyticsProperties2](docs/sdks/property/README.md#listanalyticsproperties2) - List all properties
* [PatchAnalyticsProperty2](docs/sdks/property/README.md#patchanalyticsproperty2) - Update a property
* [RemoveAnalyticsProperty2](docs/sdks/property/README.md#removeanalyticsproperty2) - Remove a property
* [UpdateAnalyticsProperty2](docs/sdks/property/README.md#updateanalyticsproperty2) - Update a property

### [Pullrequest](docs/sdks/pullrequest/README.md)

* [CreateRepoPullrequest2](docs/sdks/pullrequest/README.md#createrepopullrequest2) - Create a pullrequest
* [GetRepoPullrequest2](docs/sdks/pullrequest/README.md#getrepopullrequest2) - Retrieve a pullrequest
* [ListRepoPullrequests2](docs/sdks/pullrequest/README.md#listrepopullrequests2) - List all pullrequests
* [PatchRepoPullrequest2](docs/sdks/pullrequest/README.md#patchrepopullrequest2) - Update a pullrequest
* [RemoveRepoPullrequest2](docs/sdks/pullrequest/README.md#removerepopullrequest2) - Remove a pullrequest
* [UpdateRepoPullrequest2](docs/sdks/pullrequest/README.md#updaterepopullrequest2) - Update a pullrequest

### [Purchaseorder](docs/sdks/purchaseorder/README.md)

* [CreateAccountingPurchaseorder2](docs/sdks/purchaseorder/README.md#createaccountingpurchaseorder2) - Create a purchaseorder
* [GetAccountingPurchaseorder2](docs/sdks/purchaseorder/README.md#getaccountingpurchaseorder2) - Retrieve a purchaseorder
* [ListAccountingPurchaseorders2](docs/sdks/purchaseorder/README.md#listaccountingpurchaseorders2) - List all purchaseorders
* [PatchAccountingPurchaseorder2](docs/sdks/purchaseorder/README.md#patchaccountingpurchaseorder2) - Update a purchaseorder
* [RemoveAccountingPurchaseorder2](docs/sdks/purchaseorder/README.md#removeaccountingpurchaseorder2) - Remove a purchaseorder
* [UpdateAccountingPurchaseorder2](docs/sdks/purchaseorder/README.md#updateaccountingpurchaseorder2) - Update a purchaseorder

### [Query](docs/sdks/query/README.md)

* [CreateDatastoreQuery2](docs/sdks/query/README.md#createdatastorequery2) - Create a query

### [Rate](docs/sdks/rate/README.md)

* [CreateShippingRate2](docs/sdks/rate/README.md#createshippingrate2) - Create a rate

### [Record](docs/sdks/record/README.md)

* [CreateDatastoreRecord2](docs/sdks/record/README.md#createdatastorerecord2) - Create a record
* [GetDatastoreRecord2](docs/sdks/record/README.md#getdatastorerecord2) - Retrieve a record
* [ListDatastoreRecords2](docs/sdks/record/README.md#listdatastorerecords2) - List all records
* [PatchDatastoreRecord2](docs/sdks/record/README.md#patchdatastorerecord2) - Update a record
* [RemoveDatastoreRecord2](docs/sdks/record/README.md#removedatastorerecord2) - Remove a record
* [UpdateDatastoreRecord2](docs/sdks/record/README.md#updatedatastorerecord2) - Update a record

### [Recording](docs/sdks/recording/README.md)

* [CreateUcRecording2](docs/sdks/recording/README.md#createucrecording2) - Create a recording
* [GetCalendarRecording2](docs/sdks/recording/README.md#getcalendarrecording2) - Retrieve a recording
* [GetUcRecording2](docs/sdks/recording/README.md#getucrecording2) - Retrieve a recording
* [ListCalendarRecordings2](docs/sdks/recording/README.md#listcalendarrecordings2) - List all recordings
* [ListUcRecordings2](docs/sdks/recording/README.md#listucrecordings2) - List all recordings
* [PatchUcRecording2](docs/sdks/recording/README.md#patchucrecording2) - Update a recording
* [RemoveUcRecording2](docs/sdks/recording/README.md#removeucrecording2) - Remove a recording
* [UpdateUcRecording2](docs/sdks/recording/README.md#updateucrecording2) - Update a recording

### [Refund](docs/sdks/refund/README.md)

* [GetPaymentRefund2](docs/sdks/refund/README.md#getpaymentrefund2) - Retrieve a refund
* [ListPaymentRefunds2](docs/sdks/refund/README.md#listpaymentrefunds2) - List all refunds

### [Repo](docs/sdks/repo/README.md)

* [CreateRepoBranch2](docs/sdks/repo/README.md#createrepobranch2) - Create a branch
* [CreateRepoCommit2](docs/sdks/repo/README.md#createrepocommit2) - Create a commit
* [CreateRepoOrganization2](docs/sdks/repo/README.md#createrepoorganization2) - Create an organization
* [CreateRepoPullrequest2](docs/sdks/repo/README.md#createrepopullrequest2) - Create a pullrequest
* [CreateRepoRepository2](docs/sdks/repo/README.md#createreporepository2) - Create a repository
* [GetRepoBranch2](docs/sdks/repo/README.md#getrepobranch2) - Retrieve a branch
* [GetRepoCommit2](docs/sdks/repo/README.md#getrepocommit2) - Retrieve a commit
* [GetRepoOrganization2](docs/sdks/repo/README.md#getrepoorganization2) - Retrieve an organization
* [GetRepoPullrequest2](docs/sdks/repo/README.md#getrepopullrequest2) - Retrieve a pullrequest
* [GetRepoRepository2](docs/sdks/repo/README.md#getreporepository2) - Retrieve a repository
* [ListRepoBranches2](docs/sdks/repo/README.md#listrepobranches2) - List all branches
* [ListRepoCommits2](docs/sdks/repo/README.md#listrepocommits2) - List all commits
* [ListRepoOrganizations2](docs/sdks/repo/README.md#listrepoorganizations2) - List all organizations
* [ListRepoPullrequests2](docs/sdks/repo/README.md#listrepopullrequests2) - List all pullrequests
* [ListRepoRepositories2](docs/sdks/repo/README.md#listreporepositories2) - List all repositories
* [PatchRepoBranch2](docs/sdks/repo/README.md#patchrepobranch2) - Update a branch
* [PatchRepoCommit2](docs/sdks/repo/README.md#patchrepocommit2) - Update a commit
* [PatchRepoOrganization2](docs/sdks/repo/README.md#patchrepoorganization2) - Update an organization
* [PatchRepoPullrequest2](docs/sdks/repo/README.md#patchrepopullrequest2) - Update a pullrequest
* [PatchRepoRepository2](docs/sdks/repo/README.md#patchreporepository2) - Update a repository
* [RemoveRepoBranch2](docs/sdks/repo/README.md#removerepobranch2) - Remove a branch
* [RemoveRepoCommit2](docs/sdks/repo/README.md#removerepocommit2) - Remove a commit
* [RemoveRepoOrganization2](docs/sdks/repo/README.md#removerepoorganization2) - Remove an organization
* [RemoveRepoPullrequest2](docs/sdks/repo/README.md#removerepopullrequest2) - Remove a pullrequest
* [RemoveRepoRepository2](docs/sdks/repo/README.md#removereporepository2) - Remove a repository
* [UpdateRepoBranch2](docs/sdks/repo/README.md#updaterepobranch2) - Update a branch
* [UpdateRepoCommit2](docs/sdks/repo/README.md#updaterepocommit2) - Update a commit
* [UpdateRepoOrganization2](docs/sdks/repo/README.md#updaterepoorganization2) - Update an organization
* [UpdateRepoPullrequest2](docs/sdks/repo/README.md#updaterepopullrequest2) - Update a pullrequest
* [UpdateRepoRepository2](docs/sdks/repo/README.md#updatereporepository2) - Update a repository

### [Report](docs/sdks/report/README.md)

* [GetAccountingReport2](docs/sdks/report/README.md#getaccountingreport2) - Retrieve a report
* [ListAccountingReports2](docs/sdks/report/README.md#listaccountingreports2) - List all reports
* [ListAdsReports2](docs/sdks/report/README.md#listadsreports2) - List all reports
* [ListAnalyticsReports2](docs/sdks/report/README.md#listanalyticsreports2) - List all reports
* [ListMartechReports2](docs/sdks/report/README.md#listmartechreports2) - List all reports

### [Repository](docs/sdks/repository/README.md)

* [CreateRepoRepository2](docs/sdks/repository/README.md#createreporepository2) - Create a repository
* [GetRepoRepository2](docs/sdks/repository/README.md#getreporepository2) - Retrieve a repository
* [ListRepoRepositories2](docs/sdks/repository/README.md#listreporepositories2) - List all repositories
* [PatchRepoRepository2](docs/sdks/repository/README.md#patchreporepository2) - Update a repository
* [RemoveRepoRepository2](docs/sdks/repository/README.md#removereporepository2) - Remove a repository
* [UpdateRepoRepository2](docs/sdks/repository/README.md#updatereporepository2) - Update a repository

### [Request](docs/sdks/request/README.md)

* [CreateVerificationRequest2](docs/sdks/request/README.md#createverificationrequest2) - Create a request
* [GetVerificationRequest2](docs/sdks/request/README.md#getverificationrequest2) - Retrieve a request
* [ListVerificationRequests2](docs/sdks/request/README.md#listverificationrequests2) - List all requests
* [PatchVerificationRequest2](docs/sdks/request/README.md#patchverificationrequest2) - Update a request
* [RemoveVerificationRequest2](docs/sdks/request/README.md#removeverificationrequest2) - Remove a request
* [UpdateVerificationRequest2](docs/sdks/request/README.md#updateverificationrequest2) - Update a request

### [Reservation](docs/sdks/reservation/README.md)

* [CreateCommerceReservation2](docs/sdks/reservation/README.md#createcommercereservation2) - Create a reservation
* [GetCommerceReservation2](docs/sdks/reservation/README.md#getcommercereservation2) - Retrieve a reservation
* [ListCommerceReservations2](docs/sdks/reservation/README.md#listcommercereservations2) - List all reservations
* [PatchCommerceReservation2](docs/sdks/reservation/README.md#patchcommercereservation2) - Update a reservation
* [RemoveCommerceReservation2](docs/sdks/reservation/README.md#removecommercereservation2) - Remove a reservation
* [UpdateCommerceReservation2](docs/sdks/reservation/README.md#updatecommercereservation2) - Update a reservation

### [Review](docs/sdks/review/README.md)

* [CreateCommerceReview2](docs/sdks/review/README.md#createcommercereview2) - Create a review
* [GetCommerceReview2](docs/sdks/review/README.md#getcommercereview2) - Retrieve a review
* [ListCommerceReviews2](docs/sdks/review/README.md#listcommercereviews2) - List all reviews
* [PatchCommerceReview2](docs/sdks/review/README.md#patchcommercereview2) - Update a review
* [RemoveCommerceReview2](docs/sdks/review/README.md#removecommercereview2) - Remove a review
* [UpdateCommerceReview2](docs/sdks/review/README.md#updatecommercereview2) - Update a review

### [Saleschannel](docs/sdks/saleschannel/README.md)

* [CreateCommerceSaleschannel2](docs/sdks/saleschannel/README.md#createcommercesaleschannel2) - Create a saleschannel
* [GetCommerceSaleschannel2](docs/sdks/saleschannel/README.md#getcommercesaleschannel2) - Retrieve a saleschannel
* [ListCommerceSaleschannels2](docs/sdks/saleschannel/README.md#listcommercesaleschannels2) - List all saleschannels
* [PatchCommerceSaleschannel2](docs/sdks/saleschannel/README.md#patchcommercesaleschannel2) - Update a saleschannel
* [RemoveCommerceSaleschannel2](docs/sdks/saleschannel/README.md#removecommercesaleschannel2) - Remove a saleschannel
* [UpdateCommerceSaleschannel2](docs/sdks/saleschannel/README.md#updatecommercesaleschannel2) - Update a saleschannel

### [Salesorder](docs/sdks/salesorder/README.md)

* [CreateAccountingSalesorder2](docs/sdks/salesorder/README.md#createaccountingsalesorder2) - Create a salesorder
* [GetAccountingSalesorder2](docs/sdks/salesorder/README.md#getaccountingsalesorder2) - Retrieve a salesorder
* [ListAccountingSalesorders2](docs/sdks/salesorder/README.md#listaccountingsalesorders2) - List all salesorders
* [PatchAccountingSalesorder2](docs/sdks/salesorder/README.md#patchaccountingsalesorder2) - Update a salesorder
* [RemoveAccountingSalesorder2](docs/sdks/salesorder/README.md#removeaccountingsalesorder2) - Remove a salesorder
* [UpdateAccountingSalesorder2](docs/sdks/salesorder/README.md#updateaccountingsalesorder2) - Update a salesorder

### [Scim](docs/sdks/scim/README.md)

* [CreateScimGroups](docs/sdks/scim/README.md#createscimgroups) - Create group
* [CreateScimUsers](docs/sdks/scim/README.md#createscimusers) - Create user
* [GetScimGroups](docs/sdks/scim/README.md#getscimgroups) - Get group
* [GetScimUsers](docs/sdks/scim/README.md#getscimusers) - Get user
* [ListScimGroups](docs/sdks/scim/README.md#listscimgroups) - List groups
* [ListScimUsers](docs/sdks/scim/README.md#listscimusers) - List users
* [PatchScimGroups](docs/sdks/scim/README.md#patchscimgroups) - Update group
* [PatchScimUsers](docs/sdks/scim/README.md#patchscimusers) - Update user
* [RemoveScimGroups](docs/sdks/scim/README.md#removescimgroups) - Delete group
* [RemoveScimUsers](docs/sdks/scim/README.md#removescimusers) - Delete user
* [UpdateScimGroups](docs/sdks/scim/README.md#updatescimgroups) - Update group
* [UpdateScimUsers](docs/sdks/scim/README.md#updatescimusers) - Update user

### [Scorecard](docs/sdks/scorecard/README.md)

* [CreateAtsScorecard2](docs/sdks/scorecard/README.md#createatsscorecard2) - Create a scorecard
* [GetAtsScorecard2](docs/sdks/scorecard/README.md#getatsscorecard2) - Retrieve a scorecard
* [ListAtsScorecards2](docs/sdks/scorecard/README.md#listatsscorecards2) - List all scorecards
* [PatchAtsScorecard2](docs/sdks/scorecard/README.md#patchatsscorecard2) - Update a scorecard
* [RemoveAtsScorecard2](docs/sdks/scorecard/README.md#removeatsscorecard2) - Remove a scorecard
* [UpdateAtsScorecard2](docs/sdks/scorecard/README.md#updateatsscorecard2) - Update a scorecard

### [Session](docs/sdks/session/README.md)

* [GetAnalyticsSession2](docs/sdks/session/README.md#getanalyticssession2) - Retrieve a session
* [ListAnalyticsSessions2](docs/sdks/session/README.md#listanalyticssessions2) - List all sessions

### [Shipment](docs/sdks/shipment/README.md)

* [CreateShippingShipment2](docs/sdks/shipment/README.md#createshippingshipment2) - Create a shipment
* [GetShippingShipment2](docs/sdks/shipment/README.md#getshippingshipment2) - Retrieve a shipment
* [ListShippingShipments2](docs/sdks/shipment/README.md#listshippingshipments2) - List all shipments
* [PatchShippingShipment2](docs/sdks/shipment/README.md#patchshippingshipment2) - Update a shipment
* [RemoveShippingShipment2](docs/sdks/shipment/README.md#removeshippingshipment2) - Remove a shipment
* [UpdateShippingShipment2](docs/sdks/shipment/README.md#updateshippingshipment2) - Update a shipment

### [Shipping](docs/sdks/shipping/README.md)

* [CreateShippingLabel2](docs/sdks/shipping/README.md#createshippinglabel2) - Create a label
* [CreateShippingRate2](docs/sdks/shipping/README.md#createshippingrate2) - Create a rate
* [CreateShippingShipment2](docs/sdks/shipping/README.md#createshippingshipment2) - Create a shipment
* [GetShippingCarrier2](docs/sdks/shipping/README.md#getshippingcarrier2) - Retrieve a carrier
* [GetShippingLabel2](docs/sdks/shipping/README.md#getshippinglabel2) - Retrieve a label
* [GetShippingShipment2](docs/sdks/shipping/README.md#getshippingshipment2) - Retrieve a shipment
* [GetShippingTracking2](docs/sdks/shipping/README.md#getshippingtracking2) - Retrieve a tracking
* [ListShippingCarriers2](docs/sdks/shipping/README.md#listshippingcarriers2) - List all carriers
* [ListShippingLabels2](docs/sdks/shipping/README.md#listshippinglabels2) - List all labels
* [ListShippingShipments2](docs/sdks/shipping/README.md#listshippingshipments2) - List all shipments
* [ListShippingTrackings2](docs/sdks/shipping/README.md#listshippingtrackings2) - List all trackings
* [PatchShippingLabel2](docs/sdks/shipping/README.md#patchshippinglabel2) - Update a label
* [PatchShippingShipment2](docs/sdks/shipping/README.md#patchshippingshipment2) - Update a shipment
* [RemoveShippingLabel2](docs/sdks/shipping/README.md#removeshippinglabel2) - Remove a label
* [RemoveShippingShipment2](docs/sdks/shipping/README.md#removeshippingshipment2) - Remove a shipment
* [UpdateShippingLabel2](docs/sdks/shipping/README.md#updateshippinglabel2) - Update a label
* [UpdateShippingShipment2](docs/sdks/shipping/README.md#updateshippingshipment2) - Update a shipment

### [Signatory](docs/sdks/signatory/README.md)

* [CreateSigningSignatory2](docs/sdks/signatory/README.md#createsigningsignatory2) - Create a signatory
* [GetSigningSignatory2](docs/sdks/signatory/README.md#getsigningsignatory2) - Retrieve a signatory
* [ListSigningSignatories2](docs/sdks/signatory/README.md#listsigningsignatories2) - List all signatories
* [PatchSigningSignatory2](docs/sdks/signatory/README.md#patchsigningsignatory2) - Update a signatory
* [RemoveSigningSignatory2](docs/sdks/signatory/README.md#removesigningsignatory2) - Remove a signatory
* [UpdateSigningSignatory2](docs/sdks/signatory/README.md#updatesigningsignatory2) - Update a signatory

### [Signing](docs/sdks/signing/README.md)

* [CreateSigningDocument2](docs/sdks/signing/README.md#createsigningdocument2) - Create a document
* [CreateSigningSignatory2](docs/sdks/signing/README.md#createsigningsignatory2) - Create a signatory
* [GetSigningDocument2](docs/sdks/signing/README.md#getsigningdocument2) - Retrieve a document
* [GetSigningSignatory2](docs/sdks/signing/README.md#getsigningsignatory2) - Retrieve a signatory
* [GetSigningTemplate2](docs/sdks/signing/README.md#getsigningtemplate2) - Retrieve a template
* [ListSigningDocuments2](docs/sdks/signing/README.md#listsigningdocuments2) - List all documents
* [ListSigningSignatories2](docs/sdks/signing/README.md#listsigningsignatories2) - List all signatories
* [ListSigningTemplates2](docs/sdks/signing/README.md#listsigningtemplates2) - List all templates
* [PatchSigningDocument2](docs/sdks/signing/README.md#patchsigningdocument2) - Update a document
* [PatchSigningSignatory2](docs/sdks/signing/README.md#patchsigningsignatory2) - Update a signatory
* [RemoveSigningDocument2](docs/sdks/signing/README.md#removesigningdocument2) - Remove a document
* [RemoveSigningSignatory2](docs/sdks/signing/README.md#removesigningsignatory2) - Remove a signatory
* [UpdateSigningDocument2](docs/sdks/signing/README.md#updatesigningdocument2) - Update a document
* [UpdateSigningSignatory2](docs/sdks/signing/README.md#updatesigningsignatory2) - Update a signatory

### [Space](docs/sdks/space/README.md)

* [CreateKmsSpace2](docs/sdks/space/README.md#createkmsspace2) - Create a space
* [GetKmsSpace2](docs/sdks/space/README.md#getkmsspace2) - Retrieve a space
* [ListKmsSpaces2](docs/sdks/space/README.md#listkmsspaces2) - List all spaces
* [PatchKmsSpace2](docs/sdks/space/README.md#patchkmsspace2) - Update a space
* [RemoveKmsSpace2](docs/sdks/space/README.md#removekmsspace2) - Remove a space
* [UpdateKmsSpace2](docs/sdks/space/README.md#updatekmsspace2) - Update a space

### [Storage](docs/sdks/storage/README.md)

* [CreateStorageFile2](docs/sdks/storage/README.md#createstoragefile2) - Create a file
* [GetStorageFile2](docs/sdks/storage/README.md#getstoragefile2) - Retrieve a file
* [ListStorageFiles2](docs/sdks/storage/README.md#liststoragefiles2) - List all files
* [PatchStorageFile2](docs/sdks/storage/README.md#patchstoragefile2) - Update a file
* [RemoveStorageFile2](docs/sdks/storage/README.md#removestoragefile2) - Remove a file
* [UpdateStorageFile2](docs/sdks/storage/README.md#updatestoragefile2) - Update a file

### [Student](docs/sdks/student/README.md)

* [CreateLmsStudent2](docs/sdks/student/README.md#createlmsstudent2) - Create a student
* [GetLmsStudent2](docs/sdks/student/README.md#getlmsstudent2) - Retrieve a student
* [ListLmsStudents2](docs/sdks/student/README.md#listlmsstudents2) - List all students
* [PatchLmsStudent2](docs/sdks/student/README.md#patchlmsstudent2) - Update a student
* [RemoveLmsStudent2](docs/sdks/student/README.md#removelmsstudent2) - Remove a student
* [UpdateLmsStudent2](docs/sdks/student/README.md#updatelmsstudent2) - Update a student

### [Submission](docs/sdks/submission/README.md)

* [GetFormsSubmission2](docs/sdks/submission/README.md#getformssubmission2) - Retrieve a submission
* [ListFormsSubmissions2](docs/sdks/submission/README.md#listformssubmissions2) - List all submissions

### [Subscription](docs/sdks/subscription/README.md)

* [CreatePaymentSubscription2](docs/sdks/subscription/README.md#createpaymentsubscription2) - Create a subscription
* [GetPaymentSubscription2](docs/sdks/subscription/README.md#getpaymentsubscription2) - Retrieve a subscription
* [ListPaymentSubscriptions2](docs/sdks/subscription/README.md#listpaymentsubscriptions2) - List all subscriptions
* [PatchPaymentSubscription2](docs/sdks/subscription/README.md#patchpaymentsubscription2) - Update a subscription
* [RemovePaymentSubscription2](docs/sdks/subscription/README.md#removepaymentsubscription2) - Remove a subscription
* [UpdatePaymentSubscription2](docs/sdks/subscription/README.md#updatepaymentsubscription2) - Update a subscription

### [Table](docs/sdks/table/README.md)

* [CreateDatastoreTable2](docs/sdks/table/README.md#createdatastoretable2) - Create a table
* [GetDatastoreTable2](docs/sdks/table/README.md#getdatastoretable2) - Retrieve a table
* [ListDatastoreTables2](docs/sdks/table/README.md#listdatastoretables2) - List all tables
* [PatchDatastoreTable2](docs/sdks/table/README.md#patchdatastoretable2) - Update a table
* [RemoveDatastoreTable2](docs/sdks/table/README.md#removedatastoretable2) - Remove a table
* [UpdateDatastoreTable2](docs/sdks/table/README.md#updatedatastoretable2) - Update a table

### [Target](docs/sdks/target/README.md)

* [GetAdsTarget2](docs/sdks/target/README.md#getadstarget2) - Retrieve a target
* [ListAdsTargets2](docs/sdks/target/README.md#listadstargets2) - List all targets

### [Task](docs/sdks/task/README.md)

* [CreateTaskComment2](docs/sdks/task/README.md#createtaskcomment2) - Create a comment
* [CreateTaskProject2](docs/sdks/task/README.md#createtaskproject2) - Create a project
* [CreateTaskTask2](docs/sdks/task/README.md#createtasktask2) - Create a task
* [GetTaskChange2](docs/sdks/task/README.md#gettaskchange2) - Retrieve a change
* [GetTaskComment2](docs/sdks/task/README.md#gettaskcomment2) - Retrieve a comment
* [GetTaskProject2](docs/sdks/task/README.md#gettaskproject2) - Retrieve a project
* [GetTaskTask2](docs/sdks/task/README.md#gettasktask2) - Retrieve a task
* [ListTaskChanges2](docs/sdks/task/README.md#listtaskchanges2) - List all changes
* [ListTaskComments2](docs/sdks/task/README.md#listtaskcomments2) - List all comments
* [ListTaskProjects2](docs/sdks/task/README.md#listtaskprojects2) - List all projects
* [ListTaskTasks2](docs/sdks/task/README.md#listtasktasks2) - List all tasks
* [PatchTaskComment2](docs/sdks/task/README.md#patchtaskcomment2) - Update a comment
* [PatchTaskProject2](docs/sdks/task/README.md#patchtaskproject2) - Update a project
* [PatchTaskTask2](docs/sdks/task/README.md#patchtasktask2) - Update a task
* [RemoveTaskComment2](docs/sdks/task/README.md#removetaskcomment2) - Remove a comment
* [RemoveTaskProject2](docs/sdks/task/README.md#removetaskproject2) - Remove a project
* [RemoveTaskTask2](docs/sdks/task/README.md#removetasktask2) - Remove a task
* [UpdateTaskComment2](docs/sdks/task/README.md#updatetaskcomment2) - Update a comment
* [UpdateTaskProject2](docs/sdks/task/README.md#updatetaskproject2) - Update a project
* [UpdateTaskTask2](docs/sdks/task/README.md#updatetasktask2) - Update a task

### [Taxrate](docs/sdks/taxrate/README.md)

* [CreateAccountingTaxrate2](docs/sdks/taxrate/README.md#createaccountingtaxrate2) - Create a taxrate
* [GetAccountingTaxrate2](docs/sdks/taxrate/README.md#getaccountingtaxrate2) - Retrieve a taxrate
* [ListAccountingTaxrates2](docs/sdks/taxrate/README.md#listaccountingtaxrates2) - List all taxrates
* [PatchAccountingTaxrate2](docs/sdks/taxrate/README.md#patchaccountingtaxrate2) - Update a taxrate
* [RemoveAccountingTaxrate2](docs/sdks/taxrate/README.md#removeaccountingtaxrate2) - Remove a taxrate
* [UpdateAccountingTaxrate2](docs/sdks/taxrate/README.md#updateaccountingtaxrate2) - Update a taxrate

### [Template](docs/sdks/template/README.md)

* [GetSigningTemplate2](docs/sdks/template/README.md#getsigningtemplate2) - Retrieve a template
* [ListSigningTemplates2](docs/sdks/template/README.md#listsigningtemplates2) - List all templates

### [Ticket](docs/sdks/ticket/README.md)

* [CreateTicketingTicket2](docs/sdks/ticket/README.md#createticketingticket2) - Create a ticket
* [GetTicketingTicket2](docs/sdks/ticket/README.md#getticketingticket2) - Retrieve a ticket
* [ListTicketingTickets2](docs/sdks/ticket/README.md#listticketingtickets2) - List all tickets
* [PatchTicketingTicket2](docs/sdks/ticket/README.md#patchticketingticket2) - Update a ticket
* [RemoveTicketingTicket2](docs/sdks/ticket/README.md#removeticketingticket2) - Remove a ticket
* [UpdateTicketingTicket2](docs/sdks/ticket/README.md#updateticketingticket2) - Update a ticket

### [Ticketing](docs/sdks/ticketing/README.md)

* [CreateTicketingCategory2](docs/sdks/ticketing/README.md#createticketingcategory2) - Create a category
* [CreateTicketingCustomer2](docs/sdks/ticketing/README.md#createticketingcustomer2) - Create a customer
* [CreateTicketingNote2](docs/sdks/ticketing/README.md#createticketingnote2) - Create a note
* [CreateTicketingTicket2](docs/sdks/ticketing/README.md#createticketingticket2) - Create a ticket
* [GetTicketingCategory2](docs/sdks/ticketing/README.md#getticketingcategory2) - Retrieve a category
* [GetTicketingCustomer2](docs/sdks/ticketing/README.md#getticketingcustomer2) - Retrieve a customer
* [GetTicketingNote2](docs/sdks/ticketing/README.md#getticketingnote2) - Retrieve a note
* [GetTicketingTicket2](docs/sdks/ticketing/README.md#getticketingticket2) - Retrieve a ticket
* [ListTicketingCategories2](docs/sdks/ticketing/README.md#listticketingcategories2) - List all categories
* [ListTicketingCustomers2](docs/sdks/ticketing/README.md#listticketingcustomers2) - List all customers
* [ListTicketingNotes2](docs/sdks/ticketing/README.md#listticketingnotes2) - List all notes
* [ListTicketingTickets2](docs/sdks/ticketing/README.md#listticketingtickets2) - List all tickets
* [PatchTicketingCategory2](docs/sdks/ticketing/README.md#patchticketingcategory2) - Update a category
* [PatchTicketingCustomer2](docs/sdks/ticketing/README.md#patchticketingcustomer2) - Update a customer
* [PatchTicketingNote2](docs/sdks/ticketing/README.md#patchticketingnote2) - Update a note
* [PatchTicketingTicket2](docs/sdks/ticketing/README.md#patchticketingticket2) - Update a ticket
* [RemoveTicketingCategory2](docs/sdks/ticketing/README.md#removeticketingcategory2) - Remove a category
* [RemoveTicketingCustomer2](docs/sdks/ticketing/README.md#removeticketingcustomer2) - Remove a customer
* [RemoveTicketingNote2](docs/sdks/ticketing/README.md#removeticketingnote2) - Remove a note
* [RemoveTicketingTicket2](docs/sdks/ticketing/README.md#removeticketingticket2) - Remove a ticket
* [UpdateTicketingCategory2](docs/sdks/ticketing/README.md#updateticketingcategory2) - Update a category
* [UpdateTicketingCustomer2](docs/sdks/ticketing/README.md#updateticketingcustomer2) - Update a customer
* [UpdateTicketingNote2](docs/sdks/ticketing/README.md#updateticketingnote2) - Update a note
* [UpdateTicketingTicket2](docs/sdks/ticketing/README.md#updateticketingticket2) - Update a ticket

### [Timeoff](docs/sdks/timeoff/README.md)

* [CreateHrisTimeoff2](docs/sdks/timeoff/README.md#createhristimeoff2) - Create a timeoff
* [GetHrisTimeoff2](docs/sdks/timeoff/README.md#gethristimeoff2) - Retrieve a timeoff
* [ListHrisTimeoffs2](docs/sdks/timeoff/README.md#listhristimeoffs2) - List all timeoffs
* [PatchHrisTimeoff2](docs/sdks/timeoff/README.md#patchhristimeoff2) - Update a timeoff
* [RemoveHrisTimeoff2](docs/sdks/timeoff/README.md#removehristimeoff2) - Remove a timeoff
* [UpdateHrisTimeoff2](docs/sdks/timeoff/README.md#updatehristimeoff2) - Update a timeoff

### [Timeshift](docs/sdks/timeshift/README.md)

* [CreateHrisTimeshift2](docs/sdks/timeshift/README.md#createhristimeshift2) - Create a timeshift
* [GetHrisTimeshift2](docs/sdks/timeshift/README.md#gethristimeshift2) - Retrieve a timeshift
* [ListHrisTimeshifts2](docs/sdks/timeshift/README.md#listhristimeshifts2) - List all timeshifts
* [PatchHrisTimeshift2](docs/sdks/timeshift/README.md#patchhristimeshift2) - Update a timeshift
* [RemoveHrisTimeshift2](docs/sdks/timeshift/README.md#removehristimeshift2) - Remove a timeshift
* [UpdateHrisTimeshift2](docs/sdks/timeshift/README.md#updatehristimeshift2) - Update a timeshift

### [Tracking](docs/sdks/tracking/README.md)

* [GetShippingTracking2](docs/sdks/tracking/README.md#getshippingtracking2) - Retrieve a tracking
* [ListShippingTrackings2](docs/sdks/tracking/README.md#listshippingtrackings2) - List all trackings

### [Transaction](docs/sdks/transaction/README.md)

* [CreateAccountingTransaction2](docs/sdks/transaction/README.md#createaccountingtransaction2) - Create a transaction
* [GetAccountingTransaction2](docs/sdks/transaction/README.md#getaccountingtransaction2) - Retrieve a transaction
* [ListAccountingTransactions2](docs/sdks/transaction/README.md#listaccountingtransactions2) - List all transactions
* [PatchAccountingTransaction2](docs/sdks/transaction/README.md#patchaccountingtransaction2) - Update a transaction
* [RemoveAccountingTransaction2](docs/sdks/transaction/README.md#removeaccountingtransaction2) - Remove a transaction
* [UpdateAccountingTransaction2](docs/sdks/transaction/README.md#updateaccountingtransaction2) - Update a transaction

### [Trialbalance](docs/sdks/trialbalance/README.md)

* [GetAccountingTrialbalance2](docs/sdks/trialbalance/README.md#getaccountingtrialbalance2) - Retrieve a trialbalance
* [ListAccountingTrialbalances2](docs/sdks/trialbalance/README.md#listaccountingtrialbalances2) - List all trialbalances

### [Uc](docs/sdks/uc/README.md)

* [CreateUcComment2](docs/sdks/uc/README.md#createuccomment2) - Create a comment
* [CreateUcContact2](docs/sdks/uc/README.md#createuccontact2) - Create a contact
* [CreateUcRecording2](docs/sdks/uc/README.md#createucrecording2) - Create a recording
* [GetUcCall2](docs/sdks/uc/README.md#getuccall2) - Retrieve a call
* [GetUcComment2](docs/sdks/uc/README.md#getuccomment2) - Retrieve a comment
* [GetUcContact2](docs/sdks/uc/README.md#getuccontact2) - Retrieve a contact
* [GetUcRecording2](docs/sdks/uc/README.md#getucrecording2) - Retrieve a recording
* [ListUcCalls2](docs/sdks/uc/README.md#listuccalls2) - List all calls
* [ListUcComments2](docs/sdks/uc/README.md#listuccomments2) - List all comments
* [ListUcContacts2](docs/sdks/uc/README.md#listuccontacts2) - List all contacts
* [ListUcRecordings2](docs/sdks/uc/README.md#listucrecordings2) - List all recordings
* [PatchUcComment2](docs/sdks/uc/README.md#patchuccomment2) - Update a comment
* [PatchUcContact2](docs/sdks/uc/README.md#patchuccontact2) - Update a contact
* [PatchUcRecording2](docs/sdks/uc/README.md#patchucrecording2) - Update a recording
* [RemoveUcComment2](docs/sdks/uc/README.md#removeuccomment2) - Remove a comment
* [RemoveUcContact2](docs/sdks/uc/README.md#removeuccontact2) - Remove a contact
* [RemoveUcRecording2](docs/sdks/uc/README.md#removeucrecording2) - Remove a recording
* [UpdateUcComment2](docs/sdks/uc/README.md#updateuccomment2) - Update a comment
* [UpdateUcContact2](docs/sdks/uc/README.md#updateuccontact2) - Update a contact
* [UpdateUcRecording2](docs/sdks/uc/README.md#updateucrecording2) - Update a recording

### [Unified](docs/sdks/unified/README.md)

* [CreateUnifiedConnection](docs/sdks/unified/README.md#createunifiedconnection) - Create connection
* [CreateUnifiedEnvironment](docs/sdks/unified/README.md#createunifiedenvironment) - Create new environments
* [CreateUnifiedWebhook](docs/sdks/unified/README.md#createunifiedwebhook) - Create webhook subscription
* [GetUnifiedApicall](docs/sdks/unified/README.md#getunifiedapicall) - Retrieve specific API Call by its ID
* [GetUnifiedConnection](docs/sdks/unified/README.md#getunifiedconnection) - Retrieve connection
* [GetUnifiedIntegrationAuth](docs/sdks/unified/README.md#getunifiedintegrationauth) - Authorize new connection
* [GetUnifiedIssue](docs/sdks/unified/README.md#getunifiedissue) - Retrieve support issue
* [GetUnifiedWebhook](docs/sdks/unified/README.md#getunifiedwebhook) - Retrieve webhook by its ID
* [ListUnifiedApicalls](docs/sdks/unified/README.md#listunifiedapicalls) - Returns API Calls
* [ListUnifiedConnections](docs/sdks/unified/README.md#listunifiedconnections) - List all connections
* [ListUnifiedEnvironments](docs/sdks/unified/README.md#listunifiedenvironments) - Returns all environments
* [ListUnifiedIntegrationWorkspaces](docs/sdks/unified/README.md#listunifiedintegrationworkspaces) - Returns all activated integrations in a workspace
* [ListUnifiedIntegrations](docs/sdks/unified/README.md#listunifiedintegrations) - Returns all integrations
* [ListUnifiedIssues](docs/sdks/unified/README.md#listunifiedissues) - List support issues
* [ListUnifiedWebhooks](docs/sdks/unified/README.md#listunifiedwebhooks) - Returns all registered webhooks
* [PatchUnifiedConnection](docs/sdks/unified/README.md#patchunifiedconnection) - Update connection
* [PatchUnifiedWebhook](docs/sdks/unified/README.md#patchunifiedwebhook) - Update webhook subscription
* [PatchUnifiedWebhookTrigger](docs/sdks/unified/README.md#patchunifiedwebhooktrigger) - Trigger webhook
* [RemoveUnifiedConnection](docs/sdks/unified/README.md#removeunifiedconnection) - Remove connection
* [RemoveUnifiedEnvironment](docs/sdks/unified/README.md#removeunifiedenvironment) - Remove an environment
* [RemoveUnifiedWebhook](docs/sdks/unified/README.md#removeunifiedwebhook) - Remove webhook subscription
* [UpdateUnifiedConnection](docs/sdks/unified/README.md#updateunifiedconnection) - Update connection
* [UpdateUnifiedWebhook](docs/sdks/unified/README.md#updateunifiedwebhook) - Update webhook subscription
* [UpdateUnifiedWebhookTrigger](docs/sdks/unified/README.md#updateunifiedwebhooktrigger) - Trigger webhook

### [User](docs/sdks/user/README.md)

* [CreateScimUsers](docs/sdks/user/README.md#createscimusers) - Create user
* [GetScimUsers](docs/sdks/user/README.md#getscimusers) - Get user
* [ListScimUsers](docs/sdks/user/README.md#listscimusers) - List users
* [PatchScimUsers](docs/sdks/user/README.md#patchscimusers) - Update user
* [RemoveScimUsers](docs/sdks/user/README.md#removescimusers) - Delete user
* [UpdateScimUsers](docs/sdks/user/README.md#updatescimusers) - Update user

### [Verification](docs/sdks/verification/README.md)

* [CreateVerificationRequest2](docs/sdks/verification/README.md#createverificationrequest2) - Create a request
* [GetVerificationPackage2](docs/sdks/verification/README.md#getverificationpackage2) - Retrieve a package
* [GetVerificationRequest2](docs/sdks/verification/README.md#getverificationrequest2) - Retrieve a request
* [ListVerificationPackages2](docs/sdks/verification/README.md#listverificationpackages2) - List all packages
* [ListVerificationRequests2](docs/sdks/verification/README.md#listverificationrequests2) - List all requests
* [PatchVerificationRequest2](docs/sdks/verification/README.md#patchverificationrequest2) - Update a request
* [RemoveVerificationRequest2](docs/sdks/verification/README.md#removeverificationrequest2) - Remove a request
* [UpdateVerificationRequest2](docs/sdks/verification/README.md#updateverificationrequest2) - Update a request

### [Visitor](docs/sdks/visitor/README.md)

* [CreateAnalyticsVisitor2](docs/sdks/visitor/README.md#createanalyticsvisitor2) - Create a visitor
* [GetAnalyticsVisitor2](docs/sdks/visitor/README.md#getanalyticsvisitor2) - Retrieve a visitor
* [ListAnalyticsVisitors2](docs/sdks/visitor/README.md#listanalyticsvisitors2) - List all visitors
* [PatchAnalyticsVisitor2](docs/sdks/visitor/README.md#patchanalyticsvisitor2) - Update a visitor
* [RemoveAnalyticsVisitor2](docs/sdks/visitor/README.md#removeanalyticsvisitor2) - Remove a visitor
* [UpdateAnalyticsVisitor2](docs/sdks/visitor/README.md#updateanalyticsvisitor2) - Update a visitor

### [Webhook](docs/sdks/webhook/README.md)

* [CreateUnifiedWebhook](docs/sdks/webhook/README.md#createunifiedwebhook) - Create webhook subscription
* [GetUnifiedWebhook](docs/sdks/webhook/README.md#getunifiedwebhook) - Retrieve webhook by its ID
* [ListUnifiedWebhooks](docs/sdks/webhook/README.md#listunifiedwebhooks) - Returns all registered webhooks
* [PatchUnifiedWebhook](docs/sdks/webhook/README.md#patchunifiedwebhook) - Update webhook subscription
* [PatchUnifiedWebhookTrigger](docs/sdks/webhook/README.md#patchunifiedwebhooktrigger) - Trigger webhook
* [RemoveUnifiedWebhook](docs/sdks/webhook/README.md#removeunifiedwebhook) - Remove webhook subscription
* [UpdateUnifiedWebhook](docs/sdks/webhook/README.md#updateunifiedwebhook) - Update webhook subscription
* [UpdateUnifiedWebhookTrigger](docs/sdks/webhook/README.md#updateunifiedwebhooktrigger) - Trigger webhook

### [Webinar](docs/sdks/webinar/README.md)

* [CreateCalendarWebinar2](docs/sdks/webinar/README.md#createcalendarwebinar2) - Create a webinar
* [GetCalendarWebinar2](docs/sdks/webinar/README.md#getcalendarwebinar2) - Retrieve a webinar
* [ListCalendarWebinars2](docs/sdks/webinar/README.md#listcalendarwebinars2) - List all webinars
* [PatchCalendarWebinar2](docs/sdks/webinar/README.md#patchcalendarwebinar2) - Update a webinar
* [RemoveCalendarWebinar2](docs/sdks/webinar/README.md#removecalendarwebinar2) - Remove a webinar
* [UpdateCalendarWebinar2](docs/sdks/webinar/README.md#updatecalendarwebinar2) - Update a webinar

</details>
<!-- End Available Resources and Operations [operations] -->









<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `CreateAccountingAccount2` function may return the following errors:

| Error Type         | Status Code | Content Type |
| ------------------ | ----------- | ------------ |
| sdkerrors.SDKError | 4XX, 5XX    | \*/\*        |

### Example

```go
package main

import (
	"context"
	"errors"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/sdkerrors"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"log"
)

func main() {
	ctx := context.Background()

	s := unifiedgosdk.New(
		unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Accounting.CreateAccountingAccount2(ctx, operations.CreateAccountingAccount2Request{
		AccountingAccount: shared.AccountingAccount{},
		ConnectionID:      "<id>",
	})
	if err != nil {

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->



<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex(serverIndex int)` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| #   | Server                      | Description                |
| --- | --------------------------- | -------------------------- |
| 0   | `https://api.unified.to`    | North American data region |
| 1   | `https://api-eu.unified.to` | European data region       |
| 2   | `https://api-au.unified.to` | Australian data region     |

#### Example

```go
package main

import (
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"log"
)

func main() {
	ctx := context.Background()

	s := unifiedgosdk.New(
		unifiedgosdk.WithServerIndex(0),
		unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Accounting.CreateAccountingAccount2(ctx, operations.CreateAccountingAccount2Request{
		AccountingAccount: shared.AccountingAccount{},
		ConnectionID:      "<id>",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.AccountingAccount != nil {
		// handle response
	}
}

```

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"log"
)

func main() {
	ctx := context.Background()

	s := unifiedgosdk.New(
		unifiedgosdk.WithServerURL("https://api-au.unified.to"),
		unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Accounting.CreateAccountingAccount2(ctx, operations.CreateAccountingAccount2Request{
		AccountingAccount: shared.AccountingAccount{},
		ConnectionID:      "<id>",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.AccountingAccount != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->



<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"

	"github.com/unified-to/unified-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = unifiedgosdk.New(unifiedgosdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->



<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name  | Type   | Scheme  |
| ----- | ------ | ------- |
| `Jwt` | apiKey | API key |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"log"
)

func main() {
	ctx := context.Background()

	s := unifiedgosdk.New(
		unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Accounting.CreateAccountingAccount2(ctx, operations.CreateAccountingAccount2Request{
		AccountingAccount: shared.AccountingAccount{},
		ConnectionID:      "<id>",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.AccountingAccount != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/retry"
	"log"
	"pkg/models/operations"
)

func main() {
	ctx := context.Background()

	s := unifiedgosdk.New(
		unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Accounting.CreateAccountingAccount2(ctx, operations.CreateAccountingAccount2Request{
		AccountingAccount: shared.AccountingAccount{},
		ConnectionID:      "<id>",
	}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.AccountingAccount != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := unifiedgosdk.New(
		unifiedgosdk.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Accounting.CreateAccountingAccount2(ctx, operations.CreateAccountingAccount2Request{
		AccountingAccount: shared.AccountingAccount{},
		ConnectionID:      "<id>",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.AccountingAccount != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
