# Verification

## Overview

### Available Operations

* [CreateVerificationRequest](#createverificationrequest) - Create a request
* [GetVerificationPackage](#getverificationpackage) - Retrieve a package
* [GetVerificationRequest](#getverificationrequest) - Retrieve a request
* [ListVerificationPackages](#listverificationpackages) - List all packages
* [ListVerificationRequests](#listverificationrequests) - List all requests
* [PatchVerificationRequest](#patchverificationrequest) - Update a request
* [RemoveVerificationRequest](#removeverificationrequest) - Remove a request
* [UpdateVerificationRequest](#updateverificationrequest) - Update a request

## CreateVerificationRequest

Creates a VerificationRequest. Provide the unified VerificationRequest fields in the request body; they are mapped to the provider's API. Returns the created VerificationRequest.

### Example Usage

<!-- UsageSnippet language="go" operationID="createVerificationRequest" method="post" path="/verification/{connection_id}/request" example="verification_request" -->
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

    res, err := s.Verification.CreateVerificationRequest(ctx, operations.CreateVerificationRequestRequest{
        VerificationRequest: shared.VerificationRequest{
            CreatedAt: types.MustNewTimeFromString("2023-11-20T08:17:29.950Z"),
            ID: unifiedgosdk.Pointer("c14f6a36-7d13-4ae6-9bfe-07855cae3d9b"),
            Metadata: []shared.VerificationMetadata{},
            Parameters: []shared.VerificationParameterInput{},
            ProfileAddresses: []shared.VerificationAddress{},
            ProfileEmails: []string{},
            ProfileTelephones: []string{},
            ResponseDetails: []shared.VerificationResponseDetail{},
            ResponseDownloadUrls: []string{},
            ResponseSource: unifiedgosdk.Pointer("Implemented regional protocol"),
            ResponseStatus: shared.VerificationRequestResponseStatusPending.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-12-27T07:50:35.190Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VerificationRequest != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.CreateVerificationRequestRequest](../../pkg/models/operations/createverificationrequestrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.CreateVerificationRequestResponse](../../pkg/models/operations/createverificationrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetVerificationPackage

Retrieves a single VerificationPackage by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getVerificationPackage" method="get" path="/verification/{connection_id}/package/{id}" -->
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

    res, err := s.Verification.GetVerificationPackage(ctx, operations.GetVerificationPackageRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VerificationPackage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetVerificationPackageRequest](../../pkg/models/operations/getverificationpackagerequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.GetVerificationPackageResponse](../../pkg/models/operations/getverificationpackageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetVerificationRequest

Retrieves a single VerificationRequest by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getVerificationRequest" method="get" path="/verification/{connection_id}/request/{id}" -->
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

    res, err := s.Verification.GetVerificationRequest(ctx, operations.GetVerificationRequestRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VerificationRequest != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetVerificationRequestRequest](../../pkg/models/operations/getverificationrequestrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.GetVerificationRequestResponse](../../pkg/models/operations/getverificationrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListVerificationPackages

Returns a list of VerificationPackages. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listVerificationPackages" method="get" path="/verification/{connection_id}/package" -->
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

    res, err := s.Verification.ListVerificationPackages(ctx, operations.ListVerificationPackagesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VerificationPackages != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ListVerificationPackagesRequest](../../pkg/models/operations/listverificationpackagesrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.ListVerificationPackagesResponse](../../pkg/models/operations/listverificationpackagesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListVerificationRequests

Returns a list of VerificationRequests. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listVerificationRequests" method="get" path="/verification/{connection_id}/request" -->
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

    res, err := s.Verification.ListVerificationRequests(ctx, operations.ListVerificationRequestsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VerificationRequests != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ListVerificationRequestsRequest](../../pkg/models/operations/listverificationrequestsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.ListVerificationRequestsResponse](../../pkg/models/operations/listverificationrequestsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchVerificationRequest

Updates the VerificationRequest identified by `id`. Provide the unified VerificationRequest fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchVerificationRequest" method="patch" path="/verification/{connection_id}/request/{id}" example="verification_request" -->
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

    res, err := s.Verification.PatchVerificationRequest(ctx, operations.PatchVerificationRequestRequest{
        VerificationRequest: shared.VerificationRequest{
            CreatedAt: types.MustNewTimeFromString("2023-11-20T08:17:29.950Z"),
            ID: unifiedgosdk.Pointer("91a8b308-2168-487f-9993-9e68ed2309ca"),
            Metadata: []shared.VerificationMetadata{},
            Parameters: []shared.VerificationParameterInput{},
            ProfileAddresses: []shared.VerificationAddress{},
            ProfileEmails: []string{},
            ProfileTelephones: []string{},
            ResponseDetails: []shared.VerificationResponseDetail{},
            ResponseDownloadUrls: []string{},
            ResponseSource: unifiedgosdk.Pointer("Implemented regional protocol"),
            ResponseStatus: shared.VerificationRequestResponseStatusPending.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-12-27T07:50:35.191Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VerificationRequest != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PatchVerificationRequestRequest](../../pkg/models/operations/patchverificationrequestrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.PatchVerificationRequestResponse](../../pkg/models/operations/patchverificationrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveVerificationRequest

Deletes the VerificationRequest identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeVerificationRequest" method="delete" path="/verification/{connection_id}/request/{id}" -->
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

    res, err := s.Verification.RemoveVerificationRequest(ctx, operations.RemoveVerificationRequestRequest{
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.RemoveVerificationRequestRequest](../../pkg/models/operations/removeverificationrequestrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.RemoveVerificationRequestResponse](../../pkg/models/operations/removeverificationrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateVerificationRequest

Updates the VerificationRequest identified by `id`. Provide the unified VerificationRequest fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateVerificationRequest" method="put" path="/verification/{connection_id}/request/{id}" example="verification_request" -->
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

    res, err := s.Verification.UpdateVerificationRequest(ctx, operations.UpdateVerificationRequestRequest{
        VerificationRequest: shared.VerificationRequest{
            CreatedAt: types.MustNewTimeFromString("2023-11-20T08:17:29.950Z"),
            ID: unifiedgosdk.Pointer("91a8b308-2168-487f-9993-9e68ed2309ca"),
            Metadata: []shared.VerificationMetadata{},
            Parameters: []shared.VerificationParameterInput{},
            ProfileAddresses: []shared.VerificationAddress{},
            ProfileEmails: []string{},
            ProfileTelephones: []string{},
            ResponseDetails: []shared.VerificationResponseDetail{},
            ResponseDownloadUrls: []string{},
            ResponseSource: unifiedgosdk.Pointer("Implemented regional protocol"),
            ResponseStatus: shared.VerificationRequestResponseStatusPending.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-12-27T07:50:35.191Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VerificationRequest != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.UpdateVerificationRequestRequest](../../pkg/models/operations/updateverificationrequestrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.UpdateVerificationRequestResponse](../../pkg/models/operations/updateverificationrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |