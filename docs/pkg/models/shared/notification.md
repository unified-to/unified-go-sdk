# Notification

A notification of an event that occurred in you account.


## Fields

| Field                                                | Type                                                 | Required                                             | Description                                          |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `ConnectionID`                                       | `*string`                                            | :heavy_minus_sign:                                   | N/A                                                  |
| `CreatedAt`                                          | [*time.Time](https://pkg.go.dev/time#Time)           | :heavy_minus_sign:                                   | N/A                                                  |
| `Description`                                        | `*string`                                            | :heavy_minus_sign:                                   | N/A                                                  |
| `Event`                                              | [*shared.Event](../../../pkg/models/shared/event.md) | :heavy_minus_sign:                                   | N/A                                                  |
| `ID`                                                 | `*string`                                            | :heavy_minus_sign:                                   | N/A                                                  |
| `IntegrationName`                                    | `*string`                                            | :heavy_minus_sign:                                   | N/A                                                  |
| `IntegrationType`                                    | `*string`                                            | :heavy_minus_sign:                                   | N/A                                                  |
| `SentAt`                                             | [*time.Time](https://pkg.go.dev/time#Time)           | :heavy_minus_sign:                                   | N/A                                                  |
| `UpdatedAt`                                          | [*time.Time](https://pkg.go.dev/time#Time)           | :heavy_minus_sign:                                   | N/A                                                  |
| `UserID`                                             | `*string`                                            | :heavy_minus_sign:                                   | N/A                                                  |
| `UserName`                                           | `*string`                                            | :heavy_minus_sign:                                   | N/A                                                  |
| `WebhookID`                                          | `*string`                                            | :heavy_minus_sign:                                   | N/A                                                  |
| `WorkspaceID`                                        | `*string`                                            | :heavy_minus_sign:                                   | N/A                                                  |
| `WorkspaceName`                                      | `*string`                                            | :heavy_minus_sign:                                   | N/A                                                  |