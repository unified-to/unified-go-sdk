# FormsMetadataExtraData


## Supported Types

### 

```go
formsMetadataExtraData := shared.CreateFormsMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
formsMetadataExtraData := shared.CreateFormsMetadataExtraDataStr(string{/* values here */})
```

### 

```go
formsMetadataExtraData := shared.CreateFormsMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
formsMetadataExtraData := shared.CreateFormsMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
formsMetadataExtraData := shared.CreateFormsMetadataExtraDataArrayOfFormsMetadata5([]shared.FormsMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch formsMetadataExtraData.Type {
	case shared.FormsMetadataExtraDataTypeMapOfAny:
		// formsMetadataExtraData.MapOfAny is populated
	case shared.FormsMetadataExtraDataTypeStr:
		// formsMetadataExtraData.Str is populated
	case shared.FormsMetadataExtraDataTypeNumber:
		// formsMetadataExtraData.Number is populated
	case shared.FormsMetadataExtraDataTypeBoolean:
		// formsMetadataExtraData.Boolean is populated
	case shared.FormsMetadataExtraDataTypeArrayOfFormsMetadata5:
		// formsMetadataExtraData.ArrayOfFormsMetadata5 is populated
}
```
