# Integration

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
    res, err := s.Integration.GetUnifiedIntegration(ctx, operations.GetUnifiedIntegrationRequest{
        Active: unifiedto.Bool(false),
        Categories: []GetUnifiedIntegrationCategories{
            operations.GetUnifiedIntegrationCategoriesHris,
        },
        Limit: unifiedto.Float64(2394.27),
        Offset: unifiedto.Float64(8487.85),
        Order: unifiedto.String("ea"),
        Sort: unifiedto.String("veniam"),
        Summary: unifiedto.Bool(false),
        UpdatedGte: types.MustDateFromString("2021-01-26"),
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
    res, err := s.Integration.GetUnifiedIntegrationAuthWorkspaceIDIntegrationType(ctx, operations.GetUnifiedIntegrationAuthWorkspaceIDIntegrationTypeRequest{
        Env: unifiedto.String("earum"),
        ExternalXref: unifiedto.String("placeat"),
        FailureRedirect: unifiedto.String("saepe"),
        IntegrationType: "quod",
        Lang: unifiedto.String("odit"),
        Redirect: unifiedto.Bool(false),
        Scopes: []GetUnifiedIntegrationAuthWorkspaceIDIntegrationTypeScopes{
            operations.GetUnifiedIntegrationAuthWorkspaceIDIntegrationTypeScopesHrisGroupWrite,
        },
        State: unifiedto.String("at"),
        Subdomain: unifiedto.String("ea"),
        SuccessRedirect: unifiedto.String("provident"),
        WorkspaceID: "inventore",
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
    res, err := s.Integration.GetUnifiedIntegrationIntegrationType(ctx, operations.GetUnifiedIntegrationIntegrationTypeRequest{
        IntegrationType: "ea",
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
    res, err := s.Integration.GetUnifiedIntegrationWorkspaceWorkspaceID(ctx, operations.GetUnifiedIntegrationWorkspaceWorkspaceIDRequest{
        Active: unifiedto.Bool(false),
        Categories: []GetUnifiedIntegrationWorkspaceWorkspaceIDCategories{
            operations.GetUnifiedIntegrationWorkspaceWorkspaceIDCategoriesUc,
        },
        Env: unifiedto.String("quam"),
        Summary: unifiedto.Bool(false),
        WorkspaceID: "delectus",
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

