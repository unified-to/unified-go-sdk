# PaymentMetadataExtraData


## Supported Types

### 

```go
paymentMetadataExtraData := shared.CreatePaymentMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
paymentMetadataExtraData := shared.CreatePaymentMetadataExtraDataStr(string{/* values here */})
```

### 

```go
paymentMetadataExtraData := shared.CreatePaymentMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
paymentMetadataExtraData := shared.CreatePaymentMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
paymentMetadataExtraData := shared.CreatePaymentMetadataExtraDataArrayOfPaymentMetadata5([]shared.PaymentMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch paymentMetadataExtraData.Type {
	case shared.PaymentMetadataExtraDataTypeMapOfAny:
		// paymentMetadataExtraData.MapOfAny is populated
	case shared.PaymentMetadataExtraDataTypeStr:
		// paymentMetadataExtraData.Str is populated
	case shared.PaymentMetadataExtraDataTypeNumber:
		// paymentMetadataExtraData.Number is populated
	case shared.PaymentMetadataExtraDataTypeBoolean:
		// paymentMetadataExtraData.Boolean is populated
	case shared.PaymentMetadataExtraDataTypeArrayOfPaymentMetadata5:
		// paymentMetadataExtraData.ArrayOfPaymentMetadata5 is populated
}
```
