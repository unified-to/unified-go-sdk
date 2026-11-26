# DmsRepairorderPayType

Primary pay type of the repair order

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.DmsRepairorderPayTypeCustomer

// Open enum: custom values can be created with a direct type cast
custom := shared.DmsRepairorderPayType("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `DmsRepairorderPayTypeCustomer`        | CUSTOMER                               |
| `DmsRepairorderPayTypeWarranty`        | WARRANTY                               |
| `DmsRepairorderPayTypeInternal`        | INTERNAL                               |
| `DmsRepairorderPayTypeServiceContract` | SERVICE_CONTRACT                       |
| `DmsRepairorderPayTypeOther`           | OTHER                                  |