# FormsMetadataSchemas5


## Supported Types

### FormsMetadataSchemas1

```go
formsMetadataSchemas5 := shared.CreateFormsMetadataSchemas5FormsMetadataSchemas1(shared.FormsMetadataSchemas1{/* values here */})
```

### 

```go
formsMetadataSchemas5 := shared.CreateFormsMetadataSchemas5Str(string{/* values here */})
```

### 

```go
formsMetadataSchemas5 := shared.CreateFormsMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
formsMetadataSchemas5 := shared.CreateFormsMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch formsMetadataSchemas5.Type {
	case shared.FormsMetadataSchemas5TypeFormsMetadataSchemas1:
		// formsMetadataSchemas5.FormsMetadataSchemas1 is populated
	case shared.FormsMetadataSchemas5TypeStr:
		// formsMetadataSchemas5.Str is populated
	case shared.FormsMetadataSchemas5TypeNumber:
		// formsMetadataSchemas5.Number is populated
	case shared.FormsMetadataSchemas5TypeBoolean:
		// formsMetadataSchemas5.Boolean is populated
}
```
