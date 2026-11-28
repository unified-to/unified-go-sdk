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
            CompletedAt: types.MustNewTimeFromString("2026-11-14T20:19:18.692Z"),
            CreatedAt: types.MustNewTimeFromString("2022-10-25T15:50:48.007Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            EstimateAmount: unifiedgosdk.Pointer[float64](3264.0),
            ID: unifiedgosdk.Pointer("819eb449-c334-4695-a046-f3fd65ded588"),
            Jobs: []shared.DmsRepairorderJob{
                shared.DmsRepairorderJob{
                    ActualHours: unifiedgosdk.Pointer[float64](3.6),
                    Cause: unifiedgosdk.Pointer("Scheduled maintenance"),
                    Complaint: unifiedgosdk.Pointer("Due for service"),
                    Correction: unifiedgosdk.Pointer("Aufero cumque vicinus aptus assumenda quis aptus."),
                    ID: unifiedgosdk.Pointer("89fcd5f8-5ae5-419a-86a0-62e3c040618f"),
                    LaborAmount: unifiedgosdk.Pointer[float64](681.6),
                    Labors: []shared.DmsRepairorderLabor{
                        shared.DmsRepairorderLabor{
                            ID: unifiedgosdk.Pointer("fd7f51d5-a365-4de0-9584-fdd3c6b221a3"),
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
                    ID: unifiedgosdk.Pointer("eed10b5e-fa6b-437c-b17c-a9ff71d905a4"),
                    LaborAmount: unifiedgosdk.Pointer[float64](446.4),
                    Labors: []shared.DmsRepairorderLabor{
                        shared.DmsRepairorderLabor{
                            ID: unifiedgosdk.Pointer("ebaf93bf-dd62-4d91-bff8-aa2bf22fec24"),
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
                            ID: unifiedgosdk.Pointer("d9a8cf40-eb82-457b-be3c-aa555158c4d0"),
                            PayType: shared.DmsPartlinePayTypeWarranty.ToPointer(),
                            Quantity: unifiedgosdk.Pointer[float64](4.0),
                            QuantityOrdered: unifiedgosdk.Pointer[float64](4.0),
                            TotalAmount: unifiedgosdk.Pointer[float64](890.44),
                            UnitPriceAmount: unifiedgosdk.Pointer[float64](222.61),
                        },
                        shared.DmsPartline{
                            ID: unifiedgosdk.Pointer("3e954a02-ceb3-4782-b9ee-cceb3ae6a217"),
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
                    ID: unifiedgosdk.Pointer("695bba06-83d6-4ca1-a75f-4fdb2d06338f"),
                    LaborAmount: unifiedgosdk.Pointer[float64](232.4),
                    Labors: []shared.DmsRepairorderLabor{
                        shared.DmsRepairorderLabor{
                            ID: unifiedgosdk.Pointer("f3bde429-6333-4dd7-99a7-15acf96b8b07"),
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
                            ID: unifiedgosdk.Pointer("66be0f67-6f8a-4440-a180-41ed71e5da00"),
                            PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                            Quantity: unifiedgosdk.Pointer[float64](1.0),
                            QuantityOrdered: unifiedgosdk.Pointer[float64](1.0),
                            TotalAmount: unifiedgosdk.Pointer[float64](216.02),
                            UnitPriceAmount: unifiedgosdk.Pointer[float64](216.02),
                        },
                        shared.DmsPartline{
                            ID: unifiedgosdk.Pointer("9f7cc0f8-74c0-474d-8e79-5ed0109f6200"),
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
            OpenedAt: types.MustNewTimeFromString("2025-02-08T22:56:07.455Z"),
            PartsAmount: unifiedgosdk.Pointer[float64](1680.8000000000002),
            PayType: shared.DmsRepairorderPayTypeWarranty.ToPointer(),
            PromisedAt: types.MustNewTimeFromString("2025-02-10T06:56:07.455Z"),
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
            UpdatedAt: types.MustNewTimeFromString("2025-11-29T11:31:28.925Z"),
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
            CompletedAt: types.MustNewTimeFromString("2026-11-14T20:19:18.722Z"),
            CreatedAt: types.MustNewTimeFromString("2022-10-25T15:50:48.007Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            EstimateAmount: unifiedgosdk.Pointer[float64](3264.0),
            ID: unifiedgosdk.Pointer("5e6e5625-3b8e-47dc-acc7-5e729a4d349a"),
            Jobs: []shared.DmsRepairorderJob{
                shared.DmsRepairorderJob{
                    ActualHours: unifiedgosdk.Pointer[float64](3.6),
                    Cause: unifiedgosdk.Pointer("Scheduled maintenance"),
                    Complaint: unifiedgosdk.Pointer("Due for service"),
                    Correction: unifiedgosdk.Pointer("Aufero cumque vicinus aptus assumenda quis aptus."),
                    ID: unifiedgosdk.Pointer("fe995c55-405f-4553-bc93-d07fc263c903"),
                    LaborAmount: unifiedgosdk.Pointer[float64](681.6),
                    Labors: []shared.DmsRepairorderLabor{
                        shared.DmsRepairorderLabor{
                            ID: unifiedgosdk.Pointer("55383436-f938-4acc-8f1d-05754cd6070d"),
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
                    ID: unifiedgosdk.Pointer("c4ed33f6-d936-4305-9dea-73914166df09"),
                    LaborAmount: unifiedgosdk.Pointer[float64](446.4),
                    Labors: []shared.DmsRepairorderLabor{
                        shared.DmsRepairorderLabor{
                            ID: unifiedgosdk.Pointer("b1efda49-94db-4bb0-89bb-1efd331e8882"),
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
                            ID: unifiedgosdk.Pointer("10ee34f6-ec4a-47da-810f-41703382e2e1"),
                            PayType: shared.DmsPartlinePayTypeWarranty.ToPointer(),
                            Quantity: unifiedgosdk.Pointer[float64](4.0),
                            QuantityOrdered: unifiedgosdk.Pointer[float64](4.0),
                            TotalAmount: unifiedgosdk.Pointer[float64](890.44),
                            UnitPriceAmount: unifiedgosdk.Pointer[float64](222.61),
                        },
                        shared.DmsPartline{
                            ID: unifiedgosdk.Pointer("4be1c494-2f4d-48be-bc8c-ffddd6eafb6b"),
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
                    ID: unifiedgosdk.Pointer("f03e3ef7-410f-4292-ae2e-fc09f1ae3a6b"),
                    LaborAmount: unifiedgosdk.Pointer[float64](232.4),
                    Labors: []shared.DmsRepairorderLabor{
                        shared.DmsRepairorderLabor{
                            ID: unifiedgosdk.Pointer("fe304a1f-0f0c-4133-920f-e886548c7316"),
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
                            ID: unifiedgosdk.Pointer("88fe3de0-115d-4898-992f-20e511e7a656"),
                            PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                            Quantity: unifiedgosdk.Pointer[float64](1.0),
                            QuantityOrdered: unifiedgosdk.Pointer[float64](1.0),
                            TotalAmount: unifiedgosdk.Pointer[float64](216.02),
                            UnitPriceAmount: unifiedgosdk.Pointer[float64](216.02),
                        },
                        shared.DmsPartline{
                            ID: unifiedgosdk.Pointer("5e946f50-e229-46cf-8526-f48be45db987"),
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
            OpenedAt: types.MustNewTimeFromString("2025-02-08T22:56:07.472Z"),
            PartsAmount: unifiedgosdk.Pointer[float64](1680.8000000000002),
            PayType: shared.DmsRepairorderPayTypeWarranty.ToPointer(),
            PromisedAt: types.MustNewTimeFromString("2025-02-10T06:56:07.472Z"),
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
            UpdatedAt: types.MustNewTimeFromString("2025-11-29T11:31:28.948Z"),
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
            CompletedAt: types.MustNewTimeFromString("2026-11-14T20:19:18.722Z"),
            CreatedAt: types.MustNewTimeFromString("2022-10-25T15:50:48.007Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            EstimateAmount: unifiedgosdk.Pointer[float64](3264.0),
            ID: unifiedgosdk.Pointer("5e6e5625-3b8e-47dc-acc7-5e729a4d349a"),
            Jobs: []shared.DmsRepairorderJob{
                shared.DmsRepairorderJob{
                    ActualHours: unifiedgosdk.Pointer[float64](3.6),
                    Cause: unifiedgosdk.Pointer("Scheduled maintenance"),
                    Complaint: unifiedgosdk.Pointer("Due for service"),
                    Correction: unifiedgosdk.Pointer("Aufero cumque vicinus aptus assumenda quis aptus."),
                    ID: unifiedgosdk.Pointer("fe995c55-405f-4553-bc93-d07fc263c903"),
                    LaborAmount: unifiedgosdk.Pointer[float64](681.6),
                    Labors: []shared.DmsRepairorderLabor{
                        shared.DmsRepairorderLabor{
                            ID: unifiedgosdk.Pointer("55383436-f938-4acc-8f1d-05754cd6070d"),
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
                    ID: unifiedgosdk.Pointer("c4ed33f6-d936-4305-9dea-73914166df09"),
                    LaborAmount: unifiedgosdk.Pointer[float64](446.4),
                    Labors: []shared.DmsRepairorderLabor{
                        shared.DmsRepairorderLabor{
                            ID: unifiedgosdk.Pointer("b1efda49-94db-4bb0-89bb-1efd331e8882"),
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
                            ID: unifiedgosdk.Pointer("10ee34f6-ec4a-47da-810f-41703382e2e1"),
                            PayType: shared.DmsPartlinePayTypeWarranty.ToPointer(),
                            Quantity: unifiedgosdk.Pointer[float64](4.0),
                            QuantityOrdered: unifiedgosdk.Pointer[float64](4.0),
                            TotalAmount: unifiedgosdk.Pointer[float64](890.44),
                            UnitPriceAmount: unifiedgosdk.Pointer[float64](222.61),
                        },
                        shared.DmsPartline{
                            ID: unifiedgosdk.Pointer("4be1c494-2f4d-48be-bc8c-ffddd6eafb6b"),
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
                    ID: unifiedgosdk.Pointer("f03e3ef7-410f-4292-ae2e-fc09f1ae3a6b"),
                    LaborAmount: unifiedgosdk.Pointer[float64](232.4),
                    Labors: []shared.DmsRepairorderLabor{
                        shared.DmsRepairorderLabor{
                            ID: unifiedgosdk.Pointer("fe304a1f-0f0c-4133-920f-e886548c7316"),
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
                            ID: unifiedgosdk.Pointer("88fe3de0-115d-4898-992f-20e511e7a656"),
                            PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                            Quantity: unifiedgosdk.Pointer[float64](1.0),
                            QuantityOrdered: unifiedgosdk.Pointer[float64](1.0),
                            TotalAmount: unifiedgosdk.Pointer[float64](216.02),
                            UnitPriceAmount: unifiedgosdk.Pointer[float64](216.02),
                        },
                        shared.DmsPartline{
                            ID: unifiedgosdk.Pointer("5e946f50-e229-46cf-8526-f48be45db987"),
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
            OpenedAt: types.MustNewTimeFromString("2025-02-08T22:56:07.472Z"),
            PartsAmount: unifiedgosdk.Pointer[float64](1680.8000000000002),
            PayType: shared.DmsRepairorderPayTypeWarranty.ToPointer(),
            PromisedAt: types.MustNewTimeFromString("2025-02-10T06:56:07.472Z"),
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
            UpdatedAt: types.MustNewTimeFromString("2025-11-29T11:31:28.948Z"),
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