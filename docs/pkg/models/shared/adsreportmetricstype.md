# AdsReportMetricsType

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AdsReportMetricsTypeClicks

// Open enum: custom values can be created with a direct type cast
custom := shared.AdsReportMetricsType("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `AdsReportMetricsTypeClicks`                           | CLICKS                                                 |
| `AdsReportMetricsTypeImpressions`                      | IMPRESSIONS                                            |
| `AdsReportMetricsTypeConversions`                      | CONVERSIONS                                            |
| `AdsReportMetricsTypeCost`                             | COST                                                   |
| `AdsReportMetricsTypeCtr`                              | CTR                                                    |
| `AdsReportMetricsTypeCpc`                              | CPC                                                    |
| `AdsReportMetricsTypeConversionValue`                  | CONVERSION_VALUE                                       |
| `AdsReportMetricsTypeCpa`                              | CPA                                                    |
| `AdsReportMetricsTypeRoas`                             | ROAS                                                   |
| `AdsReportMetricsTypeCpm`                              | CPM                                                    |
| `AdsReportMetricsTypeEcpm`                             | ECPM                                                   |
| `AdsReportMetricsTypeEngagement`                       | ENGAGEMENT                                             |
| `AdsReportMetricsTypeVideoCompletions`                 | VIDEO_COMPLETIONS                                      |
| `AdsReportMetricsTypeVideoViews`                       | VIDEO_VIEWS                                            |
| `AdsReportMetricsTypeLeads`                            | LEADS                                                  |
| `AdsReportMetricsTypeEngagements`                      | ENGAGEMENTS                                            |
| `AdsReportMetricsTypeSaves`                            | SAVES                                                  |
| `AdsReportMetricsTypeLikes`                            | LIKES                                                  |
| `AdsReportMetricsTypeShares`                           | SHARES                                                 |
| `AdsReportMetricsTypeComments`                         | COMMENTS                                               |
| `AdsReportMetricsTypeFollows`                          | FOLLOWS                                                |
| `AdsReportMetricsTypePostClickConversions`             | POST_CLICK_CONVERSIONS                                 |
| `AdsReportMetricsTypeViewThroughConversions`           | VIEW_THROUGH_CONVERSIONS                               |
| `AdsReportMetricsTypeAllConversions`                   | ALL_CONVERSIONS                                        |
| `AdsReportMetricsTypeAllConversionValue`               | ALL_CONVERSION_VALUE                                   |
| `AdsReportMetricsTypeInteractions`                     | INTERACTIONS                                           |
| `AdsReportMetricsTypeRevenue`                          | REVENUE                                                |
| `AdsReportMetricsTypeMediaCost`                        | MEDIA_COST                                             |
| `AdsReportMetricsTypeTotalMediaCost`                   | TOTAL_MEDIA_COST                                       |
| `AdsReportMetricsTypeVideoPlays`                       | VIDEO_PLAYS                                            |
| `AdsReportMetricsTypeVideoQuartile25`                  | VIDEO_QUARTILE_25                                      |
| `AdsReportMetricsTypeVideoQuartile50`                  | VIDEO_QUARTILE_50                                      |
| `AdsReportMetricsTypeVideoQuartile75`                  | VIDEO_QUARTILE_75                                      |
| `AdsReportMetricsTypeVideoQuartile100`                 | VIDEO_QUARTILE_100                                     |
| `AdsReportMetricsTypeVideoAvgTimeWatched`              | VIDEO_AVG_TIME_WATCHED                                 |
| `AdsReportMetricsTypeVideoThruplay`                    | VIDEO_THRUPLAY                                         |
| `AdsReportMetricsTypeViewableImpressions`              | VIEWABLE_IMPRESSIONS                                   |
| `AdsReportMetricsTypeMeasurableImpressions`            | MEASURABLE_IMPRESSIONS                                 |
| `AdsReportMetricsTypeViewabilityRate`                  | VIEWABILITY_RATE                                       |
| `AdsReportMetricsTypeBillableImpressions`              | BILLABLE_IMPRESSIONS                                   |
| `AdsReportMetricsTypeDataFees`                         | DATA_FEES                                              |
| `AdsReportMetricsTypePlatformFees`                     | PLATFORM_FEES                                          |
| `AdsReportMetricsTypeProfit`                           | PROFIT                                                 |
| `AdsReportMetricsTypeLandingPageClicks`                | LANDING_PAGE_CLICKS                                    |
| `AdsReportMetricsTypeLeadFormOpens`                    | LEAD_FORM_OPENS                                        |
| `AdsReportMetricsTypeAdUnitClicks`                     | AD_UNIT_CLICKS                                         |
| `AdsReportMetricsTypeCardClicks`                       | CARD_CLICKS                                            |
| `AdsReportMetricsTypeCardImpressions`                  | CARD_IMPRESSIONS                                       |
| `AdsReportMetricsTypeCommentLikes`                     | COMMENT_LIKES                                          |
| `AdsReportMetricsTypeCompanyPageClicks`                | COMPANY_PAGE_CLICKS                                    |
| `AdsReportMetricsTypeActionClicks`                     | ACTION_CLICKS                                          |
| `AdsReportMetricsTypeTextURLClicks`                    | TEXT_URL_CLICKS                                        |
| `AdsReportMetricsTypeOtherEngagements`                 | OTHER_ENGAGEMENTS                                      |
| `AdsReportMetricsTypeOpens`                            | OPENS                                                  |
| `AdsReportMetricsTypeTotalEngagements`                 | TOTAL_ENGAGEMENTS                                      |
| `AdsReportMetricsTypeUniqueImpressions`                | UNIQUE_IMPRESSIONS                                     |
| `AdsReportMetricsTypeUniqueClicks`                     | UNIQUE_CLICKS                                          |
| `AdsReportMetricsTypeViralImpressions`                 | VIRAL_IMPRESSIONS                                      |
| `AdsReportMetricsTypeViralClicks`                      | VIRAL_CLICKS                                           |
| `AdsReportMetricsTypeViralLikes`                       | VIRAL_LIKES                                            |
| `AdsReportMetricsTypeViralComments`                    | VIRAL_COMMENTS                                         |
| `AdsReportMetricsTypeViralShares`                      | VIRAL_SHARES                                           |
| `AdsReportMetricsTypeViralFollows`                     | VIRAL_FOLLOWS                                          |
| `AdsReportMetricsTypeViralVideoPlays`                  | VIRAL_VIDEO_PLAYS                                      |
| `AdsReportMetricsTypeViralVideoViews`                  | VIRAL_VIDEO_VIEWS                                      |
| `AdsReportMetricsTypeViralVideoCompletions`            | VIRAL_VIDEO_COMPLETIONS                                |
| `AdsReportMetricsTypeViralVideoQuartile25`             | VIRAL_VIDEO_QUARTILE_25                                |
| `AdsReportMetricsTypeViralVideoQuartile50`             | VIRAL_VIDEO_QUARTILE_50                                |
| `AdsReportMetricsTypeViralVideoQuartile75`             | VIRAL_VIDEO_QUARTILE_75                                |
| `AdsReportMetricsTypeViralLeads`                       | VIRAL_LEADS                                            |
| `AdsReportMetricsTypeViralLeadFormOpens`               | VIRAL_LEAD_FORM_OPENS                                  |
| `AdsReportMetricsTypeViralLandingPageClicks`           | VIRAL_LANDING_PAGE_CLICKS                              |
| `AdsReportMetricsTypeViralConversions`                 | VIRAL_CONVERSIONS                                      |
| `AdsReportMetricsTypeViralPostClickConversions`        | VIRAL_POST_CLICK_CONVERSIONS                           |
| `AdsReportMetricsTypeViralViewThroughConversions`      | VIRAL_VIEW_THROUGH_CONVERSIONS                         |
| `AdsReportMetricsTypeViralEngagements`                 | VIRAL_ENGAGEMENTS                                      |
| `AdsReportMetricsTypeGmailSecondaryClicks`             | GMAIL_SECONDARY_CLICKS                                 |
| `AdsReportMetricsTypeAverageCpv`                       | AVERAGE_CPV                                            |
| `AdsReportMetricsTypeVideoViewsFromSearch`             | VIDEO_VIEWS_FROM_SEARCH                                |
| `AdsReportMetricsTypeCrossDeviceConversions`           | CROSS_DEVICE_CONVERSIONS                               |
| `AdsReportMetricsTypeAbsoluteTopImpressionShare`       | ABSOLUTE_TOP_IMPRESSION_SHARE                          |
| `AdsReportMetricsTypeTopImpressionShare`               | TOP_IMPRESSION_SHARE                                   |
| `AdsReportMetricsTypeAbsoluteTopImpressionRatePercent` | ABSOLUTE_TOP_IMPRESSION_RATE_PERCENT                   |
| `AdsReportMetricsTypeTopImpressionRatePercent`         | TOP_IMPRESSION_RATE_PERCENT                            |
| `AdsReportMetricsTypeVideoFullscreens`                 | VIDEO_FULLSCREENS                                      |
| `AdsReportMetricsTypeVideoPauses`                      | VIDEO_PAUSES                                           |
| `AdsReportMetricsTypeVideoMutes`                       | VIDEO_MUTES                                            |
| `AdsReportMetricsTypeVideoSkips`                       | VIDEO_SKIPS                                            |
| `AdsReportMetricsTypeCompanionClicks`                  | COMPANION_CLICKS                                       |
| `AdsReportMetricsTypeCompanionViews`                   | COMPANION_VIEWS                                        |
| `AdsReportMetricsTypeActiveViewAvgTime`                | ACTIVE_VIEW_AVG_TIME                                   |
| `AdsReportMetricsTypeEligibleImpressions`              | ELIGIBLE_IMPRESSIONS                                   |
| `AdsReportMetricsTypeEarnedViews`                      | EARNED_VIEWS                                           |
| `AdsReportMetricsTypeUniqueViewers`                    | UNIQUE_VIEWERS                                         |
| `AdsReportMetricsTypeCostUsd`                          | COST_USD                                               |
| `AdsReportMetricsTypeViralCardClicks`                  | VIRAL_CARD_CLICKS                                      |
| `AdsReportMetricsTypeViralCardImpressions`             | VIRAL_CARD_IMPRESSIONS                                 |
| `AdsReportMetricsTypeViralCompanyPageClicks`           | VIRAL_COMPANY_PAGE_CLICKS                              |
| `AdsReportMetricsTypeViralVideoFullscreens`            | VIRAL_VIDEO_FULLSCREENS                                |
| `AdsReportMetricsTypeViralOtherEngagements`            | VIRAL_OTHER_ENGAGEMENTS                                |
| `AdsReportMetricsTypeLeadGenContactShares`             | LEAD_GEN_CONTACT_SHARES                                |
| `AdsReportMetricsTypeLeadGenInterestedClicks`          | LEAD_GEN_INTERESTED_CLICKS                             |
| `AdsReportMetricsTypeCm360PostClickRevenue`            | CM360_POST_CLICK_REVENUE                               |
| `AdsReportMetricsTypeCm360PostViewRevenue`             | CM360_POST_VIEW_REVENUE                                |
| `AdsReportMetricsTypeScrolls`                          | SCROLLS                                                |
| `AdsReportMetricsTypeMediaFee1`                        | MEDIA_FEE_1                                            |
| `AdsReportMetricsTypeMediaFee2`                        | MEDIA_FEE_2                                            |
| `AdsReportMetricsTypeMediaFee3`                        | MEDIA_FEE_3                                            |
| `AdsReportMetricsTypeMediaFee4`                        | MEDIA_FEE_4                                            |
| `AdsReportMetricsTypeMediaFee5`                        | MEDIA_FEE_5                                            |
| `AdsReportMetricsTypeActiveViewDistUnmeasurable`       | ACTIVE_VIEW_DIST_UNMEASURABLE                          |
| `AdsReportMetricsTypeActiveViewDistUnviewable`         | ACTIVE_VIEW_DIST_UNVIEWABLE                            |
| `AdsReportMetricsTypeActiveViewDistViewable`           | ACTIVE_VIEW_DIST_VIEWABLE                              |
| `AdsReportMetricsTypeActiveViewAudibleVisibleComplete` | ACTIVE_VIEW_AUDIBLE_VISIBLE_COMPLETE                   |
| `AdsReportMetricsTypeActiveViewVisible10S`             | ACTIVE_VIEW_VISIBLE_10S                                |
| `AdsReportMetricsTypeNotMeasurableImpressions`         | NOT_MEASURABLE_IMPRESSIONS                             |
| `AdsReportMetricsTypeNotViewableImpressions`           | NOT_VIEWABLE_IMPRESSIONS                               |
| `AdsReportMetricsTypeOneDView`                         | 1D_VIEW                                                |
| `AdsReportMetricsTypeOneDClick`                        | 1D_CLICK                                               |
| `AdsReportMetricsTypeSevenDView`                       | 7D_VIEW                                                |
| `AdsReportMetricsTypeSevenDClick`                      | 7D_CLICK                                               |
| `AdsReportMetricsTypeFourteenDClick`                   | 14D_CLICK                                              |
| `AdsReportMetricsTypeTwentyEightDView`                 | 28D_VIEW                                               |
| `AdsReportMetricsTypeTwentyEightDClick`                | 28D_CLICK                                              |
| `AdsReportMetricsTypeThirtyDClick`                     | 30D_CLICK                                              |