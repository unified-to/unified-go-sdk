# VirtualWebhookStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.VirtualWebhookStatusSupportedRequired

// Open enum: custom values can be created with a direct type cast
custom := shared.VirtualWebhookStatus("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `VirtualWebhookStatusSupportedRequired` | supported-required                      |
| `VirtualWebhookStatusSupported`         | supported                               |
| `VirtualWebhookStatusNotSupported`      | not-supported                           |