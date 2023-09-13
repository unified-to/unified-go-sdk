# User

The User object represents you on the system. A user can belong to multiple workspaces (ie. organizations).


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `CreatedAt`                                                  | [*types.Date](../../types/date.md)                           | :heavy_minus_sign:                                           | N/A                                                          |
| `Email`                                                      | **string*                                                    | :heavy_minus_sign:                                           | N/A                                                          |
| `Environment`                                                | **string*                                                    | :heavy_minus_sign:                                           | N/A                                                          |
| `ID`                                                         | **string*                                                    | :heavy_minus_sign:                                           | N/A                                                          |
| `Meta`                                                       | [*PropertyUserMeta](../../models/shared/propertyusermeta.md) | :heavy_minus_sign:                                           | N/A                                                          |
| `Name`                                                       | *string*                                                     | :heavy_check_mark:                                           | N/A                                                          |
| `UpdatedAt`                                                  | [*types.Date](../../types/date.md)                           | :heavy_minus_sign:                                           | N/A                                                          |
| `WorkspaceID`                                                | *string*                                                     | :heavy_check_mark:                                           | N/A                                                          |
| `WorkspaceIds`                                               | []*string*                                                   | :heavy_check_mark:                                           | A list of all of the user's workspaces                       |