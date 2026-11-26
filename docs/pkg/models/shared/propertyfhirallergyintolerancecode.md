# PropertyFhirAllergyIntoleranceCode

Code that identifies the allergy or intolerance


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Coding`                                                              | [][shared.FhirCoding](../../../pkg/models/shared/fhircoding.md)       | :heavy_minus_sign:                                                    | Codes defined by a terminology system                                 |
| `Extension`                                                           | [][shared.FhirExtension](../../../pkg/models/shared/fhirextension.md) | :heavy_minus_sign:                                                    | Extensions, eg. data-absent-reason when a required concept is unknown |
| `Text`                                                                | `*string`                                                             | :heavy_minus_sign:                                                    | Plain text representation of the concept                              |