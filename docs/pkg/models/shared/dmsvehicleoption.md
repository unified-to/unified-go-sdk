# DmsVehicleOption

An installed vehicle option or package


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `Code`                                                                             | `*string`                                                                          | :heavy_minus_sign:                                                                 | Option or RPO code                                                                 |
| `CostAmount`                                                                       | `*float64`                                                                         | :heavy_minus_sign:                                                                 | Dealer cost                                                                        |
| `Description`                                                                      | `*string`                                                                          | :heavy_minus_sign:                                                                 | N/A                                                                                |
| `PriceAmount`                                                                      | `*float64`                                                                         | :heavy_minus_sign:                                                                 | Retail price                                                                       |
| `Type`                                                                             | [*shared.DmsVehicleOptionType](../../../pkg/models/shared/dmsvehicleoptiontype.md) | :heavy_minus_sign:                                                                 | Factory-fitted or dealer-added                                                     |