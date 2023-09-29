# Deal
(*Deal*)

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
        ConnectionID: "Fresh",
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
        CompanyID: unifiedgosdk.String("Tools Card copying"),
        ConnectionID: "Renminbi",
        ContactID: unifiedgosdk.String("till payment World"),
        Limit: unifiedgosdk.Float64(8656.16),
        Offset: unifiedgosdk.Float64(4455.8),
        Order: unifiedgosdk.String("global"),
        Query: unifiedgosdk.String("Program Bespoke Wisconsin"),
        Sort: unifiedgosdk.String("Netherlands under"),
        UpdatedGte: types.MustTimeFromString("2022-12-23T01:47:21.816Z"),
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
        ConnectionID: "Concrete Director",
        ID: "<ID>",
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
            Amount: unifiedgosdk.Float64(7725.78),
            ClosedAt: types.MustTimeFromString("2021-10-28T08:42:49.591Z"),
            CreatedAt: types.MustTimeFromString("2023-04-23T15:03:53.999Z"),
            Currency: unifiedgosdk.String("Afghani"),
            ID: unifiedgosdk.String("<ID>"),
            LostReason: unifiedgosdk.String("North"),
            Name: unifiedgosdk.String("midnight"),
            Pipeline: unifiedgosdk.String("envisioneer Functionality Loan"),
            Probability: unifiedgosdk.Float64(7051.73),
            Raw: &shared.PropertyCrmDealRaw{},
            Source: unifiedgosdk.String("Krone"),
            Stage: unifiedgosdk.String("pascal aliquam gripping"),
            Tags: []string{
                "where",
            },
            UpdatedAt: types.MustTimeFromString("2022-04-05T10:21:22.505Z"),
            WonReason: unifiedgosdk.String("Savings kilogram"),
        },
        ConnectionID: "Chair weber silver",
        ID: "<ID>",
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
            Amount: unifiedgosdk.Float64(6144.41),
            ClosedAt: types.MustTimeFromString("2022-07-10T09:55:59.977Z"),
            CreatedAt: types.MustTimeFromString("2022-01-20T07:28:03.436Z"),
            Currency: unifiedgosdk.String("Convertible Marks"),
            ID: unifiedgosdk.String("<ID>"),
            LostReason: unifiedgosdk.String("pfft female"),
            Name: unifiedgosdk.String("Expressway"),
            Pipeline: unifiedgosdk.String("withdrawal Extended busily"),
            Probability: unifiedgosdk.Float64(7998.22),
            Raw: &shared.PropertyCrmDealRaw{},
            Source: unifiedgosdk.String("spiffy sometimes"),
            Stage: unifiedgosdk.String("transmitter"),
            Tags: []string{
                "intermediate",
            },
            UpdatedAt: types.MustTimeFromString("2022-10-06T18:34:11.762Z"),
            WonReason: unifiedgosdk.String("Cisgender input HTTP"),
        },
        ConnectionID: "accusantium Checking",
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
            Amount: unifiedgosdk.Float64(4050.98),
            ClosedAt: types.MustTimeFromString("2022-01-15T04:05:31.641Z"),
            CreatedAt: types.MustTimeFromString("2023-06-04T01:28:32.466Z"),
            Currency: unifiedgosdk.String("Bermudian Dollar (customarily known as Bermuda Dollar)"),
            ID: unifiedgosdk.String("<ID>"),
            LostReason: unifiedgosdk.String("laudantium Southwest"),
            Name: unifiedgosdk.String("wail Developer"),
            Pipeline: unifiedgosdk.String("male Samarium Gourde"),
            Probability: unifiedgosdk.Float64(6728.74),
            Raw: &shared.PropertyCrmDealRaw{},
            Source: unifiedgosdk.String("Stage Gasoline Metal"),
            Stage: unifiedgosdk.String("Corporate withdrawal Tasty"),
            Tags: []string{
                "extranet",
            },
            UpdatedAt: types.MustTimeFromString("2021-10-16T22:38:02.052Z"),
            WonReason: unifiedgosdk.String("phooey"),
        },
        ConnectionID: "Jazz",
        ID: "<ID>",
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

