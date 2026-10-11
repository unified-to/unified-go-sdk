# TaskComment


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `CreatedAt`                                | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)   |
| `HasChildren`                              | `*bool`                                    | :heavy_minus_sign:                         | N/A                                        |
| `ID`                                       | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |
| `ParentID`                                 | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |
| `Raw`                                      | map[string]`any`                           | :heavy_minus_sign:                         | N/A                                        |
| `TaskID`                                   | `*string`                                  | :heavy_minus_sign:                         | (reference to TaskTask)                    |
| `Text`                                     | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |
| `UpdatedAt`                                | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)   |
| `UserID`                                   | `*string`                                  | :heavy_minus_sign:                         | (reference to HrisEmployee)                |
| `UserName`                                 | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |