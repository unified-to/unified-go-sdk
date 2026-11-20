# EmploymentStatusReason

why the employee is not actively working (eg. LEAVE, TERMINATED, RETIRED)

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.EmploymentStatusReasonLeave

// Open enum: custom values can be created with a direct type cast
custom := shared.EmploymentStatusReason("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `EmploymentStatusReasonLeave`      | LEAVE                              |
| `EmploymentStatusReasonTerminated` | TERMINATED                         |
| `EmploymentStatusReasonRetired`    | RETIRED                            |
| `EmploymentStatusReasonDeceased`   | DECEASED                           |
| `EmploymentStatusReasonDisabled`   | DISABLED                           |
| `EmploymentStatusReasonOther`      | OTHER                              |