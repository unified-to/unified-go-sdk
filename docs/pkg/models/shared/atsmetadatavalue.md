# AtsMetadataValue

The field's value, typed by `format`: DATE as ISO-8601 (`YYYY-MM-DD`, or `YYYY-MM-DDTHH:MM:SSZ` for date-times); NUMBER, PRICE and CURRENCY as a number; BOOLEAN and YES_NO as true/false; MULTIPLE_SELECT as an array of option values; otherwise a string.


## Supported Types

### 

```go
atsMetadataValue := shared.CreateAtsMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
atsMetadataValue := shared.CreateAtsMetadataValueStr(string{/* values here */})
```

### 

```go
atsMetadataValue := shared.CreateAtsMetadataValueNumber(float64{/* values here */})
```

### 

```go
atsMetadataValue := shared.CreateAtsMetadataValueBoolean(bool{/* values here */})
```

### 

```go
atsMetadataValue := shared.CreateAtsMetadataValueArrayOfAtsMetadataSchemas5([]shared.AtsMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch atsMetadataValue.Type {
	case shared.AtsMetadataValueTypeMapOfAny:
		// atsMetadataValue.MapOfAny is populated
	case shared.AtsMetadataValueTypeStr:
		// atsMetadataValue.Str is populated
	case shared.AtsMetadataValueTypeNumber:
		// atsMetadataValue.Number is populated
	case shared.AtsMetadataValueTypeBoolean:
		// atsMetadataValue.Boolean is populated
	case shared.AtsMetadataValueTypeArrayOfAtsMetadataSchemas5:
		// atsMetadataValue.ArrayOfAtsMetadataSchemas5 is populated
}
```
