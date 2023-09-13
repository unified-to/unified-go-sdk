# Group

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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.DeleteHrisConnectionIDGroupIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Group.DeleteHrisConnectionIDGroupID(ctx, operations.DeleteHrisConnectionIDGroupIDRequest{
        ConnectionID: "ipsam",
        ID: "8aaeacae-323a-431b-b7ba-1cc97716c802",
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

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.DeleteHrisConnectionIDGroupIDRequest](../../models/operations/deletehrisconnectionidgroupidrequest.md)   | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `security`                                                                                                           | [operations.DeleteHrisConnectionIDGroupIDSecurity](../../models/operations/deletehrisconnectionidgroupidsecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetHrisConnectionIDGroupSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Group.GetHrisConnectionIDGroup(ctx, operations.GetHrisConnectionIDGroupRequest{
        ConnectionID: "minus",
        Limit: unifiedto.Float64(7864.46),
        Offset: unifiedto.Float64(5742.21),
        Order: unifiedto.String("voluptates"),
        Query: unifiedto.String("alias"),
        Sort: unifiedto.String("placeat"),
        UpdatedGte: types.MustDateFromString("2022-03-07"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.HrisGroups != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetHrisConnectionIDGroupRequest](../../models/operations/gethrisconnectionidgrouprequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.GetHrisConnectionIDGroupSecurity](../../models/operations/gethrisconnectionidgroupsecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetHrisConnectionIDGroupIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Group.GetHrisConnectionIDGroupID(ctx, operations.GetHrisConnectionIDGroupIDRequest{
        ConnectionID: "iste",
        ID: "d323f1aa-63ed-49cf-9c85-6bcba51ef245",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.HrisGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetHrisConnectionIDGroupIDRequest](../../models/operations/gethrisconnectionidgroupidrequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.GetHrisConnectionIDGroupIDSecurity](../../models/operations/gethrisconnectionidgroupidsecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PatchHrisConnectionIDGroupIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Group.PatchHrisConnectionIDGroupID(ctx, operations.PatchHrisConnectionIDGroupIDRequest{
        HrisGroup: &shared.HrisGroup{
            CreatedAt: types.MustDateFromString("2022-05-13"),
            Description: unifiedto.String("aliquam"),
            EmployeeIds: []string{
                "iusto",
            },
            ID: unifiedto.String("facf116c-dd54-444a-b562-873c7dd9efaf"),
            IsActive: unifiedto.Bool(false),
            ManagerIds: []string{
                "labore",
            },
            Name: unifiedto.String("Cristina Russel"),
            ParentID: unifiedto.String("consectetur"),
            Raw: &shared.PropertyHrisGroupRaw{},
            Type: shared.HrisGroupTypeDepartment.ToPointer(),
            UpdatedAt: types.MustDateFromString("2022-12-28"),
        },
        ConnectionID: "delectus",
        ID: "3138f30d-f3db-4022-baa5-65fb8f652ebb",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.HrisGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.PatchHrisConnectionIDGroupIDRequest](../../models/operations/patchhrisconnectionidgroupidrequest.md)   | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `security`                                                                                                         | [operations.PatchHrisConnectionIDGroupIDSecurity](../../models/operations/patchhrisconnectionidgroupidsecurity.md) | :heavy_check_mark:                                                                                                 | The security requirements to use for the request.                                                                  |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PostHrisConnectionIDGroupSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Group.PostHrisConnectionIDGroup(ctx, operations.PostHrisConnectionIDGroupRequest{
        HrisGroup: &shared.HrisGroup{
            CreatedAt: types.MustDateFromString("2021-04-22"),
            Description: unifiedto.String("dolor"),
            EmployeeIds: []string{
                "praesentium",
            },
            ID: unifiedto.String("38387902-43b2-493d-ab30-e917f50fda04"),
            IsActive: unifiedto.Bool(false),
            ManagerIds: []string{
                "porro",
            },
            Name: unifiedto.String("Wm Boyer"),
            ParentID: unifiedto.String("exercitationem"),
            Raw: &shared.PropertyHrisGroupRaw{},
            Type: shared.HrisGroupTypeDepartment.ToPointer(),
            UpdatedAt: types.MustDateFromString("2022-08-30"),
        },
        ConnectionID: "unde",
    }, operationSecurity)
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
| `request`                                                                                                    | [operations.PostHrisConnectionIDGroupRequest](../../models/operations/posthrisconnectionidgrouprequest.md)   | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `security`                                                                                                   | [operations.PostHrisConnectionIDGroupSecurity](../../models/operations/posthrisconnectionidgroupsecurity.md) | :heavy_check_mark:                                                                                           | The security requirements to use for the request.                                                            |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PutHrisConnectionIDGroupIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Group.PutHrisConnectionIDGroupID(ctx, operations.PutHrisConnectionIDGroupIDRequest{
        HrisGroup: &shared.HrisGroup{
            CreatedAt: types.MustDateFromString("2022-04-06"),
            Description: unifiedto.String("aut"),
            EmployeeIds: []string{
                "expedita",
            },
            ID: unifiedto.String("c3bb7446-64eb-41d0-b388-b0d1bb17afee"),
            IsActive: unifiedto.Bool(false),
            ManagerIds: []string{
                "reprehenderit",
            },
            Name: unifiedto.String("Latoya Hodkiewicz"),
            ParentID: unifiedto.String("quidem"),
            Raw: &shared.PropertyHrisGroupRaw{},
            Type: shared.HrisGroupTypeBusinessUnit.ToPointer(),
            UpdatedAt: types.MustDateFromString("2022-08-19"),
        },
        ConnectionID: "voluptate",
        ID: "c7edaf39-d16f-4bf7-afd1-62b303e3023b",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.HrisGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PutHrisConnectionIDGroupIDRequest](../../models/operations/puthrisconnectionidgroupidrequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.PutHrisConnectionIDGroupIDSecurity](../../models/operations/puthrisconnectionidgroupidsecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


### Response

**[*operations.PutHrisConnectionIDGroupIDResponse](../../models/operations/puthrisconnectionidgroupidresponse.md), error**

