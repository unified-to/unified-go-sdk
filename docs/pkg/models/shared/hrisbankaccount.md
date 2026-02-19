# HrisBankaccount

Employee payroll bank account for direct deposit.


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `AccountNumber`                                                  | **string*                                                        | :heavy_minus_sign:                                               | N/A                                                              |
| `AccountNumberLast4`                                             | **string*                                                        | :heavy_minus_sign:                                               | N/A                                                              |
| `AccountType`                                                    | [*shared.AccountType](../../../pkg/models/shared/accounttype.md) | :heavy_minus_sign:                                               | N/A                                                              |
| `BankName`                                                       | **string*                                                        | :heavy_minus_sign:                                               | N/A                                                              |
| `CompanyID`                                                      | **string*                                                        | :heavy_minus_sign:                                               | N/A                                                              |
| `CreatedAt`                                                      | [*time.Time](https://pkg.go.dev/time#Time)                       | :heavy_minus_sign:                                               | N/A                                                              |
| `ID`                                                             | **string*                                                        | :heavy_minus_sign:                                               | N/A                                                              |
| `IsPrimary`                                                      | **bool*                                                          | :heavy_minus_sign:                                               | N/A                                                              |
| `Name`                                                           | **string*                                                        | :heavy_minus_sign:                                               | N/A                                                              |
| `Raw`                                                            | map[string]*any*                                                 | :heavy_minus_sign:                                               | N/A                                                              |
| `RoutingNumber`                                                  | **string*                                                        | :heavy_minus_sign:                                               | N/A                                                              |
| `UpdatedAt`                                                      | [*time.Time](https://pkg.go.dev/time#Time)                       | :heavy_minus_sign:                                               | N/A                                                              |
| `UserID`                                                         | **string*                                                        | :heavy_minus_sign:                                               | N/A                                                              |