# PropertyShippingRateFromAddress

Origin address


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `Address1`                                 | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |
| `Address2`                                 | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |
| `City`                                     | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |
| `CompanyName`                              | `*string`                                  | :heavy_minus_sign:                         | Company name (for commercial addresses)    |
| `Country`                                  | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |
| `CountryCode`                              | `*string`                                  | :heavy_minus_sign:                         | ISO 2-digit country code                   |
| `DeliveryInstructions`                     | `*string`                                  | :heavy_minus_sign:                         | Special delivery notes                     |
| `Email`                                    | `*string`                                  | :heavy_minus_sign:                         | Email address (for delivery notifications) |
| `IsResidential`                            | `*bool`                                    | :heavy_minus_sign:                         | Alias for address_type === 'RESIDENTIAL'   |
| `IsValidated`                              | `*bool`                                    | :heavy_minus_sign:                         | Whether address has been validated         |
| `Name`                                     | `*string`                                  | :heavy_minus_sign:                         | Recipientsender name                       |
| `PostalCode`                               | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |
| `Region`                                   | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |
| `RegionCode`                               | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |
| `Telephone`                                | `*string`                                  | :heavy_minus_sign:                         | Phone number (required by FedEx, UPS, DHL) |