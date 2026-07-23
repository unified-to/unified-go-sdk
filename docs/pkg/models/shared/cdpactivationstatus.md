# CdpActivationStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.CdpActivationStatusActive

// Open enum: custom values can be created with a direct type cast
custom := shared.CdpActivationStatus("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `CdpActivationStatusActive`  | ACTIVE                       |
| `CdpActivationStatusPaused`  | PAUSED                       |
| `CdpActivationStatusError`   | ERROR                        |
| `CdpActivationStatusPending` | PENDING                      |