# PropertyAvailabilityStatus

whether the period is available, booked or blocked

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PropertyAvailabilityStatusAvailable

// Open enum: custom values can be created with a direct type cast
custom := shared.PropertyAvailabilityStatus("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `PropertyAvailabilityStatusAvailable`   | AVAILABLE                               |
| `PropertyAvailabilityStatusUnavailable` | UNAVAILABLE                             |
| `PropertyAvailabilityStatusBlocked`     | BLOCKED                                 |