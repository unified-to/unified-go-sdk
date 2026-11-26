# FhirPatientGender

male | female | other | unknown

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.FhirPatientGenderMale

// Open enum: custom values can be created with a direct type cast
custom := shared.FhirPatientGender("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `FhirPatientGenderMale`    | male                       |
| `FhirPatientGenderFemale`  | female                     |
| `FhirPatientGenderOther`   | other                      |
| `FhirPatientGenderUnknown` | unknown                    |