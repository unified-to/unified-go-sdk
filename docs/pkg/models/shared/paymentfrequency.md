# PaymentFrequency

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PaymentFrequencyMonthly

// Open enum: custom values can be created with a direct type cast
custom := shared.PaymentFrequency("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `PaymentFrequencyMonthly`     | MONTHLY                       |
| `PaymentFrequencySemiMonthly` | SEMI_MONTHLY                  |
| `PaymentFrequencyBiweekly`    | BIWEEKLY                      |
| `PaymentFrequencyWeekly`      | WEEKLY                        |
| `PaymentFrequencyYearly`      | YEARLY                        |
| `PaymentFrequencyOther`       | OTHER                         |