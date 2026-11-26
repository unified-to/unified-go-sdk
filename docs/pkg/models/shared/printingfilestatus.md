# PrintingFileStatus

Processing status of the file

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PrintingFileStatusPending

// Open enum: custom values can be created with a direct type cast
custom := shared.PrintingFileStatus("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `PrintingFileStatusPending` | PENDING                     |
| `PrintingFileStatusReady`   | READY                       |
| `PrintingFileStatusFailed`  | FAILED                      |
| `PrintingFileStatusOther`   | OTHER                       |