# Legal

## Overview

### Available Operations

* [CreateLegalExpense](#createlegalexpense) - Create an expense
* [CreateLegalMatter](#createlegalmatter) - Create a matter
* [CreateLegalPracticearea](#createlegalpracticearea) - Create a practicearea
* [CreateLegalTimeentry](#createlegaltimeentry) - Create a timeentry
* [GetLegalExpense](#getlegalexpense) - Retrieve an expense
* [GetLegalMatter](#getlegalmatter) - Retrieve a matter
* [GetLegalPracticearea](#getlegalpracticearea) - Retrieve a practicearea
* [GetLegalTimeentry](#getlegaltimeentry) - Retrieve a timeentry
* [ListLegalExpenses](#listlegalexpenses) - List all expenses
* [ListLegalMatters](#listlegalmatters) - List all matters
* [ListLegalPracticeareas](#listlegalpracticeareas) - List all practiceareas
* [ListLegalTimeentries](#listlegaltimeentries) - List all timeentries
* [PatchLegalExpense](#patchlegalexpense) - Update an expense
* [PatchLegalMatter](#patchlegalmatter) - Update a matter
* [PatchLegalPracticearea](#patchlegalpracticearea) - Update a practicearea
* [PatchLegalTimeentry](#patchlegaltimeentry) - Update a timeentry
* [RemoveLegalExpense](#removelegalexpense) - Remove an expense
* [RemoveLegalMatter](#removelegalmatter) - Remove a matter
* [RemoveLegalPracticearea](#removelegalpracticearea) - Remove a practicearea
* [RemoveLegalTimeentry](#removelegaltimeentry) - Remove a timeentry
* [UpdateLegalExpense](#updatelegalexpense) - Update an expense
* [UpdateLegalMatter](#updatelegalmatter) - Update a matter
* [UpdateLegalPracticearea](#updatelegalpracticearea) - Update a practicearea
* [UpdateLegalTimeentry](#updatelegaltimeentry) - Update a timeentry

## CreateLegalExpense

Create an expense

### Example Usage

<!-- UsageSnippet language="go" operationID="createLegalExpense" method="post" path="/legal/{connection_id}/expense" example="legal_expense" -->
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

    res, err := s.Legal.CreateLegalExpense(ctx, operations.CreateLegalExpenseRequest{
        LegalExpense: shared.LegalExpense{
            Category: unifiedgosdk.Pointer("Kids"),
            CreatedAt: types.MustNewTimeFromString("2020-11-08T00:41:42.356Z"),
            Currency: unifiedgosdk.Pointer("MVR"),
            Description: unifiedgosdk.Pointer("Temporibus agnitio sui."),
            ID: unifiedgosdk.Pointer("4b5483e3-7b1c-4b1b-9ae2-53a8db8258ba"),
            IncurredAt: types.MustNewTimeFromString("2020-11-08T00:41:42.356Z"),
            IsBilled: unifiedgosdk.Pointer(false),
            Metadata: []shared.LegalMetadata{},
            PriceAmount: unifiedgosdk.Pointer[float64](100.0),
            Quantity: unifiedgosdk.Pointer[float64](3.0),
            TotalAmount: unifiedgosdk.Pointer[float64](300.0),
            Type: shared.LegalExpenseTypeOther.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-03-25T18:37:19.750Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalExpense != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateLegalExpenseRequest](../../pkg/models/operations/createlegalexpenserequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateLegalExpenseResponse](../../pkg/models/operations/createlegalexpenseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

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

    res, err := s.Legal.CreateLegalMatter(ctx, operations.CreateLegalMatterRequest{
        LegalMatter: shared.LegalMatter{
            BillingMethod: shared.BillingMethodOther.ToPointer(),
            ClientReference: unifiedgosdk.Pointer("qSrjIBAD"),
            CompanyIds: []string{},
            ContactIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2023-05-28T04:10:42.375Z"),
            Currency: unifiedgosdk.Pointer("BOB"),
            Description: unifiedgosdk.Pointer("Speciosus cresco aut cotidie temeritas arma."),
            ID: unifiedgosdk.Pointer("cbcdc766-83f8-4b6e-997e-23f454184f04"),
            Metadata: []shared.LegalMetadata{},
            Name: unifiedgosdk.Pointer("Walter v. Franecki - Spinka"),
            Number: unifiedgosdk.Pointer("2023-3062"),
            OpenedAt: types.MustNewTimeFromString("2023-07-04T23:36:07.862Z"),
            Parties: []shared.LegalParty{
                shared.LegalParty{
                    Description: unifiedgosdk.Pointer("Urban Walter"),
                    Role: shared.LegalPartyRoleExpert.ToPointer(),
                },
            },
            Stage: unifiedgosdk.Pointer("currus"),
            Status: shared.LegalMatterStatusOpen.ToPointer(),
            StatuteOfLimitationsAt: types.MustNewTimeFromString("2027-06-19T18:26:16.494Z"),
            UpdatedAt: types.MustNewTimeFromString("2023-07-17T01:14:18.367Z"),
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

## CreateLegalPracticearea

Create a practicearea

### Example Usage

<!-- UsageSnippet language="go" operationID="createLegalPracticearea" method="post" path="/legal/{connection_id}/practicearea" example="legal_practicearea" -->
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

    res, err := s.Legal.CreateLegalPracticearea(ctx, operations.CreateLegalPracticeareaRequest{
        LegalPracticearea: shared.LegalPracticearea{
            Code: unifiedgosdk.Pointer("QCA"),
            CreatedAt: types.MustNewTimeFromString("2023-01-28T21:00:26.264Z"),
            ID: unifiedgosdk.Pointer("16f76e0f-cea1-4921-93fa-9521017ff8b0"),
            Metadata: []shared.LegalMetadata{},
            Name: unifiedgosdk.Pointer("witty Law"),
            Type: shared.LegalPracticeareaTypeEducation.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-08-21T09:46:07.829Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalPracticearea != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.CreateLegalPracticeareaRequest](../../pkg/models/operations/createlegalpracticearearequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.CreateLegalPracticeareaResponse](../../pkg/models/operations/createlegalpracticearearesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateLegalTimeentry

Create a timeentry

### Example Usage

<!-- UsageSnippet language="go" operationID="createLegalTimeentry" method="post" path="/legal/{connection_id}/timeentry" example="legal_timeentry" -->
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

    res, err := s.Legal.CreateLegalTimeentry(ctx, operations.CreateLegalTimeentryRequest{
        LegalTimeentry: shared.LegalTimeentry{
            BillableMinutes: unifiedgosdk.Pointer[float64](390.0),
            CreatedAt: types.MustNewTimeFromString("2022-05-06T03:26:03.756Z"),
            Currency: unifiedgosdk.Pointer("ZAR"),
            Description: unifiedgosdk.Pointer("Ea derideo adulatio torrens."),
            ID: unifiedgosdk.Pointer("78f43089-9c66-4226-96a6-9d168c54d193"),
            IsBillable: unifiedgosdk.Pointer(true),
            Metadata: []shared.LegalMetadata{},
            Minutes: unifiedgosdk.Pointer[float64](390.0),
            RateAmount: unifiedgosdk.Pointer[float64](637.0),
            TotalAmount: unifiedgosdk.Pointer[float64](4140.5),
            UpdatedAt: types.MustNewTimeFromString("2022-06-29T18:26:46.380Z"),
            WorkedAt: types.MustNewTimeFromString("2022-05-06T03:26:03.756Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalTimeentry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateLegalTimeentryRequest](../../pkg/models/operations/createlegaltimeentryrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.CreateLegalTimeentryResponse](../../pkg/models/operations/createlegaltimeentryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetLegalExpense

Retrieve an expense

### Example Usage

<!-- UsageSnippet language="go" operationID="getLegalExpense" method="get" path="/legal/{connection_id}/expense/{id}" -->
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

    res, err := s.Legal.GetLegalExpense(ctx, operations.GetLegalExpenseRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalExpense != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetLegalExpenseRequest](../../pkg/models/operations/getlegalexpenserequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetLegalExpenseResponse](../../pkg/models/operations/getlegalexpenseresponse.md), error**

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

    res, err := s.Legal.GetLegalMatter(ctx, operations.GetLegalMatterRequest{
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

## GetLegalPracticearea

Retrieve a practicearea

### Example Usage

<!-- UsageSnippet language="go" operationID="getLegalPracticearea" method="get" path="/legal/{connection_id}/practicearea/{id}" -->
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

    res, err := s.Legal.GetLegalPracticearea(ctx, operations.GetLegalPracticeareaRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalPracticearea != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetLegalPracticeareaRequest](../../pkg/models/operations/getlegalpracticearearequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetLegalPracticeareaResponse](../../pkg/models/operations/getlegalpracticearearesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetLegalTimeentry

Retrieve a timeentry

### Example Usage

<!-- UsageSnippet language="go" operationID="getLegalTimeentry" method="get" path="/legal/{connection_id}/timeentry/{id}" -->
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

    res, err := s.Legal.GetLegalTimeentry(ctx, operations.GetLegalTimeentryRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalTimeentry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetLegalTimeentryRequest](../../pkg/models/operations/getlegaltimeentryrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetLegalTimeentryResponse](../../pkg/models/operations/getlegaltimeentryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListLegalExpenses

List all expenses

### Example Usage

<!-- UsageSnippet language="go" operationID="listLegalExpenses" method="get" path="/legal/{connection_id}/expense" -->
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

    res, err := s.Legal.ListLegalExpenses(ctx, operations.ListLegalExpensesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalExpenses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListLegalExpensesRequest](../../pkg/models/operations/listlegalexpensesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListLegalExpensesResponse](../../pkg/models/operations/listlegalexpensesresponse.md), error**

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

    res, err := s.Legal.ListLegalMatters(ctx, operations.ListLegalMattersRequest{
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

## ListLegalPracticeareas

List all practiceareas

### Example Usage

<!-- UsageSnippet language="go" operationID="listLegalPracticeareas" method="get" path="/legal/{connection_id}/practicearea" -->
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

    res, err := s.Legal.ListLegalPracticeareas(ctx, operations.ListLegalPracticeareasRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalPracticeareas != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListLegalPracticeareasRequest](../../pkg/models/operations/listlegalpracticeareasrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.ListLegalPracticeareasResponse](../../pkg/models/operations/listlegalpracticeareasresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListLegalTimeentries

List all timeentries

### Example Usage

<!-- UsageSnippet language="go" operationID="listLegalTimeentries" method="get" path="/legal/{connection_id}/timeentry" -->
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

    res, err := s.Legal.ListLegalTimeentries(ctx, operations.ListLegalTimeentriesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalTimeentries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListLegalTimeentriesRequest](../../pkg/models/operations/listlegaltimeentriesrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListLegalTimeentriesResponse](../../pkg/models/operations/listlegaltimeentriesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchLegalExpense

Update an expense

### Example Usage

<!-- UsageSnippet language="go" operationID="patchLegalExpense" method="patch" path="/legal/{connection_id}/expense/{id}" example="legal_expense" -->
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

    res, err := s.Legal.PatchLegalExpense(ctx, operations.PatchLegalExpenseRequest{
        LegalExpense: shared.LegalExpense{
            Category: unifiedgosdk.Pointer("Kids"),
            CreatedAt: types.MustNewTimeFromString("2020-11-08T00:41:42.356Z"),
            Currency: unifiedgosdk.Pointer("MVR"),
            Description: unifiedgosdk.Pointer("Temporibus agnitio sui."),
            ID: unifiedgosdk.Pointer("a1d1574a-5f4f-401d-8178-fff5cfce7fc4"),
            IncurredAt: types.MustNewTimeFromString("2020-11-08T00:41:42.356Z"),
            IsBilled: unifiedgosdk.Pointer(false),
            Metadata: []shared.LegalMetadata{},
            PriceAmount: unifiedgosdk.Pointer[float64](100.0),
            Quantity: unifiedgosdk.Pointer[float64](3.0),
            TotalAmount: unifiedgosdk.Pointer[float64](300.0),
            Type: shared.LegalExpenseTypeOther.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-03-25T18:37:19.762Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalExpense != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PatchLegalExpenseRequest](../../pkg/models/operations/patchlegalexpenserequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PatchLegalExpenseResponse](../../pkg/models/operations/patchlegalexpenseresponse.md), error**

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

    res, err := s.Legal.PatchLegalMatter(ctx, operations.PatchLegalMatterRequest{
        LegalMatter: shared.LegalMatter{
            BillingMethod: shared.BillingMethodOther.ToPointer(),
            ClientReference: unifiedgosdk.Pointer("qSrjIBAD"),
            CompanyIds: []string{},
            ContactIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2023-05-28T04:10:42.375Z"),
            Currency: unifiedgosdk.Pointer("BOB"),
            Description: unifiedgosdk.Pointer("Speciosus cresco aut cotidie temeritas arma."),
            ID: unifiedgosdk.Pointer("05c813d4-bdff-44e0-9693-71fc527aec71"),
            Metadata: []shared.LegalMetadata{},
            Name: unifiedgosdk.Pointer("Walter v. Franecki - Spinka"),
            Number: unifiedgosdk.Pointer("2023-3062"),
            OpenedAt: types.MustNewTimeFromString("2023-07-04T23:36:07.863Z"),
            Parties: []shared.LegalParty{
                shared.LegalParty{
                    Description: unifiedgosdk.Pointer("Urban Walter"),
                    Role: shared.LegalPartyRoleExpert.ToPointer(),
                },
            },
            Stage: unifiedgosdk.Pointer("currus"),
            Status: shared.LegalMatterStatusOpen.ToPointer(),
            StatuteOfLimitationsAt: types.MustNewTimeFromString("2027-06-19T18:26:16.524Z"),
            UpdatedAt: types.MustNewTimeFromString("2023-07-17T01:14:18.368Z"),
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

## PatchLegalPracticearea

Update a practicearea

### Example Usage

<!-- UsageSnippet language="go" operationID="patchLegalPracticearea" method="patch" path="/legal/{connection_id}/practicearea/{id}" example="legal_practicearea" -->
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

    res, err := s.Legal.PatchLegalPracticearea(ctx, operations.PatchLegalPracticeareaRequest{
        LegalPracticearea: shared.LegalPracticearea{
            Code: unifiedgosdk.Pointer("QCA"),
            CreatedAt: types.MustNewTimeFromString("2023-01-28T21:00:26.264Z"),
            ID: unifiedgosdk.Pointer("1e0e3aec-009a-4fa7-99f0-2699eebf31df"),
            Metadata: []shared.LegalMetadata{},
            Name: unifiedgosdk.Pointer("witty Law"),
            Type: shared.LegalPracticeareaTypeEducation.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-08-21T09:46:07.846Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalPracticearea != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PatchLegalPracticeareaRequest](../../pkg/models/operations/patchlegalpracticearearequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.PatchLegalPracticeareaResponse](../../pkg/models/operations/patchlegalpracticearearesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchLegalTimeentry

Update a timeentry

### Example Usage

<!-- UsageSnippet language="go" operationID="patchLegalTimeentry" method="patch" path="/legal/{connection_id}/timeentry/{id}" example="legal_timeentry" -->
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

    res, err := s.Legal.PatchLegalTimeentry(ctx, operations.PatchLegalTimeentryRequest{
        LegalTimeentry: shared.LegalTimeentry{
            BillableMinutes: unifiedgosdk.Pointer[float64](390.0),
            CreatedAt: types.MustNewTimeFromString("2022-05-06T03:26:03.756Z"),
            Currency: unifiedgosdk.Pointer("ZAR"),
            Description: unifiedgosdk.Pointer("Ea derideo adulatio torrens."),
            ID: unifiedgosdk.Pointer("7bcbf3d7-e53b-41d8-a165-f7cc3c7a73f3"),
            IsBillable: unifiedgosdk.Pointer(true),
            Metadata: []shared.LegalMetadata{},
            Minutes: unifiedgosdk.Pointer[float64](390.0),
            RateAmount: unifiedgosdk.Pointer[float64](637.0),
            TotalAmount: unifiedgosdk.Pointer[float64](4140.5),
            UpdatedAt: types.MustNewTimeFromString("2022-06-29T18:26:46.381Z"),
            WorkedAt: types.MustNewTimeFromString("2022-05-06T03:26:03.756Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalTimeentry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PatchLegalTimeentryRequest](../../pkg/models/operations/patchlegaltimeentryrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.PatchLegalTimeentryResponse](../../pkg/models/operations/patchlegaltimeentryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveLegalExpense

Remove an expense

### Example Usage

<!-- UsageSnippet language="go" operationID="removeLegalExpense" method="delete" path="/legal/{connection_id}/expense/{id}" -->
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

    res, err := s.Legal.RemoveLegalExpense(ctx, operations.RemoveLegalExpenseRequest{
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.RemoveLegalExpenseRequest](../../pkg/models/operations/removelegalexpenserequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.RemoveLegalExpenseResponse](../../pkg/models/operations/removelegalexpenseresponse.md), error**

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

    res, err := s.Legal.RemoveLegalMatter(ctx, operations.RemoveLegalMatterRequest{
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

## RemoveLegalPracticearea

Remove a practicearea

### Example Usage

<!-- UsageSnippet language="go" operationID="removeLegalPracticearea" method="delete" path="/legal/{connection_id}/practicearea/{id}" -->
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

    res, err := s.Legal.RemoveLegalPracticearea(ctx, operations.RemoveLegalPracticeareaRequest{
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
| `request`                                                                                                  | [operations.RemoveLegalPracticeareaRequest](../../pkg/models/operations/removelegalpracticearearequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.RemoveLegalPracticeareaResponse](../../pkg/models/operations/removelegalpracticearearesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveLegalTimeentry

Remove a timeentry

### Example Usage

<!-- UsageSnippet language="go" operationID="removeLegalTimeentry" method="delete" path="/legal/{connection_id}/timeentry/{id}" -->
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

    res, err := s.Legal.RemoveLegalTimeentry(ctx, operations.RemoveLegalTimeentryRequest{
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.RemoveLegalTimeentryRequest](../../pkg/models/operations/removelegaltimeentryrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.RemoveLegalTimeentryResponse](../../pkg/models/operations/removelegaltimeentryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateLegalExpense

Update an expense

### Example Usage

<!-- UsageSnippet language="go" operationID="updateLegalExpense" method="put" path="/legal/{connection_id}/expense/{id}" example="legal_expense" -->
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

    res, err := s.Legal.UpdateLegalExpense(ctx, operations.UpdateLegalExpenseRequest{
        LegalExpense: shared.LegalExpense{
            Category: unifiedgosdk.Pointer("Kids"),
            CreatedAt: types.MustNewTimeFromString("2020-11-08T00:41:42.356Z"),
            Currency: unifiedgosdk.Pointer("MVR"),
            Description: unifiedgosdk.Pointer("Temporibus agnitio sui."),
            ID: unifiedgosdk.Pointer("a1d1574a-5f4f-401d-8178-fff5cfce7fc4"),
            IncurredAt: types.MustNewTimeFromString("2020-11-08T00:41:42.356Z"),
            IsBilled: unifiedgosdk.Pointer(false),
            Metadata: []shared.LegalMetadata{},
            PriceAmount: unifiedgosdk.Pointer[float64](100.0),
            Quantity: unifiedgosdk.Pointer[float64](3.0),
            TotalAmount: unifiedgosdk.Pointer[float64](300.0),
            Type: shared.LegalExpenseTypeOther.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-03-25T18:37:19.762Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalExpense != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateLegalExpenseRequest](../../pkg/models/operations/updatelegalexpenserequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.UpdateLegalExpenseResponse](../../pkg/models/operations/updatelegalexpenseresponse.md), error**

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

    res, err := s.Legal.UpdateLegalMatter(ctx, operations.UpdateLegalMatterRequest{
        LegalMatter: shared.LegalMatter{
            BillingMethod: shared.BillingMethodOther.ToPointer(),
            ClientReference: unifiedgosdk.Pointer("qSrjIBAD"),
            CompanyIds: []string{},
            ContactIds: []string{},
            CreatedAt: types.MustNewTimeFromString("2023-05-28T04:10:42.375Z"),
            Currency: unifiedgosdk.Pointer("BOB"),
            Description: unifiedgosdk.Pointer("Speciosus cresco aut cotidie temeritas arma."),
            ID: unifiedgosdk.Pointer("05c813d4-bdff-44e0-9693-71fc527aec71"),
            Metadata: []shared.LegalMetadata{},
            Name: unifiedgosdk.Pointer("Walter v. Franecki - Spinka"),
            Number: unifiedgosdk.Pointer("2023-3062"),
            OpenedAt: types.MustNewTimeFromString("2023-07-04T23:36:07.863Z"),
            Parties: []shared.LegalParty{
                shared.LegalParty{
                    Description: unifiedgosdk.Pointer("Urban Walter"),
                    Role: shared.LegalPartyRoleExpert.ToPointer(),
                },
            },
            Stage: unifiedgosdk.Pointer("currus"),
            Status: shared.LegalMatterStatusOpen.ToPointer(),
            StatuteOfLimitationsAt: types.MustNewTimeFromString("2027-06-19T18:26:16.524Z"),
            UpdatedAt: types.MustNewTimeFromString("2023-07-17T01:14:18.368Z"),
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

## UpdateLegalPracticearea

Update a practicearea

### Example Usage

<!-- UsageSnippet language="go" operationID="updateLegalPracticearea" method="put" path="/legal/{connection_id}/practicearea/{id}" example="legal_practicearea" -->
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

    res, err := s.Legal.UpdateLegalPracticearea(ctx, operations.UpdateLegalPracticeareaRequest{
        LegalPracticearea: shared.LegalPracticearea{
            Code: unifiedgosdk.Pointer("QCA"),
            CreatedAt: types.MustNewTimeFromString("2023-01-28T21:00:26.264Z"),
            ID: unifiedgosdk.Pointer("1e0e3aec-009a-4fa7-99f0-2699eebf31df"),
            Metadata: []shared.LegalMetadata{},
            Name: unifiedgosdk.Pointer("witty Law"),
            Type: shared.LegalPracticeareaTypeEducation.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-08-21T09:46:07.846Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalPracticearea != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.UpdateLegalPracticeareaRequest](../../pkg/models/operations/updatelegalpracticearearequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.UpdateLegalPracticeareaResponse](../../pkg/models/operations/updatelegalpracticearearesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateLegalTimeentry

Update a timeentry

### Example Usage

<!-- UsageSnippet language="go" operationID="updateLegalTimeentry" method="put" path="/legal/{connection_id}/timeentry/{id}" example="legal_timeentry" -->
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

    res, err := s.Legal.UpdateLegalTimeentry(ctx, operations.UpdateLegalTimeentryRequest{
        LegalTimeentry: shared.LegalTimeentry{
            BillableMinutes: unifiedgosdk.Pointer[float64](390.0),
            CreatedAt: types.MustNewTimeFromString("2022-05-06T03:26:03.756Z"),
            Currency: unifiedgosdk.Pointer("ZAR"),
            Description: unifiedgosdk.Pointer("Ea derideo adulatio torrens."),
            ID: unifiedgosdk.Pointer("7bcbf3d7-e53b-41d8-a165-f7cc3c7a73f3"),
            IsBillable: unifiedgosdk.Pointer(true),
            Metadata: []shared.LegalMetadata{},
            Minutes: unifiedgosdk.Pointer[float64](390.0),
            RateAmount: unifiedgosdk.Pointer[float64](637.0),
            TotalAmount: unifiedgosdk.Pointer[float64](4140.5),
            UpdatedAt: types.MustNewTimeFromString("2022-06-29T18:26:46.381Z"),
            WorkedAt: types.MustNewTimeFromString("2022-05-06T03:26:03.756Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LegalTimeentry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdateLegalTimeentryRequest](../../pkg/models/operations/updatelegaltimeentryrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.UpdateLegalTimeentryResponse](../../pkg/models/operations/updatelegaltimeentryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |