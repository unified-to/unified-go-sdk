# Event

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.EventUserCreated

// Open enum: custom values can be created with a direct type cast
custom := shared.Event("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `EventUserCreated`             | USER_CREATED                   |
| `EventUserDeleted`             | USER_DELETED                   |
| `EventConnectionHealthy`       | CONNECTION_HEALTHY             |
| `EventConnectionUnhealthy`     | CONNECTION_UNHEALTHY           |
| `EventConnectionCreated`       | CONNECTION_CREATED             |
| `EventConnectionUpdated`       | CONNECTION_UPDATED             |
| `EventConnectionDeleted`       | CONNECTION_DELETED             |
| `EventConnectionPaused`        | CONNECTION_PAUSED              |
| `EventConnectionUnpaused`      | CONNECTION_UNPAUSED            |
| `EventIntegrationActivated`    | INTEGRATION_ACTIVATED          |
| `EventIntegrationDeactivated`  | INTEGRATION_DEACTIVATED        |
| `EventIntegrationUpdated`      | INTEGRATION_UPDATED            |
| `EventWorkspaceUpdated`        | WORKSPACE_UPDATED              |
| `EventWorkspaceOverLimit`      | WORKSPACE_OVER_LIMIT           |
| `EventWorkspace80PercentLimit` | WORKSPACE_80PERCENT_LIMIT      |
| `EventWebhookCreated`          | WEBHOOK_CREATED                |
| `EventWebhookDeleted`          | WEBHOOK_DELETED                |
| `EventWebhookUnhealthy`        | WEBHOOK_UNHEALTHY              |
| `EventWebhookPaused`           | WEBHOOK_PAUSED                 |
| `EventWebhookUnpaused`         | WEBHOOK_UNPAUSED               |