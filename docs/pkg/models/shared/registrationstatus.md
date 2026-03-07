# RegistrationStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.RegistrationStatusPending

// Open enum: custom values can be created with a direct type cast
custom := shared.RegistrationStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `RegistrationStatusPending`   | PENDING                       |
| `RegistrationStatusApproved`  | APPROVED                      |
| `RegistrationStatusRejected`  | REJECTED                      |
| `RegistrationStatusCancelled` | CANCELLED                     |