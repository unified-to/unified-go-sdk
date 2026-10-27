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
* [CreateHrisEarning](#createhrisearning) - Create an earning
* [CreateHrisEmployee](#createhrisemployee) - Create an employee
* [CreateHrisGroup](#createhrisgroup) - Create a group
* [CreateHrisJob](#createhrisjob) - Create a job
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
* [GetHrisEarning](#gethrisearning) - Retrieve an earning
* [GetHrisEmployee](#gethrisemployee) - Retrieve an employee
* [GetHrisGroup](#gethrisgroup) - Retrieve a group
* [GetHrisJob](#gethrisjob) - Retrieve a job
* [GetHrisLocation](#gethrislocation) - Retrieve a location
* [GetHrisPaycode](#gethrispaycode) - Retrieve a paycode
* [GetHrisPayroll](#gethrispayroll) - Retrieve a payroll
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
* [ListHrisEarnings](#listhrisearnings) - List all earnings
* [ListHrisEmployees](#listhrisemployees) - List all employees
* [ListHrisGroups](#listhrisgroups) - List all groups
* [ListHrisJobs](#listhrisjobs) - List all jobs
* [ListHrisLocations](#listhrislocations) - List all locations
* [ListHrisPaycodes](#listhrispaycodes) - List all paycodes
* [ListHrisPayrolls](#listhrispayrolls) - List all payrolls
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
* [PatchHrisEarning](#patchhrisearning) - Update an earning
* [PatchHrisEmployee](#patchhrisemployee) - Update an employee
* [PatchHrisGroup](#patchhrisgroup) - Update a group
* [PatchHrisJob](#patchhrisjob) - Update a job
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
* [RemoveHrisEarning](#removehrisearning) - Remove an earning
* [RemoveHrisEmployee](#removehrisemployee) - Remove an employee
* [RemoveHrisGroup](#removehrisgroup) - Remove a group
* [RemoveHrisJob](#removehrisjob) - Remove a job
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
* [UpdateHrisEarning](#updatehrisearning) - Update an earning
* [UpdateHrisEmployee](#updatehrisemployee) - Update an employee
* [UpdateHrisGroup](#updatehrisgroup) - Update a group
* [UpdateHrisJob](#updatehrisjob) - Update a job
* [UpdateHrisLocation](#updatehrislocation) - Update a location
* [UpdateHrisTimeoff](#updatehristimeoff) - Update a timeoff
* [UpdateHrisTimeshift](#updatehristimeshift) - Update a timeshift

## CreateHrisAttendance

Creates a HrisAttendance. Provide the unified HrisAttendance fields in the request body; they are mapped to the provider's API. Returns the created HrisAttendance.

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
            ApprovedAt: types.MustNewTimeFromString("2021-09-19T17:03:59.192Z"),
            Breaks: []shared.HrisAttendanceBreak{
                shared.HrisAttendanceBreak{
                    DurationMinutes: unifiedgosdk.Pointer[float64](51.0),
                    EndAt: types.MustNewTimeFromString("2024-02-17T09:31:50.519Z"),
                    ID: unifiedgosdk.Pointer("64d3803c-c859-4c32-8cc3-c9c7f9fa4e81"),
                    IsPaid: unifiedgosdk.Pointer(false),
                    Name: unifiedgosdk.Pointer("Lunch"),
                    StartAt: types.MustNewTimeFromString("2023-11-01T11:18:39.308Z"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2021-08-10T19:43:18.452Z"),
            Currency: unifiedgosdk.Pointer("EUR"),
            DeclaredTipsAmount: unifiedgosdk.Pointer[float64](54.0),
            EmployeeUserID: "<id>",
            EndAt: types.MustTimeFromString("2024-04-26T09:10:57.685Z"),
            HourlyRate: unifiedgosdk.Pointer[float64](16.0),
            Hours: unifiedgosdk.Pointer[float64](10.0),
            ID: unifiedgosdk.Pointer("9b6fbce6-bdbd-4181-8d18-bee35a16bf20"),
            JobName: unifiedgosdk.Pointer("Human Accountability Administrator"),
            Metadata: []shared.HrisMetadata{},
            NonCashTipsAmount: unifiedgosdk.Pointer[float64](167.0),
            StartAt: types.MustTimeFromString("2021-11-11T07:47:35.744Z"),
            Status: shared.HrisAttendanceStatusOpen.ToPointer(),
            Timezone: unifiedgosdk.Pointer("America/Atikokan"),
            UpdatedAt: types.MustNewTimeFromString("2022-01-20T09:07:50.276Z"),
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

Creates a HrisBankaccount. Provide the unified HrisBankaccount fields in the request body; they are mapped to the provider's API. Returns the created HrisBankaccount.

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
            CountryCode: unifiedgosdk.Pointer("US"),
            CreatedAt: types.MustNewTimeFromString("2019-11-16T16:43:45.976Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            ID: unifiedgosdk.Pointer("cd1f5dc0-a9aa-4644-8ea5-e30267401bb5"),
            IsPrimary: unifiedgosdk.Pointer(true),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("Granville Schumm - Checking"),
            Priority: unifiedgosdk.Pointer[float64](2.0),
            RoutingNumber: unifiedgosdk.Pointer("448650724"),
            SplitAmount: unifiedgosdk.Pointer[float64](98.0),
            SplitType: shared.SplitTypePercentage.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-07-06T00:53:08.528Z"),
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

Creates a HrisBenefit. Provide the unified HrisBenefit fields in the request body; they are mapped to the provider's API. Returns the created HrisBenefit.

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
            Currency: unifiedgosdk.Pointer("GBP"),
            Description: unifiedgosdk.Pointer("Vomito voluptas dolor sed."),
            EmployerContributionAmount: unifiedgosdk.Pointer[float64](185006.0),
            EmployerContributionMaxAmount: unifiedgosdk.Pointer[float64](179093.0),
            EmployerContributionType: shared.EmployerContributionTypePercentage.ToPointer(),
            Frequency: shared.HrisBenefitFrequencyMonth.ToPointer(),
            ID: unifiedgosdk.Pointer("f5ede8a2-2b18-48fd-91b3-27382c3dbe20"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("Frozen Wooden Ball"),
            Tax: shared.TaxPreTax.ToPointer(),
            Type: shared.HrisBenefitTypeOther.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-03-23T09:52:37.641Z"),
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

Creates a HrisCompany. Provide the unified HrisCompany fields in the request body; they are mapped to the provider's API. Returns the created HrisCompany.

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
                Address1: unifiedgosdk.Pointer("675 Lavinia Lights"),
                City: unifiedgosdk.Pointer("Rocky Mount"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("67325"),
                Region: unifiedgosdk.Pointer("Kentucky"),
                RegionCode: unifiedgosdk.Pointer("MO"),
            },
            CreatedAt: types.MustNewTimeFromString("2021-05-02T22:27:38.970Z"),
            ID: unifiedgosdk.Pointer("04c72a4e-8524-4883-a389-961d4d216b57"),
            LegalName: unifiedgosdk.Pointer("Gottlieb Group LLC"),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("Gottlieb Group"),
            UpdatedAt: types.MustNewTimeFromString("2026-10-14T12:37:49.345Z"),
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

Creates a HrisDeduction. Provide the unified HrisDeduction fields in the request body; they are mapped to the provider's API. Returns the created HrisDeduction.

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
            Amount: unifiedgosdk.Pointer[float64](564.0),
            CoverageLevel: shared.HrisDeductionCoverageLevelEmployeeChild.ToPointer(),
            CreatedAt: types.MustNewTimeFromString("2020-02-05T01:46:31.384Z"),
            EmployerAmount: unifiedgosdk.Pointer[float64](80510.0),
            EmployerType: shared.EmployerTypePercentage.ToPointer(),
            Frequency: shared.HrisDeductionFrequencyOneTime.ToPointer(),
            GarnishmentType: shared.GarnishmentTypeChildSupport.ToPointer(),
            ID: unifiedgosdk.Pointer("5f050ddd-074b-4cf5-b93f-446a07231be7"),
            IsActive: unifiedgosdk.Pointer(true),
            MaxAmount: unifiedgosdk.Pointer[float64](21794.0),
            Metadata: []shared.HrisMetadata{},
            Notes: unifiedgosdk.Pointer("Carmen desidero."),
            StartAt: types.MustNewTimeFromString("2025-03-19T19:07:11.588Z"),
            Type: shared.HrisDeductionTypeFixed.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-03-25T21:48:35.117Z"),
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

Creates a HrisDevice. Provide the unified HrisDevice fields in the request body; they are mapped to the provider's API. Returns the created HrisDevice.

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
            AssetTag: unifiedgosdk.Pointer("fJNacaaoWi"),
            CreatedAt: types.MustNewTimeFromString("2019-04-04T17:11:40.322Z"),
            HasAntivirus: unifiedgosdk.Pointer(true),
            HasFirewall: unifiedgosdk.Pointer(true),
            HasHdEncrypted: unifiedgosdk.Pointer(true),
            HasPasswordManager: unifiedgosdk.Pointer(true),
            HasScreenlock: unifiedgosdk.Pointer(false),
            ID: unifiedgosdk.Pointer("c323d451-7216-482e-8a89-422a24fb88d5"),
            IsMissing: unifiedgosdk.Pointer(false),
            Manufacturer: unifiedgosdk.Pointer("Hackett Inc"),
            Metadata: []shared.HrisMetadata{},
            Model: unifiedgosdk.Pointer("Fantastic"),
            Name: unifiedgosdk.Pointer("august-laptop"),
            Os: unifiedgosdk.Pointer("transmitter"),
            OsVersion: unifiedgosdk.Pointer("8.4.8"),
            UpdatedAt: types.MustNewTimeFromString("2023-06-11T13:27:27.349Z"),
            Version: unifiedgosdk.Pointer("4.15.7"),
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

Creates a HrisDocument. Provide the unified HrisDocument fields in the request body; they are mapped to the provider's API. Returns the created HrisDocument.

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
            ID: unifiedgosdk.Pointer("4dac01b1-b7f1-4afc-9215-27c3e06eb6f3"),
            Metadata: []shared.HrisMetadata{},
            Type: shared.HrisDocumentTypePolicy.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-10-15T22:53:39.062Z"),
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

## CreateHrisEarning

Creates a HrisEarning. Provide the unified HrisEarning fields in the request body; they are mapped to the provider's API. Returns the created HrisEarning.

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisEarning" method="post" path="/hris/{connection_id}/earning" example="hris_earning" -->
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

    res, err := s.Hris.CreateHrisEarning(ctx, operations.CreateHrisEarningRequest{
        HrisEarning: shared.HrisEarning{
            Amount: unifiedgosdk.Pointer[float64](4026.0),
            CreatedAt: types.MustNewTimeFromString("2019-01-26T12:44:52.326Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Frequency: shared.HrisEarningFrequencySemimonthly.ToPointer(),
            ID: unifiedgosdk.Pointer("558a1864-18dd-4df0-b98f-c7999cedb645"),
            Metadata: []shared.HrisMetadata{},
            Notes: unifiedgosdk.Pointer("Deporto truculenter vinum beatae."),
            StartAt: types.MustNewTimeFromString("2022-07-28T03:33:50.064Z"),
            Status: shared.HrisEarningStatusEnded.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2020-02-09T15:27:00.044Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisEarning != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateHrisEarningRequest](../../pkg/models/operations/createhrisearningrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateHrisEarningResponse](../../pkg/models/operations/createhrisearningresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateHrisEmployee

Creates a HrisEmployee. Provide the unified HrisEmployee fields in the request body; they are mapped to the provider's API. Returns the created HrisEmployee.

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
                Address1: unifiedgosdk.Pointer("7431 Hintz Parkway"),
                Address2: unifiedgosdk.Pointer("Apt. 520"),
                City: unifiedgosdk.Pointer("Port Rogerview"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("01711"),
                Region: unifiedgosdk.Pointer("Missouri"),
                RegionCode: unifiedgosdk.Pointer("TN"),
            },
            Bio: unifiedgosdk.Pointer("fun advocate  🇱🇻"),
            Compensation: []shared.HrisCompensation{
                shared.HrisCompensation{
                    Amount: unifiedgosdk.Pointer[float64](95000.0),
                    Currency: unifiedgosdk.Pointer("USD"),
                    Frequency: shared.HrisCompensationFrequencyYear.ToPointer(),
                    Notes: unifiedgosdk.Pointer("Eius."),
                    Type: shared.HrisCompensationTypeOther.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2019-09-16T15:08:53.262Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DateOfBirth: types.MustNewTimeFromString("2001-04-22"),
            Emails: []shared.HrisEmail{
                shared.HrisEmail{
                    Email: "zetta.prohaska@insidious-stump.net",
                    Type: shared.HrisEmailTypeWork.ToPointer(),
                },
            },
            EmployeeNumber: unifiedgosdk.Pointer("uCk3qaL1ei"),
            EmploymentStatus: shared.EmploymentStatusActive.ToPointer(),
            EmploymentType: shared.HrisEmployeeEmploymentTypePartTime.ToPointer(),
            FirstName: unifiedgosdk.Pointer("Zetta"),
            FlsaStatus: shared.FlsaStatusExempt.ToPointer(),
            Gender: shared.HrisEmployeeGenderMale.ToPointer(),
            HasMfa: unifiedgosdk.Pointer(true),
            HiredAt: types.MustNewTimeFromString("2019-09-16T15:08:53.262Z"),
            ID: unifiedgosdk.Pointer("89ece618-3681-468a-8f17-b6da53e1d01b"),
            ImageURL: unifiedgosdk.Pointer("https://cdn.jsdelivr.net/gh/faker-js/assets-person-portrait/female/512/85.jpg"),
            LanguageLocale: unifiedgosdk.Pointer("in"),
            LastName: unifiedgosdk.Pointer("Prohaska"),
            Locations: []shared.HrisLocation{},
            MaritalStatus: shared.HrisEmployeeMaritalStatusMarried.ToPointer(),
            Metadata: []shared.HrisMetadata{
                shared.HrisMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateHrisMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.HrisMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("b870637f-d229-4b09-9ebe-d6c2baa4179f"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateHrisMetadataValueStr(
                        "tutis",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Zetta Prohaska"),
            PayFrequency: shared.PayFrequencyWeek.ToPointer(),
            Pronouns: unifiedgosdk.Pointer("she/her"),
            Relationships: []shared.HrisEmployeerelationship{
                shared.HrisEmployeerelationship{
                    Emails: []shared.HrisEmail{
                        shared.HrisEmail{
                            Email: "keith.prohaska@gmail.com",
                        },
                    },
                    Name: unifiedgosdk.Pointer("Keith Prohaska"),
                    Type: shared.HrisEmployeerelationshipTypeSibling.ToPointer(),
                },
                shared.HrisEmployeerelationship{
                    Emails: []shared.HrisEmail{
                        shared.HrisEmail{
                            Email: "keely.prohaska@icloud.com",
                        },
                    },
                    Name: unifiedgosdk.Pointer("Keely Prohaska"),
                    Type: shared.HrisEmployeerelationshipTypeOther.ToPointer(),
                },
            },
            Salutation: unifiedgosdk.Pointer("Ms."),
            SsnSin: unifiedgosdk.Pointer("284-82-8167"),
            StorageQuotaAllocated: unifiedgosdk.Pointer[float64](3653713.0),
            StorageQuotaAvailable: unifiedgosdk.Pointer[float64](3017275.0),
            StorageQuotaUsed: unifiedgosdk.Pointer[float64](367449.0),
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "(320) 980-1370",
                    Type: shared.HrisTelephoneTypeHome.ToPointer(),
                },
                shared.HrisTelephone{
                    Telephone: "(694) 840-0166",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
            },
            TimeoffDaysTotal: unifiedgosdk.Pointer[float64](17.0),
            TimeoffDaysUsed: unifiedgosdk.Pointer[float64](9.0),
            Timezone: unifiedgosdk.Pointer("Europe/Ljubljana"),
            Title: unifiedgosdk.Pointer("District Branding Specialist"),
            UpdatedAt: types.MustNewTimeFromString("2020-07-19T20:54:20.156Z"),
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

Creates a HrisGroup. Provide the unified HrisGroup fields in the request body; they are mapped to the provider's API. Returns the created HrisGroup.

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
            ID: unifiedgosdk.Pointer("54973609-2d36-4cf5-afb1-689db46fb98c"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("Games"),
            Type: shared.HrisGroupTypeBranch.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-05-27T00:47:53.770Z"),
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

## CreateHrisJob

Creates a HrisJob. Provide the unified HrisJob fields in the request body; they are mapped to the provider's API. Returns the created HrisJob.

### Example Usage

<!-- UsageSnippet language="go" operationID="createHrisJob" method="post" path="/hris/{connection_id}/job" example="hris_job" -->
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

    res, err := s.Hris.CreateHrisJob(ctx, operations.CreateHrisJobRequest{
        HrisJob: shared.HrisJob{
            Code: unifiedgosdk.Pointer("P-05547"),
            CreatedAt: types.MustNewTimeFromString("2020-10-28T01:13:19.491Z"),
            Description: unifiedgosdk.Pointer("Alienus timor depromo bonus vobis creator tabella callide cavus triumphus caveo."),
            EmploymentType: shared.HrisJobEmploymentTypeFreelance.ToPointer(),
            ID: unifiedgosdk.Pointer("ed280128-849e-4164-b4fc-cce1aa1f781a"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("Internal Integration Specialist"),
            UpdatedAt: types.MustNewTimeFromString("2023-02-01T10:33:57.646Z"),
        },
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateHrisJobRequest](../../pkg/models/operations/createhrisjobrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreateHrisJobResponse](../../pkg/models/operations/createhrisjobresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateHrisLocation

Creates a HrisLocation. Provide the unified HrisLocation fields in the request body; they are mapped to the provider's API. Returns the created HrisLocation.

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
                Address1: unifiedgosdk.Pointer("4418 Daniella Mall"),
                Address2: unifiedgosdk.Pointer("Apt. 997"),
                City: unifiedgosdk.Pointer("Burke"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("54109"),
                Region: unifiedgosdk.Pointer("Ohio"),
                RegionCode: unifiedgosdk.Pointer("OR"),
            },
            CreatedAt: types.MustNewTimeFromString("2021-07-18T10:32:01.414Z"),
            Currency: unifiedgosdk.Pointer("MNT"),
            Description: unifiedgosdk.Pointer("Commemoro acervus."),
            ID: unifiedgosdk.Pointer("829cfb3e-efbb-4234-b106-a223403d4774"),
            IsActive: unifiedgosdk.Pointer(false),
            IsHq: unifiedgosdk.Pointer(true),
            LanguageLocale: unifiedgosdk.Pointer("it"),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("North Kimberlyland Office"),
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "(371) 245-0699",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
                shared.HrisTelephone{
                    Telephone: "(700) 845-5854",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
                shared.HrisTelephone{
                    Telephone: "(461) 437-3548",
                    Type: shared.HrisTelephoneTypeHome.ToPointer(),
                },
            },
            Timezone: unifiedgosdk.Pointer("Asia/Jerusalem"),
            UpdatedAt: types.MustNewTimeFromString("2023-06-23T06:05:23.455Z"),
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

Creates a HrisTaxonomy. Provide the unified HrisTaxonomy fields in the request body; they are mapped to the provider's API. Returns the created HrisTaxonomy.

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
            AlternativeNames: []string{
                "Facilitator",
                "National",
            },
            CreatedAt: types.MustNewTimeFromString("2022-06-23T02:10:00.789Z"),
            Description: unifiedgosdk.Pointer("Uterque vespillo adnuo ea commodo turbo suscipio condico."),
            Domain: unifiedgosdk.Pointer("Health"),
            ID: unifiedgosdk.Pointer("e655973b-8dbc-4d79-9274-423c0ff26cd8"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("International Functionality Architect"),
            Subcategory: unifiedgosdk.Pointer("Marble"),
            Type: shared.HrisTaxonomyTypeKnowledge.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-05-30T13:47:13.204Z"),
            URL: unifiedgosdk.Pointer("https://fixed-membership.biz/"),
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

Creates a HrisTimeoff. Provide the unified HrisTimeoff fields in the request body; they are mapped to the provider's API. Returns the created HrisTimeoff.

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
            ApprovedAt: types.MustNewTimeFromString("2022-02-23T19:51:05.756Z"),
            Comments: unifiedgosdk.Pointer("Blandior ventus curiositas amplitudo."),
            CreatedAt: types.MustNewTimeFromString("2021-10-06T18:00:20.615Z"),
            Duration: unifiedgosdk.Pointer[float64](4.0),
            DurationType: shared.DurationTypeDay.ToPointer(),
            EndAt: types.MustNewTimeFromString("2025-01-01T11:42:48.799Z"),
            ID: unifiedgosdk.Pointer("aa2a4a40-52ae-4acf-9866-16a22d4efafb"),
            IsPaid: unifiedgosdk.Pointer(true),
            Metadata: []shared.HrisMetadata{},
            OriginalType: unifiedgosdk.Pointer("acerbitas ut"),
            Reason: unifiedgosdk.Pointer("verto"),
            StartAt: types.MustNewTimeFromString("2023-09-07T00:56:10.762Z"),
            Status: shared.HrisTimeoffStatusDenied.ToPointer(),
            Type: shared.HrisTimeoffTypeInLieu.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-07-13T17:06:01.811Z"),
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

Creates a HrisTimeshift. Provide the unified HrisTimeshift fields in the request body; they are mapped to the provider's API. Returns the created HrisTimeshift.

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
            ApprovedAt: types.MustNewTimeFromString("2022-07-20T12:40:22.350Z"),
            Compensation: []shared.HrisCompensation{
                shared.HrisCompensation{
                    Amount: unifiedgosdk.Pointer[float64](76761.0),
                    Currency: unifiedgosdk.Pointer("GBP"),
                    Frequency: shared.HrisCompensationFrequencySemimonthly.ToPointer(),
                    Notes: unifiedgosdk.Pointer("Vehemens annus adficio."),
                    Type: shared.HrisCompensationTypeCommission.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2019-07-01T23:53:15.738Z"),
            EmployeeUserID: "<id>",
            EndAt: types.MustNewTimeFromString("2026-10-02T21:17:54.773Z"),
            Hours: unifiedgosdk.Pointer[float64](6.0),
            ID: unifiedgosdk.Pointer("9dffcf97-2f6c-4ef0-a6f8-57085913d99b"),
            IsApproved: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            StartAt: types.MustNewTimeFromString("2023-07-16T06:22:10.551Z"),
            UpdatedAt: types.MustNewTimeFromString("2021-07-03T12:32:37.806Z"),
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

Retrieves a single HrisAttendance by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single HrisBankaccount by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single HrisBenefit by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single HrisCompany by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single HrisDeduction by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single HrisDevice by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single HrisDocument by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

## GetHrisEarning

Retrieves a single HrisEarning by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisEarning" method="get" path="/hris/{connection_id}/earning/{id}" -->
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

    res, err := s.Hris.GetHrisEarning(ctx, operations.GetHrisEarningRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisEarning != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetHrisEarningRequest](../../pkg/models/operations/gethrisearningrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetHrisEarningResponse](../../pkg/models/operations/gethrisearningresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisEmployee

Retrieves a single HrisEmployee by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single HrisGroup by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

## GetHrisJob

Retrieves a single HrisJob by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisJob" method="get" path="/hris/{connection_id}/job/{id}" -->
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

    res, err := s.Hris.GetHrisJob(ctx, operations.GetHrisJobRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetHrisJobRequest](../../pkg/models/operations/gethrisjobrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.GetHrisJobResponse](../../pkg/models/operations/gethrisjobresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisLocation

Retrieves a single HrisLocation by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

## GetHrisPaycode

Retrieves a single HrisPaycode by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisPaycode" method="get" path="/hris/{connection_id}/paycode/{id}" -->
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

    res, err := s.Hris.GetHrisPaycode(ctx, operations.GetHrisPaycodeRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisPaycode != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetHrisPaycodeRequest](../../pkg/models/operations/gethrispaycoderequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetHrisPaycodeResponse](../../pkg/models/operations/gethrispaycoderesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisPayroll

Retrieves a single HrisPayroll by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

### Example Usage

<!-- UsageSnippet language="go" operationID="getHrisPayroll" method="get" path="/hris/{connection_id}/payroll/{id}" -->
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

    res, err := s.Hris.GetHrisPayroll(ctx, operations.GetHrisPayrollRequest{
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisPayroll != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetHrisPayrollRequest](../../pkg/models/operations/gethrispayrollrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetHrisPayrollResponse](../../pkg/models/operations/gethrispayrollresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetHrisPayslip

Retrieves a single HrisPayslip by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single HrisTaxonomy by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single HrisTimeoff by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Retrieves a single HrisTimeshift by its `id`. Obtain `id` values from the corresponding list endpoint. Use `fields` to return only the fields you need.

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

Returns a list of HrisAttendances. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of HrisBankaccounts. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of HrisBenefits. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of HrisCompanies. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of HrisDeductions. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of HrisDevices. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of HrisDocuments. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

## ListHrisEarnings

Returns a list of HrisEarnings. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisEarnings" method="get" path="/hris/{connection_id}/earning" -->
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

    res, err := s.Hris.ListHrisEarnings(ctx, operations.ListHrisEarningsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisEarnings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListHrisEarningsRequest](../../pkg/models/operations/listhrisearningsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListHrisEarningsResponse](../../pkg/models/operations/listhrisearningsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisEmployees

Returns a list of HrisEmployees. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of HrisGroups. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

## ListHrisJobs

Returns a list of HrisJobs. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisJobs" method="get" path="/hris/{connection_id}/job" -->
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

    res, err := s.Hris.ListHrisJobs(ctx, operations.ListHrisJobsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisJobs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListHrisJobsRequest](../../pkg/models/operations/listhrisjobsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListHrisJobsResponse](../../pkg/models/operations/listhrisjobsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisLocations

Returns a list of HrisLocations. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

## ListHrisPaycodes

Returns a list of HrisPaycodes. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisPaycodes" method="get" path="/hris/{connection_id}/paycode" -->
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

    res, err := s.Hris.ListHrisPaycodes(ctx, operations.ListHrisPaycodesRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisPaycodes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListHrisPaycodesRequest](../../pkg/models/operations/listhrispaycodesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListHrisPaycodesResponse](../../pkg/models/operations/listhrispaycodesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisPayrolls

Returns a list of HrisPayrolls. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="listHrisPayrolls" method="get" path="/hris/{connection_id}/payroll" -->
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

    res, err := s.Hris.ListHrisPayrolls(ctx, operations.ListHrisPayrollsRequest{
        ConnectionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisPayrolls != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListHrisPayrollsRequest](../../pkg/models/operations/listhrispayrollsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListHrisPayrollsResponse](../../pkg/models/operations/listhrispayrollsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListHrisPayslips

Returns a list of HrisPayslips. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of HrisTaxonomies. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of HrisTimeoffs. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Returns a list of HrisTimeshifts. Narrow the results with the query parameters before paging, then page through them with `limit` and `offset` (the default and maximum page size depend on the integration). Use `fields` to return only the fields you need and reduce the response size. When the integration supports a total count, it is returned in the `X-Total-Count` response header. Use each item's `id` to retrieve, update, or delete it where those operations are supported.

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

Updates the HrisAttendance identified by `id`. Provide the unified HrisAttendance fields to change in the request body.

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
            ApprovedAt: types.MustNewTimeFromString("2021-09-19T17:03:59.193Z"),
            Breaks: []shared.HrisAttendanceBreak{
                shared.HrisAttendanceBreak{
                    DurationMinutes: unifiedgosdk.Pointer[float64](51.0),
                    EndAt: types.MustNewTimeFromString("2024-02-17T09:31:50.528Z"),
                    ID: unifiedgosdk.Pointer("76c7a528-f157-4b1f-9ef6-ddc9cc64d1a4"),
                    IsPaid: unifiedgosdk.Pointer(false),
                    Name: unifiedgosdk.Pointer("Lunch"),
                    StartAt: types.MustNewTimeFromString("2023-11-01T11:18:39.316Z"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2021-08-10T19:43:18.452Z"),
            Currency: unifiedgosdk.Pointer("EUR"),
            DeclaredTipsAmount: unifiedgosdk.Pointer[float64](54.0),
            EmployeeUserID: "<id>",
            EndAt: types.MustTimeFromString("2024-04-26T09:10:57.694Z"),
            HourlyRate: unifiedgosdk.Pointer[float64](16.0),
            Hours: unifiedgosdk.Pointer[float64](10.0),
            ID: unifiedgosdk.Pointer("82b699ea-bfa8-4895-baa3-755d8399d4b0"),
            JobName: unifiedgosdk.Pointer("Human Accountability Administrator"),
            Metadata: []shared.HrisMetadata{},
            NonCashTipsAmount: unifiedgosdk.Pointer[float64](167.0),
            StartAt: types.MustTimeFromString("2021-11-11T07:47:35.745Z"),
            Status: shared.HrisAttendanceStatusOpen.ToPointer(),
            Timezone: unifiedgosdk.Pointer("America/Atikokan"),
            UpdatedAt: types.MustNewTimeFromString("2022-01-20T09:07:50.278Z"),
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

Updates the HrisBankaccount identified by `id`. Provide the unified HrisBankaccount fields to change in the request body.

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
            CountryCode: unifiedgosdk.Pointer("US"),
            CreatedAt: types.MustNewTimeFromString("2019-11-16T16:43:45.976Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            ID: unifiedgosdk.Pointer("a559ba11-761b-4283-8929-af6e220eef89"),
            IsPrimary: unifiedgosdk.Pointer(true),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("Granville Schumm - Checking"),
            Priority: unifiedgosdk.Pointer[float64](2.0),
            RoutingNumber: unifiedgosdk.Pointer("448650724"),
            SplitAmount: unifiedgosdk.Pointer[float64](98.0),
            SplitType: shared.SplitTypePercentage.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-07-06T00:53:08.537Z"),
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

Updates the HrisBenefit identified by `id`. Provide the unified HrisBenefit fields to change in the request body.

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
            Currency: unifiedgosdk.Pointer("GBP"),
            Description: unifiedgosdk.Pointer("Vomito voluptas dolor sed."),
            EmployerContributionAmount: unifiedgosdk.Pointer[float64](185006.0),
            EmployerContributionMaxAmount: unifiedgosdk.Pointer[float64](179093.0),
            EmployerContributionType: shared.EmployerContributionTypePercentage.ToPointer(),
            Frequency: shared.HrisBenefitFrequencyMonth.ToPointer(),
            ID: unifiedgosdk.Pointer("2ae4763b-fe19-441e-a091-ee01829c21d8"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("Frozen Wooden Ball"),
            Tax: shared.TaxPreTax.ToPointer(),
            Type: shared.HrisBenefitTypeOther.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-03-23T09:52:37.646Z"),
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

Updates the HrisCompany identified by `id`. Provide the unified HrisCompany fields to change in the request body.

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
                Address1: unifiedgosdk.Pointer("675 Lavinia Lights"),
                City: unifiedgosdk.Pointer("Rocky Mount"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("67325"),
                Region: unifiedgosdk.Pointer("Kentucky"),
                RegionCode: unifiedgosdk.Pointer("MO"),
            },
            CreatedAt: types.MustNewTimeFromString("2021-05-02T22:27:38.970Z"),
            ID: unifiedgosdk.Pointer("b71cbe84-c3bd-4aaa-b501-0299223e0528"),
            LegalName: unifiedgosdk.Pointer("Gottlieb Group LLC"),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("Gottlieb Group"),
            UpdatedAt: types.MustNewTimeFromString("2026-10-14T12:37:49.357Z"),
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

Updates the HrisDeduction identified by `id`. Provide the unified HrisDeduction fields to change in the request body.

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
            Amount: unifiedgosdk.Pointer[float64](564.0),
            CoverageLevel: shared.HrisDeductionCoverageLevelEmployeeChild.ToPointer(),
            CreatedAt: types.MustNewTimeFromString("2020-02-05T01:46:31.384Z"),
            EmployerAmount: unifiedgosdk.Pointer[float64](80510.0),
            EmployerType: shared.EmployerTypePercentage.ToPointer(),
            Frequency: shared.HrisDeductionFrequencyOneTime.ToPointer(),
            GarnishmentType: shared.GarnishmentTypeChildSupport.ToPointer(),
            ID: unifiedgosdk.Pointer("bffff85a-7553-417b-9284-702cdd61fc38"),
            IsActive: unifiedgosdk.Pointer(true),
            MaxAmount: unifiedgosdk.Pointer[float64](21794.0),
            Metadata: []shared.HrisMetadata{},
            Notes: unifiedgosdk.Pointer("Carmen desidero."),
            StartAt: types.MustNewTimeFromString("2025-03-19T19:07:11.597Z"),
            Type: shared.HrisDeductionTypeFixed.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-03-25T21:48:35.125Z"),
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

Updates the HrisDevice identified by `id`. Provide the unified HrisDevice fields to change in the request body.

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
            AssetTag: unifiedgosdk.Pointer("fJNacaaoWi"),
            CreatedAt: types.MustNewTimeFromString("2019-04-04T17:11:40.322Z"),
            HasAntivirus: unifiedgosdk.Pointer(true),
            HasFirewall: unifiedgosdk.Pointer(true),
            HasHdEncrypted: unifiedgosdk.Pointer(true),
            HasPasswordManager: unifiedgosdk.Pointer(true),
            HasScreenlock: unifiedgosdk.Pointer(false),
            ID: unifiedgosdk.Pointer("13f60916-0831-4617-8e37-f03be78a881a"),
            IsMissing: unifiedgosdk.Pointer(false),
            Manufacturer: unifiedgosdk.Pointer("Hackett Inc"),
            Metadata: []shared.HrisMetadata{},
            Model: unifiedgosdk.Pointer("Fantastic"),
            Name: unifiedgosdk.Pointer("august-laptop"),
            Os: unifiedgosdk.Pointer("transmitter"),
            OsVersion: unifiedgosdk.Pointer("8.4.8"),
            UpdatedAt: types.MustNewTimeFromString("2023-06-11T13:27:27.355Z"),
            Version: unifiedgosdk.Pointer("4.15.7"),
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

Updates the HrisDocument identified by `id`. Provide the unified HrisDocument fields to change in the request body.

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
            ID: unifiedgosdk.Pointer("5994f5e3-45e2-4c3f-8cec-201b6d6d39d2"),
            Metadata: []shared.HrisMetadata{},
            Type: shared.HrisDocumentTypePolicy.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-10-15T22:53:39.069Z"),
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

## PatchHrisEarning

Updates the HrisEarning identified by `id`. Provide the unified HrisEarning fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisEarning" method="patch" path="/hris/{connection_id}/earning/{id}" example="hris_earning" -->
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

    res, err := s.Hris.PatchHrisEarning(ctx, operations.PatchHrisEarningRequest{
        HrisEarning: shared.HrisEarning{
            Amount: unifiedgosdk.Pointer[float64](4026.0),
            CreatedAt: types.MustNewTimeFromString("2019-01-26T12:44:52.326Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Frequency: shared.HrisEarningFrequencySemimonthly.ToPointer(),
            ID: unifiedgosdk.Pointer("988b6127-3491-45b6-8c35-2ea8ab6fea93"),
            Metadata: []shared.HrisMetadata{},
            Notes: unifiedgosdk.Pointer("Deporto truculenter vinum beatae."),
            StartAt: types.MustNewTimeFromString("2022-07-28T03:33:50.069Z"),
            Status: shared.HrisEarningStatusEnded.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2020-02-09T15:27:00.045Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisEarning != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PatchHrisEarningRequest](../../pkg/models/operations/patchhrisearningrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchHrisEarningResponse](../../pkg/models/operations/patchhrisearningresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisEmployee

Updates the HrisEmployee identified by `id`. Provide the unified HrisEmployee fields to change in the request body.

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
                Address1: unifiedgosdk.Pointer("7431 Hintz Parkway"),
                Address2: unifiedgosdk.Pointer("Apt. 520"),
                City: unifiedgosdk.Pointer("Port Rogerview"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("01711"),
                Region: unifiedgosdk.Pointer("Missouri"),
                RegionCode: unifiedgosdk.Pointer("TN"),
            },
            Bio: unifiedgosdk.Pointer("fun advocate  🇱🇻"),
            Compensation: []shared.HrisCompensation{
                shared.HrisCompensation{
                    Amount: unifiedgosdk.Pointer[float64](95000.0),
                    Currency: unifiedgosdk.Pointer("USD"),
                    Frequency: shared.HrisCompensationFrequencyYear.ToPointer(),
                    Notes: unifiedgosdk.Pointer("Eius."),
                    Type: shared.HrisCompensationTypeOther.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2019-09-16T15:08:53.262Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DateOfBirth: types.MustNewTimeFromString("2001-04-22"),
            Emails: []shared.HrisEmail{
                shared.HrisEmail{
                    Email: "zetta.prohaska@insidious-stump.net",
                    Type: shared.HrisEmailTypeWork.ToPointer(),
                },
            },
            EmployeeNumber: unifiedgosdk.Pointer("uCk3qaL1ei"),
            EmploymentStatus: shared.EmploymentStatusActive.ToPointer(),
            EmploymentType: shared.HrisEmployeeEmploymentTypePartTime.ToPointer(),
            FirstName: unifiedgosdk.Pointer("Zetta"),
            FlsaStatus: shared.FlsaStatusExempt.ToPointer(),
            Gender: shared.HrisEmployeeGenderMale.ToPointer(),
            HasMfa: unifiedgosdk.Pointer(true),
            HiredAt: types.MustNewTimeFromString("2019-09-16T15:08:53.262Z"),
            ID: unifiedgosdk.Pointer("ffeb673d-ef5e-4f46-bf0f-3d93b4b2ad4d"),
            ImageURL: unifiedgosdk.Pointer("https://cdn.jsdelivr.net/gh/faker-js/assets-person-portrait/female/512/85.jpg"),
            LanguageLocale: unifiedgosdk.Pointer("in"),
            LastName: unifiedgosdk.Pointer("Prohaska"),
            Locations: []shared.HrisLocation{},
            MaritalStatus: shared.HrisEmployeeMaritalStatusMarried.ToPointer(),
            Metadata: []shared.HrisMetadata{
                shared.HrisMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateHrisMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.HrisMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("5c436a00-8fde-4c33-b3a7-bfa83a5bad6b"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateHrisMetadataValueStr(
                        "tutis",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Zetta Prohaska"),
            PayFrequency: shared.PayFrequencyWeek.ToPointer(),
            Pronouns: unifiedgosdk.Pointer("she/her"),
            Relationships: []shared.HrisEmployeerelationship{
                shared.HrisEmployeerelationship{
                    Emails: []shared.HrisEmail{
                        shared.HrisEmail{
                            Email: "keith.prohaska@gmail.com",
                        },
                    },
                    Name: unifiedgosdk.Pointer("Keith Prohaska"),
                    Type: shared.HrisEmployeerelationshipTypeSibling.ToPointer(),
                },
                shared.HrisEmployeerelationship{
                    Emails: []shared.HrisEmail{
                        shared.HrisEmail{
                            Email: "keely.prohaska@icloud.com",
                        },
                    },
                    Name: unifiedgosdk.Pointer("Keely Prohaska"),
                    Type: shared.HrisEmployeerelationshipTypeOther.ToPointer(),
                },
            },
            Salutation: unifiedgosdk.Pointer("Ms."),
            SsnSin: unifiedgosdk.Pointer("284-82-8167"),
            StorageQuotaAllocated: unifiedgosdk.Pointer[float64](3653713.0),
            StorageQuotaAvailable: unifiedgosdk.Pointer[float64](3017275.0),
            StorageQuotaUsed: unifiedgosdk.Pointer[float64](367449.0),
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "(320) 980-1370",
                    Type: shared.HrisTelephoneTypeHome.ToPointer(),
                },
                shared.HrisTelephone{
                    Telephone: "(694) 840-0166",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
            },
            TimeoffDaysTotal: unifiedgosdk.Pointer[float64](17.0),
            TimeoffDaysUsed: unifiedgosdk.Pointer[float64](9.0),
            Timezone: unifiedgosdk.Pointer("Europe/Ljubljana"),
            Title: unifiedgosdk.Pointer("District Branding Specialist"),
            UpdatedAt: types.MustNewTimeFromString("2020-07-19T20:54:20.164Z"),
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

Updates the HrisGroup identified by `id`. Provide the unified HrisGroup fields to change in the request body.

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
            ID: unifiedgosdk.Pointer("7aec714c-595f-4e0c-add5-1b8ea084410e"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("Games"),
            Type: shared.HrisGroupTypeBranch.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-05-27T00:47:53.779Z"),
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

## PatchHrisJob

Updates the HrisJob identified by `id`. Provide the unified HrisJob fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchHrisJob" method="patch" path="/hris/{connection_id}/job/{id}" example="hris_job" -->
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

    res, err := s.Hris.PatchHrisJob(ctx, operations.PatchHrisJobRequest{
        HrisJob: shared.HrisJob{
            Code: unifiedgosdk.Pointer("P-05547"),
            CreatedAt: types.MustNewTimeFromString("2020-10-28T01:13:19.491Z"),
            Description: unifiedgosdk.Pointer("Alienus timor depromo bonus vobis creator tabella callide cavus triumphus caveo."),
            EmploymentType: shared.HrisJobEmploymentTypeFreelance.ToPointer(),
            ID: unifiedgosdk.Pointer("64094bac-b5c3-4656-af73-b18cfc8bc976"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("Internal Integration Specialist"),
            UpdatedAt: types.MustNewTimeFromString("2023-02-01T10:33:57.685Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.PatchHrisJobRequest](../../pkg/models/operations/patchhrisjobrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.PatchHrisJobResponse](../../pkg/models/operations/patchhrisjobresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## PatchHrisLocation

Updates the HrisLocation identified by `id`. Provide the unified HrisLocation fields to change in the request body.

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
                Address1: unifiedgosdk.Pointer("4418 Daniella Mall"),
                Address2: unifiedgosdk.Pointer("Apt. 997"),
                City: unifiedgosdk.Pointer("Burke"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("54109"),
                Region: unifiedgosdk.Pointer("Ohio"),
                RegionCode: unifiedgosdk.Pointer("OR"),
            },
            CreatedAt: types.MustNewTimeFromString("2021-07-18T10:32:01.414Z"),
            Currency: unifiedgosdk.Pointer("MNT"),
            Description: unifiedgosdk.Pointer("Commemoro acervus."),
            ID: unifiedgosdk.Pointer("6e8f5955-b4c4-4ac5-b0a2-9db1c851c759"),
            IsActive: unifiedgosdk.Pointer(false),
            IsHq: unifiedgosdk.Pointer(true),
            LanguageLocale: unifiedgosdk.Pointer("it"),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("North Kimberlyland Office"),
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "(371) 245-0699",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
                shared.HrisTelephone{
                    Telephone: "(700) 845-5854",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
                shared.HrisTelephone{
                    Telephone: "(461) 437-3548",
                    Type: shared.HrisTelephoneTypeHome.ToPointer(),
                },
            },
            Timezone: unifiedgosdk.Pointer("Asia/Jerusalem"),
            UpdatedAt: types.MustNewTimeFromString("2023-06-23T06:05:23.461Z"),
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

Updates the HrisTimeoff identified by `id`. Provide the unified HrisTimeoff fields to change in the request body.

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
            ApprovedAt: types.MustNewTimeFromString("2022-02-23T19:51:05.757Z"),
            Comments: unifiedgosdk.Pointer("Blandior ventus curiositas amplitudo."),
            CreatedAt: types.MustNewTimeFromString("2021-10-06T18:00:20.615Z"),
            Duration: unifiedgosdk.Pointer[float64](4.0),
            DurationType: shared.DurationTypeDay.ToPointer(),
            EndAt: types.MustNewTimeFromString("2025-01-01T11:42:48.805Z"),
            ID: unifiedgosdk.Pointer("ff3b9aee-5865-41e3-b663-1ea749295508"),
            IsPaid: unifiedgosdk.Pointer(true),
            Metadata: []shared.HrisMetadata{},
            OriginalType: unifiedgosdk.Pointer("acerbitas ut"),
            Reason: unifiedgosdk.Pointer("verto"),
            StartAt: types.MustNewTimeFromString("2023-09-07T00:56:10.766Z"),
            Status: shared.HrisTimeoffStatusDenied.ToPointer(),
            Type: shared.HrisTimeoffTypeInLieu.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-07-13T17:06:01.813Z"),
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

Updates the HrisTimeshift identified by `id`. Provide the unified HrisTimeshift fields to change in the request body.

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
            ApprovedAt: types.MustNewTimeFromString("2022-07-20T12:40:22.355Z"),
            Compensation: []shared.HrisCompensation{
                shared.HrisCompensation{
                    Amount: unifiedgosdk.Pointer[float64](76761.0),
                    Currency: unifiedgosdk.Pointer("GBP"),
                    Frequency: shared.HrisCompensationFrequencySemimonthly.ToPointer(),
                    Notes: unifiedgosdk.Pointer("Vehemens annus adficio."),
                    Type: shared.HrisCompensationTypeCommission.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2019-07-01T23:53:15.738Z"),
            EmployeeUserID: "<id>",
            EndAt: types.MustNewTimeFromString("2026-10-02T21:17:54.785Z"),
            Hours: unifiedgosdk.Pointer[float64](6.0),
            ID: unifiedgosdk.Pointer("b1df487a-b3a9-472a-97de-e5d4a980cdd9"),
            IsApproved: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            StartAt: types.MustNewTimeFromString("2023-07-16T06:22:10.558Z"),
            UpdatedAt: types.MustNewTimeFromString("2021-07-03T12:32:37.810Z"),
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

Deletes the HrisAttendance identified by `id`.

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

Deletes the HrisBankaccount identified by `id`.

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

Deletes the HrisBenefit identified by `id`.

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

Deletes the HrisCompany identified by `id`.

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

Deletes the HrisDeduction identified by `id`.

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

Deletes the HrisDevice identified by `id`.

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

Deletes the HrisDocument identified by `id`.

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

## RemoveHrisEarning

Deletes the HrisEarning identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisEarning" method="delete" path="/hris/{connection_id}/earning/{id}" -->
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

    res, err := s.Hris.RemoveHrisEarning(ctx, operations.RemoveHrisEarningRequest{
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
| `request`                                                                                      | [operations.RemoveHrisEarningRequest](../../pkg/models/operations/removehrisearningrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.RemoveHrisEarningResponse](../../pkg/models/operations/removehrisearningresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisEmployee

Deletes the HrisEmployee identified by `id`.

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

Deletes the HrisGroup identified by `id`.

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

## RemoveHrisJob

Deletes the HrisJob identified by `id`.

### Example Usage

<!-- UsageSnippet language="go" operationID="removeHrisJob" method="delete" path="/hris/{connection_id}/job/{id}" -->
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

    res, err := s.Hris.RemoveHrisJob(ctx, operations.RemoveHrisJobRequest{
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.RemoveHrisJobRequest](../../pkg/models/operations/removehrisjobrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.RemoveHrisJobResponse](../../pkg/models/operations/removehrisjobresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveHrisLocation

Deletes the HrisLocation identified by `id`.

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

Deletes the HrisTimeoff identified by `id`.

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

Deletes the HrisTimeshift identified by `id`.

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

Updates the HrisAttendance identified by `id`. Provide the unified HrisAttendance fields to change in the request body.

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
            ApprovedAt: types.MustNewTimeFromString("2021-09-19T17:03:59.193Z"),
            Breaks: []shared.HrisAttendanceBreak{
                shared.HrisAttendanceBreak{
                    DurationMinutes: unifiedgosdk.Pointer[float64](51.0),
                    EndAt: types.MustNewTimeFromString("2024-02-17T09:31:50.528Z"),
                    ID: unifiedgosdk.Pointer("76c7a528-f157-4b1f-9ef6-ddc9cc64d1a4"),
                    IsPaid: unifiedgosdk.Pointer(false),
                    Name: unifiedgosdk.Pointer("Lunch"),
                    StartAt: types.MustNewTimeFromString("2023-11-01T11:18:39.316Z"),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2021-08-10T19:43:18.452Z"),
            Currency: unifiedgosdk.Pointer("EUR"),
            DeclaredTipsAmount: unifiedgosdk.Pointer[float64](54.0),
            EmployeeUserID: "<id>",
            EndAt: types.MustTimeFromString("2024-04-26T09:10:57.694Z"),
            HourlyRate: unifiedgosdk.Pointer[float64](16.0),
            Hours: unifiedgosdk.Pointer[float64](10.0),
            ID: unifiedgosdk.Pointer("82b699ea-bfa8-4895-baa3-755d8399d4b0"),
            JobName: unifiedgosdk.Pointer("Human Accountability Administrator"),
            Metadata: []shared.HrisMetadata{},
            NonCashTipsAmount: unifiedgosdk.Pointer[float64](167.0),
            StartAt: types.MustTimeFromString("2021-11-11T07:47:35.745Z"),
            Status: shared.HrisAttendanceStatusOpen.ToPointer(),
            Timezone: unifiedgosdk.Pointer("America/Atikokan"),
            UpdatedAt: types.MustNewTimeFromString("2022-01-20T09:07:50.278Z"),
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

Updates the HrisBankaccount identified by `id`. Provide the unified HrisBankaccount fields to change in the request body.

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
            CountryCode: unifiedgosdk.Pointer("US"),
            CreatedAt: types.MustNewTimeFromString("2019-11-16T16:43:45.976Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            ID: unifiedgosdk.Pointer("a559ba11-761b-4283-8929-af6e220eef89"),
            IsPrimary: unifiedgosdk.Pointer(true),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("Granville Schumm - Checking"),
            Priority: unifiedgosdk.Pointer[float64](2.0),
            RoutingNumber: unifiedgosdk.Pointer("448650724"),
            SplitAmount: unifiedgosdk.Pointer[float64](98.0),
            SplitType: shared.SplitTypePercentage.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-07-06T00:53:08.537Z"),
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

Updates the HrisBenefit identified by `id`. Provide the unified HrisBenefit fields to change in the request body.

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
            Currency: unifiedgosdk.Pointer("GBP"),
            Description: unifiedgosdk.Pointer("Vomito voluptas dolor sed."),
            EmployerContributionAmount: unifiedgosdk.Pointer[float64](185006.0),
            EmployerContributionMaxAmount: unifiedgosdk.Pointer[float64](179093.0),
            EmployerContributionType: shared.EmployerContributionTypePercentage.ToPointer(),
            Frequency: shared.HrisBenefitFrequencyMonth.ToPointer(),
            ID: unifiedgosdk.Pointer("2ae4763b-fe19-441e-a091-ee01829c21d8"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("Frozen Wooden Ball"),
            Tax: shared.TaxPreTax.ToPointer(),
            Type: shared.HrisBenefitTypeOther.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2023-03-23T09:52:37.646Z"),
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

Updates the HrisCompany identified by `id`. Provide the unified HrisCompany fields to change in the request body.

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
                Address1: unifiedgosdk.Pointer("675 Lavinia Lights"),
                City: unifiedgosdk.Pointer("Rocky Mount"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("67325"),
                Region: unifiedgosdk.Pointer("Kentucky"),
                RegionCode: unifiedgosdk.Pointer("MO"),
            },
            CreatedAt: types.MustNewTimeFromString("2021-05-02T22:27:38.970Z"),
            ID: unifiedgosdk.Pointer("b71cbe84-c3bd-4aaa-b501-0299223e0528"),
            LegalName: unifiedgosdk.Pointer("Gottlieb Group LLC"),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("Gottlieb Group"),
            UpdatedAt: types.MustNewTimeFromString("2026-10-14T12:37:49.357Z"),
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

Updates the HrisDeduction identified by `id`. Provide the unified HrisDeduction fields to change in the request body.

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
            Amount: unifiedgosdk.Pointer[float64](564.0),
            CoverageLevel: shared.HrisDeductionCoverageLevelEmployeeChild.ToPointer(),
            CreatedAt: types.MustNewTimeFromString("2020-02-05T01:46:31.384Z"),
            EmployerAmount: unifiedgosdk.Pointer[float64](80510.0),
            EmployerType: shared.EmployerTypePercentage.ToPointer(),
            Frequency: shared.HrisDeductionFrequencyOneTime.ToPointer(),
            GarnishmentType: shared.GarnishmentTypeChildSupport.ToPointer(),
            ID: unifiedgosdk.Pointer("bffff85a-7553-417b-9284-702cdd61fc38"),
            IsActive: unifiedgosdk.Pointer(true),
            MaxAmount: unifiedgosdk.Pointer[float64](21794.0),
            Metadata: []shared.HrisMetadata{},
            Notes: unifiedgosdk.Pointer("Carmen desidero."),
            StartAt: types.MustNewTimeFromString("2025-03-19T19:07:11.597Z"),
            Type: shared.HrisDeductionTypeFixed.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2024-03-25T21:48:35.125Z"),
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

Updates the HrisDevice identified by `id`. Provide the unified HrisDevice fields to change in the request body.

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
            AssetTag: unifiedgosdk.Pointer("fJNacaaoWi"),
            CreatedAt: types.MustNewTimeFromString("2019-04-04T17:11:40.322Z"),
            HasAntivirus: unifiedgosdk.Pointer(true),
            HasFirewall: unifiedgosdk.Pointer(true),
            HasHdEncrypted: unifiedgosdk.Pointer(true),
            HasPasswordManager: unifiedgosdk.Pointer(true),
            HasScreenlock: unifiedgosdk.Pointer(false),
            ID: unifiedgosdk.Pointer("13f60916-0831-4617-8e37-f03be78a881a"),
            IsMissing: unifiedgosdk.Pointer(false),
            Manufacturer: unifiedgosdk.Pointer("Hackett Inc"),
            Metadata: []shared.HrisMetadata{},
            Model: unifiedgosdk.Pointer("Fantastic"),
            Name: unifiedgosdk.Pointer("august-laptop"),
            Os: unifiedgosdk.Pointer("transmitter"),
            OsVersion: unifiedgosdk.Pointer("8.4.8"),
            UpdatedAt: types.MustNewTimeFromString("2023-06-11T13:27:27.355Z"),
            Version: unifiedgosdk.Pointer("4.15.7"),
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

Updates the HrisDocument identified by `id`. Provide the unified HrisDocument fields to change in the request body.

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
            ID: unifiedgosdk.Pointer("5994f5e3-45e2-4c3f-8cec-201b6d6d39d2"),
            Metadata: []shared.HrisMetadata{},
            Type: shared.HrisDocumentTypePolicy.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2025-10-15T22:53:39.069Z"),
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

## UpdateHrisEarning

Updates the HrisEarning identified by `id`. Provide the unified HrisEarning fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisEarning" method="put" path="/hris/{connection_id}/earning/{id}" example="hris_earning" -->
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

    res, err := s.Hris.UpdateHrisEarning(ctx, operations.UpdateHrisEarningRequest{
        HrisEarning: shared.HrisEarning{
            Amount: unifiedgosdk.Pointer[float64](4026.0),
            CreatedAt: types.MustNewTimeFromString("2019-01-26T12:44:52.326Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            Frequency: shared.HrisEarningFrequencySemimonthly.ToPointer(),
            ID: unifiedgosdk.Pointer("988b6127-3491-45b6-8c35-2ea8ab6fea93"),
            Metadata: []shared.HrisMetadata{},
            Notes: unifiedgosdk.Pointer("Deporto truculenter vinum beatae."),
            StartAt: types.MustNewTimeFromString("2022-07-28T03:33:50.069Z"),
            Status: shared.HrisEarningStatusEnded.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2020-02-09T15:27:00.045Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisEarning != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateHrisEarningRequest](../../pkg/models/operations/updatehrisearningrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateHrisEarningResponse](../../pkg/models/operations/updatehrisearningresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisEmployee

Updates the HrisEmployee identified by `id`. Provide the unified HrisEmployee fields to change in the request body.

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
                Address1: unifiedgosdk.Pointer("7431 Hintz Parkway"),
                Address2: unifiedgosdk.Pointer("Apt. 520"),
                City: unifiedgosdk.Pointer("Port Rogerview"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("01711"),
                Region: unifiedgosdk.Pointer("Missouri"),
                RegionCode: unifiedgosdk.Pointer("TN"),
            },
            Bio: unifiedgosdk.Pointer("fun advocate  🇱🇻"),
            Compensation: []shared.HrisCompensation{
                shared.HrisCompensation{
                    Amount: unifiedgosdk.Pointer[float64](95000.0),
                    Currency: unifiedgosdk.Pointer("USD"),
                    Frequency: shared.HrisCompensationFrequencyYear.ToPointer(),
                    Notes: unifiedgosdk.Pointer("Eius."),
                    Type: shared.HrisCompensationTypeOther.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2019-09-16T15:08:53.262Z"),
            Currency: unifiedgosdk.Pointer("USD"),
            DateOfBirth: types.MustNewTimeFromString("2001-04-22"),
            Emails: []shared.HrisEmail{
                shared.HrisEmail{
                    Email: "zetta.prohaska@insidious-stump.net",
                    Type: shared.HrisEmailTypeWork.ToPointer(),
                },
            },
            EmployeeNumber: unifiedgosdk.Pointer("uCk3qaL1ei"),
            EmploymentStatus: shared.EmploymentStatusActive.ToPointer(),
            EmploymentType: shared.HrisEmployeeEmploymentTypePartTime.ToPointer(),
            FirstName: unifiedgosdk.Pointer("Zetta"),
            FlsaStatus: shared.FlsaStatusExempt.ToPointer(),
            Gender: shared.HrisEmployeeGenderMale.ToPointer(),
            HasMfa: unifiedgosdk.Pointer(true),
            HiredAt: types.MustNewTimeFromString("2019-09-16T15:08:53.262Z"),
            ID: unifiedgosdk.Pointer("ffeb673d-ef5e-4f46-bf0f-3d93b4b2ad4d"),
            ImageURL: unifiedgosdk.Pointer("https://cdn.jsdelivr.net/gh/faker-js/assets-person-portrait/female/512/85.jpg"),
            LanguageLocale: unifiedgosdk.Pointer("in"),
            LastName: unifiedgosdk.Pointer("Prohaska"),
            Locations: []shared.HrisLocation{},
            MaritalStatus: shared.HrisEmployeeMaritalStatusMarried.ToPointer(),
            Metadata: []shared.HrisMetadata{
                shared.HrisMetadata{
                    ExtraData: unifiedgosdk.Pointer(shared.CreateHrisMetadataExtraDataMapOfAny(
                        map[string]any{
                            "display_name": "Custom Property",
                        },
                    )),
                    Format: shared.HrisMetadataFormatText.ToPointer(),
                    ID: unifiedgosdk.Pointer("5c436a00-8fde-4c33-b3a7-bfa83a5bad6b"),
                    Namespace: unifiedgosdk.Pointer("custom"),
                    Slug: unifiedgosdk.Pointer("custom_property"),
                    Value: unifiedgosdk.Pointer(shared.CreateHrisMetadataValueStr(
                        "tutis",
                    )),
                },
            },
            Name: unifiedgosdk.Pointer("Zetta Prohaska"),
            PayFrequency: shared.PayFrequencyWeek.ToPointer(),
            Pronouns: unifiedgosdk.Pointer("she/her"),
            Relationships: []shared.HrisEmployeerelationship{
                shared.HrisEmployeerelationship{
                    Emails: []shared.HrisEmail{
                        shared.HrisEmail{
                            Email: "keith.prohaska@gmail.com",
                        },
                    },
                    Name: unifiedgosdk.Pointer("Keith Prohaska"),
                    Type: shared.HrisEmployeerelationshipTypeSibling.ToPointer(),
                },
                shared.HrisEmployeerelationship{
                    Emails: []shared.HrisEmail{
                        shared.HrisEmail{
                            Email: "keely.prohaska@icloud.com",
                        },
                    },
                    Name: unifiedgosdk.Pointer("Keely Prohaska"),
                    Type: shared.HrisEmployeerelationshipTypeOther.ToPointer(),
                },
            },
            Salutation: unifiedgosdk.Pointer("Ms."),
            SsnSin: unifiedgosdk.Pointer("284-82-8167"),
            StorageQuotaAllocated: unifiedgosdk.Pointer[float64](3653713.0),
            StorageQuotaAvailable: unifiedgosdk.Pointer[float64](3017275.0),
            StorageQuotaUsed: unifiedgosdk.Pointer[float64](367449.0),
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "(320) 980-1370",
                    Type: shared.HrisTelephoneTypeHome.ToPointer(),
                },
                shared.HrisTelephone{
                    Telephone: "(694) 840-0166",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
            },
            TimeoffDaysTotal: unifiedgosdk.Pointer[float64](17.0),
            TimeoffDaysUsed: unifiedgosdk.Pointer[float64](9.0),
            Timezone: unifiedgosdk.Pointer("Europe/Ljubljana"),
            Title: unifiedgosdk.Pointer("District Branding Specialist"),
            UpdatedAt: types.MustNewTimeFromString("2020-07-19T20:54:20.164Z"),
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

Updates the HrisGroup identified by `id`. Provide the unified HrisGroup fields to change in the request body.

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
            ID: unifiedgosdk.Pointer("7aec714c-595f-4e0c-add5-1b8ea084410e"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("Games"),
            Type: shared.HrisGroupTypeBranch.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2026-05-27T00:47:53.779Z"),
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

## UpdateHrisJob

Updates the HrisJob identified by `id`. Provide the unified HrisJob fields to change in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateHrisJob" method="put" path="/hris/{connection_id}/job/{id}" example="hris_job" -->
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

    res, err := s.Hris.UpdateHrisJob(ctx, operations.UpdateHrisJobRequest{
        HrisJob: shared.HrisJob{
            Code: unifiedgosdk.Pointer("P-05547"),
            CreatedAt: types.MustNewTimeFromString("2020-10-28T01:13:19.491Z"),
            Description: unifiedgosdk.Pointer("Alienus timor depromo bonus vobis creator tabella callide cavus triumphus caveo."),
            EmploymentType: shared.HrisJobEmploymentTypeFreelance.ToPointer(),
            ID: unifiedgosdk.Pointer("64094bac-b5c3-4656-af73-b18cfc8bc976"),
            IsActive: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("Internal Integration Specialist"),
            UpdatedAt: types.MustNewTimeFromString("2023-02-01T10:33:57.685Z"),
        },
        ConnectionID: "<id>",
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.HrisJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateHrisJobRequest](../../pkg/models/operations/updatehrisjobrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.UpdateHrisJobResponse](../../pkg/models/operations/updatehrisjobresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateHrisLocation

Updates the HrisLocation identified by `id`. Provide the unified HrisLocation fields to change in the request body.

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
                Address1: unifiedgosdk.Pointer("4418 Daniella Mall"),
                Address2: unifiedgosdk.Pointer("Apt. 997"),
                City: unifiedgosdk.Pointer("Burke"),
                CountryCode: unifiedgosdk.Pointer("US"),
                PostalCode: unifiedgosdk.Pointer("54109"),
                Region: unifiedgosdk.Pointer("Ohio"),
                RegionCode: unifiedgosdk.Pointer("OR"),
            },
            CreatedAt: types.MustNewTimeFromString("2021-07-18T10:32:01.414Z"),
            Currency: unifiedgosdk.Pointer("MNT"),
            Description: unifiedgosdk.Pointer("Commemoro acervus."),
            ID: unifiedgosdk.Pointer("6e8f5955-b4c4-4ac5-b0a2-9db1c851c759"),
            IsActive: unifiedgosdk.Pointer(false),
            IsHq: unifiedgosdk.Pointer(true),
            LanguageLocale: unifiedgosdk.Pointer("it"),
            Metadata: []shared.HrisMetadata{},
            Name: unifiedgosdk.Pointer("North Kimberlyland Office"),
            Telephones: []shared.HrisTelephone{
                shared.HrisTelephone{
                    Telephone: "(371) 245-0699",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
                shared.HrisTelephone{
                    Telephone: "(700) 845-5854",
                    Type: shared.HrisTelephoneTypeFax.ToPointer(),
                },
                shared.HrisTelephone{
                    Telephone: "(461) 437-3548",
                    Type: shared.HrisTelephoneTypeHome.ToPointer(),
                },
            },
            Timezone: unifiedgosdk.Pointer("Asia/Jerusalem"),
            UpdatedAt: types.MustNewTimeFromString("2023-06-23T06:05:23.461Z"),
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

Updates the HrisTimeoff identified by `id`. Provide the unified HrisTimeoff fields to change in the request body.

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
            ApprovedAt: types.MustNewTimeFromString("2022-02-23T19:51:05.757Z"),
            Comments: unifiedgosdk.Pointer("Blandior ventus curiositas amplitudo."),
            CreatedAt: types.MustNewTimeFromString("2021-10-06T18:00:20.615Z"),
            Duration: unifiedgosdk.Pointer[float64](4.0),
            DurationType: shared.DurationTypeDay.ToPointer(),
            EndAt: types.MustNewTimeFromString("2025-01-01T11:42:48.805Z"),
            ID: unifiedgosdk.Pointer("ff3b9aee-5865-41e3-b663-1ea749295508"),
            IsPaid: unifiedgosdk.Pointer(true),
            Metadata: []shared.HrisMetadata{},
            OriginalType: unifiedgosdk.Pointer("acerbitas ut"),
            Reason: unifiedgosdk.Pointer("verto"),
            StartAt: types.MustNewTimeFromString("2023-09-07T00:56:10.766Z"),
            Status: shared.HrisTimeoffStatusDenied.ToPointer(),
            Type: shared.HrisTimeoffTypeInLieu.ToPointer(),
            UpdatedAt: types.MustNewTimeFromString("2022-07-13T17:06:01.813Z"),
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

Updates the HrisTimeshift identified by `id`. Provide the unified HrisTimeshift fields to change in the request body.

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
            ApprovedAt: types.MustNewTimeFromString("2022-07-20T12:40:22.355Z"),
            Compensation: []shared.HrisCompensation{
                shared.HrisCompensation{
                    Amount: unifiedgosdk.Pointer[float64](76761.0),
                    Currency: unifiedgosdk.Pointer("GBP"),
                    Frequency: shared.HrisCompensationFrequencySemimonthly.ToPointer(),
                    Notes: unifiedgosdk.Pointer("Vehemens annus adficio."),
                    Type: shared.HrisCompensationTypeCommission.ToPointer(),
                },
            },
            CreatedAt: types.MustNewTimeFromString("2019-07-01T23:53:15.738Z"),
            EmployeeUserID: "<id>",
            EndAt: types.MustNewTimeFromString("2026-10-02T21:17:54.785Z"),
            Hours: unifiedgosdk.Pointer[float64](6.0),
            ID: unifiedgosdk.Pointer("b1df487a-b3a9-472a-97de-e5d4a980cdd9"),
            IsApproved: unifiedgosdk.Pointer(false),
            Metadata: []shared.HrisMetadata{},
            StartAt: types.MustNewTimeFromString("2023-07-16T06:22:10.558Z"),
            UpdatedAt: types.MustNewTimeFromString("2021-07-03T12:32:37.810Z"),
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