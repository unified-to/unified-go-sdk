# FhirDiagnosticReportResourceType

Always DiagnosticReport

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.FhirDiagnosticReportResourceTypePatient

// Open enum: custom values can be created with a direct type cast
custom := shared.FhirDiagnosticReportResourceType("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `FhirDiagnosticReportResourceTypePatient`             | Patient                                               |
| `FhirDiagnosticReportResourceTypePractitioner`        | Practitioner                                          |
| `FhirDiagnosticReportResourceTypeEncounter`           | Encounter                                             |
| `FhirDiagnosticReportResourceTypeObservation`         | Observation                                           |
| `FhirDiagnosticReportResourceTypeCondition`           | Condition                                             |
| `FhirDiagnosticReportResourceTypeAllergyIntolerance`  | AllergyIntolerance                                    |
| `FhirDiagnosticReportResourceTypeMedicationRequest`   | MedicationRequest                                     |
| `FhirDiagnosticReportResourceTypeImmunization`        | Immunization                                          |
| `FhirDiagnosticReportResourceTypeBundle`              | Bundle                                                |
| `FhirDiagnosticReportResourceTypeOperationOutcome`    | OperationOutcome                                      |
| `FhirDiagnosticReportResourceTypeCapabilityStatement` | CapabilityStatement                                   |
| `FhirDiagnosticReportResourceTypeDiagnosticReport`    | DiagnosticReport                                      |