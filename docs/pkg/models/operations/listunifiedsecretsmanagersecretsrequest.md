# ListUnifiedSecretsmanagerSecretsRequest


## Fields

| Field                                    | Type                                     | Required                                 | Description                              |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| `Limit`                                  | `*float64`                               | :heavy_minus_sign:                       | Maximum number of secrets to return      |
| `Offset`                                 | `*float64`                               | :heavy_minus_sign:                       | Number of secrets to skip                |
| `Query`                                  | `*string`                                | :heavy_minus_sign:                       | Filter secrets by name                   |
| `SecretsmanagerID`                       | `string`                                 | :heavy_check_mark:                       | The secrets manager to list secrets from |