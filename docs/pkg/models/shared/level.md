# Level

Normalized stock level

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.LevelNone

// Open enum: custom values can be created with a direct type cast
custom := shared.Level("custom_value")
```


## Values

| Name          | Value         |
| ------------- | ------------- |
| `LevelNone`   | NONE          |
| `LevelLow`    | LOW           |
| `LevelMedium` | MEDIUM        |
| `LevelHigh`   | HIGH          |
| `LevelOther`  | OTHER         |