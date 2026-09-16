# CreateUnifiedWebhookRequest


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `Webhook`                                                                     | [shared.Webhook](../../../pkg/models/shared/webhook.md)                       | :heavy_check_mark:                                                            | A webhook is used to POST new/updated information to your server.             |
| `IncludeAll`                                                                  | `*bool`                                                                       | :heavy_minus_sign:                                                            | When true, send existing/historic data. When false, no historic data is sent. |