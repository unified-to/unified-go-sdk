# HealthcareEncounterType

Class of the encounter (e.g. ambulatory, inpatient)

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.HealthcareEncounterTypeAmbulatory

// Open enum: custom values can be created with a direct type cast
custom := shared.HealthcareEncounterType("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `HealthcareEncounterTypeAmbulatory`  | AMBULATORY                           |
| `HealthcareEncounterTypeEmergency`   | EMERGENCY                            |
| `HealthcareEncounterTypeInpatient`   | INPATIENT                            |
| `HealthcareEncounterTypeObservation` | OBSERVATION                          |
| `HealthcareEncounterTypeHomeHealth`  | HOME_HEALTH                          |
| `HealthcareEncounterTypeVirtual`     | VIRTUAL                              |
| `HealthcareEncounterTypeOther`       | OTHER                                |