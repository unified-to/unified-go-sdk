# Campaign

## Overview

### Available Operations

* [CreateAdsCampaign](#createadscampaign) - Create a campaign
* [CreateMartechCampaign](#createmartechcampaign) - Create a campaign
* [GetAdsCampaign](#getadscampaign) - Retrieve a campaign
* [GetMartechCampaign](#getmartechcampaign) - Retrieve a campaign
* [ListAdsCampaigns](#listadscampaigns) - List all campaigns
* [ListMartechCampaigns](#listmartechcampaigns) - List all campaigns
* [PatchAdsCampaign](#patchadscampaign) - Update a campaign
* [PatchMartechCampaign](#patchmartechcampaign) - Update a campaign
* [RemoveAdsCampaign](#removeadscampaign) - Remove a campaign
* [RemoveMartechCampaign](#removemartechcampaign) - Remove a campaign
* [UpdateAdsCampaign](#updateadscampaign) - Update a campaign
* [UpdateMartechCampaign](#updatemartechcampaign) - Update a campaign

## CreateAdsCampaign

Create a campaign

### Example Usage

<!-- UsageSnippet language="go" operationID="createAdsCampaign" method="post" path="/ads/{connection_id}/campaign" example="ads_campaign" -->
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

    res, err := s.Campaign.CreateAdsCampaign(ctx, operations.CreateAdsCampaignRequest{
        AdsCampaign: shared.AdsCampaign{
            BudgetAmount: unifiedgosdk.Pointer[float64](8743.179536121897),
            BudgetPeriod: shared.BudgetPeriodMonthly.ToPointer(),
            Category: unifiedgosdk.Pointer("CREDIT"),
            CreatedAt: types.MustNewTimeFromString("2022-05-21T08:51:41.868Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            EffectiveStatus: shared.EffectiveStatusNotEligible.ToPointer(),
            EndAt: types.MustNewTimeFromString("2025-05-09T08:59:51.604Z"),
            ID: unifiedgosdk.Pointer("751e775a-d107-436d-beae-0b6e29a81866"),
            Labels: []string{
                "comedo",
            },
            Name: unifiedgosdk.Pointer("Emard Inc"),
            StartAt: types.MustNewTimeFromString("2022-07-20T04:53:38.349Z"),
            Status: shared.AdsCampaignStatusProcessingFailed.ToPointer(),
            Targeting: &shared.PropertyAdsCampaignTargeting{},
            TotalSpendAmount: unifiedgosdk.Pointer[float64](2349.8642875347286),
            UpdatedAt: types.MustNewTimeFromString("2025-12-05T14:24:38.564Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsCampaign != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateAdsCampaignRequest](../../pkg/models/operations/createadscampaignrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateAdsCampaignResponse](../../pkg/models/operations/createadscampaignresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

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

    res, err := s.Campaign.CreateMartechCampaign(ctx, operations.CreateMartechCampaignRequest{
        MarketingCampaign: shared.MarketingCampaign{
            CreatedAt: types.MustNewTimeFromString("2023-08-01T22:29:12.121Z"),
            FromEmail: unifiedgosdk.Pointer("Nick.Beahan@hotmail.com"),
            FromName: unifiedgosdk.Pointer("Javier Rempel"),
            ID: unifiedgosdk.Pointer("429f23e8-28a5-4533-a66d-46de29192df8"),
            ListIds: []string{
                "bde5cab9-cf2f-4ed5-adab-b33c88bac5af",
            },
            Name: unifiedgosdk.Pointer("Consequatur atqui sustineo."),
            PreviewText: unifiedgosdk.Pointer("Bellicus tener cinis causa cavus toties."),
            ReplyToEmail: unifiedgosdk.Pointer("Antwan.Abshire@hotmail.com"),
            SendAt: types.MustNewTimeFromString("2023-03-28T12:33:25.052Z"),
            Status: shared.MarketingCampaignStatusSent.ToPointer(),
            SubjectLine: unifiedgosdk.Pointer("Depromo depulso turpis teres apparatus placeat ventus tolero cunctatio."),
            Type: unifiedgosdk.Pointer("plaintext"),
            UpdatedAt: types.MustNewTimeFromString("2023-12-17T22:11:31.702Z"),
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

## GetAdsCampaign

Retrieve a campaign

### Example Usage

<!-- UsageSnippet language="go" operationID="getAdsCampaign" method="get" path="/ads/{connection_id}/campaign/{id}" -->
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

    res, err := s.Campaign.GetAdsCampaign(ctx, operations.GetAdsCampaignRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsCampaign != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetAdsCampaignRequest](../../pkg/models/operations/getadscampaignrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetAdsCampaignResponse](../../pkg/models/operations/getadscampaignresponse.md), error**

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

    res, err := s.Campaign.GetMartechCampaign(ctx, operations.GetMartechCampaignRequest{
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

## ListAdsCampaigns

List all campaigns

### Example Usage

<!-- UsageSnippet language="go" operationID="listAdsCampaigns" method="get" path="/ads/{connection_id}/campaign" -->
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

    res, err := s.Campaign.ListAdsCampaigns(ctx, operations.ListAdsCampaignsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsCampaigns != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListAdsCampaignsRequest](../../pkg/models/operations/listadscampaignsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListAdsCampaignsResponse](../../pkg/models/operations/listadscampaignsresponse.md), error**

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

    res, err := s.Campaign.ListMartechCampaigns(ctx, operations.ListMartechCampaignsRequest{
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

## PatchAdsCampaign

Update a campaign

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAdsCampaign" method="patch" path="/ads/{connection_id}/campaign/{id}" example="ads_campaign" -->
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

    res, err := s.Campaign.PatchAdsCampaign(ctx, operations.PatchAdsCampaignRequest{
        AdsCampaign: shared.AdsCampaign{
            BudgetAmount: unifiedgosdk.Pointer[float64](8743.179536121897),
            BudgetPeriod: shared.BudgetPeriodMonthly.ToPointer(),
            Category: unifiedgosdk.Pointer("CREDIT"),
            CreatedAt: types.MustNewTimeFromString("2022-05-21T08:51:41.868Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            EffectiveStatus: shared.EffectiveStatusNotEligible.ToPointer(),
            EndAt: types.MustNewTimeFromString("2025-05-09T08:59:51.668Z"),
            ID: unifiedgosdk.Pointer("2d02c149-981a-4c0b-b778-b196fd997f46"),
            Labels: []string{
                "comedo",
            },
            Name: unifiedgosdk.Pointer("Emard Inc"),
            StartAt: types.MustNewTimeFromString("2022-07-20T04:53:38.352Z"),
            Status: shared.AdsCampaignStatusProcessingFailed.ToPointer(),
            Targeting: &shared.PropertyAdsCampaignTargeting{},
            TotalSpendAmount: unifiedgosdk.Pointer[float64](2349.8642875347286),
            UpdatedAt: types.MustNewTimeFromString("2025-12-05T14:24:38.640Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsCampaign != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchAdsCampaignRequest](../../pkg/models/operations/patchadscampaignrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchAdsCampaignResponse](../../pkg/models/operations/patchadscampaignresponse.md), error**

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

    res, err := s.Campaign.PatchMartechCampaign(ctx, operations.PatchMartechCampaignRequest{
        MarketingCampaign: shared.MarketingCampaign{
            CreatedAt: types.MustNewTimeFromString("2023-08-01T22:29:12.121Z"),
            FromEmail: unifiedgosdk.Pointer("Nick.Beahan@hotmail.com"),
            FromName: unifiedgosdk.Pointer("Javier Rempel"),
            ID: unifiedgosdk.Pointer("2062d840-fdc3-44ac-9ead-6ff30a5209df"),
            ListIds: []string{
                "bde5cab9-cf2f-4ed5-adab-b33c88bac5af",
            },
            Name: unifiedgosdk.Pointer("Consequatur atqui sustineo."),
            PreviewText: unifiedgosdk.Pointer("Bellicus tener cinis causa cavus toties."),
            ReplyToEmail: unifiedgosdk.Pointer("Antwan.Abshire@hotmail.com"),
            SendAt: types.MustNewTimeFromString("2023-03-28T12:33:25.052Z"),
            Status: shared.MarketingCampaignStatusSent.ToPointer(),
            SubjectLine: unifiedgosdk.Pointer("Depromo depulso turpis teres apparatus placeat ventus tolero cunctatio."),
            Type: unifiedgosdk.Pointer("plaintext"),
            UpdatedAt: types.MustNewTimeFromString("2023-12-17T22:11:31.702Z"),
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

## RemoveAdsCampaign

Remove a campaign

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAdsCampaign" method="delete" path="/ads/{connection_id}/campaign/{id}" -->
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

    res, err := s.Campaign.RemoveAdsCampaign(ctx, operations.RemoveAdsCampaignRequest{
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
| `request`                                                                                      | [operations.RemoveAdsCampaignRequest](../../pkg/models/operations/removeadscampaignrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemoveAdsCampaignResponse](../../pkg/models/operations/removeadscampaignresponse.md), error**

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

    res, err := s.Campaign.RemoveMartechCampaign(ctx, operations.RemoveMartechCampaignRequest{
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

## UpdateAdsCampaign

Update a campaign

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAdsCampaign" method="put" path="/ads/{connection_id}/campaign/{id}" example="ads_campaign" -->
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

    res, err := s.Campaign.UpdateAdsCampaign(ctx, operations.UpdateAdsCampaignRequest{
        AdsCampaign: shared.AdsCampaign{
            BudgetAmount: unifiedgosdk.Pointer[float64](8743.179536121897),
            BudgetPeriod: shared.BudgetPeriodMonthly.ToPointer(),
            Category: unifiedgosdk.Pointer("CREDIT"),
            CreatedAt: types.MustNewTimeFromString("2022-05-21T08:51:41.868Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            EffectiveStatus: shared.EffectiveStatusNotEligible.ToPointer(),
            EndAt: types.MustNewTimeFromString("2025-05-09T08:59:51.668Z"),
            ID: unifiedgosdk.Pointer("2d02c149-981a-4c0b-b778-b196fd997f46"),
            Labels: []string{
                "comedo",
            },
            Name: unifiedgosdk.Pointer("Emard Inc"),
            StartAt: types.MustNewTimeFromString("2022-07-20T04:53:38.352Z"),
            Status: shared.AdsCampaignStatusProcessingFailed.ToPointer(),
            Targeting: &shared.PropertyAdsCampaignTargeting{},
            TotalSpendAmount: unifiedgosdk.Pointer[float64](2349.8642875347286),
            UpdatedAt: types.MustNewTimeFromString("2025-12-05T14:24:38.640Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsCampaign != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateAdsCampaignRequest](../../pkg/models/operations/updateadscampaignrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateAdsCampaignResponse](../../pkg/models/operations/updateadscampaignresponse.md), error**

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

    res, err := s.Campaign.UpdateMartechCampaign(ctx, operations.UpdateMartechCampaignRequest{
        MarketingCampaign: shared.MarketingCampaign{
            CreatedAt: types.MustNewTimeFromString("2023-08-01T22:29:12.121Z"),
            FromEmail: unifiedgosdk.Pointer("Nick.Beahan@hotmail.com"),
            FromName: unifiedgosdk.Pointer("Javier Rempel"),
            ID: unifiedgosdk.Pointer("2062d840-fdc3-44ac-9ead-6ff30a5209df"),
            ListIds: []string{
                "bde5cab9-cf2f-4ed5-adab-b33c88bac5af",
            },
            Name: unifiedgosdk.Pointer("Consequatur atqui sustineo."),
            PreviewText: unifiedgosdk.Pointer("Bellicus tener cinis causa cavus toties."),
            ReplyToEmail: unifiedgosdk.Pointer("Antwan.Abshire@hotmail.com"),
            SendAt: types.MustNewTimeFromString("2023-03-28T12:33:25.052Z"),
            Status: shared.MarketingCampaignStatusSent.ToPointer(),
            SubjectLine: unifiedgosdk.Pointer("Depromo depulso turpis teres apparatus placeat ventus tolero cunctatio."),
            Type: unifiedgosdk.Pointer("plaintext"),
            UpdatedAt: types.MustNewTimeFromString("2023-12-17T22:11:31.702Z"),
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