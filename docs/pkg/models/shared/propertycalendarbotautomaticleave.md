# PropertyCalendarBotAutomaticLeave


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `DetectBots`                                                           | `*bool`                                                                | :heavy_minus_sign:                                                     | Whether to leave if only bots remain                                   |
| `NoParticipantsTimeout`                                                | `*float64`                                                             | :heavy_minus_sign:                                                     | Seconds to wait when no other participants before leaving (default: 2) |
| `WaitingRoomTimeout`                                                   | `*float64`                                                             | :heavy_minus_sign:                                                     | Seconds to wait in waiting room before leaving (default: 1200)         |