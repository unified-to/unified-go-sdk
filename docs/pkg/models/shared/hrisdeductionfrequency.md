# HrisDeductionFrequency

Frequency for this deduction (should always be set, matches IHrisBenefit.frequency)

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.HrisDeductionFrequencyOneTime

// Open enum: custom values can be created with a direct type cast
custom := shared.HrisDeductionFrequency("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `HrisDeductionFrequencyOneTime`     | ONE_TIME                            |
| `HrisDeductionFrequencyDay`         | DAY                                 |
| `HrisDeductionFrequencyQuarter`     | QUARTER                             |
| `HrisDeductionFrequencyYear`        | YEAR                                |
| `HrisDeductionFrequencyHour`        | HOUR                                |
| `HrisDeductionFrequencyMonth`       | MONTH                               |
| `HrisDeductionFrequencyWeek`        | WEEK                                |
| `HrisDeductionFrequencyBiweekly`    | BIWEEKLY                            |
| `HrisDeductionFrequencySemimonthly` | SEMIMONTHLY                         |