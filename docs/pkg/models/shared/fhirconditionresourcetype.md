# FhirConditionResourceType

Always Condition

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.FhirConditionResourceTypePatient

// Open enum: custom values can be created with a direct type cast
custom := shared.FhirConditionResourceType("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `FhirConditionResourceTypePatient`             | Patient                                        |
| `FhirConditionResourceTypePractitioner`        | Practitioner                                   |
| `FhirConditionResourceTypeEncounter`           | Encounter                                      |
| `FhirConditionResourceTypeObservation`         | Observation                                    |
| `FhirConditionResourceTypeCondition`           | Condition                                      |
| `FhirConditionResourceTypeAllergyIntolerance`  | AllergyIntolerance                             |
| `FhirConditionResourceTypeMedicationRequest`   | MedicationRequest                              |
| `FhirConditionResourceTypeImmunization`        | Immunization                                   |
| `FhirConditionResourceTypeBundle`              | Bundle                                         |
| `FhirConditionResourceTypeOperationOutcome`    | OperationOutcome                               |
| `FhirConditionResourceTypeCapabilityStatement` | CapabilityStatement                            |
| `FhirConditionResourceTypeDiagnosticReport`    | DiagnosticReport                               |