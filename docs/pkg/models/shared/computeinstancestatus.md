# ComputeInstanceStatus

Normalized lifecycle status; set on update to start, stop, pause, resume, restart or terminate

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.ComputeInstanceStatusProvisioning

// Open enum: custom values can be created with a direct type cast
custom := shared.ComputeInstanceStatus("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `ComputeInstanceStatusProvisioning` | PROVISIONING                        |
| `ComputeInstanceStatusStarting`     | STARTING                            |
| `ComputeInstanceStatusRunning`      | RUNNING                             |
| `ComputeInstanceStatusStopping`     | STOPPING                            |
| `ComputeInstanceStatusStopped`      | STOPPED                             |
| `ComputeInstanceStatusPausing`      | PAUSING                             |
| `ComputeInstanceStatusPaused`       | PAUSED                              |
| `ComputeInstanceStatusResuming`     | RESUMING                            |
| `ComputeInstanceStatusRestarting`   | RESTARTING                          |
| `ComputeInstanceStatusTerminating`  | TERMINATING                         |
| `ComputeInstanceStatusTerminated`   | TERMINATED                          |
| `ComputeInstanceStatusPreempted`    | PREEMPTED                           |
| `ComputeInstanceStatusError`        | ERROR                               |
| `ComputeInstanceStatusOther`        | OTHER                               |