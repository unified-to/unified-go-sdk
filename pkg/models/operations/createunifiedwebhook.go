// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateUnifiedWebhookEvents string

const (
	CreateUnifiedWebhookEventsUpdated CreateUnifiedWebhookEvents = "updated"
	CreateUnifiedWebhookEventsCreated CreateUnifiedWebhookEvents = "created"
)

func (e CreateUnifiedWebhookEvents) ToPointer() *CreateUnifiedWebhookEvents {
	return &e
}

func (e *CreateUnifiedWebhookEvents) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "updated":
		fallthrough
	case "created":
		*e = CreateUnifiedWebhookEvents(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateUnifiedWebhookEvents: %v", v)
	}
}

type CreateUnifiedWebhookRequest struct {
	Webhook *shared.Webhook `request:"mediaType=application/json"`
	// ID of the connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	// Which events to subscribe to.
	Events []CreateUnifiedWebhookEvents `queryParam:"style=form,explode=true,name=events"`
	// The object to subscribe to
	Object string `pathParam:"style=simple,explode=false,name=object"`
}

func (o *CreateUnifiedWebhookRequest) GetWebhook() *shared.Webhook {
	if o == nil {
		return nil
	}
	return o.Webhook
}

func (o *CreateUnifiedWebhookRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateUnifiedWebhookRequest) GetEvents() []CreateUnifiedWebhookEvents {
	if o == nil {
		return nil
	}
	return o.Events
}

func (o *CreateUnifiedWebhookRequest) GetObject() string {
	if o == nil {
		return ""
	}
	return o.Object
}

type CreateUnifiedWebhookResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful
	Webhook *shared.Webhook
}

func (o *CreateUnifiedWebhookResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateUnifiedWebhookResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateUnifiedWebhookResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateUnifiedWebhookResponse) GetWebhook() *shared.Webhook {
	if o == nil {
		return nil
	}
	return o.Webhook
}
