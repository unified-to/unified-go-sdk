# HealthcareMetadata5


## Supported Types

### HealthcareMetadata1

```go
healthcareMetadata5 := shared.CreateHealthcareMetadata5HealthcareMetadata1(shared.HealthcareMetadata1{/* values here */})
```

### 

```go
healthcareMetadata5 := shared.CreateHealthcareMetadata5Str(string{/* values here */})
```

### 

```go
healthcareMetadata5 := shared.CreateHealthcareMetadata5Number(float64{/* values here */})
```

### 

```go
healthcareMetadata5 := shared.CreateHealthcareMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch healthcareMetadata5.Type {
	case shared.HealthcareMetadata5TypeHealthcareMetadata1:
		// healthcareMetadata5.HealthcareMetadata1 is populated
	case shared.HealthcareMetadata5TypeStr:
		// healthcareMetadata5.Str is populated
	case shared.HealthcareMetadata5TypeNumber:
		// healthcareMetadata5.Number is populated
	case shared.HealthcareMetadata5TypeBoolean:
		// healthcareMetadata5.Boolean is populated
}
```
