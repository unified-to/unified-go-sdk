# ShippingMetadataValue


## Supported Types

### 

```go
shippingMetadataValue := shared.CreateShippingMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
shippingMetadataValue := shared.CreateShippingMetadataValueStr(string{/* values here */})
```

### 

```go
shippingMetadataValue := shared.CreateShippingMetadataValueNumber(float64{/* values here */})
```

### 

```go
shippingMetadataValue := shared.CreateShippingMetadataValueBoolean(bool{/* values here */})
```

### 

```go
shippingMetadataValue := shared.CreateShippingMetadataValueArrayOfShippingMetadataSchemas5([]shared.ShippingMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch shippingMetadataValue.Type {
	case shared.ShippingMetadataValueTypeMapOfAny:
		// shippingMetadataValue.MapOfAny is populated
	case shared.ShippingMetadataValueTypeStr:
		// shippingMetadataValue.Str is populated
	case shared.ShippingMetadataValueTypeNumber:
		// shippingMetadataValue.Number is populated
	case shared.ShippingMetadataValueTypeBoolean:
		// shippingMetadataValue.Boolean is populated
	case shared.ShippingMetadataValueTypeArrayOfShippingMetadataSchemas5:
		// shippingMetadataValue.ArrayOfShippingMetadataSchemas5 is populated
}
```
