# CalendarMetadataValue


## Supported Types

### 

```go
calendarMetadataValue := shared.CreateCalendarMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
calendarMetadataValue := shared.CreateCalendarMetadataValueStr(string{/* values here */})
```

### 

```go
calendarMetadataValue := shared.CreateCalendarMetadataValueNumber(float64{/* values here */})
```

### 

```go
calendarMetadataValue := shared.CreateCalendarMetadataValueBoolean(bool{/* values here */})
```

### 

```go
calendarMetadataValue := shared.CreateCalendarMetadataValueArrayOfCalendarMetadataSchemas5([]shared.CalendarMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch calendarMetadataValue.Type {
	case shared.CalendarMetadataValueTypeMapOfAny:
		// calendarMetadataValue.MapOfAny is populated
	case shared.CalendarMetadataValueTypeStr:
		// calendarMetadataValue.Str is populated
	case shared.CalendarMetadataValueTypeNumber:
		// calendarMetadataValue.Number is populated
	case shared.CalendarMetadataValueTypeBoolean:
		// calendarMetadataValue.Boolean is populated
	case shared.CalendarMetadataValueTypeArrayOfCalendarMetadataSchemas5:
		// calendarMetadataValue.ArrayOfCalendarMetadataSchemas5 is populated
}
```
