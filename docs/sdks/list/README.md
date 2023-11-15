# List
(*List*)

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

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedgosdk.New()

    ctx := context.Background()
    res, err := s.List.CreateMartechList(ctx, operations.CreateMartechListRequest{
        MarketingList: &shared.MarketingList{
            Raw: &shared.PropertyMarketingListRaw{},
        },
        ConnectionID: "string",
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


### Response

**[*operations.CreateMartechListResponse](../../pkg/models/operations/createmartechlistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetMartechList

Retrieve a list

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedgosdk.New()

    ctx := context.Background()
    res, err := s.List.GetMartechList(ctx, operations.GetMartechListRequest{
        ConnectionID: "string",
        Fields: []string{
            "string",
        },
        ID: "<ID>",
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


### Response

**[*operations.GetMartechListResponse](../../pkg/models/operations/getmartechlistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ListMartechLists

List all lists

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedgosdk.New()

    ctx := context.Background()
    res, err := s.List.ListMartechLists(ctx, operations.ListMartechListsRequest{
        ConnectionID: "string",
        Fields: []string{
            "string",
        },
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


### Response

**[*operations.ListMartechListsResponse](../../pkg/models/operations/listmartechlistsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PatchMartechList

Update a list

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedgosdk.New()

    ctx := context.Background()
    res, err := s.List.PatchMartechList(ctx, operations.PatchMartechListRequest{
        MarketingList: &shared.MarketingList{
            Raw: &shared.PropertyMarketingListRaw{},
        },
        ConnectionID: "string",
        ID: "<ID>",
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


### Response

**[*operations.PatchMartechListResponse](../../pkg/models/operations/patchmartechlistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## RemoveMartechList

Remove a list

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedgosdk.New()

    ctx := context.Background()
    res, err := s.List.RemoveMartechList(ctx, operations.RemoveMartechListRequest{
        ConnectionID: "string",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.RemoveMartechListRequest](../../pkg/models/operations/removemartechlistrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.RemoveMartechListResponse](../../pkg/models/operations/removemartechlistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## UpdateMartechList

Update a list

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedgosdk.New()

    ctx := context.Background()
    res, err := s.List.UpdateMartechList(ctx, operations.UpdateMartechListRequest{
        MarketingList: &shared.MarketingList{
            Raw: &shared.PropertyMarketingListRaw{},
        },
        ConnectionID: "string",
        ID: "<ID>",
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


### Response

**[*operations.UpdateMartechListResponse](../../pkg/models/operations/updatemartechlistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
