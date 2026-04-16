# SigningSignatoryStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.SigningSignatoryStatusPending

// Open enum: custom values can be created with a direct type cast
custom := shared.SigningSignatoryStatus("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `SigningSignatoryStatusPending`   | PENDING                           |
| `SigningSignatoryStatusSent`      | SENT                              |
| `SigningSignatoryStatusDelivered` | DELIVERED                         |
| `SigningSignatoryStatusSigned`    | SIGNED                            |
| `SigningSignatoryStatusDeclined`  | DECLINED                          |
| `SigningSignatoryStatusError`     | ERROR                             |