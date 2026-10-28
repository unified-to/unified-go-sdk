# Assessment

## Overview

### Available Operations

* [CreateAssessmentOrder](#createassessmentorder) - Create an order
* [CreateAssessmentPackage](#createassessmentpackage) - Create an assessment package
* [GetAssessmentOrder](#getassessmentorder) - Retrieve an order
* [GetAssessmentPackage](#getassessmentpackage) - Get an assessment package
* [ListAssessmentPackages](#listassessmentpackages) - List assessment packages
* [PatchAssessmentOrder](#patchassessmentorder) - Update an order
* [PatchAssessmentPackage](#patchassessmentpackage) - Update an assessment package
* [RemoveAssessmentPackage](#removeassessmentpackage) - Delete an assessment package
* [UpdateAssessmentOrder](#updateassessmentorder) - Update an order
* [UpdateAssessmentPackage](#updateassessmentpackage) - Update an assessment package

## CreateAssessmentOrder

Creates a AssessmentOrder. Provide the unified AssessmentOrder fields in the request body; they are mapped to the provider's API. Returns the created AssessmentOrder.

### Example Usage

<!-- UsageSnippet language="go" operationID="createAssessmentOrder" method="post" path="/assessment/{connection_id}/order" example="assessment_order" -->
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

    res, err := s.Assessment.CreateAssessmentOrder(ctx, operations.CreateAssessmentOrderRequest{
        AssessmentOrder: shared.AssessmentOrder{
            ConnectionID: "<id>",
            CreatedAt: types.MustNewTimeFromString("2021-09-18T10:33:57.803Z"),
            ID: unifiedgosdk.Pointer("1bf0ecc5-560f-4197-88a4-ca04193db9e5"),
            Metadata: []shared.AssessmentMetadata{},
            Parameters: []shared.AssessmentParameterInput{},
            ProfileAddresses: []shared.AssessmentAddress{},
            ProfileDateOfBirth: unifiedgosdk.Pointer("1989-07-22T16:18:37.650Z"),
            ProfileEmails: []string{
                "amy.kris.windler@icloud.com",
            },
            ProfileFirstName: unifiedgosdk.Pointer("Amy"),
            ProfileGender: shared.ProfileGenderNonBinary.ToPointer(),
            ProfileLastName: unifiedgosdk.Pointer("Kris-Windler"),
            ProfileName: unifiedgosdk.Pointer("Amy Kris-Windler"),
            ProfileResumeURL: unifiedgosdk.Pointer("https://enchanted-cycle.biz/"),
            ProfileSocialMediaUrls: []string{},
            ProfileTelephones: []string{
                "(341) 425-5828",
            },
            Reference: unifiedgosdk.Pointer("ab"),
            ResponseAttributes: []shared.AssessmentAttribute{},
            ResponseDetails: []shared.AssessmentResponseDetail{},
            ResponseDownloadUrls: []string{},
            ResponseMaxScore: unifiedgosdk.Pointer[float64](84.0),
            ResponseScore: unifiedgosdk.Pointer[float64](82.0),
            ResponseStatus: shared.ResponseStatusRejected.ToPointer(),
            ResponseURL: unifiedgosdk.Pointer("https://rowdy-fisherman.biz/"),
            Status: shared.AssessmentOrderStatusRejected.ToPointer(),
            TargetURL: unifiedgosdk.Pointer("https://cautious-turret.info"),
            UpdatedAt: types.MustNewTimeFromString("2023-01-27T17:29:31.775Z"),
            WorkspaceID: "<id>",
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssessmentOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CreateAssessmentOrderRequest](../../pkg/models/operations/createassessmentorderrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.CreateAssessmentOrderResponse](../../pkg/models/operations/createassessmentorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateAssessmentPackage

Create an assessment package

### Example Usage

<!-- UsageSnippet language="go" operationID="createAssessmentPackage" method="post" path="/assessment/{connection_id}/package" example="assessment_package" -->
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

    res, err := s.Assessment.CreateAssessmentPackage(ctx, operations.CreateAssessmentPackageRequest{
        AssessmentPackage: shared.AssessmentPackage{
            Aliases: []string{
                "quia",
            },
            CreatedAt: types.MustNewTimeFromString("2022-11-18T19:48:39.433Z"),
            Description: unifiedgosdk.Pointer("Eos aedificium consectetur urbs. Admitto summa accusator tabesco distinctio vapulus culpo templum ancilla."),
            HasRedirectURL: unifiedgosdk.Pointer(true),
            HasTargetURL: unifiedgosdk.Pointer(false),
            ID: unifiedgosdk.Pointer("97427f92-9a67-4d84-836c-095a8a15b895"),
            InfoURL: unifiedgosdk.Pointer("https://ugly-instance.biz/"),
            IntegrationTypes: []string{
                "viridis",
            },
            MaxScore: unifiedgosdk.Pointer[float64](22.0),
            Metadata: []shared.AssessmentMetadata{},
            Name: unifiedgosdk.Pointer("Carus sed vox doloremque vigor surgo tabella cupiditas abduco clarus."),
            NeedsIPAddress: unifiedgosdk.Pointer(true),
            Parameters: []shared.AssessmentParameter{},
            Regions: []shared.AssessmentPackageRegion{},
            Tags: []string{
                "clamo",
            },
            Type: shared.AssessmentPackageTypeVideoInterview,
            UpdatedAt: types.MustNewTimeFromString("2023-09-26T21:51:10.588Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssessmentPackage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.CreateAssessmentPackageRequest](../../pkg/models/operations/createassessmentpackagerequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.CreateAssessmentPackageResponse](../../pkg/models/operations/createassessmentpackageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAssessmentOrder

Retrieves a single AssessmentOrder by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAssessmentOrder" method="get" path="/assessment/{connection_id}/order/{id}" -->
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

    res, err := s.Assessment.GetAssessmentOrder(ctx, operations.GetAssessmentOrderRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssessmentOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetAssessmentOrderRequest](../../pkg/models/operations/getassessmentorderrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetAssessmentOrderResponse](../../pkg/models/operations/getassessmentorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAssessmentPackage

Get an assessment package

### Example Usage

<!-- UsageSnippet language="go" operationID="getAssessmentPackage" method="get" path="/assessment/{connection_id}/package/{id}" -->
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

    res, err := s.Assessment.GetAssessmentPackage(ctx, operations.GetAssessmentPackageRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssessmentPackage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetAssessmentPackageRequest](../../pkg/models/operations/getassessmentpackagerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetAssessmentPackageResponse](../../pkg/models/operations/getassessmentpackageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAssessmentPackages

List assessment packages

### Example Usage

<!-- UsageSnippet language="go" operationID="listAssessmentPackages" method="get" path="/assessment/{connection_id}/package" -->
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

    res, err := s.Assessment.ListAssessmentPackages(ctx, operations.ListAssessmentPackagesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssessmentPackages != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListAssessmentPackagesRequest](../../pkg/models/operations/listassessmentpackagesrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.ListAssessmentPackagesResponse](../../pkg/models/operations/listassessmentpackagesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAssessmentOrder

Updates the AssessmentOrder identified by `id`. Provide the unified AssessmentOrder fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAssessmentOrder" method="patch" path="/assessment/{connection_id}/order/{id}" example="assessment_order" -->
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

    res, err := s.Assessment.PatchAssessmentOrder(ctx, operations.PatchAssessmentOrderRequest{
        AssessmentOrder: shared.AssessmentOrder{
            ConnectionID: "<id>",
            CreatedAt: types.MustNewTimeFromString("2021-09-18T10:33:57.803Z"),
            ID: unifiedgosdk.Pointer("ac628dce-0756-457f-a8d6-5f3c31f19dd9"),
            Metadata: []shared.AssessmentMetadata{},
            Parameters: []shared.AssessmentParameterInput{},
            ProfileAddresses: []shared.AssessmentAddress{},
            ProfileDateOfBirth: unifiedgosdk.Pointer("1989-07-22T16:18:37.650Z"),
            ProfileEmails: []string{
                "amy.kris.windler@icloud.com",
            },
            ProfileFirstName: unifiedgosdk.Pointer("Amy"),
            ProfileGender: shared.ProfileGenderNonBinary.ToPointer(),
            ProfileLastName: unifiedgosdk.Pointer("Kris-Windler"),
            ProfileName: unifiedgosdk.Pointer("Amy Kris-Windler"),
            ProfileResumeURL: unifiedgosdk.Pointer("https://enchanted-cycle.biz/"),
            ProfileSocialMediaUrls: []string{},
            ProfileTelephones: []string{
                "(341) 425-5828",
            },
            Reference: unifiedgosdk.Pointer("ab"),
            ResponseAttributes: []shared.AssessmentAttribute{},
            ResponseDetails: []shared.AssessmentResponseDetail{},
            ResponseDownloadUrls: []string{},
            ResponseMaxScore: unifiedgosdk.Pointer[float64](84.0),
            ResponseScore: unifiedgosdk.Pointer[float64](82.0),
            ResponseStatus: shared.ResponseStatusRejected.ToPointer(),
            ResponseURL: unifiedgosdk.Pointer("https://rowdy-fisherman.biz/"),
            Status: shared.AssessmentOrderStatusRejected.ToPointer(),
            TargetURL: unifiedgosdk.Pointer("https://cautious-turret.info"),
            UpdatedAt: types.MustNewTimeFromString("2023-01-27T17:29:31.783Z"),
            WorkspaceID: "<id>",
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssessmentOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PatchAssessmentOrderRequest](../../pkg/models/operations/patchassessmentorderrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.PatchAssessmentOrderResponse](../../pkg/models/operations/patchassessmentorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAssessmentPackage

Update an assessment package

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAssessmentPackage" method="patch" path="/assessment/{connection_id}/package/{id}" example="assessment_package" -->
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

    res, err := s.Assessment.PatchAssessmentPackage(ctx, operations.PatchAssessmentPackageRequest{
        AssessmentPackage: shared.AssessmentPackage{
            Aliases: []string{
                "quia",
            },
            CreatedAt: types.MustNewTimeFromString("2022-11-18T19:48:39.433Z"),
            Description: unifiedgosdk.Pointer("Eos aedificium consectetur urbs. Admitto summa accusator tabesco distinctio vapulus culpo templum ancilla."),
            HasRedirectURL: unifiedgosdk.Pointer(true),
            HasTargetURL: unifiedgosdk.Pointer(false),
            ID: unifiedgosdk.Pointer("168ebddd-95b7-4e33-aebb-441c937365ab"),
            InfoURL: unifiedgosdk.Pointer("https://ugly-instance.biz/"),
            IntegrationTypes: []string{
                "viridis",
            },
            MaxScore: unifiedgosdk.Pointer[float64](22.0),
            Metadata: []shared.AssessmentMetadata{},
            Name: unifiedgosdk.Pointer("Carus sed vox doloremque vigor surgo tabella cupiditas abduco clarus."),
            NeedsIPAddress: unifiedgosdk.Pointer(true),
            Parameters: []shared.AssessmentParameter{},
            Regions: []shared.AssessmentPackageRegion{},
            Tags: []string{
                "clamo",
            },
            Type: shared.AssessmentPackageTypeVideoInterview,
            UpdatedAt: types.MustNewTimeFromString("2023-09-26T21:51:10.591Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssessmentPackage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PatchAssessmentPackageRequest](../../pkg/models/operations/patchassessmentpackagerequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.PatchAssessmentPackageResponse](../../pkg/models/operations/patchassessmentpackageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAssessmentPackage

Delete an assessment package

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAssessmentPackage" method="delete" path="/assessment/{connection_id}/package/{id}" -->
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

    res, err := s.Assessment.RemoveAssessmentPackage(ctx, operations.RemoveAssessmentPackageRequest{
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.RemoveAssessmentPackageRequest](../../pkg/models/operations/removeassessmentpackagerequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.RemoveAssessmentPackageResponse](../../pkg/models/operations/removeassessmentpackageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAssessmentOrder

Updates the AssessmentOrder identified by `id`. Provide the unified AssessmentOrder fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAssessmentOrder" method="put" path="/assessment/{connection_id}/order/{id}" example="assessment_order" -->
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

    res, err := s.Assessment.UpdateAssessmentOrder(ctx, operations.UpdateAssessmentOrderRequest{
        AssessmentOrder: shared.AssessmentOrder{
            ConnectionID: "<id>",
            CreatedAt: types.MustNewTimeFromString("2021-09-18T10:33:57.803Z"),
            ID: unifiedgosdk.Pointer("ac628dce-0756-457f-a8d6-5f3c31f19dd9"),
            Metadata: []shared.AssessmentMetadata{},
            Parameters: []shared.AssessmentParameterInput{},
            ProfileAddresses: []shared.AssessmentAddress{},
            ProfileDateOfBirth: unifiedgosdk.Pointer("1989-07-22T16:18:37.650Z"),
            ProfileEmails: []string{
                "amy.kris.windler@icloud.com",
            },
            ProfileFirstName: unifiedgosdk.Pointer("Amy"),
            ProfileGender: shared.ProfileGenderNonBinary.ToPointer(),
            ProfileLastName: unifiedgosdk.Pointer("Kris-Windler"),
            ProfileName: unifiedgosdk.Pointer("Amy Kris-Windler"),
            ProfileResumeURL: unifiedgosdk.Pointer("https://enchanted-cycle.biz/"),
            ProfileSocialMediaUrls: []string{},
            ProfileTelephones: []string{
                "(341) 425-5828",
            },
            Reference: unifiedgosdk.Pointer("ab"),
            ResponseAttributes: []shared.AssessmentAttribute{},
            ResponseDetails: []shared.AssessmentResponseDetail{},
            ResponseDownloadUrls: []string{},
            ResponseMaxScore: unifiedgosdk.Pointer[float64](84.0),
            ResponseScore: unifiedgosdk.Pointer[float64](82.0),
            ResponseStatus: shared.ResponseStatusRejected.ToPointer(),
            ResponseURL: unifiedgosdk.Pointer("https://rowdy-fisherman.biz/"),
            Status: shared.AssessmentOrderStatusRejected.ToPointer(),
            TargetURL: unifiedgosdk.Pointer("https://cautious-turret.info"),
            UpdatedAt: types.MustNewTimeFromString("2023-01-27T17:29:31.783Z"),
            WorkspaceID: "<id>",
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssessmentOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdateAssessmentOrderRequest](../../pkg/models/operations/updateassessmentorderrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UpdateAssessmentOrderResponse](../../pkg/models/operations/updateassessmentorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAssessmentPackage

Update an assessment package

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAssessmentPackage" method="put" path="/assessment/{connection_id}/package/{id}" example="assessment_package" -->
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

    res, err := s.Assessment.UpdateAssessmentPackage(ctx, operations.UpdateAssessmentPackageRequest{
        AssessmentPackage: shared.AssessmentPackage{
            Aliases: []string{
                "quia",
            },
            CreatedAt: types.MustNewTimeFromString("2022-11-18T19:48:39.433Z"),
            Description: unifiedgosdk.Pointer("Eos aedificium consectetur urbs. Admitto summa accusator tabesco distinctio vapulus culpo templum ancilla."),
            HasRedirectURL: unifiedgosdk.Pointer(true),
            HasTargetURL: unifiedgosdk.Pointer(false),
            ID: unifiedgosdk.Pointer("168ebddd-95b7-4e33-aebb-441c937365ab"),
            InfoURL: unifiedgosdk.Pointer("https://ugly-instance.biz/"),
            IntegrationTypes: []string{
                "viridis",
            },
            MaxScore: unifiedgosdk.Pointer[float64](22.0),
            Metadata: []shared.AssessmentMetadata{},
            Name: unifiedgosdk.Pointer("Carus sed vox doloremque vigor surgo tabella cupiditas abduco clarus."),
            NeedsIPAddress: unifiedgosdk.Pointer(true),
            Parameters: []shared.AssessmentParameter{},
            Regions: []shared.AssessmentPackageRegion{},
            Tags: []string{
                "clamo",
            },
            Type: shared.AssessmentPackageTypeVideoInterview,
            UpdatedAt: types.MustNewTimeFromString("2023-09-26T21:51:10.591Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssessmentPackage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.UpdateAssessmentPackageRequest](../../pkg/models/operations/updateassessmentpackagerequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.UpdateAssessmentPackageResponse](../../pkg/models/operations/updateassessmentpackageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |