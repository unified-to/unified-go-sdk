# PropertyCalendarBotMeetingMetadata


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `EndedAt`                                                       | [*time.Time](https://pkg.go.dev/time#Time)                      | :heavy_minus_sign:                                              | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)                        |
| `HostEmail`                                                     | `*string`                                                       | :heavy_minus_sign:                                              | N/A                                                             |
| `HostName`                                                      | `*string`                                                       | :heavy_minus_sign:                                              | N/A                                                             |
| `MeetingIdentifier`                                             | `*string`                                                       | :heavy_minus_sign:                                              | Platform-specific meeting identifier (not a unified object ref) |
| `ParticipantCount`                                              | `*float64`                                                      | :heavy_minus_sign:                                              | N/A                                                             |
| `Platform`                                                      | `*string`                                                       | :heavy_minus_sign:                                              | 'zoom', 'webex', 'google_meet', etc.                            |
| `StartedAt`                                                     | [*time.Time](https://pkg.go.dev/time#Time)                      | :heavy_minus_sign:                                              | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)                        |
| `Title`                                                         | `*string`                                                       | :heavy_minus_sign:                                              | N/A                                                             |