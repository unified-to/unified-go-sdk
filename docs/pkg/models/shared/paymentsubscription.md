# PaymentSubscription


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `CanceledAt`                                                                                 | [*time.Time](https://pkg.go.dev/time#Time)                                                   | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `ContactID`                                                                                  | **string*                                                                                    | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `CreatedAt`                                                                                  | [*time.Time](https://pkg.go.dev/time#Time)                                                   | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `Currency`                                                                                   | **string*                                                                                    | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `CurrentPeriodEndAt`                                                                         | [*time.Time](https://pkg.go.dev/time#Time)                                                   | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `CurrentPeriodStartAt`                                                                       | [*time.Time](https://pkg.go.dev/time#Time)                                                   | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `DayOfMonth`                                                                                 | **float64*                                                                                   | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `DayOfWeek`                                                                                  | **float64*                                                                                   | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `Description`                                                                                | **string*                                                                                    | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `EndAt`                                                                                      | [*time.Time](https://pkg.go.dev/time#Time)                                                   | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `ID`                                                                                         | **string*                                                                                    | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `Interval`                                                                                   | **float64*                                                                                   | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `IntervalUnit`                                                                               | [*shared.IntervalUnit](../../../pkg/models/shared/intervalunit.md)                           | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `InvoiceID`                                                                                  | **string*                                                                                    | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `Lineitems`                                                                                  | [][shared.PaymentLineitem](../../../pkg/models/shared/paymentlineitem.md)                    | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `Month`                                                                                      | **float64*                                                                                   | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `Raw`                                                                                        | map[string]*any*                                                                             | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `StartAt`                                                                                    | [*time.Time](https://pkg.go.dev/time#Time)                                                   | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `Status`                                                                                     | [*shared.PaymentSubscriptionStatus](../../../pkg/models/shared/paymentsubscriptionstatus.md) | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `UpdatedAt`                                                                                  | [*time.Time](https://pkg.go.dev/time#Time)                                                   | :heavy_minus_sign:                                                                           | N/A                                                                                          |