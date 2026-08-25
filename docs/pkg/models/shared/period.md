# Period

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PeriodDay

// Open enum: custom values can be created with a direct type cast
custom := shared.Period("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `PeriodDay`      | DAY              |
| `PeriodWeek`     | WEEK             |
| `PeriodMonth`    | MONTH            |
| `PeriodLifetime` | LIFETIME         |
| `PeriodTotal`    | TOTAL            |
| `PeriodOther`    | OTHER            |