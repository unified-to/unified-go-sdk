# Status

Account status, such as ACTIVE or ARCHIVED

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.StatusActive

// Open enum: custom values can be created with a direct type cast
custom := shared.Status("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `StatusActive`   | ACTIVE           |
| `StatusArchived` | ARCHIVED         |