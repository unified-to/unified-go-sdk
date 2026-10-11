# PropertyShippingShipmentInsurance

Insurance details


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `CoverageAmount`                                                   | `*float64`                                                         | :heavy_minus_sign:                                                 | Coverage amount (may differ from insured_value)                    |
| `CoverageType`                                                     | [*shared.CoverageType](../../../pkg/models/shared/coveragetype.md) | :heavy_minus_sign:                                                 | N/A                                                                |
| `Currency`                                                         | `*string`                                                          | :heavy_minus_sign:                                                 | N/A                                                                |
| `InsuranceCost`                                                    | `*float64`                                                         | :heavy_minus_sign:                                                 | Cost of insurance                                                  |
| `InsuranceCostCurrency`                                            | `*string`                                                          | :heavy_minus_sign:                                                 | N/A                                                                |
| `InsuranceProvider`                                                | `*string`                                                          | :heavy_minus_sign:                                                 | Insurance provider name                                            |
| `InsuranceProviderCode`                                            | `*string`                                                          | :heavy_minus_sign:                                                 | Provider code                                                      |
| `InsuredValue`                                                     | `*float64`                                                         | :heavy_minus_sign:                                                 | Insured value                                                      |