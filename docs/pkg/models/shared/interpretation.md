# Interpretation

Interpretation of the result (e.g. high, low)

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.InterpretationNormal

// Open enum: custom values can be created with a direct type cast
custom := shared.Interpretation("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `InterpretationNormal`       | NORMAL                       |
| `InterpretationAbnormal`     | ABNORMAL                     |
| `InterpretationLow`          | LOW                          |
| `InterpretationHigh`         | HIGH                         |
| `InterpretationCriticalLow`  | CRITICAL_LOW                 |
| `InterpretationCriticalHigh` | CRITICAL_HIGH                |
| `InterpretationOther`        | OTHER                        |