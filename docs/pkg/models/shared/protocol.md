# Protocol

The protocol served on this port

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.ProtocolHTTP

// Open enum: custom values can be created with a direct type cast
custom := shared.Protocol("custom_value")
```


## Values

| Name            | Value           |
| --------------- | --------------- |
| `ProtocolHTTP`  | HTTP            |
| `ProtocolHTTPS` | HTTPS           |
| `ProtocolTCP`   | TCP             |
| `ProtocolUDP`   | UDP             |
| `ProtocolOther` | OTHER           |