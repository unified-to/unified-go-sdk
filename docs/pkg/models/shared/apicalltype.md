# APICallType

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.APICallTypeLogin

// Open enum: custom values can be created with a direct type cast
custom := shared.APICallType("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `APICallTypeLogin`   | login                |
| `APICallTypeWebhook` | webhook              |
| `APICallTypeInbound` | inbound              |
| `APICallTypeMcp`     | mcp                  |