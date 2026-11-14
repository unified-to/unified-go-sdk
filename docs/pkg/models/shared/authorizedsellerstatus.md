# AuthorizedSellerStatus

ads.txt seller authorization to buy from

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AuthorizedSellerStatusAuthorizedDirectOnly

// Open enum: custom values can be created with a direct type cast
custom := shared.AuthorizedSellerStatus("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `AuthorizedSellerStatusAuthorizedDirectOnly`          | AUTHORIZED_DIRECT_ONLY                                |
| `AuthorizedSellerStatusAuthorizedDirectAndResellers`  | AUTHORIZED_DIRECT_AND_RESELLERS                       |
| `AuthorizedSellerStatusAuthorizedAndNonParticipating` | AUTHORIZED_AND_NON_PARTICIPATING                      |