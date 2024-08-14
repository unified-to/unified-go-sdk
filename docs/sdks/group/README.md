# Group
(*Group*)

### Available Operations

* [CreateHrisGroup](#createhrisgroup) - Create a group
* [GetHrisGroup](#gethrisgroup) - Retrieve a group
* [ListHrisGroups](#listhrisgroups) - List all groups
* [PatchHrisGroup](#patchhrisgroup) - Update a group
* [RemoveHrisGroup](#removehrisgroup) - Remove a group
* [UpdateHrisGroup](#updatehrisgroup) - Update a group

## CreateHrisGroup

Create a group

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
    request := operations.CreateHrisGroupRequest{
        ConnectionID: "<value>",
    }
    ctx := context.Background()
    res, err := s.Group.CreateHrisGroup(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisGroup != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateHrisGroupRequest](../../pkg/models/operations/createhrisgrouprequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |


### Response

**[*operations.CreateHrisGroupResponse](../../pkg/models/operations/createhrisgroupresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetHrisGroup

Retrieve a group

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
    request := operations.GetHrisGroupRequest{
        ConnectionID: "<value>",
        ID: "<id>",
    }
    ctx := context.Background()
    res, err := s.Group.GetHrisGroup(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisGroup != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetHrisGroupRequest](../../pkg/models/operations/gethrisgrouprequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |


### Response

**[*operations.GetHrisGroupResponse](../../pkg/models/operations/gethrisgroupresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListHrisGroups

List all groups

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
    request := operations.ListHrisGroupsRequest{
        ConnectionID: "<value>",
    }
    ctx := context.Background()
    res, err := s.Group.ListHrisGroups(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisGroups != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListHrisGroupsRequest](../../pkg/models/operations/listhrisgroupsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |


### Response

**[*operations.ListHrisGroupsResponse](../../pkg/models/operations/listhrisgroupsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PatchHrisGroup

Update a group

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
    request := operations.PatchHrisGroupRequest{
        ConnectionID: "<value>",
        ID: "<id>",
    }
    ctx := context.Background()
    res, err := s.Group.PatchHrisGroup(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisGroup != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.PatchHrisGroupRequest](../../pkg/models/operations/patchhrisgrouprequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |


### Response

**[*operations.PatchHrisGroupResponse](../../pkg/models/operations/patchhrisgroupresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RemoveHrisGroup

Remove a group

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
    request := operations.RemoveHrisGroupRequest{
        ConnectionID: "<value>",
        ID: "<id>",
    }
    ctx := context.Background()
    res, err := s.Group.RemoveHrisGroup(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.RemoveHrisGroupRequest](../../pkg/models/operations/removehrisgrouprequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |


### Response

**[*operations.RemoveHrisGroupResponse](../../pkg/models/operations/removehrisgroupresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateHrisGroup

Update a group

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
    request := operations.UpdateHrisGroupRequest{
        ConnectionID: "<value>",
        ID: "<id>",
    }
    ctx := context.Background()
    res, err := s.Group.UpdateHrisGroup(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisGroup != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateHrisGroupRequest](../../pkg/models/operations/updatehrisgrouprequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |


### Response

**[*operations.UpdateHrisGroupResponse](../../pkg/models/operations/updatehrisgroupresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
