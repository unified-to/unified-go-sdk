# ExtraData


## Supported Types

### 

```go
extraData := shared.CreateExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
extraData := shared.CreateExtraDataStr(string{/* values here */})
```

### 

```go
extraData := shared.CreateExtraDataNumber(float64{/* values here */})
```

### 

```go
extraData := shared.CreateExtraDataBoolean(bool{/* values here */})
```

### 

```go
extraData := shared.CreateExtraDataArrayOfAtsMetadata5([]shared.AtsMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch extraData.Type {
	case shared.ExtraDataTypeMapOfAny:
		// extraData.MapOfAny is populated
	case shared.ExtraDataTypeStr:
		// extraData.Str is populated
	case shared.ExtraDataTypeNumber:
		// extraData.Number is populated
	case shared.ExtraDataTypeBoolean:
		// extraData.Boolean is populated
	case shared.ExtraDataTypeArrayOfAtsMetadata5:
		// extraData.ArrayOfAtsMetadata5 is populated
}
```
