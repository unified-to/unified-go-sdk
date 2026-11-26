# PrintingShippingoption

A shipping method available for a quote, with its price


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `Amount`                                                    | `*float64`                                                  | :heavy_minus_sign:                                          | Shipping price                                              |
| `Code`                                                      | `*string`                                                   | :heavy_minus_sign:                                          | Shipping method code, used as `shipping_method` on an order |
| `Currency`                                                  | `*string`                                                   | :heavy_minus_sign:                                          | ISO 4217 currency code                                      |
| `MaxDeliveryDays`                                           | `*float64`                                                  | :heavy_minus_sign:                                          | Estimated maximum number of days to delivery                |
| `MinDeliveryDays`                                           | `*float64`                                                  | :heavy_minus_sign:                                          | Estimated minimum number of days to delivery                |
| `Name`                                                      | `*string`                                                   | :heavy_minus_sign:                                          | Display name                                                |