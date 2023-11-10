# CreateUnifiedWebhookRequest


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `Webhook`                                                       | [*shared.Webhook](../../../pkg/models/shared/webhook.md)        | :heavy_minus_sign:                                              | N/A                                                             |
| `ConnectionID`                                                  | *string*                                                        | :heavy_check_mark:                                              | ID of the connection                                            |
| `Events`                                                        | [][operations.Events](../../../pkg/models/operations/events.md) | :heavy_minus_sign:                                              | Which events to subscribe to.                                   |
| `Object`                                                        | *string*                                                        | :heavy_check_mark:                                              | The object to subscribe to                                      |