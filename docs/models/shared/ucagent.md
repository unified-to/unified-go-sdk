# UcAgent

Represents an agent


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `CreatedAt`                                                     | [*types.Date](../../types/date.md)                              | :heavy_minus_sign:                                              | N/A                                                             |
| `Emails`                                                        | [][UcEmail](../../models/shared/ucemail.md)                     | :heavy_minus_sign:                                              | An array of email addresses for this agent                      |
| `ID`                                                            | **string*                                                       | :heavy_minus_sign:                                              | N/A                                                             |
| `Name`                                                          | **string*                                                       | :heavy_minus_sign:                                              | N/A                                                             |
| `Raw`                                                           | [PropertyUcAgentRaw](../../models/shared/propertyucagentraw.md) | :heavy_check_mark:                                              | The raw data returned by the integration for this agent         |
| `Telephones`                                                    | [][UcTelephone](../../models/shared/uctelephone.md)             | :heavy_minus_sign:                                              | N/A                                                             |
| `UpdatedAt`                                                     | [*types.Date](../../types/date.md)                              | :heavy_minus_sign:                                              | N/A                                                             |