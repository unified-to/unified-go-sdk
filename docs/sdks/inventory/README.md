# Inventory
(*Inventory*)

### Available Operations

* [CreateCommerceInventory](#createcommerceinventory) - Create an inventory
* [GetCommerceInventory](#getcommerceinventory) - Retrieve an inventory
* [ListCommerceInventories](#listcommerceinventories) - List all inventories
* [PatchCommerceInventory](#patchcommerceinventory) - Update an inventory
* [RemoveCommerceInventory](#removecommerceinventory) - Remove an inventory
* [UpdateCommerceInventory](#updatecommerceinventory) - Update an inventory

## CreateCommerceInventory

Create an inventory

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


    operationSecurity := operations.CreateCommerceInventorySecurity{
            Jwt: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Inventory.CreateCommerceInventory(ctx, operations.CreateCommerceInventoryRequest{
        ConnectionID: "<value>",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CommerceInventory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.CreateCommerceInventoryRequest](../../pkg/models/operations/createcommerceinventoryrequest.md)   | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `security`                                                                                                   | [operations.CreateCommerceInventorySecurity](../../pkg/models/operations/createcommerceinventorysecurity.md) | :heavy_check_mark:                                                                                           | The security requirements to use for the request.                                                            |


### Response

**[*operations.CreateCommerceInventoryResponse](../../pkg/models/operations/createcommerceinventoryresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetCommerceInventory

Retrieve an inventory

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


    operationSecurity := operations.GetCommerceInventorySecurity{
            Jwt: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Inventory.GetCommerceInventory(ctx, operations.GetCommerceInventoryRequest{
        ConnectionID: "<value>",
        ID: "<id>",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CommerceInventory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetCommerceInventoryRequest](../../pkg/models/operations/getcommerceinventoryrequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.GetCommerceInventorySecurity](../../pkg/models/operations/getcommerceinventorysecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


### Response

**[*operations.GetCommerceInventoryResponse](../../pkg/models/operations/getcommerceinventoryresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListCommerceInventories

List all inventories

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


    operationSecurity := operations.ListCommerceInventoriesSecurity{
            Jwt: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Inventory.ListCommerceInventories(ctx, operations.ListCommerceInventoriesRequest{
        ConnectionID: "<value>",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CommerceInventories != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ListCommerceInventoriesRequest](../../pkg/models/operations/listcommerceinventoriesrequest.md)   | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `security`                                                                                                   | [operations.ListCommerceInventoriesSecurity](../../pkg/models/operations/listcommerceinventoriessecurity.md) | :heavy_check_mark:                                                                                           | The security requirements to use for the request.                                                            |


### Response

**[*operations.ListCommerceInventoriesResponse](../../pkg/models/operations/listcommerceinventoriesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PatchCommerceInventory

Update an inventory

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


    operationSecurity := operations.PatchCommerceInventorySecurity{
            Jwt: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Inventory.PatchCommerceInventory(ctx, operations.PatchCommerceInventoryRequest{
        ConnectionID: "<value>",
        ID: "<id>",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CommerceInventory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.PatchCommerceInventoryRequest](../../pkg/models/operations/patchcommerceinventoryrequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.PatchCommerceInventorySecurity](../../pkg/models/operations/patchcommerceinventorysecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |


### Response

**[*operations.PatchCommerceInventoryResponse](../../pkg/models/operations/patchcommerceinventoryresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RemoveCommerceInventory

Remove an inventory

### Example Usage

```go
package main

import(
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"context"
	"log"
	"net/http"
)

func main() {
    s := unifiedgosdk.New()


    operationSecurity := operations.RemoveCommerceInventorySecurity{
            Jwt: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Inventory.RemoveCommerceInventory(ctx, operations.RemoveCommerceInventoryRequest{
        ConnectionID: "<value>",
        ID: "<id>",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.RemoveCommerceInventoryRequest](../../pkg/models/operations/removecommerceinventoryrequest.md)   | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `security`                                                                                                   | [operations.RemoveCommerceInventorySecurity](../../pkg/models/operations/removecommerceinventorysecurity.md) | :heavy_check_mark:                                                                                           | The security requirements to use for the request.                                                            |


### Response

**[*operations.RemoveCommerceInventoryResponse](../../pkg/models/operations/removecommerceinventoryresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateCommerceInventory

Update an inventory

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


    operationSecurity := operations.UpdateCommerceInventorySecurity{
            Jwt: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Inventory.UpdateCommerceInventory(ctx, operations.UpdateCommerceInventoryRequest{
        ConnectionID: "<value>",
        ID: "<id>",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CommerceInventory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.UpdateCommerceInventoryRequest](../../pkg/models/operations/updatecommerceinventoryrequest.md)   | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `security`                                                                                                   | [operations.UpdateCommerceInventorySecurity](../../pkg/models/operations/updatecommerceinventorysecurity.md) | :heavy_check_mark:                                                                                           | The security requirements to use for the request.                                                            |


### Response

**[*operations.UpdateCommerceInventoryResponse](../../pkg/models/operations/updatecommerceinventoryresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
