# ShippingMetadataSchemas5


## Supported Types

### ShippingMetadataSchemas1

```go
shippingMetadataSchemas5 := shared.CreateShippingMetadataSchemas5ShippingMetadataSchemas1(shared.ShippingMetadataSchemas1{/* values here */})
```

### 

```go
shippingMetadataSchemas5 := shared.CreateShippingMetadataSchemas5Str(string{/* values here */})
```

### 

```go
shippingMetadataSchemas5 := shared.CreateShippingMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
shippingMetadataSchemas5 := shared.CreateShippingMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch shippingMetadataSchemas5.Type {
	case shared.ShippingMetadataSchemas5TypeShippingMetadataSchemas1:
		// shippingMetadataSchemas5.ShippingMetadataSchemas1 is populated
	case shared.ShippingMetadataSchemas5TypeStr:
		// shippingMetadataSchemas5.Str is populated
	case shared.ShippingMetadataSchemas5TypeNumber:
		// shippingMetadataSchemas5.Number is populated
	case shared.ShippingMetadataSchemas5TypeBoolean:
		// shippingMetadataSchemas5.Boolean is populated
}
```
