# TicketingMetadataSchemas5


## Supported Types

### TicketingMetadataSchemas1

```go
ticketingMetadataSchemas5 := shared.CreateTicketingMetadataSchemas5TicketingMetadataSchemas1(shared.TicketingMetadataSchemas1{/* values here */})
```

### 

```go
ticketingMetadataSchemas5 := shared.CreateTicketingMetadataSchemas5Str(string{/* values here */})
```

### 

```go
ticketingMetadataSchemas5 := shared.CreateTicketingMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
ticketingMetadataSchemas5 := shared.CreateTicketingMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch ticketingMetadataSchemas5.Type {
	case shared.TicketingMetadataSchemas5TypeTicketingMetadataSchemas1:
		// ticketingMetadataSchemas5.TicketingMetadataSchemas1 is populated
	case shared.TicketingMetadataSchemas5TypeStr:
		// ticketingMetadataSchemas5.Str is populated
	case shared.TicketingMetadataSchemas5TypeNumber:
		// ticketingMetadataSchemas5.Number is populated
	case shared.TicketingMetadataSchemas5TypeBoolean:
		// ticketingMetadataSchemas5.Boolean is populated
}
```
