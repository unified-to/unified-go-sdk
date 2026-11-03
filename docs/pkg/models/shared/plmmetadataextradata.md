# PlmMetadataExtraData


## Supported Types

### 

```go
plmMetadataExtraData := shared.CreatePlmMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
plmMetadataExtraData := shared.CreatePlmMetadataExtraDataStr(string{/* values here */})
```

### 

```go
plmMetadataExtraData := shared.CreatePlmMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
plmMetadataExtraData := shared.CreatePlmMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
plmMetadataExtraData := shared.CreatePlmMetadataExtraDataArrayOfPlmMetadata5([]shared.PlmMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch plmMetadataExtraData.Type {
	case shared.PlmMetadataExtraDataTypeMapOfAny:
		// plmMetadataExtraData.MapOfAny is populated
	case shared.PlmMetadataExtraDataTypeStr:
		// plmMetadataExtraData.Str is populated
	case shared.PlmMetadataExtraDataTypeNumber:
		// plmMetadataExtraData.Number is populated
	case shared.PlmMetadataExtraDataTypeBoolean:
		// plmMetadataExtraData.Boolean is populated
	case shared.PlmMetadataExtraDataTypeArrayOfPlmMetadata5:
		// plmMetadataExtraData.ArrayOfPlmMetadata5 is populated
}
```
