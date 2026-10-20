# DatastoreMetadata5


## Supported Types

### DatastoreMetadata1

```go
datastoreMetadata5 := shared.CreateDatastoreMetadata5DatastoreMetadata1(shared.DatastoreMetadata1{/* values here */})
```

### 

```go
datastoreMetadata5 := shared.CreateDatastoreMetadata5Str(string{/* values here */})
```

### 

```go
datastoreMetadata5 := shared.CreateDatastoreMetadata5Number(float64{/* values here */})
```

### 

```go
datastoreMetadata5 := shared.CreateDatastoreMetadata5Boolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch datastoreMetadata5.Type {
	case shared.DatastoreMetadata5TypeDatastoreMetadata1:
		// datastoreMetadata5.DatastoreMetadata1 is populated
	case shared.DatastoreMetadata5TypeStr:
		// datastoreMetadata5.Str is populated
	case shared.DatastoreMetadata5TypeNumber:
		// datastoreMetadata5.Number is populated
	case shared.DatastoreMetadata5TypeBoolean:
		// datastoreMetadata5.Boolean is populated
}
```
