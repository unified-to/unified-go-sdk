# Opcode

## Overview

### Available Operations

* [CreateDmsOpcode](#createdmsopcode) - Create an opcode
* [GetDmsOpcode](#getdmsopcode) - Retrieve an opcode
* [ListDmsOpcodes](#listdmsopcodes) - List all opcodes
* [PatchDmsOpcode](#patchdmsopcode) - Update an opcode
* [UpdateDmsOpcode](#updatedmsopcode) - Update an opcode

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

    res, err := s.Opcode.CreateDmsOpcode(ctx, operations.CreateDmsOpcodeRequest{
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

    res, err := s.Opcode.GetDmsOpcode(ctx, operations.GetDmsOpcodeRequest{
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

    res, err := s.Opcode.ListDmsOpcodes(ctx, operations.ListDmsOpcodesRequest{
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

    res, err := s.Opcode.PatchDmsOpcode(ctx, operations.PatchDmsOpcodeRequest{
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

    res, err := s.Opcode.UpdateDmsOpcode(ctx, operations.UpdateDmsOpcodeRequest{
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