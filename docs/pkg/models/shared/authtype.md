# AuthType

the authentication method to use for this workspace integration; when set, it overrides the default auth-method selection (e.g. force 'token' even when Unified OAUTH credentials exist)

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AuthTypeOauth2

// Open enum: custom values can be created with a direct type cast
custom := shared.AuthType("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `AuthTypeOauth2` | oauth2           |
| `AuthTypeOther`  | other            |
| `AuthTypeToken`  | token            |