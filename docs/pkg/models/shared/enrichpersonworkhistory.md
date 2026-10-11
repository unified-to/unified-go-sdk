# EnrichPersonWorkHistory


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `CompanyDomain`                            | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |
| `CompanyID`                                | `*string`                                  | :heavy_minus_sign:                         | (reference to EnrichCompany)               |
| `CompanyName`                              | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |
| `EndAt`                                    | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)   |
| `Location`                                 | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |
| `StartAt`                                  | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)   |
| `Title`                                    | `string`                                   | :heavy_check_mark:                         | N/A                                        |