# Space
(*Space*)

### Available Operations

* [CreateKmsSpace](#createkmsspace) - Create a space
* [GetKmsSpace](#getkmsspace) - Retrieve a space
* [ListKmsSpaces](#listkmsspaces) - List all spaces
* [PatchKmsSpace](#patchkmsspace) - Update a space
* [RemoveKmsSpace](#removekmsspace) - Remove a space
* [UpdateKmsSpace](#updatekmsspace) - Update a space

## CreateKmsSpace

Create a space

### Example Usage

```go
package main

import(
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    request := operations.CreateKmsSpaceRequest{
        ConnectionID: "<value>",
    }
    
    ctx := context.Background()
    res, err := s.Space.CreateKmsSpace(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.KmsSpace != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateKmsSpaceRequest](../../pkg/models/operations/createkmsspacerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.CreateKmsSpaceResponse](../../pkg/models/operations/createkmsspaceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetKmsSpace

Retrieve a space

### Example Usage

```go
package main

import(
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    request := operations.GetKmsSpaceRequest{
        ConnectionID: "<value>",
        ID: "<id>",
    }
    
    ctx := context.Background()
    res, err := s.Space.GetKmsSpace(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.KmsSpace != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetKmsSpaceRequest](../../pkg/models/operations/getkmsspacerequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetKmsSpaceResponse](../../pkg/models/operations/getkmsspaceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListKmsSpaces

List all spaces

### Example Usage

```go
package main

import(
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    request := operations.ListKmsSpacesRequest{
        ConnectionID: "<value>",
    }
    
    ctx := context.Background()
    res, err := s.Space.ListKmsSpaces(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.KmsSpaces != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListKmsSpacesRequest](../../pkg/models/operations/listkmsspacesrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.ListKmsSpacesResponse](../../pkg/models/operations/listkmsspacesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PatchKmsSpace

Update a space

### Example Usage

```go
package main

import(
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    request := operations.PatchKmsSpaceRequest{
        ConnectionID: "<value>",
        ID: "<id>",
    }
    
    ctx := context.Background()
    res, err := s.Space.PatchKmsSpace(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.KmsSpace != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.PatchKmsSpaceRequest](../../pkg/models/operations/patchkmsspacerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.PatchKmsSpaceResponse](../../pkg/models/operations/patchkmsspaceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RemoveKmsSpace

Remove a space

### Example Usage

```go
package main

import(
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    request := operations.RemoveKmsSpaceRequest{
        ConnectionID: "<value>",
        ID: "<id>",
    }
    
    ctx := context.Background()
    res, err := s.Space.RemoveKmsSpace(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.RemoveKmsSpaceRequest](../../pkg/models/operations/removekmsspacerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.RemoveKmsSpaceResponse](../../pkg/models/operations/removekmsspaceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateKmsSpace

Update a space

### Example Usage

```go
package main

import(
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    request := operations.UpdateKmsSpaceRequest{
        ConnectionID: "<value>",
        ID: "<id>",
    }
    
    ctx := context.Background()
    res, err := s.Space.UpdateKmsSpace(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.KmsSpace != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateKmsSpaceRequest](../../pkg/models/operations/updatekmsspacerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UpdateKmsSpaceResponse](../../pkg/models/operations/updatekmsspaceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
