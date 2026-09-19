# Calendar

## Overview

### Available Operations

* [CreateCalendarCalendar](#createcalendarcalendar) - Create a calendar
* [CreateCalendarEvent](#createcalendarevent) - Create an event
* [CreateCalendarLink](#createcalendarlink) - Create a link
* [CreateCalendarWebinar](#createcalendarwebinar) - Create a webinar
* [GetCalendarCalendar](#getcalendarcalendar) - Retrieve a calendar
* [GetCalendarEvent](#getcalendarevent) - Retrieve an event
* [GetCalendarLink](#getcalendarlink) - Retrieve a link
* [GetCalendarRecording](#getcalendarrecording) - Retrieve a recording
* [GetCalendarWebinar](#getcalendarwebinar) - Retrieve a webinar
* [ListCalendarBusies](#listcalendarbusies) - List all busies
* [ListCalendarCalendars](#listcalendarcalendars) - List all calendars
* [ListCalendarEvents](#listcalendarevents) - List all events
* [ListCalendarLinks](#listcalendarlinks) - List all links
* [ListCalendarRecordings](#listcalendarrecordings) - List all recordings
* [ListCalendarWebinars](#listcalendarwebinars) - List all webinars
* [PatchCalendarCalendar](#patchcalendarcalendar) - Update a calendar
* [PatchCalendarEvent](#patchcalendarevent) - Update an event
* [PatchCalendarLink](#patchcalendarlink) - Update a link
* [PatchCalendarWebinar](#patchcalendarwebinar) - Update a webinar
* [RemoveCalendarCalendar](#removecalendarcalendar) - Remove a calendar
* [RemoveCalendarEvent](#removecalendarevent) - Remove an event
* [RemoveCalendarLink](#removecalendarlink) - Remove a link
* [RemoveCalendarWebinar](#removecalendarwebinar) - Remove a webinar
* [UpdateCalendarCalendar](#updatecalendarcalendar) - Update a calendar
* [UpdateCalendarEvent](#updatecalendarevent) - Update an event
* [UpdateCalendarLink](#updatecalendarlink) - Update a link
* [UpdateCalendarWebinar](#updatecalendarwebinar) - Update a webinar

## CreateCalendarCalendar

Create a calendar

### Example Usage

<!-- UsageSnippet language="go" operationID="createCalendarCalendar" method="post" path="/calendar/{connection_id}/calendar" example="calendar_calendar" -->
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

    res, err := s.Calendar.CreateCalendarCalendar(ctx, operations.CreateCalendarCalendarRequest{
        CalendarCalendar: shared.CalendarCalendar{
            CreatedAt: types.MustNewTimeFromString("2020-01-09T23:11:34.147Z"),
            Description: unifiedgosdk.Pointer("Socius catena auxilium."),
            ID: unifiedgosdk.Pointer("bddcb983-4c88-4342-a300-e5a1165d1624"),
            IsPrimary: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("Acer supra vallum suasoria thesaurus omnis condico cognomen accendo vehemens."),
            Timezone: unifiedgosdk.Pointer("America/Dawson_Creek"),
            UpdatedAt: types.MustNewTimeFromString("2023-03-12T22:44:44.996Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CalendarCalendar != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.CreateCalendarCalendarRequest](../../pkg/models/operations/createcalendarcalendarrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.CreateCalendarCalendarResponse](../../pkg/models/operations/createcalendarcalendarresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateCalendarEvent

Create an event

### Example Usage

<!-- UsageSnippet language="go" operationID="createCalendarEvent" method="post" path="/calendar/{connection_id}/event" example="calendar_event" -->
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

    res, err := s.Calendar.CreateCalendarEvent(ctx, operations.CreateCalendarEventRequest{
        CalendarEvent: shared.CalendarEvent{
            Attachments: []shared.CalendarAttachment{},
            Conference: []shared.CalendarConference{},
            CreatedAt: unifiedgosdk.Pointer("2019-08-04T14:33:51.814Z"),
            EndAt: unifiedgosdk.Pointer("2020-05-20T17:12:37.976Z"),
            ID: unifiedgosdk.Pointer("dbfb8859-7c40-48b3-8e82-4d75514c5c87"),
            IsAllDay: unifiedgosdk.Pointer(false),
            IsFree: unifiedgosdk.Pointer(false),
            IsPrivate: unifiedgosdk.Pointer(false),
            Location: unifiedgosdk.Pointer("621 Boehm Prairie"),
            Notes: unifiedgosdk.Pointer("Aegre traho."),
            Recurrence: []shared.CalendarEventRecurrence{
                shared.CalendarEventRecurrence{
                    Count: unifiedgosdk.Pointer[float64](8.0),
                    EndAt: types.MustNewTimeFromString("2025-12-26T06:34:54.337Z"),
                    ExcludedDates: []string{
                        "2025-09-30T22:42:44.909Z",
                        "2023-10-09T12:14:01.306Z",
                        "2024-02-15T13:20:24.716Z",
                    },
                    Frequency: shared.CalendarEventRecurrenceFrequencyMonthly.ToPointer(),
                    IncludedDates: []string{
                        "2021-02-16T22:04:02.460Z",
                    },
                    Interval: unifiedgosdk.Pointer[float64](4.0),
                    OnDays: []shared.PropertyCalendarEventRecurrenceOnDays{
                        shared.PropertyCalendarEventRecurrenceOnDaysTh,
                        shared.PropertyCalendarEventRecurrenceOnDaysMo,
                        shared.PropertyCalendarEventRecurrenceOnDaysTh,
                    },
                    OnMonthDays: []float64{
                        -26.0,
                    },
                    OnMonths: []float64{
                        12.0,
                        9.0,
                        -1.0,
                        0.0,
                        1.0,
                        6.0,
                        -10.0,
                        9.0,
                        0.0,
                        4.0,
                        -2.0,
                    },
                    OnWeeks: []float64{
                        -7.0,
                        51.0,
                        -3.0,
                        -41.0,
                        15.0,
                        46.0,
                        -1.0,
                        46.0,
                        42.0,
                        11.0,
                        12.0,
                        -35.0,
                        -15.0,
                        -3.0,
                        -42.0,
                        50.0,
                        3.0,
                        -15.0,
                        -10.0,
                        6.0,
                        -53.0,
                        5.0,
                        -32.0,
                        -22.0,
                        43.0,
                        -44.0,
                        -23.0,
                        -21.0,
                        -18.0,
                    },
                    OnYearDays: []float64{
                        -35.0,
                        14.0,
                        -338.0,
                        175.0,
                        -87.0,
                        339.0,
                        341.0,
                        287.0,
                        -17.0,
                        319.0,
                        -3.0,
                        238.0,
                        -115.0,
                        -116.0,
                        283.0,
                        -61.0,
                        -254.0,
                        86.0,
                        -163.0,
                        5.0,
                        -171.0,
                        -99.0,
                        279.0,
                        19.0,
                        303.0,
                        -106.0,
                        90.0,
                        109.0,
                        -185.0,
                        -285.0,
                        -83.0,
                        -236.0,
                        66.0,
                        -215.0,
                        178.0,
                        64.0,
                        78.0,
                        5.0,
                        -251.0,
                        -79.0,
                        -271.0,
                        33.0,
                        320.0,
                        67.0,
                        -84.0,
                        -355.0,
                        -364.0,
                        348.0,
                        271.0,
                        -304.0,
                        -199.0,
                        106.0,
                        -345.0,
                        24.0,
                        -89.0,
                        -109.0,
                        -314.0,
                        365.0,
                        38.0,
                        -42.0,
                        123.0,
                        56.0,
                        -3.0,
                        31.0,
                        101.0,
                        326.0,
                        -160.0,
                        -101.0,
                        -267.0,
                        -309.0,
                        -363.0,
                        125.0,
                        -182.0,
                        363.0,
                        324.0,
                        36.0,
                        -269.0,
                        -79.0,
                        -60.0,
                        272.0,
                        -254.0,
                        -160.0,
                        -82.0,
                        19.0,
                        42.0,
                        69.0,
                        -104.0,
                        333.0,
                        236.0,
                        -287.0,
                        296.0,
                        261.0,
                        241.0,
                        348.0,
                        -72.0,
                        159.0,
                        -127.0,
                        229.0,
                        -158.0,
                        190.0,
                        -173.0,
                        -84.0,
                        -96.0,
                        176.0,
                        339.0,
                        -48.0,
                        287.0,
                        -46.0,
                        -101.0,
                        246.0,
                        -8.0,
                        -74.0,
                        338.0,
                        -51.0,
                        -42.0,
                        -128.0,
                        -169.0,
                        -174.0,
                        168.0,
                        -85.0,
                        37.0,
                        169.0,
                        -105.0,
                        231.0,
                        -250.0,
                        -286.0,
                        -7.0,
                        -121.0,
                        321.0,
                        278.0,
                        -120.0,
                        -96.0,
                        360.0,
                        337.0,
                        -258.0,
                        -179.0,
                        324.0,
                        -204.0,
                        327.0,
                        15.0,
                        365.0,
                        191.0,
                        -345.0,
                        -345.0,
                        56.0,
                        217.0,
                        60.0,
                        -264.0,
                        -248.0,
                        -316.0,
                        191.0,
                        -189.0,
                        -152.0,
                        -296.0,
                        194.0,
                        -42.0,
                        -21.0,
                        -218.0,
                        171.0,
                        -15.0,
                        301.0,
                        37.0,
                        -167.0,
                        18.0,
                        248.0,
                        -263.0,
                        27.0,
                        14.0,
                        59.0,
                        219.0,
                        -284.0,
                        221.0,
                        -76.0,
                        277.0,
                        183.0,
                        200.0,
                        -12.0,
                        -28.0,
                        -79.0,
                        150.0,
                        320.0,
                        -152.0,
                        -15.0,
                        -42.0,
                        -125.0,
                        -4.0,
                        269.0,
                        290.0,
                        52.0,
                        320.0,
                        344.0,
                        13.0,
                        -69.0,
                        255.0,
                        -154.0,
                        -281.0,
                        158.0,
                        25.0,
                        240.0,
                        -339.0,
                        96.0,
                        204.0,
                        324.0,
                        221.0,
                        37.0,
                        -333.0,
                        87.0,
                        354.0,
                        -365.0,
                        -203.0,
                        -341.0,
                        -79.0,
                        -208.0,
                        135.0,
                        132.0,
                        -351.0,
                        39.0,
                        -87.0,
                        -297.0,
                        -66.0,
                        346.0,
                        69.0,
                        -177.0,
                        235.0,
                        295.0,
                        -366.0,
                        -55.0,
                    },
                    Timezone: unifiedgosdk.Pointer("Asia/Ho_Chi_Minh"),
                    WeekStart: shared.WeekStartSu.ToPointer(),
                },
                shared.CalendarEventRecurrence{
                    Count: unifiedgosdk.Pointer[float64](9.0),
                    EndAt: types.MustNewTimeFromString("2025-04-30T05:52:39.868Z"),
                    ExcludedDates: []string{
                        "2020-04-29T00:43:21.543Z",
                    },
                    Frequency: shared.CalendarEventRecurrenceFrequencyDaily.ToPointer(),
                    IncludedDates: []string{
                        "2020-09-11T02:34:17.915Z",
                        "2021-11-29T01:27:53.237Z",
                        "2019-12-22T18:06:05.559Z",
                    },
                    Interval: unifiedgosdk.Pointer[float64](1.0),
                    OnDays: []shared.PropertyCalendarEventRecurrenceOnDays{
                        shared.PropertyCalendarEventRecurrenceOnDaysWe,
                        shared.PropertyCalendarEventRecurrenceOnDaysTu,
                        shared.PropertyCalendarEventRecurrenceOnDaysWe,
                        shared.PropertyCalendarEventRecurrenceOnDaysSa,
                        shared.PropertyCalendarEventRecurrenceOnDaysSa,
                        shared.PropertyCalendarEventRecurrenceOnDaysSa,
                    },
                    OnMonthDays: []float64{
                        1.0,
                    },
                    OnMonths: []float64{
                        4.0,
                        0.0,
                        -3.0,
                    },
                    OnWeeks: []float64{
                        -7.0,
                        -19.0,
                        50.0,
                        -37.0,
                        43.0,
                        -48.0,
                        -30.0,
                        34.0,
                        36.0,
                        -33.0,
                        24.0,
                        -4.0,
                    },
                    OnYearDays: []float64{
                        277.0,
                        -115.0,
                        100.0,
                        2.0,
                        81.0,
                        -66.0,
                        31.0,
                        -39.0,
                        -319.0,
                        -251.0,
                        -254.0,
                        -35.0,
                        -121.0,
                        262.0,
                        32.0,
                        190.0,
                        107.0,
                        -145.0,
                        91.0,
                        313.0,
                        -48.0,
                        277.0,
                        104.0,
                        342.0,
                        297.0,
                        -216.0,
                        346.0,
                        -257.0,
                        307.0,
                        -44.0,
                        264.0,
                        -153.0,
                        -268.0,
                        92.0,
                        152.0,
                        -182.0,
                        -334.0,
                        89.0,
                        343.0,
                        -320.0,
                        -36.0,
                        84.0,
                        340.0,
                        -88.0,
                        -278.0,
                        202.0,
                        291.0,
                        95.0,
                        -234.0,
                        -304.0,
                        -157.0,
                        -82.0,
                        -339.0,
                        83.0,
                        2.0,
                        -238.0,
                        -204.0,
                        206.0,
                        -273.0,
                        -78.0,
                        -21.0,
                        270.0,
                        -266.0,
                        -276.0,
                        154.0,
                        -97.0,
                        -43.0,
                        -3.0,
                        191.0,
                        -302.0,
                        290.0,
                        -118.0,
                        -125.0,
                        -294.0,
                        115.0,
                        -73.0,
                        -244.0,
                        127.0,
                        26.0,
                        251.0,
                        47.0,
                        -157.0,
                        22.0,
                        -361.0,
                        318.0,
                        352.0,
                        358.0,
                        167.0,
                        210.0,
                        -185.0,
                        327.0,
                        117.0,
                        350.0,
                        -170.0,
                        -144.0,
                        -14.0,
                        -37.0,
                        318.0,
                        243.0,
                        33.0,
                        90.0,
                        319.0,
                        -270.0,
                        229.0,
                        122.0,
                        287.0,
                        -90.0,
                        -69.0,
                        -134.0,
                        -184.0,
                        25.0,
                        -178.0,
                        -89.0,
                        -273.0,
                        -49.0,
                        -362.0,
                        -9.0,
                        -71.0,
                        -347.0,
                        353.0,
                        342.0,
                        133.0,
                        -116.0,
                        231.0,
                        -231.0,
                        51.0,
                        288.0,
                        186.0,
                        -328.0,
                        275.0,
                        81.0,
                        94.0,
                        -263.0,
                        114.0,
                        13.0,
                        -357.0,
                        171.0,
                        -242.0,
                        -85.0,
                        -362.0,
                        108.0,
                        164.0,
                        69.0,
                        15.0,
                        57.0,
                        -287.0,
                        100.0,
                        165.0,
                        205.0,
                        204.0,
                        -78.0,
                        360.0,
                        -80.0,
                        -120.0,
                        -255.0,
                        -77.0,
                        110.0,
                        -26.0,
                        -149.0,
                        -254.0,
                        95.0,
                        32.0,
                        -57.0,
                        -195.0,
                        100.0,
                        221.0,
                        74.0,
                        274.0,
                        15.0,
                        353.0,
                        204.0,
                        -365.0,
                        315.0,
                        344.0,
                        199.0,
                        -59.0,
                        272.0,
                        173.0,
                        -40.0,
                        -318.0,
                        -330.0,
                        -365.0,
                        -272.0,
                        -149.0,
                        -27.0,
                        -334.0,
                        -277.0,
                        344.0,
                        351.0,
                        -310.0,
                        264.0,
                        281.0,
                        176.0,
                        191.0,
                        -183.0,
                        288.0,
                        -112.0,
                        -55.0,
                        -166.0,
                        258.0,
                        194.0,
                        59.0,
                    },
                    Timezone: unifiedgosdk.Pointer("America/Guadeloupe"),
                    WeekStart: shared.WeekStartTu.ToPointer(),
                },
                shared.CalendarEventRecurrence{
                    Count: unifiedgosdk.Pointer[float64](1.0),
                    EndAt: types.MustNewTimeFromString("2020-11-04T17:53:56.516Z"),
                    ExcludedDates: []string{
                        "2023-01-11T21:34:37.555Z",
                        "2021-09-07T13:21:42.126Z",
                    },
                    Frequency: shared.CalendarEventRecurrenceFrequencyWeekly.ToPointer(),
                    IncludedDates: []string{
                        "2024-08-31T10:58:44.143Z",
                    },
                    Interval: unifiedgosdk.Pointer[float64](9.0),
                    OnDays: []shared.PropertyCalendarEventRecurrenceOnDays{
                        shared.PropertyCalendarEventRecurrenceOnDaysTu,
                        shared.PropertyCalendarEventRecurrenceOnDaysSa,
                    },
                    OnMonthDays: []float64{
                        -2.0,
                    },
                    OnMonths: []float64{
                        -4.0,
                        8.0,
                        0.0,
                        9.0,
                        4.0,
                        -11.0,
                        7.0,
                        1.0,
                        -5.0,
                    },
                    OnWeeks: []float64{
                        -36.0,
                        -31.0,
                        -16.0,
                        -6.0,
                        44.0,
                        -37.0,
                        14.0,
                        38.0,
                        -27.0,
                        -22.0,
                        -2.0,
                        24.0,
                        7.0,
                        50.0,
                        46.0,
                        52.0,
                        20.0,
                        37.0,
                        31.0,
                        48.0,
                        35.0,
                        -46.0,
                        13.0,
                        22.0,
                        53.0,
                        20.0,
                        -28.0,
                        -2.0,
                        39.0,
                        13.0,
                        4.0,
                        0.0,
                        7.0,
                        -38.0,
                        -35.0,
                        41.0,
                        49.0,
                        12.0,
                        17.0,
                        8.0,
                        49.0,
                        -47.0,
                        46.0,
                        25.0,
                        14.0,
                        -26.0,
                        -37.0,
                        -25.0,
                        -41.0,
                        27.0,
                        28.0,
                        -19.0,
                    },
                    OnYearDays: []float64{
                        -256.0,
                        -328.0,
                        -312.0,
                        50.0,
                        -251.0,
                        -338.0,
                        -315.0,
                        214.0,
                        129.0,
                        -263.0,
                        -108.0,
                        -11.0,
                        206.0,
                        -29.0,
                        -159.0,
                        -29.0,
                        -264.0,
                        295.0,
                        -231.0,
                        53.0,
                        34.0,
                        -366.0,
                        326.0,
                        -202.0,
                        151.0,
                        79.0,
                        -66.0,
                        11.0,
                        -42.0,
                        73.0,
                        338.0,
                        -155.0,
                        197.0,
                        260.0,
                        356.0,
                        -323.0,
                        -213.0,
                        -332.0,
                        -305.0,
                        -182.0,
                        -253.0,
                        -276.0,
                        -285.0,
                        96.0,
                        -336.0,
                        269.0,
                        -233.0,
                        250.0,
                        -112.0,
                        -307.0,
                        -96.0,
                        54.0,
                        267.0,
                        318.0,
                        -66.0,
                        11.0,
                        -303.0,
                        231.0,
                        165.0,
                        -297.0,
                        -348.0,
                        -355.0,
                        364.0,
                        312.0,
                        -26.0,
                        111.0,
                        162.0,
                        280.0,
                        312.0,
                        337.0,
                        235.0,
                        68.0,
                        -282.0,
                        363.0,
                        212.0,
                        -328.0,
                        9.0,
                        -24.0,
                        -163.0,
                        -101.0,
                        -79.0,
                        -264.0,
                        -157.0,
                        188.0,
                        290.0,
                        51.0,
                        -213.0,
                        216.0,
                        230.0,
                        -270.0,
                        -211.0,
                        -156.0,
                        -165.0,
                        -305.0,
                        -45.0,
                        224.0,
                        -248.0,
                        65.0,
                        9.0,
                        274.0,
                        -299.0,
                        -228.0,
                        33.0,
                        -42.0,
                        356.0,
                        -311.0,
                        241.0,
                        261.0,
                        -136.0,
                        -252.0,
                        166.0,
                        208.0,
                        -126.0,
                        64.0,
                        323.0,
                        -104.0,
                        -106.0,
                        -248.0,
                        -41.0,
                        -109.0,
                        245.0,
                        47.0,
                        205.0,
                        358.0,
                        -296.0,
                        214.0,
                        -157.0,
                        -313.0,
                        -303.0,
                        -54.0,
                        -229.0,
                        231.0,
                        -94.0,
                        -198.0,
                        338.0,
                        199.0,
                        5.0,
                        42.0,
                        309.0,
                        73.0,
                        56.0,
                        -120.0,
                        351.0,
                        6.0,
                        -193.0,
                        21.0,
                        78.0,
                        57.0,
                        -269.0,
                        -76.0,
                        -299.0,
                        295.0,
                        -278.0,
                        11.0,
                        121.0,
                        -323.0,
                        156.0,
                        67.0,
                        152.0,
                        284.0,
                        108.0,
                        -7.0,
                        329.0,
                        -32.0,
                        333.0,
                        -338.0,
                        148.0,
                        -42.0,
                        151.0,
                        145.0,
                        -34.0,
                        -36.0,
                        296.0,
                        -198.0,
                        -317.0,
                        -161.0,
                        -253.0,
                        328.0,
                        -57.0,
                        134.0,
                        -289.0,
                        229.0,
                        44.0,
                        16.0,
                        -256.0,
                        289.0,
                        -234.0,
                        197.0,
                        333.0,
                        228.0,
                        -143.0,
                        -202.0,
                        -172.0,
                        -262.0,
                        -203.0,
                        -83.0,
                        -242.0,
                        -173.0,
                        336.0,
                        298.0,
                        -319.0,
                        66.0,
                        254.0,
                        214.0,
                        -118.0,
                        -216.0,
                        -168.0,
                        44.0,
                        -243.0,
                        207.0,
                        -28.0,
                        -4.0,
                        -272.0,
                        79.0,
                    },
                    Timezone: unifiedgosdk.Pointer("Atlantic/Reykjavik"),
                    WeekStart: shared.WeekStartTu.ToPointer(),
                },
            },
            RecurringEventID: unifiedgosdk.Pointer("bd28e1ea-fb99-454e-89c0-6be3542789d1"),
            SendNotifications: unifiedgosdk.Pointer(false),
            StartAt: unifiedgosdk.Pointer("2020-05-20T08:05:45.139Z"),
            Status: shared.CalendarEventStatusConfirmed.ToPointer(),
            Subject: unifiedgosdk.Pointer("Sunt spargo tepidus bestia vigor credo coadunatio appello."),
            Timezone: unifiedgosdk.Pointer("Asia/Bangkok"),
            UpdatedAt: unifiedgosdk.Pointer("2020-06-26T04:47:47.416Z"),
            WebURL: unifiedgosdk.Pointer("https://another-pinstripe.com"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CalendarEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreateCalendarEventRequest](../../pkg/models/operations/createcalendareventrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreateCalendarEventResponse](../../pkg/models/operations/createcalendareventresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateCalendarLink

Create a link

### Example Usage

<!-- UsageSnippet language="go" operationID="createCalendarLink" method="post" path="/calendar/{connection_id}/link" example="calendar_link" -->
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

    res, err := s.Calendar.CreateCalendarLink(ctx, operations.CreateCalendarLinkRequest{
        CalendarLink: shared.CalendarLink{
            CreatedAt: unifiedgosdk.Pointer("2023-03-07T13:34:11.959Z"),
            Description: unifiedgosdk.Pointer("Vitium clibanus laboriosam uxor denuncio."),
            Duration: unifiedgosdk.Pointer[float64](74.0),
            ID: unifiedgosdk.Pointer("f8e8095c-0f1e-458a-91ab-2c7a3bc373aa"),
            IsActive: unifiedgosdk.Pointer(true),
            Name: unifiedgosdk.Pointer("Sopor sopor ancilla animus anser dignissimos vito confero utilis."),
            PriceAmount: unifiedgosdk.Pointer[float64](44.0),
            PriceCurrency: unifiedgosdk.Pointer("USD"),
            UpdatedAt: unifiedgosdk.Pointer("2024-03-06T11:31:30.143Z"),
            URL: "https://annual-apricot.info/",
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CalendarLink != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateCalendarLinkRequest](../../pkg/models/operations/createcalendarlinkrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateCalendarLinkResponse](../../pkg/models/operations/createcalendarlinkresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateCalendarWebinar

Create a webinar

### Example Usage

<!-- UsageSnippet language="go" operationID="createCalendarWebinar" method="post" path="/calendar/{connection_id}/webinar" example="calendar_webinar" -->
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

    res, err := s.Calendar.CreateCalendarWebinar(ctx, operations.CreateCalendarWebinarRequest{
        CalendarWebinar: shared.CalendarWebinar{
            Conference: []shared.CalendarConference{},
            CreatedAt: types.MustNewTimeFromString("2022-07-06T11:45:14.631Z"),
            EndAt: types.MustNewTimeFromString("2025-10-03T23:05:05.692Z"),
            HasPolls: unifiedgosdk.Pointer(false),
            HasRecording: unifiedgosdk.Pointer(false),
            ID: unifiedgosdk.Pointer("038e6b28-c41c-451b-8571-3f2727636b28"),
            IsAutoApprove: unifiedgosdk.Pointer(false),
            IsEnabled: unifiedgosdk.Pointer(true),
            IsWebcast: unifiedgosdk.Pointer(false),
            JoinURL: unifiedgosdk.Pointer("https://robust-bathhouse.biz"),
            Notes: unifiedgosdk.Pointer("Curriculum ducimus assentator aspernatur ait."),
            Organizer: &shared.PropertyCalendarWebinarOrganizer{
                Email: unifiedgosdk.Pointer("Kelton_Dicki@yahoo.com"),
                Name: unifiedgosdk.Pointer("Walter Greenfelder"),
            },
            Recurrence: []shared.CalendarEventRecurrence{
                shared.CalendarEventRecurrence{
                    Count: unifiedgosdk.Pointer[float64](10.0),
                    EndAt: types.MustNewTimeFromString("2023-08-23T00:00:57.819Z"),
                    ExcludedDates: []string{
                        "2025-01-24T12:51:50.254Z",
                    },
                    Frequency: shared.CalendarEventRecurrenceFrequencyMonthly.ToPointer(),
                    IncludedDates: []string{
                        "2024-04-14T17:49:13.787Z",
                    },
                    Interval: unifiedgosdk.Pointer[float64](8.0),
                    OnDays: []shared.PropertyCalendarEventRecurrenceOnDays{
                        shared.PropertyCalendarEventRecurrenceOnDaysSu,
                        shared.PropertyCalendarEventRecurrenceOnDaysFr,
                        shared.PropertyCalendarEventRecurrenceOnDaysSa,
                        shared.PropertyCalendarEventRecurrenceOnDaysWe,
                        shared.PropertyCalendarEventRecurrenceOnDaysMo,
                    },
                    OnMonthDays: []float64{
                        -10.0,
                    },
                    OnMonths: []float64{
                        -9.0,
                    },
                    OnWeeks: []float64{
                        10.0,
                        30.0,
                        -38.0,
                        30.0,
                        -22.0,
                        37.0,
                        -12.0,
                        27.0,
                        2.0,
                        15.0,
                        26.0,
                        18.0,
                        -43.0,
                        -33.0,
                        -27.0,
                        38.0,
                        28.0,
                        47.0,
                        -8.0,
                        24.0,
                        35.0,
                        -2.0,
                        7.0,
                        49.0,
                        38.0,
                        -41.0,
                        46.0,
                        -11.0,
                        -45.0,
                        0.0,
                        48.0,
                        34.0,
                    },
                    OnYearDays: []float64{
                        345.0,
                        -207.0,
                        230.0,
                        -10.0,
                        364.0,
                        -256.0,
                        -218.0,
                        -295.0,
                        290.0,
                        -250.0,
                        -315.0,
                        60.0,
                        205.0,
                        -247.0,
                        -318.0,
                        -211.0,
                        -13.0,
                        256.0,
                        -200.0,
                        -313.0,
                        336.0,
                        -332.0,
                        -90.0,
                        287.0,
                        -273.0,
                        156.0,
                        241.0,
                        -138.0,
                        -363.0,
                        -37.0,
                        -171.0,
                        -62.0,
                        -57.0,
                        280.0,
                        -322.0,
                        -79.0,
                        -364.0,
                        -201.0,
                        84.0,
                        341.0,
                        334.0,
                        -75.0,
                        332.0,
                        207.0,
                        337.0,
                        -244.0,
                        131.0,
                        -191.0,
                        164.0,
                        -235.0,
                        285.0,
                        -309.0,
                        -158.0,
                        306.0,
                        180.0,
                        -130.0,
                        -162.0,
                        -155.0,
                        3.0,
                        198.0,
                        26.0,
                        -366.0,
                        -191.0,
                        127.0,
                        -331.0,
                        -11.0,
                        -239.0,
                        -189.0,
                        243.0,
                        118.0,
                        346.0,
                        -174.0,
                        -146.0,
                        -161.0,
                        -330.0,
                        327.0,
                        192.0,
                        310.0,
                        316.0,
                        313.0,
                        -242.0,
                        -51.0,
                        -264.0,
                        -180.0,
                        -88.0,
                        305.0,
                        270.0,
                        358.0,
                        -173.0,
                        -298.0,
                        153.0,
                        -89.0,
                        155.0,
                        -45.0,
                        248.0,
                        -46.0,
                        -146.0,
                        300.0,
                        364.0,
                        -335.0,
                        356.0,
                        -18.0,
                        219.0,
                        324.0,
                        -239.0,
                        -106.0,
                        -298.0,
                        328.0,
                        362.0,
                        344.0,
                        -54.0,
                        133.0,
                        50.0,
                        112.0,
                        -212.0,
                        -179.0,
                        22.0,
                        -201.0,
                        -62.0,
                        -293.0,
                        9.0,
                        30.0,
                        -50.0,
                        126.0,
                        -72.0,
                        264.0,
                        28.0,
                        -1.0,
                        -207.0,
                        160.0,
                        -168.0,
                        3.0,
                        -176.0,
                        -19.0,
                        -157.0,
                        349.0,
                        100.0,
                        -201.0,
                        108.0,
                        -180.0,
                        51.0,
                        -73.0,
                        366.0,
                        74.0,
                        -226.0,
                        238.0,
                        121.0,
                        -193.0,
                        -125.0,
                        -109.0,
                        316.0,
                        -177.0,
                        -307.0,
                        31.0,
                        -76.0,
                        217.0,
                        -310.0,
                        227.0,
                        -360.0,
                        71.0,
                        255.0,
                        -325.0,
                        -214.0,
                        40.0,
                        42.0,
                        17.0,
                        -241.0,
                        -84.0,
                        -188.0,
                        302.0,
                        64.0,
                        94.0,
                        -362.0,
                        23.0,
                        166.0,
                        85.0,
                        71.0,
                        -74.0,
                        -47.0,
                        -119.0,
                        98.0,
                        40.0,
                        158.0,
                        -64.0,
                        175.0,
                        269.0,
                        127.0,
                        -143.0,
                        213.0,
                        -196.0,
                        121.0,
                        81.0,
                        -238.0,
                        288.0,
                        321.0,
                        276.0,
                        133.0,
                        22.0,
                        -213.0,
                        -157.0,
                        -280.0,
                        -35.0,
                        73.0,
                        -194.0,
                        65.0,
                        -180.0,
                        63.0,
                        -242.0,
                        -117.0,
                        148.0,
                        157.0,
                        -320.0,
                        318.0,
                        8.0,
                        210.0,
                        -21.0,
                        81.0,
                        205.0,
                        -258.0,
                        -40.0,
                        -114.0,
                        -253.0,
                        -263.0,
                        65.0,
                        185.0,
                        -24.0,
                        324.0,
                        -172.0,
                        25.0,
                        260.0,
                        211.0,
                        342.0,
                        -31.0,
                        -288.0,
                        -159.0,
                        -4.0,
                        -2.0,
                        -107.0,
                        -316.0,
                        -276.0,
                        331.0,
                        -114.0,
                        -20.0,
                        -320.0,
                        51.0,
                        -176.0,
                        -148.0,
                        -50.0,
                        -201.0,
                        -104.0,
                        153.0,
                        -273.0,
                        -189.0,
                        67.0,
                        209.0,
                        149.0,
                        49.0,
                        -136.0,
                        -125.0,
                        -169.0,
                        -324.0,
                        309.0,
                        -51.0,
                        288.0,
                        253.0,
                        175.0,
                        -146.0,
                        171.0,
                        -140.0,
                        58.0,
                        -212.0,
                        164.0,
                        270.0,
                        102.0,
                        70.0,
                        299.0,
                        89.0,
                        -280.0,
                        252.0,
                        -342.0,
                        240.0,
                        226.0,
                        68.0,
                        -30.0,
                        -232.0,
                        -358.0,
                        -166.0,
                        60.0,
                        140.0,
                        275.0,
                        13.0,
                        250.0,
                        -328.0,
                        -189.0,
                        -22.0,
                        7.0,
                        -235.0,
                        -322.0,
                        178.0,
                        167.0,
                        -104.0,
                        -61.0,
                        282.0,
                        -80.0,
                        -277.0,
                        108.0,
                        271.0,
                        -237.0,
                        297.0,
                        -135.0,
                        -135.0,
                        -323.0,
                        342.0,
                        -267.0,
                        -235.0,
                        173.0,
                        249.0,
                        -288.0,
                        257.0,
                        139.0,
                        -191.0,
                        -217.0,
                        10.0,
                        -117.0,
                        -297.0,
                        -196.0,
                        -206.0,
                        341.0,
                        166.0,
                        181.0,
                        129.0,
                        -207.0,
                        55.0,
                        86.0,
                    },
                    Timezone: unifiedgosdk.Pointer("Asia/Ust-Nera"),
                    WeekStart: shared.WeekStartMo.ToPointer(),
                },
                shared.CalendarEventRecurrence{
                    Count: unifiedgosdk.Pointer[float64](3.0),
                    EndAt: types.MustNewTimeFromString("2022-09-28T21:54:22.886Z"),
                    ExcludedDates: []string{
                        "2024-08-16T15:01:59.491Z",
                        "2024-08-01T09:41:48.731Z",
                    },
                    Frequency: shared.CalendarEventRecurrenceFrequencyDaily.ToPointer(),
                    IncludedDates: []string{
                        "2024-03-12T07:59:01.460Z",
                        "2025-12-18T01:45:08.067Z",
                        "2023-08-06T00:06:02.449Z",
                    },
                    Interval: unifiedgosdk.Pointer[float64](1.0),
                    OnDays: []shared.PropertyCalendarEventRecurrenceOnDays{
                        shared.PropertyCalendarEventRecurrenceOnDaysWe,
                        shared.PropertyCalendarEventRecurrenceOnDaysSu,
                        shared.PropertyCalendarEventRecurrenceOnDaysMo,
                        shared.PropertyCalendarEventRecurrenceOnDaysFr,
                    },
                    OnMonthDays: []float64{
                        -15.0,
                    },
                    OnMonths: []float64{
                        5.0,
                        12.0,
                        3.0,
                        12.0,
                        8.0,
                    },
                    OnWeeks: []float64{
                        -47.0,
                        44.0,
                    },
                    OnYearDays: []float64{
                        -117.0,
                        59.0,
                        -6.0,
                        187.0,
                        45.0,
                        70.0,
                        15.0,
                        255.0,
                        44.0,
                        -2.0,
                        25.0,
                        -175.0,
                        -240.0,
                        171.0,
                        -294.0,
                        19.0,
                        38.0,
                        -351.0,
                        170.0,
                        -10.0,
                        -269.0,
                        18.0,
                        -65.0,
                        -266.0,
                        -31.0,
                        328.0,
                        -361.0,
                        358.0,
                        -256.0,
                        -4.0,
                        -312.0,
                        82.0,
                        -2.0,
                        -75.0,
                        -281.0,
                        -304.0,
                        53.0,
                        -295.0,
                        366.0,
                        322.0,
                        -191.0,
                        26.0,
                        97.0,
                        53.0,
                        75.0,
                        -62.0,
                        -109.0,
                        66.0,
                        177.0,
                        -68.0,
                        175.0,
                        -280.0,
                        70.0,
                        -238.0,
                        109.0,
                        -304.0,
                        326.0,
                        -8.0,
                        -71.0,
                        -236.0,
                        225.0,
                        358.0,
                        20.0,
                        -5.0,
                        -102.0,
                        -134.0,
                        -204.0,
                        -116.0,
                        -353.0,
                        -273.0,
                        106.0,
                        284.0,
                        -137.0,
                        -324.0,
                        301.0,
                        -42.0,
                        -229.0,
                        271.0,
                        -293.0,
                        -343.0,
                        211.0,
                        47.0,
                        -254.0,
                        -154.0,
                        -182.0,
                        264.0,
                        120.0,
                        -11.0,
                        -307.0,
                        99.0,
                        227.0,
                        190.0,
                        -17.0,
                        -77.0,
                        -255.0,
                        -61.0,
                        -249.0,
                        -102.0,
                        70.0,
                        345.0,
                        -187.0,
                        -308.0,
                        194.0,
                        221.0,
                        268.0,
                        -169.0,
                        -190.0,
                        88.0,
                        10.0,
                        262.0,
                        177.0,
                        -314.0,
                        -151.0,
                        -295.0,
                    },
                    Timezone: unifiedgosdk.Pointer("Pacific/Wake"),
                    WeekStart: shared.WeekStartTu.ToPointer(),
                },
                shared.CalendarEventRecurrence{
                    Count: unifiedgosdk.Pointer[float64](8.0),
                    EndAt: types.MustNewTimeFromString("2026-06-26T05:33:58.747Z"),
                    ExcludedDates: []string{
                        "2023-06-11T12:02:36.558Z",
                        "2023-05-31T18:16:08.915Z",
                    },
                    Frequency: shared.CalendarEventRecurrenceFrequencyWeekly.ToPointer(),
                    IncludedDates: []string{
                        "2024-03-20T04:54:34.086Z",
                        "2023-08-11T16:40:30.422Z",
                        "2024-09-10T07:26:29.376Z",
                    },
                    Interval: unifiedgosdk.Pointer[float64](8.0),
                    OnDays: []shared.PropertyCalendarEventRecurrenceOnDays{
                        shared.PropertyCalendarEventRecurrenceOnDaysSu,
                        shared.PropertyCalendarEventRecurrenceOnDaysMo,
                        shared.PropertyCalendarEventRecurrenceOnDaysTu,
                        shared.PropertyCalendarEventRecurrenceOnDaysFr,
                        shared.PropertyCalendarEventRecurrenceOnDaysMo,
                        shared.PropertyCalendarEventRecurrenceOnDaysTh,
                    },
                    OnMonthDays: []float64{
                        -23.0,
                    },
                    OnMonths: []float64{
                        11.0,
                        8.0,
                        9.0,
                        5.0,
                        -12.0,
                        -7.0,
                        -5.0,
                        10.0,
                        10.0,
                        -9.0,
                        -10.0,
                    },
                    OnWeeks: []float64{
                        -49.0,
                        46.0,
                        35.0,
                        -26.0,
                        2.0,
                        15.0,
                        15.0,
                        -26.0,
                        24.0,
                        -53.0,
                        36.0,
                        -43.0,
                        51.0,
                        -19.0,
                        -7.0,
                        -12.0,
                        28.0,
                        27.0,
                        35.0,
                        12.0,
                        -28.0,
                        -8.0,
                        -4.0,
                        -45.0,
                    },
                    OnYearDays: []float64{
                        84.0,
                        -251.0,
                        71.0,
                        181.0,
                        -163.0,
                        158.0,
                        301.0,
                        -299.0,
                        -184.0,
                        -331.0,
                        -152.0,
                        -129.0,
                        -237.0,
                        -303.0,
                        -24.0,
                        126.0,
                        -103.0,
                        146.0,
                        -346.0,
                        86.0,
                        -296.0,
                        -337.0,
                        -185.0,
                        16.0,
                        -270.0,
                        -126.0,
                        -295.0,
                        -231.0,
                        356.0,
                        -293.0,
                        115.0,
                        -265.0,
                        -293.0,
                        -34.0,
                        357.0,
                        313.0,
                        -343.0,
                        180.0,
                        -22.0,
                        -161.0,
                        350.0,
                        177.0,
                        190.0,
                        223.0,
                        -152.0,
                        -360.0,
                        -225.0,
                        -60.0,
                        -35.0,
                        353.0,
                        117.0,
                        -171.0,
                        -315.0,
                        -321.0,
                        -202.0,
                        345.0,
                        -1.0,
                        -148.0,
                        -168.0,
                        181.0,
                        -17.0,
                        282.0,
                        234.0,
                        31.0,
                        47.0,
                        -236.0,
                        366.0,
                        -251.0,
                        -232.0,
                        -308.0,
                        76.0,
                        -199.0,
                        184.0,
                        198.0,
                        225.0,
                        75.0,
                        6.0,
                        227.0,
                        -148.0,
                        259.0,
                        -146.0,
                        49.0,
                        -254.0,
                        341.0,
                        93.0,
                        138.0,
                        -164.0,
                        237.0,
                        4.0,
                        -287.0,
                        161.0,
                    },
                    Timezone: unifiedgosdk.Pointer("Africa/Bissau"),
                    WeekStart: shared.WeekStartWe.ToPointer(),
                },
            },
            RegistrantPassword: unifiedgosdk.Pointer("OxwWzr0C"),
            RequireAddress: unifiedgosdk.Pointer(false),
            RequireJobTitle: unifiedgosdk.Pointer(false),
            StartAt: types.MustNewTimeFromString("2025-04-09T12:29:18.731Z"),
            Status: shared.CalendarWebinarStatusTentative.ToPointer(),
            Subject: unifiedgosdk.Pointer("Harum culpa decipio ex cubo ancilla cresco."),
            Timezone: unifiedgosdk.Pointer("Europe/Kaliningrad"),
            UpdatedAt: types.MustNewTimeFromString("2026-08-29T20:26:31.424Z"),
            WebURL: unifiedgosdk.Pointer("https://classic-recovery.biz"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CalendarWebinar != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CreateCalendarWebinarRequest](../../pkg/models/operations/createcalendarwebinarrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.CreateCalendarWebinarResponse](../../pkg/models/operations/createcalendarwebinarresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCalendarCalendar

Retrieve a calendar

### Example Usage

<!-- UsageSnippet language="go" operationID="getCalendarCalendar" method="get" path="/calendar/{connection_id}/calendar/{id}" -->
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

    res, err := s.Calendar.GetCalendarCalendar(ctx, operations.GetCalendarCalendarRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CalendarCalendar != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetCalendarCalendarRequest](../../pkg/models/operations/getcalendarcalendarrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.GetCalendarCalendarResponse](../../pkg/models/operations/getcalendarcalendarresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCalendarEvent

Retrieve an event

### Example Usage

<!-- UsageSnippet language="go" operationID="getCalendarEvent" method="get" path="/calendar/{connection_id}/event/{id}" -->
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

    res, err := s.Calendar.GetCalendarEvent(ctx, operations.GetCalendarEventRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CalendarEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetCalendarEventRequest](../../pkg/models/operations/getcalendareventrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetCalendarEventResponse](../../pkg/models/operations/getcalendareventresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCalendarLink

Retrieve a link

### Example Usage

<!-- UsageSnippet language="go" operationID="getCalendarLink" method="get" path="/calendar/{connection_id}/link/{id}" -->
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

    res, err := s.Calendar.GetCalendarLink(ctx, operations.GetCalendarLinkRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CalendarLink != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetCalendarLinkRequest](../../pkg/models/operations/getcalendarlinkrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetCalendarLinkResponse](../../pkg/models/operations/getcalendarlinkresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCalendarRecording

Retrieve a recording

### Example Usage

<!-- UsageSnippet language="go" operationID="getCalendarRecording" method="get" path="/calendar/{connection_id}/recording/{id}" -->
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

    res, err := s.Calendar.GetCalendarRecording(ctx, operations.GetCalendarRecordingRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CalendarRecording != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetCalendarRecordingRequest](../../pkg/models/operations/getcalendarrecordingrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetCalendarRecordingResponse](../../pkg/models/operations/getcalendarrecordingresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetCalendarWebinar

Retrieve a webinar

### Example Usage

<!-- UsageSnippet language="go" operationID="getCalendarWebinar" method="get" path="/calendar/{connection_id}/webinar/{id}" -->
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

    res, err := s.Calendar.GetCalendarWebinar(ctx, operations.GetCalendarWebinarRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CalendarWebinar != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetCalendarWebinarRequest](../../pkg/models/operations/getcalendarwebinarrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetCalendarWebinarResponse](../../pkg/models/operations/getcalendarwebinarresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCalendarBusies

List all busies

### Example Usage

<!-- UsageSnippet language="go" operationID="listCalendarBusies" method="get" path="/calendar/{connection_id}/busy" -->
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

    res, err := s.Calendar.ListCalendarBusies(ctx, operations.ListCalendarBusiesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CalendarBusies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListCalendarBusiesRequest](../../pkg/models/operations/listcalendarbusiesrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListCalendarBusiesResponse](../../pkg/models/operations/listcalendarbusiesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCalendarCalendars

List all calendars

### Example Usage

<!-- UsageSnippet language="go" operationID="listCalendarCalendars" method="get" path="/calendar/{connection_id}/calendar" -->
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

    res, err := s.Calendar.ListCalendarCalendars(ctx, operations.ListCalendarCalendarsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CalendarCalendars != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListCalendarCalendarsRequest](../../pkg/models/operations/listcalendarcalendarsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.ListCalendarCalendarsResponse](../../pkg/models/operations/listcalendarcalendarsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCalendarEvents

List all events

### Example Usage

<!-- UsageSnippet language="go" operationID="listCalendarEvents" method="get" path="/calendar/{connection_id}/event" -->
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

    res, err := s.Calendar.ListCalendarEvents(ctx, operations.ListCalendarEventsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CalendarEvents != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListCalendarEventsRequest](../../pkg/models/operations/listcalendareventsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListCalendarEventsResponse](../../pkg/models/operations/listcalendareventsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCalendarLinks

List all links

### Example Usage

<!-- UsageSnippet language="go" operationID="listCalendarLinks" method="get" path="/calendar/{connection_id}/link" -->
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

    res, err := s.Calendar.ListCalendarLinks(ctx, operations.ListCalendarLinksRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CalendarLinks != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListCalendarLinksRequest](../../pkg/models/operations/listcalendarlinksrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListCalendarLinksResponse](../../pkg/models/operations/listcalendarlinksresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCalendarRecordings

List all recordings

### Example Usage

<!-- UsageSnippet language="go" operationID="listCalendarRecordings" method="get" path="/calendar/{connection_id}/recording" -->
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

    res, err := s.Calendar.ListCalendarRecordings(ctx, operations.ListCalendarRecordingsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CalendarRecordings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListCalendarRecordingsRequest](../../pkg/models/operations/listcalendarrecordingsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.ListCalendarRecordingsResponse](../../pkg/models/operations/listcalendarrecordingsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListCalendarWebinars

List all webinars

### Example Usage

<!-- UsageSnippet language="go" operationID="listCalendarWebinars" method="get" path="/calendar/{connection_id}/webinar" -->
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

    res, err := s.Calendar.ListCalendarWebinars(ctx, operations.ListCalendarWebinarsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CalendarWebinars != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListCalendarWebinarsRequest](../../pkg/models/operations/listcalendarwebinarsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListCalendarWebinarsResponse](../../pkg/models/operations/listcalendarwebinarsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchCalendarCalendar

Update a calendar

### Example Usage

<!-- UsageSnippet language="go" operationID="patchCalendarCalendar" method="patch" path="/calendar/{connection_id}/calendar/{id}" example="calendar_calendar" -->
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

    res, err := s.Calendar.PatchCalendarCalendar(ctx, operations.PatchCalendarCalendarRequest{
        CalendarCalendar: shared.CalendarCalendar{
            CreatedAt: types.MustNewTimeFromString("2020-01-09T23:11:34.147Z"),
            Description: unifiedgosdk.Pointer("Socius catena auxilium."),
            ID: unifiedgosdk.Pointer("c8bd690d-d2e7-4367-a8d9-54cb99b57437"),
            IsPrimary: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("Acer supra vallum suasoria thesaurus omnis condico cognomen accendo vehemens."),
            Timezone: unifiedgosdk.Pointer("America/Dawson_Creek"),
            UpdatedAt: types.MustNewTimeFromString("2023-03-12T22:44:44.999Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CalendarCalendar != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.PatchCalendarCalendarRequest](../../pkg/models/operations/patchcalendarcalendarrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.PatchCalendarCalendarResponse](../../pkg/models/operations/patchcalendarcalendarresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchCalendarEvent

Update an event

### Example Usage

<!-- UsageSnippet language="go" operationID="patchCalendarEvent" method="patch" path="/calendar/{connection_id}/event/{id}" example="calendar_event" -->
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

    res, err := s.Calendar.PatchCalendarEvent(ctx, operations.PatchCalendarEventRequest{
        CalendarEvent: shared.CalendarEvent{
            Attachments: []shared.CalendarAttachment{},
            Conference: []shared.CalendarConference{},
            CreatedAt: unifiedgosdk.Pointer("2019-08-04T14:33:51.814Z"),
            EndAt: unifiedgosdk.Pointer("2020-05-20T17:12:37.979Z"),
            ID: unifiedgosdk.Pointer("67215e17-405a-4402-bfdf-b40ac9978bb3"),
            IsAllDay: unifiedgosdk.Pointer(false),
            IsFree: unifiedgosdk.Pointer(false),
            IsPrivate: unifiedgosdk.Pointer(false),
            Location: unifiedgosdk.Pointer("621 Boehm Prairie"),
            Notes: unifiedgosdk.Pointer("Aegre traho."),
            Recurrence: []shared.CalendarEventRecurrence{
                shared.CalendarEventRecurrence{
                    Count: unifiedgosdk.Pointer[float64](8.0),
                    EndAt: types.MustNewTimeFromString("2025-12-26T06:34:54.364Z"),
                    ExcludedDates: []string{
                        "2025-09-30T22:42:44.936Z",
                        "2023-10-09T12:14:01.324Z",
                        "2024-02-15T13:20:24.736Z",
                    },
                    Frequency: shared.CalendarEventRecurrenceFrequencyMonthly.ToPointer(),
                    IncludedDates: []string{
                        "2021-02-16T22:04:02.467Z",
                    },
                    Interval: unifiedgosdk.Pointer[float64](4.0),
                    OnDays: []shared.PropertyCalendarEventRecurrenceOnDays{
                        shared.PropertyCalendarEventRecurrenceOnDaysTh,
                        shared.PropertyCalendarEventRecurrenceOnDaysMo,
                        shared.PropertyCalendarEventRecurrenceOnDaysTh,
                    },
                    OnMonthDays: []float64{
                        -26.0,
                    },
                    OnMonths: []float64{
                        12.0,
                        9.0,
                        -1.0,
                        0.0,
                        1.0,
                        6.0,
                        -10.0,
                        9.0,
                        0.0,
                        4.0,
                        -2.0,
                    },
                    OnWeeks: []float64{
                        -7.0,
                        51.0,
                        -3.0,
                        -41.0,
                        15.0,
                        46.0,
                        -1.0,
                        46.0,
                        42.0,
                        11.0,
                        12.0,
                        -35.0,
                        -15.0,
                        -3.0,
                        -42.0,
                        50.0,
                        3.0,
                        -15.0,
                        -10.0,
                        6.0,
                        -53.0,
                        5.0,
                        -32.0,
                        -22.0,
                        43.0,
                        -44.0,
                        -23.0,
                        -21.0,
                        -18.0,
                    },
                    OnYearDays: []float64{
                        -35.0,
                        14.0,
                        -338.0,
                        175.0,
                        -87.0,
                        339.0,
                        341.0,
                        287.0,
                        -17.0,
                        319.0,
                        -3.0,
                        238.0,
                        -115.0,
                        -116.0,
                        283.0,
                        -61.0,
                        -254.0,
                        86.0,
                        -163.0,
                        5.0,
                        -171.0,
                        -99.0,
                        279.0,
                        19.0,
                        303.0,
                        -106.0,
                        90.0,
                        109.0,
                        -185.0,
                        -285.0,
                        -83.0,
                        -236.0,
                        66.0,
                        -215.0,
                        178.0,
                        64.0,
                        78.0,
                        5.0,
                        -251.0,
                        -79.0,
                        -271.0,
                        33.0,
                        320.0,
                        67.0,
                        -84.0,
                        -355.0,
                        -364.0,
                        348.0,
                        271.0,
                        -304.0,
                        -199.0,
                        106.0,
                        -345.0,
                        24.0,
                        -89.0,
                        -109.0,
                        -314.0,
                        365.0,
                        38.0,
                        -42.0,
                        123.0,
                        56.0,
                        -3.0,
                        31.0,
                        101.0,
                        326.0,
                        -160.0,
                        -101.0,
                        -267.0,
                        -309.0,
                        -363.0,
                        125.0,
                        -182.0,
                        363.0,
                        324.0,
                        36.0,
                        -269.0,
                        -79.0,
                        -60.0,
                        272.0,
                        -254.0,
                        -160.0,
                        -82.0,
                        19.0,
                        42.0,
                        69.0,
                        -104.0,
                        333.0,
                        236.0,
                        -287.0,
                        296.0,
                        261.0,
                        241.0,
                        348.0,
                        -72.0,
                        159.0,
                        -127.0,
                        229.0,
                        -158.0,
                        190.0,
                        -173.0,
                        -84.0,
                        -96.0,
                        176.0,
                        339.0,
                        -48.0,
                        287.0,
                        -46.0,
                        -101.0,
                        246.0,
                        -8.0,
                        -74.0,
                        338.0,
                        -51.0,
                        -42.0,
                        -128.0,
                        -169.0,
                        -174.0,
                        168.0,
                        -85.0,
                        37.0,
                        169.0,
                        -105.0,
                        231.0,
                        -250.0,
                        -286.0,
                        -7.0,
                        -121.0,
                        321.0,
                        278.0,
                        -120.0,
                        -96.0,
                        360.0,
                        337.0,
                        -258.0,
                        -179.0,
                        324.0,
                        -204.0,
                        327.0,
                        15.0,
                        365.0,
                        191.0,
                        -345.0,
                        -345.0,
                        56.0,
                        217.0,
                        60.0,
                        -264.0,
                        -248.0,
                        -316.0,
                        191.0,
                        -189.0,
                        -152.0,
                        -296.0,
                        194.0,
                        -42.0,
                        -21.0,
                        -218.0,
                        171.0,
                        -15.0,
                        301.0,
                        37.0,
                        -167.0,
                        18.0,
                        248.0,
                        -263.0,
                        27.0,
                        14.0,
                        59.0,
                        219.0,
                        -284.0,
                        221.0,
                        -76.0,
                        277.0,
                        183.0,
                        200.0,
                        -12.0,
                        -28.0,
                        -79.0,
                        150.0,
                        320.0,
                        -152.0,
                        -15.0,
                        -42.0,
                        -125.0,
                        -4.0,
                        269.0,
                        290.0,
                        52.0,
                        320.0,
                        344.0,
                        13.0,
                        -69.0,
                        255.0,
                        -154.0,
                        -281.0,
                        158.0,
                        25.0,
                        240.0,
                        -339.0,
                        96.0,
                        204.0,
                        324.0,
                        221.0,
                        37.0,
                        -333.0,
                        87.0,
                        354.0,
                        -365.0,
                        -203.0,
                        -341.0,
                        -79.0,
                        -208.0,
                        135.0,
                        132.0,
                        -351.0,
                        39.0,
                        -87.0,
                        -297.0,
                        -66.0,
                        346.0,
                        69.0,
                        -177.0,
                        235.0,
                        295.0,
                        -366.0,
                        -55.0,
                    },
                    Timezone: unifiedgosdk.Pointer("Asia/Ho_Chi_Minh"),
                    WeekStart: shared.WeekStartSu.ToPointer(),
                },
                shared.CalendarEventRecurrence{
                    Count: unifiedgosdk.Pointer[float64](9.0),
                    EndAt: types.MustNewTimeFromString("2025-04-30T05:52:39.892Z"),
                    ExcludedDates: []string{
                        "2020-04-29T00:43:21.546Z",
                    },
                    Frequency: shared.CalendarEventRecurrenceFrequencyDaily.ToPointer(),
                    IncludedDates: []string{
                        "2020-09-11T02:34:17.920Z",
                        "2021-11-29T01:27:53.247Z",
                        "2019-12-22T18:06:05.561Z",
                    },
                    Interval: unifiedgosdk.Pointer[float64](1.0),
                    OnDays: []shared.PropertyCalendarEventRecurrenceOnDays{
                        shared.PropertyCalendarEventRecurrenceOnDaysWe,
                        shared.PropertyCalendarEventRecurrenceOnDaysTu,
                        shared.PropertyCalendarEventRecurrenceOnDaysWe,
                        shared.PropertyCalendarEventRecurrenceOnDaysSa,
                        shared.PropertyCalendarEventRecurrenceOnDaysSa,
                        shared.PropertyCalendarEventRecurrenceOnDaysSa,
                    },
                    OnMonthDays: []float64{
                        1.0,
                    },
                    OnMonths: []float64{
                        4.0,
                        0.0,
                        -3.0,
                    },
                    OnWeeks: []float64{
                        -7.0,
                        -19.0,
                        50.0,
                        -37.0,
                        43.0,
                        -48.0,
                        -30.0,
                        34.0,
                        36.0,
                        -33.0,
                        24.0,
                        -4.0,
                    },
                    OnYearDays: []float64{
                        277.0,
                        -115.0,
                        100.0,
                        2.0,
                        81.0,
                        -66.0,
                        31.0,
                        -39.0,
                        -319.0,
                        -251.0,
                        -254.0,
                        -35.0,
                        -121.0,
                        262.0,
                        32.0,
                        190.0,
                        107.0,
                        -145.0,
                        91.0,
                        313.0,
                        -48.0,
                        277.0,
                        104.0,
                        342.0,
                        297.0,
                        -216.0,
                        346.0,
                        -257.0,
                        307.0,
                        -44.0,
                        264.0,
                        -153.0,
                        -268.0,
                        92.0,
                        152.0,
                        -182.0,
                        -334.0,
                        89.0,
                        343.0,
                        -320.0,
                        -36.0,
                        84.0,
                        340.0,
                        -88.0,
                        -278.0,
                        202.0,
                        291.0,
                        95.0,
                        -234.0,
                        -304.0,
                        -157.0,
                        -82.0,
                        -339.0,
                        83.0,
                        2.0,
                        -238.0,
                        -204.0,
                        206.0,
                        -273.0,
                        -78.0,
                        -21.0,
                        270.0,
                        -266.0,
                        -276.0,
                        154.0,
                        -97.0,
                        -43.0,
                        -3.0,
                        191.0,
                        -302.0,
                        290.0,
                        -118.0,
                        -125.0,
                        -294.0,
                        115.0,
                        -73.0,
                        -244.0,
                        127.0,
                        26.0,
                        251.0,
                        47.0,
                        -157.0,
                        22.0,
                        -361.0,
                        318.0,
                        352.0,
                        358.0,
                        167.0,
                        210.0,
                        -185.0,
                        327.0,
                        117.0,
                        350.0,
                        -170.0,
                        -144.0,
                        -14.0,
                        -37.0,
                        318.0,
                        243.0,
                        33.0,
                        90.0,
                        319.0,
                        -270.0,
                        229.0,
                        122.0,
                        287.0,
                        -90.0,
                        -69.0,
                        -134.0,
                        -184.0,
                        25.0,
                        -178.0,
                        -89.0,
                        -273.0,
                        -49.0,
                        -362.0,
                        -9.0,
                        -71.0,
                        -347.0,
                        353.0,
                        342.0,
                        133.0,
                        -116.0,
                        231.0,
                        -231.0,
                        51.0,
                        288.0,
                        186.0,
                        -328.0,
                        275.0,
                        81.0,
                        94.0,
                        -263.0,
                        114.0,
                        13.0,
                        -357.0,
                        171.0,
                        -242.0,
                        -85.0,
                        -362.0,
                        108.0,
                        164.0,
                        69.0,
                        15.0,
                        57.0,
                        -287.0,
                        100.0,
                        165.0,
                        205.0,
                        204.0,
                        -78.0,
                        360.0,
                        -80.0,
                        -120.0,
                        -255.0,
                        -77.0,
                        110.0,
                        -26.0,
                        -149.0,
                        -254.0,
                        95.0,
                        32.0,
                        -57.0,
                        -195.0,
                        100.0,
                        221.0,
                        74.0,
                        274.0,
                        15.0,
                        353.0,
                        204.0,
                        -365.0,
                        315.0,
                        344.0,
                        199.0,
                        -59.0,
                        272.0,
                        173.0,
                        -40.0,
                        -318.0,
                        -330.0,
                        -365.0,
                        -272.0,
                        -149.0,
                        -27.0,
                        -334.0,
                        -277.0,
                        344.0,
                        351.0,
                        -310.0,
                        264.0,
                        281.0,
                        176.0,
                        191.0,
                        -183.0,
                        288.0,
                        -112.0,
                        -55.0,
                        -166.0,
                        258.0,
                        194.0,
                        59.0,
                    },
                    Timezone: unifiedgosdk.Pointer("America/Guadeloupe"),
                    WeekStart: shared.WeekStartTu.ToPointer(),
                },
                shared.CalendarEventRecurrence{
                    Count: unifiedgosdk.Pointer[float64](1.0),
                    EndAt: types.MustNewTimeFromString("2020-11-04T17:53:56.521Z"),
                    ExcludedDates: []string{
                        "2023-01-11T21:34:37.570Z",
                        "2021-09-07T13:21:42.135Z",
                    },
                    Frequency: shared.CalendarEventRecurrenceFrequencyWeekly.ToPointer(),
                    IncludedDates: []string{
                        "2024-08-31T10:58:44.165Z",
                    },
                    Interval: unifiedgosdk.Pointer[float64](9.0),
                    OnDays: []shared.PropertyCalendarEventRecurrenceOnDays{
                        shared.PropertyCalendarEventRecurrenceOnDaysTu,
                        shared.PropertyCalendarEventRecurrenceOnDaysSa,
                    },
                    OnMonthDays: []float64{
                        -2.0,
                    },
                    OnMonths: []float64{
                        -4.0,
                        8.0,
                        0.0,
                        9.0,
                        4.0,
                        -11.0,
                        7.0,
                        1.0,
                        -5.0,
                    },
                    OnWeeks: []float64{
                        -36.0,
                        -31.0,
                        -16.0,
                        -6.0,
                        44.0,
                        -37.0,
                        14.0,
                        38.0,
                        -27.0,
                        -22.0,
                        -2.0,
                        24.0,
                        7.0,
                        50.0,
                        46.0,
                        52.0,
                        20.0,
                        37.0,
                        31.0,
                        48.0,
                        35.0,
                        -46.0,
                        13.0,
                        22.0,
                        53.0,
                        20.0,
                        -28.0,
                        -2.0,
                        39.0,
                        13.0,
                        4.0,
                        0.0,
                        7.0,
                        -38.0,
                        -35.0,
                        41.0,
                        49.0,
                        12.0,
                        17.0,
                        8.0,
                        49.0,
                        -47.0,
                        46.0,
                        25.0,
                        14.0,
                        -26.0,
                        -37.0,
                        -25.0,
                        -41.0,
                        27.0,
                        28.0,
                        -19.0,
                    },
                    OnYearDays: []float64{
                        -256.0,
                        -328.0,
                        -312.0,
                        50.0,
                        -251.0,
                        -338.0,
                        -315.0,
                        214.0,
                        129.0,
                        -263.0,
                        -108.0,
                        -11.0,
                        206.0,
                        -29.0,
                        -159.0,
                        -29.0,
                        -264.0,
                        295.0,
                        -231.0,
                        53.0,
                        34.0,
                        -366.0,
                        326.0,
                        -202.0,
                        151.0,
                        79.0,
                        -66.0,
                        11.0,
                        -42.0,
                        73.0,
                        338.0,
                        -155.0,
                        197.0,
                        260.0,
                        356.0,
                        -323.0,
                        -213.0,
                        -332.0,
                        -305.0,
                        -182.0,
                        -253.0,
                        -276.0,
                        -285.0,
                        96.0,
                        -336.0,
                        269.0,
                        -233.0,
                        250.0,
                        -112.0,
                        -307.0,
                        -96.0,
                        54.0,
                        267.0,
                        318.0,
                        -66.0,
                        11.0,
                        -303.0,
                        231.0,
                        165.0,
                        -297.0,
                        -348.0,
                        -355.0,
                        364.0,
                        312.0,
                        -26.0,
                        111.0,
                        162.0,
                        280.0,
                        312.0,
                        337.0,
                        235.0,
                        68.0,
                        -282.0,
                        363.0,
                        212.0,
                        -328.0,
                        9.0,
                        -24.0,
                        -163.0,
                        -101.0,
                        -79.0,
                        -264.0,
                        -157.0,
                        188.0,
                        290.0,
                        51.0,
                        -213.0,
                        216.0,
                        230.0,
                        -270.0,
                        -211.0,
                        -156.0,
                        -165.0,
                        -305.0,
                        -45.0,
                        224.0,
                        -248.0,
                        65.0,
                        9.0,
                        274.0,
                        -299.0,
                        -228.0,
                        33.0,
                        -42.0,
                        356.0,
                        -311.0,
                        241.0,
                        261.0,
                        -136.0,
                        -252.0,
                        166.0,
                        208.0,
                        -126.0,
                        64.0,
                        323.0,
                        -104.0,
                        -106.0,
                        -248.0,
                        -41.0,
                        -109.0,
                        245.0,
                        47.0,
                        205.0,
                        358.0,
                        -296.0,
                        214.0,
                        -157.0,
                        -313.0,
                        -303.0,
                        -54.0,
                        -229.0,
                        231.0,
                        -94.0,
                        -198.0,
                        338.0,
                        199.0,
                        5.0,
                        42.0,
                        309.0,
                        73.0,
                        56.0,
                        -120.0,
                        351.0,
                        6.0,
                        -193.0,
                        21.0,
                        78.0,
                        57.0,
                        -269.0,
                        -76.0,
                        -299.0,
                        295.0,
                        -278.0,
                        11.0,
                        121.0,
                        -323.0,
                        156.0,
                        67.0,
                        152.0,
                        284.0,
                        108.0,
                        -7.0,
                        329.0,
                        -32.0,
                        333.0,
                        -338.0,
                        148.0,
                        -42.0,
                        151.0,
                        145.0,
                        -34.0,
                        -36.0,
                        296.0,
                        -198.0,
                        -317.0,
                        -161.0,
                        -253.0,
                        328.0,
                        -57.0,
                        134.0,
                        -289.0,
                        229.0,
                        44.0,
                        16.0,
                        -256.0,
                        289.0,
                        -234.0,
                        197.0,
                        333.0,
                        228.0,
                        -143.0,
                        -202.0,
                        -172.0,
                        -262.0,
                        -203.0,
                        -83.0,
                        -242.0,
                        -173.0,
                        336.0,
                        298.0,
                        -319.0,
                        66.0,
                        254.0,
                        214.0,
                        -118.0,
                        -216.0,
                        -168.0,
                        44.0,
                        -243.0,
                        207.0,
                        -28.0,
                        -4.0,
                        -272.0,
                        79.0,
                    },
                    Timezone: unifiedgosdk.Pointer("Atlantic/Reykjavik"),
                    WeekStart: shared.WeekStartTu.ToPointer(),
                },
            },
            RecurringEventID: unifiedgosdk.Pointer("cf746b0c-9906-4960-97c6-ccd79704eb13"),
            SendNotifications: unifiedgosdk.Pointer(false),
            StartAt: unifiedgosdk.Pointer("2020-05-20T08:05:45.142Z"),
            Status: shared.CalendarEventStatusConfirmed.ToPointer(),
            Subject: unifiedgosdk.Pointer("Sunt spargo tepidus bestia vigor credo coadunatio appello."),
            Timezone: unifiedgosdk.Pointer("Asia/Bangkok"),
            UpdatedAt: unifiedgosdk.Pointer("2020-06-26T04:47:47.420Z"),
            WebURL: unifiedgosdk.Pointer("https://another-pinstripe.com"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CalendarEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PatchCalendarEventRequest](../../pkg/models/operations/patchcalendareventrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.PatchCalendarEventResponse](../../pkg/models/operations/patchcalendareventresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchCalendarLink

Update a link

### Example Usage

<!-- UsageSnippet language="go" operationID="patchCalendarLink" method="patch" path="/calendar/{connection_id}/link/{id}" example="calendar_link" -->
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

    res, err := s.Calendar.PatchCalendarLink(ctx, operations.PatchCalendarLinkRequest{
        CalendarLink: shared.CalendarLink{
            CreatedAt: unifiedgosdk.Pointer("2023-03-07T13:34:11.959Z"),
            Description: unifiedgosdk.Pointer("Vitium clibanus laboriosam uxor denuncio."),
            Duration: unifiedgosdk.Pointer[float64](74.0),
            ID: unifiedgosdk.Pointer("7ee7d961-69a9-4d0d-abc8-a952c590c9e6"),
            IsActive: unifiedgosdk.Pointer(true),
            Name: unifiedgosdk.Pointer("Sopor sopor ancilla animus anser dignissimos vito confero utilis."),
            PriceAmount: unifiedgosdk.Pointer[float64](44.0),
            PriceCurrency: unifiedgosdk.Pointer("USD"),
            UpdatedAt: unifiedgosdk.Pointer("2024-03-06T11:31:30.146Z"),
            URL: "https://annual-apricot.info/",
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CalendarLink != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PatchCalendarLinkRequest](../../pkg/models/operations/patchcalendarlinkrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PatchCalendarLinkResponse](../../pkg/models/operations/patchcalendarlinkresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchCalendarWebinar

Update a webinar

### Example Usage

<!-- UsageSnippet language="go" operationID="patchCalendarWebinar" method="patch" path="/calendar/{connection_id}/webinar/{id}" example="calendar_webinar" -->
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

    res, err := s.Calendar.PatchCalendarWebinar(ctx, operations.PatchCalendarWebinarRequest{
        CalendarWebinar: shared.CalendarWebinar{
            Conference: []shared.CalendarConference{},
            CreatedAt: types.MustNewTimeFromString("2022-07-06T11:45:14.631Z"),
            EndAt: types.MustNewTimeFromString("2025-10-03T23:05:05.720Z"),
            HasPolls: unifiedgosdk.Pointer(false),
            HasRecording: unifiedgosdk.Pointer(false),
            ID: unifiedgosdk.Pointer("d6c98c6e-b43a-42f2-bc60-51849586d7bb"),
            IsAutoApprove: unifiedgosdk.Pointer(false),
            IsEnabled: unifiedgosdk.Pointer(true),
            IsWebcast: unifiedgosdk.Pointer(false),
            JoinURL: unifiedgosdk.Pointer("https://robust-bathhouse.biz"),
            Notes: unifiedgosdk.Pointer("Curriculum ducimus assentator aspernatur ait."),
            Organizer: &shared.PropertyCalendarWebinarOrganizer{
                Email: unifiedgosdk.Pointer("Kelton_Dicki@yahoo.com"),
                Name: unifiedgosdk.Pointer("Walter Greenfelder"),
            },
            Recurrence: []shared.CalendarEventRecurrence{
                shared.CalendarEventRecurrence{
                    Count: unifiedgosdk.Pointer[float64](10.0),
                    EndAt: types.MustNewTimeFromString("2023-08-23T00:00:57.829Z"),
                    ExcludedDates: []string{
                        "2025-01-24T12:51:50.276Z",
                    },
                    Frequency: shared.CalendarEventRecurrenceFrequencyMonthly.ToPointer(),
                    IncludedDates: []string{
                        "2024-04-14T17:49:13.802Z",
                    },
                    Interval: unifiedgosdk.Pointer[float64](8.0),
                    OnDays: []shared.PropertyCalendarEventRecurrenceOnDays{
                        shared.PropertyCalendarEventRecurrenceOnDaysSu,
                        shared.PropertyCalendarEventRecurrenceOnDaysFr,
                        shared.PropertyCalendarEventRecurrenceOnDaysSa,
                        shared.PropertyCalendarEventRecurrenceOnDaysWe,
                        shared.PropertyCalendarEventRecurrenceOnDaysMo,
                    },
                    OnMonthDays: []float64{
                        -10.0,
                    },
                    OnMonths: []float64{
                        -9.0,
                    },
                    OnWeeks: []float64{
                        10.0,
                        30.0,
                        -38.0,
                        30.0,
                        -22.0,
                        37.0,
                        -12.0,
                        27.0,
                        2.0,
                        15.0,
                        26.0,
                        18.0,
                        -43.0,
                        -33.0,
                        -27.0,
                        38.0,
                        28.0,
                        47.0,
                        -8.0,
                        24.0,
                        35.0,
                        -2.0,
                        7.0,
                        49.0,
                        38.0,
                        -41.0,
                        46.0,
                        -11.0,
                        -45.0,
                        0.0,
                        48.0,
                        34.0,
                    },
                    OnYearDays: []float64{
                        345.0,
                        -207.0,
                        230.0,
                        -10.0,
                        364.0,
                        -256.0,
                        -218.0,
                        -295.0,
                        290.0,
                        -250.0,
                        -315.0,
                        60.0,
                        205.0,
                        -247.0,
                        -318.0,
                        -211.0,
                        -13.0,
                        256.0,
                        -200.0,
                        -313.0,
                        336.0,
                        -332.0,
                        -90.0,
                        287.0,
                        -273.0,
                        156.0,
                        241.0,
                        -138.0,
                        -363.0,
                        -37.0,
                        -171.0,
                        -62.0,
                        -57.0,
                        280.0,
                        -322.0,
                        -79.0,
                        -364.0,
                        -201.0,
                        84.0,
                        341.0,
                        334.0,
                        -75.0,
                        332.0,
                        207.0,
                        337.0,
                        -244.0,
                        131.0,
                        -191.0,
                        164.0,
                        -235.0,
                        285.0,
                        -309.0,
                        -158.0,
                        306.0,
                        180.0,
                        -130.0,
                        -162.0,
                        -155.0,
                        3.0,
                        198.0,
                        26.0,
                        -366.0,
                        -191.0,
                        127.0,
                        -331.0,
                        -11.0,
                        -239.0,
                        -189.0,
                        243.0,
                        118.0,
                        346.0,
                        -174.0,
                        -146.0,
                        -161.0,
                        -330.0,
                        327.0,
                        192.0,
                        310.0,
                        316.0,
                        313.0,
                        -242.0,
                        -51.0,
                        -264.0,
                        -180.0,
                        -88.0,
                        305.0,
                        270.0,
                        358.0,
                        -173.0,
                        -298.0,
                        153.0,
                        -89.0,
                        155.0,
                        -45.0,
                        248.0,
                        -46.0,
                        -146.0,
                        300.0,
                        364.0,
                        -335.0,
                        356.0,
                        -18.0,
                        219.0,
                        324.0,
                        -239.0,
                        -106.0,
                        -298.0,
                        328.0,
                        362.0,
                        344.0,
                        -54.0,
                        133.0,
                        50.0,
                        112.0,
                        -212.0,
                        -179.0,
                        22.0,
                        -201.0,
                        -62.0,
                        -293.0,
                        9.0,
                        30.0,
                        -50.0,
                        126.0,
                        -72.0,
                        264.0,
                        28.0,
                        -1.0,
                        -207.0,
                        160.0,
                        -168.0,
                        3.0,
                        -176.0,
                        -19.0,
                        -157.0,
                        349.0,
                        100.0,
                        -201.0,
                        108.0,
                        -180.0,
                        51.0,
                        -73.0,
                        366.0,
                        74.0,
                        -226.0,
                        238.0,
                        121.0,
                        -193.0,
                        -125.0,
                        -109.0,
                        316.0,
                        -177.0,
                        -307.0,
                        31.0,
                        -76.0,
                        217.0,
                        -310.0,
                        227.0,
                        -360.0,
                        71.0,
                        255.0,
                        -325.0,
                        -214.0,
                        40.0,
                        42.0,
                        17.0,
                        -241.0,
                        -84.0,
                        -188.0,
                        302.0,
                        64.0,
                        94.0,
                        -362.0,
                        23.0,
                        166.0,
                        85.0,
                        71.0,
                        -74.0,
                        -47.0,
                        -119.0,
                        98.0,
                        40.0,
                        158.0,
                        -64.0,
                        175.0,
                        269.0,
                        127.0,
                        -143.0,
                        213.0,
                        -196.0,
                        121.0,
                        81.0,
                        -238.0,
                        288.0,
                        321.0,
                        276.0,
                        133.0,
                        22.0,
                        -213.0,
                        -157.0,
                        -280.0,
                        -35.0,
                        73.0,
                        -194.0,
                        65.0,
                        -180.0,
                        63.0,
                        -242.0,
                        -117.0,
                        148.0,
                        157.0,
                        -320.0,
                        318.0,
                        8.0,
                        210.0,
                        -21.0,
                        81.0,
                        205.0,
                        -258.0,
                        -40.0,
                        -114.0,
                        -253.0,
                        -263.0,
                        65.0,
                        185.0,
                        -24.0,
                        324.0,
                        -172.0,
                        25.0,
                        260.0,
                        211.0,
                        342.0,
                        -31.0,
                        -288.0,
                        -159.0,
                        -4.0,
                        -2.0,
                        -107.0,
                        -316.0,
                        -276.0,
                        331.0,
                        -114.0,
                        -20.0,
                        -320.0,
                        51.0,
                        -176.0,
                        -148.0,
                        -50.0,
                        -201.0,
                        -104.0,
                        153.0,
                        -273.0,
                        -189.0,
                        67.0,
                        209.0,
                        149.0,
                        49.0,
                        -136.0,
                        -125.0,
                        -169.0,
                        -324.0,
                        309.0,
                        -51.0,
                        288.0,
                        253.0,
                        175.0,
                        -146.0,
                        171.0,
                        -140.0,
                        58.0,
                        -212.0,
                        164.0,
                        270.0,
                        102.0,
                        70.0,
                        299.0,
                        89.0,
                        -280.0,
                        252.0,
                        -342.0,
                        240.0,
                        226.0,
                        68.0,
                        -30.0,
                        -232.0,
                        -358.0,
                        -166.0,
                        60.0,
                        140.0,
                        275.0,
                        13.0,
                        250.0,
                        -328.0,
                        -189.0,
                        -22.0,
                        7.0,
                        -235.0,
                        -322.0,
                        178.0,
                        167.0,
                        -104.0,
                        -61.0,
                        282.0,
                        -80.0,
                        -277.0,
                        108.0,
                        271.0,
                        -237.0,
                        297.0,
                        -135.0,
                        -135.0,
                        -323.0,
                        342.0,
                        -267.0,
                        -235.0,
                        173.0,
                        249.0,
                        -288.0,
                        257.0,
                        139.0,
                        -191.0,
                        -217.0,
                        10.0,
                        -117.0,
                        -297.0,
                        -196.0,
                        -206.0,
                        341.0,
                        166.0,
                        181.0,
                        129.0,
                        -207.0,
                        55.0,
                        86.0,
                    },
                    Timezone: unifiedgosdk.Pointer("Asia/Ust-Nera"),
                    WeekStart: shared.WeekStartMo.ToPointer(),
                },
                shared.CalendarEventRecurrence{
                    Count: unifiedgosdk.Pointer[float64](3.0),
                    EndAt: types.MustNewTimeFromString("2022-09-28T21:54:22.888Z"),
                    ExcludedDates: []string{
                        "2024-08-16T15:01:59.509Z",
                        "2024-08-01T09:41:48.749Z",
                    },
                    Frequency: shared.CalendarEventRecurrenceFrequencyDaily.ToPointer(),
                    IncludedDates: []string{
                        "2024-03-12T07:59:01.474Z",
                        "2025-12-18T01:45:08.097Z",
                        "2023-08-06T00:06:02.458Z",
                    },
                    Interval: unifiedgosdk.Pointer[float64](1.0),
                    OnDays: []shared.PropertyCalendarEventRecurrenceOnDays{
                        shared.PropertyCalendarEventRecurrenceOnDaysWe,
                        shared.PropertyCalendarEventRecurrenceOnDaysSu,
                        shared.PropertyCalendarEventRecurrenceOnDaysMo,
                        shared.PropertyCalendarEventRecurrenceOnDaysFr,
                    },
                    OnMonthDays: []float64{
                        -15.0,
                    },
                    OnMonths: []float64{
                        5.0,
                        12.0,
                        3.0,
                        12.0,
                        8.0,
                    },
                    OnWeeks: []float64{
                        -47.0,
                        44.0,
                    },
                    OnYearDays: []float64{
                        -117.0,
                        59.0,
                        -6.0,
                        187.0,
                        45.0,
                        70.0,
                        15.0,
                        255.0,
                        44.0,
                        -2.0,
                        25.0,
                        -175.0,
                        -240.0,
                        171.0,
                        -294.0,
                        19.0,
                        38.0,
                        -351.0,
                        170.0,
                        -10.0,
                        -269.0,
                        18.0,
                        -65.0,
                        -266.0,
                        -31.0,
                        328.0,
                        -361.0,
                        358.0,
                        -256.0,
                        -4.0,
                        -312.0,
                        82.0,
                        -2.0,
                        -75.0,
                        -281.0,
                        -304.0,
                        53.0,
                        -295.0,
                        366.0,
                        322.0,
                        -191.0,
                        26.0,
                        97.0,
                        53.0,
                        75.0,
                        -62.0,
                        -109.0,
                        66.0,
                        177.0,
                        -68.0,
                        175.0,
                        -280.0,
                        70.0,
                        -238.0,
                        109.0,
                        -304.0,
                        326.0,
                        -8.0,
                        -71.0,
                        -236.0,
                        225.0,
                        358.0,
                        20.0,
                        -5.0,
                        -102.0,
                        -134.0,
                        -204.0,
                        -116.0,
                        -353.0,
                        -273.0,
                        106.0,
                        284.0,
                        -137.0,
                        -324.0,
                        301.0,
                        -42.0,
                        -229.0,
                        271.0,
                        -293.0,
                        -343.0,
                        211.0,
                        47.0,
                        -254.0,
                        -154.0,
                        -182.0,
                        264.0,
                        120.0,
                        -11.0,
                        -307.0,
                        99.0,
                        227.0,
                        190.0,
                        -17.0,
                        -77.0,
                        -255.0,
                        -61.0,
                        -249.0,
                        -102.0,
                        70.0,
                        345.0,
                        -187.0,
                        -308.0,
                        194.0,
                        221.0,
                        268.0,
                        -169.0,
                        -190.0,
                        88.0,
                        10.0,
                        262.0,
                        177.0,
                        -314.0,
                        -151.0,
                        -295.0,
                    },
                    Timezone: unifiedgosdk.Pointer("Pacific/Wake"),
                    WeekStart: shared.WeekStartTu.ToPointer(),
                },
                shared.CalendarEventRecurrence{
                    Count: unifiedgosdk.Pointer[float64](8.0),
                    EndAt: types.MustNewTimeFromString("2026-06-26T05:33:58.781Z"),
                    ExcludedDates: []string{
                        "2023-06-11T12:02:36.566Z",
                        "2023-05-31T18:16:08.923Z",
                    },
                    Frequency: shared.CalendarEventRecurrenceFrequencyWeekly.ToPointer(),
                    IncludedDates: []string{
                        "2024-03-20T04:54:34.101Z",
                        "2023-08-11T16:40:30.431Z",
                        "2024-09-10T07:26:29.395Z",
                    },
                    Interval: unifiedgosdk.Pointer[float64](8.0),
                    OnDays: []shared.PropertyCalendarEventRecurrenceOnDays{
                        shared.PropertyCalendarEventRecurrenceOnDaysSu,
                        shared.PropertyCalendarEventRecurrenceOnDaysMo,
                        shared.PropertyCalendarEventRecurrenceOnDaysTu,
                        shared.PropertyCalendarEventRecurrenceOnDaysFr,
                        shared.PropertyCalendarEventRecurrenceOnDaysMo,
                        shared.PropertyCalendarEventRecurrenceOnDaysTh,
                    },
                    OnMonthDays: []float64{
                        -23.0,
                    },
                    OnMonths: []float64{
                        11.0,
                        8.0,
                        9.0,
                        5.0,
                        -12.0,
                        -7.0,
                        -5.0,
                        10.0,
                        10.0,
                        -9.0,
                        -10.0,
                    },
                    OnWeeks: []float64{
                        -49.0,
                        46.0,
                        35.0,
                        -26.0,
                        2.0,
                        15.0,
                        15.0,
                        -26.0,
                        24.0,
                        -53.0,
                        36.0,
                        -43.0,
                        51.0,
                        -19.0,
                        -7.0,
                        -12.0,
                        28.0,
                        27.0,
                        35.0,
                        12.0,
                        -28.0,
                        -8.0,
                        -4.0,
                        -45.0,
                    },
                    OnYearDays: []float64{
                        84.0,
                        -251.0,
                        71.0,
                        181.0,
                        -163.0,
                        158.0,
                        301.0,
                        -299.0,
                        -184.0,
                        -331.0,
                        -152.0,
                        -129.0,
                        -237.0,
                        -303.0,
                        -24.0,
                        126.0,
                        -103.0,
                        146.0,
                        -346.0,
                        86.0,
                        -296.0,
                        -337.0,
                        -185.0,
                        16.0,
                        -270.0,
                        -126.0,
                        -295.0,
                        -231.0,
                        356.0,
                        -293.0,
                        115.0,
                        -265.0,
                        -293.0,
                        -34.0,
                        357.0,
                        313.0,
                        -343.0,
                        180.0,
                        -22.0,
                        -161.0,
                        350.0,
                        177.0,
                        190.0,
                        223.0,
                        -152.0,
                        -360.0,
                        -225.0,
                        -60.0,
                        -35.0,
                        353.0,
                        117.0,
                        -171.0,
                        -315.0,
                        -321.0,
                        -202.0,
                        345.0,
                        -1.0,
                        -148.0,
                        -168.0,
                        181.0,
                        -17.0,
                        282.0,
                        234.0,
                        31.0,
                        47.0,
                        -236.0,
                        366.0,
                        -251.0,
                        -232.0,
                        -308.0,
                        76.0,
                        -199.0,
                        184.0,
                        198.0,
                        225.0,
                        75.0,
                        6.0,
                        227.0,
                        -148.0,
                        259.0,
                        -146.0,
                        49.0,
                        -254.0,
                        341.0,
                        93.0,
                        138.0,
                        -164.0,
                        237.0,
                        4.0,
                        -287.0,
                        161.0,
                    },
                    Timezone: unifiedgosdk.Pointer("Africa/Bissau"),
                    WeekStart: shared.WeekStartWe.ToPointer(),
                },
            },
            RegistrantPassword: unifiedgosdk.Pointer("OxwWzr0C"),
            RequireAddress: unifiedgosdk.Pointer(false),
            RequireJobTitle: unifiedgosdk.Pointer(false),
            StartAt: types.MustNewTimeFromString("2025-04-09T12:29:18.755Z"),
            Status: shared.CalendarWebinarStatusTentative.ToPointer(),
            Subject: unifiedgosdk.Pointer("Harum culpa decipio ex cubo ancilla cresco."),
            Timezone: unifiedgosdk.Pointer("Europe/Kaliningrad"),
            UpdatedAt: types.MustNewTimeFromString("2026-08-29T20:26:31.460Z"),
            WebURL: unifiedgosdk.Pointer("https://classic-recovery.biz"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CalendarWebinar != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PatchCalendarWebinarRequest](../../pkg/models/operations/patchcalendarwebinarrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.PatchCalendarWebinarResponse](../../pkg/models/operations/patchcalendarwebinarresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveCalendarCalendar

Remove a calendar

### Example Usage

<!-- UsageSnippet language="go" operationID="removeCalendarCalendar" method="delete" path="/calendar/{connection_id}/calendar/{id}" -->
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

    res, err := s.Calendar.RemoveCalendarCalendar(ctx, operations.RemoveCalendarCalendarRequest{
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.RemoveCalendarCalendarRequest](../../pkg/models/operations/removecalendarcalendarrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.RemoveCalendarCalendarResponse](../../pkg/models/operations/removecalendarcalendarresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveCalendarEvent

Remove an event

### Example Usage

<!-- UsageSnippet language="go" operationID="removeCalendarEvent" method="delete" path="/calendar/{connection_id}/event/{id}" -->
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

    res, err := s.Calendar.RemoveCalendarEvent(ctx, operations.RemoveCalendarEventRequest{
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
| `request`                                                                                          | [operations.RemoveCalendarEventRequest](../../pkg/models/operations/removecalendareventrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.RemoveCalendarEventResponse](../../pkg/models/operations/removecalendareventresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveCalendarLink

Remove a link

### Example Usage

<!-- UsageSnippet language="go" operationID="removeCalendarLink" method="delete" path="/calendar/{connection_id}/link/{id}" -->
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

    res, err := s.Calendar.RemoveCalendarLink(ctx, operations.RemoveCalendarLinkRequest{
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.RemoveCalendarLinkRequest](../../pkg/models/operations/removecalendarlinkrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.RemoveCalendarLinkResponse](../../pkg/models/operations/removecalendarlinkresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveCalendarWebinar

Remove a webinar

### Example Usage

<!-- UsageSnippet language="go" operationID="removeCalendarWebinar" method="delete" path="/calendar/{connection_id}/webinar/{id}" -->
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

    res, err := s.Calendar.RemoveCalendarWebinar(ctx, operations.RemoveCalendarWebinarRequest{
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
| `request`                                                                                              | [operations.RemoveCalendarWebinarRequest](../../pkg/models/operations/removecalendarwebinarrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.RemoveCalendarWebinarResponse](../../pkg/models/operations/removecalendarwebinarresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateCalendarCalendar

Update a calendar

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCalendarCalendar" method="put" path="/calendar/{connection_id}/calendar/{id}" example="calendar_calendar" -->
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

    res, err := s.Calendar.UpdateCalendarCalendar(ctx, operations.UpdateCalendarCalendarRequest{
        CalendarCalendar: shared.CalendarCalendar{
            CreatedAt: types.MustNewTimeFromString("2020-01-09T23:11:34.147Z"),
            Description: unifiedgosdk.Pointer("Socius catena auxilium."),
            ID: unifiedgosdk.Pointer("c8bd690d-d2e7-4367-a8d9-54cb99b57437"),
            IsPrimary: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("Acer supra vallum suasoria thesaurus omnis condico cognomen accendo vehemens."),
            Timezone: unifiedgosdk.Pointer("America/Dawson_Creek"),
            UpdatedAt: types.MustNewTimeFromString("2023-03-12T22:44:44.999Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CalendarCalendar != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.UpdateCalendarCalendarRequest](../../pkg/models/operations/updatecalendarcalendarrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.UpdateCalendarCalendarResponse](../../pkg/models/operations/updatecalendarcalendarresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateCalendarEvent

Update an event

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCalendarEvent" method="put" path="/calendar/{connection_id}/event/{id}" example="calendar_event" -->
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

    res, err := s.Calendar.UpdateCalendarEvent(ctx, operations.UpdateCalendarEventRequest{
        CalendarEvent: shared.CalendarEvent{
            Attachments: []shared.CalendarAttachment{},
            Conference: []shared.CalendarConference{},
            CreatedAt: unifiedgosdk.Pointer("2019-08-04T14:33:51.814Z"),
            EndAt: unifiedgosdk.Pointer("2020-05-20T17:12:37.979Z"),
            ID: unifiedgosdk.Pointer("67215e17-405a-4402-bfdf-b40ac9978bb3"),
            IsAllDay: unifiedgosdk.Pointer(false),
            IsFree: unifiedgosdk.Pointer(false),
            IsPrivate: unifiedgosdk.Pointer(false),
            Location: unifiedgosdk.Pointer("621 Boehm Prairie"),
            Notes: unifiedgosdk.Pointer("Aegre traho."),
            Recurrence: []shared.CalendarEventRecurrence{
                shared.CalendarEventRecurrence{
                    Count: unifiedgosdk.Pointer[float64](8.0),
                    EndAt: types.MustNewTimeFromString("2025-12-26T06:34:54.364Z"),
                    ExcludedDates: []string{
                        "2025-09-30T22:42:44.936Z",
                        "2023-10-09T12:14:01.324Z",
                        "2024-02-15T13:20:24.736Z",
                    },
                    Frequency: shared.CalendarEventRecurrenceFrequencyMonthly.ToPointer(),
                    IncludedDates: []string{
                        "2021-02-16T22:04:02.467Z",
                    },
                    Interval: unifiedgosdk.Pointer[float64](4.0),
                    OnDays: []shared.PropertyCalendarEventRecurrenceOnDays{
                        shared.PropertyCalendarEventRecurrenceOnDaysTh,
                        shared.PropertyCalendarEventRecurrenceOnDaysMo,
                        shared.PropertyCalendarEventRecurrenceOnDaysTh,
                    },
                    OnMonthDays: []float64{
                        -26.0,
                    },
                    OnMonths: []float64{
                        12.0,
                        9.0,
                        -1.0,
                        0.0,
                        1.0,
                        6.0,
                        -10.0,
                        9.0,
                        0.0,
                        4.0,
                        -2.0,
                    },
                    OnWeeks: []float64{
                        -7.0,
                        51.0,
                        -3.0,
                        -41.0,
                        15.0,
                        46.0,
                        -1.0,
                        46.0,
                        42.0,
                        11.0,
                        12.0,
                        -35.0,
                        -15.0,
                        -3.0,
                        -42.0,
                        50.0,
                        3.0,
                        -15.0,
                        -10.0,
                        6.0,
                        -53.0,
                        5.0,
                        -32.0,
                        -22.0,
                        43.0,
                        -44.0,
                        -23.0,
                        -21.0,
                        -18.0,
                    },
                    OnYearDays: []float64{
                        -35.0,
                        14.0,
                        -338.0,
                        175.0,
                        -87.0,
                        339.0,
                        341.0,
                        287.0,
                        -17.0,
                        319.0,
                        -3.0,
                        238.0,
                        -115.0,
                        -116.0,
                        283.0,
                        -61.0,
                        -254.0,
                        86.0,
                        -163.0,
                        5.0,
                        -171.0,
                        -99.0,
                        279.0,
                        19.0,
                        303.0,
                        -106.0,
                        90.0,
                        109.0,
                        -185.0,
                        -285.0,
                        -83.0,
                        -236.0,
                        66.0,
                        -215.0,
                        178.0,
                        64.0,
                        78.0,
                        5.0,
                        -251.0,
                        -79.0,
                        -271.0,
                        33.0,
                        320.0,
                        67.0,
                        -84.0,
                        -355.0,
                        -364.0,
                        348.0,
                        271.0,
                        -304.0,
                        -199.0,
                        106.0,
                        -345.0,
                        24.0,
                        -89.0,
                        -109.0,
                        -314.0,
                        365.0,
                        38.0,
                        -42.0,
                        123.0,
                        56.0,
                        -3.0,
                        31.0,
                        101.0,
                        326.0,
                        -160.0,
                        -101.0,
                        -267.0,
                        -309.0,
                        -363.0,
                        125.0,
                        -182.0,
                        363.0,
                        324.0,
                        36.0,
                        -269.0,
                        -79.0,
                        -60.0,
                        272.0,
                        -254.0,
                        -160.0,
                        -82.0,
                        19.0,
                        42.0,
                        69.0,
                        -104.0,
                        333.0,
                        236.0,
                        -287.0,
                        296.0,
                        261.0,
                        241.0,
                        348.0,
                        -72.0,
                        159.0,
                        -127.0,
                        229.0,
                        -158.0,
                        190.0,
                        -173.0,
                        -84.0,
                        -96.0,
                        176.0,
                        339.0,
                        -48.0,
                        287.0,
                        -46.0,
                        -101.0,
                        246.0,
                        -8.0,
                        -74.0,
                        338.0,
                        -51.0,
                        -42.0,
                        -128.0,
                        -169.0,
                        -174.0,
                        168.0,
                        -85.0,
                        37.0,
                        169.0,
                        -105.0,
                        231.0,
                        -250.0,
                        -286.0,
                        -7.0,
                        -121.0,
                        321.0,
                        278.0,
                        -120.0,
                        -96.0,
                        360.0,
                        337.0,
                        -258.0,
                        -179.0,
                        324.0,
                        -204.0,
                        327.0,
                        15.0,
                        365.0,
                        191.0,
                        -345.0,
                        -345.0,
                        56.0,
                        217.0,
                        60.0,
                        -264.0,
                        -248.0,
                        -316.0,
                        191.0,
                        -189.0,
                        -152.0,
                        -296.0,
                        194.0,
                        -42.0,
                        -21.0,
                        -218.0,
                        171.0,
                        -15.0,
                        301.0,
                        37.0,
                        -167.0,
                        18.0,
                        248.0,
                        -263.0,
                        27.0,
                        14.0,
                        59.0,
                        219.0,
                        -284.0,
                        221.0,
                        -76.0,
                        277.0,
                        183.0,
                        200.0,
                        -12.0,
                        -28.0,
                        -79.0,
                        150.0,
                        320.0,
                        -152.0,
                        -15.0,
                        -42.0,
                        -125.0,
                        -4.0,
                        269.0,
                        290.0,
                        52.0,
                        320.0,
                        344.0,
                        13.0,
                        -69.0,
                        255.0,
                        -154.0,
                        -281.0,
                        158.0,
                        25.0,
                        240.0,
                        -339.0,
                        96.0,
                        204.0,
                        324.0,
                        221.0,
                        37.0,
                        -333.0,
                        87.0,
                        354.0,
                        -365.0,
                        -203.0,
                        -341.0,
                        -79.0,
                        -208.0,
                        135.0,
                        132.0,
                        -351.0,
                        39.0,
                        -87.0,
                        -297.0,
                        -66.0,
                        346.0,
                        69.0,
                        -177.0,
                        235.0,
                        295.0,
                        -366.0,
                        -55.0,
                    },
                    Timezone: unifiedgosdk.Pointer("Asia/Ho_Chi_Minh"),
                    WeekStart: shared.WeekStartSu.ToPointer(),
                },
                shared.CalendarEventRecurrence{
                    Count: unifiedgosdk.Pointer[float64](9.0),
                    EndAt: types.MustNewTimeFromString("2025-04-30T05:52:39.892Z"),
                    ExcludedDates: []string{
                        "2020-04-29T00:43:21.546Z",
                    },
                    Frequency: shared.CalendarEventRecurrenceFrequencyDaily.ToPointer(),
                    IncludedDates: []string{
                        "2020-09-11T02:34:17.920Z",
                        "2021-11-29T01:27:53.247Z",
                        "2019-12-22T18:06:05.561Z",
                    },
                    Interval: unifiedgosdk.Pointer[float64](1.0),
                    OnDays: []shared.PropertyCalendarEventRecurrenceOnDays{
                        shared.PropertyCalendarEventRecurrenceOnDaysWe,
                        shared.PropertyCalendarEventRecurrenceOnDaysTu,
                        shared.PropertyCalendarEventRecurrenceOnDaysWe,
                        shared.PropertyCalendarEventRecurrenceOnDaysSa,
                        shared.PropertyCalendarEventRecurrenceOnDaysSa,
                        shared.PropertyCalendarEventRecurrenceOnDaysSa,
                    },
                    OnMonthDays: []float64{
                        1.0,
                    },
                    OnMonths: []float64{
                        4.0,
                        0.0,
                        -3.0,
                    },
                    OnWeeks: []float64{
                        -7.0,
                        -19.0,
                        50.0,
                        -37.0,
                        43.0,
                        -48.0,
                        -30.0,
                        34.0,
                        36.0,
                        -33.0,
                        24.0,
                        -4.0,
                    },
                    OnYearDays: []float64{
                        277.0,
                        -115.0,
                        100.0,
                        2.0,
                        81.0,
                        -66.0,
                        31.0,
                        -39.0,
                        -319.0,
                        -251.0,
                        -254.0,
                        -35.0,
                        -121.0,
                        262.0,
                        32.0,
                        190.0,
                        107.0,
                        -145.0,
                        91.0,
                        313.0,
                        -48.0,
                        277.0,
                        104.0,
                        342.0,
                        297.0,
                        -216.0,
                        346.0,
                        -257.0,
                        307.0,
                        -44.0,
                        264.0,
                        -153.0,
                        -268.0,
                        92.0,
                        152.0,
                        -182.0,
                        -334.0,
                        89.0,
                        343.0,
                        -320.0,
                        -36.0,
                        84.0,
                        340.0,
                        -88.0,
                        -278.0,
                        202.0,
                        291.0,
                        95.0,
                        -234.0,
                        -304.0,
                        -157.0,
                        -82.0,
                        -339.0,
                        83.0,
                        2.0,
                        -238.0,
                        -204.0,
                        206.0,
                        -273.0,
                        -78.0,
                        -21.0,
                        270.0,
                        -266.0,
                        -276.0,
                        154.0,
                        -97.0,
                        -43.0,
                        -3.0,
                        191.0,
                        -302.0,
                        290.0,
                        -118.0,
                        -125.0,
                        -294.0,
                        115.0,
                        -73.0,
                        -244.0,
                        127.0,
                        26.0,
                        251.0,
                        47.0,
                        -157.0,
                        22.0,
                        -361.0,
                        318.0,
                        352.0,
                        358.0,
                        167.0,
                        210.0,
                        -185.0,
                        327.0,
                        117.0,
                        350.0,
                        -170.0,
                        -144.0,
                        -14.0,
                        -37.0,
                        318.0,
                        243.0,
                        33.0,
                        90.0,
                        319.0,
                        -270.0,
                        229.0,
                        122.0,
                        287.0,
                        -90.0,
                        -69.0,
                        -134.0,
                        -184.0,
                        25.0,
                        -178.0,
                        -89.0,
                        -273.0,
                        -49.0,
                        -362.0,
                        -9.0,
                        -71.0,
                        -347.0,
                        353.0,
                        342.0,
                        133.0,
                        -116.0,
                        231.0,
                        -231.0,
                        51.0,
                        288.0,
                        186.0,
                        -328.0,
                        275.0,
                        81.0,
                        94.0,
                        -263.0,
                        114.0,
                        13.0,
                        -357.0,
                        171.0,
                        -242.0,
                        -85.0,
                        -362.0,
                        108.0,
                        164.0,
                        69.0,
                        15.0,
                        57.0,
                        -287.0,
                        100.0,
                        165.0,
                        205.0,
                        204.0,
                        -78.0,
                        360.0,
                        -80.0,
                        -120.0,
                        -255.0,
                        -77.0,
                        110.0,
                        -26.0,
                        -149.0,
                        -254.0,
                        95.0,
                        32.0,
                        -57.0,
                        -195.0,
                        100.0,
                        221.0,
                        74.0,
                        274.0,
                        15.0,
                        353.0,
                        204.0,
                        -365.0,
                        315.0,
                        344.0,
                        199.0,
                        -59.0,
                        272.0,
                        173.0,
                        -40.0,
                        -318.0,
                        -330.0,
                        -365.0,
                        -272.0,
                        -149.0,
                        -27.0,
                        -334.0,
                        -277.0,
                        344.0,
                        351.0,
                        -310.0,
                        264.0,
                        281.0,
                        176.0,
                        191.0,
                        -183.0,
                        288.0,
                        -112.0,
                        -55.0,
                        -166.0,
                        258.0,
                        194.0,
                        59.0,
                    },
                    Timezone: unifiedgosdk.Pointer("America/Guadeloupe"),
                    WeekStart: shared.WeekStartTu.ToPointer(),
                },
                shared.CalendarEventRecurrence{
                    Count: unifiedgosdk.Pointer[float64](1.0),
                    EndAt: types.MustNewTimeFromString("2020-11-04T17:53:56.521Z"),
                    ExcludedDates: []string{
                        "2023-01-11T21:34:37.570Z",
                        "2021-09-07T13:21:42.135Z",
                    },
                    Frequency: shared.CalendarEventRecurrenceFrequencyWeekly.ToPointer(),
                    IncludedDates: []string{
                        "2024-08-31T10:58:44.165Z",
                    },
                    Interval: unifiedgosdk.Pointer[float64](9.0),
                    OnDays: []shared.PropertyCalendarEventRecurrenceOnDays{
                        shared.PropertyCalendarEventRecurrenceOnDaysTu,
                        shared.PropertyCalendarEventRecurrenceOnDaysSa,
                    },
                    OnMonthDays: []float64{
                        -2.0,
                    },
                    OnMonths: []float64{
                        -4.0,
                        8.0,
                        0.0,
                        9.0,
                        4.0,
                        -11.0,
                        7.0,
                        1.0,
                        -5.0,
                    },
                    OnWeeks: []float64{
                        -36.0,
                        -31.0,
                        -16.0,
                        -6.0,
                        44.0,
                        -37.0,
                        14.0,
                        38.0,
                        -27.0,
                        -22.0,
                        -2.0,
                        24.0,
                        7.0,
                        50.0,
                        46.0,
                        52.0,
                        20.0,
                        37.0,
                        31.0,
                        48.0,
                        35.0,
                        -46.0,
                        13.0,
                        22.0,
                        53.0,
                        20.0,
                        -28.0,
                        -2.0,
                        39.0,
                        13.0,
                        4.0,
                        0.0,
                        7.0,
                        -38.0,
                        -35.0,
                        41.0,
                        49.0,
                        12.0,
                        17.0,
                        8.0,
                        49.0,
                        -47.0,
                        46.0,
                        25.0,
                        14.0,
                        -26.0,
                        -37.0,
                        -25.0,
                        -41.0,
                        27.0,
                        28.0,
                        -19.0,
                    },
                    OnYearDays: []float64{
                        -256.0,
                        -328.0,
                        -312.0,
                        50.0,
                        -251.0,
                        -338.0,
                        -315.0,
                        214.0,
                        129.0,
                        -263.0,
                        -108.0,
                        -11.0,
                        206.0,
                        -29.0,
                        -159.0,
                        -29.0,
                        -264.0,
                        295.0,
                        -231.0,
                        53.0,
                        34.0,
                        -366.0,
                        326.0,
                        -202.0,
                        151.0,
                        79.0,
                        -66.0,
                        11.0,
                        -42.0,
                        73.0,
                        338.0,
                        -155.0,
                        197.0,
                        260.0,
                        356.0,
                        -323.0,
                        -213.0,
                        -332.0,
                        -305.0,
                        -182.0,
                        -253.0,
                        -276.0,
                        -285.0,
                        96.0,
                        -336.0,
                        269.0,
                        -233.0,
                        250.0,
                        -112.0,
                        -307.0,
                        -96.0,
                        54.0,
                        267.0,
                        318.0,
                        -66.0,
                        11.0,
                        -303.0,
                        231.0,
                        165.0,
                        -297.0,
                        -348.0,
                        -355.0,
                        364.0,
                        312.0,
                        -26.0,
                        111.0,
                        162.0,
                        280.0,
                        312.0,
                        337.0,
                        235.0,
                        68.0,
                        -282.0,
                        363.0,
                        212.0,
                        -328.0,
                        9.0,
                        -24.0,
                        -163.0,
                        -101.0,
                        -79.0,
                        -264.0,
                        -157.0,
                        188.0,
                        290.0,
                        51.0,
                        -213.0,
                        216.0,
                        230.0,
                        -270.0,
                        -211.0,
                        -156.0,
                        -165.0,
                        -305.0,
                        -45.0,
                        224.0,
                        -248.0,
                        65.0,
                        9.0,
                        274.0,
                        -299.0,
                        -228.0,
                        33.0,
                        -42.0,
                        356.0,
                        -311.0,
                        241.0,
                        261.0,
                        -136.0,
                        -252.0,
                        166.0,
                        208.0,
                        -126.0,
                        64.0,
                        323.0,
                        -104.0,
                        -106.0,
                        -248.0,
                        -41.0,
                        -109.0,
                        245.0,
                        47.0,
                        205.0,
                        358.0,
                        -296.0,
                        214.0,
                        -157.0,
                        -313.0,
                        -303.0,
                        -54.0,
                        -229.0,
                        231.0,
                        -94.0,
                        -198.0,
                        338.0,
                        199.0,
                        5.0,
                        42.0,
                        309.0,
                        73.0,
                        56.0,
                        -120.0,
                        351.0,
                        6.0,
                        -193.0,
                        21.0,
                        78.0,
                        57.0,
                        -269.0,
                        -76.0,
                        -299.0,
                        295.0,
                        -278.0,
                        11.0,
                        121.0,
                        -323.0,
                        156.0,
                        67.0,
                        152.0,
                        284.0,
                        108.0,
                        -7.0,
                        329.0,
                        -32.0,
                        333.0,
                        -338.0,
                        148.0,
                        -42.0,
                        151.0,
                        145.0,
                        -34.0,
                        -36.0,
                        296.0,
                        -198.0,
                        -317.0,
                        -161.0,
                        -253.0,
                        328.0,
                        -57.0,
                        134.0,
                        -289.0,
                        229.0,
                        44.0,
                        16.0,
                        -256.0,
                        289.0,
                        -234.0,
                        197.0,
                        333.0,
                        228.0,
                        -143.0,
                        -202.0,
                        -172.0,
                        -262.0,
                        -203.0,
                        -83.0,
                        -242.0,
                        -173.0,
                        336.0,
                        298.0,
                        -319.0,
                        66.0,
                        254.0,
                        214.0,
                        -118.0,
                        -216.0,
                        -168.0,
                        44.0,
                        -243.0,
                        207.0,
                        -28.0,
                        -4.0,
                        -272.0,
                        79.0,
                    },
                    Timezone: unifiedgosdk.Pointer("Atlantic/Reykjavik"),
                    WeekStart: shared.WeekStartTu.ToPointer(),
                },
            },
            RecurringEventID: unifiedgosdk.Pointer("cf746b0c-9906-4960-97c6-ccd79704eb13"),
            SendNotifications: unifiedgosdk.Pointer(false),
            StartAt: unifiedgosdk.Pointer("2020-05-20T08:05:45.142Z"),
            Status: shared.CalendarEventStatusConfirmed.ToPointer(),
            Subject: unifiedgosdk.Pointer("Sunt spargo tepidus bestia vigor credo coadunatio appello."),
            Timezone: unifiedgosdk.Pointer("Asia/Bangkok"),
            UpdatedAt: unifiedgosdk.Pointer("2020-06-26T04:47:47.420Z"),
            WebURL: unifiedgosdk.Pointer("https://another-pinstripe.com"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CalendarEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateCalendarEventRequest](../../pkg/models/operations/updatecalendareventrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateCalendarEventResponse](../../pkg/models/operations/updatecalendareventresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateCalendarLink

Update a link

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCalendarLink" method="put" path="/calendar/{connection_id}/link/{id}" example="calendar_link" -->
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

    res, err := s.Calendar.UpdateCalendarLink(ctx, operations.UpdateCalendarLinkRequest{
        CalendarLink: shared.CalendarLink{
            CreatedAt: unifiedgosdk.Pointer("2023-03-07T13:34:11.959Z"),
            Description: unifiedgosdk.Pointer("Vitium clibanus laboriosam uxor denuncio."),
            Duration: unifiedgosdk.Pointer[float64](74.0),
            ID: unifiedgosdk.Pointer("7ee7d961-69a9-4d0d-abc8-a952c590c9e6"),
            IsActive: unifiedgosdk.Pointer(true),
            Name: unifiedgosdk.Pointer("Sopor sopor ancilla animus anser dignissimos vito confero utilis."),
            PriceAmount: unifiedgosdk.Pointer[float64](44.0),
            PriceCurrency: unifiedgosdk.Pointer("USD"),
            UpdatedAt: unifiedgosdk.Pointer("2024-03-06T11:31:30.146Z"),
            URL: "https://annual-apricot.info/",
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CalendarLink != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateCalendarLinkRequest](../../pkg/models/operations/updatecalendarlinkrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.UpdateCalendarLinkResponse](../../pkg/models/operations/updatecalendarlinkresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateCalendarWebinar

Update a webinar

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCalendarWebinar" method="put" path="/calendar/{connection_id}/webinar/{id}" example="calendar_webinar" -->
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

    res, err := s.Calendar.UpdateCalendarWebinar(ctx, operations.UpdateCalendarWebinarRequest{
        CalendarWebinar: shared.CalendarWebinar{
            Conference: []shared.CalendarConference{},
            CreatedAt: types.MustNewTimeFromString("2022-07-06T11:45:14.631Z"),
            EndAt: types.MustNewTimeFromString("2025-10-03T23:05:05.720Z"),
            HasPolls: unifiedgosdk.Pointer(false),
            HasRecording: unifiedgosdk.Pointer(false),
            ID: unifiedgosdk.Pointer("d6c98c6e-b43a-42f2-bc60-51849586d7bb"),
            IsAutoApprove: unifiedgosdk.Pointer(false),
            IsEnabled: unifiedgosdk.Pointer(true),
            IsWebcast: unifiedgosdk.Pointer(false),
            JoinURL: unifiedgosdk.Pointer("https://robust-bathhouse.biz"),
            Notes: unifiedgosdk.Pointer("Curriculum ducimus assentator aspernatur ait."),
            Organizer: &shared.PropertyCalendarWebinarOrganizer{
                Email: unifiedgosdk.Pointer("Kelton_Dicki@yahoo.com"),
                Name: unifiedgosdk.Pointer("Walter Greenfelder"),
            },
            Recurrence: []shared.CalendarEventRecurrence{
                shared.CalendarEventRecurrence{
                    Count: unifiedgosdk.Pointer[float64](10.0),
                    EndAt: types.MustNewTimeFromString("2023-08-23T00:00:57.829Z"),
                    ExcludedDates: []string{
                        "2025-01-24T12:51:50.276Z",
                    },
                    Frequency: shared.CalendarEventRecurrenceFrequencyMonthly.ToPointer(),
                    IncludedDates: []string{
                        "2024-04-14T17:49:13.802Z",
                    },
                    Interval: unifiedgosdk.Pointer[float64](8.0),
                    OnDays: []shared.PropertyCalendarEventRecurrenceOnDays{
                        shared.PropertyCalendarEventRecurrenceOnDaysSu,
                        shared.PropertyCalendarEventRecurrenceOnDaysFr,
                        shared.PropertyCalendarEventRecurrenceOnDaysSa,
                        shared.PropertyCalendarEventRecurrenceOnDaysWe,
                        shared.PropertyCalendarEventRecurrenceOnDaysMo,
                    },
                    OnMonthDays: []float64{
                        -10.0,
                    },
                    OnMonths: []float64{
                        -9.0,
                    },
                    OnWeeks: []float64{
                        10.0,
                        30.0,
                        -38.0,
                        30.0,
                        -22.0,
                        37.0,
                        -12.0,
                        27.0,
                        2.0,
                        15.0,
                        26.0,
                        18.0,
                        -43.0,
                        -33.0,
                        -27.0,
                        38.0,
                        28.0,
                        47.0,
                        -8.0,
                        24.0,
                        35.0,
                        -2.0,
                        7.0,
                        49.0,
                        38.0,
                        -41.0,
                        46.0,
                        -11.0,
                        -45.0,
                        0.0,
                        48.0,
                        34.0,
                    },
                    OnYearDays: []float64{
                        345.0,
                        -207.0,
                        230.0,
                        -10.0,
                        364.0,
                        -256.0,
                        -218.0,
                        -295.0,
                        290.0,
                        -250.0,
                        -315.0,
                        60.0,
                        205.0,
                        -247.0,
                        -318.0,
                        -211.0,
                        -13.0,
                        256.0,
                        -200.0,
                        -313.0,
                        336.0,
                        -332.0,
                        -90.0,
                        287.0,
                        -273.0,
                        156.0,
                        241.0,
                        -138.0,
                        -363.0,
                        -37.0,
                        -171.0,
                        -62.0,
                        -57.0,
                        280.0,
                        -322.0,
                        -79.0,
                        -364.0,
                        -201.0,
                        84.0,
                        341.0,
                        334.0,
                        -75.0,
                        332.0,
                        207.0,
                        337.0,
                        -244.0,
                        131.0,
                        -191.0,
                        164.0,
                        -235.0,
                        285.0,
                        -309.0,
                        -158.0,
                        306.0,
                        180.0,
                        -130.0,
                        -162.0,
                        -155.0,
                        3.0,
                        198.0,
                        26.0,
                        -366.0,
                        -191.0,
                        127.0,
                        -331.0,
                        -11.0,
                        -239.0,
                        -189.0,
                        243.0,
                        118.0,
                        346.0,
                        -174.0,
                        -146.0,
                        -161.0,
                        -330.0,
                        327.0,
                        192.0,
                        310.0,
                        316.0,
                        313.0,
                        -242.0,
                        -51.0,
                        -264.0,
                        -180.0,
                        -88.0,
                        305.0,
                        270.0,
                        358.0,
                        -173.0,
                        -298.0,
                        153.0,
                        -89.0,
                        155.0,
                        -45.0,
                        248.0,
                        -46.0,
                        -146.0,
                        300.0,
                        364.0,
                        -335.0,
                        356.0,
                        -18.0,
                        219.0,
                        324.0,
                        -239.0,
                        -106.0,
                        -298.0,
                        328.0,
                        362.0,
                        344.0,
                        -54.0,
                        133.0,
                        50.0,
                        112.0,
                        -212.0,
                        -179.0,
                        22.0,
                        -201.0,
                        -62.0,
                        -293.0,
                        9.0,
                        30.0,
                        -50.0,
                        126.0,
                        -72.0,
                        264.0,
                        28.0,
                        -1.0,
                        -207.0,
                        160.0,
                        -168.0,
                        3.0,
                        -176.0,
                        -19.0,
                        -157.0,
                        349.0,
                        100.0,
                        -201.0,
                        108.0,
                        -180.0,
                        51.0,
                        -73.0,
                        366.0,
                        74.0,
                        -226.0,
                        238.0,
                        121.0,
                        -193.0,
                        -125.0,
                        -109.0,
                        316.0,
                        -177.0,
                        -307.0,
                        31.0,
                        -76.0,
                        217.0,
                        -310.0,
                        227.0,
                        -360.0,
                        71.0,
                        255.0,
                        -325.0,
                        -214.0,
                        40.0,
                        42.0,
                        17.0,
                        -241.0,
                        -84.0,
                        -188.0,
                        302.0,
                        64.0,
                        94.0,
                        -362.0,
                        23.0,
                        166.0,
                        85.0,
                        71.0,
                        -74.0,
                        -47.0,
                        -119.0,
                        98.0,
                        40.0,
                        158.0,
                        -64.0,
                        175.0,
                        269.0,
                        127.0,
                        -143.0,
                        213.0,
                        -196.0,
                        121.0,
                        81.0,
                        -238.0,
                        288.0,
                        321.0,
                        276.0,
                        133.0,
                        22.0,
                        -213.0,
                        -157.0,
                        -280.0,
                        -35.0,
                        73.0,
                        -194.0,
                        65.0,
                        -180.0,
                        63.0,
                        -242.0,
                        -117.0,
                        148.0,
                        157.0,
                        -320.0,
                        318.0,
                        8.0,
                        210.0,
                        -21.0,
                        81.0,
                        205.0,
                        -258.0,
                        -40.0,
                        -114.0,
                        -253.0,
                        -263.0,
                        65.0,
                        185.0,
                        -24.0,
                        324.0,
                        -172.0,
                        25.0,
                        260.0,
                        211.0,
                        342.0,
                        -31.0,
                        -288.0,
                        -159.0,
                        -4.0,
                        -2.0,
                        -107.0,
                        -316.0,
                        -276.0,
                        331.0,
                        -114.0,
                        -20.0,
                        -320.0,
                        51.0,
                        -176.0,
                        -148.0,
                        -50.0,
                        -201.0,
                        -104.0,
                        153.0,
                        -273.0,
                        -189.0,
                        67.0,
                        209.0,
                        149.0,
                        49.0,
                        -136.0,
                        -125.0,
                        -169.0,
                        -324.0,
                        309.0,
                        -51.0,
                        288.0,
                        253.0,
                        175.0,
                        -146.0,
                        171.0,
                        -140.0,
                        58.0,
                        -212.0,
                        164.0,
                        270.0,
                        102.0,
                        70.0,
                        299.0,
                        89.0,
                        -280.0,
                        252.0,
                        -342.0,
                        240.0,
                        226.0,
                        68.0,
                        -30.0,
                        -232.0,
                        -358.0,
                        -166.0,
                        60.0,
                        140.0,
                        275.0,
                        13.0,
                        250.0,
                        -328.0,
                        -189.0,
                        -22.0,
                        7.0,
                        -235.0,
                        -322.0,
                        178.0,
                        167.0,
                        -104.0,
                        -61.0,
                        282.0,
                        -80.0,
                        -277.0,
                        108.0,
                        271.0,
                        -237.0,
                        297.0,
                        -135.0,
                        -135.0,
                        -323.0,
                        342.0,
                        -267.0,
                        -235.0,
                        173.0,
                        249.0,
                        -288.0,
                        257.0,
                        139.0,
                        -191.0,
                        -217.0,
                        10.0,
                        -117.0,
                        -297.0,
                        -196.0,
                        -206.0,
                        341.0,
                        166.0,
                        181.0,
                        129.0,
                        -207.0,
                        55.0,
                        86.0,
                    },
                    Timezone: unifiedgosdk.Pointer("Asia/Ust-Nera"),
                    WeekStart: shared.WeekStartMo.ToPointer(),
                },
                shared.CalendarEventRecurrence{
                    Count: unifiedgosdk.Pointer[float64](3.0),
                    EndAt: types.MustNewTimeFromString("2022-09-28T21:54:22.888Z"),
                    ExcludedDates: []string{
                        "2024-08-16T15:01:59.509Z",
                        "2024-08-01T09:41:48.749Z",
                    },
                    Frequency: shared.CalendarEventRecurrenceFrequencyDaily.ToPointer(),
                    IncludedDates: []string{
                        "2024-03-12T07:59:01.474Z",
                        "2025-12-18T01:45:08.097Z",
                        "2023-08-06T00:06:02.458Z",
                    },
                    Interval: unifiedgosdk.Pointer[float64](1.0),
                    OnDays: []shared.PropertyCalendarEventRecurrenceOnDays{
                        shared.PropertyCalendarEventRecurrenceOnDaysWe,
                        shared.PropertyCalendarEventRecurrenceOnDaysSu,
                        shared.PropertyCalendarEventRecurrenceOnDaysMo,
                        shared.PropertyCalendarEventRecurrenceOnDaysFr,
                    },
                    OnMonthDays: []float64{
                        -15.0,
                    },
                    OnMonths: []float64{
                        5.0,
                        12.0,
                        3.0,
                        12.0,
                        8.0,
                    },
                    OnWeeks: []float64{
                        -47.0,
                        44.0,
                    },
                    OnYearDays: []float64{
                        -117.0,
                        59.0,
                        -6.0,
                        187.0,
                        45.0,
                        70.0,
                        15.0,
                        255.0,
                        44.0,
                        -2.0,
                        25.0,
                        -175.0,
                        -240.0,
                        171.0,
                        -294.0,
                        19.0,
                        38.0,
                        -351.0,
                        170.0,
                        -10.0,
                        -269.0,
                        18.0,
                        -65.0,
                        -266.0,
                        -31.0,
                        328.0,
                        -361.0,
                        358.0,
                        -256.0,
                        -4.0,
                        -312.0,
                        82.0,
                        -2.0,
                        -75.0,
                        -281.0,
                        -304.0,
                        53.0,
                        -295.0,
                        366.0,
                        322.0,
                        -191.0,
                        26.0,
                        97.0,
                        53.0,
                        75.0,
                        -62.0,
                        -109.0,
                        66.0,
                        177.0,
                        -68.0,
                        175.0,
                        -280.0,
                        70.0,
                        -238.0,
                        109.0,
                        -304.0,
                        326.0,
                        -8.0,
                        -71.0,
                        -236.0,
                        225.0,
                        358.0,
                        20.0,
                        -5.0,
                        -102.0,
                        -134.0,
                        -204.0,
                        -116.0,
                        -353.0,
                        -273.0,
                        106.0,
                        284.0,
                        -137.0,
                        -324.0,
                        301.0,
                        -42.0,
                        -229.0,
                        271.0,
                        -293.0,
                        -343.0,
                        211.0,
                        47.0,
                        -254.0,
                        -154.0,
                        -182.0,
                        264.0,
                        120.0,
                        -11.0,
                        -307.0,
                        99.0,
                        227.0,
                        190.0,
                        -17.0,
                        -77.0,
                        -255.0,
                        -61.0,
                        -249.0,
                        -102.0,
                        70.0,
                        345.0,
                        -187.0,
                        -308.0,
                        194.0,
                        221.0,
                        268.0,
                        -169.0,
                        -190.0,
                        88.0,
                        10.0,
                        262.0,
                        177.0,
                        -314.0,
                        -151.0,
                        -295.0,
                    },
                    Timezone: unifiedgosdk.Pointer("Pacific/Wake"),
                    WeekStart: shared.WeekStartTu.ToPointer(),
                },
                shared.CalendarEventRecurrence{
                    Count: unifiedgosdk.Pointer[float64](8.0),
                    EndAt: types.MustNewTimeFromString("2026-06-26T05:33:58.781Z"),
                    ExcludedDates: []string{
                        "2023-06-11T12:02:36.566Z",
                        "2023-05-31T18:16:08.923Z",
                    },
                    Frequency: shared.CalendarEventRecurrenceFrequencyWeekly.ToPointer(),
                    IncludedDates: []string{
                        "2024-03-20T04:54:34.101Z",
                        "2023-08-11T16:40:30.431Z",
                        "2024-09-10T07:26:29.395Z",
                    },
                    Interval: unifiedgosdk.Pointer[float64](8.0),
                    OnDays: []shared.PropertyCalendarEventRecurrenceOnDays{
                        shared.PropertyCalendarEventRecurrenceOnDaysSu,
                        shared.PropertyCalendarEventRecurrenceOnDaysMo,
                        shared.PropertyCalendarEventRecurrenceOnDaysTu,
                        shared.PropertyCalendarEventRecurrenceOnDaysFr,
                        shared.PropertyCalendarEventRecurrenceOnDaysMo,
                        shared.PropertyCalendarEventRecurrenceOnDaysTh,
                    },
                    OnMonthDays: []float64{
                        -23.0,
                    },
                    OnMonths: []float64{
                        11.0,
                        8.0,
                        9.0,
                        5.0,
                        -12.0,
                        -7.0,
                        -5.0,
                        10.0,
                        10.0,
                        -9.0,
                        -10.0,
                    },
                    OnWeeks: []float64{
                        -49.0,
                        46.0,
                        35.0,
                        -26.0,
                        2.0,
                        15.0,
                        15.0,
                        -26.0,
                        24.0,
                        -53.0,
                        36.0,
                        -43.0,
                        51.0,
                        -19.0,
                        -7.0,
                        -12.0,
                        28.0,
                        27.0,
                        35.0,
                        12.0,
                        -28.0,
                        -8.0,
                        -4.0,
                        -45.0,
                    },
                    OnYearDays: []float64{
                        84.0,
                        -251.0,
                        71.0,
                        181.0,
                        -163.0,
                        158.0,
                        301.0,
                        -299.0,
                        -184.0,
                        -331.0,
                        -152.0,
                        -129.0,
                        -237.0,
                        -303.0,
                        -24.0,
                        126.0,
                        -103.0,
                        146.0,
                        -346.0,
                        86.0,
                        -296.0,
                        -337.0,
                        -185.0,
                        16.0,
                        -270.0,
                        -126.0,
                        -295.0,
                        -231.0,
                        356.0,
                        -293.0,
                        115.0,
                        -265.0,
                        -293.0,
                        -34.0,
                        357.0,
                        313.0,
                        -343.0,
                        180.0,
                        -22.0,
                        -161.0,
                        350.0,
                        177.0,
                        190.0,
                        223.0,
                        -152.0,
                        -360.0,
                        -225.0,
                        -60.0,
                        -35.0,
                        353.0,
                        117.0,
                        -171.0,
                        -315.0,
                        -321.0,
                        -202.0,
                        345.0,
                        -1.0,
                        -148.0,
                        -168.0,
                        181.0,
                        -17.0,
                        282.0,
                        234.0,
                        31.0,
                        47.0,
                        -236.0,
                        366.0,
                        -251.0,
                        -232.0,
                        -308.0,
                        76.0,
                        -199.0,
                        184.0,
                        198.0,
                        225.0,
                        75.0,
                        6.0,
                        227.0,
                        -148.0,
                        259.0,
                        -146.0,
                        49.0,
                        -254.0,
                        341.0,
                        93.0,
                        138.0,
                        -164.0,
                        237.0,
                        4.0,
                        -287.0,
                        161.0,
                    },
                    Timezone: unifiedgosdk.Pointer("Africa/Bissau"),
                    WeekStart: shared.WeekStartWe.ToPointer(),
                },
            },
            RegistrantPassword: unifiedgosdk.Pointer("OxwWzr0C"),
            RequireAddress: unifiedgosdk.Pointer(false),
            RequireJobTitle: unifiedgosdk.Pointer(false),
            StartAt: types.MustNewTimeFromString("2025-04-09T12:29:18.755Z"),
            Status: shared.CalendarWebinarStatusTentative.ToPointer(),
            Subject: unifiedgosdk.Pointer("Harum culpa decipio ex cubo ancilla cresco."),
            Timezone: unifiedgosdk.Pointer("Europe/Kaliningrad"),
            UpdatedAt: types.MustNewTimeFromString("2026-08-29T20:26:31.460Z"),
            WebURL: unifiedgosdk.Pointer("https://classic-recovery.biz"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CalendarWebinar != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdateCalendarWebinarRequest](../../pkg/models/operations/updatecalendarwebinarrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UpdateCalendarWebinarResponse](../../pkg/models/operations/updatecalendarwebinarresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |