// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetUnifiedIntegrationLoginRequest struct {
	Env *string `queryParam:"style=form,explode=true,name=env"`
	// The URL where you want the user to be redirect to after an unsuccessful authentication. An "error" variable will be appended.
	FailureRedirect *string `queryParam:"style=form,explode=true,name=failure_redirect"`
	// Type of the supported integration
	IntegrationType string `pathParam:"style=simple,explode=false,name=integration_type"`
	Redirect        *bool  `queryParam:"style=form,explode=true,name=redirect"`
	// Extra state to send back to your success URL
	State *string `queryParam:"style=form,explode=true,name=state"`
	// The URL where you want the user to be redirect to after a successful authentication/sign-in.  A "jwt" parameter will be appended to the URL which will contain a name and email of the user who just signed-in.
	SuccessRedirect *string `queryParam:"style=form,explode=true,name=success_redirect"`
	// The ID of the workspace
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
}

func (o *GetUnifiedIntegrationLoginRequest) GetEnv() *string {
	if o == nil {
		return nil
	}
	return o.Env
}

func (o *GetUnifiedIntegrationLoginRequest) GetFailureRedirect() *string {
	if o == nil {
		return nil
	}
	return o.FailureRedirect
}

func (o *GetUnifiedIntegrationLoginRequest) GetIntegrationType() string {
	if o == nil {
		return ""
	}
	return o.IntegrationType
}

func (o *GetUnifiedIntegrationLoginRequest) GetRedirect() *bool {
	if o == nil {
		return nil
	}
	return o.Redirect
}

func (o *GetUnifiedIntegrationLoginRequest) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *GetUnifiedIntegrationLoginRequest) GetSuccessRedirect() *string {
	if o == nil {
		return nil
	}
	return o.SuccessRedirect
}

func (o *GetUnifiedIntegrationLoginRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

type GetUnifiedIntegrationLoginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	Res *string
}

func (o *GetUnifiedIntegrationLoginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetUnifiedIntegrationLoginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetUnifiedIntegrationLoginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetUnifiedIntegrationLoginResponse) GetRes() *string {
	if o == nil {
		return nil
	}
	return o.Res
}
