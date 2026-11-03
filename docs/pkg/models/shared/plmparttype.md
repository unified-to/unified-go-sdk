# PlmPartType

Normalized part type

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PlmPartTypeComponent

// Open enum: custom values can be created with a direct type cast
custom := shared.PlmPartType("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `PlmPartTypeComponent`    | COMPONENT                 |
| `PlmPartTypeAssembly`     | ASSEMBLY                  |
| `PlmPartTypeRawMaterial`  | RAW_MATERIAL              |
| `PlmPartTypeSoftware`     | SOFTWARE                  |
| `PlmPartTypeDocumentOnly` | DOCUMENT_ONLY             |
| `PlmPartTypeOther`        | OTHER                     |