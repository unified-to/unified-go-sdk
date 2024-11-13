# ScimGroup


## Fields

| Field                                                                                       | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `DisplayName`                                                                               | *string*                                                                                    | :heavy_check_mark:                                                                          | N/A                                                                                         |
| `ExternalID`                                                                                | **string*                                                                                   | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `GroupType`                                                                                 | **string*                                                                                   | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `ID`                                                                                        | **string*                                                                                   | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `Members`                                                                                   | [][shared.ScimGroupMember](../../../pkg/models/shared/scimgroupmember.md)                   | :heavy_minus_sign:                                                                          | An array of members                                                                         |
| `Meta`                                                                                      | [*shared.PropertyScimGroupMeta](../../../pkg/models/shared/propertyscimgroupmeta.md)        | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `Schemas`                                                                                   | [][shared.PropertyScimGroupSchemas](../../../pkg/models/shared/propertyscimgroupschemas.md) | :heavy_minus_sign:                                                                          | Array of schema URIs                                                                        |