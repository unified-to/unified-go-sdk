# HrisAttendanceBreak


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `DurationMinutes`                          | `*float64`                                 | :heavy_minus_sign:                         | N/A                                        |
| `EndAt`                                    | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)   |
| `ID`                                       | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |
| `IsPaid`                                   | `*bool`                                    | :heavy_minus_sign:                         | N/A                                        |
| `Name`                                     | `*string`                                  | :heavy_minus_sign:                         | break type label, e.g. "Lunch"             |
| `StartAt`                                  | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)   |