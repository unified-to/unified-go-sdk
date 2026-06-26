# SecretsManagerType

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.SecretsManagerTypeAws

// Open enum: custom values can be created with a direct type cast
custom := shared.SecretsManagerType("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `SecretsManagerTypeAws`       | aws                           |
| `SecretsManagerTypeAzure`     | azure                         |
| `SecretsManagerTypeGcp`       | gcp                           |
| `SecretsManagerTypeHashicorp` | hashicorp                     |
| `SecretsManagerTypeComposio`  | composio                      |