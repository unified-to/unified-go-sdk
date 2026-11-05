# PlatformMetadata


## Supported Types

### 

```go
platformMetadata := shared.CreatePlatformMetadataMapOfAny(map[string]any{/* values here */})
```

### 

```go
platformMetadata := shared.CreatePlatformMetadataStr(string{/* values here */})
```

### 

```go
platformMetadata := shared.CreatePlatformMetadataNumber(float64{/* values here */})
```

### 

```go
platformMetadata := shared.CreatePlatformMetadataBoolean(bool{/* values here */})
```

### 

```go
platformMetadata := shared.CreatePlatformMetadataArrayOfCalendarBotParticipant5([]shared.CalendarBotParticipant5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch platformMetadata.Type {
	case shared.PlatformMetadataTypeMapOfAny:
		// platformMetadata.MapOfAny is populated
	case shared.PlatformMetadataTypeStr:
		// platformMetadata.Str is populated
	case shared.PlatformMetadataTypeNumber:
		// platformMetadata.Number is populated
	case shared.PlatformMetadataTypeBoolean:
		// platformMetadata.Boolean is populated
	case shared.PlatformMetadataTypeArrayOfCalendarBotParticipant5:
		// platformMetadata.ArrayOfCalendarBotParticipant5 is populated
}
```
