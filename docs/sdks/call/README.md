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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetUcConnectionIDCallSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Call.GetUcConnectionIDCall(ctx, operations.GetUcConnectionIDCallRequest{
        AgentID: unifiedto.String("reiciendis"),
        ConnectionID: "repellat",
        ContactID: unifiedto.String("nulla"),
        Limit: unifiedto.Float64(6711.16),
        Offset: unifiedto.Float64(6176.57),
        Order: unifiedto.String("accusamus"),
        Query: unifiedto.String("doloremque"),
        Sort: unifiedto.String("nisi"),
        UpdatedGte: types.MustTimeFromString("2021-02-24T20:48:23.002Z"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.UcCalls != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetUcConnectionIDCallRequest](../../models/operations/getucconnectionidcallrequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.GetUcConnectionIDCallSecurity](../../models/operations/getucconnectionidcallsecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |


### Response

**[*operations.GetUcConnectionIDCallResponse](../../models/operations/getucconnectionidcallresponse.md), error**

