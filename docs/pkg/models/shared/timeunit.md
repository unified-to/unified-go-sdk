# TimeUnit

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.TimeUnitUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TimeUnit("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `TimeUnitUnspecified` | UNSPECIFIED           |
| `TimeUnitLifetime`    | LIFETIME              |
| `TimeUnitMonths`      | MONTHS                |
| `TimeUnitWeeks`       | WEEKS                 |
| `TimeUnitDays`        | DAYS                  |
| `TimeUnitHours`       | HOURS                 |
| `TimeUnitMinutes`     | MINUTES               |