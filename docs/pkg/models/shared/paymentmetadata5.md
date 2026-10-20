# PaymentMetadata5


## Supported Types

### PaymentMetadata1

```go
paymentMetadata5 := shared.CreatePaymentMetadata5PaymentMetadata1(shared.PaymentMetadata1{/* values here */})
```

### 

```go
paymentMetadata5 := shared.CreatePaymentMetadata5Str(string{/* values here */})
```

### 

```go
paymentMetadata5 := shared.CreatePaymentMetadata5Number(float64{/* values here */})
```

### 

```go
paymentMetadata5 := shared.CreatePaymentMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch paymentMetadata5.Type {
	case shared.PaymentMetadata5TypePaymentMetadata1:
		// paymentMetadata5.PaymentMetadata1 is populated
	case shared.PaymentMetadata5TypeStr:
		// paymentMetadata5.Str is populated
	case shared.PaymentMetadata5TypeNumber:
		// paymentMetadata5.Number is populated
	case shared.PaymentMetadata5TypeBoolean:
		// paymentMetadata5.Boolean is populated
}
```
