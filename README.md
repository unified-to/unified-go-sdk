# github.com/unified-to/unified-go-sdk

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://github.com/unified-to/unified-go-sdk.git/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/bolt-php/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
    
</div>

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/unified-to/unified-go-sdk
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Agent.DeleteTicketingConnectionIDAgentID(ctx, operations.DeleteTicketingConnectionIDAgentIDRequest{
        ConnectionID: "navigate",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Agent](docs/sdks/agent/README.md)

* [DeleteTicketingConnectionIDAgentID](docs/sdks/agent/README.md#deleteticketingconnectionidagentid) - Remove a agent
* [GetTicketingConnectionIDAgent](docs/sdks/agent/README.md#getticketingconnectionidagent) - List all agents
* [GetTicketingConnectionIDAgentID](docs/sdks/agent/README.md#getticketingconnectionidagentid) - Retrieve a agent
* [GetUcConnectionIDAgent](docs/sdks/agent/README.md#getucconnectionidagent) - List all agents
* [PatchTicketingConnectionIDAgentID](docs/sdks/agent/README.md#patchticketingconnectionidagentid) - Update a agent
* [PostTicketingConnectionIDAgent](docs/sdks/agent/README.md#postticketingconnectionidagent) - Create a agent
* [PutTicketingConnectionIDAgentID](docs/sdks/agent/README.md#putticketingconnectionidagentid) - Update a agent

### [Apicall](docs/sdks/apicall/README.md)

* [GetUnifiedApicall](docs/sdks/apicall/README.md#getunifiedapicall) - Returns API Calls
* [GetUnifiedApicallID](docs/sdks/apicall/README.md#getunifiedapicallid) - Retrieve specific API Call by its ID

### [Application](docs/sdks/application/README.md)

* [DeleteAtsConnectionIDApplicationID](docs/sdks/application/README.md#deleteatsconnectionidapplicationid) - Remove an application
* [GetAtsConnectionIDApplication](docs/sdks/application/README.md#getatsconnectionidapplication) - List all applications
* [GetAtsConnectionIDApplicationID](docs/sdks/application/README.md#getatsconnectionidapplicationid) - Retrieve an application
* [PatchAtsConnectionIDApplicationID](docs/sdks/application/README.md#patchatsconnectionidapplicationid) - Update an application
* [PostAtsConnectionIDApplication](docs/sdks/application/README.md#postatsconnectionidapplication) - Create an application
* [PutAtsConnectionIDApplicationID](docs/sdks/application/README.md#putatsconnectionidapplicationid) - Update an application

### [Ats](docs/sdks/ats/README.md)

* [DeleteAtsConnectionIDApplicationID](docs/sdks/ats/README.md#deleteatsconnectionidapplicationid) - Remove an application
* [DeleteAtsConnectionIDCandidateID](docs/sdks/ats/README.md#deleteatsconnectionidcandidateid) - Remove a candidate
* [DeleteAtsConnectionIDInterviewID](docs/sdks/ats/README.md#deleteatsconnectionidinterviewid) - Remove a interview
* [DeleteAtsConnectionIDJobID](docs/sdks/ats/README.md#deleteatsconnectionidjobid) - Remove a job
* [DeleteAtsConnectionIDScorecardID](docs/sdks/ats/README.md#deleteatsconnectionidscorecardid) - Remove a scorecard
* [GetAtsConnectionIDApplication](docs/sdks/ats/README.md#getatsconnectionidapplication) - List all applications
* [GetAtsConnectionIDApplicationID](docs/sdks/ats/README.md#getatsconnectionidapplicationid) - Retrieve an application
* [GetAtsConnectionIDCandidate](docs/sdks/ats/README.md#getatsconnectionidcandidate) - List all candidates
* [GetAtsConnectionIDCandidateID](docs/sdks/ats/README.md#getatsconnectionidcandidateid) - Retrieve a candidate
* [GetAtsConnectionIDInterview](docs/sdks/ats/README.md#getatsconnectionidinterview) - List all interviews
* [GetAtsConnectionIDInterviewID](docs/sdks/ats/README.md#getatsconnectionidinterviewid) - Retrieve a interview
* [GetAtsConnectionIDJob](docs/sdks/ats/README.md#getatsconnectionidjob) - List all jobs
* [GetAtsConnectionIDJobID](docs/sdks/ats/README.md#getatsconnectionidjobid) - Retrieve a job
* [GetAtsConnectionIDScorecard](docs/sdks/ats/README.md#getatsconnectionidscorecard) - List all scorecards
* [GetAtsConnectionIDScorecardID](docs/sdks/ats/README.md#getatsconnectionidscorecardid) - Retrieve a scorecard
* [PatchAtsConnectionIDApplicationID](docs/sdks/ats/README.md#patchatsconnectionidapplicationid) - Update an application
* [PatchAtsConnectionIDCandidateID](docs/sdks/ats/README.md#patchatsconnectionidcandidateid) - Update a candidate
* [PatchAtsConnectionIDInterviewID](docs/sdks/ats/README.md#patchatsconnectionidinterviewid) - Update a interview
* [PatchAtsConnectionIDJobID](docs/sdks/ats/README.md#patchatsconnectionidjobid) - Update a job
* [PatchAtsConnectionIDScorecardID](docs/sdks/ats/README.md#patchatsconnectionidscorecardid) - Update a scorecard
* [PostAtsConnectionIDApplication](docs/sdks/ats/README.md#postatsconnectionidapplication) - Create an application
* [PostAtsConnectionIDCandidate](docs/sdks/ats/README.md#postatsconnectionidcandidate) - Create a candidate
* [PostAtsConnectionIDInterview](docs/sdks/ats/README.md#postatsconnectionidinterview) - Create a interview
* [PostAtsConnectionIDJob](docs/sdks/ats/README.md#postatsconnectionidjob) - Create a job
* [PostAtsConnectionIDScorecard](docs/sdks/ats/README.md#postatsconnectionidscorecard) - Create a scorecard
* [PutAtsConnectionIDApplicationID](docs/sdks/ats/README.md#putatsconnectionidapplicationid) - Update an application
* [PutAtsConnectionIDCandidateID](docs/sdks/ats/README.md#putatsconnectionidcandidateid) - Update a candidate
* [PutAtsConnectionIDInterviewID](docs/sdks/ats/README.md#putatsconnectionidinterviewid) - Update a interview
* [PutAtsConnectionIDJobID](docs/sdks/ats/README.md#putatsconnectionidjobid) - Update a job
* [PutAtsConnectionIDScorecardID](docs/sdks/ats/README.md#putatsconnectionidscorecardid) - Update a scorecard

### [Auth](docs/sdks/auth/README.md)

* [GetUnifiedIntegrationAuthWorkspaceIDIntegrationType](docs/sdks/auth/README.md#getunifiedintegrationauthworkspaceidintegrationtype) - Create connection indirectly
* [GetUnifiedIntegrationLoginWorkspaceIDIntegrationType](docs/sdks/auth/README.md#getunifiedintegrationloginworkspaceidintegrationtype) - Sign in a user

### [Call](docs/sdks/call/README.md)

* [GetUcConnectionIDCall](docs/sdks/call/README.md#getucconnectionidcall) - List all calls

### [Candidate](docs/sdks/candidate/README.md)

* [DeleteAtsConnectionIDCandidateID](docs/sdks/candidate/README.md#deleteatsconnectionidcandidateid) - Remove a candidate
* [GetAtsConnectionIDCandidate](docs/sdks/candidate/README.md#getatsconnectionidcandidate) - List all candidates
* [GetAtsConnectionIDCandidateID](docs/sdks/candidate/README.md#getatsconnectionidcandidateid) - Retrieve a candidate
* [PatchAtsConnectionIDCandidateID](docs/sdks/candidate/README.md#patchatsconnectionidcandidateid) - Update a candidate
* [PostAtsConnectionIDCandidate](docs/sdks/candidate/README.md#postatsconnectionidcandidate) - Create a candidate
* [PutAtsConnectionIDCandidateID](docs/sdks/candidate/README.md#putatsconnectionidcandidateid) - Update a candidate

### [Company](docs/sdks/company/README.md)

* [DeleteCrmConnectionIDCompanyID](docs/sdks/company/README.md#deletecrmconnectionidcompanyid) - Remove a company
* [DeleteCrmConnectionIDCompanyIDDealDealID](docs/sdks/company/README.md#deletecrmconnectionidcompanyiddealdealid) - Remove deal association from a company
* [GetCrmConnectionIDCompany](docs/sdks/company/README.md#getcrmconnectionidcompany) - List all companies
* [GetCrmConnectionIDCompanyID](docs/sdks/company/README.md#getcrmconnectionidcompanyid) - Retrieve a company
* [GetEnrichConnectionIDCompany](docs/sdks/company/README.md#getenrichconnectionidcompany) - Retrieve enrichment information for a company
* [PatchCrmConnectionIDCompanyID](docs/sdks/company/README.md#patchcrmconnectionidcompanyid) - Update a company
* [PatchCrmConnectionIDCompanyIDDealDealID](docs/sdks/company/README.md#patchcrmconnectionidcompanyiddealdealid) - Associate a deal with a company
* [PostCrmConnectionIDCompany](docs/sdks/company/README.md#postcrmconnectionidcompany) - Create a company
* [PutCrmConnectionIDCompanyID](docs/sdks/company/README.md#putcrmconnectionidcompanyid) - Update a company
* [PutCrmConnectionIDCompanyIDDealDealID](docs/sdks/company/README.md#putcrmconnectionidcompanyiddealdealid) - Associate a deal with a company

### [Connection](docs/sdks/connection/README.md)

* [DeleteUnifiedConnectionID](docs/sdks/connection/README.md#deleteunifiedconnectionid) - Remove connection
* [GetUnifiedConnection](docs/sdks/connection/README.md#getunifiedconnection) - List all connections
* [GetUnifiedConnectionID](docs/sdks/connection/README.md#getunifiedconnectionid) - Retrieve connection
* [PatchUnifiedConnectionID](docs/sdks/connection/README.md#patchunifiedconnectionid) - Update connection
* [PostUnifiedConnection](docs/sdks/connection/README.md#postunifiedconnection) - Create connection
* [PutUnifiedConnectionID](docs/sdks/connection/README.md#putunifiedconnectionid) - Update connection

### [Contact](docs/sdks/contact/README.md)

* [DeleteCrmConnectionIDContactID](docs/sdks/contact/README.md#deletecrmconnectionidcontactid) - Remove a contact
* [DeleteCrmConnectionIDContactIDCompanyCompanyID](docs/sdks/contact/README.md#deletecrmconnectionidcontactidcompanycompanyid) - Remove company association from a contact
* [DeleteCrmConnectionIDContactIDDealDealID](docs/sdks/contact/README.md#deletecrmconnectionidcontactiddealdealid) - Remove deal association from a contact
* [DeleteUcConnectionIDContactID](docs/sdks/contact/README.md#deleteucconnectionidcontactid) - Remove a contact
* [GetCrmConnectionIDContact](docs/sdks/contact/README.md#getcrmconnectionidcontact) - List all contacts
* [GetCrmConnectionIDContactID](docs/sdks/contact/README.md#getcrmconnectionidcontactid) - Retrieve a contact
* [GetUcConnectionIDContact](docs/sdks/contact/README.md#getucconnectionidcontact) - List all contacts
* [GetUcConnectionIDContactID](docs/sdks/contact/README.md#getucconnectionidcontactid) - Retrieve a contact
* [PatchCrmConnectionIDContactID](docs/sdks/contact/README.md#patchcrmconnectionidcontactid) - Update a contact
* [PatchCrmConnectionIDContactIDCompanyCompanyID](docs/sdks/contact/README.md#patchcrmconnectionidcontactidcompanycompanyid) - Associate a company with a contact
* [PatchCrmConnectionIDContactIDDealDealID](docs/sdks/contact/README.md#patchcrmconnectionidcontactiddealdealid) - Associate a deal with a contact
* [PatchUcConnectionIDContactID](docs/sdks/contact/README.md#patchucconnectionidcontactid) - Update a contact
* [PostCrmConnectionIDContact](docs/sdks/contact/README.md#postcrmconnectionidcontact) - Create a contact
* [PostUcConnectionIDContact](docs/sdks/contact/README.md#postucconnectionidcontact) - Create a contact
* [PutCrmConnectionIDContactID](docs/sdks/contact/README.md#putcrmconnectionidcontactid) - Update a contact
* [PutCrmConnectionIDContactIDCompanyCompanyID](docs/sdks/contact/README.md#putcrmconnectionidcontactidcompanycompanyid) - Associate a company with a contact
* [PutCrmConnectionIDContactIDDealDealID](docs/sdks/contact/README.md#putcrmconnectionidcontactiddealdealid) - Associate a deal with a contact
* [PutUcConnectionIDContactID](docs/sdks/contact/README.md#putucconnectionidcontactid) - Update a contact

### [Crm](docs/sdks/crm/README.md)

* [DeleteCrmConnectionIDCompanyID](docs/sdks/crm/README.md#deletecrmconnectionidcompanyid) - Remove a company
* [DeleteCrmConnectionIDCompanyIDDealDealID](docs/sdks/crm/README.md#deletecrmconnectionidcompanyiddealdealid) - Remove deal association from a company
* [DeleteCrmConnectionIDContactID](docs/sdks/crm/README.md#deletecrmconnectionidcontactid) - Remove a contact
* [DeleteCrmConnectionIDContactIDCompanyCompanyID](docs/sdks/crm/README.md#deletecrmconnectionidcontactidcompanycompanyid) - Remove company association from a contact
* [DeleteCrmConnectionIDContactIDDealDealID](docs/sdks/crm/README.md#deletecrmconnectionidcontactiddealdealid) - Remove deal association from a contact
* [DeleteCrmConnectionIDDealID](docs/sdks/crm/README.md#deletecrmconnectioniddealid) - Remove a deal
* [DeleteCrmConnectionIDEventID](docs/sdks/crm/README.md#deletecrmconnectionideventid) - Remove a event
* [DeleteCrmConnectionIDEventIDCompanyCompanyID](docs/sdks/crm/README.md#deletecrmconnectionideventidcompanycompanyid) - Remove company association from an event
* [DeleteCrmConnectionIDEventIDContactContactID](docs/sdks/crm/README.md#deletecrmconnectionideventidcontactcontactid) - Remove contact association from an event
* [DeleteCrmConnectionIDEventIDDealDealID](docs/sdks/crm/README.md#deletecrmconnectionideventiddealdealid) - Remove deal association from an event
* [DeleteCrmConnectionIDFileID](docs/sdks/crm/README.md#deletecrmconnectionidfileid) - Remove a file
* [DeleteCrmConnectionIDLeadID](docs/sdks/crm/README.md#deletecrmconnectionidleadid) - Remove a lead
* [DeleteCrmConnectionIDPipelineID](docs/sdks/crm/README.md#deletecrmconnectionidpipelineid) - Remove a pipeline
* [DeleteCrmConnectionIDTeamID](docs/sdks/crm/README.md#deletecrmconnectionidteamid) - Remove a team
* [DeleteCrmConnectionIDUserID](docs/sdks/crm/README.md#deletecrmconnectioniduserid) - Remove a user
* [GetCrmConnectionIDCompany](docs/sdks/crm/README.md#getcrmconnectionidcompany) - List all companies
* [GetCrmConnectionIDCompanyID](docs/sdks/crm/README.md#getcrmconnectionidcompanyid) - Retrieve a company
* [GetCrmConnectionIDContact](docs/sdks/crm/README.md#getcrmconnectionidcontact) - List all contacts
* [GetCrmConnectionIDContactID](docs/sdks/crm/README.md#getcrmconnectionidcontactid) - Retrieve a contact
* [GetCrmConnectionIDDeal](docs/sdks/crm/README.md#getcrmconnectioniddeal) - List all deals
* [GetCrmConnectionIDDealID](docs/sdks/crm/README.md#getcrmconnectioniddealid) - Retrieve a deal
* [GetCrmConnectionIDEvent](docs/sdks/crm/README.md#getcrmconnectionidevent) - List all events
* [GetCrmConnectionIDEventID](docs/sdks/crm/README.md#getcrmconnectionideventid) - Retrieve a event
* [GetCrmConnectionIDFile](docs/sdks/crm/README.md#getcrmconnectionidfile) - List all files
* [GetCrmConnectionIDFileID](docs/sdks/crm/README.md#getcrmconnectionidfileid) - Retrieve a file
* [GetCrmConnectionIDLead](docs/sdks/crm/README.md#getcrmconnectionidlead) - List all leads
* [GetCrmConnectionIDLeadID](docs/sdks/crm/README.md#getcrmconnectionidleadid) - Retrieve a lead
* [GetCrmConnectionIDPipeline](docs/sdks/crm/README.md#getcrmconnectionidpipeline) - List all pipelines
* [GetCrmConnectionIDPipelineID](docs/sdks/crm/README.md#getcrmconnectionidpipelineid) - Retrieve a pipeline
* [GetCrmConnectionIDTeam](docs/sdks/crm/README.md#getcrmconnectionidteam) - List all teams
* [GetCrmConnectionIDTeamID](docs/sdks/crm/README.md#getcrmconnectionidteamid) - Retrieve a team
* [GetCrmConnectionIDUser](docs/sdks/crm/README.md#getcrmconnectioniduser) - List all users
* [GetCrmConnectionIDUserID](docs/sdks/crm/README.md#getcrmconnectioniduserid) - Retrieve a user
* [PatchCrmConnectionIDCompanyID](docs/sdks/crm/README.md#patchcrmconnectionidcompanyid) - Update a company
* [PatchCrmConnectionIDCompanyIDDealDealID](docs/sdks/crm/README.md#patchcrmconnectionidcompanyiddealdealid) - Associate a deal with a company
* [PatchCrmConnectionIDContactID](docs/sdks/crm/README.md#patchcrmconnectionidcontactid) - Update a contact
* [PatchCrmConnectionIDContactIDCompanyCompanyID](docs/sdks/crm/README.md#patchcrmconnectionidcontactidcompanycompanyid) - Associate a company with a contact
* [PatchCrmConnectionIDContactIDDealDealID](docs/sdks/crm/README.md#patchcrmconnectionidcontactiddealdealid) - Associate a deal with a contact
* [PatchCrmConnectionIDDealID](docs/sdks/crm/README.md#patchcrmconnectioniddealid) - Update a deal
* [PatchCrmConnectionIDEventID](docs/sdks/crm/README.md#patchcrmconnectionideventid) - Update a event
* [PatchCrmConnectionIDEventIDCompanyCompanyID](docs/sdks/crm/README.md#patchcrmconnectionideventidcompanycompanyid) - Associate a company with an event
* [PatchCrmConnectionIDEventIDContactContactID](docs/sdks/crm/README.md#patchcrmconnectionideventidcontactcontactid) - Associate a contact with an event
* [PatchCrmConnectionIDEventIDDealDealID](docs/sdks/crm/README.md#patchcrmconnectionideventiddealdealid) - Associate a deal with an event
* [PatchCrmConnectionIDFileID](docs/sdks/crm/README.md#patchcrmconnectionidfileid) - Update a file
* [PatchCrmConnectionIDLeadID](docs/sdks/crm/README.md#patchcrmconnectionidleadid) - Update a lead
* [PatchCrmConnectionIDPipelineID](docs/sdks/crm/README.md#patchcrmconnectionidpipelineid) - Update a pipeline
* [PatchCrmConnectionIDTeamID](docs/sdks/crm/README.md#patchcrmconnectionidteamid) - Update a team
* [PatchCrmConnectionIDUserID](docs/sdks/crm/README.md#patchcrmconnectioniduserid) - Update a user
* [PostCrmConnectionIDCompany](docs/sdks/crm/README.md#postcrmconnectionidcompany) - Create a company
* [PostCrmConnectionIDContact](docs/sdks/crm/README.md#postcrmconnectionidcontact) - Create a contact
* [PostCrmConnectionIDDeal](docs/sdks/crm/README.md#postcrmconnectioniddeal) - Create a deal
* [PostCrmConnectionIDEvent](docs/sdks/crm/README.md#postcrmconnectionidevent) - Create a event
* [PostCrmConnectionIDFile](docs/sdks/crm/README.md#postcrmconnectionidfile) - Create a file
* [PostCrmConnectionIDLead](docs/sdks/crm/README.md#postcrmconnectionidlead) - Create a lead
* [PostCrmConnectionIDPipeline](docs/sdks/crm/README.md#postcrmconnectionidpipeline) - Create a pipeline
* [PostCrmConnectionIDTeam](docs/sdks/crm/README.md#postcrmconnectionidteam) - Create a team
* [PostCrmConnectionIDUser](docs/sdks/crm/README.md#postcrmconnectioniduser) - Create a user
* [PutCrmConnectionIDCompanyID](docs/sdks/crm/README.md#putcrmconnectionidcompanyid) - Update a company
* [PutCrmConnectionIDCompanyIDDealDealID](docs/sdks/crm/README.md#putcrmconnectionidcompanyiddealdealid) - Associate a deal with a company
* [PutCrmConnectionIDContactID](docs/sdks/crm/README.md#putcrmconnectionidcontactid) - Update a contact
* [PutCrmConnectionIDContactIDCompanyCompanyID](docs/sdks/crm/README.md#putcrmconnectionidcontactidcompanycompanyid) - Associate a company with a contact
* [PutCrmConnectionIDContactIDDealDealID](docs/sdks/crm/README.md#putcrmconnectionidcontactiddealdealid) - Associate a deal with a contact
* [PutCrmConnectionIDDealID](docs/sdks/crm/README.md#putcrmconnectioniddealid) - Update a deal
* [PutCrmConnectionIDEventID](docs/sdks/crm/README.md#putcrmconnectionideventid) - Update a event
* [PutCrmConnectionIDEventIDCompanyCompanyID](docs/sdks/crm/README.md#putcrmconnectionideventidcompanycompanyid) - Associate a company with an event
* [PutCrmConnectionIDEventIDContactContactID](docs/sdks/crm/README.md#putcrmconnectionideventidcontactcontactid) - Associate a contact with an event
* [PutCrmConnectionIDEventIDDealDealID](docs/sdks/crm/README.md#putcrmconnectionideventiddealdealid) - Associate a deal with an event
* [PutCrmConnectionIDFileID](docs/sdks/crm/README.md#putcrmconnectionidfileid) - Update a file
* [PutCrmConnectionIDLeadID](docs/sdks/crm/README.md#putcrmconnectionidleadid) - Update a lead
* [PutCrmConnectionIDPipelineID](docs/sdks/crm/README.md#putcrmconnectionidpipelineid) - Update a pipeline
* [PutCrmConnectionIDTeamID](docs/sdks/crm/README.md#putcrmconnectionidteamid) - Update a team
* [PutCrmConnectionIDUserID](docs/sdks/crm/README.md#putcrmconnectioniduserid) - Update a user

### [Customer](docs/sdks/customer/README.md)

* [DeleteTicketingConnectionIDCustomerID](docs/sdks/customer/README.md#deleteticketingconnectionidcustomerid) - Remove a customer
* [GetTicketingConnectionIDCustomer](docs/sdks/customer/README.md#getticketingconnectionidcustomer) - List all customers
* [GetTicketingConnectionIDCustomerID](docs/sdks/customer/README.md#getticketingconnectionidcustomerid) - Retrieve a customer
* [PatchTicketingConnectionIDCustomerID](docs/sdks/customer/README.md#patchticketingconnectionidcustomerid) - Update a customer
* [PostTicketingConnectionIDCustomer](docs/sdks/customer/README.md#postticketingconnectionidcustomer) - Create a customer
* [PutTicketingConnectionIDCustomerID](docs/sdks/customer/README.md#putticketingconnectionidcustomerid) - Update a customer

### [Deal](docs/sdks/deal/README.md)

* [DeleteCrmConnectionIDDealID](docs/sdks/deal/README.md#deletecrmconnectioniddealid) - Remove a deal
* [GetCrmConnectionIDDeal](docs/sdks/deal/README.md#getcrmconnectioniddeal) - List all deals
* [GetCrmConnectionIDDealID](docs/sdks/deal/README.md#getcrmconnectioniddealid) - Retrieve a deal
* [PatchCrmConnectionIDDealID](docs/sdks/deal/README.md#patchcrmconnectioniddealid) - Update a deal
* [PostCrmConnectionIDDeal](docs/sdks/deal/README.md#postcrmconnectioniddeal) - Create a deal
* [PutCrmConnectionIDDealID](docs/sdks/deal/README.md#putcrmconnectioniddealid) - Update a deal

### [Document](docs/sdks/document/README.md)

* [DeleteAtsConnectionIDScorecardID](docs/sdks/document/README.md#deleteatsconnectionidscorecardid) - Remove a scorecard
* [GetAtsConnectionIDScorecard](docs/sdks/document/README.md#getatsconnectionidscorecard) - List all scorecards
* [GetAtsConnectionIDScorecardID](docs/sdks/document/README.md#getatsconnectionidscorecardid) - Retrieve a scorecard
* [PatchAtsConnectionIDScorecardID](docs/sdks/document/README.md#patchatsconnectionidscorecardid) - Update a scorecard
* [PostAtsConnectionIDScorecard](docs/sdks/document/README.md#postatsconnectionidscorecard) - Create a scorecard
* [PutAtsConnectionIDScorecardID](docs/sdks/document/README.md#putatsconnectionidscorecardid) - Update a scorecard

### [Employee](docs/sdks/employee/README.md)

* [DeleteHrisConnectionIDEmployeeID](docs/sdks/employee/README.md#deletehrisconnectionidemployeeid) - Remove a Employee
* [GetHrisConnectionIDEmployee](docs/sdks/employee/README.md#gethrisconnectionidemployee) - List all Employees
* [GetHrisConnectionIDEmployeeID](docs/sdks/employee/README.md#gethrisconnectionidemployeeid) - Retrieve a Employee
* [PatchHrisConnectionIDEmployeeID](docs/sdks/employee/README.md#patchhrisconnectionidemployeeid) - Update a Employee
* [PostHrisConnectionIDEmployee](docs/sdks/employee/README.md#posthrisconnectionidemployee) - Create a Employee
* [PutHrisConnectionIDEmployeeID](docs/sdks/employee/README.md#puthrisconnectionidemployeeid) - Update a Employee

### [Enrich](docs/sdks/enrich/README.md)

* [GetEnrichConnectionIDCompany](docs/sdks/enrich/README.md#getenrichconnectionidcompany) - Retrieve enrichment information for a company
* [GetEnrichConnectionIDPerson](docs/sdks/enrich/README.md#getenrichconnectionidperson) - Retrieve enrichment information for a person

### [Event](docs/sdks/event/README.md)

* [DeleteCrmConnectionIDEventID](docs/sdks/event/README.md#deletecrmconnectionideventid) - Remove a event
* [DeleteCrmConnectionIDEventIDCompanyCompanyID](docs/sdks/event/README.md#deletecrmconnectionideventidcompanycompanyid) - Remove company association from an event
* [DeleteCrmConnectionIDEventIDContactContactID](docs/sdks/event/README.md#deletecrmconnectionideventidcontactcontactid) - Remove contact association from an event
* [DeleteCrmConnectionIDEventIDDealDealID](docs/sdks/event/README.md#deletecrmconnectionideventiddealdealid) - Remove deal association from an event
* [GetCrmConnectionIDEvent](docs/sdks/event/README.md#getcrmconnectionidevent) - List all events
* [GetCrmConnectionIDEventID](docs/sdks/event/README.md#getcrmconnectionideventid) - Retrieve a event
* [PatchCrmConnectionIDEventID](docs/sdks/event/README.md#patchcrmconnectionideventid) - Update a event
* [PatchCrmConnectionIDEventIDCompanyCompanyID](docs/sdks/event/README.md#patchcrmconnectionideventidcompanycompanyid) - Associate a company with an event
* [PatchCrmConnectionIDEventIDContactContactID](docs/sdks/event/README.md#patchcrmconnectionideventidcontactcontactid) - Associate a contact with an event
* [PatchCrmConnectionIDEventIDDealDealID](docs/sdks/event/README.md#patchcrmconnectionideventiddealdealid) - Associate a deal with an event
* [PostCrmConnectionIDEvent](docs/sdks/event/README.md#postcrmconnectionidevent) - Create a event
* [PutCrmConnectionIDEventID](docs/sdks/event/README.md#putcrmconnectionideventid) - Update a event
* [PutCrmConnectionIDEventIDCompanyCompanyID](docs/sdks/event/README.md#putcrmconnectionideventidcompanycompanyid) - Associate a company with an event
* [PutCrmConnectionIDEventIDContactContactID](docs/sdks/event/README.md#putcrmconnectionideventidcontactcontactid) - Associate a contact with an event
* [PutCrmConnectionIDEventIDDealDealID](docs/sdks/event/README.md#putcrmconnectionideventiddealdealid) - Associate a deal with an event

### [File](docs/sdks/file/README.md)

* [DeleteCrmConnectionIDFileID](docs/sdks/file/README.md#deletecrmconnectionidfileid) - Remove a file
* [GetCrmConnectionIDFile](docs/sdks/file/README.md#getcrmconnectionidfile) - List all files
* [GetCrmConnectionIDFileID](docs/sdks/file/README.md#getcrmconnectionidfileid) - Retrieve a file
* [PatchCrmConnectionIDFileID](docs/sdks/file/README.md#patchcrmconnectionidfileid) - Update a file
* [PostCrmConnectionIDFile](docs/sdks/file/README.md#postcrmconnectionidfile) - Create a file
* [PutCrmConnectionIDFileID](docs/sdks/file/README.md#putcrmconnectionidfileid) - Update a file

### [Group](docs/sdks/group/README.md)

* [DeleteHrisConnectionIDGroupID](docs/sdks/group/README.md#deletehrisconnectionidgroupid) - Remove a Group
* [GetHrisConnectionIDGroup](docs/sdks/group/README.md#gethrisconnectionidgroup) - List all Groups
* [GetHrisConnectionIDGroupID](docs/sdks/group/README.md#gethrisconnectionidgroupid) - Retrieve a Group
* [PatchHrisConnectionIDGroupID](docs/sdks/group/README.md#patchhrisconnectionidgroupid) - Update a Group
* [PostHrisConnectionIDGroup](docs/sdks/group/README.md#posthrisconnectionidgroup) - Create a Group
* [PutHrisConnectionIDGroupID](docs/sdks/group/README.md#puthrisconnectionidgroupid) - Update a Group

### [Hris](docs/sdks/hris/README.md)

* [DeleteHrisConnectionIDEmployeeID](docs/sdks/hris/README.md#deletehrisconnectionidemployeeid) - Remove a Employee
* [DeleteHrisConnectionIDGroupID](docs/sdks/hris/README.md#deletehrisconnectionidgroupid) - Remove a Group
* [GetHrisConnectionIDEmployee](docs/sdks/hris/README.md#gethrisconnectionidemployee) - List all Employees
* [GetHrisConnectionIDEmployeeID](docs/sdks/hris/README.md#gethrisconnectionidemployeeid) - Retrieve a Employee
* [GetHrisConnectionIDGroup](docs/sdks/hris/README.md#gethrisconnectionidgroup) - List all Groups
* [GetHrisConnectionIDGroupID](docs/sdks/hris/README.md#gethrisconnectionidgroupid) - Retrieve a Group
* [PatchHrisConnectionIDEmployeeID](docs/sdks/hris/README.md#patchhrisconnectionidemployeeid) - Update a Employee
* [PatchHrisConnectionIDGroupID](docs/sdks/hris/README.md#patchhrisconnectionidgroupid) - Update a Group
* [PostHrisConnectionIDEmployee](docs/sdks/hris/README.md#posthrisconnectionidemployee) - Create a Employee
* [PostHrisConnectionIDGroup](docs/sdks/hris/README.md#posthrisconnectionidgroup) - Create a Group
* [PutHrisConnectionIDEmployeeID](docs/sdks/hris/README.md#puthrisconnectionidemployeeid) - Update a Employee
* [PutHrisConnectionIDGroupID](docs/sdks/hris/README.md#puthrisconnectionidgroupid) - Update a Group

### [Integration](docs/sdks/integration/README.md)

* [GetUnifiedIntegration](docs/sdks/integration/README.md#getunifiedintegration) - Returns all integrations
* [GetUnifiedIntegrationAuthWorkspaceIDIntegrationType](docs/sdks/integration/README.md#getunifiedintegrationauthworkspaceidintegrationtype) - Create connection indirectly
* [GetUnifiedIntegrationIntegrationType](docs/sdks/integration/README.md#getunifiedintegrationintegrationtype) - Retrieve an integration
* [GetUnifiedIntegrationWorkspaceWorkspaceID](docs/sdks/integration/README.md#getunifiedintegrationworkspaceworkspaceid) - Returns all activated integrations in a workspace

### [Interview](docs/sdks/interview/README.md)

* [DeleteAtsConnectionIDInterviewID](docs/sdks/interview/README.md#deleteatsconnectionidinterviewid) - Remove a interview
* [GetAtsConnectionIDInterview](docs/sdks/interview/README.md#getatsconnectionidinterview) - List all interviews
* [GetAtsConnectionIDInterviewID](docs/sdks/interview/README.md#getatsconnectionidinterviewid) - Retrieve a interview
* [PatchAtsConnectionIDInterviewID](docs/sdks/interview/README.md#patchatsconnectionidinterviewid) - Update a interview
* [PostAtsConnectionIDInterview](docs/sdks/interview/README.md#postatsconnectionidinterview) - Create a interview
* [PutAtsConnectionIDInterviewID](docs/sdks/interview/README.md#putatsconnectionidinterviewid) - Update a interview

### [Job](docs/sdks/job/README.md)

* [DeleteAtsConnectionIDJobID](docs/sdks/job/README.md#deleteatsconnectionidjobid) - Remove a job
* [GetAtsConnectionIDJob](docs/sdks/job/README.md#getatsconnectionidjob) - List all jobs
* [GetAtsConnectionIDJobID](docs/sdks/job/README.md#getatsconnectionidjobid) - Retrieve a job
* [PatchAtsConnectionIDJobID](docs/sdks/job/README.md#patchatsconnectionidjobid) - Update a job
* [PostAtsConnectionIDJob](docs/sdks/job/README.md#postatsconnectionidjob) - Create a job
* [PutAtsConnectionIDJobID](docs/sdks/job/README.md#putatsconnectionidjobid) - Update a job

### [Lead](docs/sdks/lead/README.md)

* [DeleteCrmConnectionIDLeadID](docs/sdks/lead/README.md#deletecrmconnectionidleadid) - Remove a lead
* [GetCrmConnectionIDLead](docs/sdks/lead/README.md#getcrmconnectionidlead) - List all leads
* [GetCrmConnectionIDLeadID](docs/sdks/lead/README.md#getcrmconnectionidleadid) - Retrieve a lead
* [PatchCrmConnectionIDLeadID](docs/sdks/lead/README.md#patchcrmconnectionidleadid) - Update a lead
* [PostCrmConnectionIDLead](docs/sdks/lead/README.md#postcrmconnectionidlead) - Create a lead
* [PutCrmConnectionIDLeadID](docs/sdks/lead/README.md#putcrmconnectionidleadid) - Update a lead

### [List](docs/sdks/list/README.md)

* [DeleteMartechConnectionIDListID](docs/sdks/list/README.md#deletemartechconnectionidlistid) - Remove a list
* [GetMartechConnectionIDList](docs/sdks/list/README.md#getmartechconnectionidlist) - List all lists
* [GetMartechConnectionIDListID](docs/sdks/list/README.md#getmartechconnectionidlistid) - Retrieve a list
* [PatchMartechConnectionIDListID](docs/sdks/list/README.md#patchmartechconnectionidlistid) - Update a list
* [PostMartechConnectionIDList](docs/sdks/list/README.md#postmartechconnectionidlist) - Create a list
* [PutMartechConnectionIDListID](docs/sdks/list/README.md#putmartechconnectionidlistid) - Update a list

### [Login](docs/sdks/login/README.md)

* [GetUnifiedIntegrationLoginWorkspaceIDIntegrationType](docs/sdks/login/README.md#getunifiedintegrationloginworkspaceidintegrationtype) - Sign in a user

### [Martech](docs/sdks/martech/README.md)

* [DeleteMartechConnectionIDListID](docs/sdks/martech/README.md#deletemartechconnectionidlistid) - Remove a list
* [DeleteMartechConnectionIDListIDMemberID](docs/sdks/martech/README.md#deletemartechconnectionidlistidmemberid) - Remove member from a list
* [GetMartechConnectionIDList](docs/sdks/martech/README.md#getmartechconnectionidlist) - List all lists
* [GetMartechConnectionIDListID](docs/sdks/martech/README.md#getmartechconnectionidlistid) - Retrieve a list
* [GetMartechConnectionIDListIDMember](docs/sdks/martech/README.md#getmartechconnectionidlistidmember) - List all members in a list
* [GetMartechConnectionIDListIDMemberID](docs/sdks/martech/README.md#getmartechconnectionidlistidmemberid) - Retrieve a member from a list
* [PatchMartechConnectionIDListID](docs/sdks/martech/README.md#patchmartechconnectionidlistid) - Update a list
* [PatchMartechConnectionIDListIDMemberID](docs/sdks/martech/README.md#patchmartechconnectionidlistidmemberid) - Update a member in a list
* [PostMartechConnectionIDList](docs/sdks/martech/README.md#postmartechconnectionidlist) - Create a list
* [PostMartechConnectionIDListIDMember](docs/sdks/martech/README.md#postmartechconnectionidlistidmember) - Create a member in a list
* [PutMartechConnectionIDListID](docs/sdks/martech/README.md#putmartechconnectionidlistid) - Update a list
* [PutMartechConnectionIDListIDMemberID](docs/sdks/martech/README.md#putmartechconnectionidlistidmemberid) - Update a member in a list

### [Member](docs/sdks/member/README.md)

* [DeleteMartechConnectionIDListIDMemberID](docs/sdks/member/README.md#deletemartechconnectionidlistidmemberid) - Remove member from a list
* [GetMartechConnectionIDListIDMember](docs/sdks/member/README.md#getmartechconnectionidlistidmember) - List all members in a list
* [GetMartechConnectionIDListIDMemberID](docs/sdks/member/README.md#getmartechconnectionidlistidmemberid) - Retrieve a member from a list
* [PatchMartechConnectionIDListIDMemberID](docs/sdks/member/README.md#patchmartechconnectionidlistidmemberid) - Update a member in a list
* [PostMartechConnectionIDListIDMember](docs/sdks/member/README.md#postmartechconnectionidlistidmember) - Create a member in a list
* [PutMartechConnectionIDListIDMemberID](docs/sdks/member/README.md#putmartechconnectionidlistidmemberid) - Update a member in a list

### [Note](docs/sdks/note/README.md)

* [DeleteTicketingConnectionIDNoteTicketIDID](docs/sdks/note/README.md#deleteticketingconnectionidnoteticketidid) - Remove a note
* [GetTicketingConnectionIDNoteTicketID](docs/sdks/note/README.md#getticketingconnectionidnoteticketid) - List all notes
* [GetTicketingConnectionIDNoteTicketIDID](docs/sdks/note/README.md#getticketingconnectionidnoteticketidid) - Retrieve a note
* [PatchTicketingConnectionIDNoteTicketIDID](docs/sdks/note/README.md#patchticketingconnectionidnoteticketidid) - Update a note
* [PostTicketingConnectionIDNoteTicketID](docs/sdks/note/README.md#postticketingconnectionidnoteticketid) - Create a note
* [PutTicketingConnectionIDNoteTicketIDID](docs/sdks/note/README.md#putticketingconnectionidnoteticketidid) - Update a note

### [Passthrough](docs/sdks/passthrough/README.md)

* [DeletePassthroughConnectionIDPath](docs/sdks/passthrough/README.md#deletepassthroughconnectionidpath) - Passthrough DELETE
* [GetPassthroughConnectionIDPath](docs/sdks/passthrough/README.md#getpassthroughconnectionidpath) - Passthrough GET
* [PatchPassthroughConnectionIDPath](docs/sdks/passthrough/README.md#patchpassthroughconnectionidpath) - Passthrough PUT
* [PostPassthroughConnectionIDPath](docs/sdks/passthrough/README.md#postpassthroughconnectionidpath) - Passthrough POST
* [PutPassthroughConnectionIDPath](docs/sdks/passthrough/README.md#putpassthroughconnectionidpath) - Passthrough PUT

### [Person](docs/sdks/person/README.md)

* [GetEnrichConnectionIDPerson](docs/sdks/person/README.md#getenrichconnectionidperson) - Retrieve enrichment information for a person

### [Pipeline](docs/sdks/pipeline/README.md)

* [DeleteCrmConnectionIDPipelineID](docs/sdks/pipeline/README.md#deletecrmconnectionidpipelineid) - Remove a pipeline
* [GetCrmConnectionIDPipeline](docs/sdks/pipeline/README.md#getcrmconnectionidpipeline) - List all pipelines
* [GetCrmConnectionIDPipelineID](docs/sdks/pipeline/README.md#getcrmconnectionidpipelineid) - Retrieve a pipeline
* [PatchCrmConnectionIDPipelineID](docs/sdks/pipeline/README.md#patchcrmconnectionidpipelineid) - Update a pipeline
* [PostCrmConnectionIDPipeline](docs/sdks/pipeline/README.md#postcrmconnectionidpipeline) - Create a pipeline
* [PutCrmConnectionIDPipelineID](docs/sdks/pipeline/README.md#putcrmconnectionidpipelineid) - Update a pipeline

### [Team](docs/sdks/team/README.md)

* [DeleteCrmConnectionIDTeamID](docs/sdks/team/README.md#deletecrmconnectionidteamid) - Remove a team
* [GetCrmConnectionIDTeam](docs/sdks/team/README.md#getcrmconnectionidteam) - List all teams
* [GetCrmConnectionIDTeamID](docs/sdks/team/README.md#getcrmconnectionidteamid) - Retrieve a team
* [PatchCrmConnectionIDTeamID](docs/sdks/team/README.md#patchcrmconnectionidteamid) - Update a team
* [PostCrmConnectionIDTeam](docs/sdks/team/README.md#postcrmconnectionidteam) - Create a team
* [PutCrmConnectionIDTeamID](docs/sdks/team/README.md#putcrmconnectionidteamid) - Update a team

### [Ticket](docs/sdks/ticket/README.md)

* [DeleteTicketingConnectionIDTicketID](docs/sdks/ticket/README.md#deleteticketingconnectionidticketid) - Remove a ticket
* [GetTicketingConnectionIDTicket](docs/sdks/ticket/README.md#getticketingconnectionidticket) - List all tickets
* [GetTicketingConnectionIDTicketID](docs/sdks/ticket/README.md#getticketingconnectionidticketid) - Retrieve a ticket
* [PatchTicketingConnectionIDTicketID](docs/sdks/ticket/README.md#patchticketingconnectionidticketid) - Update a ticket
* [PostTicketingConnectionIDTicket](docs/sdks/ticket/README.md#postticketingconnectionidticket) - Create a ticket
* [PutTicketingConnectionIDTicketID](docs/sdks/ticket/README.md#putticketingconnectionidticketid) - Update a ticket

### [Ticketing](docs/sdks/ticketing/README.md)

* [DeleteTicketingConnectionIDAgentID](docs/sdks/ticketing/README.md#deleteticketingconnectionidagentid) - Remove a agent
* [DeleteTicketingConnectionIDCustomerID](docs/sdks/ticketing/README.md#deleteticketingconnectionidcustomerid) - Remove a customer
* [DeleteTicketingConnectionIDNoteTicketIDID](docs/sdks/ticketing/README.md#deleteticketingconnectionidnoteticketidid) - Remove a note
* [DeleteTicketingConnectionIDTicketID](docs/sdks/ticketing/README.md#deleteticketingconnectionidticketid) - Remove a ticket
* [GetTicketingConnectionIDAgent](docs/sdks/ticketing/README.md#getticketingconnectionidagent) - List all agents
* [GetTicketingConnectionIDAgentID](docs/sdks/ticketing/README.md#getticketingconnectionidagentid) - Retrieve a agent
* [GetTicketingConnectionIDCustomer](docs/sdks/ticketing/README.md#getticketingconnectionidcustomer) - List all customers
* [GetTicketingConnectionIDCustomerID](docs/sdks/ticketing/README.md#getticketingconnectionidcustomerid) - Retrieve a customer
* [GetTicketingConnectionIDNoteTicketID](docs/sdks/ticketing/README.md#getticketingconnectionidnoteticketid) - List all notes
* [GetTicketingConnectionIDNoteTicketIDID](docs/sdks/ticketing/README.md#getticketingconnectionidnoteticketidid) - Retrieve a note
* [GetTicketingConnectionIDTicket](docs/sdks/ticketing/README.md#getticketingconnectionidticket) - List all tickets
* [GetTicketingConnectionIDTicketID](docs/sdks/ticketing/README.md#getticketingconnectionidticketid) - Retrieve a ticket
* [PatchTicketingConnectionIDAgentID](docs/sdks/ticketing/README.md#patchticketingconnectionidagentid) - Update a agent
* [PatchTicketingConnectionIDCustomerID](docs/sdks/ticketing/README.md#patchticketingconnectionidcustomerid) - Update a customer
* [PatchTicketingConnectionIDNoteTicketIDID](docs/sdks/ticketing/README.md#patchticketingconnectionidnoteticketidid) - Update a note
* [PatchTicketingConnectionIDTicketID](docs/sdks/ticketing/README.md#patchticketingconnectionidticketid) - Update a ticket
* [PostTicketingConnectionIDAgent](docs/sdks/ticketing/README.md#postticketingconnectionidagent) - Create a agent
* [PostTicketingConnectionIDCustomer](docs/sdks/ticketing/README.md#postticketingconnectionidcustomer) - Create a customer
* [PostTicketingConnectionIDNoteTicketID](docs/sdks/ticketing/README.md#postticketingconnectionidnoteticketid) - Create a note
* [PostTicketingConnectionIDTicket](docs/sdks/ticketing/README.md#postticketingconnectionidticket) - Create a ticket
* [PutTicketingConnectionIDAgentID](docs/sdks/ticketing/README.md#putticketingconnectionidagentid) - Update a agent
* [PutTicketingConnectionIDCustomerID](docs/sdks/ticketing/README.md#putticketingconnectionidcustomerid) - Update a customer
* [PutTicketingConnectionIDNoteTicketIDID](docs/sdks/ticketing/README.md#putticketingconnectionidnoteticketidid) - Update a note
* [PutTicketingConnectionIDTicketID](docs/sdks/ticketing/README.md#putticketingconnectionidticketid) - Update a ticket

### [Uc](docs/sdks/uc/README.md)

* [DeleteUcConnectionIDContactID](docs/sdks/uc/README.md#deleteucconnectionidcontactid) - Remove a contact
* [GetUcConnectionIDAgent](docs/sdks/uc/README.md#getucconnectionidagent) - List all agents
* [GetUcConnectionIDCall](docs/sdks/uc/README.md#getucconnectionidcall) - List all calls
* [GetUcConnectionIDContact](docs/sdks/uc/README.md#getucconnectionidcontact) - List all contacts
* [GetUcConnectionIDContactID](docs/sdks/uc/README.md#getucconnectionidcontactid) - Retrieve a contact
* [PatchUcConnectionIDContactID](docs/sdks/uc/README.md#patchucconnectionidcontactid) - Update a contact
* [PostUcConnectionIDContact](docs/sdks/uc/README.md#postucconnectionidcontact) - Create a contact
* [PutUcConnectionIDContactID](docs/sdks/uc/README.md#putucconnectionidcontactid) - Update a contact

### [Unified](docs/sdks/unified/README.md)

* [DeleteUnifiedConnectionID](docs/sdks/unified/README.md#deleteunifiedconnectionid) - Remove connection
* [DeleteUnifiedUser](docs/sdks/unified/README.md#deleteunifieduser) - Delete your user object
* [DeleteUnifiedWebhookID](docs/sdks/unified/README.md#deleteunifiedwebhookid) - Remove webhook subscription
* [GetUnifiedApicall](docs/sdks/unified/README.md#getunifiedapicall) - Returns API Calls
* [GetUnifiedApicallID](docs/sdks/unified/README.md#getunifiedapicallid) - Retrieve specific API Call by its ID
* [GetUnifiedConnection](docs/sdks/unified/README.md#getunifiedconnection) - List all connections
* [GetUnifiedConnectionID](docs/sdks/unified/README.md#getunifiedconnectionid) - Retrieve connection
* [GetUnifiedIntegration](docs/sdks/unified/README.md#getunifiedintegration) - Returns all integrations
* [GetUnifiedIntegrationAuthWorkspaceIDIntegrationType](docs/sdks/unified/README.md#getunifiedintegrationauthworkspaceidintegrationtype) - Create connection indirectly
* [GetUnifiedIntegrationIntegrationType](docs/sdks/unified/README.md#getunifiedintegrationintegrationtype) - Retrieve an integration
* [GetUnifiedIntegrationWorkspaceWorkspaceID](docs/sdks/unified/README.md#getunifiedintegrationworkspaceworkspaceid) - Returns all activated integrations in a workspace
* [GetUnifiedUser](docs/sdks/unified/README.md#getunifieduser) - Retrieve your user object
* [GetUnifiedUserToken](docs/sdks/unified/README.md#getunifiedusertoken) - Retrieve your user token
* [GetUnifiedWebhook](docs/sdks/unified/README.md#getunifiedwebhook) - Returns all registered webhooks
* [GetUnifiedWebhookID](docs/sdks/unified/README.md#getunifiedwebhookid) - Retrieve webhook by its ID
* [PatchUnifiedConnectionID](docs/sdks/unified/README.md#patchunifiedconnectionid) - Update connection
* [PatchUnifiedUser](docs/sdks/unified/README.md#patchunifieduser) - Updates your user object
* [PostUnifiedConnection](docs/sdks/unified/README.md#postunifiedconnection) - Create connection
* [PostUnifiedWebhookConnectionIDObject](docs/sdks/unified/README.md#postunifiedwebhookconnectionidobject) - Create webhook subscription
* [PutUnifiedConnectionID](docs/sdks/unified/README.md#putunifiedconnectionid) - Update connection
* [PutUnifiedUser](docs/sdks/unified/README.md#putunifieduser) - Updates your user object

### [User](docs/sdks/user/README.md)

* [DeleteCrmConnectionIDUserID](docs/sdks/user/README.md#deletecrmconnectioniduserid) - Remove a user
* [DeleteUnifiedUser](docs/sdks/user/README.md#deleteunifieduser) - Delete your user object
* [GetCrmConnectionIDUser](docs/sdks/user/README.md#getcrmconnectioniduser) - List all users
* [GetCrmConnectionIDUserID](docs/sdks/user/README.md#getcrmconnectioniduserid) - Retrieve a user
* [GetUnifiedUser](docs/sdks/user/README.md#getunifieduser) - Retrieve your user object
* [GetUnifiedUserToken](docs/sdks/user/README.md#getunifiedusertoken) - Retrieve your user token
* [PatchCrmConnectionIDUserID](docs/sdks/user/README.md#patchcrmconnectioniduserid) - Update a user
* [PatchUnifiedUser](docs/sdks/user/README.md#patchunifieduser) - Updates your user object
* [PostCrmConnectionIDUser](docs/sdks/user/README.md#postcrmconnectioniduser) - Create a user
* [PutCrmConnectionIDUserID](docs/sdks/user/README.md#putcrmconnectioniduserid) - Update a user
* [PutUnifiedUser](docs/sdks/user/README.md#putunifieduser) - Updates your user object

### [Webhook](docs/sdks/webhook/README.md)

* [DeleteUnifiedWebhookID](docs/sdks/webhook/README.md#deleteunifiedwebhookid) - Remove webhook subscription
* [GetUnifiedWebhook](docs/sdks/webhook/README.md#getunifiedwebhook) - Returns all registered webhooks
* [GetUnifiedWebhookID](docs/sdks/webhook/README.md#getunifiedwebhookid) - Retrieve webhook by its ID
* [PostUnifiedWebhookConnectionIDObject](docs/sdks/webhook/README.md#postunifiedwebhookconnectionidobject) - Create webhook subscription
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->

<!-- End Dev Containers -->



<!-- Start Pagination -->
# Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
<!-- End Pagination -->



<!-- Start Go Types -->

<!-- End Go Types -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
