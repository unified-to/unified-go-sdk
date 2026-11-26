# FhirObservationResourceType

Always Observation

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.FhirObservationResourceTypePatient

// Open enum: custom values can be created with a direct type cast
custom := shared.FhirObservationResourceType("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `FhirObservationResourceTypePatient`             | Patient                                          |
| `FhirObservationResourceTypePractitioner`        | Practitioner                                     |
| `FhirObservationResourceTypeEncounter`           | Encounter                                        |
| `FhirObservationResourceTypeObservation`         | Observation                                      |
| `FhirObservationResourceTypeCondition`           | Condition                                        |
| `FhirObservationResourceTypeAllergyIntolerance`  | AllergyIntolerance                               |
| `FhirObservationResourceTypeMedicationRequest`   | MedicationRequest                                |
| `FhirObservationResourceTypeImmunization`        | Immunization                                     |
| `FhirObservationResourceTypeBundle`              | Bundle                                           |
| `FhirObservationResourceTypeOperationOutcome`    | OperationOutcome                                 |
| `FhirObservationResourceTypeCapabilityStatement` | CapabilityStatement                              |
| `FhirObservationResourceTypeDiagnosticReport`    | DiagnosticReport                                 |