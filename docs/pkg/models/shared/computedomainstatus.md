# ComputeDomainStatus

DNS/ownership verification status

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.ComputeDomainStatusPending

// Open enum: custom values can be created with a direct type cast
custom := shared.ComputeDomainStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `ComputeDomainStatusPending`  | PENDING                       |
| `ComputeDomainStatusVerified` | VERIFIED                      |
| `ComputeDomainStatusFailed`   | FAILED                        |
| `ComputeDomainStatusOther`    | OTHER                         |