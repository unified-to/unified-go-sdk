# Storage
(*Storage*)

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

```go
package main

import(
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := unifiedgosdk.New()


    operationSecurity := operations.CreateStorageFileSecurity{
            Jwt: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Storage.CreateStorageFile(ctx, operations.CreateStorageFileRequest{
        ConnectionID: "<value>",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.StorageFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateStorageFileRequest](../../pkg/models/operations/createstoragefilerequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.CreateStorageFileSecurity](../../pkg/models/operations/createstoragefilesecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.CreateStorageFileResponse](../../pkg/models/operations/createstoragefileresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetStorageFile

Retrieve a file

### Example Usage

```go
package main

import(
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := unifiedgosdk.New()


    operationSecurity := operations.GetStorageFileSecurity{
            Jwt: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Storage.GetStorageFile(ctx, operations.GetStorageFileRequest{
        ConnectionID: "<value>",
        ID: "<id>",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.StorageFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetStorageFileRequest](../../pkg/models/operations/getstoragefilerequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.GetStorageFileSecurity](../../pkg/models/operations/getstoragefilesecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.GetStorageFileResponse](../../pkg/models/operations/getstoragefileresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListStorageFiles

List all files

### Example Usage

```go
package main

import(
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := unifiedgosdk.New()


    operationSecurity := operations.ListStorageFilesSecurity{
            Jwt: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Storage.ListStorageFiles(ctx, operations.ListStorageFilesRequest{
        ConnectionID: "<value>",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.StorageFiles != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListStorageFilesRequest](../../pkg/models/operations/liststoragefilesrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.ListStorageFilesSecurity](../../pkg/models/operations/liststoragefilessecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.ListStorageFilesResponse](../../pkg/models/operations/liststoragefilesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PatchStorageFile

Update a file

### Example Usage

```go
package main

import(
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := unifiedgosdk.New()


    operationSecurity := operations.PatchStorageFileSecurity{
            Jwt: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Storage.PatchStorageFile(ctx, operations.PatchStorageFileRequest{
        ConnectionID: "<value>",
        ID: "<id>",
    }, operationSecurity)
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
| `request`                                                                                      | [operations.PatchStorageFileRequest](../../pkg/models/operations/patchstoragefilerequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.PatchStorageFileSecurity](../../pkg/models/operations/patchstoragefilesecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.PatchStorageFileResponse](../../pkg/models/operations/patchstoragefileresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RemoveStorageFile

Remove a file

### Example Usage

```go
package main

import(
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"context"
	"log"
	"net/http"
)

func main() {
    s := unifiedgosdk.New()


    operationSecurity := operations.RemoveStorageFileSecurity{
            Jwt: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Storage.RemoveStorageFile(ctx, operations.RemoveStorageFileRequest{
        ConnectionID: "<value>",
        ID: "<id>",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.RemoveStorageFileRequest](../../pkg/models/operations/removestoragefilerequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.RemoveStorageFileSecurity](../../pkg/models/operations/removestoragefilesecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.RemoveStorageFileResponse](../../pkg/models/operations/removestoragefileresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateStorageFile

Update a file

### Example Usage

```go
package main

import(
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := unifiedgosdk.New()


    operationSecurity := operations.UpdateStorageFileSecurity{
            Jwt: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Storage.UpdateStorageFile(ctx, operations.UpdateStorageFileRequest{
        ConnectionID: "<value>",
        ID: "<id>",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.StorageFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateStorageFileRequest](../../pkg/models/operations/updatestoragefilerequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.UpdateStorageFileSecurity](../../pkg/models/operations/updatestoragefilesecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.UpdateStorageFileResponse](../../pkg/models/operations/updatestoragefileresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
