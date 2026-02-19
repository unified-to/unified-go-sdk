# TaskMetadataSchemas5


## Supported Types

### TaskMetadataSchemas1

```go
taskMetadataSchemas5 := shared.CreateTaskMetadataSchemas5TaskMetadataSchemas1(shared.TaskMetadataSchemas1{/* values here */})
```

### 

```go
taskMetadataSchemas5 := shared.CreateTaskMetadataSchemas5Str(string{/* values here */})
```

### 

```go
taskMetadataSchemas5 := shared.CreateTaskMetadataSchemas5Number(float64{/* values here */})
```

### 

```go
taskMetadataSchemas5 := shared.CreateTaskMetadataSchemas5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch taskMetadataSchemas5.Type {
	case shared.TaskMetadataSchemas5TypeTaskMetadataSchemas1:
		// taskMetadataSchemas5.TaskMetadataSchemas1 is populated
	case shared.TaskMetadataSchemas5TypeStr:
		// taskMetadataSchemas5.Str is populated
	case shared.TaskMetadataSchemas5TypeNumber:
		// taskMetadataSchemas5.Number is populated
	case shared.TaskMetadataSchemas5TypeBoolean:
		// taskMetadataSchemas5.Boolean is populated
}
```
