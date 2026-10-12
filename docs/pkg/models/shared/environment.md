# Environment

Target environment

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.EnvironmentProduction

// Open enum: custom values can be created with a direct type cast
custom := shared.Environment("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `EnvironmentProduction`  | PRODUCTION               |
| `EnvironmentPreview`     | PREVIEW                  |
| `EnvironmentDevelopment` | DEVELOPMENT              |
| `EnvironmentOther`       | OTHER                    |