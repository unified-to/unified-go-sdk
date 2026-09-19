# Hris

## Overview

### Available Operations

* [CreateHrisAttendance](#createhrisattendance) - Create an attendance
* [CreateHrisBankaccount](#createhrisbankaccount) - Create a bankaccount
* [CreateHrisBenefit](#createhrisbenefit) - Create a benefit
* [CreateHrisCompany](#createhriscompany) - Create a company
* [CreateHrisDeduction](#createhrisdeduction) - Create a deduction
* [CreateHrisDevice](#createhrisdevice) - Create a device
* [CreateHrisDocument](#createhrisdocument) - Create a document
* [CreateHrisEmployee](#createhrisemployee) - Create an employee
* [CreateHrisGroup](#createhrisgroup) - Create a group
* [CreateHrisLocation](#createhrislocation) - Create a location
* [CreateHrisTaxonomy](#createhristaxonomy) - Create a taxonomy
* [CreateHrisTimeoff](#createhristimeoff) - Create a timeoff
* [CreateHrisTimeshift](#createhristimeshift) - Create a timeshift
* [GetHrisAttendance](#gethrisattendance) - Retrieve an attendance
* [GetHrisBankaccount](#gethrisbankaccount) - Retrieve a bankaccount
* [GetHrisBenefit](#gethrisbenefit) - Retrieve a benefit
* [GetHrisCompany](#gethriscompany) - Retrieve a company
* [GetHrisDeduction](#gethrisdeduction) - Retrieve a deduction
* [GetHrisDevice](#gethrisdevice) - Retrieve a device
* [GetHrisDocument](#gethrisdocument) - Retrieve a document
* [GetHrisEmployee](#gethrisemployee) - Retrieve an employee
* [GetHrisGroup](#gethrisgroup) - Retrieve a group
* [GetHrisLocation](#gethrislocation) - Retrieve a location
* [GetHrisPayslip](#gethrispayslip) - Retrieve a payslip
* [GetHrisTaxonomy](#gethristaxonomy) - Retrieve a taxonomy
* [GetHrisTimeoff](#gethristimeoff) - Retrieve a timeoff
* [GetHrisTimeshift](#gethristimeshift) - Retrieve a timeshift
* [ListHrisAttendances](#listhrisattendances) - List all attendances
* [ListHrisBankaccounts](#listhrisbankaccounts) - List all bankaccounts
* [ListHrisBenefits](#listhrisbenefits) - List all benefits
* [ListHrisCompanies](#listhriscompanies) - List all companies
* [ListHrisDeductions](#listhrisdeductions) - List all deductions
* [ListHrisDevices](#listhrisdevices) - List all devices
* [ListHrisDocuments](#listhrisdocuments) - List all documents
* [ListHrisEmployees](#listhrisemployees) - List all employees
* [ListHrisGroups](#listhrisgroups) - List all groups
* [ListHrisLocations](#listhrislocations) - List all locations
* [ListHrisPayslips](#listhrispayslips) - List all payslips
* [ListHrisTaxonomies](#listhristaxonomies) - List all taxonomies
* [ListHrisTimeoffs](#listhristimeoffs) - List all timeoffs
* [ListHrisTimeshifts](#listhristimeshifts) - List all timeshifts
* [PatchHrisAttendance](#patchhrisattendance) - Update an attendance
* [PatchHrisBankaccount](#patchhrisbankaccount) - Update a bankaccount
* [PatchHrisBenefit](#patchhrisbenefit) - Update a benefit
* [PatchHrisCompany](#patchhriscompany) - Update a company
* [PatchHrisDeduction](#patchhrisdeduction) - Update a deduction
* [PatchHrisDevice](#patchhrisdevice) - Update a device
* [PatchHrisDocument](#patchhrisdocument) - Update a document
* [PatchHrisEmployee](#patchhrisemployee) - Update an employee
* [PatchHrisGroup](#patchhrisgroup) - Update a group
* [PatchHrisLocation](#patchhrislocation) - Update a location
* [PatchHrisTimeoff](#patchhristimeoff) - Update a timeoff
* [PatchHrisTimeshift](#patchhristimeshift) - Update a timeshift
* [RemoveHrisAttendance](#removehrisattendance) - Remove an attendance
* [RemoveHrisBankaccount](#removehrisbankaccount) - Remove a bankaccount
* [RemoveHrisBenefit](#removehrisbenefit) - Remove a benefit
* [RemoveHrisCompany](#removehriscompany) - Remove a company
* [RemoveHrisDeduction](#removehrisdeduction) - Remove a deduction
* [RemoveHrisDevice](#removehrisdevice) - Remove a device
* [RemoveHrisDocument](#removehrisdocument) - Remove a document
* [RemoveHrisEmployee](#removehrisemployee) - Remove an employee
* [RemoveHrisGroup](#removehrisgroup) - Remove a group
* [RemoveHrisLocation](#removehrislocation) - Remove a location
* [RemoveHrisTimeoff](#removehristimeoff) - Remove a timeoff
* [RemoveHrisTimeshift](#removehristimeshift) - Remove a timeshift
* [UpdateHrisAttendance](#updatehrisattendance) - Update an attendance
* [UpdateHrisBankaccount](#updatehrisbankaccount) - Update a bankaccount
* [UpdateHrisBenefit](#updatehrisbenefit) - Update a benefit
* [UpdateHrisCompany](#updatehriscompany) - Update a company
* [UpdateHrisDeduction](#updatehrisdeduction) - Update a deduction
* [UpdateHrisDevice](#updatehrisdevice) - Update a device
* [UpdateHrisDocument](#updatehrisdocument) - Update a document
* [UpdateHrisEmployee](#updatehrisemployee) - Update an employee
* [UpdateHrisGroup](#updatehrisgroup) - Update a group
* [UpdateHrisLocation](#updatehrislocation) - Update a location
* [UpdateHrisTimeoff](#updatehristimeoff) - Update a timeoff
* [UpdateHrisTimeshift](#updatehristimeshift) - Update a timeshift

## CreateHrisAttendance

Create an attendance

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisAttendance" method="post" path="/hris/{connection_id}/attendance" example="hris_attendance" -->
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

    res, err := s.Hris.CreateHrisAttendance(ctx, operations.CreateHrisAttendanceRequest{
        HrisAttendance: shared.HrisAttendance{
            Address: &shared.PropertyHrisAttendanceAddress{
                Address1: unifiedgosdk.Pointer("14108 Allie Flats"),
                City: unifiedgosdk.Pointer("Kearaborough"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("23844-2344"),
                Region: unifiedgosdk.Pointer("Tennessee"),
                RegionCode: unifiedgosdk.Pointer("CA"),
            },
            ApprovedAt: types.MustNewTimeFromString("2021-08-13T10:37:55.766Z"),
            Breaks: []shared.HrisAttendanceBreak{
                shared.HrisAttendanceBreak{
                    DurationMinutes: unifiedgosdk.Pointer[float64](12.0),
                    EndAt: types.MustNewTimeFromString("2023-10-23T02:00:25.579Z"),
                    ID: unifiedgosdk.Pointer("d60a1001-5a8a-4991-8c21-f4da6036cc87"),
                    IsPaid: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Lunch"),
                    StartAt: types.MustNewTimeFromString("2023-10-16T06:21:50.732Z"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2021-08-10T19:43:18.452Z"),
            Currency: unifiedgosdk.Pointer("UGX"),
            DeclaredTipsAmount: unifiedgosdk.Pointer[float64](161.0),
            EmployeeUserID: "<id>",
            EndAt: types.MustTimeFromString("2024-04-06T15:33:47.979Z"),
            HourlyRate: unifiedgosdk.Pointer[float64](53.0),
            Hours: unifiedgosdk.Pointer[float64](10.0),
            ID: unifiedgosdk.Pointer("fec3216d-df3b-4e51-90f6-dedb81e89342"),
            JobName: unifiedgosdk.Pointer("Global Creative Supervisor"),
            NonCashTipsAmount: unifiedgosdk.Pointer[float64](54.0),
            StartAt: types.MustTimeFromString("2021-11-09T11:31:11.587Z"),
            Status: shared.HrisAttendanceStatusClosed.ToPointer(),
            Timezone: unifiedgosdk.Pointer("America/Atikokan"),
            UpdatedAt: types.MustNewTimeFromString("2022-01-17T03:19:38.951Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisAttendance != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateHrisAttendanceRequest](../../pkg/models/operations/createhrisattendancerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.CreateHrisAttendanceResponse](../../pkg/models/operations/createhrisattendanceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateHrisBankaccount

Create a bankaccount

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisBankaccount" method="post" path="/hris/{connection_id}/bankaccount" example="hris_bankaccount" -->
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

    res, err := s.Hris.CreateHrisBankaccount(ctx, operations.CreateHrisBankaccountRequest{
        HrisBankaccount: shared.HrisBankaccount{
            AccountNumber: unifiedgosdk.Pointer("****3777"),
            AccountNumberLast4: unifiedgosdk.Pointer("3777"),
            AccountType: shared.HrisBankaccountAccountTypeChecking.ToPointer(),
            BankName: unifiedgosdk.Pointer("Huel Group"),
            CreatedAt: types.MustNewTimeFromString("2019-11-16T16:43:45.976Z"),
            ID: unifiedgosdk.Pointer("a3516f34-9336-4e6c-89b4-b15a52a387ec"),
            IsPrimary: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("Checking Account"),
            RoutingNumber: unifiedgosdk.Pointer("448650724"),
            UpdatedAt: types.MustNewTimeFromString("2025-06-05T05:39:50.306Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisBankaccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CreateHrisBankaccountRequest](../../pkg/models/operations/createhrisbankaccountrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.CreateHrisBankaccountResponse](../../pkg/models/operations/createhrisbankaccountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateHrisBenefit

Create a benefit

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisBenefit" method="post" path="/hris/{connection_id}/benefit" example="hris_benefit" -->
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

    res, err := s.Hris.CreateHrisBenefit(ctx, operations.CreateHrisBenefitRequest{
        HrisBenefit: shared.HrisBenefit{
            CoverageLevel: shared.CoverageLevelEmployeeSpouse.ToPointer(),
            CreatedAt: types.MustNewTimeFromString("2020-06-11T01:24:05.654Z"),
            Currency: unifiedgosdk.Pointer("JOD"),
            Description: unifiedgosdk.Pointer("Vomito voluptas dolor sed."),
            EmployerContributionAmount: unifiedgosdk.Pointer[float64](185006.0),
            EmployerContributionMaxAmount: unifiedgosdk.Pointer[float64](179093.0),
            EmployerContributionType: shared.EmployerContributionTypePercentage.ToPointer(),
            Frequency: shared.HrisBenefitFrequencyHour.ToPointer(),
            ID: unifiedgosdk.Pointer("8b32b37f-4d01-4664-a74f-c47ac3822568"),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("Frozen Wooden Ball"),
            Tax: shared.TaxPreTax.ToPointer(),
            Type: shared.HrisBenefitTypeGarnishment.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-03-06T20:45:33.631Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisBenefit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateHrisBenefitRequest](../../pkg/models/operations/createhrisbenefitrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateHrisBenefitResponse](../../pkg/models/operations/createhrisbenefitresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateHrisCompany

Create a company

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisCompany" method="post" path="/hris/{connection_id}/company" example="hris_company" -->
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

    res, err := s.Hris.CreateHrisCompany(ctx, operations.CreateHrisCompanyRequest{
        HrisCompany: shared.HrisCompany{
            Address: &shared.PropertyHrisCompanyAddress{
                Address1: unifiedgosdk.Pointer("2549 Church Walk"),
                City: unifiedgosdk.Pointer("Lake Nettiebury"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("32877-4898"),
                Region: unifiedgosdk.Pointer("Idaho"),
                RegionCode: unifiedgosdk.Pointer("PA"),
            },
            CreatedAt: types.MustNewTimeFromString("2021-05-02T22:27:38.970Z"),
            ID: unifiedgosdk.Pointer("cce02aac-1494-4c07-b74a-3bcaf14ab2d7"),
            LegalName: unifiedgosdk.Pointer("Schultz LLC"),
            Name: unifiedgosdk.Pointer("Gottlieb Group"),
            UpdatedAt: types.MustNewTimeFromString("2026-09-06T19:24:54.141Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisCompany != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateHrisCompanyRequest](../../pkg/models/operations/createhriscompanyrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateHrisCompanyResponse](../../pkg/models/operations/createhriscompanyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateHrisDeduction

Create a deduction

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisDeduction" method="post" path="/hris/{connection_id}/deduction" example="hris_deduction" -->
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

    res, err := s.Hris.CreateHrisDeduction(ctx, operations.CreateHrisDeductionRequest{
        HrisDeduction: shared.HrisDeduction{
            Amount: unifiedgosdk.Pointer[float64](139655.0),
            CoverageLevel: shared.HrisDeductionCoverageLevelEmployeeOnly.ToPointer(),
            CreatedAt: types.MustNewTimeFromString("2020-02-05T01:46:31.384Z"),
            EndAt: types.MustNewTimeFromString("2026-05-23T20:08:22.523Z"),
            Frequency: shared.HrisDeductionFrequencyMonth.ToPointer(),
            ID: unifiedgosdk.Pointer("b52562c5-6d0c-435f-a91a-f74dc660c77f"),
            IsActive: unifiedgosdk.Pointer(false),
            Notes: unifiedgosdk.Pointer("Carmen desidero."),
            StartAt: types.MustNewTimeFromString("2025-02-18T21:39:35.472Z"),
            Type: shared.HrisDeductionTypeFixed.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-03-02T13:27:26.612Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDeduction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreateHrisDeductionRequest](../../pkg/models/operations/createhrisdeductionrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreateHrisDeductionResponse](../../pkg/models/operations/createhrisdeductionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateHrisDevice

Create a device

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisDevice" method="post" path="/hris/{connection_id}/device" example="hris_device" -->
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

    res, err := s.Hris.CreateHrisDevice(ctx, operations.CreateHrisDeviceRequest{
        HrisDevice: shared.HrisDevice{
            AdminUserIds: []string{},
            AssetTag: unifiedgosdk.Pointer("dpho9OuFNG"),
            CreatedAt: types.MustNewTimeFromString("2019-04-04T17:11:40.322Z"),
            HasAntivirus: unifiedgosdk.Pointer(false),
            HasFirewall: unifiedgosdk.Pointer(true),
            HasHdEncrypted: unifiedgosdk.Pointer(true),
            HasPasswordManager: unifiedgosdk.Pointer(true),
            HasScreenlock: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("af3e715d-70e2-46d2-b996-0d44152db42c"),
            IsMissing: unifiedgosdk.Pointer(false),
            Manufacturer: unifiedgosdk.Pointer("Sanford - Hamill"),
            Model: unifiedgosdk.Pointer("Refined"),
            Name: unifiedgosdk.Pointer("cross_contamination_if.rar"),
            Os: unifiedgosdk.Pointer("monitor"),
            OsVersion: unifiedgosdk.Pointer("1.12.16"),
            UpdatedAt: types.MustNewTimeFromString("2023-05-21T13:19:00.267Z"),
            Version: unifiedgosdk.Pointer("2.20.17"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDevice != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateHrisDeviceRequest](../../pkg/models/operations/createhrisdevicerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateHrisDeviceResponse](../../pkg/models/operations/createhrisdeviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateHrisDocument

Create a document

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisDocument" method="post" path="/hris/{connection_id}/document" example="hris_document" -->
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

    res, err := s.Hris.CreateHrisDocument(ctx, operations.CreateHrisDocumentRequest{
        HrisDocument: shared.HrisDocument{
            CreatedAt: types.MustNewTimeFromString("2022-10-27T11:47:26.086Z"),
            DocumentURL: unifiedgosdk.Pointer("https://sore-decision.biz/"),
            Filename: unifiedgosdk.Pointer("ridge_forager.xsl"),
            ID: unifiedgosdk.Pointer("80620b12-53a2-48ee-ab80-ea0693705467"),
            Type: shared.HrisDocumentTypePolicy.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-09-17T18:31:17.139Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDocument != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateHrisDocumentRequest](../../pkg/models/operations/createhrisdocumentrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateHrisDocumentResponse](../../pkg/models/operations/createhrisdocumentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateHrisEmployee

Create an employee

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisEmployee" method="post" path="/hris/{connection_id}/employee" example="hris_employee" -->
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

    res, err := s.Hris.CreateHrisEmployee(ctx, operations.CreateHrisEmployeeRequest{
        HrisEmployee: shared.HrisEmployee{
            Address: &shared.PropertyHrisEmployeeAddress{
                Address1: unifiedgosdk.Pointer("52008 Lansdowne Road"),
                Address2: unifiedgosdk.Pointer("Apt. 101"),
                City: unifiedgosdk.Pointer("Connellyberg"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("18978"),
                Region: unifiedgosdk.Pointer("South Dakota"),
                RegionCode: unifiedgosdk.Pointer("NM"),
            },
            Bio: unifiedgosdk.Pointer("sushi devotee, singer"),
            Compensation: []shared.HrisCompensation{
                shared.HrisCompensation{
                    Amount: unifiedgosdk.Pointer[float64](69148.0),
                    Currency: unifiedgosdk.Pointer("CRC"),
                    Frequency: shared.HrisCompensationFrequencyQuarter.ToPointer(),
                    Notes: unifiedgosdk.Pointer("Tergeo laborum laboriosam tutis."),
                    Type: shared.HrisCompensationTypeEquity.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2019-09-16T15:08:53.262Z"),
            Currency: unifiedgosdk.Pointer("IDR"),
            DateOfBirth: types.MustNewTimeFromString("2001-04-22"),
            Emails: []shared.HrisEmail{
                shared.HrisEmail{
                    Email: "Zetta_Prohaska67@hotmail.com",
                    Type: shared.HrisEmailTypeHome.ToPointer(),
                },
            },
            EmployeeNumber: unifiedgosdk.Pointer("YuOt169CGu"),
            EmploymentStatus: shared.EmploymentStatusActive.ToPointer(),
            EmploymentType: shared.HrisEmployeeEmploymentTypeVolunteer.ToPointer(),
            FirstName: unifiedgosdk.Pointer("Zetta"),
            Gender: shared.HrisEmployeeGenderIntersex.ToPointer(),
            HasMfa: unifiedgosdk.Pointer(true),
            HiredAt: types.MustNewTimeFromString("2023-05-11T03:53:15.152Z"),
            ID: unifiedgosdk.Pointer("8b185b20-7dbe-48ef-a8c8-c83140a90eed"),
            ImageURL: unifiedgosdk.Pointer("https://loremflickr.com/3684/2116?lock=4686991638584456"),
            LanguageLocale: unifiedgosdk.Pointer("es"),
            LastName: unifiedgosdk.Pointer("Prohaska"),
            Locations: []shared.HrisLocation{},
            MaritalStatus: shared.MaritalStatusMarried.ToPointer(),
            Metadata: []shared.HrisMetadata{
                shared.HrisMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateHrisMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.HrisMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("6b50dd19-cdec-448f-8ce2-1e98da53c872"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateHrisMetadataValueStr(
                        "tenetur",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Zetta Prohaska"),
            Pronouns: unifiedgosdk.Pointer("she/her"),
            Relationships: []shared.HrisEmployeerelationship{
                shared.HrisEmployeerelationship{
                    Emails: []shared.HrisEmail{
                        shared.HrisEmail{
                            Email: "Deshaun.Sanford24@yahoo.com",
                        },
                        shared.HrisEmail{
                            Email: "Rebeca.Dibbert11@hotmail.com",
                        },
                        shared.HrisEmail{
                            Email: "Hester80@gmail.com",
                        },
                    },
                    Name: unifiedgosdk.Pointer("Automotive"),
                    Type: shared.HrisEmployeerelationshipTypeEmergency.ToPointer(),
                },
                shared.HrisEmployeerelationship{
                    Emails: []shared.HrisEmail{
                        shared.HrisEmail{
                            Email: "Benedict_Wisozk83@hotmail.com",
                        },
                        shared.HrisEmail{
                            Email: "Princess_Rath43@gmail.com",
                        },
                        shared.HrisEmail{
                            Email: "Elmira92@yahoo.com",
                        },
                    },
                    Name: unifiedgosdk.Pointer("Music"),
                    Type: shared.HrisEmployeerelationshipTypeFriend.ToPointer(),
                },
                shared.HrisEmployeerelationship{
                    Emails: []shared.HrisEmail{
                        shared.HrisEmail{
                            Email: "Jane30@gmail.com",
                        },
                    },
                    Name: unifiedgosdk.Pointer("Jewelry"),
                    Type: shared.HrisEmployeerelationshipTypeSibling.ToPointer(),
                },
            },
            Salutation: unifiedgosdk.Pointer("Miss"),
            SsnSin: unifiedgosdk.Pointer("yMRtj0Q3xO"),
            StorageQuotaAllocated: unifiedgosdk.Pointer[float64](3674489.0),
            StorageQuotaAvailable: unifiedgosdk.Pointer[float64](7748057.0),
            StorageQuotaUsed: unifiedgosdk.Pointer[float64](301727.0),
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "(409) 801-3705",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
            },
            TerminationReason: unifiedgosdk.Pointer("Communis adnuo damnatio atavus terebro acies canis cogito triumphus creber temptatio defendo cubo amissio paulatim corroboro."),
            TimeoffDaysTotal: unifiedgosdk.Pointer[float64](12.0),
            TimeoffDaysUsed: unifiedgosdk.Pointer[float64](6.0),
            Timezone: unifiedgosdk.Pointer("Africa/Harare"),
            Title: unifiedgosdk.Pointer("Investor Paradigm Liaison"),
            UpdatedAt: types.MustNewTimeFromString("2022-02-19T14:46:57.314Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisEmployee != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateHrisEmployeeRequest](../../pkg/models/operations/createhrisemployeerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateHrisEmployeeResponse](../../pkg/models/operations/createhrisemployeeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateHrisGroup

Create a group

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisGroup" method="post" path="/hris/{connection_id}/group" example="hris_group" -->
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

    res, err := s.Hris.CreateHrisGroup(ctx, operations.CreateHrisGroupRequest{
        HrisGroup: shared.HrisGroup{
            CreatedAt: types.MustNewTimeFromString("2023-11-01T13:13:40.714Z"),
            Description: unifiedgosdk.Pointer("Absorbeo casso."),
            ID: unifiedgosdk.Pointer("e67b8f74-b470-430b-964e-b95e6b690f67"),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("Games"),
            Type: shared.HrisGroupTypeBusinessUnit.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-04-24T09:38:41.461Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateHrisGroupRequest](../../pkg/models/operations/createhrisgrouprequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreateHrisGroupResponse](../../pkg/models/operations/createhrisgroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateHrisLocation

Create a location

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisLocation" method="post" path="/hris/{connection_id}/location" example="hris_location" -->
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

    res, err := s.Hris.CreateHrisLocation(ctx, operations.CreateHrisLocationRequest{
        HrisLocation: shared.HrisLocation{
            Address: &shared.PropertyHrisLocationAddress{
                Address1: unifiedgosdk.Pointer("2743 Connelly Summit"),
                Address2: unifiedgosdk.Pointer("Apt. 350"),
                City: unifiedgosdk.Pointer("Titusville"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("16154-1095"),
                Region: unifiedgosdk.Pointer("Oregon"),
                RegionCode: unifiedgosdk.Pointer("AL"),
            },
            CreatedAt: types.MustNewTimeFromString("2021-07-18T10:32:01.414Z"),
            Currency: unifiedgosdk.Pointer("MUR"),
            Description: unifiedgosdk.Pointer("Acervus caries."),
            ExternalIdentifier: unifiedgosdk.Pointer("f453c41f-5f25-44cc-8cdc-8e33a5b31a9c"),
            ID: unifiedgosdk.Pointer("6c11765c-9349-4d92-a5b0-5943f9cb851b"),
            IsActive: unifiedgosdk.Pointer(true),
            IsHq: unifiedgosdk.Pointer(false),
            LanguageLocale: unifiedgosdk.Pointer("fr"),
            Name: unifiedgosdk.Pointer("adhuc"),
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "(710) 550-6997",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
                shared.HrisTelephone{
                    Telephone: "(208) 555-8542",
                    Type: shared.HrisTelephoneTypeHome.ToPointer(),
                },
                shared.HrisTelephone{
                    Telephone: "(712) 473-5482",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
            },
            Timezone: unifiedgosdk.Pointer("America/Guyana"),
            UpdatedAt: types.MustNewTimeFromString("2023-06-09T08:50:56.332Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisLocation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateHrisLocationRequest](../../pkg/models/operations/createhrislocationrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateHrisLocationResponse](../../pkg/models/operations/createhrislocationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateHrisTaxonomy

Create a taxonomy

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisTaxonomy" method="post" path="/hris/{connection_id}/taxonomy" example="hris_taxonomy" -->
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

    res, err := s.Hris.CreateHrisTaxonomy(ctx, operations.CreateHrisTaxonomyRequest{
        HrisTaxonomy: shared.HrisTaxonomy{
            CreatedAt: types.MustNewTimeFromString("2022-06-23T02:10:00.789Z"),
            Description: unifiedgosdk.Pointer("Apto demonstro audacia adstringo cursim tristis solio careo."),
            Domain: unifiedgosdk.Pointer("Electronics"),
            ID: unifiedgosdk.Pointer("ede085db-5709-4d53-a490-746f3de5be17"),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("International Functionality Architect"),
            ParentID: unifiedgosdk.Pointer("6524b2a7-6520-4e15-8c4e-1aa6793db837"),
            RoleIds: []string{
                "2b1ef757-eb4c-4207-8af1-929afe49cd65",
            },
            Subcategory: unifiedgosdk.Pointer("Bamboo"),
            Type: shared.HrisTaxonomyTypeKnowledge.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-05-22T09:46:16.697Z"),
            URL: unifiedgosdk.Pointer("https://our-polarisation.name"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTaxonomy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateHrisTaxonomyRequest](../../pkg/models/operations/createhristaxonomyrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateHrisTaxonomyResponse](../../pkg/models/operations/createhristaxonomyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateHrisTimeoff

Create a timeoff

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisTimeoff" method="post" path="/hris/{connection_id}/timeoff" example="hris_timeoff" -->
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

    res, err := s.Hris.CreateHrisTimeoff(ctx, operations.CreateHrisTimeoffRequest{
        HrisTimeoff: shared.HrisTimeoff{
            ApprovedAt: types.MustNewTimeFromString("2022-02-20T22:44:28.294Z"),
            Comments: unifiedgosdk.Pointer("Blandior ventus curiositas amplitudo."),
            CreatedAt: types.MustNewTimeFromString("2021-10-06T18:00:20.615Z"),
            Duration: unifiedgosdk.Pointer[float64](4.0),
            DurationType: shared.DurationTypeDay.ToPointer(),
            EndAt: types.MustNewTimeFromString("2024-12-08T04:10:50.220Z"),
            ID: unifiedgosdk.Pointer("a588fadd-8d1a-4a6e-8248-9047b712f880"),
            IsPaid: unifiedgosdk.Pointer(true),
            OriginalType: unifiedgosdk.Pointer("acerbitas ut"),
            Reason: unifiedgosdk.Pointer("verto"),
            StartAt: types.MustNewTimeFromString("2023-08-23T15:25:53.572Z"),
            Status: shared.HrisTimeoffStatusDenied.ToPointer(),
            Type: shared.HrisTimeoffTypeInLieu.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-07-07T22:58:27.090Z"),
            UserID: "<id>",
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeoff != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateHrisTimeoffRequest](../../pkg/models/operations/createhristimeoffrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateHrisTimeoffResponse](../../pkg/models/operations/createhristimeoffresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateHrisTimeshift

Create a timeshift

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisTimeshift" method="post" path="/hris/{connection_id}/timeshift" example="hris_timeshift" -->
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

    res, err := s.Hris.CreateHrisTimeshift(ctx, operations.CreateHrisTimeshiftRequest{
        HrisTimeshift: shared.HrisTimeshift{
            ApprovedAt: types.MustNewTimeFromString("2023-06-06T02:42:24.608Z"),
            Compensation: []shared.HrisCompensation{
                shared.HrisCompensation{
                    Amount: unifiedgosdk.Pointer[float64](76761.0),
                    Currency: unifiedgosdk.Pointer("JPY"),
                    Frequency: shared.HrisCompensationFrequencyHour.ToPointer(),
                    Notes: unifiedgosdk.Pointer("Annus adficio suasoria architecto aggero."),
                    Type: shared.HrisCompensationTypeOther.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2019-07-01T23:53:15.738Z"),
            EmployeeUserID: "<id>",
            EndAt: types.MustNewTimeFromString("2026-08-26T06:38:23.084Z"),
            Hours: unifiedgosdk.Pointer[float64](8.0),
            ID: unifiedgosdk.Pointer("5c67d6d3-4511-4afd-a007-7f1393aecc5d"),
            IsApproved: unifiedgosdk.Pointer(true),
            StartAt: types.MustNewTimeFromString("2023-06-25T07:50:01.852Z"),
            UpdatedAt: types.MustNewTimeFromString("2021-06-23T03:00:57.047Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeshift != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreateHrisTimeshiftRequest](../../pkg/models/operations/createhristimeshiftrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreateHrisTimeshiftResponse](../../pkg/models/operations/createhristimeshiftresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisAttendance

Retrieve an attendance

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisAttendance" method="get" path="/hris/{connection_id}/attendance/{id}" -->
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

    res, err := s.Hris.GetHrisAttendance(ctx, operations.GetHrisAttendanceRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisAttendance != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetHrisAttendanceRequest](../../pkg/models/operations/gethrisattendancerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetHrisAttendanceResponse](../../pkg/models/operations/gethrisattendanceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisBankaccount

Retrieve a bankaccount

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisBankaccount" method="get" path="/hris/{connection_id}/bankaccount/{id}" -->
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

    res, err := s.Hris.GetHrisBankaccount(ctx, operations.GetHrisBankaccountRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisBankaccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetHrisBankaccountRequest](../../pkg/models/operations/gethrisbankaccountrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetHrisBankaccountResponse](../../pkg/models/operations/gethrisbankaccountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisBenefit

Retrieve a benefit

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisBenefit" method="get" path="/hris/{connection_id}/benefit/{id}" -->
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

    res, err := s.Hris.GetHrisBenefit(ctx, operations.GetHrisBenefitRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisBenefit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetHrisBenefitRequest](../../pkg/models/operations/gethrisbenefitrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetHrisBenefitResponse](../../pkg/models/operations/gethrisbenefitresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisCompany

Retrieve a company

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisCompany" method="get" path="/hris/{connection_id}/company/{id}" -->
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

    res, err := s.Hris.GetHrisCompany(ctx, operations.GetHrisCompanyRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisCompany != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetHrisCompanyRequest](../../pkg/models/operations/gethriscompanyrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetHrisCompanyResponse](../../pkg/models/operations/gethriscompanyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisDeduction

Retrieve a deduction

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisDeduction" method="get" path="/hris/{connection_id}/deduction/{id}" -->
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

    res, err := s.Hris.GetHrisDeduction(ctx, operations.GetHrisDeductionRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDeduction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetHrisDeductionRequest](../../pkg/models/operations/gethrisdeductionrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetHrisDeductionResponse](../../pkg/models/operations/gethrisdeductionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisDevice

Retrieve a device

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisDevice" method="get" path="/hris/{connection_id}/device/{id}" -->
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

    res, err := s.Hris.GetHrisDevice(ctx, operations.GetHrisDeviceRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDevice != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetHrisDeviceRequest](../../pkg/models/operations/gethrisdevicerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetHrisDeviceResponse](../../pkg/models/operations/gethrisdeviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisDocument

Retrieve a document

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisDocument" method="get" path="/hris/{connection_id}/document/{id}" -->
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

    res, err := s.Hris.GetHrisDocument(ctx, operations.GetHrisDocumentRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDocument != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetHrisDocumentRequest](../../pkg/models/operations/gethrisdocumentrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetHrisDocumentResponse](../../pkg/models/operations/gethrisdocumentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisEmployee

Retrieve an employee

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisEmployee" method="get" path="/hris/{connection_id}/employee/{id}" -->
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

    res, err := s.Hris.GetHrisEmployee(ctx, operations.GetHrisEmployeeRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisEmployee != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetHrisEmployeeRequest](../../pkg/models/operations/gethrisemployeerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetHrisEmployeeResponse](../../pkg/models/operations/gethrisemployeeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisGroup

Retrieve a group

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisGroup" method="get" path="/hris/{connection_id}/group/{id}" -->
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

    res, err := s.Hris.GetHrisGroup(ctx, operations.GetHrisGroupRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetHrisGroupRequest](../../pkg/models/operations/gethrisgrouprequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.GetHrisGroupResponse](../../pkg/models/operations/gethrisgroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisLocation

Retrieve a location

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisLocation" method="get" path="/hris/{connection_id}/location/{id}" -->
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

    res, err := s.Hris.GetHrisLocation(ctx, operations.GetHrisLocationRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisLocation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetHrisLocationRequest](../../pkg/models/operations/gethrislocationrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetHrisLocationResponse](../../pkg/models/operations/gethrislocationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisPayslip

Retrieve a payslip

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisPayslip" method="get" path="/hris/{connection_id}/payslip/{id}" -->
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

    res, err := s.Hris.GetHrisPayslip(ctx, operations.GetHrisPayslipRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisPayslip != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetHrisPayslipRequest](../../pkg/models/operations/gethrispaysliprequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetHrisPayslipResponse](../../pkg/models/operations/gethrispayslipresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisTaxonomy

Retrieve a taxonomy

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisTaxonomy" method="get" path="/hris/{connection_id}/taxonomy/{id}" -->
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

    res, err := s.Hris.GetHrisTaxonomy(ctx, operations.GetHrisTaxonomyRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTaxonomy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetHrisTaxonomyRequest](../../pkg/models/operations/gethristaxonomyrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetHrisTaxonomyResponse](../../pkg/models/operations/gethristaxonomyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisTimeoff

Retrieve a timeoff

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisTimeoff" method="get" path="/hris/{connection_id}/timeoff/{id}" -->
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

    res, err := s.Hris.GetHrisTimeoff(ctx, operations.GetHrisTimeoffRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeoff != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetHrisTimeoffRequest](../../pkg/models/operations/gethristimeoffrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetHrisTimeoffResponse](../../pkg/models/operations/gethristimeoffresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisTimeshift

Retrieve a timeshift

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisTimeshift" method="get" path="/hris/{connection_id}/timeshift/{id}" -->
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

    res, err := s.Hris.GetHrisTimeshift(ctx, operations.GetHrisTimeshiftRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeshift != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetHrisTimeshiftRequest](../../pkg/models/operations/gethristimeshiftrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetHrisTimeshiftResponse](../../pkg/models/operations/gethristimeshiftresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisAttendances

List all attendances

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisAttendances" method="get" path="/hris/{connection_id}/attendance" -->
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

    res, err := s.Hris.ListHrisAttendances(ctx, operations.ListHrisAttendancesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisAttendances != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListHrisAttendancesRequest](../../pkg/models/operations/listhrisattendancesrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.ListHrisAttendancesResponse](../../pkg/models/operations/listhrisattendancesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisBankaccounts

List all bankaccounts

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisBankaccounts" method="get" path="/hris/{connection_id}/bankaccount" -->
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

    res, err := s.Hris.ListHrisBankaccounts(ctx, operations.ListHrisBankaccountsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisBankaccounts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListHrisBankaccountsRequest](../../pkg/models/operations/listhrisbankaccountsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListHrisBankaccountsResponse](../../pkg/models/operations/listhrisbankaccountsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisBenefits

List all benefits

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisBenefits" method="get" path="/hris/{connection_id}/benefit" -->
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

    res, err := s.Hris.ListHrisBenefits(ctx, operations.ListHrisBenefitsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisBenefits != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListHrisBenefitsRequest](../../pkg/models/operations/listhrisbenefitsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListHrisBenefitsResponse](../../pkg/models/operations/listhrisbenefitsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisCompanies

List all companies

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisCompanies" method="get" path="/hris/{connection_id}/company" -->
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

    res, err := s.Hris.ListHrisCompanies(ctx, operations.ListHrisCompaniesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisCompanies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListHrisCompaniesRequest](../../pkg/models/operations/listhriscompaniesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListHrisCompaniesResponse](../../pkg/models/operations/listhriscompaniesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisDeductions

List all deductions

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisDeductions" method="get" path="/hris/{connection_id}/deduction" -->
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

    res, err := s.Hris.ListHrisDeductions(ctx, operations.ListHrisDeductionsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDeductions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListHrisDeductionsRequest](../../pkg/models/operations/listhrisdeductionsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListHrisDeductionsResponse](../../pkg/models/operations/listhrisdeductionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisDevices

List all devices

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisDevices" method="get" path="/hris/{connection_id}/device" -->
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

    res, err := s.Hris.ListHrisDevices(ctx, operations.ListHrisDevicesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDevices != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListHrisDevicesRequest](../../pkg/models/operations/listhrisdevicesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListHrisDevicesResponse](../../pkg/models/operations/listhrisdevicesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisDocuments

List all documents

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisDocuments" method="get" path="/hris/{connection_id}/document" -->
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

    res, err := s.Hris.ListHrisDocuments(ctx, operations.ListHrisDocumentsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDocuments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListHrisDocumentsRequest](../../pkg/models/operations/listhrisdocumentsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListHrisDocumentsResponse](../../pkg/models/operations/listhrisdocumentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisEmployees

List all employees

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisEmployees" method="get" path="/hris/{connection_id}/employee" -->
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

    res, err := s.Hris.ListHrisEmployees(ctx, operations.ListHrisEmployeesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisEmployees != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListHrisEmployeesRequest](../../pkg/models/operations/listhrisemployeesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListHrisEmployeesResponse](../../pkg/models/operations/listhrisemployeesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisGroups

List all groups

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisGroups" method="get" path="/hris/{connection_id}/group" -->
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

    res, err := s.Hris.ListHrisGroups(ctx, operations.ListHrisGroupsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisGroups != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListHrisGroupsRequest](../../pkg/models/operations/listhrisgroupsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ListHrisGroupsResponse](../../pkg/models/operations/listhrisgroupsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisLocations

List all locations

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisLocations" method="get" path="/hris/{connection_id}/location" -->
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

    res, err := s.Hris.ListHrisLocations(ctx, operations.ListHrisLocationsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisLocations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListHrisLocationsRequest](../../pkg/models/operations/listhrislocationsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListHrisLocationsResponse](../../pkg/models/operations/listhrislocationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisPayslips

List all payslips

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisPayslips" method="get" path="/hris/{connection_id}/payslip" -->
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

    res, err := s.Hris.ListHrisPayslips(ctx, operations.ListHrisPayslipsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisPayslips != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListHrisPayslipsRequest](../../pkg/models/operations/listhrispayslipsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListHrisPayslipsResponse](../../pkg/models/operations/listhrispayslipsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisTaxonomies

List all taxonomies

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisTaxonomies" method="get" path="/hris/{connection_id}/taxonomy" -->
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

    res, err := s.Hris.ListHrisTaxonomies(ctx, operations.ListHrisTaxonomiesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTaxonomies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListHrisTaxonomiesRequest](../../pkg/models/operations/listhristaxonomiesrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListHrisTaxonomiesResponse](../../pkg/models/operations/listhristaxonomiesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisTimeoffs

List all timeoffs

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisTimeoffs" method="get" path="/hris/{connection_id}/timeoff" -->
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

    res, err := s.Hris.ListHrisTimeoffs(ctx, operations.ListHrisTimeoffsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeoffs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListHrisTimeoffsRequest](../../pkg/models/operations/listhristimeoffsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListHrisTimeoffsResponse](../../pkg/models/operations/listhristimeoffsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisTimeshifts

List all timeshifts

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisTimeshifts" method="get" path="/hris/{connection_id}/timeshift" -->
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

    res, err := s.Hris.ListHrisTimeshifts(ctx, operations.ListHrisTimeshiftsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeshifts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListHrisTimeshiftsRequest](../../pkg/models/operations/listhristimeshiftsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListHrisTimeshiftsResponse](../../pkg/models/operations/listhristimeshiftsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisAttendance

Update an attendance

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisAttendance" method="patch" path="/hris/{connection_id}/attendance/{id}" example="hris_attendance" -->
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

    res, err := s.Hris.PatchHrisAttendance(ctx, operations.PatchHrisAttendanceRequest{
        HrisAttendance: shared.HrisAttendance{
            Address: &shared.PropertyHrisAttendanceAddress{
                Address1: unifiedgosdk.Pointer("14108 Allie Flats"),
                City: unifiedgosdk.Pointer("Kearaborough"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("23844-2344"),
                Region: unifiedgosdk.Pointer("Tennessee"),
                RegionCode: unifiedgosdk.Pointer("CA"),
            },
            ApprovedAt: types.MustNewTimeFromString("2021-08-13T10:37:55.766Z"),
            Breaks: []shared.HrisAttendanceBreak{
                shared.HrisAttendanceBreak{
                    DurationMinutes: unifiedgosdk.Pointer[float64](12.0),
                    EndAt: types.MustNewTimeFromString("2023-10-23T02:00:25.590Z"),
                    ID: unifiedgosdk.Pointer("d60a1001-5a8a-4991-8c21-f4da6036cc87"),
                    IsPaid: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Lunch"),
                    StartAt: types.MustNewTimeFromString("2023-10-16T06:21:50.743Z"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2021-08-10T19:43:18.452Z"),
            Currency: unifiedgosdk.Pointer("UGX"),
            DeclaredTipsAmount: unifiedgosdk.Pointer[float64](161.0),
            EmployeeUserID: "<id>",
            EndAt: types.MustTimeFromString("2024-04-06T15:33:47.993Z"),
            HourlyRate: unifiedgosdk.Pointer[float64](53.0),
            Hours: unifiedgosdk.Pointer[float64](10.0),
            ID: unifiedgosdk.Pointer("b2360f22-cde1-41a5-9cea-b76990466cf3"),
            JobName: unifiedgosdk.Pointer("Global Creative Supervisor"),
            NonCashTipsAmount: unifiedgosdk.Pointer[float64](54.0),
            StartAt: types.MustTimeFromString("2021-11-09T11:31:11.588Z"),
            Status: shared.HrisAttendanceStatusClosed.ToPointer(),
            Timezone: unifiedgosdk.Pointer("America/Atikokan"),
            UpdatedAt: types.MustNewTimeFromString("2022-01-17T03:19:38.954Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisAttendance != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PatchHrisAttendanceRequest](../../pkg/models/operations/patchhrisattendancerequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.PatchHrisAttendanceResponse](../../pkg/models/operations/patchhrisattendanceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisBankaccount

Update a bankaccount

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisBankaccount" method="patch" path="/hris/{connection_id}/bankaccount/{id}" example="hris_bankaccount" -->
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

    res, err := s.Hris.PatchHrisBankaccount(ctx, operations.PatchHrisBankaccountRequest{
        HrisBankaccount: shared.HrisBankaccount{
            AccountNumber: unifiedgosdk.Pointer("****3777"),
            AccountNumberLast4: unifiedgosdk.Pointer("3777"),
            AccountType: shared.HrisBankaccountAccountTypeChecking.ToPointer(),
            BankName: unifiedgosdk.Pointer("Huel Group"),
            CreatedAt: types.MustNewTimeFromString("2019-11-16T16:43:45.976Z"),
            ID: unifiedgosdk.Pointer("ee0b7b15-c45b-4cad-ba3d-ae501748d719"),
            IsPrimary: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("Checking Account"),
            RoutingNumber: unifiedgosdk.Pointer("448650724"),
            UpdatedAt: types.MustNewTimeFromString("2025-06-05T05:39:50.317Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisBankaccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PatchHrisBankaccountRequest](../../pkg/models/operations/patchhrisbankaccountrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.PatchHrisBankaccountResponse](../../pkg/models/operations/patchhrisbankaccountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisBenefit

Update a benefit

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisBenefit" method="patch" path="/hris/{connection_id}/benefit/{id}" example="hris_benefit" -->
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

    res, err := s.Hris.PatchHrisBenefit(ctx, operations.PatchHrisBenefitRequest{
        HrisBenefit: shared.HrisBenefit{
            CoverageLevel: shared.CoverageLevelEmployeeSpouse.ToPointer(),
            CreatedAt: types.MustNewTimeFromString("2020-06-11T01:24:05.654Z"),
            Currency: unifiedgosdk.Pointer("JOD"),
            Description: unifiedgosdk.Pointer("Vomito voluptas dolor sed."),
            EmployerContributionAmount: unifiedgosdk.Pointer[float64](185006.0),
            EmployerContributionMaxAmount: unifiedgosdk.Pointer[float64](179093.0),
            EmployerContributionType: shared.EmployerContributionTypePercentage.ToPointer(),
            Frequency: shared.HrisBenefitFrequencyHour.ToPointer(),
            ID: unifiedgosdk.Pointer("7b6897d5-1178-4063-9d82-be1ed0beebce"),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("Frozen Wooden Ball"),
            Tax: shared.TaxPreTax.ToPointer(),
            Type: shared.HrisBenefitTypeGarnishment.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-03-06T20:45:33.638Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisBenefit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchHrisBenefitRequest](../../pkg/models/operations/patchhrisbenefitrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchHrisBenefitResponse](../../pkg/models/operations/patchhrisbenefitresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisCompany

Update a company

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisCompany" method="patch" path="/hris/{connection_id}/company/{id}" example="hris_company" -->
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

    res, err := s.Hris.PatchHrisCompany(ctx, operations.PatchHrisCompanyRequest{
        HrisCompany: shared.HrisCompany{
            Address: &shared.PropertyHrisCompanyAddress{
                Address1: unifiedgosdk.Pointer("2549 Church Walk"),
                City: unifiedgosdk.Pointer("Lake Nettiebury"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("32877-4898"),
                Region: unifiedgosdk.Pointer("Idaho"),
                RegionCode: unifiedgosdk.Pointer("PA"),
            },
            CreatedAt: types.MustNewTimeFromString("2021-05-02T22:27:38.970Z"),
            ID: unifiedgosdk.Pointer("50181e51-b4bf-4da1-a86f-e9d457a25773"),
            LegalName: unifiedgosdk.Pointer("Schultz LLC"),
            Name: unifiedgosdk.Pointer("Gottlieb Group"),
            UpdatedAt: types.MustNewTimeFromString("2026-09-06T19:24:54.155Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisCompany != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchHrisCompanyRequest](../../pkg/models/operations/patchhriscompanyrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchHrisCompanyResponse](../../pkg/models/operations/patchhriscompanyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisDeduction

Update a deduction

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisDeduction" method="patch" path="/hris/{connection_id}/deduction/{id}" example="hris_deduction" -->
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

    res, err := s.Hris.PatchHrisDeduction(ctx, operations.PatchHrisDeductionRequest{
        HrisDeduction: shared.HrisDeduction{
            Amount: unifiedgosdk.Pointer[float64](139655.0),
            CoverageLevel: shared.HrisDeductionCoverageLevelEmployeeOnly.ToPointer(),
            CreatedAt: types.MustNewTimeFromString("2020-02-05T01:46:31.384Z"),
            EndAt: types.MustNewTimeFromString("2026-05-23T20:08:22.536Z"),
            Frequency: shared.HrisDeductionFrequencyMonth.ToPointer(),
            ID: unifiedgosdk.Pointer("c80210bf-e463-44e9-bccf-a0313d59aad8"),
            IsActive: unifiedgosdk.Pointer(false),
            Notes: unifiedgosdk.Pointer("Carmen desidero."),
            StartAt: types.MustNewTimeFromString("2025-02-18T21:39:35.482Z"),
            Type: shared.HrisDeductionTypeFixed.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-03-02T13:27:26.620Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDeduction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PatchHrisDeductionRequest](../../pkg/models/operations/patchhrisdeductionrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.PatchHrisDeductionResponse](../../pkg/models/operations/patchhrisdeductionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisDevice

Update a device

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisDevice" method="patch" path="/hris/{connection_id}/device/{id}" example="hris_device" -->
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

    res, err := s.Hris.PatchHrisDevice(ctx, operations.PatchHrisDeviceRequest{
        HrisDevice: shared.HrisDevice{
            AdminUserIds: []string{},
            AssetTag: unifiedgosdk.Pointer("dpho9OuFNG"),
            CreatedAt: types.MustNewTimeFromString("2019-04-04T17:11:40.322Z"),
            HasAntivirus: unifiedgosdk.Pointer(false),
            HasFirewall: unifiedgosdk.Pointer(true),
            HasHdEncrypted: unifiedgosdk.Pointer(true),
            HasPasswordManager: unifiedgosdk.Pointer(true),
            HasScreenlock: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("5ca72f98-4e3a-4bc4-8d76-bfc94f272ecc"),
            IsMissing: unifiedgosdk.Pointer(false),
            Manufacturer: unifiedgosdk.Pointer("Sanford - Hamill"),
            Model: unifiedgosdk.Pointer("Refined"),
            Name: unifiedgosdk.Pointer("cross_contamination_if.rar"),
            Os: unifiedgosdk.Pointer("monitor"),
            OsVersion: unifiedgosdk.Pointer("1.12.16"),
            UpdatedAt: types.MustNewTimeFromString("2023-05-21T13:19:00.275Z"),
            Version: unifiedgosdk.Pointer("2.20.17"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDevice != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PatchHrisDeviceRequest](../../pkg/models/operations/patchhrisdevicerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.PatchHrisDeviceResponse](../../pkg/models/operations/patchhrisdeviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisDocument

Update a document

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisDocument" method="patch" path="/hris/{connection_id}/document/{id}" example="hris_document" -->
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

    res, err := s.Hris.PatchHrisDocument(ctx, operations.PatchHrisDocumentRequest{
        HrisDocument: shared.HrisDocument{
            CreatedAt: types.MustNewTimeFromString("2022-10-27T11:47:26.086Z"),
            DocumentURL: unifiedgosdk.Pointer("https://sore-decision.biz/"),
            Filename: unifiedgosdk.Pointer("ridge_forager.xsl"),
            ID: unifiedgosdk.Pointer("1ab16669-88e1-458b-a1c6-6e8ce6719e6e"),
            Type: shared.HrisDocumentTypePolicy.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-09-17T18:31:17.147Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDocument != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PatchHrisDocumentRequest](../../pkg/models/operations/patchhrisdocumentrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PatchHrisDocumentResponse](../../pkg/models/operations/patchhrisdocumentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisEmployee

Update an employee

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisEmployee" method="patch" path="/hris/{connection_id}/employee/{id}" example="hris_employee" -->
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

    res, err := s.Hris.PatchHrisEmployee(ctx, operations.PatchHrisEmployeeRequest{
        HrisEmployee: shared.HrisEmployee{
            Address: &shared.PropertyHrisEmployeeAddress{
                Address1: unifiedgosdk.Pointer("52008 Lansdowne Road"),
                Address2: unifiedgosdk.Pointer("Apt. 101"),
                City: unifiedgosdk.Pointer("Connellyberg"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("18978"),
                Region: unifiedgosdk.Pointer("South Dakota"),
                RegionCode: unifiedgosdk.Pointer("NM"),
            },
            Bio: unifiedgosdk.Pointer("sushi devotee, singer"),
            Compensation: []shared.HrisCompensation{
                shared.HrisCompensation{
                    Amount: unifiedgosdk.Pointer[float64](69148.0),
                    Currency: unifiedgosdk.Pointer("CRC"),
                    Frequency: shared.HrisCompensationFrequencyQuarter.ToPointer(),
                    Notes: unifiedgosdk.Pointer("Tergeo laborum laboriosam tutis."),
                    Type: shared.HrisCompensationTypeEquity.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2019-09-16T15:08:53.262Z"),
            Currency: unifiedgosdk.Pointer("IDR"),
            DateOfBirth: types.MustNewTimeFromString("2001-04-22"),
            Emails: []shared.HrisEmail{
                shared.HrisEmail{
                    Email: "Zetta_Prohaska67@hotmail.com",
                    Type: shared.HrisEmailTypeHome.ToPointer(),
                },
            },
            EmployeeNumber: unifiedgosdk.Pointer("YuOt169CGu"),
            EmploymentStatus: shared.EmploymentStatusActive.ToPointer(),
            EmploymentType: shared.HrisEmployeeEmploymentTypeVolunteer.ToPointer(),
            FirstName: unifiedgosdk.Pointer("Zetta"),
            Gender: shared.HrisEmployeeGenderIntersex.ToPointer(),
            HasMfa: unifiedgosdk.Pointer(true),
            HiredAt: types.MustNewTimeFromString("2023-05-11T03:53:15.207Z"),
            ID: unifiedgosdk.Pointer("a3db698c-6df7-4684-a87f-ebba77407059"),
            ImageURL: unifiedgosdk.Pointer("https://loremflickr.com/3684/2116?lock=4686991638584456"),
            LanguageLocale: unifiedgosdk.Pointer("es"),
            LastName: unifiedgosdk.Pointer("Prohaska"),
            Locations: []shared.HrisLocation{},
            MaritalStatus: shared.MaritalStatusMarried.ToPointer(),
            Metadata: []shared.HrisMetadata{
                shared.HrisMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateHrisMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.HrisMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("a4f44f4f-e771-4d99-8542-ec1711248755"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateHrisMetadataValueStr(
                        "tenetur",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Zetta Prohaska"),
            Pronouns: unifiedgosdk.Pointer("she/her"),
            Relationships: []shared.HrisEmployeerelationship{
                shared.HrisEmployeerelationship{
                    Emails: []shared.HrisEmail{
                        shared.HrisEmail{
                            Email: "Deshaun.Sanford24@yahoo.com",
                        },
                        shared.HrisEmail{
                            Email: "Rebeca.Dibbert11@hotmail.com",
                        },
                        shared.HrisEmail{
                            Email: "Hester80@gmail.com",
                        },
                    },
                    Name: unifiedgosdk.Pointer("Automotive"),
                    Type: shared.HrisEmployeerelationshipTypeEmergency.ToPointer(),
                },
                shared.HrisEmployeerelationship{
                    Emails: []shared.HrisEmail{
                        shared.HrisEmail{
                            Email: "Benedict_Wisozk83@hotmail.com",
                        },
                        shared.HrisEmail{
                            Email: "Princess_Rath43@gmail.com",
                        },
                        shared.HrisEmail{
                            Email: "Elmira92@yahoo.com",
                        },
                    },
                    Name: unifiedgosdk.Pointer("Music"),
                    Type: shared.HrisEmployeerelationshipTypeFriend.ToPointer(),
                },
                shared.HrisEmployeerelationship{
                    Emails: []shared.HrisEmail{
                        shared.HrisEmail{
                            Email: "Jane30@gmail.com",
                        },
                    },
                    Name: unifiedgosdk.Pointer("Jewelry"),
                    Type: shared.HrisEmployeerelationshipTypeSibling.ToPointer(),
                },
            },
            Salutation: unifiedgosdk.Pointer("Miss"),
            SsnSin: unifiedgosdk.Pointer("yMRtj0Q3xO"),
            StorageQuotaAllocated: unifiedgosdk.Pointer[float64](3674489.0),
            StorageQuotaAvailable: unifiedgosdk.Pointer[float64](7748057.0),
            StorageQuotaUsed: unifiedgosdk.Pointer[float64](301727.0),
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "(409) 801-3705",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
            },
            TerminationReason: unifiedgosdk.Pointer("Communis adnuo damnatio atavus terebro acies canis cogito triumphus creber temptatio defendo cubo amissio paulatim corroboro."),
            TimeoffDaysTotal: unifiedgosdk.Pointer[float64](12.0),
            TimeoffDaysUsed: unifiedgosdk.Pointer[float64](6.0),
            Timezone: unifiedgosdk.Pointer("Africa/Harare"),
            Title: unifiedgosdk.Pointer("Investor Paradigm Liaison"),
            UpdatedAt: types.MustNewTimeFromString("2022-02-19T14:46:57.350Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisEmployee != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PatchHrisEmployeeRequest](../../pkg/models/operations/patchhrisemployeerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PatchHrisEmployeeResponse](../../pkg/models/operations/patchhrisemployeeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisGroup

Update a group

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisGroup" method="patch" path="/hris/{connection_id}/group/{id}" example="hris_group" -->
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

    res, err := s.Hris.PatchHrisGroup(ctx, operations.PatchHrisGroupRequest{
        HrisGroup: shared.HrisGroup{
            CreatedAt: types.MustNewTimeFromString("2023-11-01T13:13:40.714Z"),
            Description: unifiedgosdk.Pointer("Absorbeo casso."),
            ID: unifiedgosdk.Pointer("2c881674-f9e4-4cf7-bb1b-09837efdfdaa"),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("Games"),
            Type: shared.HrisGroupTypeBusinessUnit.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-04-24T09:38:41.471Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.PatchHrisGroupRequest](../../pkg/models/operations/patchhrisgrouprequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.PatchHrisGroupResponse](../../pkg/models/operations/patchhrisgroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisLocation

Update a location

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisLocation" method="patch" path="/hris/{connection_id}/location/{id}" example="hris_location" -->
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

    res, err := s.Hris.PatchHrisLocation(ctx, operations.PatchHrisLocationRequest{
        HrisLocation: shared.HrisLocation{
            Address: &shared.PropertyHrisLocationAddress{
                Address1: unifiedgosdk.Pointer("2743 Connelly Summit"),
                Address2: unifiedgosdk.Pointer("Apt. 350"),
                City: unifiedgosdk.Pointer("Titusville"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("16154-1095"),
                Region: unifiedgosdk.Pointer("Oregon"),
                RegionCode: unifiedgosdk.Pointer("AL"),
            },
            CreatedAt: types.MustNewTimeFromString("2021-07-18T10:32:01.414Z"),
            Currency: unifiedgosdk.Pointer("MUR"),
            Description: unifiedgosdk.Pointer("Acervus caries."),
            ExternalIdentifier: unifiedgosdk.Pointer("83f19f90-c3ce-4617-9ba5-6babae37a9a0"),
            ID: unifiedgosdk.Pointer("7cc227c3-51b2-44b4-b67f-53bc5b23b44a"),
            IsActive: unifiedgosdk.Pointer(true),
            IsHq: unifiedgosdk.Pointer(false),
            LanguageLocale: unifiedgosdk.Pointer("fr"),
            Name: unifiedgosdk.Pointer("adhuc"),
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "(710) 550-6997",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
                shared.HrisTelephone{
                    Telephone: "(208) 555-8542",
                    Type: shared.HrisTelephoneTypeHome.ToPointer(),
                },
                shared.HrisTelephone{
                    Telephone: "(712) 473-5482",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
            },
            Timezone: unifiedgosdk.Pointer("America/Guyana"),
            UpdatedAt: types.MustNewTimeFromString("2023-06-09T08:50:56.340Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisLocation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PatchHrisLocationRequest](../../pkg/models/operations/patchhrislocationrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PatchHrisLocationResponse](../../pkg/models/operations/patchhrislocationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisTimeoff

Update a timeoff

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisTimeoff" method="patch" path="/hris/{connection_id}/timeoff/{id}" example="hris_timeoff" -->
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

    res, err := s.Hris.PatchHrisTimeoff(ctx, operations.PatchHrisTimeoffRequest{
        HrisTimeoff: shared.HrisTimeoff{
            ApprovedAt: types.MustNewTimeFromString("2022-02-20T22:44:28.295Z"),
            Comments: unifiedgosdk.Pointer("Blandior ventus curiositas amplitudo."),
            CreatedAt: types.MustNewTimeFromString("2021-10-06T18:00:20.615Z"),
            Duration: unifiedgosdk.Pointer[float64](4.0),
            DurationType: shared.DurationTypeDay.ToPointer(),
            EndAt: types.MustNewTimeFromString("2024-12-08T04:10:50.230Z"),
            ID: unifiedgosdk.Pointer("dcc4b0a9-b917-4783-ad9d-f8b0ae043999"),
            IsPaid: unifiedgosdk.Pointer(true),
            OriginalType: unifiedgosdk.Pointer("acerbitas ut"),
            Reason: unifiedgosdk.Pointer("verto"),
            StartAt: types.MustNewTimeFromString("2023-08-23T15:25:53.578Z"),
            Status: shared.HrisTimeoffStatusDenied.ToPointer(),
            Type: shared.HrisTimeoffTypeInLieu.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-07-07T22:58:27.092Z"),
            UserID: "<id>",
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeoff != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchHrisTimeoffRequest](../../pkg/models/operations/patchhristimeoffrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchHrisTimeoffResponse](../../pkg/models/operations/patchhristimeoffresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisTimeshift

Update a timeshift

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisTimeshift" method="patch" path="/hris/{connection_id}/timeshift/{id}" example="hris_timeshift" -->
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

    res, err := s.Hris.PatchHrisTimeshift(ctx, operations.PatchHrisTimeshiftRequest{
        HrisTimeshift: shared.HrisTimeshift{
            ApprovedAt: types.MustNewTimeFromString("2023-06-06T02:42:24.617Z"),
            Compensation: []shared.HrisCompensation{
                shared.HrisCompensation{
                    Amount: unifiedgosdk.Pointer[float64](76761.0),
                    Currency: unifiedgosdk.Pointer("JPY"),
                    Frequency: shared.HrisCompensationFrequencyHour.ToPointer(),
                    Notes: unifiedgosdk.Pointer("Annus adficio suasoria architecto aggero."),
                    Type: shared.HrisCompensationTypeOther.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2019-07-01T23:53:15.738Z"),
            EmployeeUserID: "<id>",
            EndAt: types.MustNewTimeFromString("2026-08-26T06:38:23.101Z"),
            Hours: unifiedgosdk.Pointer[float64](8.0),
            ID: unifiedgosdk.Pointer("1bcedceb-7227-49c8-8729-babbe7c0f8fa"),
            IsApproved: unifiedgosdk.Pointer(true),
            StartAt: types.MustNewTimeFromString("2023-06-25T07:50:01.861Z"),
            UpdatedAt: types.MustNewTimeFromString("2021-06-23T03:00:57.052Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeshift != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PatchHrisTimeshiftRequest](../../pkg/models/operations/patchhristimeshiftrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.PatchHrisTimeshiftResponse](../../pkg/models/operations/patchhristimeshiftresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisAttendance

Remove an attendance

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisAttendance" method="delete" path="/hris/{connection_id}/attendance/{id}" -->
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

    res, err := s.Hris.RemoveHrisAttendance(ctx, operations.RemoveHrisAttendanceRequest{
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.RemoveHrisAttendanceRequest](../../pkg/models/operations/removehrisattendancerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.RemoveHrisAttendanceResponse](../../pkg/models/operations/removehrisattendanceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisBankaccount

Remove a bankaccount

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisBankaccount" method="delete" path="/hris/{connection_id}/bankaccount/{id}" -->
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

    res, err := s.Hris.RemoveHrisBankaccount(ctx, operations.RemoveHrisBankaccountRequest{
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
| `request`                                                                                              | [operations.RemoveHrisBankaccountRequest](../../pkg/models/operations/removehrisbankaccountrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.RemoveHrisBankaccountResponse](../../pkg/models/operations/removehrisbankaccountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisBenefit

Remove a benefit

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisBenefit" method="delete" path="/hris/{connection_id}/benefit/{id}" -->
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

    res, err := s.Hris.RemoveHrisBenefit(ctx, operations.RemoveHrisBenefitRequest{
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
| `request`                                                                                      | [operations.RemoveHrisBenefitRequest](../../pkg/models/operations/removehrisbenefitrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemoveHrisBenefitResponse](../../pkg/models/operations/removehrisbenefitresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisCompany

Remove a company

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisCompany" method="delete" path="/hris/{connection_id}/company/{id}" -->
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

    res, err := s.Hris.RemoveHrisCompany(ctx, operations.RemoveHrisCompanyRequest{
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
| `request`                                                                                      | [operations.RemoveHrisCompanyRequest](../../pkg/models/operations/removehriscompanyrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemoveHrisCompanyResponse](../../pkg/models/operations/removehriscompanyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisDeduction

Remove a deduction

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisDeduction" method="delete" path="/hris/{connection_id}/deduction/{id}" -->
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

    res, err := s.Hris.RemoveHrisDeduction(ctx, operations.RemoveHrisDeductionRequest{
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
| `request`                                                                                          | [operations.RemoveHrisDeductionRequest](../../pkg/models/operations/removehrisdeductionrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.RemoveHrisDeductionResponse](../../pkg/models/operations/removehrisdeductionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisDevice

Remove a device

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisDevice" method="delete" path="/hris/{connection_id}/device/{id}" -->
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

    res, err := s.Hris.RemoveHrisDevice(ctx, operations.RemoveHrisDeviceRequest{
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.RemoveHrisDeviceRequest](../../pkg/models/operations/removehrisdevicerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.RemoveHrisDeviceResponse](../../pkg/models/operations/removehrisdeviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisDocument

Remove a document

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisDocument" method="delete" path="/hris/{connection_id}/document/{id}" -->
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

    res, err := s.Hris.RemoveHrisDocument(ctx, operations.RemoveHrisDocumentRequest{
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
| `request`                                                                                        | [operations.RemoveHrisDocumentRequest](../../pkg/models/operations/removehrisdocumentrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.RemoveHrisDocumentResponse](../../pkg/models/operations/removehrisdocumentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisEmployee

Remove an employee

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisEmployee" method="delete" path="/hris/{connection_id}/employee/{id}" -->
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

    res, err := s.Hris.RemoveHrisEmployee(ctx, operations.RemoveHrisEmployeeRequest{
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
| `request`                                                                                        | [operations.RemoveHrisEmployeeRequest](../../pkg/models/operations/removehrisemployeerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.RemoveHrisEmployeeResponse](../../pkg/models/operations/removehrisemployeeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisGroup

Remove a group

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisGroup" method="delete" path="/hris/{connection_id}/group/{id}" -->
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

    res, err := s.Hris.RemoveHrisGroup(ctx, operations.RemoveHrisGroupRequest{
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
| `request`                                                                                  | [operations.RemoveHrisGroupRequest](../../pkg/models/operations/removehrisgrouprequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.RemoveHrisGroupResponse](../../pkg/models/operations/removehrisgroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisLocation

Remove a location

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisLocation" method="delete" path="/hris/{connection_id}/location/{id}" -->
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

    res, err := s.Hris.RemoveHrisLocation(ctx, operations.RemoveHrisLocationRequest{
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
| `request`                                                                                        | [operations.RemoveHrisLocationRequest](../../pkg/models/operations/removehrislocationrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.RemoveHrisLocationResponse](../../pkg/models/operations/removehrislocationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisTimeoff

Remove a timeoff

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisTimeoff" method="delete" path="/hris/{connection_id}/timeoff/{id}" -->
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

    res, err := s.Hris.RemoveHrisTimeoff(ctx, operations.RemoveHrisTimeoffRequest{
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
| `request`                                                                                      | [operations.RemoveHrisTimeoffRequest](../../pkg/models/operations/removehristimeoffrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemoveHrisTimeoffResponse](../../pkg/models/operations/removehristimeoffresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisTimeshift

Remove a timeshift

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisTimeshift" method="delete" path="/hris/{connection_id}/timeshift/{id}" -->
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

    res, err := s.Hris.RemoveHrisTimeshift(ctx, operations.RemoveHrisTimeshiftRequest{
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
| `request`                                                                                          | [operations.RemoveHrisTimeshiftRequest](../../pkg/models/operations/removehristimeshiftrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.RemoveHrisTimeshiftResponse](../../pkg/models/operations/removehristimeshiftresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisAttendance

Update an attendance

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisAttendance" method="put" path="/hris/{connection_id}/attendance/{id}" example="hris_attendance" -->
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

    res, err := s.Hris.UpdateHrisAttendance(ctx, operations.UpdateHrisAttendanceRequest{
        HrisAttendance: shared.HrisAttendance{
            Address: &shared.PropertyHrisAttendanceAddress{
                Address1: unifiedgosdk.Pointer("14108 Allie Flats"),
                City: unifiedgosdk.Pointer("Kearaborough"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("23844-2344"),
                Region: unifiedgosdk.Pointer("Tennessee"),
                RegionCode: unifiedgosdk.Pointer("CA"),
            },
            ApprovedAt: types.MustNewTimeFromString("2021-08-13T10:37:55.766Z"),
            Breaks: []shared.HrisAttendanceBreak{
                shared.HrisAttendanceBreak{
                    DurationMinutes: unifiedgosdk.Pointer[float64](12.0),
                    EndAt: types.MustNewTimeFromString("2023-10-23T02:00:25.590Z"),
                    ID: unifiedgosdk.Pointer("d60a1001-5a8a-4991-8c21-f4da6036cc87"),
                    IsPaid: unifiedgosdk.Pointer(true),
                    Name: unifiedgosdk.Pointer("Lunch"),
                    StartAt: types.MustNewTimeFromString("2023-10-16T06:21:50.743Z"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2021-08-10T19:43:18.452Z"),
            Currency: unifiedgosdk.Pointer("UGX"),
            DeclaredTipsAmount: unifiedgosdk.Pointer[float64](161.0),
            EmployeeUserID: "<id>",
            EndAt: types.MustTimeFromString("2024-04-06T15:33:47.993Z"),
            HourlyRate: unifiedgosdk.Pointer[float64](53.0),
            Hours: unifiedgosdk.Pointer[float64](10.0),
            ID: unifiedgosdk.Pointer("b2360f22-cde1-41a5-9cea-b76990466cf3"),
            JobName: unifiedgosdk.Pointer("Global Creative Supervisor"),
            NonCashTipsAmount: unifiedgosdk.Pointer[float64](54.0),
            StartAt: types.MustTimeFromString("2021-11-09T11:31:11.588Z"),
            Status: shared.HrisAttendanceStatusClosed.ToPointer(),
            Timezone: unifiedgosdk.Pointer("America/Atikokan"),
            UpdatedAt: types.MustNewTimeFromString("2022-01-17T03:19:38.954Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisAttendance != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdateHrisAttendanceRequest](../../pkg/models/operations/updatehrisattendancerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.UpdateHrisAttendanceResponse](../../pkg/models/operations/updatehrisattendanceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisBankaccount

Update a bankaccount

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisBankaccount" method="put" path="/hris/{connection_id}/bankaccount/{id}" example="hris_bankaccount" -->
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

    res, err := s.Hris.UpdateHrisBankaccount(ctx, operations.UpdateHrisBankaccountRequest{
        HrisBankaccount: shared.HrisBankaccount{
            AccountNumber: unifiedgosdk.Pointer("****3777"),
            AccountNumberLast4: unifiedgosdk.Pointer("3777"),
            AccountType: shared.HrisBankaccountAccountTypeChecking.ToPointer(),
            BankName: unifiedgosdk.Pointer("Huel Group"),
            CreatedAt: types.MustNewTimeFromString("2019-11-16T16:43:45.976Z"),
            ID: unifiedgosdk.Pointer("ee0b7b15-c45b-4cad-ba3d-ae501748d719"),
            IsPrimary: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("Checking Account"),
            RoutingNumber: unifiedgosdk.Pointer("448650724"),
            UpdatedAt: types.MustNewTimeFromString("2025-06-05T05:39:50.317Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisBankaccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdateHrisBankaccountRequest](../../pkg/models/operations/updatehrisbankaccountrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UpdateHrisBankaccountResponse](../../pkg/models/operations/updatehrisbankaccountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisBenefit

Update a benefit

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisBenefit" method="put" path="/hris/{connection_id}/benefit/{id}" example="hris_benefit" -->
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

    res, err := s.Hris.UpdateHrisBenefit(ctx, operations.UpdateHrisBenefitRequest{
        HrisBenefit: shared.HrisBenefit{
            CoverageLevel: shared.CoverageLevelEmployeeSpouse.ToPointer(),
            CreatedAt: types.MustNewTimeFromString("2020-06-11T01:24:05.654Z"),
            Currency: unifiedgosdk.Pointer("JOD"),
            Description: unifiedgosdk.Pointer("Vomito voluptas dolor sed."),
            EmployerContributionAmount: unifiedgosdk.Pointer[float64](185006.0),
            EmployerContributionMaxAmount: unifiedgosdk.Pointer[float64](179093.0),
            EmployerContributionType: shared.EmployerContributionTypePercentage.ToPointer(),
            Frequency: shared.HrisBenefitFrequencyHour.ToPointer(),
            ID: unifiedgosdk.Pointer("7b6897d5-1178-4063-9d82-be1ed0beebce"),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("Frozen Wooden Ball"),
            Tax: shared.TaxPreTax.ToPointer(),
            Type: shared.HrisBenefitTypeGarnishment.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-03-06T20:45:33.638Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisBenefit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateHrisBenefitRequest](../../pkg/models/operations/updatehrisbenefitrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateHrisBenefitResponse](../../pkg/models/operations/updatehrisbenefitresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisCompany

Update a company

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisCompany" method="put" path="/hris/{connection_id}/company/{id}" example="hris_company" -->
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

    res, err := s.Hris.UpdateHrisCompany(ctx, operations.UpdateHrisCompanyRequest{
        HrisCompany: shared.HrisCompany{
            Address: &shared.PropertyHrisCompanyAddress{
                Address1: unifiedgosdk.Pointer("2549 Church Walk"),
                City: unifiedgosdk.Pointer("Lake Nettiebury"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("32877-4898"),
                Region: unifiedgosdk.Pointer("Idaho"),
                RegionCode: unifiedgosdk.Pointer("PA"),
            },
            CreatedAt: types.MustNewTimeFromString("2021-05-02T22:27:38.970Z"),
            ID: unifiedgosdk.Pointer("50181e51-b4bf-4da1-a86f-e9d457a25773"),
            LegalName: unifiedgosdk.Pointer("Schultz LLC"),
            Name: unifiedgosdk.Pointer("Gottlieb Group"),
            UpdatedAt: types.MustNewTimeFromString("2026-09-06T19:24:54.155Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisCompany != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateHrisCompanyRequest](../../pkg/models/operations/updatehriscompanyrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateHrisCompanyResponse](../../pkg/models/operations/updatehriscompanyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisDeduction

Update a deduction

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisDeduction" method="put" path="/hris/{connection_id}/deduction/{id}" example="hris_deduction" -->
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

    res, err := s.Hris.UpdateHrisDeduction(ctx, operations.UpdateHrisDeductionRequest{
        HrisDeduction: shared.HrisDeduction{
            Amount: unifiedgosdk.Pointer[float64](139655.0),
            CoverageLevel: shared.HrisDeductionCoverageLevelEmployeeOnly.ToPointer(),
            CreatedAt: types.MustNewTimeFromString("2020-02-05T01:46:31.384Z"),
            EndAt: types.MustNewTimeFromString("2026-05-23T20:08:22.536Z"),
            Frequency: shared.HrisDeductionFrequencyMonth.ToPointer(),
            ID: unifiedgosdk.Pointer("c80210bf-e463-44e9-bccf-a0313d59aad8"),
            IsActive: unifiedgosdk.Pointer(false),
            Notes: unifiedgosdk.Pointer("Carmen desidero."),
            StartAt: types.MustNewTimeFromString("2025-02-18T21:39:35.482Z"),
            Type: shared.HrisDeductionTypeFixed.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-03-02T13:27:26.620Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDeduction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateHrisDeductionRequest](../../pkg/models/operations/updatehrisdeductionrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateHrisDeductionResponse](../../pkg/models/operations/updatehrisdeductionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisDevice

Update a device

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisDevice" method="put" path="/hris/{connection_id}/device/{id}" example="hris_device" -->
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

    res, err := s.Hris.UpdateHrisDevice(ctx, operations.UpdateHrisDeviceRequest{
        HrisDevice: shared.HrisDevice{
            AdminUserIds: []string{},
            AssetTag: unifiedgosdk.Pointer("dpho9OuFNG"),
            CreatedAt: types.MustNewTimeFromString("2019-04-04T17:11:40.322Z"),
            HasAntivirus: unifiedgosdk.Pointer(false),
            HasFirewall: unifiedgosdk.Pointer(true),
            HasHdEncrypted: unifiedgosdk.Pointer(true),
            HasPasswordManager: unifiedgosdk.Pointer(true),
            HasScreenlock: unifiedgosdk.Pointer(true),
            ID: unifiedgosdk.Pointer("5ca72f98-4e3a-4bc4-8d76-bfc94f272ecc"),
            IsMissing: unifiedgosdk.Pointer(false),
            Manufacturer: unifiedgosdk.Pointer("Sanford - Hamill"),
            Model: unifiedgosdk.Pointer("Refined"),
            Name: unifiedgosdk.Pointer("cross_contamination_if.rar"),
            Os: unifiedgosdk.Pointer("monitor"),
            OsVersion: unifiedgosdk.Pointer("1.12.16"),
            UpdatedAt: types.MustNewTimeFromString("2023-05-21T13:19:00.275Z"),
            Version: unifiedgosdk.Pointer("2.20.17"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDevice != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateHrisDeviceRequest](../../pkg/models/operations/updatehrisdevicerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpdateHrisDeviceResponse](../../pkg/models/operations/updatehrisdeviceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisDocument

Update a document

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisDocument" method="put" path="/hris/{connection_id}/document/{id}" example="hris_document" -->
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

    res, err := s.Hris.UpdateHrisDocument(ctx, operations.UpdateHrisDocumentRequest{
        HrisDocument: shared.HrisDocument{
            CreatedAt: types.MustNewTimeFromString("2022-10-27T11:47:26.086Z"),
            DocumentURL: unifiedgosdk.Pointer("https://sore-decision.biz/"),
            Filename: unifiedgosdk.Pointer("ridge_forager.xsl"),
            ID: unifiedgosdk.Pointer("1ab16669-88e1-458b-a1c6-6e8ce6719e6e"),
            Type: shared.HrisDocumentTypePolicy.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-09-17T18:31:17.147Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisDocument != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateHrisDocumentRequest](../../pkg/models/operations/updatehrisdocumentrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.UpdateHrisDocumentResponse](../../pkg/models/operations/updatehrisdocumentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisEmployee

Update an employee

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisEmployee" method="put" path="/hris/{connection_id}/employee/{id}" example="hris_employee" -->
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

    res, err := s.Hris.UpdateHrisEmployee(ctx, operations.UpdateHrisEmployeeRequest{
        HrisEmployee: shared.HrisEmployee{
            Address: &shared.PropertyHrisEmployeeAddress{
                Address1: unifiedgosdk.Pointer("52008 Lansdowne Road"),
                Address2: unifiedgosdk.Pointer("Apt. 101"),
                City: unifiedgosdk.Pointer("Connellyberg"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("18978"),
                Region: unifiedgosdk.Pointer("South Dakota"),
                RegionCode: unifiedgosdk.Pointer("NM"),
            },
            Bio: unifiedgosdk.Pointer("sushi devotee, singer"),
            Compensation: []shared.HrisCompensation{
                shared.HrisCompensation{
                    Amount: unifiedgosdk.Pointer[float64](69148.0),
                    Currency: unifiedgosdk.Pointer("CRC"),
                    Frequency: shared.HrisCompensationFrequencyQuarter.ToPointer(),
                    Notes: unifiedgosdk.Pointer("Tergeo laborum laboriosam tutis."),
                    Type: shared.HrisCompensationTypeEquity.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2019-09-16T15:08:53.262Z"),
            Currency: unifiedgosdk.Pointer("IDR"),
            DateOfBirth: types.MustNewTimeFromString("2001-04-22"),
            Emails: []shared.HrisEmail{
                shared.HrisEmail{
                    Email: "Zetta_Prohaska67@hotmail.com",
                    Type: shared.HrisEmailTypeHome.ToPointer(),
                },
            },
            EmployeeNumber: unifiedgosdk.Pointer("YuOt169CGu"),
            EmploymentStatus: shared.EmploymentStatusActive.ToPointer(),
            EmploymentType: shared.HrisEmployeeEmploymentTypeVolunteer.ToPointer(),
            FirstName: unifiedgosdk.Pointer("Zetta"),
            Gender: shared.HrisEmployeeGenderIntersex.ToPointer(),
            HasMfa: unifiedgosdk.Pointer(true),
            HiredAt: types.MustNewTimeFromString("2023-05-11T03:53:15.207Z"),
            ID: unifiedgosdk.Pointer("a3db698c-6df7-4684-a87f-ebba77407059"),
            ImageURL: unifiedgosdk.Pointer("https://loremflickr.com/3684/2116?lock=4686991638584456"),
            LanguageLocale: unifiedgosdk.Pointer("es"),
            LastName: unifiedgosdk.Pointer("Prohaska"),
            Locations: []shared.HrisLocation{},
            MaritalStatus: shared.MaritalStatusMarried.ToPointer(),
            Metadata: []shared.HrisMetadata{
                shared.HrisMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateHrisMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.HrisMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("a4f44f4f-e771-4d99-8542-ec1711248755"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateHrisMetadataValueStr(
                        "tenetur",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Zetta Prohaska"),
            Pronouns: unifiedgosdk.Pointer("she/her"),
            Relationships: []shared.HrisEmployeerelationship{
                shared.HrisEmployeerelationship{
                    Emails: []shared.HrisEmail{
                        shared.HrisEmail{
                            Email: "Deshaun.Sanford24@yahoo.com",
                        },
                        shared.HrisEmail{
                            Email: "Rebeca.Dibbert11@hotmail.com",
                        },
                        shared.HrisEmail{
                            Email: "Hester80@gmail.com",
                        },
                    },
                    Name: unifiedgosdk.Pointer("Automotive"),
                    Type: shared.HrisEmployeerelationshipTypeEmergency.ToPointer(),
                },
                shared.HrisEmployeerelationship{
                    Emails: []shared.HrisEmail{
                        shared.HrisEmail{
                            Email: "Benedict_Wisozk83@hotmail.com",
                        },
                        shared.HrisEmail{
                            Email: "Princess_Rath43@gmail.com",
                        },
                        shared.HrisEmail{
                            Email: "Elmira92@yahoo.com",
                        },
                    },
                    Name: unifiedgosdk.Pointer("Music"),
                    Type: shared.HrisEmployeerelationshipTypeFriend.ToPointer(),
                },
                shared.HrisEmployeerelationship{
                    Emails: []shared.HrisEmail{
                        shared.HrisEmail{
                            Email: "Jane30@gmail.com",
                        },
                    },
                    Name: unifiedgosdk.Pointer("Jewelry"),
                    Type: shared.HrisEmployeerelationshipTypeSibling.ToPointer(),
                },
            },
            Salutation: unifiedgosdk.Pointer("Miss"),
            SsnSin: unifiedgosdk.Pointer("yMRtj0Q3xO"),
            StorageQuotaAllocated: unifiedgosdk.Pointer[float64](3674489.0),
            StorageQuotaAvailable: unifiedgosdk.Pointer[float64](7748057.0),
            StorageQuotaUsed: unifiedgosdk.Pointer[float64](301727.0),
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "(409) 801-3705",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
            },
            TerminationReason: unifiedgosdk.Pointer("Communis adnuo damnatio atavus terebro acies canis cogito triumphus creber temptatio defendo cubo amissio paulatim corroboro."),
            TimeoffDaysTotal: unifiedgosdk.Pointer[float64](12.0),
            TimeoffDaysUsed: unifiedgosdk.Pointer[float64](6.0),
            Timezone: unifiedgosdk.Pointer("Africa/Harare"),
            Title: unifiedgosdk.Pointer("Investor Paradigm Liaison"),
            UpdatedAt: types.MustNewTimeFromString("2022-02-19T14:46:57.350Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisEmployee != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateHrisEmployeeRequest](../../pkg/models/operations/updatehrisemployeerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.UpdateHrisEmployeeResponse](../../pkg/models/operations/updatehrisemployeeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisGroup

Update a group

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisGroup" method="put" path="/hris/{connection_id}/group/{id}" example="hris_group" -->
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

    res, err := s.Hris.UpdateHrisGroup(ctx, operations.UpdateHrisGroupRequest{
        HrisGroup: shared.HrisGroup{
            CreatedAt: types.MustNewTimeFromString("2023-11-01T13:13:40.714Z"),
            Description: unifiedgosdk.Pointer("Absorbeo casso."),
            ID: unifiedgosdk.Pointer("2c881674-f9e4-4cf7-bb1b-09837efdfdaa"),
            IsActive: unifiedgosdk.Pointer(false),
            Name: unifiedgosdk.Pointer("Games"),
            Type: shared.HrisGroupTypeBusinessUnit.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-04-24T09:38:41.471Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateHrisGroupRequest](../../pkg/models/operations/updatehrisgrouprequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.UpdateHrisGroupResponse](../../pkg/models/operations/updatehrisgroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisLocation

Update a location

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisLocation" method="put" path="/hris/{connection_id}/location/{id}" example="hris_location" -->
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

    res, err := s.Hris.UpdateHrisLocation(ctx, operations.UpdateHrisLocationRequest{
        HrisLocation: shared.HrisLocation{
            Address: &shared.PropertyHrisLocationAddress{
                Address1: unifiedgosdk.Pointer("2743 Connelly Summit"),
                Address2: unifiedgosdk.Pointer("Apt. 350"),
                City: unifiedgosdk.Pointer("Titusville"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("16154-1095"),
                Region: unifiedgosdk.Pointer("Oregon"),
                RegionCode: unifiedgosdk.Pointer("AL"),
            },
            CreatedAt: types.MustNewTimeFromString("2021-07-18T10:32:01.414Z"),
            Currency: unifiedgosdk.Pointer("MUR"),
            Description: unifiedgosdk.Pointer("Acervus caries."),
            ExternalIdentifier: unifiedgosdk.Pointer("83f19f90-c3ce-4617-9ba5-6babae37a9a0"),
            ID: unifiedgosdk.Pointer("7cc227c3-51b2-44b4-b67f-53bc5b23b44a"),
            IsActive: unifiedgosdk.Pointer(true),
            IsHq: unifiedgosdk.Pointer(false),
            LanguageLocale: unifiedgosdk.Pointer("fr"),
            Name: unifiedgosdk.Pointer("adhuc"),
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "(710) 550-6997",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
                shared.HrisTelephone{
                    Telephone: "(208) 555-8542",
                    Type: shared.HrisTelephoneTypeHome.ToPointer(),
                },
                shared.HrisTelephone{
                    Telephone: "(712) 473-5482",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
            },
            Timezone: unifiedgosdk.Pointer("America/Guyana"),
            UpdatedAt: types.MustNewTimeFromString("2023-06-09T08:50:56.340Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisLocation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateHrisLocationRequest](../../pkg/models/operations/updatehrislocationrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.UpdateHrisLocationResponse](../../pkg/models/operations/updatehrislocationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisTimeoff

Update a timeoff

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisTimeoff" method="put" path="/hris/{connection_id}/timeoff/{id}" example="hris_timeoff" -->
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

    res, err := s.Hris.UpdateHrisTimeoff(ctx, operations.UpdateHrisTimeoffRequest{
        HrisTimeoff: shared.HrisTimeoff{
            ApprovedAt: types.MustNewTimeFromString("2022-02-20T22:44:28.295Z"),
            Comments: unifiedgosdk.Pointer("Blandior ventus curiositas amplitudo."),
            CreatedAt: types.MustNewTimeFromString("2021-10-06T18:00:20.615Z"),
            Duration: unifiedgosdk.Pointer[float64](4.0),
            DurationType: shared.DurationTypeDay.ToPointer(),
            EndAt: types.MustNewTimeFromString("2024-12-08T04:10:50.230Z"),
            ID: unifiedgosdk.Pointer("dcc4b0a9-b917-4783-ad9d-f8b0ae043999"),
            IsPaid: unifiedgosdk.Pointer(true),
            OriginalType: unifiedgosdk.Pointer("acerbitas ut"),
            Reason: unifiedgosdk.Pointer("verto"),
            StartAt: types.MustNewTimeFromString("2023-08-23T15:25:53.578Z"),
            Status: shared.HrisTimeoffStatusDenied.ToPointer(),
            Type: shared.HrisTimeoffTypeInLieu.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-07-07T22:58:27.092Z"),
            UserID: "<id>",
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeoff != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateHrisTimeoffRequest](../../pkg/models/operations/updatehristimeoffrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateHrisTimeoffResponse](../../pkg/models/operations/updatehristimeoffresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisTimeshift

Update a timeshift

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisTimeshift" method="put" path="/hris/{connection_id}/timeshift/{id}" example="hris_timeshift" -->
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

    res, err := s.Hris.UpdateHrisTimeshift(ctx, operations.UpdateHrisTimeshiftRequest{
        HrisTimeshift: shared.HrisTimeshift{
            ApprovedAt: types.MustNewTimeFromString("2023-06-06T02:42:24.617Z"),
            Compensation: []shared.HrisCompensation{
                shared.HrisCompensation{
                    Amount: unifiedgosdk.Pointer[float64](76761.0),
                    Currency: unifiedgosdk.Pointer("JPY"),
                    Frequency: shared.HrisCompensationFrequencyHour.ToPointer(),
                    Notes: unifiedgosdk.Pointer("Annus adficio suasoria architecto aggero."),
                    Type: shared.HrisCompensationTypeOther.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2019-07-01T23:53:15.738Z"),
            EmployeeUserID: "<id>",
            EndAt: types.MustNewTimeFromString("2026-08-26T06:38:23.101Z"),
            Hours: unifiedgosdk.Pointer[float64](8.0),
            ID: unifiedgosdk.Pointer("1bcedceb-7227-49c8-8729-babbe7c0f8fa"),
            IsApproved: unifiedgosdk.Pointer(true),
            StartAt: types.MustNewTimeFromString("2023-06-25T07:50:01.861Z"),
            UpdatedAt: types.MustNewTimeFromString("2021-06-23T03:00:57.052Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisTimeshift != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateHrisTimeshiftRequest](../../pkg/models/operations/updatehristimeshiftrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateHrisTimeshiftResponse](../../pkg/models/operations/updatehristimeshiftresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |