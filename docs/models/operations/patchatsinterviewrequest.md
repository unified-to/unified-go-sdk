# PatchAtsInterviewRequest


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `AtsInterview`                                              | [*shared.AtsInterview](../../models/shared/atsinterview.md) | :heavy_minus_sign:                                          | An interview between a candidate for a specific job         |
| `ConnectionID`                                              | *string*                                                    | :heavy_check_mark:                                          | ID of the connection                                        |
| `Fields`                                                    | []*string*                                                  | :heavy_minus_sign:                                          | Comma-delimited fields to return                            |
| `ID`                                                        | *string*                                                    | :heavy_check_mark:                                          | ID of the Interview                                         |