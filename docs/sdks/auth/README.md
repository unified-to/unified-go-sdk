# Auth

### Available Operations

* [GetUnifiedIntegrationAuthWorkspaceIDIntegrationType](#getunifiedintegrationauthworkspaceidintegrationtype) - Create connection indirectly
* [GetUnifiedIntegrationLoginWorkspaceIDIntegrationType](#getunifiedintegrationloginworkspaceidintegrationtype) - Sign in a user

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
    res, err := s.Auth.GetUnifiedIntegrationAuthWorkspaceIDIntegrationType(ctx, operations.GetUnifiedIntegrationAuthWorkspaceIDIntegrationTypeRequest{
        Env: unifiedto.String("beatae"),
        ExternalXref: unifiedto.String("aliquid"),
        FailureRedirect: unifiedto.String("modi"),
        IntegrationType: "optio",
        Lang: unifiedto.String("voluptatibus"),
        Redirect: unifiedto.Bool(false),
        Scopes: []GetUnifiedIntegrationAuthWorkspaceIDIntegrationTypeScopes{
            operations.GetUnifiedIntegrationAuthWorkspaceIDIntegrationTypeScopesMartechMemberRead,
        },
        State: unifiedto.String("officia"),
        Subdomain: unifiedto.String("libero"),
        SuccessRedirect: unifiedto.String("totam"),
        WorkspaceID: "sequi",
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


## GetUnifiedIntegrationLoginWorkspaceIDIntegrationType

Returns an authentication URL for the specified integration.  Once a successful authentication occurs, the name and emails are returned.

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
    res, err := s.Auth.GetUnifiedIntegrationLoginWorkspaceIDIntegrationType(ctx, operations.GetUnifiedIntegrationLoginWorkspaceIDIntegrationTypeRequest{
        Env: unifiedto.String("aliquid"),
        FailureRedirect: unifiedto.String("ea"),
        IntegrationType: "impedit",
        Redirect: unifiedto.Bool(false),
        State: unifiedto.String("ducimus"),
        SuccessRedirect: unifiedto.String("odit"),
        WorkspaceID: "velit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetUnifiedIntegrationLoginWorkspaceIDIntegrationType200ApplicationJSONString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                        | Type                                                                                                                                                             | Required                                                                                                                                                         | Description                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                            | :heavy_check_mark:                                                                                                                                               | The context to use for the request.                                                                                                                              |
| `request`                                                                                                                                                        | [operations.GetUnifiedIntegrationLoginWorkspaceIDIntegrationTypeRequest](../../models/operations/getunifiedintegrationloginworkspaceidintegrationtyperequest.md) | :heavy_check_mark:                                                                                                                                               | The request object to use for the request.                                                                                                                       |


### Response

**[*operations.GetUnifiedIntegrationLoginWorkspaceIDIntegrationTypeResponse](../../models/operations/getunifiedintegrationloginworkspaceidintegrationtyperesponse.md), error**

