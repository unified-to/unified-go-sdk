# WebhookType

"native" registers with the integration, which sends events as they happen. "virtual" checks for changes on "interval". Optional on create; if omitted, Unified.to picks a supported type for the object. Requesting a type the integration doesn't support returns 400. Set on create only: changes sent on update are ignored without an error. Delete and recreate the webhook to change type.

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.WebhookTypeVirtual

// Open enum: custom values can be created with a direct type cast
custom := shared.WebhookType("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `WebhookTypeVirtual` | virtual              |
| `WebhookTypeNative`  | native               |