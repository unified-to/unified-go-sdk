# FormFieldType

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.FormFieldTypeText

// Open enum: custom values can be created with a direct type cast
custom := shared.FormFieldType("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `FormFieldTypeText`           | TEXT                          |
| `FormFieldTypeTextarea`       | TEXTAREA                      |
| `FormFieldTypeNumber`         | NUMBER                        |
| `FormFieldTypeEmail`          | EMAIL                         |
| `FormFieldTypeURL`            | URL                           |
| `FormFieldTypeDate`           | DATE                          |
| `FormFieldTypeTime`           | TIME                          |
| `FormFieldTypeDatetime`       | DATETIME                      |
| `FormFieldTypePhone`          | PHONE                         |
| `FormFieldTypeBoolean`        | BOOLEAN                       |
| `FormFieldTypeSingleSelect`   | SINGLE_SELECT                 |
| `FormFieldTypeMultipleSelect` | MULTIPLE_SELECT               |
| `FormFieldTypeFileUpload`     | FILE_UPLOAD                   |
| `FormFieldTypeRating`         | RATING                        |
| `FormFieldTypeScale`          | SCALE                         |
| `FormFieldTypeMatrix`         | MATRIX                        |
| `FormFieldTypeSectionHeader`  | SECTION_HEADER                |
| `FormFieldTypeOther`          | OTHER                         |