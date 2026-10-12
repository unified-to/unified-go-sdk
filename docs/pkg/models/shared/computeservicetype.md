# ComputeServiceType

The kind of service

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.ComputeServiceTypeWeb

// Open enum: custom values can be created with a direct type cast
custom := shared.ComputeServiceType("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `ComputeServiceTypeWeb`     | WEB                         |
| `ComputeServiceTypeWorker`  | WORKER                      |
| `ComputeServiceTypeCron`    | CRON                        |
| `ComputeServiceTypeStatic`  | STATIC                      |
| `ComputeServiceTypePrivate` | PRIVATE                     |
| `ComputeServiceTypeOther`   | OTHER                       |