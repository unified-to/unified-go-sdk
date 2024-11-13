# CreatePassthroughJSONResponse


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `ContentType`                                           | *string*                                                | :heavy_check_mark:                                      | HTTP response content type for this operation           |
| `Headers`                                               | map[string][]*string*                                   | :heavy_check_mark:                                      | N/A                                                     |
| `StatusCode`                                            | *int*                                                   | :heavy_check_mark:                                      | HTTP response status code for this operation            |
| `RawResponse`                                           | [*http.Response](https://pkg.go.dev/net/http#Response)  | :heavy_check_mark:                                      | Raw HTTP response; suitable for custom response parsing |
| `DefaultWildcardWildcardResponseStream`                 | *io.ReadCloser*                                         | :heavy_minus_sign:                                      | Successful                                              |
| `DefaultApplicationJSONAny`                             | *any*                                                   | :heavy_minus_sign:                                      | Successful                                              |
| `DefaultApplicationXMLRes`                              | **string*                                               | :heavy_minus_sign:                                      | Successful                                              |
| `DefaultTextCsvRes`                                     | **string*                                               | :heavy_minus_sign:                                      | Successful                                              |
| `DefaultTextPlainRes`                                   | **string*                                               | :heavy_minus_sign:                                      | Successful                                              |