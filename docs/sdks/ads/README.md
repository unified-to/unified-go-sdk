# Ads

## Overview

### Available Operations

* [CreateAdsAd](#createadsad) - Create an ad
* [CreateAdsAsset](#createadsasset) - Create an asset
* [CreateAdsCampaign](#createadscampaign) - Create a campaign
* [CreateAdsCreative](#createadscreative) - Create a creative
* [CreateAdsGroup](#createadsgroup) - Create a group
* [CreateAdsInsertionorder](#createadsinsertionorder) - Create an insertionorder
* [CreateAdsOrganization](#createadsorganization) - Create an organization
* [GetAdsAd](#getadsad) - Retrieve an ad
* [GetAdsAsset](#getadsasset) - Retrieve an asset
* [GetAdsCampaign](#getadscampaign) - Retrieve a campaign
* [GetAdsCreative](#getadscreative) - Retrieve a creative
* [GetAdsGroup](#getadsgroup) - Retrieve a group
* [GetAdsInsertionorder](#getadsinsertionorder) - Retrieve an insertionorder
* [GetAdsOrganization](#getadsorganization) - Retrieve an organization
* [GetAdsPromoted](#getadspromoted) - Retrieve a promoted
* [GetAdsTarget](#getadstarget) - Retrieve a target
* [ListAdsAds](#listadsads) - List all ads
* [ListAdsAssets](#listadsassets) - List all assets
* [ListAdsCampaigns](#listadscampaigns) - List all campaigns
* [ListAdsCreatives](#listadscreatives) - List all creatives
* [ListAdsGroups](#listadsgroups) - List all groups
* [ListAdsInsertionorders](#listadsinsertionorders) - List all insertionorders
* [ListAdsOrganizations](#listadsorganizations) - List all organizations
* [ListAdsPromoteds](#listadspromoteds) - List all promoteds
* [ListAdsReports](#listadsreports) - List all reports
* [ListAdsTargets](#listadstargets) - List all targets
* [PatchAdsAd](#patchadsad) - Update an ad
* [PatchAdsCampaign](#patchadscampaign) - Update a campaign
* [PatchAdsCreative](#patchadscreative) - Update a creative
* [PatchAdsGroup](#patchadsgroup) - Update a group
* [PatchAdsInsertionorder](#patchadsinsertionorder) - Update an insertionorder
* [PatchAdsOrganization](#patchadsorganization) - Update an organization
* [RemoveAdsAd](#removeadsad) - Remove an ad
* [RemoveAdsCampaign](#removeadscampaign) - Remove a campaign
* [RemoveAdsCreative](#removeadscreative) - Remove a creative
* [RemoveAdsGroup](#removeadsgroup) - Remove a group
* [RemoveAdsInsertionorder](#removeadsinsertionorder) - Remove an insertionorder
* [RemoveAdsOrganization](#removeadsorganization) - Remove an organization
* [UpdateAdsAd](#updateadsad) - Update an ad
* [UpdateAdsCampaign](#updateadscampaign) - Update a campaign
* [UpdateAdsCreative](#updateadscreative) - Update a creative
* [UpdateAdsGroup](#updateadsgroup) - Update a group
* [UpdateAdsInsertionorder](#updateadsinsertionorder) - Update an insertionorder
* [UpdateAdsOrganization](#updateadsorganization) - Update an organization

## CreateAdsAd

Create an ad

### Example Usage

<!-- UsageSnippet language="go" operationID="createAdsAd" method="post" path="/ads/{connection_id}/ad" example="ads_ad" -->
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

    res, err := s.Ads.CreateAdsAd(ctx, operations.CreateAdsAdRequest{
        AdsAd: shared.AdsAd{
            AdCopy: unifiedgosdk.Pointer("Ascisco tolero caute sapiente. Valens unde comedo cursus crinis nobis thema. Cohaero nisi ullam tum unde ultio vilicus auditor capio."),
            AdType: shared.AdTypeSocial.ToPointer(),
            AdvertiserName: unifiedgosdk.Pointer("Robel, Nader and Rau"),
            CreatedAt: types.MustNewTimeFromString("2022-11-08T03:38:20.978Z"),
            CreativeAssetURL: unifiedgosdk.Pointer("https://picsum.photos/seed/LwOzrpr9/948/2793"),
            Description: unifiedgosdk.Pointer("Accedo vespillo carpo dolor decet stillicidium comptus tenuis."),
            FinalURL: unifiedgosdk.Pointer("https://improbable-sanity.com"),
            ID: unifiedgosdk.Pointer("7d6bc006-9a9e-4304-a7d3-61419c806e1a"),
            Name: unifiedgosdk.Pointer("Hermiston Group"),
            Status: shared.AdsAdStatusArchived.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-06-05T03:02:50.902Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsAd != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CreateAdsAdRequest](../../pkg/models/operations/createadsadrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.CreateAdsAdResponse](../../pkg/models/operations/createadsadresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateAdsAsset

Create an asset

### Example Usage

<!-- UsageSnippet language="go" operationID="createAdsAsset" method="post" path="/ads/{connection_id}/asset" example="ads_asset" -->
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

    res, err := s.Ads.CreateAdsAsset(ctx, operations.CreateAdsAssetRequest{
        AdsAsset: shared.AdsAsset{
            CreatedAt: types.MustNewTimeFromString("2020-03-27T20:14:38.603Z"),
            Height: unifiedgosdk.Pointer[float64](400.0),
            ID: unifiedgosdk.Pointer("69ac339c-ea18-45a6-8c03-fadb8ed8e424"),
            MimeType: unifiedgosdk.Pointer("IMAGE_PNG"),
            Name: unifiedgosdk.Pointer("Lockman - DuBuque"),
            Type: shared.AdsAssetTypeImage.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-03-15T04:43:22.538Z"),
            URL: unifiedgosdk.Pointer("https://informal-perfection.com/"),
            Width: unifiedgosdk.Pointer[float64](600.0),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsAsset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateAdsAssetRequest](../../pkg/models/operations/createadsassetrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.CreateAdsAssetResponse](../../pkg/models/operations/createadsassetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

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

    res, err := s.Ads.CreateAdsCampaign(ctx, operations.CreateAdsCampaignRequest{
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

## CreateAdsCreative

Create a creative

### Example Usage

<!-- UsageSnippet language="go" operationID="createAdsCreative" method="post" path="/ads/{connection_id}/creative" example="ads_creative" -->
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

    res, err := s.Ads.CreateAdsCreative(ctx, operations.CreateAdsCreativeRequest{
        AdsCreative: shared.AdsCreative{
            CreatedAt: types.MustNewTimeFromString("2020-02-17T11:24:51.093Z"),
            ID: unifiedgosdk.Pointer("4d31950c-af56-406c-bac7-75ce978929c1"),
            Labels: []string{
                "coma",
                "accedo",
                "termes",
            },
            Name: unifiedgosdk.Pointer("Brekke, Bradtke and Robel"),
            Status: shared.AdsCreativeStatusPaused.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2021-06-21T01:13:41.795Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsCreative != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateAdsCreativeRequest](../../pkg/models/operations/createadscreativerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateAdsCreativeResponse](../../pkg/models/operations/createadscreativeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateAdsGroup

Create a group

### Example Usage

<!-- UsageSnippet language="go" operationID="createAdsGroup" method="post" path="/ads/{connection_id}/group" example="ads_group" -->
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

    res, err := s.Ads.CreateAdsGroup(ctx, operations.CreateAdsGroupRequest{
        AdsGroup: shared.AdsGroup{
            BidAmount: unifiedgosdk.Pointer[float64](26.16030164062977),
            BudgetAmount: unifiedgosdk.Pointer[float64](5099.175239447504),
            BudgetPeriod: shared.AdsGroupBudgetPeriodMonthly.ToPointer(),
            CreatedAt: types.MustNewTimeFromString("2019-08-29T17:59:41.045Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            EffectiveStatus: shared.AdsGroupEffectiveStatusPaused.ToPointer(),
            EndAt: types.MustNewTimeFromString("2026-05-24T14:15:58.302Z"),
            ID: unifiedgosdk.Pointer("497eb0be-612c-416e-be9f-14f379e66106"),
            LanguageLocale: unifiedgosdk.Pointer("fr-FR"),
            Name: unifiedgosdk.Pointer("Stark - Baumbach"),
            StartAt: types.MustNewTimeFromString("2025-12-10T22:04:10.683Z"),
            Status: shared.AdsGroupStatusProcessing.ToPointer(),
            Targeting: &shared.PropertyAdsGroupTargeting{},
            UpdatedAt: types.MustNewTimeFromString("2022-01-02T17:05:47.220Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateAdsGroupRequest](../../pkg/models/operations/createadsgrouprequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.CreateAdsGroupResponse](../../pkg/models/operations/createadsgroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateAdsInsertionorder

Create an insertionorder

### Example Usage

<!-- UsageSnippet language="go" operationID="createAdsInsertionorder" method="post" path="/ads/{connection_id}/insertionorder" example="ads_insertionorder" -->
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

    res, err := s.Ads.CreateAdsInsertionorder(ctx, operations.CreateAdsInsertionorderRequest{
        AdsInsertionorder: shared.AdsInsertionorder{
            CreatedAt: types.MustNewTimeFromString("2021-04-10T06:57:36.611Z"),
            ID: unifiedgosdk.Pointer("0025631a-c198-4691-9215-01f54de73cae"),
            Name: unifiedgosdk.Pointer("Kunde, Smith and Reinger"),
            Status: shared.AdsInsertionorderStatusUnspecified.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2021-04-28T12:31:55.184Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsInsertionorder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.CreateAdsInsertionorderRequest](../../pkg/models/operations/createadsinsertionorderrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.CreateAdsInsertionorderResponse](../../pkg/models/operations/createadsinsertionorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateAdsOrganization

Create an organization

### Example Usage

<!-- UsageSnippet language="go" operationID="createAdsOrganization" method="post" path="/ads/{connection_id}/organization" example="ads_organization" -->
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

    res, err := s.Ads.CreateAdsOrganization(ctx, operations.CreateAdsOrganizationRequest{
        AdsOrganization: shared.AdsOrganization{
            AccountNumber: unifiedgosdk.Pointer("LQUJx8zQBW"),
            CreatedAt: types.MustNewTimeFromString("2020-07-23T21:47:11.440Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            ID: unifiedgosdk.Pointer("09562003-d3da-4510-9573-e932b1798d06"),
            Managers: []shared.AdsManager{
                shared.AdsManager{
                    ID: unifiedgosdk.Pointer("e4fd87df-9f8b-4fa0-a77b-b7d18669e350"),
                    Name: unifiedgosdk.Pointer("Parker, Leannon and Gibson"),
                },
            },
            Name: unifiedgosdk.Pointer("Ankunding Inc"),
            Status: shared.AdsOrganizationStatusProcessing.ToPointer(),
            Timezone: unifiedgosdk.Pointer("Europe/Chisinau"),
            UpdatedAt: types.MustNewTimeFromString("2026-02-27T03:26:08.210Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsOrganization != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CreateAdsOrganizationRequest](../../pkg/models/operations/createadsorganizationrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.CreateAdsOrganizationResponse](../../pkg/models/operations/createadsorganizationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAdsAd

Retrieve an ad

### Example Usage

<!-- UsageSnippet language="go" operationID="getAdsAd" method="get" path="/ads/{connection_id}/ad/{id}" -->
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

    res, err := s.Ads.GetAdsAd(ctx, operations.GetAdsAdRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsAd != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetAdsAdRequest](../../pkg/models/operations/getadsadrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                 | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.GetAdsAdResponse](../../pkg/models/operations/getadsadresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAdsAsset

Retrieve an asset

### Example Usage

<!-- UsageSnippet language="go" operationID="getAdsAsset" method="get" path="/ads/{connection_id}/asset/{id}" -->
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

    res, err := s.Ads.GetAdsAsset(ctx, operations.GetAdsAssetRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsAsset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetAdsAssetRequest](../../pkg/models/operations/getadsassetrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.GetAdsAssetResponse](../../pkg/models/operations/getadsassetresponse.md), error**

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

    res, err := s.Ads.GetAdsCampaign(ctx, operations.GetAdsCampaignRequest{
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

## GetAdsCreative

Retrieve a creative

### Example Usage

<!-- UsageSnippet language="go" operationID="getAdsCreative" method="get" path="/ads/{connection_id}/creative/{id}" -->
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

    res, err := s.Ads.GetAdsCreative(ctx, operations.GetAdsCreativeRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsCreative != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetAdsCreativeRequest](../../pkg/models/operations/getadscreativerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetAdsCreativeResponse](../../pkg/models/operations/getadscreativeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAdsGroup

Retrieve a group

### Example Usage

<!-- UsageSnippet language="go" operationID="getAdsGroup" method="get" path="/ads/{connection_id}/group/{id}" -->
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

    res, err := s.Ads.GetAdsGroup(ctx, operations.GetAdsGroupRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetAdsGroupRequest](../../pkg/models/operations/getadsgrouprequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.GetAdsGroupResponse](../../pkg/models/operations/getadsgroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAdsInsertionorder

Retrieve an insertionorder

### Example Usage

<!-- UsageSnippet language="go" operationID="getAdsInsertionorder" method="get" path="/ads/{connection_id}/insertionorder/{id}" -->
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

    res, err := s.Ads.GetAdsInsertionorder(ctx, operations.GetAdsInsertionorderRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsInsertionorder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetAdsInsertionorderRequest](../../pkg/models/operations/getadsinsertionorderrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetAdsInsertionorderResponse](../../pkg/models/operations/getadsinsertionorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAdsOrganization

Retrieve an organization

### Example Usage

<!-- UsageSnippet language="go" operationID="getAdsOrganization" method="get" path="/ads/{connection_id}/organization/{id}" -->
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

    res, err := s.Ads.GetAdsOrganization(ctx, operations.GetAdsOrganizationRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsOrganization != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetAdsOrganizationRequest](../../pkg/models/operations/getadsorganizationrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetAdsOrganizationResponse](../../pkg/models/operations/getadsorganizationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAdsPromoted

Retrieve a promoted

### Example Usage

<!-- UsageSnippet language="go" operationID="getAdsPromoted" method="get" path="/ads/{connection_id}/promoted/{id}" -->
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

    res, err := s.Ads.GetAdsPromoted(ctx, operations.GetAdsPromotedRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsPromoted != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetAdsPromotedRequest](../../pkg/models/operations/getadspromotedrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetAdsPromotedResponse](../../pkg/models/operations/getadspromotedresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAdsTarget

Retrieve a target

### Example Usage

<!-- UsageSnippet language="go" operationID="getAdsTarget" method="get" path="/ads/{connection_id}/target/{id}" -->
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

    res, err := s.Ads.GetAdsTarget(ctx, operations.GetAdsTargetRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsTarget != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetAdsTargetRequest](../../pkg/models/operations/getadstargetrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.GetAdsTargetResponse](../../pkg/models/operations/getadstargetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAdsAds

List all ads

### Example Usage

<!-- UsageSnippet language="go" operationID="listAdsAds" method="get" path="/ads/{connection_id}/ad" -->
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

    res, err := s.Ads.ListAdsAds(ctx, operations.ListAdsAdsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsAds != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListAdsAdsRequest](../../pkg/models/operations/listadsadsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.ListAdsAdsResponse](../../pkg/models/operations/listadsadsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAdsAssets

List all assets

### Example Usage

<!-- UsageSnippet language="go" operationID="listAdsAssets" method="get" path="/ads/{connection_id}/asset" -->
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

    res, err := s.Ads.ListAdsAssets(ctx, operations.ListAdsAssetsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsAssets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListAdsAssetsRequest](../../pkg/models/operations/listadsassetsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ListAdsAssetsResponse](../../pkg/models/operations/listadsassetsresponse.md), error**

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

    res, err := s.Ads.ListAdsCampaigns(ctx, operations.ListAdsCampaignsRequest{
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

## ListAdsCreatives

List all creatives

### Example Usage

<!-- UsageSnippet language="go" operationID="listAdsCreatives" method="get" path="/ads/{connection_id}/creative" -->
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

    res, err := s.Ads.ListAdsCreatives(ctx, operations.ListAdsCreativesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsCreatives != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListAdsCreativesRequest](../../pkg/models/operations/listadscreativesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListAdsCreativesResponse](../../pkg/models/operations/listadscreativesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAdsGroups

List all groups

### Example Usage

<!-- UsageSnippet language="go" operationID="listAdsGroups" method="get" path="/ads/{connection_id}/group" -->
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

    res, err := s.Ads.ListAdsGroups(ctx, operations.ListAdsGroupsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsGroups != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListAdsGroupsRequest](../../pkg/models/operations/listadsgroupsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ListAdsGroupsResponse](../../pkg/models/operations/listadsgroupsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAdsInsertionorders

List all insertionorders

### Example Usage

<!-- UsageSnippet language="go" operationID="listAdsInsertionorders" method="get" path="/ads/{connection_id}/insertionorder" -->
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

    res, err := s.Ads.ListAdsInsertionorders(ctx, operations.ListAdsInsertionordersRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsInsertionorders != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListAdsInsertionordersRequest](../../pkg/models/operations/listadsinsertionordersrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.ListAdsInsertionordersResponse](../../pkg/models/operations/listadsinsertionordersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAdsOrganizations

List all organizations

### Example Usage

<!-- UsageSnippet language="go" operationID="listAdsOrganizations" method="get" path="/ads/{connection_id}/organization" -->
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

    res, err := s.Ads.ListAdsOrganizations(ctx, operations.ListAdsOrganizationsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsOrganizations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListAdsOrganizationsRequest](../../pkg/models/operations/listadsorganizationsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListAdsOrganizationsResponse](../../pkg/models/operations/listadsorganizationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAdsPromoteds

List all promoteds

### Example Usage

<!-- UsageSnippet language="go" operationID="listAdsPromoteds" method="get" path="/ads/{connection_id}/promoted" -->
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

    res, err := s.Ads.ListAdsPromoteds(ctx, operations.ListAdsPromotedsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsPromoteds != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListAdsPromotedsRequest](../../pkg/models/operations/listadspromotedsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListAdsPromotedsResponse](../../pkg/models/operations/listadspromotedsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAdsReports

List all reports

### Example Usage

<!-- UsageSnippet language="go" operationID="listAdsReports" method="get" path="/ads/{connection_id}/report" -->
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

    res, err := s.Ads.ListAdsReports(ctx, operations.ListAdsReportsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsReports != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListAdsReportsRequest](../../pkg/models/operations/listadsreportsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ListAdsReportsResponse](../../pkg/models/operations/listadsreportsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAdsTargets

List all targets

### Example Usage

<!-- UsageSnippet language="go" operationID="listAdsTargets" method="get" path="/ads/{connection_id}/target" -->
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

    res, err := s.Ads.ListAdsTargets(ctx, operations.ListAdsTargetsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsTargets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListAdsTargetsRequest](../../pkg/models/operations/listadstargetsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ListAdsTargetsResponse](../../pkg/models/operations/listadstargetsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAdsAd

Update an ad

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAdsAd" method="patch" path="/ads/{connection_id}/ad/{id}" example="ads_ad" -->
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

    res, err := s.Ads.PatchAdsAd(ctx, operations.PatchAdsAdRequest{
        AdsAd: shared.AdsAd{
            AdCopy: unifiedgosdk.Pointer("Ascisco tolero caute sapiente. Valens unde comedo cursus crinis nobis thema. Cohaero nisi ullam tum unde ultio vilicus auditor capio."),
            AdType: shared.AdTypeSocial.ToPointer(),
            AdvertiserName: unifiedgosdk.Pointer("Robel, Nader and Rau"),
            CreatedAt: types.MustNewTimeFromString("2022-11-08T03:38:20.978Z"),
            CreativeAssetURL: unifiedgosdk.Pointer("https://picsum.photos/seed/LwOzrpr9/948/2793"),
            Description: unifiedgosdk.Pointer("Accedo vespillo carpo dolor decet stillicidium comptus tenuis."),
            FinalURL: unifiedgosdk.Pointer("https://improbable-sanity.com"),
            ID: unifiedgosdk.Pointer("8c3e2ac1-38e0-499b-9b46-12e44c25b14c"),
            Name: unifiedgosdk.Pointer("Hermiston Group"),
            Status: shared.AdsAdStatusArchived.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-06-05T03:02:50.909Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsAd != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.PatchAdsAdRequest](../../pkg/models/operations/patchadsadrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.PatchAdsAdResponse](../../pkg/models/operations/patchadsadresponse.md), error**

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

    res, err := s.Ads.PatchAdsCampaign(ctx, operations.PatchAdsCampaignRequest{
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

## PatchAdsCreative

Update a creative

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAdsCreative" method="patch" path="/ads/{connection_id}/creative/{id}" example="ads_creative" -->
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

    res, err := s.Ads.PatchAdsCreative(ctx, operations.PatchAdsCreativeRequest{
        AdsCreative: shared.AdsCreative{
            CreatedAt: types.MustNewTimeFromString("2020-02-17T11:24:51.093Z"),
            ID: unifiedgosdk.Pointer("fba66c03-d68d-4007-8ae9-cf369d858311"),
            Labels: []string{
                "coma",
                "accedo",
                "termes",
            },
            Name: unifiedgosdk.Pointer("Brekke, Bradtke and Robel"),
            Status: shared.AdsCreativeStatusPaused.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2021-06-21T01:13:41.798Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsCreative != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchAdsCreativeRequest](../../pkg/models/operations/patchadscreativerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchAdsCreativeResponse](../../pkg/models/operations/patchadscreativeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAdsGroup

Update a group

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAdsGroup" method="patch" path="/ads/{connection_id}/group/{id}" example="ads_group" -->
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

    res, err := s.Ads.PatchAdsGroup(ctx, operations.PatchAdsGroupRequest{
        AdsGroup: shared.AdsGroup{
            BidAmount: unifiedgosdk.Pointer[float64](26.16030164062977),
            BudgetAmount: unifiedgosdk.Pointer[float64](5099.175239447504),
            BudgetPeriod: shared.AdsGroupBudgetPeriodMonthly.ToPointer(),
            CreatedAt: types.MustNewTimeFromString("2019-08-29T17:59:41.045Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            EffectiveStatus: shared.AdsGroupEffectiveStatusPaused.ToPointer(),
            EndAt: types.MustNewTimeFromString("2026-05-24T14:15:58.401Z"),
            ID: unifiedgosdk.Pointer("4e613cf5-3cb2-4316-a124-d5d7b5cb0cce"),
            LanguageLocale: unifiedgosdk.Pointer("fr-FR"),
            Name: unifiedgosdk.Pointer("Stark - Baumbach"),
            StartAt: types.MustNewTimeFromString("2025-12-10T22:04:10.775Z"),
            Status: shared.AdsGroupStatusProcessing.ToPointer(),
            Targeting: &shared.PropertyAdsGroupTargeting{},
            UpdatedAt: types.MustNewTimeFromString("2022-01-02T17:05:47.254Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.PatchAdsGroupRequest](../../pkg/models/operations/patchadsgrouprequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.PatchAdsGroupResponse](../../pkg/models/operations/patchadsgroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAdsInsertionorder

Update an insertionorder

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAdsInsertionorder" method="patch" path="/ads/{connection_id}/insertionorder/{id}" example="ads_insertionorder" -->
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

    res, err := s.Ads.PatchAdsInsertionorder(ctx, operations.PatchAdsInsertionorderRequest{
        AdsInsertionorder: shared.AdsInsertionorder{
            CreatedAt: types.MustNewTimeFromString("2021-04-10T06:57:36.611Z"),
            ID: unifiedgosdk.Pointer("658ad6f8-cf35-434b-9652-34e030c35b64"),
            Name: unifiedgosdk.Pointer("Kunde, Smith and Reinger"),
            Status: shared.AdsInsertionorderStatusUnspecified.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2021-04-28T12:31:55.184Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsInsertionorder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PatchAdsInsertionorderRequest](../../pkg/models/operations/patchadsinsertionorderrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.PatchAdsInsertionorderResponse](../../pkg/models/operations/patchadsinsertionorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAdsOrganization

Update an organization

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAdsOrganization" method="patch" path="/ads/{connection_id}/organization/{id}" example="ads_organization" -->
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

    res, err := s.Ads.PatchAdsOrganization(ctx, operations.PatchAdsOrganizationRequest{
        AdsOrganization: shared.AdsOrganization{
            AccountNumber: unifiedgosdk.Pointer("LQUJx8zQBW"),
            CreatedAt: types.MustNewTimeFromString("2020-07-23T21:47:11.440Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            ID: unifiedgosdk.Pointer("1e7c6471-0081-404e-9c97-a7cd17942a76"),
            Managers: []shared.AdsManager{
                shared.AdsManager{
                    ID: unifiedgosdk.Pointer("e4fd87df-9f8b-4fa0-a77b-b7d18669e350"),
                    Name: unifiedgosdk.Pointer("Parker, Leannon and Gibson"),
                },
            },
            Name: unifiedgosdk.Pointer("Ankunding Inc"),
            Status: shared.AdsOrganizationStatusProcessing.ToPointer(),
            Timezone: unifiedgosdk.Pointer("Europe/Chisinau"),
            UpdatedAt: types.MustNewTimeFromString("2026-02-27T03:26:08.216Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsOrganization != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PatchAdsOrganizationRequest](../../pkg/models/operations/patchadsorganizationrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.PatchAdsOrganizationResponse](../../pkg/models/operations/patchadsorganizationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAdsAd

Remove an ad

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAdsAd" method="delete" path="/ads/{connection_id}/ad/{id}" -->
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

    res, err := s.Ads.RemoveAdsAd(ctx, operations.RemoveAdsAdRequest{
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.RemoveAdsAdRequest](../../pkg/models/operations/removeadsadrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.RemoveAdsAdResponse](../../pkg/models/operations/removeadsadresponse.md), error**

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

    res, err := s.Ads.RemoveAdsCampaign(ctx, operations.RemoveAdsCampaignRequest{
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

## RemoveAdsCreative

Remove a creative

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAdsCreative" method="delete" path="/ads/{connection_id}/creative/{id}" -->
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

    res, err := s.Ads.RemoveAdsCreative(ctx, operations.RemoveAdsCreativeRequest{
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
| `request`                                                                                      | [operations.RemoveAdsCreativeRequest](../../pkg/models/operations/removeadscreativerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemoveAdsCreativeResponse](../../pkg/models/operations/removeadscreativeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAdsGroup

Remove a group

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAdsGroup" method="delete" path="/ads/{connection_id}/group/{id}" -->
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

    res, err := s.Ads.RemoveAdsGroup(ctx, operations.RemoveAdsGroupRequest{
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.RemoveAdsGroupRequest](../../pkg/models/operations/removeadsgrouprequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.RemoveAdsGroupResponse](../../pkg/models/operations/removeadsgroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAdsInsertionorder

Remove an insertionorder

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAdsInsertionorder" method="delete" path="/ads/{connection_id}/insertionorder/{id}" -->
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

    res, err := s.Ads.RemoveAdsInsertionorder(ctx, operations.RemoveAdsInsertionorderRequest{
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
| `request`                                                                                                  | [operations.RemoveAdsInsertionorderRequest](../../pkg/models/operations/removeadsinsertionorderrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.RemoveAdsInsertionorderResponse](../../pkg/models/operations/removeadsinsertionorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAdsOrganization

Remove an organization

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAdsOrganization" method="delete" path="/ads/{connection_id}/organization/{id}" -->
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

    res, err := s.Ads.RemoveAdsOrganization(ctx, operations.RemoveAdsOrganizationRequest{
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
| `request`                                                                                              | [operations.RemoveAdsOrganizationRequest](../../pkg/models/operations/removeadsorganizationrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.RemoveAdsOrganizationResponse](../../pkg/models/operations/removeadsorganizationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAdsAd

Update an ad

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAdsAd" method="put" path="/ads/{connection_id}/ad/{id}" example="ads_ad" -->
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

    res, err := s.Ads.UpdateAdsAd(ctx, operations.UpdateAdsAdRequest{
        AdsAd: shared.AdsAd{
            AdCopy: unifiedgosdk.Pointer("Ascisco tolero caute sapiente. Valens unde comedo cursus crinis nobis thema. Cohaero nisi ullam tum unde ultio vilicus auditor capio."),
            AdType: shared.AdTypeSocial.ToPointer(),
            AdvertiserName: unifiedgosdk.Pointer("Robel, Nader and Rau"),
            CreatedAt: types.MustNewTimeFromString("2022-11-08T03:38:20.978Z"),
            CreativeAssetURL: unifiedgosdk.Pointer("https://picsum.photos/seed/LwOzrpr9/948/2793"),
            Description: unifiedgosdk.Pointer("Accedo vespillo carpo dolor decet stillicidium comptus tenuis."),
            FinalURL: unifiedgosdk.Pointer("https://improbable-sanity.com"),
            ID: unifiedgosdk.Pointer("8c3e2ac1-38e0-499b-9b46-12e44c25b14c"),
            Name: unifiedgosdk.Pointer("Hermiston Group"),
            Status: shared.AdsAdStatusArchived.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-06-05T03:02:50.909Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsAd != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateAdsAdRequest](../../pkg/models/operations/updateadsadrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.UpdateAdsAdResponse](../../pkg/models/operations/updateadsadresponse.md), error**

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

    res, err := s.Ads.UpdateAdsCampaign(ctx, operations.UpdateAdsCampaignRequest{
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

## UpdateAdsCreative

Update a creative

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAdsCreative" method="put" path="/ads/{connection_id}/creative/{id}" example="ads_creative" -->
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

    res, err := s.Ads.UpdateAdsCreative(ctx, operations.UpdateAdsCreativeRequest{
        AdsCreative: shared.AdsCreative{
            CreatedAt: types.MustNewTimeFromString("2020-02-17T11:24:51.093Z"),
            ID: unifiedgosdk.Pointer("fba66c03-d68d-4007-8ae9-cf369d858311"),
            Labels: []string{
                "coma",
                "accedo",
                "termes",
            },
            Name: unifiedgosdk.Pointer("Brekke, Bradtke and Robel"),
            Status: shared.AdsCreativeStatusPaused.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2021-06-21T01:13:41.798Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsCreative != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateAdsCreativeRequest](../../pkg/models/operations/updateadscreativerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateAdsCreativeResponse](../../pkg/models/operations/updateadscreativeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAdsGroup

Update a group

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAdsGroup" method="put" path="/ads/{connection_id}/group/{id}" example="ads_group" -->
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

    res, err := s.Ads.UpdateAdsGroup(ctx, operations.UpdateAdsGroupRequest{
        AdsGroup: shared.AdsGroup{
            BidAmount: unifiedgosdk.Pointer[float64](26.16030164062977),
            BudgetAmount: unifiedgosdk.Pointer[float64](5099.175239447504),
            BudgetPeriod: shared.AdsGroupBudgetPeriodMonthly.ToPointer(),
            CreatedAt: types.MustNewTimeFromString("2019-08-29T17:59:41.045Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            EffectiveStatus: shared.AdsGroupEffectiveStatusPaused.ToPointer(),
            EndAt: types.MustNewTimeFromString("2026-05-24T14:15:58.401Z"),
            ID: unifiedgosdk.Pointer("4e613cf5-3cb2-4316-a124-d5d7b5cb0cce"),
            LanguageLocale: unifiedgosdk.Pointer("fr-FR"),
            Name: unifiedgosdk.Pointer("Stark - Baumbach"),
            StartAt: types.MustNewTimeFromString("2025-12-10T22:04:10.775Z"),
            Status: shared.AdsGroupStatusProcessing.ToPointer(),
            Targeting: &shared.PropertyAdsGroupTargeting{},
            UpdatedAt: types.MustNewTimeFromString("2022-01-02T17:05:47.254Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateAdsGroupRequest](../../pkg/models/operations/updateadsgrouprequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.UpdateAdsGroupResponse](../../pkg/models/operations/updateadsgroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAdsInsertionorder

Update an insertionorder

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAdsInsertionorder" method="put" path="/ads/{connection_id}/insertionorder/{id}" example="ads_insertionorder" -->
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

    res, err := s.Ads.UpdateAdsInsertionorder(ctx, operations.UpdateAdsInsertionorderRequest{
        AdsInsertionorder: shared.AdsInsertionorder{
            CreatedAt: types.MustNewTimeFromString("2021-04-10T06:57:36.611Z"),
            ID: unifiedgosdk.Pointer("658ad6f8-cf35-434b-9652-34e030c35b64"),
            Name: unifiedgosdk.Pointer("Kunde, Smith and Reinger"),
            Status: shared.AdsInsertionorderStatusUnspecified.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2021-04-28T12:31:55.184Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsInsertionorder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.UpdateAdsInsertionorderRequest](../../pkg/models/operations/updateadsinsertionorderrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.UpdateAdsInsertionorderResponse](../../pkg/models/operations/updateadsinsertionorderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAdsOrganization

Update an organization

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAdsOrganization" method="put" path="/ads/{connection_id}/organization/{id}" example="ads_organization" -->
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

    res, err := s.Ads.UpdateAdsOrganization(ctx, operations.UpdateAdsOrganizationRequest{
        AdsOrganization: shared.AdsOrganization{
            AccountNumber: unifiedgosdk.Pointer("LQUJx8zQBW"),
            CreatedAt: types.MustNewTimeFromString("2020-07-23T21:47:11.440Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            ID: unifiedgosdk.Pointer("1e7c6471-0081-404e-9c97-a7cd17942a76"),
            Managers: []shared.AdsManager{
                shared.AdsManager{
                    ID: unifiedgosdk.Pointer("e4fd87df-9f8b-4fa0-a77b-b7d18669e350"),
                    Name: unifiedgosdk.Pointer("Parker, Leannon and Gibson"),
                },
            },
            Name: unifiedgosdk.Pointer("Ankunding Inc"),
            Status: shared.AdsOrganizationStatusProcessing.ToPointer(),
            Timezone: unifiedgosdk.Pointer("Europe/Chisinau"),
            UpdatedAt: types.MustNewTimeFromString("2026-02-27T03:26:08.216Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AdsOrganization != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdateAdsOrganizationRequest](../../pkg/models/operations/updateadsorganizationrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UpdateAdsOrganizationResponse](../../pkg/models/operations/updateadsorganizationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |