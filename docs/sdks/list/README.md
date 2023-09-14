# List

### Available Operations

* [DeleteMartechConnectionIDListID](#deletemartechconnectionidlistid) - Remove a list
* [GetMartechConnectionIDList](#getmartechconnectionidlist) - List all lists
* [GetMartechConnectionIDListID](#getmartechconnectionidlistid) - Retrieve a list
* [PatchMartechConnectionIDListID](#patchmartechconnectionidlistid) - Update a list
* [PostMartechConnectionIDList](#postmartechconnectionidlist) - Create a list
* [PutMartechConnectionIDListID](#putmartechconnectionidlistid) - Update a list

## DeleteMartechConnectionIDListID

Remove a list

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.DeleteMartechConnectionIDListIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.List.DeleteMartechConnectionIDListID(ctx, operations.DeleteMartechConnectionIDListIDRequest{
        ConnectionID: "molestiae",
        ID: "866db8a7-49e3-4984-911c-c75e4f0c004b",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.DeleteMartechConnectionIDListIDRequest](../../models/operations/deletemartechconnectionidlistidrequest.md)   | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `security`                                                                                                               | [operations.DeleteMartechConnectionIDListIDSecurity](../../models/operations/deletemartechconnectionidlistidsecurity.md) | :heavy_check_mark:                                                                                                       | The security requirements to use for the request.                                                                        |


### Response

**[*operations.DeleteMartechConnectionIDListIDResponse](../../models/operations/deletemartechconnectionidlistidresponse.md), error**


## GetMartechConnectionIDList

List all lists

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetMartechConnectionIDListSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.List.GetMartechConnectionIDList(ctx, operations.GetMartechConnectionIDListRequest{
        ConnectionID: "minima",
        Limit: unifiedto.Float64(7315.15),
        Offset: unifiedto.Float64(6991.28),
        Order: unifiedto.String("molestiae"),
        Query: unifiedto.String("ipsam"),
        Sort: unifiedto.String("quos"),
        UpdatedGte: types.MustTimeFromString("2020-09-30T08:18:10.798Z"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.MarketingLists != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetMartechConnectionIDListRequest](../../models/operations/getmartechconnectionidlistrequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.GetMartechConnectionIDListSecurity](../../models/operations/getmartechconnectionidlistsecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


### Response

**[*operations.GetMartechConnectionIDListResponse](../../models/operations/getmartechconnectionidlistresponse.md), error**


## GetMartechConnectionIDListID

Retrieve a list

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetMartechConnectionIDListIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.List.GetMartechConnectionIDListID(ctx, operations.GetMartechConnectionIDListIDRequest{
        ConnectionID: "cupiditate",
        ID: "4562f006-9685-4fcd-9a17-3d84bbe24f29",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.MarketingList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetMartechConnectionIDListIDRequest](../../models/operations/getmartechconnectionidlistidrequest.md)   | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `security`                                                                                                         | [operations.GetMartechConnectionIDListIDSecurity](../../models/operations/getmartechconnectionidlistidsecurity.md) | :heavy_check_mark:                                                                                                 | The security requirements to use for the request.                                                                  |


### Response

**[*operations.GetMartechConnectionIDListIDResponse](../../models/operations/getmartechconnectionidlistidresponse.md), error**


## PatchMartechConnectionIDListID

Update a list

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PatchMartechConnectionIDListIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.List.PatchMartechConnectionIDListID(ctx, operations.PatchMartechConnectionIDListIDRequest{
        MarketingList: &shared.MarketingList{
            CreatedAt: unifiedto.String("praesentium"),
            ID: unifiedto.String("34afb073-5cb6-4285-94a2-9aaa1e169156"),
            Name: unifiedto.String("Adrian Schuster"),
            Raw: &shared.PropertyMarketingListRaw{},
            UpdatedAt: types.MustTimeFromString("2022-06-17T09:25:28.057Z"),
        },
        ConnectionID: "perferendis",
        ID: "9505bf03-a93e-4944-80ca-37fb10789032",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.MarketingList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.PatchMartechConnectionIDListIDRequest](../../models/operations/patchmartechconnectionidlistidrequest.md)   | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `security`                                                                                                             | [operations.PatchMartechConnectionIDListIDSecurity](../../models/operations/patchmartechconnectionidlistidsecurity.md) | :heavy_check_mark:                                                                                                     | The security requirements to use for the request.                                                                      |


### Response

**[*operations.PatchMartechConnectionIDListIDResponse](../../models/operations/patchmartechconnectionidlistidresponse.md), error**


## PostMartechConnectionIDList

Create a list

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PostMartechConnectionIDListSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.List.PostMartechConnectionIDList(ctx, operations.PostMartechConnectionIDListRequest{
        MarketingList: &shared.MarketingList{
            CreatedAt: unifiedto.String("deserunt"),
            ID: unifiedto.String("c333172e-2dd7-49ec-b4ba-7e88ddb36fd1"),
            Name: unifiedto.String("Lucas Schneider"),
            Raw: &shared.PropertyMarketingListRaw{},
            UpdatedAt: types.MustTimeFromString("2022-03-20T02:47:07.461Z"),
        },
        ConnectionID: "quas",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.MarketingList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.PostMartechConnectionIDListRequest](../../models/operations/postmartechconnectionidlistrequest.md)   | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `security`                                                                                                       | [operations.PostMartechConnectionIDListSecurity](../../models/operations/postmartechconnectionidlistsecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |


### Response

**[*operations.PostMartechConnectionIDListResponse](../../models/operations/postmartechconnectionidlistresponse.md), error**


## PutMartechConnectionIDListID

Update a list

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PutMartechConnectionIDListIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.List.PutMartechConnectionIDListID(ctx, operations.PutMartechConnectionIDListIDRequest{
        MarketingList: &shared.MarketingList{
            CreatedAt: unifiedto.String("autem"),
            ID: unifiedto.String("573474f0-a740-4fb4-ab44-1c3a09e76399"),
            Name: unifiedto.String("Ms. Eula Leffler"),
            Raw: &shared.PropertyMarketingListRaw{},
            UpdatedAt: types.MustTimeFromString("2021-03-29T15:45:25.588Z"),
        },
        ConnectionID: "odio",
        ID: "94455ebc-550a-41c4-a6b5-9c8366fdcc13",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.MarketingList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.PutMartechConnectionIDListIDRequest](../../models/operations/putmartechconnectionidlistidrequest.md)   | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `security`                                                                                                         | [operations.PutMartechConnectionIDListIDSecurity](../../models/operations/putmartechconnectionidlistidsecurity.md) | :heavy_check_mark:                                                                                                 | The security requirements to use for the request.                                                                  |


### Response

**[*operations.PutMartechConnectionIDListIDResponse](../../models/operations/putmartechconnectionidlistidresponse.md), error**

