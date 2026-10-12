# OnExpire

What happens when the instance expires

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.OnExpireTerminate

// Open enum: custom values can be created with a direct type cast
custom := shared.OnExpire("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `OnExpireTerminate` | TERMINATE           |
| `OnExpirePause`     | PAUSE               |
| `OnExpireStop`      | STOP                |
| `OnExpireOther`     | OTHER               |