# Connection
(*Connection*)

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
            operations.GetUnifiedConnectionCategoriesAts,
        },
        Env: unifiedgosdk.String("copying invoice coulomb"),
        ExternalXref: unifiedgosdk.String("Islands West"),
        Limit: unifiedgosdk.Float64(7809.21),
        Offset: unifiedgosdk.Float64(2750.2),
        Order: unifiedgosdk.String("Volkswagen architect"),
        Sort: unifiedgosdk.String("consequently synthesizing Technician"),
        UpdatedGte: types.MustTimeFromString("2021-11-09T20:41:53.442Z"),
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
        ID: "<ID>",
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
                AccessToken: unifiedgosdk.String("Northwest Cupertino"),
                APIURL: unifiedgosdk.String("Center Curium Electric"),
                AppID: unifiedgosdk.String("female fragrant"),
                AuthorizeURL: unifiedgosdk.String("Electric Bicycle payment"),
                ClientID: unifiedgosdk.String("transmitting North"),
                ClientSecret: unifiedgosdk.String("mole Gasoline morph"),
                ConsumerKey: unifiedgosdk.String("Keyboard Antimony primary"),
                ConsumerSecret: unifiedgosdk.String("yearly"),
                Emails: []string{
                    "athwart",
                },
                ExpiresIn: unifiedgosdk.Float64(3185.09),
                ExpiryDate: types.MustTimeFromString("2022-08-12T13:21:47.977Z"),
                Key: unifiedgosdk.String("<key>"),
                Meta: &shared.PropertyPropertyConnectionAuthMeta{},
                Name: unifiedgosdk.String("inside Rupee"),
                OtherAuthInfo: []string{
                    "Future",
                },
                Pem: unifiedgosdk.String("guard Internal"),
                RefreshToken: unifiedgosdk.String("Diesel"),
                State: unifiedgosdk.String("copy Cotton Bicycle"),
                Token: unifiedgosdk.String("drive gold"),
                TokenURL: unifiedgosdk.String("now"),
            },
            AuthAwsArn: unifiedgosdk.String("yum"),
            Categories: []shared.PropertyConnectionCategories{
                shared.PropertyConnectionCategoriesHris,
            },
            CreatedAt: types.MustTimeFromString("2021-06-18T22:02:30.822Z"),
            Environment: unifiedgosdk.String("Northwest Balanced"),
            ExternalXref: unifiedgosdk.String("boo"),
            ID: unifiedgosdk.String("<ID>"),
            IntegrationType: "Soft",
            IsPaused: unifiedgosdk.Bool(false),
            Permissions: []shared.PropertyConnectionPermissions{
                shared.PropertyConnectionPermissionsCrmLeadRead,
            },
            UpdatedAt: types.MustTimeFromString("2022-07-27T15:43:44.767Z"),
            WorkspaceID: unifiedgosdk.String("extend"),
        },
        ID: "<ID>",
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
            AccessToken: unifiedgosdk.String("asperiores indexing"),
            APIURL: unifiedgosdk.String("plus pace global"),
            AppID: unifiedgosdk.String("And port"),
            AuthorizeURL: unifiedgosdk.String("West whiteboard"),
            ClientID: unifiedgosdk.String("Folk"),
            ClientSecret: unifiedgosdk.String("Northwest Modern"),
            ConsumerKey: unifiedgosdk.String("Southeast deposit"),
            ConsumerSecret: unifiedgosdk.String("Falls irritating up"),
            Emails: []string{
                "intuitive",
            },
            ExpiresIn: unifiedgosdk.Float64(4121.5),
            ExpiryDate: types.MustTimeFromString("2021-01-21T03:25:42.786Z"),
            Key: unifiedgosdk.String("<key>"),
            Meta: &shared.PropertyPropertyConnectionAuthMeta{},
            Name: unifiedgosdk.String("membership Classical schnitzel"),
            OtherAuthInfo: []string{
                "Convertible",
            },
            Pem: unifiedgosdk.String("magenta Riel bolívar"),
            RefreshToken: unifiedgosdk.String("Pula"),
            State: unifiedgosdk.String("white"),
            Token: unifiedgosdk.String("Northwest"),
            TokenURL: unifiedgosdk.String("unbutton"),
        },
        AuthAwsArn: unifiedgosdk.String("Investor circuit"),
        Categories: []shared.PropertyConnectionCategories{
            shared.PropertyConnectionCategoriesAts,
        },
        CreatedAt: types.MustTimeFromString("2023-01-25T14:37:40.202Z"),
        Environment: unifiedgosdk.String("Hat watt"),
        ExternalXref: unifiedgosdk.String("Sausages tan"),
        ID: unifiedgosdk.String("<ID>"),
        IntegrationType: "Principal Extended velit",
        IsPaused: unifiedgosdk.Bool(false),
        Permissions: []shared.PropertyConnectionPermissions{
            shared.PropertyConnectionPermissionsCrmFileWrite,
        },
        UpdatedAt: types.MustTimeFromString("2021-09-22T05:13:05.778Z"),
        WorkspaceID: unifiedgosdk.String("Auto"),
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
                AccessToken: unifiedgosdk.String("female Buckinghamshire"),
                APIURL: unifiedgosdk.String("Web"),
                AppID: unifiedgosdk.String("pumpkin Account"),
                AuthorizeURL: unifiedgosdk.String("dolorem Hybrid white"),
                ClientID: unifiedgosdk.String("ohm"),
                ClientSecret: unifiedgosdk.String("Pennsylvania Executive"),
                ConsumerKey: unifiedgosdk.String("silver Account Accountability"),
                ConsumerSecret: unifiedgosdk.String("Mouse"),
                Emails: []string{
                    "oh",
                },
                ExpiresIn: unifiedgosdk.Float64(8946.31),
                ExpiryDate: types.MustTimeFromString("2022-01-29T12:35:08.478Z"),
                Key: unifiedgosdk.String("<key>"),
                Meta: &shared.PropertyPropertyConnectionAuthMeta{},
                Name: unifiedgosdk.String("incidentally shrimp bypass"),
                OtherAuthInfo: []string{
                    "invoice",
                },
                Pem: unifiedgosdk.String("recent midst Northeast"),
                RefreshToken: unifiedgosdk.String("Product"),
                State: unifiedgosdk.String("circuit precious"),
                Token: unifiedgosdk.String("gee collaborative withdrawal"),
                TokenURL: unifiedgosdk.String("Platinum"),
            },
            AuthAwsArn: unifiedgosdk.String("suddenly Fiat"),
            Categories: []shared.PropertyConnectionCategories{
                shared.PropertyConnectionCategoriesUc,
            },
            CreatedAt: types.MustTimeFromString("2022-09-20T19:51:21.025Z"),
            Environment: unifiedgosdk.String("redundant Southeast Camren"),
            ExternalXref: unifiedgosdk.String("firewall"),
            ID: unifiedgosdk.String("<ID>"),
            IntegrationType: "Beauty",
            IsPaused: unifiedgosdk.Bool(false),
            Permissions: []shared.PropertyConnectionPermissions{
                shared.PropertyConnectionPermissionsWebhook,
            },
            UpdatedAt: types.MustTimeFromString("2023-12-30T14:20:47.994Z"),
            WorkspaceID: unifiedgosdk.String("parse Peso Investment"),
        },
        ID: "<ID>",
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

