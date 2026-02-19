# CommerceMetadata5


## Supported Types

### CommerceMetadata1

```go
commerceMetadata5 := shared.CreateCommerceMetadata5CommerceMetadata1(shared.CommerceMetadata1{/* values here */})
```

### 

```go
commerceMetadata5 := shared.CreateCommerceMetadata5Str(string{/* values here */})
```

### 

```go
commerceMetadata5 := shared.CreateCommerceMetadata5Number(float64{/* values here */})
```

### 

```go
commerceMetadata5 := shared.CreateCommerceMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch commerceMetadata5.Type {
	case shared.CommerceMetadata5TypeCommerceMetadata1:
		// commerceMetadata5.CommerceMetadata1 is populated
	case shared.CommerceMetadata5TypeStr:
		// commerceMetadata5.Str is populated
	case shared.CommerceMetadata5TypeNumber:
		// commerceMetadata5.Number is populated
	case shared.CommerceMetadata5TypeBoolean:
		// commerceMetadata5.Boolean is populated
}
```
