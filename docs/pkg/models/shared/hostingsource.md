# HostingSource

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.HostingSourceUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.HostingSource("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `HostingSourceUnspecified`     | UNSPECIFIED                    |
| `HostingSourceCm`              | CM                             |
| `HostingSourceThirdParty`      | THIRD_PARTY                    |
| `HostingSourceHosted`          | HOSTED                         |
| `HostingSourceRichMedia`       | RICH_MEDIA                     |
| `HostingSourcePublisherHosted` | PUBLISHER_HOSTED               |