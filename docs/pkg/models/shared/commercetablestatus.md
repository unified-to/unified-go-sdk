# CommerceTableStatus

current occupancy

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.CommerceTableStatusAvailable

// Open enum: custom values can be created with a direct type cast
custom := shared.CommerceTableStatus("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `CommerceTableStatusAvailable` | AVAILABLE                      |
| `CommerceTableStatusOccupied`  | OCCUPIED                       |
| `CommerceTableStatusReserved`  | RESERVED                       |
| `CommerceTableStatusInactive`  | INACTIVE                       |
| `CommerceTableStatusOther`     | OTHER                          |