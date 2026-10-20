# ShippingMetadataExtraData


## Supported Types

### 

```go
shippingMetadataExtraData := shared.CreateShippingMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
shippingMetadataExtraData := shared.CreateShippingMetadataExtraDataStr(string{/* values here */})
```

### 

```go
shippingMetadataExtraData := shared.CreateShippingMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
shippingMetadataExtraData := shared.CreateShippingMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
shippingMetadataExtraData := shared.CreateShippingMetadataExtraDataArrayOfShippingMetadata5([]shared.ShippingMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch shippingMetadataExtraData.Type {
	case shared.ShippingMetadataExtraDataTypeMapOfAny:
		// shippingMetadataExtraData.MapOfAny is populated
	case shared.ShippingMetadataExtraDataTypeStr:
		// shippingMetadataExtraData.Str is populated
	case shared.ShippingMetadataExtraDataTypeNumber:
		// shippingMetadataExtraData.Number is populated
	case shared.ShippingMetadataExtraDataTypeBoolean:
		// shippingMetadataExtraData.Boolean is populated
	case shared.ShippingMetadataExtraDataTypeArrayOfShippingMetadata5:
		// shippingMetadataExtraData.ArrayOfShippingMetadata5 is populated
}
```
