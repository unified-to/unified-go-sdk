# PatchPaymentLinkRequest


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `PaymentLink`                                                   | [shared.PaymentLink](../../../pkg/models/shared/paymentlink.md) | :heavy_check_mark:                                              | N/A                                                             |
| `ConnectionID`                                                  | *string*                                                        | :heavy_check_mark:                                              | ID of the connection                                            |
| `Fields`                                                        | []*string*                                                      | :heavy_minus_sign:                                              | Comma-delimited fields to return                                |
| `ID`                                                            | *string*                                                        | :heavy_check_mark:                                              | ID of the Link                                                  |