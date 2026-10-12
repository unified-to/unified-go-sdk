# OnIdle

What happens when the instance is idle

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.OnIdleTerminate

// Open enum: custom values can be created with a direct type cast
custom := shared.OnIdle("custom_value")
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `OnIdleTerminate` | TERMINATE         |
| `OnIdlePause`     | PAUSE             |
| `OnIdleStop`      | STOP              |
| `OnIdleOther`     | OTHER             |