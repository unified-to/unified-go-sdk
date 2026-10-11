# GenaiSkill


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `CreatedAt`                                              | [*time.Time](https://pkg.go.dev/time#Time)               | :heavy_minus_sign:                                       | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)                 |
| `Description`                                            | `*string`                                                | :heavy_minus_sign:                                       | What the skill does and when to use it.                  |
| `ID`                                                     | `*string`                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `Name`                                                   | `*string`                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `Raw`                                                    | map[string]`any`                                         | :heavy_minus_sign:                                       | N/A                                                      |
| `UpdatedAt`                                              | [*time.Time](https://pkg.go.dev/time#Time)               | :heavy_minus_sign:                                       | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)                 |
| `Version`                                                | `*string`                                                | :heavy_minus_sign:                                       | The version that is used when a prompt does not pin one. |