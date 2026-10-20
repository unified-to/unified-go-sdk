# AnalyticsMetadataValue


## Supported Types

### 

```go
analyticsMetadataValue := shared.CreateAnalyticsMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
analyticsMetadataValue := shared.CreateAnalyticsMetadataValueStr(string{/* values here */})
```

### 

```go
analyticsMetadataValue := shared.CreateAnalyticsMetadataValueNumber(float64{/* values here */})
```

### 

```go
analyticsMetadataValue := shared.CreateAnalyticsMetadataValueBoolean(bool{/* values here */})
```

### 

```go
analyticsMetadataValue := shared.CreateAnalyticsMetadataValueArrayOfAnalyticsMetadataSchemas5([]shared.AnalyticsMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch analyticsMetadataValue.Type {
	case shared.AnalyticsMetadataValueTypeMapOfAny:
		// analyticsMetadataValue.MapOfAny is populated
	case shared.AnalyticsMetadataValueTypeStr:
		// analyticsMetadataValue.Str is populated
	case shared.AnalyticsMetadataValueTypeNumber:
		// analyticsMetadataValue.Number is populated
	case shared.AnalyticsMetadataValueTypeBoolean:
		// analyticsMetadataValue.Boolean is populated
	case shared.AnalyticsMetadataValueTypeArrayOfAnalyticsMetadataSchemas5:
		// analyticsMetadataValue.ArrayOfAnalyticsMetadataSchemas5 is populated
}
```
