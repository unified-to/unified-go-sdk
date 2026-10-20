# TicketingMetadataExtraData


## Supported Types

### 

```go
ticketingMetadataExtraData := shared.CreateTicketingMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
ticketingMetadataExtraData := shared.CreateTicketingMetadataExtraDataStr(string{/* values here */})
```

### 

```go
ticketingMetadataExtraData := shared.CreateTicketingMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
ticketingMetadataExtraData := shared.CreateTicketingMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
ticketingMetadataExtraData := shared.CreateTicketingMetadataExtraDataArrayOfTicketingMetadata5([]shared.TicketingMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch ticketingMetadataExtraData.Type {
	case shared.TicketingMetadataExtraDataTypeMapOfAny:
		// ticketingMetadataExtraData.MapOfAny is populated
	case shared.TicketingMetadataExtraDataTypeStr:
		// ticketingMetadataExtraData.Str is populated
	case shared.TicketingMetadataExtraDataTypeNumber:
		// ticketingMetadataExtraData.Number is populated
	case shared.TicketingMetadataExtraDataTypeBoolean:
		// ticketingMetadataExtraData.Boolean is populated
	case shared.TicketingMetadataExtraDataTypeArrayOfTicketingMetadata5:
		// ticketingMetadataExtraData.ArrayOfTicketingMetadata5 is populated
}
```
