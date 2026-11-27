# Repairorder

## Overview

### Available Operations

* [CreateDmsRepairorder](#createdmsrepairorder) - Create a repairorder
* [GetDmsRepairorder](#getdmsrepairorder) - Retrieve a repairorder
* [ListDmsRepairorders](#listdmsrepairorders) - List all repairorders
* [PatchDmsRepairorder](#patchdmsrepairorder) - Update a repairorder
* [UpdateDmsRepairorder](#updatedmsrepairorder) - Update a repairorder

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

    res, err := s.Repairorder.CreateDmsRepairorder(ctx, operations.CreateDmsRepairorderRequest{
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

    res, err := s.Repairorder.GetDmsRepairorder(ctx, operations.GetDmsRepairorderRequest{
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

    res, err := s.Repairorder.ListDmsRepairorders(ctx, operations.ListDmsRepairordersRequest{
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

    res, err := s.Repairorder.PatchDmsRepairorder(ctx, operations.PatchDmsRepairorderRequest{
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

    res, err := s.Repairorder.UpdateDmsRepairorder(ctx, operations.UpdateDmsRepairorderRequest{
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