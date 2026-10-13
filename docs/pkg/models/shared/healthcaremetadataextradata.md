# HealthcareMetadataExtraData

Additional data about the custom field


## Supported Types

### 

```go
healthcareMetadataExtraData := shared.CreateHealthcareMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
healthcareMetadataExtraData := shared.CreateHealthcareMetadataExtraDataStr(string{/* values here */})
```

### 

```go
healthcareMetadataExtraData := shared.CreateHealthcareMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
healthcareMetadataExtraData := shared.CreateHealthcareMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
healthcareMetadataExtraData := shared.CreateHealthcareMetadataExtraDataArrayOfHealthcareMetadata5([]shared.HealthcareMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch healthcareMetadataExtraData.Type {
	case shared.HealthcareMetadataExtraDataTypeMapOfAny:
		// healthcareMetadataExtraData.MapOfAny is populated
	case shared.HealthcareMetadataExtraDataTypeStr:
		// healthcareMetadataExtraData.Str is populated
	case shared.HealthcareMetadataExtraDataTypeNumber:
		// healthcareMetadataExtraData.Number is populated
	case shared.HealthcareMetadataExtraDataTypeBoolean:
		// healthcareMetadataExtraData.Boolean is populated
	case shared.HealthcareMetadataExtraDataTypeArrayOfHealthcareMetadata5:
		// healthcareMetadataExtraData.ArrayOfHealthcareMetadata5 is populated
}
```
