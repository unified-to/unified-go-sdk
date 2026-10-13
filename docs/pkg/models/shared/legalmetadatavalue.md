# LegalMetadataValue

the custom field value


## Supported Types

### 

```go
legalMetadataValue := shared.CreateLegalMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
legalMetadataValue := shared.CreateLegalMetadataValueStr(string{/* values here */})
```

### 

```go
legalMetadataValue := shared.CreateLegalMetadataValueNumber(float64{/* values here */})
```

### 

```go
legalMetadataValue := shared.CreateLegalMetadataValueBoolean(bool{/* values here */})
```

### 

```go
legalMetadataValue := shared.CreateLegalMetadataValueArrayOfLegalMetadataSchemas5([]shared.LegalMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch legalMetadataValue.Type {
	case shared.LegalMetadataValueTypeMapOfAny:
		// legalMetadataValue.MapOfAny is populated
	case shared.LegalMetadataValueTypeStr:
		// legalMetadataValue.Str is populated
	case shared.LegalMetadataValueTypeNumber:
		// legalMetadataValue.Number is populated
	case shared.LegalMetadataValueTypeBoolean:
		// legalMetadataValue.Boolean is populated
	case shared.LegalMetadataValueTypeArrayOfLegalMetadataSchemas5:
		// legalMetadataValue.ArrayOfLegalMetadataSchemas5 is populated
}
```
