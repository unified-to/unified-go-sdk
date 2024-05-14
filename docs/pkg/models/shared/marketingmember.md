# MarketingMember

A member represents a person


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `CreatedAt`                                                             | [*time.Time](https://pkg.go.dev/time#Time)                              | :heavy_minus_sign:                                                      | N/A                                                                     |
| `Emails`                                                                | [][shared.MarketingEmail](../../../pkg/models/shared/marketingemail.md) | :heavy_minus_sign:                                                      | An array of email addresses for this member                             |
| `ID`                                                                    | **string*                                                               | :heavy_minus_sign:                                                      | N/A                                                                     |
| `ListIds`                                                               | []*string*                                                              | :heavy_minus_sign:                                                      | An array of list IDs associated with this member                        |
| `Name`                                                                  | **string*                                                               | :heavy_minus_sign:                                                      | N/A                                                                     |
| `Raw`                                                                   | map[string]*any*                                                        | :heavy_minus_sign:                                                      | The raw data returned by the integration for this member                |
| `Tags`                                                                  | []*string*                                                              | :heavy_minus_sign:                                                      | An array of tags associated with this member                            |
| `UpdatedAt`                                                             | [*time.Time](https://pkg.go.dev/time#Time)                              | :heavy_minus_sign:                                                      | N/A                                                                     |