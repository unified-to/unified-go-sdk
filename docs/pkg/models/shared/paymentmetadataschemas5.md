# PaymentMetadataSchemas5


## Supported Types

### PaymentMetadataSchemas1

```go
paymentMetadataSchemas5 := shared.CreatePaymentMetadataSchemas5PaymentMetadataSchemas1(shared.PaymentMetadataSchemas1{/* values here */})
```

### 

```go
paymentMetadataSchemas5 := shared.CreatePaymentMetadataSchemas5Str(string{/* values here */})
```

### 

```go
paymentMetadataSchemas5 := shared.CreatePaymentMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
paymentMetadataSchemas5 := shared.CreatePaymentMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch paymentMetadataSchemas5.Type {
	case shared.PaymentMetadataSchemas5TypePaymentMetadataSchemas1:
		// paymentMetadataSchemas5.PaymentMetadataSchemas1 is populated
	case shared.PaymentMetadataSchemas5TypeStr:
		// paymentMetadataSchemas5.Str is populated
	case shared.PaymentMetadataSchemas5TypeNumber:
		// paymentMetadataSchemas5.Number is populated
	case shared.PaymentMetadataSchemas5TypeBoolean:
		// paymentMetadataSchemas5.Boolean is populated
}
```
