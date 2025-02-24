# CreatePaymentSubscriptionRequest


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `PaymentSubscription`                                                           | [shared.PaymentSubscription](../../../pkg/models/shared/paymentsubscription.md) | :heavy_check_mark:                                                              | N/A                                                                             |
| `ConnectionID`                                                                  | *string*                                                                        | :heavy_check_mark:                                                              | ID of the connection                                                            |
| `Fields`                                                                        | []*string*                                                                      | :heavy_minus_sign:                                                              | Comma-delimited fields to return                                                |