# PropertyHrisPayslipDeductionFrequency

Frequency for this deduction (should always be set, matches IHrisBenefit.frequency)

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PropertyHrisPayslipDeductionFrequencyOneTime

// Open enum: custom values can be created with a direct type cast
custom := shared.PropertyHrisPayslipDeductionFrequency("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `PropertyHrisPayslipDeductionFrequencyOneTime`     | ONE_TIME                                           |
| `PropertyHrisPayslipDeductionFrequencyDay`         | DAY                                                |
| `PropertyHrisPayslipDeductionFrequencyQuarter`     | QUARTER                                            |
| `PropertyHrisPayslipDeductionFrequencyYear`        | YEAR                                               |
| `PropertyHrisPayslipDeductionFrequencyHour`        | HOUR                                               |
| `PropertyHrisPayslipDeductionFrequencyMonth`       | MONTH                                              |
| `PropertyHrisPayslipDeductionFrequencyWeek`        | WEEK                                               |
| `PropertyHrisPayslipDeductionFrequencyBiweekly`    | BIWEEKLY                                           |
| `PropertyHrisPayslipDeductionFrequencySemimonthly` | SEMIMONTHLY                                        |