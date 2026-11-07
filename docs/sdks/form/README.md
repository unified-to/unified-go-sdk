# Form

## Overview

### Available Operations

* [CreateFormsForm](#createformsform) - Create a form
* [GetFormsForm](#getformsform) - Retrieve a form
* [ListFormsForms](#listformsforms) - List all forms
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

    res, err := s.Form.CreateFormsForm(ctx, operations.CreateFormsFormRequest{
        FormsForm: shared.FormsForm{
            ConfirmationMessage: unifiedgosdk.Pointer("Alo crebro vado utique."),
            CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
            Description: unifiedgosdk.Pointer("Qui pauper triumphus adipisci comitatus tamen stabilis clementia culpo tolero."),
            Fields: []shared.FormField{
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    ID: unifiedgosdk.Pointer("88560670-5e98-4336-b19e-61933844ebfa"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(true),
                    MaxLength: unifiedgosdk.Pointer[float64](146.0),
                    Name: "cornu creta confugo",
                    Order: unifiedgosdk.Pointer[float64](0.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2025-05-07T10:06:04.418Z"),
                },
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    ID: unifiedgosdk.Pointer("1c170be5-be68-4603-bfca-9bbecd90faf3"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "tracto nostrum desidero",
                    Order: unifiedgosdk.Pointer[float64](1.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2024-09-07T06:14:15.757Z"),
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
                    ID: unifiedgosdk.Pointer("964c832c-73cd-4f7a-8909-80a1d8060d48"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "officia supra consequatur",
                    Order: unifiedgosdk.Pointer[float64](2.0),
                    Type: shared.FormFieldTypeMultipleSelect,
                    UpdatedAt: types.MustNewTimeFromString("2024-02-28T13:16:03.572Z"),
                },
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    Description: unifiedgosdk.Pointer("Antepono cursim aurum voluptatum placeat audax tripudio amo."),
                    ID: unifiedgosdk.Pointer("b6adf1ea-f972-4645-802c-dcf0fe122516"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "curtus fugiat audentia",
                    Order: unifiedgosdk.Pointer[float64](3.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2024-12-03T20:00:02.087Z"),
                },
            },
            HasMultipleSubmissions: unifiedgosdk.Pointer(false),
            HasProgressBar: unifiedgosdk.Pointer(false),
            HasShuffleQuestions: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("9ba31079-10f6-4b2f-a4fd-a639577aaba6"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.FormsMetadata{},
            Name: unifiedgosdk.Pointer("tandem tracto vivo Form"),
            PublishedURL: unifiedgosdk.Pointer("https://massive-blowgun.info"),
            ResponseCount: unifiedgosdk.Pointer[float64](423.0),
            UpdatedAt: types.MustNewTimeFromString("2024-08-29T18:25:20.442Z"),
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

    res, err := s.Form.GetFormsForm(ctx, operations.GetFormsFormRequest{
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

    res, err := s.Form.ListFormsForms(ctx, operations.ListFormsFormsRequest{
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

    res, err := s.Form.PatchFormsForm(ctx, operations.PatchFormsFormRequest{
        FormsForm: shared.FormsForm{
            ConfirmationMessage: unifiedgosdk.Pointer("Alo crebro vado utique."),
            CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
            Description: unifiedgosdk.Pointer("Qui pauper triumphus adipisci comitatus tamen stabilis clementia culpo tolero."),
            Fields: []shared.FormField{
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    ID: unifiedgosdk.Pointer("f8d02501-4c26-4c2a-af31-c9215bac31a3"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(true),
                    MaxLength: unifiedgosdk.Pointer[float64](146.0),
                    Name: "cornu creta confugo",
                    Order: unifiedgosdk.Pointer[float64](0.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2025-05-07T10:06:04.429Z"),
                },
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    ID: unifiedgosdk.Pointer("6edaf256-f86e-495e-a938-be4bcc7b70cf"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "tracto nostrum desidero",
                    Order: unifiedgosdk.Pointer[float64](1.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2024-09-07T06:14:15.763Z"),
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
                    ID: unifiedgosdk.Pointer("1abac48d-4402-46db-a621-402f98d99c74"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "officia supra consequatur",
                    Order: unifiedgosdk.Pointer[float64](2.0),
                    Type: shared.FormFieldTypeMultipleSelect,
                    UpdatedAt: types.MustNewTimeFromString("2024-02-28T13:16:03.574Z"),
                },
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    Description: unifiedgosdk.Pointer("Antepono cursim aurum voluptatum placeat audax tripudio amo."),
                    ID: unifiedgosdk.Pointer("fe8d0ddc-cb95-4e51-a3d3-6d62f2c760ac"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "curtus fugiat audentia",
                    Order: unifiedgosdk.Pointer[float64](3.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2024-12-03T20:00:02.095Z"),
                },
            },
            HasMultipleSubmissions: unifiedgosdk.Pointer(false),
            HasProgressBar: unifiedgosdk.Pointer(false),
            HasShuffleQuestions: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("3c04ea71-fcee-4450-a7e2-2ba1c1a54a3f"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.FormsMetadata{},
            Name: unifiedgosdk.Pointer("tandem tracto vivo Form"),
            PublishedURL: unifiedgosdk.Pointer("https://massive-blowgun.info"),
            ResponseCount: unifiedgosdk.Pointer[float64](423.0),
            UpdatedAt: types.MustNewTimeFromString("2024-08-29T18:25:20.448Z"),
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

    res, err := s.Form.RemoveFormsForm(ctx, operations.RemoveFormsFormRequest{
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

    res, err := s.Form.UpdateFormsForm(ctx, operations.UpdateFormsFormRequest{
        FormsForm: shared.FormsForm{
            ConfirmationMessage: unifiedgosdk.Pointer("Alo crebro vado utique."),
            CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
            Description: unifiedgosdk.Pointer("Qui pauper triumphus adipisci comitatus tamen stabilis clementia culpo tolero."),
            Fields: []shared.FormField{
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    ID: unifiedgosdk.Pointer("f8d02501-4c26-4c2a-af31-c9215bac31a3"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(true),
                    MaxLength: unifiedgosdk.Pointer[float64](146.0),
                    Name: "cornu creta confugo",
                    Order: unifiedgosdk.Pointer[float64](0.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2025-05-07T10:06:04.429Z"),
                },
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    ID: unifiedgosdk.Pointer("6edaf256-f86e-495e-a938-be4bcc7b70cf"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "tracto nostrum desidero",
                    Order: unifiedgosdk.Pointer[float64](1.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2024-09-07T06:14:15.763Z"),
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
                    ID: unifiedgosdk.Pointer("1abac48d-4402-46db-a621-402f98d99c74"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "officia supra consequatur",
                    Order: unifiedgosdk.Pointer[float64](2.0),
                    Type: shared.FormFieldTypeMultipleSelect,
                    UpdatedAt: types.MustNewTimeFromString("2024-02-28T13:16:03.574Z"),
                },
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    Description: unifiedgosdk.Pointer("Antepono cursim aurum voluptatum placeat audax tripudio amo."),
                    ID: unifiedgosdk.Pointer("fe8d0ddc-cb95-4e51-a3d3-6d62f2c760ac"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "curtus fugiat audentia",
                    Order: unifiedgosdk.Pointer[float64](3.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2024-12-03T20:00:02.095Z"),
                },
            },
            HasMultipleSubmissions: unifiedgosdk.Pointer(false),
            HasProgressBar: unifiedgosdk.Pointer(false),
            HasShuffleQuestions: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("3c04ea71-fcee-4450-a7e2-2ba1c1a54a3f"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.FormsMetadata{},
            Name: unifiedgosdk.Pointer("tandem tracto vivo Form"),
            PublishedURL: unifiedgosdk.Pointer("https://massive-blowgun.info"),
            ResponseCount: unifiedgosdk.Pointer[float64](423.0),
            UpdatedAt: types.MustNewTimeFromString("2024-08-29T18:25:20.448Z"),
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