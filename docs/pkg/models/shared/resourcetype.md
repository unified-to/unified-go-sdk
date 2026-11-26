# ResourceType

Always AllergyIntolerance

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.ResourceTypePatient

// Open enum: custom values can be created with a direct type cast
custom := shared.ResourceType("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `ResourceTypePatient`             | Patient                           |
| `ResourceTypePractitioner`        | Practitioner                      |
| `ResourceTypeEncounter`           | Encounter                         |
| `ResourceTypeObservation`         | Observation                       |
| `ResourceTypeCondition`           | Condition                         |
| `ResourceTypeAllergyIntolerance`  | AllergyIntolerance                |
| `ResourceTypeMedicationRequest`   | MedicationRequest                 |
| `ResourceTypeImmunization`        | Immunization                      |
| `ResourceTypeBundle`              | Bundle                            |
| `ResourceTypeOperationOutcome`    | OperationOutcome                  |
| `ResourceTypeCapabilityStatement` | CapabilityStatement               |
| `ResourceTypeDiagnosticReport`    | DiagnosticReport                  |