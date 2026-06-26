# Secretsmanager

## Overview

### Available Operations

* [CreateUnifiedWorkspaceSecretsmanager](#createunifiedworkspacesecretsmanager) - Create secrets manager
* [GetUnifiedWorkspaceSecretsmanager](#getunifiedworkspacesecretsmanager) - Retrieve secrets manager
* [ListUnifiedWorkspaceSecretsmanagers](#listunifiedworkspacesecretsmanagers) - List secrets managers
* [RemoveUnifiedWorkspaceSecretsmanager](#removeunifiedworkspacesecretsmanager) - Remove secrets manager

## CreateUnifiedWorkspaceSecretsmanager

Create secrets manager

### Example Usage

<!-- UsageSnippet language="go" operationID="createUnifiedWorkspaceSecretsmanager" method="post" path="/unified/workspace/secretsmanager" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Secretsmanager.CreateUnifiedWorkspaceSecretsmanager(ctx, shared.SecretsManager{
        Auth: map[string]string{

        },
        Name: "<value>",
        Type: shared.SecretsManagerTypeHashicorp,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SecretsManager != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [shared.SecretsManager](../../pkg/models/shared/secretsmanager.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |
| `opts`                                                             | [][operations.Option](../../pkg/models/operations/option.md)       | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.CreateUnifiedWorkspaceSecretsmanagerResponse](../../pkg/models/operations/createunifiedworkspacesecretsmanagerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetUnifiedWorkspaceSecretsmanager

Retrieve secrets manager

### Example Usage

<!-- UsageSnippet language="go" operationID="getUnifiedWorkspaceSecretsmanager" method="get" path="/unified/workspace/secretsmanager/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Secretsmanager.GetUnifiedWorkspaceSecretsmanager(ctx, operations.GetUnifiedWorkspaceSecretsmanagerRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SecretsManager != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.GetUnifiedWorkspaceSecretsmanagerRequest](../../pkg/models/operations/getunifiedworkspacesecretsmanagerrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.GetUnifiedWorkspaceSecretsmanagerResponse](../../pkg/models/operations/getunifiedworkspacesecretsmanagerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListUnifiedWorkspaceSecretsmanagers

List secrets managers

### Example Usage

<!-- UsageSnippet language="go" operationID="listUnifiedWorkspaceSecretsmanagers" method="get" path="/unified/workspace/secretsmanager" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Secretsmanager.ListUnifiedWorkspaceSecretsmanagers(ctx, operations.ListUnifiedWorkspaceSecretsmanagersRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.SecretsManagers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.ListUnifiedWorkspaceSecretsmanagersRequest](../../pkg/models/operations/listunifiedworkspacesecretsmanagersrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.ListUnifiedWorkspaceSecretsmanagersResponse](../../pkg/models/operations/listunifiedworkspacesecretsmanagersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveUnifiedWorkspaceSecretsmanager

Remove secrets manager

### Example Usage

<!-- UsageSnippet language="go" operationID="removeUnifiedWorkspaceSecretsmanager" method="delete" path="/unified/workspace/secretsmanager/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Secretsmanager.RemoveUnifiedWorkspaceSecretsmanager(ctx, operations.RemoveUnifiedWorkspaceSecretsmanagerRequest{
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

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.RemoveUnifiedWorkspaceSecretsmanagerRequest](../../pkg/models/operations/removeunifiedworkspacesecretsmanagerrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |
| `opts`                                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |

### Response

**[*operations.RemoveUnifiedWorkspaceSecretsmanagerResponse](../../pkg/models/operations/removeunifiedworkspacesecretsmanagerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |