# PropertyFhirEncounterClass

Classification of patient encounter


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Code`                                                                | `*string`                                                             | :heavy_minus_sign:                                                    | Classification of patient encounter, eg. AMB, EMER, IMP, VR           |
| `Extension`                                                           | [][shared.FhirExtension](../../../pkg/models/shared/fhirextension.md) | :heavy_minus_sign:                                                    | data-absent-reason when the encounter class is unknown                |
| `System`                                                              | `*string`                                                             | :heavy_minus_sign:                                                    | Identity of the terminology system (v3-ActCode)                       |