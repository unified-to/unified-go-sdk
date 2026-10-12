# ComputeKeyType

The kind of credential

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.ComputeKeyTypeSSH

// Open enum: custom values can be created with a direct type cast
custom := shared.ComputeKeyType("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `ComputeKeyTypeSSH`      | SSH                      |
| `ComputeKeyTypeRegistry` | REGISTRY                 |
| `ComputeKeyTypeSecret`   | SECRET                   |
| `ComputeKeyTypeOther`    | OTHER                    |