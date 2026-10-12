# LegalMetadataSchemas5


## Supported Types

### LegalMetadataSchemas1

```go
legalMetadataSchemas5 := shared.CreateLegalMetadataSchemas5LegalMetadataSchemas1(shared.LegalMetadataSchemas1{/* values here */})
```

### 

```go
legalMetadataSchemas5 := shared.CreateLegalMetadataSchemas5Str(string{/* values here */})
```

### 

```go
legalMetadataSchemas5 := shared.CreateLegalMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
legalMetadataSchemas5 := shared.CreateLegalMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch legalMetadataSchemas5.Type {
	case shared.LegalMetadataSchemas5TypeLegalMetadataSchemas1:
		// legalMetadataSchemas5.LegalMetadataSchemas1 is populated
	case shared.LegalMetadataSchemas5TypeStr:
		// legalMetadataSchemas5.Str is populated
	case shared.LegalMetadataSchemas5TypeNumber:
		// legalMetadataSchemas5.Number is populated
	case shared.LegalMetadataSchemas5TypeBoolean:
		// legalMetadataSchemas5.Boolean is populated
}
```
