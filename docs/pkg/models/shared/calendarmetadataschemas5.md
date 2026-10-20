# CalendarMetadataSchemas5


## Supported Types

### CalendarMetadataSchemas1

```go
calendarMetadataSchemas5 := shared.CreateCalendarMetadataSchemas5CalendarMetadataSchemas1(shared.CalendarMetadataSchemas1{/* values here */})
```

### 

```go
calendarMetadataSchemas5 := shared.CreateCalendarMetadataSchemas5Str(string{/* values here */})
```

### 

```go
calendarMetadataSchemas5 := shared.CreateCalendarMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
calendarMetadataSchemas5 := shared.CreateCalendarMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch calendarMetadataSchemas5.Type {
	case shared.CalendarMetadataSchemas5TypeCalendarMetadataSchemas1:
		// calendarMetadataSchemas5.CalendarMetadataSchemas1 is populated
	case shared.CalendarMetadataSchemas5TypeStr:
		// calendarMetadataSchemas5.Str is populated
	case shared.CalendarMetadataSchemas5TypeNumber:
		// calendarMetadataSchemas5.Number is populated
	case shared.CalendarMetadataSchemas5TypeBoolean:
		// calendarMetadataSchemas5.Boolean is populated
}
```
