# PlmChangeType

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PlmChangeTypeRequest

// Open enum: custom values can be created with a direct type cast
custom := shared.PlmChangeType("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `PlmChangeTypeRequest`   | REQUEST                  |
| `PlmChangeTypeOrder`     | ORDER                    |
| `PlmChangeTypeNotice`    | NOTICE                   |
| `PlmChangeTypeDeviation` | DEVIATION                |
| `PlmChangeTypeOther`     | OTHER                    |