# TaskMetadataValue


## Supported Types

### 

```go
taskMetadataValue := shared.CreateTaskMetadataValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
taskMetadataValue := shared.CreateTaskMetadataValueStr(string{/* values here */})
```

### 

```go
taskMetadataValue := shared.CreateTaskMetadataValueNumber(float64{/* values here */})
```

### 

```go
taskMetadataValue := shared.CreateTaskMetadataValueBoolean(bool{/* values here */})
```

### 

```go
taskMetadataValue := shared.CreateTaskMetadataValueArrayOfTaskMetadataSchemas5([]shared.TaskMetadataSchemas5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch taskMetadataValue.Type {
	case shared.TaskMetadataValueTypeMapOfAny:
		// taskMetadataValue.MapOfAny is populated
	case shared.TaskMetadataValueTypeStr:
		// taskMetadataValue.Str is populated
	case shared.TaskMetadataValueTypeNumber:
		// taskMetadataValue.Number is populated
	case shared.TaskMetadataValueTypeBoolean:
		// taskMetadataValue.Boolean is populated
	case shared.TaskMetadataValueTypeArrayOfTaskMetadataSchemas5:
		// taskMetadataValue.ArrayOfTaskMetadataSchemas5 is populated
}
```
