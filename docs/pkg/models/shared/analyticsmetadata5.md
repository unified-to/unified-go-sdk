# AnalyticsMetadata5


## Supported Types

### AnalyticsMetadata1

```go
analyticsMetadata5 := shared.CreateAnalyticsMetadata5AnalyticsMetadata1(shared.AnalyticsMetadata1{/* values here */})
```

### 

```go
analyticsMetadata5 := shared.CreateAnalyticsMetadata5Str(string{/* values here */})
```

### 

```go
analyticsMetadata5 := shared.CreateAnalyticsMetadata5Number(float64{/* values here */})
```

### 

```go
analyticsMetadata5 := shared.CreateAnalyticsMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch analyticsMetadata5.Type {
	case shared.AnalyticsMetadata5TypeAnalyticsMetadata1:
		// analyticsMetadata5.AnalyticsMetadata1 is populated
	case shared.AnalyticsMetadata5TypeStr:
		// analyticsMetadata5.Str is populated
	case shared.AnalyticsMetadata5TypeNumber:
		// analyticsMetadata5.Number is populated
	case shared.AnalyticsMetadata5TypeBoolean:
		// analyticsMetadata5.Boolean is populated
}
```
