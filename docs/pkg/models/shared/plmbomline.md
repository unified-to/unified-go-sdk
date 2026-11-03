# PlmBomline

One parent to child usage in a bill of materials


## Fields

| Field                                         | Type                                          | Required                                      | Description                                   |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| `ChildPartID`                                 | `*string`                                     | :heavy_minus_sign:                            | The component's PlmPart ID (a revision ID)    |
| `ChildRevision`                               | `*string`                                     | :heavy_minus_sign:                            | Pinned child revision, when the line pins one |
| `CreatedAt`                                   | [*time.Time](https://pkg.go.dev/time#Time)    | :heavy_minus_sign:                            | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)      |
| `EffectiveAt`                                 | [*time.Time](https://pkg.go.dev/time#Time)    | :heavy_minus_sign:                            | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)      |
| `ExpiresAt`                                   | [*time.Time](https://pkg.go.dev/time#Time)    | :heavy_minus_sign:                            | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)      |
| `ID`                                          | `*string`                                     | :heavy_minus_sign:                            | N/A                                           |
| `LineNumber`                                  | `*string`                                     | :heavy_minus_sign:                            | Find or item number on the BOM                |
| `Notes`                                       | `*string`                                     | :heavy_minus_sign:                            | N/A                                           |
| `ParentPartID`                                | `*string`                                     | :heavy_minus_sign:                            | The assembly                                  |
| `Quantity`                                    | `*float64`                                    | :heavy_minus_sign:                            | N/A                                           |
| `Raw`                                         | map[string]`any`                              | :heavy_minus_sign:                            | Raw data from integration                     |
| `ReferenceDesignators`                        | []`string`                                    | :heavy_minus_sign:                            | N/A                                           |
| `UnitOfMeasure`                               | `*string`                                     | :heavy_minus_sign:                            | N/A                                           |
| `UpdatedAt`                                   | [*time.Time](https://pkg.go.dev/time#Time)    | :heavy_minus_sign:                            | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)      |