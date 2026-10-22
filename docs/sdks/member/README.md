# Member

## Overview

### Available Operations

* [CreateMartechMember](#createmartechmember) - Create a member
* [GetClubsMember](#getclubsmember) - Retrieve a member
* [GetMartechMember](#getmartechmember) - Retrieve a member
* [ListClubsMembers](#listclubsmembers) - List all members
* [ListMartechMembers](#listmartechmembers) - List all members
* [PatchMartechMember](#patchmartechmember) - Update a member
* [RemoveMartechMember](#removemartechmember) - Remove a member
* [UpdateMartechMember](#updatemartechmember) - Update a member

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

    res, err := s.Member.CreateMartechMember(ctx, operations.CreateMartechMemberRequest{
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
            ID: unifiedgosdk.Pointer("e53d71a3-ca93-4d27-b463-3e298d43176a"),
            LastName: unifiedgosdk.Pointer("Leffler"),
            Metadata: []shared.MarketingMetadata{},
            Name: unifiedgosdk.Pointer("Jude Leffler"),
            Status: shared.MarketingMemberStatusUnsubscribed.ToPointer(),
            Tags: []string{
                "adnuo",
                "verumtamen",
            },
            UpdatedAt: types.MustNewTimeFromString("2025-07-09T12:53:01.027Z"),
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

## GetClubsMember

Retrieves a single ClubsMember by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getClubsMember" method="get" path="/clubs/{connection_id}/member/{id}" -->
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

    res, err := s.Member.GetClubsMember(ctx, operations.GetClubsMemberRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ClubsMember != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetClubsMemberRequest](../../pkg/models/operations/getclubsmemberrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetClubsMemberResponse](../../pkg/models/operations/getclubsmemberresponse.md), error**

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

    res, err := s.Member.GetMartechMember(ctx, operations.GetMartechMemberRequest{
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

## ListClubsMembers

Returns a list of ClubsMembers. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listClubsMembers" method="get" path="/clubs/{connection_id}/member" -->
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

    res, err := s.Member.ListClubsMembers(ctx, operations.ListClubsMembersRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ClubsMembers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListClubsMembersRequest](../../pkg/models/operations/listclubsmembersrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListClubsMembersResponse](../../pkg/models/operations/listclubsmembersresponse.md), error**

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

    res, err := s.Member.ListMartechMembers(ctx, operations.ListMartechMembersRequest{
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

    res, err := s.Member.PatchMartechMember(ctx, operations.PatchMartechMemberRequest{
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
            ID: unifiedgosdk.Pointer("40880c0e-6de0-452f-8e79-4d94aef8d633"),
            LastName: unifiedgosdk.Pointer("Leffler"),
            Metadata: []shared.MarketingMetadata{},
            Name: unifiedgosdk.Pointer("Jude Leffler"),
            Status: shared.MarketingMemberStatusUnsubscribed.ToPointer(),
            Tags: []string{
                "adnuo",
                "verumtamen",
            },
            UpdatedAt: types.MustNewTimeFromString("2025-07-09T12:53:01.040Z"),
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

    res, err := s.Member.RemoveMartechMember(ctx, operations.RemoveMartechMemberRequest{
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

    res, err := s.Member.UpdateMartechMember(ctx, operations.UpdateMartechMemberRequest{
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
            ID: unifiedgosdk.Pointer("40880c0e-6de0-452f-8e79-4d94aef8d633"),
            LastName: unifiedgosdk.Pointer("Leffler"),
            Metadata: []shared.MarketingMetadata{},
            Name: unifiedgosdk.Pointer("Jude Leffler"),
            Status: shared.MarketingMemberStatusUnsubscribed.ToPointer(),
            Tags: []string{
                "adnuo",
                "verumtamen",
            },
            UpdatedAt: types.MustNewTimeFromString("2025-07-09T12:53:01.040Z"),
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