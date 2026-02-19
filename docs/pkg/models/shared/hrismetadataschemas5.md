# HrisMetadataSchemas5


## Supported Types

### HrisMetadataSchemas1

```go
hrisMetadataSchemas5 := shared.CreateHrisMetadataSchemas5HrisMetadataSchemas1(shared.HrisMetadataSchemas1{/* values here */})
```

### 

```go
hrisMetadataSchemas5 := shared.CreateHrisMetadataSchemas5Str(string{/* values here */})
```

### 

```go
hrisMetadataSchemas5 := shared.CreateHrisMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
hrisMetadataSchemas5 := shared.CreateHrisMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch hrisMetadataSchemas5.Type {
	case shared.HrisMetadataSchemas5TypeHrisMetadataSchemas1:
		// hrisMetadataSchemas5.HrisMetadataSchemas1 is populated
	case shared.HrisMetadataSchemas5TypeStr:
		// hrisMetadataSchemas5.Str is populated
	case shared.HrisMetadataSchemas5TypeNumber:
		// hrisMetadataSchemas5.Number is populated
	case shared.HrisMetadataSchemas5TypeBoolean:
		// hrisMetadataSchemas5.Boolean is populated
}
```
