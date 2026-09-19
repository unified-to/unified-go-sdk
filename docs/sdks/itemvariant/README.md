# Itemvariant

## Overview

### Available Operations

* [CreateCommerceItemvariant](#createcommerceitemvariant) - Create an itemvariant
* [GetCommerceItemvariant](#getcommerceitemvariant) - Retrieve an itemvariant
* [ListCommerceItemvariants](#listcommerceitemvariants) - List all itemvariants
* [PatchCommerceItemvariant](#patchcommerceitemvariant) - Update an itemvariant
* [RemoveCommerceItemvariant](#removecommerceitemvariant) - Remove an itemvariant
* [UpdateCommerceItemvariant](#updatecommerceitemvariant) - Update an itemvariant

## CreateCommerceItemvariant

Create an itemvariant

### Example Usage

<!-- UsageSnippet language="go" operationID="createCommerceItemvariant" method="post" path="/commerce/{connection_id}/itemvariant" example="commerce_itemvariant" -->
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

    res, err := s.Itemvariant.CreateCommerceItemvariant(ctx, operations.CreateCommerceItemvariantRequest{
        CommerceItemvariant: shared.CommerceItemvariant{
            AvailableAt: types.MustNewTimeFromString("2022-02-02T16:10:33.503Z"),
            CreatedAt: types.MustNewTimeFromString("2022-01-20T13:49:12.968Z"),
            Description: unifiedgosdk.Pointer("Featuring Helium-enhanced technology, our Chips offers unparalleled helpful performance"),
            Height: unifiedgosdk.Pointer[float64](52.0),
            ID: unifiedgosdk.Pointer("618b9cb5-8c37-4b36-9b2b-e045d8975715"),
            IsActive: unifiedgosdk.Pointer(false),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsVisible: unifiedgosdk.Pointer(false),
            Length: unifiedgosdk.Pointer[float64](94.0),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Calcar delibero cursim summisse."),
                    Height: unifiedgosdk.Pointer[float64](394.0),
                    ID: unifiedgosdk.Pointer("ae4a1d86-3fb7-4ce8-8e90-c39ebe5a2530"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("3f9e367f-1476-4682-b8cf-83443f762dd0"),
                            Slug: unifiedgosdk.Pointer("illo"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "quia",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](92.0),
                    Type: shared.CommerceItemMediaTypeImage.ToPointer(),
                    URL: "https://picsum.photos/seed/u0YdHqlRu/2007/3208",
                    Width: unifiedgosdk.Pointer[float64](54.0),
                },
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Civitas acies substantia tergo."),
                    Height: unifiedgosdk.Pointer[float64](351.0),
                    ID: unifiedgosdk.Pointer("64dc46c4-476e-4488-8824-df1948249132"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("ddd56205-a20a-4da8-ac42-3e91d0bb516b"),
                            Slug: unifiedgosdk.Pointer("libero"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "capitulus",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](44.0),
                    Type: shared.CommerceItemMediaTypeImage.ToPointer(),
                    URL: "https://loremflickr.com/2230/1237?lock=8628070842159966",
                    Width: unifiedgosdk.Pointer[float64](55.0),
                },
            },
            Metadata: []shared.CommerceMetadata{
                shared.CommerceMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCommerceMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CommerceMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("9afdbab5-7438-4811-b9b0-6a4cb735e369"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                        "nihil",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Keyboard"),
            Options: []shared.CommerceItemOption{
                shared.CommerceItemOption{
                    ID: unifiedgosdk.Pointer("a9481e93-92ce-4a86-8d75-6ea96db69c6a"),
                    Name: "Steel",
                    Position: unifiedgosdk.Pointer[float64](97.0),
                    Values: []string{
                        "Granite",
                        "Plastic",
                    },
                },
            },
            Prices: []shared.CommerceItemPrice{
                shared.CommerceItemPrice{
                    CompareAtPrice: unifiedgosdk.Pointer[float64](3745.0),
                    Currency: unifiedgosdk.Pointer("COP"),
                    Price: 4913.0,
                },
                shared.CommerceItemPrice{
                    CompareAtPrice: unifiedgosdk.Pointer[float64](438.0),
                    Currency: unifiedgosdk.Pointer("PHP"),
                    Price: 1378.0,
                },
                shared.CommerceItemPrice{
                    CompareAtPrice: unifiedgosdk.Pointer[float64](1614.0),
                    Currency: unifiedgosdk.Pointer("PHP"),
                    Price: 8702.0,
                },
            },
            PublicDescription: unifiedgosdk.Pointer("Stylish Soap designed to make you stand out with insistent looks"),
            PublicName: unifiedgosdk.Pointer("Keyboard"),
            RequiresShipping: unifiedgosdk.Pointer(false),
            SizeUnit: shared.SizeUnitCm.ToPointer(),
            Sku: unifiedgosdk.Pointer("978-0-7051-0955-0"),
            Tags: []string{
                "vomito",
                "custodia",
            },
            TotalStock: unifiedgosdk.Pointer[float64](929.0),
            UpdatedAt: types.MustNewTimeFromString("2025-05-25T01:46:03.416Z"),
            Weight: unifiedgosdk.Pointer[float64](61.0),
            WeightUnit: shared.CommerceItemvariantWeightUnitOz.ToPointer(),
            Width: unifiedgosdk.Pointer[float64](26.0),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceItemvariant != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.CreateCommerceItemvariantRequest](../../pkg/models/operations/createcommerceitemvariantrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.CreateCommerceItemvariantResponse](../../pkg/models/operations/createcommerceitemvariantresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCommerceItemvariant

Retrieve an itemvariant

### Example Usage

<!-- UsageSnippet language="go" operationID="getCommerceItemvariant" method="get" path="/commerce/{connection_id}/itemvariant/{id}" -->
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

    res, err := s.Itemvariant.GetCommerceItemvariant(ctx, operations.GetCommerceItemvariantRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceItemvariant != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetCommerceItemvariantRequest](../../pkg/models/operations/getcommerceitemvariantrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.GetCommerceItemvariantResponse](../../pkg/models/operations/getcommerceitemvariantresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCommerceItemvariants

List all itemvariants

### Example Usage

<!-- UsageSnippet language="go" operationID="listCommerceItemvariants" method="get" path="/commerce/{connection_id}/itemvariant" -->
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

    res, err := s.Itemvariant.ListCommerceItemvariants(ctx, operations.ListCommerceItemvariantsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceItemvariants != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ListCommerceItemvariantsRequest](../../pkg/models/operations/listcommerceitemvariantsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.ListCommerceItemvariantsResponse](../../pkg/models/operations/listcommerceitemvariantsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchCommerceItemvariant

Update an itemvariant

### Example Usage

<!-- UsageSnippet language="go" operationID="patchCommerceItemvariant" method="patch" path="/commerce/{connection_id}/itemvariant/{id}" example="commerce_itemvariant" -->
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

    res, err := s.Itemvariant.PatchCommerceItemvariant(ctx, operations.PatchCommerceItemvariantRequest{
        CommerceItemvariant: shared.CommerceItemvariant{
            AvailableAt: types.MustNewTimeFromString("2022-02-02T16:10:33.503Z"),
            CreatedAt: types.MustNewTimeFromString("2022-01-20T13:49:12.968Z"),
            Description: unifiedgosdk.Pointer("Featuring Helium-enhanced technology, our Chips offers unparalleled helpful performance"),
            Height: unifiedgosdk.Pointer[float64](52.0),
            ID: unifiedgosdk.Pointer("eadbfffd-73f1-4b0c-a9d2-863bc87199e3"),
            IsActive: unifiedgosdk.Pointer(false),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsVisible: unifiedgosdk.Pointer(false),
            Length: unifiedgosdk.Pointer[float64](94.0),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Calcar delibero cursim summisse."),
                    Height: unifiedgosdk.Pointer[float64](394.0),
                    ID: unifiedgosdk.Pointer("c820a804-1b36-4c9a-8287-1709262a432b"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("a522f4d0-7438-4f9e-8258-83396f447e36"),
                            Slug: unifiedgosdk.Pointer("illo"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "quia",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](92.0),
                    Type: shared.CommerceItemMediaTypeImage.ToPointer(),
                    URL: "https://picsum.photos/seed/u0YdHqlRu/2007/3208",
                    Width: unifiedgosdk.Pointer[float64](54.0),
                },
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Civitas acies substantia tergo."),
                    Height: unifiedgosdk.Pointer[float64](351.0),
                    ID: unifiedgosdk.Pointer("0503d1dd-5f34-487e-853c-4ff521c430e1"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("67a03d22-fb35-4d92-be27-db0e33f3ada6"),
                            Slug: unifiedgosdk.Pointer("libero"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "capitulus",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](44.0),
                    Type: shared.CommerceItemMediaTypeImage.ToPointer(),
                    URL: "https://loremflickr.com/2230/1237?lock=8628070842159966",
                    Width: unifiedgosdk.Pointer[float64](55.0),
                },
            },
            Metadata: []shared.CommerceMetadata{
                shared.CommerceMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCommerceMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CommerceMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("f84ce569-8d6b-4c86-9c94-17b304106466"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                        "nihil",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Keyboard"),
            Options: []shared.CommerceItemOption{
                shared.CommerceItemOption{
                    ID: unifiedgosdk.Pointer("ecc3d943-9805-4b90-ad6b-73fbdb1aabe3"),
                    Name: "Steel",
                    Position: unifiedgosdk.Pointer[float64](97.0),
                    Values: []string{
                        "Granite",
                        "Plastic",
                    },
                },
            },
            Prices: []shared.CommerceItemPrice{
                shared.CommerceItemPrice{
                    CompareAtPrice: unifiedgosdk.Pointer[float64](3745.0),
                    Currency: unifiedgosdk.Pointer("COP"),
                    Price: 4913.0,
                },
                shared.CommerceItemPrice{
                    CompareAtPrice: unifiedgosdk.Pointer[float64](438.0),
                    Currency: unifiedgosdk.Pointer("PHP"),
                    Price: 1378.0,
                },
                shared.CommerceItemPrice{
                    CompareAtPrice: unifiedgosdk.Pointer[float64](1614.0),
                    Currency: unifiedgosdk.Pointer("PHP"),
                    Price: 8702.0,
                },
            },
            PublicDescription: unifiedgosdk.Pointer("Stylish Soap designed to make you stand out with insistent looks"),
            PublicName: unifiedgosdk.Pointer("Keyboard"),
            RequiresShipping: unifiedgosdk.Pointer(false),
            SizeUnit: shared.SizeUnitCm.ToPointer(),
            Sku: unifiedgosdk.Pointer("978-0-7051-0955-0"),
            Tags: []string{
                "vomito",
                "custodia",
            },
            TotalStock: unifiedgosdk.Pointer[float64](929.0),
            UpdatedAt: types.MustNewTimeFromString("2025-05-25T01:46:03.440Z"),
            Weight: unifiedgosdk.Pointer[float64](61.0),
            WeightUnit: shared.CommerceItemvariantWeightUnitOz.ToPointer(),
            Width: unifiedgosdk.Pointer[float64](26.0),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceItemvariant != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PatchCommerceItemvariantRequest](../../pkg/models/operations/patchcommerceitemvariantrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.PatchCommerceItemvariantResponse](../../pkg/models/operations/patchcommerceitemvariantresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveCommerceItemvariant

Remove an itemvariant

### Example Usage

<!-- UsageSnippet language="go" operationID="removeCommerceItemvariant" method="delete" path="/commerce/{connection_id}/itemvariant/{id}" -->
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

    res, err := s.Itemvariant.RemoveCommerceItemvariant(ctx, operations.RemoveCommerceItemvariantRequest{
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.RemoveCommerceItemvariantRequest](../../pkg/models/operations/removecommerceitemvariantrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.RemoveCommerceItemvariantResponse](../../pkg/models/operations/removecommerceitemvariantresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateCommerceItemvariant

Update an itemvariant

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCommerceItemvariant" method="put" path="/commerce/{connection_id}/itemvariant/{id}" example="commerce_itemvariant" -->
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

    res, err := s.Itemvariant.UpdateCommerceItemvariant(ctx, operations.UpdateCommerceItemvariantRequest{
        CommerceItemvariant: shared.CommerceItemvariant{
            AvailableAt: types.MustNewTimeFromString("2022-02-02T16:10:33.503Z"),
            CreatedAt: types.MustNewTimeFromString("2022-01-20T13:49:12.968Z"),
            Description: unifiedgosdk.Pointer("Featuring Helium-enhanced technology, our Chips offers unparalleled helpful performance"),
            Height: unifiedgosdk.Pointer[float64](52.0),
            ID: unifiedgosdk.Pointer("eadbfffd-73f1-4b0c-a9d2-863bc87199e3"),
            IsActive: unifiedgosdk.Pointer(false),
            IsFeatured: unifiedgosdk.Pointer(false),
            IsVisible: unifiedgosdk.Pointer(false),
            Length: unifiedgosdk.Pointer[float64](94.0),
            Media: []shared.CommerceItemMedia{
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Calcar delibero cursim summisse."),
                    Height: unifiedgosdk.Pointer[float64](394.0),
                    ID: unifiedgosdk.Pointer("c820a804-1b36-4c9a-8287-1709262a432b"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("a522f4d0-7438-4f9e-8258-83396f447e36"),
                            Slug: unifiedgosdk.Pointer("illo"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "quia",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](92.0),
                    Type: shared.CommerceItemMediaTypeImage.ToPointer(),
                    URL: "https://picsum.photos/seed/u0YdHqlRu/2007/3208",
                    Width: unifiedgosdk.Pointer[float64](54.0),
                },
                shared.CommerceItemMedia{
                    Alt: unifiedgosdk.Pointer("Civitas acies substantia tergo."),
                    Height: unifiedgosdk.Pointer[float64](351.0),
                    ID: unifiedgosdk.Pointer("0503d1dd-5f34-487e-853c-4ff521c430e1"),
                    Metadata: []shared.CommerceMetadata{
                        shared.CommerceMetadata{
                            ID: unifiedgosdk.Pointer("67a03d22-fb35-4d92-be27-db0e33f3ada6"),
                            Slug: unifiedgosdk.Pointer("libero"),
                            Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                                "capitulus",
                            )),
                        },
                    },
                    Position: unifiedgosdk.Pointer[float64](44.0),
                    Type: shared.CommerceItemMediaTypeImage.ToPointer(),
                    URL: "https://loremflickr.com/2230/1237?lock=8628070842159966",
                    Width: unifiedgosdk.Pointer[float64](55.0),
                },
            },
            Metadata: []shared.CommerceMetadata{
                shared.CommerceMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCommerceMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CommerceMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("f84ce569-8d6b-4c86-9c94-17b304106466"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCommerceMetadataValueStr(
                        "nihil",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Keyboard"),
            Options: []shared.CommerceItemOption{
                shared.CommerceItemOption{
                    ID: unifiedgosdk.Pointer("ecc3d943-9805-4b90-ad6b-73fbdb1aabe3"),
                    Name: "Steel",
                    Position: unifiedgosdk.Pointer[float64](97.0),
                    Values: []string{
                        "Granite",
                        "Plastic",
                    },
                },
            },
            Prices: []shared.CommerceItemPrice{
                shared.CommerceItemPrice{
                    CompareAtPrice: unifiedgosdk.Pointer[float64](3745.0),
                    Currency: unifiedgosdk.Pointer("COP"),
                    Price: 4913.0,
                },
                shared.CommerceItemPrice{
                    CompareAtPrice: unifiedgosdk.Pointer[float64](438.0),
                    Currency: unifiedgosdk.Pointer("PHP"),
                    Price: 1378.0,
                },
                shared.CommerceItemPrice{
                    CompareAtPrice: unifiedgosdk.Pointer[float64](1614.0),
                    Currency: unifiedgosdk.Pointer("PHP"),
                    Price: 8702.0,
                },
            },
            PublicDescription: unifiedgosdk.Pointer("Stylish Soap designed to make you stand out with insistent looks"),
            PublicName: unifiedgosdk.Pointer("Keyboard"),
            RequiresShipping: unifiedgosdk.Pointer(false),
            SizeUnit: shared.SizeUnitCm.ToPointer(),
            Sku: unifiedgosdk.Pointer("978-0-7051-0955-0"),
            Tags: []string{
                "vomito",
                "custodia",
            },
            TotalStock: unifiedgosdk.Pointer[float64](929.0),
            UpdatedAt: types.MustNewTimeFromString("2025-05-25T01:46:03.440Z"),
            Weight: unifiedgosdk.Pointer[float64](61.0),
            WeightUnit: shared.CommerceItemvariantWeightUnitOz.ToPointer(),
            Width: unifiedgosdk.Pointer[float64](26.0),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommerceItemvariant != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.UpdateCommerceItemvariantRequest](../../pkg/models/operations/updatecommerceitemvariantrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.UpdateCommerceItemvariantResponse](../../pkg/models/operations/updatecommerceitemvariantresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |