# CdpConsentStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.CdpConsentStatusGranted

// Open enum: custom values can be created with a direct type cast
custom := shared.CdpConsentStatus("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `CdpConsentStatusGranted` | GRANTED                   |
| `CdpConsentStatusDenied`  | DENIED                    |
| `CdpConsentStatusPending` | PENDING                   |