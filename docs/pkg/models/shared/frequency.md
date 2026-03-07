# Frequency

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.FrequencyOneTime

// Open enum: custom values can be created with a direct type cast
custom := shared.Frequency("custom_value")
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `FrequencyOneTime` | ONE_TIME           |
| `FrequencyDay`     | DAY                |
| `FrequencyQuarter` | QUARTER            |
| `FrequencyYear`    | YEAR               |
| `FrequencyHour`    | HOUR               |
| `FrequencyMonth`   | MONTH              |
| `FrequencyWeek`    | WEEK               |