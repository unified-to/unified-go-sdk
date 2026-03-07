# FromWebhook

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.FromWebhookSupportedRequired

// Open enum: custom values can be created with a direct type cast
custom := shared.FromWebhook("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `FromWebhookSupportedRequired` | supported-required             |
| `FromWebhookSupported`         | supported                      |
| `FromWebhookNotSupported`      | not-supported                  |