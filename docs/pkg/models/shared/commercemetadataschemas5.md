# CommerceMetadataSchemas5


## Supported Types

### CommerceMetadataSchemas1

```go
commerceMetadataSchemas5 := shared.CreateCommerceMetadataSchemas5CommerceMetadataSchemas1(shared.CommerceMetadataSchemas1{/* values here */})
```

### 

```go
commerceMetadataSchemas5 := shared.CreateCommerceMetadataSchemas5Str(string{/* values here */})
```

### 

```go
commerceMetadataSchemas5 := shared.CreateCommerceMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
commerceMetadataSchemas5 := shared.CreateCommerceMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch commerceMetadataSchemas5.Type {
	case shared.CommerceMetadataSchemas5TypeCommerceMetadataSchemas1:
		// commerceMetadataSchemas5.CommerceMetadataSchemas1 is populated
	case shared.CommerceMetadataSchemas5TypeStr:
		// commerceMetadataSchemas5.Str is populated
	case shared.CommerceMetadataSchemas5TypeNumber:
		// commerceMetadataSchemas5.Number is populated
	case shared.CommerceMetadataSchemas5TypeBoolean:
		// commerceMetadataSchemas5.Boolean is populated
}
```
