# PaymentMetadataValue


## Supported Types

### 

```go
paymentMetadataValue := shared.CreatePaymentMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
paymentMetadataValue := shared.CreatePaymentMetadataValueStr(string{/* values here */})
```

### 

```go
paymentMetadataValue := shared.CreatePaymentMetadataValueNumber(float64{/* values here */})
```

### 

```go
paymentMetadataValue := shared.CreatePaymentMetadataValueBoolean(bool{/* values here */})
```

### 

```go
paymentMetadataValue := shared.CreatePaymentMetadataValueArrayOfPaymentMetadataSchemas5([]shared.PaymentMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch paymentMetadataValue.Type {
	case shared.PaymentMetadataValueTypeMapOfAny:
		// paymentMetadataValue.MapOfAny is populated
	case shared.PaymentMetadataValueTypeStr:
		// paymentMetadataValue.Str is populated
	case shared.PaymentMetadataValueTypeNumber:
		// paymentMetadataValue.Number is populated
	case shared.PaymentMetadataValueTypeBoolean:
		// paymentMetadataValue.Boolean is populated
	case shared.PaymentMetadataValueTypeArrayOfPaymentMetadataSchemas5:
		// paymentMetadataValue.ArrayOfPaymentMetadataSchemas5 is populated
}
```
