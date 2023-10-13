# Integration
(*Integration*)

### Available Operations

* [GetUnifiedIntegration](#getunifiedintegration) - Retrieve an integration
* [GetUnifiedIntegrationAuth](#getunifiedintegrationauth) - Create connection indirectly
* [ListUnifiedIntegrationWorkspaces](#listunifiedintegrationworkspaces) - Returns all activated integrations in a workspace
* [ListUnifiedIntegrations](#listunifiedintegrations) - Returns all integrations

## GetUnifiedIntegration

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Integration.GetUnifiedIntegration(ctx, operations.GetUnifiedIntegrationRequest{
        IntegrationType: "Berkelium panel",
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetUnifiedIntegrationRequest](../../models/operations/getunifiedintegrationrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.GetUnifiedIntegrationResponse](../../models/operations/getunifiedintegrationresponse.md), error**


## GetUnifiedIntegrationAuth

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Integration.GetUnifiedIntegrationAuth(ctx, operations.GetUnifiedIntegrationAuthRequest{
        IntegrationType: "Reggae Van pascal",
        Scopes: []operations.GetUnifiedIntegrationAuthScopes{
            operations.GetUnifiedIntegrationAuthScopesAtsScorecardRead,
        },
        WorkspaceID: "Xenogender North groupware",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetUnifiedIntegrationAuth200ApplicationJSONString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetUnifiedIntegrationAuthRequest](../../models/operations/getunifiedintegrationauthrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.GetUnifiedIntegrationAuthResponse](../../models/operations/getunifiedintegrationauthresponse.md), error**


## ListUnifiedIntegrationWorkspaces

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Integration.ListUnifiedIntegrationWorkspaces(ctx, operations.ListUnifiedIntegrationWorkspacesRequest{
        Categories: []operations.ListUnifiedIntegrationWorkspacesCategories{
            operations.ListUnifiedIntegrationWorkspacesCategoriesMartech,
        },
        WorkspaceID: "Country Market Representative",
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

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.ListUnifiedIntegrationWorkspacesRequest](../../models/operations/listunifiedintegrationworkspacesrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


### Response

**[*operations.ListUnifiedIntegrationWorkspacesResponse](../../models/operations/listunifiedintegrationworkspacesresponse.md), error**


## ListUnifiedIntegrations

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
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Integration.ListUnifiedIntegrations(ctx, operations.ListUnifiedIntegrationsRequest{
        Categories: []operations.ListUnifiedIntegrationsCategories{
            operations.ListUnifiedIntegrationsCategoriesAuth,
        },
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListUnifiedIntegrationsRequest](../../models/operations/listunifiedintegrationsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.ListUnifiedIntegrationsResponse](../../models/operations/listunifiedintegrationsresponse.md), error**

