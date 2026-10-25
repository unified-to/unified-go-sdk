# HealthcareCoverageStatus

Status of the coverage

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.HealthcareCoverageStatusActive

// Open enum: custom values can be created with a direct type cast
custom := shared.HealthcareCoverageStatus("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `HealthcareCoverageStatusActive`         | ACTIVE                                   |
| `HealthcareCoverageStatusCancelled`      | CANCELLED                                |
| `HealthcareCoverageStatusDraft`          | DRAFT                                    |
| `HealthcareCoverageStatusEnteredInError` | ENTERED_IN_ERROR                         |
| `HealthcareCoverageStatusOther`          | OTHER                                    |