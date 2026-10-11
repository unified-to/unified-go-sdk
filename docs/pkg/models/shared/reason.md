# Reason

Why the refund was issued

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.ReasonDuplicate

// Open enum: custom values can be created with a direct type cast
custom := shared.Reason("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `ReasonDuplicate`           | DUPLICATE                   |
| `ReasonFraudulent`          | FRAUDULENT                  |
| `ReasonRequestedByCustomer` | REQUESTED_BY_CUSTOMER       |
| `ReasonOther`               | OTHER                       |