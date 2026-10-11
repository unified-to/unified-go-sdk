# AccountingTaxrateComponent


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `IsCompound`                                                       | `*bool`                                                            | :heavy_minus_sign:                                                 | True when this component is charged on top of the other components |
| `Name`                                                             | `*string`                                                          | :heavy_minus_sign:                                                 | Name of the tax component, e.g. GST or PST                         |
| `Rate`                                                             | `*float64`                                                         | :heavy_minus_sign:                                                 | Rate of this component, in %                                       |