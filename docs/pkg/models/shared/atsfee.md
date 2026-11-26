# AtsFee


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `Amount`                                                        | `*float64`                                                      | :heavy_minus_sign:                                              | Flat fee amount, in `currency`                                  |
| `Currency`                                                      | `*string`                                                       | :heavy_minus_sign:                                              | ISO 4217 currency code of `amount`                              |
| `Percentage`                                                    | `*float64`                                                      | :heavy_minus_sign:                                              | Fee as a percentage of the candidate's salary (e.g. 20 for 20%) |
| `Type`                                                          | [*shared.AtsFeeType](../../../pkg/models/shared/atsfeetype.md)  | :heavy_minus_sign:                                              | The fee arrangement (e.g. contingency or retained search)       |