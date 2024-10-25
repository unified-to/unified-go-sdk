# AccountingOrder


## Fields

| Field                                                                                                                  | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `AccountID`                                                                                                            | **string*                                                                                                              | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `BillingAddress`                                                                                                       | [*shared.PropertyAccountingOrderBillingAddress](../../../pkg/models/shared/propertyaccountingorderbillingaddress.md)   | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `ContactID`                                                                                                            | **string*                                                                                                              | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `CreatedAt`                                                                                                            | [*time.Time](https://pkg.go.dev/time#Time)                                                                             | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `Currency`                                                                                                             | **string*                                                                                                              | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `ID`                                                                                                                   | **string*                                                                                                              | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `Lineitems`                                                                                                            | [][shared.AccountingLineitem](../../../pkg/models/shared/accountinglineitem.md)                                        | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `Raw`                                                                                                                  | map[string]*any*                                                                                                       | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `ShippingAddress`                                                                                                      | [*shared.PropertyAccountingOrderShippingAddress](../../../pkg/models/shared/propertyaccountingordershippingaddress.md) | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `Status`                                                                                                               | [*shared.AccountingOrderStatus](../../../pkg/models/shared/accountingorderstatus.md)                                   | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `TotalAmount`                                                                                                          | **float64*                                                                                                             | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `Type`                                                                                                                 | [*shared.AccountingOrderType](../../../pkg/models/shared/accountingordertype.md)                                       | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `UpdatedAt`                                                                                                            | [*time.Time](https://pkg.go.dev/time#Time)                                                                             | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |