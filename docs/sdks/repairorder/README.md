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
            CompletedAt: types.MustNewTimeFromString("2026-11-12T19:17:42.253Z"),
            CreatedAt: types.MustNewTimeFromString("2022-10-25T15:50:48.007Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            EstimateAmount: unifiedgosdk.Pointer[float64](3264.0),
            ID: unifiedgosdk.Pointer("7fb2dd80-bcfa-4c94-8c60-b2d2b9a4343a"),
            Jobs: []shared.DmsRepairorderJob{
                shared.DmsRepairorderJob{
                    ActualHours: unifiedgosdk.Pointer[float64](3.6),
                    Cause: unifiedgosdk.Pointer("Scheduled maintenance"),
                    Complaint: unifiedgosdk.Pointer("Due for service"),
                    Correction: unifiedgosdk.Pointer("Aufero cumque vicinus aptus assumenda quis aptus."),
                    ID: unifiedgosdk.Pointer("a3e10710-eaf3-42c0-a284-3b2bb7a069e7"),
                    LaborAmount: unifiedgosdk.Pointer[float64](681.6),
                    Labors: []shared.DmsRepairorderLabor{
                        shared.DmsRepairorderLabor{
                            ID: unifiedgosdk.Pointer("9dfd5a09-1f6c-421b-b16c-aed61a5c14e6"),
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
                    ID: unifiedgosdk.Pointer("45a15ebd-8869-4c49-9e59-c8ab00015b01"),
                    LaborAmount: unifiedgosdk.Pointer[float64](446.4),
                    Labors: []shared.DmsRepairorderLabor{
                        shared.DmsRepairorderLabor{
                            ID: unifiedgosdk.Pointer("d18cb9c4-81bd-4cd5-bf32-eff004254921"),
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
                            ID: unifiedgosdk.Pointer("a9f379cc-9f18-4ac2-9723-7fda4061e45e"),
                            PayType: shared.DmsPartlinePayTypeWarranty.ToPointer(),
                            Quantity: unifiedgosdk.Pointer[float64](4.0),
                            QuantityOrdered: unifiedgosdk.Pointer[float64](4.0),
                            TotalAmount: unifiedgosdk.Pointer[float64](890.44),
                            UnitPriceAmount: unifiedgosdk.Pointer[float64](222.61),
                        },
                        shared.DmsPartline{
                            ID: unifiedgosdk.Pointer("884c9571-00e2-402a-9c24-6522d26444c3"),
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
                    ID: unifiedgosdk.Pointer("e20d5136-0721-4308-84ca-8e748a8afbe4"),
                    LaborAmount: unifiedgosdk.Pointer[float64](232.4),
                    Labors: []shared.DmsRepairorderLabor{
                        shared.DmsRepairorderLabor{
                            ID: unifiedgosdk.Pointer("18657ec3-723a-4418-9319-fdb6e7cfd8e9"),
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
                            ID: unifiedgosdk.Pointer("b26e6eaa-eadf-444f-b4be-3438e950d554"),
                            PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                            Quantity: unifiedgosdk.Pointer[float64](1.0),
                            QuantityOrdered: unifiedgosdk.Pointer[float64](1.0),
                            TotalAmount: unifiedgosdk.Pointer[float64](216.02),
                            UnitPriceAmount: unifiedgosdk.Pointer[float64](216.02),
                        },
                        shared.DmsPartline{
                            ID: unifiedgosdk.Pointer("775ba6c5-dc39-4dfd-9c5e-f2139c2f652f"),
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
            OpenedAt: types.MustNewTimeFromString("2025-02-07T19:13:16.323Z"),
            PartsAmount: unifiedgosdk.Pointer[float64](1680.8000000000002),
            PayType: shared.DmsRepairorderPayTypeWarranty.ToPointer(),
            PromisedAt: types.MustNewTimeFromString("2025-02-09T03:13:16.323Z"),
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
            UpdatedAt: types.MustNewTimeFromString("2025-11-27T22:05:41.752Z"),
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
            CompletedAt: types.MustNewTimeFromString("2026-11-12T19:17:42.307Z"),
            CreatedAt: types.MustNewTimeFromString("2022-10-25T15:50:48.007Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            EstimateAmount: unifiedgosdk.Pointer[float64](3264.0),
            ID: unifiedgosdk.Pointer("205d65d6-f300-436e-a728-71ed9ac8cb5c"),
            Jobs: []shared.DmsRepairorderJob{
                shared.DmsRepairorderJob{
                    ActualHours: unifiedgosdk.Pointer[float64](3.6),
                    Cause: unifiedgosdk.Pointer("Scheduled maintenance"),
                    Complaint: unifiedgosdk.Pointer("Due for service"),
                    Correction: unifiedgosdk.Pointer("Aufero cumque vicinus aptus assumenda quis aptus."),
                    ID: unifiedgosdk.Pointer("426874d7-2e52-4162-8b6c-8639c58d18b5"),
                    LaborAmount: unifiedgosdk.Pointer[float64](681.6),
                    Labors: []shared.DmsRepairorderLabor{
                        shared.DmsRepairorderLabor{
                            ID: unifiedgosdk.Pointer("e1009e6c-2d8d-446c-8269-bcc66c32cd67"),
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
                    ID: unifiedgosdk.Pointer("7a124d15-1bb6-491d-adc1-579be2c00466"),
                    LaborAmount: unifiedgosdk.Pointer[float64](446.4),
                    Labors: []shared.DmsRepairorderLabor{
                        shared.DmsRepairorderLabor{
                            ID: unifiedgosdk.Pointer("bf400c19-1912-4255-b38c-bdc5cb87abf6"),
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
                            ID: unifiedgosdk.Pointer("5d8d7099-cdda-4146-b08c-46d84fba20a8"),
                            PayType: shared.DmsPartlinePayTypeWarranty.ToPointer(),
                            Quantity: unifiedgosdk.Pointer[float64](4.0),
                            QuantityOrdered: unifiedgosdk.Pointer[float64](4.0),
                            TotalAmount: unifiedgosdk.Pointer[float64](890.44),
                            UnitPriceAmount: unifiedgosdk.Pointer[float64](222.61),
                        },
                        shared.DmsPartline{
                            ID: unifiedgosdk.Pointer("9a7d805a-b9d9-4663-86e4-0e9b32dba051"),
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
                    ID: unifiedgosdk.Pointer("8fe8640a-c7a2-4ad5-86ef-a016db57c58a"),
                    LaborAmount: unifiedgosdk.Pointer[float64](232.4),
                    Labors: []shared.DmsRepairorderLabor{
                        shared.DmsRepairorderLabor{
                            ID: unifiedgosdk.Pointer("1e1d8015-f3f9-4d90-bede-5f2b99783e47"),
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
                            ID: unifiedgosdk.Pointer("9237d282-5b87-4e09-bb1f-363fc2a6296e"),
                            PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                            Quantity: unifiedgosdk.Pointer[float64](1.0),
                            QuantityOrdered: unifiedgosdk.Pointer[float64](1.0),
                            TotalAmount: unifiedgosdk.Pointer[float64](216.02),
                            UnitPriceAmount: unifiedgosdk.Pointer[float64](216.02),
                        },
                        shared.DmsPartline{
                            ID: unifiedgosdk.Pointer("884766ed-b7f1-4902-9a06-1a0837b87653"),
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
            OpenedAt: types.MustNewTimeFromString("2025-02-07T19:13:16.353Z"),
            PartsAmount: unifiedgosdk.Pointer[float64](1680.8000000000002),
            PayType: shared.DmsRepairorderPayTypeWarranty.ToPointer(),
            PromisedAt: types.MustNewTimeFromString("2025-02-09T03:13:16.353Z"),
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
            UpdatedAt: types.MustNewTimeFromString("2025-11-27T22:05:41.793Z"),
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
            CompletedAt: types.MustNewTimeFromString("2026-11-12T19:17:42.307Z"),
            CreatedAt: types.MustNewTimeFromString("2022-10-25T15:50:48.007Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            EstimateAmount: unifiedgosdk.Pointer[float64](3264.0),
            ID: unifiedgosdk.Pointer("205d65d6-f300-436e-a728-71ed9ac8cb5c"),
            Jobs: []shared.DmsRepairorderJob{
                shared.DmsRepairorderJob{
                    ActualHours: unifiedgosdk.Pointer[float64](3.6),
                    Cause: unifiedgosdk.Pointer("Scheduled maintenance"),
                    Complaint: unifiedgosdk.Pointer("Due for service"),
                    Correction: unifiedgosdk.Pointer("Aufero cumque vicinus aptus assumenda quis aptus."),
                    ID: unifiedgosdk.Pointer("426874d7-2e52-4162-8b6c-8639c58d18b5"),
                    LaborAmount: unifiedgosdk.Pointer[float64](681.6),
                    Labors: []shared.DmsRepairorderLabor{
                        shared.DmsRepairorderLabor{
                            ID: unifiedgosdk.Pointer("e1009e6c-2d8d-446c-8269-bcc66c32cd67"),
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
                    ID: unifiedgosdk.Pointer("7a124d15-1bb6-491d-adc1-579be2c00466"),
                    LaborAmount: unifiedgosdk.Pointer[float64](446.4),
                    Labors: []shared.DmsRepairorderLabor{
                        shared.DmsRepairorderLabor{
                            ID: unifiedgosdk.Pointer("bf400c19-1912-4255-b38c-bdc5cb87abf6"),
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
                            ID: unifiedgosdk.Pointer("5d8d7099-cdda-4146-b08c-46d84fba20a8"),
                            PayType: shared.DmsPartlinePayTypeWarranty.ToPointer(),
                            Quantity: unifiedgosdk.Pointer[float64](4.0),
                            QuantityOrdered: unifiedgosdk.Pointer[float64](4.0),
                            TotalAmount: unifiedgosdk.Pointer[float64](890.44),
                            UnitPriceAmount: unifiedgosdk.Pointer[float64](222.61),
                        },
                        shared.DmsPartline{
                            ID: unifiedgosdk.Pointer("9a7d805a-b9d9-4663-86e4-0e9b32dba051"),
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
                    ID: unifiedgosdk.Pointer("8fe8640a-c7a2-4ad5-86ef-a016db57c58a"),
                    LaborAmount: unifiedgosdk.Pointer[float64](232.4),
                    Labors: []shared.DmsRepairorderLabor{
                        shared.DmsRepairorderLabor{
                            ID: unifiedgosdk.Pointer("1e1d8015-f3f9-4d90-bede-5f2b99783e47"),
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
                            ID: unifiedgosdk.Pointer("9237d282-5b87-4e09-bb1f-363fc2a6296e"),
                            PayType: shared.DmsPartlinePayTypeCustomer.ToPointer(),
                            Quantity: unifiedgosdk.Pointer[float64](1.0),
                            QuantityOrdered: unifiedgosdk.Pointer[float64](1.0),
                            TotalAmount: unifiedgosdk.Pointer[float64](216.02),
                            UnitPriceAmount: unifiedgosdk.Pointer[float64](216.02),
                        },
                        shared.DmsPartline{
                            ID: unifiedgosdk.Pointer("884766ed-b7f1-4902-9a06-1a0837b87653"),
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
            OpenedAt: types.MustNewTimeFromString("2025-02-07T19:13:16.353Z"),
            PartsAmount: unifiedgosdk.Pointer[float64](1680.8000000000002),
            PayType: shared.DmsRepairorderPayTypeWarranty.ToPointer(),
            PromisedAt: types.MustNewTimeFromString("2025-02-09T03:13:16.353Z"),
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
            UpdatedAt: types.MustNewTimeFromString("2025-11-27T22:05:41.793Z"),
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