# Webhook

A webhook is used to POST new/updated information to your server.


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `CheckedAt`                                                             | [*time.Time](https://pkg.go.dev/time#Time)                              | :heavy_minus_sign:                                                      | N/A                                                                     |
| `ConnectionID`                                                          | *string*                                                                | :heavy_check_mark:                                                      | N/A                                                                     |
| `CreatedAt`                                                             | [*time.Time](https://pkg.go.dev/time#Time)                              | :heavy_minus_sign:                                                      | N/A                                                                     |
| `Environment`                                                           | **string*                                                               | :heavy_minus_sign:                                                      | N/A                                                                     |
| `Events`                                                                | [][PropertyWebhookEvents](../../models/shared/propertywebhookevents.md) | :heavy_check_mark:                                                      | N/A                                                                     |
| `HookURL`                                                               | *string*                                                                | :heavy_check_mark:                                                      | N/A                                                                     |
| `ID`                                                                    | **string*                                                               | :heavy_minus_sign:                                                      | N/A                                                                     |
| `IncludeRaw`                                                            | **bool*                                                                 | :heavy_minus_sign:                                                      | N/A                                                                     |
| `IntegrationType`                                                       | *string*                                                                | :heavy_check_mark:                                                      | N/A                                                                     |
| `Interval`                                                              | *float64*                                                               | :heavy_check_mark:                                                      | N/A                                                                     |
| `ObjectType`                                                            | [WebhookObjectType](../../models/shared/webhookobjecttype.md)           | :heavy_check_mark:                                                      | N/A                                                                     |
| `Subscriptions`                                                         | []*string*                                                              | :heavy_minus_sign:                                                      | integration-specific subscriptions IDs                                  |
| `UpdatedAt`                                                             | [*time.Time](https://pkg.go.dev/time#Time)                              | :heavy_minus_sign:                                                      | N/A                                                                     |
| `WorkspaceID`                                                           | *string*                                                                | :heavy_check_mark:                                                      | N/A                                                                     |