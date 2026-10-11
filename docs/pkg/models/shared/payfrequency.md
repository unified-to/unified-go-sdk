# PayFrequency

how often the employee is paid (eg. BIWEEKLY)

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PayFrequencyOneTime

// Open enum: custom values can be created with a direct type cast
custom := shared.PayFrequency("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `PayFrequencyOneTime`     | ONE_TIME                  |
| `PayFrequencyDay`         | DAY                       |
| `PayFrequencyQuarter`     | QUARTER                   |
| `PayFrequencyYear`        | YEAR                      |
| `PayFrequencyHour`        | HOUR                      |
| `PayFrequencyMonth`       | MONTH                     |
| `PayFrequencyWeek`        | WEEK                      |
| `PayFrequencyBiweekly`    | BIWEEKLY                  |
| `PayFrequencySemimonthly` | SEMIMONTHLY               |