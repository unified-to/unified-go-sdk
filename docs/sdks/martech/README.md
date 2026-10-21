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

Create a campaign

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
            FromEmail: unifiedgosdk.Pointer("graham.mertz@gmail.com"),
            FromName: unifiedgosdk.Pointer("Graham Mertz"),
            ID: unifiedgosdk.Pointer("152b6a39-cc03-4bf3-820b-401d75975804"),
            Metadata: []shared.MarketingMetadata{},
            Name: unifiedgosdk.Pointer("Consequatur atqui sustineo."),
            PreviewText: unifiedgosdk.Pointer("Auctus votum ver umbra corroboro turpis spoliatio tamen strues auxilium benigne tempus error eaque suasoria."),
            ReplyToEmail: unifiedgosdk.Pointer("graham.mertz@icloud.com"),
            SendAt: types.MustNewTimeFromString("2024-01-23T21:19:26.708Z"),
            Status: shared.MarketingCampaignStatusScheduled.ToPointer(),
            SubjectLine: unifiedgosdk.Pointer("Trepide utpote conculco tibi solio supplanto peior tibi."),
            Type: unifiedgosdk.Pointer("absplit"),
            UpdatedAt: types.MustNewTimeFromString("2022-06-22T02:50:23.723Z"),
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

Create a list

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
                Address1: unifiedgosdk.Pointer("76564 Coronation Road"),
                Address2: unifiedgosdk.Pointer("Suite 922"),
                City: unifiedgosdk.Pointer("West Jordon"),
                Country: unifiedgosdk.Pointer("Mexico"),
                PostalCode: unifiedgosdk.Pointer("04183"),
                Region: unifiedgosdk.Pointer("NC"),
            },
            CreatedAt: types.MustNewTimeFromString("2019-09-18T02:01:36.950Z"),
            Description: unifiedgosdk.Pointer("Vilis."),
            ID: unifiedgosdk.Pointer("fa25e1e6-f0a8-45b9-b31f-6dfd003964d6"),
            IsActive: unifiedgosdk.Pointer(true),
            Language: unifiedgosdk.Pointer("ru"),
            Metadata: []shared.MarketingMetadata{},
            Name: unifiedgosdk.Pointer("Electronics Newsletter"),
            SenderCompany: unifiedgosdk.Pointer("Koepp and Sons"),
            SenderEmail: unifiedgosdk.Pointer("blake.kuhn@hotmail.com"),
            SenderName: unifiedgosdk.Pointer("Blake Kuhn"),
            SenderPhone: unifiedgosdk.Pointer("359.305.7962 x28115"),
            Subject: unifiedgosdk.Pointer("Theca caecus copia crebro valeo ulciscor."),
            UpdatedAt: types.MustNewTimeFromString("2022-09-13T10:08:43.347Z"),
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

