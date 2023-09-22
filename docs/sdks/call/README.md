# Call

### Available Operations

* [GetUcConnectionIDCall](#getucconnectionidcall) - List all calls

## GetUcConnectionIDCall

List all calls

### Example Usage

```go
package main

import(
	"context"
	"log"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Call.GetUcConnectionIDCall(ctx, operations.GetUcConnectionIDCallRequest{
        AgentID: unifiedgosdk.String("a"),
        ConnectionID: "iste",
        ContactID: unifiedgosdk.String("dicta"),
        Limit: unifiedgosdk.Float64(5524.39),
        Offset: unifiedgosdk.Float64(3563.15),
        Order: unifiedgosdk.String("dolore"),
        Query: unifiedgosdk.String("modi"),
        Sort: unifiedgosdk.String("itaque"),
        UpdatedGte: types.MustTimeFromString("2022-03-15T19:59:59.350Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UcCalls != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetUcConnectionIDCallRequest](../../models/operations/getucconnectionidcallrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.GetUcConnectionIDCallResponse](../../models/operations/getucconnectionidcallresponse.md), error**

