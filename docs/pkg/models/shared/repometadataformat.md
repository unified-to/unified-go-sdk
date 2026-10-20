# RepoMetadataFormat

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.RepoMetadataFormatText

// Open enum: custom values can be created with a direct type cast
custom := shared.RepoMetadataFormat("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `RepoMetadataFormatText`           | TEXT                               |
| `RepoMetadataFormatNumber`         | NUMBER                             |
| `RepoMetadataFormatDate`           | DATE                               |
| `RepoMetadataFormatBoolean`        | BOOLEAN                            |
| `RepoMetadataFormatFile`           | FILE                               |
| `RepoMetadataFormatTextarea`       | TEXTAREA                           |
| `RepoMetadataFormatSingleSelect`   | SINGLE_SELECT                      |
| `RepoMetadataFormatMultipleSelect` | MULTIPLE_SELECT                    |
| `RepoMetadataFormatMeasurement`    | MEASUREMENT                        |
| `RepoMetadataFormatPrice`          | PRICE                              |
| `RepoMetadataFormatYesNo`          | YES_NO                             |
| `RepoMetadataFormatCurrency`       | CURRENCY                           |
| `RepoMetadataFormatURL`            | URL                                |