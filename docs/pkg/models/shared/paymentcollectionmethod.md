# PaymentCollectionMethod

When set to charging_automatically, an automated attempt will occur to pay this invoice using the default payment source attached to the contactcustomer. When set to send_invoice, an will email will be sent with this invoice to the contact/customer with payment instructions.

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