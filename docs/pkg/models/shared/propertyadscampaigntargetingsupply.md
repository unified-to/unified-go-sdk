# PropertyAdsCampaignTargetingSupply

Buy-side inventory: exchanges, deals, deal groups and seller authorization (DSPs)


## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `AuthorizedSellerStatus`                                                               | [*shared.AuthorizedSellerStatus](../../../pkg/models/shared/authorizedsellerstatus.md) | :heavy_minus_sign:                                                                     | ads.txt seller authorization to buy from                                               |
| `DealGroups`                                                                           | [][shared.TargetRef](../../../pkg/models/shared/targetref.md)                          | :heavy_minus_sign:                                                                     | Groups of deals (ads_target type DEAL_GROUP)                                           |
| `Deals`                                                                                | [][shared.TargetRef](../../../pkg/models/shared/targetref.md)                          | :heavy_minus_sign:                                                                     | Deals / inventory packages (ads_target type DEAL)                                      |
| `Exchanges`                                                                            | [][shared.TargetRef](../../../pkg/models/shared/targetref.md)                          | :heavy_minus_sign:                                                                     | Public exchanges to buy on (ads_target type EXCHANGE)                                  |