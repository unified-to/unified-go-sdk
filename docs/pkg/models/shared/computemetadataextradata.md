# ComputeMetadataExtraData

Additional data about the value


## Supported Types

### 

```go
computeMetadataExtraData := shared.CreateComputeMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
computeMetadataExtraData := shared.CreateComputeMetadataExtraDataStr(string{/* values here */})
```

### 

```go
computeMetadataExtraData := shared.CreateComputeMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
computeMetadataExtraData := shared.CreateComputeMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
computeMetadataExtraData := shared.CreateComputeMetadataExtraDataArrayOfComputeMetadata5([]shared.ComputeMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch computeMetadataExtraData.Type {
	case shared.ComputeMetadataExtraDataTypeMapOfAny:
		// computeMetadataExtraData.MapOfAny is populated
	case shared.ComputeMetadataExtraDataTypeStr:
		// computeMetadataExtraData.Str is populated
	case shared.ComputeMetadataExtraDataTypeNumber:
		// computeMetadataExtraData.Number is populated
	case shared.ComputeMetadataExtraDataTypeBoolean:
		// computeMetadataExtraData.Boolean is populated
	case shared.ComputeMetadataExtraDataTypeArrayOfComputeMetadata5:
		// computeMetadataExtraData.ArrayOfComputeMetadata5 is populated
}
```
