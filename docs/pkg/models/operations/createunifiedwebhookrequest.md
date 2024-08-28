# CreateUnifiedWebhookRequest


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `Webhook`                                                         | [*shared.Webhook](../../../pkg/models/shared/webhook.md)          | :heavy_minus_sign:                                                | A webhook is used to POST new/updated information to your server. |
| `IncludeAll`                                                      | **bool*                                                           | :heavy_minus_sign:                                                | When set, all of the existing data will sent back to your server. |