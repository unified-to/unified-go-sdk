# LegalMetadataFormat

the data type of the value

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.LegalMetadataFormatText

// Open enum: custom values can be created with a direct type cast
custom := shared.LegalMetadataFormat("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `LegalMetadataFormatText`           | TEXT                                |
| `LegalMetadataFormatNumber`         | NUMBER                              |
| `LegalMetadataFormatDate`           | DATE                                |
| `LegalMetadataFormatBoolean`        | BOOLEAN                             |
| `LegalMetadataFormatFile`           | FILE                                |
| `LegalMetadataFormatTextarea`       | TEXTAREA                            |
| `LegalMetadataFormatSingleSelect`   | SINGLE_SELECT                       |
| `LegalMetadataFormatMultipleSelect` | MULTIPLE_SELECT                     |
| `LegalMetadataFormatMeasurement`    | MEASUREMENT                         |
| `LegalMetadataFormatPrice`          | PRICE                               |
| `LegalMetadataFormatYesNo`          | YES_NO                              |
| `LegalMetadataFormatCurrency`       | CURRENCY                            |
| `LegalMetadataFormatURL`            | URL                                 |