# CreateUcContactRequest


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `UcContact`                                                             | [*shared.UcContact](../../../pkg/models/shared/uccontact.md)            | :heavy_minus_sign:                                                      | A contact represents a person that optionally is associated with a call |
| `ConnectionID`                                                          | *string*                                                                | :heavy_check_mark:                                                      | ID of the connection                                                    |
| `Fields`                                                                | []*string*                                                              | :heavy_minus_sign:                                                      | Comma-delimited fields to return                                        |