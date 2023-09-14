# Team

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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.DeleteCrmConnectionIDTeamIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Team.DeleteCrmConnectionIDTeamID(ctx, operations.DeleteCrmConnectionIDTeamIDRequest{
        ConnectionID: "ab",
        ID: "bbf05527-1b25-411d-9606-dd1b28272bc9",
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
| `request`                                                                                                        | [operations.DeleteCrmConnectionIDTeamIDRequest](../../models/operations/deletecrmconnectionidteamidrequest.md)   | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `security`                                                                                                       | [operations.DeleteCrmConnectionIDTeamIDSecurity](../../models/operations/deletecrmconnectionidteamidsecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetCrmConnectionIDTeamSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Team.GetCrmConnectionIDTeam(ctx, operations.GetCrmConnectionIDTeamRequest{
        ConnectionID: "placeat",
        Limit: unifiedto.Float64(1884.9),
        Offset: unifiedto.Float64(1694.68),
        Order: unifiedto.String("sunt"),
        Query: unifiedto.String("vitae"),
        Sort: unifiedto.String("ex"),
        UpdatedGte: types.MustTimeFromString("2022-01-05T07:41:51.025Z"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmTeams != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetCrmConnectionIDTeamRequest](../../models/operations/getcrmconnectionidteamrequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.GetCrmConnectionIDTeamSecurity](../../models/operations/getcrmconnectionidteamsecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetCrmConnectionIDTeamIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Team.GetCrmConnectionIDTeamID(ctx, operations.GetCrmConnectionIDTeamIDRequest{
        ConnectionID: "rerum",
        ID: "1880fcbb-2b93-4c15-b670-bd1784831653",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmTeam != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetCrmConnectionIDTeamIDRequest](../../models/operations/getcrmconnectionidteamidrequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.GetCrmConnectionIDTeamIDSecurity](../../models/operations/getcrmconnectionidteamidsecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PatchCrmConnectionIDTeamIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Team.PatchCrmConnectionIDTeamID(ctx, operations.PatchCrmConnectionIDTeamIDRequest{
        CrmTeam: &shared.CrmTeam{
            CreatedAt: types.MustTimeFromString("2020-04-24T00:39:17.172Z"),
            Description: unifiedto.String("harum"),
            ID: unifiedto.String("3b6e241c-3109-4983-a63c-66dcbb7df6cb"),
            Name: unifiedto.String("Jenny Rolfson"),
            Raw: &shared.PropertyCrmTeamRaw{},
            UpdatedAt: types.MustTimeFromString("2022-12-14T00:49:36.543Z"),
            UserIds: []string{
                "praesentium",
            },
        },
        ConnectionID: "recusandae",
        ID: "0713774d-e4fe-4e10-9d97-80a10c47b950",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmTeam != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PatchCrmConnectionIDTeamIDRequest](../../models/operations/patchcrmconnectionidteamidrequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.PatchCrmConnectionIDTeamIDSecurity](../../models/operations/patchcrmconnectionidteamidsecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PostCrmConnectionIDTeamSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Team.PostCrmConnectionIDTeam(ctx, operations.PostCrmConnectionIDTeamRequest{
        CrmTeam: &shared.CrmTeam{
            CreatedAt: types.MustTimeFromString("2022-12-18T11:11:12.745Z"),
            Description: unifiedto.String("possimus"),
            ID: unifiedto.String("6c8b2a5f-0022-407e-8048-f90009ed2902"),
            Name: unifiedto.String("Brandy Tillman"),
            Raw: &shared.PropertyCrmTeamRaw{},
            UpdatedAt: types.MustTimeFromString("2021-02-19T10:57:16.366Z"),
            UserIds: []string{
                "iste",
            },
        },
        ConnectionID: "pariatur",
    }, operationSecurity)
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
| `request`                                                                                                | [operations.PostCrmConnectionIDTeamRequest](../../models/operations/postcrmconnectionidteamrequest.md)   | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.PostCrmConnectionIDTeamSecurity](../../models/operations/postcrmconnectionidteamsecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PutCrmConnectionIDTeamIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Team.PutCrmConnectionIDTeamID(ctx, operations.PutCrmConnectionIDTeamIDRequest{
        CrmTeam: &shared.CrmTeam{
            CreatedAt: types.MustTimeFromString("2022-09-20T15:32:20.854Z"),
            Description: unifiedto.String("sunt"),
            ID: unifiedto.String("61e91500-323b-42c0-9b92-4771f5669e5b"),
            Name: unifiedto.String("Tricia Sawayn"),
            Raw: &shared.PropertyCrmTeamRaw{},
            UpdatedAt: types.MustTimeFromString("2022-07-25T07:35:50.345Z"),
            UserIds: []string{
                "ea",
            },
        },
        ConnectionID: "labore",
        ID: "9d84eb9e-4cfd-4227-ae0b-88fb87d6fa5b",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrmTeam != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.PutCrmConnectionIDTeamIDRequest](../../models/operations/putcrmconnectionidteamidrequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.PutCrmConnectionIDTeamIDSecurity](../../models/operations/putcrmconnectionidteamidsecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |


### Response

**[*operations.PutCrmConnectionIDTeamIDResponse](../../models/operations/putcrmconnectionidteamidresponse.md), error**

