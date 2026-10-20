# LmsMetadataExtraData


## Supported Types

### 

```go
lmsMetadataExtraData := shared.CreateLmsMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
lmsMetadataExtraData := shared.CreateLmsMetadataExtraDataStr(string{/* values here */})
```

### 

```go
lmsMetadataExtraData := shared.CreateLmsMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
lmsMetadataExtraData := shared.CreateLmsMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
lmsMetadataExtraData := shared.CreateLmsMetadataExtraDataArrayOfLmsMetadata5([]shared.LmsMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch lmsMetadataExtraData.Type {
	case shared.LmsMetadataExtraDataTypeMapOfAny:
		// lmsMetadataExtraData.MapOfAny is populated
	case shared.LmsMetadataExtraDataTypeStr:
		// lmsMetadataExtraData.Str is populated
	case shared.LmsMetadataExtraDataTypeNumber:
		// lmsMetadataExtraData.Number is populated
	case shared.LmsMetadataExtraDataTypeBoolean:
		// lmsMetadataExtraData.Boolean is populated
	case shared.LmsMetadataExtraDataTypeArrayOfLmsMetadata5:
		// lmsMetadataExtraData.ArrayOfLmsMetadata5 is populated
}
```
