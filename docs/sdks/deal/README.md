# Deal

### Available Operations

* [DeleteCrmConnectionIDDealID](#deletecrmconnectioniddealid) - Remove a deal
* [GetCrmConnectionIDDeal](#getcrmconnectioniddeal) - List all deals
* [GetCrmConnectionIDDealID](#getcrmconnectioniddealid) - Retrieve a deal
* [PatchCrmConnectionIDDealID](#patchcrmconnectioniddealid) - Update a deal
* [PostCrmConnectionIDDeal](#postcrmconnectioniddeal) - Create a deal
* [PutCrmConnectionIDDealID](#putcrmconnectioniddealid) - Update a deal

## DeleteCrmConnectionIDDealID

Remove a deal

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
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Deal.DeleteCrmConnectionIDDealID(ctx, operations.DeleteCrmConnectionIDDealIDRequest{
        ConnectionID: "ducimus",
        ID: "ed565076-21c5-48f4-9739-6564c20a0711",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.DeleteCrmConnectionIDDealIDRequest](../../models/operations/deletecrmconnectioniddealidrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.DeleteCrmConnectionIDDealIDResponse](../../models/operations/deletecrmconnectioniddealidresponse.md), error**


## GetCrmConnectionIDDeal

List all deals

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Deal.GetCrmConnectionIDDeal(ctx, operations.GetCrmConnectionIDDealRequest{
        CompanyID: unifiedgosdk.String("similique"),
        ConnectionID: "omnis",
        ContactID: unifiedgosdk.String("commodi"),
        Limit: unifiedgosdk.Float64(1166.35),
        Offset: unifiedgosdk.Float64(8489.26),
        Order: unifiedgosdk.String("aspernatur"),
        Query: unifiedgosdk.String("ut"),
        Sort: unifiedgosdk.String("deserunt"),
        UpdatedGte: types.MustTimeFromString("2022-02-20T22:48:15.284Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmDeals != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetCrmConnectionIDDealRequest](../../models/operations/getcrmconnectioniddealrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.GetCrmConnectionIDDealResponse](../../models/operations/getcrmconnectioniddealresponse.md), error**


## GetCrmConnectionIDDealID

Retrieve a deal

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
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Deal.GetCrmConnectionIDDealID(ctx, operations.GetCrmConnectionIDDealIDRequest{
        ConnectionID: "facilis",
        ID: "b8f532d8-92cf-4781-acb5-12c878240bf5",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmDeal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetCrmConnectionIDDealIDRequest](../../models/operations/getcrmconnectioniddealidrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.GetCrmConnectionIDDealIDResponse](../../models/operations/getcrmconnectioniddealidresponse.md), error**


## PatchCrmConnectionIDDealID

Update a deal

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Deal.PatchCrmConnectionIDDealID(ctx, operations.PatchCrmConnectionIDDealIDRequest{
        CrmDeal: &shared.CrmDeal{
            Amount: unifiedgosdk.Float64(2743.68),
            ClosedAt: types.MustTimeFromString("2021-02-04T20:36:14.764Z"),
            CreatedAt: types.MustTimeFromString("2021-12-30T15:49:38.515Z"),
            Currency: unifiedgosdk.String("hic"),
            ID: unifiedgosdk.String("8f1bf0bc-8e1f-4206-95d8-31d0081090f6"),
            LostReason: unifiedgosdk.String("nihil"),
            Name: unifiedgosdk.String("Loretta Howe"),
            Pipeline: unifiedgosdk.String("doloribus"),
            Probability: unifiedgosdk.Float64(1877.7),
            Raw: &shared.PropertyCrmDealRaw{},
            Source: unifiedgosdk.String("id"),
            Stage: unifiedgosdk.String("ex"),
            Tags: []string{
                "quos",
            },
            UpdatedAt: types.MustTimeFromString("2022-03-17T20:43:59.276Z"),
            WonReason: unifiedgosdk.String("exercitationem"),
        },
        ConnectionID: "molestiae",
        ID: "68dce742-409a-4215-a086-01489a5f63e3",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmDeal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PatchCrmConnectionIDDealIDRequest](../../models/operations/patchcrmconnectioniddealidrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.PatchCrmConnectionIDDealIDResponse](../../models/operations/patchcrmconnectioniddealidresponse.md), error**


## PostCrmConnectionIDDeal

Create a deal

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Deal.PostCrmConnectionIDDeal(ctx, operations.PostCrmConnectionIDDealRequest{
        CrmDeal: &shared.CrmDeal{
            Amount: unifiedgosdk.Float64(6489.85),
            ClosedAt: types.MustTimeFromString("2022-04-04T11:17:39.742Z"),
            CreatedAt: types.MustTimeFromString("2020-05-19T07:26:52.477Z"),
            Currency: unifiedgosdk.String("natus"),
            ID: unifiedgosdk.String("dda33dcd-6348-43e4-a7a9-8e4df37e45b8"),
            LostReason: unifiedgosdk.String("omnis"),
            Name: unifiedgosdk.String("Bernice Schultz I"),
            Pipeline: unifiedgosdk.String("recusandae"),
            Probability: unifiedgosdk.Float64(784.86),
            Raw: &shared.PropertyCrmDealRaw{},
            Source: unifiedgosdk.String("ipsum"),
            Stage: unifiedgosdk.String("error"),
            Tags: []string{
                "numquam",
            },
            UpdatedAt: types.MustTimeFromString("2022-08-25T17:34:42.796Z"),
            WonReason: unifiedgosdk.String("consectetur"),
        },
        ConnectionID: "dicta",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmDeal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.PostCrmConnectionIDDealRequest](../../models/operations/postcrmconnectioniddealrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.PostCrmConnectionIDDealResponse](../../models/operations/postcrmconnectioniddealresponse.md), error**


## PutCrmConnectionIDDealID

Update a deal

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Deal.PutCrmConnectionIDDealID(ctx, operations.PutCrmConnectionIDDealIDRequest{
        CrmDeal: &shared.CrmDeal{
            Amount: unifiedgosdk.Float64(551.24),
            ClosedAt: types.MustTimeFromString("2022-12-18T04:51:10.637Z"),
            CreatedAt: types.MustTimeFromString("2022-04-06T10:46:32.109Z"),
            Currency: unifiedgosdk.String("facere"),
            ID: unifiedgosdk.String("354c092b-d734-4f02-849d-86f4bb20fe5d"),
            LostReason: unifiedgosdk.String("provident"),
            Name: unifiedgosdk.String("Ashley Schmeler"),
            Pipeline: unifiedgosdk.String("itaque"),
            Probability: unifiedgosdk.Float64(4920.7),
            Raw: &shared.PropertyCrmDealRaw{},
            Source: unifiedgosdk.String("magnam"),
            Stage: unifiedgosdk.String("excepturi"),
            Tags: []string{
                "placeat",
            },
            UpdatedAt: types.MustTimeFromString("2021-01-19T07:13:22.442Z"),
            WonReason: unifiedgosdk.String("modi"),
        },
        ConnectionID: "enim",
        ID: "a27f69e2-c9e6-4d10-a9db-3ad4c6b03108",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmDeal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PutCrmConnectionIDDealIDRequest](../../models/operations/putcrmconnectioniddealidrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.PutCrmConnectionIDDealIDResponse](../../models/operations/putcrmconnectioniddealidresponse.md), error**

