# Unified
(*Unified*)

### Available Operations

* [CreateUnifiedConnection](#createunifiedconnection) - Create connection
* [CreateUnifiedWebhook](#createunifiedwebhook) - Create webhook subscription
* [GetUnifiedApicall](#getunifiedapicall) - Retrieve specific API Call by its ID
* [GetUnifiedConnection](#getunifiedconnection) - Retrieve connection
* [GetUnifiedIntegrationAuth](#getunifiedintegrationauth) - Create connection indirectly
* [GetUnifiedWebhook](#getunifiedwebhook) - Retrieve webhook by its ID
* [ListUnifiedApicalls](#listunifiedapicalls) - Returns API Calls
* [ListUnifiedConnections](#listunifiedconnections) - List all connections
* [ListUnifiedIntegrationWorkspaces](#listunifiedintegrationworkspaces) - Returns all activated integrations in a workspace
* [ListUnifiedIntegrations](#listunifiedintegrations) - Returns all integrations
* [ListUnifiedIssues](#listunifiedissues) - List support issues
* [ListUnifiedSupports](#listunifiedsupports) - Get support info
* [ListUnifiedWebhooks](#listunifiedwebhooks) - Returns all registered webhooks
* [PatchUnifiedConnection](#patchunifiedconnection) - Update connection
* [PatchUnifiedWebhookTrigger](#patchunifiedwebhooktrigger) - Trigger webhook
* [RemoveUnifiedConnection](#removeunifiedconnection) - Remove connection
* [RemoveUnifiedWebhook](#removeunifiedwebhook) - Remove webhook subscription
* [UpdateUnifiedConnection](#updateunifiedconnection) - Update connection
* [UpdateUnifiedWebhookTrigger](#updateunifiedwebhooktrigger) - Trigger webhook

## CreateUnifiedConnection

Create connection

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Unified.CreateUnifiedConnection(ctx, &shared.Connection{
        Categories: []shared.PropertyConnectionCategories{
            shared.PropertyConnectionCategoriesCommerce,
        },
        IntegrationType: "<value>",
        Permissions: []shared.PropertyConnectionPermissions{
            shared.PropertyConnectionPermissionsAtsCandidateWrite,
        },
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

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `request`                                                  | [shared.Connection](../../pkg/models/shared/connection.md) | :heavy_check_mark:                                         | The request object to use for the request.                 |


### Response

**[*operations.CreateUnifiedConnectionResponse](../../pkg/models/operations/createunifiedconnectionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateUnifiedWebhook

The data payload received by your server is described at https://docs.unified.to/unified/overview. The `interval` field can be set as low as 1 minute for paid accounts, and 60 minutes for free accounts.

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Unified.CreateUnifiedWebhook(ctx, operations.CreateUnifiedWebhookRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Webhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateUnifiedWebhookRequest](../../pkg/models/operations/createunifiedwebhookrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.CreateUnifiedWebhookResponse](../../pkg/models/operations/createunifiedwebhookresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetUnifiedApicall

Retrieve specific API Call by its ID

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Unified.GetUnifiedApicall(ctx, operations.GetUnifiedApicallRequest{
        ID: "<id>",
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
| `request`                                                                                      | [operations.GetUnifiedApicallRequest](../../pkg/models/operations/getunifiedapicallrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetUnifiedApicallResponse](../../pkg/models/operations/getunifiedapicallresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetUnifiedConnection

Retrieve connection

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Unified.GetUnifiedConnection(ctx, operations.GetUnifiedConnectionRequest{
        ID: "<id>",
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
| `request`                                                                                            | [operations.GetUnifiedConnectionRequest](../../pkg/models/operations/getunifiedconnectionrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.GetUnifiedConnectionResponse](../../pkg/models/operations/getunifiedconnectionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetUnifiedIntegrationAuth

Returns an authorization URL for the specified integration.  Once a successful authorization occurs, a new connection is created.

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Unified.GetUnifiedIntegrationAuth(ctx, operations.GetUnifiedIntegrationAuthRequest{
        IntegrationType: "<value>",
        WorkspaceID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetUnifiedIntegrationAuthRequest](../../pkg/models/operations/getunifiedintegrationauthrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.GetUnifiedIntegrationAuthResponse](../../pkg/models/operations/getunifiedintegrationauthresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetUnifiedWebhook

Retrieve webhook by its ID

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Unified.GetUnifiedWebhook(ctx, operations.GetUnifiedWebhookRequest{
        ID: "<id>",
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
| `request`                                                                                      | [operations.GetUnifiedWebhookRequest](../../pkg/models/operations/getunifiedwebhookrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetUnifiedWebhookResponse](../../pkg/models/operations/getunifiedwebhookresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListUnifiedApicalls

Returns API Calls

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Unified.ListUnifiedApicalls(ctx, operations.ListUnifiedApicallsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.APICalls != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListUnifiedApicallsRequest](../../pkg/models/operations/listunifiedapicallsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.ListUnifiedApicallsResponse](../../pkg/models/operations/listunifiedapicallsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListUnifiedConnections

List all connections

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Unified.ListUnifiedConnections(ctx, operations.ListUnifiedConnectionsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Connections != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListUnifiedConnectionsRequest](../../pkg/models/operations/listunifiedconnectionsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.ListUnifiedConnectionsResponse](../../pkg/models/operations/listunifiedconnectionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListUnifiedIntegrationWorkspaces

No authentication required as this is to be used by front-end interface

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Unified.ListUnifiedIntegrationWorkspaces(ctx, operations.ListUnifiedIntegrationWorkspacesRequest{
        WorkspaceID: "<value>",
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

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.ListUnifiedIntegrationWorkspacesRequest](../../pkg/models/operations/listunifiedintegrationworkspacesrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.ListUnifiedIntegrationWorkspacesResponse](../../pkg/models/operations/listunifiedintegrationworkspacesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListUnifiedIntegrations

Returns all integrations

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Unified.ListUnifiedIntegrations(ctx, operations.ListUnifiedIntegrationsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Integrations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ListUnifiedIntegrationsRequest](../../pkg/models/operations/listunifiedintegrationsrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.ListUnifiedIntegrationsResponse](../../pkg/models/operations/listunifiedintegrationsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListUnifiedIssues

List support issues

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Unified.ListUnifiedIssues(ctx, operations.ListUnifiedIssuesRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Issues != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListUnifiedIssuesRequest](../../pkg/models/operations/listunifiedissuesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.ListUnifiedIssuesResponse](../../pkg/models/operations/listunifiedissuesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListUnifiedSupports

Get support info

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Unified.ListUnifiedSupports(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Undefined != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListUnifiedSupportsResponse](../../pkg/models/operations/listunifiedsupportsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListUnifiedWebhooks

Returns all registered webhooks

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Unified.ListUnifiedWebhooks(ctx, operations.ListUnifiedWebhooksRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Webhooks != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListUnifiedWebhooksRequest](../../pkg/models/operations/listunifiedwebhooksrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.ListUnifiedWebhooksResponse](../../pkg/models/operations/listunifiedwebhooksresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PatchUnifiedConnection

Update connection

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Unified.PatchUnifiedConnection(ctx, operations.PatchUnifiedConnectionRequest{
        ID: "<id>",
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
| `request`                                                                                                | [operations.PatchUnifiedConnectionRequest](../../pkg/models/operations/patchunifiedconnectionrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.PatchUnifiedConnectionResponse](../../pkg/models/operations/patchunifiedconnectionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PatchUnifiedWebhookTrigger

Trigger webhook

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Unified.PatchUnifiedWebhookTrigger(ctx, operations.PatchUnifiedWebhookTriggerRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.PatchUnifiedWebhookTriggerRequest](../../pkg/models/operations/patchunifiedwebhooktriggerrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.PatchUnifiedWebhookTriggerResponse](../../pkg/models/operations/patchunifiedwebhooktriggerresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RemoveUnifiedConnection

Remove connection

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Unified.RemoveUnifiedConnection(ctx, operations.RemoveUnifiedConnectionRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.RemoveUnifiedConnectionRequest](../../pkg/models/operations/removeunifiedconnectionrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.RemoveUnifiedConnectionResponse](../../pkg/models/operations/removeunifiedconnectionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RemoveUnifiedWebhook

Remove webhook subscription

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Unified.RemoveUnifiedWebhook(ctx, operations.RemoveUnifiedWebhookRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.RemoveUnifiedWebhookRequest](../../pkg/models/operations/removeunifiedwebhookrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.RemoveUnifiedWebhookResponse](../../pkg/models/operations/removeunifiedwebhookresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateUnifiedConnection

Update connection

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Unified.UpdateUnifiedConnection(ctx, operations.UpdateUnifiedConnectionRequest{
        ID: "<id>",
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.UpdateUnifiedConnectionRequest](../../pkg/models/operations/updateunifiedconnectionrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.UpdateUnifiedConnectionResponse](../../pkg/models/operations/updateunifiedconnectionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateUnifiedWebhookTrigger

Trigger webhook

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Unified.UpdateUnifiedWebhookTrigger(ctx, operations.UpdateUnifiedWebhookTriggerRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.UpdateUnifiedWebhookTriggerRequest](../../pkg/models/operations/updateunifiedwebhooktriggerrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.UpdateUnifiedWebhookTriggerResponse](../../pkg/models/operations/updateunifiedwebhooktriggerresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
