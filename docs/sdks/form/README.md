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
                    ID: unifiedgosdk.Pointer("e51b3a5e-b796-4194-8160-78443406249d"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(true),
                    MaxLength: unifiedgosdk.Pointer[float64](146.0),
                    Name: "voluntarius argumentum decumbo",
                    Order: unifiedgosdk.Pointer[float64](0.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2025-05-17T03:22:32.979Z"),
                },
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    ID: unifiedgosdk.Pointer("274ca2a2-3a78-4233-9271-df58454219f4"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "desidero clibanus calcar",
                    Order: unifiedgosdk.Pointer[float64](1.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2024-09-12T22:02:07.454Z"),
                },
                shared.FormField{
                    Choices: []string{
                        "crur",
                        "vinculum",
                        "tredecim",
                        "somniculosus",
                        "curtus",
                    },
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    DefaultValue: unifiedgosdk.Pointer("fugiat"),
                    ID: unifiedgosdk.Pointer("ccb3c460-2eba-44b3-a22f-e12d30cbda92"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "consequatur error cena",
                    Order: unifiedgosdk.Pointer[float64](2.0),
                    Type: shared.FormFieldTypeMultipleSelect,
                    UpdatedAt: types.MustNewTimeFromString("2024-03-01T23:53:51.432Z"),
                },
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    Description: unifiedgosdk.Pointer("Aurum voluptatum placeat audax tripudio amo."),
                    ID: unifiedgosdk.Pointer("b6c42c22-65f5-4a01-94da-dc3a974e2031"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "audentia subseco antepono",
                    Order: unifiedgosdk.Pointer[float64](3.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2024-12-10T23:02:56.977Z"),
                },
            },
            HasMultipleSubmissions: unifiedgosdk.Pointer(false),
            HasProgressBar: unifiedgosdk.Pointer(false),
            HasShuffleQuestions: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("22433036-6bef-426b-89b7-b38b46125a64"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.FormsMetadata{},
            Name: unifiedgosdk.Pointer("tandem tracto vivo Form"),
            PublishedURL: unifiedgosdk.Pointer("https://massive-blowgun.info"),
            ResponseCount: unifiedgosdk.Pointer[float64](423.0),
            UpdatedAt: types.MustNewTimeFromString("2024-09-04T06:48:05.889Z"),
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
                    ID: unifiedgosdk.Pointer("79894811-f797-4f7f-b210-c7ee1443ac8d"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(true),
                    MaxLength: unifiedgosdk.Pointer[float64](146.0),
                    Name: "voluntarius argumentum decumbo",
                    Order: unifiedgosdk.Pointer[float64](0.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2025-05-17T03:22:32.991Z"),
                },
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    ID: unifiedgosdk.Pointer("1808686e-8a59-4266-9c11-2580e8c68383"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "desidero clibanus calcar",
                    Order: unifiedgosdk.Pointer[float64](1.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2024-09-12T22:02:07.460Z"),
                },
                shared.FormField{
                    Choices: []string{
                        "crur",
                        "vinculum",
                        "tredecim",
                        "somniculosus",
                        "curtus",
                    },
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    DefaultValue: unifiedgosdk.Pointer("fugiat"),
                    ID: unifiedgosdk.Pointer("49ce19df-408a-4878-af8f-da2399a45d46"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "consequatur error cena",
                    Order: unifiedgosdk.Pointer[float64](2.0),
                    Type: shared.FormFieldTypeMultipleSelect,
                    UpdatedAt: types.MustNewTimeFromString("2024-03-01T23:53:51.435Z"),
                },
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    Description: unifiedgosdk.Pointer("Aurum voluptatum placeat audax tripudio amo."),
                    ID: unifiedgosdk.Pointer("88d5c961-dd16-41e4-af8a-767bb39f2e7e"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "audentia subseco antepono",
                    Order: unifiedgosdk.Pointer[float64](3.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2024-12-10T23:02:56.986Z"),
                },
            },
            HasMultipleSubmissions: unifiedgosdk.Pointer(false),
            HasProgressBar: unifiedgosdk.Pointer(false),
            HasShuffleQuestions: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("14c63393-a398-499a-96f0-dd8f6adb0b22"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.FormsMetadata{},
            Name: unifiedgosdk.Pointer("tandem tracto vivo Form"),
            PublishedURL: unifiedgosdk.Pointer("https://massive-blowgun.info"),
            ResponseCount: unifiedgosdk.Pointer[float64](423.0),
            UpdatedAt: types.MustNewTimeFromString("2024-09-04T06:48:05.896Z"),
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
                    ID: unifiedgosdk.Pointer("79894811-f797-4f7f-b210-c7ee1443ac8d"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(true),
                    MaxLength: unifiedgosdk.Pointer[float64](146.0),
                    Name: "voluntarius argumentum decumbo",
                    Order: unifiedgosdk.Pointer[float64](0.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2025-05-17T03:22:32.991Z"),
                },
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    ID: unifiedgosdk.Pointer("1808686e-8a59-4266-9c11-2580e8c68383"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "desidero clibanus calcar",
                    Order: unifiedgosdk.Pointer[float64](1.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2024-09-12T22:02:07.460Z"),
                },
                shared.FormField{
                    Choices: []string{
                        "crur",
                        "vinculum",
                        "tredecim",
                        "somniculosus",
                        "curtus",
                    },
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    DefaultValue: unifiedgosdk.Pointer("fugiat"),
                    ID: unifiedgosdk.Pointer("49ce19df-408a-4878-af8f-da2399a45d46"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "consequatur error cena",
                    Order: unifiedgosdk.Pointer[float64](2.0),
                    Type: shared.FormFieldTypeMultipleSelect,
                    UpdatedAt: types.MustNewTimeFromString("2024-03-01T23:53:51.435Z"),
                },
                shared.FormField{
                    CreatedAt: types.MustNewTimeFromString("2023-10-05T21:34:29.094Z"),
                    Description: unifiedgosdk.Pointer("Aurum voluptatum placeat audax tripudio amo."),
                    ID: unifiedgosdk.Pointer("88d5c961-dd16-41e4-af8a-767bb39f2e7e"),
                    IsActive: unifiedgosdk.Pointer(true),
                    IsRequired: unifiedgosdk.Pointer(false),
                    Name: "audentia subseco antepono",
                    Order: unifiedgosdk.Pointer[float64](3.0),
                    Type: shared.FormFieldTypeTextarea,
                    UpdatedAt: types.MustNewTimeFromString("2024-12-10T23:02:56.986Z"),
                },
            },
            HasMultipleSubmissions: unifiedgosdk.Pointer(false),
            HasProgressBar: unifiedgosdk.Pointer(false),
            HasShuffleQuestions: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("14c63393-a398-499a-96f0-dd8f6adb0b22"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.FormsMetadata{},
            Name: unifiedgosdk.Pointer("tandem tracto vivo Form"),
            PublishedURL: unifiedgosdk.Pointer("https://massive-blowgun.info"),
            ResponseCount: unifiedgosdk.Pointer[float64](423.0),
            UpdatedAt: types.MustNewTimeFromString("2024-09-04T06:48:05.896Z"),
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