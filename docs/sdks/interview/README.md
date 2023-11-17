# Interview
(*Interview*)

### Available Operations

* [CreateAtsInterview](#createatsinterview) - Create a interview
* [GetAtsInterview](#getatsinterview) - Retrieve a interview
* [ListAtsInterviews](#listatsinterviews) - List all interviews
* [PatchAtsInterview](#patchatsinterview) - Update a interview
* [RemoveAtsInterview](#removeatsinterview) - Remove a interview
* [UpdateAtsInterview](#updateatsinterview) - Update a interview

## CreateAtsInterview

Create a interview

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Interview.CreateAtsInterview(ctx, operations.CreateAtsInterviewRequest{
        AtsInterview: &shared.AtsInterview{
            Raw: &shared.PropertyAtsInterviewRaw{},
            UserIds: []string{
                "string",
            },
        },
        ConnectionID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsInterview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateAtsInterviewRequest](../../pkg/models/operations/createatsinterviewrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.CreateAtsInterviewResponse](../../pkg/models/operations/createatsinterviewresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetAtsInterview

Retrieve a interview

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Interview.GetAtsInterview(ctx, operations.GetAtsInterviewRequest{
        ConnectionID: "string",
        Fields: []string{
            "string",
        },
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsInterview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetAtsInterviewRequest](../../pkg/models/operations/getatsinterviewrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetAtsInterviewResponse](../../pkg/models/operations/getatsinterviewresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ListAtsInterviews

List all interviews

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Interview.ListAtsInterviews(ctx, operations.ListAtsInterviewsRequest{
        ConnectionID: "string",
        Fields: []string{
            "string",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsInterviews != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListAtsInterviewsRequest](../../pkg/models/operations/listatsinterviewsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.ListAtsInterviewsResponse](../../pkg/models/operations/listatsinterviewsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PatchAtsInterview

Update a interview

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Interview.PatchAtsInterview(ctx, operations.PatchAtsInterviewRequest{
        AtsInterview: &shared.AtsInterview{
            Raw: &shared.PropertyAtsInterviewRaw{},
            UserIds: []string{
                "string",
            },
        },
        ConnectionID: "string",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsInterview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PatchAtsInterviewRequest](../../pkg/models/operations/patchatsinterviewrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.PatchAtsInterviewResponse](../../pkg/models/operations/patchatsinterviewresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## RemoveAtsInterview

Remove a interview

### Example Usage

```go
package main

import(
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"context"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Interview.RemoveAtsInterview(ctx, operations.RemoveAtsInterviewRequest{
        ConnectionID: "string",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.RemoveAtsInterviewRequest](../../pkg/models/operations/removeatsinterviewrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.RemoveAtsInterviewResponse](../../pkg/models/operations/removeatsinterviewresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## UpdateAtsInterview

Update a interview

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
        unifiedgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Interview.UpdateAtsInterview(ctx, operations.UpdateAtsInterviewRequest{
        AtsInterview: &shared.AtsInterview{
            Raw: &shared.PropertyAtsInterviewRaw{},
            UserIds: []string{
                "string",
            },
        },
        ConnectionID: "string",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsInterview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateAtsInterviewRequest](../../pkg/models/operations/updateatsinterviewrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.UpdateAtsInterviewResponse](../../pkg/models/operations/updateatsinterviewresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
