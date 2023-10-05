# Integration
(*Integration*)

### Available Operations

* [GetUnifiedIntegration](#getunifiedintegration) - Returns all integrations
* [GetUnifiedIntegrationAuthWorkspaceIDIntegrationType](#getunifiedintegrationauthworkspaceidintegrationtype) - Create connection indirectly
* [GetUnifiedIntegrationIntegrationType](#getunifiedintegrationintegrationtype) - Retrieve an integration
* [GetUnifiedIntegrationWorkspaceWorkspaceID](#getunifiedintegrationworkspaceworkspaceid) - Returns all activated integrations in a workspace

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
    res, err := s.Integration.GetUnifiedIntegration(ctx, operations.GetUnifiedIntegrationRequest{
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
    res, err := s.Integration.GetUnifiedIntegrationAuthWorkspaceIDIntegrationType(ctx, operations.GetUnifiedIntegrationAuthWorkspaceIDIntegrationTypeRequest{
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
    res, err := s.Integration.GetUnifiedIntegrationIntegrationType(ctx, operations.GetUnifiedIntegrationIntegrationTypeRequest{
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
    res, err := s.Integration.GetUnifiedIntegrationWorkspaceWorkspaceID(ctx, operations.GetUnifiedIntegrationWorkspaceWorkspaceIDRequest{
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
