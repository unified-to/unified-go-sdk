# Martech

## Overview

### Available Operations

* [CreateMartechCampaign](#createmartechcampaign) - Create a campaign
* [CreateMartechList](#createmartechlist) - Create a list
* [CreateMartechMember](#createmartechmember) - Create a member
* [GetMartechCampaign](#getmartechcampaign) - Retrieve a campaign
* [GetMartechList](#getmartechlist) - Retrieve a list
* [GetMartechMember](#getmartechmember) - Retrieve a member
* [ListMartechCampaigns](#listmartechcampaigns) - List all campaigns
* [ListMartechLists](#listmartechlists) - List all lists
* [ListMartechMembers](#listmartechmembers) - List all members
* [ListMartechReports](#listmartechreports) - List all reports
* [PatchMartechCampaign](#patchmartechcampaign) - Update a campaign
* [PatchMartechList](#patchmartechlist) - Update a list
* [PatchMartechMember](#patchmartechmember) - Update a member
* [RemoveMartechCampaign](#removemartechcampaign) - Remove a campaign
* [RemoveMartechList](#removemartechlist) - Remove a list
* [RemoveMartechMember](#removemartechmember) - Remove a member
* [UpdateMartechCampaign](#updatemartechcampaign) - Update a campaign
* [UpdateMartechList](#updatemartechlist) - Update a list
* [UpdateMartechMember](#updatemartechmember) - Update a member

## CreateMartechCampaign

Creates a MarketingCampaign. Provide the unified MarketingCampaign fields in the request body; they are mapped to the provider's API. Returns the created MarketingCampaign.

### Example Usage

<!-- UsageSnippet language="go" operationID="createMartechCampaign" method="post" path="/martech/{connection_id}/campaign" example="martech_campaign" -->
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

    res, err := s.Martech.CreateMartechCampaign(ctx, operations.CreateMartechCampaignRequest{
        MarketingCampaign: shared.MarketingCampaign{
            CreatedAt: types.MustNewTimeFromString("2019-05-19T11:18:15.589Z"),
            FromEmail: unifiedgosdk.Pointer("liana.dare@gmail.com"),
            FromName: unifiedgosdk.Pointer("Liana Dare"),
            ID: unifiedgosdk.Pointer("d8df22cc-f476-4c7e-bda3-20c9da01c601"),
            Metadata: []shared.MarketingMetadata{},
            Name: unifiedgosdk.Pointer("Sumo trepide utpote conculco tibi solio supplanto."),
            PreviewText: unifiedgosdk.Pointer("Spoliatio tamen strues auxilium benigne tempus error eaque suasoria soleo tergeo contra solutio vilitas."),
            ReplyToEmail: unifiedgosdk.Pointer("liana.dare@icloud.com"),
            SendAt: types.MustNewTimeFromString("2024-01-23T21:19:26.708Z"),
            Status: shared.MarketingCampaignStatusScheduled.ToPointer(),
            SubjectLine: unifiedgosdk.Pointer("Tibi volo auctus votum ver umbra corroboro."),
            Type: unifiedgosdk.Pointer("absplit"),
            UpdatedAt: types.MustNewTimeFromString("2022-07-07T22:52:22.542Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MarketingCampaign != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CreateMartechCampaignRequest](../../pkg/models/operations/createmartechcampaignrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.CreateMartechCampaignResponse](../../pkg/models/operations/createmartechcampaignresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateMartechList

Creates a MarketingList. Provide the unified MarketingList fields in the request body; they are mapped to the provider's API. Returns the created MarketingList.

### Example Usage

<!-- UsageSnippet language="go" operationID="createMartechList" method="post" path="/martech/{connection_id}/list" example="martech_list" -->
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

    res, err := s.Martech.CreateMartechList(ctx, operations.CreateMartechListRequest{
        MarketingList: shared.MarketingList{
            Address: &shared.PropertyMarketingListAddress{
                Address1: unifiedgosdk.Pointer("33559 Darian Center"),
                Address2: unifiedgosdk.Pointer("Suite 796"),
                City: unifiedgosdk.Pointer("East Braulio"),
                Country: unifiedgosdk.Pointer("Liechtenstein"),
                PostalCode: unifiedgosdk.Pointer("65647-1692"),
                Region: unifiedgosdk.Pointer("AZ"),
            },
            CreatedAt: types.MustNewTimeFromString("2019-09-18T02:01:36.950Z"),
            Description: unifiedgosdk.Pointer("Adamo currus at tutamen cubo."),
            ID: unifiedgosdk.Pointer("823ed272-56de-4e52-9c5b-cadf6eed95b2"),
            IsActive: unifiedgosdk.Pointer(true),
            Language: unifiedgosdk.Pointer("pl"),
            Metadata: []shared.MarketingMetadata{},
            Name: unifiedgosdk.Pointer("Kids Newsletter"),
            SenderCompany: unifiedgosdk.Pointer("Monahan, Kerluke and Grant"),
            SenderEmail: unifiedgosdk.Pointer("kaci.weissnat@hotmail.com"),
            SenderName: unifiedgosdk.Pointer("Kaci Weissnat"),
            SenderPhone: unifiedgosdk.Pointer("595.823.3882 x724"),
            Subject: unifiedgosdk.Pointer("Celo calco contra asporto vallum solutio deficio paens demergo."),
            UpdatedAt: types.MustNewTimeFromString("2022-09-29T10:31:49.834Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MarketingList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateMartechListRequest](../../pkg/models/operations/createmartechlistrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateMartechListResponse](../../pkg/models/operations/createmartechlistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateMartechMember

Creates a MarketingMember. Provide the unified MarketingMember fields in the request body; they are mapped to the provider's API. Returns the created MarketingMember.

### Example Usage

<!-- UsageSnippet language="go" operationID="createMartechMember" method="post" path="/martech/{connection_id}/member" example="martech_member" -->
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

    res, err := s.Martech.CreateMartechMember(ctx, operations.CreateMartechMemberRequest{
        MarketingMember: shared.MarketingMember{
            Company: unifiedgosdk.Pointer("Abernathy - Carter"),
            CreatedAt: types.MustNewTimeFromString("2022-04-15T15:32:38.496Z"),
            Emails: []shared.MarketingEmail{
                shared.MarketingEmail{
                    Email: "jeremy.franecki@hotmail.com",
                    Type: shared.MarketingEmailTypeHome.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Jeremy"),
            ID: unifiedgosdk.Pointer("0bac7910-eea4-4eec-bf38-0aca12eccf9f"),
            LastName: unifiedgosdk.Pointer("Franecki"),
            Metadata: []shared.MarketingMetadata{},
            Name: unifiedgosdk.Pointer("Jeremy Franecki"),
            Status: shared.MarketingMemberStatusUnsubscribed.ToPointer(),
            Tags: []string{
                "deleniti",
                "subvenio",
            },
            UpdatedAt: types.MustNewTimeFromString("2025-08-05T00:35:51.716Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MarketingMember != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreateMartechMemberRequest](../../pkg/models/operations/createmartechmemberrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreateMartechMemberResponse](../../pkg/models/operations/createmartechmemberresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetMartechCampaign

Retrieves a single MarketingCampaign by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getMartechCampaign" method="get" path="/martech/{connection_id}/campaign/{id}" -->
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

    res, err := s.Martech.GetMartechCampaign(ctx, operations.GetMartechCampaignRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MarketingCampaign != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetMartechCampaignRequest](../../pkg/models/operations/getmartechcampaignrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetMartechCampaignResponse](../../pkg/models/operations/getmartechcampaignresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetMartechList

Retrieves a single MarketingList by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getMartechList" method="get" path="/martech/{connection_id}/list/{id}" -->
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

    res, err := s.Martech.GetMartechList(ctx, operations.GetMartechListRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MarketingList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetMartechListRequest](../../pkg/models/operations/getmartechlistrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetMartechListResponse](../../pkg/models/operations/getmartechlistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetMartechMember

Retrieves a single MarketingMember by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getMartechMember" method="get" path="/martech/{connection_id}/member/{id}" -->
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

    res, err := s.Martech.GetMartechMember(ctx, operations.GetMartechMemberRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MarketingMember != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetMartechMemberRequest](../../pkg/models/operations/getmartechmemberrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetMartechMemberResponse](../../pkg/models/operations/getmartechmemberresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListMartechCampaigns

Returns a list of MarketingCampaigns. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listMartechCampaigns" method="get" path="/martech/{connection_id}/campaign" -->
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

    res, err := s.Martech.ListMartechCampaigns(ctx, operations.ListMartechCampaignsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MarketingCampaigns != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListMartechCampaignsRequest](../../pkg/models/operations/listmartechcampaignsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListMartechCampaignsResponse](../../pkg/models/operations/listmartechcampaignsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListMartechLists

Returns a list of MarketingLists. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listMartechLists" method="get" path="/martech/{connection_id}/list" -->
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

    res, err := s.Martech.ListMartechLists(ctx, operations.ListMartechListsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MarketingLists != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListMartechListsRequest](../../pkg/models/operations/listmartechlistsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListMartechListsResponse](../../pkg/models/operations/listmartechlistsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListMartechMembers

Returns a list of MarketingMembers. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listMartechMembers" method="get" path="/martech/{connection_id}/member" -->
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

    res, err := s.Martech.ListMartechMembers(ctx, operations.ListMartechMembersRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MarketingMembers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListMartechMembersRequest](../../pkg/models/operations/listmartechmembersrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListMartechMembersResponse](../../pkg/models/operations/listmartechmembersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListMartechReports

Returns a list of MarketingReports. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listMartechReports" method="get" path="/martech/{connection_id}/report" -->
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

    res, err := s.Martech.ListMartechReports(ctx, operations.ListMartechReportsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MarketingReports != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListMartechReportsRequest](../../pkg/models/operations/listmartechreportsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListMartechReportsResponse](../../pkg/models/operations/listmartechreportsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchMartechCampaign

Updates the MarketingCampaign identified by `id`. Provide the unified MarketingCampaign fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchMartechCampaign" method="patch" path="/martech/{connection_id}/campaign/{id}" example="martech_campaign" -->
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

    res, err := s.Martech.PatchMartechCampaign(ctx, operations.PatchMartechCampaignRequest{
        MarketingCampaign: shared.MarketingCampaign{
            CreatedAt: types.MustNewTimeFromString("2019-05-19T11:18:15.589Z"),
            FromEmail: unifiedgosdk.Pointer("liana.dare@gmail.com"),
            FromName: unifiedgosdk.Pointer("Liana Dare"),
            ID: unifiedgosdk.Pointer("5702a811-b6e1-4409-abbc-eaf7aacdea38"),
            Metadata: []shared.MarketingMetadata{},
            Name: unifiedgosdk.Pointer("Sumo trepide utpote conculco tibi solio supplanto."),
            PreviewText: unifiedgosdk.Pointer("Spoliatio tamen strues auxilium benigne tempus error eaque suasoria soleo tergeo contra solutio vilitas."),
            ReplyToEmail: unifiedgosdk.Pointer("liana.dare@icloud.com"),
            SendAt: types.MustNewTimeFromString("2024-01-23T21:19:26.708Z"),
            Status: shared.MarketingCampaignStatusScheduled.ToPointer(),
            SubjectLine: unifiedgosdk.Pointer("Tibi volo auctus votum ver umbra corroboro."),
            Type: unifiedgosdk.Pointer("absplit"),
            UpdatedAt: types.MustNewTimeFromString("2022-07-07T22:52:22.545Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MarketingCampaign != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PatchMartechCampaignRequest](../../pkg/models/operations/patchmartechcampaignrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.PatchMartechCampaignResponse](../../pkg/models/operations/patchmartechcampaignresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchMartechList

Updates the MarketingList identified by `id`. Provide the unified MarketingList fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchMartechList" method="patch" path="/martech/{connection_id}/list/{id}" example="martech_list" -->
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

    res, err := s.Martech.PatchMartechList(ctx, operations.PatchMartechListRequest{
        MarketingList: shared.MarketingList{
            Address: &shared.PropertyMarketingListAddress{
                Address1: unifiedgosdk.Pointer("33559 Darian Center"),
                Address2: unifiedgosdk.Pointer("Suite 796"),
                City: unifiedgosdk.Pointer("East Braulio"),
                Country: unifiedgosdk.Pointer("Liechtenstein"),
                PostalCode: unifiedgosdk.Pointer("65647-1692"),
                Region: unifiedgosdk.Pointer("AZ"),
            },
            CreatedAt: types.MustNewTimeFromString("2019-09-18T02:01:36.950Z"),
            Description: unifiedgosdk.Pointer("Adamo currus at tutamen cubo."),
            ID: unifiedgosdk.Pointer("a24ed362-3d8a-403c-8491-f509ffec5274"),
            IsActive: unifiedgosdk.Pointer(true),
            Language: unifiedgosdk.Pointer("pl"),
            Metadata: []shared.MarketingMetadata{},
            Name: unifiedgosdk.Pointer("Kids Newsletter"),
            SenderCompany: unifiedgosdk.Pointer("Monahan, Kerluke and Grant"),
            SenderEmail: unifiedgosdk.Pointer("kaci.weissnat@hotmail.com"),
            SenderName: unifiedgosdk.Pointer("Kaci Weissnat"),
            SenderPhone: unifiedgosdk.Pointer("595.823.3882 x724"),
            Subject: unifiedgosdk.Pointer("Celo calco contra asporto vallum solutio deficio paens demergo."),
            UpdatedAt: types.MustNewTimeFromString("2022-09-29T10:31:49.840Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MarketingList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchMartechListRequest](../../pkg/models/operations/patchmartechlistrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchMartechListResponse](../../pkg/models/operations/patchmartechlistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchMartechMember

Updates the MarketingMember identified by `id`. Provide the unified MarketingMember fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchMartechMember" method="patch" path="/martech/{connection_id}/member/{id}" example="martech_member" -->
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

    res, err := s.Martech.PatchMartechMember(ctx, operations.PatchMartechMemberRequest{
        MarketingMember: shared.MarketingMember{
            Company: unifiedgosdk.Pointer("Abernathy - Carter"),
            CreatedAt: types.MustNewTimeFromString("2022-04-15T15:32:38.496Z"),
            Emails: []shared.MarketingEmail{
                shared.MarketingEmail{
                    Email: "jeremy.franecki@hotmail.com",
                    Type: shared.MarketingEmailTypeHome.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Jeremy"),
            ID: unifiedgosdk.Pointer("ced778b7-a7a5-41bd-8f27-e5644f6cd0fc"),
            LastName: unifiedgosdk.Pointer("Franecki"),
            Metadata: []shared.MarketingMetadata{},
            Name: unifiedgosdk.Pointer("Jeremy Franecki"),
            Status: shared.MarketingMemberStatusUnsubscribed.ToPointer(),
            Tags: []string{
                "deleniti",
                "subvenio",
            },
            UpdatedAt: types.MustNewTimeFromString("2025-08-05T00:35:51.724Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MarketingMember != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PatchMartechMemberRequest](../../pkg/models/operations/patchmartechmemberrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.PatchMartechMemberResponse](../../pkg/models/operations/patchmartechmemberresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveMartechCampaign

Deletes the MarketingCampaign identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeMartechCampaign" method="delete" path="/martech/{connection_id}/campaign/{id}" -->
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

    res, err := s.Martech.RemoveMartechCampaign(ctx, operations.RemoveMartechCampaignRequest{
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
| `request`                                                                                              | [operations.RemoveMartechCampaignRequest](../../pkg/models/operations/removemartechcampaignrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.RemoveMartechCampaignResponse](../../pkg/models/operations/removemartechcampaignresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveMartechList

Deletes the MarketingList identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeMartechList" method="delete" path="/martech/{connection_id}/list/{id}" -->
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

    res, err := s.Martech.RemoveMartechList(ctx, operations.RemoveMartechListRequest{
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
| `request`                                                                                      | [operations.RemoveMartechListRequest](../../pkg/models/operations/removemartechlistrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemoveMartechListResponse](../../pkg/models/operations/removemartechlistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveMartechMember

Deletes the MarketingMember identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeMartechMember" method="delete" path="/martech/{connection_id}/member/{id}" -->
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

    res, err := s.Martech.RemoveMartechMember(ctx, operations.RemoveMartechMemberRequest{
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
| `request`                                                                                          | [operations.RemoveMartechMemberRequest](../../pkg/models/operations/removemartechmemberrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.RemoveMartechMemberResponse](../../pkg/models/operations/removemartechmemberresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateMartechCampaign

Updates the MarketingCampaign identified by `id`. Provide the unified MarketingCampaign fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateMartechCampaign" method="put" path="/martech/{connection_id}/campaign/{id}" example="martech_campaign" -->
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

    res, err := s.Martech.UpdateMartechCampaign(ctx, operations.UpdateMartechCampaignRequest{
        MarketingCampaign: shared.MarketingCampaign{
            CreatedAt: types.MustNewTimeFromString("2019-05-19T11:18:15.589Z"),
            FromEmail: unifiedgosdk.Pointer("liana.dare@gmail.com"),
            FromName: unifiedgosdk.Pointer("Liana Dare"),
            ID: unifiedgosdk.Pointer("5702a811-b6e1-4409-abbc-eaf7aacdea38"),
            Metadata: []shared.MarketingMetadata{},
            Name: unifiedgosdk.Pointer("Sumo trepide utpote conculco tibi solio supplanto."),
            PreviewText: unifiedgosdk.Pointer("Spoliatio tamen strues auxilium benigne tempus error eaque suasoria soleo tergeo contra solutio vilitas."),
            ReplyToEmail: unifiedgosdk.Pointer("liana.dare@icloud.com"),
            SendAt: types.MustNewTimeFromString("2024-01-23T21:19:26.708Z"),
            Status: shared.MarketingCampaignStatusScheduled.ToPointer(),
            SubjectLine: unifiedgosdk.Pointer("Tibi volo auctus votum ver umbra corroboro."),
            Type: unifiedgosdk.Pointer("absplit"),
            UpdatedAt: types.MustNewTimeFromString("2022-07-07T22:52:22.545Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MarketingCampaign != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdateMartechCampaignRequest](../../pkg/models/operations/updatemartechcampaignrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UpdateMartechCampaignResponse](../../pkg/models/operations/updatemartechcampaignresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateMartechList

Updates the MarketingList identified by `id`. Provide the unified MarketingList fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateMartechList" method="put" path="/martech/{connection_id}/list/{id}" example="martech_list" -->
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

    res, err := s.Martech.UpdateMartechList(ctx, operations.UpdateMartechListRequest{
        MarketingList: shared.MarketingList{
            Address: &shared.PropertyMarketingListAddress{
                Address1: unifiedgosdk.Pointer("33559 Darian Center"),
                Address2: unifiedgosdk.Pointer("Suite 796"),
                City: unifiedgosdk.Pointer("East Braulio"),
                Country: unifiedgosdk.Pointer("Liechtenstein"),
                PostalCode: unifiedgosdk.Pointer("65647-1692"),
                Region: unifiedgosdk.Pointer("AZ"),
            },
            CreatedAt: types.MustNewTimeFromString("2019-09-18T02:01:36.950Z"),
            Description: unifiedgosdk.Pointer("Adamo currus at tutamen cubo."),
            ID: unifiedgosdk.Pointer("a24ed362-3d8a-403c-8491-f509ffec5274"),
            IsActive: unifiedgosdk.Pointer(true),
            Language: unifiedgosdk.Pointer("pl"),
            Metadata: []shared.MarketingMetadata{},
            Name: unifiedgosdk.Pointer("Kids Newsletter"),
            SenderCompany: unifiedgosdk.Pointer("Monahan, Kerluke and Grant"),
            SenderEmail: unifiedgosdk.Pointer("kaci.weissnat@hotmail.com"),
            SenderName: unifiedgosdk.Pointer("Kaci Weissnat"),
            SenderPhone: unifiedgosdk.Pointer("595.823.3882 x724"),
            Subject: unifiedgosdk.Pointer("Celo calco contra asporto vallum solutio deficio paens demergo."),
            UpdatedAt: types.MustNewTimeFromString("2022-09-29T10:31:49.840Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MarketingList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateMartechListRequest](../../pkg/models/operations/updatemartechlistrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateMartechListResponse](../../pkg/models/operations/updatemartechlistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateMartechMember

Updates the MarketingMember identified by `id`. Provide the unified MarketingMember fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateMartechMember" method="put" path="/martech/{connection_id}/member/{id}" example="martech_member" -->
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

    res, err := s.Martech.UpdateMartechMember(ctx, operations.UpdateMartechMemberRequest{
        MarketingMember: shared.MarketingMember{
            Company: unifiedgosdk.Pointer("Abernathy - Carter"),
            CreatedAt: types.MustNewTimeFromString("2022-04-15T15:32:38.496Z"),
            Emails: []shared.MarketingEmail{
                shared.MarketingEmail{
                    Email: "jeremy.franecki@hotmail.com",
                    Type: shared.MarketingEmailTypeHome.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Jeremy"),
            ID: unifiedgosdk.Pointer("ced778b7-a7a5-41bd-8f27-e5644f6cd0fc"),
            LastName: unifiedgosdk.Pointer("Franecki"),
            Metadata: []shared.MarketingMetadata{},
            Name: unifiedgosdk.Pointer("Jeremy Franecki"),
            Status: shared.MarketingMemberStatusUnsubscribed.ToPointer(),
            Tags: []string{
                "deleniti",
                "subvenio",
            },
            UpdatedAt: types.MustNewTimeFromString("2025-08-05T00:35:51.724Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MarketingMember != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateMartechMemberRequest](../../pkg/models/operations/updatemartechmemberrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateMartechMemberResponse](../../pkg/models/operations/updatemartechmemberresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |