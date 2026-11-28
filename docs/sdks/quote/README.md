# Quote

## Overview

### Available Operations

* [CreateAccountingQuote](#createaccountingquote) - Create a quote
* [CreatePrintingQuote](#createprintingquote) - Create a quote
* [CreatePropertyQuote](#createpropertyquote) - Create a quote
* [GetAccountingQuote](#getaccountingquote) - Retrieve a quote
* [ListAccountingQuotes](#listaccountingquotes) - List all quotes
* [PatchAccountingQuote](#patchaccountingquote) - Update a quote
* [RemoveAccountingQuote](#removeaccountingquote) - Remove a quote
* [UpdateAccountingQuote](#updateaccountingquote) - Update a quote

## CreateAccountingQuote

Creates a AccountingQuote. Provide the unified AccountingQuote fields in the request body; they are mapped to the provider's API. Returns the created AccountingQuote.

### Example Usage

<!-- UsageSnippet language="go" operationID="createAccountingQuote" method="post" path="/accounting/{connection_id}/quote" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Quote.CreateAccountingQuote(ctx, operations.CreateAccountingQuoteRequest{
        AccountingQuote: shared.AccountingQuote{},
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingQuote != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CreateAccountingQuoteRequest](../../pkg/models/operations/createaccountingquoterequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.CreateAccountingQuoteResponse](../../pkg/models/operations/createaccountingquoteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreatePrintingQuote

Creates a PrintingQuote. Provide the unified PrintingQuote fields in the request body; they are mapped to the provider's API. Returns the created PrintingQuote.

### Example Usage

<!-- UsageSnippet language="go" operationID="createPrintingQuote" method="post" path="/printing/{connection_id}/quote" example="printing_quote" -->
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

    res, err := s.Quote.CreatePrintingQuote(ctx, operations.CreatePrintingQuoteRequest{
        PrintingQuote: shared.PrintingQuote{
            CreatedAt: types.MustNewTimeFromString("2019-01-02T03:42:06.944Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            ID: unifiedgosdk.Pointer("b837c070-dba3-4031-aee8-91fcdca1a61c"),
            Lineitems: []shared.PrintingLineitem{
                shared.PrintingLineitem{
                    Files: []shared.PrintingLineitemFile{
                        shared.PrintingLineitemFile{
                            Placement: unifiedgosdk.Pointer("default"),
                            URL: unifiedgosdk.Pointer("https://picsum.photos/seed/0Y5Bjnoiwx/3057/3113"),
                        },
                    },
                    ID: unifiedgosdk.Pointer("1"),
                    Name: unifiedgosdk.Pointer("Practical Wooden Salad"),
                    Quantity: unifiedgosdk.Pointer[float64](2.0),
                    RetailPriceAmount: unifiedgosdk.Pointer[float64](50.83),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](21.72),
                },
                shared.PrintingLineitem{
                    Files: []shared.PrintingLineitemFile{
                        shared.PrintingLineitemFile{
                            Placement: unifiedgosdk.Pointer("default"),
                            URL: unifiedgosdk.Pointer("https://picsum.photos/seed/rcPUF/2295/1801"),
                        },
                    },
                    ID: unifiedgosdk.Pointer("2"),
                    Name: unifiedgosdk.Pointer("Handmade Gold Car"),
                    Quantity: unifiedgosdk.Pointer[float64](4.0),
                    RetailPriceAmount: unifiedgosdk.Pointer[float64](35.47),
                    Sku: unifiedgosdk.Pointer("I79GIBCK"),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](17.64),
                },
            },
            Metadata: []shared.PrintingMetadata{},
            ShippingAddress: &shared.PropertyPrintingQuoteShippingAddress{
                Address1: unifiedgosdk.Pointer("4290 Garden Street"),
                Address2: unifiedgosdk.Pointer("Suite 874"),
                City: unifiedgosdk.Pointer("Cartwrightshire"),
                CompanyName: unifiedgosdk.Pointer("Kris and Sons"),
                Country: unifiedgosdk.Pointer("United States"),
                CountryCode: unifiedgosdk.Pointer("US"),
                Email: unifiedgosdk.Pointer("jada73@hotmail.com"),
                Name: unifiedgosdk.Pointer("Heaven Crist"),
                PostalCode: unifiedgosdk.Pointer("92871"),
                Region: unifiedgosdk.Pointer("North Carolina"),
                RegionCode: unifiedgosdk.Pointer("KY"),
                Telephone: unifiedgosdk.Pointer("653.412.8900 x9387"),
            },
            ShippingAmount: unifiedgosdk.Pointer[float64](4.69),
            ShippingMethod: unifiedgosdk.Pointer("STANDARD"),
            ShippingOptions: []shared.PrintingShippingoption{
                shared.PrintingShippingoption{
                    Amount: unifiedgosdk.Pointer[float64](4.69),
                    Code: unifiedgosdk.Pointer("STANDARD"),
                    Currency: unifiedgosdk.Pointer("USD"),
                    MaxDeliveryDays: unifiedgosdk.Pointer[float64](5.0),
                    MinDeliveryDays: unifiedgosdk.Pointer[float64](3.0),
                    Name: unifiedgosdk.Pointer("Flat Rate (3-5 business days after fulfillment)"),
                },
                shared.PrintingShippingoption{
                    Amount: unifiedgosdk.Pointer[float64](12.99),
                    Code: unifiedgosdk.Pointer("EXPRESS"),
                    Currency: unifiedgosdk.Pointer("USD"),
                    MaxDeliveryDays: unifiedgosdk.Pointer[float64](3.0),
                    MinDeliveryDays: unifiedgosdk.Pointer[float64](1.0),
                    Name: unifiedgosdk.Pointer("Express (1-3 business days after fulfillment)"),
                },
            },
            SubtotalAmount: unifiedgosdk.Pointer[float64](114.0),
            TaxAmount: unifiedgosdk.Pointer[float64](9.12),
            TotalAmount: unifiedgosdk.Pointer[float64](127.81),
            UpdatedAt: types.MustNewTimeFromString("2019-01-02T03:42:06.944Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrintingQuote != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreatePrintingQuoteRequest](../../pkg/models/operations/createprintingquoterequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreatePrintingQuoteResponse](../../pkg/models/operations/createprintingquoteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreatePropertyQuote

Creates a PropertyQuote. Provide the unified PropertyQuote fields in the request body; they are mapped to the provider's API. Returns the created PropertyQuote.

### Example Usage

<!-- UsageSnippet language="go" operationID="createPropertyQuote" method="post" path="/property/{connection_id}/quote" example="property_quote" -->
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

    res, err := s.Quote.CreatePropertyQuote(ctx, operations.CreatePropertyQuoteRequest{
        PropertyQuote: shared.PropertyQuote{
            Adults: unifiedgosdk.Pointer[float64](1.0),
            Children: unifiedgosdk.Pointer[float64](1.0),
            CreatedAt: types.MustNewTimeFromString("2020-12-18T19:14:11.415Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DepositAmount: unifiedgosdk.Pointer[float64](594.54),
            EndAt: types.MustNewTimeFromString("2026-12-12"),
            Guests: unifiedgosdk.Pointer[float64](2.0),
            ID: unifiedgosdk.Pointer("903419a1-9889-4de8-9146-59339ddb133d"),
            LineItems: []shared.PropertyLineitem{
                shared.PropertyLineitem{
                    Amount: unifiedgosdk.Pointer[float64](1665.0),
                    Description: unifiedgosdk.Pointer("9 nights"),
                    IsMandatory: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Accommodation"),
                    Type: shared.PropertyLineitemTypeBase.ToPointer(),
                },
                shared.PropertyLineitem{
                    Amount: unifiedgosdk.Pointer[float64](117.0),
                    IsMandatory: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Cleaning fee"),
                    Type: shared.PropertyLineitemTypeFee.ToPointer(),
                },
                shared.PropertyLineitem{
                    Amount: unifiedgosdk.Pointer[float64](199.8),
                    IsMandatory: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Lodging tax"),
                    Type: shared.PropertyLineitemTypeTax.ToPointer(),
                },
            },
            StartAt: types.MustNewTimeFromString("2026-12-03"),
            TotalAmount: unifiedgosdk.Pointer[float64](1981.8),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PropertyQuote != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreatePropertyQuoteRequest](../../pkg/models/operations/createpropertyquoterequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreatePropertyQuoteResponse](../../pkg/models/operations/createpropertyquoteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAccountingQuote

Retrieves a single AccountingQuote by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAccountingQuote" method="get" path="/accounting/{connection_id}/quote/{id}" -->
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

    res, err := s.Quote.GetAccountingQuote(ctx, operations.GetAccountingQuoteRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingQuote != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetAccountingQuoteRequest](../../pkg/models/operations/getaccountingquoterequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetAccountingQuoteResponse](../../pkg/models/operations/getaccountingquoteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAccountingQuotes

Returns a list of AccountingQuotes. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listAccountingQuotes" method="get" path="/accounting/{connection_id}/quote" -->
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

    res, err := s.Quote.ListAccountingQuotes(ctx, operations.ListAccountingQuotesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingQuotes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListAccountingQuotesRequest](../../pkg/models/operations/listaccountingquotesrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListAccountingQuotesResponse](../../pkg/models/operations/listaccountingquotesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAccountingQuote

Updates the AccountingQuote identified by `id`. Provide the unified AccountingQuote fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAccountingQuote" method="patch" path="/accounting/{connection_id}/quote/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Quote.PatchAccountingQuote(ctx, operations.PatchAccountingQuoteRequest{
        AccountingQuote: shared.AccountingQuote{},
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingQuote != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PatchAccountingQuoteRequest](../../pkg/models/operations/patchaccountingquoterequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.PatchAccountingQuoteResponse](../../pkg/models/operations/patchaccountingquoteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAccountingQuote

Deletes the AccountingQuote identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAccountingQuote" method="delete" path="/accounting/{connection_id}/quote/{id}" -->
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

    res, err := s.Quote.RemoveAccountingQuote(ctx, operations.RemoveAccountingQuoteRequest{
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.RemoveAccountingQuoteRequest](../../pkg/models/operations/removeaccountingquoterequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.RemoveAccountingQuoteResponse](../../pkg/models/operations/removeaccountingquoteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAccountingQuote

Updates the AccountingQuote identified by `id`. Provide the unified AccountingQuote fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAccountingQuote" method="put" path="/accounting/{connection_id}/quote/{id}" -->
```go
package main

import(
	"context"
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Quote.UpdateAccountingQuote(ctx, operations.UpdateAccountingQuoteRequest{
        AccountingQuote: shared.AccountingQuote{},
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountingQuote != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdateAccountingQuoteRequest](../../pkg/models/operations/updateaccountingquoterequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UpdateAccountingQuoteResponse](../../pkg/models/operations/updateaccountingquoteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |