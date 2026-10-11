# PropertySocialReviewReply

The business's reply to the review


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `Content`                                                                | `*string`                                                                | :heavy_minus_sign:                                                       | The reply text from the business                                         |
| `CreatedAt`                                                              | [*time.Time](https://pkg.go.dev/time#Time)                               | :heavy_minus_sign:                                                       | When the reply was created (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)      |
| `UpdatedAt`                                                              | [*time.Time](https://pkg.go.dev/time#Time)                               | :heavy_minus_sign:                                                       | When the reply was last updated (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format) |