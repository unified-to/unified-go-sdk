# TicketingMetadataValue


## Supported Types

### 

```go
ticketingMetadataValue := shared.CreateTicketingMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
ticketingMetadataValue := shared.CreateTicketingMetadataValueStr(string{/* values here */})
```

### 

```go
ticketingMetadataValue := shared.CreateTicketingMetadataValueNumber(float64{/* values here */})
```

### 

```go
ticketingMetadataValue := shared.CreateTicketingMetadataValueBoolean(bool{/* values here */})
```

### 

```go
ticketingMetadataValue := shared.CreateTicketingMetadataValueArrayOfTicketingMetadataSchemas5([]shared.TicketingMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch ticketingMetadataValue.Type {
	case shared.TicketingMetadataValueTypeMapOfAny:
		// ticketingMetadataValue.MapOfAny is populated
	case shared.TicketingMetadataValueTypeStr:
		// ticketingMetadataValue.Str is populated
	case shared.TicketingMetadataValueTypeNumber:
		// ticketingMetadataValue.Number is populated
	case shared.TicketingMetadataValueTypeBoolean:
		// ticketingMetadataValue.Boolean is populated
	case shared.TicketingMetadataValueTypeArrayOfTicketingMetadataSchemas5:
		// ticketingMetadataValue.ArrayOfTicketingMetadataSchemas5 is populated
}
```
