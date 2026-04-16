# SigningSignatoryRole

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.SigningSignatoryRoleSigner

// Open enum: custom values can be created with a direct type cast
custom := shared.SigningSignatoryRole("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `SigningSignatoryRoleSigner`         | SIGNER                               |
| `SigningSignatoryRoleCc`             | CC                                   |
| `SigningSignatoryRoleApprover`       | APPROVER                             |
| `SigningSignatoryRoleInPersonSigner` | IN_PERSON_SIGNER                     |
| `SigningSignatoryRoleViewer`         | VIEWER                               |