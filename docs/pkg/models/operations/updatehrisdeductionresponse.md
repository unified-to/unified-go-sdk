# UpdateHrisDeductionResponse


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ContentType`                                                        | *string*                                                             | :heavy_check_mark:                                                   | HTTP response content type for this operation                        |
| `HrisDeduction`                                                      | [*shared.HrisDeduction](../../../pkg/models/shared/hrisdeduction.md) | :heavy_minus_sign:                                                   | Successful                                                           |
| `StatusCode`                                                         | *int*                                                                | :heavy_check_mark:                                                   | HTTP response status code for this operation                         |
| `RawResponse`                                                        | [*http.Response](https://pkg.go.dev/net/http#Response)               | :heavy_check_mark:                                                   | Raw HTTP response; suitable for custom response parsing              |