# SigningMetadataSchemas5


## Supported Types

### SigningMetadataSchemas1

```go
signingMetadataSchemas5 := shared.CreateSigningMetadataSchemas5SigningMetadataSchemas1(shared.SigningMetadataSchemas1{/* values here */})
```

### 

```go
signingMetadataSchemas5 := shared.CreateSigningMetadataSchemas5Str(string{/* values here */})
```

### 

```go
signingMetadataSchemas5 := shared.CreateSigningMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
signingMetadataSchemas5 := shared.CreateSigningMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch signingMetadataSchemas5.Type {
	case shared.SigningMetadataSchemas5TypeSigningMetadataSchemas1:
		// signingMetadataSchemas5.SigningMetadataSchemas1 is populated
	case shared.SigningMetadataSchemas5TypeStr:
		// signingMetadataSchemas5.Str is populated
	case shared.SigningMetadataSchemas5TypeNumber:
		// signingMetadataSchemas5.Number is populated
	case shared.SigningMetadataSchemas5TypeBoolean:
		// signingMetadataSchemas5.Boolean is populated
}
```
