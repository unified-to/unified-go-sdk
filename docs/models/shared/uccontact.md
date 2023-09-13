# UcContact

A contact represents a person that optionally is associated with a call


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `Company`                                                            | **string*                                                            | :heavy_minus_sign:                                                   | N/A                                                                  |
| `CreatedAt`                                                          | [*types.Date](../../types/date.md)                                   | :heavy_minus_sign:                                                   | N/A                                                                  |
| `Emails`                                                             | [][UcEmail](../../models/shared/ucemail.md)                          | :heavy_minus_sign:                                                   | An array of email addresses for this contact                         |
| `ID`                                                                 | **string*                                                            | :heavy_minus_sign:                                                   | N/A                                                                  |
| `Name`                                                               | **string*                                                            | :heavy_minus_sign:                                                   | N/A                                                                  |
| `Raw`                                                                | [*PropertyUcContactRaw](../../models/shared/propertyuccontactraw.md) | :heavy_minus_sign:                                                   | The raw data returned by the integration for this contact            |
| `Telephones`                                                         | [][UcTelephone](../../models/shared/uctelephone.md)                  | :heavy_minus_sign:                                                   | An array of telephones for this contact                              |
| `Title`                                                              | **string*                                                            | :heavy_minus_sign:                                                   | N/A                                                                  |
| `UpdatedAt`                                                          | [*types.Date](../../types/date.md)                                   | :heavy_minus_sign:                                                   | N/A                                                                  |