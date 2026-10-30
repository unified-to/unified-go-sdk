# SenderType

whether the guest, host or system sent the message

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.SenderTypeGuest

// Open enum: custom values can be created with a direct type cast
custom := shared.SenderType("custom_value")
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `SenderTypeGuest`  | GUEST              |
| `SenderTypeHost`   | HOST               |
| `SenderTypeSystem` | SYSTEM             |