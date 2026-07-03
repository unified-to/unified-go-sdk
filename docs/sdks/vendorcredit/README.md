# Vendorcredit

## Overview

### Available Operations

* [CreateAccountingVendorcredit](#createaccountingvendorcredit) - Create a vendorcredit
* [GetAccountingVendorcredit](#getaccountingvendorcredit) - Retrieve a vendorcredit
* [ListAccountingVendorcredits](#listaccountingvendorcredits) - List all vendorcredits
* [PatchAccountingVendorcredit](#patchaccountingvendorcredit) - Update a vendorcredit
* [RemoveAccountingVendorcredit](#removeaccountingvendorcredit) - Remove a vendorcredit
* [UpdateAccountingVendorcredit](#updateaccountingvendorcredit) - Update a vendorcredit

## CreateAccountingVendorcredit

Create a vendorcredit

### Example Usage

<!-- UsageSnippet language="go" operationID="createAccountingVendorcredit" method="post" path="/accounting/{connection_id}/vendorcredit" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Vendorcredit.CreateAccountingVendorcredit(ctx, operations.CreateAccountingVendorcreditRequest{
        AccountingVendorcredit: shared.AccountingVendorcredit{},
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingVendorcredit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.CreateAccountingVendorcreditRequest](../../pkg/models/operations/createaccountingvendorcreditrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.CreateAccountingVendorcreditResponse](../../pkg/models/operations/createaccountingvendorcreditresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAccountingVendorcredit

Retrieve a vendorcredit

### Example Usage

<!-- UsageSnippet language="go" operationID="getAccountingVendorcredit" method="get" path="/accounting/{connection_id}/vendorcredit/{id}" -->
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

    res, err := s.Vendorcredit.GetAccountingVendorcredit(ctx, operations.GetAccountingVendorcreditRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingVendorcredit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetAccountingVendorcreditRequest](../../pkg/models/operations/getaccountingvendorcreditrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.GetAccountingVendorcreditResponse](../../pkg/models/operations/getaccountingvendorcreditresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAccountingVendorcredits

List all vendorcredits

### Example Usage

<!-- UsageSnippet language="go" operationID="listAccountingVendorcredits" method="get" path="/accounting/{connection_id}/vendorcredit" -->
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

    res, err := s.Vendorcredit.ListAccountingVendorcredits(ctx, operations.ListAccountingVendorcreditsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingVendorcredits != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.ListAccountingVendorcreditsRequest](../../pkg/models/operations/listaccountingvendorcreditsrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.ListAccountingVendorcreditsResponse](../../pkg/models/operations/listaccountingvendorcreditsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAccountingVendorcredit

Update a vendorcredit

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAccountingVendorcredit" method="patch" path="/accounting/{connection_id}/vendorcredit/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Vendorcredit.PatchAccountingVendorcredit(ctx, operations.PatchAccountingVendorcreditRequest{
        AccountingVendorcredit: shared.AccountingVendorcredit{},
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingVendorcredit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.PatchAccountingVendorcreditRequest](../../pkg/models/operations/patchaccountingvendorcreditrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.PatchAccountingVendorcreditResponse](../../pkg/models/operations/patchaccountingvendorcreditresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAccountingVendorcredit

Remove a vendorcredit

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAccountingVendorcredit" method="delete" path="/accounting/{connection_id}/vendorcredit/{id}" -->
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

    res, err := s.Vendorcredit.RemoveAccountingVendorcredit(ctx, operations.RemoveAccountingVendorcreditRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.RemoveAccountingVendorcreditRequest](../../pkg/models/operations/removeaccountingvendorcreditrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.RemoveAccountingVendorcreditResponse](../../pkg/models/operations/removeaccountingvendorcreditresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAccountingVendorcredit

Update a vendorcredit

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAccountingVendorcredit" method="put" path="/accounting/{connection_id}/vendorcredit/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Vendorcredit.UpdateAccountingVendorcredit(ctx, operations.UpdateAccountingVendorcreditRequest{
        AccountingVendorcredit: shared.AccountingVendorcredit{},
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingVendorcredit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.UpdateAccountingVendorcreditRequest](../../pkg/models/operations/updateaccountingvendorcreditrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.UpdateAccountingVendorcreditResponse](../../pkg/models/operations/updateaccountingvendorcreditresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |