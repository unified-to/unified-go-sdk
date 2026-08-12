# AtsMetadataExtraData


## Supported Types

### 

```go
atsMetadataExtraData := shared.CreateAtsMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
atsMetadataExtraData := shared.CreateAtsMetadataExtraDataStr(string{/* values here */})
```

### 

```go
atsMetadataExtraData := shared.CreateAtsMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
atsMetadataExtraData := shared.CreateAtsMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
atsMetadataExtraData := shared.CreateAtsMetadataExtraDataArrayOfAtsMetadata5([]shared.AtsMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch atsMetadataExtraData.Type {
	case shared.AtsMetadataExtraDataTypeMapOfAny:
		// atsMetadataExtraData.MapOfAny is populated
	case shared.AtsMetadataExtraDataTypeStr:
		// atsMetadataExtraData.Str is populated
	case shared.AtsMetadataExtraDataTypeNumber:
		// atsMetadataExtraData.Number is populated
	case shared.AtsMetadataExtraDataTypeBoolean:
		// atsMetadataExtraData.Boolean is populated
	case shared.AtsMetadataExtraDataTypeArrayOfAtsMetadata5:
		// atsMetadataExtraData.ArrayOfAtsMetadata5 is populated
}
```
