# EncodingFormat

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.EncodingFormatFloat

// Open enum: custom values can be created with a direct type cast
custom := shared.EncodingFormat("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `EncodingFormatFloat`   | FLOAT                   |
| `EncodingFormatUint8`   | UINT8                   |
| `EncodingFormatInt8`    | INT8                    |
| `EncodingFormatBinary`  | BINARY                  |
| `EncodingFormatUbinary` | UBINARY                 |
| `EncodingFormatBase64`  | BASE64                  |