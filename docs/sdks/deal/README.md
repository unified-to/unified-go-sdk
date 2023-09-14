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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.DeleteCrmConnectionIDDealIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Deal.DeleteCrmConnectionIDDealID(ctx, operations.DeleteCrmConnectionIDDealIDRequest{
        ConnectionID: "ducimus",
        ID: "ed565076-21c5-48f4-9739-6564c20a0711",
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.DeleteCrmConnectionIDDealIDRequest](../../models/operations/deletecrmconnectioniddealidrequest.md)   | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `security`                                                                                                       | [operations.DeleteCrmConnectionIDDealIDSecurity](../../models/operations/deletecrmconnectioniddealidsecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetCrmConnectionIDDealSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Deal.GetCrmConnectionIDDeal(ctx, operations.GetCrmConnectionIDDealRequest{
        CompanyID: unifiedto.String("similique"),
        ConnectionID: "omnis",
        ContactID: unifiedto.String("commodi"),
        Limit: unifiedto.Float64(1166.35),
        Offset: unifiedto.Float64(8489.26),
        Order: unifiedto.String("aspernatur"),
        Query: unifiedto.String("ut"),
        Sort: unifiedto.String("deserunt"),
        UpdatedGte: types.MustTimeFromString("2022-02-20T22:48:15.284Z"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmDeals != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetCrmConnectionIDDealRequest](../../models/operations/getcrmconnectioniddealrequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.GetCrmConnectionIDDealSecurity](../../models/operations/getcrmconnectioniddealsecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetCrmConnectionIDDealIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Deal.GetCrmConnectionIDDealID(ctx, operations.GetCrmConnectionIDDealIDRequest{
        ConnectionID: "facilis",
        ID: "b8f532d8-92cf-4781-acb5-12c878240bf5",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmDeal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetCrmConnectionIDDealIDRequest](../../models/operations/getcrmconnectioniddealidrequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.GetCrmConnectionIDDealIDSecurity](../../models/operations/getcrmconnectioniddealidsecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PatchCrmConnectionIDDealIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Deal.PatchCrmConnectionIDDealID(ctx, operations.PatchCrmConnectionIDDealIDRequest{
        CrmDeal: &shared.CrmDeal{
            Amount: unifiedto.Float64(2743.68),
            ClosedAt: types.MustTimeFromString("2021-02-04T20:36:14.764Z"),
            CreatedAt: types.MustTimeFromString("2021-12-30T15:49:38.515Z"),
            Currency: unifiedto.String("hic"),
            ID: unifiedto.String("8f1bf0bc-8e1f-4206-95d8-31d0081090f6"),
            LostReason: unifiedto.String("nihil"),
            Name: unifiedto.String("Loretta Howe"),
            Pipeline: unifiedto.String("doloribus"),
            Probability: unifiedto.Float64(1877.7),
            Raw: &shared.PropertyCrmDealRaw{},
            Source: unifiedto.String("id"),
            Stage: unifiedto.String("ex"),
            Tags: []string{
                "quos",
            },
            UpdatedAt: types.MustTimeFromString("2022-03-17T20:43:59.276Z"),
            WonReason: unifiedto.String("exercitationem"),
        },
        ConnectionID: "molestiae",
        ID: "68dce742-409a-4215-a086-01489a5f63e3",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmDeal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PatchCrmConnectionIDDealIDRequest](../../models/operations/patchcrmconnectioniddealidrequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.PatchCrmConnectionIDDealIDSecurity](../../models/operations/patchcrmconnectioniddealidsecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PostCrmConnectionIDDealSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Deal.PostCrmConnectionIDDeal(ctx, operations.PostCrmConnectionIDDealRequest{
        CrmDeal: &shared.CrmDeal{
            Amount: unifiedto.Float64(6489.85),
            ClosedAt: types.MustTimeFromString("2022-04-04T11:17:39.742Z"),
            CreatedAt: types.MustTimeFromString("2020-05-19T07:26:52.477Z"),
            Currency: unifiedto.String("natus"),
            ID: unifiedto.String("dda33dcd-6348-43e4-a7a9-8e4df37e45b8"),
            LostReason: unifiedto.String("omnis"),
            Name: unifiedto.String("Bernice Schultz I"),
            Pipeline: unifiedto.String("recusandae"),
            Probability: unifiedto.Float64(784.86),
            Raw: &shared.PropertyCrmDealRaw{},
            Source: unifiedto.String("ipsum"),
            Stage: unifiedto.String("error"),
            Tags: []string{
                "numquam",
            },
            UpdatedAt: types.MustTimeFromString("2022-08-25T17:34:42.796Z"),
            WonReason: unifiedto.String("consectetur"),
        },
        ConnectionID: "dicta",
    }, operationSecurity)
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
| `request`                                                                                                | [operations.PostCrmConnectionIDDealRequest](../../models/operations/postcrmconnectioniddealrequest.md)   | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.PostCrmConnectionIDDealSecurity](../../models/operations/postcrmconnectioniddealsecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PutCrmConnectionIDDealIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Deal.PutCrmConnectionIDDealID(ctx, operations.PutCrmConnectionIDDealIDRequest{
        CrmDeal: &shared.CrmDeal{
            Amount: unifiedto.Float64(551.24),
            ClosedAt: types.MustTimeFromString("2022-12-18T04:51:10.637Z"),
            CreatedAt: types.MustTimeFromString("2022-04-06T10:46:32.109Z"),
            Currency: unifiedto.String("facere"),
            ID: unifiedto.String("354c092b-d734-4f02-849d-86f4bb20fe5d"),
            LostReason: unifiedto.String("provident"),
            Name: unifiedto.String("Ashley Schmeler"),
            Pipeline: unifiedto.String("itaque"),
            Probability: unifiedto.Float64(4920.7),
            Raw: &shared.PropertyCrmDealRaw{},
            Source: unifiedto.String("magnam"),
            Stage: unifiedto.String("excepturi"),
            Tags: []string{
                "placeat",
            },
            UpdatedAt: types.MustTimeFromString("2021-01-19T07:13:22.442Z"),
            WonReason: unifiedto.String("modi"),
        },
        ConnectionID: "enim",
        ID: "a27f69e2-c9e6-4d10-a9db-3ad4c6b03108",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmDeal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.PutCrmConnectionIDDealIDRequest](../../models/operations/putcrmconnectioniddealidrequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.PutCrmConnectionIDDealIDSecurity](../../models/operations/putcrmconnectioniddealidsecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |


### Response

**[*operations.PutCrmConnectionIDDealIDResponse](../../models/operations/putcrmconnectioniddealidresponse.md), error**

