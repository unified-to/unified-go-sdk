# HealthcareMetadataValue

The custom field value


## Supported Types

### 

```go
healthcareMetadataValue := shared.CreateHealthcareMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
healthcareMetadataValue := shared.CreateHealthcareMetadataValueStr(string{/* values here */})
```

### 

```go
healthcareMetadataValue := shared.CreateHealthcareMetadataValueNumber(float64{/* values here */})
```

### 

```go
healthcareMetadataValue := shared.CreateHealthcareMetadataValueBoolean(bool{/* values here */})
```

### 

```go
healthcareMetadataValue := shared.CreateHealthcareMetadataValueArrayOfHealthcareMetadataSchemas5([]shared.HealthcareMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch healthcareMetadataValue.Type {
	case shared.HealthcareMetadataValueTypeMapOfAny:
		// healthcareMetadataValue.MapOfAny is populated
	case shared.HealthcareMetadataValueTypeStr:
		// healthcareMetadataValue.Str is populated
	case shared.HealthcareMetadataValueTypeNumber:
		// healthcareMetadataValue.Number is populated
	case shared.HealthcareMetadataValueTypeBoolean:
		// healthcareMetadataValue.Boolean is populated
	case shared.HealthcareMetadataValueTypeArrayOfHealthcareMetadataSchemas5:
		// healthcareMetadataValue.ArrayOfHealthcareMetadataSchemas5 is populated
}
```
