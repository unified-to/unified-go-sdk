# DatastoreFieldType

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.DatastoreFieldTypeText

// Open enum: custom values can be created with a direct type cast
custom := shared.DatastoreFieldType("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `DatastoreFieldTypeText`           | TEXT                               |
| `DatastoreFieldTypeNumber`         | NUMBER                             |
| `DatastoreFieldTypeDate`           | DATE                               |
| `DatastoreFieldTypeBoolean`        | BOOLEAN                            |
| `DatastoreFieldTypeFile`           | FILE                               |
| `DatastoreFieldTypeTextarea`       | TEXTAREA                           |
| `DatastoreFieldTypeSingleSelect`   | SINGLE_SELECT                      |
| `DatastoreFieldTypeMultipleSelect` | MULTIPLE_SELECT                    |
| `DatastoreFieldTypeCurrency`       | CURRENCY                           |
| `DatastoreFieldTypeURL`            | URL                                |
| `DatastoreFieldTypeEmail`          | EMAIL                              |
| `DatastoreFieldTypePhone`          | PHONE                              |
| `DatastoreFieldTypeLinkedRecord`   | LINKED_RECORD                      |
| `DatastoreFieldTypeRelation`       | RELATION                           |