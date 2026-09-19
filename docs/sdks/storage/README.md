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
            Data: unifiedgosdk.Pointer("data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZlcnNpb249IjEuMSIgYmFzZVByb2ZpbGU9ImZ1bGwiIHdpZHRoPSI4MzIiIGhlaWdodD0iMTg2MSI+PHJlY3Qgd2lkdGg9IjEwMCUiIGhlaWdodD0iMTAwJSIgZmlsbD0iIzFmYzM1NSIvPjx0ZXh0IHg9IjQxNiIgeT0iOTMwLjUiIGZvbnQtc2l6ZT0iMjAiIGFsaWdubWVudC1iYXNlbGluZT0ibWlkZGxlIiB0ZXh0LWFuY2hvcj0ibWlkZGxlIiBmaWxsPSJ3aGl0ZSI+ODMyeDE4NjE8L3RleHQ+PC9zdmc+"),
            Description: unifiedgosdk.Pointer("Crastinus cupiditate debilito cimentarius virgo."),
            DownloadURL: unifiedgosdk.Pointer("https://stingy-casement.name/"),
            Hash: unifiedgosdk.Pointer("fe6a659e-75cd-4079-9b76-351f9af2205a"),
            ID: unifiedgosdk.Pointer("e21d67dd-a021-423a-bcc1-1ff0764566b3"),
            MimeType: unifiedgosdk.Pointer("FOLDER"),
            Name: unifiedgosdk.Pointer("softly.tiff"),
            References: []shared.StorageReference{
                shared.StorageReference{
                    ID: unifiedgosdk.Pointer("ab705f3b-e368-4a94-8b22-d5f693c14a76"),
                    Name: unifiedgosdk.Pointer("tamisium viduo odio cauda"),
                    Type: unifiedgosdk.Pointer("accounting_bill"),
                },
                shared.StorageReference{
                    ID: unifiedgosdk.Pointer("9f0f694e-b6f4-4c12-b5f6-ab08d4e81140"),
                    Name: unifiedgosdk.Pointer("quia"),
                    Type: unifiedgosdk.Pointer("accounting_expense"),
                },
            },
            Size: unifiedgosdk.Pointer[float64](10276.0),
            Tags: []string{
                "spoliatio",
            },
            Type: shared.StorageFileTypeFile.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-01-27T18:05:36.893Z"),
            Version: unifiedgosdk.Pointer("1"),
            WebURL: unifiedgosdk.Pointer("https://sandy-distinction.info/"),
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
            Data: unifiedgosdk.Pointer("data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZlcnNpb249IjEuMSIgYmFzZVByb2ZpbGU9ImZ1bGwiIHdpZHRoPSI4MzIiIGhlaWdodD0iMTg2MSI+PHJlY3Qgd2lkdGg9IjEwMCUiIGhlaWdodD0iMTAwJSIgZmlsbD0iIzFmYzM1NSIvPjx0ZXh0IHg9IjQxNiIgeT0iOTMwLjUiIGZvbnQtc2l6ZT0iMjAiIGFsaWdubWVudC1iYXNlbGluZT0ibWlkZGxlIiB0ZXh0LWFuY2hvcj0ibWlkZGxlIiBmaWxsPSJ3aGl0ZSI+ODMyeDE4NjE8L3RleHQ+PC9zdmc+"),
            Description: unifiedgosdk.Pointer("Crastinus cupiditate debilito cimentarius virgo."),
            DownloadURL: unifiedgosdk.Pointer("https://stingy-casement.name/"),
            Hash: unifiedgosdk.Pointer("fe6a659e-75cd-4079-9b76-351f9af2205a"),
            ID: unifiedgosdk.Pointer("154ddfa8-27f8-4e14-aac9-6d94f83b72a9"),
            MimeType: unifiedgosdk.Pointer("FOLDER"),
            Name: unifiedgosdk.Pointer("softly.tiff"),
            References: []shared.StorageReference{
                shared.StorageReference{
                    ID: unifiedgosdk.Pointer("ab705f3b-e368-4a94-8b22-d5f693c14a76"),
                    Name: unifiedgosdk.Pointer("tamisium viduo odio cauda"),
                    Type: unifiedgosdk.Pointer("accounting_bill"),
                },
                shared.StorageReference{
                    ID: unifiedgosdk.Pointer("9f0f694e-b6f4-4c12-b5f6-ab08d4e81140"),
                    Name: unifiedgosdk.Pointer("quia"),
                    Type: unifiedgosdk.Pointer("accounting_expense"),
                },
            },
            Size: unifiedgosdk.Pointer[float64](10276.0),
            Tags: []string{
                "spoliatio",
            },
            Type: shared.StorageFileTypeFile.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-01-27T18:05:36.898Z"),
            Version: unifiedgosdk.Pointer("1"),
            WebURL: unifiedgosdk.Pointer("https://sandy-distinction.info/"),
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
            Data: unifiedgosdk.Pointer("data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZlcnNpb249IjEuMSIgYmFzZVByb2ZpbGU9ImZ1bGwiIHdpZHRoPSI4MzIiIGhlaWdodD0iMTg2MSI+PHJlY3Qgd2lkdGg9IjEwMCUiIGhlaWdodD0iMTAwJSIgZmlsbD0iIzFmYzM1NSIvPjx0ZXh0IHg9IjQxNiIgeT0iOTMwLjUiIGZvbnQtc2l6ZT0iMjAiIGFsaWdubWVudC1iYXNlbGluZT0ibWlkZGxlIiB0ZXh0LWFuY2hvcj0ibWlkZGxlIiBmaWxsPSJ3aGl0ZSI+ODMyeDE4NjE8L3RleHQ+PC9zdmc+"),
            Description: unifiedgosdk.Pointer("Crastinus cupiditate debilito cimentarius virgo."),
            DownloadURL: unifiedgosdk.Pointer("https://stingy-casement.name/"),
            Hash: unifiedgosdk.Pointer("fe6a659e-75cd-4079-9b76-351f9af2205a"),
            ID: unifiedgosdk.Pointer("154ddfa8-27f8-4e14-aac9-6d94f83b72a9"),
            MimeType: unifiedgosdk.Pointer("FOLDER"),
            Name: unifiedgosdk.Pointer("softly.tiff"),
            References: []shared.StorageReference{
                shared.StorageReference{
                    ID: unifiedgosdk.Pointer("ab705f3b-e368-4a94-8b22-d5f693c14a76"),
                    Name: unifiedgosdk.Pointer("tamisium viduo odio cauda"),
                    Type: unifiedgosdk.Pointer("accounting_bill"),
                },
                shared.StorageReference{
                    ID: unifiedgosdk.Pointer("9f0f694e-b6f4-4c12-b5f6-ab08d4e81140"),
                    Name: unifiedgosdk.Pointer("quia"),
                    Type: unifiedgosdk.Pointer("accounting_expense"),
                },
            },
            Size: unifiedgosdk.Pointer[float64](10276.0),
            Tags: []string{
                "spoliatio",
            },
            Type: shared.StorageFileTypeFile.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-01-27T18:05:36.898Z"),
            Version: unifiedgosdk.Pointer("1"),
            WebURL: unifiedgosdk.Pointer("https://sandy-distinction.info/"),
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