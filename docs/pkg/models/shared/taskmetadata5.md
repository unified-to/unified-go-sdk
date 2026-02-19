# TaskMetadata5


## Supported Types

### TaskMetadata1

```go
taskMetadata5 := shared.CreateTaskMetadata5TaskMetadata1(shared.TaskMetadata1{/* values here */})
```

### 

```go
taskMetadata5 := shared.CreateTaskMetadata5Str(string{/* values here */})
```

### 

```go
taskMetadata5 := shared.CreateTaskMetadata5Number(float64{/* values here */})
```

### 

```go
taskMetadata5 := shared.CreateTaskMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch taskMetadata5.Type {
	case shared.TaskMetadata5TypeTaskMetadata1:
		// taskMetadata5.TaskMetadata1 is populated
	case shared.TaskMetadata5TypeStr:
		// taskMetadata5.Str is populated
	case shared.TaskMetadata5TypeNumber:
		// taskMetadata5.Number is populated
	case shared.TaskMetadata5TypeBoolean:
		// taskMetadata5.Boolean is populated
}
```
