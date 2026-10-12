# ComputeMetadataSchemas5


## Supported Types

### ComputeMetadataSchemas1

```go
computeMetadataSchemas5 := shared.CreateComputeMetadataSchemas5ComputeMetadataSchemas1(shared.ComputeMetadataSchemas1{/* values here */})
```

### 

```go
computeMetadataSchemas5 := shared.CreateComputeMetadataSchemas5Str(string{/* values here */})
```

### 

```go
computeMetadataSchemas5 := shared.CreateComputeMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
computeMetadataSchemas5 := shared.CreateComputeMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch computeMetadataSchemas5.Type {
	case shared.ComputeMetadataSchemas5TypeComputeMetadataSchemas1:
		// computeMetadataSchemas5.ComputeMetadataSchemas1 is populated
	case shared.ComputeMetadataSchemas5TypeStr:
		// computeMetadataSchemas5.Str is populated
	case shared.ComputeMetadataSchemas5TypeNumber:
		// computeMetadataSchemas5.Number is populated
	case shared.ComputeMetadataSchemas5TypeBoolean:
		// computeMetadataSchemas5.Boolean is populated
}
```
