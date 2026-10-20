# LmsMetadataSchemas5


## Supported Types

### LmsMetadataSchemas1

```go
lmsMetadataSchemas5 := shared.CreateLmsMetadataSchemas5LmsMetadataSchemas1(shared.LmsMetadataSchemas1{/* values here */})
```

### 

```go
lmsMetadataSchemas5 := shared.CreateLmsMetadataSchemas5Str(string{/* values here */})
```

### 

```go
lmsMetadataSchemas5 := shared.CreateLmsMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
lmsMetadataSchemas5 := shared.CreateLmsMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch lmsMetadataSchemas5.Type {
	case shared.LmsMetadataSchemas5TypeLmsMetadataSchemas1:
		// lmsMetadataSchemas5.LmsMetadataSchemas1 is populated
	case shared.LmsMetadataSchemas5TypeStr:
		// lmsMetadataSchemas5.Str is populated
	case shared.LmsMetadataSchemas5TypeNumber:
		// lmsMetadataSchemas5.Number is populated
	case shared.LmsMetadataSchemas5TypeBoolean:
		// lmsMetadataSchemas5.Boolean is populated
}
```
