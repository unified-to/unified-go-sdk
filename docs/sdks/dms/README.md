# Dms

## Overview

### Available Operations

* [CreateDmsAppointment](#createdmsappointment) - Create an appointment
* [CreateDmsOpcode](#createdmsopcode) - Create an opcode
* [CreateDmsPart](#createdmspart) - Create a part
* [CreateDmsPartsale](#createdmspartsale) - Create a partsale
* [CreateDmsRepairorder](#createdmsrepairorder) - Create a repairorder
* [CreateDmsVehicle](#createdmsvehicle) - Create a vehicle
* [CreateDmsVehiclesale](#createdmsvehiclesale) - Create a vehiclesale
* [GetDmsAppointment](#getdmsappointment) - Retrieve an appointment
* [GetDmsOpcode](#getdmsopcode) - Retrieve an opcode
* [GetDmsPart](#getdmspart) - Retrieve a part
* [GetDmsPartsale](#getdmspartsale) - Retrieve a partsale
* [GetDmsRepairorder](#getdmsrepairorder) - Retrieve a repairorder
* [GetDmsVehicle](#getdmsvehicle) - Retrieve a vehicle
* [GetDmsVehiclesale](#getdmsvehiclesale) - Retrieve a vehiclesale
* [ListDmsAppointments](#listdmsappointments) - List all appointments
* [ListDmsOpcodes](#listdmsopcodes) - List all opcodes
* [ListDmsParts](#listdmsparts) - List all parts
* [ListDmsPartsales](#listdmspartsales) - List all partsales
* [ListDmsRepairorders](#listdmsrepairorders) - List all repairorders
* [ListDmsVehicles](#listdmsvehicles) - List all vehicles
* [ListDmsVehiclesales](#listdmsvehiclesales) - List all vehiclesales
* [PatchDmsAppointment](#patchdmsappointment) - Update an appointment
* [PatchDmsOpcode](#patchdmsopcode) - Update an opcode
* [PatchDmsPart](#patchdmspart) - Update a part
* [PatchDmsPartsale](#patchdmspartsale) - Update a partsale
* [PatchDmsRepairorder](#patchdmsrepairorder) - Update a repairorder
* [PatchDmsVehicle](#patchdmsvehicle) - Update a vehicle
* [PatchDmsVehiclesale](#patchdmsvehiclesale) - Update a vehiclesale
* [RemoveDmsAppointment](#removedmsappointment) - Remove an appointment
* [UpdateDmsAppointment](#updatedmsappointment) - Update an appointment
* [UpdateDmsOpcode](#updatedmsopcode) - Update an opcode
* [UpdateDmsPart](#updatedmspart) - Update a part
* [UpdateDmsPartsale](#updatedmspartsale) - Update a partsale
* [UpdateDmsRepairorder](#updatedmsrepairorder) - Update a repairorder
* [UpdateDmsVehicle](#updatedmsvehicle) - Update a vehicle
* [UpdateDmsVehiclesale](#updatedmsvehiclesale) - Update a vehiclesale

## CreateDmsAppointment

Creates a DmsAppointment. Provide the unified DmsAppointment fields in the request body; they are mapped to the provider's API. Returns the created DmsAppointment.

### Example Usage

<!-- UsageSnippet language="go" operationID="createDmsAppointment" method="post" path="/dms/{connection_id}/appointment" example="dms_appointment" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.CreateDmsAppointment(ctx, operations.CreateDmsAppointmentRequest{
        DmsAppointment: shared.DmsAppointment{
            CreatedAt: types.MustNewTimeFromString("2023-10-16T02:56:45.127Z"),
            EndAt: types.MustNewTimeFromString("2026-05-08T02:39:01.774Z"),
            ID: unifiedgosdk.Pointer("a45b899b-460f-4dd4-a8b9-c35116dc94e3"),
            Metadata: []shared.DmsMetadata{},
            Notes: unifiedgosdk.Pointer("Soluta conitor timidus confero."),
            Number: unifiedgosdk.Pointer("A541909"),
            Services: []shared.DmsAppointmentService{
                shared.DmsAppointmentService{
                    ID: unifiedgosdk.Pointer("2d588067-f689-484b-84fe-da3488c15dcc"),
                    Name: unifiedgosdk.Pointer("Customer concern"),
                    PayType: shared.PayTypeWarranty.ToPointer(),
                    TechnicianUserIds: []string{},
                },
                shared.DmsAppointmentService{
                    Complaint: unifiedgosdk.Pointer("Check engine light on"),
                    ID: unifiedgosdk.Pointer("f581c607-3f12-45d1-a7cf-8d5e988f484f"),
                    Name: unifiedgosdk.Pointer("Customer concern"),
                    PayType: shared.PayTypeCustomer.ToPointer(),
                    TechnicianUserIds: []string{},
                },
                shared.DmsAppointmentService{
                    Complaint: unifiedgosdk.Pointer("Due for service"),
                    ID: unifiedgosdk.Pointer("fc06cd83-06cb-47d2-81f2-598ce378b902"),
                    Name: unifiedgosdk.Pointer("Customer concern"),
                    PayType: shared.PayTypeCustomer.ToPointer(),
                    TechnicianUserIds: []string{},
                },
            },
            StartAt: types.MustNewTimeFromString("2026-05-07T19:39:01.774Z"),
            Status: shared.DmsAppointmentStatusCancelled.ToPointer(),
            TransportType: shared.TransportTypeLoaner.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-04-06T16:29:51.843Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsAppointment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateDmsAppointmentRequest](../../pkg/models/operations/createdmsappointmentrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.CreateDmsAppointmentResponse](../../pkg/models/operations/createdmsappointmentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateDmsOpcode

Creates a DmsOpcode. Provide the unified DmsOpcode fields in the request body; they are mapped to the provider's API. Returns the created DmsOpcode.

### Example Usage

<!-- UsageSnippet language="go" operationID="createDmsOpcode" method="post" path="/dms/{connection_id}/opcode" example="dms_opcode" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.CreateDmsOpcode(ctx, operations.CreateDmsOpcodeRequest{
        DmsOpcode: shared.DmsOpcode{
            Category: unifiedgosdk.Pointer("Electrical"),
            Code: unifiedgosdk.Pointer("02AUS"),
            CreatedAt: types.MustNewTimeFromString("2021-01-12T19:03:17.545Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Description: unifiedgosdk.Pointer("Oil and filter change"),
            EstimatedHours: unifiedgosdk.Pointer[float64](3.2),
            FlatRateHours: unifiedgosdk.Pointer[float64](3.5),
            ID: unifiedgosdk.Pointer("0a4d6e8d-e0f4-4f7c-8e7a-42fd56bcddfd"),
            IsActive: unifiedgosdk.Pointer(true),
            LaborRateAmount: unifiedgosdk.Pointer[float64](216.0),
            Metadata: []shared.DmsMetadata{},
            OemCode: unifiedgosdk.Pointer("05604A3"),
            Parts: []shared.DmsPartline{},
            PayType: shared.DmsOpcodePayTypeCustomer.ToPointer(),
            SkillCode: unifiedgosdk.Pointer("A"),
            Type: shared.DmsOpcodeTypeLabor.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2021-11-03T15:48:51.322Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsOpcode != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateDmsOpcodeRequest](../../pkg/models/operations/createdmsopcoderequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreateDmsOpcodeResponse](../../pkg/models/operations/createdmsopcoderesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateDmsPart

Creates a DmsPart. Provide the unified DmsPart fields in the request body; they are mapped to the provider's API. Returns the created DmsPart.

### Example Usage

<!-- UsageSnippet language="go" operationID="createDmsPart" method="post" path="/dms/{connection_id}/part" example="dms_part" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.CreateDmsPart(ctx, operations.CreateDmsPartRequest{
        DmsPart: shared.DmsPart{
            BinLocations: []string{
                "S18-6",
            },
            Category: unifiedgosdk.Pointer("Ignition"),
            CostAmount: unifiedgosdk.Pointer[float64](25.26),
            CreatedAt: types.MustNewTimeFromString("2020-07-02T08:18:03.858Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Description: unifiedgosdk.Pointer("Cabin filter"),
            ID: unifiedgosdk.Pointer("44573705-82b3-43dc-a21f-6a8d3e8fdce2"),
            LastSoldAt: types.MustNewTimeFromString("2020-11-23T04:33:10.426Z"),
            ListPriceAmount: unifiedgosdk.Pointer[float64](39.61),
            ManufacturerCode: unifiedgosdk.Pointer("TOY"),
            MaxStockQuantity: unifiedgosdk.Pointer[float64](21.0),
            Metadata: []shared.DmsMetadata{},
            MinStockQuantity: unifiedgosdk.Pointer[float64](3.0),
            PackQuantity: unifiedgosdk.Pointer[float64](1.0),
            PartNumber: unifiedgosdk.Pointer("73992-IM3FI"),
            QuantityAvailable: unifiedgosdk.Pointer[float64](13.0),
            QuantityOnHand: unifiedgosdk.Pointer[float64](15.0),
            QuantityOnOrder: unifiedgosdk.Pointer[float64](14.0),
            SourceCode: unifiedgosdk.Pointer("30"),
            TradePriceAmount: unifiedgosdk.Pointer[float64](29.05),
            UpdatedAt: types.MustNewTimeFromString("2023-11-23T07:05:30.247Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsPart != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateDmsPartRequest](../../pkg/models/operations/createdmspartrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreateDmsPartResponse](../../pkg/models/operations/createdmspartresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateDmsPartsale

Creates a DmsPartsale. Provide the unified DmsPartsale fields in the request body; they are mapped to the provider's API. Returns the created DmsPartsale.

### Example Usage

<!-- UsageSnippet language="go" operationID="createDmsPartsale" method="post" path="/dms/{connection_id}/partsale" example="dms_partsale" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.CreateDmsPartsale(ctx, operations.CreateDmsPartsaleRequest{
        DmsPartsale: shared.DmsPartsale{
            CostAmount: unifiedgosdk.Pointer[float64](0.0),
            CreatedAt: types.MustNewTimeFromString("2023-05-22T03:39:30.335Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            FreightAmount: unifiedgosdk.Pointer[float64](0.0),
            ID: unifiedgosdk.Pointer("bd08d54d-448b-4d62-8768-b6502973e230"),
            Lines: []shared.DmsPartline{
                shared.DmsPartline{
                    ID: unifiedgosdk.Pointer("caca2735-8164-423c-84f4-488aa8c75ac9"),
                    PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                    Quantity: unifiedgosdk.Pointer[float64](5.0),
                    QuantityOrdered: unifiedgosdk.Pointer[float64](5.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](491.45),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](98.29),
                },
                shared.DmsPartline{
                    ID: unifiedgosdk.Pointer("8e6e84b8-fb04-45e5-a51e-06a1cd4dc803"),
                    PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                    Quantity: unifiedgosdk.Pointer[float64](2.0),
                    QuantityOrdered: unifiedgosdk.Pointer[float64](2.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](240.9),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](120.45),
                },
                shared.DmsPartline{
                    ID: unifiedgosdk.Pointer("dc2e64c1-cc98-4b72-a4fa-f67dd4941ab9"),
                    PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                    Quantity: unifiedgosdk.Pointer[float64](3.0),
                    QuantityOrdered: unifiedgosdk.Pointer[float64](3.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](249.54),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](83.18),
                },
            },
            Metadata: []shared.DmsMetadata{},
            Number: unifiedgosdk.Pointer("P292867"),
            PayType: shared.DmsPartsalePayTypeCustomer.ToPointer(),
            Status: shared.DmsPartsaleStatusVoided.ToPointer(),
            SubtotalAmount: unifiedgosdk.Pointer[float64](981.89),
            TaxAmount: unifiedgosdk.Pointer[float64](68.73),
            TotalAmount: unifiedgosdk.Pointer[float64](1050.62),
            Type: shared.DmsPartsaleTypeSale.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-04-18T13:15:50.883Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsPartsale != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateDmsPartsaleRequest](../../pkg/models/operations/createdmspartsalerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateDmsPartsaleResponse](../../pkg/models/operations/createdmspartsaleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateDmsRepairorder

Creates a DmsRepairorder. Provide the unified DmsRepairorder fields in the request body; they are mapped to the provider's API. Returns the created DmsRepairorder.

### Example Usage

<!-- UsageSnippet language="go" operationID="createDmsRepairorder" method="post" path="/dms/{connection_id}/repairorder" example="dms_repairorder" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.CreateDmsRepairorder(ctx, operations.CreateDmsRepairorderRequest{
        DmsRepairorder: shared.DmsRepairorder{
            BalanceAmount: unifiedgosdk.Pointer[float64](3193.86),
            CompletedAt: types.MustNewTimeFromString("2026-11-13T20:43:09.427Z"),
            CreatedAt: types.MustNewTimeFromString("2022-10-25T15:50:48.007Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            EstimateAmount: unifiedgosdk.Pointer[float64](3264.0),
            ID: unifiedgosdk.Pointer("5b68208c-b532-4191-a308-b3aea18b349f"),
            Jobs: []shared.DmsRepairorderJob{
                shared.DmsRepairorderJob{
                    ActualHours: unifiedgosdk.Pointer[float64](3.6),
                    Cause: unifiedgosdk.Pointer("Scheduled maintenance"),
                    Complaint: unifiedgosdk.Pointer("Due for service"),
                    Correction: unifiedgosdk.Pointer("Aufero cumque vicinus aptus assumenda quis aptus."),
                    ID: unifiedgosdk.Pointer("61881098-ee1e-4f1e-bb79-518744ec672e"),
                    LaborAmount: unifiedgosdk.Pointer[float64](681.6),
                    Labors: []shared.DmsRepairorderLabor{
                        shared.DmsRepairorderLabor{
                            ID: unifiedgosdk.Pointer("6d675122-a8ed-4635-b907-760be1b4d601"),
                            PayType: shared.DmsRepairorderLaborPayTypeWarranty.ToPointer(),
                            RateAmount: unifiedgosdk.Pointer[float64](213.0),
                            SoldHours: unifiedgosdk.Pointer[float64](3.2),
                            TechnicianUserIds: []string{},
                            TotalAmount: unifiedgosdk.Pointer[float64](681.6),
                        },
                    },
                    Name: unifiedgosdk.Pointer("Customer concern"),
                    Number: unifiedgosdk.Pointer("A"),
                    Parts: []shared.DmsPartline{},
                    PartsAmount: unifiedgosdk.Pointer[float64](0.0),
                    PayType: shared.DmsRepairorderJobPayTypeWarranty.ToPointer(),
                    SoldHours: unifiedgosdk.Pointer[float64](3.2),
                    Status: shared.DmsRepairorderJobStatusCompleted.ToPointer(),
                    TechnicianUserIds: []string{},
                    TotalAmount: unifiedgosdk.Pointer[float64](681.6),
                },
                shared.DmsRepairorderJob{
                    ActualHours: unifiedgosdk.Pointer[float64](3.0),
                    Cause: unifiedgosdk.Pointer("Worn brake pads"),
                    Complaint: unifiedgosdk.Pointer("Vibration at highway speed"),
                    Correction: unifiedgosdk.Pointer("Somnus avaritia utroque vulnero ullam ancilla."),
                    ID: unifiedgosdk.Pointer("8aa00e52-e835-4147-b63e-f0f101d96411"),
                    LaborAmount: unifiedgosdk.Pointer[float64](446.4),
                    Labors: []shared.DmsRepairorderLabor{
                        shared.DmsRepairorderLabor{
                            ID: unifiedgosdk.Pointer("e1d2504c-84ea-4552-a854-300c19edaea6"),
                            PayType: shared.DmsRepairorderLaborPayTypeWarranty.ToPointer(),
                            RateAmount: unifiedgosdk.Pointer[float64](144.0),
                            SoldHours: unifiedgosdk.Pointer[float64](3.1),
                            TechnicianUserIds: []string{},
                            TotalAmount: unifiedgosdk.Pointer[float64](446.4),
                        },
                    },
                    Name: unifiedgosdk.Pointer("Customer concern"),
                    Number: unifiedgosdk.Pointer("B"),
                    Parts: []shared.DmsPartline{
                        shared.DmsPartline{
                            ID: unifiedgosdk.Pointer("bd071217-0d8d-4a6f-a6af-d7ba827d186e"),
                            PayType: shared.DmsPartlinePayTypeWarranty.ToPointer(),
                            Quantity: unifiedgosdk.Pointer[float64](4.0),
                            QuantityOrdered: unifiedgosdk.Pointer[float64](4.0),
                            TotalAmount: unifiedgosdk.Pointer[float64](890.44),
                            UnitPriceAmount: unifiedgosdk.Pointer[float64](222.61),
                        },
                        shared.DmsPartline{
                            ID: unifiedgosdk.Pointer("bfa84261-d5a5-4311-a017-130b324efef6"),
                            PayType: shared.DmsPartlinePayTypeWarranty.ToPointer(),
                            Quantity: unifiedgosdk.Pointer[float64](2.0),
                            QuantityOrdered: unifiedgosdk.Pointer[float64](2.0),
                            TotalAmount: unifiedgosdk.Pointer[float64](216.18),
                            UnitPriceAmount: unifiedgosdk.Pointer[float64](108.09),
                        },
                    },
                    PartsAmount: unifiedgosdk.Pointer[float64](1106.6200000000001),
                    PayType: shared.DmsRepairorderJobPayTypeWarranty.ToPointer(),
                    SoldHours: unifiedgosdk.Pointer[float64](3.1),
                    Status: shared.DmsRepairorderJobStatusCompleted.ToPointer(),
                    TechnicianUserIds: []string{},
                    TotalAmount: unifiedgosdk.Pointer[float64](1553.02),
                },
                shared.DmsRepairorderJob{
                    ActualHours: unifiedgosdk.Pointer[float64](1.5),
                    Cause: unifiedgosdk.Pointer("Scheduled maintenance"),
                    Complaint: unifiedgosdk.Pointer("Noise when braking"),
                    Correction: unifiedgosdk.Pointer("Degero theca natus."),
                    ID: unifiedgosdk.Pointer("6a2a7fef-cf18-42b3-804a-3ef80fc119c8"),
                    LaborAmount: unifiedgosdk.Pointer[float64](232.4),
                    Labors: []shared.DmsRepairorderLabor{
                        shared.DmsRepairorderLabor{
                            ID: unifiedgosdk.Pointer("13bb1d26-07df-4f63-b39f-42d39e786858"),
                            PayType: shared.DmsRepairorderLaborPayTypeCustomer.ToPointer(),
                            RateAmount: unifiedgosdk.Pointer[float64](166.0),
                            SoldHours: unifiedgosdk.Pointer[float64](1.4),
                            TechnicianUserIds: []string{},
                            TotalAmount: unifiedgosdk.Pointer[float64](232.4),
                        },
                    },
                    Name: unifiedgosdk.Pointer("Customer concern"),
                    Number: unifiedgosdk.Pointer("C"),
                    Parts: []shared.DmsPartline{
                        shared.DmsPartline{
                            ID: unifiedgosdk.Pointer("3d2f54ed-394f-47c9-9ffd-8ebc6a9db242"),
                            PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                            Quantity: unifiedgosdk.Pointer[float64](1.0),
                            QuantityOrdered: unifiedgosdk.Pointer[float64](1.0),
                            TotalAmount: unifiedgosdk.Pointer[float64](216.02),
                            UnitPriceAmount: unifiedgosdk.Pointer[float64](216.02),
                        },
                        shared.DmsPartline{
                            ID: unifiedgosdk.Pointer("3440a836-55c5-4033-b4af-1c3a9a9cdd6a"),
                            PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                            Quantity: unifiedgosdk.Pointer[float64](2.0),
                            QuantityOrdered: unifiedgosdk.Pointer[float64](2.0),
                            TotalAmount: unifiedgosdk.Pointer[float64](358.16),
                            UnitPriceAmount: unifiedgosdk.Pointer[float64](179.08),
                        },
                    },
                    PartsAmount: unifiedgosdk.Pointer[float64](574.1800000000001),
                    PayType: shared.DmsRepairorderJobPayTypeCustomer.ToPointer(),
                    SoldHours: unifiedgosdk.Pointer[float64](1.4),
                    Status: shared.DmsRepairorderJobStatusCompleted.ToPointer(),
                    TechnicianUserIds: []string{},
                    TotalAmount: unifiedgosdk.Pointer[float64](806.58),
                },
            },
            LaborAmount: unifiedgosdk.Pointer[float64](1360.4),
            Metadata: []shared.DmsMetadata{},
            MiscAmount: unifiedgosdk.Pointer[float64](35.0),
            Number: unifiedgosdk.Pointer("143398"),
            OdometerIn: unifiedgosdk.Pointer[float64](7776.0),
            OdometerOut: unifiedgosdk.Pointer[float64](7776.0),
            OdometerUnit: shared.DmsRepairorderOdometerUnitMiles.ToPointer(),
            OpenedAt: types.MustNewTimeFromString("2025-02-08T09:35:35.439Z"),
            PartsAmount: unifiedgosdk.Pointer[float64](1680.8000000000002),
            PayType: shared.DmsRepairorderPayTypeWarranty.ToPointer(),
            PromisedAt: types.MustNewTimeFromString("2025-02-09T17:35:35.439Z"),
            PurchaseOrderNumber: unifiedgosdk.Pointer("182743"),
            Status: shared.DmsRepairorderStatusCompleted.ToPointer(),
            SubletAmount: unifiedgosdk.Pointer[float64](0.0),
            TagNumber: unifiedgosdk.Pointer("758"),
            TaxAmount: unifiedgosdk.Pointer[float64](117.66),
            TechnicianUserIds: []string{},
            TotalAmount: unifiedgosdk.Pointer[float64](3193.86),
            Totals: []shared.DmsRepairorderTotal{
                shared.DmsRepairorderTotal{
                    DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
                    LaborAmount: unifiedgosdk.Pointer[float64](232.4),
                    PartsAmount: unifiedgosdk.Pointer[float64](574.1800000000001),
                    PayType: shared.DmsRepairorderTotalPayTypeCustomer.ToPointer(),
                    SubletAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](806.58),
                },
                shared.DmsRepairorderTotal{
                    DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
                    LaborAmount: unifiedgosdk.Pointer[float64](1128.0),
                    PartsAmount: unifiedgosdk.Pointer[float64](1106.6200000000001),
                    PayType: shared.DmsRepairorderTotalPayTypeWarranty.ToPointer(),
                    SubletAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](2234.62),
                },
            },
            TransportType: shared.DmsRepairorderTransportTypeLoaner.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-11-28T17:30:18.674Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsRepairorder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateDmsRepairorderRequest](../../pkg/models/operations/createdmsrepairorderrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.CreateDmsRepairorderResponse](../../pkg/models/operations/createdmsrepairorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateDmsVehicle

Creates a DmsVehicle. Provide the unified DmsVehicle fields in the request body; they are mapped to the provider's API. Returns the created DmsVehicle.

### Example Usage

<!-- UsageSnippet language="go" operationID="createDmsVehicle" method="post" path="/dms/{connection_id}/vehicle" example="dms_vehicle" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.CreateDmsVehicle(ctx, operations.CreateDmsVehicleRequest{
        DmsVehicle: shared.DmsVehicle{
            BodyStyle: unifiedgosdk.Pointer("Pickup"),
            Condition: shared.ConditionNew.ToPointer(),
            CostAmount: unifiedgosdk.Pointer[float64](26498.0),
            CreatedAt: types.MustNewTimeFromString("2021-02-03T08:01:58.559Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Description: unifiedgosdk.Pointer("2023 Citroën Silverado"),
            ExteriorColor: unifiedgosdk.Pointer("lavender"),
            FuelType: shared.FuelTypeElectric.ToPointer(),
            ID: unifiedgosdk.Pointer("55b78751-25a9-4229-b88f-86b3ea022bf0"),
            ImageUrls: []string{
                "https://picsum.photos/seed/tOUr5K/1358/1877",
            },
            InteriorColor: unifiedgosdk.Pointer("Black"),
            InternetPriceAmount: unifiedgosdk.Pointer[float64](29526.0),
            IsActive: unifiedgosdk.Pointer(true),
            IsCertified: unifiedgosdk.Pointer(false),
            IsInventory: unifiedgosdk.Pointer(true),
            IsPublished: unifiedgosdk.Pointer(false),
            Make: unifiedgosdk.Pointer("Citroën"),
            Metadata: []shared.DmsMetadata{},
            Model: unifiedgosdk.Pointer("Silverado"),
            ModelCode: unifiedgosdk.Pointer("PMU8KI"),
            MsrpAmount: unifiedgosdk.Pointer[float64](33100.0),
            Odometer: unifiedgosdk.Pointer[float64](1.0),
            OdometerUnit: shared.DmsVehicleOdometerUnitMiles.ToPointer(),
            Options: []shared.DmsVehicleOption{
                shared.DmsVehicleOption{
                    Code: unifiedgosdk.Pointer("SXE"),
                    CostAmount: unifiedgosdk.Pointer[float64](250.0),
                    Description: unifiedgosdk.Pointer("Tow package"),
                    PriceAmount: unifiedgosdk.Pointer[float64](1400.0),
                    Type: shared.DmsVehicleOptionTypeDealer.ToPointer(),
                },
                shared.DmsVehicleOption{
                    Code: unifiedgosdk.Pointer("458"),
                    CostAmount: unifiedgosdk.Pointer[float64](700.0),
                    Description: unifiedgosdk.Pointer("Tow package"),
                    PriceAmount: unifiedgosdk.Pointer[float64](600.0),
                    Type: shared.DmsVehicleOptionTypeDealer.ToPointer(),
                },
            },
            PriceAmount: unifiedgosdk.Pointer[float64](29926.0),
            ReceivedAt: types.MustNewTimeFromString("2021-02-03T08:01:58.559Z"),
            Status: shared.DmsVehicleStatusOther.ToPointer(),
            StockNumber: unifiedgosdk.Pointer("N24959"),
            Transmission: unifiedgosdk.Pointer("Manual"),
            Trim: unifiedgosdk.Pointer("Base"),
            Type: shared.DmsVehicleTypeSuv.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-11-03T13:50:06.421Z"),
            Vin: unifiedgosdk.Pointer("C29ZDU474DDR70955"),
            WarrantyExpiresAt: types.MustNewTimeFromString("2022-03-28T09:10:09.889Z"),
            Year: unifiedgosdk.Pointer[float64](2023.0),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsVehicle != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateDmsVehicleRequest](../../pkg/models/operations/createdmsvehiclerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateDmsVehicleResponse](../../pkg/models/operations/createdmsvehicleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateDmsVehiclesale

Creates a DmsVehiclesale. Provide the unified DmsVehiclesale fields in the request body; they are mapped to the provider's API. Returns the created DmsVehiclesale.

### Example Usage

<!-- UsageSnippet language="go" operationID="createDmsVehiclesale" method="post" path="/dms/{connection_id}/vehiclesale" example="dms_vehiclesale" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.CreateDmsVehiclesale(ctx, operations.CreateDmsVehiclesaleRequest{
        DmsVehiclesale: shared.DmsVehiclesale{
            BackGrossAmount: unifiedgosdk.Pointer[float64](4565.0),
            CreatedAt: types.MustNewTimeFromString("2020-04-02T08:14:56.150Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DownPaymentAmount: unifiedgosdk.Pointer[float64](6700.0),
            Fees: []shared.DmsVehiclesaleFee{
                shared.DmsVehiclesaleFee{
                    CostAmount: unifiedgosdk.Pointer[float64](0.0),
                    IsCapitalized: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Documentation fee"),
                    PriceAmount: unifiedgosdk.Pointer[float64](499.0),
                },
                shared.DmsVehiclesaleFee{
                    IsCapitalized: unifiedgosdk.Pointer(false),
                    Name: unifiedgosdk.Pointer("Title and registration"),
                    PriceAmount: unifiedgosdk.Pointer[float64](392.0),
                },
            },
            Finance: &shared.PropertyDmsVehiclesaleFinance{
                FinancedAmount: unifiedgosdk.Pointer[float64](0.0),
                FirstPaymentAt: types.MustNewTimeFromString("2022-08-16T11:36:15.388Z"),
                InterestRate: unifiedgosdk.Pointer[float64](6.68),
                LenderName: unifiedgosdk.Pointer("Mueller, Yundt and Murray Financial"),
                PaymentAmount: unifiedgosdk.Pointer[float64](339.0),
                PaymentFrequency: shared.PaymentFrequencyMonthly.ToPointer(),
                TermMonths: unifiedgosdk.Pointer[float64](72.0),
            },
            FinanceType: shared.FinanceTypeFinance.ToPointer(),
            ID: unifiedgosdk.Pointer("4309118f-90dd-4a07-94f2-ca8368edb908"),
            Metadata: []shared.DmsMetadata{},
            Number: unifiedgosdk.Pointer("D84591"),
            Products: []shared.DmsVehiclesaleProduct{
                shared.DmsVehiclesaleProduct{
                    ContractNumber: unifiedgosdk.Pointer("B7JN10GVB7"),
                    CostAmount: unifiedgosdk.Pointer[float64](1320.0),
                    Name: unifiedgosdk.Pointer("Tire & Wheel"),
                    PriceAmount: unifiedgosdk.Pointer[float64](3300.0),
                    Provider: unifiedgosdk.Pointer("Renner and Sons"),
                    TermMonths: unifiedgosdk.Pointer[float64](60.0),
                    Type: shared.DmsVehiclesaleProductTypeProtection.ToPointer(),
                },
                shared.DmsVehiclesaleProduct{
                    ContractNumber: unifiedgosdk.Pointer("VIZU7K3VRU"),
                    CostAmount: unifiedgosdk.Pointer[float64](1115.0),
                    DeductibleAmount: unifiedgosdk.Pointer[float64](100.0),
                    Name: unifiedgosdk.Pointer("Platinum Vehicle Service Contract"),
                    PriceAmount: unifiedgosdk.Pointer[float64](3700.0),
                    Provider: unifiedgosdk.Pointer("Frami, Wisozk and Gleichner"),
                    TermMonths: unifiedgosdk.Pointer[float64](72.0),
                    Type: shared.DmsVehiclesaleProductTypeServiceContract.ToPointer(),
                },
            },
            RebateAmount: unifiedgosdk.Pointer[float64](500.0),
            SoldAt: types.MustNewTimeFromString("2022-08-14T20:20:41.101Z"),
            Source: unifiedgosdk.Pointer("Website"),
            Status: shared.DmsVehiclesaleStatusSold.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            Taxes: []shared.DmsVehiclesaleTax{
                shared.DmsVehiclesaleTax{
                    Name: unifiedgosdk.Pointer("State sales tax"),
                    Rate: unifiedgosdk.Pointer[float64](8.25),
                    TotalAmount: unifiedgosdk.Pointer[float64](0.0),
                },
            },
            TotalAmount: unifiedgosdk.Pointer[float64](7891.0),
            TradeIns: []shared.DmsVehiclesaleTradein{},
            Type: shared.DmsVehiclesaleTypeRetail.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-04-16T21:50:54.863Z"),
            Vehicles: []shared.DmsVehiclesaleVehicle{},
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsVehiclesale != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateDmsVehiclesaleRequest](../../pkg/models/operations/createdmsvehiclesalerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.CreateDmsVehiclesaleResponse](../../pkg/models/operations/createdmsvehiclesaleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetDmsAppointment

Retrieves a single DmsAppointment by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getDmsAppointment" method="get" path="/dms/{connection_id}/appointment/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.GetDmsAppointment(ctx, operations.GetDmsAppointmentRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsAppointment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetDmsAppointmentRequest](../../pkg/models/operations/getdmsappointmentrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetDmsAppointmentResponse](../../pkg/models/operations/getdmsappointmentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetDmsOpcode

Retrieves a single DmsOpcode by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getDmsOpcode" method="get" path="/dms/{connection_id}/opcode/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.GetDmsOpcode(ctx, operations.GetDmsOpcodeRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsOpcode != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetDmsOpcodeRequest](../../pkg/models/operations/getdmsopcoderequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.GetDmsOpcodeResponse](../../pkg/models/operations/getdmsopcoderesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetDmsPart

Retrieves a single DmsPart by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getDmsPart" method="get" path="/dms/{connection_id}/part/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.GetDmsPart(ctx, operations.GetDmsPartRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsPart != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetDmsPartRequest](../../pkg/models/operations/getdmspartrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.GetDmsPartResponse](../../pkg/models/operations/getdmspartresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetDmsPartsale

Retrieves a single DmsPartsale by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getDmsPartsale" method="get" path="/dms/{connection_id}/partsale/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.GetDmsPartsale(ctx, operations.GetDmsPartsaleRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsPartsale != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetDmsPartsaleRequest](../../pkg/models/operations/getdmspartsalerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetDmsPartsaleResponse](../../pkg/models/operations/getdmspartsaleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetDmsRepairorder

Retrieves a single DmsRepairorder by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getDmsRepairorder" method="get" path="/dms/{connection_id}/repairorder/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.GetDmsRepairorder(ctx, operations.GetDmsRepairorderRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsRepairorder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetDmsRepairorderRequest](../../pkg/models/operations/getdmsrepairorderrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetDmsRepairorderResponse](../../pkg/models/operations/getdmsrepairorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetDmsVehicle

Retrieves a single DmsVehicle by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getDmsVehicle" method="get" path="/dms/{connection_id}/vehicle/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.GetDmsVehicle(ctx, operations.GetDmsVehicleRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsVehicle != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetDmsVehicleRequest](../../pkg/models/operations/getdmsvehiclerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetDmsVehicleResponse](../../pkg/models/operations/getdmsvehicleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetDmsVehiclesale

Retrieves a single DmsVehiclesale by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getDmsVehiclesale" method="get" path="/dms/{connection_id}/vehiclesale/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.GetDmsVehiclesale(ctx, operations.GetDmsVehiclesaleRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsVehiclesale != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetDmsVehiclesaleRequest](../../pkg/models/operations/getdmsvehiclesalerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetDmsVehiclesaleResponse](../../pkg/models/operations/getdmsvehiclesaleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListDmsAppointments

Returns a list of DmsAppointments. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listDmsAppointments" method="get" path="/dms/{connection_id}/appointment" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.ListDmsAppointments(ctx, operations.ListDmsAppointmentsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsAppointments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListDmsAppointmentsRequest](../../pkg/models/operations/listdmsappointmentsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListDmsAppointmentsResponse](../../pkg/models/operations/listdmsappointmentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListDmsOpcodes

Returns a list of DmsOpcodes. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listDmsOpcodes" method="get" path="/dms/{connection_id}/opcode" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.ListDmsOpcodes(ctx, operations.ListDmsOpcodesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsOpcodes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListDmsOpcodesRequest](../../pkg/models/operations/listdmsopcodesrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ListDmsOpcodesResponse](../../pkg/models/operations/listdmsopcodesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListDmsParts

Returns a list of DmsParts. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listDmsParts" method="get" path="/dms/{connection_id}/part" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.ListDmsParts(ctx, operations.ListDmsPartsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsParts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListDmsPartsRequest](../../pkg/models/operations/listdmspartsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListDmsPartsResponse](../../pkg/models/operations/listdmspartsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListDmsPartsales

Returns a list of DmsPartsales. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listDmsPartsales" method="get" path="/dms/{connection_id}/partsale" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.ListDmsPartsales(ctx, operations.ListDmsPartsalesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsPartsales != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListDmsPartsalesRequest](../../pkg/models/operations/listdmspartsalesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListDmsPartsalesResponse](../../pkg/models/operations/listdmspartsalesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListDmsRepairorders

Returns a list of DmsRepairorders. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listDmsRepairorders" method="get" path="/dms/{connection_id}/repairorder" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.ListDmsRepairorders(ctx, operations.ListDmsRepairordersRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsRepairorders != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListDmsRepairordersRequest](../../pkg/models/operations/listdmsrepairordersrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListDmsRepairordersResponse](../../pkg/models/operations/listdmsrepairordersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListDmsVehicles

Returns a list of DmsVehicles. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listDmsVehicles" method="get" path="/dms/{connection_id}/vehicle" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.ListDmsVehicles(ctx, operations.ListDmsVehiclesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsVehicles != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListDmsVehiclesRequest](../../pkg/models/operations/listdmsvehiclesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListDmsVehiclesResponse](../../pkg/models/operations/listdmsvehiclesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListDmsVehiclesales

Returns a list of DmsVehiclesales. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listDmsVehiclesales" method="get" path="/dms/{connection_id}/vehiclesale" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.ListDmsVehiclesales(ctx, operations.ListDmsVehiclesalesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsVehiclesales != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListDmsVehiclesalesRequest](../../pkg/models/operations/listdmsvehiclesalesrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListDmsVehiclesalesResponse](../../pkg/models/operations/listdmsvehiclesalesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchDmsAppointment

Updates the DmsAppointment identified by `id`. Provide the unified DmsAppointment fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchDmsAppointment" method="patch" path="/dms/{connection_id}/appointment/{id}" example="dms_appointment" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.PatchDmsAppointment(ctx, operations.PatchDmsAppointmentRequest{
        DmsAppointment: shared.DmsAppointment{
            CreatedAt: types.MustNewTimeFromString("2023-10-16T02:56:45.127Z"),
            EndAt: types.MustNewTimeFromString("2026-05-08T02:39:01.800Z"),
            ID: unifiedgosdk.Pointer("c0e5894c-5d3a-4c4a-90a3-f8511fcaf3e2"),
            Metadata: []shared.DmsMetadata{},
            Notes: unifiedgosdk.Pointer("Soluta conitor timidus confero."),
            Number: unifiedgosdk.Pointer("A541909"),
            Services: []shared.DmsAppointmentService{
                shared.DmsAppointmentService{
                    ID: unifiedgosdk.Pointer("f74a8fff-8526-4417-b048-4c972bd4b31b"),
                    Name: unifiedgosdk.Pointer("Customer concern"),
                    PayType: shared.PayTypeWarranty.ToPointer(),
                    TechnicianUserIds: []string{},
                },
                shared.DmsAppointmentService{
                    Complaint: unifiedgosdk.Pointer("Check engine light on"),
                    ID: unifiedgosdk.Pointer("90c323d7-deaa-447d-a591-e1ce6f4d9803"),
                    Name: unifiedgosdk.Pointer("Customer concern"),
                    PayType: shared.PayTypeCustomer.ToPointer(),
                    TechnicianUserIds: []string{},
                },
                shared.DmsAppointmentService{
                    Complaint: unifiedgosdk.Pointer("Due for service"),
                    ID: unifiedgosdk.Pointer("e73e1b90-395a-4afb-8ca7-1b108fb48bc3"),
                    Name: unifiedgosdk.Pointer("Customer concern"),
                    PayType: shared.PayTypeCustomer.ToPointer(),
                    TechnicianUserIds: []string{},
                },
            },
            StartAt: types.MustNewTimeFromString("2026-05-07T19:39:01.800Z"),
            Status: shared.DmsAppointmentStatusCancelled.ToPointer(),
            TransportType: shared.TransportTypeLoaner.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-04-06T16:29:51.858Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsAppointment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PatchDmsAppointmentRequest](../../pkg/models/operations/patchdmsappointmentrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.PatchDmsAppointmentResponse](../../pkg/models/operations/patchdmsappointmentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchDmsOpcode

Updates the DmsOpcode identified by `id`. Provide the unified DmsOpcode fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchDmsOpcode" method="patch" path="/dms/{connection_id}/opcode/{id}" example="dms_opcode" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.PatchDmsOpcode(ctx, operations.PatchDmsOpcodeRequest{
        DmsOpcode: shared.DmsOpcode{
            Category: unifiedgosdk.Pointer("Electrical"),
            Code: unifiedgosdk.Pointer("02AUS"),
            CreatedAt: types.MustNewTimeFromString("2021-01-12T19:03:17.545Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Description: unifiedgosdk.Pointer("Oil and filter change"),
            EstimatedHours: unifiedgosdk.Pointer[float64](3.2),
            FlatRateHours: unifiedgosdk.Pointer[float64](3.5),
            ID: unifiedgosdk.Pointer("95bbbacf-b8ea-4841-8a1a-03f73f25cbda"),
            IsActive: unifiedgosdk.Pointer(true),
            LaborRateAmount: unifiedgosdk.Pointer[float64](216.0),
            Metadata: []shared.DmsMetadata{},
            OemCode: unifiedgosdk.Pointer("05604A3"),
            Parts: []shared.DmsPartline{},
            PayType: shared.DmsOpcodePayTypeCustomer.ToPointer(),
            SkillCode: unifiedgosdk.Pointer("A"),
            Type: shared.DmsOpcodeTypeLabor.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2021-11-03T15:48:51.325Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsOpcode != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.PatchDmsOpcodeRequest](../../pkg/models/operations/patchdmsopcoderequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.PatchDmsOpcodeResponse](../../pkg/models/operations/patchdmsopcoderesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchDmsPart

Updates the DmsPart identified by `id`. Provide the unified DmsPart fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchDmsPart" method="patch" path="/dms/{connection_id}/part/{id}" example="dms_part" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.PatchDmsPart(ctx, operations.PatchDmsPartRequest{
        DmsPart: shared.DmsPart{
            BinLocations: []string{
                "S18-6",
            },
            Category: unifiedgosdk.Pointer("Ignition"),
            CostAmount: unifiedgosdk.Pointer[float64](25.26),
            CreatedAt: types.MustNewTimeFromString("2020-07-02T08:18:03.858Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Description: unifiedgosdk.Pointer("Cabin filter"),
            ID: unifiedgosdk.Pointer("44358212-7c73-4818-bd09-04036b6311db"),
            LastSoldAt: types.MustNewTimeFromString("2020-11-23T04:33:10.427Z"),
            ListPriceAmount: unifiedgosdk.Pointer[float64](39.61),
            ManufacturerCode: unifiedgosdk.Pointer("TOY"),
            MaxStockQuantity: unifiedgosdk.Pointer[float64](21.0),
            Metadata: []shared.DmsMetadata{},
            MinStockQuantity: unifiedgosdk.Pointer[float64](3.0),
            PackQuantity: unifiedgosdk.Pointer[float64](1.0),
            PartNumber: unifiedgosdk.Pointer("73992-IM3FI"),
            QuantityAvailable: unifiedgosdk.Pointer[float64](13.0),
            QuantityOnHand: unifiedgosdk.Pointer[float64](15.0),
            QuantityOnOrder: unifiedgosdk.Pointer[float64](14.0),
            SourceCode: unifiedgosdk.Pointer("30"),
            TradePriceAmount: unifiedgosdk.Pointer[float64](29.05),
            UpdatedAt: types.MustNewTimeFromString("2023-11-23T07:05:30.256Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsPart != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.PatchDmsPartRequest](../../pkg/models/operations/patchdmspartrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.PatchDmsPartResponse](../../pkg/models/operations/patchdmspartresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchDmsPartsale

Updates the DmsPartsale identified by `id`. Provide the unified DmsPartsale fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchDmsPartsale" method="patch" path="/dms/{connection_id}/partsale/{id}" example="dms_partsale" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.PatchDmsPartsale(ctx, operations.PatchDmsPartsaleRequest{
        DmsPartsale: shared.DmsPartsale{
            CostAmount: unifiedgosdk.Pointer[float64](0.0),
            CreatedAt: types.MustNewTimeFromString("2023-05-22T03:39:30.335Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            FreightAmount: unifiedgosdk.Pointer[float64](0.0),
            ID: unifiedgosdk.Pointer("c8f3d6d4-f9f0-4bf3-9721-4d7e489e98ba"),
            Lines: []shared.DmsPartline{
                shared.DmsPartline{
                    ID: unifiedgosdk.Pointer("13b3eda3-f590-4de6-a13d-499d910cfc23"),
                    PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                    Quantity: unifiedgosdk.Pointer[float64](5.0),
                    QuantityOrdered: unifiedgosdk.Pointer[float64](5.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](491.45),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](98.29),
                },
                shared.DmsPartline{
                    ID: unifiedgosdk.Pointer("939aa1f5-3dfd-4b82-838b-240da4c49ea6"),
                    PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                    Quantity: unifiedgosdk.Pointer[float64](2.0),
                    QuantityOrdered: unifiedgosdk.Pointer[float64](2.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](240.9),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](120.45),
                },
                shared.DmsPartline{
                    ID: unifiedgosdk.Pointer("f65e363c-3828-4710-909c-2b319567d354"),
                    PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                    Quantity: unifiedgosdk.Pointer[float64](3.0),
                    QuantityOrdered: unifiedgosdk.Pointer[float64](3.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](249.54),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](83.18),
                },
            },
            Metadata: []shared.DmsMetadata{},
            Number: unifiedgosdk.Pointer("P292867"),
            PayType: shared.DmsPartsalePayTypeCustomer.ToPointer(),
            Status: shared.DmsPartsaleStatusVoided.ToPointer(),
            SubtotalAmount: unifiedgosdk.Pointer[float64](981.89),
            TaxAmount: unifiedgosdk.Pointer[float64](68.73),
            TotalAmount: unifiedgosdk.Pointer[float64](1050.62),
            Type: shared.DmsPartsaleTypeSale.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-04-18T13:15:50.920Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsPartsale != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchDmsPartsaleRequest](../../pkg/models/operations/patchdmspartsalerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchDmsPartsaleResponse](../../pkg/models/operations/patchdmspartsaleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchDmsRepairorder

Updates the DmsRepairorder identified by `id`. Provide the unified DmsRepairorder fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchDmsRepairorder" method="patch" path="/dms/{connection_id}/repairorder/{id}" example="dms_repairorder" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.PatchDmsRepairorder(ctx, operations.PatchDmsRepairorderRequest{
        DmsRepairorder: shared.DmsRepairorder{
            BalanceAmount: unifiedgosdk.Pointer[float64](3193.86),
            CompletedAt: types.MustNewTimeFromString("2026-11-13T20:43:09.493Z"),
            CreatedAt: types.MustNewTimeFromString("2022-10-25T15:50:48.007Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            EstimateAmount: unifiedgosdk.Pointer[float64](3264.0),
            ID: unifiedgosdk.Pointer("0187b8d9-283c-4987-a821-3a99a34b88cf"),
            Jobs: []shared.DmsRepairorderJob{
                shared.DmsRepairorderJob{
                    ActualHours: unifiedgosdk.Pointer[float64](3.6),
                    Cause: unifiedgosdk.Pointer("Scheduled maintenance"),
                    Complaint: unifiedgosdk.Pointer("Due for service"),
                    Correction: unifiedgosdk.Pointer("Aufero cumque vicinus aptus assumenda quis aptus."),
                    ID: unifiedgosdk.Pointer("03beb52d-f18e-42ce-a286-39ca1d7d2652"),
                    LaborAmount: unifiedgosdk.Pointer[float64](681.6),
                    Labors: []shared.DmsRepairorderLabor{
                        shared.DmsRepairorderLabor{
                            ID: unifiedgosdk.Pointer("5ebedfb0-e6f8-4d3d-a820-632a6edccb14"),
                            PayType: shared.DmsRepairorderLaborPayTypeWarranty.ToPointer(),
                            RateAmount: unifiedgosdk.Pointer[float64](213.0),
                            SoldHours: unifiedgosdk.Pointer[float64](3.2),
                            TechnicianUserIds: []string{},
                            TotalAmount: unifiedgosdk.Pointer[float64](681.6),
                        },
                    },
                    Name: unifiedgosdk.Pointer("Customer concern"),
                    Number: unifiedgosdk.Pointer("A"),
                    Parts: []shared.DmsPartline{},
                    PartsAmount: unifiedgosdk.Pointer[float64](0.0),
                    PayType: shared.DmsRepairorderJobPayTypeWarranty.ToPointer(),
                    SoldHours: unifiedgosdk.Pointer[float64](3.2),
                    Status: shared.DmsRepairorderJobStatusCompleted.ToPointer(),
                    TechnicianUserIds: []string{},
                    TotalAmount: unifiedgosdk.Pointer[float64](681.6),
                },
                shared.DmsRepairorderJob{
                    ActualHours: unifiedgosdk.Pointer[float64](3.0),
                    Cause: unifiedgosdk.Pointer("Worn brake pads"),
                    Complaint: unifiedgosdk.Pointer("Vibration at highway speed"),
                    Correction: unifiedgosdk.Pointer("Somnus avaritia utroque vulnero ullam ancilla."),
                    ID: unifiedgosdk.Pointer("f2f8379e-6a83-4a06-aa7a-d6b1e9e53763"),
                    LaborAmount: unifiedgosdk.Pointer[float64](446.4),
                    Labors: []shared.DmsRepairorderLabor{
                        shared.DmsRepairorderLabor{
                            ID: unifiedgosdk.Pointer("e1c638b6-da08-4e1f-aad7-59db110190fc"),
                            PayType: shared.DmsRepairorderLaborPayTypeWarranty.ToPointer(),
                            RateAmount: unifiedgosdk.Pointer[float64](144.0),
                            SoldHours: unifiedgosdk.Pointer[float64](3.1),
                            TechnicianUserIds: []string{},
                            TotalAmount: unifiedgosdk.Pointer[float64](446.4),
                        },
                    },
                    Name: unifiedgosdk.Pointer("Customer concern"),
                    Number: unifiedgosdk.Pointer("B"),
                    Parts: []shared.DmsPartline{
                        shared.DmsPartline{
                            ID: unifiedgosdk.Pointer("d246b71c-d4a0-443f-bbde-7382f2328ffa"),
                            PayType: shared.DmsPartlinePayTypeWarranty.ToPointer(),
                            Quantity: unifiedgosdk.Pointer[float64](4.0),
                            QuantityOrdered: unifiedgosdk.Pointer[float64](4.0),
                            TotalAmount: unifiedgosdk.Pointer[float64](890.44),
                            UnitPriceAmount: unifiedgosdk.Pointer[float64](222.61),
                        },
                        shared.DmsPartline{
                            ID: unifiedgosdk.Pointer("7f781967-2d75-49f8-b0f5-10cc295c186e"),
                            PayType: shared.DmsPartlinePayTypeWarranty.ToPointer(),
                            Quantity: unifiedgosdk.Pointer[float64](2.0),
                            QuantityOrdered: unifiedgosdk.Pointer[float64](2.0),
                            TotalAmount: unifiedgosdk.Pointer[float64](216.18),
                            UnitPriceAmount: unifiedgosdk.Pointer[float64](108.09),
                        },
                    },
                    PartsAmount: unifiedgosdk.Pointer[float64](1106.6200000000001),
                    PayType: shared.DmsRepairorderJobPayTypeWarranty.ToPointer(),
                    SoldHours: unifiedgosdk.Pointer[float64](3.1),
                    Status: shared.DmsRepairorderJobStatusCompleted.ToPointer(),
                    TechnicianUserIds: []string{},
                    TotalAmount: unifiedgosdk.Pointer[float64](1553.02),
                },
                shared.DmsRepairorderJob{
                    ActualHours: unifiedgosdk.Pointer[float64](1.5),
                    Cause: unifiedgosdk.Pointer("Scheduled maintenance"),
                    Complaint: unifiedgosdk.Pointer("Noise when braking"),
                    Correction: unifiedgosdk.Pointer("Degero theca natus."),
                    ID: unifiedgosdk.Pointer("d6aa4d3a-4e31-4252-b3fc-144c6b59b698"),
                    LaborAmount: unifiedgosdk.Pointer[float64](232.4),
                    Labors: []shared.DmsRepairorderLabor{
                        shared.DmsRepairorderLabor{
                            ID: unifiedgosdk.Pointer("6d15a307-c8c5-43c6-8b29-6ea8013bee52"),
                            PayType: shared.DmsRepairorderLaborPayTypeCustomer.ToPointer(),
                            RateAmount: unifiedgosdk.Pointer[float64](166.0),
                            SoldHours: unifiedgosdk.Pointer[float64](1.4),
                            TechnicianUserIds: []string{},
                            TotalAmount: unifiedgosdk.Pointer[float64](232.4),
                        },
                    },
                    Name: unifiedgosdk.Pointer("Customer concern"),
                    Number: unifiedgosdk.Pointer("C"),
                    Parts: []shared.DmsPartline{
                        shared.DmsPartline{
                            ID: unifiedgosdk.Pointer("50dc4642-0ca8-474f-9a11-b1e86d34c9c6"),
                            PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                            Quantity: unifiedgosdk.Pointer[float64](1.0),
                            QuantityOrdered: unifiedgosdk.Pointer[float64](1.0),
                            TotalAmount: unifiedgosdk.Pointer[float64](216.02),
                            UnitPriceAmount: unifiedgosdk.Pointer[float64](216.02),
                        },
                        shared.DmsPartline{
                            ID: unifiedgosdk.Pointer("246d6be0-ee9c-47c5-9631-c9d202925390"),
                            PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                            Quantity: unifiedgosdk.Pointer[float64](2.0),
                            QuantityOrdered: unifiedgosdk.Pointer[float64](2.0),
                            TotalAmount: unifiedgosdk.Pointer[float64](358.16),
                            UnitPriceAmount: unifiedgosdk.Pointer[float64](179.08),
                        },
                    },
                    PartsAmount: unifiedgosdk.Pointer[float64](574.1800000000001),
                    PayType: shared.DmsRepairorderJobPayTypeCustomer.ToPointer(),
                    SoldHours: unifiedgosdk.Pointer[float64](1.4),
                    Status: shared.DmsRepairorderJobStatusCompleted.ToPointer(),
                    TechnicianUserIds: []string{},
                    TotalAmount: unifiedgosdk.Pointer[float64](806.58),
                },
            },
            LaborAmount: unifiedgosdk.Pointer[float64](1360.4),
            Metadata: []shared.DmsMetadata{},
            MiscAmount: unifiedgosdk.Pointer[float64](35.0),
            Number: unifiedgosdk.Pointer("143398"),
            OdometerIn: unifiedgosdk.Pointer[float64](7776.0),
            OdometerOut: unifiedgosdk.Pointer[float64](7776.0),
            OdometerUnit: shared.DmsRepairorderOdometerUnitMiles.ToPointer(),
            OpenedAt: types.MustNewTimeFromString("2025-02-08T09:35:35.476Z"),
            PartsAmount: unifiedgosdk.Pointer[float64](1680.8000000000002),
            PayType: shared.DmsRepairorderPayTypeWarranty.ToPointer(),
            PromisedAt: types.MustNewTimeFromString("2025-02-09T17:35:35.476Z"),
            PurchaseOrderNumber: unifiedgosdk.Pointer("182743"),
            Status: shared.DmsRepairorderStatusCompleted.ToPointer(),
            SubletAmount: unifiedgosdk.Pointer[float64](0.0),
            TagNumber: unifiedgosdk.Pointer("758"),
            TaxAmount: unifiedgosdk.Pointer[float64](117.66),
            TechnicianUserIds: []string{},
            TotalAmount: unifiedgosdk.Pointer[float64](3193.86),
            Totals: []shared.DmsRepairorderTotal{
                shared.DmsRepairorderTotal{
                    DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
                    LaborAmount: unifiedgosdk.Pointer[float64](232.4),
                    PartsAmount: unifiedgosdk.Pointer[float64](574.1800000000001),
                    PayType: shared.DmsRepairorderTotalPayTypeCustomer.ToPointer(),
                    SubletAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](806.58),
                },
                shared.DmsRepairorderTotal{
                    DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
                    LaborAmount: unifiedgosdk.Pointer[float64](1128.0),
                    PartsAmount: unifiedgosdk.Pointer[float64](1106.6200000000001),
                    PayType: shared.DmsRepairorderTotalPayTypeWarranty.ToPointer(),
                    SubletAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](2234.62),
                },
            },
            TransportType: shared.DmsRepairorderTransportTypeLoaner.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-11-28T17:30:18.724Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsRepairorder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PatchDmsRepairorderRequest](../../pkg/models/operations/patchdmsrepairorderrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.PatchDmsRepairorderResponse](../../pkg/models/operations/patchdmsrepairorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchDmsVehicle

Updates the DmsVehicle identified by `id`. Provide the unified DmsVehicle fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchDmsVehicle" method="patch" path="/dms/{connection_id}/vehicle/{id}" example="dms_vehicle" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.PatchDmsVehicle(ctx, operations.PatchDmsVehicleRequest{
        DmsVehicle: shared.DmsVehicle{
            BodyStyle: unifiedgosdk.Pointer("Pickup"),
            Condition: shared.ConditionNew.ToPointer(),
            CostAmount: unifiedgosdk.Pointer[float64](26498.0),
            CreatedAt: types.MustNewTimeFromString("2021-02-03T08:01:58.559Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Description: unifiedgosdk.Pointer("2023 Citroën Silverado"),
            ExteriorColor: unifiedgosdk.Pointer("lavender"),
            FuelType: shared.FuelTypeElectric.ToPointer(),
            ID: unifiedgosdk.Pointer("13ebcf8a-fc93-446c-b117-c623a1a150f6"),
            ImageUrls: []string{
                "https://picsum.photos/seed/tOUr5K/1358/1877",
            },
            InteriorColor: unifiedgosdk.Pointer("Black"),
            InternetPriceAmount: unifiedgosdk.Pointer[float64](29526.0),
            IsActive: unifiedgosdk.Pointer(true),
            IsCertified: unifiedgosdk.Pointer(false),
            IsInventory: unifiedgosdk.Pointer(true),
            IsPublished: unifiedgosdk.Pointer(false),
            Make: unifiedgosdk.Pointer("Citroën"),
            Metadata: []shared.DmsMetadata{},
            Model: unifiedgosdk.Pointer("Silverado"),
            ModelCode: unifiedgosdk.Pointer("PMU8KI"),
            MsrpAmount: unifiedgosdk.Pointer[float64](33100.0),
            Odometer: unifiedgosdk.Pointer[float64](1.0),
            OdometerUnit: shared.DmsVehicleOdometerUnitMiles.ToPointer(),
            Options: []shared.DmsVehicleOption{
                shared.DmsVehicleOption{
                    Code: unifiedgosdk.Pointer("SXE"),
                    CostAmount: unifiedgosdk.Pointer[float64](250.0),
                    Description: unifiedgosdk.Pointer("Tow package"),
                    PriceAmount: unifiedgosdk.Pointer[float64](1400.0),
                    Type: shared.DmsVehicleOptionTypeDealer.ToPointer(),
                },
                shared.DmsVehicleOption{
                    Code: unifiedgosdk.Pointer("458"),
                    CostAmount: unifiedgosdk.Pointer[float64](700.0),
                    Description: unifiedgosdk.Pointer("Tow package"),
                    PriceAmount: unifiedgosdk.Pointer[float64](600.0),
                    Type: shared.DmsVehicleOptionTypeDealer.ToPointer(),
                },
            },
            PriceAmount: unifiedgosdk.Pointer[float64](29926.0),
            ReceivedAt: types.MustNewTimeFromString("2021-02-03T08:01:58.559Z"),
            Status: shared.DmsVehicleStatusOther.ToPointer(),
            StockNumber: unifiedgosdk.Pointer("N24959"),
            Transmission: unifiedgosdk.Pointer("Manual"),
            Trim: unifiedgosdk.Pointer("Base"),
            Type: shared.DmsVehicleTypeSuv.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-11-03T13:50:06.439Z"),
            Vin: unifiedgosdk.Pointer("C29ZDU474DDR70955"),
            WarrantyExpiresAt: types.MustNewTimeFromString("2022-03-28T09:10:09.889Z"),
            Year: unifiedgosdk.Pointer[float64](2023.0),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsVehicle != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PatchDmsVehicleRequest](../../pkg/models/operations/patchdmsvehiclerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.PatchDmsVehicleResponse](../../pkg/models/operations/patchdmsvehicleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchDmsVehiclesale

Updates the DmsVehiclesale identified by `id`. Provide the unified DmsVehiclesale fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchDmsVehiclesale" method="patch" path="/dms/{connection_id}/vehiclesale/{id}" example="dms_vehiclesale" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.PatchDmsVehiclesale(ctx, operations.PatchDmsVehiclesaleRequest{
        DmsVehiclesale: shared.DmsVehiclesale{
            BackGrossAmount: unifiedgosdk.Pointer[float64](4565.0),
            CreatedAt: types.MustNewTimeFromString("2020-04-02T08:14:56.150Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DownPaymentAmount: unifiedgosdk.Pointer[float64](6700.0),
            Fees: []shared.DmsVehiclesaleFee{
                shared.DmsVehiclesaleFee{
                    CostAmount: unifiedgosdk.Pointer[float64](0.0),
                    IsCapitalized: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Documentation fee"),
                    PriceAmount: unifiedgosdk.Pointer[float64](499.0),
                },
                shared.DmsVehiclesaleFee{
                    IsCapitalized: unifiedgosdk.Pointer(false),
                    Name: unifiedgosdk.Pointer("Title and registration"),
                    PriceAmount: unifiedgosdk.Pointer[float64](392.0),
                },
            },
            Finance: &shared.PropertyDmsVehiclesaleFinance{
                FinancedAmount: unifiedgosdk.Pointer[float64](0.0),
                FirstPaymentAt: types.MustNewTimeFromString("2022-08-16T11:36:15.407Z"),
                InterestRate: unifiedgosdk.Pointer[float64](6.68),
                LenderName: unifiedgosdk.Pointer("Mueller, Yundt and Murray Financial"),
                PaymentAmount: unifiedgosdk.Pointer[float64](339.0),
                PaymentFrequency: shared.PaymentFrequencyMonthly.ToPointer(),
                TermMonths: unifiedgosdk.Pointer[float64](72.0),
            },
            FinanceType: shared.FinanceTypeFinance.ToPointer(),
            ID: unifiedgosdk.Pointer("feb296bd-f5b4-4b47-a430-6ef8d0f9a7a5"),
            Metadata: []shared.DmsMetadata{},
            Number: unifiedgosdk.Pointer("D84591"),
            Products: []shared.DmsVehiclesaleProduct{
                shared.DmsVehiclesaleProduct{
                    ContractNumber: unifiedgosdk.Pointer("B7JN10GVB7"),
                    CostAmount: unifiedgosdk.Pointer[float64](1320.0),
                    Name: unifiedgosdk.Pointer("Tire & Wheel"),
                    PriceAmount: unifiedgosdk.Pointer[float64](3300.0),
                    Provider: unifiedgosdk.Pointer("Renner and Sons"),
                    TermMonths: unifiedgosdk.Pointer[float64](60.0),
                    Type: shared.DmsVehiclesaleProductTypeProtection.ToPointer(),
                },
                shared.DmsVehiclesaleProduct{
                    ContractNumber: unifiedgosdk.Pointer("VIZU7K3VRU"),
                    CostAmount: unifiedgosdk.Pointer[float64](1115.0),
                    DeductibleAmount: unifiedgosdk.Pointer[float64](100.0),
                    Name: unifiedgosdk.Pointer("Platinum Vehicle Service Contract"),
                    PriceAmount: unifiedgosdk.Pointer[float64](3700.0),
                    Provider: unifiedgosdk.Pointer("Frami, Wisozk and Gleichner"),
                    TermMonths: unifiedgosdk.Pointer[float64](72.0),
                    Type: shared.DmsVehiclesaleProductTypeServiceContract.ToPointer(),
                },
            },
            RebateAmount: unifiedgosdk.Pointer[float64](500.0),
            SoldAt: types.MustNewTimeFromString("2022-08-14T20:20:41.120Z"),
            Source: unifiedgosdk.Pointer("Website"),
            Status: shared.DmsVehiclesaleStatusSold.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            Taxes: []shared.DmsVehiclesaleTax{
                shared.DmsVehiclesaleTax{
                    Name: unifiedgosdk.Pointer("State sales tax"),
                    Rate: unifiedgosdk.Pointer[float64](8.25),
                    TotalAmount: unifiedgosdk.Pointer[float64](0.0),
                },
            },
            TotalAmount: unifiedgosdk.Pointer[float64](7891.0),
            TradeIns: []shared.DmsVehiclesaleTradein{},
            Type: shared.DmsVehiclesaleTypeRetail.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-04-16T21:50:54.903Z"),
            Vehicles: []shared.DmsVehiclesaleVehicle{},
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsVehiclesale != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PatchDmsVehiclesaleRequest](../../pkg/models/operations/patchdmsvehiclesalerequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.PatchDmsVehiclesaleResponse](../../pkg/models/operations/patchdmsvehiclesaleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveDmsAppointment

Deletes the DmsAppointment identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeDmsAppointment" method="delete" path="/dms/{connection_id}/appointment/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.RemoveDmsAppointment(ctx, operations.RemoveDmsAppointmentRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.RemoveDmsAppointmentRequest](../../pkg/models/operations/removedmsappointmentrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.RemoveDmsAppointmentResponse](../../pkg/models/operations/removedmsappointmentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateDmsAppointment

Updates the DmsAppointment identified by `id`. Provide the unified DmsAppointment fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateDmsAppointment" method="put" path="/dms/{connection_id}/appointment/{id}" example="dms_appointment" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.UpdateDmsAppointment(ctx, operations.UpdateDmsAppointmentRequest{
        DmsAppointment: shared.DmsAppointment{
            CreatedAt: types.MustNewTimeFromString("2023-10-16T02:56:45.127Z"),
            EndAt: types.MustNewTimeFromString("2026-05-08T02:39:01.800Z"),
            ID: unifiedgosdk.Pointer("c0e5894c-5d3a-4c4a-90a3-f8511fcaf3e2"),
            Metadata: []shared.DmsMetadata{},
            Notes: unifiedgosdk.Pointer("Soluta conitor timidus confero."),
            Number: unifiedgosdk.Pointer("A541909"),
            Services: []shared.DmsAppointmentService{
                shared.DmsAppointmentService{
                    ID: unifiedgosdk.Pointer("f74a8fff-8526-4417-b048-4c972bd4b31b"),
                    Name: unifiedgosdk.Pointer("Customer concern"),
                    PayType: shared.PayTypeWarranty.ToPointer(),
                    TechnicianUserIds: []string{},
                },
                shared.DmsAppointmentService{
                    Complaint: unifiedgosdk.Pointer("Check engine light on"),
                    ID: unifiedgosdk.Pointer("90c323d7-deaa-447d-a591-e1ce6f4d9803"),
                    Name: unifiedgosdk.Pointer("Customer concern"),
                    PayType: shared.PayTypeCustomer.ToPointer(),
                    TechnicianUserIds: []string{},
                },
                shared.DmsAppointmentService{
                    Complaint: unifiedgosdk.Pointer("Due for service"),
                    ID: unifiedgosdk.Pointer("e73e1b90-395a-4afb-8ca7-1b108fb48bc3"),
                    Name: unifiedgosdk.Pointer("Customer concern"),
                    PayType: shared.PayTypeCustomer.ToPointer(),
                    TechnicianUserIds: []string{},
                },
            },
            StartAt: types.MustNewTimeFromString("2026-05-07T19:39:01.800Z"),
            Status: shared.DmsAppointmentStatusCancelled.ToPointer(),
            TransportType: shared.TransportTypeLoaner.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-04-06T16:29:51.858Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsAppointment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdateDmsAppointmentRequest](../../pkg/models/operations/updatedmsappointmentrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.UpdateDmsAppointmentResponse](../../pkg/models/operations/updatedmsappointmentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateDmsOpcode

Updates the DmsOpcode identified by `id`. Provide the unified DmsOpcode fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateDmsOpcode" method="put" path="/dms/{connection_id}/opcode/{id}" example="dms_opcode" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.UpdateDmsOpcode(ctx, operations.UpdateDmsOpcodeRequest{
        DmsOpcode: shared.DmsOpcode{
            Category: unifiedgosdk.Pointer("Electrical"),
            Code: unifiedgosdk.Pointer("02AUS"),
            CreatedAt: types.MustNewTimeFromString("2021-01-12T19:03:17.545Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Description: unifiedgosdk.Pointer("Oil and filter change"),
            EstimatedHours: unifiedgosdk.Pointer[float64](3.2),
            FlatRateHours: unifiedgosdk.Pointer[float64](3.5),
            ID: unifiedgosdk.Pointer("95bbbacf-b8ea-4841-8a1a-03f73f25cbda"),
            IsActive: unifiedgosdk.Pointer(true),
            LaborRateAmount: unifiedgosdk.Pointer[float64](216.0),
            Metadata: []shared.DmsMetadata{},
            OemCode: unifiedgosdk.Pointer("05604A3"),
            Parts: []shared.DmsPartline{},
            PayType: shared.DmsOpcodePayTypeCustomer.ToPointer(),
            SkillCode: unifiedgosdk.Pointer("A"),
            Type: shared.DmsOpcodeTypeLabor.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2021-11-03T15:48:51.325Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsOpcode != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateDmsOpcodeRequest](../../pkg/models/operations/updatedmsopcoderequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.UpdateDmsOpcodeResponse](../../pkg/models/operations/updatedmsopcoderesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateDmsPart

Updates the DmsPart identified by `id`. Provide the unified DmsPart fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateDmsPart" method="put" path="/dms/{connection_id}/part/{id}" example="dms_part" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.UpdateDmsPart(ctx, operations.UpdateDmsPartRequest{
        DmsPart: shared.DmsPart{
            BinLocations: []string{
                "S18-6",
            },
            Category: unifiedgosdk.Pointer("Ignition"),
            CostAmount: unifiedgosdk.Pointer[float64](25.26),
            CreatedAt: types.MustNewTimeFromString("2020-07-02T08:18:03.858Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Description: unifiedgosdk.Pointer("Cabin filter"),
            ID: unifiedgosdk.Pointer("44358212-7c73-4818-bd09-04036b6311db"),
            LastSoldAt: types.MustNewTimeFromString("2020-11-23T04:33:10.427Z"),
            ListPriceAmount: unifiedgosdk.Pointer[float64](39.61),
            ManufacturerCode: unifiedgosdk.Pointer("TOY"),
            MaxStockQuantity: unifiedgosdk.Pointer[float64](21.0),
            Metadata: []shared.DmsMetadata{},
            MinStockQuantity: unifiedgosdk.Pointer[float64](3.0),
            PackQuantity: unifiedgosdk.Pointer[float64](1.0),
            PartNumber: unifiedgosdk.Pointer("73992-IM3FI"),
            QuantityAvailable: unifiedgosdk.Pointer[float64](13.0),
            QuantityOnHand: unifiedgosdk.Pointer[float64](15.0),
            QuantityOnOrder: unifiedgosdk.Pointer[float64](14.0),
            SourceCode: unifiedgosdk.Pointer("30"),
            TradePriceAmount: unifiedgosdk.Pointer[float64](29.05),
            UpdatedAt: types.MustNewTimeFromString("2023-11-23T07:05:30.256Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsPart != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateDmsPartRequest](../../pkg/models/operations/updatedmspartrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.UpdateDmsPartResponse](../../pkg/models/operations/updatedmspartresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateDmsPartsale

Updates the DmsPartsale identified by `id`. Provide the unified DmsPartsale fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateDmsPartsale" method="put" path="/dms/{connection_id}/partsale/{id}" example="dms_partsale" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.UpdateDmsPartsale(ctx, operations.UpdateDmsPartsaleRequest{
        DmsPartsale: shared.DmsPartsale{
            CostAmount: unifiedgosdk.Pointer[float64](0.0),
            CreatedAt: types.MustNewTimeFromString("2023-05-22T03:39:30.335Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            FreightAmount: unifiedgosdk.Pointer[float64](0.0),
            ID: unifiedgosdk.Pointer("c8f3d6d4-f9f0-4bf3-9721-4d7e489e98ba"),
            Lines: []shared.DmsPartline{
                shared.DmsPartline{
                    ID: unifiedgosdk.Pointer("13b3eda3-f590-4de6-a13d-499d910cfc23"),
                    PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                    Quantity: unifiedgosdk.Pointer[float64](5.0),
                    QuantityOrdered: unifiedgosdk.Pointer[float64](5.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](491.45),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](98.29),
                },
                shared.DmsPartline{
                    ID: unifiedgosdk.Pointer("939aa1f5-3dfd-4b82-838b-240da4c49ea6"),
                    PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                    Quantity: unifiedgosdk.Pointer[float64](2.0),
                    QuantityOrdered: unifiedgosdk.Pointer[float64](2.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](240.9),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](120.45),
                },
                shared.DmsPartline{
                    ID: unifiedgosdk.Pointer("f65e363c-3828-4710-909c-2b319567d354"),
                    PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                    Quantity: unifiedgosdk.Pointer[float64](3.0),
                    QuantityOrdered: unifiedgosdk.Pointer[float64](3.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](249.54),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](83.18),
                },
            },
            Metadata: []shared.DmsMetadata{},
            Number: unifiedgosdk.Pointer("P292867"),
            PayType: shared.DmsPartsalePayTypeCustomer.ToPointer(),
            Status: shared.DmsPartsaleStatusVoided.ToPointer(),
            SubtotalAmount: unifiedgosdk.Pointer[float64](981.89),
            TaxAmount: unifiedgosdk.Pointer[float64](68.73),
            TotalAmount: unifiedgosdk.Pointer[float64](1050.62),
            Type: shared.DmsPartsaleTypeSale.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-04-18T13:15:50.920Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsPartsale != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateDmsPartsaleRequest](../../pkg/models/operations/updatedmspartsalerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateDmsPartsaleResponse](../../pkg/models/operations/updatedmspartsaleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateDmsRepairorder

Updates the DmsRepairorder identified by `id`. Provide the unified DmsRepairorder fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateDmsRepairorder" method="put" path="/dms/{connection_id}/repairorder/{id}" example="dms_repairorder" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.UpdateDmsRepairorder(ctx, operations.UpdateDmsRepairorderRequest{
        DmsRepairorder: shared.DmsRepairorder{
            BalanceAmount: unifiedgosdk.Pointer[float64](3193.86),
            CompletedAt: types.MustNewTimeFromString("2026-11-13T20:43:09.493Z"),
            CreatedAt: types.MustNewTimeFromString("2022-10-25T15:50:48.007Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            EstimateAmount: unifiedgosdk.Pointer[float64](3264.0),
            ID: unifiedgosdk.Pointer("0187b8d9-283c-4987-a821-3a99a34b88cf"),
            Jobs: []shared.DmsRepairorderJob{
                shared.DmsRepairorderJob{
                    ActualHours: unifiedgosdk.Pointer[float64](3.6),
                    Cause: unifiedgosdk.Pointer("Scheduled maintenance"),
                    Complaint: unifiedgosdk.Pointer("Due for service"),
                    Correction: unifiedgosdk.Pointer("Aufero cumque vicinus aptus assumenda quis aptus."),
                    ID: unifiedgosdk.Pointer("03beb52d-f18e-42ce-a286-39ca1d7d2652"),
                    LaborAmount: unifiedgosdk.Pointer[float64](681.6),
                    Labors: []shared.DmsRepairorderLabor{
                        shared.DmsRepairorderLabor{
                            ID: unifiedgosdk.Pointer("5ebedfb0-e6f8-4d3d-a820-632a6edccb14"),
                            PayType: shared.DmsRepairorderLaborPayTypeWarranty.ToPointer(),
                            RateAmount: unifiedgosdk.Pointer[float64](213.0),
                            SoldHours: unifiedgosdk.Pointer[float64](3.2),
                            TechnicianUserIds: []string{},
                            TotalAmount: unifiedgosdk.Pointer[float64](681.6),
                        },
                    },
                    Name: unifiedgosdk.Pointer("Customer concern"),
                    Number: unifiedgosdk.Pointer("A"),
                    Parts: []shared.DmsPartline{},
                    PartsAmount: unifiedgosdk.Pointer[float64](0.0),
                    PayType: shared.DmsRepairorderJobPayTypeWarranty.ToPointer(),
                    SoldHours: unifiedgosdk.Pointer[float64](3.2),
                    Status: shared.DmsRepairorderJobStatusCompleted.ToPointer(),
                    TechnicianUserIds: []string{},
                    TotalAmount: unifiedgosdk.Pointer[float64](681.6),
                },
                shared.DmsRepairorderJob{
                    ActualHours: unifiedgosdk.Pointer[float64](3.0),
                    Cause: unifiedgosdk.Pointer("Worn brake pads"),
                    Complaint: unifiedgosdk.Pointer("Vibration at highway speed"),
                    Correction: unifiedgosdk.Pointer("Somnus avaritia utroque vulnero ullam ancilla."),
                    ID: unifiedgosdk.Pointer("f2f8379e-6a83-4a06-aa7a-d6b1e9e53763"),
                    LaborAmount: unifiedgosdk.Pointer[float64](446.4),
                    Labors: []shared.DmsRepairorderLabor{
                        shared.DmsRepairorderLabor{
                            ID: unifiedgosdk.Pointer("e1c638b6-da08-4e1f-aad7-59db110190fc"),
                            PayType: shared.DmsRepairorderLaborPayTypeWarranty.ToPointer(),
                            RateAmount: unifiedgosdk.Pointer[float64](144.0),
                            SoldHours: unifiedgosdk.Pointer[float64](3.1),
                            TechnicianUserIds: []string{},
                            TotalAmount: unifiedgosdk.Pointer[float64](446.4),
                        },
                    },
                    Name: unifiedgosdk.Pointer("Customer concern"),
                    Number: unifiedgosdk.Pointer("B"),
                    Parts: []shared.DmsPartline{
                        shared.DmsPartline{
                            ID: unifiedgosdk.Pointer("d246b71c-d4a0-443f-bbde-7382f2328ffa"),
                            PayType: shared.DmsPartlinePayTypeWarranty.ToPointer(),
                            Quantity: unifiedgosdk.Pointer[float64](4.0),
                            QuantityOrdered: unifiedgosdk.Pointer[float64](4.0),
                            TotalAmount: unifiedgosdk.Pointer[float64](890.44),
                            UnitPriceAmount: unifiedgosdk.Pointer[float64](222.61),
                        },
                        shared.DmsPartline{
                            ID: unifiedgosdk.Pointer("7f781967-2d75-49f8-b0f5-10cc295c186e"),
                            PayType: shared.DmsPartlinePayTypeWarranty.ToPointer(),
                            Quantity: unifiedgosdk.Pointer[float64](2.0),
                            QuantityOrdered: unifiedgosdk.Pointer[float64](2.0),
                            TotalAmount: unifiedgosdk.Pointer[float64](216.18),
                            UnitPriceAmount: unifiedgosdk.Pointer[float64](108.09),
                        },
                    },
                    PartsAmount: unifiedgosdk.Pointer[float64](1106.6200000000001),
                    PayType: shared.DmsRepairorderJobPayTypeWarranty.ToPointer(),
                    SoldHours: unifiedgosdk.Pointer[float64](3.1),
                    Status: shared.DmsRepairorderJobStatusCompleted.ToPointer(),
                    TechnicianUserIds: []string{},
                    TotalAmount: unifiedgosdk.Pointer[float64](1553.02),
                },
                shared.DmsRepairorderJob{
                    ActualHours: unifiedgosdk.Pointer[float64](1.5),
                    Cause: unifiedgosdk.Pointer("Scheduled maintenance"),
                    Complaint: unifiedgosdk.Pointer("Noise when braking"),
                    Correction: unifiedgosdk.Pointer("Degero theca natus."),
                    ID: unifiedgosdk.Pointer("d6aa4d3a-4e31-4252-b3fc-144c6b59b698"),
                    LaborAmount: unifiedgosdk.Pointer[float64](232.4),
                    Labors: []shared.DmsRepairorderLabor{
                        shared.DmsRepairorderLabor{
                            ID: unifiedgosdk.Pointer("6d15a307-c8c5-43c6-8b29-6ea8013bee52"),
                            PayType: shared.DmsRepairorderLaborPayTypeCustomer.ToPointer(),
                            RateAmount: unifiedgosdk.Pointer[float64](166.0),
                            SoldHours: unifiedgosdk.Pointer[float64](1.4),
                            TechnicianUserIds: []string{},
                            TotalAmount: unifiedgosdk.Pointer[float64](232.4),
                        },
                    },
                    Name: unifiedgosdk.Pointer("Customer concern"),
                    Number: unifiedgosdk.Pointer("C"),
                    Parts: []shared.DmsPartline{
                        shared.DmsPartline{
                            ID: unifiedgosdk.Pointer("50dc4642-0ca8-474f-9a11-b1e86d34c9c6"),
                            PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                            Quantity: unifiedgosdk.Pointer[float64](1.0),
                            QuantityOrdered: unifiedgosdk.Pointer[float64](1.0),
                            TotalAmount: unifiedgosdk.Pointer[float64](216.02),
                            UnitPriceAmount: unifiedgosdk.Pointer[float64](216.02),
                        },
                        shared.DmsPartline{
                            ID: unifiedgosdk.Pointer("246d6be0-ee9c-47c5-9631-c9d202925390"),
                            PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                            Quantity: unifiedgosdk.Pointer[float64](2.0),
                            QuantityOrdered: unifiedgosdk.Pointer[float64](2.0),
                            TotalAmount: unifiedgosdk.Pointer[float64](358.16),
                            UnitPriceAmount: unifiedgosdk.Pointer[float64](179.08),
                        },
                    },
                    PartsAmount: unifiedgosdk.Pointer[float64](574.1800000000001),
                    PayType: shared.DmsRepairorderJobPayTypeCustomer.ToPointer(),
                    SoldHours: unifiedgosdk.Pointer[float64](1.4),
                    Status: shared.DmsRepairorderJobStatusCompleted.ToPointer(),
                    TechnicianUserIds: []string{},
                    TotalAmount: unifiedgosdk.Pointer[float64](806.58),
                },
            },
            LaborAmount: unifiedgosdk.Pointer[float64](1360.4),
            Metadata: []shared.DmsMetadata{},
            MiscAmount: unifiedgosdk.Pointer[float64](35.0),
            Number: unifiedgosdk.Pointer("143398"),
            OdometerIn: unifiedgosdk.Pointer[float64](7776.0),
            OdometerOut: unifiedgosdk.Pointer[float64](7776.0),
            OdometerUnit: shared.DmsRepairorderOdometerUnitMiles.ToPointer(),
            OpenedAt: types.MustNewTimeFromString("2025-02-08T09:35:35.476Z"),
            PartsAmount: unifiedgosdk.Pointer[float64](1680.8000000000002),
            PayType: shared.DmsRepairorderPayTypeWarranty.ToPointer(),
            PromisedAt: types.MustNewTimeFromString("2025-02-09T17:35:35.476Z"),
            PurchaseOrderNumber: unifiedgosdk.Pointer("182743"),
            Status: shared.DmsRepairorderStatusCompleted.ToPointer(),
            SubletAmount: unifiedgosdk.Pointer[float64](0.0),
            TagNumber: unifiedgosdk.Pointer("758"),
            TaxAmount: unifiedgosdk.Pointer[float64](117.66),
            TechnicianUserIds: []string{},
            TotalAmount: unifiedgosdk.Pointer[float64](3193.86),
            Totals: []shared.DmsRepairorderTotal{
                shared.DmsRepairorderTotal{
                    DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
                    LaborAmount: unifiedgosdk.Pointer[float64](232.4),
                    PartsAmount: unifiedgosdk.Pointer[float64](574.1800000000001),
                    PayType: shared.DmsRepairorderTotalPayTypeCustomer.ToPointer(),
                    SubletAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](806.58),
                },
                shared.DmsRepairorderTotal{
                    DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
                    LaborAmount: unifiedgosdk.Pointer[float64](1128.0),
                    PartsAmount: unifiedgosdk.Pointer[float64](1106.6200000000001),
                    PayType: shared.DmsRepairorderTotalPayTypeWarranty.ToPointer(),
                    SubletAmount: unifiedgosdk.Pointer[float64](0.0),
                    TotalAmount: unifiedgosdk.Pointer[float64](2234.62),
                },
            },
            TransportType: shared.DmsRepairorderTransportTypeLoaner.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-11-28T17:30:18.724Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsRepairorder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdateDmsRepairorderRequest](../../pkg/models/operations/updatedmsrepairorderrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.UpdateDmsRepairorderResponse](../../pkg/models/operations/updatedmsrepairorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateDmsVehicle

Updates the DmsVehicle identified by `id`. Provide the unified DmsVehicle fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateDmsVehicle" method="put" path="/dms/{connection_id}/vehicle/{id}" example="dms_vehicle" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.UpdateDmsVehicle(ctx, operations.UpdateDmsVehicleRequest{
        DmsVehicle: shared.DmsVehicle{
            BodyStyle: unifiedgosdk.Pointer("Pickup"),
            Condition: shared.ConditionNew.ToPointer(),
            CostAmount: unifiedgosdk.Pointer[float64](26498.0),
            CreatedAt: types.MustNewTimeFromString("2021-02-03T08:01:58.559Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Description: unifiedgosdk.Pointer("2023 Citroën Silverado"),
            ExteriorColor: unifiedgosdk.Pointer("lavender"),
            FuelType: shared.FuelTypeElectric.ToPointer(),
            ID: unifiedgosdk.Pointer("13ebcf8a-fc93-446c-b117-c623a1a150f6"),
            ImageUrls: []string{
                "https://picsum.photos/seed/tOUr5K/1358/1877",
            },
            InteriorColor: unifiedgosdk.Pointer("Black"),
            InternetPriceAmount: unifiedgosdk.Pointer[float64](29526.0),
            IsActive: unifiedgosdk.Pointer(true),
            IsCertified: unifiedgosdk.Pointer(false),
            IsInventory: unifiedgosdk.Pointer(true),
            IsPublished: unifiedgosdk.Pointer(false),
            Make: unifiedgosdk.Pointer("Citroën"),
            Metadata: []shared.DmsMetadata{},
            Model: unifiedgosdk.Pointer("Silverado"),
            ModelCode: unifiedgosdk.Pointer("PMU8KI"),
            MsrpAmount: unifiedgosdk.Pointer[float64](33100.0),
            Odometer: unifiedgosdk.Pointer[float64](1.0),
            OdometerUnit: shared.DmsVehicleOdometerUnitMiles.ToPointer(),
            Options: []shared.DmsVehicleOption{
                shared.DmsVehicleOption{
                    Code: unifiedgosdk.Pointer("SXE"),
                    CostAmount: unifiedgosdk.Pointer[float64](250.0),
                    Description: unifiedgosdk.Pointer("Tow package"),
                    PriceAmount: unifiedgosdk.Pointer[float64](1400.0),
                    Type: shared.DmsVehicleOptionTypeDealer.ToPointer(),
                },
                shared.DmsVehicleOption{
                    Code: unifiedgosdk.Pointer("458"),
                    CostAmount: unifiedgosdk.Pointer[float64](700.0),
                    Description: unifiedgosdk.Pointer("Tow package"),
                    PriceAmount: unifiedgosdk.Pointer[float64](600.0),
                    Type: shared.DmsVehicleOptionTypeDealer.ToPointer(),
                },
            },
            PriceAmount: unifiedgosdk.Pointer[float64](29926.0),
            ReceivedAt: types.MustNewTimeFromString("2021-02-03T08:01:58.559Z"),
            Status: shared.DmsVehicleStatusOther.ToPointer(),
            StockNumber: unifiedgosdk.Pointer("N24959"),
            Transmission: unifiedgosdk.Pointer("Manual"),
            Trim: unifiedgosdk.Pointer("Base"),
            Type: shared.DmsVehicleTypeSuv.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-11-03T13:50:06.439Z"),
            Vin: unifiedgosdk.Pointer("C29ZDU474DDR70955"),
            WarrantyExpiresAt: types.MustNewTimeFromString("2022-03-28T09:10:09.889Z"),
            Year: unifiedgosdk.Pointer[float64](2023.0),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsVehicle != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateDmsVehicleRequest](../../pkg/models/operations/updatedmsvehiclerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpdateDmsVehicleResponse](../../pkg/models/operations/updatedmsvehicleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateDmsVehiclesale

Updates the DmsVehiclesale identified by `id`. Provide the unified DmsVehiclesale fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateDmsVehiclesale" method="put" path="/dms/{connection_id}/vehiclesale/{id}" example="dms_vehiclesale" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Dms.UpdateDmsVehiclesale(ctx, operations.UpdateDmsVehiclesaleRequest{
        DmsVehiclesale: shared.DmsVehiclesale{
            BackGrossAmount: unifiedgosdk.Pointer[float64](4565.0),
            CreatedAt: types.MustNewTimeFromString("2020-04-02T08:14:56.150Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DownPaymentAmount: unifiedgosdk.Pointer[float64](6700.0),
            Fees: []shared.DmsVehiclesaleFee{
                shared.DmsVehiclesaleFee{
                    CostAmount: unifiedgosdk.Pointer[float64](0.0),
                    IsCapitalized: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Documentation fee"),
                    PriceAmount: unifiedgosdk.Pointer[float64](499.0),
                },
                shared.DmsVehiclesaleFee{
                    IsCapitalized: unifiedgosdk.Pointer(false),
                    Name: unifiedgosdk.Pointer("Title and registration"),
                    PriceAmount: unifiedgosdk.Pointer[float64](392.0),
                },
            },
            Finance: &shared.PropertyDmsVehiclesaleFinance{
                FinancedAmount: unifiedgosdk.Pointer[float64](0.0),
                FirstPaymentAt: types.MustNewTimeFromString("2022-08-16T11:36:15.407Z"),
                InterestRate: unifiedgosdk.Pointer[float64](6.68),
                LenderName: unifiedgosdk.Pointer("Mueller, Yundt and Murray Financial"),
                PaymentAmount: unifiedgosdk.Pointer[float64](339.0),
                PaymentFrequency: shared.PaymentFrequencyMonthly.ToPointer(),
                TermMonths: unifiedgosdk.Pointer[float64](72.0),
            },
            FinanceType: shared.FinanceTypeFinance.ToPointer(),
            ID: unifiedgosdk.Pointer("feb296bd-f5b4-4b47-a430-6ef8d0f9a7a5"),
            Metadata: []shared.DmsMetadata{},
            Number: unifiedgosdk.Pointer("D84591"),
            Products: []shared.DmsVehiclesaleProduct{
                shared.DmsVehiclesaleProduct{
                    ContractNumber: unifiedgosdk.Pointer("B7JN10GVB7"),
                    CostAmount: unifiedgosdk.Pointer[float64](1320.0),
                    Name: unifiedgosdk.Pointer("Tire & Wheel"),
                    PriceAmount: unifiedgosdk.Pointer[float64](3300.0),
                    Provider: unifiedgosdk.Pointer("Renner and Sons"),
                    TermMonths: unifiedgosdk.Pointer[float64](60.0),
                    Type: shared.DmsVehiclesaleProductTypeProtection.ToPointer(),
                },
                shared.DmsVehiclesaleProduct{
                    ContractNumber: unifiedgosdk.Pointer("VIZU7K3VRU"),
                    CostAmount: unifiedgosdk.Pointer[float64](1115.0),
                    DeductibleAmount: unifiedgosdk.Pointer[float64](100.0),
                    Name: unifiedgosdk.Pointer("Platinum Vehicle Service Contract"),
                    PriceAmount: unifiedgosdk.Pointer[float64](3700.0),
                    Provider: unifiedgosdk.Pointer("Frami, Wisozk and Gleichner"),
                    TermMonths: unifiedgosdk.Pointer[float64](72.0),
                    Type: shared.DmsVehiclesaleProductTypeServiceContract.ToPointer(),
                },
            },
            RebateAmount: unifiedgosdk.Pointer[float64](500.0),
            SoldAt: types.MustNewTimeFromString("2022-08-14T20:20:41.120Z"),
            Source: unifiedgosdk.Pointer("Website"),
            Status: shared.DmsVehiclesaleStatusSold.ToPointer(),
            TaxAmount: unifiedgosdk.Pointer[float64](0.0),
            Taxes: []shared.DmsVehiclesaleTax{
                shared.DmsVehiclesaleTax{
                    Name: unifiedgosdk.Pointer("State sales tax"),
                    Rate: unifiedgosdk.Pointer[float64](8.25),
                    TotalAmount: unifiedgosdk.Pointer[float64](0.0),
                },
            },
            TotalAmount: unifiedgosdk.Pointer[float64](7891.0),
            TradeIns: []shared.DmsVehiclesaleTradein{},
            Type: shared.DmsVehiclesaleTypeRetail.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-04-16T21:50:54.903Z"),
            Vehicles: []shared.DmsVehiclesaleVehicle{},
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DmsVehiclesale != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdateDmsVehiclesaleRequest](../../pkg/models/operations/updatedmsvehiclesalerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.UpdateDmsVehiclesaleResponse](../../pkg/models/operations/updatedmsvehiclesaleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |