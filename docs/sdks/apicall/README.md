# Apicall
(*Apicall*)

### Available Operations

* [GetUnifiedApicall](#getunifiedapicall) - Retrieve specific API Call by its ID
* [ListUnifiedApicalls](#listunifiedapicalls) - Returns API Calls

## GetUnifiedApicall

Retrieve specific API Call by its ID

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
    s := unifiedgosdk.New()


    operationSecurity := operations.GetUnifiedApicallSecurity{
            Jwt: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Apicall.GetUnifiedApicall(ctx, operations.GetUnifiedApicallRequest{
        ID: "<id>",
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
| `request`                                                                                        | [operations.GetUnifiedApicallRequest](../../pkg/models/operations/getunifiedapicallrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.GetUnifiedApicallSecurity](../../pkg/models/operations/getunifiedapicallsecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.GetUnifiedApicallResponse](../../pkg/models/operations/getunifiedapicallresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListUnifiedApicalls

Returns API Calls

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
    s := unifiedgosdk.New()


    operationSecurity := operations.ListUnifiedApicallsSecurity{
            Jwt: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Apicall.ListUnifiedApicalls(ctx, operations.ListUnifiedApicallsRequest{}, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.APICalls != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListUnifiedApicallsRequest](../../pkg/models/operations/listunifiedapicallsrequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.ListUnifiedApicallsSecurity](../../pkg/models/operations/listunifiedapicallssecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |


### Response

**[*operations.ListUnifiedApicallsResponse](../../pkg/models/operations/listunifiedapicallsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
