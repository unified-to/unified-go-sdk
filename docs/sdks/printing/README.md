# Printing

## Overview

### Available Operations

* [CreatePrintingFile](#createprintingfile) - Create a file
* [CreatePrintingOrder](#createprintingorder) - Create an order
* [CreatePrintingQuote](#createprintingquote) - Create a quote
* [GetPrintingFile](#getprintingfile) - Retrieve a file
* [GetPrintingOrder](#getprintingorder) - Retrieve an order
* [GetPrintingProduct](#getprintingproduct) - Retrieve a product
* [GetPrintingTemplate](#getprintingtemplate) - Retrieve a template
* [ListPrintingFiles](#listprintingfiles) - List all files
* [ListPrintingOrders](#listprintingorders) - List all orders
* [ListPrintingProducts](#listprintingproducts) - List all products
* [ListPrintingTemplates](#listprintingtemplates) - List all templates
* [PatchPrintingOrder](#patchprintingorder) - Update an order
* [RemovePrintingOrder](#removeprintingorder) - Remove an order
* [UpdatePrintingOrder](#updateprintingorder) - Update an order

## CreatePrintingFile

Creates a PrintingFile. Provide the unified PrintingFile fields in the request body; they are mapped to the provider's API. Returns the created PrintingFile.

### Example Usage

<!-- UsageSnippet language="go" operationID="createPrintingFile" method="post" path="/printing/{connection_id}/file" example="printing_file" -->
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

    res, err := s.Printing.CreatePrintingFile(ctx, operations.CreatePrintingFileRequest{
        PrintingFile: shared.PrintingFile{
            CreatedAt: types.MustNewTimeFromString("2020-10-22T04:25:39.825Z"),
            Dpi: unifiedgosdk.Pointer[float64](300.0),
            Height: unifiedgosdk.Pointer[float64](5400.0),
            ID: unifiedgosdk.Pointer("c35be3e9-c3be-4478-b0df-b883fb76e24d"),
            Metadata: []shared.PrintingMetadata{},
            MimeType: unifiedgosdk.Pointer("image/jpeg"),
            Name: unifiedgosdk.Pointer("blond-ecliptic.jpg"),
            PreviewURL: unifiedgosdk.Pointer("https://picsum.photos/seed/aATvhmNA/400/400"),
            Size: unifiedgosdk.Pointer[float64](1956654.0),
            Status: shared.PrintingFileStatusReady.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2020-10-22T04:25:39.825Z"),
            URL: unifiedgosdk.Pointer("https://picsum.photos/seed/XBj97B/2783/172"),
            Width: unifiedgosdk.Pointer[float64](4500.0),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrintingFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreatePrintingFileRequest](../../pkg/models/operations/createprintingfilerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreatePrintingFileResponse](../../pkg/models/operations/createprintingfileresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreatePrintingOrder

Creates a PrintingOrder. Provide the unified PrintingOrder fields in the request body; they are mapped to the provider's API. Returns the created PrintingOrder.

### Example Usage

<!-- UsageSnippet language="go" operationID="createPrintingOrder" method="post" path="/printing/{connection_id}/order" example="printing_order" -->
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

    res, err := s.Printing.CreatePrintingOrder(ctx, operations.CreatePrintingOrderRequest{
        PrintingOrder: shared.PrintingOrder{
            CreatedAt: types.MustNewTimeFromString("2022-03-14T11:16:15.147Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            ID: unifiedgosdk.Pointer("9d6516d5-784b-4442-9751-127e629dba39"),
            Lineitems: []shared.PrintingLineitem{
                shared.PrintingLineitem{
                    Files: []shared.PrintingLineitemFile{
                        shared.PrintingLineitemFile{
                            Placement: unifiedgosdk.Pointer("default"),
                            URL: unifiedgosdk.Pointer("https://picsum.photos/seed/y54e6o9hYi/191/3357"),
                        },
                    },
                    ID: unifiedgosdk.Pointer("1"),
                    Name: unifiedgosdk.Pointer("Bespoke Marble Chair"),
                    Quantity: unifiedgosdk.Pointer[float64](2.0),
                    RetailPriceAmount: unifiedgosdk.Pointer[float64](54.89),
                    Sku: unifiedgosdk.Pointer("IBDIE3XR"),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](26.37),
                },
                shared.PrintingLineitem{
                    Files: []shared.PrintingLineitemFile{
                        shared.PrintingLineitemFile{
                            Placement: unifiedgosdk.Pointer("default"),
                            URL: unifiedgosdk.Pointer("https://picsum.photos/seed/T7S5xf/3521/1221"),
                        },
                    },
                    ID: unifiedgosdk.Pointer("2"),
                    Name: unifiedgosdk.Pointer("Electronic Cotton Hat"),
                    Quantity: unifiedgosdk.Pointer[float64](1.0),
                    RetailPriceAmount: unifiedgosdk.Pointer[float64](23.94),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](10.04),
                },
            },
            Metadata: []shared.PrintingMetadata{},
            Shipments: []shared.PrintingShipment{},
            ShippingAddress: &shared.PropertyPrintingOrderShippingAddress{
                Address1: unifiedgosdk.Pointer("651 Main Street"),
                City: unifiedgosdk.Pointer("Euclid"),
                Country: unifiedgosdk.Pointer("United States"),
                CountryCode: unifiedgosdk.Pointer("US"),
                Email: unifiedgosdk.Pointer("gary_leffler28@gmail.com"),
                Name: unifiedgosdk.Pointer("Angelita Denesik"),
                PostalCode: unifiedgosdk.Pointer("80301"),
                Region: unifiedgosdk.Pointer("Colorado"),
                RegionCode: unifiedgosdk.Pointer("AR"),
                Telephone: unifiedgosdk.Pointer("808-831-6489 x62979"),
            },
            ShippingAmount: unifiedgosdk.Pointer[float64](4.69),
            ShippingMethod: unifiedgosdk.Pointer("STANDARD"),
            Status: shared.PrintingOrderStatusOther.ToPointer(),
            SubtotalAmount: unifiedgosdk.Pointer[float64](62.78),
            TaxAmount: unifiedgosdk.Pointer[float64](5.02),
            TotalAmount: unifiedgosdk.Pointer[float64](72.49),
            UpdatedAt: types.MustNewTimeFromString("2025-01-23T00:20:58.017Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrintingOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreatePrintingOrderRequest](../../pkg/models/operations/createprintingorderrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreatePrintingOrderResponse](../../pkg/models/operations/createprintingorderresponse.md), error**

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

    res, err := s.Printing.CreatePrintingQuote(ctx, operations.CreatePrintingQuoteRequest{
        PrintingQuote: shared.PrintingQuote{
            CreatedAt: types.MustNewTimeFromString("2019-01-02T03:42:06.944Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            ID: unifiedgosdk.Pointer("c0ca1481-d794-43b1-a1f3-80d0a7364419"),
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

## GetPrintingFile

Retrieves a single PrintingFile by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPrintingFile" method="get" path="/printing/{connection_id}/file/{id}" -->
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

    res, err := s.Printing.GetPrintingFile(ctx, operations.GetPrintingFileRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrintingFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetPrintingFileRequest](../../pkg/models/operations/getprintingfilerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetPrintingFileResponse](../../pkg/models/operations/getprintingfileresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPrintingOrder

Retrieves a single PrintingOrder by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPrintingOrder" method="get" path="/printing/{connection_id}/order/{id}" -->
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

    res, err := s.Printing.GetPrintingOrder(ctx, operations.GetPrintingOrderRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrintingOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetPrintingOrderRequest](../../pkg/models/operations/getprintingorderrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetPrintingOrderResponse](../../pkg/models/operations/getprintingorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPrintingProduct

Retrieves a single PrintingProduct by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPrintingProduct" method="get" path="/printing/{connection_id}/product/{id}" -->
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

    res, err := s.Printing.GetPrintingProduct(ctx, operations.GetPrintingProductRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrintingProduct != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetPrintingProductRequest](../../pkg/models/operations/getprintingproductrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetPrintingProductResponse](../../pkg/models/operations/getprintingproductresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetPrintingTemplate

Retrieves a single PrintingTemplate by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPrintingTemplate" method="get" path="/printing/{connection_id}/template/{id}" -->
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

    res, err := s.Printing.GetPrintingTemplate(ctx, operations.GetPrintingTemplateRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrintingTemplate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetPrintingTemplateRequest](../../pkg/models/operations/getprintingtemplaterequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.GetPrintingTemplateResponse](../../pkg/models/operations/getprintingtemplateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPrintingFiles

Returns a list of PrintingFiles. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPrintingFiles" method="get" path="/printing/{connection_id}/file" -->
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

    res, err := s.Printing.ListPrintingFiles(ctx, operations.ListPrintingFilesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrintingFiles != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListPrintingFilesRequest](../../pkg/models/operations/listprintingfilesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListPrintingFilesResponse](../../pkg/models/operations/listprintingfilesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPrintingOrders

Returns a list of PrintingOrders. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPrintingOrders" method="get" path="/printing/{connection_id}/order" -->
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

    res, err := s.Printing.ListPrintingOrders(ctx, operations.ListPrintingOrdersRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrintingOrders != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListPrintingOrdersRequest](../../pkg/models/operations/listprintingordersrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListPrintingOrdersResponse](../../pkg/models/operations/listprintingordersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPrintingProducts

Returns a list of PrintingProducts. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPrintingProducts" method="get" path="/printing/{connection_id}/product" -->
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

    res, err := s.Printing.ListPrintingProducts(ctx, operations.ListPrintingProductsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrintingProducts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListPrintingProductsRequest](../../pkg/models/operations/listprintingproductsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListPrintingProductsResponse](../../pkg/models/operations/listprintingproductsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPrintingTemplates

Returns a list of PrintingTemplates. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPrintingTemplates" method="get" path="/printing/{connection_id}/template" -->
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

    res, err := s.Printing.ListPrintingTemplates(ctx, operations.ListPrintingTemplatesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrintingTemplates != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListPrintingTemplatesRequest](../../pkg/models/operations/listprintingtemplatesrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.ListPrintingTemplatesResponse](../../pkg/models/operations/listprintingtemplatesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchPrintingOrder

Updates the PrintingOrder identified by `id`. Provide the unified PrintingOrder fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchPrintingOrder" method="patch" path="/printing/{connection_id}/order/{id}" example="printing_order" -->
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

    res, err := s.Printing.PatchPrintingOrder(ctx, operations.PatchPrintingOrderRequest{
        PrintingOrder: shared.PrintingOrder{
            CreatedAt: types.MustNewTimeFromString("2022-03-14T11:16:15.147Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            ID: unifiedgosdk.Pointer("af220e4d-5540-45c4-bb37-aeb092251e06"),
            Lineitems: []shared.PrintingLineitem{
                shared.PrintingLineitem{
                    Files: []shared.PrintingLineitemFile{
                        shared.PrintingLineitemFile{
                            Placement: unifiedgosdk.Pointer("default"),
                            URL: unifiedgosdk.Pointer("https://picsum.photos/seed/y54e6o9hYi/191/3357"),
                        },
                    },
                    ID: unifiedgosdk.Pointer("1"),
                    Name: unifiedgosdk.Pointer("Bespoke Marble Chair"),
                    Quantity: unifiedgosdk.Pointer[float64](2.0),
                    RetailPriceAmount: unifiedgosdk.Pointer[float64](54.89),
                    Sku: unifiedgosdk.Pointer("IBDIE3XR"),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](26.37),
                },
                shared.PrintingLineitem{
                    Files: []shared.PrintingLineitemFile{
                        shared.PrintingLineitemFile{
                            Placement: unifiedgosdk.Pointer("default"),
                            URL: unifiedgosdk.Pointer("https://picsum.photos/seed/T7S5xf/3521/1221"),
                        },
                    },
                    ID: unifiedgosdk.Pointer("2"),
                    Name: unifiedgosdk.Pointer("Electronic Cotton Hat"),
                    Quantity: unifiedgosdk.Pointer[float64](1.0),
                    RetailPriceAmount: unifiedgosdk.Pointer[float64](23.94),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](10.04),
                },
            },
            Metadata: []shared.PrintingMetadata{},
            Shipments: []shared.PrintingShipment{},
            ShippingAddress: &shared.PropertyPrintingOrderShippingAddress{
                Address1: unifiedgosdk.Pointer("651 Main Street"),
                City: unifiedgosdk.Pointer("Euclid"),
                Country: unifiedgosdk.Pointer("United States"),
                CountryCode: unifiedgosdk.Pointer("US"),
                Email: unifiedgosdk.Pointer("gary_leffler28@gmail.com"),
                Name: unifiedgosdk.Pointer("Angelita Denesik"),
                PostalCode: unifiedgosdk.Pointer("80301"),
                Region: unifiedgosdk.Pointer("Colorado"),
                RegionCode: unifiedgosdk.Pointer("AR"),
                Telephone: unifiedgosdk.Pointer("808-831-6489 x62979"),
            },
            ShippingAmount: unifiedgosdk.Pointer[float64](4.69),
            ShippingMethod: unifiedgosdk.Pointer("STANDARD"),
            Status: shared.PrintingOrderStatusOther.ToPointer(),
            SubtotalAmount: unifiedgosdk.Pointer[float64](62.78),
            TaxAmount: unifiedgosdk.Pointer[float64](5.02),
            TotalAmount: unifiedgosdk.Pointer[float64](72.49),
            UpdatedAt: types.MustNewTimeFromString("2025-01-23T00:20:58.037Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrintingOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PatchPrintingOrderRequest](../../pkg/models/operations/patchprintingorderrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.PatchPrintingOrderResponse](../../pkg/models/operations/patchprintingorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemovePrintingOrder

Deletes the PrintingOrder identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removePrintingOrder" method="delete" path="/printing/{connection_id}/order/{id}" -->
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

    res, err := s.Printing.RemovePrintingOrder(ctx, operations.RemovePrintingOrderRequest{
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.RemovePrintingOrderRequest](../../pkg/models/operations/removeprintingorderrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.RemovePrintingOrderResponse](../../pkg/models/operations/removeprintingorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdatePrintingOrder

Updates the PrintingOrder identified by `id`. Provide the unified PrintingOrder fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updatePrintingOrder" method="put" path="/printing/{connection_id}/order/{id}" example="printing_order" -->
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

    res, err := s.Printing.UpdatePrintingOrder(ctx, operations.UpdatePrintingOrderRequest{
        PrintingOrder: shared.PrintingOrder{
            CreatedAt: types.MustNewTimeFromString("2022-03-14T11:16:15.147Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DiscountAmount: unifiedgosdk.Pointer[float64](0.0),
            ID: unifiedgosdk.Pointer("af220e4d-5540-45c4-bb37-aeb092251e06"),
            Lineitems: []shared.PrintingLineitem{
                shared.PrintingLineitem{
                    Files: []shared.PrintingLineitemFile{
                        shared.PrintingLineitemFile{
                            Placement: unifiedgosdk.Pointer("default"),
                            URL: unifiedgosdk.Pointer("https://picsum.photos/seed/y54e6o9hYi/191/3357"),
                        },
                    },
                    ID: unifiedgosdk.Pointer("1"),
                    Name: unifiedgosdk.Pointer("Bespoke Marble Chair"),
                    Quantity: unifiedgosdk.Pointer[float64](2.0),
                    RetailPriceAmount: unifiedgosdk.Pointer[float64](54.89),
                    Sku: unifiedgosdk.Pointer("IBDIE3XR"),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](26.37),
                },
                shared.PrintingLineitem{
                    Files: []shared.PrintingLineitemFile{
                        shared.PrintingLineitemFile{
                            Placement: unifiedgosdk.Pointer("default"),
                            URL: unifiedgosdk.Pointer("https://picsum.photos/seed/T7S5xf/3521/1221"),
                        },
                    },
                    ID: unifiedgosdk.Pointer("2"),
                    Name: unifiedgosdk.Pointer("Electronic Cotton Hat"),
                    Quantity: unifiedgosdk.Pointer[float64](1.0),
                    RetailPriceAmount: unifiedgosdk.Pointer[float64](23.94),
                    UnitPriceAmount: unifiedgosdk.Pointer[float64](10.04),
                },
            },
            Metadata: []shared.PrintingMetadata{},
            Shipments: []shared.PrintingShipment{},
            ShippingAddress: &shared.PropertyPrintingOrderShippingAddress{
                Address1: unifiedgosdk.Pointer("651 Main Street"),
                City: unifiedgosdk.Pointer("Euclid"),
                Country: unifiedgosdk.Pointer("United States"),
                CountryCode: unifiedgosdk.Pointer("US"),
                Email: unifiedgosdk.Pointer("gary_leffler28@gmail.com"),
                Name: unifiedgosdk.Pointer("Angelita Denesik"),
                PostalCode: unifiedgosdk.Pointer("80301"),
                Region: unifiedgosdk.Pointer("Colorado"),
                RegionCode: unifiedgosdk.Pointer("AR"),
                Telephone: unifiedgosdk.Pointer("808-831-6489 x62979"),
            },
            ShippingAmount: unifiedgosdk.Pointer[float64](4.69),
            ShippingMethod: unifiedgosdk.Pointer("STANDARD"),
            Status: shared.PrintingOrderStatusOther.ToPointer(),
            SubtotalAmount: unifiedgosdk.Pointer[float64](62.78),
            TaxAmount: unifiedgosdk.Pointer[float64](5.02),
            TotalAmount: unifiedgosdk.Pointer[float64](72.49),
            UpdatedAt: types.MustNewTimeFromString("2025-01-23T00:20:58.037Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PrintingOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdatePrintingOrderRequest](../../pkg/models/operations/updateprintingorderrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdatePrintingOrderResponse](../../pkg/models/operations/updateprintingorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |