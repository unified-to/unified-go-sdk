# TicketingMetadataValue

The field's value, typed by `format`: DATE as ISO-8601 (`YYYY-MM-DD`, or `YYYY-MM-DDTHH:MM:SSZ` for date-times); NUMBER, PRICE and CURRENCY as a number; BOOLEAN and YES_NO as true/false; MULTIPLE_SELECT as an array of option values; otherwise a string.


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
