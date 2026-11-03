# Source

Whether a part is made in-house or bought

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.SourceMake

// Open enum: custom values can be created with a direct type cast
custom := shared.Source("custom_value")
```


## Values

| Name          | Value         |
| ------------- | ------------- |
| `SourceMake`  | MAKE          |
| `SourceBuy`   | BUY           |
| `SourceOther` | OTHER         |