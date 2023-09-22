# Connection

### Available Operations

* [DeleteUnifiedConnectionID](#deleteunifiedconnectionid) - Remove connection
* [GetUnifiedConnection](#getunifiedconnection) - List all connections
* [GetUnifiedConnectionID](#getunifiedconnectionid) - Retrieve connection
* [PatchUnifiedConnectionID](#patchunifiedconnectionid) - Update connection
* [PostUnifiedConnection](#postunifiedconnection) - Create connection
* [PutUnifiedConnectionID](#putunifiedconnectionid) - Update connection

## DeleteUnifiedConnectionID

Remove connection

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
    res, err := s.Connection.DeleteUnifiedConnectionID(ctx, operations.DeleteUnifiedConnectionIDRequest{
        ID: "b75d2367-fe1a-40cc-8df7-9f0a396d90c3",
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.DeleteUnifiedConnectionIDRequest](../../models/operations/deleteunifiedconnectionidrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.DeleteUnifiedConnectionIDResponse](../../models/operations/deleteunifiedconnectionidresponse.md), error**


## GetUnifiedConnection

List all connections

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
    res, err := s.Connection.GetUnifiedConnection(ctx, operations.GetUnifiedConnectionRequest{
        Categories: []GetUnifiedConnectionCategories{
            operations.GetUnifiedConnectionCategoriesAuth,
        },
        Env: unifiedgosdk.String("labore"),
        ExternalXref: unifiedgosdk.String("expedita"),
        Limit: unifiedgosdk.Float64(4467.37),
        Offset: unifiedgosdk.Float64(7898.7),
        Order: unifiedgosdk.String("sunt"),
        Sort: unifiedgosdk.String("enim"),
        UpdatedGte: types.MustTimeFromString("2020-01-24T16:46:41.592Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Connections != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetUnifiedConnectionRequest](../../models/operations/getunifiedconnectionrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.GetUnifiedConnectionResponse](../../models/operations/getunifiedconnectionresponse.md), error**


## GetUnifiedConnectionID

Retrieve connection

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
    res, err := s.Connection.GetUnifiedConnectionID(ctx, operations.GetUnifiedConnectionIDRequest{
        ID: "bace188b-1c4e-4e2c-8c6c-e611feeb1c7c",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetUnifiedConnectionIDRequest](../../models/operations/getunifiedconnectionidrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.GetUnifiedConnectionIDResponse](../../models/operations/getunifiedconnectionidresponse.md), error**


## PatchUnifiedConnectionID

Update connection

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
    res, err := s.Connection.PatchUnifiedConnectionID(ctx, operations.PatchUnifiedConnectionIDRequest{
        Connection: &shared.Connection{
            Auth: &shared.PropertyConnectionAuth{
                AccessToken: unifiedgosdk.String("distinctio"),
                APIURL: unifiedgosdk.String("possimus"),
                AppID: unifiedgosdk.String("cum"),
                AuthorizeURL: unifiedgosdk.String("suscipit"),
                ClientID: unifiedgosdk.String("saepe"),
                ClientSecret: unifiedgosdk.String("earum"),
                ConsumerKey: unifiedgosdk.String("quod"),
                ConsumerSecret: unifiedgosdk.String("nihil"),
                Emails: []string{
                    "quaerat",
                },
                ExpiresIn: unifiedgosdk.Float64(2159.59),
                ExpiryDate: types.MustTimeFromString("2022-06-27T03:53:09.176Z"),
                Key: unifiedgosdk.String("rerum"),
                Meta: &shared.PropertyPropertyConnectionAuthMeta{},
                Name: unifiedgosdk.String("Harry Hammes IV"),
                OtherAuthInfo: []string{
                    "esse",
                },
                Pem: unifiedgosdk.String("magnam"),
                RefreshToken: unifiedgosdk.String("odio"),
                State: unifiedgosdk.String("nulla"),
                Token: unifiedgosdk.String("impedit"),
                TokenURL: unifiedgosdk.String("cupiditate"),
            },
            AuthAwsArn: unifiedgosdk.String("illo"),
            Categories: []shared.PropertyConnectionCategories{
                shared.PropertyConnectionCategoriesAuth,
            },
            CreatedAt: types.MustTimeFromString("2021-04-21T08:26:42.931Z"),
            Environment: unifiedgosdk.String("fugit"),
            ExternalXref: unifiedgosdk.String("maxime"),
            ID: unifiedgosdk.String("af5dd672-3dc0-4f5a-a2f3-a6b700878756"),
            IntegrationType: "ab",
            IsPaused: unifiedgosdk.Bool(false),
            Permissions: []shared.PropertyConnectionPermissions{
                shared.PropertyConnectionPermissionsCrmUserWrite,
            },
            UpdatedAt: types.MustTimeFromString("2022-01-16T17:11:29.866Z"),
            WorkspaceID: unifiedgosdk.String("corporis"),
        },
        ID: "a6c98b55-5540-480d-80bc-acc6cbd6b5f3",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PatchUnifiedConnectionIDRequest](../../models/operations/patchunifiedconnectionidrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.PatchUnifiedConnectionIDResponse](../../models/operations/patchunifiedconnectionidresponse.md), error**


## PostUnifiedConnection

Create connection

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Connection.PostUnifiedConnection(ctx, shared.Connection{
        Auth: &shared.PropertyConnectionAuth{
            AccessToken: unifiedgosdk.String("necessitatibus"),
            APIURL: unifiedgosdk.String("maxime"),
            AppID: unifiedgosdk.String("cupiditate"),
            AuthorizeURL: unifiedgosdk.String("voluptatem"),
            ClientID: unifiedgosdk.String("provident"),
            ClientSecret: unifiedgosdk.String("adipisci"),
            ConsumerKey: unifiedgosdk.String("accusantium"),
            ConsumerSecret: unifiedgosdk.String("magnam"),
            Emails: []string{
                "repellat",
            },
            ExpiresIn: unifiedgosdk.Float64(6076.31),
            ExpiryDate: types.MustTimeFromString("2022-07-29T13:50:17.340Z"),
            Key: unifiedgosdk.String("cum"),
            Meta: &shared.PropertyPropertyConnectionAuthMeta{},
            Name: unifiedgosdk.String("Rufus Conroy"),
            OtherAuthInfo: []string{
                "sequi",
            },
            Pem: unifiedgosdk.String("voluptatum"),
            RefreshToken: unifiedgosdk.String("quasi"),
            State: unifiedgosdk.String("error"),
            Token: unifiedgosdk.String("nobis"),
            TokenURL: unifiedgosdk.String("tempora"),
        },
        AuthAwsArn: unifiedgosdk.String("voluptate"),
        Categories: []shared.PropertyConnectionCategories{
            shared.PropertyConnectionCategoriesAts,
        },
        CreatedAt: types.MustTimeFromString("2022-11-28T03:21:08.707Z"),
        Environment: unifiedgosdk.String("voluptates"),
        ExternalXref: unifiedgosdk.String("possimus"),
        ID: unifiedgosdk.String("20e56248-fff6-439a-910a-bdcab6267669"),
        IntegrationType: "laboriosam",
        IsPaused: unifiedgosdk.Bool(false),
        Permissions: []shared.PropertyConnectionPermissions{
            shared.PropertyConnectionPermissionsAtsCandidateWrite,
        },
        UpdatedAt: types.MustTimeFromString("2022-01-26T21:22:49.757Z"),
        WorkspaceID: unifiedgosdk.String("quisquam"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                              | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `ctx`                                                  | [context.Context](https://pkg.go.dev/context#Context)  | :heavy_check_mark:                                     | The context to use for the request.                    |
| `request`                                              | [shared.Connection](../../models/shared/connection.md) | :heavy_check_mark:                                     | The request object to use for the request.             |


### Response

**[*operations.PostUnifiedConnectionResponse](../../models/operations/postunifiedconnectionresponse.md), error**


## PutUnifiedConnectionID

Update connection

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
    res, err := s.Connection.PutUnifiedConnectionID(ctx, operations.PutUnifiedConnectionIDRequest{
        Connection: &shared.Connection{
            Auth: &shared.PropertyConnectionAuth{
                AccessToken: unifiedgosdk.String("eaque"),
                APIURL: unifiedgosdk.String("alias"),
                AppID: unifiedgosdk.String("qui"),
                AuthorizeURL: unifiedgosdk.String("consequuntur"),
                ClientID: unifiedgosdk.String("vitae"),
                ClientSecret: unifiedgosdk.String("quidem"),
                ConsumerKey: unifiedgosdk.String("sequi"),
                ConsumerSecret: unifiedgosdk.String("amet"),
                Emails: []string{
                    "exercitationem",
                },
                ExpiresIn: unifiedgosdk.Float64(8470.18),
                ExpiryDate: types.MustTimeFromString("2021-10-18T19:31:06.005Z"),
                Key: unifiedgosdk.String("similique"),
                Meta: &shared.PropertyPropertyConnectionAuthMeta{},
                Name: unifiedgosdk.String("Garry Fahey"),
                OtherAuthInfo: []string{
                    "asperiores",
                },
                Pem: unifiedgosdk.String("temporibus"),
                RefreshToken: unifiedgosdk.String("id"),
                State: unifiedgosdk.String("atque"),
                Token: unifiedgosdk.String("quibusdam"),
                TokenURL: unifiedgosdk.String("sit"),
            },
            AuthAwsArn: unifiedgosdk.String("quo"),
            Categories: []shared.PropertyConnectionCategories{
                shared.PropertyConnectionCategoriesAts,
            },
            CreatedAt: types.MustTimeFromString("2022-05-30T05:35:32.331Z"),
            Environment: unifiedgosdk.String("vero"),
            ExternalXref: unifiedgosdk.String("earum"),
            ID: unifiedgosdk.String("03004978-a61f-4a1c-b206-88f77c1ffc71"),
            IntegrationType: "facere",
            IsPaused: unifiedgosdk.Bool(false),
            Permissions: []shared.PropertyConnectionPermissions{
                shared.PropertyConnectionPermissionsHrisGroupRead,
            },
            UpdatedAt: types.MustTimeFromString("2022-11-13T09:13:55.156Z"),
            WorkspaceID: unifiedgosdk.String("ex"),
        },
        ID: "3f2a3c80-a97f-4f33-8cdd-f857a9e61876",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PutUnifiedConnectionIDRequest](../../models/operations/putunifiedconnectionidrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.PutUnifiedConnectionIDResponse](../../models/operations/putunifiedconnectionidresponse.md), error**

