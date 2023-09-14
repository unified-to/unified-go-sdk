# Candidate

### Available Operations

* [DeleteAtsConnectionIDCandidateID](#deleteatsconnectionidcandidateid) - Remove a candidate
* [GetAtsConnectionIDCandidate](#getatsconnectionidcandidate) - List all candidates
* [GetAtsConnectionIDCandidateID](#getatsconnectionidcandidateid) - Retrieve a candidate
* [PatchAtsConnectionIDCandidateID](#patchatsconnectionidcandidateid) - Update a candidate
* [PostAtsConnectionIDCandidate](#postatsconnectionidcandidate) - Create a candidate
* [PutAtsConnectionIDCandidateID](#putatsconnectionidcandidateid) - Update a candidate

## DeleteAtsConnectionIDCandidateID

Remove a candidate

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.DeleteAtsConnectionIDCandidateIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Candidate.DeleteAtsConnectionIDCandidateID(ctx, operations.DeleteAtsConnectionIDCandidateIDRequest{
        ConnectionID: "voluptates",
        ID: "4825c1fc-0e11-45c8-8bff-918544ec42de",
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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.DeleteAtsConnectionIDCandidateIDRequest](../../models/operations/deleteatsconnectionidcandidateidrequest.md)   | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `security`                                                                                                                 | [operations.DeleteAtsConnectionIDCandidateIDSecurity](../../models/operations/deleteatsconnectionidcandidateidsecurity.md) | :heavy_check_mark:                                                                                                         | The security requirements to use for the request.                                                                          |


### Response

**[*operations.DeleteAtsConnectionIDCandidateIDResponse](../../models/operations/deleteatsconnectionidcandidateidresponse.md), error**


## GetAtsConnectionIDCandidate

List all candidates

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetAtsConnectionIDCandidateSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Candidate.GetAtsConnectionIDCandidate(ctx, operations.GetAtsConnectionIDCandidateRequest{
        ConnectionID: "doloribus",
        Limit: unifiedto.Float64(7737.11),
        Offset: unifiedto.Float64(7833.97),
        Order: unifiedto.String("accusamus"),
        Query: unifiedto.String("totam"),
        Sort: unifiedto.String("reiciendis"),
        UpdatedGte: types.MustTimeFromString("2022-06-05T16:37:51.499Z"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsCandidates != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetAtsConnectionIDCandidateRequest](../../models/operations/getatsconnectionidcandidaterequest.md)   | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `security`                                                                                                       | [operations.GetAtsConnectionIDCandidateSecurity](../../models/operations/getatsconnectionidcandidatesecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |


### Response

**[*operations.GetAtsConnectionIDCandidateResponse](../../models/operations/getatsconnectionidcandidateresponse.md), error**


## GetAtsConnectionIDCandidateID

Retrieve a candidate

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.GetAtsConnectionIDCandidateIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Candidate.GetAtsConnectionIDCandidateID(ctx, operations.GetAtsConnectionIDCandidateIDRequest{
        ConnectionID: "nihil",
        ID: "7773e635-62a7-4b40-8f05-e3d48fdaf313",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsCandidate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.GetAtsConnectionIDCandidateIDRequest](../../models/operations/getatsconnectionidcandidateidrequest.md)   | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `security`                                                                                                           | [operations.GetAtsConnectionIDCandidateIDSecurity](../../models/operations/getatsconnectionidcandidateidsecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |


### Response

**[*operations.GetAtsConnectionIDCandidateIDResponse](../../models/operations/getatsconnectionidcandidateidresponse.md), error**


## PatchAtsConnectionIDCandidateID

Update a candidate

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PatchAtsConnectionIDCandidateIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Candidate.PatchAtsConnectionIDCandidateID(ctx, operations.PatchAtsConnectionIDCandidateIDRequest{
        AtsCandidate: &shared.AtsCandidate{
            Address: &shared.PropertyAtsCandidateAddress{
                Address1: unifiedto.String("similique"),
                Address2: unifiedto.String("illo"),
                City: unifiedto.String("Greeley"),
                Country: unifiedto.String("Yemen"),
                CountryCode: unifiedto.String("SO"),
                PostalCode: unifiedto.String("21357-0614"),
                Region: unifiedto.String("sapiente"),
                RegionCode: unifiedto.String("consequuntur"),
            },
            CompanyName: unifiedto.String("veniam"),
            CreatedAt: types.MustTimeFromString("2021-01-31T23:06:28.366Z"),
            Emails: []shared.AtsEmail{
                shared.AtsEmail{
                    Email: "Dylan.Gerhold72@yahoo.com",
                    Type: shared.AtsEmailTypeHome.ToPointer(),
                },
            },
            ExternalID: unifiedto.String("minima"),
            ID: unifiedto.String("6c11f6c3-7a51-4262-8383-5bbc05a23a45"),
            ImageURL: unifiedto.String("quod"),
            Name: unifiedto.String("Tommie Schamberger"),
            Raw: &shared.PropertyAtsCandidateRaw{},
            Tags: []string{
                "assumenda",
            },
            Telephones: []shared.AtsTelephone{
                shared.AtsTelephone{
                    Telephone: "officiis",
                    Type: shared.AtsTelephoneTypeWork.ToPointer(),
                },
            },
            Title: unifiedto.String("Mr."),
            UpdatedAt: types.MustTimeFromString("2022-11-17T22:55:28.189Z"),
        },
        ConnectionID: "nobis",
        ID: "e2169e51-0019-4c6d-85e3-4762799bfbbe",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsCandidate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.PatchAtsConnectionIDCandidateIDRequest](../../models/operations/patchatsconnectionidcandidateidrequest.md)   | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `security`                                                                                                               | [operations.PatchAtsConnectionIDCandidateIDSecurity](../../models/operations/patchatsconnectionidcandidateidsecurity.md) | :heavy_check_mark:                                                                                                       | The security requirements to use for the request.                                                                        |


### Response

**[*operations.PatchAtsConnectionIDCandidateIDResponse](../../models/operations/patchatsconnectionidcandidateidresponse.md), error**


## PostAtsConnectionIDCandidate

Create a candidate

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PostAtsConnectionIDCandidateSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Candidate.PostAtsConnectionIDCandidate(ctx, operations.PostAtsConnectionIDCandidateRequest{
        AtsCandidate: &shared.AtsCandidate{
            Address: &shared.PropertyAtsCandidateAddress{
                Address1: unifiedto.String("laboriosam"),
                Address2: unifiedto.String("unde"),
                City: unifiedto.String("New Tyrel"),
                Country: unifiedto.String("Russian Federation"),
                CountryCode: unifiedto.String("BV"),
                PostalCode: unifiedto.String("72976-9471"),
                Region: unifiedto.String("illum"),
                RegionCode: unifiedto.String("nemo"),
            },
            CompanyName: unifiedto.String("illum"),
            CreatedAt: types.MustTimeFromString("2022-07-04T05:44:09.732Z"),
            Emails: []shared.AtsEmail{
                shared.AtsEmail{
                    Email: "Reynold_Walter32@yahoo.com",
                    Type: shared.AtsEmailTypeOther.ToPointer(),
                },
            },
            ExternalID: unifiedto.String("laborum"),
            ID: unifiedto.String("ea4c506a-8aa9-44c0-a644-cf5e9d9a4578"),
            ImageURL: unifiedto.String("fuga"),
            Name: unifiedto.String("Edmund Boyle"),
            Raw: &shared.PropertyAtsCandidateRaw{},
            Tags: []string{
                "laboriosam",
            },
            Telephones: []shared.AtsTelephone{
                shared.AtsTelephone{
                    Telephone: "doloremque",
                    Type: shared.AtsTelephoneTypeWork.ToPointer(),
                },
            },
            Title: unifiedto.String("Dr."),
            UpdatedAt: types.MustTimeFromString("2020-07-30T17:13:23.320Z"),
        },
        ConnectionID: "consequatur",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsCandidate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.PostAtsConnectionIDCandidateRequest](../../models/operations/postatsconnectionidcandidaterequest.md)   | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `security`                                                                                                         | [operations.PostAtsConnectionIDCandidateSecurity](../../models/operations/postatsconnectionidcandidatesecurity.md) | :heavy_check_mark:                                                                                                 | The security requirements to use for the request.                                                                  |


### Response

**[*operations.PostAtsConnectionIDCandidateResponse](../../models/operations/postatsconnectionidcandidateresponse.md), error**


## PutAtsConnectionIDCandidateID

Update a candidate

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedto.New()
    operationSecurity := operations.PutAtsConnectionIDCandidateIDSecurity{
            Jwt: "",
        }

    ctx := context.Background()
    res, err := s.Candidate.PutAtsConnectionIDCandidateID(ctx, operations.PutAtsConnectionIDCandidateIDRequest{
        AtsCandidate: &shared.AtsCandidate{
            Address: &shared.PropertyAtsCandidateAddress{
                Address1: unifiedto.String("eaque"),
                Address2: unifiedto.String("architecto"),
                City: unifiedto.String("Sawaynmouth"),
                Country: unifiedto.String("Azerbaijan"),
                CountryCode: unifiedto.String("CH"),
                PostalCode: unifiedto.String("18705-9959"),
                Region: unifiedto.String("alias"),
                RegionCode: unifiedto.String("asperiores"),
            },
            CompanyName: unifiedto.String("rem"),
            CreatedAt: types.MustTimeFromString("2022-08-15T01:59:03.683Z"),
            Emails: []shared.AtsEmail{
                shared.AtsEmail{
                    Email: "Winona45@hotmail.com",
                    Type: shared.AtsEmailTypeHome.ToPointer(),
                },
            },
            ExternalID: unifiedto.String("eligendi"),
            ID: unifiedto.String("13e902c1-4125-4b09-a0a6-68151a472af9"),
            ImageURL: unifiedto.String("sed"),
            Name: unifiedto.String("Kendra Hauck"),
            Raw: &shared.PropertyAtsCandidateRaw{},
            Tags: []string{
                "excepturi",
            },
            Telephones: []shared.AtsTelephone{
                shared.AtsTelephone{
                    Telephone: "maiores",
                    Type: shared.AtsTelephoneTypeOther.ToPointer(),
                },
            },
            Title: unifiedto.String("Mrs."),
            UpdatedAt: types.MustTimeFromString("2022-04-23T19:34:08.217Z"),
        },
        ConnectionID: "nemo",
        ID: "0cf876ff-b901-4c6e-8bb4-e243cf789ffa",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsCandidate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.PutAtsConnectionIDCandidateIDRequest](../../models/operations/putatsconnectionidcandidateidrequest.md)   | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `security`                                                                                                           | [operations.PutAtsConnectionIDCandidateIDSecurity](../../models/operations/putatsconnectionidcandidateidsecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |


### Response

**[*operations.PutAtsConnectionIDCandidateIDResponse](../../models/operations/putatsconnectionidcandidateidresponse.md), error**

