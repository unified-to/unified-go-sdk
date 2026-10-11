# UcRecordingTranscript


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ContactID`                                           | `*string`                                             | :heavy_minus_sign:                                    | (reference to UcContact)                              |
| `EndAt`                                               | [*time.Time](https://pkg.go.dev/time#Time)            | :heavy_minus_sign:                                    | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)              |
| `StartAt`                                             | [*time.Time](https://pkg.go.dev/time#Time)            | :heavy_minus_sign:                                    | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)              |
| `Text`                                                | `string`                                              | :heavy_check_mark:                                    | N/A                                                   |
| `UserID`                                              | `*string`                                             | :heavy_minus_sign:                                    | reference to HrisEmployee (reference to HrisEmployee) |