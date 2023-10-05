# Integration

Informational object for supported integrations.


## Fields

| Field                                                                                   | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `APIDocsURL`                                                                            | **string*                                                                               | :heavy_minus_sign:                                                                      | N/A                                                                                     |
| `Beta`                                                                                  | **bool*                                                                                 | :heavy_minus_sign:                                                                      | N/A                                                                                     |
| `Categories`                                                                            | [][PropertyIntegrationCategories](../../models/shared/propertyintegrationcategories.md) | :heavy_check_mark:                                                                      | The categories of support solutions that this integration has                           |
| `Color`                                                                                 | **string*                                                                               | :heavy_minus_sign:                                                                      | N/A                                                                                     |
| `CreatedAt`                                                                             | [*time.Time](https://pkg.go.dev/time#Time)                                              | :heavy_minus_sign:                                                                      | N/A                                                                                     |
| `FaIcon`                                                                                | **string*                                                                               | :heavy_minus_sign:                                                                      | N/A                                                                                     |
| `InProgress`                                                                            | *bool*                                                                                  | :heavy_check_mark:                                                                      | N/A                                                                                     |
| `IsActive`                                                                              | **bool*                                                                                 | :heavy_minus_sign:                                                                      | N/A                                                                                     |
| `LogoURL`                                                                               | **string*                                                                               | :heavy_minus_sign:                                                                      | N/A                                                                                     |
| `Name`                                                                                  | *string*                                                                                | :heavy_check_mark:                                                                      | N/A                                                                                     |
| `RateLimitDescription`                                                                  | **string*                                                                               | :heavy_minus_sign:                                                                      | N/A                                                                                     |
| `Support`                                                                               | map[string][IntegrationSupport](../../models/shared/integrationsupport.md)              | :heavy_check_mark:                                                                      | N/A                                                                                     |
| `TestedAt`                                                                              | [*time.Time](https://pkg.go.dev/time#Time)                                              | :heavy_minus_sign:                                                                      | N/A                                                                                     |
| `TextColor`                                                                             | **string*                                                                               | :heavy_minus_sign:                                                                      | N/A                                                                                     |
| `TokenInstructions`                                                                     | []*string*                                                                              | :heavy_minus_sign:                                                                      | instructions for the user on how to find the token/key                                  |
| `TokenNames`                                                                            | []*string*                                                                              | :heavy_minus_sign:                                                                      | if auth_types = 'token'                                                                 |
| `Type`                                                                                  | *string*                                                                                | :heavy_check_mark:                                                                      | N/A                                                                                     |
| `UpdatedAt`                                                                             | [*time.Time](https://pkg.go.dev/time#Time)                                              | :heavy_minus_sign:                                                                      | N/A                                                                                     |
| `WebURL`                                                                                | **string*                                                                               | :heavy_minus_sign:                                                                      | N/A                                                                                     |