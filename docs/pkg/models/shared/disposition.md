# Disposition

What to do with the nonconforming material

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.DispositionUseAsIs

// Open enum: custom values can be created with a direct type cast
custom := shared.Disposition("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `DispositionUseAsIs`        | USE_AS_IS                   |
| `DispositionRework`         | REWORK                      |
| `DispositionRepair`         | REPAIR                      |
| `DispositionScrap`          | SCRAP                       |
| `DispositionReturnToVendor` | RETURN_TO_VENDOR            |
| `DispositionOther`          | OTHER                       |