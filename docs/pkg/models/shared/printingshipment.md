# PrintingShipment

A shipment of a print order


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `Carrier`                                  | `*string`                                  | :heavy_minus_sign:                         | Carrier name                               |
| `DeliveredAt`                              | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)   |
| `ID`                                       | `*string`                                  | :heavy_minus_sign:                         | Shipment identifier                        |
| `Service`                                  | `*string`                                  | :heavy_minus_sign:                         | Delivery service name                      |
| `ShippedAt`                                | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | (ISO-8601 / YYYY-MM-DDTHH:MM:SSZ format)   |
| `TrackingNumber`                           | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |
| `TrackingURL`                              | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |