# LegalMatterStatus

the matter status

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.LegalMatterStatusPending

// Open enum: custom values can be created with a direct type cast
custom := shared.LegalMatterStatus("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `LegalMatterStatusPending` | PENDING                    |
| `LegalMatterStatusOpen`    | OPEN                       |
| `LegalMatterStatusClosed`  | CLOSED                     |
| `LegalMatterStatusOther`   | OTHER                      |