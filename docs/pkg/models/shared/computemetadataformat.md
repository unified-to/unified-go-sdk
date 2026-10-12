# ComputeMetadataFormat

The data type of the value

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.ComputeMetadataFormatText

// Open enum: custom values can be created with a direct type cast
custom := shared.ComputeMetadataFormat("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `ComputeMetadataFormatText`    | TEXT                           |
| `ComputeMetadataFormatNumber`  | NUMBER                         |
| `ComputeMetadataFormatDate`    | DATE                           |
| `ComputeMetadataFormatBoolean` | BOOLEAN                        |
| `ComputeMetadataFormatArray`   | ARRAY                          |
| `ComputeMetadataFormatObject`  | OBJECT                         |