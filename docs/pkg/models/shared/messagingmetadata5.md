# MessagingMetadata5


## Supported Types

### MessagingMetadata1

```go
messagingMetadata5 := shared.CreateMessagingMetadata5MessagingMetadata1(shared.MessagingMetadata1{/* values here */})
```

### 

```go
messagingMetadata5 := shared.CreateMessagingMetadata5Str(string{/* values here */})
```

### 

```go
messagingMetadata5 := shared.CreateMessagingMetadata5Number(float64{/* values here */})
```

### 

```go
messagingMetadata5 := shared.CreateMessagingMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch messagingMetadata5.Type {
	case shared.MessagingMetadata5TypeMessagingMetadata1:
		// messagingMetadata5.MessagingMetadata1 is populated
	case shared.MessagingMetadata5TypeStr:
		// messagingMetadata5.Str is populated
	case shared.MessagingMetadata5TypeNumber:
		// messagingMetadata5.Number is populated
	case shared.MessagingMetadata5TypeBoolean:
		// messagingMetadata5.Boolean is populated
}
```
