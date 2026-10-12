# ComputeAvailability


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `Level`                                                    | [*shared.Level](../../../pkg/models/shared/level.md)       | :heavy_minus_sign:                                         | Normalized stock level                                     |
| `Quantity`                                                 | `*float64`                                                 | :heavy_minus_sign:                                         | Number of units available, when the vendor reports a count |
| `RegionID`                                                 | `*string`                                                  | :heavy_minus_sign:                                         | The region this availability applies to                    |