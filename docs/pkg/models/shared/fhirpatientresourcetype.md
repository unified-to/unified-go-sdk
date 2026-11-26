# FhirPatientResourceType

Always Patient

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.FhirPatientResourceTypePatient

// Open enum: custom values can be created with a direct type cast
custom := shared.FhirPatientResourceType("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `FhirPatientResourceTypePatient`             | Patient                                      |
| `FhirPatientResourceTypePractitioner`        | Practitioner                                 |
| `FhirPatientResourceTypeEncounter`           | Encounter                                    |
| `FhirPatientResourceTypeObservation`         | Observation                                  |
| `FhirPatientResourceTypeCondition`           | Condition                                    |
| `FhirPatientResourceTypeAllergyIntolerance`  | AllergyIntolerance                           |
| `FhirPatientResourceTypeMedicationRequest`   | MedicationRequest                            |
| `FhirPatientResourceTypeImmunization`        | Immunization                                 |
| `FhirPatientResourceTypeBundle`              | Bundle                                       |
| `FhirPatientResourceTypeOperationOutcome`    | OperationOutcome                             |
| `FhirPatientResourceTypeCapabilityStatement` | CapabilityStatement                          |
| `FhirPatientResourceTypeDiagnosticReport`    | DiagnosticReport                             |