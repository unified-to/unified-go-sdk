# GenaiMetadataExtraData


## Supported Types

### 

```go
genaiMetadataExtraData := shared.CreateGenaiMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
genaiMetadataExtraData := shared.CreateGenaiMetadataExtraDataStr(string{/* values here */})
```

### 

```go
genaiMetadataExtraData := shared.CreateGenaiMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
genaiMetadataExtraData := shared.CreateGenaiMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
genaiMetadataExtraData := shared.CreateGenaiMetadataExtraDataArrayOfGenaiMetadata5([]shared.GenaiMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch genaiMetadataExtraData.Type {
	case shared.GenaiMetadataExtraDataTypeMapOfAny:
		// genaiMetadataExtraData.MapOfAny is populated
	case shared.GenaiMetadataExtraDataTypeStr:
		// genaiMetadataExtraData.Str is populated
	case shared.GenaiMetadataExtraDataTypeNumber:
		// genaiMetadataExtraData.Number is populated
	case shared.GenaiMetadataExtraDataTypeBoolean:
		// genaiMetadataExtraData.Boolean is populated
	case shared.GenaiMetadataExtraDataTypeArrayOfGenaiMetadata5:
		// genaiMetadataExtraData.ArrayOfGenaiMetadata5 is populated
}
```
