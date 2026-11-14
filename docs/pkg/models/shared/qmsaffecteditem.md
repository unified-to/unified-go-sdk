# QmsAffectedItem

A part, lot or serial affected by a quality record


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `DocumentID`                                                      | `*string`                                                         | :heavy_minus_sign:                                                | The PlmDocument ID, when the same vendor also provides PLM        |
| `LotNumber`                                                       | `*string`                                                         | :heavy_minus_sign:                                                | N/A                                                               |
| `PartID`                                                          | `*string`                                                         | :heavy_minus_sign:                                                | The PlmPart (revision) ID, when the same vendor also provides PLM |
| `PartNumber`                                                      | `*string`                                                         | :heavy_minus_sign:                                                | N/A                                                               |
| `Quantity`                                                        | `*float64`                                                        | :heavy_minus_sign:                                                | Quantity affected                                                 |
| `Revision`                                                        | `*string`                                                         | :heavy_minus_sign:                                                | N/A                                                               |
| `SerialNumber`                                                    | `*string`                                                         | :heavy_minus_sign:                                                | N/A                                                               |