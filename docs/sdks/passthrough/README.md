# Passthrough

### Available Operations

* [DeletePassthroughConnectionIDPath](#deletepassthroughconnectionidpath) - Passthrough DELETE
* [GetPassthroughConnectionIDPath](#getpassthroughconnectionidpath) - Passthrough GET
* [PatchPassthroughConnectionIDPath](#patchpassthroughconnectionidpath) - Passthrough PUT
* [PostPassthroughConnectionIDPath](#postpassthroughconnectionidpath) - Passthrough POST
* [PutPassthroughConnectionIDPath](#putpassthroughconnectionidpath) - Passthrough PUT

## DeletePassthroughConnectionIDPath

Passthrough DELETE

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
    operationSecurity := operations.DeletePassthroughConnectionIDPathSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Passthrough.DeletePassthroughConnectionIDPath(ctx, operations.DeletePassthroughConnectionIDPathRequest{
        ConnectionID: "fugiat",
        Path: "debitis",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Undefined != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.DeletePassthroughConnectionIDPathRequest](../../models/operations/deletepassthroughconnectionidpathrequest.md)   | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `security`                                                                                                                   | [operations.DeletePassthroughConnectionIDPathSecurity](../../models/operations/deletepassthroughconnectionidpathsecurity.md) | :heavy_check_mark:                                                                                                           | The security requirements to use for the request.                                                                            |


### Response

**[*operations.DeletePassthroughConnectionIDPathResponse](../../models/operations/deletepassthroughconnectionidpathresponse.md), error**


## GetPassthroughConnectionIDPath

Passthrough GET

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
    operationSecurity := operations.GetPassthroughConnectionIDPathSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Passthrough.GetPassthroughConnectionIDPath(ctx, operations.GetPassthroughConnectionIDPathRequest{
        ConnectionID: "minima",
        Path: "ducimus",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Undefined != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.GetPassthroughConnectionIDPathRequest](../../models/operations/getpassthroughconnectionidpathrequest.md)   | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `security`                                                                                                             | [operations.GetPassthroughConnectionIDPathSecurity](../../models/operations/getpassthroughconnectionidpathsecurity.md) | :heavy_check_mark:                                                                                                     | The security requirements to use for the request.                                                                      |


### Response

**[*operations.GetPassthroughConnectionIDPathResponse](../../models/operations/getpassthroughconnectionidpathresponse.md), error**


## PatchPassthroughConnectionIDPath

Passthrough PUT

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PatchPassthroughConnectionIDPathSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Passthrough.PatchPassthroughConnectionIDPath(ctx, operations.PatchPassthroughConnectionIDPathRequest{
        ConnectionID: "est",
        Path: "dicta",
        Undefined: &shared.Undefined{},
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Undefined != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.PatchPassthroughConnectionIDPathRequest](../../models/operations/patchpassthroughconnectionidpathrequest.md)   | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `security`                                                                                                                 | [operations.PatchPassthroughConnectionIDPathSecurity](../../models/operations/patchpassthroughconnectionidpathsecurity.md) | :heavy_check_mark:                                                                                                         | The security requirements to use for the request.                                                                          |


### Response

**[*operations.PatchPassthroughConnectionIDPathResponse](../../models/operations/patchpassthroughconnectionidpathresponse.md), error**


## PostPassthroughConnectionIDPath

Passthrough POST

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PostPassthroughConnectionIDPathSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Passthrough.PostPassthroughConnectionIDPath(ctx, operations.PostPassthroughConnectionIDPathRequest{
        ConnectionID: "architecto",
        Path: "fugiat",
        Undefined: &shared.Undefined{},
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Undefined != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.PostPassthroughConnectionIDPathRequest](../../models/operations/postpassthroughconnectionidpathrequest.md)   | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `security`                                                                                                               | [operations.PostPassthroughConnectionIDPathSecurity](../../models/operations/postpassthroughconnectionidpathsecurity.md) | :heavy_check_mark:                                                                                                       | The security requirements to use for the request.                                                                        |


### Response

**[*operations.PostPassthroughConnectionIDPathResponse](../../models/operations/postpassthroughconnectionidpathresponse.md), error**


## PutPassthroughConnectionIDPath

Passthrough PUT

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PutPassthroughConnectionIDPathSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Passthrough.PutPassthroughConnectionIDPath(ctx, operations.PutPassthroughConnectionIDPathRequest{
        ConnectionID: "eum",
        Path: "vitae",
        Undefined: &shared.Undefined{},
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Undefined != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.PutPassthroughConnectionIDPathRequest](../../models/operations/putpassthroughconnectionidpathrequest.md)   | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `security`                                                                                                             | [operations.PutPassthroughConnectionIDPathSecurity](../../models/operations/putpassthroughconnectionidpathsecurity.md) | :heavy_check_mark:                                                                                                     | The security requirements to use for the request.                                                                      |


### Response

**[*operations.PutPassthroughConnectionIDPathResponse](../../models/operations/putpassthroughconnectionidpathresponse.md), error**

