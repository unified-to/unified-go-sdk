# FhirMedicationRequestResourceType

Always MedicationRequest

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.FhirMedicationRequestResourceTypePatient

// Open enum: custom values can be created with a direct type cast
custom := shared.FhirMedicationRequestResourceType("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `FhirMedicationRequestResourceTypePatient`             | Patient                                                |
| `FhirMedicationRequestResourceTypePractitioner`        | Practitioner                                           |
| `FhirMedicationRequestResourceTypeEncounter`           | Encounter                                              |
| `FhirMedicationRequestResourceTypeObservation`         | Observation                                            |
| `FhirMedicationRequestResourceTypeCondition`           | Condition                                              |
| `FhirMedicationRequestResourceTypeAllergyIntolerance`  | AllergyIntolerance                                     |
| `FhirMedicationRequestResourceTypeMedicationRequest`   | MedicationRequest                                      |
| `FhirMedicationRequestResourceTypeImmunization`        | Immunization                                           |
| `FhirMedicationRequestResourceTypeBundle`              | Bundle                                                 |
| `FhirMedicationRequestResourceTypeOperationOutcome`    | OperationOutcome                                       |
| `FhirMedicationRequestResourceTypeCapabilityStatement` | CapabilityStatement                                    |
| `FhirMedicationRequestResourceTypeDiagnosticReport`    | DiagnosticReport                                       |