# VerificationRequestResponseStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.VerificationRequestResponseStatusCompleted

// Open enum: custom values can be created with a direct type cast
custom := shared.VerificationRequestResponseStatus("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `VerificationRequestResponseStatusCompleted` | COMPLETED                                    |
| `VerificationRequestResponseStatusFailed`    | FAILED                                       |
| `VerificationRequestResponseStatusPassed`    | PASSED                                       |
| `VerificationRequestResponseStatusPending`   | PENDING                                      |