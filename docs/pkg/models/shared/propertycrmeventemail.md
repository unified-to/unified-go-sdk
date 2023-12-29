# PropertyCrmEventEmail

The email object, when type = email


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `Body`                                                     | **string*                                                  | :heavy_minus_sign:                                         | N/A                                                        |
| `Cc`                                                       | []*string*                                                 | :heavy_minus_sign:                                         | The event email's cc name & email (name <test@test.com>)   |
| `CreatedAt`                                                | [*time.Time](https://pkg.go.dev/time#Time)                 | :heavy_minus_sign:                                         | N/A                                                        |
| `From`                                                     | **string*                                                  | :heavy_minus_sign:                                         | N/A                                                        |
| `Subject`                                                  | **string*                                                  | :heavy_minus_sign:                                         | N/A                                                        |
| `To`                                                       | []*string*                                                 | :heavy_minus_sign:                                         | The event email's "to" name & email (name <test@test.com>) |