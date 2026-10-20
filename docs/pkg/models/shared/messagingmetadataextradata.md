# MessagingMetadataExtraData


## Supported Types

### 

```go
messagingMetadataExtraData := shared.CreateMessagingMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
messagingMetadataExtraData := shared.CreateMessagingMetadataExtraDataStr(string{/* values here */})
```

### 

```go
messagingMetadataExtraData := shared.CreateMessagingMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
messagingMetadataExtraData := shared.CreateMessagingMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
messagingMetadataExtraData := shared.CreateMessagingMetadataExtraDataArrayOfMessagingMetadata5([]shared.MessagingMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch messagingMetadataExtraData.Type {
	case shared.MessagingMetadataExtraDataTypeMapOfAny:
		// messagingMetadataExtraData.MapOfAny is populated
	case shared.MessagingMetadataExtraDataTypeStr:
		// messagingMetadataExtraData.Str is populated
	case shared.MessagingMetadataExtraDataTypeNumber:
		// messagingMetadataExtraData.Number is populated
	case shared.MessagingMetadataExtraDataTypeBoolean:
		// messagingMetadataExtraData.Boolean is populated
	case shared.MessagingMetadataExtraDataTypeArrayOfMessagingMetadata5:
		// messagingMetadataExtraData.ArrayOfMessagingMetadata5 is populated
}
```
