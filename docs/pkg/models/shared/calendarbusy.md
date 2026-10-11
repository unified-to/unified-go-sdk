# CalendarBusy


## Fields

| Field                                     | Type                                      | Required                                  | Description                               |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `Description`                             | `*string`                                 | :heavy_minus_sign:                        | N/A                                       |
| `EndAt`                                   | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)  |
| `ID`                                      | `*string`                                 | :heavy_minus_sign:                        | N/A                                       |
| `Raw`                                     | map[string]`any`                          | :heavy_minus_sign:                        | N/A                                       |
| `StartAt`                                 | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)  |
| `Timezone`                                | `*string`                                 | :heavy_minus_sign:                        | N/A                                       |