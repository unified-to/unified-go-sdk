# Group
(*Group*)

### Available Operations

* [DeleteHrisConnectionIDGroupID](#deletehrisconnectionidgroupid) - Remove a Group
* [GetHrisConnectionIDGroup](#gethrisconnectionidgroup) - List all Groups
* [GetHrisConnectionIDGroupID](#gethrisconnectionidgroupid) - Retrieve a Group
* [PatchHrisConnectionIDGroupID](#patchhrisconnectionidgroupid) - Update a Group
* [PostHrisConnectionIDGroup](#posthrisconnectionidgroup) - Create a Group
* [PutHrisConnectionIDGroupID](#puthrisconnectionidgroupid) - Update a Group

## DeleteHrisConnectionIDGroupID

Remove a Group

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
    res, err := s.Group.DeleteHrisConnectionIDGroupID(ctx, operations.DeleteHrisConnectionIDGroupIDRequest{
        ConnectionID: "consequently platforms Metal",
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.DeleteHrisConnectionIDGroupIDRequest](../../models/operations/deletehrisconnectionidgroupidrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.DeleteHrisConnectionIDGroupIDResponse](../../models/operations/deletehrisconnectionidgroupidresponse.md), error**


## GetHrisConnectionIDGroup

List all Groups

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
    res, err := s.Group.GetHrisConnectionIDGroup(ctx, operations.GetHrisConnectionIDGroupRequest{
        ConnectionID: "Loan",
        Limit: unifiedgosdk.Float64(3486.96),
        Offset: unifiedgosdk.Float64(9705.73),
        Order: unifiedgosdk.String("Coordinator"),
        Query: unifiedgosdk.String("World"),
        Sort: unifiedgosdk.String("Dollar"),
        UpdatedGte: types.MustTimeFromString("2021-01-15T16:06:13.340Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.HrisGroups != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetHrisConnectionIDGroupRequest](../../models/operations/gethrisconnectionidgrouprequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.GetHrisConnectionIDGroupResponse](../../models/operations/gethrisconnectionidgroupresponse.md), error**


## GetHrisConnectionIDGroupID

Retrieve a Group

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
    res, err := s.Group.GetHrisConnectionIDGroupID(ctx, operations.GetHrisConnectionIDGroupIDRequest{
        ConnectionID: "behind",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.HrisGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetHrisConnectionIDGroupIDRequest](../../models/operations/gethrisconnectionidgroupidrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.GetHrisConnectionIDGroupIDResponse](../../models/operations/gethrisconnectionidgroupidresponse.md), error**


## PatchHrisConnectionIDGroupID

Update a Group

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
    res, err := s.Group.PatchHrisConnectionIDGroupID(ctx, operations.PatchHrisConnectionIDGroupIDRequest{
        HrisGroup: &shared.HrisGroup{
            CreatedAt: types.MustTimeFromString("2023-10-19T05:30:26.390Z"),
            Description: unifiedgosdk.String("Stand-alone asymmetric orchestration"),
            EmployeeIds: []string{
                "shootdown",
            },
            ID: unifiedgosdk.String("<ID>"),
            IsActive: unifiedgosdk.Bool(false),
            ManagerIds: []string{
                "24/7",
            },
            Name: unifiedgosdk.String("Agender trainer"),
            ParentID: unifiedgosdk.String("Configuration Kids Sedan"),
            Raw: &shared.PropertyHrisGroupRaw{},
            Type: shared.HrisGroupTypeDivision.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2021-08-18T16:48:12.885Z"),
        },
        ConnectionID: "Intersex",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.HrisGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.PatchHrisConnectionIDGroupIDRequest](../../models/operations/patchhrisconnectionidgroupidrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.PatchHrisConnectionIDGroupIDResponse](../../models/operations/patchhrisconnectionidgroupidresponse.md), error**


## PostHrisConnectionIDGroup

Create a Group

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
    res, err := s.Group.PostHrisConnectionIDGroup(ctx, operations.PostHrisConnectionIDGroupRequest{
        HrisGroup: &shared.HrisGroup{
            CreatedAt: types.MustTimeFromString("2021-02-23T15:35:38.483Z"),
            Description: unifiedgosdk.String("Configurable stable product"),
            EmployeeIds: []string{
                "Auto",
            },
            ID: unifiedgosdk.String("<ID>"),
            IsActive: unifiedgosdk.Bool(false),
            ManagerIds: []string{
                "JSON",
            },
            Name: unifiedgosdk.String("whereas Usability transmitting"),
            ParentID: unifiedgosdk.String("invoice Cyclocross Electric"),
            Raw: &shared.PropertyHrisGroupRaw{},
            Type: shared.HrisGroupTypeGroup.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2021-04-30T12:40:50.129Z"),
        },
        ConnectionID: "Hybrid Schenectady",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.HrisGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.PostHrisConnectionIDGroupRequest](../../models/operations/posthrisconnectionidgrouprequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.PostHrisConnectionIDGroupResponse](../../models/operations/posthrisconnectionidgroupresponse.md), error**


## PutHrisConnectionIDGroupID

Update a Group

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
    res, err := s.Group.PutHrisConnectionIDGroupID(ctx, operations.PutHrisConnectionIDGroupIDRequest{
        HrisGroup: &shared.HrisGroup{
            CreatedAt: types.MustTimeFromString("2022-08-10T12:11:42.375Z"),
            Description: unifiedgosdk.String("Decentralized methodical projection"),
            EmployeeIds: []string{
                "Credit",
            },
            ID: unifiedgosdk.String("<ID>"),
            IsActive: unifiedgosdk.Bool(false),
            ManagerIds: []string{
                "South",
            },
            Name: unifiedgosdk.String("Jeep brr Northwest"),
            ParentID: unifiedgosdk.String("quickly Licensed"),
            Raw: &shared.PropertyHrisGroupRaw{},
            Type: shared.HrisGroupTypeBusinessUnit.ToPointer(),
            UpdatedAt: types.MustTimeFromString("2021-11-08T00:11:45.458Z"),
        },
        ConnectionID: "vortals interface Gasoline",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.HrisGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PutHrisConnectionIDGroupIDRequest](../../models/operations/puthrisconnectionidgroupidrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.PutHrisConnectionIDGroupIDResponse](../../models/operations/puthrisconnectionidgroupidresponse.md), error**

