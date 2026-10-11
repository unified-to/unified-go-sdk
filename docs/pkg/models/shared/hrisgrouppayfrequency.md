# HrisGroupPayFrequency

pay cycle when type is PAY_GROUP

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.HrisGroupPayFrequencyOneTime

// Open enum: custom values can be created with a direct type cast
custom := shared.HrisGroupPayFrequency("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `HrisGroupPayFrequencyOneTime`     | ONE_TIME                           |
| `HrisGroupPayFrequencyDay`         | DAY                                |
| `HrisGroupPayFrequencyQuarter`     | QUARTER                            |
| `HrisGroupPayFrequencyYear`        | YEAR                               |
| `HrisGroupPayFrequencyHour`        | HOUR                               |
| `HrisGroupPayFrequencyMonth`       | MONTH                              |
| `HrisGroupPayFrequencyWeek`        | WEEK                               |
| `HrisGroupPayFrequencyBiweekly`    | BIWEEKLY                           |
| `HrisGroupPayFrequencySemimonthly` | SEMIMONTHLY                        |