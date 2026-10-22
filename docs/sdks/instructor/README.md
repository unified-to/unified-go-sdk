# Instructor

## Overview

### Available Operations

* [CreateLmsInstructor](#createlmsinstructor) - Create an instructor
* [GetLmsInstructor](#getlmsinstructor) - Retrieve an instructor
* [ListLmsInstructors](#listlmsinstructors) - List all instructors
* [PatchLmsInstructor](#patchlmsinstructor) - Update an instructor
* [RemoveLmsInstructor](#removelmsinstructor) - Remove an instructor
* [UpdateLmsInstructor](#updatelmsinstructor) - Update an instructor

## CreateLmsInstructor

Creates a LmsInstructor. Provide the unified LmsInstructor fields in the request body; they are mapped to the provider's API. Returns the created LmsInstructor.

### Example Usage

<!-- UsageSnippet language="go" operationID="createLmsInstructor" method="post" path="/lms/{connection_id}/instructor" example="lms_instructor" -->
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

    res, err := s.Instructor.CreateLmsInstructor(ctx, operations.CreateLmsInstructorRequest{
        LmsInstructor: shared.LmsInstructor{
            CreatedAt: types.MustNewTimeFromString("2021-10-12T16:38:54.979Z"),
            Emails: []shared.LmsEmail{
                shared.LmsEmail{},
            },
            FirstName: unifiedgosdk.Pointer("Deangelo"),
            ID: unifiedgosdk.Pointer("57402488-af3f-404b-8ada-780956f69f5b"),
            ImageURL: unifiedgosdk.Pointer("https://avatars.githubusercontent.com/u/95131132"),
            LastName: unifiedgosdk.Pointer("Ritchie"),
            Metadata: []shared.LmsMetadata{},
            Name: unifiedgosdk.Pointer("Deangelo Ritchie"),
            Telephones: []shared.LmsTelephone{
                shared.LmsTelephone{
                    Telephone: "(605) 752-8163",
                    Type: shared.LmsTelephoneTypeWork.ToPointer(),
                },
                shared.LmsTelephone{
                    Telephone: "(430) 415-2451",
                    Type: shared.LmsTelephoneTypeHome.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("Regional Security Agent"),
            UpdatedAt: types.MustNewTimeFromString("2025-07-25T04:13:25.928Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsInstructor != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreateLmsInstructorRequest](../../pkg/models/operations/createlmsinstructorrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreateLmsInstructorResponse](../../pkg/models/operations/createlmsinstructorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetLmsInstructor

Retrieves a single LmsInstructor by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getLmsInstructor" method="get" path="/lms/{connection_id}/instructor/{id}" -->
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

    res, err := s.Instructor.GetLmsInstructor(ctx, operations.GetLmsInstructorRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsInstructor != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetLmsInstructorRequest](../../pkg/models/operations/getlmsinstructorrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetLmsInstructorResponse](../../pkg/models/operations/getlmsinstructorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListLmsInstructors

Returns a list of LmsInstructors. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listLmsInstructors" method="get" path="/lms/{connection_id}/instructor" -->
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

    res, err := s.Instructor.ListLmsInstructors(ctx, operations.ListLmsInstructorsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsInstructors != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListLmsInstructorsRequest](../../pkg/models/operations/listlmsinstructorsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListLmsInstructorsResponse](../../pkg/models/operations/listlmsinstructorsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchLmsInstructor

Updates the LmsInstructor identified by `id`. Provide the unified LmsInstructor fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchLmsInstructor" method="patch" path="/lms/{connection_id}/instructor/{id}" example="lms_instructor" -->
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

    res, err := s.Instructor.PatchLmsInstructor(ctx, operations.PatchLmsInstructorRequest{
        LmsInstructor: shared.LmsInstructor{
            CreatedAt: types.MustNewTimeFromString("2021-10-12T16:38:54.979Z"),
            Emails: []shared.LmsEmail{
                shared.LmsEmail{},
            },
            FirstName: unifiedgosdk.Pointer("Deangelo"),
            ID: unifiedgosdk.Pointer("7ef6015d-4d96-4456-8b89-51b9e4139eb2"),
            ImageURL: unifiedgosdk.Pointer("https://avatars.githubusercontent.com/u/95131132"),
            LastName: unifiedgosdk.Pointer("Ritchie"),
            Metadata: []shared.LmsMetadata{},
            Name: unifiedgosdk.Pointer("Deangelo Ritchie"),
            Telephones: []shared.LmsTelephone{
                shared.LmsTelephone{
                    Telephone: "(605) 752-8163",
                    Type: shared.LmsTelephoneTypeWork.ToPointer(),
                },
                shared.LmsTelephone{
                    Telephone: "(430) 415-2451",
                    Type: shared.LmsTelephoneTypeHome.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("Regional Security Agent"),
            UpdatedAt: types.MustNewTimeFromString("2025-07-25T04:13:25.942Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsInstructor != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PatchLmsInstructorRequest](../../pkg/models/operations/patchlmsinstructorrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.PatchLmsInstructorResponse](../../pkg/models/operations/patchlmsinstructorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveLmsInstructor

Deletes the LmsInstructor identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeLmsInstructor" method="delete" path="/lms/{connection_id}/instructor/{id}" -->
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

    res, err := s.Instructor.RemoveLmsInstructor(ctx, operations.RemoveLmsInstructorRequest{
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.RemoveLmsInstructorRequest](../../pkg/models/operations/removelmsinstructorrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.RemoveLmsInstructorResponse](../../pkg/models/operations/removelmsinstructorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateLmsInstructor

Updates the LmsInstructor identified by `id`. Provide the unified LmsInstructor fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateLmsInstructor" method="put" path="/lms/{connection_id}/instructor/{id}" example="lms_instructor" -->
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

    res, err := s.Instructor.UpdateLmsInstructor(ctx, operations.UpdateLmsInstructorRequest{
        LmsInstructor: shared.LmsInstructor{
            CreatedAt: types.MustNewTimeFromString("2021-10-12T16:38:54.979Z"),
            Emails: []shared.LmsEmail{
                shared.LmsEmail{},
            },
            FirstName: unifiedgosdk.Pointer("Deangelo"),
            ID: unifiedgosdk.Pointer("7ef6015d-4d96-4456-8b89-51b9e4139eb2"),
            ImageURL: unifiedgosdk.Pointer("https://avatars.githubusercontent.com/u/95131132"),
            LastName: unifiedgosdk.Pointer("Ritchie"),
            Metadata: []shared.LmsMetadata{},
            Name: unifiedgosdk.Pointer("Deangelo Ritchie"),
            Telephones: []shared.LmsTelephone{
                shared.LmsTelephone{
                    Telephone: "(605) 752-8163",
                    Type: shared.LmsTelephoneTypeWork.ToPointer(),
                },
                shared.LmsTelephone{
                    Telephone: "(430) 415-2451",
                    Type: shared.LmsTelephoneTypeHome.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("Regional Security Agent"),
            UpdatedAt: types.MustNewTimeFromString("2025-07-25T04:13:25.942Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsInstructor != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateLmsInstructorRequest](../../pkg/models/operations/updatelmsinstructorrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateLmsInstructorResponse](../../pkg/models/operations/updatelmsinstructorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |