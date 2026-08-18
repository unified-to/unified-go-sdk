# UcMetadataSchemas5


## Supported Types

### UcMetadataSchemas1

```go
ucMetadataSchemas5 := shared.CreateUcMetadataSchemas5UcMetadataSchemas1(shared.UcMetadataSchemas1{/* values here */})
```

### 

```go
ucMetadataSchemas5 := shared.CreateUcMetadataSchemas5Str(string{/* values here */})
```

### 

```go
ucMetadataSchemas5 := shared.CreateUcMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
ucMetadataSchemas5 := shared.CreateUcMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch ucMetadataSchemas5.Type {
	case shared.UcMetadataSchemas5TypeUcMetadataSchemas1:
		// ucMetadataSchemas5.UcMetadataSchemas1 is populated
	case shared.UcMetadataSchemas5TypeStr:
		// ucMetadataSchemas5.Str is populated
	case shared.UcMetadataSchemas5TypeNumber:
		// ucMetadataSchemas5.Number is populated
	case shared.UcMetadataSchemas5TypeBoolean:
		// ucMetadataSchemas5.Boolean is populated
}
```
