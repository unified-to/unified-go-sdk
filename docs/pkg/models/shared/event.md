# Event

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.EventUpdated

// Open enum: custom values can be created with a direct type cast
custom := shared.Event("custom_value")
```


## Values

| Name           | Value          |
| -------------- | -------------- |
| `EventUpdated` | updated        |
| `EventCreated` | created        |
| `EventDeleted` | deleted        |