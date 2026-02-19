# KmsPageMetadataSchemas5


## Supported Types

### KmsPageMetadataSchemas1

```go
kmsPageMetadataSchemas5 := shared.CreateKmsPageMetadataSchemas5KmsPageMetadataSchemas1(shared.KmsPageMetadataSchemas1{/* values here */})
```

### 

```go
kmsPageMetadataSchemas5 := shared.CreateKmsPageMetadataSchemas5Str(string{/* values here */})
```

### 

```go
kmsPageMetadataSchemas5 := shared.CreateKmsPageMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
kmsPageMetadataSchemas5 := shared.CreateKmsPageMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch kmsPageMetadataSchemas5.Type {
	case shared.KmsPageMetadataSchemas5TypeKmsPageMetadataSchemas1:
		// kmsPageMetadataSchemas5.KmsPageMetadataSchemas1 is populated
	case shared.KmsPageMetadataSchemas5TypeStr:
		// kmsPageMetadataSchemas5.Str is populated
	case shared.KmsPageMetadataSchemas5TypeNumber:
		// kmsPageMetadataSchemas5.Number is populated
	case shared.KmsPageMetadataSchemas5TypeBoolean:
		// kmsPageMetadataSchemas5.Boolean is populated
}
```
