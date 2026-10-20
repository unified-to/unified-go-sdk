# ShippingMetadata5


## Supported Types

### ShippingMetadata1

```go
shippingMetadata5 := shared.CreateShippingMetadata5ShippingMetadata1(shared.ShippingMetadata1{/* values here */})
```

### 

```go
shippingMetadata5 := shared.CreateShippingMetadata5Str(string{/* values here */})
```

### 

```go
shippingMetadata5 := shared.CreateShippingMetadata5Number(float64{/* values here */})
```

### 

```go
shippingMetadata5 := shared.CreateShippingMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch shippingMetadata5.Type {
	case shared.ShippingMetadata5TypeShippingMetadata1:
		// shippingMetadata5.ShippingMetadata1 is populated
	case shared.ShippingMetadata5TypeStr:
		// shippingMetadata5.Str is populated
	case shared.ShippingMetadata5TypeNumber:
		// shippingMetadata5.Number is populated
	case shared.ShippingMetadata5TypeBoolean:
		// shippingMetadata5.Boolean is populated
}
```
