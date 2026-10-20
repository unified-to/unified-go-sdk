# PerformanceMetadataValue


## Supported Types

### 

```go
performanceMetadataValue := shared.CreatePerformanceMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
performanceMetadataValue := shared.CreatePerformanceMetadataValueStr(string{/* values here */})
```

### 

```go
performanceMetadataValue := shared.CreatePerformanceMetadataValueNumber(float64{/* values here */})
```

### 

```go
performanceMetadataValue := shared.CreatePerformanceMetadataValueBoolean(bool{/* values here */})
```

### 

```go
performanceMetadataValue := shared.CreatePerformanceMetadataValueArrayOfPerformanceMetadataSchemas5([]shared.PerformanceMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch performanceMetadataValue.Type {
	case shared.PerformanceMetadataValueTypeMapOfAny:
		// performanceMetadataValue.MapOfAny is populated
	case shared.PerformanceMetadataValueTypeStr:
		// performanceMetadataValue.Str is populated
	case shared.PerformanceMetadataValueTypeNumber:
		// performanceMetadataValue.Number is populated
	case shared.PerformanceMetadataValueTypeBoolean:
		// performanceMetadataValue.Boolean is populated
	case shared.PerformanceMetadataValueTypeArrayOfPerformanceMetadataSchemas5:
		// performanceMetadataValue.ArrayOfPerformanceMetadataSchemas5 is populated
}
```
