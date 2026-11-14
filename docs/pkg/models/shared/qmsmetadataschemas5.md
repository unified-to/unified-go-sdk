# QmsMetadataSchemas5


## Supported Types

### QmsMetadataSchemas1

```go
qmsMetadataSchemas5 := shared.CreateQmsMetadataSchemas5QmsMetadataSchemas1(shared.QmsMetadataSchemas1{/* values here */})
```

### 

```go
qmsMetadataSchemas5 := shared.CreateQmsMetadataSchemas5Str(string{/* values here */})
```

### 

```go
qmsMetadataSchemas5 := shared.CreateQmsMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
qmsMetadataSchemas5 := shared.CreateQmsMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch qmsMetadataSchemas5.Type {
	case shared.QmsMetadataSchemas5TypeQmsMetadataSchemas1:
		// qmsMetadataSchemas5.QmsMetadataSchemas1 is populated
	case shared.QmsMetadataSchemas5TypeStr:
		// qmsMetadataSchemas5.Str is populated
	case shared.QmsMetadataSchemas5TypeNumber:
		// qmsMetadataSchemas5.Number is populated
	case shared.QmsMetadataSchemas5TypeBoolean:
		// qmsMetadataSchemas5.Boolean is populated
}
```
