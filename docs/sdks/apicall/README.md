# Apicall
(*Apicall*)

### Available Operations

* [GetUnifiedApicall](#getunifiedapicall) - Returns API Calls
* [GetUnifiedApicallID](#getunifiedapicallid) - Retrieve specific API Call by its ID

## GetUnifiedApicall

Returns API Calls

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
    res, err := s.Apicall.GetUnifiedApicall(ctx, operations.GetUnifiedApicallRequest{
        ConnectionID: unifiedgosdk.String("delectus green Hybrid"),
        CreatedLte: types.MustTimeFromString("2021-04-02T21:36:49.952Z"),
        Env: unifiedgosdk.String("Fantastic Iodine indexing"),
        Error: unifiedgosdk.Bool(false),
        ExternalXref: unifiedgosdk.String("Music"),
        IntegrationType: unifiedgosdk.String("Soft"),
        Limit: unifiedgosdk.Float64(2390.64),
        Offset: unifiedgosdk.Float64(3757.34),
        Order: unifiedgosdk.String("mobile envisioneer"),
        Sort: unifiedgosdk.String("North payment opposite"),
        UpdatedGte: types.MustTimeFromString("2021-08-11T16:18:13.644Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APICalls != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetUnifiedApicallRequest](../../models/operations/getunifiedapicallrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetUnifiedApicallResponse](../../models/operations/getunifiedapicallresponse.md), error**


## GetUnifiedApicallID

Retrieve specific API Call by its ID

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
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Apicall.GetUnifiedApicallID(ctx, operations.GetUnifiedApicallIDRequest{
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APICall != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetUnifiedApicallIDRequest](../../models/operations/getunifiedapicallidrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetUnifiedApicallIDResponse](../../models/operations/getunifiedapicallidresponse.md), error**

