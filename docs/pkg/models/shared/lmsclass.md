# LmsClass


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `CourseID`                                                  | *string*                                                    | :heavy_check_mark:                                          | N/A                                                         |
| `CreatedAt`                                                 | [*time.Time](https://pkg.go.dev/time#Time)                  | :heavy_minus_sign:                                          | N/A                                                         |
| `Description`                                               | **string*                                                   | :heavy_minus_sign:                                          | N/A                                                         |
| `ID`                                                        | **string*                                                   | :heavy_minus_sign:                                          | N/A                                                         |
| `InstructorIds`                                             | []*string*                                                  | :heavy_minus_sign:                                          | N/A                                                         |
| `Languages`                                                 | []*string*                                                  | :heavy_minus_sign:                                          | N/A                                                         |
| `Media`                                                     | [][shared.LmsMedia](../../../pkg/models/shared/lmsmedia.md) | :heavy_minus_sign:                                          | N/A                                                         |
| `Name`                                                      | *string*                                                    | :heavy_check_mark:                                          | N/A                                                         |
| `Raw`                                                       | map[string]*any*                                            | :heavy_minus_sign:                                          | N/A                                                         |
| `StudentIds`                                                | []*string*                                                  | :heavy_minus_sign:                                          | N/A                                                         |
| `UpdatedAt`                                                 | [*time.Time](https://pkg.go.dev/time#Time)                  | :heavy_minus_sign:                                          | N/A                                                         |