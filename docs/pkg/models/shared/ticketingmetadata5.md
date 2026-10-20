# TicketingMetadata5


## Supported Types

### TicketingMetadata1

```go
ticketingMetadata5 := shared.CreateTicketingMetadata5TicketingMetadata1(shared.TicketingMetadata1{/* values here */})
```

### 

```go
ticketingMetadata5 := shared.CreateTicketingMetadata5Str(string{/* values here */})
```

### 

```go
ticketingMetadata5 := shared.CreateTicketingMetadata5Number(float64{/* values here */})
```

### 

```go
ticketingMetadata5 := shared.CreateTicketingMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch ticketingMetadata5.Type {
	case shared.TicketingMetadata5TypeTicketingMetadata1:
		// ticketingMetadata5.TicketingMetadata1 is populated
	case shared.TicketingMetadata5TypeStr:
		// ticketingMetadata5.Str is populated
	case shared.TicketingMetadata5TypeNumber:
		// ticketingMetadata5.Number is populated
	case shared.TicketingMetadata5TypeBoolean:
		// ticketingMetadata5.Boolean is populated
}
```
