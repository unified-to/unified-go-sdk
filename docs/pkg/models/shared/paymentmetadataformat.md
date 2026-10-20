# PaymentMetadataFormat

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PaymentMetadataFormatText

// Open enum: custom values can be created with a direct type cast
custom := shared.PaymentMetadataFormat("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `PaymentMetadataFormatText`           | TEXT                                  |
| `PaymentMetadataFormatNumber`         | NUMBER                                |
| `PaymentMetadataFormatDate`           | DATE                                  |
| `PaymentMetadataFormatBoolean`        | BOOLEAN                               |
| `PaymentMetadataFormatFile`           | FILE                                  |
| `PaymentMetadataFormatTextarea`       | TEXTAREA                              |
| `PaymentMetadataFormatSingleSelect`   | SINGLE_SELECT                         |
| `PaymentMetadataFormatMultipleSelect` | MULTIPLE_SELECT                       |
| `PaymentMetadataFormatMeasurement`    | MEASUREMENT                           |
| `PaymentMetadataFormatPrice`          | PRICE                                 |
| `PaymentMetadataFormatYesNo`          | YES_NO                                |
| `PaymentMetadataFormatCurrency`       | CURRENCY                              |
| `PaymentMetadataFormatURL`            | URL                                   |