# UcComment


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `CallID`                                                        | `*string`                                                       | :heavy_minus_sign:                                              | N/A                                                             |
| `Content`                                                       | `*string`                                                       | :heavy_minus_sign:                                              | N/A                                                             |
| `CreatedAt`                                                     | `*string`                                                       | :heavy_minus_sign:                                              | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)                        |
| `ID`                                                            | `*string`                                                       | :heavy_minus_sign:                                              | N/A                                                             |
| `Metadata`                                                      | [][shared.UcMetadata](../../../pkg/models/shared/ucmetadata.md) | :heavy_minus_sign:                                              | N/A                                                             |
| `Raw`                                                           | map[string]`any`                                                | :heavy_minus_sign:                                              | N/A                                                             |
| `UpdatedAt`                                                     | `*string`                                                       | :heavy_minus_sign:                                              | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)                        |
| `UserID`                                                        | `*string`                                                       | :heavy_minus_sign:                                              | reference to HrisEmployee (reference to HrisEmployee)           |