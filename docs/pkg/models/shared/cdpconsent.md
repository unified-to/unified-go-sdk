# CdpConsent


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `Purpose`                                                                         | `*string`                                                                         | :heavy_minus_sign:                                                                | The consent purpose or objective (eg. marketing, analytics)                       |
| `Status`                                                                          | [*shared.CdpConsentStatus](../../../pkg/models/shared/cdpconsentstatus.md)        | :heavy_minus_sign:                                                                | The consent status for this purpose                                               |
| `UpdatedAt`                                                                       | [*time.Time](https://pkg.go.dev/time#Time)                                        | :heavy_minus_sign:                                                                | When the consent status was last updated (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format) |