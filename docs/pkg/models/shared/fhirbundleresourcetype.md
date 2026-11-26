# FhirBundleResourceType

Always Bundle

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.FhirBundleResourceTypePatient

// Open enum: custom values can be created with a direct type cast
custom := shared.FhirBundleResourceType("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `FhirBundleResourceTypePatient`             | Patient                                     |
| `FhirBundleResourceTypePractitioner`        | Practitioner                                |
| `FhirBundleResourceTypeEncounter`           | Encounter                                   |
| `FhirBundleResourceTypeObservation`         | Observation                                 |
| `FhirBundleResourceTypeCondition`           | Condition                                   |
| `FhirBundleResourceTypeAllergyIntolerance`  | AllergyIntolerance                          |
| `FhirBundleResourceTypeMedicationRequest`   | MedicationRequest                           |
| `FhirBundleResourceTypeImmunization`        | Immunization                                |
| `FhirBundleResourceTypeBundle`              | Bundle                                      |
| `FhirBundleResourceTypeOperationOutcome`    | OperationOutcome                            |
| `FhirBundleResourceTypeCapabilityStatement` | CapabilityStatement                         |
| `FhirBundleResourceTypeDiagnosticReport`    | DiagnosticReport                            |