# PropertyCrmEventEmail

The email object, when type = email


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `AttachmentFileIds`                                                | []`string`                                                         | :heavy_minus_sign:                                                 | N/A                                                                |
| `Body`                                                             | `*string`                                                          | :heavy_minus_sign:                                                 | The event email's body                                             |
| `Cc`                                                               | []`string`                                                         | :heavy_minus_sign:                                                 | The event email's cc name & email (name <test@test.com>)           |
| `From`                                                             | `*string`                                                          | :heavy_minus_sign:                                                 | The event email's from name & email address (name <test@test.com>) |
| `Subject`                                                          | `*string`                                                          | :heavy_minus_sign:                                                 | The event email's subject                                          |
| `To`                                                               | []`string`                                                         | :heavy_minus_sign:                                                 | The event email's "to" name & email (name <test@test.com>)         |