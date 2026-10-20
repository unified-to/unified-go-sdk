# AnalyticsMetadataSchemas5


## Supported Types

### AnalyticsMetadataSchemas1

```go
analyticsMetadataSchemas5 := shared.CreateAnalyticsMetadataSchemas5AnalyticsMetadataSchemas1(shared.AnalyticsMetadataSchemas1{/* values here */})
```

### 

```go
analyticsMetadataSchemas5 := shared.CreateAnalyticsMetadataSchemas5Str(string{/* values here */})
```

### 

```go
analyticsMetadataSchemas5 := shared.CreateAnalyticsMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
analyticsMetadataSchemas5 := shared.CreateAnalyticsMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch analyticsMetadataSchemas5.Type {
	case shared.AnalyticsMetadataSchemas5TypeAnalyticsMetadataSchemas1:
		// analyticsMetadataSchemas5.AnalyticsMetadataSchemas1 is populated
	case shared.AnalyticsMetadataSchemas5TypeStr:
		// analyticsMetadataSchemas5.Str is populated
	case shared.AnalyticsMetadataSchemas5TypeNumber:
		// analyticsMetadataSchemas5.Number is populated
	case shared.AnalyticsMetadataSchemas5TypeBoolean:
		// analyticsMetadataSchemas5.Boolean is populated
}
```
