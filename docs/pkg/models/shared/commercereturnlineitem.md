# CommerceReturnLineitem


## Fields

| Field                       | Type                        | Required                    | Description                 |
| --------------------------- | --------------------------- | --------------------------- | --------------------------- |
| `ID`                        | `*string`                   | :heavy_minus_sign:          | N/A                         |
| `ItemID`                    | `*string`                   | :heavy_minus_sign:          | (reference to CommerceItem) |
| `Name`                      | `*string`                   | :heavy_minus_sign:          | N/A                         |
| `OrderLineID`               | `*string`                   | :heavy_minus_sign:          | N/A                         |
| `Quantity`                  | `*float64`                  | :heavy_minus_sign:          | N/A                         |
| `Raw`                       | map[string]`any`            | :heavy_minus_sign:          | N/A                         |
| `Reason`                    | `*string`                   | :heavy_minus_sign:          | N/A                         |
| `RefundAmount`              | `*float64`                  | :heavy_minus_sign:          | N/A                         |
| `Sku`                       | `*string`                   | :heavy_minus_sign:          | N/A                         |
| `Status`                    | `*string`                   | :heavy_minus_sign:          | N/A                         |