// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/unified-to/unified-go-sdk/pkg/retry"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

var ErrUnsupportedOption = errors.New("unsupported option")

const (
	SupportedOptionServerURL            = "serverURL"
	SupportedOptionRetries              = "retries"
	SupportedOptionTimeout              = "timeout"
	SupportedOptionAcceptHeaderOverride = "acceptHeaderOverride"
	SupportedOptionURLOverride          = "urlOverride"
)

type AcceptHeaderEnum string

const (
	AcceptHeaderEnumApplicationJson  AcceptHeaderEnum = "application/json"
	AcceptHeaderEnumWildcardWildcard AcceptHeaderEnum = "*/*"
	AcceptHeaderEnumTextCsv          AcceptHeaderEnum = "text/csv"
	AcceptHeaderEnumTextPlain        AcceptHeaderEnum = "text/plain"
	AcceptHeaderEnumApplicationXml   AcceptHeaderEnum = "application/xml"
)

func (e AcceptHeaderEnum) ToPointer() *AcceptHeaderEnum {
	return &e
}

type Options struct {
	ServerURL            *string
	Retries              *retry.Config
	Timeout              *time.Duration
	AcceptHeaderOverride *AcceptHeaderEnum
	URLOverride          *string
}

type Option func(*Options, ...string) error

// WithServerURL allows providing an alternative server URL.
func WithServerURL(serverURL string) Option {
	return func(opts *Options, supportedOptions ...string) error {
		if !utils.Contains(supportedOptions, SupportedOptionServerURL) {
			return ErrUnsupportedOption
		}

		opts.ServerURL = &serverURL
		return nil
	}
}

// WithTemplatedServerURL allows providing an alternative server URL with templated parameters.
func WithTemplatedServerURL(serverURL string, params map[string]string) Option {
	return func(opts *Options, supportedOptions ...string) error {
		if !utils.Contains(supportedOptions, SupportedOptionServerURL) {
			return ErrUnsupportedOption
		}

		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		opts.ServerURL = &serverURL
		return nil
	}
}

// WithRetries allows customizing the default retry configuration.
func WithRetries(config retry.Config) Option {
	return func(opts *Options, supportedOptions ...string) error {
		if !utils.Contains(supportedOptions, SupportedOptionRetries) {
			return ErrUnsupportedOption
		}

		opts.Retries = &config
		return nil
	}
}

// WithOperationTimeout allows setting the request timeout applied for an operation.
func WithOperationTimeout(timeout time.Duration) Option {
	return func(opts *Options, supportedOptions ...string) error {
		if !utils.Contains(supportedOptions, SupportedOptionRetries) {
			return ErrUnsupportedOption
		}

		opts.Timeout = &timeout
		return nil
	}
}

func WithAcceptHeaderOverride(acceptHeaderOverride AcceptHeaderEnum) Option {
	return func(opts *Options, supportedOptions ...string) error {
		if !utils.Contains(supportedOptions, SupportedOptionAcceptHeaderOverride) {
			return ErrUnsupportedOption
		}

		opts.AcceptHeaderOverride = &acceptHeaderOverride
		return nil
	}
}

// WithURLOverride allows overriding the URL.
func WithURLOverride(urlOverride string) Option {
	return func(opts *Options, supportedOptions ...string) error {
		if !utils.Contains(supportedOptions, SupportedOptionURLOverride) {
			return ErrUnsupportedOption
		}

		opts.URLOverride = &urlOverride
		return nil
	}
}
