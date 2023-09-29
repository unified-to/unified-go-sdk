# Unified
(*Unified*)

### Available Operations

* [DeleteUnifiedConnectionID](#deleteunifiedconnectionid) - Remove connection
* [DeleteUnifiedUser](#deleteunifieduser) - Delete your user object
* [DeleteUnifiedWebhookID](#deleteunifiedwebhookid) - Remove webhook subscription
* [GetUnifiedApicall](#getunifiedapicall) - Returns API Calls
* [GetUnifiedApicallID](#getunifiedapicallid) - Retrieve specific API Call by its ID
* [GetUnifiedConnection](#getunifiedconnection) - List all connections
* [GetUnifiedConnectionID](#getunifiedconnectionid) - Retrieve connection
* [GetUnifiedIntegration](#getunifiedintegration) - Returns all integrations
* [GetUnifiedIntegrationAuthWorkspaceIDIntegrationType](#getunifiedintegrationauthworkspaceidintegrationtype) - Create connection indirectly
* [GetUnifiedIntegrationIntegrationType](#getunifiedintegrationintegrationtype) - Retrieve an integration
* [GetUnifiedIntegrationWorkspaceWorkspaceID](#getunifiedintegrationworkspaceworkspaceid) - Returns all activated integrations in a workspace
* [GetUnifiedUser](#getunifieduser) - Retrieve your user object
* [GetUnifiedUserToken](#getunifiedusertoken) - Retrieve your user token
* [GetUnifiedWebhook](#getunifiedwebhook) - Returns all registered webhooks
* [GetUnifiedWebhookID](#getunifiedwebhookid) - Retrieve webhook by its ID
* [PatchUnifiedConnectionID](#patchunifiedconnectionid) - Update connection
* [PatchUnifiedUser](#patchunifieduser) - Updates your user object
* [PostUnifiedConnection](#postunifiedconnection) - Create connection
* [PostUnifiedWebhookConnectionIDObject](#postunifiedwebhookconnectionidobject) - Create webhook subscription
* [PutUnifiedConnectionID](#putunifiedconnectionid) - Update connection
* [PutUnifiedUser](#putunifieduser) - Updates your user object

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
    res, err := s.Unified.DeleteUnifiedConnectionID(ctx, operations.DeleteUnifiedConnectionIDRequest{
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


## DeleteUnifiedUser

Delete your user object

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Unified.DeleteUnifiedUser(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.DeleteUnifiedUserResponse](../../models/operations/deleteunifieduserresponse.md), error**


## DeleteUnifiedWebhookID

Remove webhook subscription

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
    res, err := s.Unified.DeleteUnifiedWebhookID(ctx, operations.DeleteUnifiedWebhookIDRequest{
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.DeleteUnifiedWebhookIDRequest](../../models/operations/deleteunifiedwebhookidrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.DeleteUnifiedWebhookIDResponse](../../models/operations/deleteunifiedwebhookidresponse.md), error**


## GetUnifiedApicall

Returns API Calls

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
    res, err := s.Unified.GetUnifiedApicall(ctx, operations.GetUnifiedApicallRequest{
        ConnectionID: unifiedgosdk.String("delectus green Hybrid"),
        CreatedLte: types.MustTimeFromString("2021-04-02T21:36:49.952Z"),
        Env: unifiedgosdk.String("Fantastic Iodine indexing"),
        Error: unifiedgosdk.Bool(false),
        ExternalXref: unifiedgosdk.String("Music"),
        IntegrationType: unifiedgosdk.String("Soft"),
        Limit: unifiedgosdk.Float64(2390.64),
        Offset: unifiedgosdk.Float64(3757.34),
        Order: unifiedgosdk.String("mobile envisioneer"),
        Sort: unifiedgosdk.String("North payment opposite"),
        UpdatedGte: types.MustTimeFromString("2021-08-11T16:18:13.644Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APICalls != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetUnifiedApicallRequest](../../models/operations/getunifiedapicallrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetUnifiedApicallResponse](../../models/operations/getunifiedapicallresponse.md), error**


## GetUnifiedApicallID

Retrieve specific API Call by its ID

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
    res, err := s.Unified.GetUnifiedApicallID(ctx, operations.GetUnifiedApicallIDRequest{
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APICall != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetUnifiedApicallIDRequest](../../models/operations/getunifiedapicallidrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetUnifiedApicallIDResponse](../../models/operations/getunifiedapicallidresponse.md), error**


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
    res, err := s.Unified.GetUnifiedConnection(ctx, operations.GetUnifiedConnectionRequest{
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
    res, err := s.Unified.GetUnifiedConnectionID(ctx, operations.GetUnifiedConnectionIDRequest{
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


## GetUnifiedIntegration

Returns all integrations

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
    res, err := s.Unified.GetUnifiedIntegration(ctx, operations.GetUnifiedIntegrationRequest{
        Active: unifiedgosdk.Bool(false),
        Categories: []GetUnifiedIntegrationCategories{
            operations.GetUnifiedIntegrationCategoriesEnrich,
        },
        Limit: unifiedgosdk.Float64(7363.95),
        Offset: unifiedgosdk.Float64(8214.4),
        Order: unifiedgosdk.String("Nelda Implemented"),
        Sort: unifiedgosdk.String("cabinet"),
        Summary: unifiedgosdk.Bool(false),
        UpdatedGte: types.MustTimeFromString("2022-02-05T00:16:37.455Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Integrations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetUnifiedIntegrationRequest](../../models/operations/getunifiedintegrationrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.GetUnifiedIntegrationResponse](../../models/operations/getunifiedintegrationresponse.md), error**


## GetUnifiedIntegrationAuthWorkspaceIDIntegrationType

Returns an authorization URL for the specified integration.  Once a successful authorization occurs, a new connection is created.

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
    res, err := s.Unified.GetUnifiedIntegrationAuthWorkspaceIDIntegrationType(ctx, operations.GetUnifiedIntegrationAuthWorkspaceIDIntegrationTypeRequest{
        Env: unifiedgosdk.String("Algerian"),
        ExternalXref: unifiedgosdk.String("Cambridgeshire Surinam"),
        FailureRedirect: unifiedgosdk.String("Designer Drive"),
        IntegrationType: "program Home",
        Lang: unifiedgosdk.String("Plastic program"),
        Redirect: unifiedgosdk.Bool(false),
        Scopes: []GetUnifiedIntegrationAuthWorkspaceIDIntegrationTypeScopes{
            operations.GetUnifiedIntegrationAuthWorkspaceIDIntegrationTypeScopesCrmFileRead,
        },
        State: unifiedgosdk.String("Functionality Product"),
        Subdomain: unifiedgosdk.String("payment Developer Dynamic"),
        SuccessRedirect: unifiedgosdk.String("Northeast"),
        WorkspaceID: "duh empower Kwanza",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetUnifiedIntegrationAuthWorkspaceIDIntegrationType200ApplicationJSONString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                      | Type                                                                                                                                                           | Required                                                                                                                                                       | Description                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                          | :heavy_check_mark:                                                                                                                                             | The context to use for the request.                                                                                                                            |
| `request`                                                                                                                                                      | [operations.GetUnifiedIntegrationAuthWorkspaceIDIntegrationTypeRequest](../../models/operations/getunifiedintegrationauthworkspaceidintegrationtyperequest.md) | :heavy_check_mark:                                                                                                                                             | The request object to use for the request.                                                                                                                     |


### Response

**[*operations.GetUnifiedIntegrationAuthWorkspaceIDIntegrationTypeResponse](../../models/operations/getunifiedintegrationauthworkspaceidintegrationtyperesponse.md), error**


## GetUnifiedIntegrationIntegrationType

Retrieve an integration

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
    res, err := s.Unified.GetUnifiedIntegrationIntegrationType(ctx, operations.GetUnifiedIntegrationIntegrationTypeRequest{
        IntegrationType: "Pizza Electric",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Integration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.GetUnifiedIntegrationIntegrationTypeRequest](../../models/operations/getunifiedintegrationintegrationtyperequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.GetUnifiedIntegrationIntegrationTypeResponse](../../models/operations/getunifiedintegrationintegrationtyperesponse.md), error**


## GetUnifiedIntegrationWorkspaceWorkspaceID

No authentication required as this is to be used by front-end interface

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
    res, err := s.Unified.GetUnifiedIntegrationWorkspaceWorkspaceID(ctx, operations.GetUnifiedIntegrationWorkspaceWorkspaceIDRequest{
        Active: unifiedgosdk.Bool(false),
        Categories: []GetUnifiedIntegrationWorkspaceWorkspaceIDCategories{
            operations.GetUnifiedIntegrationWorkspaceWorkspaceIDCategoriesHris,
        },
        Env: unifiedgosdk.String("North Southeast exercitationem"),
        Summary: unifiedgosdk.Bool(false),
        WorkspaceID: "Bronze Plastic",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Integrations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.GetUnifiedIntegrationWorkspaceWorkspaceIDRequest](../../models/operations/getunifiedintegrationworkspaceworkspaceidrequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |


### Response

**[*operations.GetUnifiedIntegrationWorkspaceWorkspaceIDResponse](../../models/operations/getunifiedintegrationworkspaceworkspaceidresponse.md), error**


## GetUnifiedUser

Retrieve your user object

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Unified.GetUnifiedUser(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetUnifiedUserResponse](../../models/operations/getunifieduserresponse.md), error**


## GetUnifiedUserToken

Retrieve your user token

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Unified.GetUnifiedUserToken(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetUnifiedUserToken200ApplicationJSONString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetUnifiedUserTokenResponse](../../models/operations/getunifiedusertokenresponse.md), error**


## GetUnifiedWebhook

Returns all registered webhooks

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
    res, err := s.Unified.GetUnifiedWebhook(ctx, operations.GetUnifiedWebhookRequest{
        Env: unifiedgosdk.String("Investor methodical Fitness"),
        Limit: unifiedgosdk.Float64(8087.22),
        Object: unifiedgosdk.String("Franc past salmon"),
        Offset: unifiedgosdk.Float64(5240.75),
        Order: unifiedgosdk.String("program"),
        Sort: unifiedgosdk.String("below JSON"),
        UpdatedGte: types.MustTimeFromString("2022-05-29T13:22:55.562Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Webhooks != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetUnifiedWebhookRequest](../../models/operations/getunifiedwebhookrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetUnifiedWebhookResponse](../../models/operations/getunifiedwebhookresponse.md), error**


## GetUnifiedWebhookID

Retrieve webhook by its ID

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
    res, err := s.Unified.GetUnifiedWebhookID(ctx, operations.GetUnifiedWebhookIDRequest{
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Webhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetUnifiedWebhookIDRequest](../../models/operations/getunifiedwebhookidrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetUnifiedWebhookIDResponse](../../models/operations/getunifiedwebhookidresponse.md), error**


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
    res, err := s.Unified.PatchUnifiedConnectionID(ctx, operations.PatchUnifiedConnectionIDRequest{
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


## PatchUnifiedUser

Only the name field is updated.

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
    res, err := s.Unified.PatchUnifiedUser(ctx, shared.User{
        CreatedAt: types.MustTimeFromString("2022-04-24T15:25:24.483Z"),
        Email: unifiedgosdk.String("Emmalee.Quitzon@yahoo.com"),
        Environment: unifiedgosdk.String("Bicycle"),
        ID: unifiedgosdk.String("<ID>"),
        Meta: &shared.PropertyUserMeta{},
        Name: "vice compressing",
        UpdatedAt: types.MustTimeFromString("2023-05-05T16:52:20.023Z"),
        WorkspaceID: "Hybrid methodologies",
        WorkspaceIds: []string{
            "Potassium",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.User](../../models/shared/user.md)            | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PatchUnifiedUserResponse](../../models/operations/patchunifieduserresponse.md), error**


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
    res, err := s.Unified.PostUnifiedConnection(ctx, shared.Connection{
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


## PostUnifiedWebhookConnectionIDObject

To maintain compatibility with the webhooks specification and Zapier webhooks, only the hook_url field is required in the payload when creating a Webhook.  When updated/new objects are found, the array of objects will be POSTed to the hook_url with the paramater hookId=<hookId>.

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
    res, err := s.Unified.PostUnifiedWebhookConnectionIDObject(ctx, operations.PostUnifiedWebhookConnectionIDObjectRequest{
        Webhook: &shared.Webhook{
            CheckedAt: types.MustTimeFromString("2021-02-25T07:12:08.980Z"),
            ConnectionID: "deposit 1080p Passenger",
            CreatedAt: types.MustTimeFromString("2023-02-21T14:58:56.193Z"),
            Environment: unifiedgosdk.String("Minnesota Soap"),
            Events: []shared.PropertyWebhookEvents{
                shared.PropertyWebhookEventsUpdated,
            },
            HookURL: "Table female ken",
            ID: unifiedgosdk.String("<ID>"),
            IncludeRaw: unifiedgosdk.Bool(false),
            IntegrationType: "chocolate",
            Interval: 1710.16,
            ObjectType: shared.WebhookObjectTypeEnrichCompany,
            Subscriptions: []string{
                "female",
            },
            UpdatedAt: types.MustTimeFromString("2022-08-02T17:13:06.397Z"),
            WorkspaceID: "hertz",
        },
        ConnectionID: "Borders",
        Events: []PostUnifiedWebhookConnectionIDObjectEvents{
            operations.PostUnifiedWebhookConnectionIDObjectEventsCreated,
        },
        Object: "scalable",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Webhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.PostUnifiedWebhookConnectionIDObjectRequest](../../models/operations/postunifiedwebhookconnectionidobjectrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.PostUnifiedWebhookConnectionIDObjectResponse](../../models/operations/postunifiedwebhookconnectionidobjectresponse.md), error**


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
    res, err := s.Unified.PutUnifiedConnectionID(ctx, operations.PutUnifiedConnectionIDRequest{
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


## PutUnifiedUser

Only the name field is updated.

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
    res, err := s.Unified.PutUnifiedUser(ctx, shared.User{
        CreatedAt: types.MustTimeFromString("2023-07-31T04:46:29.769Z"),
        Email: unifiedgosdk.String("Selena59@yahoo.com"),
        Environment: unifiedgosdk.String("Bedfordshire Lucia"),
        ID: unifiedgosdk.String("<ID>"),
        Meta: &shared.PropertyUserMeta{},
        Name: "Bicycle hacking South",
        UpdatedAt: types.MustTimeFromString("2023-03-15T15:08:26.238Z"),
        WorkspaceID: "Card defect",
        WorkspaceIds: []string{
            "repudiandae",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.User](../../models/shared/user.md)            | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PutUnifiedUserResponse](../../models/operations/putunifieduserresponse.md), error**

