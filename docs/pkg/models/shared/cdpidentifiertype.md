# CdpIdentifierType

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.CdpIdentifierTypeEmail

// Open enum: custom values can be created with a direct type cast
custom := shared.CdpIdentifierType("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `CdpIdentifierTypeEmail`       | EMAIL                          |
| `CdpIdentifierTypeUserID`      | USER_ID                        |
| `CdpIdentifierTypeAnonymousID` | ANONYMOUS_ID                   |
| `CdpIdentifierTypeDeviceID`    | DEVICE_ID                      |
| `CdpIdentifierTypePhone`       | PHONE                          |
| `CdpIdentifierTypeCrmID`       | CRM_ID                         |
| `CdpIdentifierTypeLoyaltyID`   | LOYALTY_ID                     |
| `CdpIdentifierTypeOther`       | OTHER                          |