# PlmMetadataSchemas5


## Supported Types

### PlmMetadataSchemas1

```go
plmMetadataSchemas5 := shared.CreatePlmMetadataSchemas5PlmMetadataSchemas1(shared.PlmMetadataSchemas1{/* values here */})
```

### 

```go
plmMetadataSchemas5 := shared.CreatePlmMetadataSchemas5Str(string{/* values here */})
```

### 

```go
plmMetadataSchemas5 := shared.CreatePlmMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
plmMetadataSchemas5 := shared.CreatePlmMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch plmMetadataSchemas5.Type {
	case shared.PlmMetadataSchemas5TypePlmMetadataSchemas1:
		// plmMetadataSchemas5.PlmMetadataSchemas1 is populated
	case shared.PlmMetadataSchemas5TypeStr:
		// plmMetadataSchemas5.Str is populated
	case shared.PlmMetadataSchemas5TypeNumber:
		// plmMetadataSchemas5.Number is populated
	case shared.PlmMetadataSchemas5TypeBoolean:
		// plmMetadataSchemas5.Boolean is populated
}
```
