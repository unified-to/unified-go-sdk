# Forms

## Overview

### Available Operations

* [CreateFormsForm](#createformsform) - Create a form
* [GetFormsForm](#getformsform) - Retrieve a form
* [GetFormsSubmission](#getformssubmission) - Retrieve a submission
* [ListFormsForms](#listformsforms) - List all forms
* [ListFormsSubmissions](#listformssubmissions) - List all submissions
* [PatchFormsForm](#patchformsform) - Update a form
* [RemoveFormsForm](#removeformsform) - Remove a form
* [UpdateFormsForm](#updateformsform) - Update a form

## CreateFormsForm

Create a form

### Example Usage

<!-- UsageSnippet language="go" operationID="createFormsForm" method="post" path="/forms/{connection_id}/form" example="forms_form" -->
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

    res, err := s.Forms.CreateFormsForm(ctx, operations.CreateFormsFormRequest{
        FormsForm: shared.FormsForm{
            ConfirmationMessage: unifiedgosdk.Pointer("Cultura temeritas aptus celebrer volo pecus culpa annus aurum."),
            CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
            Description: unifiedgosdk.Pointer("Sodalitas cupiditas terebro conduco."),
            Fields: []shared.FormField{
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    ID: unifiedgosdk.Pointer("565f27cf-2cf7-4c30-ad97-4340d859b584"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(true),
                    MaxLength: unifiedgosdk.Pointer[float64](146.0),
                    Name: "vulgivagus audio accendo",
                    Order: unifiedgosdk.Pointer[float64](0.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2025-04-11T17:48:08.242Z"),
                },
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    ID: unifiedgosdk.Pointer("82b263f9-2d16-4cdf-8e99-d05ba46ce817"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "alo crebro vado",
                    Order: unifiedgosdk.Pointer[float64](1.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2024-08-23T07:28:03.048Z"),
                },
                shared.FormField{
                    Choices: []string{
                        "vallum",
                        "vae",
                        "nesciunt",
                        "commodi",
                        "appositus",
                    },
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    DefaultValue: unifiedgosdk.Pointer("cattus"),
                    ID: unifiedgosdk.Pointer("d7c963e5-2d3d-4436-a57d-a6e337d9d170"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "casso tenus nesciunt",
                    Order: unifiedgosdk.Pointer[float64](2.0),
                    Type: shared.FormFieldTypeMultipleSelect,
                    UpdatedAt: types.MustNewTimeFromString("2024-02-22T02:22:13.646Z"),
                },
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    Description: unifiedgosdk.Pointer("Sequi antea delectatio."),
                    ID: unifiedgosdk.Pointer("e45efb8f-439b-40f6-8370-99c8ec66b065"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "comburo utique ipsa",
                    Order: unifiedgosdk.Pointer[float64](3.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2024-11-15T00:05:57.974Z"),
                },
            },
            HasMultipleSubmissions: unifiedgosdk.Pointer(false),
            HasProgressBar: unifiedgosdk.Pointer(false),
            HasShuffleQuestions: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("9845d639-1cac-4ede-abd7-e7e05ab9ede9"),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("voluptatibus omnis audax Form"),
            PublishedURL: unifiedgosdk.Pointer("https://impartial-institute.org/"),
            ResponseCount: unifiedgosdk.Pointer[float64](423.0),
            UpdatedAt: types.MustNewTimeFromString("2024-08-15T04:41:00.185Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FormsForm != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateFormsFormRequest](../../pkg/models/operations/createformsformrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreateFormsFormResponse](../../pkg/models/operations/createformsformresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetFormsForm

Retrieve a form

### Example Usage

<!-- UsageSnippet language="go" operationID="getFormsForm" method="get" path="/forms/{connection_id}/form/{id}" -->
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

    res, err := s.Forms.GetFormsForm(ctx, operations.GetFormsFormRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FormsForm != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetFormsFormRequest](../../pkg/models/operations/getformsformrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.GetFormsFormResponse](../../pkg/models/operations/getformsformresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetFormsSubmission

Retrieve a submission

### Example Usage

<!-- UsageSnippet language="go" operationID="getFormsSubmission" method="get" path="/forms/{connection_id}/submission/{id}" -->
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

    res, err := s.Forms.GetFormsSubmission(ctx, operations.GetFormsSubmissionRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FormsSubmission != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetFormsSubmissionRequest](../../pkg/models/operations/getformssubmissionrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetFormsSubmissionResponse](../../pkg/models/operations/getformssubmissionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListFormsForms

List all forms

### Example Usage

<!-- UsageSnippet language="go" operationID="listFormsForms" method="get" path="/forms/{connection_id}/form" -->
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

    res, err := s.Forms.ListFormsForms(ctx, operations.ListFormsFormsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FormsForms != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListFormsFormsRequest](../../pkg/models/operations/listformsformsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ListFormsFormsResponse](../../pkg/models/operations/listformsformsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListFormsSubmissions

List all submissions

### Example Usage

<!-- UsageSnippet language="go" operationID="listFormsSubmissions" method="get" path="/forms/{connection_id}/submission" -->
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

    res, err := s.Forms.ListFormsSubmissions(ctx, operations.ListFormsSubmissionsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FormsSubmissions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListFormsSubmissionsRequest](../../pkg/models/operations/listformssubmissionsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListFormsSubmissionsResponse](../../pkg/models/operations/listformssubmissionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchFormsForm

Update a form

### Example Usage

<!-- UsageSnippet language="go" operationID="patchFormsForm" method="patch" path="/forms/{connection_id}/form/{id}" example="forms_form" -->
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

    res, err := s.Forms.PatchFormsForm(ctx, operations.PatchFormsFormRequest{
        FormsForm: shared.FormsForm{
            ConfirmationMessage: unifiedgosdk.Pointer("Cultura temeritas aptus celebrer volo pecus culpa annus aurum."),
            CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
            Description: unifiedgosdk.Pointer("Sodalitas cupiditas terebro conduco."),
            Fields: []shared.FormField{
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    ID: unifiedgosdk.Pointer("565f27cf-2cf7-4c30-ad97-4340d859b584"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(true),
                    MaxLength: unifiedgosdk.Pointer[float64](146.0),
                    Name: "vulgivagus audio accendo",
                    Order: unifiedgosdk.Pointer[float64](0.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2025-04-11T17:48:08.248Z"),
                },
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    ID: unifiedgosdk.Pointer("82b263f9-2d16-4cdf-8e99-d05ba46ce817"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "alo crebro vado",
                    Order: unifiedgosdk.Pointer[float64](1.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2024-08-23T07:28:03.051Z"),
                },
                shared.FormField{
                    Choices: []string{
                        "vallum",
                        "vae",
                        "nesciunt",
                        "commodi",
                        "appositus",
                    },
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    DefaultValue: unifiedgosdk.Pointer("cattus"),
                    ID: unifiedgosdk.Pointer("d7c963e5-2d3d-4436-a57d-a6e337d9d170"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "casso tenus nesciunt",
                    Order: unifiedgosdk.Pointer[float64](2.0),
                    Type: shared.FormFieldTypeMultipleSelect,
                    UpdatedAt: types.MustNewTimeFromString("2024-02-22T02:22:13.648Z"),
                },
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    Description: unifiedgosdk.Pointer("Sequi antea delectatio."),
                    ID: unifiedgosdk.Pointer("e45efb8f-439b-40f6-8370-99c8ec66b065"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "comburo utique ipsa",
                    Order: unifiedgosdk.Pointer[float64](3.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2024-11-15T00:05:57.979Z"),
                },
            },
            HasMultipleSubmissions: unifiedgosdk.Pointer(false),
            HasProgressBar: unifiedgosdk.Pointer(false),
            HasShuffleQuestions: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("9757be4d-9b8d-4864-b0db-487339418161"),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("voluptatibus omnis audax Form"),
            PublishedURL: unifiedgosdk.Pointer("https://impartial-institute.org/"),
            ResponseCount: unifiedgosdk.Pointer[float64](423.0),
            UpdatedAt: types.MustNewTimeFromString("2024-08-15T04:41:00.189Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FormsForm != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.PatchFormsFormRequest](../../pkg/models/operations/patchformsformrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.PatchFormsFormResponse](../../pkg/models/operations/patchformsformresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveFormsForm

Remove a form

### Example Usage

<!-- UsageSnippet language="go" operationID="removeFormsForm" method="delete" path="/forms/{connection_id}/form/{id}" -->
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

    res, err := s.Forms.RemoveFormsForm(ctx, operations.RemoveFormsFormRequest{
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.RemoveFormsFormRequest](../../pkg/models/operations/removeformsformrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.RemoveFormsFormResponse](../../pkg/models/operations/removeformsformresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateFormsForm

Update a form

### Example Usage

<!-- UsageSnippet language="go" operationID="updateFormsForm" method="put" path="/forms/{connection_id}/form/{id}" example="forms_form" -->
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

    res, err := s.Forms.UpdateFormsForm(ctx, operations.UpdateFormsFormRequest{
        FormsForm: shared.FormsForm{
            ConfirmationMessage: unifiedgosdk.Pointer("Cultura temeritas aptus celebrer volo pecus culpa annus aurum."),
            CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
            Description: unifiedgosdk.Pointer("Sodalitas cupiditas terebro conduco."),
            Fields: []shared.FormField{
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    ID: unifiedgosdk.Pointer("565f27cf-2cf7-4c30-ad97-4340d859b584"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(true),
                    MaxLength: unifiedgosdk.Pointer[float64](146.0),
                    Name: "vulgivagus audio accendo",
                    Order: unifiedgosdk.Pointer[float64](0.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2025-04-11T17:48:08.248Z"),
                },
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    ID: unifiedgosdk.Pointer("82b263f9-2d16-4cdf-8e99-d05ba46ce817"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "alo crebro vado",
                    Order: unifiedgosdk.Pointer[float64](1.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2024-08-23T07:28:03.051Z"),
                },
                shared.FormField{
                    Choices: []string{
                        "vallum",
                        "vae",
                        "nesciunt",
                        "commodi",
                        "appositus",
                    },
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    DefaultValue: unifiedgosdk.Pointer("cattus"),
                    ID: unifiedgosdk.Pointer("d7c963e5-2d3d-4436-a57d-a6e337d9d170"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "casso tenus nesciunt",
                    Order: unifiedgosdk.Pointer[float64](2.0),
                    Type: shared.FormFieldTypeMultipleSelect,
                    UpdatedAt: types.MustNewTimeFromString("2024-02-22T02:22:13.648Z"),
                },
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    Description: unifiedgosdk.Pointer("Sequi antea delectatio."),
                    ID: unifiedgosdk.Pointer("e45efb8f-439b-40f6-8370-99c8ec66b065"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "comburo utique ipsa",
                    Order: unifiedgosdk.Pointer[float64](3.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2024-11-15T00:05:57.979Z"),
                },
            },
            HasMultipleSubmissions: unifiedgosdk.Pointer(false),
            HasProgressBar: unifiedgosdk.Pointer(false),
            HasShuffleQuestions: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("9757be4d-9b8d-4864-b0db-487339418161"),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("voluptatibus omnis audax Form"),
            PublishedURL: unifiedgosdk.Pointer("https://impartial-institute.org/"),
            ResponseCount: unifiedgosdk.Pointer[float64](423.0),
            UpdatedAt: types.MustNewTimeFromString("2024-08-15T04:41:00.189Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FormsForm != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateFormsFormRequest](../../pkg/models/operations/updateformsformrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.UpdateFormsFormResponse](../../pkg/models/operations/updateformsformresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |