# Storage

## Overview

### Available Operations

* [CreateStorageFile](#createstoragefile) - Create a file
* [GetStorageFile](#getstoragefile) - Retrieve a file
* [ListStorageFiles](#liststoragefiles) - List all files
* [PatchStorageFile](#patchstoragefile) - Update a file
* [RemoveStorageFile](#removestoragefile) - Remove a file
* [UpdateStorageFile](#updatestoragefile) - Update a file

## CreateStorageFile

Create a file

### Example Usage

<!-- UsageSnippet language="go" operationID="createStorageFile" method="post" path="/storage/{connection_id}/file" example="storage_file" -->
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

    res, err := s.Storage.CreateStorageFile(ctx, operations.CreateStorageFileRequest{
        StorageFile: shared.StorageFile{
            CreatedAt: types.MustNewTimeFromString("2021-09-12T16:48:23.774Z"),
            Description: unifiedgosdk.Pointer("Deprimo crastinus cupiditate debilito cimentarius virgo acidus."),
            DownloadURL: unifiedgosdk.Pointer("https://clueless-laughter.biz"),
            Hash: unifiedgosdk.Pointer("d9d88dda7bc0acaa8472fdef2307f94a1fc355ce"),
            ID: unifiedgosdk.Pointer("5a3b2522-ba9a-4fc0-b63a-50257b278c9b"),
            Metadata: []shared.StorageMetadata{},
            MimeType: unifiedgosdk.Pointer("application/pdf"),
            Name: unifiedgosdk.Pointer("softly.pdf"),
            Size: unifiedgosdk.Pointer[float64](1134456.0),
            Tags: []string{
                "suasoria",
            },
            Type: shared.StorageFileTypeFolder.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-02-05T13:00:18.919Z"),
            Version: unifiedgosdk.Pointer("1"),
            WebURL: unifiedgosdk.Pointer("https://faraway-inspection.org"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.StorageFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateStorageFileRequest](../../pkg/models/operations/createstoragefilerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateStorageFileResponse](../../pkg/models/operations/createstoragefileresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetStorageFile

Retrieve a file

### Example Usage

<!-- UsageSnippet language="go" operationID="getStorageFile" method="get" path="/storage/{connection_id}/file/{id}" -->
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

    res, err := s.Storage.GetStorageFile(ctx, operations.GetStorageFileRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.StorageFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetStorageFileRequest](../../pkg/models/operations/getstoragefilerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetStorageFileResponse](../../pkg/models/operations/getstoragefileresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListStorageFiles

List all files

### Example Usage

<!-- UsageSnippet language="go" operationID="listStorageFiles" method="get" path="/storage/{connection_id}/file" -->
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

    res, err := s.Storage.ListStorageFiles(ctx, operations.ListStorageFilesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.StorageFiles != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListStorageFilesRequest](../../pkg/models/operations/liststoragefilesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListStorageFilesResponse](../../pkg/models/operations/liststoragefilesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchStorageFile

Update a file

### Example Usage

<!-- UsageSnippet language="go" operationID="patchStorageFile" method="patch" path="/storage/{connection_id}/file/{id}" example="storage_file" -->
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

    res, err := s.Storage.PatchStorageFile(ctx, operations.PatchStorageFileRequest{
        StorageFile: shared.StorageFile{
            CreatedAt: types.MustNewTimeFromString("2021-09-12T16:48:23.774Z"),
            Description: unifiedgosdk.Pointer("Deprimo crastinus cupiditate debilito cimentarius virgo acidus."),
            DownloadURL: unifiedgosdk.Pointer("https://clueless-laughter.biz"),
            Hash: unifiedgosdk.Pointer("d9d88dda7bc0acaa8472fdef2307f94a1fc355ce"),
            ID: unifiedgosdk.Pointer("78dc5fa1-2bf0-4000-a4ed-49fb5198ac5e"),
            Metadata: []shared.StorageMetadata{},
            MimeType: unifiedgosdk.Pointer("application/pdf"),
            Name: unifiedgosdk.Pointer("softly.pdf"),
            Size: unifiedgosdk.Pointer[float64](1134456.0),
            Tags: []string{
                "suasoria",
            },
            Type: shared.StorageFileTypeFolder.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-02-05T13:00:18.925Z"),
            Version: unifiedgosdk.Pointer("1"),
            WebURL: unifiedgosdk.Pointer("https://faraway-inspection.org"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.StorageFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchStorageFileRequest](../../pkg/models/operations/patchstoragefilerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchStorageFileResponse](../../pkg/models/operations/patchstoragefileresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveStorageFile

Remove a file

### Example Usage

<!-- UsageSnippet language="go" operationID="removeStorageFile" method="delete" path="/storage/{connection_id}/file/{id}" -->
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

    res, err := s.Storage.RemoveStorageFile(ctx, operations.RemoveStorageFileRequest{
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
| `request`                                                                                      | [operations.RemoveStorageFileRequest](../../pkg/models/operations/removestoragefilerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemoveStorageFileResponse](../../pkg/models/operations/removestoragefileresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateStorageFile

Update a file

### Example Usage

<!-- UsageSnippet language="go" operationID="updateStorageFile" method="put" path="/storage/{connection_id}/file/{id}" example="storage_file" -->
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

    res, err := s.Storage.UpdateStorageFile(ctx, operations.UpdateStorageFileRequest{
        StorageFile: shared.StorageFile{
            CreatedAt: types.MustNewTimeFromString("2021-09-12T16:48:23.774Z"),
            Description: unifiedgosdk.Pointer("Deprimo crastinus cupiditate debilito cimentarius virgo acidus."),
            DownloadURL: unifiedgosdk.Pointer("https://clueless-laughter.biz"),
            Hash: unifiedgosdk.Pointer("d9d88dda7bc0acaa8472fdef2307f94a1fc355ce"),
            ID: unifiedgosdk.Pointer("78dc5fa1-2bf0-4000-a4ed-49fb5198ac5e"),
            Metadata: []shared.StorageMetadata{},
            MimeType: unifiedgosdk.Pointer("application/pdf"),
            Name: unifiedgosdk.Pointer("softly.pdf"),
            Size: unifiedgosdk.Pointer[float64](1134456.0),
            Tags: []string{
                "suasoria",
            },
            Type: shared.StorageFileTypeFolder.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-02-05T13:00:18.925Z"),
            Version: unifiedgosdk.Pointer("1"),
            WebURL: unifiedgosdk.Pointer("https://faraway-inspection.org"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.StorageFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateStorageFileRequest](../../pkg/models/operations/updatestoragefilerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateStorageFileResponse](../../pkg/models/operations/updatestoragefileresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |