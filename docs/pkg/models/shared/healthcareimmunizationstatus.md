# HealthcareImmunizationStatus

Status of the immunization

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.HealthcareImmunizationStatusCompleted

// Open enum: custom values can be created with a direct type cast
custom := shared.HealthcareImmunizationStatus("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `HealthcareImmunizationStatusCompleted`      | COMPLETED                                    |
| `HealthcareImmunizationStatusNotDone`        | NOT_DONE                                     |
| `HealthcareImmunizationStatusEnteredInError` | ENTERED_IN_ERROR                             |
| `HealthcareImmunizationStatusOther`          | OTHER                                        |