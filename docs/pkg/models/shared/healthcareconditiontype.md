# HealthcareConditionType

Category of the condition

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.HealthcareConditionTypeProblem

// Open enum: custom values can be created with a direct type cast
custom := shared.HealthcareConditionType("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `HealthcareConditionTypeProblem`       | PROBLEM                                |
| `HealthcareConditionTypeDiagnosis`     | DIAGNOSIS                              |
| `HealthcareConditionTypeHealthConcern` | HEALTH_CONCERN                         |
| `HealthcareConditionTypeOther`         | OTHER                                  |