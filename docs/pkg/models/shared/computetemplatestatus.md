# ComputeTemplateStatus

Normalized build status

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.ComputeTemplateStatusPending

// Open enum: custom values can be created with a direct type cast
custom := shared.ComputeTemplateStatus("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `ComputeTemplateStatusPending`  | PENDING                         |
| `ComputeTemplateStatusBuilding` | BUILDING                        |
| `ComputeTemplateStatusReady`    | READY                           |
| `ComputeTemplateStatusFailed`   | FAILED                          |
| `ComputeTemplateStatusOther`    | OTHER                           |