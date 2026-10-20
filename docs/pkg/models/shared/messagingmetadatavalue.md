# MessagingMetadataValue


## Supported Types

### 

```go
messagingMetadataValue := shared.CreateMessagingMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
messagingMetadataValue := shared.CreateMessagingMetadataValueStr(string{/* values here */})
```

### 

```go
messagingMetadataValue := shared.CreateMessagingMetadataValueNumber(float64{/* values here */})
```

### 

```go
messagingMetadataValue := shared.CreateMessagingMetadataValueBoolean(bool{/* values here */})
```

### 

```go
messagingMetadataValue := shared.CreateMessagingMetadataValueArrayOfMessagingMetadataSchemas5([]shared.MessagingMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch messagingMetadataValue.Type {
	case shared.MessagingMetadataValueTypeMapOfAny:
		// messagingMetadataValue.MapOfAny is populated
	case shared.MessagingMetadataValueTypeStr:
		// messagingMetadataValue.Str is populated
	case shared.MessagingMetadataValueTypeNumber:
		// messagingMetadataValue.Number is populated
	case shared.MessagingMetadataValueTypeBoolean:
		// messagingMetadataValue.Boolean is populated
	case shared.MessagingMetadataValueTypeArrayOfMessagingMetadataSchemas5:
		// messagingMetadataValue.ArrayOfMessagingMetadataSchemas5 is populated
}
```
