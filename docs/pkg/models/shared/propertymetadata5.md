# PropertyMetadata5


## Supported Types

### PropertyMetadata1

```go
propertyMetadata5 := shared.CreatePropertyMetadata5PropertyMetadata1(shared.PropertyMetadata1{/* values here */})
```

### 

```go
propertyMetadata5 := shared.CreatePropertyMetadata5Str(string{/* values here */})
```

### 

```go
propertyMetadata5 := shared.CreatePropertyMetadata5Number(float64{/* values here */})
```

### 

```go
propertyMetadata5 := shared.CreatePropertyMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch propertyMetadata5.Type {
	case shared.PropertyMetadata5TypePropertyMetadata1:
		// propertyMetadata5.PropertyMetadata1 is populated
	case shared.PropertyMetadata5TypeStr:
		// propertyMetadata5.Str is populated
	case shared.PropertyMetadata5TypeNumber:
		// propertyMetadata5.Number is populated
	case shared.PropertyMetadata5TypeBoolean:
		// propertyMetadata5.Boolean is populated
}
```
