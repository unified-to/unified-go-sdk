# WebhookEvent

The event to subscribe to. "deleted" is supported on native webhooks, and on virtual webhooks where the integration supports it. Check the integration's Feature Support tab.

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.WebhookEventUpdated

// Open enum: custom values can be created with a direct type cast
custom := shared.WebhookEvent("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `WebhookEventUpdated` | updated               |
| `WebhookEventCreated` | created               |
| `WebhookEventDeleted` | deleted               |