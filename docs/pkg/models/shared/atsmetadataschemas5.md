# AtsMetadataSchemas5


## Supported Types

### AtsMetadataSchemas1

```go
atsMetadataSchemas5 := shared.CreateAtsMetadataSchemas5AtsMetadataSchemas1(shared.AtsMetadataSchemas1{/* values here */})
```

### 

```go
atsMetadataSchemas5 := shared.CreateAtsMetadataSchemas5Str(string{/* values here */})
```

### 

```go
atsMetadataSchemas5 := shared.CreateAtsMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
atsMetadataSchemas5 := shared.CreateAtsMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch atsMetadataSchemas5.Type {
	case shared.AtsMetadataSchemas5TypeAtsMetadataSchemas1:
		// atsMetadataSchemas5.AtsMetadataSchemas1 is populated
	case shared.AtsMetadataSchemas5TypeStr:
		// atsMetadataSchemas5.Str is populated
	case shared.AtsMetadataSchemas5TypeNumber:
		// atsMetadataSchemas5.Number is populated
	case shared.AtsMetadataSchemas5TypeBoolean:
		// atsMetadataSchemas5.Boolean is populated
}
```
