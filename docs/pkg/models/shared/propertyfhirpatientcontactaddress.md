# PropertyFhirPatientContactAddress

Address for the contact person


## Fields

| Field                                             | Type                                              | Required                                          | Description                                       |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `City`                                            | `*string`                                         | :heavy_minus_sign:                                | Name of city, town etc.                           |
| `Country`                                         | `*string`                                         | :heavy_minus_sign:                                | Country (e.g. can be ISO 3166 2 or 3 letter code) |
| `Line`                                            | []`string`                                        | :heavy_minus_sign:                                | Street name, number, direction and P.O. Box etc.  |
| `PostalCode`                                      | `*string`                                         | :heavy_minus_sign:                                | Postal code for area                              |
| `State`                                           | `*string`                                         | :heavy_minus_sign:                                | Sub-unit of country (abbreviations ok)            |
| `Type`                                            | `*string`                                         | :heavy_minus_sign:                                | postal \| physical \| both                        |
| `Use`                                             | `*string`                                         | :heavy_minus_sign:                                | home \| work \| temp \| old \| billing            |