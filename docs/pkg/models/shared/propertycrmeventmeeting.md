# PropertyCrmEventMeeting

The meeting object, when type = meeting


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `Description`                                                               | `*string`                                                                   | :heavy_minus_sign:                                                          | The event meeting's agendadescription                                       |
| `EndAt`                                                                     | [*time.Time](https://pkg.go.dev/time#Time)                                  | :heavy_minus_sign:                                                          | The event meeting's end datetime (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)   |
| `StartAt`                                                                   | [*time.Time](https://pkg.go.dev/time#Time)                                  | :heavy_minus_sign:                                                          | The event meeting's start datetime (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format) |
| `Title`                                                                     | `*string`                                                                   | :heavy_minus_sign:                                                          | The event meeting's topicsubject                                            |