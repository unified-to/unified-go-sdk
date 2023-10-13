# Passthrough
(*Passthrough*)

### Available Operations

* [CreatePassthrough](#createpassthrough) - Passthrough POST
* [ListPassthroughs](#listpassthroughs) - Passthrough GET
* [PatchPassthrough](#patchpassthrough) - Passthrough PUT
* [RemovePassthrough](#removepassthrough) - Passthrough DELETE
* [UpdatePassthrough](#updatepassthrough) - Passthrough PUT

## CreatePassthrough

Passthrough POST

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
    res, err := s.Passthrough.CreatePassthrough(ctx, operations.CreatePassthroughRequest{
        ConnectionID: "UTF8",
        Path: "/home",
        Undefined: &shared.Undefined{},
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Undefined != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreatePassthroughRequest](../../models/operations/createpassthroughrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.CreatePassthroughResponse](../../models/operations/createpassthroughresponse.md), error**


## ListPassthroughs

Passthrough GET

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
    res, err := s.Passthrough.ListPassthroughs(ctx, operations.ListPassthroughsRequest{
        ConnectionID: "circuit Loan",
        Path: "/usr",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Undefined != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListPassthroughsRequest](../../models/operations/listpassthroughsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.ListPassthroughsResponse](../../models/operations/listpassthroughsresponse.md), error**


## PatchPassthrough

Passthrough PUT

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
    res, err := s.Passthrough.PatchPassthrough(ctx, operations.PatchPassthroughRequest{
        ConnectionID: "VGA",
        Path: "/opt/lib",
        Undefined: &shared.Undefined{},
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Undefined != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.PatchPassthroughRequest](../../models/operations/patchpassthroughrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.PatchPassthroughResponse](../../models/operations/patchpassthroughresponse.md), error**


## RemovePassthrough

Passthrough DELETE

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
    res, err := s.Passthrough.RemovePassthrough(ctx, operations.RemovePassthroughRequest{
        ConnectionID: "Maine",
        Path: "/boot",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Undefined != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.RemovePassthroughRequest](../../models/operations/removepassthroughrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.RemovePassthroughResponse](../../models/operations/removepassthroughresponse.md), error**


## UpdatePassthrough

Passthrough PUT

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
    res, err := s.Passthrough.UpdatePassthrough(ctx, operations.UpdatePassthroughRequest{
        ConnectionID: "Manager",
        Path: "/private/var",
        Undefined: &shared.Undefined{},
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Undefined != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdatePassthroughRequest](../../models/operations/updatepassthroughrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.UpdatePassthroughResponse](../../models/operations/updatepassthroughresponse.md), error**

