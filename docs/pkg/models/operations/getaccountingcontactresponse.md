# GetAccountingContactResponse


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `AccountingContact`                                                          | [*shared.AccountingContact](../../../pkg/models/shared/accountingcontact.md) | :heavy_minus_sign:                                                           | Successful                                                                   |
| `ContentType`                                                                | *string*                                                                     | :heavy_check_mark:                                                           | HTTP response content type for this operation                                |
| `StatusCode`                                                                 | *int*                                                                        | :heavy_check_mark:                                                           | HTTP response status code for this operation                                 |
| `RawResponse`                                                                | [*http.Response](https://pkg.go.dev/net/http#Response)                       | :heavy_check_mark:                                                           | Raw HTTP response; suitable for custom response parsing                      |