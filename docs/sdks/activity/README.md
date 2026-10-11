# Activity

## Overview

### Available Operations

* [CreateAtsActivity](#createatsactivity) - Create an activity
* [CreateLmsActivity](#createlmsactivity) - Create an activity
* [GetAtsActivity](#getatsactivity) - Retrieve an activity
* [GetClubsActivity](#getclubsactivity) - Retrieve an activity
* [GetLmsActivity](#getlmsactivity) - Retrieve an activity
* [ListAtsActivities](#listatsactivities) - List all activities
* [ListClubsActivities](#listclubsactivities) - List all activities
* [ListLmsActivities](#listlmsactivities) - List all activities
* [PatchAtsActivity](#patchatsactivity) - Update an activity
* [PatchLmsActivity](#patchlmsactivity) - Update an activity
* [RemoveAtsActivity](#removeatsactivity) - Remove an activity
* [RemoveLmsActivity](#removelmsactivity) - Remove an activity
* [UpdateAtsActivity](#updateatsactivity) - Update an activity
* [UpdateLmsActivity](#updatelmsactivity) - Update an activity

## CreateAtsActivity

Create an activity

### Example Usage

<!-- UsageSnippet language="go" operationID="createAtsActivity" method="post" path="/ats/{connection_id}/activity" example="ats_activity" -->
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

    res, err := s.Activity.CreateAtsActivity(ctx, operations.CreateAtsActivityRequest{
        AtsActivity: shared.AtsActivity{
            Bcc: []shared.AtsEmail{
                shared.AtsEmail{
                    Email: "Mabel_Schuppe-Schowalter42@hotmail.com",
                    Name: unifiedgosdk.Pointer("Rochelle Franey-Bechtelar"),
                    Type: shared.AtsEmailTypeHome.ToPointer(),
                },
            },
            Cc: []shared.AtsEmail{
                shared.AtsEmail{
                    Email: "Sasha24@hotmail.com",
                    Name: unifiedgosdk.Pointer("Dr. Elbert Kuvalis"),
                    Type: shared.AtsEmailTypeHome.ToPointer(),
                },
                shared.AtsEmail{
                    Email: "Rosetta_Donnelly@gmail.com",
                    Name: unifiedgosdk.Pointer("Ramon Daniel"),
                    Type: shared.AtsEmailTypeOther.ToPointer(),
                },
                shared.AtsEmail{
                    Email: "Kathryne_Jast@yahoo.com",
                    Name: unifiedgosdk.Pointer("Christian Jacobson"),
                    Type: shared.AtsEmailTypeOther.ToPointer(),
                },
                shared.AtsEmail{
                    Email: "Eldred95@yahoo.com",
                    Name: unifiedgosdk.Pointer("Edna Bogan"),
                    Type: shared.AtsEmailTypeOther.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2022-08-07T03:16:43.865Z"),
            Description: unifiedgosdk.Pointer("Amplus."),
            From: &shared.PropertyAtsActivityFrom{
                Email: "Norwood.Wiza47@yahoo.com",
                Name: unifiedgosdk.Pointer("Toby Grant"),
                Type: shared.PropertyAtsActivityFromTypeOther.ToPointer(),
            },
            ID: unifiedgosdk.Pointer("ef9dfaf0-426d-4aed-bf1d-9f98c1559b90"),
            IsPrivate: unifiedgosdk.Pointer(false),
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("5f9f0f86-6634-4803-95dd-c9a7b0d79b45"),
                    Namespace: unifiedgosdk.Pointer("activity"),
                    Slug: unifiedgosdk.Pointer("acer"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "Pauci eius cena adamo summisse arguo pectus communis arcesso tergeo.",
                    )),
                },
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("cde549f1-354a-4d72-86c4-365de7fed7fa"),
                    Namespace: unifiedgosdk.Pointer("activity"),
                    Slug: unifiedgosdk.Pointer("tremo"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "Amita delectus dicta temptatio utroque ex.",
                    )),
                },
            },
            SubType: unifiedgosdk.Pointer("TASK"),
            Title: unifiedgosdk.Pointer("Senior Interactions Manager"),
            To: []shared.AtsEmail{
                shared.AtsEmail{
                    Email: "Sister91@hotmail.com",
                    Name: unifiedgosdk.Pointer("Eddie Nienow PhD"),
                    Type: shared.AtsEmailTypeWork.ToPointer(),
                },
            },
            Type: shared.AtsActivityTypeTask.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-03-26T13:55:33.902Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsActivity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateAtsActivityRequest](../../pkg/models/operations/createatsactivityrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateAtsActivityResponse](../../pkg/models/operations/createatsactivityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateLmsActivity

Create an activity

### Example Usage

<!-- UsageSnippet language="go" operationID="createLmsActivity" method="post" path="/lms/{connection_id}/activity" example="lms_activity" -->
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

    res, err := s.Activity.CreateLmsActivity(ctx, operations.CreateLmsActivityRequest{
        LmsActivity: shared.LmsActivity{
            AssignedGrade: unifiedgosdk.Pointer("summopere"),
            CompletedAt: types.MustNewTimeFromString("2025-04-23T20:06:36.025Z"),
            CreatedAt: types.MustNewTimeFromString("2020-10-17T01:25:21.745Z"),
            DurationMinutes: unifiedgosdk.Pointer[float64](55.0),
            ID: unifiedgosdk.Pointer("d0e82a3b-42f0-48e4-b262-a872942dd28d"),
            IsCompleted: unifiedgosdk.Pointer(true),
            ProgressPercentage: unifiedgosdk.Pointer[float64](100.0),
            StartedAt: types.MustNewTimeFromString("2023-12-24T04:54:05.825Z"),
            UpdatedAt: types.MustNewTimeFromString("2022-01-28T20:12:59.333Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsActivity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateLmsActivityRequest](../../pkg/models/operations/createlmsactivityrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateLmsActivityResponse](../../pkg/models/operations/createlmsactivityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAtsActivity

Retrieve an activity

### Example Usage

<!-- UsageSnippet language="go" operationID="getAtsActivity" method="get" path="/ats/{connection_id}/activity/{id}" -->
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

    res, err := s.Activity.GetAtsActivity(ctx, operations.GetAtsActivityRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsActivity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetAtsActivityRequest](../../pkg/models/operations/getatsactivityrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetAtsActivityResponse](../../pkg/models/operations/getatsactivityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetClubsActivity

Retrieve an activity

### Example Usage

<!-- UsageSnippet language="go" operationID="getClubsActivity" method="get" path="/clubs/{connection_id}/activity/{id}" -->
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

    res, err := s.Activity.GetClubsActivity(ctx, operations.GetClubsActivityRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ClubsActivity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetClubsActivityRequest](../../pkg/models/operations/getclubsactivityrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetClubsActivityResponse](../../pkg/models/operations/getclubsactivityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetLmsActivity

Retrieve an activity

### Example Usage

<!-- UsageSnippet language="go" operationID="getLmsActivity" method="get" path="/lms/{connection_id}/activity/{id}" -->
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

    res, err := s.Activity.GetLmsActivity(ctx, operations.GetLmsActivityRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsActivity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetLmsActivityRequest](../../pkg/models/operations/getlmsactivityrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetLmsActivityResponse](../../pkg/models/operations/getlmsactivityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAtsActivities

List all activities

### Example Usage

<!-- UsageSnippet language="go" operationID="listAtsActivities" method="get" path="/ats/{connection_id}/activity" -->
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

    res, err := s.Activity.ListAtsActivities(ctx, operations.ListAtsActivitiesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsActivities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListAtsActivitiesRequest](../../pkg/models/operations/listatsactivitiesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListAtsActivitiesResponse](../../pkg/models/operations/listatsactivitiesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListClubsActivities

List all activities

### Example Usage

<!-- UsageSnippet language="go" operationID="listClubsActivities" method="get" path="/clubs/{connection_id}/activity" -->
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

    res, err := s.Activity.ListClubsActivities(ctx, operations.ListClubsActivitiesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ClubsActivities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListClubsActivitiesRequest](../../pkg/models/operations/listclubsactivitiesrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListClubsActivitiesResponse](../../pkg/models/operations/listclubsactivitiesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListLmsActivities

List all activities

### Example Usage

<!-- UsageSnippet language="go" operationID="listLmsActivities" method="get" path="/lms/{connection_id}/activity" -->
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

    res, err := s.Activity.ListLmsActivities(ctx, operations.ListLmsActivitiesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsActivities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListLmsActivitiesRequest](../../pkg/models/operations/listlmsactivitiesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListLmsActivitiesResponse](../../pkg/models/operations/listlmsactivitiesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchAtsActivity

Update an activity

### Example Usage

<!-- UsageSnippet language="go" operationID="patchAtsActivity" method="patch" path="/ats/{connection_id}/activity/{id}" example="ats_activity" -->
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

    res, err := s.Activity.PatchAtsActivity(ctx, operations.PatchAtsActivityRequest{
        AtsActivity: shared.AtsActivity{
            Bcc: []shared.AtsEmail{
                shared.AtsEmail{
                    Email: "Mabel_Schuppe-Schowalter42@hotmail.com",
                    Name: unifiedgosdk.Pointer("Rochelle Franey-Bechtelar"),
                    Type: shared.AtsEmailTypeHome.ToPointer(),
                },
            },
            Cc: []shared.AtsEmail{
                shared.AtsEmail{
                    Email: "Sasha24@hotmail.com",
                    Name: unifiedgosdk.Pointer("Dr. Elbert Kuvalis"),
                    Type: shared.AtsEmailTypeHome.ToPointer(),
                },
                shared.AtsEmail{
                    Email: "Rosetta_Donnelly@gmail.com",
                    Name: unifiedgosdk.Pointer("Ramon Daniel"),
                    Type: shared.AtsEmailTypeOther.ToPointer(),
                },
                shared.AtsEmail{
                    Email: "Kathryne_Jast@yahoo.com",
                    Name: unifiedgosdk.Pointer("Christian Jacobson"),
                    Type: shared.AtsEmailTypeOther.ToPointer(),
                },
                shared.AtsEmail{
                    Email: "Eldred95@yahoo.com",
                    Name: unifiedgosdk.Pointer("Edna Bogan"),
                    Type: shared.AtsEmailTypeOther.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2022-08-07T03:16:43.865Z"),
            Description: unifiedgosdk.Pointer("Amplus."),
            From: &shared.PropertyAtsActivityFrom{
                Email: "Norwood.Wiza47@yahoo.com",
                Name: unifiedgosdk.Pointer("Toby Grant"),
                Type: shared.PropertyAtsActivityFromTypeOther.ToPointer(),
            },
            ID: unifiedgosdk.Pointer("fc9ba3df-0d13-4d4b-8f42-342fc48d6194"),
            IsPrivate: unifiedgosdk.Pointer(false),
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("7d680aa0-5d11-4c16-971a-1d7fbe99a8e8"),
                    Namespace: unifiedgosdk.Pointer("activity"),
                    Slug: unifiedgosdk.Pointer("acer"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "Pauci eius cena adamo summisse arguo pectus communis arcesso tergeo.",
                    )),
                },
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("961fc2c5-f2b0-4db6-839d-79c93436486f"),
                    Namespace: unifiedgosdk.Pointer("activity"),
                    Slug: unifiedgosdk.Pointer("tremo"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "Amita delectus dicta temptatio utroque ex.",
                    )),
                },
            },
            SubType: unifiedgosdk.Pointer("TASK"),
            Title: unifiedgosdk.Pointer("Senior Interactions Manager"),
            To: []shared.AtsEmail{
                shared.AtsEmail{
                    Email: "Sister91@hotmail.com",
                    Name: unifiedgosdk.Pointer("Eddie Nienow PhD"),
                    Type: shared.AtsEmailTypeWork.ToPointer(),
                },
            },
            Type: shared.AtsActivityTypeTask.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-03-26T13:55:33.926Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsActivity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchAtsActivityRequest](../../pkg/models/operations/patchatsactivityrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchAtsActivityResponse](../../pkg/models/operations/patchatsactivityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchLmsActivity

Update an activity

### Example Usage

<!-- UsageSnippet language="go" operationID="patchLmsActivity" method="patch" path="/lms/{connection_id}/activity/{id}" example="lms_activity" -->
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

    res, err := s.Activity.PatchLmsActivity(ctx, operations.PatchLmsActivityRequest{
        LmsActivity: shared.LmsActivity{
            AssignedGrade: unifiedgosdk.Pointer("summopere"),
            CompletedAt: types.MustNewTimeFromString("2025-04-23T20:06:36.029Z"),
            CreatedAt: types.MustNewTimeFromString("2020-10-17T01:25:21.745Z"),
            DurationMinutes: unifiedgosdk.Pointer[float64](55.0),
            ID: unifiedgosdk.Pointer("5237b4ab-6179-4efd-906f-95427b763e46"),
            IsCompleted: unifiedgosdk.Pointer(true),
            ProgressPercentage: unifiedgosdk.Pointer[float64](100.0),
            StartedAt: types.MustNewTimeFromString("2023-12-24T04:54:05.825Z"),
            UpdatedAt: types.MustNewTimeFromString("2022-01-28T20:12:59.335Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsActivity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchLmsActivityRequest](../../pkg/models/operations/patchlmsactivityrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchLmsActivityResponse](../../pkg/models/operations/patchlmsactivityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAtsActivity

Remove an activity

### Example Usage

<!-- UsageSnippet language="go" operationID="removeAtsActivity" method="delete" path="/ats/{connection_id}/activity/{id}" -->
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

    res, err := s.Activity.RemoveAtsActivity(ctx, operations.RemoveAtsActivityRequest{
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
| `request`                                                                                      | [operations.RemoveAtsActivityRequest](../../pkg/models/operations/removeatsactivityrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemoveAtsActivityResponse](../../pkg/models/operations/removeatsactivityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveLmsActivity

Remove an activity

### Example Usage

<!-- UsageSnippet language="go" operationID="removeLmsActivity" method="delete" path="/lms/{connection_id}/activity/{id}" -->
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

    res, err := s.Activity.RemoveLmsActivity(ctx, operations.RemoveLmsActivityRequest{
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
| `request`                                                                                      | [operations.RemoveLmsActivityRequest](../../pkg/models/operations/removelmsactivityrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemoveLmsActivityResponse](../../pkg/models/operations/removelmsactivityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateAtsActivity

Update an activity

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAtsActivity" method="put" path="/ats/{connection_id}/activity/{id}" example="ats_activity" -->
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

    res, err := s.Activity.UpdateAtsActivity(ctx, operations.UpdateAtsActivityRequest{
        AtsActivity: shared.AtsActivity{
            Bcc: []shared.AtsEmail{
                shared.AtsEmail{
                    Email: "Mabel_Schuppe-Schowalter42@hotmail.com",
                    Name: unifiedgosdk.Pointer("Rochelle Franey-Bechtelar"),
                    Type: shared.AtsEmailTypeHome.ToPointer(),
                },
            },
            Cc: []shared.AtsEmail{
                shared.AtsEmail{
                    Email: "Sasha24@hotmail.com",
                    Name: unifiedgosdk.Pointer("Dr. Elbert Kuvalis"),
                    Type: shared.AtsEmailTypeHome.ToPointer(),
                },
                shared.AtsEmail{
                    Email: "Rosetta_Donnelly@gmail.com",
                    Name: unifiedgosdk.Pointer("Ramon Daniel"),
                    Type: shared.AtsEmailTypeOther.ToPointer(),
                },
                shared.AtsEmail{
                    Email: "Kathryne_Jast@yahoo.com",
                    Name: unifiedgosdk.Pointer("Christian Jacobson"),
                    Type: shared.AtsEmailTypeOther.ToPointer(),
                },
                shared.AtsEmail{
                    Email: "Eldred95@yahoo.com",
                    Name: unifiedgosdk.Pointer("Edna Bogan"),
                    Type: shared.AtsEmailTypeOther.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2022-08-07T03:16:43.865Z"),
            Description: unifiedgosdk.Pointer("Amplus."),
            From: &shared.PropertyAtsActivityFrom{
                Email: "Norwood.Wiza47@yahoo.com",
                Name: unifiedgosdk.Pointer("Toby Grant"),
                Type: shared.PropertyAtsActivityFromTypeOther.ToPointer(),
            },
            ID: unifiedgosdk.Pointer("fc9ba3df-0d13-4d4b-8f42-342fc48d6194"),
            IsPrivate: unifiedgosdk.Pointer(false),
            Metadata: []shared.AtsMetadata{
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("7d680aa0-5d11-4c16-971a-1d7fbe99a8e8"),
                    Namespace: unifiedgosdk.Pointer("activity"),
                    Slug: unifiedgosdk.Pointer("acer"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "Pauci eius cena adamo summisse arguo pectus communis arcesso tergeo.",
                    )),
                },
                shared.AtsMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateAtsMetadataExtraDataMapOfAny(
                        map[string]any{

                        },
                    )),
                    Format: shared.AtsMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("961fc2c5-f2b0-4db6-839d-79c93436486f"),
                    Namespace: unifiedgosdk.Pointer("activity"),
                    Slug: unifiedgosdk.Pointer("tremo"),
                    Value: unifiedgosdk.Pointer(shared.CreateAtsMetadataValueStr(
                        "Amita delectus dicta temptatio utroque ex.",
                    )),
                },
            },
            SubType: unifiedgosdk.Pointer("TASK"),
            Title: unifiedgosdk.Pointer("Senior Interactions Manager"),
            To: []shared.AtsEmail{
                shared.AtsEmail{
                    Email: "Sister91@hotmail.com",
                    Name: unifiedgosdk.Pointer("Eddie Nienow PhD"),
                    Type: shared.AtsEmailTypeWork.ToPointer(),
                },
            },
            Type: shared.AtsActivityTypeTask.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-03-26T13:55:33.926Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AtsActivity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateAtsActivityRequest](../../pkg/models/operations/updateatsactivityrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateAtsActivityResponse](../../pkg/models/operations/updateatsactivityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateLmsActivity

Update an activity

### Example Usage

<!-- UsageSnippet language="go" operationID="updateLmsActivity" method="put" path="/lms/{connection_id}/activity/{id}" example="lms_activity" -->
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

    res, err := s.Activity.UpdateLmsActivity(ctx, operations.UpdateLmsActivityRequest{
        LmsActivity: shared.LmsActivity{
            AssignedGrade: unifiedgosdk.Pointer("summopere"),
            CompletedAt: types.MustNewTimeFromString("2025-04-23T20:06:36.029Z"),
            CreatedAt: types.MustNewTimeFromString("2020-10-17T01:25:21.745Z"),
            DurationMinutes: unifiedgosdk.Pointer[float64](55.0),
            ID: unifiedgosdk.Pointer("5237b4ab-6179-4efd-906f-95427b763e46"),
            IsCompleted: unifiedgosdk.Pointer(true),
            ProgressPercentage: unifiedgosdk.Pointer[float64](100.0),
            StartedAt: types.MustNewTimeFromString("2023-12-24T04:54:05.825Z"),
            UpdatedAt: types.MustNewTimeFromString("2022-01-28T20:12:59.335Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LmsActivity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateLmsActivityRequest](../../pkg/models/operations/updatelmsactivityrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateLmsActivityResponse](../../pkg/models/operations/updatelmsactivityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |