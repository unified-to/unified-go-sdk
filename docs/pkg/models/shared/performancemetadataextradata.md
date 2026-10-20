# PerformanceMetadataExtraData


## Supported Types

### 

```go
performanceMetadataExtraData := shared.CreatePerformanceMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
performanceMetadataExtraData := shared.CreatePerformanceMetadataExtraDataStr(string{/* values here */})
```

### 

```go
performanceMetadataExtraData := shared.CreatePerformanceMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
performanceMetadataExtraData := shared.CreatePerformanceMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
performanceMetadataExtraData := shared.CreatePerformanceMetadataExtraDataArrayOfPerformanceMetadata5([]shared.PerformanceMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch performanceMetadataExtraData.Type {
	case shared.PerformanceMetadataExtraDataTypeMapOfAny:
		// performanceMetadataExtraData.MapOfAny is populated
	case shared.PerformanceMetadataExtraDataTypeStr:
		// performanceMetadataExtraData.Str is populated
	case shared.PerformanceMetadataExtraDataTypeNumber:
		// performanceMetadataExtraData.Number is populated
	case shared.PerformanceMetadataExtraDataTypeBoolean:
		// performanceMetadataExtraData.Boolean is populated
	case shared.PerformanceMetadataExtraDataTypeArrayOfPerformanceMetadata5:
		// performanceMetadataExtraData.ArrayOfPerformanceMetadata5 is populated
}
```
