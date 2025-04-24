# HrisDevice


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `AdminUserIds`                             | []*string*                                 | :heavy_minus_sign:                         | N/A                                        |
| `AssetTag`                                 | **string*                                  | :heavy_minus_sign:                         | N/A                                        |
| `CreatedAt`                                | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `HasAntivirus`                             | **bool*                                    | :heavy_minus_sign:                         | N/A                                        |
| `HasFirewall`                              | **bool*                                    | :heavy_minus_sign:                         | N/A                                        |
| `HasHdEncrypted`                           | **bool*                                    | :heavy_minus_sign:                         | N/A                                        |
| `HasPasswordManager`                       | **bool*                                    | :heavy_minus_sign:                         | N/A                                        |
| `HasScreenlock`                            | **bool*                                    | :heavy_minus_sign:                         | N/A                                        |
| `ID`                                       | **string*                                  | :heavy_minus_sign:                         | N/A                                        |
| `IsMissing`                                | **bool*                                    | :heavy_minus_sign:                         | N/A                                        |
| `LocationID`                               | **string*                                  | :heavy_minus_sign:                         | N/A                                        |
| `Manufacturer`                             | **string*                                  | :heavy_minus_sign:                         | N/A                                        |
| `Model`                                    | **string*                                  | :heavy_minus_sign:                         | N/A                                        |
| `Name`                                     | *string*                                   | :heavy_check_mark:                         | N/A                                        |
| `Os`                                       | **string*                                  | :heavy_minus_sign:                         | N/A                                        |
| `OsVersion`                                | **string*                                  | :heavy_minus_sign:                         | N/A                                        |
| `Raw`                                      | map[string]*any*                           | :heavy_minus_sign:                         | N/A                                        |
| `UpdatedAt`                                | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `UserIds`                                  | []*string*                                 | :heavy_minus_sign:                         | users who have this device                 |
| `Version`                                  | **string*                                  | :heavy_minus_sign:                         | N/A                                        |