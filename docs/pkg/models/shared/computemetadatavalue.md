# ComputeMetadataValue

The label, tag or metadata value


## Supported Types

### 

```go
computeMetadataValue := shared.CreateComputeMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
computeMetadataValue := shared.CreateComputeMetadataValueStr(string{/* values here */})
```

### 

```go
computeMetadataValue := shared.CreateComputeMetadataValueNumber(float64{/* values here */})
```

### 

```go
computeMetadataValue := shared.CreateComputeMetadataValueBoolean(bool{/* values here */})
```

### 

```go
computeMetadataValue := shared.CreateComputeMetadataValueArrayOfComputeMetadataSchemas5([]shared.ComputeMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch computeMetadataValue.Type {
	case shared.ComputeMetadataValueTypeMapOfAny:
		// computeMetadataValue.MapOfAny is populated
	case shared.ComputeMetadataValueTypeStr:
		// computeMetadataValue.Str is populated
	case shared.ComputeMetadataValueTypeNumber:
		// computeMetadataValue.Number is populated
	case shared.ComputeMetadataValueTypeBoolean:
		// computeMetadataValue.Boolean is populated
	case shared.ComputeMetadataValueTypeArrayOfComputeMetadataSchemas5:
		// computeMetadataValue.ArrayOfComputeMetadataSchemas5 is populated
}
```
