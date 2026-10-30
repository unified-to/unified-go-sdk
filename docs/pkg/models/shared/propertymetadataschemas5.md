# PropertyMetadataSchemas5


## Supported Types

### PropertyMetadataSchemas1

```go
propertyMetadataSchemas5 := shared.CreatePropertyMetadataSchemas5PropertyMetadataSchemas1(shared.PropertyMetadataSchemas1{/* values here */})
```

### 

```go
propertyMetadataSchemas5 := shared.CreatePropertyMetadataSchemas5Str(string{/* values here */})
```

### 

```go
propertyMetadataSchemas5 := shared.CreatePropertyMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
propertyMetadataSchemas5 := shared.CreatePropertyMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch propertyMetadataSchemas5.Type {
	case shared.PropertyMetadataSchemas5TypePropertyMetadataSchemas1:
		// propertyMetadataSchemas5.PropertyMetadataSchemas1 is populated
	case shared.PropertyMetadataSchemas5TypeStr:
		// propertyMetadataSchemas5.Str is populated
	case shared.PropertyMetadataSchemas5TypeNumber:
		// propertyMetadataSchemas5.Number is populated
	case shared.PropertyMetadataSchemas5TypeBoolean:
		// propertyMetadataSchemas5.Boolean is populated
}
```
