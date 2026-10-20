# StorageMetadataFormat

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.StorageMetadataFormatText

// Open enum: custom values can be created with a direct type cast
custom := shared.StorageMetadataFormat("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `StorageMetadataFormatText`           | TEXT                                  |
| `StorageMetadataFormatNumber`         | NUMBER                                |
| `StorageMetadataFormatDate`           | DATE                                  |
| `StorageMetadataFormatBoolean`        | BOOLEAN                               |
| `StorageMetadataFormatFile`           | FILE                                  |
| `StorageMetadataFormatTextarea`       | TEXTAREA                              |
| `StorageMetadataFormatSingleSelect`   | SINGLE_SELECT                         |
| `StorageMetadataFormatMultipleSelect` | MULTIPLE_SELECT                       |
| `StorageMetadataFormatMeasurement`    | MEASUREMENT                           |
| `StorageMetadataFormatPrice`          | PRICE                                 |
| `StorageMetadataFormatYesNo`          | YES_NO                                |
| `StorageMetadataFormatCurrency`       | CURRENCY                              |
| `StorageMetadataFormatURL`            | URL                                   |