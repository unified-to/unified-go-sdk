# Value

The field's value, typed by `format`: DATE as ISO-8601 (`YYYY-MM-DD`, or `YYYY-MM-DDTHH:MM:SSZ` for date-times); NUMBER, PRICE and CURRENCY as a number; BOOLEAN and YES_NO as true/false; MULTIPLE_SELECT as an array of option values; otherwise a string.


## Supported Types

### 

```go
value := shared.CreateValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
value := shared.CreateValueStr(string{/* values here */})
```

### 

```go
value := shared.CreateValueNumber(float64{/* values here */})
```

### 

```go
value := shared.CreateValueBoolean(bool{/* values here */})
```

### 

```go
value := shared.CreateValueArrayOfAccountingMetadata5([]shared.AccountingMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch value.Type {
	case shared.ValueTypeMapOfAny:
		// value.MapOfAny is populated
	case shared.ValueTypeStr:
		// value.Str is populated
	case shared.ValueTypeNumber:
		// value.Number is populated
	case shared.ValueTypeBoolean:
		// value.Boolean is populated
	case shared.ValueTypeArrayOfAccountingMetadata5:
		// value.ArrayOfAccountingMetadata5 is populated
}
```
