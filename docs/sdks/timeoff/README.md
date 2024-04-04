# Timeoff
(*Timeoff*)

### Available Operations

* [GetHrisTimeoff](#gethristimeoff) - Retrieve a timeoff
* [ListHrisTimeoffs](#listhristimeoffs) - List all timeoffs

## GetHrisTimeoff

Retrieve a timeoff

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
    res, err := s.Timeoff.GetHrisTimeoff(ctx, operations.GetHrisTimeoffRequest{
        ConnectionID: "<value>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeoff != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetHrisTimeoffRequest](../../pkg/models/operations/gethristimeoffrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetHrisTimeoffResponse](../../pkg/models/operations/gethristimeoffresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListHrisTimeoffs

List all timeoffs

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
    res, err := s.Timeoff.ListHrisTimeoffs(ctx, operations.ListHrisTimeoffsRequest{
        ConnectionID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeoffs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListHrisTimeoffsRequest](../../pkg/models/operations/listhristimeoffsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.ListHrisTimeoffsResponse](../../pkg/models/operations/listhristimeoffsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
