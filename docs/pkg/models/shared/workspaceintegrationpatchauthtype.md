# WorkspaceIntegrationPatchAuthType

the authentication method to use for this workspace integration; when set, it overrides the default auth-method selection (e.g. force 'token' even when Unified OAUTH credentials exist)

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.WorkspaceIntegrationPatchAuthTypeOauth2

// Open enum: custom values can be created with a direct type cast
custom := shared.WorkspaceIntegrationPatchAuthType("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `WorkspaceIntegrationPatchAuthTypeOauth2` | oauth2                                    |
| `WorkspaceIntegrationPatchAuthTypeOther`  | other                                     |
| `WorkspaceIntegrationPatchAuthTypeToken`  | token                                     |