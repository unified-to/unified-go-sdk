# DmsVehiclesaleFee

A fee on a vehicle sale (documentation, title, registration, lender, dealer)


## Fields

| Field                                     | Type                                      | Required                                  | Description                               |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `CostAmount`                              | `*float64`                                | :heavy_minus_sign:                        | Dealer cost                               |
| `IsCapitalized`                           | `*bool`                                   | :heavy_minus_sign:                        | Rolled into the financed or leased amount |
| `Name`                                    | `*string`                                 | :heavy_minus_sign:                        | N/A                                       |
| `PriceAmount`                             | `*float64`                                | :heavy_minus_sign:                        | Amount charged to the buyer               |