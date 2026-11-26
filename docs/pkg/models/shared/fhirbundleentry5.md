# FhirBundleEntry5


## Supported Types

### FhirBundleEntry1

```go
fhirBundleEntry5 := shared.CreateFhirBundleEntry5FhirBundleEntry1(shared.FhirBundleEntry1{/* values here */})
```

### 

```go
fhirBundleEntry5 := shared.CreateFhirBundleEntry5Str(string{/* values here */})
```

### 

```go
fhirBundleEntry5 := shared.CreateFhirBundleEntry5Number(float64{/* values here */})
```

### 

```go
fhirBundleEntry5 := shared.CreateFhirBundleEntry5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch fhirBundleEntry5.Type {
	case shared.FhirBundleEntry5TypeFhirBundleEntry1:
		// fhirBundleEntry5.FhirBundleEntry1 is populated
	case shared.FhirBundleEntry5TypeStr:
		// fhirBundleEntry5.Str is populated
	case shared.FhirBundleEntry5TypeNumber:
		// fhirBundleEntry5.Number is populated
	case shared.FhirBundleEntry5TypeBoolean:
		// fhirBundleEntry5.Boolean is populated
}
```
