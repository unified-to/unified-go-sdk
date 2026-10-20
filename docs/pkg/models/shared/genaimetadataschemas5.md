# GenaiMetadataSchemas5


## Supported Types

### GenaiMetadataSchemas1

```go
genaiMetadataSchemas5 := shared.CreateGenaiMetadataSchemas5GenaiMetadataSchemas1(shared.GenaiMetadataSchemas1{/* values here */})
```

### 

```go
genaiMetadataSchemas5 := shared.CreateGenaiMetadataSchemas5Str(string{/* values here */})
```

### 

```go
genaiMetadataSchemas5 := shared.CreateGenaiMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
genaiMetadataSchemas5 := shared.CreateGenaiMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch genaiMetadataSchemas5.Type {
	case shared.GenaiMetadataSchemas5TypeGenaiMetadataSchemas1:
		// genaiMetadataSchemas5.GenaiMetadataSchemas1 is populated
	case shared.GenaiMetadataSchemas5TypeStr:
		// genaiMetadataSchemas5.Str is populated
	case shared.GenaiMetadataSchemas5TypeNumber:
		// genaiMetadataSchemas5.Number is populated
	case shared.GenaiMetadataSchemas5TypeBoolean:
		// genaiMetadataSchemas5.Boolean is populated
}
```
