# HealthcareMetadataSchemas5


## Supported Types

### HealthcareMetadataSchemas1

```go
healthcareMetadataSchemas5 := shared.CreateHealthcareMetadataSchemas5HealthcareMetadataSchemas1(shared.HealthcareMetadataSchemas1{/* values here */})
```

### 

```go
healthcareMetadataSchemas5 := shared.CreateHealthcareMetadataSchemas5Str(string{/* values here */})
```

### 

```go
healthcareMetadataSchemas5 := shared.CreateHealthcareMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
healthcareMetadataSchemas5 := shared.CreateHealthcareMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch healthcareMetadataSchemas5.Type {
	case shared.HealthcareMetadataSchemas5TypeHealthcareMetadataSchemas1:
		// healthcareMetadataSchemas5.HealthcareMetadataSchemas1 is populated
	case shared.HealthcareMetadataSchemas5TypeStr:
		// healthcareMetadataSchemas5.Str is populated
	case shared.HealthcareMetadataSchemas5TypeNumber:
		// healthcareMetadataSchemas5.Number is populated
	case shared.HealthcareMetadataSchemas5TypeBoolean:
		// healthcareMetadataSchemas5.Boolean is populated
}
```
