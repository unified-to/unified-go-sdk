# HrisEarningFrequency

ONE_TIME for a single addition to the next unprocessed pay run, otherwise recurring

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.HrisEarningFrequencyOneTime

// Open enum: custom values can be created with a direct type cast
custom := shared.HrisEarningFrequency("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `HrisEarningFrequencyOneTime`     | ONE_TIME                          |
| `HrisEarningFrequencyDay`         | DAY                               |
| `HrisEarningFrequencyQuarter`     | QUARTER                           |
| `HrisEarningFrequencyYear`        | YEAR                              |
| `HrisEarningFrequencyHour`        | HOUR                              |
| `HrisEarningFrequencyMonth`       | MONTH                             |
| `HrisEarningFrequencyWeek`        | WEEK                              |
| `HrisEarningFrequencyBiweekly`    | BIWEEKLY                          |
| `HrisEarningFrequencySemimonthly` | SEMIMONTHLY                       |