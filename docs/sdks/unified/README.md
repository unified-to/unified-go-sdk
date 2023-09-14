# Unified

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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.DeleteUnifiedConnectionIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Unified.DeleteUnifiedConnectionID(ctx, operations.DeleteUnifiedConnectionIDRequest{
        ID: "d2a7c7d1-ea0e-479f-a9bb-e5f179f650b1",
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.DeleteUnifiedConnectionIDRequest](../../models/operations/deleteunifiedconnectionidrequest.md)   | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `security`                                                                                                   | [operations.DeleteUnifiedConnectionIDSecurity](../../models/operations/deleteunifiedconnectionidsecurity.md) | :heavy_check_mark:                                                                                           | The security requirements to use for the request.                                                            |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.DeleteUnifiedUserSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Unified.DeleteUnifiedUser(ctx, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `security`                                                                                   | [operations.DeleteUnifiedUserSecurity](../../models/operations/deleteunifiedusersecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.DeleteUnifiedWebhookIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Unified.DeleteUnifiedWebhookID(ctx, operations.DeleteUnifiedWebhookIDRequest{
        ID: "e707e7e4-3967-413b-acce-072abd61918d",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.DeleteUnifiedWebhookIDRequest](../../models/operations/deleteunifiedwebhookidrequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.DeleteUnifiedWebhookIDSecurity](../../models/operations/deleteunifiedwebhookidsecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetUnifiedApicallSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Unified.GetUnifiedApicall(ctx, operations.GetUnifiedApicallRequest{
        ConnectionID: unifiedto.String("magni"),
        CreatedLte: types.MustTimeFromString("2022-05-24T09:06:17.139Z"),
        Env: unifiedto.String("maxime"),
        Error: unifiedto.Bool(false),
        ExternalXref: unifiedto.String("vitae"),
        IntegrationType: unifiedto.String("alias"),
        Limit: unifiedto.Float64(8070.07),
        Offset: unifiedto.Float64(1150.28),
        Order: unifiedto.String("blanditiis"),
        Sort: unifiedto.String("ipsam"),
        UpdatedGte: types.MustTimeFromString("2022-08-08T20:31:22.144Z"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.APICalls != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetUnifiedApicallRequest](../../models/operations/getunifiedapicallrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.GetUnifiedApicallSecurity](../../models/operations/getunifiedapicallsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetUnifiedApicallIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Unified.GetUnifiedApicallID(ctx, operations.GetUnifiedApicallIDRequest{
        ID: "fd78be26-2127-4262-8fa5-03962867e72b",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.APICall != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetUnifiedApicallIDRequest](../../models/operations/getunifiedapicallidrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.GetUnifiedApicallIDSecurity](../../models/operations/getunifiedapicallidsecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetUnifiedConnectionSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Unified.GetUnifiedConnection(ctx, operations.GetUnifiedConnectionRequest{
        Categories: []GetUnifiedConnectionCategories{
            operations.GetUnifiedConnectionCategoriesHris,
        },
        Env: unifiedto.String("laborum"),
        ExternalXref: unifiedto.String("autem"),
        Limit: unifiedto.Float64(3273.73),
        Offset: unifiedto.Float64(603.93),
        Order: unifiedto.String("qui"),
        Sort: unifiedto.String("labore"),
        UpdatedGte: types.MustTimeFromString("2022-11-05T10:44:25.961Z"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Connections != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetUnifiedConnectionRequest](../../models/operations/getunifiedconnectionrequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.GetUnifiedConnectionSecurity](../../models/operations/getunifiedconnectionsecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetUnifiedConnectionIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Unified.GetUnifiedConnectionID(ctx, operations.GetUnifiedConnectionIDRequest{
        ID: "57f9bb6e-f72a-4508-b1d9-9b661a7def16",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetUnifiedConnectionIDRequest](../../models/operations/getunifiedconnectionidrequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.GetUnifiedConnectionIDSecurity](../../models/operations/getunifiedconnectionidsecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetUnifiedIntegrationSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Unified.GetUnifiedIntegration(ctx, operations.GetUnifiedIntegrationRequest{
        Active: unifiedto.Bool(false),
        Categories: []GetUnifiedIntegrationCategories{
            operations.GetUnifiedIntegrationCategoriesCrm,
        },
        Limit: unifiedto.Float64(7216.29),
        Offset: unifiedto.Float64(4263.23),
        Order: unifiedto.String("impedit"),
        Sort: unifiedto.String("optio"),
        Summary: unifiedto.Bool(false),
        UpdatedGte: types.MustTimeFromString("2022-09-12T15:50:36.450Z"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Integrations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetUnifiedIntegrationRequest](../../models/operations/getunifiedintegrationrequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.GetUnifiedIntegrationSecurity](../../models/operations/getunifiedintegrationsecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()

    ctx := context.Background()
    res, err := s.Unified.GetUnifiedIntegrationAuthWorkspaceIDIntegrationType(ctx, operations.GetUnifiedIntegrationAuthWorkspaceIDIntegrationTypeRequest{
        Env: unifiedto.String("deleniti"),
        ExternalXref: unifiedto.String("dolores"),
        FailureRedirect: unifiedto.String("dolores"),
        IntegrationType: "distinctio",
        Lang: unifiedto.String("modi"),
        Redirect: unifiedto.Bool(false),
        Scopes: []GetUnifiedIntegrationAuthWorkspaceIDIntegrationTypeScopes{
            operations.GetUnifiedIntegrationAuthWorkspaceIDIntegrationTypeScopesTicketingTicketRead,
        },
        State: unifiedto.String("perspiciatis"),
        Subdomain: unifiedto.String("totam"),
        SuccessRedirect: unifiedto.String("ipsam"),
        WorkspaceID: "alias",
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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetUnifiedIntegrationIntegrationTypeSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Unified.GetUnifiedIntegrationIntegrationType(ctx, operations.GetUnifiedIntegrationIntegrationTypeRequest{
        IntegrationType: "repudiandae",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Integration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.GetUnifiedIntegrationIntegrationTypeRequest](../../models/operations/getunifiedintegrationintegrationtyperequest.md)   | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `security`                                                                                                                         | [operations.GetUnifiedIntegrationIntegrationTypeSecurity](../../models/operations/getunifiedintegrationintegrationtypesecurity.md) | :heavy_check_mark:                                                                                                                 | The security requirements to use for the request.                                                                                  |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()

    ctx := context.Background()
    res, err := s.Unified.GetUnifiedIntegrationWorkspaceWorkspaceID(ctx, operations.GetUnifiedIntegrationWorkspaceWorkspaceIDRequest{
        Active: unifiedto.Bool(false),
        Categories: []GetUnifiedIntegrationWorkspaceWorkspaceIDCategories{
            operations.GetUnifiedIntegrationWorkspaceWorkspaceIDCategoriesTicketing,
        },
        Env: unifiedto.String("magni"),
        Summary: unifiedto.Bool(false),
        WorkspaceID: "doloribus",
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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetUnifiedUserSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Unified.GetUnifiedUser(ctx, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `security`                                                                             | [operations.GetUnifiedUserSecurity](../../models/operations/getunifiedusersecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetUnifiedUserTokenSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Unified.GetUnifiedUserToken(ctx, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetUnifiedUserToken200ApplicationJSONString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `security`                                                                                       | [operations.GetUnifiedUserTokenSecurity](../../models/operations/getunifiedusertokensecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetUnifiedWebhookSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Unified.GetUnifiedWebhook(ctx, operations.GetUnifiedWebhookRequest{
        Env: unifiedto.String("dolore"),
        Limit: unifiedto.Float64(6674.44),
        Object: unifiedto.String("veritatis"),
        Offset: unifiedto.Float64(9332.28),
        Order: unifiedto.String("excepturi"),
        Sort: unifiedto.String("eligendi"),
        UpdatedGte: types.MustTimeFromString("2022-04-26T06:44:29.423Z"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Webhooks != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetUnifiedWebhookRequest](../../models/operations/getunifiedwebhookrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.GetUnifiedWebhookSecurity](../../models/operations/getunifiedwebhooksecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetUnifiedWebhookIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Unified.GetUnifiedWebhookID(ctx, operations.GetUnifiedWebhookIDRequest{
        ID: "e55140e7-5726-4e00-bc2f-0294192518ce",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Webhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetUnifiedWebhookIDRequest](../../models/operations/getunifiedwebhookidrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.GetUnifiedWebhookIDSecurity](../../models/operations/getunifiedwebhookidsecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PatchUnifiedConnectionIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Unified.PatchUnifiedConnectionID(ctx, operations.PatchUnifiedConnectionIDRequest{
        Connection: &shared.Connection{
            Auth: &shared.PropertyConnectionAuth{
                AccessToken: unifiedto.String("accusamus"),
                APIURL: unifiedto.String("incidunt"),
                AppID: unifiedto.String("dicta"),
                AuthorizeURL: unifiedto.String("quo"),
                ClientID: unifiedto.String("natus"),
                ClientSecret: unifiedto.String("excepturi"),
                ConsumerKey: unifiedto.String("natus"),
                ConsumerSecret: unifiedto.String("hic"),
                Emails: []string{
                    "ut",
                },
                ExpiresIn: unifiedto.Float64(3924.24),
                ExpiryDate: types.MustTimeFromString("2021-01-28T23:49:19.065Z"),
                Key: unifiedto.String("eum"),
                Meta: &shared.PropertyPropertyConnectionAuthMeta{},
                Name: unifiedto.String("Albert Schmidt MD"),
                OtherAuthInfo: []string{
                    "adipisci",
                },
                Pem: unifiedto.String("a"),
                RefreshToken: unifiedto.String("ipsa"),
                State: unifiedto.String("sed"),
                Token: unifiedto.String("sequi"),
                TokenURL: unifiedto.String("minus"),
            },
            AuthAwsArn: unifiedto.String("suscipit"),
            Categories: []shared.PropertyConnectionCategories{
                shared.PropertyConnectionCategoriesAuth,
            },
            CreatedAt: types.MustTimeFromString("2021-02-21T07:53:55.382Z"),
            Environment: unifiedto.String("laboriosam"),
            ExternalXref: unifiedto.String("harum"),
            ID: unifiedto.String("626012eb-a057-4988-8672-0c3103f1a40c"),
            IntegrationType: "doloremque",
            IsPaused: unifiedto.Bool(false),
            Permissions: []shared.PropertyConnectionPermissions{
                shared.PropertyConnectionPermissionsAtsJobRead,
            },
            UpdatedAt: types.MustTimeFromString("2022-02-08T06:53:00.234Z"),
            WorkspaceID: unifiedto.String("quo"),
        },
        ID: "8688fd8e-c6fc-4031-a8f0-aaaeee004eba",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.PatchUnifiedConnectionIDRequest](../../models/operations/patchunifiedconnectionidrequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.PatchUnifiedConnectionIDSecurity](../../models/operations/patchunifiedconnectionidsecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PatchUnifiedUserSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Unified.PatchUnifiedUser(ctx, shared.User{
        CreatedAt: types.MustTimeFromString("2022-04-05T01:32:39.778Z"),
        Email: unifiedto.String("Kavon12@hotmail.com"),
        Environment: unifiedto.String("harum"),
        ID: unifiedto.String("e509c508-7131-4f06-b0bc-e55a8687143c"),
        Meta: &shared.PropertyUserMeta{},
        Name: "Mrs. Darryl Morar",
        UpdatedAt: types.MustTimeFromString("2021-07-05T20:55:18.206Z"),
        WorkspaceID: "sint",
        WorkspaceIds: []string{
            "odio",
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [shared.User](../../models/shared/user.md)                                                 | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.PatchUnifiedUserSecurity](../../models/operations/patchunifiedusersecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PostUnifiedConnectionSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Unified.PostUnifiedConnection(ctx, shared.Connection{
        Auth: &shared.PropertyConnectionAuth{
            AccessToken: unifiedto.String("animi"),
            APIURL: unifiedto.String("exercitationem"),
            AppID: unifiedto.String("repellendus"),
            AuthorizeURL: unifiedto.String("culpa"),
            ClientID: unifiedto.String("vel"),
            ClientSecret: unifiedto.String("ex"),
            ConsumerKey: unifiedto.String("non"),
            ConsumerSecret: unifiedto.String("nobis"),
            Emails: []string{
                "in",
            },
            ExpiresIn: unifiedto.Float64(8765.83),
            ExpiryDate: types.MustTimeFromString("2022-07-06T05:57:07.995Z"),
            Key: unifiedto.String("voluptatum"),
            Meta: &shared.PropertyPropertyConnectionAuthMeta{},
            Name: unifiedto.String("Neil Grimes"),
            OtherAuthInfo: []string{
                "culpa",
            },
            Pem: unifiedto.String("culpa"),
            RefreshToken: unifiedto.String("odit"),
            State: unifiedto.String("laudantium"),
            Token: unifiedto.String("dolor"),
            TokenURL: unifiedto.String("consequuntur"),
        },
        AuthAwsArn: unifiedto.String("libero"),
        Categories: []shared.PropertyConnectionCategories{
            shared.PropertyConnectionCategoriesMartech,
        },
        CreatedAt: types.MustTimeFromString("2022-09-06T07:44:31.963Z"),
        Environment: unifiedto.String("totam"),
        ExternalXref: unifiedto.String("laboriosam"),
        ID: unifiedto.String("2d2a31f9-b14a-4a6b-9ec7-f444232e9a5d"),
        IntegrationType: "eveniet",
        IsPaused: unifiedto.Bool(false),
        Permissions: []shared.PropertyConnectionPermissions{
            shared.PropertyConnectionPermissionsAtsInterviewRead,
        },
        UpdatedAt: types.MustTimeFromString("2022-05-01T01:22:05.452Z"),
        WorkspaceID: unifiedto.String("cumque"),
    }, operationSecurity)
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
| `request`                                                                                            | [shared.Connection](../../models/shared/connection.md)                                               | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.PostUnifiedConnectionSecurity](../../models/operations/postunifiedconnectionsecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PostUnifiedWebhookConnectionIDObjectSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Unified.PostUnifiedWebhookConnectionIDObject(ctx, operations.PostUnifiedWebhookConnectionIDObjectRequest{
        Webhook: &shared.Webhook{
            CheckedAt: types.MustTimeFromString("2021-07-13T18:24:14.540Z"),
            ConnectionID: "sed",
            CreatedAt: types.MustTimeFromString("2021-12-08T04:18:23.721Z"),
            Environment: unifiedto.String("cupiditate"),
            Events: []shared.PropertyWebhookEvents{
                shared.PropertyWebhookEventsCreated,
            },
            HookURL: "voluptatum",
            ID: unifiedto.String("1b58fe68-2e1c-42db-a23d-58e8247d122c"),
            IncludeRaw: unifiedto.Bool(false),
            IntegrationType: "provident",
            Interval: 9868.06,
            ObjectType: shared.WebhookObjectTypeCrmLead,
            Subscriptions: []string{
                "iusto",
            },
            UpdatedAt: types.MustTimeFromString("2022-07-12T13:27:28.756Z"),
            WorkspaceID: "praesentium",
        },
        ConnectionID: "maiores",
        Events: []PostUnifiedWebhookConnectionIDObjectEvents{
            operations.PostUnifiedWebhookConnectionIDObjectEventsCreated,
        },
        Object: "dolores",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Webhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.PostUnifiedWebhookConnectionIDObjectRequest](../../models/operations/postunifiedwebhookconnectionidobjectrequest.md)   | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `security`                                                                                                                         | [operations.PostUnifiedWebhookConnectionIDObjectSecurity](../../models/operations/postunifiedwebhookconnectionidobjectsecurity.md) | :heavy_check_mark:                                                                                                                 | The security requirements to use for the request.                                                                                  |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PutUnifiedConnectionIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Unified.PutUnifiedConnectionID(ctx, operations.PutUnifiedConnectionIDRequest{
        Connection: &shared.Connection{
            Auth: &shared.PropertyConnectionAuth{
                AccessToken: unifiedto.String("ducimus"),
                APIURL: unifiedto.String("occaecati"),
                AppID: unifiedto.String("nostrum"),
                AuthorizeURL: unifiedto.String("atque"),
                ClientID: unifiedto.String("sequi"),
                ClientSecret: unifiedto.String("commodi"),
                ConsumerKey: unifiedto.String("quam"),
                ConsumerSecret: unifiedto.String("dolor"),
                Emails: []string{
                    "voluptas",
                },
                ExpiresIn: unifiedto.Float64(2226.69),
                ExpiryDate: types.MustTimeFromString("2020-12-30T11:42:07.241Z"),
                Key: unifiedto.String("aut"),
                Meta: &shared.PropertyPropertyConnectionAuthMeta{},
                Name: unifiedto.String("Velma Baumbach"),
                OtherAuthInfo: []string{
                    "doloribus",
                },
                Pem: unifiedto.String("deserunt"),
                RefreshToken: unifiedto.String("officiis"),
                State: unifiedto.String("nam"),
                Token: unifiedto.String("totam"),
                TokenURL: unifiedto.String("ex"),
            },
            AuthAwsArn: unifiedto.String("labore"),
            Categories: []shared.PropertyConnectionCategories{
                shared.PropertyConnectionCategoriesCrm,
            },
            CreatedAt: types.MustTimeFromString("2022-07-24T05:34:55.315Z"),
            Environment: unifiedto.String("adipisci"),
            ExternalXref: unifiedto.String("voluptatem"),
            ID: unifiedto.String("d8f8b89d-9ca6-4075-a56f-c0ebe67155e2"),
            IntegrationType: "assumenda",
            IsPaused: unifiedto.Bool(false),
            Permissions: []shared.PropertyConnectionPermissions{
                shared.PropertyConnectionPermissionsAtsScorecardWrite,
            },
            UpdatedAt: types.MustTimeFromString("2022-04-26T12:57:03.822Z"),
            WorkspaceID: unifiedto.String("ipsum"),
        },
        ID: "070d6e29-7f58-41fa-baaa-7d801088076f",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.PutUnifiedConnectionIDRequest](../../models/operations/putunifiedconnectionidrequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.PutUnifiedConnectionIDSecurity](../../models/operations/putunifiedconnectionidsecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PutUnifiedUserSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Unified.PutUnifiedUser(ctx, shared.User{
        CreatedAt: types.MustTimeFromString("2021-12-05T11:57:07.239Z"),
        Email: unifiedto.String("Gillian.Walsh62@hotmail.com"),
        Environment: unifiedto.String("laudantium"),
        ID: unifiedto.String("14088269-b6a7-40b0-9d82-f94fffbd1e1e"),
        Meta: &shared.PropertyUserMeta{},
        Name: "Debra Stiedemann",
        UpdatedAt: types.MustTimeFromString("2022-06-05T07:44:58.092Z"),
        WorkspaceID: "doloremque",
        WorkspaceIds: []string{
            "sequi",
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [shared.User](../../models/shared/user.md)                                             | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.PutUnifiedUserSecurity](../../models/operations/putunifiedusersecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.PutUnifiedUserResponse](../../models/operations/putunifieduserresponse.md), error**

