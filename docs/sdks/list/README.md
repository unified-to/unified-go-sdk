# List

## Overview

### Available Operations

* [CreateMartechList](#createmartechlist) - Create a list
* [GetMartechList](#getmartechlist) - Retrieve a list
* [ListMartechLists](#listmartechlists) - List all lists
* [PatchMartechList](#patchmartechlist) - Update a list
* [RemoveMartechList](#removemartechlist) - Remove a list
* [UpdateMartechList](#updatemartechlist) - Update a list

## CreateMartechList

Create a list

### Example Usage

<!-- UsageSnippet language="go" operationID="createMartechList" method="post" path="/martech/{connection_id}/list" example="martech_list" -->
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

    res, err := s.List.CreateMartechList(ctx, operations.CreateMartechListRequest{
        MarketingList: shared.MarketingList{
            Address: &shared.PropertyMarketingListAddress{
                Address1: unifiedgosdk.Pointer("76564 Coronation Road"),
                Address2: unifiedgosdk.Pointer("Suite 922"),
                City: unifiedgosdk.Pointer("West Jordon"),
                Country: unifiedgosdk.Pointer("Mexico"),
                PostalCode: unifiedgosdk.Pointer("04183"),
                Region: unifiedgosdk.Pointer("NC"),
            },
            CreatedAt: types.MustNewTimeFromString("2019-09-18T02:01:36.950Z"),
            Description: unifiedgosdk.Pointer("Vilis."),
            ID: unifiedgosdk.Pointer("fa25e1e6-f0a8-45b9-b31f-6dfd003964d6"),
            IsActive: unifiedgosdk.Pointer(true),
            Language: unifiedgosdk.Pointer("ru"),
            Metadata: []shared.MarketingMetadata{},
            Name: unifiedgosdk.Pointer("Electronics Newsletter"),
            SenderCompany: unifiedgosdk.Pointer("Koepp and Sons"),
            SenderEmail: unifiedgosdk.Pointer("blake.kuhn@hotmail.com"),
            SenderName: unifiedgosdk.Pointer("Blake Kuhn"),
            SenderPhone: unifiedgosdk.Pointer("359.305.7962 x28115"),
            Subject: unifiedgosdk.Pointer("Theca caecus copia crebro valeo ulciscor."),
            UpdatedAt: types.MustNewTimeFromString("2022-09-13T10:08:43.347Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MarketingList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateMartechListRequest](../../pkg/models/operations/createmartechlistrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateMartechListResponse](../../pkg/models/operations/createmartechlistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetMartechList

Retrieve a list

### Example Usage

<!-- UsageSnippet language="go" operationID="getMartechList" method="get" path="/martech/{connection_id}/list/{id}" -->
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

    res, err := s.List.GetMartechList(ctx, operations.GetMartechListRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MarketingList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetMartechListRequest](../../pkg/models/operations/getmartechlistrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetMartechListResponse](../../pkg/models/operations/getmartechlistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListMartechLists

List all lists

### Example Usage

<!-- UsageSnippet language="go" operationID="listMartechLists" method="get" path="/martech/{connection_id}/list" -->
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

    res, err := s.List.ListMartechLists(ctx, operations.ListMartechListsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MarketingLists != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListMartechListsRequest](../../pkg/models/operations/listmartechlistsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListMartechListsResponse](../../pkg/models/operations/listmartechlistsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchMartechList

Update a list

### Example Usage

<!-- UsageSnippet language="go" operationID="patchMartechList" method="patch" path="/martech/{connection_id}/list/{id}" example="martech_list" -->
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

    res, err := s.List.PatchMartechList(ctx, operations.PatchMartechListRequest{
        MarketingList: shared.MarketingList{
            Address: &shared.PropertyMarketingListAddress{
                Address1: unifiedgosdk.Pointer("76564 Coronation Road"),
                Address2: unifiedgosdk.Pointer("Suite 922"),
                City: unifiedgosdk.Pointer("West Jordon"),
                Country: unifiedgosdk.Pointer("Mexico"),
                PostalCode: unifiedgosdk.Pointer("04183"),
                Region: unifiedgosdk.Pointer("NC"),
            },
            CreatedAt: types.MustNewTimeFromString("2019-09-18T02:01:36.950Z"),
            Description: unifiedgosdk.Pointer("Vilis."),
            ID: unifiedgosdk.Pointer("fbf9dae2-4ca6-4dd6-8113-05be28e2eadc"),
            IsActive: unifiedgosdk.Pointer(true),
            Language: unifiedgosdk.Pointer("ru"),
            Metadata: []shared.MarketingMetadata{},
            Name: unifiedgosdk.Pointer("Electronics Newsletter"),
            SenderCompany: unifiedgosdk.Pointer("Koepp and Sons"),
            SenderEmail: unifiedgosdk.Pointer("blake.kuhn@hotmail.com"),
            SenderName: unifiedgosdk.Pointer("Blake Kuhn"),
            SenderPhone: unifiedgosdk.Pointer("359.305.7962 x28115"),
            Subject: unifiedgosdk.Pointer("Theca caecus copia crebro valeo ulciscor."),
            UpdatedAt: types.MustNewTimeFromString("2022-09-13T10:08:43.358Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MarketingList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchMartechListRequest](../../pkg/models/operations/patchmartechlistrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchMartechListResponse](../../pkg/models/operations/patchmartechlistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveMartechList

Remove a list

### Example Usage

<!-- UsageSnippet language="go" operationID="removeMartechList" method="delete" path="/martech/{connection_id}/list/{id}" -->
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

    res, err := s.List.RemoveMartechList(ctx, operations.RemoveMartechListRequest{
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.RemoveMartechListRequest](../../pkg/models/operations/removemartechlistrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemoveMartechListResponse](../../pkg/models/operations/removemartechlistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateMartechList

Update a list

### Example Usage

<!-- UsageSnippet language="go" operationID="updateMartechList" method="put" path="/martech/{connection_id}/list/{id}" example="martech_list" -->
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

    res, err := s.List.UpdateMartechList(ctx, operations.UpdateMartechListRequest{
        MarketingList: shared.MarketingList{
            Address: &shared.PropertyMarketingListAddress{
                Address1: unifiedgosdk.Pointer("76564 Coronation Road"),
                Address2: unifiedgosdk.Pointer("Suite 922"),
                City: unifiedgosdk.Pointer("West Jordon"),
                Country: unifiedgosdk.Pointer("Mexico"),
                PostalCode: unifiedgosdk.Pointer("04183"),
                Region: unifiedgosdk.Pointer("NC"),
            },
            CreatedAt: types.MustNewTimeFromString("2019-09-18T02:01:36.950Z"),
            Description: unifiedgosdk.Pointer("Vilis."),
            ID: unifiedgosdk.Pointer("fbf9dae2-4ca6-4dd6-8113-05be28e2eadc"),
            IsActive: unifiedgosdk.Pointer(true),
            Language: unifiedgosdk.Pointer("ru"),
            Metadata: []shared.MarketingMetadata{},
            Name: unifiedgosdk.Pointer("Electronics Newsletter"),
            SenderCompany: unifiedgosdk.Pointer("Koepp and Sons"),
            SenderEmail: unifiedgosdk.Pointer("blake.kuhn@hotmail.com"),
            SenderName: unifiedgosdk.Pointer("Blake Kuhn"),
            SenderPhone: unifiedgosdk.Pointer("359.305.7962 x28115"),
            Subject: unifiedgosdk.Pointer("Theca caecus copia crebro valeo ulciscor."),
            UpdatedAt: types.MustNewTimeFromString("2022-09-13T10:08:43.358Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MarketingList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateMartechListRequest](../../pkg/models/operations/updatemartechlistrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateMartechListResponse](../../pkg/models/operations/updatemartechlistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |