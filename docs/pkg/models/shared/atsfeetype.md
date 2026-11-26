# AtsFeeType

The fee arrangement (e.g. contingency or retained search)

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AtsFeeTypeContingency

// Open enum: custom values can be created with a direct type cast
custom := shared.AtsFeeType("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `AtsFeeTypeContingency` | CONTINGENCY             |
| `AtsFeeTypeRetained`    | RETAINED                |
| `AtsFeeTypeOther`       | OTHER                   |