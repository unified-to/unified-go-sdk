# ListUnifiedIssuesRequest


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `Limit`                                                                  | **float64*                                                               | :heavy_minus_sign:                                                       | N/A                                                                      |
| `Offset`                                                                 | **float64*                                                               | :heavy_minus_sign:                                                       | N/A                                                                      |
| `Order`                                                                  | **string*                                                                | :heavy_minus_sign:                                                       | N/A                                                                      |
| `Sort`                                                                   | **string*                                                                | :heavy_minus_sign:                                                       | N/A                                                                      |
| `UpdatedGte`                                                             | [*time.Time](https://pkg.go.dev/time#Time)                               | :heavy_minus_sign:                                                       | Return only results whose updated date is equal or greater to this value |