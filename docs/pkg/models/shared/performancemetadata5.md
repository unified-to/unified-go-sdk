# PerformanceMetadata5


## Supported Types

### PerformanceMetadata1

```go
performanceMetadata5 := shared.CreatePerformanceMetadata5PerformanceMetadata1(shared.PerformanceMetadata1{/* values here */})
```

### 

```go
performanceMetadata5 := shared.CreatePerformanceMetadata5Str(string{/* values here */})
```

### 

```go
performanceMetadata5 := shared.CreatePerformanceMetadata5Number(float64{/* values here */})
```

### 

```go
performanceMetadata5 := shared.CreatePerformanceMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch performanceMetadata5.Type {
	case shared.PerformanceMetadata5TypePerformanceMetadata1:
		// performanceMetadata5.PerformanceMetadata1 is populated
	case shared.PerformanceMetadata5TypeStr:
		// performanceMetadata5.Str is populated
	case shared.PerformanceMetadata5TypeNumber:
		// performanceMetadata5.Number is populated
	case shared.PerformanceMetadata5TypeBoolean:
		// performanceMetadata5.Boolean is populated
}
```
