# UcContact

A contact represents a person that optionally is associated with a call


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `Company`                                                         | **string*                                                         | :heavy_minus_sign:                                                | N/A                                                               |
| `CreatedAt`                                                       | [*time.Time](https://pkg.go.dev/time#Time)                        | :heavy_minus_sign:                                                | N/A                                                               |
| `Emails`                                                          | [][shared.UcEmail](../../../pkg/models/shared/ucemail.md)         | :heavy_minus_sign:                                                | An array of email addresses for this contact                      |
| `ID`                                                              | **string*                                                         | :heavy_minus_sign:                                                | N/A                                                               |
| `Name`                                                            | **string*                                                         | :heavy_minus_sign:                                                | N/A                                                               |
| `Raw`                                                             | map[string]*any*                                                  | :heavy_minus_sign:                                                | N/A                                                               |
| `Telephones`                                                      | [][shared.UcTelephone](../../../pkg/models/shared/uctelephone.md) | :heavy_minus_sign:                                                | An array of telephones for this contact                           |
| `Title`                                                           | **string*                                                         | :heavy_minus_sign:                                                | N/A                                                               |
| `UpdatedAt`                                                       | [*time.Time](https://pkg.go.dev/time#Time)                        | :heavy_minus_sign:                                                | N/A                                                               |