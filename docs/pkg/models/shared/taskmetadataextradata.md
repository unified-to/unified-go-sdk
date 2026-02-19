# TaskMetadataExtraData


## Supported Types

### 

```go
taskMetadataExtraData := shared.CreateTaskMetadataExtraDataMapOfAny(map[string]any{/* values here */})
```

### 

```go
taskMetadataExtraData := shared.CreateTaskMetadataExtraDataStr(string{/* values here */})
```

### 

```go
taskMetadataExtraData := shared.CreateTaskMetadataExtraDataNumber(float64{/* values here */})
```

### 

```go
taskMetadataExtraData := shared.CreateTaskMetadataExtraDataBoolean(bool{/* values here */})
```

### 

```go
taskMetadataExtraData := shared.CreateTaskMetadataExtraDataArrayOfTaskMetadata5([]shared.TaskMetadata5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch taskMetadataExtraData.Type {
	case shared.TaskMetadataExtraDataTypeMapOfAny:
		// taskMetadataExtraData.MapOfAny is populated
	case shared.TaskMetadataExtraDataTypeStr:
		// taskMetadataExtraData.Str is populated
	case shared.TaskMetadataExtraDataTypeNumber:
		// taskMetadataExtraData.Number is populated
	case shared.TaskMetadataExtraDataTypeBoolean:
		// taskMetadataExtraData.Boolean is populated
	case shared.TaskMetadataExtraDataTypeArrayOfTaskMetadata5:
		// taskMetadataExtraData.ArrayOfTaskMetadata5 is populated
}
```
