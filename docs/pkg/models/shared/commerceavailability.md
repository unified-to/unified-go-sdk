# CommerceAvailability


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `EndAt`                                    | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)   |
| `ItemID`                                   | `*string`                                  | :heavy_minus_sign:                         | (reference to CommerceItem)                |
| `LocationID`                               | `*string`                                  | :heavy_minus_sign:                         | (reference to CommerceLocation)            |
| `Raw`                                      | map[string]`any`                           | :heavy_minus_sign:                         | N/A                                        |
| `Size`                                     | `*float64`                                 | :heavy_minus_sign:                         | N/A                                        |
| `StaffUserID`                              | `*string`                                  | :heavy_minus_sign:                         | (reference to HrisEmployee)                |
| `StartAt`                                  | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)   |