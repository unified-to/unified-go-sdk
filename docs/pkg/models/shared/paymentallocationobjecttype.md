# PaymentAllocationObjectType

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PaymentAllocationObjectTypeInvoice

// Open enum: custom values can be created with a direct type cast
custom := shared.PaymentAllocationObjectType("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `PaymentAllocationObjectTypeInvoice`       | INVOICE                                    |
| `PaymentAllocationObjectTypeBill`          | BILL                                       |
| `PaymentAllocationObjectTypeCreditmemo`    | CREDITMEMO                                 |
| `PaymentAllocationObjectTypeVendorcredit`  | VENDORCREDIT                               |
| `PaymentAllocationObjectTypeSalesorder`    | SALESORDER                                 |
| `PaymentAllocationObjectTypePurchaseorder` | PURCHASEORDER                              |