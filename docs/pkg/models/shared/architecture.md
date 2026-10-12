# Architecture

CPU architecture

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.ArchitectureX8664

// Open enum: custom values can be created with a direct type cast
custom := shared.Architecture("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `ArchitectureX8664` | X86_64              |
| `ArchitectureArm64` | ARM64               |
| `ArchitectureOther` | OTHER               |