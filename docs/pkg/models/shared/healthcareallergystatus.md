# HealthcareAllergyStatus

Clinical status of the allergy

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.HealthcareAllergyStatusActive

// Open enum: custom values can be created with a direct type cast
custom := shared.HealthcareAllergyStatus("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `HealthcareAllergyStatusActive`   | ACTIVE                            |
| `HealthcareAllergyStatusInactive` | INACTIVE                          |
| `HealthcareAllergyStatusResolved` | RESOLVED                          |
| `HealthcareAllergyStatusOther`    | OTHER                             |