# QmsMetadata5


## Supported Types

### QmsMetadata1

```go
qmsMetadata5 := shared.CreateQmsMetadata5QmsMetadata1(shared.QmsMetadata1{/* values here */})
```

### 

```go
qmsMetadata5 := shared.CreateQmsMetadata5Str(string{/* values here */})
```

### 

```go
qmsMetadata5 := shared.CreateQmsMetadata5Number(float64{/* values here */})
```

### 

```go
qmsMetadata5 := shared.CreateQmsMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch qmsMetadata5.Type {
	case shared.QmsMetadata5TypeQmsMetadata1:
		// qmsMetadata5.QmsMetadata1 is populated
	case shared.QmsMetadata5TypeStr:
		// qmsMetadata5.Str is populated
	case shared.QmsMetadata5TypeNumber:
		// qmsMetadata5.Number is populated
	case shared.QmsMetadata5TypeBoolean:
		// qmsMetadata5.Boolean is populated
}
```
