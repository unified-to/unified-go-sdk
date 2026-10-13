# Matter

## Overview

### Available Operations

* [CreateLegalMatter](#createlegalmatter) - Create a matter
* [GetLegalMatter](#getlegalmatter) - Retrieve a matter
* [ListLegalMatters](#listlegalmatters) - List all matters
* [PatchLegalMatter](#patchlegalmatter) - Update a matter
* [RemoveLegalMatter](#removelegalmatter) - Remove a matter
* [UpdateLegalMatter](#updatelegalmatter) - Update a matter

## CreateLegalMatter

Create a matter

### Example Usage

<!-- UsageSnippet language="go" operationID="createLegalMatter" method="post" path="/legal/{connection_id}/matter" example="legal_matter" -->
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

    res, err := s.Matter.CreateLegalMatter(ctx, operations.CreateLegalMatterRequest{
        LegalMatter: shared.LegalMatter{
            BillingMethod: shared.BillingMethodOther.ToPointer(),
            ClientReference: unifiedgosdk.Pointer("qSrjIBAD"),
            CompanyIds: []string{},
            ContactIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2023-05-28T04:10:42.375Z"),
            Currency: unifiedgosdk.Pointer("BOB"),
            Description: unifiedgosdk.Pointer("Speciosus cresco aut cotidie temeritas arma."),
            ID: unifiedgosdk.Pointer("85d057d1-a7a1-463e-ba59-b87d85857a97"),
            Metadata: []shared.LegalMetadata{},
            Name: unifiedgosdk.Pointer("Walter v. Franecki - Spinka"),
            Number: unifiedgosdk.Pointer("2023-3062"),
            OpenedAt: types.MustNewTimeFromString("2023-07-05T00:10:52.688Z"),
            Parties: []shared.LegalParty{
                shared.LegalParty{
                    Description: unifiedgosdk.Pointer("Urban Walter"),
                    Role: shared.LegalPartyRoleExpert.ToPointer(),
                },
            },
            Stage: unifiedgosdk.Pointer("currus"),
            Status: shared.LegalMatterStatusOpen.ToPointer(),
            StatuteOfLimitationsAt: types.MustNewTimeFromString("2027-06-20T13:19:08.077Z"),
            UpdatedAt: types.MustNewTimeFromString("2023-07-17T02:00:08.639Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalMatter != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateLegalMatterRequest](../../pkg/models/operations/createlegalmatterrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateLegalMatterResponse](../../pkg/models/operations/createlegalmatterresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetLegalMatter

Retrieve a matter

### Example Usage

<!-- UsageSnippet language="go" operationID="getLegalMatter" method="get" path="/legal/{connection_id}/matter/{id}" -->
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

    res, err := s.Matter.GetLegalMatter(ctx, operations.GetLegalMatterRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalMatter != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetLegalMatterRequest](../../pkg/models/operations/getlegalmatterrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetLegalMatterResponse](../../pkg/models/operations/getlegalmatterresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListLegalMatters

List all matters

### Example Usage

<!-- UsageSnippet language="go" operationID="listLegalMatters" method="get" path="/legal/{connection_id}/matter" -->
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

    res, err := s.Matter.ListLegalMatters(ctx, operations.ListLegalMattersRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalMatters != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListLegalMattersRequest](../../pkg/models/operations/listlegalmattersrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListLegalMattersResponse](../../pkg/models/operations/listlegalmattersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchLegalMatter

Update a matter

### Example Usage

<!-- UsageSnippet language="go" operationID="patchLegalMatter" method="patch" path="/legal/{connection_id}/matter/{id}" example="legal_matter" -->
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

    res, err := s.Matter.PatchLegalMatter(ctx, operations.PatchLegalMatterRequest{
        LegalMatter: shared.LegalMatter{
            BillingMethod: shared.BillingMethodOther.ToPointer(),
            ClientReference: unifiedgosdk.Pointer("qSrjIBAD"),
            CompanyIds: []string{},
            ContactIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2023-05-28T04:10:42.375Z"),
            Currency: unifiedgosdk.Pointer("BOB"),
            Description: unifiedgosdk.Pointer("Speciosus cresco aut cotidie temeritas arma."),
            ID: unifiedgosdk.Pointer("22792f41-4200-43e3-9282-2bb50ed285e5"),
            Metadata: []shared.LegalMetadata{},
            Name: unifiedgosdk.Pointer("Walter v. Franecki - Spinka"),
            Number: unifiedgosdk.Pointer("2023-3062"),
            OpenedAt: types.MustNewTimeFromString("2023-07-05T00:10:52.689Z"),
            Parties: []shared.LegalParty{
                shared.LegalParty{
                    Description: unifiedgosdk.Pointer("Urban Walter"),
                    Role: shared.LegalPartyRoleExpert.ToPointer(),
                },
            },
            Stage: unifiedgosdk.Pointer("currus"),
            Status: shared.LegalMatterStatusOpen.ToPointer(),
            StatuteOfLimitationsAt: types.MustNewTimeFromString("2027-06-20T13:19:08.104Z"),
            UpdatedAt: types.MustNewTimeFromString("2023-07-17T02:00:08.640Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalMatter != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchLegalMatterRequest](../../pkg/models/operations/patchlegalmatterrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchLegalMatterResponse](../../pkg/models/operations/patchlegalmatterresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveLegalMatter

Remove a matter

### Example Usage

<!-- UsageSnippet language="go" operationID="removeLegalMatter" method="delete" path="/legal/{connection_id}/matter/{id}" -->
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

    res, err := s.Matter.RemoveLegalMatter(ctx, operations.RemoveLegalMatterRequest{
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.RemoveLegalMatterRequest](../../pkg/models/operations/removelegalmatterrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemoveLegalMatterResponse](../../pkg/models/operations/removelegalmatterresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateLegalMatter

Update a matter

### Example Usage

<!-- UsageSnippet language="go" operationID="updateLegalMatter" method="put" path="/legal/{connection_id}/matter/{id}" example="legal_matter" -->
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

    res, err := s.Matter.UpdateLegalMatter(ctx, operations.UpdateLegalMatterRequest{
        LegalMatter: shared.LegalMatter{
            BillingMethod: shared.BillingMethodOther.ToPointer(),
            ClientReference: unifiedgosdk.Pointer("qSrjIBAD"),
            CompanyIds: []string{},
            ContactIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2023-05-28T04:10:42.375Z"),
            Currency: unifiedgosdk.Pointer("BOB"),
            Description: unifiedgosdk.Pointer("Speciosus cresco aut cotidie temeritas arma."),
            ID: unifiedgosdk.Pointer("22792f41-4200-43e3-9282-2bb50ed285e5"),
            Metadata: []shared.LegalMetadata{},
            Name: unifiedgosdk.Pointer("Walter v. Franecki - Spinka"),
            Number: unifiedgosdk.Pointer("2023-3062"),
            OpenedAt: types.MustNewTimeFromString("2023-07-05T00:10:52.689Z"),
            Parties: []shared.LegalParty{
                shared.LegalParty{
                    Description: unifiedgosdk.Pointer("Urban Walter"),
                    Role: shared.LegalPartyRoleExpert.ToPointer(),
                },
            },
            Stage: unifiedgosdk.Pointer("currus"),
            Status: shared.LegalMatterStatusOpen.ToPointer(),
            StatuteOfLimitationsAt: types.MustNewTimeFromString("2027-06-20T13:19:08.104Z"),
            UpdatedAt: types.MustNewTimeFromString("2023-07-17T02:00:08.640Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalMatter != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateLegalMatterRequest](../../pkg/models/operations/updatelegalmatterrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateLegalMatterResponse](../../pkg/models/operations/updatelegalmatterresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |