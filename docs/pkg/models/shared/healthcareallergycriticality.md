# HealthcareAllergyCriticality

Potential severity of future reactions

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.HealthcareAllergyCriticalityLow

// Open enum: custom values can be created with a direct type cast
custom := shared.HealthcareAllergyCriticality("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `HealthcareAllergyCriticalityLow`            | LOW                                          |
| `HealthcareAllergyCriticalityHigh`           | HIGH                                         |
| `HealthcareAllergyCriticalityUnableToAssess` | UNABLE_TO_ASSESS                             |
| `HealthcareAllergyCriticalityOther`          | OTHER                                        |