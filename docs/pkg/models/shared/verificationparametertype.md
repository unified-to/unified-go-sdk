# VerificationParameterType

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.VerificationParameterTypeText

// Open enum: custom values can be created with a direct type cast
custom := shared.VerificationParameterType("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `VerificationParameterTypeText`           | TEXT                                      |
| `VerificationParameterTypeNumber`         | NUMBER                                    |
| `VerificationParameterTypeMultipleChoice` | MULTIPLE_CHOICE                           |
| `VerificationParameterTypeMultipleSelect` | MULTIPLE_SELECT                           |
| `VerificationParameterTypeDate`           | DATE                                      |
| `VerificationParameterTypeFile`           | FILE                                      |