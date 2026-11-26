# FhirPractitionerResourceType

Always Practitioner

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.FhirPractitionerResourceTypePatient

// Open enum: custom values can be created with a direct type cast
custom := shared.FhirPractitionerResourceType("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `FhirPractitionerResourceTypePatient`             | Patient                                           |
| `FhirPractitionerResourceTypePractitioner`        | Practitioner                                      |
| `FhirPractitionerResourceTypeEncounter`           | Encounter                                         |
| `FhirPractitionerResourceTypeObservation`         | Observation                                       |
| `FhirPractitionerResourceTypeCondition`           | Condition                                         |
| `FhirPractitionerResourceTypeAllergyIntolerance`  | AllergyIntolerance                                |
| `FhirPractitionerResourceTypeMedicationRequest`   | MedicationRequest                                 |
| `FhirPractitionerResourceTypeImmunization`        | Immunization                                      |
| `FhirPractitionerResourceTypeBundle`              | Bundle                                            |
| `FhirPractitionerResourceTypeOperationOutcome`    | OperationOutcome                                  |
| `FhirPractitionerResourceTypeCapabilityStatement` | CapabilityStatement                               |
| `FhirPractitionerResourceTypeDiagnosticReport`    | DiagnosticReport                                  |