# DmsRepairorderJobPayType

Who pays for this job

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.DmsRepairorderJobPayTypeCustomer

// Open enum: custom values can be created with a direct type cast
custom := shared.DmsRepairorderJobPayType("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `DmsRepairorderJobPayTypeCustomer`        | CUSTOMER                                  |
| `DmsRepairorderJobPayTypeWarranty`        | WARRANTY                                  |
| `DmsRepairorderJobPayTypeInternal`        | INTERNAL                                  |
| `DmsRepairorderJobPayTypeServiceContract` | SERVICE_CONTRACT                          |
| `DmsRepairorderJobPayTypeOther`           | OTHER                                     |