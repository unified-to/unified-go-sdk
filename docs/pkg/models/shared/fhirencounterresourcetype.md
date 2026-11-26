# FhirEncounterResourceType

Always Encounter

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.FhirEncounterResourceTypePatient

// Open enum: custom values can be created with a direct type cast
custom := shared.FhirEncounterResourceType("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `FhirEncounterResourceTypePatient`             | Patient                                        |
| `FhirEncounterResourceTypePractitioner`        | Practitioner                                   |
| `FhirEncounterResourceTypeEncounter`           | Encounter                                      |
| `FhirEncounterResourceTypeObservation`         | Observation                                    |
| `FhirEncounterResourceTypeCondition`           | Condition                                      |
| `FhirEncounterResourceTypeAllergyIntolerance`  | AllergyIntolerance                             |
| `FhirEncounterResourceTypeMedicationRequest`   | MedicationRequest                              |
| `FhirEncounterResourceTypeImmunization`        | Immunization                                   |
| `FhirEncounterResourceTypeBundle`              | Bundle                                         |
| `FhirEncounterResourceTypeOperationOutcome`    | OperationOutcome                               |
| `FhirEncounterResourceTypeCapabilityStatement` | CapabilityStatement                            |
| `FhirEncounterResourceTypeDiagnosticReport`    | DiagnosticReport                               |