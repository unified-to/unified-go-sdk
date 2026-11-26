# Criticality

low | high | unable-to-assess

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.CriticalityLow

// Open enum: custom values can be created with a direct type cast
custom := shared.Criticality("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `CriticalityLow`            | low                         |
| `CriticalityHigh`           | high                        |
| `CriticalityUnableToAssess` | unable-to-assess            |