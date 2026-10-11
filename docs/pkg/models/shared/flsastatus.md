# FlsaStatus

US Fair Labor Standards Act classification

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.FlsaStatusExempt

// Open enum: custom values can be created with a direct type cast
custom := shared.FlsaStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `FlsaStatusExempt`            | EXEMPT                        |
| `FlsaStatusNonExempt`         | NON_EXEMPT                    |
| `FlsaStatusSalariedNonexempt` | SALARIED_NONEXEMPT            |
| `FlsaStatusOwner`             | OWNER                         |