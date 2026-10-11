# HrisEarningStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.HrisEarningStatusPending

// Open enum: custom values can be created with a direct type cast
custom := shared.HrisEarningStatus("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `HrisEarningStatusPending`   | PENDING                      |
| `HrisEarningStatusActive`    | ACTIVE                       |
| `HrisEarningStatusProcessed` | PROCESSED                    |
| `HrisEarningStatusEnded`     | ENDED                        |
| `HrisEarningStatusCancelled` | CANCELLED                    |
| `HrisEarningStatusFailed`    | FAILED                       |