# QmsCharacteristic

A characteristic measured or checked during an inspection


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `LowerLimit`                                             | `*float64`                                               | :heavy_minus_sign:                                       | N/A                                                      |
| `MeasuredValue`                                          | `*string`                                                | :heavy_minus_sign:                                       | Measured or observed value; text so attribute checks fit |
| `Name`                                                   | `*string`                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `NominalValue`                                           | `*float64`                                               | :heavy_minus_sign:                                       | N/A                                                      |
| `Notes`                                                  | `*string`                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `Result`                                                 | [*shared.Result](../../../pkg/models/shared/result.md)   | :heavy_minus_sign:                                       | N/A                                                      |
| `Specification`                                          | `*string`                                                | :heavy_minus_sign:                                       | Requirement or drawing reference                         |
| `UnitOfMeasure`                                          | `*string`                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `UpperLimit`                                             | `*float64`                                               | :heavy_minus_sign:                                       | N/A                                                      |