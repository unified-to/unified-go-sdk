# CreateMartechMemberRequest


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `MarketingMember`                                                       | [shared.MarketingMember](../../../pkg/models/shared/marketingmember.md) | :heavy_check_mark:                                                      | A member represents a person                                            |
| `ConnectionID`                                                          | *string*                                                                | :heavy_check_mark:                                                      | ID of the connection                                                    |
| `Fields`                                                                | []*string*                                                              | :heavy_minus_sign:                                                      | Comma-delimited fields to return                                        |