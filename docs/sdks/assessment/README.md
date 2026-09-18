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

Create an order

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
            ID: unifiedgosdk.Pointer("d2c7a88c-4973-4f3a-977c-36e91d6bbb66"),
            Parameters: []shared.AssessmentParameterInput{},
            ProfileAddresses: []shared.AssessmentAddress{},
            ProfileDateOfBirth: unifiedgosdk.Pointer("1989-07-22T16:18:37.650Z"),
            ProfileEmails: []string{
                "Cleta.Daugherty@gmail.com",
            },
            ProfileFirstName: unifiedgosdk.Pointer("Amy"),
            ProfileGender: shared.ProfileGenderNonBinary.ToPointer(),
            ProfileLastName: unifiedgosdk.Pointer("Kris-Windler"),
            ProfileName: unifiedgosdk.Pointer("Amy Kris-Windler"),
            ProfileResumeURL: unifiedgosdk.Pointer("https://enchanted-cycle.biz/"),
            ProfileSocialMediaUrls: []string{},
            ProfileTelephones: []string{
                "(828) 263-1594 x5248",
            },
            Reference: unifiedgosdk.Pointer("ab"),
            ResponseAttributes: []shared.AssessmentAttribute{},
            ResponseDetails: []shared.AssessmentResponseDetail{},
            ResponseDownloadUrls: []string{},
            ResponseMaxScore: unifiedgosdk.Pointer[float64](82.0),
            ResponseScore: unifiedgosdk.Pointer[float64](92.0),
            ResponseStatus: shared.ResponseStatusFailed.ToPointer(),
            ResponseURL: unifiedgosdk.Pointer("https://irresponsible-trench.info/"),
            Status: shared.AssessmentOrderStatusRejected.ToPointer(),
            TargetURL: unifiedgosdk.Pointer("https://cautious-turret.info"),
            UpdatedAt: types.MustNewTimeFromString("2023-01-17T02:08:14.501Z"),
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
            ID: unifiedgosdk.Pointer("b5c3a4cc-2da0-49d3-aa18-cbd03d5213ce"),
            InfoURL: unifiedgosdk.Pointer("https://ugly-instance.biz/"),
            IntegrationTypes: []string{
                "viridis",
            },
            MaxScore: unifiedgosdk.Pointer[float64](22.0),
            Name: unifiedgosdk.Pointer("Carus sed vox doloremque vigor surgo tabella cupiditas abduco clarus."),
            NeedsIPAddress: unifiedgosdk.Pointer(true),
            Parameters: []shared.AssessmentParameter{},
            Regions: []shared.AssessmentPackageRegion{},
            Tags: []string{
                "clamo",
            },
            Type: shared.AssessmentPackageTypeVideoInterview,
            UpdatedAt: types.MustNewTimeFromString("2023-09-18T05:42:09.559Z"),
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

Retrieve an order

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

