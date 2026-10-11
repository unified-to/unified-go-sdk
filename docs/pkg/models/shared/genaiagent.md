# GenaiAgent


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `CreatedAt`                                                   | [*time.Time](https://pkg.go.dev/time#Time)                    | :heavy_minus_sign:                                            | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)                      |
| `Description`                                                 | `*string`                                                     | :heavy_minus_sign:                                            | N/A                                                           |
| `ID`                                                          | `*string`                                                     | :heavy_minus_sign:                                            | N/A                                                           |
| `Instructions`                                                | `*string`                                                     | :heavy_minus_sign:                                            | The system prompt or stored procedure that the agent follows. |
| `IsActive`                                                    | `*bool`                                                       | :heavy_minus_sign:                                            | N/A                                                           |
| `ModelID`                                                     | `*string`                                                     | :heavy_minus_sign:                                            | The model that the agent runs on.                             |
| `Name`                                                        | `*string`                                                     | :heavy_minus_sign:                                            | N/A                                                           |
| `Raw`                                                         | map[string]`any`                                              | :heavy_minus_sign:                                            | N/A                                                           |
| `UpdatedAt`                                                   | [*time.Time](https://pkg.go.dev/time#Time)                    | :heavy_minus_sign:                                            | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)                      |