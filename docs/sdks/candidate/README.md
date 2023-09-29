# Candidate
(*Candidate*)

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
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Candidate.DeleteAtsConnectionIDCandidateID(ctx, operations.DeleteAtsConnectionIDCandidateIDRequest{
        ConnectionID: "multimedia",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.DeleteAtsConnectionIDCandidateIDRequest](../../models/operations/deleteatsconnectionidcandidateidrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


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
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Candidate.GetAtsConnectionIDCandidate(ctx, operations.GetAtsConnectionIDCandidateRequest{
        ConnectionID: "Northwest forceful Moore",
        Limit: unifiedgosdk.Float64(2623.89),
        Offset: unifiedgosdk.Float64(7811.91),
        Order: unifiedgosdk.String("Mouse whether deploy"),
        Query: unifiedgosdk.String("pink"),
        Sort: unifiedgosdk.String("huzzah thistle"),
        UpdatedGte: types.MustTimeFromString("2022-03-13T15:14:03.645Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsCandidates != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetAtsConnectionIDCandidateRequest](../../models/operations/getatsconnectionidcandidaterequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


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
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Candidate.GetAtsConnectionIDCandidateID(ctx, operations.GetAtsConnectionIDCandidateIDRequest{
        ConnectionID: "ha Loan",
        ID: "<ID>",
    })
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
| `request`                                                                                                          | [operations.GetAtsConnectionIDCandidateIDRequest](../../models/operations/getatsconnectionidcandidateidrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


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
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Candidate.PatchAtsConnectionIDCandidateID(ctx, operations.PatchAtsConnectionIDCandidateIDRequest{
        AtsCandidate: &shared.AtsCandidate{
            Address: &shared.PropertyAtsCandidateAddress{
                Address1: unifiedgosdk.String("closely Goyette plus"),
                Address2: unifiedgosdk.String("culpa"),
                City: unifiedgosdk.String("Darrinshire"),
                Country: unifiedgosdk.String("Mongolia"),
                CountryCode: unifiedgosdk.String("GW"),
                PostalCode: unifiedgosdk.String("05275"),
                Region: unifiedgosdk.String("TLS calculating"),
                RegionCode: unifiedgosdk.String("up Argon Internal"),
            },
            CompanyName: unifiedgosdk.String("Fadel, Schulist and Koss"),
            CreatedAt: types.MustTimeFromString("2022-12-09T07:16:54.728Z"),
            Emails: []shared.AtsEmail{
                shared.AtsEmail{
                    Email: "Gregory63@gmail.com",
                    Type: shared.AtsEmailTypeOther.ToPointer(),
                },
            },
            ExternalID: unifiedgosdk.String("Elegant"),
            ID: unifiedgosdk.String("<ID>"),
            ImageURL: unifiedgosdk.String("Tricycle Yttrium Hybrid"),
            Name: unifiedgosdk.String("ornery whether"),
            Raw: &shared.PropertyAtsCandidateRaw{},
            Tags: []string{
                "Cadillac",
            },
            Telephones: []shared.AtsTelephone{
                shared.AtsTelephone{
                    Telephone: "Marketing Cotton",
                    Type: shared.AtsTelephoneTypeHome.ToPointer(),
                },
            },
            Title: unifiedgosdk.String("East"),
            UpdatedAt: types.MustTimeFromString("2023-10-31T11:53:36.953Z"),
        },
        ConnectionID: "redundant Tricycle unloose",
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsCandidate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.PatchAtsConnectionIDCandidateIDRequest](../../models/operations/patchatsconnectionidcandidateidrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


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
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Candidate.PostAtsConnectionIDCandidate(ctx, operations.PostAtsConnectionIDCandidateRequest{
        AtsCandidate: &shared.AtsCandidate{
            Address: &shared.PropertyAtsCandidateAddress{
                Address1: unifiedgosdk.String("incubate"),
                Address2: unifiedgosdk.String("azure Trans"),
                City: unifiedgosdk.String("Port Rory"),
                Country: unifiedgosdk.String("El Salvador"),
                CountryCode: unifiedgosdk.String("CX"),
                PostalCode: unifiedgosdk.String("54222-0235"),
                Region: unifiedgosdk.String("modi fooey"),
                RegionCode: unifiedgosdk.String("Metal TCP incidunt"),
            },
            CompanyName: unifiedgosdk.String("McCullough, Rosenbaum and Daugherty"),
            CreatedAt: types.MustTimeFromString("2023-02-07T05:55:59.357Z"),
            Emails: []shared.AtsEmail{
                shared.AtsEmail{
                    Email: "Eleanora.Rogahn44@hotmail.com",
                    Type: shared.AtsEmailTypeHome.ToPointer(),
                },
            },
            ExternalID: unifiedgosdk.String("South though"),
            ID: unifiedgosdk.String("<ID>"),
            ImageURL: unifiedgosdk.String("Pants"),
            Name: unifiedgosdk.String("Raleigh"),
            Raw: &shared.PropertyAtsCandidateRaw{},
            Tags: []string{
                "morph",
            },
            Telephones: []shared.AtsTelephone{
                shared.AtsTelephone{
                    Telephone: "lavender Sedan Folk",
                    Type: shared.AtsTelephoneTypeOther.ToPointer(),
                },
            },
            Title: unifiedgosdk.String("Savings panel"),
            UpdatedAt: types.MustTimeFromString("2022-02-09T15:32:35.578Z"),
        },
        ConnectionID: "Ngultrum red glean",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AtsCandidate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.PostAtsConnectionIDCandidateRequest](../../models/operations/postatsconnectionidcandidaterequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


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
	unifiedgosdk "github.com/unified-to/unified-go-sdk"
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
	"github.com/unified-to/unified-go-sdk/pkg/models/operations"
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

func main() {
    s := unifiedgosdk.New(
        unifiedgosdk.WithSecurity(shared.Security{
            Jwt: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Candidate.PutAtsConnectionIDCandidateID(ctx, operations.PutAtsConnectionIDCandidateIDRequest{
        AtsCandidate: &shared.AtsCandidate{
            Address: &shared.PropertyAtsCandidateAddress{
                Address1: unifiedgosdk.String("archive"),
                Address2: unifiedgosdk.String("Specialist Kyat"),
                City: unifiedgosdk.String("New Dennis"),
                Country: unifiedgosdk.String("Mauritius"),
                CountryCode: unifiedgosdk.String("TL"),
                PostalCode: unifiedgosdk.String("49105-9909"),
                Region: unifiedgosdk.String("copy olive"),
                RegionCode: unifiedgosdk.String("withdrawal cumque person"),
            },
            CompanyName: unifiedgosdk.String("Kuhn and Sons"),
            CreatedAt: types.MustTimeFromString("2022-01-28T10:51:00.922Z"),
            Emails: []shared.AtsEmail{
                shared.AtsEmail{
                    Email: "Hester.Jenkins@gmail.com",
                    Type: shared.AtsEmailTypeHome.ToPointer(),
                },
            },
            ExternalID: unifiedgosdk.String("Loan EXE"),
            ID: unifiedgosdk.String("<ID>"),
            ImageURL: unifiedgosdk.String("deliver executive RSS"),
            Name: unifiedgosdk.String("because aha black"),
            Raw: &shared.PropertyAtsCandidateRaw{},
            Tags: []string{
                "program",
            },
            Telephones: []shared.AtsTelephone{
                shared.AtsTelephone{
                    Telephone: "empower exit Pangender",
                    Type: shared.AtsTelephoneTypeMobile.ToPointer(),
                },
            },
            Title: unifiedgosdk.String("Corporate anenst Electronic"),
            UpdatedAt: types.MustTimeFromString("2022-03-30T08:00:53.284Z"),
        },
        ConnectionID: "Flerovium azure",
        ID: "<ID>",
    })
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
| `request`                                                                                                          | [operations.PutAtsConnectionIDCandidateIDRequest](../../models/operations/putatsconnectionidcandidateidrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.PutAtsConnectionIDCandidateIDResponse](../../models/operations/putatsconnectionidcandidateidresponse.md), error**

