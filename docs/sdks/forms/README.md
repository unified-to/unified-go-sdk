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

Creates a FormsForm. Provide the unified FormsForm fields in the request body; they are mapped to the provider's API. Returns the created FormsForm.

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
            ConfirmationMessage: unifiedgosdk.Pointer("Alo crebro vado utique."),
            CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
            Description: unifiedgosdk.Pointer("Qui pauper triumphus adipisci comitatus tamen stabilis clementia culpo tolero."),
            Fields: []shared.FormField{
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    ID: unifiedgosdk.Pointer("25a17bfb-1b75-488a-ae39-2fdc5e5c9d5b"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(true),
                    MaxLength: unifiedgosdk.Pointer[float64](146.0),
                    Name: "cornu creta confugo",
                    Order: unifiedgosdk.Pointer[float64](0.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2025-04-30T15:47:42.501Z"),
                },
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    ID: unifiedgosdk.Pointer("d48d113d-6728-4fee-a5ab-f5d6fc44d784"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "tracto nostrum desidero",
                    Order: unifiedgosdk.Pointer[float64](1.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2024-09-03T07:45:12.198Z"),
                },
                shared.FormField{
                    Choices: []string{
                        "conor",
                        "patior",
                        "crur",
                        "vinculum",
                        "tredecim",
                    },
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    DefaultValue: unifiedgosdk.Pointer("somniculosus"),
                    ID: unifiedgosdk.Pointer("4ca13e85-1378-4cb0-81b3-c29bfb1d4f7b"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "officia supra consequatur",
                    Order: unifiedgosdk.Pointer[float64](2.0),
                    Type: shared.FormFieldTypeMultipleSelect,
                    UpdatedAt: types.MustNewTimeFromString("2024-02-26T20:28:28.326Z"),
                },
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    Description: unifiedgosdk.Pointer("Antepono cursim aurum voluptatum placeat audax tripudio amo."),
                    ID: unifiedgosdk.Pointer("98f7cbb6-830e-4302-b8db-2a3080ba8ce3"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "curtus fugiat audentia",
                    Order: unifiedgosdk.Pointer[float64](3.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2024-11-28T20:59:22.757Z"),
                },
            },
            HasMultipleSubmissions: unifiedgosdk.Pointer(false),
            HasProgressBar: unifiedgosdk.Pointer(false),
            HasShuffleQuestions: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("63fbacf8-c6b1-4b15-b95c-61a0a9feaa14"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.FormsMetadata{},
            Name: unifiedgosdk.Pointer("tandem tracto vivo Form"),
            PublishedURL: unifiedgosdk.Pointer("https://massive-blowgun.info"),
            ResponseCount: unifiedgosdk.Pointer[float64](423.0),
            UpdatedAt: types.MustNewTimeFromString("2024-08-25T22:18:59.237Z"),
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

Retrieves a single FormsForm by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single FormsSubmission by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Returns a list of FormsForms. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of FormsSubmissions. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Updates the FormsForm identified by `id`. Provide the unified FormsForm fields to change in the request body.

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
            ConfirmationMessage: unifiedgosdk.Pointer("Alo crebro vado utique."),
            CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
            Description: unifiedgosdk.Pointer("Qui pauper triumphus adipisci comitatus tamen stabilis clementia culpo tolero."),
            Fields: []shared.FormField{
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    ID: unifiedgosdk.Pointer("bd2c1cf6-b525-42aa-bbaa-fba71140731e"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(true),
                    MaxLength: unifiedgosdk.Pointer[float64](146.0),
                    Name: "cornu creta confugo",
                    Order: unifiedgosdk.Pointer[float64](0.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2025-04-30T15:47:42.510Z"),
                },
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    ID: unifiedgosdk.Pointer("4c515266-8aa1-4a46-b780-1516c197393f"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "tracto nostrum desidero",
                    Order: unifiedgosdk.Pointer[float64](1.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2024-09-03T07:45:12.203Z"),
                },
                shared.FormField{
                    Choices: []string{
                        "conor",
                        "patior",
                        "crur",
                        "vinculum",
                        "tredecim",
                    },
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    DefaultValue: unifiedgosdk.Pointer("somniculosus"),
                    ID: unifiedgosdk.Pointer("83d33c05-b704-4d91-9532-7e9d0a00e705"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "officia supra consequatur",
                    Order: unifiedgosdk.Pointer[float64](2.0),
                    Type: shared.FormFieldTypeMultipleSelect,
                    UpdatedAt: types.MustNewTimeFromString("2024-02-26T20:28:28.329Z"),
                },
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    Description: unifiedgosdk.Pointer("Antepono cursim aurum voluptatum placeat audax tripudio amo."),
                    ID: unifiedgosdk.Pointer("37dacd72-0bce-467f-b44e-c77df300e80a"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "curtus fugiat audentia",
                    Order: unifiedgosdk.Pointer[float64](3.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2024-11-28T20:59:22.764Z"),
                },
            },
            HasMultipleSubmissions: unifiedgosdk.Pointer(false),
            HasProgressBar: unifiedgosdk.Pointer(false),
            HasShuffleQuestions: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("dd12b49e-fdd1-4b33-8d38-7d4d3edca908"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.FormsMetadata{},
            Name: unifiedgosdk.Pointer("tandem tracto vivo Form"),
            PublishedURL: unifiedgosdk.Pointer("https://massive-blowgun.info"),
            ResponseCount: unifiedgosdk.Pointer[float64](423.0),
            UpdatedAt: types.MustNewTimeFromString("2024-08-25T22:18:59.242Z"),
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

Deletes the FormsForm identified by `id`.

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

Updates the FormsForm identified by `id`. Provide the unified FormsForm fields to change in the request body.

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
            ConfirmationMessage: unifiedgosdk.Pointer("Alo crebro vado utique."),
            CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
            Description: unifiedgosdk.Pointer("Qui pauper triumphus adipisci comitatus tamen stabilis clementia culpo tolero."),
            Fields: []shared.FormField{
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    ID: unifiedgosdk.Pointer("bd2c1cf6-b525-42aa-bbaa-fba71140731e"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(true),
                    MaxLength: unifiedgosdk.Pointer[float64](146.0),
                    Name: "cornu creta confugo",
                    Order: unifiedgosdk.Pointer[float64](0.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2025-04-30T15:47:42.510Z"),
                },
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    ID: unifiedgosdk.Pointer("4c515266-8aa1-4a46-b780-1516c197393f"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "tracto nostrum desidero",
                    Order: unifiedgosdk.Pointer[float64](1.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2024-09-03T07:45:12.203Z"),
                },
                shared.FormField{
                    Choices: []string{
                        "conor",
                        "patior",
                        "crur",
                        "vinculum",
                        "tredecim",
                    },
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    DefaultValue: unifiedgosdk.Pointer("somniculosus"),
                    ID: unifiedgosdk.Pointer("83d33c05-b704-4d91-9532-7e9d0a00e705"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "officia supra consequatur",
                    Order: unifiedgosdk.Pointer[float64](2.0),
                    Type: shared.FormFieldTypeMultipleSelect,
                    UpdatedAt: types.MustNewTimeFromString("2024-02-26T20:28:28.329Z"),
                },
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    Description: unifiedgosdk.Pointer("Antepono cursim aurum voluptatum placeat audax tripudio amo."),
                    ID: unifiedgosdk.Pointer("37dacd72-0bce-467f-b44e-c77df300e80a"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "curtus fugiat audentia",
                    Order: unifiedgosdk.Pointer[float64](3.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2024-11-28T20:59:22.764Z"),
                },
            },
            HasMultipleSubmissions: unifiedgosdk.Pointer(false),
            HasProgressBar: unifiedgosdk.Pointer(false),
            HasShuffleQuestions: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("dd12b49e-fdd1-4b33-8d38-7d4d3edca908"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.FormsMetadata{},
            Name: unifiedgosdk.Pointer("tandem tracto vivo Form"),
            PublishedURL: unifiedgosdk.Pointer("https://massive-blowgun.info"),
            ResponseCount: unifiedgosdk.Pointer[float64](423.0),
            UpdatedAt: types.MustNewTimeFromString("2024-08-25T22:18:59.242Z"),
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