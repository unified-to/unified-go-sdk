# Bankaccount

## Overview

### Available Operations

* [CreateHrisBankaccount](#createhrisbankaccount) - Create a bankaccount
* [GetHrisBankaccount](#gethrisbankaccount) - Retrieve a bankaccount
* [ListHrisBankaccounts](#listhrisbankaccounts) - List all bankaccounts
* [PatchHrisBankaccount](#patchhrisbankaccount) - Update a bankaccount
* [RemoveHrisBankaccount](#removehrisbankaccount) - Remove a bankaccount
* [UpdateHrisBankaccount](#updatehrisbankaccount) - Update a bankaccount

## CreateHrisBankaccount

Creates a HrisBankaccount. Provide the unified HrisBankaccount fields in the request body; they are mapped to the provider's API. Returns the created HrisBankaccount.

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisBankaccount" method="post" path="/hris/{connection_id}/bankaccount" example="hris_bankaccount" -->
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

    res, err := s.Bankaccount.CreateHrisBankaccount(ctx, operations.CreateHrisBankaccountRequest{
        HrisBankaccount: shared.HrisBankaccount{
            AccountNumber: unifiedgosdk.Pointer("****3777"),
            AccountNumberLast4: unifiedgosdk.Pointer("3777"),
            AccountType: shared.HrisBankaccountAccountTypeChecking.ToPointer(),
            BankName: unifiedgosdk.Pointer("Huel Group"),
            CountryCode: unifiedgosdk.Pointer("US"),
            CreatedAt: types.MustNewTimeFromString("2019-11-16T16:43:45.976Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            ID: unifiedgosdk.Pointer("8dfd9e9a-787c-4b25-a1b3-a85bc78aca87"),
            IsPrimary: unifiedgosdk.Pointer(true),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("Granville Schumm - Checking"),
            Priority: unifiedgosdk.Pointer[float64](2.0),
            RoutingNumber: unifiedgosdk.Pointer("448650724"),
            SplitAmount: unifiedgosdk.Pointer[float64](98.0),
            SplitType: shared.SplitTypePercentage.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-07-05T04:16:55.985Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisBankaccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CreateHrisBankaccountRequest](../../pkg/models/operations/createhrisbankaccountrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.CreateHrisBankaccountResponse](../../pkg/models/operations/createhrisbankaccountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisBankaccount

Retrieves a single HrisBankaccount by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisBankaccount" method="get" path="/hris/{connection_id}/bankaccount/{id}" -->
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

    res, err := s.Bankaccount.GetHrisBankaccount(ctx, operations.GetHrisBankaccountRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisBankaccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetHrisBankaccountRequest](../../pkg/models/operations/gethrisbankaccountrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetHrisBankaccountResponse](../../pkg/models/operations/gethrisbankaccountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisBankaccounts

Returns a list of HrisBankaccounts. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisBankaccounts" method="get" path="/hris/{connection_id}/bankaccount" -->
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

    res, err := s.Bankaccount.ListHrisBankaccounts(ctx, operations.ListHrisBankaccountsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisBankaccounts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListHrisBankaccountsRequest](../../pkg/models/operations/listhrisbankaccountsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListHrisBankaccountsResponse](../../pkg/models/operations/listhrisbankaccountsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisBankaccount

Updates the HrisBankaccount identified by `id`. Provide the unified HrisBankaccount fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisBankaccount" method="patch" path="/hris/{connection_id}/bankaccount/{id}" example="hris_bankaccount" -->
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

    res, err := s.Bankaccount.PatchHrisBankaccount(ctx, operations.PatchHrisBankaccountRequest{
        HrisBankaccount: shared.HrisBankaccount{
            AccountNumber: unifiedgosdk.Pointer("****3777"),
            AccountNumberLast4: unifiedgosdk.Pointer("3777"),
            AccountType: shared.HrisBankaccountAccountTypeChecking.ToPointer(),
            BankName: unifiedgosdk.Pointer("Huel Group"),
            CountryCode: unifiedgosdk.Pointer("US"),
            CreatedAt: types.MustNewTimeFromString("2019-11-16T16:43:45.976Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            ID: unifiedgosdk.Pointer("69cd2c9e-29bd-4806-bf44-f41765e10a44"),
            IsPrimary: unifiedgosdk.Pointer(true),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("Granville Schumm - Checking"),
            Priority: unifiedgosdk.Pointer[float64](2.0),
            RoutingNumber: unifiedgosdk.Pointer("448650724"),
            SplitAmount: unifiedgosdk.Pointer[float64](98.0),
            SplitType: shared.SplitTypePercentage.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-07-05T04:16:55.995Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisBankaccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PatchHrisBankaccountRequest](../../pkg/models/operations/patchhrisbankaccountrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.PatchHrisBankaccountResponse](../../pkg/models/operations/patchhrisbankaccountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisBankaccount

Deletes the HrisBankaccount identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisBankaccount" method="delete" path="/hris/{connection_id}/bankaccount/{id}" -->
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

    res, err := s.Bankaccount.RemoveHrisBankaccount(ctx, operations.RemoveHrisBankaccountRequest{
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.RemoveHrisBankaccountRequest](../../pkg/models/operations/removehrisbankaccountrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.RemoveHrisBankaccountResponse](../../pkg/models/operations/removehrisbankaccountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisBankaccount

Updates the HrisBankaccount identified by `id`. Provide the unified HrisBankaccount fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisBankaccount" method="put" path="/hris/{connection_id}/bankaccount/{id}" example="hris_bankaccount" -->
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

    res, err := s.Bankaccount.UpdateHrisBankaccount(ctx, operations.UpdateHrisBankaccountRequest{
        HrisBankaccount: shared.HrisBankaccount{
            AccountNumber: unifiedgosdk.Pointer("****3777"),
            AccountNumberLast4: unifiedgosdk.Pointer("3777"),
            AccountType: shared.HrisBankaccountAccountTypeChecking.ToPointer(),
            BankName: unifiedgosdk.Pointer("Huel Group"),
            CountryCode: unifiedgosdk.Pointer("US"),
            CreatedAt: types.MustNewTimeFromString("2019-11-16T16:43:45.976Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            ID: unifiedgosdk.Pointer("69cd2c9e-29bd-4806-bf44-f41765e10a44"),
            IsPrimary: unifiedgosdk.Pointer(true),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("Granville Schumm - Checking"),
            Priority: unifiedgosdk.Pointer[float64](2.0),
            RoutingNumber: unifiedgosdk.Pointer("448650724"),
            SplitAmount: unifiedgosdk.Pointer[float64](98.0),
            SplitType: shared.SplitTypePercentage.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-07-05T04:16:55.995Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisBankaccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdateHrisBankaccountRequest](../../pkg/models/operations/updatehrisbankaccountrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UpdateHrisBankaccountResponse](../../pkg/models/operations/updatehrisbankaccountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |