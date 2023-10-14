# PatchMartechMemberRequest


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `MarketingMember`                                                 | [*shared.MarketingMember](../../models/shared/marketingmember.md) | :heavy_minus_sign:                                                | A member represents a person                                      |
| `ConnectionID`                                                    | *string*                                                          | :heavy_check_mark:                                                | ID of the connection                                              |
| `Fields`                                                          | []*string*                                                        | :heavy_minus_sign:                                                | Comma-delimited fields to return                                  |
| `ID`                                                              | *string*                                                          | :heavy_check_mark:                                                | ID of the Member                                                  |
| `ListID`                                                          | *string*                                                          | :heavy_check_mark:                                                | ID of the list                                                    |