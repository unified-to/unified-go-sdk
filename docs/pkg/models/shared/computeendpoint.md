# ComputeEndpoint


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `IsAuthRequired`                                           | `*bool`                                                    | :heavy_minus_sign:                                         | Whether the URL requires a vendor access token             |
| `Port`                                                     | `*float64`                                                 | :heavy_minus_sign:                                         | The port inside the instance                               |
| `Protocol`                                                 | [*shared.Protocol](../../../pkg/models/shared/protocol.md) | :heavy_minus_sign:                                         | The protocol served on this port                           |
| `URL`                                                      | `*string`                                                  | :heavy_minus_sign:                                         | Public or proxied URL for this port                        |