Update an order

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
            ID: unifiedgosdk.Pointer("ab8d64eb-a6c2-4128-a202-df4bd71d26a7"),
            Parameters: []shared.AssessmentParameterInput{},
            ProfileAddresses: []shared.AssessmentAddress{},
            ProfileDateOfBirth: unifiedgosdk.Pointer("1989-07-22T16:18:37.650Z"),
            ProfileEmails: []string{
                "Cleta.Daugherty@gmail.com",
            },
            ProfileFirstName: unifiedgosdk.Pointer("Amy"),
            ProfileGender: shared.ProfileGenderNonBinary.ToPointer(),
            ProfileLastName: unifiedgosdk.Pointer("Kris-Windler"),
            ProfileName: unifiedgosdk.Pointer("Amy Kris-Windler"),
            ProfileResumeURL: unifiedgosdk.Pointer("https://enchanted-cycle.biz/"),
            ProfileSocialMediaUrls: []string{},
            ProfileTelephones: []string{
                "(828) 263-1594 x5248",
            },
            Reference: unifiedgosdk.Pointer("ab"),
            ResponseAttributes: []shared.AssessmentAttribute{},
            ResponseDetails: []shared.AssessmentResponseDetail{},
            ResponseDownloadUrls: []string{},
            ResponseMaxScore: unifiedgosdk.Pointer[float64](82.0),
            ResponseScore: unifiedgosdk.Pointer[float64](92.0),
            ResponseStatus: shared.ResponseStatusFailed.ToPointer(),
            ResponseURL: unifiedgosdk.Pointer("https://irresponsible-trench.info/"),
            Status: shared.AssessmentOrderStatusRejected.ToPointer(),
            TargetURL: unifiedgosdk.Pointer("https://cautious-turret.info"),
            UpdatedAt: types.MustNewTimeFromString("2023-01-17T02:08:14.507Z"),
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
            ID: unifiedgosdk.Pointer("6632ab60-bde2-454d-a228-59b2e47451f4"),
            InfoURL: unifiedgosdk.Pointer("https://ugly-instance.biz/"),
            IntegrationTypes: []string{
                "viridis",
            },
            MaxScore: unifiedgosdk.Pointer[float64](22.0),
            Name: unifiedgosdk.Pointer("Carus sed vox doloremque vigor surgo tabella cupiditas abduco clarus."),
            NeedsIPAddress: unifiedgosdk.Pointer(true),
            Parameters: []shared.AssessmentParameter{},
            Regions: []shared.AssessmentPackageRegion{},
            Tags: []string{
                "clamo",
            },
            Type: shared.AssessmentPackageTypeVideoInterview,
            UpdatedAt: types.MustNewTimeFromString("2023-09-18T05:42:09.562Z"),
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

Update an order

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
            ID: unifiedgosdk.Pointer("ab8d64eb-a6c2-4128-a202-df4bd71d26a7"),
            Parameters: []shared.AssessmentParameterInput{},
            ProfileAddresses: []shared.AssessmentAddress{},
            ProfileDateOfBirth: unifiedgosdk.Pointer("1989-07-22T16:18:37.650Z"),
            ProfileEmails: []string{
                "Cleta.Daugherty@gmail.com",
            },
            ProfileFirstName: unifiedgosdk.Pointer("Amy"),
            ProfileGender: shared.ProfileGenderNonBinary.ToPointer(),
            ProfileLastName: unifiedgosdk.Pointer("Kris-Windler"),
            ProfileName: unifiedgosdk.Pointer("Amy Kris-Windler"),
            ProfileResumeURL: unifiedgosdk.Pointer("https://enchanted-cycle.biz/"),
            ProfileSocialMediaUrls: []string{},
            ProfileTelephones: []string{
                "(828) 263-1594 x5248",
            },
            Reference: unifiedgosdk.Pointer("ab"),
            ResponseAttributes: []shared.AssessmentAttribute{},
            ResponseDetails: []shared.AssessmentResponseDetail{},
            ResponseDownloadUrls: []string{},
            ResponseMaxScore: unifiedgosdk.Pointer[float64](82.0),
            ResponseScore: unifiedgosdk.Pointer[float64](92.0),
            ResponseStatus: shared.ResponseStatusFailed.ToPointer(),
            ResponseURL: unifiedgosdk.Pointer("https://irresponsible-trench.info/"),
            Status: shared.AssessmentOrderStatusRejected.ToPointer(),
            TargetURL: unifiedgosdk.Pointer("https://cautious-turret.info"),
            UpdatedAt: types.MustNewTimeFromString("2023-01-17T02:08:14.507Z"),
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
            ID: unifiedgosdk.Pointer("6632ab60-bde2-454d-a228-59b2e47451f4"),
            InfoURL: unifiedgosdk.Pointer("https://ugly-instance.biz/"),
            IntegrationTypes: []string{
                "viridis",
            },
            MaxScore: unifiedgosdk.Pointer[float64](22.0),
            Name: unifiedgosdk.Pointer("Carus sed vox doloremque vigor surgo tabella cupiditas abduco clarus."),
            NeedsIPAddress: unifiedgosdk.Pointer(true),
            Parameters: []shared.AssessmentParameter{},
            Regions: []shared.AssessmentPackageRegion{},
            Tags: []string{
                "clamo",
            },
            Type: shared.AssessmentPackageTypeVideoInterview,
            UpdatedAt: types.MustNewTimeFromString("2023-09-18T05:42:09.562Z"),
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