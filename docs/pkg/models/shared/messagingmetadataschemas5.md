# MessagingMetadataSchemas5


## Supported Types

### MessagingMetadataSchemas1

```go
messagingMetadataSchemas5 := shared.CreateMessagingMetadataSchemas5MessagingMetadataSchemas1(shared.MessagingMetadataSchemas1{/* values here */})
```

### 

```go
messagingMetadataSchemas5 := shared.CreateMessagingMetadataSchemas5Str(string{/* values here */})
```

### 

```go
messagingMetadataSchemas5 := shared.CreateMessagingMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
messagingMetadataSchemas5 := shared.CreateMessagingMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch messagingMetadataSchemas5.Type {
	case shared.MessagingMetadataSchemas5TypeMessagingMetadataSchemas1:
		// messagingMetadataSchemas5.MessagingMetadataSchemas1 is populated
	case shared.MessagingMetadataSchemas5TypeStr:
		// messagingMetadataSchemas5.Str is populated
	case shared.MessagingMetadataSchemas5TypeNumber:
		// messagingMetadataSchemas5.Number is populated
	case shared.MessagingMetadataSchemas5TypeBoolean:
		// messagingMetadataSchemas5.Boolean is populated
}
```
