# CalendarMetadataExtraData


## Supported Types

### 

```go
calendarMetadataExtraData := shared.CreateCalendarMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
calendarMetadataExtraData := shared.CreateCalendarMetadataExtraDataStr(string{/* values here */})
```

### 

```go
calendarMetadataExtraData := shared.CreateCalendarMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
calendarMetadataExtraData := shared.CreateCalendarMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
calendarMetadataExtraData := shared.CreateCalendarMetadataExtraDataArrayOfCalendarMetadata5([]shared.CalendarMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch calendarMetadataExtraData.Type {
	case shared.CalendarMetadataExtraDataTypeMapOfAny:
		// calendarMetadataExtraData.MapOfAny is populated
	case shared.CalendarMetadataExtraDataTypeStr:
		// calendarMetadataExtraData.Str is populated
	case shared.CalendarMetadataExtraDataTypeNumber:
		// calendarMetadataExtraData.Number is populated
	case shared.CalendarMetadataExtraDataTypeBoolean:
		// calendarMetadataExtraData.Boolean is populated
	case shared.CalendarMetadataExtraDataTypeArrayOfCalendarMetadata5:
		// calendarMetadataExtraData.ArrayOfCalendarMetadata5 is populated
}
```
