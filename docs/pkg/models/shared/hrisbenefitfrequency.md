# HrisBenefitFrequency

Frequency that costs accrue for this benefit. For insurance/benefits, usually "MONTHLY". For compensation, matches pay cycle (WEEKLY, BIWEEKLY, etc).

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.HrisBenefitFrequencyOneTime

// Open enum: custom values can be created with a direct type cast
custom := shared.HrisBenefitFrequency("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `HrisBenefitFrequencyOneTime`     | ONE_TIME                          |
| `HrisBenefitFrequencyDay`         | DAY                               |
| `HrisBenefitFrequencyQuarter`     | QUARTER                           |
| `HrisBenefitFrequencyYear`        | YEAR                              |
| `HrisBenefitFrequencyHour`        | HOUR                              |
| `HrisBenefitFrequencyMonth`       | MONTH                             |
| `HrisBenefitFrequencyWeek`        | WEEK                              |
| `HrisBenefitFrequencyBiweekly`    | BIWEEKLY                          |
| `HrisBenefitFrequencySemimonthly` | SEMIMONTHLY                       |