# CalendarMetadata5


## Supported Types

### CalendarMetadata1

```go
calendarMetadata5 := shared.CreateCalendarMetadata5CalendarMetadata1(shared.CalendarMetadata1{/* values here */})
```

### 

```go
calendarMetadata5 := shared.CreateCalendarMetadata5Str(string{/* values here */})
```

### 

```go
calendarMetadata5 := shared.CreateCalendarMetadata5Number(float64{/* values here */})
```

### 

```go
calendarMetadata5 := shared.CreateCalendarMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch calendarMetadata5.Type {
	case shared.CalendarMetadata5TypeCalendarMetadata1:
		// calendarMetadata5.CalendarMetadata1 is populated
	case shared.CalendarMetadata5TypeStr:
		// calendarMetadata5.Str is populated
	case shared.CalendarMetadata5TypeNumber:
		// calendarMetadata5.Number is populated
	case shared.CalendarMetadata5TypeBoolean:
		// calendarMetadata5.Boolean is populated
}
```
