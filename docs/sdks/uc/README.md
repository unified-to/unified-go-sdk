# Uc

## Overview

### Available Operations

* [CreateUcCall](#createuccall) - Create a call
* [CreateUcComment](#createuccomment) - Create a comment
* [CreateUcContact](#createuccontact) - Create a contact
* [CreateUcRecording](#createucrecording) - Create a recording
* [GetUcCall](#getuccall) - Retrieve a call
* [GetUcComment](#getuccomment) - Retrieve a comment
* [GetUcContact](#getuccontact) - Retrieve a contact
* [GetUcRecording](#getucrecording) - Retrieve a recording
* [GetUcScorecard](#getucscorecard) - Retrieve a scorecard
* [GetUcScorecardtemplate](#getucscorecardtemplate) - Retrieve a scorecardtemplate
* [ListUcCalls](#listuccalls) - List all calls
* [ListUcComments](#listuccomments) - List all comments
* [ListUcContacts](#listuccontacts) - List all contacts
* [ListUcRecordings](#listucrecordings) - List all recordings
* [ListUcScorecards](#listucscorecards) - List all scorecards
* [ListUcScorecardtemplates](#listucscorecardtemplates) - List all scorecardtemplates
* [PatchUcCall](#patchuccall) - Update a call
* [PatchUcComment](#patchuccomment) - Update a comment
* [PatchUcContact](#patchuccontact) - Update a contact
* [PatchUcRecording](#patchucrecording) - Update a recording
* [RemoveUcComment](#removeuccomment) - Remove a comment
* [RemoveUcContact](#removeuccontact) - Remove a contact
* [RemoveUcRecording](#removeucrecording) - Remove a recording
* [UpdateUcCall](#updateuccall) - Update a call
* [UpdateUcComment](#updateuccomment) - Update a comment
* [UpdateUcContact](#updateuccontact) - Update a contact
* [UpdateUcRecording](#updateucrecording) - Update a recording

## CreateUcCall

Creates a UcCall. Provide the unified UcCall fields in the request body; they are mapped to the provider's API. Returns the created UcCall.

### Example Usage

<!-- UsageSnippet language="go" operationID="createUcCall" method="post" path="/uc/{connection_id}/call" example="uc_call" -->
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

    res, err := s.Uc.CreateUcCall(ctx, operations.CreateUcCallRequest{
        UcCall: shared.UcCall{
            AnswerAt: types.MustNewTimeFromString("2025-05-18T19:57:43.114Z"),
            Contacts: []shared.UcContact{},
            ConversationIdentifier: unifiedgosdk.Pointer("8354a922-f8b1-4ec9-84c9-b9b3936cca8a"),
            CostAmount: unifiedgosdk.Pointer[float64](1.564),
            CreatedAt: types.MustNewTimeFromString("2023-04-22T19:05:34.882Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DurationSeconds: unifiedgosdk.Pointer[float64](2556.0),
            EndAt: types.MustNewTimeFromString("2024-06-01T20:25:25.237Z"),
            ID: unifiedgosdk.Pointer("8354a922-f8b1-4ec9-84c9-b9b3936cca8a"),
            IsPrivate: unifiedgosdk.Pointer(false),
            Metadata: []shared.UcMetadata{
                shared.UcMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateUcMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.UcMetadataFormatPrice.ToPointer(),
                    ID: unifiedgosdk.Pointer("96afd034-f8e3-48c5-a1e5-229ab7d3a20b"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_field"),
                    Value: unifiedgosdk.Pointer(shared.CreateUcMetadataValueStr(
                        "tui",
                    )),
                },
            },
            StartAt: types.MustNewTimeFromString("2023-12-19T01:34:36.131Z"),
            Status: shared.UcCallStatusInProgress.ToPointer(),
            Telephone: &shared.PropertyUcCallTelephone{
                Telephone: "(493) 534-0768",
                Type: shared.PropertyUcCallTelephoneTypeMobile.ToPointer(),
            },
            Type: shared.UcCallTypeInbound.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-08-04T01:34:36.809Z"),
            UserPhone: unifiedgosdk.Pointer("819-894-3317 x163"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcCall != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateUcCallRequest](../../pkg/models/operations/createuccallrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.CreateUcCallResponse](../../pkg/models/operations/createuccallresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateUcComment

Creates a UcComment. Provide the unified UcComment fields in the request body; they are mapped to the provider's API. Returns the created UcComment.

### Example Usage

<!-- UsageSnippet language="go" operationID="createUcComment" method="post" path="/uc/{connection_id}/comment" example="uc_comment" -->
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

    res, err := s.Uc.CreateUcComment(ctx, operations.CreateUcCommentRequest{
        UcComment: shared.UcComment{
            Content: unifiedgosdk.Pointer("Vociferor vitiosus."),
            CreatedAt: unifiedgosdk.Pointer("2023-04-02T23:42:31.571Z"),
            ID: unifiedgosdk.Pointer("b91b94ee-70e3-48cb-af03-0f777aaee0c0"),
            Metadata: []shared.UcMetadata{},
            UpdatedAt: unifiedgosdk.Pointer("2024-02-11T21:33:10.072Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcComment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateUcCommentRequest](../../pkg/models/operations/createuccommentrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreateUcCommentResponse](../../pkg/models/operations/createuccommentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateUcContact

Creates a UcContact. Provide the unified UcContact fields in the request body; they are mapped to the provider's API. Returns the created UcContact.

### Example Usage

<!-- UsageSnippet language="go" operationID="createUcContact" method="post" path="/uc/{connection_id}/contact" example="uc_contact" -->
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

    res, err := s.Uc.CreateUcContact(ctx, operations.CreateUcContactRequest{
        UcContact: shared.UcContact{
            Company: unifiedgosdk.Pointer("Tillman Group"),
            CreatedAt: types.MustNewTimeFromString("2019-10-28T11:06:56.460Z"),
            Emails: []shared.UcEmail{
                shared.UcEmail{
                    Email: "Luther_Rogahn32@yahoo.com",
                    Type: shared.UcEmailTypeWork.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Luther"),
            ID: unifiedgosdk.Pointer("01d10f86-19f2-4a36-a113-d4cf95b86a9c"),
            LastName: unifiedgosdk.Pointer("Rogahn"),
            Metadata: []shared.UcMetadata{},
            Name: unifiedgosdk.Pointer("Luther Rogahn"),
            Telephones: []shared.UcTelephone{
                shared.UcTelephone{
                    Telephone: "(809) 992-1681",
                    Type: shared.UcTelephoneTypeFax.ToPointer(),
                },
                shared.UcTelephone{
                    Telephone: "(868) 238-2746",
                    Type: shared.UcTelephoneTypeHome.ToPointer(),
                },
                shared.UcTelephone{
                    Telephone: "(219) 736-0357",
                    Type: shared.UcTelephoneTypeMobile.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("Chief Optimization Executive"),
            UpdatedAt: types.MustNewTimeFromString("2023-12-13T13:49:07.852Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateUcContactRequest](../../pkg/models/operations/createuccontactrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreateUcContactResponse](../../pkg/models/operations/createuccontactresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateUcRecording

Creates a UcRecording. Provide the unified UcRecording fields in the request body; they are mapped to the provider's API. Returns the created UcRecording.

### Example Usage

<!-- UsageSnippet language="go" operationID="createUcRecording" method="post" path="/uc/{connection_id}/recording" example="uc_recording" -->
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

    res, err := s.Uc.CreateUcRecording(ctx, operations.CreateUcRecordingRequest{
        UcRecording: shared.UcRecording{
            Contacts: []shared.UcContact{},
            ConversationIdentifier: unifiedgosdk.Pointer("85bf7066-a183-4fed-80a4-b6761b3f287d"),
            CreatedAt: types.MustNewTimeFromString("2022-09-17T19:41:46.956Z"),
            EndAt: types.MustNewTimeFromString("2024-05-08T12:15:15.297Z"),
            ExpiresAt: types.MustNewTimeFromString("2026-05-04T13:03:57.607Z"),
            ID: unifiedgosdk.Pointer("85bf7066-a183-4fed-80a4-b6761b3f287d"),
            Media: []shared.UcRecordingMedia{
                shared.UcRecordingMedia{
                    EndAt: types.MustNewTimeFromString("2024-04-05T03:13:17.505Z"),
                    Language: unifiedgosdk.Pointer("en"),
                    RecordingDownloadURL: unifiedgosdk.Pointer("https://parched-wasabi.com/"),
                    StartAt: types.MustNewTimeFromString("2026-05-23T07:46:34.139Z"),
                    TranscriptDownloadURL: unifiedgosdk.Pointer("https://colossal-cuckoo.name"),
                    Transcripts: []shared.UcRecordingTranscript{
                        shared.UcRecordingTranscript{
                            EndAt: types.MustNewTimeFromString("2026-03-03T22:12:33.733Z"),
                            StartAt: types.MustNewTimeFromString("2024-06-05T22:24:00.419Z"),
                            Text: "Turpis maiores ducimus tero speculum.",
                        },
                        shared.UcRecordingTranscript{
                            EndAt: types.MustNewTimeFromString("2024-10-17T00:16:52.078Z"),
                            StartAt: types.MustNewTimeFromString("2026-08-11T05:36:37.844Z"),
                            Text: "Magnam consuasor uxor tergiversatio subseco.",
                        },
                    },
                },
            },
            Metadata: []shared.UcMetadata{},
            StartAt: types.MustNewTimeFromString("2023-04-29T01:30:54.260Z"),
            Type: shared.UcRecordingTypeOutbound.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-03-21T20:09:31.243Z"),
            UserPhone: unifiedgosdk.Pointer("1-915-327-0429 x509"),
            WebURL: unifiedgosdk.Pointer("https://spherical-comparison.org"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcRecording != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateUcRecordingRequest](../../pkg/models/operations/createucrecordingrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateUcRecordingResponse](../../pkg/models/operations/createucrecordingresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetUcCall

Retrieves a single UcCall by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getUcCall" method="get" path="/uc/{connection_id}/call/{id}" -->
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

    res, err := s.Uc.GetUcCall(ctx, operations.GetUcCallRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcCall != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetUcCallRequest](../../pkg/models/operations/getuccallrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.GetUcCallResponse](../../pkg/models/operations/getuccallresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetUcComment

Retrieves a single UcComment by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getUcComment" method="get" path="/uc/{connection_id}/comment/{id}" -->
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

    res, err := s.Uc.GetUcComment(ctx, operations.GetUcCommentRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcComment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetUcCommentRequest](../../pkg/models/operations/getuccommentrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.GetUcCommentResponse](../../pkg/models/operations/getuccommentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetUcContact

Retrieves a single UcContact by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getUcContact" method="get" path="/uc/{connection_id}/contact/{id}" -->
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

    res, err := s.Uc.GetUcContact(ctx, operations.GetUcContactRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetUcContactRequest](../../pkg/models/operations/getuccontactrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.GetUcContactResponse](../../pkg/models/operations/getuccontactresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetUcRecording

Retrieves a single UcRecording by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getUcRecording" method="get" path="/uc/{connection_id}/recording/{id}" -->
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

    res, err := s.Uc.GetUcRecording(ctx, operations.GetUcRecordingRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcRecording != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetUcRecordingRequest](../../pkg/models/operations/getucrecordingrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetUcRecordingResponse](../../pkg/models/operations/getucrecordingresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetUcScorecard

Retrieves a single UcScorecard by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getUcScorecard" method="get" path="/uc/{connection_id}/scorecard/{id}" -->
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

    res, err := s.Uc.GetUcScorecard(ctx, operations.GetUcScorecardRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcScorecard != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetUcScorecardRequest](../../pkg/models/operations/getucscorecardrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetUcScorecardResponse](../../pkg/models/operations/getucscorecardresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetUcScorecardtemplate

Retrieves a single UcScorecardtemplate by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getUcScorecardtemplate" method="get" path="/uc/{connection_id}/scorecardtemplate/{id}" -->
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

    res, err := s.Uc.GetUcScorecardtemplate(ctx, operations.GetUcScorecardtemplateRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcScorecardtemplate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetUcScorecardtemplateRequest](../../pkg/models/operations/getucscorecardtemplaterequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.GetUcScorecardtemplateResponse](../../pkg/models/operations/getucscorecardtemplateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListUcCalls

Returns a list of UcCalls. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listUcCalls" method="get" path="/uc/{connection_id}/call" -->
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

    res, err := s.Uc.ListUcCalls(ctx, operations.ListUcCallsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcCalls != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListUcCallsRequest](../../pkg/models/operations/listuccallsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.ListUcCallsResponse](../../pkg/models/operations/listuccallsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListUcComments

Returns a list of UcComments. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listUcComments" method="get" path="/uc/{connection_id}/comment" -->
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

    res, err := s.Uc.ListUcComments(ctx, operations.ListUcCommentsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcComments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListUcCommentsRequest](../../pkg/models/operations/listuccommentsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ListUcCommentsResponse](../../pkg/models/operations/listuccommentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListUcContacts

Returns a list of UcContacts. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listUcContacts" method="get" path="/uc/{connection_id}/contact" -->
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

    res, err := s.Uc.ListUcContacts(ctx, operations.ListUcContactsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcContacts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListUcContactsRequest](../../pkg/models/operations/listuccontactsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ListUcContactsResponse](../../pkg/models/operations/listuccontactsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListUcRecordings

Returns a list of UcRecordings. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listUcRecordings" method="get" path="/uc/{connection_id}/recording" -->
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

    res, err := s.Uc.ListUcRecordings(ctx, operations.ListUcRecordingsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcRecordings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListUcRecordingsRequest](../../pkg/models/operations/listucrecordingsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListUcRecordingsResponse](../../pkg/models/operations/listucrecordingsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListUcScorecards

Returns a list of UcScorecards. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listUcScorecards" method="get" path="/uc/{connection_id}/scorecard" -->
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

    res, err := s.Uc.ListUcScorecards(ctx, operations.ListUcScorecardsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcScorecards != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListUcScorecardsRequest](../../pkg/models/operations/listucscorecardsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListUcScorecardsResponse](../../pkg/models/operations/listucscorecardsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListUcScorecardtemplates

Returns a list of UcScorecardtemplates. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listUcScorecardtemplates" method="get" path="/uc/{connection_id}/scorecardtemplate" -->
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

    res, err := s.Uc.ListUcScorecardtemplates(ctx, operations.ListUcScorecardtemplatesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcScorecardtemplates != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ListUcScorecardtemplatesRequest](../../pkg/models/operations/listucscorecardtemplatesrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.ListUcScorecardtemplatesResponse](../../pkg/models/operations/listucscorecardtemplatesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchUcCall

Updates the UcCall identified by `id`. Provide the unified UcCall fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchUcCall" method="patch" path="/uc/{connection_id}/call/{id}" example="uc_call" -->
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

    res, err := s.Uc.PatchUcCall(ctx, operations.PatchUcCallRequest{
        UcCall: shared.UcCall{
            AnswerAt: types.MustNewTimeFromString("2025-05-18T19:57:43.144Z"),
            Contacts: []shared.UcContact{},
            ConversationIdentifier: unifiedgosdk.Pointer("24c9f9b8-c64c-4a9a-8617-cd5276f393cc"),
            CostAmount: unifiedgosdk.Pointer[float64](1.564),
            CreatedAt: types.MustNewTimeFromString("2023-04-22T19:05:34.882Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DurationSeconds: unifiedgosdk.Pointer[float64](2556.0),
            EndAt: types.MustNewTimeFromString("2024-06-01T20:25:25.253Z"),
            ID: unifiedgosdk.Pointer("24c9f9b8-c64c-4a9a-8617-cd5276f393cc"),
            IsPrivate: unifiedgosdk.Pointer(false),
            Metadata: []shared.UcMetadata{
                shared.UcMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateUcMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.UcMetadataFormatPrice.ToPointer(),
                    ID: unifiedgosdk.Pointer("9b42f590-5959-4138-9a75-613d46756c7e"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_field"),
                    Value: unifiedgosdk.Pointer(shared.CreateUcMetadataValueStr(
                        "tui",
                    )),
                },
            },
            StartAt: types.MustNewTimeFromString("2023-12-19T01:34:36.140Z"),
            Status: shared.UcCallStatusInProgress.ToPointer(),
            Telephone: &shared.PropertyUcCallTelephone{
                Telephone: "(493) 534-0768",
                Type: shared.PropertyUcCallTelephoneTypeMobile.ToPointer(),
            },
            Type: shared.UcCallTypeInbound.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-08-04T01:34:36.814Z"),
            UserPhone: unifiedgosdk.Pointer("819-894-3317 x163"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcCall != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.PatchUcCallRequest](../../pkg/models/operations/patchuccallrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.PatchUcCallResponse](../../pkg/models/operations/patchuccallresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchUcComment

Updates the UcComment identified by `id`. Provide the unified UcComment fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchUcComment" method="patch" path="/uc/{connection_id}/comment/{id}" example="uc_comment" -->
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

    res, err := s.Uc.PatchUcComment(ctx, operations.PatchUcCommentRequest{
        UcComment: shared.UcComment{
            Content: unifiedgosdk.Pointer("Vociferor vitiosus."),
            CreatedAt: unifiedgosdk.Pointer("2023-04-02T23:42:31.571Z"),
            ID: unifiedgosdk.Pointer("9e630159-6f8d-4a34-aa47-97b48f12d1cf"),
            Metadata: []shared.UcMetadata{},
            UpdatedAt: unifiedgosdk.Pointer("2024-02-11T21:33:10.075Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcComment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.PatchUcCommentRequest](../../pkg/models/operations/patchuccommentrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.PatchUcCommentResponse](../../pkg/models/operations/patchuccommentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchUcContact

Updates the UcContact identified by `id`. Provide the unified UcContact fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchUcContact" method="patch" path="/uc/{connection_id}/contact/{id}" example="uc_contact" -->
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

    res, err := s.Uc.PatchUcContact(ctx, operations.PatchUcContactRequest{
        UcContact: shared.UcContact{
            Company: unifiedgosdk.Pointer("Tillman Group"),
            CreatedAt: types.MustNewTimeFromString("2019-10-28T11:06:56.460Z"),
            Emails: []shared.UcEmail{
                shared.UcEmail{
                    Email: "Luther_Rogahn32@yahoo.com",
                    Type: shared.UcEmailTypeWork.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Luther"),
            ID: unifiedgosdk.Pointer("622ded20-117c-4d9e-9b8b-08676b0860c0"),
            LastName: unifiedgosdk.Pointer("Rogahn"),
            Metadata: []shared.UcMetadata{},
            Name: unifiedgosdk.Pointer("Luther Rogahn"),
            Telephones: []shared.UcTelephone{
                shared.UcTelephone{
                    Telephone: "(809) 992-1681",
                    Type: shared.UcTelephoneTypeFax.ToPointer(),
                },
                shared.UcTelephone{
                    Telephone: "(868) 238-2746",
                    Type: shared.UcTelephoneTypeHome.ToPointer(),
                },
                shared.UcTelephone{
                    Telephone: "(219) 736-0357",
                    Type: shared.UcTelephoneTypeMobile.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("Chief Optimization Executive"),
            UpdatedAt: types.MustNewTimeFromString("2023-12-13T13:49:07.864Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.PatchUcContactRequest](../../pkg/models/operations/patchuccontactrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.PatchUcContactResponse](../../pkg/models/operations/patchuccontactresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchUcRecording

Updates the UcRecording identified by `id`. Provide the unified UcRecording fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchUcRecording" method="patch" path="/uc/{connection_id}/recording/{id}" example="uc_recording" -->
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

    res, err := s.Uc.PatchUcRecording(ctx, operations.PatchUcRecordingRequest{
        UcRecording: shared.UcRecording{
            Contacts: []shared.UcContact{},
            ConversationIdentifier: unifiedgosdk.Pointer("7a629e4b-e67d-4959-a527-3aab8cca39a5"),
            CreatedAt: types.MustNewTimeFromString("2022-09-17T19:41:46.956Z"),
            EndAt: types.MustNewTimeFromString("2024-05-08T12:15:15.315Z"),
            ExpiresAt: types.MustNewTimeFromString("2026-05-04T13:03:57.646Z"),
            ID: unifiedgosdk.Pointer("7a629e4b-e67d-4959-a527-3aab8cca39a5"),
            Media: []shared.UcRecordingMedia{
                shared.UcRecordingMedia{
                    EndAt: types.MustNewTimeFromString("2024-04-05T03:13:17.521Z"),
                    Language: unifiedgosdk.Pointer("en"),
                    RecordingDownloadURL: unifiedgosdk.Pointer("https://parched-wasabi.com/"),
                    StartAt: types.MustNewTimeFromString("2026-05-23T07:46:34.178Z"),
                    TranscriptDownloadURL: unifiedgosdk.Pointer("https://colossal-cuckoo.name"),
                    Transcripts: []shared.UcRecordingTranscript{
                        shared.UcRecordingTranscript{
                            EndAt: types.MustNewTimeFromString("2026-03-03T22:12:33.770Z"),
                            StartAt: types.MustNewTimeFromString("2024-06-05T22:24:00.438Z"),
                            Text: "Turpis maiores ducimus tero speculum.",
                        },
                        shared.UcRecordingTranscript{
                            EndAt: types.MustNewTimeFromString("2024-10-17T00:16:52.100Z"),
                            StartAt: types.MustNewTimeFromString("2026-08-11T05:36:37.886Z"),
                            Text: "Magnam consuasor uxor tergiversatio subseco.",
                        },
                    },
                },
            },
            Metadata: []shared.UcMetadata{},
            StartAt: types.MustNewTimeFromString("2023-04-29T01:30:54.267Z"),
            Type: shared.UcRecordingTypeOutbound.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-03-21T20:09:31.270Z"),
            UserPhone: unifiedgosdk.Pointer("1-915-327-0429 x509"),
            WebURL: unifiedgosdk.Pointer("https://spherical-comparison.org"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcRecording != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchUcRecordingRequest](../../pkg/models/operations/patchucrecordingrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchUcRecordingResponse](../../pkg/models/operations/patchucrecordingresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveUcComment

Deletes the UcComment identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeUcComment" method="delete" path="/uc/{connection_id}/comment/{id}" -->
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

    res, err := s.Uc.RemoveUcComment(ctx, operations.RemoveUcCommentRequest{
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.RemoveUcCommentRequest](../../pkg/models/operations/removeuccommentrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.RemoveUcCommentResponse](../../pkg/models/operations/removeuccommentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveUcContact

Deletes the UcContact identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeUcContact" method="delete" path="/uc/{connection_id}/contact/{id}" -->
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

    res, err := s.Uc.RemoveUcContact(ctx, operations.RemoveUcContactRequest{
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.RemoveUcContactRequest](../../pkg/models/operations/removeuccontactrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.RemoveUcContactResponse](../../pkg/models/operations/removeuccontactresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveUcRecording

Deletes the UcRecording identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeUcRecording" method="delete" path="/uc/{connection_id}/recording/{id}" -->
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

    res, err := s.Uc.RemoveUcRecording(ctx, operations.RemoveUcRecordingRequest{
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
| `request`                                                                                      | [operations.RemoveUcRecordingRequest](../../pkg/models/operations/removeucrecordingrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemoveUcRecordingResponse](../../pkg/models/operations/removeucrecordingresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateUcCall

Updates the UcCall identified by `id`. Provide the unified UcCall fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateUcCall" method="put" path="/uc/{connection_id}/call/{id}" example="uc_call" -->
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

    res, err := s.Uc.UpdateUcCall(ctx, operations.UpdateUcCallRequest{
        UcCall: shared.UcCall{
            AnswerAt: types.MustNewTimeFromString("2025-05-18T19:57:43.144Z"),
            Contacts: []shared.UcContact{},
            ConversationIdentifier: unifiedgosdk.Pointer("24c9f9b8-c64c-4a9a-8617-cd5276f393cc"),
            CostAmount: unifiedgosdk.Pointer[float64](1.564),
            CreatedAt: types.MustNewTimeFromString("2023-04-22T19:05:34.882Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DurationSeconds: unifiedgosdk.Pointer[float64](2556.0),
            EndAt: types.MustNewTimeFromString("2024-06-01T20:25:25.253Z"),
            ID: unifiedgosdk.Pointer("24c9f9b8-c64c-4a9a-8617-cd5276f393cc"),
            IsPrivate: unifiedgosdk.Pointer(false),
            Metadata: []shared.UcMetadata{
                shared.UcMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateUcMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.UcMetadataFormatPrice.ToPointer(),
                    ID: unifiedgosdk.Pointer("9b42f590-5959-4138-9a75-613d46756c7e"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_field"),
                    Value: unifiedgosdk.Pointer(shared.CreateUcMetadataValueStr(
                        "tui",
                    )),
                },
            },
            StartAt: types.MustNewTimeFromString("2023-12-19T01:34:36.140Z"),
            Status: shared.UcCallStatusInProgress.ToPointer(),
            Telephone: &shared.PropertyUcCallTelephone{
                Telephone: "(493) 534-0768",
                Type: shared.PropertyUcCallTelephoneTypeMobile.ToPointer(),
            },
            Type: shared.UcCallTypeInbound.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-08-04T01:34:36.814Z"),
            UserPhone: unifiedgosdk.Pointer("819-894-3317 x163"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcCall != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateUcCallRequest](../../pkg/models/operations/updateuccallrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.UpdateUcCallResponse](../../pkg/models/operations/updateuccallresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateUcComment

Updates the UcComment identified by `id`. Provide the unified UcComment fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateUcComment" method="put" path="/uc/{connection_id}/comment/{id}" example="uc_comment" -->
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

    res, err := s.Uc.UpdateUcComment(ctx, operations.UpdateUcCommentRequest{
        UcComment: shared.UcComment{
            Content: unifiedgosdk.Pointer("Vociferor vitiosus."),
            CreatedAt: unifiedgosdk.Pointer("2023-04-02T23:42:31.571Z"),
            ID: unifiedgosdk.Pointer("9e630159-6f8d-4a34-aa47-97b48f12d1cf"),
            Metadata: []shared.UcMetadata{},
            UpdatedAt: unifiedgosdk.Pointer("2024-02-11T21:33:10.075Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcComment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateUcCommentRequest](../../pkg/models/operations/updateuccommentrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.UpdateUcCommentResponse](../../pkg/models/operations/updateuccommentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateUcContact

Updates the UcContact identified by `id`. Provide the unified UcContact fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateUcContact" method="put" path="/uc/{connection_id}/contact/{id}" example="uc_contact" -->
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

    res, err := s.Uc.UpdateUcContact(ctx, operations.UpdateUcContactRequest{
        UcContact: shared.UcContact{
            Company: unifiedgosdk.Pointer("Tillman Group"),
            CreatedAt: types.MustNewTimeFromString("2019-10-28T11:06:56.460Z"),
            Emails: []shared.UcEmail{
                shared.UcEmail{
                    Email: "Luther_Rogahn32@yahoo.com",
                    Type: shared.UcEmailTypeWork.ToPointer(),
                },
            },
            FirstName: unifiedgosdk.Pointer("Luther"),
            ID: unifiedgosdk.Pointer("622ded20-117c-4d9e-9b8b-08676b0860c0"),
            LastName: unifiedgosdk.Pointer("Rogahn"),
            Metadata: []shared.UcMetadata{},
            Name: unifiedgosdk.Pointer("Luther Rogahn"),
            Telephones: []shared.UcTelephone{
                shared.UcTelephone{
                    Telephone: "(809) 992-1681",
                    Type: shared.UcTelephoneTypeFax.ToPointer(),
                },
                shared.UcTelephone{
                    Telephone: "(868) 238-2746",
                    Type: shared.UcTelephoneTypeHome.ToPointer(),
                },
                shared.UcTelephone{
                    Telephone: "(219) 736-0357",
                    Type: shared.UcTelephoneTypeMobile.ToPointer(),
                },
            },
            Title: unifiedgosdk.Pointer("Chief Optimization Executive"),
            UpdatedAt: types.MustNewTimeFromString("2023-12-13T13:49:07.864Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcContact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateUcContactRequest](../../pkg/models/operations/updateuccontactrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.UpdateUcContactResponse](../../pkg/models/operations/updateuccontactresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateUcRecording

Updates the UcRecording identified by `id`. Provide the unified UcRecording fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateUcRecording" method="put" path="/uc/{connection_id}/recording/{id}" example="uc_recording" -->
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

    res, err := s.Uc.UpdateUcRecording(ctx, operations.UpdateUcRecordingRequest{
        UcRecording: shared.UcRecording{
            Contacts: []shared.UcContact{},
            ConversationIdentifier: unifiedgosdk.Pointer("7a629e4b-e67d-4959-a527-3aab8cca39a5"),
            CreatedAt: types.MustNewTimeFromString("2022-09-17T19:41:46.956Z"),
            EndAt: types.MustNewTimeFromString("2024-05-08T12:15:15.315Z"),
            ExpiresAt: types.MustNewTimeFromString("2026-05-04T13:03:57.646Z"),
            ID: unifiedgosdk.Pointer("7a629e4b-e67d-4959-a527-3aab8cca39a5"),
            Media: []shared.UcRecordingMedia{
                shared.UcRecordingMedia{
                    EndAt: types.MustNewTimeFromString("2024-04-05T03:13:17.521Z"),
                    Language: unifiedgosdk.Pointer("en"),
                    RecordingDownloadURL: unifiedgosdk.Pointer("https://parched-wasabi.com/"),
                    StartAt: types.MustNewTimeFromString("2026-05-23T07:46:34.178Z"),
                    TranscriptDownloadURL: unifiedgosdk.Pointer("https://colossal-cuckoo.name"),
                    Transcripts: []shared.UcRecordingTranscript{
                        shared.UcRecordingTranscript{
                            EndAt: types.MustNewTimeFromString("2026-03-03T22:12:33.770Z"),
                            StartAt: types.MustNewTimeFromString("2024-06-05T22:24:00.438Z"),
                            Text: "Turpis maiores ducimus tero speculum.",
                        },
                        shared.UcRecordingTranscript{
                            EndAt: types.MustNewTimeFromString("2024-10-17T00:16:52.100Z"),
                            StartAt: types.MustNewTimeFromString("2026-08-11T05:36:37.886Z"),
                            Text: "Magnam consuasor uxor tergiversatio subseco.",
                        },
                    },
                },
            },
            Metadata: []shared.UcMetadata{},
            StartAt: types.MustNewTimeFromString("2023-04-29T01:30:54.267Z"),
            Type: shared.UcRecordingTypeOutbound.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-03-21T20:09:31.270Z"),
            UserPhone: unifiedgosdk.Pointer("1-915-327-0429 x509"),
            WebURL: unifiedgosdk.Pointer("https://spherical-comparison.org"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UcRecording != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateUcRecordingRequest](../../pkg/models/operations/updateucrecordingrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateUcRecordingResponse](../../pkg/models/operations/updateucrecordingresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |