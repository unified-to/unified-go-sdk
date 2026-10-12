# ComputeMetadata5


## Supported Types

### ComputeMetadata1

```go
computeMetadata5 := shared.CreateComputeMetadata5ComputeMetadata1(shared.ComputeMetadata1{/* values here */})
```

### 

```go
computeMetadata5 := shared.CreateComputeMetadata5Str(string{/* values here */})
```

### 

```go
computeMetadata5 := shared.CreateComputeMetadata5Number(float64{/* values here */})
```

### 

```go
computeMetadata5 := shared.CreateComputeMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch computeMetadata5.Type {
	case shared.ComputeMetadata5TypeComputeMetadata1:
		// computeMetadata5.ComputeMetadata1 is populated
	case shared.ComputeMetadata5TypeStr:
		// computeMetadata5.Str is populated
	case shared.ComputeMetadata5TypeNumber:
		// computeMetadata5.Number is populated
	case shared.ComputeMetadata5TypeBoolean:
		// computeMetadata5.Boolean is populated
}
```
