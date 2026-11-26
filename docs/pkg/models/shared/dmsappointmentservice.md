# DmsAppointmentService

A service requested on an appointment


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `Complaint`                                              | `*string`                                                | :heavy_minus_sign:                                       | Customer's own words                                     |
| `EstimateAmount`                                         | `*float64`                                               | :heavy_minus_sign:                                       | Estimated price                                          |
| `EstimatedHours`                                         | `*float64`                                               | :heavy_minus_sign:                                       | N/A                                                      |
| `ID`                                                     | `*string`                                                | :heavy_minus_sign:                                       | Line identifier                                          |
| `Name`                                                   | `*string`                                                | :heavy_minus_sign:                                       | Service or operation description                         |
| `OpcodeID`                                               | `*string`                                                | :heavy_minus_sign:                                       | The DmsOpcode ID                                         |
| `PayType`                                                | [*shared.PayType](../../../pkg/models/shared/paytype.md) | :heavy_minus_sign:                                       | N/A                                                      |
| `TechnicianUserIds`                                      | []`string`                                               | :heavy_minus_sign:                                       | Pre-assigned technicians, as HR Employees                |