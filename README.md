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

	res, err := s.Accounting.CreateAccountingAccount(ctx, operations.CreateAccountingAccountRequest{
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

* [CreateAccountingAccount](docs/sdks/account/README.md#createaccountingaccount) - Create an account
* [GetAccountingAccount](docs/sdks/account/README.md#getaccountingaccount) - Retrieve an account
* [ListAccountingAccounts](docs/sdks/account/README.md#listaccountingaccounts) - List all accounts
* [PatchAccountingAccount](docs/sdks/account/README.md#patchaccountingaccount) - Update an account
* [RemoveAccountingAccount](docs/sdks/account/README.md#removeaccountingaccount) - Remove an account
* [UpdateAccountingAccount](docs/sdks/account/README.md#updateaccountingaccount) - Update an account

### [Accounting](docs/sdks/accounting/README.md)

* [CreateAccountingAccount](docs/sdks/accounting/README.md#createaccountingaccount) - Create an account
* [CreateAccountingBill](docs/sdks/accounting/README.md#createaccountingbill) - Create a bill
* [CreateAccountingCategory](docs/sdks/accounting/README.md#createaccountingcategory) - Create a category
* [CreateAccountingContact](docs/sdks/accounting/README.md#createaccountingcontact) - Create a contact
* [CreateAccountingCreditmemo](docs/sdks/accounting/README.md#createaccountingcreditmemo) - Create a creditmemo
* [CreateAccountingExpense](docs/sdks/accounting/README.md#createaccountingexpense) - Create an expense
* [CreateAccountingInvoice](docs/sdks/accounting/README.md#createaccountinginvoice) - Create an invoice
* [CreateAccountingJournal](docs/sdks/accounting/README.md#createaccountingjournal) - Create a journal
* [CreateAccountingOrder](docs/sdks/accounting/README.md#createaccountingorder) - Create an order
* [CreateAccountingProject](docs/sdks/accounting/README.md#createaccountingproject) - Create a project
* [CreateAccountingPurchaseorder](docs/sdks/accounting/README.md#createaccountingpurchaseorder) - Create a purchaseorder
* [CreateAccountingQuote](docs/sdks/accounting/README.md#createaccountingquote) - Create a quote
* [CreateAccountingSalesorder](docs/sdks/accounting/README.md#createaccountingsalesorder) - Create a salesorder
* [CreateAccountingTaxrate](docs/sdks/accounting/README.md#createaccountingtaxrate) - Create a taxrate
* [CreateAccountingTransaction](docs/sdks/accounting/README.md#createaccountingtransaction) - Create a transaction
* [CreateAccountingVendorcredit](docs/sdks/accounting/README.md#createaccountingvendorcredit) - Create a vendorcredit
* [GetAccountingAccount](docs/sdks/accounting/README.md#getaccountingaccount) - Retrieve an account
* [GetAccountingAgedpayable](docs/sdks/accounting/README.md#getaccountingagedpayable) - Retrieve an agedpayable
* [GetAccountingAgedreceivable](docs/sdks/accounting/README.md#getaccountingagedreceivable) - Retrieve an agedreceivable
* [GetAccountingBalancesheet](docs/sdks/accounting/README.md#getaccountingbalancesheet) - Retrieve a balancesheet
* [GetAccountingBill](docs/sdks/accounting/README.md#getaccountingbill) - Retrieve a bill
* [GetAccountingCashflow](docs/sdks/accounting/README.md#getaccountingcashflow) - Retrieve a cashflow
* [GetAccountingCategory](docs/sdks/accounting/README.md#getaccountingcategory) - Retrieve a category
* [GetAccountingContact](docs/sdks/accounting/README.md#getaccountingcontact) - Retrieve a contact
* [GetAccountingCreditmemo](docs/sdks/accounting/README.md#getaccountingcreditmemo) - Retrieve a creditmemo
* [GetAccountingExpense](docs/sdks/accounting/README.md#getaccountingexpense) - Retrieve an expense
* [GetAccountingInvoice](docs/sdks/accounting/README.md#getaccountinginvoice) - Retrieve an invoice
* [GetAccountingJournal](docs/sdks/accounting/README.md#getaccountingjournal) - Retrieve a journal
* [GetAccountingOrder](docs/sdks/accounting/README.md#getaccountingorder) - Retrieve an order
* [GetAccountingOrganization](docs/sdks/accounting/README.md#getaccountingorganization) - Retrieve an organization
* [GetAccountingProfitloss](docs/sdks/accounting/README.md#getaccountingprofitloss) - Retrieve a profitloss
* [GetAccountingProject](docs/sdks/accounting/README.md#getaccountingproject) - Retrieve a project
* [GetAccountingPurchaseorder](docs/sdks/accounting/README.md#getaccountingpurchaseorder) - Retrieve a purchaseorder
* [GetAccountingQuote](docs/sdks/accounting/README.md#getaccountingquote) - Retrieve a quote
* [GetAccountingReport](docs/sdks/accounting/README.md#getaccountingreport) - Retrieve a report
* [GetAccountingSalesorder](docs/sdks/accounting/README.md#getaccountingsalesorder) - Retrieve a salesorder
* [GetAccountingTaxrate](docs/sdks/accounting/README.md#getaccountingtaxrate) - Retrieve a taxrate
* [GetAccountingTransaction](docs/sdks/accounting/README.md#getaccountingtransaction) - Retrieve a transaction
* [GetAccountingTrialbalance](docs/sdks/accounting/README.md#getaccountingtrialbalance) - Retrieve a trialbalance
* [GetAccountingVendorcredit](docs/sdks/accounting/README.md#getaccountingvendorcredit) - Retrieve a vendorcredit
* [ListAccountingAccounts](docs/sdks/accounting/README.md#listaccountingaccounts) - List all accounts
* [ListAccountingAgedpayables](docs/sdks/accounting/README.md#listaccountingagedpayables) - List all agedpayables
* [ListAccountingAgedreceivables](docs/sdks/accounting/README.md#listaccountingagedreceivables) - List all agedreceivables
* [ListAccountingBalancesheets](docs/sdks/accounting/README.md#listaccountingbalancesheets) - List all balancesheets
* [ListAccountingBills](docs/sdks/accounting/README.md#listaccountingbills) - List all bills
* [ListAccountingCashflows](docs/sdks/accounting/README.md#listaccountingcashflows) - List all cashflows
* [ListAccountingCategories](docs/sdks/accounting/README.md#listaccountingcategories) - List all categories
* [ListAccountingContacts](docs/sdks/accounting/README.md#listaccountingcontacts) - List all contacts
* [ListAccountingCreditmemoes](docs/sdks/accounting/README.md#listaccountingcreditmemoes) - List all creditmemoes
* [ListAccountingExpenses](docs/sdks/accounting/README.md#listaccountingexpenses) - List all expenses
* [ListAccountingInvoices](docs/sdks/accounting/README.md#listaccountinginvoices) - List all invoices
* [ListAccountingJournals](docs/sdks/accounting/README.md#listaccountingjournals) - List all journals
* [ListAccountingOrders](docs/sdks/accounting/README.md#listaccountingorders) - List all orders
* [ListAccountingOrganizations](docs/sdks/accounting/README.md#listaccountingorganizations) - List all organizations
* [ListAccountingProfitlosses](docs/sdks/accounting/README.md#listaccountingprofitlosses) - List all profitlosses
* [ListAccountingProjects](docs/sdks/accounting/README.md#listaccountingprojects) - List all projects
* [ListAccountingPurchaseorders](docs/sdks/accounting/README.md#listaccountingpurchaseorders) - List all purchaseorders
* [ListAccountingQuotes](docs/sdks/accounting/README.md#listaccountingquotes) - List all quotes
* [ListAccountingReports](docs/sdks/accounting/README.md#listaccountingreports) - List all reports
* [ListAccountingSalesorders](docs/sdks/accounting/README.md#listaccountingsalesorders) - List all salesorders
* [ListAccountingTaxrates](docs/sdks/accounting/README.md#listaccountingtaxrates) - List all taxrates
* [ListAccountingTransactions](docs/sdks/accounting/README.md#listaccountingtransactions) - List all transactions
* [ListAccountingTrialbalances](docs/sdks/accounting/README.md#listaccountingtrialbalances) - List all trialbalances
* [ListAccountingVendorcredits](docs/sdks/accounting/README.md#listaccountingvendorcredits) - List all vendorcredits
* [PatchAccountingAccount](docs/sdks/accounting/README.md#patchaccountingaccount) - Update an account
* [PatchAccountingBill](docs/sdks/accounting/README.md#patchaccountingbill) - Update a bill
* [PatchAccountingCategory](docs/sdks/accounting/README.md#patchaccountingcategory) - Update a category
* [PatchAccountingContact](docs/sdks/accounting/README.md#patchaccountingcontact) - Update a contact
* [PatchAccountingCreditmemo](docs/sdks/accounting/README.md#patchaccountingcreditmemo) - Update a creditmemo
* [PatchAccountingExpense](docs/sdks/accounting/README.md#patchaccountingexpense) - Update an expense
* [PatchAccountingInvoice](docs/sdks/accounting/README.md#patchaccountinginvoice) - Update an invoice
* [PatchAccountingJournal](docs/sdks/accounting/README.md#patchaccountingjournal) - Update a journal
* [PatchAccountingOrder](docs/sdks/accounting/README.md#patchaccountingorder) - Update an order
* [PatchAccountingProject](docs/sdks/accounting/README.md#patchaccountingproject) - Update a project
* [PatchAccountingPurchaseorder](docs/sdks/accounting/README.md#patchaccountingpurchaseorder) - Update a purchaseorder
* [PatchAccountingQuote](docs/sdks/accounting/README.md#patchaccountingquote) - Update a quote
* [PatchAccountingSalesorder](docs/sdks/accounting/README.md#patchaccountingsalesorder) - Update a salesorder
* [PatchAccountingTaxrate](docs/sdks/accounting/README.md#patchaccountingtaxrate) - Update a taxrate
* [PatchAccountingTransaction](docs/sdks/accounting/README.md#patchaccountingtransaction) - Update a transaction
* [PatchAccountingVendorcredit](docs/sdks/accounting/README.md#patchaccountingvendorcredit) - Update a vendorcredit
* [RemoveAccountingAccount](docs/sdks/accounting/README.md#removeaccountingaccount) - Remove an account
* [RemoveAccountingBill](docs/sdks/accounting/README.md#removeaccountingbill) - Remove a bill
* [RemoveAccountingCategory](docs/sdks/accounting/README.md#removeaccountingcategory) - Remove a category
* [RemoveAccountingContact](docs/sdks/accounting/README.md#removeaccountingcontact) - Remove a contact
* [RemoveAccountingCreditmemo](docs/sdks/accounting/README.md#removeaccountingcreditmemo) - Remove a creditmemo
* [RemoveAccountingExpense](docs/sdks/accounting/README.md#removeaccountingexpense) - Remove an expense
* [RemoveAccountingInvoice](docs/sdks/accounting/README.md#removeaccountinginvoice) - Remove an invoice
* [RemoveAccountingJournal](docs/sdks/accounting/README.md#removeaccountingjournal) - Remove a journal
* [RemoveAccountingOrder](docs/sdks/accounting/README.md#removeaccountingorder) - Remove an order
* [RemoveAccountingProject](docs/sdks/accounting/README.md#removeaccountingproject) - Remove a project
* [RemoveAccountingPurchaseorder](docs/sdks/accounting/README.md#removeaccountingpurchaseorder) - Remove a purchaseorder
* [RemoveAccountingQuote](docs/sdks/accounting/README.md#removeaccountingquote) - Remove a quote
* [RemoveAccountingSalesorder](docs/sdks/accounting/README.md#removeaccountingsalesorder) - Remove a salesorder
* [RemoveAccountingTaxrate](docs/sdks/accounting/README.md#removeaccountingtaxrate) - Remove a taxrate
* [RemoveAccountingTransaction](docs/sdks/accounting/README.md#removeaccountingtransaction) - Remove a transaction
* [RemoveAccountingVendorcredit](docs/sdks/accounting/README.md#removeaccountingvendorcredit) - Remove a vendorcredit
* [UpdateAccountingAccount](docs/sdks/accounting/README.md#updateaccountingaccount) - Update an account
* [UpdateAccountingBill](docs/sdks/accounting/README.md#updateaccountingbill) - Update a bill
* [UpdateAccountingCategory](docs/sdks/accounting/README.md#updateaccountingcategory) - Update a category
* [UpdateAccountingContact](docs/sdks/accounting/README.md#updateaccountingcontact) - Update a contact
* [UpdateAccountingCreditmemo](docs/sdks/accounting/README.md#updateaccountingcreditmemo) - Update a creditmemo
* [UpdateAccountingExpense](docs/sdks/accounting/README.md#updateaccountingexpense) - Update an expense
* [UpdateAccountingInvoice](docs/sdks/accounting/README.md#updateaccountinginvoice) - Update an invoice
* [UpdateAccountingJournal](docs/sdks/accounting/README.md#updateaccountingjournal) - Update a journal
* [UpdateAccountingOrder](docs/sdks/accounting/README.md#updateaccountingorder) - Update an order
* [UpdateAccountingProject](docs/sdks/accounting/README.md#updateaccountingproject) - Update a project
* [UpdateAccountingPurchaseorder](docs/sdks/accounting/README.md#updateaccountingpurchaseorder) - Update a purchaseorder
* [UpdateAccountingQuote](docs/sdks/accounting/README.md#updateaccountingquote) - Update a quote
* [UpdateAccountingSalesorder](docs/sdks/accounting/README.md#updateaccountingsalesorder) - Update a salesorder
* [UpdateAccountingTaxrate](docs/sdks/accounting/README.md#updateaccountingtaxrate) - Update a taxrate
* [UpdateAccountingTransaction](docs/sdks/accounting/README.md#updateaccountingtransaction) - Update a transaction
* [UpdateAccountingVendorcredit](docs/sdks/accounting/README.md#updateaccountingvendorcredit) - Update a vendorcredit

### [Activation](docs/sdks/activation/README.md)

* [CreateCdpActivation](docs/sdks/activation/README.md#createcdpactivation) - Create an activation
* [GetCdpActivation](docs/sdks/activation/README.md#getcdpactivation) - Retrieve an activation
* [ListCdpActivations](docs/sdks/activation/README.md#listcdpactivations) - List all activations
* [PatchCdpActivation](docs/sdks/activation/README.md#patchcdpactivation) - Update an activation
* [RemoveCdpActivation](docs/sdks/activation/README.md#removecdpactivation) - Remove an activation
* [UpdateCdpActivation](docs/sdks/activation/README.md#updatecdpactivation) - Update an activation

### [Activity](docs/sdks/activity/README.md)

* [CreateAtsActivity](docs/sdks/activity/README.md#createatsactivity) - Create an activity
* [CreateLmsActivity](docs/sdks/activity/README.md#createlmsactivity) - Create an activity
* [GetAtsActivity](docs/sdks/activity/README.md#getatsactivity) - Retrieve an activity
* [GetClubsActivity](docs/sdks/activity/README.md#getclubsactivity) - Retrieve an activity
* [GetLmsActivity](docs/sdks/activity/README.md#getlmsactivity) - Retrieve an activity
* [ListAtsActivities](docs/sdks/activity/README.md#listatsactivities) - List all activities
* [ListClubsActivities](docs/sdks/activity/README.md#listclubsactivities) - List all activities
* [ListLmsActivities](docs/sdks/activity/README.md#listlmsactivities) - List all activities
* [PatchAtsActivity](docs/sdks/activity/README.md#patchatsactivity) - Update an activity
* [PatchLmsActivity](docs/sdks/activity/README.md#patchlmsactivity) - Update an activity
* [RemoveAtsActivity](docs/sdks/activity/README.md#removeatsactivity) - Remove an activity
* [RemoveLmsActivity](docs/sdks/activity/README.md#removelmsactivity) - Remove an activity
* [UpdateAtsActivity](docs/sdks/activity/README.md#updateatsactivity) - Update an activity
* [UpdateLmsActivity](docs/sdks/activity/README.md#updatelmsactivity) - Update an activity

### [Ad](docs/sdks/ad/README.md)

* [CreateAdsAd](docs/sdks/ad/README.md#createadsad) - Create an ad
* [GetAdsAd](docs/sdks/ad/README.md#getadsad) - Retrieve an ad
* [ListAdsAds](docs/sdks/ad/README.md#listadsads) - List all ads
* [PatchAdsAd](docs/sdks/ad/README.md#patchadsad) - Update an ad
* [RemoveAdsAd](docs/sdks/ad/README.md#removeadsad) - Remove an ad
* [UpdateAdsAd](docs/sdks/ad/README.md#updateadsad) - Update an ad

### [Ads](docs/sdks/ads/README.md)

* [CreateAdsAd](docs/sdks/ads/README.md#createadsad) - Create an ad
* [CreateAdsAsset](docs/sdks/ads/README.md#createadsasset) - Create an asset
* [CreateAdsCampaign](docs/sdks/ads/README.md#createadscampaign) - Create a campaign
* [CreateAdsCreative](docs/sdks/ads/README.md#createadscreative) - Create a creative
* [CreateAdsGroup](docs/sdks/ads/README.md#createadsgroup) - Create a group
* [CreateAdsInsertionorder](docs/sdks/ads/README.md#createadsinsertionorder) - Create an insertionorder
* [CreateAdsOrganization](docs/sdks/ads/README.md#createadsorganization) - Create an organization
* [GetAdsAd](docs/sdks/ads/README.md#getadsad) - Retrieve an ad
* [GetAdsAsset](docs/sdks/ads/README.md#getadsasset) - Retrieve an asset
* [GetAdsCampaign](docs/sdks/ads/README.md#getadscampaign) - Retrieve a campaign
* [GetAdsCreative](docs/sdks/ads/README.md#getadscreative) - Retrieve a creative
* [GetAdsGroup](docs/sdks/ads/README.md#getadsgroup) - Retrieve a group
* [GetAdsInsertionorder](docs/sdks/ads/README.md#getadsinsertionorder) - Retrieve an insertionorder
* [GetAdsOrganization](docs/sdks/ads/README.md#getadsorganization) - Retrieve an organization
* [GetAdsPromoted](docs/sdks/ads/README.md#getadspromoted) - Retrieve a promoted
* [GetAdsTarget](docs/sdks/ads/README.md#getadstarget) - Retrieve a target
* [ListAdsAds](docs/sdks/ads/README.md#listadsads) - List all ads
* [ListAdsAssets](docs/sdks/ads/README.md#listadsassets) - List all assets
* [ListAdsCampaigns](docs/sdks/ads/README.md#listadscampaigns) - List all campaigns
* [ListAdsCreatives](docs/sdks/ads/README.md#listadscreatives) - List all creatives
* [ListAdsGroups](docs/sdks/ads/README.md#listadsgroups) - List all groups
* [ListAdsInsertionorders](docs/sdks/ads/README.md#listadsinsertionorders) - List all insertionorders
* [ListAdsOrganizations](docs/sdks/ads/README.md#listadsorganizations) - List all organizations
* [ListAdsPromoteds](docs/sdks/ads/README.md#listadspromoteds) - List all promoteds
* [ListAdsReports](docs/sdks/ads/README.md#listadsreports) - List all reports
* [ListAdsTargets](docs/sdks/ads/README.md#listadstargets) - List all targets
* [PatchAdsAd](docs/sdks/ads/README.md#patchadsad) - Update an ad
* [PatchAdsCampaign](docs/sdks/ads/README.md#patchadscampaign) - Update a campaign
* [PatchAdsCreative](docs/sdks/ads/README.md#patchadscreative) - Update a creative
* [PatchAdsGroup](docs/sdks/ads/README.md#patchadsgroup) - Update a group
* [PatchAdsInsertionorder](docs/sdks/ads/README.md#patchadsinsertionorder) - Update an insertionorder
* [PatchAdsOrganization](docs/sdks/ads/README.md#patchadsorganization) - Update an organization
* [RemoveAdsAd](docs/sdks/ads/README.md#removeadsad) - Remove an ad
* [RemoveAdsCampaign](docs/sdks/ads/README.md#removeadscampaign) - Remove a campaign
* [RemoveAdsCreative](docs/sdks/ads/README.md#removeadscreative) - Remove a creative
* [RemoveAdsGroup](docs/sdks/ads/README.md#removeadsgroup) - Remove a group
* [RemoveAdsInsertionorder](docs/sdks/ads/README.md#removeadsinsertionorder) - Remove an insertionorder
* [RemoveAdsOrganization](docs/sdks/ads/README.md#removeadsorganization) - Remove an organization
* [UpdateAdsAd](docs/sdks/ads/README.md#updateadsad) - Update an ad
* [UpdateAdsCampaign](docs/sdks/ads/README.md#updateadscampaign) - Update a campaign
* [UpdateAdsCreative](docs/sdks/ads/README.md#updateadscreative) - Update a creative
* [UpdateAdsGroup](docs/sdks/ads/README.md#updateadsgroup) - Update a group
* [UpdateAdsInsertionorder](docs/sdks/ads/README.md#updateadsinsertionorder) - Update an insertionorder
* [UpdateAdsOrganization](docs/sdks/ads/README.md#updateadsorganization) - Update an organization

### [Agedpayable](docs/sdks/agedpayable/README.md)

* [GetAccountingAgedpayable](docs/sdks/agedpayable/README.md#getaccountingagedpayable) - Retrieve an agedpayable
* [ListAccountingAgedpayables](docs/sdks/agedpayable/README.md#listaccountingagedpayables) - List all agedpayables

### [Agedreceivable](docs/sdks/agedreceivable/README.md)

* [GetAccountingAgedreceivable](docs/sdks/agedreceivable/README.md#getaccountingagedreceivable) - Retrieve an agedreceivable
* [ListAccountingAgedreceivables](docs/sdks/agedreceivable/README.md#listaccountingagedreceivables) - List all agedreceivables

### [Analytics](docs/sdks/analytics/README.md)

* [CreateAnalyticsEvent](docs/sdks/analytics/README.md#createanalyticsevent) - Create an event
* [CreateAnalyticsProperty](docs/sdks/analytics/README.md#createanalyticsproperty) - Create a property
* [CreateAnalyticsVisitor](docs/sdks/analytics/README.md#createanalyticsvisitor) - Create a visitor
* [GetAnalyticsEvent](docs/sdks/analytics/README.md#getanalyticsevent) - Retrieve an event
* [GetAnalyticsProperty](docs/sdks/analytics/README.md#getanalyticsproperty) - Retrieve a property
* [GetAnalyticsSession](docs/sdks/analytics/README.md#getanalyticssession) - Retrieve a session
* [GetAnalyticsVisitor](docs/sdks/analytics/README.md#getanalyticsvisitor) - Retrieve a visitor
* [ListAnalyticsEvents](docs/sdks/analytics/README.md#listanalyticsevents) - List all events
* [ListAnalyticsProperties](docs/sdks/analytics/README.md#listanalyticsproperties) - List all properties
* [ListAnalyticsReports](docs/sdks/analytics/README.md#listanalyticsreports) - List all reports
* [ListAnalyticsSessions](docs/sdks/analytics/README.md#listanalyticssessions) - List all sessions
* [ListAnalyticsVisitors](docs/sdks/analytics/README.md#listanalyticsvisitors) - List all visitors
* [PatchAnalyticsProperty](docs/sdks/analytics/README.md#patchanalyticsproperty) - Update a property
* [PatchAnalyticsVisitor](docs/sdks/analytics/README.md#patchanalyticsvisitor) - Update a visitor
* [RemoveAnalyticsProperty](docs/sdks/analytics/README.md#removeanalyticsproperty) - Remove a property
* [RemoveAnalyticsVisitor](docs/sdks/analytics/README.md#removeanalyticsvisitor) - Remove a visitor
* [UpdateAnalyticsProperty](docs/sdks/analytics/README.md#updateanalyticsproperty) - Update a property
* [UpdateAnalyticsVisitor](docs/sdks/analytics/README.md#updateanalyticsvisitor) - Update a visitor

### [Apicall](docs/sdks/apicall/README.md)

* [GetUnifiedApicall](docs/sdks/apicall/README.md#getunifiedapicall) - Retrieve specific API Call by its ID
* [ListUnifiedApicalls](docs/sdks/apicall/README.md#listunifiedapicalls) - Returns API Calls

### [Application](docs/sdks/application/README.md)

* [CreateAtsApplication](docs/sdks/application/README.md#createatsapplication) - Create an application
* [GetAtsApplication](docs/sdks/application/README.md#getatsapplication) - Retrieve an application
* [ListAtsApplications](docs/sdks/application/README.md#listatsapplications) - List all applications
* [PatchAtsApplication](docs/sdks/application/README.md#patchatsapplication) - Update an application
* [RemoveAtsApplication](docs/sdks/application/README.md#removeatsapplication) - Remove an application
* [UpdateAtsApplication](docs/sdks/application/README.md#updateatsapplication) - Update an application

### [Applicationstatus](docs/sdks/applicationstatus/README.md)

* [ListAtsApplicationstatuses](docs/sdks/applicationstatus/README.md#listatsapplicationstatuses) - List all applicationstatuses

### [Assessment](docs/sdks/assessment/README.md)

* [CreateAssessmentPackage](docs/sdks/assessment/README.md#createassessmentpackage) - Create an assessment package
* [GetAssessmentPackage](docs/sdks/assessment/README.md#getassessmentpackage) - Get an assessment package
* [ListAssessmentPackages](docs/sdks/assessment/README.md#listassessmentpackages) - List assessment packages
* [PatchAssessmentOrder](docs/sdks/assessment/README.md#patchassessmentorder) - Update an order
* [PatchAssessmentPackage](docs/sdks/assessment/README.md#patchassessmentpackage) - Update an assessment package
* [RemoveAssessmentPackage](docs/sdks/assessment/README.md#removeassessmentpackage) - Delete an assessment package
* [UpdateAssessmentOrder](docs/sdks/assessment/README.md#updateassessmentorder) - Update an order
* [UpdateAssessmentPackage](docs/sdks/assessment/README.md#updateassessmentpackage) - Update an assessment package

### [Asset](docs/sdks/asset/README.md)

* [CreateAdsAsset](docs/sdks/asset/README.md#createadsasset) - Create an asset
* [GetAdsAsset](docs/sdks/asset/README.md#getadsasset) - Retrieve an asset
* [ListAdsAssets](docs/sdks/asset/README.md#listadsassets) - List all assets

### [Ats](docs/sdks/ats/README.md)

* [CreateAtsActivity](docs/sdks/ats/README.md#createatsactivity) - Create an activity
* [CreateAtsApplication](docs/sdks/ats/README.md#createatsapplication) - Create an application
* [CreateAtsCandidate](docs/sdks/ats/README.md#createatscandidate) - Create a candidate
* [CreateAtsCompany](docs/sdks/ats/README.md#createatscompany) - Create a company
* [CreateAtsDocument](docs/sdks/ats/README.md#createatsdocument) - Create a document
* [CreateAtsInterview](docs/sdks/ats/README.md#createatsinterview) - Create an interview
* [CreateAtsJob](docs/sdks/ats/README.md#createatsjob) - Create a job
* [CreateAtsScorecard](docs/sdks/ats/README.md#createatsscorecard) - Create a scorecard
* [GetAtsActivity](docs/sdks/ats/README.md#getatsactivity) - Retrieve an activity
* [GetAtsApplication](docs/sdks/ats/README.md#getatsapplication) - Retrieve an application
* [GetAtsCandidate](docs/sdks/ats/README.md#getatscandidate) - Retrieve a candidate
* [GetAtsCompany](docs/sdks/ats/README.md#getatscompany) - Retrieve a company
* [GetAtsDocument](docs/sdks/ats/README.md#getatsdocument) - Retrieve a document
* [GetAtsInterview](docs/sdks/ats/README.md#getatsinterview) - Retrieve an interview
* [GetAtsJob](docs/sdks/ats/README.md#getatsjob) - Retrieve a job
* [GetAtsScorecard](docs/sdks/ats/README.md#getatsscorecard) - Retrieve a scorecard
* [ListAtsActivities](docs/sdks/ats/README.md#listatsactivities) - List all activities
* [ListAtsApplications](docs/sdks/ats/README.md#listatsapplications) - List all applications
* [ListAtsApplicationstatuses](docs/sdks/ats/README.md#listatsapplicationstatuses) - List all applicationstatuses
* [ListAtsCandidates](docs/sdks/ats/README.md#listatscandidates) - List all candidates
* [ListAtsCompanies](docs/sdks/ats/README.md#listatscompanies) - List all companies
* [ListAtsDocuments](docs/sdks/ats/README.md#listatsdocuments) - List all documents
* [ListAtsInterviews](docs/sdks/ats/README.md#listatsinterviews) - List all interviews
* [ListAtsJobs](docs/sdks/ats/README.md#listatsjobs) - List all jobs
* [ListAtsScorecards](docs/sdks/ats/README.md#listatsscorecards) - List all scorecards
* [PatchAtsActivity](docs/sdks/ats/README.md#patchatsactivity) - Update an activity
* [PatchAtsApplication](docs/sdks/ats/README.md#patchatsapplication) - Update an application
* [PatchAtsCandidate](docs/sdks/ats/README.md#patchatscandidate) - Update a candidate
* [PatchAtsCompany](docs/sdks/ats/README.md#patchatscompany) - Update a company
* [PatchAtsDocument](docs/sdks/ats/README.md#patchatsdocument) - Update a document
* [PatchAtsInterview](docs/sdks/ats/README.md#patchatsinterview) - Update an interview
* [PatchAtsJob](docs/sdks/ats/README.md#patchatsjob) - Update a job
* [PatchAtsScorecard](docs/sdks/ats/README.md#patchatsscorecard) - Update a scorecard
* [RemoveAtsActivity](docs/sdks/ats/README.md#removeatsactivity) - Remove an activity
* [RemoveAtsApplication](docs/sdks/ats/README.md#removeatsapplication) - Remove an application
* [RemoveAtsCandidate](docs/sdks/ats/README.md#removeatscandidate) - Remove a candidate
* [RemoveAtsCompany](docs/sdks/ats/README.md#removeatscompany) - Remove a company
* [RemoveAtsDocument](docs/sdks/ats/README.md#removeatsdocument) - Remove a document
* [RemoveAtsInterview](docs/sdks/ats/README.md#removeatsinterview) - Remove an interview
* [RemoveAtsJob](docs/sdks/ats/README.md#removeatsjob) - Remove a job
* [RemoveAtsScorecard](docs/sdks/ats/README.md#removeatsscorecard) - Remove a scorecard
* [UpdateAtsActivity](docs/sdks/ats/README.md#updateatsactivity) - Update an activity
* [UpdateAtsApplication](docs/sdks/ats/README.md#updateatsapplication) - Update an application
* [UpdateAtsCandidate](docs/sdks/ats/README.md#updateatscandidate) - Update a candidate
* [UpdateAtsCompany](docs/sdks/ats/README.md#updateatscompany) - Update a company
* [UpdateAtsDocument](docs/sdks/ats/README.md#updateatsdocument) - Update a document
* [UpdateAtsInterview](docs/sdks/ats/README.md#updateatsinterview) - Update an interview
* [UpdateAtsJob](docs/sdks/ats/README.md#updateatsjob) - Update a job
* [UpdateAtsScorecard](docs/sdks/ats/README.md#updateatsscorecard) - Update a scorecard

### [Auth](docs/sdks/auth/README.md)

* [GetUnifiedIntegrationAuth](docs/sdks/auth/README.md#getunifiedintegrationauth) - Authorize new connection
* [GetUnifiedIntegrationLogin](docs/sdks/auth/README.md#getunifiedintegrationlogin) - Sign in a user
* [GetUnifiedIntegrationSaml](docs/sdks/auth/README.md#getunifiedintegrationsaml) - Sign in a user via SAML

### [Availability](docs/sdks/availability/README.md)

* [ListCommerceAvailabilities](docs/sdks/availability/README.md#listcommerceavailabilities) - List all availabilities

### [Balancesheet](docs/sdks/balancesheet/README.md)

* [GetAccountingBalancesheet](docs/sdks/balancesheet/README.md#getaccountingbalancesheet) - Retrieve a balancesheet
* [ListAccountingBalancesheets](docs/sdks/balancesheet/README.md#listaccountingbalancesheets) - List all balancesheets

### [Bankaccount](docs/sdks/bankaccount/README.md)

* [CreateHrisBankaccount](docs/sdks/bankaccount/README.md#createhrisbankaccount) - Create a bankaccount
* [GetHrisBankaccount](docs/sdks/bankaccount/README.md#gethrisbankaccount) - Retrieve a bankaccount
* [ListHrisBankaccounts](docs/sdks/bankaccount/README.md#listhrisbankaccounts) - List all bankaccounts
* [PatchHrisBankaccount](docs/sdks/bankaccount/README.md#patchhrisbankaccount) - Update a bankaccount
* [RemoveHrisBankaccount](docs/sdks/bankaccount/README.md#removehrisbankaccount) - Remove a bankaccount
* [UpdateHrisBankaccount](docs/sdks/bankaccount/README.md#updatehrisbankaccount) - Update a bankaccount

### [Benefit](docs/sdks/benefit/README.md)

* [CreateHrisBenefit](docs/sdks/benefit/README.md#createhrisbenefit) - Create a benefit
* [GetHrisBenefit](docs/sdks/benefit/README.md#gethrisbenefit) - Retrieve a benefit
* [ListHrisBenefits](docs/sdks/benefit/README.md#listhrisbenefits) - List all benefits
* [PatchHrisBenefit](docs/sdks/benefit/README.md#patchhrisbenefit) - Update a benefit
* [RemoveHrisBenefit](docs/sdks/benefit/README.md#removehrisbenefit) - Remove a benefit
* [UpdateHrisBenefit](docs/sdks/benefit/README.md#updatehrisbenefit) - Update a benefit

### [Bill](docs/sdks/bill/README.md)

* [CreateAccountingBill](docs/sdks/bill/README.md#createaccountingbill) - Create a bill
* [GetAccountingBill](docs/sdks/bill/README.md#getaccountingbill) - Retrieve a bill
* [ListAccountingBills](docs/sdks/bill/README.md#listaccountingbills) - List all bills
* [PatchAccountingBill](docs/sdks/bill/README.md#patchaccountingbill) - Update a bill
* [RemoveAccountingBill](docs/sdks/bill/README.md#removeaccountingbill) - Remove a bill
* [UpdateAccountingBill](docs/sdks/bill/README.md#updateaccountingbill) - Update a bill

### [Branch](docs/sdks/branch/README.md)

* [CreateRepoBranch](docs/sdks/branch/README.md#createrepobranch) - Create a branch
* [GetRepoBranch](docs/sdks/branch/README.md#getrepobranch) - Retrieve a branch
* [ListRepoBranches](docs/sdks/branch/README.md#listrepobranches) - List all branches
* [PatchRepoBranch](docs/sdks/branch/README.md#patchrepobranch) - Update a branch
* [RemoveRepoBranch](docs/sdks/branch/README.md#removerepobranch) - Remove a branch
* [UpdateRepoBranch](docs/sdks/branch/README.md#updaterepobranch) - Update a branch

### [Busy](docs/sdks/busy/README.md)

* [ListCalendarBusies](docs/sdks/busy/README.md#listcalendarbusies) - List all busies

### [Calendar](docs/sdks/calendar/README.md)

* [CreateCalendarCalendar](docs/sdks/calendar/README.md#createcalendarcalendar) - Create a calendar
* [CreateCalendarEvent](docs/sdks/calendar/README.md#createcalendarevent) - Create an event
* [CreateCalendarLink](docs/sdks/calendar/README.md#createcalendarlink) - Create a link
* [CreateCalendarWebinar](docs/sdks/calendar/README.md#createcalendarwebinar) - Create a webinar
* [GetCalendarCalendar](docs/sdks/calendar/README.md#getcalendarcalendar) - Retrieve a calendar
* [GetCalendarEvent](docs/sdks/calendar/README.md#getcalendarevent) - Retrieve an event
* [GetCalendarLink](docs/sdks/calendar/README.md#getcalendarlink) - Retrieve a link
* [GetCalendarRecording](docs/sdks/calendar/README.md#getcalendarrecording) - Retrieve a recording
* [GetCalendarWebinar](docs/sdks/calendar/README.md#getcalendarwebinar) - Retrieve a webinar
* [ListCalendarBusies](docs/sdks/calendar/README.md#listcalendarbusies) - List all busies
* [ListCalendarCalendars](docs/sdks/calendar/README.md#listcalendarcalendars) - List all calendars
* [ListCalendarEvents](docs/sdks/calendar/README.md#listcalendarevents) - List all events
* [ListCalendarLinks](docs/sdks/calendar/README.md#listcalendarlinks) - List all links
* [ListCalendarRecordings](docs/sdks/calendar/README.md#listcalendarrecordings) - List all recordings
* [ListCalendarWebinars](docs/sdks/calendar/README.md#listcalendarwebinars) - List all webinars
* [PatchCalendarCalendar](docs/sdks/calendar/README.md#patchcalendarcalendar) - Update a calendar
* [PatchCalendarEvent](docs/sdks/calendar/README.md#patchcalendarevent) - Update an event
* [PatchCalendarLink](docs/sdks/calendar/README.md#patchcalendarlink) - Update a link
* [PatchCalendarWebinar](docs/sdks/calendar/README.md#patchcalendarwebinar) - Update a webinar
* [RemoveCalendarCalendar](docs/sdks/calendar/README.md#removecalendarcalendar) - Remove a calendar
* [RemoveCalendarEvent](docs/sdks/calendar/README.md#removecalendarevent) - Remove an event
* [RemoveCalendarLink](docs/sdks/calendar/README.md#removecalendarlink) - Remove a link
* [RemoveCalendarWebinar](docs/sdks/calendar/README.md#removecalendarwebinar) - Remove a webinar
* [UpdateCalendarCalendar](docs/sdks/calendar/README.md#updatecalendarcalendar) - Update a calendar
* [UpdateCalendarEvent](docs/sdks/calendar/README.md#updatecalendarevent) - Update an event
* [UpdateCalendarLink](docs/sdks/calendar/README.md#updatecalendarlink) - Update a link
* [UpdateCalendarWebinar](docs/sdks/calendar/README.md#updatecalendarwebinar) - Update a webinar

### [Call](docs/sdks/call/README.md)

* [GetUcCall](docs/sdks/call/README.md#getuccall) - Retrieve a call
* [ListUcCalls](docs/sdks/call/README.md#listuccalls) - List all calls

### [Campaign](docs/sdks/campaign/README.md)

* [CreateAdsCampaign](docs/sdks/campaign/README.md#createadscampaign) - Create a campaign
* [CreateMartechCampaign](docs/sdks/campaign/README.md#createmartechcampaign) - Create a campaign
* [GetAdsCampaign](docs/sdks/campaign/README.md#getadscampaign) - Retrieve a campaign
* [GetMartechCampaign](docs/sdks/campaign/README.md#getmartechcampaign) - Retrieve a campaign
* [ListAdsCampaigns](docs/sdks/campaign/README.md#listadscampaigns) - List all campaigns
* [ListMartechCampaigns](docs/sdks/campaign/README.md#listmartechcampaigns) - List all campaigns
* [PatchAdsCampaign](docs/sdks/campaign/README.md#patchadscampaign) - Update a campaign
* [PatchMartechCampaign](docs/sdks/campaign/README.md#patchmartechcampaign) - Update a campaign
* [RemoveAdsCampaign](docs/sdks/campaign/README.md#removeadscampaign) - Remove a campaign
* [RemoveMartechCampaign](docs/sdks/campaign/README.md#removemartechcampaign) - Remove a campaign
* [UpdateAdsCampaign](docs/sdks/campaign/README.md#updateadscampaign) - Update a campaign
* [UpdateMartechCampaign](docs/sdks/campaign/README.md#updatemartechcampaign) - Update a campaign

### [Candidate](docs/sdks/candidate/README.md)

* [CreateAtsCandidate](docs/sdks/candidate/README.md#createatscandidate) - Create a candidate
* [GetAtsCandidate](docs/sdks/candidate/README.md#getatscandidate) - Retrieve a candidate
* [ListAtsCandidates](docs/sdks/candidate/README.md#listatscandidates) - List all candidates
* [PatchAtsCandidate](docs/sdks/candidate/README.md#patchatscandidate) - Update a candidate
* [RemoveAtsCandidate](docs/sdks/candidate/README.md#removeatscandidate) - Remove a candidate
* [UpdateAtsCandidate](docs/sdks/candidate/README.md#updateatscandidate) - Update a candidate

### [Carrier](docs/sdks/carrier/README.md)

* [GetShippingCarrier](docs/sdks/carrier/README.md#getshippingcarrier) - Retrieve a carrier
* [ListShippingCarriers](docs/sdks/carrier/README.md#listshippingcarriers) - List all carriers

### [Cashflow](docs/sdks/cashflow/README.md)

* [GetAccountingCashflow](docs/sdks/cashflow/README.md#getaccountingcashflow) - Retrieve a cashflow
* [ListAccountingCashflows](docs/sdks/cashflow/README.md#listaccountingcashflows) - List all cashflows

### [Category](docs/sdks/category/README.md)

* [CreateAccountingCategory](docs/sdks/category/README.md#createaccountingcategory) - Create a category
* [CreateTicketingCategory](docs/sdks/category/README.md#createticketingcategory) - Create a category
* [GetAccountingCategory](docs/sdks/category/README.md#getaccountingcategory) - Retrieve a category
* [GetTicketingCategory](docs/sdks/category/README.md#getticketingcategory) - Retrieve a category
* [ListAccountingCategories](docs/sdks/category/README.md#listaccountingcategories) - List all categories
* [ListTicketingCategories](docs/sdks/category/README.md#listticketingcategories) - List all categories
* [PatchAccountingCategory](docs/sdks/category/README.md#patchaccountingcategory) - Update a category
* [PatchTicketingCategory](docs/sdks/category/README.md#patchticketingcategory) - Update a category
* [RemoveAccountingCategory](docs/sdks/category/README.md#removeaccountingcategory) - Remove a category
* [RemoveTicketingCategory](docs/sdks/category/README.md#removeticketingcategory) - Remove a category
* [UpdateAccountingCategory](docs/sdks/category/README.md#updateaccountingcategory) - Update a category
* [UpdateTicketingCategory](docs/sdks/category/README.md#updateticketingcategory) - Update a category

### [Cdp](docs/sdks/cdp/README.md)

* [CreateCdpActivation](docs/sdks/cdp/README.md#createcdpactivation) - Create an activation
* [CreateCdpDestination](docs/sdks/cdp/README.md#createcdpdestination) - Create a destination
* [CreateCdpEvent](docs/sdks/cdp/README.md#createcdpevent) - Create an event
* [CreateCdpProfile](docs/sdks/cdp/README.md#createcdpprofile) - Create a profile
* [CreateCdpSegment](docs/sdks/cdp/README.md#createcdpsegment) - Create a segment
* [CreateCdpSource](docs/sdks/cdp/README.md#createcdpsource) - Create a source
* [GetCdpActivation](docs/sdks/cdp/README.md#getcdpactivation) - Retrieve an activation
* [GetCdpDestination](docs/sdks/cdp/README.md#getcdpdestination) - Retrieve a destination
* [GetCdpEvent](docs/sdks/cdp/README.md#getcdpevent) - Retrieve an event
* [GetCdpProfile](docs/sdks/cdp/README.md#getcdpprofile) - Retrieve a profile
* [GetCdpSegment](docs/sdks/cdp/README.md#getcdpsegment) - Retrieve a segment
* [GetCdpSource](docs/sdks/cdp/README.md#getcdpsource) - Retrieve a source
* [ListCdpActivations](docs/sdks/cdp/README.md#listcdpactivations) - List all activations
* [ListCdpDestinations](docs/sdks/cdp/README.md#listcdpdestinations) - List all destinations
* [ListCdpEvents](docs/sdks/cdp/README.md#listcdpevents) - List all events
* [ListCdpProfiles](docs/sdks/cdp/README.md#listcdpprofiles) - List all profiles
* [ListCdpSegments](docs/sdks/cdp/README.md#listcdpsegments) - List all segments
* [ListCdpSources](docs/sdks/cdp/README.md#listcdpsources) - List all sources
* [PatchCdpActivation](docs/sdks/cdp/README.md#patchcdpactivation) - Update an activation
* [PatchCdpDestination](docs/sdks/cdp/README.md#patchcdpdestination) - Update a destination
* [PatchCdpEvent](docs/sdks/cdp/README.md#patchcdpevent) - Update an event
* [PatchCdpProfile](docs/sdks/cdp/README.md#patchcdpprofile) - Update a profile
* [PatchCdpSegment](docs/sdks/cdp/README.md#patchcdpsegment) - Update a segment
* [PatchCdpSource](docs/sdks/cdp/README.md#patchcdpsource) - Update a source
* [RemoveCdpActivation](docs/sdks/cdp/README.md#removecdpactivation) - Remove an activation
* [RemoveCdpDestination](docs/sdks/cdp/README.md#removecdpdestination) - Remove a destination
* [RemoveCdpEvent](docs/sdks/cdp/README.md#removecdpevent) - Remove an event
* [RemoveCdpProfile](docs/sdks/cdp/README.md#removecdpprofile) - Remove a profile
* [RemoveCdpSegment](docs/sdks/cdp/README.md#removecdpsegment) - Remove a segment
* [RemoveCdpSource](docs/sdks/cdp/README.md#removecdpsource) - Remove a source
* [UpdateCdpActivation](docs/sdks/cdp/README.md#updatecdpactivation) - Update an activation
* [UpdateCdpDestination](docs/sdks/cdp/README.md#updatecdpdestination) - Update a destination
* [UpdateCdpEvent](docs/sdks/cdp/README.md#updatecdpevent) - Update an event
* [UpdateCdpProfile](docs/sdks/cdp/README.md#updatecdpprofile) - Update a profile
* [UpdateCdpSegment](docs/sdks/cdp/README.md#updatecdpsegment) - Update a segment
* [UpdateCdpSource](docs/sdks/cdp/README.md#updatecdpsource) - Update a source

### [Change](docs/sdks/change/README.md)

* [GetTaskChange](docs/sdks/change/README.md#gettaskchange) - Retrieve a change
* [ListTaskChanges](docs/sdks/change/README.md#listtaskchanges) - List all changes

### [Channel](docs/sdks/channel/README.md)

* [CreateMessagingChannel](docs/sdks/channel/README.md#createmessagingchannel) - Create a channel
* [GetMessagingChannel](docs/sdks/channel/README.md#getmessagingchannel) - Retrieve a channel
* [ListMessagingChannels](docs/sdks/channel/README.md#listmessagingchannels) - List all channels
* [PatchMessagingChannel](docs/sdks/channel/README.md#patchmessagingchannel) - Update a channel
* [RemoveMessagingChannel](docs/sdks/channel/README.md#removemessagingchannel) - Remove a channel
* [UpdateMessagingChannel](docs/sdks/channel/README.md#updatemessagingchannel) - Update a channel

### [Class](docs/sdks/class/README.md)

* [CreateLmsClass](docs/sdks/class/README.md#createlmsclass) - Create a class
* [GetLmsClass](docs/sdks/class/README.md#getlmsclass) - Retrieve a class
* [ListLmsClasses](docs/sdks/class/README.md#listlmsclasses) - List all classes
* [PatchLmsClass](docs/sdks/class/README.md#patchlmsclass) - Update a class
* [RemoveLmsClass](docs/sdks/class/README.md#removelmsclass) - Remove a class
* [UpdateLmsClass](docs/sdks/class/README.md#updatelmsclass) - Update a class

### [Clubs](docs/sdks/clubs/README.md)

* [GetClubsActivity](docs/sdks/clubs/README.md#getclubsactivity) - Retrieve an activity
* [GetClubsEvent](docs/sdks/clubs/README.md#getclubsevent) - Retrieve an event
* [GetClubsGroup](docs/sdks/clubs/README.md#getclubsgroup) - Retrieve a group
* [GetClubsLocation](docs/sdks/clubs/README.md#getclubslocation) - Retrieve a location
* [GetClubsMember](docs/sdks/clubs/README.md#getclubsmember) - Retrieve a member
* [ListClubsActivities](docs/sdks/clubs/README.md#listclubsactivities) - List all activities
* [ListClubsEvents](docs/sdks/clubs/README.md#listclubsevents) - List all events
* [ListClubsGroups](docs/sdks/clubs/README.md#listclubsgroups) - List all groups
* [ListClubsLocations](docs/sdks/clubs/README.md#listclubslocations) - List all locations
* [ListClubsMembers](docs/sdks/clubs/README.md#listclubsmembers) - List all members

### [Collection](docs/sdks/collection/README.md)

* [CreateCommerceCollection](docs/sdks/collection/README.md#createcommercecollection) - Create a collection
* [CreateLmsCollection](docs/sdks/collection/README.md#createlmscollection) - Create a collection
* [GetCommerceCollection](docs/sdks/collection/README.md#getcommercecollection) - Retrieve a collection
* [GetLmsCollection](docs/sdks/collection/README.md#getlmscollection) - Retrieve a collection
* [ListCommerceCollections](docs/sdks/collection/README.md#listcommercecollections) - List all collections
* [ListLmsCollections](docs/sdks/collection/README.md#listlmscollections) - List all collections
* [PatchCommerceCollection](docs/sdks/collection/README.md#patchcommercecollection) - Update a collection
* [PatchLmsCollection](docs/sdks/collection/README.md#patchlmscollection) - Update a collection
* [RemoveCommerceCollection](docs/sdks/collection/README.md#removecommercecollection) - Remove a collection
* [RemoveLmsCollection](docs/sdks/collection/README.md#removelmscollection) - Remove a collection
* [UpdateCommerceCollection](docs/sdks/collection/README.md#updatecommercecollection) - Update a collection
* [UpdateLmsCollection](docs/sdks/collection/README.md#updatelmscollection) - Update a collection

### [Comment](docs/sdks/comment/README.md)

* [CreateKmsComment](docs/sdks/comment/README.md#createkmscomment) - Create a comment
* [CreateTaskComment](docs/sdks/comment/README.md#createtaskcomment) - Create a comment
* [CreateUcComment](docs/sdks/comment/README.md#createuccomment) - Create a comment
* [GetKmsComment](docs/sdks/comment/README.md#getkmscomment) - Retrieve a comment
* [GetTaskComment](docs/sdks/comment/README.md#gettaskcomment) - Retrieve a comment
* [GetUcComment](docs/sdks/comment/README.md#getuccomment) - Retrieve a comment
* [ListKmsComments](docs/sdks/comment/README.md#listkmscomments) - List all comments
* [ListTaskComments](docs/sdks/comment/README.md#listtaskcomments) - List all comments
* [ListUcComments](docs/sdks/comment/README.md#listuccomments) - List all comments
* [PatchKmsComment](docs/sdks/comment/README.md#patchkmscomment) - Update a comment
* [PatchTaskComment](docs/sdks/comment/README.md#patchtaskcomment) - Update a comment
* [PatchUcComment](docs/sdks/comment/README.md#patchuccomment) - Update a comment
* [RemoveKmsComment](docs/sdks/comment/README.md#removekmscomment) - Remove a comment
* [RemoveTaskComment](docs/sdks/comment/README.md#removetaskcomment) - Remove a comment
* [RemoveUcComment](docs/sdks/comment/README.md#removeuccomment) - Remove a comment
* [UpdateKmsComment](docs/sdks/comment/README.md#updatekmscomment) - Update a comment
* [UpdateTaskComment](docs/sdks/comment/README.md#updatetaskcomment) - Update a comment
* [UpdateUcComment](docs/sdks/comment/README.md#updateuccomment) - Update a comment

### [Commerce](docs/sdks/commerce/README.md)

* [CreateCommerceCollection](docs/sdks/commerce/README.md#createcommercecollection) - Create a collection
* [CreateCommerceInventory](docs/sdks/commerce/README.md#createcommerceinventory) - Create an inventory
* [CreateCommerceItem](docs/sdks/commerce/README.md#createcommerceitem) - Create an item
* [CreateCommerceItemvariant](docs/sdks/commerce/README.md#createcommerceitemvariant) - Create an itemvariant
* [CreateCommerceLocation](docs/sdks/commerce/README.md#createcommercelocation) - Create a location
* [CreateCommerceReservation](docs/sdks/commerce/README.md#createcommercereservation) - Create a reservation
* [CreateCommerceReview](docs/sdks/commerce/README.md#createcommercereview) - Create a review
* [CreateCommerceSaleschannel](docs/sdks/commerce/README.md#createcommercesaleschannel) - Create a saleschannel
* [GetCommerceCollection](docs/sdks/commerce/README.md#getcommercecollection) - Retrieve a collection
* [GetCommerceInventory](docs/sdks/commerce/README.md#getcommerceinventory) - Retrieve an inventory
* [GetCommerceItem](docs/sdks/commerce/README.md#getcommerceitem) - Retrieve an item
* [GetCommerceItemvariant](docs/sdks/commerce/README.md#getcommerceitemvariant) - Retrieve an itemvariant
* [GetCommerceLocation](docs/sdks/commerce/README.md#getcommercelocation) - Retrieve a location
* [GetCommerceReservation](docs/sdks/commerce/README.md#getcommercereservation) - Retrieve a reservation
* [GetCommerceReview](docs/sdks/commerce/README.md#getcommercereview) - Retrieve a review
* [GetCommerceSaleschannel](docs/sdks/commerce/README.md#getcommercesaleschannel) - Retrieve a saleschannel
* [ListCommerceAvailabilities](docs/sdks/commerce/README.md#listcommerceavailabilities) - List all availabilities
* [ListCommerceCollections](docs/sdks/commerce/README.md#listcommercecollections) - List all collections
* [ListCommerceInventories](docs/sdks/commerce/README.md#listcommerceinventories) - List all inventories
* [ListCommerceItems](docs/sdks/commerce/README.md#listcommerceitems) - List all items
* [ListCommerceItemvariants](docs/sdks/commerce/README.md#listcommerceitemvariants) - List all itemvariants
* [ListCommerceLocations](docs/sdks/commerce/README.md#listcommercelocations) - List all locations
* [ListCommerceReservations](docs/sdks/commerce/README.md#listcommercereservations) - List all reservations
* [ListCommerceReviews](docs/sdks/commerce/README.md#listcommercereviews) - List all reviews
* [ListCommerceSaleschannels](docs/sdks/commerce/README.md#listcommercesaleschannels) - List all saleschannels
* [PatchCommerceCollection](docs/sdks/commerce/README.md#patchcommercecollection) - Update a collection
* [PatchCommerceInventory](docs/sdks/commerce/README.md#patchcommerceinventory) - Update an inventory
* [PatchCommerceItem](docs/sdks/commerce/README.md#patchcommerceitem) - Update an item
* [PatchCommerceItemvariant](docs/sdks/commerce/README.md#patchcommerceitemvariant) - Update an itemvariant
* [PatchCommerceLocation](docs/sdks/commerce/README.md#patchcommercelocation) - Update a location
* [PatchCommerceReservation](docs/sdks/commerce/README.md#patchcommercereservation) - Update a reservation
* [PatchCommerceReview](docs/sdks/commerce/README.md#patchcommercereview) - Update a review
* [PatchCommerceSaleschannel](docs/sdks/commerce/README.md#patchcommercesaleschannel) - Update a saleschannel
* [RemoveCommerceCollection](docs/sdks/commerce/README.md#removecommercecollection) - Remove a collection
* [RemoveCommerceInventory](docs/sdks/commerce/README.md#removecommerceinventory) - Remove an inventory
* [RemoveCommerceItem](docs/sdks/commerce/README.md#removecommerceitem) - Remove an item
* [RemoveCommerceItemvariant](docs/sdks/commerce/README.md#removecommerceitemvariant) - Remove an itemvariant
* [RemoveCommerceLocation](docs/sdks/commerce/README.md#removecommercelocation) - Remove a location
* [RemoveCommerceReservation](docs/sdks/commerce/README.md#removecommercereservation) - Remove a reservation
* [RemoveCommerceReview](docs/sdks/commerce/README.md#removecommercereview) - Remove a review
* [RemoveCommerceSaleschannel](docs/sdks/commerce/README.md#removecommercesaleschannel) - Remove a saleschannel
* [UpdateCommerceCollection](docs/sdks/commerce/README.md#updatecommercecollection) - Update a collection
* [UpdateCommerceInventory](docs/sdks/commerce/README.md#updatecommerceinventory) - Update an inventory
* [UpdateCommerceItem](docs/sdks/commerce/README.md#updatecommerceitem) - Update an item
* [UpdateCommerceItemvariant](docs/sdks/commerce/README.md#updatecommerceitemvariant) - Update an itemvariant
* [UpdateCommerceLocation](docs/sdks/commerce/README.md#updatecommercelocation) - Update a location
* [UpdateCommerceReservation](docs/sdks/commerce/README.md#updatecommercereservation) - Update a reservation
* [UpdateCommerceReview](docs/sdks/commerce/README.md#updatecommercereview) - Update a review
* [UpdateCommerceSaleschannel](docs/sdks/commerce/README.md#updatecommercesaleschannel) - Update a saleschannel

### [Commit](docs/sdks/commit/README.md)

* [CreateRepoCommit](docs/sdks/commit/README.md#createrepocommit) - Create a commit
* [GetRepoCommit](docs/sdks/commit/README.md#getrepocommit) - Retrieve a commit
* [ListRepoCommits](docs/sdks/commit/README.md#listrepocommits) - List all commits
* [PatchRepoCommit](docs/sdks/commit/README.md#patchrepocommit) - Update a commit
* [RemoveRepoCommit](docs/sdks/commit/README.md#removerepocommit) - Remove a commit
* [UpdateRepoCommit](docs/sdks/commit/README.md#updaterepocommit) - Update a commit

### [Company](docs/sdks/company/README.md)

* [CreateAtsCompany](docs/sdks/company/README.md#createatscompany) - Create a company
* [CreateCrmCompany](docs/sdks/company/README.md#createcrmcompany) - Create a company
* [CreateHrisCompany](docs/sdks/company/README.md#createhriscompany) - Create a company
* [GetAtsCompany](docs/sdks/company/README.md#getatscompany) - Retrieve a company
* [GetCrmCompany](docs/sdks/company/README.md#getcrmcompany) - Retrieve a company
* [GetHrisCompany](docs/sdks/company/README.md#gethriscompany) - Retrieve a company
* [ListAtsCompanies](docs/sdks/company/README.md#listatscompanies) - List all companies
* [ListCrmCompanies](docs/sdks/company/README.md#listcrmcompanies) - List all companies
* [ListEnrichCompanies](docs/sdks/company/README.md#listenrichcompanies) - Retrieve enrichment information for a company
* [ListHrisCompanies](docs/sdks/company/README.md#listhriscompanies) - List all companies
* [PatchAtsCompany](docs/sdks/company/README.md#patchatscompany) - Update a company
* [PatchCrmCompany](docs/sdks/company/README.md#patchcrmcompany) - Update a company
* [PatchHrisCompany](docs/sdks/company/README.md#patchhriscompany) - Update a company
* [RemoveAtsCompany](docs/sdks/company/README.md#removeatscompany) - Remove a company
* [RemoveCrmCompany](docs/sdks/company/README.md#removecrmcompany) - Remove a company
* [RemoveHrisCompany](docs/sdks/company/README.md#removehriscompany) - Remove a company
* [UpdateAtsCompany](docs/sdks/company/README.md#updateatscompany) - Update a company
* [UpdateCrmCompany](docs/sdks/company/README.md#updatecrmcompany) - Update a company
* [UpdateHrisCompany](docs/sdks/company/README.md#updatehriscompany) - Update a company

### [Connection](docs/sdks/connection/README.md)

* [CreateUnifiedConnection](docs/sdks/connection/README.md#createunifiedconnection) - Create connection
* [GetUnifiedConnection](docs/sdks/connection/README.md#getunifiedconnection) - Retrieve connection
* [ListUnifiedConnections](docs/sdks/connection/README.md#listunifiedconnections) - List all connections
* [PatchUnifiedConnection](docs/sdks/connection/README.md#patchunifiedconnection) - Update connection
* [RemoveUnifiedConnection](docs/sdks/connection/README.md#removeunifiedconnection) - Remove connection
* [UpdateUnifiedConnection](docs/sdks/connection/README.md#updateunifiedconnection) - Update connection

### [Contact](docs/sdks/contact/README.md)

* [CreateAccountingContact](docs/sdks/contact/README.md#createaccountingcontact) - Create a contact
* [CreateCrmContact](docs/sdks/contact/README.md#createcrmcontact) - Create a contact
* [CreateUcContact](docs/sdks/contact/README.md#createuccontact) - Create a contact
* [GetAccountingContact](docs/sdks/contact/README.md#getaccountingcontact) - Retrieve a contact
* [GetCrmContact](docs/sdks/contact/README.md#getcrmcontact) - Retrieve a contact
* [GetUcContact](docs/sdks/contact/README.md#getuccontact) - Retrieve a contact
* [ListAccountingContacts](docs/sdks/contact/README.md#listaccountingcontacts) - List all contacts
* [ListCrmContacts](docs/sdks/contact/README.md#listcrmcontacts) - List all contacts
* [ListUcContacts](docs/sdks/contact/README.md#listuccontacts) - List all contacts
* [PatchAccountingContact](docs/sdks/contact/README.md#patchaccountingcontact) - Update a contact
* [PatchCrmContact](docs/sdks/contact/README.md#patchcrmcontact) - Update a contact
* [PatchUcContact](docs/sdks/contact/README.md#patchuccontact) - Update a contact
* [RemoveAccountingContact](docs/sdks/contact/README.md#removeaccountingcontact) - Remove a contact
* [RemoveCrmContact](docs/sdks/contact/README.md#removecrmcontact) - Remove a contact
* [RemoveUcContact](docs/sdks/contact/README.md#removeuccontact) - Remove a contact
* [UpdateAccountingContact](docs/sdks/contact/README.md#updateaccountingcontact) - Update a contact
* [UpdateCrmContact](docs/sdks/contact/README.md#updatecrmcontact) - Update a contact
* [UpdateUcContact](docs/sdks/contact/README.md#updateuccontact) - Update a contact

### [Content](docs/sdks/content/README.md)

* [CreateLmsContent](docs/sdks/content/README.md#createlmscontent) - Create a content
* [GetLmsContent](docs/sdks/content/README.md#getlmscontent) - Retrieve a content
* [ListLmsContents](docs/sdks/content/README.md#listlmscontents) - List all contents
* [PatchLmsContent](docs/sdks/content/README.md#patchlmscontent) - Update a content
* [RemoveLmsContent](docs/sdks/content/README.md#removelmscontent) - Remove a content
* [UpdateLmsContent](docs/sdks/content/README.md#updatelmscontent) - Update a content

### [Course](docs/sdks/course/README.md)

* [CreateLmsCourse](docs/sdks/course/README.md#createlmscourse) - Create a course
* [GetLmsCourse](docs/sdks/course/README.md#getlmscourse) - Retrieve a course
* [ListLmsCourses](docs/sdks/course/README.md#listlmscourses) - List all courses
* [PatchLmsCourse](docs/sdks/course/README.md#patchlmscourse) - Update a course
* [RemoveLmsCourse](docs/sdks/course/README.md#removelmscourse) - Remove a course
* [UpdateLmsCourse](docs/sdks/course/README.md#updatelmscourse) - Update a course

### [Creative](docs/sdks/creative/README.md)

* [CreateAdsCreative](docs/sdks/creative/README.md#createadscreative) - Create a creative
* [GetAdsCreative](docs/sdks/creative/README.md#getadscreative) - Retrieve a creative
* [ListAdsCreatives](docs/sdks/creative/README.md#listadscreatives) - List all creatives
* [PatchAdsCreative](docs/sdks/creative/README.md#patchadscreative) - Update a creative
* [RemoveAdsCreative](docs/sdks/creative/README.md#removeadscreative) - Remove a creative
* [UpdateAdsCreative](docs/sdks/creative/README.md#updateadscreative) - Update a creative

### [Creditmemo](docs/sdks/creditmemo/README.md)

* [CreateAccountingCreditmemo](docs/sdks/creditmemo/README.md#createaccountingcreditmemo) - Create a creditmemo
* [GetAccountingCreditmemo](docs/sdks/creditmemo/README.md#getaccountingcreditmemo) - Retrieve a creditmemo
* [ListAccountingCreditmemoes](docs/sdks/creditmemo/README.md#listaccountingcreditmemoes) - List all creditmemoes
* [PatchAccountingCreditmemo](docs/sdks/creditmemo/README.md#patchaccountingcreditmemo) - Update a creditmemo
* [RemoveAccountingCreditmemo](docs/sdks/creditmemo/README.md#removeaccountingcreditmemo) - Remove a creditmemo
* [UpdateAccountingCreditmemo](docs/sdks/creditmemo/README.md#updateaccountingcreditmemo) - Update a creditmemo

### [Crm](docs/sdks/crm/README.md)

* [CreateCrmCompany](docs/sdks/crm/README.md#createcrmcompany) - Create a company
* [CreateCrmContact](docs/sdks/crm/README.md#createcrmcontact) - Create a contact
* [CreateCrmDeal](docs/sdks/crm/README.md#createcrmdeal) - Create a deal
* [CreateCrmEvent](docs/sdks/crm/README.md#createcrmevent) - Create an event
* [CreateCrmLead](docs/sdks/crm/README.md#createcrmlead) - Create a lead
* [CreateCrmPipeline](docs/sdks/crm/README.md#createcrmpipeline) - Create a pipeline
* [GetCrmCompany](docs/sdks/crm/README.md#getcrmcompany) - Retrieve a company
* [GetCrmContact](docs/sdks/crm/README.md#getcrmcontact) - Retrieve a contact
* [GetCrmDeal](docs/sdks/crm/README.md#getcrmdeal) - Retrieve a deal
* [GetCrmEvent](docs/sdks/crm/README.md#getcrmevent) - Retrieve an event
* [GetCrmLead](docs/sdks/crm/README.md#getcrmlead) - Retrieve a lead
* [GetCrmPipeline](docs/sdks/crm/README.md#getcrmpipeline) - Retrieve a pipeline
* [ListCrmCompanies](docs/sdks/crm/README.md#listcrmcompanies) - List all companies
* [ListCrmContacts](docs/sdks/crm/README.md#listcrmcontacts) - List all contacts
* [ListCrmDeals](docs/sdks/crm/README.md#listcrmdeals) - List all deals
* [ListCrmEvents](docs/sdks/crm/README.md#listcrmevents) - List all events
* [ListCrmLeads](docs/sdks/crm/README.md#listcrmleads) - List all leads
* [ListCrmPicklists](docs/sdks/crm/README.md#listcrmpicklists) - List all picklists
* [ListCrmPipelines](docs/sdks/crm/README.md#listcrmpipelines) - List all pipelines
* [PatchCrmCompany](docs/sdks/crm/README.md#patchcrmcompany) - Update a company
* [PatchCrmContact](docs/sdks/crm/README.md#patchcrmcontact) - Update a contact
* [PatchCrmDeal](docs/sdks/crm/README.md#patchcrmdeal) - Update a deal
* [PatchCrmEvent](docs/sdks/crm/README.md#patchcrmevent) - Update an event
* [PatchCrmLead](docs/sdks/crm/README.md#patchcrmlead) - Update a lead
* [PatchCrmPipeline](docs/sdks/crm/README.md#patchcrmpipeline) - Update a pipeline
* [RemoveCrmCompany](docs/sdks/crm/README.md#removecrmcompany) - Remove a company
* [RemoveCrmContact](docs/sdks/crm/README.md#removecrmcontact) - Remove a contact
* [RemoveCrmDeal](docs/sdks/crm/README.md#removecrmdeal) - Remove a deal
* [RemoveCrmEvent](docs/sdks/crm/README.md#removecrmevent) - Remove an event
* [RemoveCrmLead](docs/sdks/crm/README.md#removecrmlead) - Remove a lead
* [RemoveCrmPipeline](docs/sdks/crm/README.md#removecrmpipeline) - Remove a pipeline
* [UpdateCrmCompany](docs/sdks/crm/README.md#updatecrmcompany) - Update a company
* [UpdateCrmContact](docs/sdks/crm/README.md#updatecrmcontact) - Update a contact
* [UpdateCrmDeal](docs/sdks/crm/README.md#updatecrmdeal) - Update a deal
* [UpdateCrmEvent](docs/sdks/crm/README.md#updatecrmevent) - Update an event
* [UpdateCrmLead](docs/sdks/crm/README.md#updatecrmlead) - Update a lead
* [UpdateCrmPipeline](docs/sdks/crm/README.md#updatecrmpipeline) - Update a pipeline

### [Customer](docs/sdks/customer/README.md)

* [CreateTicketingCustomer](docs/sdks/customer/README.md#createticketingcustomer) - Create a customer
* [GetTicketingCustomer](docs/sdks/customer/README.md#getticketingcustomer) - Retrieve a customer
* [ListTicketingCustomers](docs/sdks/customer/README.md#listticketingcustomers) - List all customers
* [PatchTicketingCustomer](docs/sdks/customer/README.md#patchticketingcustomer) - Update a customer
* [RemoveTicketingCustomer](docs/sdks/customer/README.md#removeticketingcustomer) - Remove a customer
* [UpdateTicketingCustomer](docs/sdks/customer/README.md#updateticketingcustomer) - Update a customer

### [Cycle](docs/sdks/cycle/README.md)

* [GetPerformanceCycle](docs/sdks/cycle/README.md#getperformancecycle) - Retrieve a cycle
* [ListPerformanceCycles](docs/sdks/cycle/README.md#listperformancecycles) - List all cycles

### [Database](docs/sdks/database/README.md)

* [CreateDatastoreDatabase](docs/sdks/database/README.md#createdatastoredatabase) - Create a database
* [GetDatastoreDatabase](docs/sdks/database/README.md#getdatastoredatabase) - Retrieve a database
* [ListDatastoreDatabases](docs/sdks/database/README.md#listdatastoredatabases) - List all databases
* [PatchDatastoreDatabase](docs/sdks/database/README.md#patchdatastoredatabase) - Update a database
* [RemoveDatastoreDatabase](docs/sdks/database/README.md#removedatastoredatabase) - Remove a database
* [UpdateDatastoreDatabase](docs/sdks/database/README.md#updatedatastoredatabase) - Update a database

### [Datastore](docs/sdks/datastore/README.md)

* [CreateDatastoreDatabase](docs/sdks/datastore/README.md#createdatastoredatabase) - Create a database
* [CreateDatastoreQuery](docs/sdks/datastore/README.md#createdatastorequery) - Create a query
* [CreateDatastoreRecord](docs/sdks/datastore/README.md#createdatastorerecord) - Create a record
* [CreateDatastoreTable](docs/sdks/datastore/README.md#createdatastoretable) - Create a table
* [GetDatastoreDatabase](docs/sdks/datastore/README.md#getdatastoredatabase) - Retrieve a database
* [GetDatastoreRecord](docs/sdks/datastore/README.md#getdatastorerecord) - Retrieve a record
* [GetDatastoreTable](docs/sdks/datastore/README.md#getdatastoretable) - Retrieve a table
* [ListDatastoreDatabases](docs/sdks/datastore/README.md#listdatastoredatabases) - List all databases
* [ListDatastoreRecords](docs/sdks/datastore/README.md#listdatastorerecords) - List all records
* [ListDatastoreTables](docs/sdks/datastore/README.md#listdatastoretables) - List all tables
* [PatchDatastoreDatabase](docs/sdks/datastore/README.md#patchdatastoredatabase) - Update a database
* [PatchDatastoreRecord](docs/sdks/datastore/README.md#patchdatastorerecord) - Update a record
* [PatchDatastoreTable](docs/sdks/datastore/README.md#patchdatastoretable) - Update a table
* [RemoveDatastoreDatabase](docs/sdks/datastore/README.md#removedatastoredatabase) - Remove a database
* [RemoveDatastoreRecord](docs/sdks/datastore/README.md#removedatastorerecord) - Remove a record
* [RemoveDatastoreTable](docs/sdks/datastore/README.md#removedatastoretable) - Remove a table
* [UpdateDatastoreDatabase](docs/sdks/datastore/README.md#updatedatastoredatabase) - Update a database
* [UpdateDatastoreRecord](docs/sdks/datastore/README.md#updatedatastorerecord) - Update a record
* [UpdateDatastoreTable](docs/sdks/datastore/README.md#updatedatastoretable) - Update a table

### [Deal](docs/sdks/deal/README.md)

* [CreateCrmDeal](docs/sdks/deal/README.md#createcrmdeal) - Create a deal
* [GetCrmDeal](docs/sdks/deal/README.md#getcrmdeal) - Retrieve a deal
* [ListCrmDeals](docs/sdks/deal/README.md#listcrmdeals) - List all deals
* [PatchCrmDeal](docs/sdks/deal/README.md#patchcrmdeal) - Update a deal
* [RemoveCrmDeal](docs/sdks/deal/README.md#removecrmdeal) - Remove a deal
* [UpdateCrmDeal](docs/sdks/deal/README.md#updatecrmdeal) - Update a deal

### [Deduction](docs/sdks/deduction/README.md)

* [CreateHrisDeduction](docs/sdks/deduction/README.md#createhrisdeduction) - Create a deduction
* [GetHrisDeduction](docs/sdks/deduction/README.md#gethrisdeduction) - Retrieve a deduction
* [ListHrisDeductions](docs/sdks/deduction/README.md#listhrisdeductions) - List all deductions
* [PatchHrisDeduction](docs/sdks/deduction/README.md#patchhrisdeduction) - Update a deduction
* [RemoveHrisDeduction](docs/sdks/deduction/README.md#removehrisdeduction) - Remove a deduction
* [UpdateHrisDeduction](docs/sdks/deduction/README.md#updatehrisdeduction) - Update a deduction

### [Destination](docs/sdks/destination/README.md)

* [CreateCdpDestination](docs/sdks/destination/README.md#createcdpdestination) - Create a destination
* [GetCdpDestination](docs/sdks/destination/README.md#getcdpdestination) - Retrieve a destination
* [ListCdpDestinations](docs/sdks/destination/README.md#listcdpdestinations) - List all destinations
* [PatchCdpDestination](docs/sdks/destination/README.md#patchcdpdestination) - Update a destination
* [RemoveCdpDestination](docs/sdks/destination/README.md#removecdpdestination) - Remove a destination
* [UpdateCdpDestination](docs/sdks/destination/README.md#updatecdpdestination) - Update a destination

### [Device](docs/sdks/device/README.md)

* [CreateHrisDevice](docs/sdks/device/README.md#createhrisdevice) - Create a device
* [GetHrisDevice](docs/sdks/device/README.md#gethrisdevice) - Retrieve a device
* [ListHrisDevices](docs/sdks/device/README.md#listhrisdevices) - List all devices
* [PatchHrisDevice](docs/sdks/device/README.md#patchhrisdevice) - Update a device
* [RemoveHrisDevice](docs/sdks/device/README.md#removehrisdevice) - Remove a device
* [UpdateHrisDevice](docs/sdks/device/README.md#updatehrisdevice) - Update a device

### [Document](docs/sdks/document/README.md)

* [CreateAtsDocument](docs/sdks/document/README.md#createatsdocument) - Create a document
* [CreateHrisDocument](docs/sdks/document/README.md#createhrisdocument) - Create a document
* [CreateSigningDocument](docs/sdks/document/README.md#createsigningdocument) - Create a document
* [GetAtsDocument](docs/sdks/document/README.md#getatsdocument) - Retrieve a document
* [GetHrisDocument](docs/sdks/document/README.md#gethrisdocument) - Retrieve a document
* [GetSigningDocument](docs/sdks/document/README.md#getsigningdocument) - Retrieve a document
* [ListAtsDocuments](docs/sdks/document/README.md#listatsdocuments) - List all documents
* [ListHrisDocuments](docs/sdks/document/README.md#listhrisdocuments) - List all documents
* [ListSigningDocuments](docs/sdks/document/README.md#listsigningdocuments) - List all documents
* [PatchAtsDocument](docs/sdks/document/README.md#patchatsdocument) - Update a document
* [PatchHrisDocument](docs/sdks/document/README.md#patchhrisdocument) - Update a document
* [PatchSigningDocument](docs/sdks/document/README.md#patchsigningdocument) - Update a document
* [RemoveAtsDocument](docs/sdks/document/README.md#removeatsdocument) - Remove a document
* [RemoveHrisDocument](docs/sdks/document/README.md#removehrisdocument) - Remove a document
* [RemoveSigningDocument](docs/sdks/document/README.md#removesigningdocument) - Remove a document
* [UpdateAtsDocument](docs/sdks/document/README.md#updateatsdocument) - Update a document
* [UpdateHrisDocument](docs/sdks/document/README.md#updatehrisdocument) - Update a document
* [UpdateSigningDocument](docs/sdks/document/README.md#updatesigningdocument) - Update a document

### [Embedding](docs/sdks/embedding/README.md)

* [CreateGenaiEmbedding](docs/sdks/embedding/README.md#creategenaiembedding) - Create an embedding

### [Employee](docs/sdks/employee/README.md)

* [CreateHrisEmployee](docs/sdks/employee/README.md#createhrisemployee) - Create an employee
* [GetHrisEmployee](docs/sdks/employee/README.md#gethrisemployee) - Retrieve an employee
* [ListHrisEmployees](docs/sdks/employee/README.md#listhrisemployees) - List all employees
* [PatchHrisEmployee](docs/sdks/employee/README.md#patchhrisemployee) - Update an employee
* [RemoveHrisEmployee](docs/sdks/employee/README.md#removehrisemployee) - Remove an employee
* [UpdateHrisEmployee](docs/sdks/employee/README.md#updatehrisemployee) - Update an employee

### [Enrich](docs/sdks/enrich/README.md)

* [ListEnrichCompanies](docs/sdks/enrich/README.md#listenrichcompanies) - Retrieve enrichment information for a company
* [ListEnrichPeople](docs/sdks/enrich/README.md#listenrichpeople) - Retrieve enrichment information for a person

### [Environment](docs/sdks/environment/README.md)

* [CreateUnifiedEnvironment](docs/sdks/environment/README.md#createunifiedenvironment) - Create new environments
* [ListUnifiedEnvironments](docs/sdks/environment/README.md#listunifiedenvironments) - Returns all environments
* [RemoveUnifiedEnvironment](docs/sdks/environment/README.md#removeunifiedenvironment) - Remove an environment

### [Event](docs/sdks/event/README.md)

* [CreateAnalyticsEvent](docs/sdks/event/README.md#createanalyticsevent) - Create an event
* [CreateCalendarEvent](docs/sdks/event/README.md#createcalendarevent) - Create an event
* [CreateCdpEvent](docs/sdks/event/README.md#createcdpevent) - Create an event
* [CreateCrmEvent](docs/sdks/event/README.md#createcrmevent) - Create an event
* [GetAnalyticsEvent](docs/sdks/event/README.md#getanalyticsevent) - Retrieve an event
* [GetCalendarEvent](docs/sdks/event/README.md#getcalendarevent) - Retrieve an event
* [GetCdpEvent](docs/sdks/event/README.md#getcdpevent) - Retrieve an event
* [GetClubsEvent](docs/sdks/event/README.md#getclubsevent) - Retrieve an event
* [GetCrmEvent](docs/sdks/event/README.md#getcrmevent) - Retrieve an event
* [ListAnalyticsEvents](docs/sdks/event/README.md#listanalyticsevents) - List all events
* [ListCalendarEvents](docs/sdks/event/README.md#listcalendarevents) - List all events
* [ListCdpEvents](docs/sdks/event/README.md#listcdpevents) - List all events
* [ListClubsEvents](docs/sdks/event/README.md#listclubsevents) - List all events
* [ListCrmEvents](docs/sdks/event/README.md#listcrmevents) - List all events
* [PatchCalendarEvent](docs/sdks/event/README.md#patchcalendarevent) - Update an event
* [PatchCdpEvent](docs/sdks/event/README.md#patchcdpevent) - Update an event
* [PatchCrmEvent](docs/sdks/event/README.md#patchcrmevent) - Update an event
* [PatchMessagingEvent](docs/sdks/event/README.md#patchmessagingevent) - Update an event
* [RemoveCalendarEvent](docs/sdks/event/README.md#removecalendarevent) - Remove an event
* [RemoveCdpEvent](docs/sdks/event/README.md#removecdpevent) - Remove an event
* [RemoveCrmEvent](docs/sdks/event/README.md#removecrmevent) - Remove an event
* [UpdateCalendarEvent](docs/sdks/event/README.md#updatecalendarevent) - Update an event
* [UpdateCdpEvent](docs/sdks/event/README.md#updatecdpevent) - Update an event
* [UpdateCrmEvent](docs/sdks/event/README.md#updatecrmevent) - Update an event
* [UpdateMessagingEvent](docs/sdks/event/README.md#updatemessagingevent) - Update an event

### [Expense](docs/sdks/expense/README.md)

* [CreateAccountingExpense](docs/sdks/expense/README.md#createaccountingexpense) - Create an expense
* [GetAccountingExpense](docs/sdks/expense/README.md#getaccountingexpense) - Retrieve an expense
* [ListAccountingExpenses](docs/sdks/expense/README.md#listaccountingexpenses) - List all expenses
* [PatchAccountingExpense](docs/sdks/expense/README.md#patchaccountingexpense) - Update an expense
* [RemoveAccountingExpense](docs/sdks/expense/README.md#removeaccountingexpense) - Remove an expense
* [UpdateAccountingExpense](docs/sdks/expense/README.md#updateaccountingexpense) - Update an expense

### [Feedback](docs/sdks/feedback/README.md)

* [CreatePerformanceFeedback](docs/sdks/feedback/README.md#createperformancefeedback) - Create a feedback
* [GetPerformanceFeedback](docs/sdks/feedback/README.md#getperformancefeedback) - Retrieve a feedback
* [ListPerformanceFeedbacks](docs/sdks/feedback/README.md#listperformancefeedbacks) - List all feedbacks

### [File](docs/sdks/file/README.md)

* [CreateStorageFile](docs/sdks/file/README.md#createstoragefile) - Create a file
* [GetStorageFile](docs/sdks/file/README.md#getstoragefile) - Retrieve a file
* [ListStorageFiles](docs/sdks/file/README.md#liststoragefiles) - List all files
* [PatchStorageFile](docs/sdks/file/README.md#patchstoragefile) - Update a file
* [RemoveStorageFile](docs/sdks/file/README.md#removestoragefile) - Remove a file
* [UpdateStorageFile](docs/sdks/file/README.md#updatestoragefile) - Update a file

### [Form](docs/sdks/form/README.md)

* [CreateFormsForm](docs/sdks/form/README.md#createformsform) - Create a form
* [GetFormsForm](docs/sdks/form/README.md#getformsform) - Retrieve a form
* [ListFormsForms](docs/sdks/form/README.md#listformsforms) - List all forms
* [PatchFormsForm](docs/sdks/form/README.md#patchformsform) - Update a form
* [RemoveFormsForm](docs/sdks/form/README.md#removeformsform) - Remove a form
* [UpdateFormsForm](docs/sdks/form/README.md#updateformsform) - Update a form

### [Forms](docs/sdks/forms/README.md)

* [CreateFormsForm](docs/sdks/forms/README.md#createformsform) - Create a form
* [GetFormsForm](docs/sdks/forms/README.md#getformsform) - Retrieve a form
* [GetFormsSubmission](docs/sdks/forms/README.md#getformssubmission) - Retrieve a submission
* [ListFormsForms](docs/sdks/forms/README.md#listformsforms) - List all forms
* [ListFormsSubmissions](docs/sdks/forms/README.md#listformssubmissions) - List all submissions
* [PatchFormsForm](docs/sdks/forms/README.md#patchformsform) - Update a form
* [RemoveFormsForm](docs/sdks/forms/README.md#removeformsform) - Remove a form
* [UpdateFormsForm](docs/sdks/forms/README.md#updateformsform) - Update a form

### [Genai](docs/sdks/genai/README.md)

* [CreateGenaiEmbedding](docs/sdks/genai/README.md#creategenaiembedding) - Create an embedding
* [CreateGenaiPrompt](docs/sdks/genai/README.md#creategenaiprompt) - Create a prompt
* [GetGenaiModel](docs/sdks/genai/README.md#getgenaimodel) - Retrieve a model
* [ListGenaiModels](docs/sdks/genai/README.md#listgenaimodels) - List all models

### [Goal](docs/sdks/goal/README.md)

* [CreatePerformanceGoal](docs/sdks/goal/README.md#createperformancegoal) - Create a goal
* [GetPerformanceGoal](docs/sdks/goal/README.md#getperformancegoal) - Retrieve a goal
* [ListPerformanceGoals](docs/sdks/goal/README.md#listperformancegoals) - List all goals
* [PatchPerformanceGoal](docs/sdks/goal/README.md#patchperformancegoal) - Update a goal
* [RemovePerformanceGoal](docs/sdks/goal/README.md#removeperformancegoal) - Remove a goal
* [UpdatePerformanceGoal](docs/sdks/goal/README.md#updateperformancegoal) - Update a goal

### [Group](docs/sdks/group/README.md)

* [CreateAdsGroup](docs/sdks/group/README.md#createadsgroup) - Create a group
* [CreateHrisGroup](docs/sdks/group/README.md#createhrisgroup) - Create a group
* [CreateScimGroups](docs/sdks/group/README.md#createscimgroups) - Create group
* [GetAdsGroup](docs/sdks/group/README.md#getadsgroup) - Retrieve a group
* [GetClubsGroup](docs/sdks/group/README.md#getclubsgroup) - Retrieve a group
* [GetHrisGroup](docs/sdks/group/README.md#gethrisgroup) - Retrieve a group
* [GetScimGroups](docs/sdks/group/README.md#getscimgroups) - Get group
* [ListAdsGroups](docs/sdks/group/README.md#listadsgroups) - List all groups
* [ListClubsGroups](docs/sdks/group/README.md#listclubsgroups) - List all groups
* [ListHrisGroups](docs/sdks/group/README.md#listhrisgroups) - List all groups
* [ListScimGroups](docs/sdks/group/README.md#listscimgroups) - List groups
* [PatchAdsGroup](docs/sdks/group/README.md#patchadsgroup) - Update a group
* [PatchHrisGroup](docs/sdks/group/README.md#patchhrisgroup) - Update a group
* [PatchScimGroups](docs/sdks/group/README.md#patchscimgroups) - Update group
* [RemoveAdsGroup](docs/sdks/group/README.md#removeadsgroup) - Remove a group
* [RemoveHrisGroup](docs/sdks/group/README.md#removehrisgroup) - Remove a group
* [RemoveScimGroups](docs/sdks/group/README.md#removescimgroups) - Delete group
* [UpdateAdsGroup](docs/sdks/group/README.md#updateadsgroup) - Update a group
* [UpdateHrisGroup](docs/sdks/group/README.md#updatehrisgroup) - Update a group
* [UpdateScimGroups](docs/sdks/group/README.md#updatescimgroups) - Update group

### [Hris](docs/sdks/hris/README.md)

* [CreateHrisBankaccount](docs/sdks/hris/README.md#createhrisbankaccount) - Create a bankaccount
* [CreateHrisBenefit](docs/sdks/hris/README.md#createhrisbenefit) - Create a benefit
* [CreateHrisCompany](docs/sdks/hris/README.md#createhriscompany) - Create a company
* [CreateHrisDeduction](docs/sdks/hris/README.md#createhrisdeduction) - Create a deduction
* [CreateHrisDevice](docs/sdks/hris/README.md#createhrisdevice) - Create a device
* [CreateHrisDocument](docs/sdks/hris/README.md#createhrisdocument) - Create a document
* [CreateHrisEmployee](docs/sdks/hris/README.md#createhrisemployee) - Create an employee
* [CreateHrisGroup](docs/sdks/hris/README.md#createhrisgroup) - Create a group
* [CreateHrisLocation](docs/sdks/hris/README.md#createhrislocation) - Create a location
* [CreateHrisTimeoff](docs/sdks/hris/README.md#createhristimeoff) - Create a timeoff
* [CreateHrisTimeshift](docs/sdks/hris/README.md#createhristimeshift) - Create a timeshift
* [GetHrisBankaccount](docs/sdks/hris/README.md#gethrisbankaccount) - Retrieve a bankaccount
* [GetHrisBenefit](docs/sdks/hris/README.md#gethrisbenefit) - Retrieve a benefit
* [GetHrisCompany](docs/sdks/hris/README.md#gethriscompany) - Retrieve a company
* [GetHrisDeduction](docs/sdks/hris/README.md#gethrisdeduction) - Retrieve a deduction
* [GetHrisDevice](docs/sdks/hris/README.md#gethrisdevice) - Retrieve a device
* [GetHrisDocument](docs/sdks/hris/README.md#gethrisdocument) - Retrieve a document
* [GetHrisEmployee](docs/sdks/hris/README.md#gethrisemployee) - Retrieve an employee
* [GetHrisGroup](docs/sdks/hris/README.md#gethrisgroup) - Retrieve a group
* [GetHrisLocation](docs/sdks/hris/README.md#gethrislocation) - Retrieve a location
* [GetHrisPayslip](docs/sdks/hris/README.md#gethrispayslip) - Retrieve a payslip
* [GetHrisTaxonomy](docs/sdks/hris/README.md#gethristaxonomy) - Retrieve a taxonomy
* [GetHrisTimeoff](docs/sdks/hris/README.md#gethristimeoff) - Retrieve a timeoff
* [GetHrisTimeshift](docs/sdks/hris/README.md#gethristimeshift) - Retrieve a timeshift
* [ListHrisBankaccounts](docs/sdks/hris/README.md#listhrisbankaccounts) - List all bankaccounts
* [ListHrisBenefits](docs/sdks/hris/README.md#listhrisbenefits) - List all benefits
* [ListHrisCompanies](docs/sdks/hris/README.md#listhriscompanies) - List all companies
* [ListHrisDeductions](docs/sdks/hris/README.md#listhrisdeductions) - List all deductions
* [ListHrisDevices](docs/sdks/hris/README.md#listhrisdevices) - List all devices
* [ListHrisDocuments](docs/sdks/hris/README.md#listhrisdocuments) - List all documents
* [ListHrisEmployees](docs/sdks/hris/README.md#listhrisemployees) - List all employees
* [ListHrisGroups](docs/sdks/hris/README.md#listhrisgroups) - List all groups
* [ListHrisLocations](docs/sdks/hris/README.md#listhrislocations) - List all locations
* [ListHrisPayslips](docs/sdks/hris/README.md#listhrispayslips) - List all payslips
* [ListHrisTaxonomies](docs/sdks/hris/README.md#listhristaxonomies) - List all taxonomies
* [ListHrisTimeoffs](docs/sdks/hris/README.md#listhristimeoffs) - List all timeoffs
* [ListHrisTimeshifts](docs/sdks/hris/README.md#listhristimeshifts) - List all timeshifts
* [PatchHrisBankaccount](docs/sdks/hris/README.md#patchhrisbankaccount) - Update a bankaccount
* [PatchHrisBenefit](docs/sdks/hris/README.md#patchhrisbenefit) - Update a benefit
* [PatchHrisCompany](docs/sdks/hris/README.md#patchhriscompany) - Update a company
* [PatchHrisDeduction](docs/sdks/hris/README.md#patchhrisdeduction) - Update a deduction
* [PatchHrisDevice](docs/sdks/hris/README.md#patchhrisdevice) - Update a device
* [PatchHrisDocument](docs/sdks/hris/README.md#patchhrisdocument) - Update a document
* [PatchHrisEmployee](docs/sdks/hris/README.md#patchhrisemployee) - Update an employee
* [PatchHrisGroup](docs/sdks/hris/README.md#patchhrisgroup) - Update a group
* [PatchHrisLocation](docs/sdks/hris/README.md#patchhrislocation) - Update a location
* [PatchHrisTimeoff](docs/sdks/hris/README.md#patchhristimeoff) - Update a timeoff
* [PatchHrisTimeshift](docs/sdks/hris/README.md#patchhristimeshift) - Update a timeshift
* [RemoveHrisBankaccount](docs/sdks/hris/README.md#removehrisbankaccount) - Remove a bankaccount
* [RemoveHrisBenefit](docs/sdks/hris/README.md#removehrisbenefit) - Remove a benefit
* [RemoveHrisCompany](docs/sdks/hris/README.md#removehriscompany) - Remove a company
* [RemoveHrisDeduction](docs/sdks/hris/README.md#removehrisdeduction) - Remove a deduction
* [RemoveHrisDevice](docs/sdks/hris/README.md#removehrisdevice) - Remove a device
* [RemoveHrisDocument](docs/sdks/hris/README.md#removehrisdocument) - Remove a document
* [RemoveHrisEmployee](docs/sdks/hris/README.md#removehrisemployee) - Remove an employee
* [RemoveHrisGroup](docs/sdks/hris/README.md#removehrisgroup) - Remove a group
* [RemoveHrisLocation](docs/sdks/hris/README.md#removehrislocation) - Remove a location
* [RemoveHrisTimeoff](docs/sdks/hris/README.md#removehristimeoff) - Remove a timeoff
* [RemoveHrisTimeshift](docs/sdks/hris/README.md#removehristimeshift) - Remove a timeshift
* [UpdateHrisBankaccount](docs/sdks/hris/README.md#updatehrisbankaccount) - Update a bankaccount
* [UpdateHrisBenefit](docs/sdks/hris/README.md#updatehrisbenefit) - Update a benefit
* [UpdateHrisCompany](docs/sdks/hris/README.md#updatehriscompany) - Update a company
* [UpdateHrisDeduction](docs/sdks/hris/README.md#updatehrisdeduction) - Update a deduction
* [UpdateHrisDevice](docs/sdks/hris/README.md#updatehrisdevice) - Update a device
* [UpdateHrisDocument](docs/sdks/hris/README.md#updatehrisdocument) - Update a document
* [UpdateHrisEmployee](docs/sdks/hris/README.md#updatehrisemployee) - Update an employee
* [UpdateHrisGroup](docs/sdks/hris/README.md#updatehrisgroup) - Update a group
* [UpdateHrisLocation](docs/sdks/hris/README.md#updatehrislocation) - Update a location
* [UpdateHrisTimeoff](docs/sdks/hris/README.md#updatehristimeoff) - Update a timeoff
* [UpdateHrisTimeshift](docs/sdks/hris/README.md#updatehristimeshift) - Update a timeshift

### [Insertionorder](docs/sdks/insertionorder/README.md)

* [CreateAdsInsertionorder](docs/sdks/insertionorder/README.md#createadsinsertionorder) - Create an insertionorder
* [GetAdsInsertionorder](docs/sdks/insertionorder/README.md#getadsinsertionorder) - Retrieve an insertionorder
* [ListAdsInsertionorders](docs/sdks/insertionorder/README.md#listadsinsertionorders) - List all insertionorders
* [PatchAdsInsertionorder](docs/sdks/insertionorder/README.md#patchadsinsertionorder) - Update an insertionorder
* [RemoveAdsInsertionorder](docs/sdks/insertionorder/README.md#removeadsinsertionorder) - Remove an insertionorder
* [UpdateAdsInsertionorder](docs/sdks/insertionorder/README.md#updateadsinsertionorder) - Update an insertionorder

### [Instructor](docs/sdks/instructor/README.md)

* [CreateLmsInstructor](docs/sdks/instructor/README.md#createlmsinstructor) - Create an instructor
* [GetLmsInstructor](docs/sdks/instructor/README.md#getlmsinstructor) - Retrieve an instructor
* [ListLmsInstructors](docs/sdks/instructor/README.md#listlmsinstructors) - List all instructors
* [PatchLmsInstructor](docs/sdks/instructor/README.md#patchlmsinstructor) - Update an instructor
* [RemoveLmsInstructor](docs/sdks/instructor/README.md#removelmsinstructor) - Remove an instructor
* [UpdateLmsInstructor](docs/sdks/instructor/README.md#updatelmsinstructor) - Update an instructor

### [Integration](docs/sdks/integration/README.md)

* [GetUnifiedIntegrationAuth](docs/sdks/integration/README.md#getunifiedintegrationauth) - Authorize new connection
* [ListUnifiedIntegrationWorkspaces](docs/sdks/integration/README.md#listunifiedintegrationworkspaces) - Returns all activated integrations in a workspace
* [ListUnifiedIntegrations](docs/sdks/integration/README.md#listunifiedintegrations) - Returns all integrations

### [Interview](docs/sdks/interview/README.md)

* [CreateAtsInterview](docs/sdks/interview/README.md#createatsinterview) - Create an interview
* [GetAtsInterview](docs/sdks/interview/README.md#getatsinterview) - Retrieve an interview
* [ListAtsInterviews](docs/sdks/interview/README.md#listatsinterviews) - List all interviews
* [PatchAtsInterview](docs/sdks/interview/README.md#patchatsinterview) - Update an interview
* [RemoveAtsInterview](docs/sdks/interview/README.md#removeatsinterview) - Remove an interview
* [UpdateAtsInterview](docs/sdks/interview/README.md#updateatsinterview) - Update an interview

### [Inventory](docs/sdks/inventory/README.md)

* [CreateCommerceInventory](docs/sdks/inventory/README.md#createcommerceinventory) - Create an inventory
* [GetCommerceInventory](docs/sdks/inventory/README.md#getcommerceinventory) - Retrieve an inventory
* [ListCommerceInventories](docs/sdks/inventory/README.md#listcommerceinventories) - List all inventories
* [PatchCommerceInventory](docs/sdks/inventory/README.md#patchcommerceinventory) - Update an inventory
* [RemoveCommerceInventory](docs/sdks/inventory/README.md#removecommerceinventory) - Remove an inventory
* [UpdateCommerceInventory](docs/sdks/inventory/README.md#updatecommerceinventory) - Update an inventory

### [Invoice](docs/sdks/invoice/README.md)

* [CreateAccountingInvoice](docs/sdks/invoice/README.md#createaccountinginvoice) - Create an invoice
* [GetAccountingInvoice](docs/sdks/invoice/README.md#getaccountinginvoice) - Retrieve an invoice
* [ListAccountingInvoices](docs/sdks/invoice/README.md#listaccountinginvoices) - List all invoices
* [PatchAccountingInvoice](docs/sdks/invoice/README.md#patchaccountinginvoice) - Update an invoice
* [RemoveAccountingInvoice](docs/sdks/invoice/README.md#removeaccountinginvoice) - Remove an invoice
* [UpdateAccountingInvoice](docs/sdks/invoice/README.md#updateaccountinginvoice) - Update an invoice

### [Issue](docs/sdks/issue/README.md)

* [GetUnifiedIssue](docs/sdks/issue/README.md#getunifiedissue) - Retrieve support issue
* [ListUnifiedIssues](docs/sdks/issue/README.md#listunifiedissues) - List support issues

### [Item](docs/sdks/item/README.md)

* [CreateCommerceItem](docs/sdks/item/README.md#createcommerceitem) - Create an item
* [GetCommerceItem](docs/sdks/item/README.md#getcommerceitem) - Retrieve an item
* [ListCommerceItems](docs/sdks/item/README.md#listcommerceitems) - List all items
* [PatchCommerceItem](docs/sdks/item/README.md#patchcommerceitem) - Update an item
* [RemoveCommerceItem](docs/sdks/item/README.md#removecommerceitem) - Remove an item
* [UpdateCommerceItem](docs/sdks/item/README.md#updatecommerceitem) - Update an item

### [Itemvariant](docs/sdks/itemvariant/README.md)

* [CreateCommerceItemvariant](docs/sdks/itemvariant/README.md#createcommerceitemvariant) - Create an itemvariant
* [GetCommerceItemvariant](docs/sdks/itemvariant/README.md#getcommerceitemvariant) - Retrieve an itemvariant
* [ListCommerceItemvariants](docs/sdks/itemvariant/README.md#listcommerceitemvariants) - List all itemvariants
* [PatchCommerceItemvariant](docs/sdks/itemvariant/README.md#patchcommerceitemvariant) - Update an itemvariant
* [RemoveCommerceItemvariant](docs/sdks/itemvariant/README.md#removecommerceitemvariant) - Remove an itemvariant
* [UpdateCommerceItemvariant](docs/sdks/itemvariant/README.md#updatecommerceitemvariant) - Update an itemvariant

### [Job](docs/sdks/job/README.md)

* [CreateAtsJob](docs/sdks/job/README.md#createatsjob) - Create a job
* [GetAtsJob](docs/sdks/job/README.md#getatsjob) - Retrieve a job
* [ListAtsJobs](docs/sdks/job/README.md#listatsjobs) - List all jobs
* [PatchAtsJob](docs/sdks/job/README.md#patchatsjob) - Update a job
* [RemoveAtsJob](docs/sdks/job/README.md#removeatsjob) - Remove a job
* [UpdateAtsJob](docs/sdks/job/README.md#updateatsjob) - Update a job

### [Journal](docs/sdks/journal/README.md)

* [CreateAccountingJournal](docs/sdks/journal/README.md#createaccountingjournal) - Create a journal
* [GetAccountingJournal](docs/sdks/journal/README.md#getaccountingjournal) - Retrieve a journal
* [ListAccountingJournals](docs/sdks/journal/README.md#listaccountingjournals) - List all journals
* [PatchAccountingJournal](docs/sdks/journal/README.md#patchaccountingjournal) - Update a journal
* [RemoveAccountingJournal](docs/sdks/journal/README.md#removeaccountingjournal) - Remove a journal
* [UpdateAccountingJournal](docs/sdks/journal/README.md#updateaccountingjournal) - Update a journal

### [Kms](docs/sdks/kms/README.md)

* [CreateKmsComment](docs/sdks/kms/README.md#createkmscomment) - Create a comment
* [CreateKmsPage](docs/sdks/kms/README.md#createkmspage) - Create a page
* [CreateKmsSpace](docs/sdks/kms/README.md#createkmsspace) - Create a space
* [GetKmsComment](docs/sdks/kms/README.md#getkmscomment) - Retrieve a comment
* [GetKmsPage](docs/sdks/kms/README.md#getkmspage) - Retrieve a page
* [GetKmsSpace](docs/sdks/kms/README.md#getkmsspace) - Retrieve a space
* [ListKmsComments](docs/sdks/kms/README.md#listkmscomments) - List all comments
* [ListKmsPages](docs/sdks/kms/README.md#listkmspages) - List all pages
* [ListKmsSpaces](docs/sdks/kms/README.md#listkmsspaces) - List all spaces
* [PatchKmsComment](docs/sdks/kms/README.md#patchkmscomment) - Update a comment
* [PatchKmsPage](docs/sdks/kms/README.md#patchkmspage) - Update a page
* [PatchKmsSpace](docs/sdks/kms/README.md#patchkmsspace) - Update a space
* [RemoveKmsComment](docs/sdks/kms/README.md#removekmscomment) - Remove a comment
* [RemoveKmsPage](docs/sdks/kms/README.md#removekmspage) - Remove a page
* [RemoveKmsSpace](docs/sdks/kms/README.md#removekmsspace) - Remove a space
* [UpdateKmsComment](docs/sdks/kms/README.md#updatekmscomment) - Update a comment
* [UpdateKmsPage](docs/sdks/kms/README.md#updatekmspage) - Update a page
* [UpdateKmsSpace](docs/sdks/kms/README.md#updatekmsspace) - Update a space

### [Label](docs/sdks/label/README.md)

* [CreateShippingLabel](docs/sdks/label/README.md#createshippinglabel) - Create a label
* [GetShippingLabel](docs/sdks/label/README.md#getshippinglabel) - Retrieve a label
* [ListShippingLabels](docs/sdks/label/README.md#listshippinglabels) - List all labels
* [PatchShippingLabel](docs/sdks/label/README.md#patchshippinglabel) - Update a label
* [RemoveShippingLabel](docs/sdks/label/README.md#removeshippinglabel) - Remove a label
* [UpdateShippingLabel](docs/sdks/label/README.md#updateshippinglabel) - Update a label

### [Lead](docs/sdks/lead/README.md)

* [CreateCrmLead](docs/sdks/lead/README.md#createcrmlead) - Create a lead
* [GetCrmLead](docs/sdks/lead/README.md#getcrmlead) - Retrieve a lead
* [ListCrmLeads](docs/sdks/lead/README.md#listcrmleads) - List all leads
* [PatchCrmLead](docs/sdks/lead/README.md#patchcrmlead) - Update a lead
* [RemoveCrmLead](docs/sdks/lead/README.md#removecrmlead) - Remove a lead
* [UpdateCrmLead](docs/sdks/lead/README.md#updatecrmlead) - Update a lead

### [Link](docs/sdks/link/README.md)

* [CreateCalendarLink](docs/sdks/link/README.md#createcalendarlink) - Create a link
* [CreatePaymentLink](docs/sdks/link/README.md#createpaymentlink) - Create a link
* [GetCalendarLink](docs/sdks/link/README.md#getcalendarlink) - Retrieve a link
* [GetPaymentLink](docs/sdks/link/README.md#getpaymentlink) - Retrieve a link
* [ListCalendarLinks](docs/sdks/link/README.md#listcalendarlinks) - List all links
* [ListPaymentLinks](docs/sdks/link/README.md#listpaymentlinks) - List all links
* [PatchCalendarLink](docs/sdks/link/README.md#patchcalendarlink) - Update a link
* [PatchPaymentLink](docs/sdks/link/README.md#patchpaymentlink) - Update a link
* [RemoveCalendarLink](docs/sdks/link/README.md#removecalendarlink) - Remove a link
* [RemovePaymentLink](docs/sdks/link/README.md#removepaymentlink) - Remove a link
* [UpdateCalendarLink](docs/sdks/link/README.md#updatecalendarlink) - Update a link
* [UpdatePaymentLink](docs/sdks/link/README.md#updatepaymentlink) - Update a link

### [List](docs/sdks/list/README.md)

* [CreateMartechList](docs/sdks/list/README.md#createmartechlist) - Create a list
* [GetMartechList](docs/sdks/list/README.md#getmartechlist) - Retrieve a list
* [ListMartechLists](docs/sdks/list/README.md#listmartechlists) - List all lists
* [PatchMartechList](docs/sdks/list/README.md#patchmartechlist) - Update a list
* [RemoveMartechList](docs/sdks/list/README.md#removemartechlist) - Remove a list
* [UpdateMartechList](docs/sdks/list/README.md#updatemartechlist) - Update a list

### [Lms](docs/sdks/lms/README.md)

* [CreateLmsActivity](docs/sdks/lms/README.md#createlmsactivity) - Create an activity
* [CreateLmsClass](docs/sdks/lms/README.md#createlmsclass) - Create a class
* [CreateLmsCollection](docs/sdks/lms/README.md#createlmscollection) - Create a collection
* [CreateLmsContent](docs/sdks/lms/README.md#createlmscontent) - Create a content
* [CreateLmsCourse](docs/sdks/lms/README.md#createlmscourse) - Create a course
* [CreateLmsInstructor](docs/sdks/lms/README.md#createlmsinstructor) - Create an instructor
* [CreateLmsStudent](docs/sdks/lms/README.md#createlmsstudent) - Create a student
* [GetLmsActivity](docs/sdks/lms/README.md#getlmsactivity) - Retrieve an activity
* [GetLmsClass](docs/sdks/lms/README.md#getlmsclass) - Retrieve a class
* [GetLmsCollection](docs/sdks/lms/README.md#getlmscollection) - Retrieve a collection
* [GetLmsContent](docs/sdks/lms/README.md#getlmscontent) - Retrieve a content
* [GetLmsCourse](docs/sdks/lms/README.md#getlmscourse) - Retrieve a course
* [GetLmsInstructor](docs/sdks/lms/README.md#getlmsinstructor) - Retrieve an instructor
* [GetLmsStudent](docs/sdks/lms/README.md#getlmsstudent) - Retrieve a student
* [ListLmsActivities](docs/sdks/lms/README.md#listlmsactivities) - List all activities
* [ListLmsClasses](docs/sdks/lms/README.md#listlmsclasses) - List all classes
* [ListLmsCollections](docs/sdks/lms/README.md#listlmscollections) - List all collections
* [ListLmsContents](docs/sdks/lms/README.md#listlmscontents) - List all contents
* [ListLmsCourses](docs/sdks/lms/README.md#listlmscourses) - List all courses
* [ListLmsInstructors](docs/sdks/lms/README.md#listlmsinstructors) - List all instructors
* [ListLmsStudents](docs/sdks/lms/README.md#listlmsstudents) - List all students
* [PatchLmsActivity](docs/sdks/lms/README.md#patchlmsactivity) - Update an activity
* [PatchLmsClass](docs/sdks/lms/README.md#patchlmsclass) - Update a class
* [PatchLmsCollection](docs/sdks/lms/README.md#patchlmscollection) - Update a collection
* [PatchLmsContent](docs/sdks/lms/README.md#patchlmscontent) - Update a content
* [PatchLmsCourse](docs/sdks/lms/README.md#patchlmscourse) - Update a course
* [PatchLmsInstructor](docs/sdks/lms/README.md#patchlmsinstructor) - Update an instructor
* [PatchLmsStudent](docs/sdks/lms/README.md#patchlmsstudent) - Update a student
* [RemoveLmsActivity](docs/sdks/lms/README.md#removelmsactivity) - Remove an activity
* [RemoveLmsClass](docs/sdks/lms/README.md#removelmsclass) - Remove a class
* [RemoveLmsCollection](docs/sdks/lms/README.md#removelmscollection) - Remove a collection
* [RemoveLmsContent](docs/sdks/lms/README.md#removelmscontent) - Remove a content
* [RemoveLmsCourse](docs/sdks/lms/README.md#removelmscourse) - Remove a course
* [RemoveLmsInstructor](docs/sdks/lms/README.md#removelmsinstructor) - Remove an instructor
* [RemoveLmsStudent](docs/sdks/lms/README.md#removelmsstudent) - Remove a student
* [UpdateLmsActivity](docs/sdks/lms/README.md#updatelmsactivity) - Update an activity
* [UpdateLmsClass](docs/sdks/lms/README.md#updatelmsclass) - Update a class
* [UpdateLmsCollection](docs/sdks/lms/README.md#updatelmscollection) - Update a collection
* [UpdateLmsContent](docs/sdks/lms/README.md#updatelmscontent) - Update a content
* [UpdateLmsCourse](docs/sdks/lms/README.md#updatelmscourse) - Update a course
* [UpdateLmsInstructor](docs/sdks/lms/README.md#updatelmsinstructor) - Update an instructor
* [UpdateLmsStudent](docs/sdks/lms/README.md#updatelmsstudent) - Update a student

### [Location](docs/sdks/location/README.md)

* [CreateCommerceLocation](docs/sdks/location/README.md#createcommercelocation) - Create a location
* [CreateHrisLocation](docs/sdks/location/README.md#createhrislocation) - Create a location
* [GetClubsLocation](docs/sdks/location/README.md#getclubslocation) - Retrieve a location
* [GetCommerceLocation](docs/sdks/location/README.md#getcommercelocation) - Retrieve a location
* [GetHrisLocation](docs/sdks/location/README.md#gethrislocation) - Retrieve a location
* [ListClubsLocations](docs/sdks/location/README.md#listclubslocations) - List all locations
* [ListCommerceLocations](docs/sdks/location/README.md#listcommercelocations) - List all locations
* [ListHrisLocations](docs/sdks/location/README.md#listhrislocations) - List all locations
* [PatchCommerceLocation](docs/sdks/location/README.md#patchcommercelocation) - Update a location
* [PatchHrisLocation](docs/sdks/location/README.md#patchhrislocation) - Update a location
* [RemoveCommerceLocation](docs/sdks/location/README.md#removecommercelocation) - Remove a location
* [RemoveHrisLocation](docs/sdks/location/README.md#removehrislocation) - Remove a location
* [UpdateCommerceLocation](docs/sdks/location/README.md#updatecommercelocation) - Update a location
* [UpdateHrisLocation](docs/sdks/location/README.md#updatehrislocation) - Update a location

### [Login](docs/sdks/login/README.md)

* [GetUnifiedIntegrationLogin](docs/sdks/login/README.md#getunifiedintegrationlogin) - Sign in a user

### [Martech](docs/sdks/martech/README.md)

* [CreateMartechCampaign](docs/sdks/martech/README.md#createmartechcampaign) - Create a campaign
* [CreateMartechList](docs/sdks/martech/README.md#createmartechlist) - Create a list
* [CreateMartechMember](docs/sdks/martech/README.md#createmartechmember) - Create a member
* [GetMartechCampaign](docs/sdks/martech/README.md#getmartechcampaign) - Retrieve a campaign
* [GetMartechList](docs/sdks/martech/README.md#getmartechlist) - Retrieve a list
* [GetMartechMember](docs/sdks/martech/README.md#getmartechmember) - Retrieve a member
* [ListMartechCampaigns](docs/sdks/martech/README.md#listmartechcampaigns) - List all campaigns
* [ListMartechLists](docs/sdks/martech/README.md#listmartechlists) - List all lists
* [ListMartechMembers](docs/sdks/martech/README.md#listmartechmembers) - List all members
* [ListMartechReports](docs/sdks/martech/README.md#listmartechreports) - List all reports
* [PatchMartechCampaign](docs/sdks/martech/README.md#patchmartechcampaign) - Update a campaign
* [PatchMartechList](docs/sdks/martech/README.md#patchmartechlist) - Update a list
* [PatchMartechMember](docs/sdks/martech/README.md#patchmartechmember) - Update a member
* [RemoveMartechCampaign](docs/sdks/martech/README.md#removemartechcampaign) - Remove a campaign
* [RemoveMartechList](docs/sdks/martech/README.md#removemartechlist) - Remove a list
* [RemoveMartechMember](docs/sdks/martech/README.md#removemartechmember) - Remove a member
* [UpdateMartechCampaign](docs/sdks/martech/README.md#updatemartechcampaign) - Update a campaign
* [UpdateMartechList](docs/sdks/martech/README.md#updatemartechlist) - Update a list
* [UpdateMartechMember](docs/sdks/martech/README.md#updatemartechmember) - Update a member

### [Member](docs/sdks/member/README.md)

* [CreateMartechMember](docs/sdks/member/README.md#createmartechmember) - Create a member
* [GetClubsMember](docs/sdks/member/README.md#getclubsmember) - Retrieve a member
* [GetMartechMember](docs/sdks/member/README.md#getmartechmember) - Retrieve a member
* [ListClubsMembers](docs/sdks/member/README.md#listclubsmembers) - List all members
* [ListMartechMembers](docs/sdks/member/README.md#listmartechmembers) - List all members
* [PatchMartechMember](docs/sdks/member/README.md#patchmartechmember) - Update a member
* [RemoveMartechMember](docs/sdks/member/README.md#removemartechmember) - Remove a member
* [UpdateMartechMember](docs/sdks/member/README.md#updatemartechmember) - Update a member

### [Message](docs/sdks/message/README.md)

* [CreateMessagingMessage](docs/sdks/message/README.md#createmessagingmessage) - Create a message
* [GetMessagingMessage](docs/sdks/message/README.md#getmessagingmessage) - Retrieve a message
* [ListMessagingMessages](docs/sdks/message/README.md#listmessagingmessages) - List all messages
* [PatchMessagingMessage](docs/sdks/message/README.md#patchmessagingmessage) - Update a message
* [RemoveMessagingMessage](docs/sdks/message/README.md#removemessagingmessage) - Remove a message
* [UpdateMessagingMessage](docs/sdks/message/README.md#updatemessagingmessage) - Update a message

### [Messaging](docs/sdks/messaging/README.md)

* [CreateMessagingChannel](docs/sdks/messaging/README.md#createmessagingchannel) - Create a channel
* [CreateMessagingMessage](docs/sdks/messaging/README.md#createmessagingmessage) - Create a message
* [GetMessagingChannel](docs/sdks/messaging/README.md#getmessagingchannel) - Retrieve a channel
* [GetMessagingMessage](docs/sdks/messaging/README.md#getmessagingmessage) - Retrieve a message
* [ListMessagingChannels](docs/sdks/messaging/README.md#listmessagingchannels) - List all channels
* [ListMessagingMessages](docs/sdks/messaging/README.md#listmessagingmessages) - List all messages
* [PatchMessagingChannel](docs/sdks/messaging/README.md#patchmessagingchannel) - Update a channel
* [PatchMessagingEvent](docs/sdks/messaging/README.md#patchmessagingevent) - Update an event
* [PatchMessagingMessage](docs/sdks/messaging/README.md#patchmessagingmessage) - Update a message
* [RemoveMessagingChannel](docs/sdks/messaging/README.md#removemessagingchannel) - Remove a channel
* [RemoveMessagingMessage](docs/sdks/messaging/README.md#removemessagingmessage) - Remove a message
* [UpdateMessagingChannel](docs/sdks/messaging/README.md#updatemessagingchannel) - Update a channel
* [UpdateMessagingEvent](docs/sdks/messaging/README.md#updatemessagingevent) - Update an event
* [UpdateMessagingMessage](docs/sdks/messaging/README.md#updatemessagingmessage) - Update a message

### [Metadata](docs/sdks/metadata/README.md)

* [CreateMetadataMetadata](docs/sdks/metadata/README.md#createmetadatametadata) - Create a metadata
* [GetMetadataMetadata](docs/sdks/metadata/README.md#getmetadatametadata) - Retrieve a metadata
* [ListMetadataMetadatas](docs/sdks/metadata/README.md#listmetadatametadatas) - List all metadatas
* [PatchMetadataMetadata](docs/sdks/metadata/README.md#patchmetadatametadata) - Update a metadata
* [RemoveMetadataMetadata](docs/sdks/metadata/README.md#removemetadatametadata) - Remove a metadata
* [UpdateMetadataMetadata](docs/sdks/metadata/README.md#updatemetadatametadata) - Update a metadata

### [Model](docs/sdks/model/README.md)

* [GetGenaiModel](docs/sdks/model/README.md#getgenaimodel) - Retrieve a model
* [ListGenaiModels](docs/sdks/model/README.md#listgenaimodels) - List all models

### [Note](docs/sdks/note/README.md)

* [CreateTicketingNote](docs/sdks/note/README.md#createticketingnote) - Create a note
* [GetTicketingNote](docs/sdks/note/README.md#getticketingnote) - Retrieve a note
* [ListTicketingNotes](docs/sdks/note/README.md#listticketingnotes) - List all notes
* [PatchTicketingNote](docs/sdks/note/README.md#patchticketingnote) - Update a note
* [RemoveTicketingNote](docs/sdks/note/README.md#removeticketingnote) - Remove a note
* [UpdateTicketingNote](docs/sdks/note/README.md#updateticketingnote) - Update a note

### [Order](docs/sdks/order/README.md)

* [CreateAccountingOrder](docs/sdks/order/README.md#createaccountingorder) - Create an order
* [GetAccountingOrder](docs/sdks/order/README.md#getaccountingorder) - Retrieve an order
* [ListAccountingOrders](docs/sdks/order/README.md#listaccountingorders) - List all orders
* [PatchAccountingOrder](docs/sdks/order/README.md#patchaccountingorder) - Update an order
* [PatchAssessmentOrder](docs/sdks/order/README.md#patchassessmentorder) - Update an order
* [RemoveAccountingOrder](docs/sdks/order/README.md#removeaccountingorder) - Remove an order
* [UpdateAccountingOrder](docs/sdks/order/README.md#updateaccountingorder) - Update an order
* [UpdateAssessmentOrder](docs/sdks/order/README.md#updateassessmentorder) - Update an order

### [Organization](docs/sdks/organization/README.md)

* [CreateAdsOrganization](docs/sdks/organization/README.md#createadsorganization) - Create an organization
* [CreateRepoOrganization](docs/sdks/organization/README.md#createrepoorganization) - Create an organization
* [GetAccountingOrganization](docs/sdks/organization/README.md#getaccountingorganization) - Retrieve an organization
* [GetAdsOrganization](docs/sdks/organization/README.md#getadsorganization) - Retrieve an organization
* [GetRepoOrganization](docs/sdks/organization/README.md#getrepoorganization) - Retrieve an organization
* [ListAccountingOrganizations](docs/sdks/organization/README.md#listaccountingorganizations) - List all organizations
* [ListAdsOrganizations](docs/sdks/organization/README.md#listadsorganizations) - List all organizations
* [ListRepoOrganizations](docs/sdks/organization/README.md#listrepoorganizations) - List all organizations
* [PatchAdsOrganization](docs/sdks/organization/README.md#patchadsorganization) - Update an organization
* [PatchRepoOrganization](docs/sdks/organization/README.md#patchrepoorganization) - Update an organization
* [RemoveAdsOrganization](docs/sdks/organization/README.md#removeadsorganization) - Remove an organization
* [RemoveRepoOrganization](docs/sdks/organization/README.md#removerepoorganization) - Remove an organization
* [UpdateAdsOrganization](docs/sdks/organization/README.md#updateadsorganization) - Update an organization
* [UpdateRepoOrganization](docs/sdks/organization/README.md#updaterepoorganization) - Update an organization

### [Package](docs/sdks/package/README.md)

* [CreateAssessmentPackage](docs/sdks/package/README.md#createassessmentpackage) - Create an assessment package
* [GetAssessmentPackage](docs/sdks/package/README.md#getassessmentpackage) - Get an assessment package
* [GetVerificationPackage](docs/sdks/package/README.md#getverificationpackage) - Retrieve a package
* [ListAssessmentPackages](docs/sdks/package/README.md#listassessmentpackages) - List assessment packages
* [ListVerificationPackages](docs/sdks/package/README.md#listverificationpackages) - List all packages
* [PatchAssessmentPackage](docs/sdks/package/README.md#patchassessmentpackage) - Update an assessment package
* [RemoveAssessmentPackage](docs/sdks/package/README.md#removeassessmentpackage) - Delete an assessment package
* [UpdateAssessmentPackage](docs/sdks/package/README.md#updateassessmentpackage) - Update an assessment package

### [Page](docs/sdks/page/README.md)

* [CreateKmsPage](docs/sdks/page/README.md#createkmspage) - Create a page
* [GetKmsPage](docs/sdks/page/README.md#getkmspage) - Retrieve a page
* [ListKmsPages](docs/sdks/page/README.md#listkmspages) - List all pages
* [PatchKmsPage](docs/sdks/page/README.md#patchkmspage) - Update a page
* [RemoveKmsPage](docs/sdks/page/README.md#removekmspage) - Remove a page
* [UpdateKmsPage](docs/sdks/page/README.md#updatekmspage) - Update a page

### [Passthrough](docs/sdks/passthrough/README.md)

* [CreatePassthroughJSON](docs/sdks/passthrough/README.md#createpassthroughjson) - Passthrough POST
* [CreatePassthroughRaw](docs/sdks/passthrough/README.md#createpassthroughraw) - Passthrough POST
* [ListPassthroughs](docs/sdks/passthrough/README.md#listpassthroughs) - Passthrough GET
* [PatchPassthroughJSON](docs/sdks/passthrough/README.md#patchpassthroughjson) - Passthrough PUT
* [PatchPassthroughRaw](docs/sdks/passthrough/README.md#patchpassthroughraw) - Passthrough PUT
* [RemovePassthrough](docs/sdks/passthrough/README.md#removepassthrough) - Passthrough DELETE
* [UpdatePassthroughJSON](docs/sdks/passthrough/README.md#updatepassthroughjson) - Passthrough PUT
* [UpdatePassthroughRaw](docs/sdks/passthrough/README.md#updatepassthroughraw) - Passthrough PUT

### [Payment](docs/sdks/payment/README.md)

* [CreatePaymentLink](docs/sdks/payment/README.md#createpaymentlink) - Create a link
* [CreatePaymentPayment](docs/sdks/payment/README.md#createpaymentpayment) - Create a payment
* [CreatePaymentSubscription](docs/sdks/payment/README.md#createpaymentsubscription) - Create a subscription
* [GetPaymentLink](docs/sdks/payment/README.md#getpaymentlink) - Retrieve a link
* [GetPaymentPayment](docs/sdks/payment/README.md#getpaymentpayment) - Retrieve a payment
* [GetPaymentPayout](docs/sdks/payment/README.md#getpaymentpayout) - Retrieve a payout
* [GetPaymentRefund](docs/sdks/payment/README.md#getpaymentrefund) - Retrieve a refund
* [GetPaymentSubscription](docs/sdks/payment/README.md#getpaymentsubscription) - Retrieve a subscription
* [ListPaymentLinks](docs/sdks/payment/README.md#listpaymentlinks) - List all links
* [ListPaymentPayments](docs/sdks/payment/README.md#listpaymentpayments) - List all payments
* [ListPaymentPayouts](docs/sdks/payment/README.md#listpaymentpayouts) - List all payouts
* [ListPaymentRefunds](docs/sdks/payment/README.md#listpaymentrefunds) - List all refunds
* [ListPaymentSubscriptions](docs/sdks/payment/README.md#listpaymentsubscriptions) - List all subscriptions
* [PatchPaymentLink](docs/sdks/payment/README.md#patchpaymentlink) - Update a link
* [PatchPaymentPayment](docs/sdks/payment/README.md#patchpaymentpayment) - Update a payment
* [PatchPaymentSubscription](docs/sdks/payment/README.md#patchpaymentsubscription) - Update a subscription
* [RemovePaymentLink](docs/sdks/payment/README.md#removepaymentlink) - Remove a link
* [RemovePaymentPayment](docs/sdks/payment/README.md#removepaymentpayment) - Remove a payment
* [RemovePaymentSubscription](docs/sdks/payment/README.md#removepaymentsubscription) - Remove a subscription
* [UpdatePaymentLink](docs/sdks/payment/README.md#updatepaymentlink) - Update a link
* [UpdatePaymentPayment](docs/sdks/payment/README.md#updatepaymentpayment) - Update a payment
* [UpdatePaymentSubscription](docs/sdks/payment/README.md#updatepaymentsubscription) - Update a subscription

### [Payout](docs/sdks/payout/README.md)

* [GetPaymentPayout](docs/sdks/payout/README.md#getpaymentpayout) - Retrieve a payout
* [ListPaymentPayouts](docs/sdks/payout/README.md#listpaymentpayouts) - List all payouts

### [Payslip](docs/sdks/payslip/README.md)

* [GetHrisPayslip](docs/sdks/payslip/README.md#gethrispayslip) - Retrieve a payslip
* [ListHrisPayslips](docs/sdks/payslip/README.md#listhrispayslips) - List all payslips

### [Performance](docs/sdks/performance/README.md)

* [CreatePerformanceFeedback](docs/sdks/performance/README.md#createperformancefeedback) - Create a feedback
* [CreatePerformanceGoal](docs/sdks/performance/README.md#createperformancegoal) - Create a goal
* [GetPerformanceCycle](docs/sdks/performance/README.md#getperformancecycle) - Retrieve a cycle
* [GetPerformanceFeedback](docs/sdks/performance/README.md#getperformancefeedback) - Retrieve a feedback
* [GetPerformanceGoal](docs/sdks/performance/README.md#getperformancegoal) - Retrieve a goal
* [GetPerformanceReview](docs/sdks/performance/README.md#getperformancereview) - Retrieve a review
* [ListPerformanceCycles](docs/sdks/performance/README.md#listperformancecycles) - List all cycles
* [ListPerformanceFeedbacks](docs/sdks/performance/README.md#listperformancefeedbacks) - List all feedbacks
* [ListPerformanceGoals](docs/sdks/performance/README.md#listperformancegoals) - List all goals
* [ListPerformanceReviews](docs/sdks/performance/README.md#listperformancereviews) - List all reviews
* [PatchPerformanceGoal](docs/sdks/performance/README.md#patchperformancegoal) - Update a goal
* [RemovePerformanceGoal](docs/sdks/performance/README.md#removeperformancegoal) - Remove a goal
* [UpdatePerformanceGoal](docs/sdks/performance/README.md#updateperformancegoal) - Update a goal

### [Person](docs/sdks/person/README.md)

* [ListEnrichPeople](docs/sdks/person/README.md#listenrichpeople) - Retrieve enrichment information for a person

### [Picklist](docs/sdks/picklist/README.md)

* [ListCrmPicklists](docs/sdks/picklist/README.md#listcrmpicklists) - List all picklists

### [Pipeline](docs/sdks/pipeline/README.md)

* [CreateCrmPipeline](docs/sdks/pipeline/README.md#createcrmpipeline) - Create a pipeline
* [GetCrmPipeline](docs/sdks/pipeline/README.md#getcrmpipeline) - Retrieve a pipeline
* [ListCrmPipelines](docs/sdks/pipeline/README.md#listcrmpipelines) - List all pipelines
* [PatchCrmPipeline](docs/sdks/pipeline/README.md#patchcrmpipeline) - Update a pipeline
* [RemoveCrmPipeline](docs/sdks/pipeline/README.md#removecrmpipeline) - Remove a pipeline
* [UpdateCrmPipeline](docs/sdks/pipeline/README.md#updatecrmpipeline) - Update a pipeline

### [Profile](docs/sdks/profile/README.md)

* [CreateCdpProfile](docs/sdks/profile/README.md#createcdpprofile) - Create a profile
* [GetCdpProfile](docs/sdks/profile/README.md#getcdpprofile) - Retrieve a profile
* [ListCdpProfiles](docs/sdks/profile/README.md#listcdpprofiles) - List all profiles
* [PatchCdpProfile](docs/sdks/profile/README.md#patchcdpprofile) - Update a profile
* [RemoveCdpProfile](docs/sdks/profile/README.md#removecdpprofile) - Remove a profile
* [UpdateCdpProfile](docs/sdks/profile/README.md#updatecdpprofile) - Update a profile

### [Profitloss](docs/sdks/profitloss/README.md)

* [GetAccountingProfitloss](docs/sdks/profitloss/README.md#getaccountingprofitloss) - Retrieve a profitloss
* [ListAccountingProfitlosses](docs/sdks/profitloss/README.md#listaccountingprofitlosses) - List all profitlosses

### [Project](docs/sdks/project/README.md)

* [CreateAccountingProject](docs/sdks/project/README.md#createaccountingproject) - Create a project
* [CreateTaskProject](docs/sdks/project/README.md#createtaskproject) - Create a project
* [GetAccountingProject](docs/sdks/project/README.md#getaccountingproject) - Retrieve a project
* [GetTaskProject](docs/sdks/project/README.md#gettaskproject) - Retrieve a project
* [ListAccountingProjects](docs/sdks/project/README.md#listaccountingprojects) - List all projects
* [ListTaskProjects](docs/sdks/project/README.md#listtaskprojects) - List all projects
* [PatchAccountingProject](docs/sdks/project/README.md#patchaccountingproject) - Update a project
* [PatchTaskProject](docs/sdks/project/README.md#patchtaskproject) - Update a project
* [RemoveAccountingProject](docs/sdks/project/README.md#removeaccountingproject) - Remove a project
* [RemoveTaskProject](docs/sdks/project/README.md#removetaskproject) - Remove a project
* [UpdateAccountingProject](docs/sdks/project/README.md#updateaccountingproject) - Update a project
* [UpdateTaskProject](docs/sdks/project/README.md#updatetaskproject) - Update a project

### [Promoted](docs/sdks/promoted/README.md)

* [GetAdsPromoted](docs/sdks/promoted/README.md#getadspromoted) - Retrieve a promoted
* [ListAdsPromoteds](docs/sdks/promoted/README.md#listadspromoteds) - List all promoteds

### [Prompt](docs/sdks/prompt/README.md)

* [CreateGenaiPrompt](docs/sdks/prompt/README.md#creategenaiprompt) - Create a prompt

### [Property](docs/sdks/property/README.md)

* [CreateAnalyticsProperty](docs/sdks/property/README.md#createanalyticsproperty) - Create a property
* [GetAnalyticsProperty](docs/sdks/property/README.md#getanalyticsproperty) - Retrieve a property
* [ListAnalyticsProperties](docs/sdks/property/README.md#listanalyticsproperties) - List all properties
* [PatchAnalyticsProperty](docs/sdks/property/README.md#patchanalyticsproperty) - Update a property
* [RemoveAnalyticsProperty](docs/sdks/property/README.md#removeanalyticsproperty) - Remove a property
* [UpdateAnalyticsProperty](docs/sdks/property/README.md#updateanalyticsproperty) - Update a property

### [Pullrequest](docs/sdks/pullrequest/README.md)

* [CreateRepoPullrequest](docs/sdks/pullrequest/README.md#createrepopullrequest) - Create a pullrequest
* [GetRepoPullrequest](docs/sdks/pullrequest/README.md#getrepopullrequest) - Retrieve a pullrequest
* [ListRepoPullrequests](docs/sdks/pullrequest/README.md#listrepopullrequests) - List all pullrequests
* [PatchRepoPullrequest](docs/sdks/pullrequest/README.md#patchrepopullrequest) - Update a pullrequest
* [RemoveRepoPullrequest](docs/sdks/pullrequest/README.md#removerepopullrequest) - Remove a pullrequest
* [UpdateRepoPullrequest](docs/sdks/pullrequest/README.md#updaterepopullrequest) - Update a pullrequest

### [Purchaseorder](docs/sdks/purchaseorder/README.md)

* [CreateAccountingPurchaseorder](docs/sdks/purchaseorder/README.md#createaccountingpurchaseorder) - Create a purchaseorder
* [GetAccountingPurchaseorder](docs/sdks/purchaseorder/README.md#getaccountingpurchaseorder) - Retrieve a purchaseorder
* [ListAccountingPurchaseorders](docs/sdks/purchaseorder/README.md#listaccountingpurchaseorders) - List all purchaseorders
* [PatchAccountingPurchaseorder](docs/sdks/purchaseorder/README.md#patchaccountingpurchaseorder) - Update a purchaseorder
* [RemoveAccountingPurchaseorder](docs/sdks/purchaseorder/README.md#removeaccountingpurchaseorder) - Remove a purchaseorder
* [UpdateAccountingPurchaseorder](docs/sdks/purchaseorder/README.md#updateaccountingpurchaseorder) - Update a purchaseorder

### [Query](docs/sdks/query/README.md)

* [CreateDatastoreQuery](docs/sdks/query/README.md#createdatastorequery) - Create a query

### [Quote](docs/sdks/quote/README.md)

* [CreateAccountingQuote](docs/sdks/quote/README.md#createaccountingquote) - Create a quote
* [GetAccountingQuote](docs/sdks/quote/README.md#getaccountingquote) - Retrieve a quote
* [ListAccountingQuotes](docs/sdks/quote/README.md#listaccountingquotes) - List all quotes
* [PatchAccountingQuote](docs/sdks/quote/README.md#patchaccountingquote) - Update a quote
* [RemoveAccountingQuote](docs/sdks/quote/README.md#removeaccountingquote) - Remove a quote
* [UpdateAccountingQuote](docs/sdks/quote/README.md#updateaccountingquote) - Update a quote

### [Rate](docs/sdks/rate/README.md)

* [CreateShippingRate](docs/sdks/rate/README.md#createshippingrate) - Create a rate

### [Record](docs/sdks/record/README.md)

* [CreateDatastoreRecord](docs/sdks/record/README.md#createdatastorerecord) - Create a record
* [GetDatastoreRecord](docs/sdks/record/README.md#getdatastorerecord) - Retrieve a record
* [ListDatastoreRecords](docs/sdks/record/README.md#listdatastorerecords) - List all records
* [PatchDatastoreRecord](docs/sdks/record/README.md#patchdatastorerecord) - Update a record
* [RemoveDatastoreRecord](docs/sdks/record/README.md#removedatastorerecord) - Remove a record
* [UpdateDatastoreRecord](docs/sdks/record/README.md#updatedatastorerecord) - Update a record

### [Recording](docs/sdks/recording/README.md)

* [CreateUcRecording](docs/sdks/recording/README.md#createucrecording) - Create a recording
* [GetCalendarRecording](docs/sdks/recording/README.md#getcalendarrecording) - Retrieve a recording
* [GetUcRecording](docs/sdks/recording/README.md#getucrecording) - Retrieve a recording
* [ListCalendarRecordings](docs/sdks/recording/README.md#listcalendarrecordings) - List all recordings
* [ListUcRecordings](docs/sdks/recording/README.md#listucrecordings) - List all recordings
* [PatchUcRecording](docs/sdks/recording/README.md#patchucrecording) - Update a recording
* [RemoveUcRecording](docs/sdks/recording/README.md#removeucrecording) - Remove a recording
* [UpdateUcRecording](docs/sdks/recording/README.md#updateucrecording) - Update a recording

### [Refund](docs/sdks/refund/README.md)

* [GetPaymentRefund](docs/sdks/refund/README.md#getpaymentrefund) - Retrieve a refund
* [ListPaymentRefunds](docs/sdks/refund/README.md#listpaymentrefunds) - List all refunds

### [Repo](docs/sdks/repo/README.md)

* [CreateRepoBranch](docs/sdks/repo/README.md#createrepobranch) - Create a branch
* [CreateRepoCommit](docs/sdks/repo/README.md#createrepocommit) - Create a commit
* [CreateRepoOrganization](docs/sdks/repo/README.md#createrepoorganization) - Create an organization
* [CreateRepoPullrequest](docs/sdks/repo/README.md#createrepopullrequest) - Create a pullrequest
* [CreateRepoRepository](docs/sdks/repo/README.md#createreporepository) - Create a repository
* [GetRepoBranch](docs/sdks/repo/README.md#getrepobranch) - Retrieve a branch
* [GetRepoCommit](docs/sdks/repo/README.md#getrepocommit) - Retrieve a commit
* [GetRepoOrganization](docs/sdks/repo/README.md#getrepoorganization) - Retrieve an organization
* [GetRepoPullrequest](docs/sdks/repo/README.md#getrepopullrequest) - Retrieve a pullrequest
* [GetRepoRepository](docs/sdks/repo/README.md#getreporepository) - Retrieve a repository
* [ListRepoBranches](docs/sdks/repo/README.md#listrepobranches) - List all branches
* [ListRepoCommits](docs/sdks/repo/README.md#listrepocommits) - List all commits
* [ListRepoOrganizations](docs/sdks/repo/README.md#listrepoorganizations) - List all organizations
* [ListRepoPullrequests](docs/sdks/repo/README.md#listrepopullrequests) - List all pullrequests
* [ListRepoRepositories](docs/sdks/repo/README.md#listreporepositories) - List all repositories
* [PatchRepoBranch](docs/sdks/repo/README.md#patchrepobranch) - Update a branch
* [PatchRepoCommit](docs/sdks/repo/README.md#patchrepocommit) - Update a commit
* [PatchRepoOrganization](docs/sdks/repo/README.md#patchrepoorganization) - Update an organization
* [PatchRepoPullrequest](docs/sdks/repo/README.md#patchrepopullrequest) - Update a pullrequest
* [PatchRepoRepository](docs/sdks/repo/README.md#patchreporepository) - Update a repository
* [RemoveRepoBranch](docs/sdks/repo/README.md#removerepobranch) - Remove a branch
* [RemoveRepoCommit](docs/sdks/repo/README.md#removerepocommit) - Remove a commit
* [RemoveRepoOrganization](docs/sdks/repo/README.md#removerepoorganization) - Remove an organization
* [RemoveRepoPullrequest](docs/sdks/repo/README.md#removerepopullrequest) - Remove a pullrequest
* [RemoveRepoRepository](docs/sdks/repo/README.md#removereporepository) - Remove a repository
* [UpdateRepoBranch](docs/sdks/repo/README.md#updaterepobranch) - Update a branch
* [UpdateRepoCommit](docs/sdks/repo/README.md#updaterepocommit) - Update a commit
* [UpdateRepoOrganization](docs/sdks/repo/README.md#updaterepoorganization) - Update an organization
* [UpdateRepoPullrequest](docs/sdks/repo/README.md#updaterepopullrequest) - Update a pullrequest
* [UpdateRepoRepository](docs/sdks/repo/README.md#updatereporepository) - Update a repository

### [Report](docs/sdks/report/README.md)

* [GetAccountingReport](docs/sdks/report/README.md#getaccountingreport) - Retrieve a report
* [ListAccountingReports](docs/sdks/report/README.md#listaccountingreports) - List all reports
* [ListAdsReports](docs/sdks/report/README.md#listadsreports) - List all reports
* [ListAnalyticsReports](docs/sdks/report/README.md#listanalyticsreports) - List all reports
* [ListMartechReports](docs/sdks/report/README.md#listmartechreports) - List all reports

### [Repository](docs/sdks/repository/README.md)

* [CreateRepoRepository](docs/sdks/repository/README.md#createreporepository) - Create a repository
* [GetRepoRepository](docs/sdks/repository/README.md#getreporepository) - Retrieve a repository
* [ListRepoRepositories](docs/sdks/repository/README.md#listreporepositories) - List all repositories
* [PatchRepoRepository](docs/sdks/repository/README.md#patchreporepository) - Update a repository
* [RemoveRepoRepository](docs/sdks/repository/README.md#removereporepository) - Remove a repository
* [UpdateRepoRepository](docs/sdks/repository/README.md#updatereporepository) - Update a repository

### [Request](docs/sdks/request/README.md)

* [CreateVerificationRequest](docs/sdks/request/README.md#createverificationrequest) - Create a request
* [GetVerificationRequest](docs/sdks/request/README.md#getverificationrequest) - Retrieve a request
* [ListVerificationRequests](docs/sdks/request/README.md#listverificationrequests) - List all requests
* [PatchVerificationRequest](docs/sdks/request/README.md#patchverificationrequest) - Update a request
* [RemoveVerificationRequest](docs/sdks/request/README.md#removeverificationrequest) - Remove a request
* [UpdateVerificationRequest](docs/sdks/request/README.md#updateverificationrequest) - Update a request

### [Reservation](docs/sdks/reservation/README.md)

* [CreateCommerceReservation](docs/sdks/reservation/README.md#createcommercereservation) - Create a reservation
* [GetCommerceReservation](docs/sdks/reservation/README.md#getcommercereservation) - Retrieve a reservation
* [ListCommerceReservations](docs/sdks/reservation/README.md#listcommercereservations) - List all reservations
* [PatchCommerceReservation](docs/sdks/reservation/README.md#patchcommercereservation) - Update a reservation
* [RemoveCommerceReservation](docs/sdks/reservation/README.md#removecommercereservation) - Remove a reservation
* [UpdateCommerceReservation](docs/sdks/reservation/README.md#updatecommercereservation) - Update a reservation

### [Review](docs/sdks/review/README.md)

* [CreateCommerceReview](docs/sdks/review/README.md#createcommercereview) - Create a review
* [GetCommerceReview](docs/sdks/review/README.md#getcommercereview) - Retrieve a review
* [GetPerformanceReview](docs/sdks/review/README.md#getperformancereview) - Retrieve a review
* [ListCommerceReviews](docs/sdks/review/README.md#listcommercereviews) - List all reviews
* [ListPerformanceReviews](docs/sdks/review/README.md#listperformancereviews) - List all reviews
* [PatchCommerceReview](docs/sdks/review/README.md#patchcommercereview) - Update a review
* [RemoveCommerceReview](docs/sdks/review/README.md#removecommercereview) - Remove a review
* [UpdateCommerceReview](docs/sdks/review/README.md#updatecommercereview) - Update a review

### [Saleschannel](docs/sdks/saleschannel/README.md)

* [CreateCommerceSaleschannel](docs/sdks/saleschannel/README.md#createcommercesaleschannel) - Create a saleschannel
* [GetCommerceSaleschannel](docs/sdks/saleschannel/README.md#getcommercesaleschannel) - Retrieve a saleschannel
* [ListCommerceSaleschannels](docs/sdks/saleschannel/README.md#listcommercesaleschannels) - List all saleschannels
* [PatchCommerceSaleschannel](docs/sdks/saleschannel/README.md#patchcommercesaleschannel) - Update a saleschannel
* [RemoveCommerceSaleschannel](docs/sdks/saleschannel/README.md#removecommercesaleschannel) - Remove a saleschannel
* [UpdateCommerceSaleschannel](docs/sdks/saleschannel/README.md#updatecommercesaleschannel) - Update a saleschannel

### [Salesorder](docs/sdks/salesorder/README.md)

* [CreateAccountingSalesorder](docs/sdks/salesorder/README.md#createaccountingsalesorder) - Create a salesorder
* [GetAccountingSalesorder](docs/sdks/salesorder/README.md#getaccountingsalesorder) - Retrieve a salesorder
* [ListAccountingSalesorders](docs/sdks/salesorder/README.md#listaccountingsalesorders) - List all salesorders
* [PatchAccountingSalesorder](docs/sdks/salesorder/README.md#patchaccountingsalesorder) - Update a salesorder
* [RemoveAccountingSalesorder](docs/sdks/salesorder/README.md#removeaccountingsalesorder) - Remove a salesorder
* [UpdateAccountingSalesorder](docs/sdks/salesorder/README.md#updateaccountingsalesorder) - Update a salesorder

### [Saml](docs/sdks/saml/README.md)

* [GetUnifiedIntegrationSaml](docs/sdks/saml/README.md#getunifiedintegrationsaml) - Sign in a user via SAML

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

* [CreateAtsScorecard](docs/sdks/scorecard/README.md#createatsscorecard) - Create a scorecard
* [GetAtsScorecard](docs/sdks/scorecard/README.md#getatsscorecard) - Retrieve a scorecard
* [ListAtsScorecards](docs/sdks/scorecard/README.md#listatsscorecards) - List all scorecards
* [PatchAtsScorecard](docs/sdks/scorecard/README.md#patchatsscorecard) - Update a scorecard
* [RemoveAtsScorecard](docs/sdks/scorecard/README.md#removeatsscorecard) - Remove a scorecard
* [UpdateAtsScorecard](docs/sdks/scorecard/README.md#updateatsscorecard) - Update a scorecard

### [Secretsmanager](docs/sdks/secretsmanager/README.md)

* [CreateUnifiedWorkspaceSecretsmanager](docs/sdks/secretsmanager/README.md#createunifiedworkspacesecretsmanager) - Create secrets manager
* [GetUnifiedWorkspaceSecretsmanager](docs/sdks/secretsmanager/README.md#getunifiedworkspacesecretsmanager) - Retrieve secrets manager
* [ListUnifiedWorkspaceSecretsmanagers](docs/sdks/secretsmanager/README.md#listunifiedworkspacesecretsmanagers) - List secrets managers
* [RemoveUnifiedWorkspaceSecretsmanager](docs/sdks/secretsmanager/README.md#removeunifiedworkspacesecretsmanager) - Remove secrets manager

### [Segment](docs/sdks/segment/README.md)

* [CreateCdpSegment](docs/sdks/segment/README.md#createcdpsegment) - Create a segment
* [GetCdpSegment](docs/sdks/segment/README.md#getcdpsegment) - Retrieve a segment
* [ListCdpSegments](docs/sdks/segment/README.md#listcdpsegments) - List all segments
* [PatchCdpSegment](docs/sdks/segment/README.md#patchcdpsegment) - Update a segment
* [RemoveCdpSegment](docs/sdks/segment/README.md#removecdpsegment) - Remove a segment
* [UpdateCdpSegment](docs/sdks/segment/README.md#updatecdpsegment) - Update a segment

### [Session](docs/sdks/session/README.md)

* [GetAnalyticsSession](docs/sdks/session/README.md#getanalyticssession) - Retrieve a session
* [ListAnalyticsSessions](docs/sdks/session/README.md#listanalyticssessions) - List all sessions

### [Shipment](docs/sdks/shipment/README.md)

* [CreateShippingShipment](docs/sdks/shipment/README.md#createshippingshipment) - Create a shipment
* [GetShippingShipment](docs/sdks/shipment/README.md#getshippingshipment) - Retrieve a shipment
* [ListShippingShipments](docs/sdks/shipment/README.md#listshippingshipments) - List all shipments
* [PatchShippingShipment](docs/sdks/shipment/README.md#patchshippingshipment) - Update a shipment
* [RemoveShippingShipment](docs/sdks/shipment/README.md#removeshippingshipment) - Remove a shipment
* [UpdateShippingShipment](docs/sdks/shipment/README.md#updateshippingshipment) - Update a shipment

### [Shipping](docs/sdks/shipping/README.md)

* [CreateShippingLabel](docs/sdks/shipping/README.md#createshippinglabel) - Create a label
* [CreateShippingRate](docs/sdks/shipping/README.md#createshippingrate) - Create a rate
* [CreateShippingShipment](docs/sdks/shipping/README.md#createshippingshipment) - Create a shipment
* [GetShippingCarrier](docs/sdks/shipping/README.md#getshippingcarrier) - Retrieve a carrier
* [GetShippingLabel](docs/sdks/shipping/README.md#getshippinglabel) - Retrieve a label
* [GetShippingShipment](docs/sdks/shipping/README.md#getshippingshipment) - Retrieve a shipment
* [GetShippingTracking](docs/sdks/shipping/README.md#getshippingtracking) - Retrieve a tracking
* [ListShippingCarriers](docs/sdks/shipping/README.md#listshippingcarriers) - List all carriers
* [ListShippingLabels](docs/sdks/shipping/README.md#listshippinglabels) - List all labels
* [ListShippingShipments](docs/sdks/shipping/README.md#listshippingshipments) - List all shipments
* [ListShippingTrackings](docs/sdks/shipping/README.md#listshippingtrackings) - List all trackings
* [PatchShippingLabel](docs/sdks/shipping/README.md#patchshippinglabel) - Update a label
* [PatchShippingShipment](docs/sdks/shipping/README.md#patchshippingshipment) - Update a shipment
* [RemoveShippingLabel](docs/sdks/shipping/README.md#removeshippinglabel) - Remove a label
* [RemoveShippingShipment](docs/sdks/shipping/README.md#removeshippingshipment) - Remove a shipment
* [UpdateShippingLabel](docs/sdks/shipping/README.md#updateshippinglabel) - Update a label
* [UpdateShippingShipment](docs/sdks/shipping/README.md#updateshippingshipment) - Update a shipment

### [Signatory](docs/sdks/signatory/README.md)

* [CreateSigningSignatory](docs/sdks/signatory/README.md#createsigningsignatory) - Create a signatory
* [GetSigningSignatory](docs/sdks/signatory/README.md#getsigningsignatory) - Retrieve a signatory
* [ListSigningSignatories](docs/sdks/signatory/README.md#listsigningsignatories) - List all signatories
* [PatchSigningSignatory](docs/sdks/signatory/README.md#patchsigningsignatory) - Update a signatory
* [RemoveSigningSignatory](docs/sdks/signatory/README.md#removesigningsignatory) - Remove a signatory
* [UpdateSigningSignatory](docs/sdks/signatory/README.md#updatesigningsignatory) - Update a signatory

### [Signing](docs/sdks/signing/README.md)

* [CreateSigningDocument](docs/sdks/signing/README.md#createsigningdocument) - Create a document
* [CreateSigningSignatory](docs/sdks/signing/README.md#createsigningsignatory) - Create a signatory
* [GetSigningDocument](docs/sdks/signing/README.md#getsigningdocument) - Retrieve a document
* [GetSigningSignatory](docs/sdks/signing/README.md#getsigningsignatory) - Retrieve a signatory
* [GetSigningTemplate](docs/sdks/signing/README.md#getsigningtemplate) - Retrieve a template
* [ListSigningDocuments](docs/sdks/signing/README.md#listsigningdocuments) - List all documents
* [ListSigningSignatories](docs/sdks/signing/README.md#listsigningsignatories) - List all signatories
* [ListSigningTemplates](docs/sdks/signing/README.md#listsigningtemplates) - List all templates
* [PatchSigningDocument](docs/sdks/signing/README.md#patchsigningdocument) - Update a document
* [PatchSigningSignatory](docs/sdks/signing/README.md#patchsigningsignatory) - Update a signatory
* [RemoveSigningDocument](docs/sdks/signing/README.md#removesigningdocument) - Remove a document
* [RemoveSigningSignatory](docs/sdks/signing/README.md#removesigningsignatory) - Remove a signatory
* [UpdateSigningDocument](docs/sdks/signing/README.md#updatesigningdocument) - Update a document
* [UpdateSigningSignatory](docs/sdks/signing/README.md#updatesigningsignatory) - Update a signatory

### [Source](docs/sdks/source/README.md)

* [CreateCdpSource](docs/sdks/source/README.md#createcdpsource) - Create a source
* [GetCdpSource](docs/sdks/source/README.md#getcdpsource) - Retrieve a source
* [ListCdpSources](docs/sdks/source/README.md#listcdpsources) - List all sources
* [PatchCdpSource](docs/sdks/source/README.md#patchcdpsource) - Update a source
* [RemoveCdpSource](docs/sdks/source/README.md#removecdpsource) - Remove a source
* [UpdateCdpSource](docs/sdks/source/README.md#updatecdpsource) - Update a source

### [Space](docs/sdks/space/README.md)

* [CreateKmsSpace](docs/sdks/space/README.md#createkmsspace) - Create a space
* [GetKmsSpace](docs/sdks/space/README.md#getkmsspace) - Retrieve a space
* [ListKmsSpaces](docs/sdks/space/README.md#listkmsspaces) - List all spaces
* [PatchKmsSpace](docs/sdks/space/README.md#patchkmsspace) - Update a space
* [RemoveKmsSpace](docs/sdks/space/README.md#removekmsspace) - Remove a space
* [UpdateKmsSpace](docs/sdks/space/README.md#updatekmsspace) - Update a space

### [Storage](docs/sdks/storage/README.md)

* [CreateStorageFile](docs/sdks/storage/README.md#createstoragefile) - Create a file
* [GetStorageFile](docs/sdks/storage/README.md#getstoragefile) - Retrieve a file
* [ListStorageFiles](docs/sdks/storage/README.md#liststoragefiles) - List all files
* [PatchStorageFile](docs/sdks/storage/README.md#patchstoragefile) - Update a file
* [RemoveStorageFile](docs/sdks/storage/README.md#removestoragefile) - Remove a file
* [UpdateStorageFile](docs/sdks/storage/README.md#updatestoragefile) - Update a file

### [Student](docs/sdks/student/README.md)

* [CreateLmsStudent](docs/sdks/student/README.md#createlmsstudent) - Create a student
* [GetLmsStudent](docs/sdks/student/README.md#getlmsstudent) - Retrieve a student
* [ListLmsStudents](docs/sdks/student/README.md#listlmsstudents) - List all students
* [PatchLmsStudent](docs/sdks/student/README.md#patchlmsstudent) - Update a student
* [RemoveLmsStudent](docs/sdks/student/README.md#removelmsstudent) - Remove a student
* [UpdateLmsStudent](docs/sdks/student/README.md#updatelmsstudent) - Update a student

### [Submission](docs/sdks/submission/README.md)

* [GetFormsSubmission](docs/sdks/submission/README.md#getformssubmission) - Retrieve a submission
* [ListFormsSubmissions](docs/sdks/submission/README.md#listformssubmissions) - List all submissions

### [Subscription](docs/sdks/subscription/README.md)

* [CreatePaymentSubscription](docs/sdks/subscription/README.md#createpaymentsubscription) - Create a subscription
* [GetPaymentSubscription](docs/sdks/subscription/README.md#getpaymentsubscription) - Retrieve a subscription
* [ListPaymentSubscriptions](docs/sdks/subscription/README.md#listpaymentsubscriptions) - List all subscriptions
* [PatchPaymentSubscription](docs/sdks/subscription/README.md#patchpaymentsubscription) - Update a subscription
* [RemovePaymentSubscription](docs/sdks/subscription/README.md#removepaymentsubscription) - Remove a subscription
* [UpdatePaymentSubscription](docs/sdks/subscription/README.md#updatepaymentsubscription) - Update a subscription

### [Table](docs/sdks/table/README.md)

* [CreateDatastoreTable](docs/sdks/table/README.md#createdatastoretable) - Create a table
* [GetDatastoreTable](docs/sdks/table/README.md#getdatastoretable) - Retrieve a table
* [ListDatastoreTables](docs/sdks/table/README.md#listdatastoretables) - List all tables
* [PatchDatastoreTable](docs/sdks/table/README.md#patchdatastoretable) - Update a table
* [RemoveDatastoreTable](docs/sdks/table/README.md#removedatastoretable) - Remove a table
* [UpdateDatastoreTable](docs/sdks/table/README.md#updatedatastoretable) - Update a table

### [Target](docs/sdks/target/README.md)

* [GetAdsTarget](docs/sdks/target/README.md#getadstarget) - Retrieve a target
* [ListAdsTargets](docs/sdks/target/README.md#listadstargets) - List all targets

### [Task](docs/sdks/task/README.md)

* [CreateTaskComment](docs/sdks/task/README.md#createtaskcomment) - Create a comment
* [CreateTaskProject](docs/sdks/task/README.md#createtaskproject) - Create a project
* [CreateTaskTask](docs/sdks/task/README.md#createtasktask) - Create a task
* [GetTaskChange](docs/sdks/task/README.md#gettaskchange) - Retrieve a change
* [GetTaskComment](docs/sdks/task/README.md#gettaskcomment) - Retrieve a comment
* [GetTaskProject](docs/sdks/task/README.md#gettaskproject) - Retrieve a project
* [GetTaskTask](docs/sdks/task/README.md#gettasktask) - Retrieve a task
* [ListTaskChanges](docs/sdks/task/README.md#listtaskchanges) - List all changes
* [ListTaskComments](docs/sdks/task/README.md#listtaskcomments) - List all comments
* [ListTaskProjects](docs/sdks/task/README.md#listtaskprojects) - List all projects
* [ListTaskTasks](docs/sdks/task/README.md#listtasktasks) - List all tasks
* [PatchTaskComment](docs/sdks/task/README.md#patchtaskcomment) - Update a comment
* [PatchTaskProject](docs/sdks/task/README.md#patchtaskproject) - Update a project
* [PatchTaskTask](docs/sdks/task/README.md#patchtasktask) - Update a task
* [RemoveTaskComment](docs/sdks/task/README.md#removetaskcomment) - Remove a comment
* [RemoveTaskProject](docs/sdks/task/README.md#removetaskproject) - Remove a project
* [RemoveTaskTask](docs/sdks/task/README.md#removetasktask) - Remove a task
* [UpdateTaskComment](docs/sdks/task/README.md#updatetaskcomment) - Update a comment
* [UpdateTaskProject](docs/sdks/task/README.md#updatetaskproject) - Update a project
* [UpdateTaskTask](docs/sdks/task/README.md#updatetasktask) - Update a task

### [Taxonomy](docs/sdks/taxonomy/README.md)

* [GetHrisTaxonomy](docs/sdks/taxonomy/README.md#gethristaxonomy) - Retrieve a taxonomy
* [ListHrisTaxonomies](docs/sdks/taxonomy/README.md#listhristaxonomies) - List all taxonomies

### [Taxrate](docs/sdks/taxrate/README.md)

* [CreateAccountingTaxrate](docs/sdks/taxrate/README.md#createaccountingtaxrate) - Create a taxrate
* [GetAccountingTaxrate](docs/sdks/taxrate/README.md#getaccountingtaxrate) - Retrieve a taxrate
* [ListAccountingTaxrates](docs/sdks/taxrate/README.md#listaccountingtaxrates) - List all taxrates
* [PatchAccountingTaxrate](docs/sdks/taxrate/README.md#patchaccountingtaxrate) - Update a taxrate
* [RemoveAccountingTaxrate](docs/sdks/taxrate/README.md#removeaccountingtaxrate) - Remove a taxrate
* [UpdateAccountingTaxrate](docs/sdks/taxrate/README.md#updateaccountingtaxrate) - Update a taxrate

### [Template](docs/sdks/template/README.md)

* [GetSigningTemplate](docs/sdks/template/README.md#getsigningtemplate) - Retrieve a template
* [ListSigningTemplates](docs/sdks/template/README.md#listsigningtemplates) - List all templates

### [Ticket](docs/sdks/ticket/README.md)

* [CreateTicketingTicket](docs/sdks/ticket/README.md#createticketingticket) - Create a ticket
* [GetTicketingTicket](docs/sdks/ticket/README.md#getticketingticket) - Retrieve a ticket
* [ListTicketingTickets](docs/sdks/ticket/README.md#listticketingtickets) - List all tickets
* [PatchTicketingTicket](docs/sdks/ticket/README.md#patchticketingticket) - Update a ticket
* [RemoveTicketingTicket](docs/sdks/ticket/README.md#removeticketingticket) - Remove a ticket
* [UpdateTicketingTicket](docs/sdks/ticket/README.md#updateticketingticket) - Update a ticket

### [Ticketing](docs/sdks/ticketing/README.md)

* [CreateTicketingCategory](docs/sdks/ticketing/README.md#createticketingcategory) - Create a category
* [CreateTicketingCustomer](docs/sdks/ticketing/README.md#createticketingcustomer) - Create a customer
* [CreateTicketingNote](docs/sdks/ticketing/README.md#createticketingnote) - Create a note
* [CreateTicketingTicket](docs/sdks/ticketing/README.md#createticketingticket) - Create a ticket
* [GetTicketingCategory](docs/sdks/ticketing/README.md#getticketingcategory) - Retrieve a category
* [GetTicketingCustomer](docs/sdks/ticketing/README.md#getticketingcustomer) - Retrieve a customer
* [GetTicketingNote](docs/sdks/ticketing/README.md#getticketingnote) - Retrieve a note
* [GetTicketingTicket](docs/sdks/ticketing/README.md#getticketingticket) - Retrieve a ticket
* [ListTicketingCategories](docs/sdks/ticketing/README.md#listticketingcategories) - List all categories
* [ListTicketingCustomers](docs/sdks/ticketing/README.md#listticketingcustomers) - List all customers
* [ListTicketingNotes](docs/sdks/ticketing/README.md#listticketingnotes) - List all notes
* [ListTicketingTickets](docs/sdks/ticketing/README.md#listticketingtickets) - List all tickets
* [PatchTicketingCategory](docs/sdks/ticketing/README.md#patchticketingcategory) - Update a category
* [PatchTicketingCustomer](docs/sdks/ticketing/README.md#patchticketingcustomer) - Update a customer
* [PatchTicketingNote](docs/sdks/ticketing/README.md#patchticketingnote) - Update a note
* [PatchTicketingTicket](docs/sdks/ticketing/README.md#patchticketingticket) - Update a ticket
* [RemoveTicketingCategory](docs/sdks/ticketing/README.md#removeticketingcategory) - Remove a category
* [RemoveTicketingCustomer](docs/sdks/ticketing/README.md#removeticketingcustomer) - Remove a customer
* [RemoveTicketingNote](docs/sdks/ticketing/README.md#removeticketingnote) - Remove a note
* [RemoveTicketingTicket](docs/sdks/ticketing/README.md#removeticketingticket) - Remove a ticket
* [UpdateTicketingCategory](docs/sdks/ticketing/README.md#updateticketingcategory) - Update a category
* [UpdateTicketingCustomer](docs/sdks/ticketing/README.md#updateticketingcustomer) - Update a customer
* [UpdateTicketingNote](docs/sdks/ticketing/README.md#updateticketingnote) - Update a note
* [UpdateTicketingTicket](docs/sdks/ticketing/README.md#updateticketingticket) - Update a ticket

### [Timeoff](docs/sdks/timeoff/README.md)

* [CreateHrisTimeoff](docs/sdks/timeoff/README.md#createhristimeoff) - Create a timeoff
* [GetHrisTimeoff](docs/sdks/timeoff/README.md#gethristimeoff) - Retrieve a timeoff
* [ListHrisTimeoffs](docs/sdks/timeoff/README.md#listhristimeoffs) - List all timeoffs
* [PatchHrisTimeoff](docs/sdks/timeoff/README.md#patchhristimeoff) - Update a timeoff
* [RemoveHrisTimeoff](docs/sdks/timeoff/README.md#removehristimeoff) - Remove a timeoff
* [UpdateHrisTimeoff](docs/sdks/timeoff/README.md#updatehristimeoff) - Update a timeoff

### [Timeshift](docs/sdks/timeshift/README.md)

* [CreateHrisTimeshift](docs/sdks/timeshift/README.md#createhristimeshift) - Create a timeshift
* [GetHrisTimeshift](docs/sdks/timeshift/README.md#gethristimeshift) - Retrieve a timeshift
* [ListHrisTimeshifts](docs/sdks/timeshift/README.md#listhristimeshifts) - List all timeshifts
* [PatchHrisTimeshift](docs/sdks/timeshift/README.md#patchhristimeshift) - Update a timeshift
* [RemoveHrisTimeshift](docs/sdks/timeshift/README.md#removehristimeshift) - Remove a timeshift
* [UpdateHrisTimeshift](docs/sdks/timeshift/README.md#updatehristimeshift) - Update a timeshift

### [Tracking](docs/sdks/tracking/README.md)

* [GetShippingTracking](docs/sdks/tracking/README.md#getshippingtracking) - Retrieve a tracking
* [ListShippingTrackings](docs/sdks/tracking/README.md#listshippingtrackings) - List all trackings

### [Transaction](docs/sdks/transaction/README.md)

* [CreateAccountingTransaction](docs/sdks/transaction/README.md#createaccountingtransaction) - Create a transaction
* [GetAccountingTransaction](docs/sdks/transaction/README.md#getaccountingtransaction) - Retrieve a transaction
* [ListAccountingTransactions](docs/sdks/transaction/README.md#listaccountingtransactions) - List all transactions
* [PatchAccountingTransaction](docs/sdks/transaction/README.md#patchaccountingtransaction) - Update a transaction
* [RemoveAccountingTransaction](docs/sdks/transaction/README.md#removeaccountingtransaction) - Remove a transaction
* [UpdateAccountingTransaction](docs/sdks/transaction/README.md#updateaccountingtransaction) - Update a transaction

### [Trialbalance](docs/sdks/trialbalance/README.md)

* [GetAccountingTrialbalance](docs/sdks/trialbalance/README.md#getaccountingtrialbalance) - Retrieve a trialbalance
* [ListAccountingTrialbalances](docs/sdks/trialbalance/README.md#listaccountingtrialbalances) - List all trialbalances

### [Uc](docs/sdks/uc/README.md)

* [CreateUcComment](docs/sdks/uc/README.md#createuccomment) - Create a comment
* [CreateUcContact](docs/sdks/uc/README.md#createuccontact) - Create a contact
* [CreateUcRecording](docs/sdks/uc/README.md#createucrecording) - Create a recording
* [GetUcCall](docs/sdks/uc/README.md#getuccall) - Retrieve a call
* [GetUcComment](docs/sdks/uc/README.md#getuccomment) - Retrieve a comment
* [GetUcContact](docs/sdks/uc/README.md#getuccontact) - Retrieve a contact
* [GetUcRecording](docs/sdks/uc/README.md#getucrecording) - Retrieve a recording
* [ListUcCalls](docs/sdks/uc/README.md#listuccalls) - List all calls
* [ListUcComments](docs/sdks/uc/README.md#listuccomments) - List all comments
* [ListUcContacts](docs/sdks/uc/README.md#listuccontacts) - List all contacts
* [ListUcRecordings](docs/sdks/uc/README.md#listucrecordings) - List all recordings
* [PatchUcComment](docs/sdks/uc/README.md#patchuccomment) - Update a comment
* [PatchUcContact](docs/sdks/uc/README.md#patchuccontact) - Update a contact
* [PatchUcRecording](docs/sdks/uc/README.md#patchucrecording) - Update a recording
* [RemoveUcComment](docs/sdks/uc/README.md#removeuccomment) - Remove a comment
* [RemoveUcContact](docs/sdks/uc/README.md#removeuccontact) - Remove a contact
* [RemoveUcRecording](docs/sdks/uc/README.md#removeucrecording) - Remove a recording
* [UpdateUcComment](docs/sdks/uc/README.md#updateuccomment) - Update a comment
* [UpdateUcContact](docs/sdks/uc/README.md#updateuccontact) - Update a contact
* [UpdateUcRecording](docs/sdks/uc/README.md#updateucrecording) - Update a recording

### [Unified](docs/sdks/unified/README.md)

* [CreateUnifiedConnection](docs/sdks/unified/README.md#createunifiedconnection) - Create connection
* [CreateUnifiedEnvironment](docs/sdks/unified/README.md#createunifiedenvironment) - Create new environments
* [CreateUnifiedWebhook](docs/sdks/unified/README.md#createunifiedwebhook) - Create webhook subscription
* [CreateUnifiedWorkspaceSecretsmanager](docs/sdks/unified/README.md#createunifiedworkspacesecretsmanager) - Create secrets manager
* [GetUnifiedApicall](docs/sdks/unified/README.md#getunifiedapicall) - Retrieve specific API Call by its ID
* [GetUnifiedConnection](docs/sdks/unified/README.md#getunifiedconnection) - Retrieve connection
* [GetUnifiedIntegrationAuth](docs/sdks/unified/README.md#getunifiedintegrationauth) - Authorize new connection
* [GetUnifiedIssue](docs/sdks/unified/README.md#getunifiedissue) - Retrieve support issue
* [GetUnifiedWebhook](docs/sdks/unified/README.md#getunifiedwebhook) - Retrieve webhook by its ID
* [GetUnifiedWorkspaceSecretsmanager](docs/sdks/unified/README.md#getunifiedworkspacesecretsmanager) - Retrieve secrets manager
* [ListUnifiedApicalls](docs/sdks/unified/README.md#listunifiedapicalls) - Returns API Calls
* [ListUnifiedConnections](docs/sdks/unified/README.md#listunifiedconnections) - List all connections
* [ListUnifiedEnvironments](docs/sdks/unified/README.md#listunifiedenvironments) - Returns all environments
* [ListUnifiedIntegrationWorkspaces](docs/sdks/unified/README.md#listunifiedintegrationworkspaces) - Returns all activated integrations in a workspace
* [ListUnifiedIntegrations](docs/sdks/unified/README.md#listunifiedintegrations) - Returns all integrations
* [ListUnifiedIssues](docs/sdks/unified/README.md#listunifiedissues) - List support issues
* [ListUnifiedWebhooks](docs/sdks/unified/README.md#listunifiedwebhooks) - Returns all registered webhooks
* [ListUnifiedWorkspaceSecretsmanagers](docs/sdks/unified/README.md#listunifiedworkspacesecretsmanagers) - List secrets managers
* [PatchUnifiedConnection](docs/sdks/unified/README.md#patchunifiedconnection) - Update connection
* [PatchUnifiedWebhook](docs/sdks/unified/README.md#patchunifiedwebhook) - Update webhook subscription
* [PatchUnifiedWebhookTrigger](docs/sdks/unified/README.md#patchunifiedwebhooktrigger) - Trigger webhook
* [RemoveUnifiedConnection](docs/sdks/unified/README.md#removeunifiedconnection) - Remove connection
* [RemoveUnifiedEnvironment](docs/sdks/unified/README.md#removeunifiedenvironment) - Remove an environment
* [RemoveUnifiedWebhook](docs/sdks/unified/README.md#removeunifiedwebhook) - Remove webhook subscription
* [RemoveUnifiedWorkspaceSecretsmanager](docs/sdks/unified/README.md#removeunifiedworkspacesecretsmanager) - Remove secrets manager
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

### [Vendorcredit](docs/sdks/vendorcredit/README.md)

* [CreateAccountingVendorcredit](docs/sdks/vendorcredit/README.md#createaccountingvendorcredit) - Create a vendorcredit
* [GetAccountingVendorcredit](docs/sdks/vendorcredit/README.md#getaccountingvendorcredit) - Retrieve a vendorcredit
* [ListAccountingVendorcredits](docs/sdks/vendorcredit/README.md#listaccountingvendorcredits) - List all vendorcredits
* [PatchAccountingVendorcredit](docs/sdks/vendorcredit/README.md#patchaccountingvendorcredit) - Update a vendorcredit
* [RemoveAccountingVendorcredit](docs/sdks/vendorcredit/README.md#removeaccountingvendorcredit) - Remove a vendorcredit
* [UpdateAccountingVendorcredit](docs/sdks/vendorcredit/README.md#updateaccountingvendorcredit) - Update a vendorcredit

### [Verification](docs/sdks/verification/README.md)

* [CreateVerificationRequest](docs/sdks/verification/README.md#createverificationrequest) - Create a request
* [GetVerificationPackage](docs/sdks/verification/README.md#getverificationpackage) - Retrieve a package
* [GetVerificationRequest](docs/sdks/verification/README.md#getverificationrequest) - Retrieve a request
* [ListVerificationPackages](docs/sdks/verification/README.md#listverificationpackages) - List all packages
* [ListVerificationRequests](docs/sdks/verification/README.md#listverificationrequests) - List all requests
* [PatchVerificationRequest](docs/sdks/verification/README.md#patchverificationrequest) - Update a request
* [RemoveVerificationRequest](docs/sdks/verification/README.md#removeverificationrequest) - Remove a request
* [UpdateVerificationRequest](docs/sdks/verification/README.md#updateverificationrequest) - Update a request

### [Visitor](docs/sdks/visitor/README.md)

* [CreateAnalyticsVisitor](docs/sdks/visitor/README.md#createanalyticsvisitor) - Create a visitor
* [GetAnalyticsVisitor](docs/sdks/visitor/README.md#getanalyticsvisitor) - Retrieve a visitor
* [ListAnalyticsVisitors](docs/sdks/visitor/README.md#listanalyticsvisitors) - List all visitors
* [PatchAnalyticsVisitor](docs/sdks/visitor/README.md#patchanalyticsvisitor) - Update a visitor
* [RemoveAnalyticsVisitor](docs/sdks/visitor/README.md#removeanalyticsvisitor) - Remove a visitor
* [UpdateAnalyticsVisitor](docs/sdks/visitor/README.md#updateanalyticsvisitor) - Update a visitor

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

* [CreateCalendarWebinar](docs/sdks/webinar/README.md#createcalendarwebinar) - Create a webinar
* [GetCalendarWebinar](docs/sdks/webinar/README.md#getcalendarwebinar) - Retrieve a webinar
* [ListCalendarWebinars](docs/sdks/webinar/README.md#listcalendarwebinars) - List all webinars
* [PatchCalendarWebinar](docs/sdks/webinar/README.md#patchcalendarwebinar) - Update a webinar
* [RemoveCalendarWebinar](docs/sdks/webinar/README.md#removecalendarwebinar) - Remove a webinar
* [UpdateCalendarWebinar](docs/sdks/webinar/README.md#updatecalendarwebinar) - Update a webinar

</details>
<!-- End Available Resources and Operations [operations] -->









<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `CreateAccountingAccount` function may return the following errors:

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

	res, err := s.Accounting.CreateAccountingAccount(ctx, operations.CreateAccountingAccountRequest{
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

	res, err := s.Accounting.CreateAccountingAccount(ctx, operations.CreateAccountingAccountRequest{
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

	res, err := s.Accounting.CreateAccountingAccount(ctx, operations.CreateAccountingAccountRequest{
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

	res, err := s.Accounting.CreateAccountingAccount(ctx, operations.CreateAccountingAccountRequest{
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

	res, err := s.Accounting.CreateAccountingAccount(ctx, operations.CreateAccountingAccountRequest{
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

	res, err := s.Accounting.CreateAccountingAccount(ctx, operations.CreateAccountingAccountRequest{
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
