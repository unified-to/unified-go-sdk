// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package unifiedgosdk

import (
	"bytes"
	"context"
	"fmt"
	"github.com/unified-to/unified-go-sdk/internal/hooks"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/sdkerrors"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/retry"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"net/http"
)

type Message struct {
	sdkConfiguration sdkConfiguration
}

func newMessage(sdkConfig sdkConfiguration) *Message {
	return &Message{
		sdkConfiguration: sdkConfig,
	}
}

// CreateMessagingMessage - Create a message
func (s *Message) CreateMessagingMessage(ctx context.Context, request operations.CreateMessagingMessageRequest, opts ...operations.Option) (*operations.CreateMessagingMessageResponse, error) {
	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	var baseURL string
	if o.ServerURL == nil {
		baseURL = utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	} else {
		baseURL = *o.ServerURL
	}
	opURL, err := utils.GenerateURL(ctx, baseURL, "/messaging/{connection_id}/message", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	hookCtx := hooks.HookContext{
		BaseURL:        baseURL,
		Context:        ctx,
		OperationID:    "createMessagingMessage",
		SecuritySource: s.sdkConfiguration.Security,
	}
	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, false, "MessagingMessage", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "POST", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	if reqContentType != "" {
		req.Header.Set("Content-Type", reqContentType)
	}

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	for k, v := range o.SetHeaders {
		req.Header.Set(k, v)
	}

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig != nil {
			retryConfig = globalRetryConfig
		}
	}

	var httpRes *http.Response
	if retryConfig != nil {
		httpRes, err = utils.Retry(ctx, utils.Retries{
			Config: retryConfig,
			StatusCodes: []string{
				"429",
				"500",
				"502",
				"503",
				"504",
			},
		}, func() (*http.Response, error) {
			if req.Body != nil {
				copyBody, err := req.GetBody()
				if err != nil {
					return nil, err
				}
				req.Body = copyBody
			}

			req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
			if err != nil {
				if retry.IsPermanentError(err) || retry.IsTemporaryError(err) {
					return nil, err
				}

				return nil, retry.Permanent(err)
			}

			httpRes, err := s.sdkConfiguration.Client.Do(req)
			if err != nil || httpRes == nil {
				if err != nil {
					err = fmt.Errorf("error sending request: %w", err)
				} else {
					err = fmt.Errorf("error sending request: no response")
				}

				_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			}
			return httpRes, err
		})

		if err != nil {
			return nil, err
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	} else {
		req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
		if err != nil {
			return nil, err
		}

		httpRes, err = s.sdkConfiguration.Client.Do(req)
		if err != nil || httpRes == nil {
			if err != nil {
				err = fmt.Errorf("error sending request: %w", err)
			} else {
				err = fmt.Errorf("error sending request: no response")
			}

			_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			return nil, err
		} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
			_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
			if err != nil {
				return nil, err
			} else if _httpRes != nil {
				httpRes = _httpRes
			}
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &operations.CreateMessagingMessageResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out shared.MessagingMessage
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.MessagingMessage = &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}

// GetMessagingMessage - Retrieve a message
func (s *Message) GetMessagingMessage(ctx context.Context, request operations.GetMessagingMessageRequest, opts ...operations.Option) (*operations.GetMessagingMessageResponse, error) {
	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	var baseURL string
	if o.ServerURL == nil {
		baseURL = utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	} else {
		baseURL = *o.ServerURL
	}
	opURL, err := utils.GenerateURL(ctx, baseURL, "/messaging/{connection_id}/message/{id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	hookCtx := hooks.HookContext{
		BaseURL:        baseURL,
		Context:        ctx,
		OperationID:    "getMessagingMessage",
		SecuritySource: s.sdkConfiguration.Security,
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	for k, v := range o.SetHeaders {
		req.Header.Set(k, v)
	}

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig != nil {
			retryConfig = globalRetryConfig
		}
	}

	var httpRes *http.Response
	if retryConfig != nil {
		httpRes, err = utils.Retry(ctx, utils.Retries{
			Config: retryConfig,
			StatusCodes: []string{
				"429",
				"500",
				"502",
				"503",
				"504",
			},
		}, func() (*http.Response, error) {
			if req.Body != nil {
				copyBody, err := req.GetBody()
				if err != nil {
					return nil, err
				}
				req.Body = copyBody
			}

			req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
			if err != nil {
				if retry.IsPermanentError(err) || retry.IsTemporaryError(err) {
					return nil, err
				}

				return nil, retry.Permanent(err)
			}

			httpRes, err := s.sdkConfiguration.Client.Do(req)
			if err != nil || httpRes == nil {
				if err != nil {
					err = fmt.Errorf("error sending request: %w", err)
				} else {
					err = fmt.Errorf("error sending request: no response")
				}

				_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			}
			return httpRes, err
		})

		if err != nil {
			return nil, err
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	} else {
		req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
		if err != nil {
			return nil, err
		}

		httpRes, err = s.sdkConfiguration.Client.Do(req)
		if err != nil || httpRes == nil {
			if err != nil {
				err = fmt.Errorf("error sending request: %w", err)
			} else {
				err = fmt.Errorf("error sending request: no response")
			}

			_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			return nil, err
		} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
			_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
			if err != nil {
				return nil, err
			} else if _httpRes != nil {
				httpRes = _httpRes
			}
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &operations.GetMessagingMessageResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out shared.MessagingMessage
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.MessagingMessage = &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}

// ListMessagingMessages - List all messages
func (s *Message) ListMessagingMessages(ctx context.Context, request operations.ListMessagingMessagesRequest, opts ...operations.Option) (*operations.ListMessagingMessagesResponse, error) {
	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	var baseURL string
	if o.ServerURL == nil {
		baseURL = utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	} else {
		baseURL = *o.ServerURL
	}
	opURL, err := utils.GenerateURL(ctx, baseURL, "/messaging/{connection_id}/message", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	hookCtx := hooks.HookContext{
		BaseURL:        baseURL,
		Context:        ctx,
		OperationID:    "listMessagingMessages",
		SecuritySource: s.sdkConfiguration.Security,
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	for k, v := range o.SetHeaders {
		req.Header.Set(k, v)
	}

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig != nil {
			retryConfig = globalRetryConfig
		}
	}

	var httpRes *http.Response
	if retryConfig != nil {
		httpRes, err = utils.Retry(ctx, utils.Retries{
			Config: retryConfig,
			StatusCodes: []string{
				"429",
				"500",
				"502",
				"503",
				"504",
			},
		}, func() (*http.Response, error) {
			if req.Body != nil {
				copyBody, err := req.GetBody()
				if err != nil {
					return nil, err
				}
				req.Body = copyBody
			}

			req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
			if err != nil {
				if retry.IsPermanentError(err) || retry.IsTemporaryError(err) {
					return nil, err
				}

				return nil, retry.Permanent(err)
			}

			httpRes, err := s.sdkConfiguration.Client.Do(req)
			if err != nil || httpRes == nil {
				if err != nil {
					err = fmt.Errorf("error sending request: %w", err)
				} else {
					err = fmt.Errorf("error sending request: no response")
				}

				_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			}
			return httpRes, err
		})

		if err != nil {
			return nil, err
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	} else {
		req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
		if err != nil {
			return nil, err
		}

		httpRes, err = s.sdkConfiguration.Client.Do(req)
		if err != nil || httpRes == nil {
			if err != nil {
				err = fmt.Errorf("error sending request: %w", err)
			} else {
				err = fmt.Errorf("error sending request: no response")
			}

			_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			return nil, err
		} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
			_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
			if err != nil {
				return nil, err
			} else if _httpRes != nil {
				httpRes = _httpRes
			}
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &operations.ListMessagingMessagesResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out []shared.MessagingMessage
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.MessagingMessages = out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}

// PatchMessagingMessage - Update a message
func (s *Message) PatchMessagingMessage(ctx context.Context, request operations.PatchMessagingMessageRequest, opts ...operations.Option) (*operations.PatchMessagingMessageResponse, error) {
	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	var baseURL string
	if o.ServerURL == nil {
		baseURL = utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	} else {
		baseURL = *o.ServerURL
	}
	opURL, err := utils.GenerateURL(ctx, baseURL, "/messaging/{connection_id}/message/{id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	hookCtx := hooks.HookContext{
		BaseURL:        baseURL,
		Context:        ctx,
		OperationID:    "patchMessagingMessage",
		SecuritySource: s.sdkConfiguration.Security,
	}
	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, false, "MessagingMessage", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "PATCH", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	if reqContentType != "" {
		req.Header.Set("Content-Type", reqContentType)
	}

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	for k, v := range o.SetHeaders {
		req.Header.Set(k, v)
	}

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig != nil {
			retryConfig = globalRetryConfig
		}
	}

	var httpRes *http.Response
	if retryConfig != nil {
		httpRes, err = utils.Retry(ctx, utils.Retries{
			Config: retryConfig,
			StatusCodes: []string{
				"429",
				"500",
				"502",
				"503",
				"504",
			},
		}, func() (*http.Response, error) {
			if req.Body != nil {
				copyBody, err := req.GetBody()
				if err != nil {
					return nil, err
				}
				req.Body = copyBody
			}

			req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
			if err != nil {
				if retry.IsPermanentError(err) || retry.IsTemporaryError(err) {
					return nil, err
				}

				return nil, retry.Permanent(err)
			}

			httpRes, err := s.sdkConfiguration.Client.Do(req)
			if err != nil || httpRes == nil {
				if err != nil {
					err = fmt.Errorf("error sending request: %w", err)
				} else {
					err = fmt.Errorf("error sending request: no response")
				}

				_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			}
			return httpRes, err
		})

		if err != nil {
			return nil, err
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	} else {
		req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
		if err != nil {
			return nil, err
		}

		httpRes, err = s.sdkConfiguration.Client.Do(req)
		if err != nil || httpRes == nil {
			if err != nil {
				err = fmt.Errorf("error sending request: %w", err)
			} else {
				err = fmt.Errorf("error sending request: no response")
			}

			_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			return nil, err
		} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
			_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
			if err != nil {
				return nil, err
			} else if _httpRes != nil {
				httpRes = _httpRes
			}
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &operations.PatchMessagingMessageResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out shared.MessagingMessage
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.MessagingMessage = &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}

// RemoveMessagingMessage - Remove a message
func (s *Message) RemoveMessagingMessage(ctx context.Context, request operations.RemoveMessagingMessageRequest, opts ...operations.Option) (*operations.RemoveMessagingMessageResponse, error) {
	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	var baseURL string
	if o.ServerURL == nil {
		baseURL = utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	} else {
		baseURL = *o.ServerURL
	}
	opURL, err := utils.GenerateURL(ctx, baseURL, "/messaging/{connection_id}/message/{id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	hookCtx := hooks.HookContext{
		BaseURL:        baseURL,
		Context:        ctx,
		OperationID:    "removeMessagingMessage",
		SecuritySource: s.sdkConfiguration.Security,
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	for k, v := range o.SetHeaders {
		req.Header.Set(k, v)
	}

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig != nil {
			retryConfig = globalRetryConfig
		}
	}

	var httpRes *http.Response
	if retryConfig != nil {
		httpRes, err = utils.Retry(ctx, utils.Retries{
			Config: retryConfig,
			StatusCodes: []string{
				"429",
				"500",
				"502",
				"503",
				"504",
			},
		}, func() (*http.Response, error) {
			if req.Body != nil {
				copyBody, err := req.GetBody()
				if err != nil {
					return nil, err
				}
				req.Body = copyBody
			}

			req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
			if err != nil {
				if retry.IsPermanentError(err) || retry.IsTemporaryError(err) {
					return nil, err
				}

				return nil, retry.Permanent(err)
			}

			httpRes, err := s.sdkConfiguration.Client.Do(req)
			if err != nil || httpRes == nil {
				if err != nil {
					err = fmt.Errorf("error sending request: %w", err)
				} else {
					err = fmt.Errorf("error sending request: no response")
				}

				_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			}
			return httpRes, err
		})

		if err != nil {
			return nil, err
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	} else {
		req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
		if err != nil {
			return nil, err
		}

		httpRes, err = s.sdkConfiguration.Client.Do(req)
		if err != nil || httpRes == nil {
			if err != nil {
				err = fmt.Errorf("error sending request: %w", err)
			} else {
				err = fmt.Errorf("error sending request: no response")
			}

			_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			return nil, err
		} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
			_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
			if err != nil {
				return nil, err
			} else if _httpRes != nil {
				httpRes = _httpRes
			}
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &operations.RemoveMessagingMessageResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	switch {
	case httpRes.StatusCode == 200:
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		res.Headers = httpRes.Header

	}

	return res, nil

}

// UpdateMessagingMessage - Update a message
func (s *Message) UpdateMessagingMessage(ctx context.Context, request operations.UpdateMessagingMessageRequest, opts ...operations.Option) (*operations.UpdateMessagingMessageResponse, error) {
	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	var baseURL string
	if o.ServerURL == nil {
		baseURL = utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	} else {
		baseURL = *o.ServerURL
	}
	opURL, err := utils.GenerateURL(ctx, baseURL, "/messaging/{connection_id}/message/{id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	hookCtx := hooks.HookContext{
		BaseURL:        baseURL,
		Context:        ctx,
		OperationID:    "updateMessagingMessage",
		SecuritySource: s.sdkConfiguration.Security,
	}
	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, false, "MessagingMessage", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	if reqContentType != "" {
		req.Header.Set("Content-Type", reqContentType)
	}

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	for k, v := range o.SetHeaders {
		req.Header.Set(k, v)
	}

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig != nil {
			retryConfig = globalRetryConfig
		}
	}

	var httpRes *http.Response
	if retryConfig != nil {
		httpRes, err = utils.Retry(ctx, utils.Retries{
			Config: retryConfig,
			StatusCodes: []string{
				"429",
				"500",
				"502",
				"503",
				"504",
			},
		}, func() (*http.Response, error) {
			if req.Body != nil {
				copyBody, err := req.GetBody()
				if err != nil {
					return nil, err
				}
				req.Body = copyBody
			}

			req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
			if err != nil {
				if retry.IsPermanentError(err) || retry.IsTemporaryError(err) {
					return nil, err
				}

				return nil, retry.Permanent(err)
			}

			httpRes, err := s.sdkConfiguration.Client.Do(req)
			if err != nil || httpRes == nil {
				if err != nil {
					err = fmt.Errorf("error sending request: %w", err)
				} else {
					err = fmt.Errorf("error sending request: no response")
				}

				_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			}
			return httpRes, err
		})

		if err != nil {
			return nil, err
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	} else {
		req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
		if err != nil {
			return nil, err
		}

		httpRes, err = s.sdkConfiguration.Client.Do(req)
		if err != nil || httpRes == nil {
			if err != nil {
				err = fmt.Errorf("error sending request: %w", err)
			} else {
				err = fmt.Errorf("error sending request: no response")
			}

			_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			return nil, err
		} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
			_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
			if err != nil {
				return nil, err
			} else if _httpRes != nil {
				httpRes = _httpRes
			}
		} else {
			httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &operations.UpdateMessagingMessageResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out shared.MessagingMessage
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.MessagingMessage = &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}
