// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package unifiedgosdk

import (
	"context"
	"fmt"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
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

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// UnifiedTo - Unified.to API: One API to Rule Them All
type UnifiedTo struct {
	Apicall     *apicall
	Application *application
	Ats         *ats
	Auth        *auth
	Call        *call
	Candidate   *candidate
	Company     *company
	Connection  *connection
	Contact     *contact
	Crm         *crm
	Customer    *customer
	Deal        *deal
	Document    *document
	Employee    *employee
	Enrich      *enrich
	Event       *event
	File        *file
	Group       *group
	Hris        *hris
	Integration *integration
	Interview   *interview
	Job         *job
	Lead        *lead
	List        *list
	Login       *login
	Martech     *martech
	Member      *member
	Note        *note
	Passthrough *passthrough
	Person      *person
	Pipeline    *pipeline
	Ticket      *ticket
	Ticketing   *ticketing
	Uc          *uc
	Unified     *unified
	Webhook     *webhook

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
		sdk.sdkConfiguration.DefaultClient = client
	}
}

func withSecurity(security interface{}) func(context.Context) (interface{}, error) {
	return func(context.Context) (interface{}, error) {
		return &security, nil
	}
}

// WithSecurity configures the SDK to use the provided security details

func WithSecurity(jwt string) SDKOption {
	return func(sdk *UnifiedTo) {
		security := shared.Security{Jwt: jwt}
		sdk.sdkConfiguration.Security = withSecurity(&security)
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *UnifiedTo) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *UnifiedTo {
	sdk := &UnifiedTo{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "1.0",
			SDKVersion:        "0.6.1",
			GenVersion:        "2.169.3",
			UserAgent:         "speakeasy-sdk/go 0.6.1 2.169.3 1.0 github.com/unified-to/unified-go-sdk",
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.Apicall = newApicall(sdk.sdkConfiguration)

	sdk.Application = newApplication(sdk.sdkConfiguration)

	sdk.Ats = newAts(sdk.sdkConfiguration)

	sdk.Auth = newAuth(sdk.sdkConfiguration)

	sdk.Call = newCall(sdk.sdkConfiguration)

	sdk.Candidate = newCandidate(sdk.sdkConfiguration)

	sdk.Company = newCompany(sdk.sdkConfiguration)

	sdk.Connection = newConnection(sdk.sdkConfiguration)

	sdk.Contact = newContact(sdk.sdkConfiguration)

	sdk.Crm = newCrm(sdk.sdkConfiguration)

	sdk.Customer = newCustomer(sdk.sdkConfiguration)

	sdk.Deal = newDeal(sdk.sdkConfiguration)

	sdk.Document = newDocument(sdk.sdkConfiguration)

	sdk.Employee = newEmployee(sdk.sdkConfiguration)

	sdk.Enrich = newEnrich(sdk.sdkConfiguration)

	sdk.Event = newEvent(sdk.sdkConfiguration)

	sdk.File = newFile(sdk.sdkConfiguration)

	sdk.Group = newGroup(sdk.sdkConfiguration)

	sdk.Hris = newHris(sdk.sdkConfiguration)

	sdk.Integration = newIntegration(sdk.sdkConfiguration)

	sdk.Interview = newInterview(sdk.sdkConfiguration)

	sdk.Job = newJob(sdk.sdkConfiguration)

	sdk.Lead = newLead(sdk.sdkConfiguration)

	sdk.List = newList(sdk.sdkConfiguration)

	sdk.Login = newLogin(sdk.sdkConfiguration)

	sdk.Martech = newMartech(sdk.sdkConfiguration)

	sdk.Member = newMember(sdk.sdkConfiguration)

	sdk.Note = newNote(sdk.sdkConfiguration)

	sdk.Passthrough = newPassthrough(sdk.sdkConfiguration)

	sdk.Person = newPerson(sdk.sdkConfiguration)

	sdk.Pipeline = newPipeline(sdk.sdkConfiguration)

	sdk.Ticket = newTicket(sdk.sdkConfiguration)

	sdk.Ticketing = newTicketing(sdk.sdkConfiguration)

	sdk.Uc = newUc(sdk.sdkConfiguration)

	sdk.Unified = newUnified(sdk.sdkConfiguration)

	sdk.Webhook = newWebhook(sdk.sdkConfiguration)

	return sdk
}
