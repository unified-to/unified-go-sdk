# CdpMetadataFormat

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.CdpMetadataFormatText

// Open enum: custom values can be created with a direct type cast
custom := shared.CdpMetadataFormat("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `CdpMetadataFormatText`    | TEXT                       |
| `CdpMetadataFormatNumber`  | NUMBER                     |
| `CdpMetadataFormatDate`    | DATE                       |
| `CdpMetadataFormatBoolean` | BOOLEAN                    |
| `CdpMetadataFormatArray`   | ARRAY                      |
| `CdpMetadataFormatObject`  | OBJECT                     |