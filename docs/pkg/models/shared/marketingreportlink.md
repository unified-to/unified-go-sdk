# MarketingReportLink


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ClickPercentage`                                             | `*float64`                                                    | :heavy_minus_sign:                                            | Percentage of total clicks                                    |
| `LastClickAt`                                                 | [*time.Time](https://pkg.go.dev/time#Time)                    | :heavy_minus_sign:                                            | Last click timestamp (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format) |
| `TotalClicks`                                                 | `*float64`                                                    | :heavy_minus_sign:                                            | Total clicks on this link                                     |
| `UniqueClickPercentage`                                       | `*float64`                                                    | :heavy_minus_sign:                                            | Percentage of unique clicks                                   |
| `UniqueClicks`                                                | `*float64`                                                    | :heavy_minus_sign:                                            | Unique clicks on this link                                    |
| `URL`                                                         | `*string`                                                     | :heavy_minus_sign:                                            | The URL                                                       |