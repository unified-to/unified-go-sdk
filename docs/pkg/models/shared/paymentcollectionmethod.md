# PaymentCollectionMethod

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PaymentCollectionMethodSendInvoice

// Open enum: custom values can be created with a direct type cast
custom := shared.PaymentCollectionMethod("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `PaymentCollectionMethodSendInvoice`         | send_invoice                                 |
| `PaymentCollectionMethodChargeAutomatically` | charge_automatically                         |