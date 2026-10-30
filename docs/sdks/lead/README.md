# Lead

## Overview

### Available Operations

* [CreateCrmLead](#createcrmlead) - Create a lead
* [GetCrmLead](#getcrmlead) - Retrieve a lead
* [ListCrmLeads](#listcrmleads) - List all leads
* [PatchCrmLead](#patchcrmlead) - Update a lead
* [RemoveCrmLead](#removecrmlead) - Remove a lead
* [UpdateCrmLead](#updatecrmlead) - Update a lead

## CreateCrmLead

Creates a CrmLead. Provide the unified CrmLead fields in the request body; they are mapped to the provider's API. Returns the created CrmLead.

### Example Usage

<!-- UsageSnippet language="go" operationID="createCrmLead" method="post" path="/crm/{connection_id}/lead" example="crm_lead" -->
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

    res, err := s.Lead.CreateCrmLead(ctx, operations.CreateCrmLeadRequest{
        CrmLead: shared.CrmLead{
            Address: &shared.PropertyCrmLeadAddress{
                Address1: unifiedgosdk.Pointer("87206 Hudson Street"),
                City: unifiedgosdk.Pointer("Leslieboro"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("49909-1183"),
                Region: unifiedgosdk.Pointer("Michigan"),
                RegionCode: unifiedgosdk.Pointer("MO"),
            },
            CompanyName: unifiedgosdk.Pointer("Tillman - Wiegand"),
            CreatedAt: types.MustNewTimeFromString("2019-10-12T11:27:59.003Z"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("velda.sporer@icloud.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Velda"),
            ID: unifiedgosdk.Pointer("f2068b3a-f69f-43fc-a6d2-66a2fde33990"),
            IsActive: unifiedgosdk.Pointer(true),
            LastName: unifiedgosdk.Pointer("Sporer"),
            LinkUrls: []string{
                "https://www.linkedin.com/in/velda-sporer",
            },
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("b49fd907-193c-4a5d-891c-b0d8daf526d2"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "vitiosus",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Velda Sporer"),
            Source: unifiedgosdk.Pointer("Advertisement"),
            Status: unifiedgosdk.Pointer("New"),
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "(841) 271-0595",
                    Type: shared.CrmTelephoneTypeFax.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(664) 498-4952",
                    Type: shared.CrmTelephoneTypeHome.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(281) 288-0037",
                    Type: shared.CrmTelephoneTypeOther.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("International Factors Administrator"),
            UpdatedAt: types.MustNewTimeFromString("2020-05-18T15:36:13.157Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CrmLead != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateCrmLeadRequest](../../pkg/models/operations/createcrmleadrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreateCrmLeadResponse](../../pkg/models/operations/createcrmleadresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCrmLead

Retrieves a single CrmLead by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getCrmLead" method="get" path="/crm/{connection_id}/lead/{id}" -->
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

    res, err := s.Lead.GetCrmLead(ctx, operations.GetCrmLeadRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CrmLead != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetCrmLeadRequest](../../pkg/models/operations/getcrmleadrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.GetCrmLeadResponse](../../pkg/models/operations/getcrmleadresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCrmLeads

Returns a list of CrmLeads. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listCrmLeads" method="get" path="/crm/{connection_id}/lead" -->
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

    res, err := s.Lead.ListCrmLeads(ctx, operations.ListCrmLeadsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CrmLeads != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListCrmLeadsRequest](../../pkg/models/operations/listcrmleadsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListCrmLeadsResponse](../../pkg/models/operations/listcrmleadsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchCrmLead

Updates the CrmLead identified by `id`. Provide the unified CrmLead fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchCrmLead" method="patch" path="/crm/{connection_id}/lead/{id}" example="crm_lead" -->
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

    res, err := s.Lead.PatchCrmLead(ctx, operations.PatchCrmLeadRequest{
        CrmLead: shared.CrmLead{
            Address: &shared.PropertyCrmLeadAddress{
                Address1: unifiedgosdk.Pointer("87206 Hudson Street"),
                City: unifiedgosdk.Pointer("Leslieboro"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("49909-1183"),
                Region: unifiedgosdk.Pointer("Michigan"),
                RegionCode: unifiedgosdk.Pointer("MO"),
            },
            CompanyName: unifiedgosdk.Pointer("Tillman - Wiegand"),
            CreatedAt: types.MustNewTimeFromString("2019-10-12T11:27:59.003Z"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("velda.sporer@icloud.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Velda"),
            ID: unifiedgosdk.Pointer("070541e0-d71c-4748-a6a1-3759a0166f62"),
            IsActive: unifiedgosdk.Pointer(true),
            LastName: unifiedgosdk.Pointer("Sporer"),
            LinkUrls: []string{
                "https://www.linkedin.com/in/velda-sporer",
            },
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("f164ed4a-dc03-44c8-93db-ae542c73814e"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "vitiosus",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Velda Sporer"),
            Source: unifiedgosdk.Pointer("Advertisement"),
            Status: unifiedgosdk.Pointer("New"),
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "(841) 271-0595",
                    Type: shared.CrmTelephoneTypeFax.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(664) 498-4952",
                    Type: shared.CrmTelephoneTypeHome.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(281) 288-0037",
                    Type: shared.CrmTelephoneTypeOther.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("International Factors Administrator"),
            UpdatedAt: types.MustNewTimeFromString("2020-05-18T15:36:13.160Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CrmLead != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.PatchCrmLeadRequest](../../pkg/models/operations/patchcrmleadrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.PatchCrmLeadResponse](../../pkg/models/operations/patchcrmleadresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveCrmLead

Deletes the CrmLead identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeCrmLead" method="delete" path="/crm/{connection_id}/lead/{id}" -->
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

    res, err := s.Lead.RemoveCrmLead(ctx, operations.RemoveCrmLeadRequest{
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.RemoveCrmLeadRequest](../../pkg/models/operations/removecrmleadrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.RemoveCrmLeadResponse](../../pkg/models/operations/removecrmleadresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateCrmLead

Updates the CrmLead identified by `id`. Provide the unified CrmLead fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCrmLead" method="put" path="/crm/{connection_id}/lead/{id}" example="crm_lead" -->
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

    res, err := s.Lead.UpdateCrmLead(ctx, operations.UpdateCrmLeadRequest{
        CrmLead: shared.CrmLead{
            Address: &shared.PropertyCrmLeadAddress{
                Address1: unifiedgosdk.Pointer("87206 Hudson Street"),
                City: unifiedgosdk.Pointer("Leslieboro"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("49909-1183"),
                Region: unifiedgosdk.Pointer("Michigan"),
                RegionCode: unifiedgosdk.Pointer("MO"),
            },
            CompanyName: unifiedgosdk.Pointer("Tillman - Wiegand"),
            CreatedAt: types.MustNewTimeFromString("2019-10-12T11:27:59.003Z"),
            Emails: []shared.CrmEmail{
                shared.CrmEmail{
                    Email: unifiedgosdk.Pointer("velda.sporer@icloud.com"),
                    Type: shared.CrmEmailTypeWork.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Velda"),
            ID: unifiedgosdk.Pointer("070541e0-d71c-4748-a6a1-3759a0166f62"),
            IsActive: unifiedgosdk.Pointer(true),
            LastName: unifiedgosdk.Pointer("Sporer"),
            LinkUrls: []string{
                "https://www.linkedin.com/in/velda-sporer",
            },
            Metadata: []shared.CrmMetadata{
                shared.CrmMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateCrmMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.CrmMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("f164ed4a-dc03-44c8-93db-ae542c73814e"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateCrmMetadataValueStr(
                        "vitiosus",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Velda Sporer"),
            Source: unifiedgosdk.Pointer("Advertisement"),
            Status: unifiedgosdk.Pointer("New"),
            Telephones: []shared.CrmTelephone{
                shared.CrmTelephone{
                    Telephone: "(841) 271-0595",
                    Type: shared.CrmTelephoneTypeFax.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(664) 498-4952",
                    Type: shared.CrmTelephoneTypeHome.ToPointer(),
                },
                shared.CrmTelephone{
                    Telephone: "(281) 288-0037",
                    Type: shared.CrmTelephoneTypeOther.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("International Factors Administrator"),
            UpdatedAt: types.MustNewTimeFromString("2020-05-18T15:36:13.160Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CrmLead != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateCrmLeadRequest](../../pkg/models/operations/updatecrmleadrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.UpdateCrmLeadResponse](../../pkg/models/operations/updatecrmleadresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |