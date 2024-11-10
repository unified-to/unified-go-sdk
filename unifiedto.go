// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package unifiedgosdk

import (
	"context"
	"fmt"
	"github.com/unified-to/unified-go-sdk/internal/hooks"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/retry"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	// North American data region
	"https://api.unified.to",
	// European data region
	"https://api-eu.unified.to",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

// Pointer provides a helper function to return a pointer to a type
func Pointer[T any](v T) *T { return &v }

type sdkConfiguration struct {
	Client            HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *retry.Config
	Hooks             *hooks.Hooks
	Timeout           *time.Duration
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// UnifiedTo - Unified.to API: One API to Rule Them All
type UnifiedTo struct {
	Accounting        *Accounting
	Account           *Account
	Contact           *Contact
	Invoice           *Invoice
	Journal           *Journal
	Order             *Order
	Organization      *Organization
	Taxrate           *Taxrate
	Transaction       *Transaction
	Ats               *Ats
	Activity          *Activity
	Application       *Application
	Applicationstatus *Applicationstatus
	Candidate         *Candidate
	Company           *Company
	Document          *Document
	Interview         *Interview
	Job               *Job
	Scorecard         *Scorecard
	Commerce          *Commerce
	Collection        *Collection
	Inventory         *Inventory
	Item              *Item
	Location          *Location
	Crm               *Crm
	Deal              *Deal
	Event             *Event
	Lead              *Lead
	Pipeline          *Pipeline
	Enrich            *Enrich
	Person            *Person
	Genai             *Genai
	Model             *Model
	Prompt            *Prompt
	Hris              *Hris
	Employee          *Employee
	Group             *Group
	Payslip           *Payslip
	Timeoff           *Timeoff
	Kms               *Kms
	Page              *Page
	Space             *Space
	Lms               *Lms
	Class             *Class
	Course            *Course
	Instructor        *Instructor
	Student           *Student
	Martech           *Martech
	List              *List
	Member            *Member
	Messaging         *Messaging
	Channel           *Channel
	Message           *Message
	Passthrough       *Passthrough
	Payment           *Payment
	Link              *Link
	Payout            *Payout
	Refund            *Refund
	Repo              *Repo
	Branch            *Branch
	Commit            *Commit
	Pullrequest       *Pullrequest
	Repository        *Repository
	Scim              *Scim
	User              *User
	Storage           *Storage
	File              *File
	Task              *Task
	Project           *Project
	Ticketing         *Ticketing
	Customer          *Customer
	Note              *Note
	Ticket            *Ticket
	Uc                *Uc
	Call              *Call
	Unified           *Unified
	Apicall           *Apicall
	Connection        *Connection
	Integration       *Integration
	Auth              *Auth
	Login             *Login
	Issue             *Issue
	Webhook           *Webhook

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*UnifiedTo)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *UnifiedTo) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *UnifiedTo) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *UnifiedTo) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *UnifiedTo) {
		sdk.sdkConfiguration.Client = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(jwt string) SDKOption {
	return func(sdk *UnifiedTo) {
		security := shared.Security{Jwt: &jwt}
		sdk.sdkConfiguration.Security = utils.AsSecuritySource(&security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *UnifiedTo) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig retry.Config) SDKOption {
	return func(sdk *UnifiedTo) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// WithTimeout Optional request timeout applied to each operation
func WithTimeout(timeout time.Duration) SDKOption {
	return func(sdk *UnifiedTo) {
		sdk.sdkConfiguration.Timeout = &timeout
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *UnifiedTo {
	sdk := &UnifiedTo{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "1.0",
			SDKVersion:        "0.20.16",
			GenVersion:        "2.455.2",
			UserAgent:         "speakeasy-sdk/go 0.20.16 2.455.2 1.0 github.com/unified-to/unified-go-sdk",
			Hooks:             hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.Client = sdk.sdkConfiguration.Hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if serverURL != currentServerURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	sdk.Accounting = newAccounting(sdk.sdkConfiguration)

	sdk.Account = newAccount(sdk.sdkConfiguration)

	sdk.Contact = newContact(sdk.sdkConfiguration)

	sdk.Invoice = newInvoice(sdk.sdkConfiguration)

	sdk.Journal = newJournal(sdk.sdkConfiguration)

	sdk.Order = newOrder(sdk.sdkConfiguration)

	sdk.Organization = newOrganization(sdk.sdkConfiguration)

	sdk.Taxrate = newTaxrate(sdk.sdkConfiguration)

	sdk.Transaction = newTransaction(sdk.sdkConfiguration)

	sdk.Ats = newAts(sdk.sdkConfiguration)

	sdk.Activity = newActivity(sdk.sdkConfiguration)

	sdk.Application = newApplication(sdk.sdkConfiguration)

	sdk.Applicationstatus = newApplicationstatus(sdk.sdkConfiguration)

	sdk.Candidate = newCandidate(sdk.sdkConfiguration)

	sdk.Company = newCompany(sdk.sdkConfiguration)

	sdk.Document = newDocument(sdk.sdkConfiguration)

	sdk.Interview = newInterview(sdk.sdkConfiguration)

	sdk.Job = newJob(sdk.sdkConfiguration)

	sdk.Scorecard = newScorecard(sdk.sdkConfiguration)

	sdk.Commerce = newCommerce(sdk.sdkConfiguration)

	sdk.Collection = newCollection(sdk.sdkConfiguration)

	sdk.Inventory = newInventory(sdk.sdkConfiguration)

	sdk.Item = newItem(sdk.sdkConfiguration)

	sdk.Location = newLocation(sdk.sdkConfiguration)

	sdk.Crm = newCrm(sdk.sdkConfiguration)

	sdk.Deal = newDeal(sdk.sdkConfiguration)

	sdk.Event = newEvent(sdk.sdkConfiguration)

	sdk.Lead = newLead(sdk.sdkConfiguration)

	sdk.Pipeline = newPipeline(sdk.sdkConfiguration)

	sdk.Enrich = newEnrich(sdk.sdkConfiguration)

	sdk.Person = newPerson(sdk.sdkConfiguration)

	sdk.Genai = newGenai(sdk.sdkConfiguration)

	sdk.Model = newModel(sdk.sdkConfiguration)

	sdk.Prompt = newPrompt(sdk.sdkConfiguration)

	sdk.Hris = newHris(sdk.sdkConfiguration)

	sdk.Employee = newEmployee(sdk.sdkConfiguration)

	sdk.Group = newGroup(sdk.sdkConfiguration)

	sdk.Payslip = newPayslip(sdk.sdkConfiguration)

	sdk.Timeoff = newTimeoff(sdk.sdkConfiguration)

	sdk.Kms = newKms(sdk.sdkConfiguration)

	sdk.Page = newPage(sdk.sdkConfiguration)

	sdk.Space = newSpace(sdk.sdkConfiguration)

	sdk.Lms = newLms(sdk.sdkConfiguration)

	sdk.Class = newClass(sdk.sdkConfiguration)

	sdk.Course = newCourse(sdk.sdkConfiguration)

	sdk.Instructor = newInstructor(sdk.sdkConfiguration)

	sdk.Student = newStudent(sdk.sdkConfiguration)

	sdk.Martech = newMartech(sdk.sdkConfiguration)

	sdk.List = newList(sdk.sdkConfiguration)

	sdk.Member = newMember(sdk.sdkConfiguration)

	sdk.Messaging = newMessaging(sdk.sdkConfiguration)

	sdk.Channel = newChannel(sdk.sdkConfiguration)

	sdk.Message = newMessage(sdk.sdkConfiguration)

	sdk.Passthrough = newPassthrough(sdk.sdkConfiguration)

	sdk.Payment = newPayment(sdk.sdkConfiguration)

	sdk.Link = newLink(sdk.sdkConfiguration)

	sdk.Payout = newPayout(sdk.sdkConfiguration)

	sdk.Refund = newRefund(sdk.sdkConfiguration)

	sdk.Repo = newRepo(sdk.sdkConfiguration)

	sdk.Branch = newBranch(sdk.sdkConfiguration)

	sdk.Commit = newCommit(sdk.sdkConfiguration)

	sdk.Pullrequest = newPullrequest(sdk.sdkConfiguration)

	sdk.Repository = newRepository(sdk.sdkConfiguration)

	sdk.Scim = newScim(sdk.sdkConfiguration)

	sdk.User = newUser(sdk.sdkConfiguration)

	sdk.Storage = newStorage(sdk.sdkConfiguration)

	sdk.File = newFile(sdk.sdkConfiguration)

	sdk.Task = newTask(sdk.sdkConfiguration)

	sdk.Project = newProject(sdk.sdkConfiguration)

	sdk.Ticketing = newTicketing(sdk.sdkConfiguration)

	sdk.Customer = newCustomer(sdk.sdkConfiguration)

	sdk.Note = newNote(sdk.sdkConfiguration)

	sdk.Ticket = newTicket(sdk.sdkConfiguration)

	sdk.Uc = newUc(sdk.sdkConfiguration)

	sdk.Call = newCall(sdk.sdkConfiguration)

	sdk.Unified = newUnified(sdk.sdkConfiguration)

	sdk.Apicall = newApicall(sdk.sdkConfiguration)

	sdk.Connection = newConnection(sdk.sdkConfiguration)

	sdk.Integration = newIntegration(sdk.sdkConfiguration)

	sdk.Auth = newAuth(sdk.sdkConfiguration)

	sdk.Login = newLogin(sdk.sdkConfiguration)

	sdk.Issue = newIssue(sdk.sdkConfiguration)

	sdk.Webhook = newWebhook(sdk.sdkConfiguration)

	return sdk
}
