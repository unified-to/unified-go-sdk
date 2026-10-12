# HealthcareObservationComponentInterpretation

Interpretation of the component value

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.HealthcareObservationComponentInterpretationNormal

// Open enum: custom values can be created with a direct type cast
custom := shared.HealthcareObservationComponentInterpretation("custom_value")
```


## Values

| Name                                                       | Value                                                      |
| ---------------------------------------------------------- | ---------------------------------------------------------- |
| `HealthcareObservationComponentInterpretationNormal`       | NORMAL                                                     |
| `HealthcareObservationComponentInterpretationAbnormal`     | ABNORMAL                                                   |
| `HealthcareObservationComponentInterpretationLow`          | LOW                                                        |
| `HealthcareObservationComponentInterpretationHigh`         | HIGH                                                       |
| `HealthcareObservationComponentInterpretationCriticalLow`  | CRITICAL_LOW                                               |
| `HealthcareObservationComponentInterpretationCriticalHigh` | CRITICAL_HIGH                                              |
| `HealthcareObservationComponentInterpretationOther`        | OTHER                                                      |