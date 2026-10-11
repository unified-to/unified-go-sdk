# PerformanceGoalMilestone


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `CurrentValue`                                         | `*float64`                                             | :heavy_minus_sign:                                     | N/A                                                    |
| `Description`                                          | `*string`                                              | :heavy_minus_sign:                                     | N/A                                                    |
| `DueAt`                                                | [*time.Time](https://pkg.go.dev/time#Time)             | :heavy_minus_sign:                                     | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)               |
| `ID`                                                   | `*string`                                              | :heavy_minus_sign:                                     | N/A                                                    |
| `IsCompleted`                                          | `*bool`                                                | :heavy_minus_sign:                                     | N/A                                                    |
| `Name`                                                 | `string`                                               | :heavy_check_mark:                                     | N/A                                                    |
| `TargetValue`                                          | `*float64`                                             | :heavy_minus_sign:                                     | N/A                                                    |
| `Unit`                                                 | `*string`                                              | :heavy_minus_sign:                                     | unit of the target/current values (e.g. %, USD, count) |
| `Weight`                                               | `*float64`                                             | :heavy_minus_sign:                                     | N/A                                                    |