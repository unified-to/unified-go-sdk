# AssessmentParameterType

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AssessmentParameterTypeText

// Open enum: custom values can be created with a direct type cast
custom := shared.AssessmentParameterType("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `AssessmentParameterTypeText`           | TEXT                                    |
| `AssessmentParameterTypeNumber`         | NUMBER                                  |
| `AssessmentParameterTypeMultipleChoice` | MULTIPLE_CHOICE                         |
| `AssessmentParameterTypeMultipleSelect` | MULTIPLE_SELECT                         |
| `AssessmentParameterTypeDate`           | DATE                                    |
| `AssessmentParameterTypeFile`           | FILE                                    |