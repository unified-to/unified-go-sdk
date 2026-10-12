# ComputeServiceStatus

Normalized status; set to SUSPENDED or LIVE on update to suspend or resume

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.ComputeServiceStatusCreating

// Open enum: custom values can be created with a direct type cast
custom := shared.ComputeServiceStatus("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `ComputeServiceStatusCreating`  | CREATING                        |
| `ComputeServiceStatusDeploying` | DEPLOYING                       |
| `ComputeServiceStatusLive`      | LIVE                            |
| `ComputeServiceStatusSuspended` | SUSPENDED                       |
| `ComputeServiceStatusFailed`    | FAILED                          |
| `ComputeServiceStatusOther`     | OTHER                           |