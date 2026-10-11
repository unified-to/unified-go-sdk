# AdSchedule


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `BidModifier`                                               | `*float64`                                                  | :heavy_minus_sign:                                          | 0.1-10.0                                                    |
| `DayOfWeek`                                                 | [shared.DayOfWeek](../../../pkg/models/shared/dayofweek.md) | :heavy_check_mark:                                          | N/A                                                         |
| `EndHour`                                                   | `float64`                                                   | :heavy_check_mark:                                          | 0-24 (24 = end of day)                                      |
| `EndMinute`                                                 | `*float64`                                                  | :heavy_minus_sign:                                          | 0, 15, 30, 45                                               |
| `StartHour`                                                 | `float64`                                                   | :heavy_check_mark:                                          | 0-23                                                        |
| `StartMinute`                                               | `*float64`                                                  | :heavy_minus_sign:                                          | 0, 15, 30, 45                                               |