# PerformanceMetadataSchemas5


## Supported Types

### PerformanceMetadataSchemas1

```go
performanceMetadataSchemas5 := shared.CreatePerformanceMetadataSchemas5PerformanceMetadataSchemas1(shared.PerformanceMetadataSchemas1{/* values here */})
```

### 

```go
performanceMetadataSchemas5 := shared.CreatePerformanceMetadataSchemas5Str(string{/* values here */})
```

### 

```go
performanceMetadataSchemas5 := shared.CreatePerformanceMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
performanceMetadataSchemas5 := shared.CreatePerformanceMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch performanceMetadataSchemas5.Type {
	case shared.PerformanceMetadataSchemas5TypePerformanceMetadataSchemas1:
		// performanceMetadataSchemas5.PerformanceMetadataSchemas1 is populated
	case shared.PerformanceMetadataSchemas5TypeStr:
		// performanceMetadataSchemas5.Str is populated
	case shared.PerformanceMetadataSchemas5TypeNumber:
		// performanceMetadataSchemas5.Number is populated
	case shared.PerformanceMetadataSchemas5TypeBoolean:
		// performanceMetadataSchemas5.Boolean is populated
}
```
