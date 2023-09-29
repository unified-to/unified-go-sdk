# Team
(*Team*)

### Available Operations

* [DeleteCrmConnectionIDTeamID](#deletecrmconnectionidteamid) - Remove a team
* [GetCrmConnectionIDTeam](#getcrmconnectionidteam) - List all teams
* [GetCrmConnectionIDTeamID](#getcrmconnectionidteamid) - Retrieve a team
* [PatchCrmConnectionIDTeamID](#patchcrmconnectionidteamid) - Update a team
* [PostCrmConnectionIDTeam](#postcrmconnectionidteam) - Create a team
* [PutCrmConnectionIDTeamID](#putcrmconnectionidteamid) - Update a team

## DeleteCrmConnectionIDTeamID

Remove a team

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
    res, err := s.Team.DeleteCrmConnectionIDTeamID(ctx, operations.DeleteCrmConnectionIDTeamIDRequest{
        ConnectionID: "Diverse",
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
| `request`                                                                                                      | [operations.DeleteCrmConnectionIDTeamIDRequest](../../models/operations/deletecrmconnectionidteamidrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.DeleteCrmConnectionIDTeamIDResponse](../../models/operations/deletecrmconnectionidteamidresponse.md), error**


## GetCrmConnectionIDTeam

List all teams

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
    res, err := s.Team.GetCrmConnectionIDTeam(ctx, operations.GetCrmConnectionIDTeamRequest{
        ConnectionID: "bath Lamborghini",
        Limit: unifiedgosdk.Float64(1042.31),
        Offset: unifiedgosdk.Float64(1586.42),
        Order: unifiedgosdk.String("Diesel Bike virtual"),
        Query: unifiedgosdk.String("bakery"),
        Sort: unifiedgosdk.String("Senior"),
        UpdatedGte: types.MustTimeFromString("2021-12-04T23:56:00.028Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmTeams != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetCrmConnectionIDTeamRequest](../../models/operations/getcrmconnectionidteamrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.GetCrmConnectionIDTeamResponse](../../models/operations/getcrmconnectionidteamresponse.md), error**


## GetCrmConnectionIDTeamID

Retrieve a team

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
    res, err := s.Team.GetCrmConnectionIDTeamID(ctx, operations.GetCrmConnectionIDTeamIDRequest{
        ConnectionID: "Intelligent invoice Tesla",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmTeam != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetCrmConnectionIDTeamIDRequest](../../models/operations/getcrmconnectionidteamidrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.GetCrmConnectionIDTeamIDResponse](../../models/operations/getcrmconnectionidteamidresponse.md), error**


## PatchCrmConnectionIDTeamID

Update a team

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
    res, err := s.Team.PatchCrmConnectionIDTeamID(ctx, operations.PatchCrmConnectionIDTeamIDRequest{
        CrmTeam: &shared.CrmTeam{
            CreatedAt: types.MustTimeFromString("2021-05-20T12:47:48.451Z"),
            Description: unifiedgosdk.String("Automated executive emulation"),
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("Internal experiences"),
            Raw: &shared.PropertyCrmTeamRaw{},
            UpdatedAt: types.MustTimeFromString("2022-05-22T09:41:53.599Z"),
            UserIds: []string{
                "lumen",
            },
        },
        ConnectionID: "up Candace",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmTeam != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PatchCrmConnectionIDTeamIDRequest](../../models/operations/patchcrmconnectionidteamidrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.PatchCrmConnectionIDTeamIDResponse](../../models/operations/patchcrmconnectionidteamidresponse.md), error**


## PostCrmConnectionIDTeam

Create a team

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
    res, err := s.Team.PostCrmConnectionIDTeam(ctx, operations.PostCrmConnectionIDTeamRequest{
        CrmTeam: &shared.CrmTeam{
            CreatedAt: types.MustTimeFromString("2022-02-12T08:57:03.070Z"),
            Description: unifiedgosdk.String("Organic transitional portal"),
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("male bandwidth"),
            Raw: &shared.PropertyCrmTeamRaw{},
            UpdatedAt: types.MustTimeFromString("2022-12-29T15:50:04.365Z"),
            UserIds: []string{
                "meter",
            },
        },
        ConnectionID: "Guaynabo AGP East",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmTeam != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.PostCrmConnectionIDTeamRequest](../../models/operations/postcrmconnectionidteamrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.PostCrmConnectionIDTeamResponse](../../models/operations/postcrmconnectionidteamresponse.md), error**


## PutCrmConnectionIDTeamID

Update a team

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
    res, err := s.Team.PutCrmConnectionIDTeamID(ctx, operations.PutCrmConnectionIDTeamIDRequest{
        CrmTeam: &shared.CrmTeam{
            CreatedAt: types.MustTimeFromString("2023-08-14T23:28:53.515Z"),
            Description: unifiedgosdk.String("Inverse multi-tasking task-force"),
            ID: unifiedgosdk.String("<ID>"),
            Name: unifiedgosdk.String("Indonesia Orchestrator Division"),
            Raw: &shared.PropertyCrmTeamRaw{},
            UpdatedAt: types.MustTimeFromString("2022-10-23T23:13:25.973Z"),
            UserIds: []string{
                "thoroughly",
            },
        },
        ConnectionID: "delectus",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmTeam != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PutCrmConnectionIDTeamIDRequest](../../models/operations/putcrmconnectionidteamidrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.PutCrmConnectionIDTeamIDResponse](../../models/operations/putcrmconnectionidteamidresponse.md), error**

