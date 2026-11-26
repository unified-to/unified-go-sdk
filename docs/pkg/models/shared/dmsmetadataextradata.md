# DmsMetadataExtraData


## Supported Types

### 

```go
dmsMetadataExtraData := shared.CreateDmsMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
dmsMetadataExtraData := shared.CreateDmsMetadataExtraDataStr(string{/* values here */})
```

### 

```go
dmsMetadataExtraData := shared.CreateDmsMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
dmsMetadataExtraData := shared.CreateDmsMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
dmsMetadataExtraData := shared.CreateDmsMetadataExtraDataArrayOfDmsMetadata5([]shared.DmsMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch dmsMetadataExtraData.Type {
	case shared.DmsMetadataExtraDataTypeMapOfAny:
		// dmsMetadataExtraData.MapOfAny is populated
	case shared.DmsMetadataExtraDataTypeStr:
		// dmsMetadataExtraData.Str is populated
	case shared.DmsMetadataExtraDataTypeNumber:
		// dmsMetadataExtraData.Number is populated
	case shared.DmsMetadataExtraDataTypeBoolean:
		// dmsMetadataExtraData.Boolean is populated
	case shared.DmsMetadataExtraDataTypeArrayOfDmsMetadata5:
		// dmsMetadataExtraData.ArrayOfDmsMetadata5 is populated
}
```
