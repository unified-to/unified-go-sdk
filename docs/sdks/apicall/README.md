# Apicall

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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetUnifiedApicallSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Apicall.GetUnifiedApicall(ctx, operations.GetUnifiedApicallRequest{
        ConnectionID: unifiedto.String("voluptate"),
        CreatedLte: types.MustTimeFromString("2022-04-12T10:47:34.158Z"),
        Env: unifiedto.String("eaque"),
        Error: unifiedto.Bool(false),
        ExternalXref: unifiedto.String("pariatur"),
        IntegrationType: unifiedto.String("nemo"),
        Limit: unifiedto.Float64(9755.22),
        Offset: unifiedto.Float64(166.27),
        Order: unifiedto.String("fugiat"),
        Sort: unifiedto.String("amet"),
        UpdatedGte: types.MustTimeFromString("2022-03-27T19:22:24.458Z"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.APICalls != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetUnifiedApicallRequest](../../models/operations/getunifiedapicallrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.GetUnifiedApicallSecurity](../../models/operations/getunifiedapicallsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


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
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetUnifiedApicallIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Apicall.GetUnifiedApicallID(ctx, operations.GetUnifiedApicallIDRequest{
        ID: "5fbb2587-0532-402c-b3d5-fe9b90c28909",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.APICall != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetUnifiedApicallIDRequest](../../models/operations/getunifiedapicallidrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.GetUnifiedApicallIDSecurity](../../models/operations/getunifiedapicallidsecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.GetUnifiedApicallIDResponse](../../models/operations/getunifiedapicallidresponse.md), error**

