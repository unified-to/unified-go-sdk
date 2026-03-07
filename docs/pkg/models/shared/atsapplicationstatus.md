# AtsApplicationStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AtsApplicationStatusNew

// Open enum: custom values can be created with a direct type cast
custom := shared.AtsApplicationStatus("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `AtsApplicationStatusNew`             | NEW                                   |
| `AtsApplicationStatusReviewing`       | REVIEWING                             |
| `AtsApplicationStatusScreening`       | SCREENING                             |
| `AtsApplicationStatusSubmitted`       | SUBMITTED                             |
| `AtsApplicationStatusFirstInterview`  | FIRST_INTERVIEW                       |
| `AtsApplicationStatusSecondInterview` | SECOND_INTERVIEW                      |
| `AtsApplicationStatusThirdInterview`  | THIRD_INTERVIEW                       |
| `AtsApplicationStatusBackgroundCheck` | BACKGROUND_CHECK                      |
| `AtsApplicationStatusOffered`         | OFFERED                               |
| `AtsApplicationStatusAccepted`        | ACCEPTED                              |
| `AtsApplicationStatusHired`           | HIRED                                 |
| `AtsApplicationStatusRejected`        | REJECTED                              |
| `AtsApplicationStatusDeclined`        | DECLINED                              |
| `AtsApplicationStatusWithdrawn`       | WITHDRAWN                             |