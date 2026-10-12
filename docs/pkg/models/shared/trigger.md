# Trigger

What started the deployment

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.TriggerGit

// Open enum: custom values can be created with a direct type cast
custom := shared.Trigger("custom_value")
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `TriggerGit`      | GIT               |
| `TriggerManual`   | MANUAL            |
| `TriggerAPI`      | API               |
| `TriggerRollback` | ROLLBACK          |
| `TriggerOther`    | OTHER             |