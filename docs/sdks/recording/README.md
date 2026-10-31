# Recording

## Overview

### Available Operations

* [CreateUcRecording](#createucrecording) - Create a recording
* [GetCalendarRecording](#getcalendarrecording) - Retrieve a recording
* [GetUcRecording](#getucrecording) - Retrieve a recording
* [ListCalendarRecordings](#listcalendarrecordings) - List all recordings
* [ListUcRecordings](#listucrecordings) - List all recordings
* [PatchUcRecording](#patchucrecording) - Update a recording
* [RemoveUcRecording](#removeucrecording) - Remove a recording
* [UpdateUcRecording](#updateucrecording) - Update a recording

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

    res, err := s.Recording.CreateUcRecording(ctx, operations.CreateUcRecordingRequest{
        UcRecording: shared.UcRecording{
            Contacts: []shared.UcContact{},
            ConversationIdentifier: unifiedgosdk.Pointer("409675fd-b6f6-44fa-bf8b-6e20e3dc32de"),
            CreatedAt: types.MustNewTimeFromString("2022-09-17T19:41:46.956Z"),
            EndAt: types.MustNewTimeFromString("2024-05-08T23:47:01.710Z"),
            ExpiresAt: types.MustNewTimeFromString("2026-05-05T14:34:39.025Z"),
            ID: unifiedgosdk.Pointer("409675fd-b6f6-44fa-bf8b-6e20e3dc32de"),
            Media: []shared.UcRecordingMedia{
                shared.UcRecordingMedia{
                    EndAt: types.MustNewTimeFromString("2024-04-05T14:06:29.977Z"),
                    Language: unifiedgosdk.Pointer("en"),
                    RecordingDownloadURL: unifiedgosdk.Pointer("https://parched-wasabi.com/"),
                    StartAt: types.MustNewTimeFromString("2026-05-24T09:38:57.522Z"),
                    TranscriptDownloadURL: unifiedgosdk.Pointer("https://colossal-cuckoo.name"),
                    Transcripts: []shared.UcRecordingTranscript{
                        shared.UcRecordingTranscript{
                            EndAt: types.MustNewTimeFromString("2026-03-04T22:32:03.183Z"),
                            StartAt: types.MustNewTimeFromString("2024-06-06T10:28:37.346Z"),
                            Text: "Turpis maiores ducimus tero speculum.",
                        },
                        shared.UcRecordingTranscript{
                            EndAt: types.MustNewTimeFromString("2024-10-17T14:55:15.159Z"),
                            StartAt: types.MustNewTimeFromString("2026-08-12T09:01:21.271Z"),
                            Text: "Magnam consuasor uxor tergiversatio subseco.",
                        },
                    },
                },
            },
            Metadata: []shared.UcMetadata{},
            StartAt: types.MustNewTimeFromString("2023-04-29T05:48:51.376Z"),
            Type: shared.UcRecordingTypeOutbound.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-03-22T13:47:57.698Z"),
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

## GetCalendarRecording

Retrieves a single CalendarRecording by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

    res, err := s.Recording.GetCalendarRecording(ctx, operations.GetCalendarRecordingRequest{
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

    res, err := s.Recording.GetUcRecording(ctx, operations.GetUcRecordingRequest{
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

## ListCalendarRecordings

Returns a list of CalendarRecordings. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

    res, err := s.Recording.ListCalendarRecordings(ctx, operations.ListCalendarRecordingsRequest{
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

    res, err := s.Recording.ListUcRecordings(ctx, operations.ListUcRecordingsRequest{
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

    res, err := s.Recording.PatchUcRecording(ctx, operations.PatchUcRecordingRequest{
        UcRecording: shared.UcRecording{
            Contacts: []shared.UcContact{},
            ConversationIdentifier: unifiedgosdk.Pointer("2a237e64-913c-4af7-b737-c324f8983052"),
            CreatedAt: types.MustNewTimeFromString("2022-09-17T19:41:46.956Z"),
            EndAt: types.MustNewTimeFromString("2024-05-08T23:47:01.721Z"),
            ExpiresAt: types.MustNewTimeFromString("2026-05-05T14:34:39.048Z"),
            ID: unifiedgosdk.Pointer("2a237e64-913c-4af7-b737-c324f8983052"),
            Media: []shared.UcRecordingMedia{
                shared.UcRecordingMedia{
                    EndAt: types.MustNewTimeFromString("2024-04-05T14:06:29.987Z"),
                    Language: unifiedgosdk.Pointer("en"),
                    RecordingDownloadURL: unifiedgosdk.Pointer("https://parched-wasabi.com/"),
                    StartAt: types.MustNewTimeFromString("2026-05-24T09:38:57.545Z"),
                    TranscriptDownloadURL: unifiedgosdk.Pointer("https://colossal-cuckoo.name"),
                    Transcripts: []shared.UcRecordingTranscript{
                        shared.UcRecordingTranscript{
                            EndAt: types.MustNewTimeFromString("2026-03-04T22:32:03.204Z"),
                            StartAt: types.MustNewTimeFromString("2024-06-06T10:28:37.356Z"),
                            Text: "Turpis maiores ducimus tero speculum.",
                        },
                        shared.UcRecordingTranscript{
                            EndAt: types.MustNewTimeFromString("2024-10-17T14:55:15.172Z"),
                            StartAt: types.MustNewTimeFromString("2026-08-12T09:01:21.294Z"),
                            Text: "Magnam consuasor uxor tergiversatio subseco.",
                        },
                    },
                },
            },
            Metadata: []shared.UcMetadata{},
            StartAt: types.MustNewTimeFromString("2023-04-29T05:48:51.380Z"),
            Type: shared.UcRecordingTypeOutbound.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-03-22T13:47:57.713Z"),
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

    res, err := s.Recording.RemoveUcRecording(ctx, operations.RemoveUcRecordingRequest{
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

    res, err := s.Recording.UpdateUcRecording(ctx, operations.UpdateUcRecordingRequest{
        UcRecording: shared.UcRecording{
            Contacts: []shared.UcContact{},
            ConversationIdentifier: unifiedgosdk.Pointer("2a237e64-913c-4af7-b737-c324f8983052"),
            CreatedAt: types.MustNewTimeFromString("2022-09-17T19:41:46.956Z"),
            EndAt: types.MustNewTimeFromString("2024-05-08T23:47:01.721Z"),
            ExpiresAt: types.MustNewTimeFromString("2026-05-05T14:34:39.048Z"),
            ID: unifiedgosdk.Pointer("2a237e64-913c-4af7-b737-c324f8983052"),
            Media: []shared.UcRecordingMedia{
                shared.UcRecordingMedia{
                    EndAt: types.MustNewTimeFromString("2024-04-05T14:06:29.987Z"),
                    Language: unifiedgosdk.Pointer("en"),
                    RecordingDownloadURL: unifiedgosdk.Pointer("https://parched-wasabi.com/"),
                    StartAt: types.MustNewTimeFromString("2026-05-24T09:38:57.545Z"),
                    TranscriptDownloadURL: unifiedgosdk.Pointer("https://colossal-cuckoo.name"),
                    Transcripts: []shared.UcRecordingTranscript{
                        shared.UcRecordingTranscript{
                            EndAt: types.MustNewTimeFromString("2026-03-04T22:32:03.204Z"),
                            StartAt: types.MustNewTimeFromString("2024-06-06T10:28:37.356Z"),
                            Text: "Turpis maiores ducimus tero speculum.",
                        },
                        shared.UcRecordingTranscript{
                            EndAt: types.MustNewTimeFromString("2024-10-17T14:55:15.172Z"),
                            StartAt: types.MustNewTimeFromString("2026-08-12T09:01:21.294Z"),
                            Text: "Magnam consuasor uxor tergiversatio subseco.",
                        },
                    },
                },
            },
            Metadata: []shared.UcMetadata{},
            StartAt: types.MustNewTimeFromString("2023-04-29T05:48:51.380Z"),
            Type: shared.UcRecordingTypeOutbound.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-03-22T13:47:57.713Z"),
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