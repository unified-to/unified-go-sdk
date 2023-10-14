# Team
(*Team*)

### Available Operations

* [CreateCrmTeam](#createcrmteam) - Create a team
* [GetCrmTeam](#getcrmteam) - Retrieve a team
* [ListCrmTeams](#listcrmteams) - List all teams
* [PatchCrmTeam](#patchcrmteam) - Update a team
* [RemoveCrmTeam](#removecrmteam) - Remove a team
* [UpdateCrmTeam](#updatecrmteam) - Update a team

## CreateCrmTeam

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
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Team.CreateCrmTeam(ctx, operations.CreateCrmTeamRequest{
        CrmTeam: &shared.CrmTeam{
            Raw: &shared.PropertyCrmTeamRaw{},
            UserIds: []string{
                "exercitationem",
            },
        },
        ConnectionID: "as New Senior",
        Fields: []string{
            "Buckinghamshire",
        },
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CreateCrmTeamRequest](../../models/operations/createcrmteamrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.CreateCrmTeamResponse](../../models/operations/createcrmteamresponse.md), error**


## GetCrmTeam

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Team.GetCrmTeam(ctx, operations.GetCrmTeamRequest{
        ConnectionID: "digital awful",
        Fields: []string{
            "Peru",
        },
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetCrmTeamRequest](../../models/operations/getcrmteamrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetCrmTeamResponse](../../models/operations/getcrmteamresponse.md), error**


## ListCrmTeams

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
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Team.ListCrmTeams(ctx, operations.ListCrmTeamsRequest{
        ConnectionID: "Classical microchip Wooden",
        Fields: []string{
            "Lutetium",
        },
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListCrmTeamsRequest](../../models/operations/listcrmteamsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.ListCrmTeamsResponse](../../models/operations/listcrmteamsresponse.md), error**


## PatchCrmTeam

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
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Team.PatchCrmTeam(ctx, operations.PatchCrmTeamRequest{
        CrmTeam: &shared.CrmTeam{
            Raw: &shared.PropertyCrmTeamRaw{},
            UserIds: []string{
                "Account",
            },
        },
        ConnectionID: "Transexual compress redefine",
        Fields: []string{
            "gold",
        },
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.PatchCrmTeamRequest](../../models/operations/patchcrmteamrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.PatchCrmTeamResponse](../../models/operations/patchcrmteamresponse.md), error**


## RemoveCrmTeam

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Team.RemoveCrmTeam(ctx, operations.RemoveCrmTeamRequest{
        ConnectionID: "Sol",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.RemoveCrmTeamRequest](../../models/operations/removecrmteamrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.RemoveCrmTeamResponse](../../models/operations/removecrmteamresponse.md), error**


## UpdateCrmTeam

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
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Team.UpdateCrmTeam(ctx, operations.UpdateCrmTeamRequest{
        CrmTeam: &shared.CrmTeam{
            Raw: &shared.PropertyCrmTeamRaw{},
            UserIds: []string{
                "Carbon",
            },
        },
        ConnectionID: "Dakota",
        Fields: []string{
            "female",
        },
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateCrmTeamRequest](../../models/operations/updatecrmteamrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.UpdateCrmTeamResponse](../../models/operations/updatecrmteamresponse.md), error**

