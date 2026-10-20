# AnalyticsMetadataExtraData


## Supported Types

### 

```go
analyticsMetadataExtraData := shared.CreateAnalyticsMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
analyticsMetadataExtraData := shared.CreateAnalyticsMetadataExtraDataStr(string{/* values here */})
```

### 

```go
analyticsMetadataExtraData := shared.CreateAnalyticsMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
analyticsMetadataExtraData := shared.CreateAnalyticsMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
analyticsMetadataExtraData := shared.CreateAnalyticsMetadataExtraDataArrayOfAnalyticsMetadata5([]shared.AnalyticsMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch analyticsMetadataExtraData.Type {
	case shared.AnalyticsMetadataExtraDataTypeMapOfAny:
		// analyticsMetadataExtraData.MapOfAny is populated
	case shared.AnalyticsMetadataExtraDataTypeStr:
		// analyticsMetadataExtraData.Str is populated
	case shared.AnalyticsMetadataExtraDataTypeNumber:
		// analyticsMetadataExtraData.Number is populated
	case shared.AnalyticsMetadataExtraDataTypeBoolean:
		// analyticsMetadataExtraData.Boolean is populated
	case shared.AnalyticsMetadataExtraDataTypeArrayOfAnalyticsMetadata5:
		// analyticsMetadataExtraData.ArrayOfAnalyticsMetadata5 is populated
}
```