Create a member

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
            Company: unifiedgosdk.Pointer("Miller - Franecki"),
            CreatedAt: types.MustNewTimeFromString("2022-04-15T15:32:38.496Z"),
            Emails: []shared.MarketingEmail{
                shared.MarketingEmail{
                    Email: "jude.leffler@hotmail.com",
                    Type: shared.MarketingEmailTypeHome.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Jude"),
            ID: unifiedgosdk.Pointer("d7b88a89-ae37-4349-a4d9-8a4beafb4569"),
            LastName: unifiedgosdk.Pointer("Leffler"),
            Metadata: []shared.MarketingMetadata{},
            Name: unifiedgosdk.Pointer("Jude Leffler"),
            Status: shared.MarketingMemberStatusUnsubscribed.ToPointer(),
            Tags: []string{
                "adnuo",
                "verumtamen",
            },
            UpdatedAt: types.MustNewTimeFromString("2025-07-08T19:52:27.940Z"),
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

Retrieve a campaign

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

Retrieve a list

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

Retrieve a member

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

List all campaigns

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

List all lists

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

List all members

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

List all reports

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

Update a campaign

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
            FromEmail: unifiedgosdk.Pointer("graham.mertz@gmail.com"),
            FromName: unifiedgosdk.Pointer("Graham Mertz"),
            ID: unifiedgosdk.Pointer("8d59a242-5d0d-4977-9089-a8e78c0a7e7e"),
            Metadata: []shared.MarketingMetadata{},
            Name: unifiedgosdk.Pointer("Consequatur atqui sustineo."),
            PreviewText: unifiedgosdk.Pointer("Auctus votum ver umbra corroboro turpis spoliatio tamen strues auxilium benigne tempus error eaque suasoria."),
            ReplyToEmail: unifiedgosdk.Pointer("graham.mertz@icloud.com"),
            SendAt: types.MustNewTimeFromString("2024-01-23T21:19:26.708Z"),
            Status: shared.MarketingCampaignStatusScheduled.ToPointer(),
            SubjectLine: unifiedgosdk.Pointer("Trepide utpote conculco tibi solio supplanto peior tibi."),
            Type: unifiedgosdk.Pointer("absplit"),
            UpdatedAt: types.MustNewTimeFromString("2022-06-22T02:50:23.732Z"),
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

Update a list

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
                Address1: unifiedgosdk.Pointer("76564 Coronation Road"),
                Address2: unifiedgosdk.Pointer("Suite 922"),
                City: unifiedgosdk.Pointer("West Jordon"),
                Country: unifiedgosdk.Pointer("Mexico"),
                PostalCode: unifiedgosdk.Pointer("04183"),
                Region: unifiedgosdk.Pointer("NC"),
            },
            CreatedAt: types.MustNewTimeFromString("2019-09-18T02:01:36.950Z"),
            Description: unifiedgosdk.Pointer("Vilis."),
            ID: unifiedgosdk.Pointer("fbf9dae2-4ca6-4dd6-8113-05be28e2eadc"),
            IsActive: unifiedgosdk.Pointer(true),
            Language: unifiedgosdk.Pointer("ru"),
            Metadata: []shared.MarketingMetadata{},
            Name: unifiedgosdk.Pointer("Electronics Newsletter"),
            SenderCompany: unifiedgosdk.Pointer("Koepp and Sons"),
            SenderEmail: unifiedgosdk.Pointer("blake.kuhn@hotmail.com"),
            SenderName: unifiedgosdk.Pointer("Blake Kuhn"),
            SenderPhone: unifiedgosdk.Pointer("359.305.7962 x28115"),
            Subject: unifiedgosdk.Pointer("Theca caecus copia crebro valeo ulciscor."),
            UpdatedAt: types.MustNewTimeFromString("2022-09-13T10:08:43.358Z"),
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

Update a member

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
            Company: unifiedgosdk.Pointer("Miller - Franecki"),
            CreatedAt: types.MustNewTimeFromString("2022-04-15T15:32:38.496Z"),
            Emails: []shared.MarketingEmail{
                shared.MarketingEmail{
                    Email: "jude.leffler@hotmail.com",
                    Type: shared.MarketingEmailTypeHome.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Jude"),
            ID: unifiedgosdk.Pointer("e1a7a651-8834-4eb6-b703-61cf8cec3e72"),
            LastName: unifiedgosdk.Pointer("Leffler"),
            Metadata: []shared.MarketingMetadata{},
            Name: unifiedgosdk.Pointer("Jude Leffler"),
            Status: shared.MarketingMemberStatusUnsubscribed.ToPointer(),
            Tags: []string{
                "adnuo",
                "verumtamen",
            },
            UpdatedAt: types.MustNewTimeFromString("2025-07-08T19:52:27.952Z"),
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

Remove a campaign

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

Remove a list

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

Remove a member

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

Update a campaign

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
            FromEmail: unifiedgosdk.Pointer("graham.mertz@gmail.com"),
            FromName: unifiedgosdk.Pointer("Graham Mertz"),
            ID: unifiedgosdk.Pointer("8d59a242-5d0d-4977-9089-a8e78c0a7e7e"),
            Metadata: []shared.MarketingMetadata{},
            Name: unifiedgosdk.Pointer("Consequatur atqui sustineo."),
            PreviewText: unifiedgosdk.Pointer("Auctus votum ver umbra corroboro turpis spoliatio tamen strues auxilium benigne tempus error eaque suasoria."),
            ReplyToEmail: unifiedgosdk.Pointer("graham.mertz@icloud.com"),
            SendAt: types.MustNewTimeFromString("2024-01-23T21:19:26.708Z"),
            Status: shared.MarketingCampaignStatusScheduled.ToPointer(),
            SubjectLine: unifiedgosdk.Pointer("Trepide utpote conculco tibi solio supplanto peior tibi."),
            Type: unifiedgosdk.Pointer("absplit"),
            UpdatedAt: types.MustNewTimeFromString("2022-06-22T02:50:23.732Z"),
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

Update a list

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
                Address1: unifiedgosdk.Pointer("76564 Coronation Road"),
                Address2: unifiedgosdk.Pointer("Suite 922"),
                City: unifiedgosdk.Pointer("West Jordon"),
                Country: unifiedgosdk.Pointer("Mexico"),
                PostalCode: unifiedgosdk.Pointer("04183"),
                Region: unifiedgosdk.Pointer("NC"),
            },
            CreatedAt: types.MustNewTimeFromString("2019-09-18T02:01:36.950Z"),
            Description: unifiedgosdk.Pointer("Vilis."),
            ID: unifiedgosdk.Pointer("fbf9dae2-4ca6-4dd6-8113-05be28e2eadc"),
            IsActive: unifiedgosdk.Pointer(true),
            Language: unifiedgosdk.Pointer("ru"),
            Metadata: []shared.MarketingMetadata{},
            Name: unifiedgosdk.Pointer("Electronics Newsletter"),
            SenderCompany: unifiedgosdk.Pointer("Koepp and Sons"),
            SenderEmail: unifiedgosdk.Pointer("blake.kuhn@hotmail.com"),
            SenderName: unifiedgosdk.Pointer("Blake Kuhn"),
            SenderPhone: unifiedgosdk.Pointer("359.305.7962 x28115"),
            Subject: unifiedgosdk.Pointer("Theca caecus copia crebro valeo ulciscor."),
            UpdatedAt: types.MustNewTimeFromString("2022-09-13T10:08:43.358Z"),
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

Update a member

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
            Company: unifiedgosdk.Pointer("Miller - Franecki"),
            CreatedAt: types.MustNewTimeFromString("2022-04-15T15:32:38.496Z"),
            Emails: []shared.MarketingEmail{
                shared.MarketingEmail{
                    Email: "jude.leffler@hotmail.com",
                    Type: shared.MarketingEmailTypeHome.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Jude"),
            ID: unifiedgosdk.Pointer("e1a7a651-8834-4eb6-b703-61cf8cec3e72"),
            LastName: unifiedgosdk.Pointer("Leffler"),
            Metadata: []shared.MarketingMetadata{},
            Name: unifiedgosdk.Pointer("Jude Leffler"),
            Status: shared.MarketingMemberStatusUnsubscribed.ToPointer(),
            Tags: []string{
                "adnuo",
                "verumtamen",
            },
            UpdatedAt: types.MustNewTimeFromString("2025-07-08T19:52:27.952Z"),
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