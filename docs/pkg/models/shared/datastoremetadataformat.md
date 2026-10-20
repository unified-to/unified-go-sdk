# DatastoreMetadataFormat

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.DatastoreMetadataFormatText

// Open enum: custom values can be created with a direct type cast
custom := shared.DatastoreMetadataFormat("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `DatastoreMetadataFormatText`           | TEXT                                    |
| `DatastoreMetadataFormatNumber`         | NUMBER                                  |
| `DatastoreMetadataFormatDate`           | DATE                                    |
| `DatastoreMetadataFormatBoolean`        | BOOLEAN                                 |
| `DatastoreMetadataFormatFile`           | FILE                                    |
| `DatastoreMetadataFormatTextarea`       | TEXTAREA                                |
| `DatastoreMetadataFormatSingleSelect`   | SINGLE_SELECT                           |
| `DatastoreMetadataFormatMultipleSelect` | MULTIPLE_SELECT                         |
| `DatastoreMetadataFormatMeasurement`    | MEASUREMENT                             |
| `DatastoreMetadataFormatPrice`          | PRICE                                   |
| `DatastoreMetadataFormatYesNo`          | YES_NO                                  |
| `DatastoreMetadataFormatCurrency`       | CURRENCY                                |
| `DatastoreMetadataFormatURL`            | URL                                     |