# HrisPayrollStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.HrisPayrollStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := shared.HrisPayrollStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `HrisPayrollStatusDraft`      | DRAFT                         |
| `HrisPayrollStatusPending`    | PENDING                       |
| `HrisPayrollStatusApproved`   | APPROVED                      |
| `HrisPayrollStatusProcessing` | PROCESSING                    |
| `HrisPayrollStatusPaid`       | PAID                          |
| `HrisPayrollStatusFailed`     | FAILED                        |
| `HrisPayrollStatusCancelled`  | CANCELLED                     |