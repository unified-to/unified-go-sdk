# ListComputeUsagesResponse


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ComputeUsages`                                                     | [][shared.ComputeUsage](../../../pkg/models/shared/computeusage.md) | :heavy_minus_sign:                                                  | Successful                                                          |
| `ContentType`                                                       | `string`                                                            | :heavy_check_mark:                                                  | HTTP response content type for this operation                       |
| `StatusCode`                                                        | `int`                                                               | :heavy_check_mark:                                                  | HTTP response status code for this operation                        |
| `RawResponse`                                                       | [*http.Response](https://pkg.go.dev/net/http#Response)              | :heavy_check_mark:                                                  | Raw HTTP response; suitable for custom response parsing             |