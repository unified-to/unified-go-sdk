# AccountingTransactionLineItem


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `AccountID`                                                    | `*string`                                                      | :heavy_minus_sign:                                             | (reference to AccountingAccount)                               |
| `CategoryIds`                                                  | []`string`                                                     | :heavy_minus_sign:                                             | N/A                                                            |
| `Description`                                                  | `*string`                                                      | :heavy_minus_sign:                                             | N/A                                                            |
| `ID`                                                           | `*string`                                                      | :heavy_minus_sign:                                             | N/A                                                            |
| `Name`                                                         | `*string`                                                      | :heavy_minus_sign:                                             | N/A                                                            |
| `ObjectType`                                                   | `*string`                                                      | :heavy_minus_sign:                                             | N/A                                                            |
| `TaxrateID`                                                    | `*string`                                                      | :heavy_minus_sign:                                             | ref -> AccountingTaxrate applied to the line                   |
| `TotalAmount`                                                  | `*float64`                                                     | :heavy_minus_sign:                                             | will be a positive value for a debit and negative for a credit |
| `UnitAmount`                                                   | `*float64`                                                     | :heavy_minus_sign:                                             | N/A                                                            |
| `UnitQuantity`                                                 | `*float64`                                                     | :heavy_minus_sign:                                             | N/A                                                            |