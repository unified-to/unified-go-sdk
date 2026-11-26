# DmsVehiclesaleTradein

A vehicle traded in on a vehicle sale


## Fields

| Field                                            | Type                                             | Required                                         | Description                                      |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| `ActualCashValueAmount`                          | `*float64`                                       | :heavy_minus_sign:                               | Actual cash value (dealer valuation)             |
| `AllowanceAmount`                                | `*float64`                                       | :heavy_minus_sign:                               | Trade allowance given to the buyer               |
| `LienholderName`                                 | `*string`                                        | :heavy_minus_sign:                               | Lender owed the payoff                           |
| `Make`                                           | `*string`                                        | :heavy_minus_sign:                               | N/A                                              |
| `Model`                                          | `*string`                                        | :heavy_minus_sign:                               | N/A                                              |
| `Odometer`                                       | `*float64`                                       | :heavy_minus_sign:                               | N/A                                              |
| `PayoffAmount`                                   | `*float64`                                       | :heavy_minus_sign:                               | Outstanding loan payoff                          |
| `VehicleID`                                      | `*string`                                        | :heavy_minus_sign:                               | The DmsVehicle ID of the trade-in, when recorded |
| `Vin`                                            | `*string`                                        | :heavy_minus_sign:                               | N/A                                              |
| `Year`                                           | `*float64`                                       | :heavy_minus_sign:                               | N/A                                              |