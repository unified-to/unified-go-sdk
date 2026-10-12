# HealthcareObservationType

Category of the observation

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.HealthcareObservationTypeLaboratory

// Open enum: custom values can be created with a direct type cast
custom := shared.HealthcareObservationType("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `HealthcareObservationTypeLaboratory`    | LABORATORY                               |
| `HealthcareObservationTypeVitalSigns`    | VITAL_SIGNS                              |
| `HealthcareObservationTypeSocialHistory` | SOCIAL_HISTORY                           |
| `HealthcareObservationTypeImaging`       | IMAGING                                  |
| `HealthcareObservationTypeSurvey`        | SURVEY                                   |
| `HealthcareObservationTypeExam`          | EXAM                                     |
| `HealthcareObservationTypeProcedure`     | PROCEDURE                                |
| `HealthcareObservationTypeTherapy`       | THERAPY                                  |
| `HealthcareObservationTypeActivity`      | ACTIVITY                                 |
| `HealthcareObservationTypeOther`         | OTHER                                    |