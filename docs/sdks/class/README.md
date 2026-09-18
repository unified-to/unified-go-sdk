# Class

## Overview

### Available Operations

* [CreateLmsClass](#createlmsclass) - Create a class
* [GetLmsClass](#getlmsclass) - Retrieve a class
* [ListLmsClasses](#listlmsclasses) - List all classes
* [PatchLmsClass](#patchlmsclass) - Update a class
* [RemoveLmsClass](#removelmsclass) - Remove a class
* [UpdateLmsClass](#updatelmsclass) - Update a class

## CreateLmsClass

Create a class

### Example Usage

<!-- UsageSnippet language="go" operationID="createLmsClass" method="post" path="/lms/{connection_id}/class" example="lms_class" -->
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

    res, err := s.Class.CreateLmsClass(ctx, operations.CreateLmsClassRequest{
        LmsClass: shared.LmsClass{
            CreatedAt: types.MustNewTimeFromString("2020-02-20T14:48:51.845Z"),
            Description: unifiedgosdk.Pointer("Anser sperno decerno."),
            ID: unifiedgosdk.Pointer("6216bb21-e805-4860-8073-3f3432c321e1"),
            Instructors: []shared.LmsReference{},
            Languages: []string{
                "in",
            },
            Media: []shared.LmsMedia{
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Defetiscor aetas acies benevolentia ulterius. Creta bis beneficium canis. Bonus valeo vulgo creator arca peior ceno earum culpa. Tabesco apostolus talis. Ultra accommodo deinde sono culpo arto cruciamentum triduana."),
                    Description: unifiedgosdk.Pointer("Esse confido."),
                    Languages: []string{
                        "fa",
                        "da",
                    },
                    Name: unifiedgosdk.Pointer("illo"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://loremflickr.com/199/1934?lock=4323325966476891"),
                    Type: shared.LmsMediaTypeVideo.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://loremflickr.com/487/921?lock=5127962071241632"),
                },
            },
            Name: unifiedgosdk.Pointer("virtus"),
            Students: []shared.LmsReference{},
            UpdatedAt: types.MustNewTimeFromString("2025-07-07T22:25:54.088Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsClass != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateLmsClassRequest](../../pkg/models/operations/createlmsclassrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.CreateLmsClassResponse](../../pkg/models/operations/createlmsclassresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetLmsClass

Retrieve a class

### Example Usage

<!-- UsageSnippet language="go" operationID="getLmsClass" method="get" path="/lms/{connection_id}/class/{id}" -->
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

    res, err := s.Class.GetLmsClass(ctx, operations.GetLmsClassRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsClass != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetLmsClassRequest](../../pkg/models/operations/getlmsclassrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.GetLmsClassResponse](../../pkg/models/operations/getlmsclassresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListLmsClasses

List all classes

### Example Usage

<!-- UsageSnippet language="go" operationID="listLmsClasses" method="get" path="/lms/{connection_id}/class" -->
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

    res, err := s.Class.ListLmsClasses(ctx, operations.ListLmsClassesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsClasses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListLmsClassesRequest](../../pkg/models/operations/listlmsclassesrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ListLmsClassesResponse](../../pkg/models/operations/listlmsclassesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchLmsClass

Update a class

### Example Usage

<!-- UsageSnippet language="go" operationID="patchLmsClass" method="patch" path="/lms/{connection_id}/class/{id}" example="lms_class" -->
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

    res, err := s.Class.PatchLmsClass(ctx, operations.PatchLmsClassRequest{
        LmsClass: shared.LmsClass{
            CreatedAt: types.MustNewTimeFromString("2020-02-20T14:48:51.845Z"),
            Description: unifiedgosdk.Pointer("Anser sperno decerno."),
            ID: unifiedgosdk.Pointer("fcbd51ed-fc99-4bce-9a89-92b0c5dc6a74"),
            Instructors: []shared.LmsReference{},
            Languages: []string{
                "in",
            },
            Media: []shared.LmsMedia{
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Defetiscor aetas acies benevolentia ulterius. Creta bis beneficium canis. Bonus valeo vulgo creator arca peior ceno earum culpa. Tabesco apostolus talis. Ultra accommodo deinde sono culpo arto cruciamentum triduana."),
                    Description: unifiedgosdk.Pointer("Esse confido."),
                    Languages: []string{
                        "fa",
                        "da",
                    },
                    Name: unifiedgosdk.Pointer("illo"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://loremflickr.com/199/1934?lock=4323325966476891"),
                    Type: shared.LmsMediaTypeVideo.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://loremflickr.com/487/921?lock=5127962071241632"),
                },
            },
            Name: unifiedgosdk.Pointer("virtus"),
            Students: []shared.LmsReference{},
            UpdatedAt: types.MustNewTimeFromString("2025-07-07T22:25:54.098Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsClass != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.PatchLmsClassRequest](../../pkg/models/operations/patchlmsclassrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.PatchLmsClassResponse](../../pkg/models/operations/patchlmsclassresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveLmsClass

Remove a class

### Example Usage

<!-- UsageSnippet language="go" operationID="removeLmsClass" method="delete" path="/lms/{connection_id}/class/{id}" -->
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

    res, err := s.Class.RemoveLmsClass(ctx, operations.RemoveLmsClassRequest{
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.RemoveLmsClassRequest](../../pkg/models/operations/removelmsclassrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.RemoveLmsClassResponse](../../pkg/models/operations/removelmsclassresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateLmsClass

Update a class

### Example Usage

<!-- UsageSnippet language="go" operationID="updateLmsClass" method="put" path="/lms/{connection_id}/class/{id}" example="lms_class" -->
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

    res, err := s.Class.UpdateLmsClass(ctx, operations.UpdateLmsClassRequest{
        LmsClass: shared.LmsClass{
            CreatedAt: types.MustNewTimeFromString("2020-02-20T14:48:51.845Z"),
            Description: unifiedgosdk.Pointer("Anser sperno decerno."),
            ID: unifiedgosdk.Pointer("fcbd51ed-fc99-4bce-9a89-92b0c5dc6a74"),
            Instructors: []shared.LmsReference{},
            Languages: []string{
                "in",
            },
            Media: []shared.LmsMedia{
                shared.LmsMedia{
                    Content: unifiedgosdk.Pointer("Defetiscor aetas acies benevolentia ulterius. Creta bis beneficium canis. Bonus valeo vulgo creator arca peior ceno earum culpa. Tabesco apostolus talis. Ultra accommodo deinde sono culpo arto cruciamentum triduana."),
                    Description: unifiedgosdk.Pointer("Esse confido."),
                    Languages: []string{
                        "fa",
                        "da",
                    },
                    Name: unifiedgosdk.Pointer("illo"),
                    ThumbnailURL: unifiedgosdk.Pointer("https://loremflickr.com/199/1934?lock=4323325966476891"),
                    Type: shared.LmsMediaTypeVideo.ToPointer(),
                    URL: unifiedgosdk.Pointer("https://loremflickr.com/487/921?lock=5127962071241632"),
                },
            },
            Name: unifiedgosdk.Pointer("virtus"),
            Students: []shared.LmsReference{},
            UpdatedAt: types.MustNewTimeFromString("2025-07-07T22:25:54.098Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsClass != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateLmsClassRequest](../../pkg/models/operations/updatelmsclassrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.UpdateLmsClassResponse](../../pkg/models/operations/updatelmsclassresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |