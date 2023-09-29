# Call
(*Call*)

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
        AgentID: unifiedgosdk.String("Directives"),
        ConnectionID: "female than",
        ContactID: unifiedgosdk.String("reintermediate Enid Applications"),
        Limit: unifiedgosdk.Float64(1980.39),
        Offset: unifiedgosdk.Float64(3478),
        Order: unifiedgosdk.String("white Oklahoma Functionality"),
        Query: unifiedgosdk.String("pricing whether Hillsboro"),
        Sort: unifiedgosdk.String("Wooden desensitize SCSI"),
        UpdatedGte: types.MustTimeFromString("2021-11-03T12:40:46.997Z"),
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

