# LegalParty


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `CompanyID`                                                            | `*string`                                                              | :heavy_minus_sign:                                                     | pointer to a CRM Company                                               |
| `ContactID`                                                            | `*string`                                                              | :heavy_minus_sign:                                                     | pointer to a CRM Contact                                               |
| `Description`                                                          | `*string`                                                              | :heavy_minus_sign:                                                     | the vendor's free-text relationship description                        |
| `Role`                                                                 | [*shared.LegalPartyRole](../../../pkg/models/shared/legalpartyrole.md) | :heavy_minus_sign:                                                     | the party role on the matter                                           |