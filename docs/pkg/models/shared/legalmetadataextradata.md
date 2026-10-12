# LegalMetadataExtraData


## Supported Types

### 

```go
legalMetadataExtraData := shared.CreateLegalMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
legalMetadataExtraData := shared.CreateLegalMetadataExtraDataStr(string{/* values here */})
```

### 

```go
legalMetadataExtraData := shared.CreateLegalMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
legalMetadataExtraData := shared.CreateLegalMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
legalMetadataExtraData := shared.CreateLegalMetadataExtraDataArrayOfLegalMetadata5([]shared.LegalMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch legalMetadataExtraData.Type {
	case shared.LegalMetadataExtraDataTypeMapOfAny:
		// legalMetadataExtraData.MapOfAny is populated
	case shared.LegalMetadataExtraDataTypeStr:
		// legalMetadataExtraData.Str is populated
	case shared.LegalMetadataExtraDataTypeNumber:
		// legalMetadataExtraData.Number is populated
	case shared.LegalMetadataExtraDataTypeBoolean:
		// legalMetadataExtraData.Boolean is populated
	case shared.LegalMetadataExtraDataTypeArrayOfLegalMetadata5:
		// legalMetadataExtraData.ArrayOfLegalMetadata5 is populated
}
```
