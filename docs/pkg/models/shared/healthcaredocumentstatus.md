# HealthcareDocumentStatus

Status of the document

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.HealthcareDocumentStatusCurrent

// Open enum: custom values can be created with a direct type cast
custom := shared.HealthcareDocumentStatus("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `HealthcareDocumentStatusCurrent`        | CURRENT                                  |
| `HealthcareDocumentStatusSuperseded`     | SUPERSEDED                               |
| `HealthcareDocumentStatusEnteredInError` | ENTERED_IN_ERROR                         |
| `HealthcareDocumentStatusOther`          | OTHER                                    |