# CrmDeal

A deal represents an opportunity with companies and/or contacts


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `Amount`                                               | **float64*                                             | :heavy_minus_sign:                                     | N/A                                                    |
| `ClosedAt`                                             | [*time.Time](https://pkg.go.dev/time#Time)             | :heavy_minus_sign:                                     | N/A                                                    |
| `CreatedAt`                                            | [*time.Time](https://pkg.go.dev/time#Time)             | :heavy_minus_sign:                                     | N/A                                                    |
| `Currency`                                             | **string*                                              | :heavy_minus_sign:                                     | N/A                                                    |
| `ID`                                                   | **string*                                              | :heavy_minus_sign:                                     | N/A                                                    |
| `LostReason`                                           | **string*                                              | :heavy_minus_sign:                                     | N/A                                                    |
| `Name`                                                 | **string*                                              | :heavy_minus_sign:                                     | N/A                                                    |
| `Pipeline`                                             | **string*                                              | :heavy_minus_sign:                                     | N/A                                                    |
| `Probability`                                          | **float64*                                             | :heavy_minus_sign:                                     | N/A                                                    |
| `Raw`                                                  | map[string]*interface{}*                               | :heavy_minus_sign:                                     | The raw data returned by the integration for this deal |
| `Source`                                               | **string*                                              | :heavy_minus_sign:                                     | N/A                                                    |
| `Stage`                                                | **string*                                              | :heavy_minus_sign:                                     | N/A                                                    |
| `Tags`                                                 | []*string*                                             | :heavy_minus_sign:                                     | N/A                                                    |
| `UpdatedAt`                                            | [*time.Time](https://pkg.go.dev/time#Time)             | :heavy_minus_sign:                                     | N/A                                                    |
| `UserID`                                               | **string*                                              | :heavy_minus_sign:                                     | N/A                                                    |
| `WonReason`                                            | **string*                                              | :heavy_minus_sign:                                     | N/A                                                    |