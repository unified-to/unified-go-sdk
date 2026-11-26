# DmsMetadataSchemas5


## Supported Types

### DmsMetadataSchemas1

```go
dmsMetadataSchemas5 := shared.CreateDmsMetadataSchemas5DmsMetadataSchemas1(shared.DmsMetadataSchemas1{/* values here */})
```

### 

```go
dmsMetadataSchemas5 := shared.CreateDmsMetadataSchemas5Str(string{/* values here */})
```

### 

```go
dmsMetadataSchemas5 := shared.CreateDmsMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
dmsMetadataSchemas5 := shared.CreateDmsMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch dmsMetadataSchemas5.Type {
	case shared.DmsMetadataSchemas5TypeDmsMetadataSchemas1:
		// dmsMetadataSchemas5.DmsMetadataSchemas1 is populated
	case shared.DmsMetadataSchemas5TypeStr:
		// dmsMetadataSchemas5.Str is populated
	case shared.DmsMetadataSchemas5TypeNumber:
		// dmsMetadataSchemas5.Number is populated
	case shared.DmsMetadataSchemas5TypeBoolean:
		// dmsMetadataSchemas5.Boolean is populated
}
```
