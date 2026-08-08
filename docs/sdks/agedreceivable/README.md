# Agedreceivable

## Overview

### Available Operations

* [GetAccountingAgedreceivable](#getaccountingagedreceivable) - Retrieve an agedreceivable
* [ListAccountingAgedreceivables](#listaccountingagedreceivables) - List all agedreceivables

## GetAccountingAgedreceivable

Retrieve an agedreceivable

### Example Usage

<!-- UsageSnippet language="go" operationID="getAccountingAgedreceivable" method="get" path="/accounting/{connection_id}/agedreceivable/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Agedreceivable.GetAccountingAgedreceivable(ctx, operations.GetAccountingAgedreceivableRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingAgedreceivable != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetAccountingAgedreceivableRequest](../../pkg/models/operations/getaccountingagedreceivablerequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.GetAccountingAgedreceivableResponse](../../pkg/models/operations/getaccountingagedreceivableresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAccountingAgedreceivables

List all agedreceivables

### Example Usage

<!-- UsageSnippet language="go" operationID="listAccountingAgedreceivables" method="get" path="/accounting/{connection_id}/agedreceivable" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Agedreceivable.ListAccountingAgedreceivables(ctx, operations.ListAccountingAgedreceivablesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingAgedreceivables != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.ListAccountingAgedreceivablesRequest](../../pkg/models/operations/listaccountingagedreceivablesrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.ListAccountingAgedreceivablesResponse](../../pkg/models/operations/listaccountingagedreceivablesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |