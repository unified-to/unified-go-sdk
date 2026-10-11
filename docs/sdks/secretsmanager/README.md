# Secretsmanager

## Overview

### Available Operations

* [CreateUnifiedSecretsmanager](#createunifiedsecretsmanager) - Create secrets manager
* [GetUnifiedSecretsmanager](#getunifiedsecretsmanager) - Retrieve secrets manager
* [ListUnifiedSecretsmanagerInstructions](#listunifiedsecretsmanagerinstructions) - List secrets manager setup instructions
* [ListUnifiedSecretsmanagerSecrets](#listunifiedsecretsmanagersecrets) - List secrets in a secrets manager
* [ListUnifiedSecretsmanagers](#listunifiedsecretsmanagers) - List secrets managers
* [RemoveUnifiedSecretsmanager](#removeunifiedsecretsmanager) - Remove secrets manager

## CreateUnifiedSecretsmanager

Create secrets manager

### Example Usage

<!-- UsageSnippet language="go" operationID="createUnifiedSecretsmanager" method="post" path="/unified/secretsmanager" -->
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

    res, err := s.Secretsmanager.CreateUnifiedSecretsmanager(ctx, shared.SecretsManager{
        Auth: map[string]string{
            "key": "<value>",
            "key1": "<value>",
        },
        Name: "<value>",
        Type: shared.SecretsManagerTypeAkeyless,
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

**[*operations.CreateUnifiedSecretsmanagerResponse](../../pkg/models/operations/createunifiedsecretsmanagerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetUnifiedSecretsmanager

Retrieve secrets manager

### Example Usage

<!-- UsageSnippet language="go" operationID="getUnifiedSecretsmanager" method="get" path="/unified/secretsmanager/{id}" -->
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

    res, err := s.Secretsmanager.GetUnifiedSecretsmanager(ctx, operations.GetUnifiedSecretsmanagerRequest{
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetUnifiedSecretsmanagerRequest](../../pkg/models/operations/getunifiedsecretsmanagerrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.GetUnifiedSecretsmanagerResponse](../../pkg/models/operations/getunifiedsecretsmanagerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListUnifiedSecretsmanagerInstructions

List secrets manager setup instructions

### Example Usage

<!-- UsageSnippet language="go" operationID="listUnifiedSecretsmanagerInstructions" method="get" path="/unified/secretsmanager/instructions" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Secretsmanager.ListUnifiedSecretsmanagerInstructions(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.SecretsManagerInstructionsList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `opts`                                                       | [][operations.Option](../../pkg/models/operations/option.md) | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.ListUnifiedSecretsmanagerInstructionsResponse](../../pkg/models/operations/listunifiedsecretsmanagerinstructionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListUnifiedSecretsmanagerSecrets

List secrets in a secrets manager

### Example Usage

<!-- UsageSnippet language="go" operationID="listUnifiedSecretsmanagerSecrets" method="get" path="/unified/secretsmanager/secrets" -->
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

    res, err := s.Secretsmanager.ListUnifiedSecretsmanagerSecrets(ctx, operations.ListUnifiedSecretsmanagerSecretsRequest{
        SecretsmanagerID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SecretsManagerSecrets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.ListUnifiedSecretsmanagerSecretsRequest](../../pkg/models/operations/listunifiedsecretsmanagersecretsrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.ListUnifiedSecretsmanagerSecretsResponse](../../pkg/models/operations/listunifiedsecretsmanagersecretsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListUnifiedSecretsmanagers

List secrets managers

### Example Usage

<!-- UsageSnippet language="go" operationID="listUnifiedSecretsmanagers" method="get" path="/unified/secretsmanager" -->
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

    res, err := s.Secretsmanager.ListUnifiedSecretsmanagers(ctx, operations.ListUnifiedSecretsmanagersRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.SecretsManagers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.ListUnifiedSecretsmanagersRequest](../../pkg/models/operations/listunifiedsecretsmanagersrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.ListUnifiedSecretsmanagersResponse](../../pkg/models/operations/listunifiedsecretsmanagersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveUnifiedSecretsmanager

Remove secrets manager

### Example Usage

<!-- UsageSnippet language="go" operationID="removeUnifiedSecretsmanager" method="delete" path="/unified/secretsmanager/{id}" -->
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

    res, err := s.Secretsmanager.RemoveUnifiedSecretsmanager(ctx, operations.RemoveUnifiedSecretsmanagerRequest{
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
| `request`                                                                                                          | [operations.RemoveUnifiedSecretsmanagerRequest](../../pkg/models/operations/removeunifiedsecretsmanagerrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.RemoveUnifiedSecretsmanagerResponse](../../pkg/models/operations/removeunifiedsecretsmanagerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |