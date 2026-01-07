# Environment

## Overview

### Available Operations

* [CreateUnifiedEnvironment](#createunifiedenvironment) - Create new environments
* [ListUnifiedEnvironments](#listunifiedenvironments) - Returns all environments
* [RemoveUnifiedEnvironment](#removeunifiedenvironment) - Remove an environment

## CreateUnifiedEnvironment

Create new environments

### Example Usage

<!-- UsageSnippet language="go" operationID="createUnifiedEnvironment" method="post" path="/unified/environment" -->
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

    res, err := s.Environment.CreateUnifiedEnvironment(ctx, []string{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Environments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `request`                                                    | [[]string](../../.md)                                        | :heavy_check_mark:                                           | The request object to use for the request.                   |
| `opts`                                                       | [][operations.Option](../../pkg/models/operations/option.md) | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.CreateUnifiedEnvironmentResponse](../../pkg/models/operations/createunifiedenvironmentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListUnifiedEnvironments

Returns all environments

### Example Usage

<!-- UsageSnippet language="go" operationID="listUnifiedEnvironments" method="get" path="/unified/environment" -->
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

    res, err := s.Environment.ListUnifiedEnvironments(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Environments != nil {
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

**[*operations.ListUnifiedEnvironmentsResponse](../../pkg/models/operations/listunifiedenvironmentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveUnifiedEnvironment

Remove an environment

### Example Usage

<!-- UsageSnippet language="go" operationID="removeUnifiedEnvironment" method="delete" path="/unified/environment/{env}" -->
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

    res, err := s.Environment.RemoveUnifiedEnvironment(ctx, operations.RemoveUnifiedEnvironmentRequest{
        Env: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Environments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.RemoveUnifiedEnvironmentRequest](../../pkg/models/operations/removeunifiedenvironmentrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.RemoveUnifiedEnvironmentResponse](../../pkg/models/operations/removeunifiedenvironmentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |