# FilingStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.FilingStatusSingle

// Open enum: custom values can be created with a direct type cast
custom := shared.FilingStatus("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `FilingStatusSingle`             | SINGLE                           |
| `FilingStatusMarriedJointly`     | MARRIED_JOINTLY                  |
| `FilingStatusMarriedSeparately`  | MARRIED_SEPARATELY               |
| `FilingStatusHeadOfHousehold`    | HEAD_OF_HOUSEHOLD                |
| `FilingStatusQualifyingSurvivor` | QUALIFYING_SURVIVOR              |
| `FilingStatusOther`              | OTHER                            |