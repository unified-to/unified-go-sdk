# ComputeUsageCategory

What was consumed

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.ComputeUsageCategoryGpu

// Open enum: custom values can be created with a direct type cast
custom := shared.ComputeUsageCategory("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `ComputeUsageCategoryGpu`     | GPU                           |
| `ComputeUsageCategoryCPU`     | CPU                           |
| `ComputeUsageCategoryMemory`  | MEMORY                        |
| `ComputeUsageCategoryStorage` | STORAGE                       |
| `ComputeUsageCategoryNetwork` | NETWORK                       |
| `ComputeUsageCategoryOther`   | OTHER                         |