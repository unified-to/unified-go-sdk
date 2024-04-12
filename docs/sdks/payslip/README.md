# Payslip
(*Payslip*)

### Available Operations

* [GetHrisPayslip](#gethrispayslip) - Retrieve a payslip
* [ListHrisPayslips](#listhrispayslips) - List all payslips

## GetHrisPayslip

Retrieve a payslip

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
    res, err := s.Payslip.GetHrisPayslip(ctx, operations.GetHrisPayslipRequest{
        ConnectionID: "<value>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisPayslip != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetHrisPayslipRequest](../../pkg/models/operations/gethrispaysliprequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetHrisPayslipResponse](../../pkg/models/operations/gethrispayslipresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListHrisPayslips

List all payslips

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
    res, err := s.Payslip.ListHrisPayslips(ctx, operations.ListHrisPayslipsRequest{
        ConnectionID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisPayslips != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListHrisPayslipsRequest](../../pkg/models/operations/listhrispayslipsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.ListHrisPayslipsResponse](../../pkg/models/operations/listhrispayslipsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
