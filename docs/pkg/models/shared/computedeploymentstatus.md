# ComputeDeploymentStatus

Normalized deployment status; set to CANCELED on update to cancel

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.ComputeDeploymentStatusQueued

// Open enum: custom values can be created with a direct type cast
custom := shared.ComputeDeploymentStatus("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `ComputeDeploymentStatusQueued`     | QUEUED                              |
| `ComputeDeploymentStatusBuilding`   | BUILDING                            |
| `ComputeDeploymentStatusDeploying`  | DEPLOYING                           |
| `ComputeDeploymentStatusLive`       | LIVE                                |
| `ComputeDeploymentStatusFailed`     | FAILED                              |
| `ComputeDeploymentStatusCanceled`   | CANCELED                            |
| `ComputeDeploymentStatusSuperseded` | SUPERSEDED                          |
| `ComputeDeploymentStatusOther`      | OTHER                               |