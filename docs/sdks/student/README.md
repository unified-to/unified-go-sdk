# Student

## Overview

### Available Operations

* [CreateLmsStudent](#createlmsstudent) - Create a student
* [GetLmsStudent](#getlmsstudent) - Retrieve a student
* [ListLmsStudents](#listlmsstudents) - List all students
* [PatchLmsStudent](#patchlmsstudent) - Update a student
* [RemoveLmsStudent](#removelmsstudent) - Remove a student
* [UpdateLmsStudent](#updatelmsstudent) - Update a student

## CreateLmsStudent

Creates a LmsStudent. Provide the unified LmsStudent fields in the request body; they are mapped to the provider's API. Returns the created LmsStudent.

### Example Usage

<!-- UsageSnippet language="go" operationID="createLmsStudent" method="post" path="/lms/{connection_id}/student" example="lms_student" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Student.CreateLmsStudent(ctx, operations.CreateLmsStudentRequest{
        LmsStudent: shared.LmsStudent{
            Address: &shared.PropertyLmsStudentAddress{
                Address1: unifiedgosdk.Pointer("15033 Corwin Circles"),
                City: unifiedgosdk.Pointer("New Craigside"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("73966"),
                Region: unifiedgosdk.Pointer("New Jersey"),
                RegionCode: unifiedgosdk.Pointer("TX"),
            },
            CreatedAt: types.MustNewTimeFromString("2020-03-23T06:59:29.777Z"),
            Emails: []shared.LmsEmail{
                shared.LmsEmail{},
            },
            FirstName: unifiedgosdk.Pointer("Marcella"),
            ID: unifiedgosdk.Pointer("036ae77f-7586-4e77-9008-cc5a7ff832d3"),
            ImageURL: unifiedgosdk.Pointer("https://cdn.jsdelivr.net/gh/faker-js/assets-person-portrait/male/512/56.jpg"),
            LastName: unifiedgosdk.Pointer("Murazik"),
            Metadata: []shared.LmsMetadata{},
            Name: unifiedgosdk.Pointer("Marcella Murazik"),
            Telephones: []shared.LmsTelephone{
                shared.LmsTelephone{
                    Telephone: "(899) 827-6400",
                    Type: shared.LmsTelephoneTypeWork.ToPointer(),
                },
                shared.LmsTelephone{
                    Telephone: "(487) 213-8236",
                    Type: shared.LmsTelephoneTypeFax.ToPointer(),
                },
                shared.LmsTelephone{
                    Telephone: "(980) 775-2940",
                    Type: shared.LmsTelephoneTypeFax.ToPointer(),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2022-07-02T15:24:37.219Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsStudent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateLmsStudentRequest](../../pkg/models/operations/createlmsstudentrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateLmsStudentResponse](../../pkg/models/operations/createlmsstudentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetLmsStudent

Retrieves a single LmsStudent by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getLmsStudent" method="get" path="/lms/{connection_id}/student/{id}" -->
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

    res, err := s.Student.GetLmsStudent(ctx, operations.GetLmsStudentRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsStudent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetLmsStudentRequest](../../pkg/models/operations/getlmsstudentrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetLmsStudentResponse](../../pkg/models/operations/getlmsstudentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListLmsStudents

Returns a list of LmsStudents. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listLmsStudents" method="get" path="/lms/{connection_id}/student" -->
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

    res, err := s.Student.ListLmsStudents(ctx, operations.ListLmsStudentsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsStudents != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListLmsStudentsRequest](../../pkg/models/operations/listlmsstudentsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListLmsStudentsResponse](../../pkg/models/operations/listlmsstudentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchLmsStudent

Updates the LmsStudent identified by `id`. Provide the unified LmsStudent fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchLmsStudent" method="patch" path="/lms/{connection_id}/student/{id}" example="lms_student" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Student.PatchLmsStudent(ctx, operations.PatchLmsStudentRequest{
        LmsStudent: shared.LmsStudent{
            Address: &shared.PropertyLmsStudentAddress{
                Address1: unifiedgosdk.Pointer("15033 Corwin Circles"),
                City: unifiedgosdk.Pointer("New Craigside"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("73966"),
                Region: unifiedgosdk.Pointer("New Jersey"),
                RegionCode: unifiedgosdk.Pointer("TX"),
            },
            CreatedAt: types.MustNewTimeFromString("2020-03-23T06:59:29.777Z"),
            Emails: []shared.LmsEmail{
                shared.LmsEmail{},
            },
            FirstName: unifiedgosdk.Pointer("Marcella"),
            ID: unifiedgosdk.Pointer("4bd49f70-72c1-4e00-a4ab-d9f29e87ab25"),
            ImageURL: unifiedgosdk.Pointer("https://cdn.jsdelivr.net/gh/faker-js/assets-person-portrait/male/512/56.jpg"),
            LastName: unifiedgosdk.Pointer("Murazik"),
            Metadata: []shared.LmsMetadata{},
            Name: unifiedgosdk.Pointer("Marcella Murazik"),
            Telephones: []shared.LmsTelephone{
                shared.LmsTelephone{
                    Telephone: "(899) 827-6400",
                    Type: shared.LmsTelephoneTypeWork.ToPointer(),
                },
                shared.LmsTelephone{
                    Telephone: "(487) 213-8236",
                    Type: shared.LmsTelephoneTypeFax.ToPointer(),
                },
                shared.LmsTelephone{
                    Telephone: "(980) 775-2940",
                    Type: shared.LmsTelephoneTypeFax.ToPointer(),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2022-07-02T15:24:37.224Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsStudent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PatchLmsStudentRequest](../../pkg/models/operations/patchlmsstudentrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.PatchLmsStudentResponse](../../pkg/models/operations/patchlmsstudentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveLmsStudent

Deletes the LmsStudent identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeLmsStudent" method="delete" path="/lms/{connection_id}/student/{id}" -->
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

    res, err := s.Student.RemoveLmsStudent(ctx, operations.RemoveLmsStudentRequest{
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.RemoveLmsStudentRequest](../../pkg/models/operations/removelmsstudentrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.RemoveLmsStudentResponse](../../pkg/models/operations/removelmsstudentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateLmsStudent

Updates the LmsStudent identified by `id`. Provide the unified LmsStudent fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateLmsStudent" method="put" path="/lms/{connection_id}/student/{id}" example="lms_student" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Student.UpdateLmsStudent(ctx, operations.UpdateLmsStudentRequest{
        LmsStudent: shared.LmsStudent{
            Address: &shared.PropertyLmsStudentAddress{
                Address1: unifiedgosdk.Pointer("15033 Corwin Circles"),
                City: unifiedgosdk.Pointer("New Craigside"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("73966"),
                Region: unifiedgosdk.Pointer("New Jersey"),
                RegionCode: unifiedgosdk.Pointer("TX"),
            },
            CreatedAt: types.MustNewTimeFromString("2020-03-23T06:59:29.777Z"),
            Emails: []shared.LmsEmail{
                shared.LmsEmail{},
            },
            FirstName: unifiedgosdk.Pointer("Marcella"),
            ID: unifiedgosdk.Pointer("4bd49f70-72c1-4e00-a4ab-d9f29e87ab25"),
            ImageURL: unifiedgosdk.Pointer("https://cdn.jsdelivr.net/gh/faker-js/assets-person-portrait/male/512/56.jpg"),
            LastName: unifiedgosdk.Pointer("Murazik"),
            Metadata: []shared.LmsMetadata{},
            Name: unifiedgosdk.Pointer("Marcella Murazik"),
            Telephones: []shared.LmsTelephone{
                shared.LmsTelephone{
                    Telephone: "(899) 827-6400",
                    Type: shared.LmsTelephoneTypeWork.ToPointer(),
                },
                shared.LmsTelephone{
                    Telephone: "(487) 213-8236",
                    Type: shared.LmsTelephoneTypeFax.ToPointer(),
                },
                shared.LmsTelephone{
                    Telephone: "(980) 775-2940",
                    Type: shared.LmsTelephoneTypeFax.ToPointer(),
                },
            },
            UpdatedAt: types.MustNewTimeFromString("2022-07-02T15:24:37.224Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsStudent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateLmsStudentRequest](../../pkg/models/operations/updatelmsstudentrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpdateLmsStudentResponse](../../pkg/models/operations/updatelmsstudentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |