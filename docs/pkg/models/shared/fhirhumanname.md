# FhirHumanName


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Family`                                                      | `*string`                                                     | :heavy_minus_sign:                                            | Family name (often called surname)                            |
| `Given`                                                       | []`string`                                                    | :heavy_minus_sign:                                            | Given names (not always first), including middle names        |
| `Prefix`                                                      | []`string`                                                    | :heavy_minus_sign:                                            | Parts that come before the name, eg. Dr                       |
| `Text`                                                        | `*string`                                                     | :heavy_minus_sign:                                            | Text representation of the full name                          |
| `Use`                                                         | `*string`                                                     | :heavy_minus_sign:                                            | usual \| official \| temp \| nickname \| anonymous \| old \| maiden |