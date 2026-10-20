# RepoBranch


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `CreatedAt`                                                         | [*time.Time](https://pkg.go.dev/time#Time)                          | :heavy_minus_sign:                                                  | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)                            |
| `ID`                                                                | `*string`                                                           | :heavy_minus_sign:                                                  | N/A                                                                 |
| `Metadata`                                                          | [][shared.RepoMetadata](../../../pkg/models/shared/repometadata.md) | :heavy_minus_sign:                                                  | N/A                                                                 |
| `Name`                                                              | `string`                                                            | :heavy_check_mark:                                                  | N/A                                                                 |
| `Raw`                                                               | map[string]`any`                                                    | :heavy_minus_sign:                                                  | N/A                                                                 |
| `RepoID`                                                            | `string`                                                            | :heavy_check_mark:                                                  | (reference to RepoRepository)                                       |
| `UpdatedAt`                                                         | [*time.Time](https://pkg.go.dev/time#Time)                          | :heavy_minus_sign:                                                  | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)                            |