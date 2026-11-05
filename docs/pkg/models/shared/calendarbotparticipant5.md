# CalendarBotParticipant5


## Supported Types

### CalendarBotParticipant1

```go
calendarBotParticipant5 := shared.CreateCalendarBotParticipant5CalendarBotParticipant1(shared.CalendarBotParticipant1{/* values here */})
```

### 

```go
calendarBotParticipant5 := shared.CreateCalendarBotParticipant5Str(string{/* values here */})
```

### 

```go
calendarBotParticipant5 := shared.CreateCalendarBotParticipant5Number(float64{/* values here */})
```

### 

```go
calendarBotParticipant5 := shared.CreateCalendarBotParticipant5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch calendarBotParticipant5.Type {
	case shared.CalendarBotParticipant5TypeCalendarBotParticipant1:
		// calendarBotParticipant5.CalendarBotParticipant1 is populated
	case shared.CalendarBotParticipant5TypeStr:
		// calendarBotParticipant5.Str is populated
	case shared.CalendarBotParticipant5TypeNumber:
		// calendarBotParticipant5.Number is populated
	case shared.CalendarBotParticipant5TypeBoolean:
		// calendarBotParticipant5.Boolean is populated
}
```
