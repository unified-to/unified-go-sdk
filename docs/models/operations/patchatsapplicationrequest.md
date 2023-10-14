# PatchAtsApplicationRequest


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `AtsApplication`                                                      | [*shared.AtsApplication](../../models/shared/atsapplication.md)       | :heavy_minus_sign:                                                    | An application is an association object between a candidate and a job |
| `ConnectionID`                                                        | *string*                                                              | :heavy_check_mark:                                                    | ID of the connection                                                  |
| `Fields`                                                              | []*string*                                                            | :heavy_minus_sign:                                                    | Comma-delimited fields to return                                      |
| `ID`                                                                  | *string*                                                              | :heavy_check_mark:                                                    | ID of the Application                                                 |