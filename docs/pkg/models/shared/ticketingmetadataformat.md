# TicketingMetadataFormat

The data type of the custom field, which determines how `value` is formatted.

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.TicketingMetadataFormatText

// Open enum: custom values can be created with a direct type cast
custom := shared.TicketingMetadataFormat("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `TicketingMetadataFormatText`           | TEXT                                    |
| `TicketingMetadataFormatNumber`         | NUMBER                                  |
| `TicketingMetadataFormatDate`           | DATE                                    |
| `TicketingMetadataFormatBoolean`        | BOOLEAN                                 |
| `TicketingMetadataFormatFile`           | FILE                                    |
| `TicketingMetadataFormatTextarea`       | TEXTAREA                                |
| `TicketingMetadataFormatSingleSelect`   | SINGLE_SELECT                           |
| `TicketingMetadataFormatMultipleSelect` | MULTIPLE_SELECT                         |
| `TicketingMetadataFormatMeasurement`    | MEASUREMENT                             |
| `TicketingMetadataFormatPrice`          | PRICE                                   |
| `TicketingMetadataFormatYesNo`          | YES_NO                                  |
| `TicketingMetadataFormatCurrency`       | CURRENCY                                |
| `TicketingMetadataFormatURL`            | URL                                     |