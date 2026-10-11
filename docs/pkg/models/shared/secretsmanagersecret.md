# SecretsManagerSecret


## Fields

| Field                                        | Type                                         | Required                                     | Description                                  |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| `CreatedAt`                                  | [*time.Time](https://pkg.go.dev/time#Time)   | :heavy_minus_sign:                           | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)     |
| `Description`                                | `*string`                                    | :heavy_minus_sign:                           | N/A                                          |
| `ID`                                         | `*string`                                    | :heavy_minus_sign:                           | maps to secretsmanager_id in ISecretsManager |
| `Name`                                       | `*string`                                    | :heavy_minus_sign:                           | N/A                                          |
| `UpdatedAt`                                  | [*time.Time](https://pkg.go.dev/time#Time)   | :heavy_minus_sign:                           | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)     |