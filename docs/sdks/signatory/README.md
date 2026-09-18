# Signatory

## Overview

### Available Operations

* [CreateSigningSignatory](#createsigningsignatory) - Create a signatory
* [GetSigningSignatory](#getsigningsignatory) - Retrieve a signatory
* [ListSigningSignatories](#listsigningsignatories) - List all signatories
* [PatchSigningSignatory](#patchsigningsignatory) - Update a signatory
* [RemoveSigningSignatory](#removesigningsignatory) - Remove a signatory
* [UpdateSigningSignatory](#updatesigningsignatory) - Update a signatory

## CreateSigningSignatory

Create a signatory

### Example Usage

<!-- UsageSnippet language="go" operationID="createSigningSignatory" method="post" path="/signing/{connection_id}/signatory" example="signing_signatory" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Signatory.CreateSigningSignatory(ctx, operations.CreateSigningSignatoryRequest{
        SigningSignatory: shared.SigningSignatory{
            CreatedAt: types.MustNewTimeFromString("2022-04-16T19:25:01.966Z"),
            Email: unifiedgosdk.Pointer("Hardy.Wehner@gmail.com"),
            ID: unifiedgosdk.Pointer("c4ca0b65-9176-4cfe-a09d-e27c905b9848"),
            Order: unifiedgosdk.Pointer[float64](5.0),
            Role: shared.SigningSignatoryRoleSigner.ToPointer(),
            Status: shared.SigningSignatoryStatusSigned.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-08-10T20:02:03.718Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SigningSignatory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.CreateSigningSignatoryRequest](../../pkg/models/operations/createsigningsignatoryrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.CreateSigningSignatoryResponse](../../pkg/models/operations/createsigningsignatoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetSigningSignatory

Retrieve a signatory

### Example Usage

<!-- UsageSnippet language="go" operationID="getSigningSignatory" method="get" path="/signing/{connection_id}/signatory/{id}" -->
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

    res, err := s.Signatory.GetSigningSignatory(ctx, operations.GetSigningSignatoryRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SigningSignatory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetSigningSignatoryRequest](../../pkg/models/operations/getsigningsignatoryrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.GetSigningSignatoryResponse](../../pkg/models/operations/getsigningsignatoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListSigningSignatories

List all signatories

### Example Usage

<!-- UsageSnippet language="go" operationID="listSigningSignatories" method="get" path="/signing/{connection_id}/signatory" -->
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

    res, err := s.Signatory.ListSigningSignatories(ctx, operations.ListSigningSignatoriesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SigningSignatories != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListSigningSignatoriesRequest](../../pkg/models/operations/listsigningsignatoriesrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.ListSigningSignatoriesResponse](../../pkg/models/operations/listsigningsignatoriesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchSigningSignatory

Update a signatory

### Example Usage

<!-- UsageSnippet language="go" operationID="patchSigningSignatory" method="patch" path="/signing/{connection_id}/signatory/{id}" example="signing_signatory" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Signatory.PatchSigningSignatory(ctx, operations.PatchSigningSignatoryRequest{
        SigningSignatory: shared.SigningSignatory{
            CreatedAt: types.MustNewTimeFromString("2022-04-16T19:25:01.966Z"),
            Email: unifiedgosdk.Pointer("Hardy.Wehner@gmail.com"),
            ID: unifiedgosdk.Pointer("a326b9e7-ff5b-42d5-bc98-1531014c2de9"),
            Order: unifiedgosdk.Pointer[float64](5.0),
            Role: shared.SigningSignatoryRoleSigner.ToPointer(),
            Status: shared.SigningSignatoryStatusSigned.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-08-10T20:02:03.724Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SigningSignatory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.PatchSigningSignatoryRequest](../../pkg/models/operations/patchsigningsignatoryrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.PatchSigningSignatoryResponse](../../pkg/models/operations/patchsigningsignatoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveSigningSignatory

Remove a signatory

### Example Usage

<!-- UsageSnippet language="go" operationID="removeSigningSignatory" method="delete" path="/signing/{connection_id}/signatory/{id}" -->
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

    res, err := s.Signatory.RemoveSigningSignatory(ctx, operations.RemoveSigningSignatoryRequest{
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.RemoveSigningSignatoryRequest](../../pkg/models/operations/removesigningsignatoryrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.RemoveSigningSignatoryResponse](../../pkg/models/operations/removesigningsignatoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateSigningSignatory

Update a signatory

### Example Usage

<!-- UsageSnippet language="go" operationID="updateSigningSignatory" method="put" path="/signing/{connection_id}/signatory/{id}" example="signing_signatory" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Signatory.UpdateSigningSignatory(ctx, operations.UpdateSigningSignatoryRequest{
        SigningSignatory: shared.SigningSignatory{
            CreatedAt: types.MustNewTimeFromString("2022-04-16T19:25:01.966Z"),
            Email: unifiedgosdk.Pointer("Hardy.Wehner@gmail.com"),
            ID: unifiedgosdk.Pointer("a326b9e7-ff5b-42d5-bc98-1531014c2de9"),
            Order: unifiedgosdk.Pointer[float64](5.0),
            Role: shared.SigningSignatoryRoleSigner.ToPointer(),
            Status: shared.SigningSignatoryStatusSigned.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-08-10T20:02:03.724Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SigningSignatory != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.UpdateSigningSignatoryRequest](../../pkg/models/operations/updatesigningsignatoryrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.UpdateSigningSignatoryResponse](../../pkg/models/operations/updatesigningsignatoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |