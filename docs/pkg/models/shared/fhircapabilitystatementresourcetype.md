# FhirCapabilityStatementResourceType

Always CapabilityStatement

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.FhirCapabilityStatementResourceTypePatient

// Open enum: custom values can be created with a direct type cast
custom := shared.FhirCapabilityStatementResourceType("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `FhirCapabilityStatementResourceTypePatient`             | Patient                                                  |
| `FhirCapabilityStatementResourceTypePractitioner`        | Practitioner                                             |
| `FhirCapabilityStatementResourceTypeEncounter`           | Encounter                                                |
| `FhirCapabilityStatementResourceTypeObservation`         | Observation                                              |
| `FhirCapabilityStatementResourceTypeCondition`           | Condition                                                |
| `FhirCapabilityStatementResourceTypeAllergyIntolerance`  | AllergyIntolerance                                       |
| `FhirCapabilityStatementResourceTypeMedicationRequest`   | MedicationRequest                                        |
| `FhirCapabilityStatementResourceTypeImmunization`        | Immunization                                             |
| `FhirCapabilityStatementResourceTypeBundle`              | Bundle                                                   |
| `FhirCapabilityStatementResourceTypeOperationOutcome`    | OperationOutcome                                         |
| `FhirCapabilityStatementResourceTypeCapabilityStatement` | CapabilityStatement                                      |
| `FhirCapabilityStatementResourceTypeDiagnosticReport`    | DiagnosticReport                                         |