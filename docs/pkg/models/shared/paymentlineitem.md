# PaymentLineitem


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `AccountID`                                | `*string`                                  | :heavy_minus_sign:                         | (reference to AccountingAccount)           |
| `CreatedAt`                                | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)   |
| `DiscountAmount`                           | `*float64`                                 | :heavy_minus_sign:                         | N/A                                        |
| `ID`                                       | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |
| `ItemDescription`                          | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |
| `ItemID`                                   | `*string`                                  | :heavy_minus_sign:                         | (reference to CommerceItem)                |
| `ItemName`                                 | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |
| `ItemSku`                                  | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |
| `Notes`                                    | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |
| `RefundAmount`                             | `*float64`                                 | :heavy_minus_sign:                         | new field                                  |
| `RefundedAt`                               | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)   |
| `TaxAmount`                                | `*float64`                                 | :heavy_minus_sign:                         | N/A                                        |
| `TaxrateID`                                | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |
| `TotalAmount`                              | `*float64`                                 | :heavy_minus_sign:                         | N/A                                        |
| `UnitAmount`                               | `*float64`                                 | :heavy_minus_sign:                         | N/A                                        |
| `UnitQuantity`                             | `*float64`                                 | :heavy_minus_sign:                         | N/A                                        |
| `UpdatedAt`                                | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)   |